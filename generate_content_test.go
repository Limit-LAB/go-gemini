package gemini_test

import (
	"encoding/json"
	"fmt"
	"github.com/Limit-LAB/go-gemini"
	"github.com/Limit-LAB/go-gemini/models"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestGenerateContent(t *testing.T) {
	godotenv.Load()
	key := os.Getenv("GEMINI")
	cli := gemini.NewClient(key)
	rst, err := cli.GenerateContent(models.GeminiPro,
		models.NewGenerateContentRequest(nil).
			AppendContent(
				models.RoleUser,
				models.NewParts(nil).
					AppendPart(models.NewTextPart("Hi! Use 3 words to describe yourself.")),
			),
	)

	if err != nil {
		panic(err)
	}
	jsonPrint(rst)
	fmt.Println(rst.Candidates[0].Content.Parts[0].GetText())
}

func jsonPrint(v any) {
	bs, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(bs))
}
