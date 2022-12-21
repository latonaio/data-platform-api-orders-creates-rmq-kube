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
	InvoicePeriodStartDate          *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate            *string  `json:"InvoicePeriodEndDate"`
	TotalNetAmount                  *float32 `json:"TotalNetAmount"`
	TotalTaxAmount                  *float32 `json:"TotalTaxAmount"`
	TotalGrossAmount                *float32 `json:"TotalGrossAmount"`
	HeaderDeliveryStatus            string   `json:"HeaderDeliveryStatus"`
	HeaderBlockStatus               *bool    `json:"HeaderBlockStatus"`
	HeaderBillingStatus             string   `json:"HeaderBillingStatus"`
	HeaderDocReferenceStatus        string   `json:"HeaderDocReferenceStatus"`
	TransactionCurrency             string   `json:"TransactionCurrency"`
	PricingDate                     *string  `json:"PricingDate"`
	PriceDetnExchangeRate           *float32 `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate           *string  `json:"RequestedDeliveryDate"`
	HeaderCompleteDeliveryIsDefined *bool    `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderBillingBlockStatus        *bool    `json:"HeaderBillingBlockStatus"`
	HeaderDeliveryBlockStatus       *bool    `json:"HeaderDeliveryBlockStatus"`
	Incoterms                       string   `json:"Incoterms"`
	BillFromParty                   *int     `json:"BillFromParty"`
	BillToParty                     *int     `json:"BillToParty"`
	BillFromCountry                 string   `json:"BillFromCountry"`
	BillToCountry                   string   `json:"BillToCountry"`
	Payer                           *int     `json:"Payer"`
	Payee                           *int     `json:"Payee"`
	PaymentTerms                    string   `json:"PaymentTerms"`
	PaymentMethod                   string   `json:"PaymentMethod"`
	ReferenceDocument               *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem           *int     `json:"ReferenceDocumentItem"`
	BPAccountAssignmentGroup        string   `json:"BPAccountAssignmentGroup"`
	AccountingExchangeRate          *float32 `json:"AccountingExchangeRate"`
	InvoiceDocumentDate             *string  `json:"InvoiceDocumentDate"`
	IsExportImportDelivery          *bool    `json:"IsExportImportDelivery"`
	HeaderText                      string   `json:"HeaderText"`
}
