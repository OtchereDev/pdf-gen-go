package generator

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aymerick/raymond"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

type ChromeDp struct{}

func (c *ChromeDp) CompileTemplate(name string, data map[string]interface{}) (string, error) {

	RegisterHelpers()

	templatePath := filepath.Join("..", "templates", name+".hbs")

	tmplContent, err := os.ReadFile(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to read template file: %w", err)
	}

	// Compile the template with Raymond
	result, err := raymond.Render(string(tmplContent), data)
	if err != nil {
		return "", fmt.Errorf("failed to render template: %w", err)
	}

	return result, nil
}

func (c *ChromeDp) GeneratePDF(p GenerationParam) (string, error) {
	// Compile the template with the given data
	htmlContent, err := c.CompileTemplate(p.TemplateName, p.Data)
	if err != nil {
		return "", err
	}

	// Create a Chrome browser context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Configure the PDF options
	var pdfBuffer []byte
	headerTemplate := ""
	footerTemplate := ""

	if p.WithHeader {
		headerTemplate = `
		<div style="display: flex; justify-content: flex-end; padding-left: 40px; padding-right: 40px;" class="flex justify-end px-10">
			<img src="LOGO_URL" style="width: 30%" alt="logo" />
		</div>`
	}

	footerTemplate = `
		<div style="font-size: 12px; margin-left: auto; margin-right: auto; font-family: Verdana, sans-serif">
			<p style="color: #e76f0f; font-weight: 700; text-align: center">
				No. 12 Opoku Adjei Avenue Patasi, off the Friends Garden Junction, Kumasi
			</p>
			<p style="text-align: center">
				<strong>(T)</strong> +233(0) 322 299 310, +233(0)507 677669.
				<strong>(E)</strong> info@spectrahealthgh.com
				<strong>www</strong>.spectrahealth.com
			</p>
		</div>`

	err = chromedp.Run(ctx,
		chromedp.Navigate("data:text/html,"+htmlContent), // Set the HTML content
		// chromedp.Emulate(device.),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var marginTop, marginBottom float64
			if p.RemoveMargins {
				marginTop = 0
				marginBottom = 0
			} else {
				marginTop = 130
				marginBottom = 100
			}

			pdfBuffer, _, err = page.PrintToPDF().WithPrintBackground(true).
				WithMarginTop(marginTop).WithMarginBottom(marginBottom).
				WithDisplayHeaderFooter(p.WithHeader).WithHeaderTemplate(headerTemplate).
				WithFooterTemplate(footerTemplate).Do(ctx)

			os.WriteFile("sample.pdf", pdfBuffer, 0644)
			return err
		}),
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate PDF: %w", err)
	}

	return base64.StdEncoding.EncodeToString(pdfBuffer), nil
}
