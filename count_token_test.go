package gemini_test

import (
	"github.com/Limit-LAB/go-gemini"
	"github.com/Limit-LAB/go-gemini/models"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestCountToken(t *testing.T) {
	godotenv.Load()
	key := os.Getenv("GEMINI")
	cli := gemini.NewClient(key)
	rst, err := cli.CountToken(models.GeminiPro, models.CountTokenRequest{
		Contents: []models.Content{
			{
				Parts: models.NewParts().AppendPart(models.NewTextPart("Hello, world!")),
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	jsonPrint(rst)
}
