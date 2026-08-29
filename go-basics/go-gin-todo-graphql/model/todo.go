package models

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var Todos = []Todo{
	{ID: 1, Title: "Learn Go", Done: false},
	{ID: 2, Title: "Build GraphQL API", Done: true},
}

func GetAllTodos() []Todo {
	return Todos
}

func GetTodoByID(id int) *Todo {
	for _, t := range Todos {
		if t.ID == id {
			return &t
		}
	}
	return nil
}

func CreateTodo(title string) Todo {
	newTodo := Todo{ID: len(Todos) + 1, Title: title, Done: false}
	Todos = append(Todos, newTodo)
	return newTodo
}

func ToggleTodoDone(id int) *Todo {
	for i, t := range Todos {
		if t.ID == id {
			Todos[i].Done = !t.Done
			return &Todos[i]
		}
	}
	return nil
}
