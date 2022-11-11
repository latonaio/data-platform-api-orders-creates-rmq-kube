package requests

type HeaderPartnerContact struct {
	OrderID           *int   `json:"OrderID"`
	PartnerFunction   string `json:"PartnerFunction"`
	BusinessPartner   *int   `json:"BusinessPartner"`
	ContactID         *int   `json:"ContactID"`
	ContactPersonName string `json:"ContactPersonName"`
	EmailAddress      string `json:"EmailAddress"`
	PhoneNumber       string `json:"PhoneNumber"`
	MobilePhoneNumber string `json:"MobilePhoneNumber"`
	FaxNumber         string `json:"FaxNumber"`
	ContactTag1       string `json:"ContactTag1"`
	ContactTag2       string `json:"ContactTag2"`
	ContactTag3       string `json:"ContactTag3"`
	ContactTag4       string `json:"ContactTag4"`
}
