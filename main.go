package StudentManagementWebServer

import (
	"github.com/labstack/echo/v4"
)

type Student struct {
	ID    int    `json: "id"`
	Name  string `json: "name"`
	Age   int    `json: "age"`
	Grade int    `json: "grade"`
}

var students []Student

func main() {
	e := echo.New()

}
