package http

type ReceiptData struct {
	fileName      string
	name          string
	patientId     string
	paymentDate   string
	receiptNumber string
	paymentMethod string
	discount      string
	subTotal      string
	total         string
	paymentItems  []struct {
		code        string
		description string
		quantity    float64
		unitPrice   string
		amount      string
	}
}

type RequestFormData struct {
	patientName        string
	sex                string
	date               string
	age                string
	phoneNumber        string
	address            string
	requestingDoctor   string
	requestingFacility string
	examination        string
	query              string
	fileName           string
}

type ReportData struct {
	name              string
	dob               string
	age               string
	patientId         string
	referringDoctor   string
	referringFacility string
	detail            string
	procedure         string
	examinationDate   string
	radiologist       string
	verifiedAt        string
	fileName          string
	titleOnReport     string
	impressions       string
	findings          string
	approvedBy        string
	approverTitle     string
	withHeader        bool
}
