package generator_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	. "github.com/OtchereDev/pdf-gen-go/internal/generator"
	"github.com/matryer/is"
)

func Test_Can_Open_Template(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		var (
			is = is.NewRelaxed(t)

			name = "invoice"
		)

		templatePath := filepath.Join("..", "templates", name+".hbs")
		_, err := os.ReadFile(templatePath)

		is.NoErr(err)

	})
}

func Test_CompileTemplate(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		var (
			is   = is.NewRelaxed(t)
			name = "invoice"

			g = newGenerator()

			data = map[string]interface{}{
				"name":            "John Doe",
				"date":            time.Now(),
				"tailwindcss":     "http://localhost:3000/css/styles.css",
				"logo":            "http://localhost:3000/images/logo.png",
				"examinationDate": "2025-01-20",
			}
		)

		htmlContent, err := g.CompileTemplate(name, data)

		is.NoErr(err)
		t.Log(htmlContent)
		fmt.Println("content: ", htmlContent)
		is.True(len(htmlContent) > 1)

	})
}

func Test_GeneratePDF(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		var (
			is           = is.NewRelaxed(t)
			g            = newGenerator()
			templateName = "invoice"
			data         = map[string]interface{}{
				"name":        "John Doe",
				"date":        time.Now(),
				"tailwindcss": "http://localhost:3000/css/styles.css",
				"logo":        "http://localhost:3000/images/logo.png",
				"items": []map[string]interface{}{
					{"name": "Item 1", "quantity": 2, "price": 10.00, "total": 20.00},
					{"name": "Item 2", "quantity": 1, "price": 15.00, "total": 15.00},
				},
				"total": 35.00,
			}
		)

		pdfUrl, err := g.GeneratePDF(GenerationParam{
			TemplateName: templateName,
			Data:         data,
		})

		is.NoErr(err)
		t.Log(pdfUrl)
		fmt.Println("content: ", pdfUrl)
		is.True(len(pdfUrl) > 1)
	})

}

func newGenerator() *ChromeDp {
	return &ChromeDp{}
}
