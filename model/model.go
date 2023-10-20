package model

type Todo struct {
	Id    uint64 `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func CreateTodo(title string) error {
	statement := "insert into todos (title) values ($1)"

	_, err := db.Query(statement, title)

	return err
}

func GetTodos() ([]Todo, error) {
	todos := []Todo{}

	statement := "select id, title, done from todos"

	rows, err := db.Query(statement)

	if err != nil {
		return todos, err
	}

	defer db.Close()

	for rows.Next() {
		todo := Todo{}
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Done)

		if err != nil {
			return todos, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}