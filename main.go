package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Student struct {
	ID    int    `json: "id"`
	Name  string `json: "name"`
	Age   int    `json: "age"`
	Grade int    `json: "grade"`
}

var students []Student

func getStudents(c echo.Context) error {
	return c.JSON(http.StatusOK, students)
}

func getStudentById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, student := range students {
		if student.ID == id {
			return c.JSON(http.StatusOK, student)
		}
	}

	return c.JSON(http.StatusNotFound, "Student not found!")
}

func createStudent(c echo.Context) error {
	student := new(Student)

	if err := c.Bind(student); err != nil {
		return err
	}

	student.ID = len(students) + 1
	students = append(students, *student)

	return c.JSON(http.StatusCreated, student)
}

func updateStudent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	for i := range students {
		if students[i].ID == id {
			updateStudent := new(Student)

			if err := c.Bind(updateStudent); err != nil {
				return err
			}

			students[i].Name = updateStudent.Name
			students[i].Age = updateStudent.Age
			students[i].Grade = updateStudent.Grade

			return c.JSON(http.StatusOK, students[i])
		}
	}

	return c.JSON(http.StatusNotFound, "Student not found!")
}

func deleteStudent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	for i := range students {
		if students[i].ID == id {
			students = append(students[:i], students[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}

	return c.JSON(http.StatusNotFound, "Student not found!")
}

func main() {
	e := echo.New()

	e.GET("/students", getStudents)
	e.POST("/students", createStudent)
	e.GET("/students/:id", getStudentById)
	e.PUT("/students/:id", updateStudent)
	e.DELETE("/students/:id", deleteStudent)

	e.Logger.Fatal(e.Start(":1323"))
}
