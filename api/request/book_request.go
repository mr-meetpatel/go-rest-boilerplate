package request

type BookRequest struct {
	Title  string `validate:"required min=1,max=100" json:"title"`
	Author string `validate:"required min=1,max=100" json:"author"`
}
