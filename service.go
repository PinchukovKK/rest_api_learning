package taskService

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task Task) (Task, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetAllTask() ([]Task, error) {
	return s.repo.GetAllTask()
}

func (s *TaskService) UpdateTask(task Task, id uint) (Task, error) {
	return s.repo.UpdateTask(id, task)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.DeleteTask(id)
}
