package handlers

import (
	"reflect"

	"github.com/gofiber/fiber/v2"
)

func ListAPIEndpoints(c *fiber.Ctx) error {

	type Result struct {
		Note []string
		User []string
	}

	// Create an instance of the Result struct
	result := Result{
		Note: listMethods(&NoteHandler{}),
		User: listMethods(&UserHandler{}),
	}
	return c.JSON(result)
}

func listMethods(obj interface{}) []string {
	t := reflect.TypeOf(obj)
	m := make([]string, 0, t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		m = append(m, method.Name)
	}
	return m
}
