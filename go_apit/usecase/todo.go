package usecase

import (
	"todo/domain/model"
	"todo/domain/repository"
)

type TodoUsecase interface {
	Search(string) (todo []*model.Todo, err error)
	View() (todo []*model.Todo, err error)
	Add(*model.Todo) (err error)
	Edit(*model.Todo) (err error)
}

type todoUsecase struct {
	todoRepo repository.TodoRepository
}

func NewTodoUsecase(todoRepo repository.TodoRepository) TodoUsecase {
	todoUsecase := todoUsecase{todoRepo: todoRepo}
	return &todoUsecase
}

// Search 入力された内容でTodoを検索する
func (usecase *todoUsecase) Search(word string) (todo []*model.Todo, err error) {
	todo, err = usecase.todoRepo.Find(word)
	return
}

// View はTodoの一覧を表示する
func (usecase *todoUsecase) View() (todo []*model.Todo, err error) {
	todo, err = usecase.todoRepo.FindAll()
	return
}

// Add はTodoを新規追加する
func (usecase *todoUsecase) Add(todo *model.Todo) (err error) {
	_, err = usecase.todoRepo.Create(todo)
	return
}

// Edit はTodoを編集する
func (usecase *todoUsecase) Edit(todo *model.Todo) (err error) {
	_, err = usecase.todoRepo.Update(todo)
	return
}
