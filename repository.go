package taskService

import "gorm.io/gorm"

type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	GetAllTask() ([]Task, error)
	UpdateTask(id uint, task Task) (Task, error)
	DeleteTask(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTask() ([]Task, error) {
	var task []Task
	err := r.db.Find(&task).Error
	return task, err
}

func (r *taskRepository) UpdateTask(id uint, task Task) (Task, error) {
	result := r.db.First(&task, id)
	if result.Error != nil {
		return Task{}, result.Error
	}

	var updateTask Task
	task.Task = updateTask.Task
	task.IsDone = updateTask.IsDone

	err := r.db.Save(&task).Error
	return task, err
}

func (r *taskRepository) DeleteTask(id uint) error {
	var task Task
	result := r.db.First(&task, id)
	if result.Error != nil {
		return result.Error
	}

	err := r.db.Delete(&task).Error
	if err != nil {
		return err
	}
	return nil
}
