package handlers

import "github.com/gin-gonic/gin"

// ListCategories returns available categories (system and user-defined).
func ListCategories(c *gin.Context) {}

// CreateCategory creates a new category for the authenticated user.
func CreateCategory(c *gin.Context) {}

// UpdateCategory updates a user's category.
func UpdateCategory(c *gin.Context) {}

// DeleteCategory deletes a user's category.
func DeleteCategory(c *gin.Context) {}
