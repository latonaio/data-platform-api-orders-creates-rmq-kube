package existence_conf

type BusinessPartnerReq struct {
	ConnectionKey     string          `json:"connection_key"`
	Result            bool            `json:"result"`
	RedisKey          string          `json:"redis_key"`
	RuntimeSessionID  string          `json:"runtime_session_id"`
	BusinessPartnerID *int            `json:"business_partner"`
	Filepath          string          `json:"filepath"`
	ServiceLabel      string          `json:"service_label"`
	BusinessPartner   BusinessPartner `json:"BusinessPartner"`
	APISchema         string          `json:"api_schema"`
	Accepter          []string        `json:"accepter"`
	OrderID           *int            `json:"order_id"`
	Deleted           bool            `json:"deleted"`
}
type BusinessPartner struct {
	BusinessPartner int `json:"BusinessPartner"`
}
