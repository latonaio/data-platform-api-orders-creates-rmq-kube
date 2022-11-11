package requests

type ItemPartnerPlant struct {
	OrderID         *int   `json:"OrderID"`
	OrderItem       *int   `json:"OrderItem"`
	PartnerFunction string `json:"PartnerFunction"`
	BusinessPartner *int   `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
}
