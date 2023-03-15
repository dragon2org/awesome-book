package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"

	"github.com/gin-gonic/gin"
)

// https://github.com/go-playground/validator/blob/master/_examples/simple/main.go
type CreateTitle struct {
	Title   string `validate:"required"`
	Content string `validate:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/api/v1/book", func(ctx *gin.Context) {
		var req CreateTitle
		if err := ctx.BindJSON(&req); err != nil {
			panic(err)
		}
		validate := validator.New()
		err := validate.Struct(req)
		if err != nil {
			// this check is only needed when your code could produce
			// an invalid value for validation such as interface with nil
			// value most including myself do not usually have code like this.
			if _, ok := err.(*validator.InvalidValidationError); ok {
				fmt.Println(err)
				return
			}
			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println(err.Namespace())
				fmt.Println(err.Field())
				fmt.Println(err.StructNamespace())
				fmt.Println(err.StructField())
				fmt.Println(err.Tag())
				fmt.Println(err.ActualTag())
				fmt.Println(err.Kind())
				fmt.Println(err.Type())
				fmt.Println(err.Value())
				fmt.Println(err.Param())
				fmt.Println()
			}
			// from here you can create your own error messages in whatever language you wish
			return
		}
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
