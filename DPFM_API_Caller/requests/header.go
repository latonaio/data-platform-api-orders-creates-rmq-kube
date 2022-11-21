package requests

type Header struct {
	OrderID                         *int     `json:"OrderID"`
	OrderDate                       *string  `json:"OrderDate"`
	OrderType                       string   `json:"OrderType"`
	Buyer                           *int     `json:"Buyer"`
	Seller                          *int     `json:"Seller"`
	CreationDate                    *string  `json:"CreationDate"`
	LastChangeDate                  *string  `json:"LastChangeDate"`
	ContractType                    string   `json:"ContractType"`
	ValidityStartDate               *string  `json:"ValidityStartDate"`
	ValidityEndDate                 *string  `json:"ValidityEndDate"`
	InvoiceScheduleStartDate        *string  `json:"InvoiceScheduleStartDate"`
	InvoiceScheduleEndDate          *string  `json:"InvoiceScheduleEndDate"`
	TotalNetAmount                  *float32 `json:"TotalNetAmount"`
	TotalTaxAmount                  *float32 `json:"TotalTaxAmount"`
	TotalGrossAmount                *float32 `json:"TotalGrossAmount"`
	OverallDeliveryStatus           string   `json:"OverallDeliveryStatus"`
	TotalBlockStatus                *bool    `json:"TotalBlockStatus"`
	OverallOrdReltdBillgStatus      string   `json:"OverallOrdReltdBillgStatus"`
	OverallDocReferenceStatus       string   `json:"OverallDocReferenceStatus"`
	TransactionCurrency             string   `json:"TransactionCurrency"`
	PricingDate                     *string  `json:"PricingDate"`
	PriceDetnExchangeRate           *float32 `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate           *string  `json:"RequestedDeliveryDate"`
	HeaderCompleteDeliveryIsDefined *bool    `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderBillingBlockReason        *bool    `json:"HeaderBillingBlockReason"`
	DeliveryBlockReason             *bool    `json:"DeliveryBlockReason"`
	Incoterms                       string   `json:"Incoterms"`
	PaymentTerms                    string   `json:"PaymentTerms"`
	PaymentMethod                   string   `json:"PaymentMethod"`
	ReferenceDocument               *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem           *int     `json:"ReferenceDocumentItem"`
	BPAccountAssignmentGroup        string   `json:"BPAccountAssignmentGroup"`
	AccountingExchangeRate          *float32 `json:"AccountingExchangeRate"`
	BillingDocumentDate             *string  `json:"BillingDocumentDate"`
	IsExportImportDelivery          *bool    `json:"IsExportImportDelivery"`
	HeaderText                      string   `json:"HeaderText"`
}
