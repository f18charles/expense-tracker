# internal

Contains application business logic and private packages not intended for external consumption.

Guidelines

- Keep APIs small and focused.
- No public export from `internal` for external modules to use.
- Prefer small, testable packages with clear responsibilities.
- Avoid framework lock-in in core logic; adaptors/handlers should live elsewhere.

Testing

- Add unit tests for behavior; integration tests may live outside `internal`.

This file intentionally remains brief and leaves implementation details to the code.