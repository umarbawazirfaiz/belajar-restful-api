package helper

import "belajar-restful-api/exception"

func PanicIfError(err error) {
	if err != nil {
		panic(exception.NewInternalServerError(err))
	}
}
