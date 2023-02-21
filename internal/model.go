package internal

type ToDo struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Status string `json:"status"`
}

type TodoList struct {
	List []ToDo `json:"list"`
}
