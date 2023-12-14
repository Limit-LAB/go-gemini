package gemini_test

import (
	"github.com/Limit-LAB/go-gemini"
	"github.com/Limit-LAB/go-gemini/models"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestGetModelInfo(t *testing.T) {
	godotenv.Load()
	key := os.Getenv("GEMINI")
	cli := gemini.NewClient(key)
	rst, err := cli.GetModelInfo(models.GeminiPro)
	if err != nil {
		t.Fatal(err)
	}
	jsonPrint(rst)
}

func TestGetModelList(t *testing.T) {
	godotenv.Load()
	key := os.Getenv("GEMINI")
	cli := gemini.NewClient(key)
	rst, err := cli.GetModelList()
	if err != nil {
		t.Fatal(err)
	}
	jsonPrint(rst)
}
