package requests

type ItemPricingElement struct {
	OrderID                    *int     `json:"OrderID"`
	OrderItem                  *int     `json:"OrderItem"`
	PricingProcedureStep       *int     `json:"PricingProcedureStep"`
	PricingProcedureCounter    *int     `json:"PricingProcedureCounter"`
	ConditionType              string   `json:"ConditionType"`
	PricingDate                *string  `json:"PricingDate"`
	ConditionRateValue         *float32 `json:"ConditionRateValue"`
	ConditionCurrency          string   `json:"ConditionCurrency"`
	ConditionQuantity          *float32 `json:"ConditionQuantity"`
	ConditionQuantityUnit      string   `json:"ConditionQuantityUnit"`
	ConditionRecord            *int     `json:"ConditionRecord"`
	ConditionSequentialNumber  *int     `json:"ConditionSequentialNumber"`
	TaxCode                    string   `json:"TaxCode"`
	ConditionAmount            *float32 `json:"ConditionAmount"`
	TransactionCurrency        string   `json:"TransactionCurrency"`
	ConditionIsManuallyChanged *bool    `json:"ConditionIsManuallyChanged"`
}
