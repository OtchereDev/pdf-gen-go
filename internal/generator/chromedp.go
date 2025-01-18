package generator

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/aymerick/raymond"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

type ChromeDp struct{}

func (c *ChromeDp) RegisterAsset() error {
	RegisterHelpers()
	err := RegisterParials()

	return err
}

func (c *ChromeDp) CompileTemplate(name string, data map[string]interface{}) (string, error) {

	templatePath := filepath.Join("..", "templates", name+".hbs")

	tmplContent, err := os.ReadFile(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to read template file: %w", err)
	}

	data["logo"] = LOGO

	// Compile the template with Raymond
	result, err := raymond.Render(string(tmplContent), data)
	if err != nil {
		return "", fmt.Errorf("failed to render template: %w", err)
	}

	return result, nil
}

func (c *ChromeDp) GeneratePDF(p GenerationParam) (string, error) {

	html, err := c.CompileTemplate(p.TemplateName, p.Data)
	if err != nil {
		return "", err
	}

	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf))
	defer cancel()

	headerTemplate := ""
	footerTemplate := ""

	if p.WithHeader {
		headerTemplate = fmt.Sprintf("<div style='display: flex; justify-content: flex-end; padding-left: 40px; padding-right: 40px;' class='flex justify-end px-10'><img src='%s' style='width: 30%%' alt='logo' /></div>", LOGO)
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

	var buf []byte

	if err := chromedp.Run(ctx,
		// the navigation will trigger the "page.EventLoadEventFired" event too,
		// so we should add the listener after the navigation.
		chromedp.Navigate("about:blank"),
		// set the page content and wait until the page is loaded (including its resources).
		chromedp.ActionFunc(actionLoadHTMLContent(html)),
		chromedp.ActionFunc(func(ctx context.Context) error {

			q := page.PrintToPDF().WithPrintBackground(true).
				WithDisplayHeaderFooter(p.WithHeader).WithHeaderTemplate(headerTemplate).
				WithFooterTemplate(footerTemplate).WithPaperWidth(8.27)

			if !p.RemoveMargins {
				q = q.WithMarginTop(130 / 96).WithMarginBottom(100 / 96).
					WithMarginLeft(0).WithMarginRight(0)
			}

			buf, _, err = q.Do(ctx)

			if err != nil {
				return err
			}
			return os.WriteFile("sample.pdf", buf, 0644)

		}),
	); err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(buf), nil
}

func actionLoadHTMLContent(html string) chromedp.ActionFunc {
	return func(ctx context.Context) error {
		isLoaded, isSetLock := false, sync.Mutex{}
		listenerCtx, cancel := context.WithCancel(ctx)
		defer cancel()

		chromedp.ListenTarget(listenerCtx, func(ev interface{}) {
			if _, ok := ev.(*page.EventLoadEventFired); ok {
				// stop listener
				cancel()

				isSetLock.Lock()
				isLoaded = true
				isSetLock.Unlock()
			}
		})

		frameTree, err := page.GetFrameTree().Do(ctx)
		if err != nil {
			return err
		}

		if err := page.SetDocumentContent(frameTree.Frame.ID, html).Do(ctx); err != nil {
			return err
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-listenerCtx.Done():
			isSetLock.Lock()
			defer isSetLock.Unlock()
			if isLoaded {
				return nil
			}
			return listenerCtx.Err()
		}
	}
}

func Connect() (c *ChromeDp, err error) {
	a := ChromeDp{}

	err = c.RegisterAsset()

	c = &a

	return
}
