package gemini_test

import (
	"fmt"
	"github.com/Limit-LAB/go-gemini"
	"github.com/Limit-LAB/go-gemini/models"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestGeminiStreamGenerateContent(t *testing.T) {
	godotenv.Load()
	key := os.Getenv("GEMINI")
	req := models.NewGenerateContentRequest().
		AppendContent(
			models.RoleUser,
			models.NewParts().
				AppendPart(models.NewTextPart("Hi! Use 20 words to describe yourself.")),
		)
	cli := gemini.NewClient(key)
	gcs, err := cli.GenerateContentStream(models.GeminiPro, req)
	if err != nil {
		panic(err)
	}
	defer gcs.Close()
	for {
		content, err := gcs.Receive()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("==============LOOP==============")
		jsonPrint(content)
	}
}
