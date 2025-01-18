package main

import (
	"fmt"

	. "github.com/OtchereDev/pdf-gen-go/internal/generator"
	. "github.com/OtchereDev/pdf-gen-go/internal/net/http"
)

func main() {
	g, err := Connect()

	if err != nil {
		panic(err)
	}

	r, err := g.GeneratePDF(GenerationParam{
		WithHeader:   true,
		TemplateName: RequestFormTemplate,
		Data: map[string]interface{}{
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
		},
	})
	if err != nil {
		fmt.Println("Error occured", err)
		return
	}
	fmt.Println("Server up: Testing generation")
	fmt.Println(r)
}
