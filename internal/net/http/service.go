package http

import "github.com/OtchereDev/pdf-gen-go/internal/generator"

const (
	ReceiptTemplate     = "receipt"
	RequestFormTemplate = "request_form"
	ReportTemplate      = "report"
)

func GenerateReceipt(d ReceiptData, g generator.Generator) (r string, err error) {
	r, err = g.GeneratePDF(generator.GenerationParam{
		WithHeader:    false,
		RemoveMargins: true,
		TemplateName:  ReceiptTemplate,
		Data: map[string]interface{}{
			"fileName":      d.fileName,
			"name":          d.name,
			"patientId":     d.patientId,
			"paymentDate":   d.paymentDate,
			"receiptNumber": d.receiptNumber,
			"paymentMethod": d.paymentMethod,
			"discount":      d.discount,
			"subTotal":      d.subTotal,
			"total":         d.total,
			"paymentItems":  d.paymentItems,
		},
	})

	return r, err
}

func GenerateRequestForm(d RequestFormData, g generator.Generator) (r string, err error) {
	r, err = g.GeneratePDF(generator.GenerationParam{
		WithHeader:   true,
		TemplateName: RequestFormTemplate,
		Data: map[string]interface{}{
			"patientName":        d.patientName,
			"sex":                d.sex,
			"date":               d.date,
			"age":                d.age,
			"phoneNumber":        d.phoneNumber,
			"address":            d.address,
			"requestingDoctor":   d.requestingDoctor,
			"requestingFacility": d.requestingFacility,
			"examination":        d.examination,
			"query":              d.query,
			"fileName":           d.fileName,
		},
	})

	return r, err
}

func GenerateReport(d ReportData, g generator.Generator) (r string, err error) {
	r, err = g.GeneratePDF(generator.GenerationParam{
		WithHeader:   d.withHeader,
		TemplateName: ReportTemplate,
		Data: map[string]interface{}{
			"name":              d.name,
			"dob":               d.dob,
			"age":               d.age,
			"patientId":         d.patientId,
			"referringDoctor":   d.referringDoctor,
			"referringFacility": d.referringFacility,
			"detail":            d.detail,
			"procedure":         d.procedure,
			"examinationDate":   d.examinationDate,
			"radiologist":       d.radiologist,
			"verifiedAt":        d.verifiedAt,
			"fileName":          d.fileName,
			"titleOnReport":     d.titleOnReport,
			"impressions":       d.impressions,
			"findings":          d.findings,
			"approvedBy":        d.approvedBy,
			"approverTitle":     d.approverTitle,
		},
	})

	return r, err
}
