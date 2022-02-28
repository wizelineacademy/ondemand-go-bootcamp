package controller

type Context interface {
	JSON(code int, v interface{}) error
	Bind(i interface{}) error
}
