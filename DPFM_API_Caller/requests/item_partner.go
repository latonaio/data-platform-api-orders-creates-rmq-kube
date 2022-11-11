package requests

type ItemPartner struct {
	OrderID         *int   `json:"OrderID"`
	OrderItem       *int   `json:"OrderItem"`
	PartnerFunction string `json:"PartnerFunction"`
	BusinessPartner *int   `json:"BusinessPartner"`
}
