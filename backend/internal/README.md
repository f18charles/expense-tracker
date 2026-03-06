# Internal Directory Architecture

This directory contains the core business logic and private packages for the expense tracker application. It implements a **layered architecture pattern** that separates concerns across different layers, making the code modular, testable, and maintainable.

---

## Table of Contents

1. [Architecture Overview](#architecture-overview)
2. [Layered Architecture Pattern](#layered-architecture-pattern)
3. [Real Example: Goals Feature](#real-example-goals-feature)
4. [Separation of Concerns](#separation-of-concerns)
5. [How to Implement This Pattern](#how-to-implement-this-pattern)
6. [Template for New Features](#template-for-new-features)
7. [Request Flow Diagram](#request-flow-diagram)
8. [Why This Structure is Good](#why-this-structure-is-good)

---

## Architecture Overview

The `internal` directory is organized into vertical slices where each feature (like Goals, Accounts, Budgets) flows through multiple horizontal layers:

```
api/handlers/          ← HTTP Layer (receives requests from clients)
    ↓
services/              ← Business Logic Layer (implements rules and workflows)
    ↓
repository/            ← Data Access Layer (talks to database)
    ↓
models/                ← Data Structures (defines the shape of data)
    ↓
database/              ← Database Configuration
```

Each layer has a **single responsibility**, and layers communicate only with the layer directly below them. This is called the **Dependency Rule**.

---

## Layered Architecture Pattern

### **The 4 Layers (Bottom to Top)**

#### 1. **Models Layer** (`models/`)
**Responsibility:** Define the shape and structure of your data.

```go
// models/goal.go
type Goal struct {
    ID            uuid.UUID  // unique identifier
    UserID        uuid.UUID  // who owns this goal
    Name          string     // goal name
    TargetAmount  float64    // how much we want to save
    CurrentAmount float64    // how much we've saved so far
    Deadline      *time.Time // when we want to reach it
    CreatedAt     time.Time  // when this goal was created
}
```

**Why:** This is the "contract" or "blueprint" for what a Goal looks like. It's purely data—no business logic here.

---

#### 2. **Repository Layer** (`repository/`)
**Responsibility:** Handle all database operations. This is the **Data Access Layer**.

```go
// repository/goals_repo.go
type GoalRepo struct {}

func (gr *GoalRepo) CreateGoal(goal *models.Goal) error {
    // Insert goal into database
    return database.DB.Create(goal).Error
}

func (gr *GoalRepo) GetGoalByID(goalID uuid.UUID) (*models.Goal, error) {
    // Fetch goal from database by ID
    var goal models.Goal
    result := database.DB.Where("id = ?", goalID).First(&goal)
    return &goal, nil
}

func (gr *GoalRepo) UpdateGoal(goal *models.Goal) error {
    // Update goal in database
    return database.DB.Save(goal).Error
}

func (gr *GoalRepo) ListGoalsByUser(user_id uuid.UUID) ([]models.Goal, error) {
    // Fetch all goals for a specific user
    goals := []models.Goal{}
    database.DB.Where("user_id = ?", user_id).Find(&goals)
    return goals, nil
}

func (gr *GoalRepo) DeleteGoal(id uuid.UUID) error {
    // Remove goal from database
    return database.DB.Delete(&models.Goal{}, "id = ?", id)
}
```

**Why:** This isolates all database queries in one place. If you ever need to switch from PostgreSQL to MongoDB, you only need to rewrite the repository. The rest of your code doesn't care.

**Key point:** Repository only knows HOW to talk to the database. It doesn't know WHY or what rules to apply.

---

#### 3. **Service Layer** (`services/`)
**Responsibility:** Implement business logic and rules. This is where decisions happen.

```go
// services/goal_service.go
type GoalService struct {
    goalRepo *repository.GoalRepo // Has access to database operations
}

func (gs *GoalService) GoalCreate(user_id uuid.UUID, req GoalCreateRequest) (*models.Goal, error) {
    // Create new goal object with user data
    goal := &models.Goal{
        UserID:        user_id,
        Name:          req.Name,
        TargetAmount:  req.TargetAmount,
        CurrentAmount: req.CurrentAmount, // Defaults to 0 if not provided
        Deadline:      req.Deadline,
    }
    
    // Tell repository to save it
    if err := gs.goalRepo.CreateGoal(goal); err != nil {
        return nil, err
    }
    return goal, nil
}

func (gs *GoalService) GetGoal(user_id, goal_id uuid.UUID) (*models.Goal, error) {
    // Get goal from database
    goal, err := gs.goalRepo.GetGoalByID(goal_id)
    if err != nil {
        return nil, err
    }
    
    // BUSINESS RULE: A user can only see their own goals
    if goal.UserID != user_id {
        return nil, utils.ErrForbidden
    }
    return goal, nil
}

func (gs *GoalService) GoalUpdate(user_id, goal_id uuid.UUID, req GoalUpdateRequest) (*models.Goal, error) {
    // Get the goal first
    goal, err := gs.goalRepo.GetGoalByID(goal_id)
    if err != nil {
        return nil, err
    }
    
    // BUSINESS RULE: Ownership check
    if goal.UserID != user_id {
        return nil, utils.ErrForbidden
    }
    
    // BUSINESS LOGIC: Update only provided fields
    if req.Name != "" {
        goal.Name = req.Name
    }
    if req.TargetAmount != 0 {
        goal.TargetAmount = req.TargetAmount
    }
    // ... etc
    
    // Save updated goal
    if err := gs.goalRepo.UpdateGoal(goal); err != nil {
        return nil, err
    }
    return goal, nil
}

func (gs *GoalService) GoalDelete(user_id, goal_id uuid.UUID) error {
    goal, err := gs.goalRepo.GetGoalByID(goal_id)
    if err != nil {
        return err
    }
    
    // BUSINESS RULE: Ownership check before deleting
    if goal.UserID != user_id {
        return utils.ErrForbidden
    }
    
    return gs.goalRepo.DeleteGoal(goal_id)
}
```

**Why:** Services contain the "business rules"—the decisions that make your app unique. Examples:
- "Users can only see their own goals" (a rule)
- "When updating a goal, only update fields that were provided" (a workflow)
- "Check authorization before allowing deletion" (a policy)

Services don't know about HTTP. They don't care if requests come from REST API, GraphQL, or a CLI tool. They only care about business logic.

---

#### 4. **Handler Layer** (`api/handlers/`)
**Responsibility:** Handle HTTP requests and responses. This is the **API Layer**.

```go
// api/handlers/goals.go
type GoalHandler struct {
    goalService services.GoalService // Uses the service for logic
}

func (gh *GoalHandler) ListGoals(c *gin.Context) {
    // Step 1: Extract who is making this request
    id := utils.ConfirmAuthedUser(c)
    if id == uuid.Nil {
        return // If not authenticated, the middleware already sent error response
    }
    
    // Step 2: Ask service to get goals for this user
    goals, err := gh.goalService.GoalList(id)
    if err != nil {
        utils.ErrorResponse(c, http.StatusNotFound, "no goals found")
        return
    }
    
    // Step 3: Return goals as JSON to client
    utils.SuccessResponse(c, http.StatusOK, goals)
}

func (gh *GoalHandler) CreateGoal(c *gin.Context) {
    // Step 1: Who is making this request?
    id := utils.ConfirmAuthedUser(c)
    if id == uuid.Nil {
        return
    }
    
    // Step 2: Parse the JSON request body
    var goalreq services.GoalCreateRequest
    if err := c.ShouldBindJSON(&goalreq); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }
    
    // Step 3: Ask service to create the goal
    goal, err := gh.goalService.GoalCreate(id, goalreq)
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, "failed to create goal")
        return
    }
    
    // Step 4: Return the created goal as JSON
    utils.SuccessResponse(c, http.StatusOK, goal)
}

func (gh *GoalHandler) UpdateGoal(c *gin.Context) {
    // Similar pattern: authenticate → parse request → call service → return response
}

func (gh *GoalHandler) DeleteGoal(c *gin.Context) {
    // Similar pattern
}
```

**Why:** Handlers only care about HTTP specifics:
- Parsing JSON from the request
- Checking authentication
- Returning status codes and HTTP responses

Handlers are "thin"—they mostly delegate to the service. They're the "traffic controller" for HTTP requests.

---

## Real Example: Goals Feature

Let's trace what happens when a user creates a goal via HTTP POST `/api/goals`:

```
1. HTTP Request arrives
   POST /api/goals
   Body: { "name": "Save for vacation", "target_amount": 5000 }

2. Router sends to GoalHandler.CreateGoal()

3. GoalHandler:
   - Confirms user is authenticated (gets userID)
   - Parses JSON into GoalCreateRequest struct
   - Calls goalService.GoalCreate(userID, request)

4. GoalService.GoalCreate():
   - Creates a Goal struct with the provided data
   - Calls goalRepo.CreateGoal(goal)

5. GoalRepo.CreateGoal():
   - Executes SQL: INSERT INTO goals (id, user_id, name, ...) VALUES (...)
   - Returns any database error

6. Back up the stack:
   - GoalService returns the created goal
   - GoalHandler formats it as JSON and returns HTTP 200 OK
   - Client receives: { "id": "...", "name": "Save for vacation", ... }
```

Notice how each layer does exactly one thing and passes the result up the chain.

---

## Separation of Concerns

### **Why Are They Separated?**

| Layer | Responsibility | Example Concern |
|-------|---|---|
| **Handler** | HTTP request/response | "How do I parse JSON and return HTTP status codes?" |
| **Service** | Business logic | "Is this user allowed to do this? What rules apply?" |
| **Repository** | Database queries | "How do I get data from PostgreSQL?" |
| **Models** | Data structure | "What fields does a Goal have?" |

### **The Key Insight: Each Layer Has Different Reasons to Change**

- You change **Models** when the shape of data changes (add a new field to Goal)
- You change **Repository** when you switch databases (PostgreSQL → MongoDB)
- You change **Services** when business rules change ("Users can now have 10 goals, not 5")
- You change **Handlers** when the API changes ("Return XML instead of JSON")

Each reason to change is isolated. Changing business rules doesn't affect database code. Switching databases doesn't affect business logic. This is powerful.

---

## How to Implement This Pattern

### **When Adding a New Feature (e.g., Categories)**

Follow this sequence **from bottom to top**:

1. **Start with Models** (`models/category.go`)
   - Define the data structure: fields, types, database tags

2. **Write Repository** (`repository/category_repo.go`)
   - Implement CRUD operations (Create, Read, Update, Delete)
   - Handle all database queries

3. **Write Service** (`services/category_service.go`)
   - Implement business logic
   - Use the repository to access data
   - Add security checks (authorization, validation)

4. **Write Handler** (`api/handlers/categories.go`)
   - Implement HTTP endpoints
   - Call the service for logic
   - Return HTTP responses

5. **Register in Router** (`api/router/router.go`)
   - Add routes that point to handlers

### **Why Bottom-to-Top?**

- You define data first (models)
- Then how to access it (repository)
- Then what rules apply to it (service)
- Then how to expose it (handler)

This order ensures each layer has what it needs.

---

## Template for New Features

### **File: `models/newfeature.go`**
```go
package models

type NewFeature struct {
    ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
    UserID    uuid.UUID `gorm:"type:uuid;not null"`
    Name      string    `gorm:"not null"`
    CreatedAt time.Time
}
```

### **File: `repository/newfeature_repo.go`**
```go
package repository

type NewFeatureRepo struct {}

func (r *NewFeatureRepo) Create(item *models.NewFeature) error {
    return database.DB.Create(item).Error
}

func (r *NewFeatureRepo) GetByID(id uuid.UUID) (*models.NewFeature, error) {
    var item models.NewFeature
    result := database.DB.Where("id = ?", id).First(&item)
    return &item, result.Error
}

func (r *NewFeatureRepo) Update(item *models.NewFeature) error {
    return database.DB.Save(item).Error
}

func (r *NewFeatureRepo) ListByUser(userID uuid.UUID) ([]models.NewFeature, error) {
    var items []models.NewFeature
    result := database.DB.Where("user_id = ?", userID).Find(&items)
    return items, result.Error
}

func (r *NewFeatureRepo) Delete(id uuid.UUID) error {
    return database.DB.Delete(&models.NewFeature{}, "id = ?", id).Error
}
```

### **File: `services/newfeature_service.go`**
```go
package services

type CreateRequest struct {
    Name string `json:"name" binding:"required"`
}

type UpdateRequest struct {
    Name string `json:"name"`
}

type NewFeatureService struct {
    repo *repository.NewFeatureRepo
}

func NewNewFeatureService() *NewFeatureService {
    return &NewFeatureService{
        repo: repository.NewNewFeatureRepo(),
    }
}

func (s *NewFeatureService) Create(userID uuid.UUID, req CreateRequest) (*models.NewFeature, error) {
    item := &models.NewFeature{
        UserID: userID,
        Name:   req.Name,
    }
    if err := s.repo.Create(item); err != nil {
        return nil, err
    }
    return item, nil
}

func (s *NewFeatureService) Get(userID, itemID uuid.UUID) (*models.NewFeature, error) {
    item, err := s.repo.GetByID(itemID)
    if err != nil {
        return nil, err
    }
    // BUSINESS RULE: Ownership check
    if item.UserID != userID {
        return nil, utils.ErrForbidden
    }
    return item, nil
}

func (s *NewFeatureService) Update(userID, itemID uuid.UUID, req UpdateRequest) (*models.NewFeature, error) {
    item, err := s.repo.GetByID(itemID)
    if err != nil {
        return nil, err
    }
    if item.UserID != userID {
        return nil, utils.ErrForbidden
    }
    
    if req.Name != "" {
        item.Name = req.Name
    }
    
    if err := s.repo.Update(item); err != nil {
        return nil, err
    }
    return item, nil
}

func (s *NewFeatureService) Delete(userID, itemID uuid.UUID) error {
    item, err := s.repo.GetByID(itemID)
    if err != nil {
        return err
    }
    if item.UserID != userID {
        return utils.ErrForbidden
    }
    return s.repo.Delete(itemID)
}
```

### **File: `api/handlers/newfeature.go`**
```go
package handlers

type NewFeatureHandler struct {
    service services.NewFeatureService
}

func NewNewFeatureHandler() *NewFeatureHandler {
    return &NewFeatureHandler{
        service: *services.NewNewFeatureService(),
    }
}

func (h *NewFeatureHandler) Create(c *gin.Context) {
    userID := utils.ConfirmAuthedUser(c)
    if userID == uuid.Nil {
        return
    }
    
    var req services.CreateRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }
    
    item, err := h.service.Create(userID, req)
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, "failed to create")
        return
    }
    
    utils.SuccessResponse(c, http.StatusOK, item)
}

// Implement List, Get, Update, Delete similarly...
```

---

## Request Flow Diagram

Here's how a request flows through the system:

```
┌─────────────────────────────────────────────────────────────────┐
│                     CLIENT (Browser/App)                        │
│              Sends: POST /api/goals                             │
│              Body: { "name": "Save", "target_amount": 5000 }    │
└────────────────────────────┬────────────────────────────────────┘
                             │ HTTP Request
                             ↓
┌─────────────────────────────────────────────────────────────────┐
│         HANDLER LAYER (api/handlers/goals.go)                   │
│  ✓ Parse JSON request                                           │
│  ✓ Check authentication (via middleware)                        │
│  ✓ Extract userID                                               │
│  ✗ Does NOT check business rules                               │
│  ✗ Does NOT talk to database                                    │
└────────────────────────────┬────────────────────────────────────┘
                             │ Delegate to Service
                             ↓
┌─────────────────────────────────────────────────────────────────┐
│       SERVICE LAYER (services/goal_service.go)                  │
│  ✓ Implement business logic                                     │
│  ✓ Check ownership (user can only see their goals)              │
│  ✓ Validate business rules                                      │
│  ✓ Orchestrate data operations                                  │
│  ✗ Does NOT know about HTTP                                     │
│  ✗ Does NOT execute SQL queries directly                        │
└────────────────────────────┬────────────────────────────────────┘
                             │ Use Repository
                             ↓
┌─────────────────────────────────────────────────────────────────┐
│    REPOSITORY LAYER (repository/goals_repo.go)                  │
│  ✓ Execute database queries (INSERT, SELECT, UPDATE, DELETE)    │
│  ✓ Handle database errors                                       │
│  ✗ Does NOT implement business logic                            │
│  ✗ Does NOT check authorization                                 │
└────────────────────────────┬────────────────────────────────────┘
                             │ Query Database
                             ↓
┌─────────────────────────────────────────────────────────────────┐
│              DATABASE (PostgreSQL)                              │
│  INSERT INTO goals (id, user_id, name, target_amount, ...)      │
│  VALUES ('uuid', 'user-uuid', 'Save', 5000, ...)                │
└────────────────────────────┬────────────────────────────────────┘
                             │ Return created row
                             ↓
              ... Response flows back up ...
                             ↓
┌─────────────────────────────────────────────────────────────────┐
│                     HANDLER (return response)                   │
│  Format Goal object as JSON                                     │
│  Set HTTP status 200 OK                                         │
└────────────────────────────┬────────────────────────────────────┘
                             │ HTTP Response
                             ↓
┌─────────────────────────────────────────────────────────────────┐
│                     CLIENT (Browser/App)                        │
│  Received: 200 OK                                               │
│  Body: { "id": "123e...", "name": "Save", "target_amount": ...} │
└─────────────────────────────────────────────────────────────────┘
```

---

## Function Interconnection Map

Here's how functions call each other in the Goals feature:

```
HTTP Request: POST /api/goals
        │
        ↓
GoalHandler.CreateGoal()
        │
        ├─→ utils.ConfirmAuthedUser()          [Check who is making request]
        │
        ├─→ c.ShouldBindJSON(&goalreq)         [Parse JSON from request]
        │
        ├─→ GoalService.GoalCreate()           [Delegate to business logic]
        │   │
        │   ├─→ models.Goal{}                  [Create data structure]
        │   │
        │   └─→ GoalRepo.CreateGoal()          [Save to database]
        │       │
        │       └─→ database.DB.Create()       [Execute SQL INSERT]
        │
        └─→ utils.SuccessResponse()            [Format and return JSON response]
                │
                └─→ HTTP Response (200 OK with Goal JSON)
```

Same pattern for GET, UPDATE, DELETE - just different repository methods called.

---

## Why This Structure is Good

### **1. Testability**
You can test each layer independently:
```go
// Test service WITHOUT hitting database
func TestGoalCreate(t *testing.T) {
    service := services.NewGoalService()
    goal, err := service.GoalCreate(userID, request)
    
    // Service logic works even if database is down
}

// Test handler WITHOUT running full HTTP server
func TestCreateGoalHandler(t *testing.T) {
    handler := NewGoalHandler()
    // Mock the service
    // Call handler function
    // Check HTTP response
}
```

### **2. Maintainability**
When something breaks, you know which layer to look at:
- **"Goals aren't saved"** → Check Repository (database queries)
- **"User can see other people's goals"** → Check Service (authorization logic)
- **"Weird JSON responses"** → Check Handler (response formatting)

### **3. Reusability**
The service can be used by multiple handlers:
```go
// REST API handler
func (h *GoalHandler) CreateGoal(c *gin.Context) {
    goal, _ := h.service.GoalCreate(userID, request)
}

// CLI handler
func CreateGoalCLI(args []string) {
    goal, _ := service.GoalCreate(userID, request)  // Same service!
}

// GraphQL handler
func (r *Resolver) CreateGoal(ctx context.Context, input GoalInput) {
    goal, _ := r.service.GoalCreate(userID, request)  // Same service!
}
```

The service is used everywhere because it's not tied to HTTP.

### **4. Database Independence**
Need to switch from PostgreSQL to MongoDB? Only rewrite the repository:

```go
// Old PostgreSQL version
func (gr *GoalRepo) GetGoalByID(goalID uuid.UUID) (*models.Goal, error) {
    var goal models.Goal
    database.DB.Where("id = ?", goalID).First(&goal)
    return &goal, nil
}

// New MongoDB version
func (gr *GoalRepo) GetGoalByID(goalID uuid.UUID) (*models.Goal, error) {
    var goal models.Goal
    mongoClient.Database("piggy").Collection("goals").FindOne(ctx, bson.M{"_id": goalID}).Decode(&goal)
    return &goal, nil
}
```

The service and handlers don't need to change AT ALL. They only care about the Repository interface.

### **5. Clear Code Organization**
Every file has one job:
- Need to understand what a Goal is? Look at `models/goal.go`
- Need to add a database query? Go to `repository/goals_repo.go`
- Need to add a business rule? Update `services/goal_service.go`
- Need to expose an API endpoint? Edit `api/handlers/goals.go`

No confusion about where things belong.

### **6. Scales to Complexity**
As features get more complex, each layer grows independently:

```go
// Simple service (few rules)
func (s *GoalService) GoalCreate(userID uuid.UUID, req GoalCreateRequest) (*models.Goal, error) {
    goal := &models.Goal{UserID: userID, Name: req.Name, ...}
    return goal, s.repo.CreateGoal(goal)
}

// Complex service (many rules)
func (s *GoalService) GoalCreate(userID uuid.UUID, req GoalCreateRequest) (*models.Goal, error) {
    // Rule 1: Validate user isn't deleted
    user, err := s.userService.GetUser(userID)
    if user.Status == "deleted" {
        return nil, ErrUserDeleted
    }
    
    // Rule 2: Validate user hasn't exceeded quota
    goals, err := s.repo.ListGoalsByUser(userID)
    if len(goals) >= user.GoalQuota {
        return nil, ErrQuotaExceeded
    }
    
    // Rule 3: Validate target amount is reasonable
    if req.TargetAmount < 100 {
        return nil, ErrTargetTooLow
    }
    
    // Rule 4: Create and save the goal
    goal := &models.Goal{...}
    if err := s.repo.CreateGoal(goal); err != nil {
        return nil, err
    }
    
    // Rule 5: Send notification to user
    s.notificationService.SendGoalCreatedNotification(userID, goal.ID)
    
    return goal, nil
}
```

The handler and repository don't need to know about any of these added rules.

---

## Guidelines

✅ **DO:**
- Keep each layer focused on its responsibility
- Call downward (Handler → Service → Repository → Models)
- Pass data through explicit structs (like `GoalCreateRequest`)
- Add business logic to the Service layer
- Keep database queries in the Repository layer
- Write unit tests for Services (easiest to test)

❌ **DON'T:**
- Put business logic in handlers (handlers should be thin)
- Put business logic in repository (repositories should be simple CRUD)
- Have handlers directly use repository (business logic gets skipped)
- Call upward (Repository calling Service, etc.)
- Access database directly in handlers or services (use repository)
- Keep framework logic in business logic (Service shouldn't know about HTTP/Gin)

---

## Summary

This architecture is like a well-designed restaurant:
- **Customers** (Client) place orders
- **Waiters** (Handlers) take orders and deliver food, handling customer interactions
- **Managers** (Services) make business decisions: "Do we have ingredients? Is this person a regular? Apply discount?"
- **Chefs** (Repositories) actually prepare/cook the food (database queries)
- **Recipes** (Models) define what ingredients go in each dish

Each person knows their job. If customers want a different menu (Handler change), waiters don't need to relearn how to cook. If we run out of an ingredient (business rule change), the customers don't need to know about it.