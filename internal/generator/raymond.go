package generator

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/aymerick/raymond"
)

func convertMomentToGoFormat(format string) string {
	replacements := map[string]string{
		"dddd": "Monday",  // Full weekday name
		"ddd":  "Mon",     // Short weekday name
		"DD":   "02",      // Day of the month (zero-padded)
		"D":    "2",       // Day of the month
		"MMMM": "January", // Full month name
		"MMM":  "Jan",     // Short month name
		"MM":   "01",      // Month (zero-padded)
		"M":    "1",       // Month
		"YYYY": "2006",    // 4-digit year
		"YY":   "06",      // 2-digit year
		"HH":   "15",      // Hour (24-hour clock, zero-padded)
		"H":    "3",       // Hour (24-hour clock)
		"mm":   "04",      // Minutes (zero-padded)
		"m":    "4",       // Minutes
		"ss":   "05",      // Seconds (zero-padded)
		"s":    "5",       // Seconds
	}

	for momentFormat, goFormat := range replacements {
		format = strings.ReplaceAll(format, momentFormat, goFormat)
	}
	return format
}

func RegisterHelpers() {
	raymond.RegisterHelper("dateFormat", func(date string, format string) string {
		parsedTime, err := time.Parse(time.RFC3339, date) // Assuming RFC3339 input
		if err != nil {
			return "Invalid date"
		}

		// Convert moment.js-style format to Go format
		goFormat := convertMomentToGoFormat(format)

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
