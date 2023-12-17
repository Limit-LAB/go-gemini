package gemini_test

import (
	"github.com/Limit-LAB/go-gemini"
	"github.com/Limit-LAB/go-gemini/models"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestEmbd(t *testing.T) {
	godotenv.Load()
	key := os.Getenv("GEMINI")
	cli := gemini.NewClient(key)
	emb, err := cli.EmbedContent(
		models.Embedding001,
		models.NewParts().AppendPart(models.NewTextPart("Write a story about a magic backpack.")),
	)
	if err != nil {
		t.Fatal(err)
	}
	jsonPrint(emb)
}
