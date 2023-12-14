package main

import (
	"fmt"
	gemini "github.com/Limit-Lab/go-gemini"
	"github.com/Limit-Lab/go-gemini/models"
)

func main() {
	cli := gemini.NewClient("")
	rst, err := cli.GenerateContent(models.GeminiPro,
		models.NewGenerateContentRequest(nil).
			AppendContent(
				models.RoleUser,
				models.NewParts(nil).
					AppendPart(models.NewTextPart("Hello, world!")),
			),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(rst.Candidates[0].Content.Parts[0])
}
