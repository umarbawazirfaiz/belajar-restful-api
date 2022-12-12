package domain

import "belajar-restful-api/model/response"

type Category struct {
	id   int
	name string
}

func (category *Category) ToResponseCategory() response.ResponseCategory {
	return response.ResponseCategory{
		Id:   category.id,
		Name: category.name,
	}
}

func (category *Category) GetId() *int {
	return &category.id
}

func (category *Category) GetName() *string {
	return &category.name
}

func (category *Category) SetId(id *int) {
	category.id = *id
}

func (category *Category) SetName(name *string) {
	category.name = *name
}
