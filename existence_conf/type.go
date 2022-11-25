package existence_conf

type BusinessPartnerReq struct {
	ConnectionKey     string          `json:"connection_key"`
	Result            bool            `json:"result"`
	RedisKey          string          `json:"redis_key"`
	RuntimeSessionID  string          `json:"runtime_session_id"`
	BusinessPartnerID *int            `json:"business_partner"`
	Filepath          string          `json:"filepath"`
	ServiceLabel      string          `json:"service_label"`
	BusinessPartner   BusinessPartner `json:"BusinessPartnerGeneral"`
	APISchema         string          `json:"api_schema"`
	Accepter          []string        `json:"accepter"`
	Deleted           bool            `json:"deleted"`
}
type BusinessPartner struct {
	BusinessPartner int `json:"BusinessPartner"`
}

type PlantReq struct {
	ConnectionKey     string   `json:"connection_key"`
	Result            bool     `json:"result"`
	RedisKey          string   `json:"redis_key"`
	RuntimeSessionID  string   `json:"runtime_session_id"`
	BusinessPartnerID *int     `json:"business_partner"`
	Filepath          string   `json:"filepath"`
	ServiceLabel      string   `json:"service_label"`
	Plant             Plant    `json:"PlantGeneral"`
	APISchema         string   `json:"api_schema"`
	Accepter          []string `json:"accepter"`
	Deleted           bool     `json:"deleted"`
}

type Plant struct {
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
}
