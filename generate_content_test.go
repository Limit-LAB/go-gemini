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
		models.NewGenerateContentRequest(
			models.NewContent(models.RoleUser, models.NewTextPart("你好")),
			models.NewContent(models.RoleModel, models.NewTextPart("你好！很高兴为您服务。请问您需要什么帮助？")),
			models.NewContent(models.RoleUser, models.NewTextPart("你是谁")),
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
