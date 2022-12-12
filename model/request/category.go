package request

type RequestCreateCategory struct {
	Name string `json:"nama" validate:"required,min=1,max=200"`
}

type RequestUpdateCategory struct {
	Id   int    `json:"id" validate:"required,numeric"`
	Name string `json:"name" validate:"required,min=1,max=200"`
}

type RequestDeleteCategory struct {
	Id int `json:"id" validate:"required,numeric"`
}
