package service_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/yehezkiel1086/go-rest-mockery-testify/internal/core/domain"
	"github.com/yehezkiel1086/go-rest-mockery-testify/internal/core/service"
	"github.com/yehezkiel1086/go-rest-mockery-testify/mocks"
)

type createTaskTestedInput struct {
	task *domain.Task
}

type createTaskExpectedOutput struct {
	task *domain.Task
	err error
}

func TestTaskService_CreateTask(t *testing.T) {
	ctx := context.Background()
	taskName := gofakeit.Adjective()
	taskDescription := gofakeit.Sentence(10)

	taskInput := &domain.Task{
		Name: taskName,
		Description: taskDescription,
	}

	taskOutput := &domain.Task{
		Name: taskName,
		Description: taskDescription,
		Status: domain.StatusNotCompleted,
	}

	testCases := []struct {
		desc string
		mocks func(
			taskRepo *mocks.TaskRepository,
		)
		input createTaskTestedInput
		output createTaskExpectedOutput
	}{
		{
			desc: "Success",
			mocks: func(
				taskRepo *mocks.TaskRepository,
			) {
				taskRepo.On("CreateTask", ctx, taskInput).Return(taskOutput, nil)
			},
			input: createTaskTestedInput{
				task: taskInput,
			},
			output: createTaskExpectedOutput{
				task: taskOutput,
				err: nil,
			},
		},
		{
			desc: "Fail_InternalError",
			mocks: func(
				taskRepo *mocks.TaskRepository,
			) {
				taskRepo.On("CreateTask", ctx, taskInput).Return(nil, domain.ErrInternal)
			},
			input: createTaskTestedInput{
				task: taskInput,
			},
			output: createTaskExpectedOutput{
				task: nil,
				err: domain.ErrInternal,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			taskRepo := mocks.NewTaskRepository(t)
			tc.mocks(taskRepo)

			taskService := service.NewTaskService(taskRepo)

			task, err := taskService.CreateTask(ctx, tc.input.task)
			assert.Equal(t, tc.output.err, err, "Error mismatch")
			assert.Equal(t, tc.output.task, task, "Task mismatch")
		})
	}
}

type getTaskByIDExpectedOutput struct {
	task *domain.Task
	err error
}

func TestTaskService_GetTaskByID(t *testing.T) {
	ctx := context.Background()
	taskName := gofakeit.Adjective()
	taskDescription := gofakeit.Sentence(10)
	taskStatus := domain.StatusNotCompleted

	taskOutput := &domain.Task{
		Name: taskName,
		Description: taskDescription,
		Status: taskStatus,
	}

	testCases := []struct{
		desc string
		mocks func(
			taskRepo *mocks.TaskRepository,
		)
		output getTaskByIDExpectedOutput
	}{
		{
			desc: "Success",
			mocks: func(
				taskRepo *mocks.TaskRepository,
			) {
				taskRepo.On("GetTaskByID", ctx, uint(1)).Return(taskOutput, nil)
			},
			output: getTaskByIDExpectedOutput{
				task: taskOutput,
				err: nil,
			},
		},
		{
			desc: "Fail_InternalError",
			mocks: func(
				taskRepo *mocks.TaskRepository,
			) {
				taskRepo.On("GetTaskByID", ctx, uint(1)).Return(nil, domain.ErrInternal)
			},
			output: getTaskByIDExpectedOutput{
				task: nil,
				err: domain.ErrInternal,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			taskRepo := mocks.NewTaskRepository(t)
			tc.mocks(taskRepo)

			taskService := service.NewTaskService(taskRepo)

			task, err := taskService.GetTaskByID(ctx, 1)
			assert.Equal(t, tc.output.err, err, "Error mismatch")
			assert.Equal(t, tc.output.task, task, "Task mismatch")
		})
	}
}

type getTasksExpectedOutput struct {
	tasks []domain.Task
	err error
}

func TestTaskService_GetTasks(t *testing.T) {
	ctx := context.Background()
	var tasks []domain.Task

	for i := 0; i < 10; i++ {
		task := domain.Task{
			Name: gofakeit.Adjective(),
			Description: gofakeit.Sentence(10),
		}
		tasks = append(tasks, task)
	}

	testCases := []struct{
		desc string
		mocks func(
			taskRepo *mocks.TaskRepository,
		)
		output getTasksExpectedOutput
	}{
		{
			desc: "Success",
			mocks: func(
				taskRepo *mocks.TaskRepository,
			) {
				taskRepo.On("GetTasks", ctx).Return(tasks, nil)
			},
			output: getTasksExpectedOutput{
				tasks: tasks,
				err: nil,
			},
		},
		{
			desc: "Fail_InternalError",
			mocks: func(
				taskRepo *mocks.TaskRepository,
			) {
				taskRepo.On("GetTasks", ctx).Return(nil, domain.ErrInternal)
			},
			output: getTasksExpectedOutput{
				tasks: nil,
				err: domain.ErrInternal,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			taskRepo := mocks.NewTaskRepository(t)
			tc.mocks(taskRepo)

			taskService := service.NewTaskService(taskRepo)

			tasks, err := taskService.GetTasks(ctx)
			assert.Equal(t, tc.output.err, err, "Error mismatch")
			assert.Equal(t, tc.output.tasks, tasks, "Tasks mismatch")
		})
	}
}

type updateTaskTestedInput struct {
	id   uint
	task *domain.Task
}

type updateTaskExpectedOutput struct {
	task *domain.Task
	err  error
}

func TestTaskService_UpdateTask(t *testing.T) {
	ctx := context.Background()
	id := uint(1)
	taskName := gofakeit.Adjective()
	taskDescription := gofakeit.Sentence(10)

	taskInput := &domain.Task{
		Name:        taskName,
		Description: taskDescription,
	}

	existingTask := &domain.Task{
		Name:        "Old Name",
		Description: "Old Description",
		Status:      domain.StatusNotCompleted,
	}

	taskOutput := &domain.Task{
		Name:        taskName,
		Description: taskDescription,
		Status:      domain.StatusNotCompleted,
	}

	testCases := []struct {
		desc   string
		mocks  func(taskRepo *mocks.TaskRepository)
		input  updateTaskTestedInput
		output updateTaskExpectedOutput
	}{
		{
			desc: "Success",
			mocks: func(taskRepo *mocks.TaskRepository) {
				taskRepo.On("GetTaskByID", ctx, id).Return(existingTask, nil)
				taskRepo.On("UpdateTask", ctx, id, taskInput).Return(taskOutput, nil)
			},
			input: updateTaskTestedInput{
				id:   id,
				task: taskInput,
			},
			output: updateTaskExpectedOutput{
				task: taskOutput,
				err:  nil,
			},
		},
		{
			desc: "Fail_GetTaskByID_InternalError",
			mocks: func(taskRepo *mocks.TaskRepository) {
				taskRepo.On("GetTaskByID", ctx, id).Return(nil, domain.ErrInternal)
			},
			input: updateTaskTestedInput{
				id:   id,
				task: taskInput,
			},
			output: updateTaskExpectedOutput{
				task: nil,
				err:  domain.ErrInternal,
			},
		},
		{
			desc: "Fail_UpdateTask_InternalError",
			mocks: func(taskRepo *mocks.TaskRepository) {
				taskRepo.On("GetTaskByID", ctx, id).Return(existingTask, nil)
				taskRepo.On("UpdateTask", ctx, id, taskInput).Return(nil, domain.ErrInternal)
			},
			input: updateTaskTestedInput{
				id:   id,
				task: taskInput,
			},
			output: updateTaskExpectedOutput{
				task: nil,
				err:  domain.ErrInternal,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			taskRepo := mocks.NewTaskRepository(t)
			tc.mocks(taskRepo)

			taskService := service.NewTaskService(taskRepo)

			task, err := taskService.UpdateTask(ctx, tc.input.id, tc.input.task)
			assert.Equal(t, tc.output.err, err, "Error mismatch")
			assert.Equal(t, tc.output.task, task, "Task mismatch")
		})
	}
}

func TestTaskService_DeleteTask(t *testing.T) {
	ctx := context.Background()
	id := uint(1)

	testCases := []struct {
		desc   string
		mocks  func(taskRepo *mocks.TaskRepository)
		input  uint
		output error
	}{
		{
			desc: "Success",
			mocks: func(taskRepo *mocks.TaskRepository) {
				taskRepo.On("DeleteTask", ctx, id).Return(nil)
			},
			input:  id,
			output: nil,
		},
		{
			desc: "Fail_InternalError",
			mocks: func(taskRepo *mocks.TaskRepository) {
				taskRepo.On("DeleteTask", ctx, id).Return(domain.ErrInternal)
			},
			input:  id,
			output: domain.ErrInternal,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			taskRepo := mocks.NewTaskRepository(t)
			tc.mocks(taskRepo)

			taskService := service.NewTaskService(taskRepo)

			err := taskService.DeleteTask(ctx, tc.input)
			assert.Equal(t, tc.output, err, "Error mismatch")
		})
	}
}
