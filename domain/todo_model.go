package domain

type Todo struct {
	Title string `json: "title"`
	Done  bool   `json: "done"`
}

type TodoPageList struct {
	PageTitle string `json: "page_title"`
	Todos     []Todo
}
