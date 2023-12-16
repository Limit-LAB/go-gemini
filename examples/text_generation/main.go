package main

import (
	"fmt"
	gemini "github.com/Limit-LAB/go-gemini"
	"github.com/Limit-LAB/go-gemini/models"
)

func main() {
	cli := gemini.NewClient("")
	rst, err := cli.GenerateContent(models.GeminiPro,
		models.NewGenerateContentRequest(
			models.NewContent(models.RoleUser, models.NewTextPart("How are you?")),
		),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(rst.Candidates[0].Content.Parts[0])
}
