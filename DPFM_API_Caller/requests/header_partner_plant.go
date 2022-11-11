package requests

type HeaderPartnerPlant struct {
	OrderID         *int   `json:"OrderID"`
	PartnerFunction string `json:"PartnerFunction"`
	BusinessPartner *int   `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
}
