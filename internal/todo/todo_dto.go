package todo

type AddRequestDto struct {
	Title string
	Done  bool
}

type UpdateRequestDto struct {
	ID    uint `json:"id"`
	Title string
	Done  bool
}
