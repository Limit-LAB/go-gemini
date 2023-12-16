# Go Gemini - A Go SDK for Google Gemini LLM


[![Go Reference](https://pkg.go.dev/badge/github.com/Limit-LAB/go-gemini.svg)](https://pkg.go.dev/github.com/Limit-LAB/go-gemini)

This library provides unofficial Go clients for [Gemini API](https://ai.google.dev/tutorials/rest_quickstart). We support:

* Gemini Pro
* Gemini Pro Vision

## Get Started

```go
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
```
