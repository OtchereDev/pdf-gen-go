package generator

import (
	"os"
	"path/filepath"
	"time"

	"github.com/aymerick/raymond"
)

func ConvertMomentToGoFormat(format string) string {
	formats := map[string]string{
		"DD-MM-YYYY":         "02-01-2006",
		"DD/MM/YYYY":         "02/01/2006",
		"DD/MM/YYYY HH:mm":   "02/01/2006 15:04",
		"dddd, DD MMMM YYYY": "Monday, 02 January 2006",
	}

	return formats[format]
}

func RegisterHelpers() {
	raymond.RegisterHelper("dateFormat", func(date string, format string) string {
		parsedTime, err := time.Parse(time.RFC3339, date) // Assuming RFC3339 input
		if err != nil {
			return "Invalid date"
		}

		// Convert moment.js-style format to Go format
		goFormat := ConvertMomentToGoFormat(format)

		// Format the date
		return parsedTime.Format(goFormat)
	})
}

func RegisterParials() error {
	templatePath := filepath.Join("..", "templates", "tailwindcss.hbs")

	cssPartial, err := os.ReadFile(templatePath)
	if err != nil {
		return err
	}

	// Register the CSS partial
	raymond.RegisterPartial("styles", string(cssPartial))

	return nil
}
