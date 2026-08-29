# Go Rest Mockery Testify

This project demonstrates how to write effective unit tests in Go, specifically focusing on the service layer. It showcases the integration of several powerful libraries to create clean, maintainable, and robust tests following the **AAA (Arrange, Act, Assert)** pattern.

## Tools & Libraries

*   **Mockery:** Automatically generates mock implementations for interfaces, allowing us to isolate the unit under test (the Service) from its dependencies (the Repository).
*   **GoFakeIt:** Generates random, realistic data for test inputs. This prevents tests from relying on hardcoded values and ensures better coverage.
*   **Testify:** Provides the `assert` package for writing fluent and readable assertions.
*   **Go Standard Library:** Uses the built-in `testing` package for running tests.

## Methodology: AAA (Arrange, Act, Assert)

The tests in `internal/core/service/task_test.go` are structured using the AAA pattern:

1.  **Arrange:** Prepare the inputs, expected outputs, and configure the mock behaviors.
2.  **Act:** Execute the function being tested.
3.  **Assert:** Verify the results match expectations and ensure mocks were called correctly.

### Example from `task_test.go`

Here is a simplified view of how the `CreateTask` test is structured:

```go
func TestTaskService_CreateTask(t *testing.T) {
    // ... setup ...

    t.Run("Success", func(t *testing.T) {
        // ARRANGE
        taskRepo := mocks.NewTaskRepository(t)
        taskRepo.On("CreateTask", ctx, taskInput).Return(taskOutput, nil)
        
        taskService := service.NewTaskService(taskRepo)

        // ACT
        task, err := taskService.CreateTask(ctx, taskInput)

        // ASSERT
        assert.Nil(t, err)
        assert.Equal(t, taskOutput, task)
    })
}
```

## Running Tests

```bash
go test ./...
```
