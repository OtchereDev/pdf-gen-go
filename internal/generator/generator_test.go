package generator_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	. "github.com/OtchereDev/pdf-gen-go/internal/generator"
	"github.com/matryer/is"
)

func TestConvertMomentToGoFormat(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		var (
			is = is.NewRelaxed(t)

			format1 = "dddd, DD MMMM YYYY"
			format2 = "DD-MM-YYYY"
			format3 = "DD/MM/YYYY"
			format4 = "DD/MM/YYYY HH:mm"
			date    = "2012-04-24T18:25:43.511Z"
		)

		s1 := ConvertMomentToGoFormat(format1)
		s2 := ConvertMomentToGoFormat(format2)
		s3 := ConvertMomentToGoFormat(format3)
		s4 := ConvertMomentToGoFormat(format4)

		p, err := time.Parse(time.RFC3339, date)

		is.NoErr(err)

		r := p.Format(s1)
		is.Equal(r, "Tuesday, 24 April 2012")

		r = p.Format(s2)
		is.Equal(r, "24-04-2012")

		r = p.Format(s3)
		is.Equal(r, "24/04/2012")

		r = p.Format(s4)
		is.Equal(r, "24/04/2012 18:25")
	})
}

func TestCanOpenTemplate(t *testing.T) {
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

func TestCompileTemplate(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		var (
			is   = is.NewRelaxed(t)
			name = "invoice"

			g = newGenerator(t)

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
		is.True(len(htmlContent) > 1)

	})
}

func TestGeneratePDF(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		var (
			is           = is.NewRelaxed(t)
			g            = newGenerator(t)
			templateName = "request_form"
			data         = map[string]interface{}{
				"patientName":        "Oliver Otcher",
				"sex":                "M",
				"date":               "2012-04-23T18:25:43.511Z",
				"age":                "21",
				"phoneNumber":        "052394748393",
				"address":            "Anywhere",
				"requestingDoctor":   "Dr tesr",
				"requestingFacility": "Test fac",
				"examination":        "ECR",
				"query":              "Location",
			}
		)

		pdfUrl, err := g.GeneratePDF(GenerationParam{
			TemplateName:  templateName,
			Data:          data,
			RemoveMargins: true,
		})

		is.NoErr(err)
		is.True(len(pdfUrl) > 1)
	})

}

func newGenerator(t *testing.T) *ChromeDp {
	is := is.NewRelaxed(t)
	c, err := Connect()
	is.NoErr(err)
	return c
}

// func TestMain(m *testing.M) {
//     setup()
//     code := m.Run()
//     shutdown()
//     os.Exit(code)
// }
