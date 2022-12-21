package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey         string                `json:"connection_key"`
	Result                bool                  `json:"result"`
	RedisKey              string                `json:"redis_key"`
	Filepath              string                `json:"filepath"`
	APIStatusCode         int                   `json:"api_status_code"`
	RuntimeSessionID      string                `json:"runtime_session_id"`
	BusinessPartner       *int                  `json:"business_partner"`
	ServiceLabel          string                `json:"service_label"`
	APIType               string                `json:"api_type"`
	OrdersInputParameters OrdersInputParameters `json:"OrdersInputParameters"`
	Header                Header                `json:"Orders"`
	APISchema             string                `json:"api_schema"`
	Accepter              []string              `json:"accepter"`
	Deleted               bool                  `json:"deleted"`
}

type OrdersInputParameters struct {
	ReferenceDocument     *int `json:"ReferenceDocument"`
	ReferenceDocumentItem *int `json:"ReferenceDocumentItem"`
}

type Header struct {
	OrderID                         *int            `json:"OrderID"`
	OrderDate                       *string         `json:"OrderDate"`
	OrderType                       string          `json:"OrderType"`
	Buyer                           *int            `json:"Buyer"`
	Seller                          *int            `json:"Seller"`
	CreationDate                    *string         `json:"CreationDate"`
	LastChangeDate                  *string         `json:"LastChangeDate"`
	ContractType                    string          `json:"ContractType"`
	ValidityStartDate               *string         `json:"ValidityStartDate"`
	ValidityEndDate                 *string         `json:"ValidityEndDate"`
	InvoicePeriodStartDate          *string         `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate            *string         `json:"InvoicePeriodEndDate"`
	TotalNetAmount                  *float32        `json:"TotalNetAmount"`
	TotalTaxAmount                  *float32        `json:"TotalTaxAmount"`
	TotalGrossAmount                *float32        `json:"TotalGrossAmount"`
	HeaderDeliveryStatus            string          `json:"HeaderDeliveryStatus"`
	HeaderBlockStatus               *bool           `json:"HeaderBlockStatus"`
	HeaderBillingStatus             string          `json:"HeaderBillingStatus"`
	HeaderDocReferenceStatus        string          `json:"HeaderDocReferenceStatus"`
	TransactionCurrency             string          `json:"TransactionCurrency"`
	PricingDate                     *string         `json:"PricingDate"`
	PriceDetnExchangeRate           *float32        `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate           *string         `json:"RequestedDeliveryDate"`
	HeaderCompleteDeliveryIsDefined *bool           `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderBillingBlockStatus        *bool           `json:"HeaderBillingBlockStatus"`
	HeaderDeliveryBlockStatus       *bool           `json:"HeaderDeliveryBlockStatus"`
	Incoterms                       string          `json:"Incoterms"`
	BillFromParty                   *int            `json:"BillFromParty"`
	BillToParty                     *int            `json:"BillToParty"`
	BillFromCountry                 string          `json:"BillFromCountry"`
	BillToCountry                   string          `json:"BillToCountry"`
	Payer                           *int            `json:"Payer"`
	Payee                           *int            `json:"Payee"`
	PaymentTerms                    string          `json:"PaymentTerms"`
	PaymentMethod                   string          `json:"PaymentMethod"`
	ReferenceDocument               *int            `json:"ReferenceDocument"`
	ReferenceDocumentItem           *int            `json:"ReferenceDocumentItem"`
	BPAccountAssignmentGroup        string          `json:"BPAccountAssignmentGroup"`
	AccountingExchangeRate          *float32        `json:"AccountingExchangeRate"`
	InvoiceDocumentDate             *string         `json:"InvoiceDocumentDate"`
	IsExportImportDelivery          *bool           `json:"IsExportImportDelivery"`
	HeaderText                      string          `json:"HeaderText"`
	HeaderPartner                   []HeaderPartner `json:"HeaderPartner"`
	Address                         []Address       `json:"Address"`
	HeaderPDF                       []HeaderPDF     `json:"HeaderPDF"`
	Item                            []Item          `json:"Item"`
}

type HeaderUpdates struct {
	TotalNetAmount                  *float32 `json:"TotalNetAmount"`
	TotalTaxAmount                  *float32 `json:"TotalTaxAmount"`
	TotalGrossAmount                *float32 `json:"TotalGrossAmount"`
	HeaderBlockStatus               *bool    `json:"HeaderBlockStatus"`
	TransactionCurrency             string   `json:"TransactionCurrency"`
	PricingDate                     *string  `json:"PricingDate"`
	PriceDetnExchangeRate           *string  `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate           *string  `json:"RequestedDeliveryDate"`
	HeaderBillingBlockStatus        *bool    `json:"HeaderBillingBlockStatus"`
	HeaderDeliveryBlockStatus       *bool    `json:"HeaderDeliveryBlockStatus"`
	HeaderCompleteDeliveryIsDefined *bool    `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderBillingBlockReason        *bool    `json:"HeaderBillingBlockReason"`
	Incoterms                       string   `json:"Incoterms"`
	BillFromParty                   *int     `json:"BillFromParty"`
	BillToParty                     *int     `json:"BillToParty"`
	BillFromCountry                 string   `json:"BillFromCountry"`
	BillToCountry                   string   `json:"BillToCountry"`
	Payer                           *int     `json:"Payer"`
	Payee                           *int     `json:"Payee"`
	PaymentTerms                    string   `json:"PaymentTerms"`
	PaymentMethod                   string   `json:"PaymentMethod"`
	InvoiceDocumentDate             *string  `json:"InvoiceDocumentDate"`
	HeaderText                      string   `json:"HeaderText"`
}

type HeaderPartner struct {
	PartnerFunction         string                 `json:"PartnerFunction"`
	BusinessPartner         *int                   `json:"BusinessPartner"`
	BusinessPartnerFullName string                 `json:"BusinessPartnerFullName"`
	BusinessPartnerName     string                 `json:"BusinessPartnerName"`
	Organization            string                 `json:"Organization"`
	Country                 string                 `json:"Country"`
	Language                string                 `json:"Language"`
	Currency                string                 `json:"Currency"`
	ExternalDocumentID      string                 `json:"ExternalDocumentID"`
	AddressID               *int                   `json:"AddressID"`
	HeaderPartnerContact    []HeaderPartnerContact `json:"HeaderPartnerContact"`
	HeaderPartnerPlant      []HeaderPartnerPlant   `json:"HeaderPartnerPlant"`
}

type HeaderPartnerUpdates struct {
	BusinessPartnerFullName string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     string `json:"BusinessPartnerName"`
	Organization            string `json:"Organization"`
	Country                 string `json:"Country"`
	Language                string `json:"Language"`
	Currency                string `json:"Currency"`
	ExternalDocumentID      string `json:"ExternalDocumentID"`
	AddressID               *int   `json:"AddressID"`
}

type HeaderPartnerPlant struct {
	Plant string `json:"Plant"`
}

type HeaderPartnerPlantUpdates struct {
	Plant string `json:"Plant"`
}

type HeaderPartnerContact struct {
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

type HeaderPartnerContactUpdates struct {
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

type Address struct {
	AddressID   *int   `json:"AddressID"`
	PostalCode  string `json:"PostalCode"`
	LocalRegion string `json:"LocalRegion"`
	Country     string `json:"Country"`
	District    string `json:"District"`
	StreetName  string `json:"StreetName"`
	CityName    string `json:"CityName"`
	Building    string `json:"Building"`
	Floor       *int   `json:"Floor"`
	Room        *int   `json:"Room"`
}

type AddressUpdates struct {
	PostalCode  string `json:"PostalCode"`
	LocalRegion string `json:"LocalRegion"`
	Country     string `json:"Country"`
	District    string `json:"District"`
	StreetName  string `json:"StreetName"`
	CityName    string `json:"CityName"`
	Building    string `json:"Building"`
	Floor       *int   `json:"Floor"`
	Room        *int   `json:"Room"`
}

type HeaderPDF struct {
	DocType                  string `json:"DocType"`
	DocVersionID             *int   `json:"DocVersionID"`
	DocID                    string `json:"DocID"`
	DocIssuerBusinessPartner *int   `json:"DocIssuerBusinessPartner"`
	FileName                 string `json:"FileName"`
}

type Item struct {
	OrderItem                                     *int                 `json:"OrderItem"`
	OrderItemCategory                             string               `json:"OrderItemCategory"`
	OrderItemText                                 string               `json:"OrderItemText"`
	OrderItemTextByBuyer                          string               `json:"OrderItemTextByBuyer"`
	OrderItemTextBySeller                         string               `json:"OrderItemTextBySeller"`
	Product                                       string               `json:"Product"`
	ProductStandardID                             string               `json:"ProductStandardID"`
	ProductGroup                                  string               `json:"ProductGroup"`
	BaseUnit                                      string               `json:"BaseUnit"`
	PricingDate                                   *string              `json:"PricingDate"`
	PriceDetnExchangeRate                         *float32             `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate                         *string              `json:"RequestedDeliveryDate"`
	DeliverFromParty                              *int                 `json:"DeliverFromParty"`
	DeliverToParty                                *int                 `json:"DeliverToParty"`
	StockConfirmationPartnerFunction              string               `json:"StockConfirmationPartnerFunction"`
	StockConfirmationBusinessPartner              *int                 `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                        string               `json:"StockConfirmationPlant"`
	StockConfirmationPlantBatch                   string               `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate  *string              `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityEndDate    *string              `json:"StockConfirmationPlantBatchValidityEndDate"`
	ProductIsBatchManagedInStockConfirmationPlant *bool                `json:"ProductIsBatchManagedInStockConfirmationPlant"`
	OrderQuantityInBaseUnit                       *float32             `json:"OrderQuantityInBaseUnit"`
	OrderQuantityInIssuingUnit                    *float32             `json:"OrderQuantityInIssuingUnit"`
	OrderQuantityInReceivingUnit                  *float32             `json:"OrderQuantityInReceivingUnit"`
	OrderIssuingUnit                              string               `json:"OrderIssuingUnit"`
	OrderReceivingUnit                            string               `json:"OrderReceivingUnit"`
	StockConfirmationPolicy                       string               `json:"StockConfirmationPolicy"`
	StockConfirmationStatus                       string               `json:"StockConfirmationStatus"`
	ConfdDelivQtyInOrderQtyUnit                   *float32             `json:"ConfdDelivQtyInOrderQtyUnit"`
	ItemWeightUnit                                string               `json:"ItemWeightUnit"`
	ProductGrossWeight                            *float32             `json:"ProductGrossWeight"`
	ItemGrossWeight                               *float32             `json:"ItemGrossWeight"`
	ProductNetWeight                              *float32             `json:"ProductNetWeight"`
	ItemNetWeight                                 *float32             `json:"ItemNetWeight"`
	NetAmount                                     *float32             `json:"NetAmount"`
	TaxAmount                                     *float32             `json:"TaxAmount"`
	GrossAmount                                   *float32             `json:"GrossAmount"`
	InvoiceDocumentDate                           *string              `json:"InvoiceDocumentDate"`
	ProductionPlantPartnerFunction                string               `json:"ProductionPlantPartnerFunction"`
	ProductionPlantBusinessPartner                *int                 `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                               string               `json:"ProductionPlant"`
	ProductionPlantTimeZone                       string               `json:"ProductionPlantTimeZone"`
	ProductionPlantStorageLocation                string               `json:"ProductionPlantStorageLocation"`
	IssuingPlantPartnerFunction                   string               `json:"IssuingPlantPartnerFunction"`
	IssuingPlantBusinessPartner                   *int                 `json:"IssuingPlantBusinessPartner"`
	IssuingPlant                                  string               `json:"IssuingPlant"`
	IssuingPlantTimeZone                          string               `json:"IssuingPlantTimeZone"`
	IssuingPlantStorageLocation                   string               `json:"IssuingPlantStorageLocation"`
	ReceivingPlantPartnerFunction                 string               `json:"ReceivingPlantPartnerFunction"`
	ReceivingPlantBusinessPartner                 *int                 `json:"ReceivingPlantBusinessPartner"`
	ReceivingPlant                                string               `json:"ReceivingPlant"`
	ReceivingPlantTimeZone                        string               `json:"ReceivingPlantTimeZone"`
	ReceivingPlantStorageLocation                 string               `json:"ReceivingPlantStorageLocation"`
	ProductIsBatchManagedInProductionPlant        *bool                `json:"ProductIsBatchManagedInProductionPlant"`
	ProductIsBatchManagedInIssuingPlant           *bool                `json:"ProductIsBatchManagedInIssuingPlant"`
	ProductIsBatchManagedInReceivingPlant         *bool                `json:"ProductIsBatchManagedInReceivingPlant"`
	BatchMgmtPolicyInProductionPlant              string               `json:"BatchMgmtPolicyInProductionPlant"`
	BatchMgmtPolicyInIssuingPlant                 string               `json:"BatchMgmtPolicyInIssuingPlant"`
	BatchMgmtPolicyInReceivingPlant               string               `json:"BatchMgmtPolicyInReceivingPlant"`
	ProductionPlantBatch                          string               `json:"ProductionPlantBatch"`
	IssuingPlantBatch                             string               `json:"IssuingPlantBatch"`
	ReceivingPlantBatch                           string               `json:"ReceivingPlantBatch"`
	ProductionPlantBatchValidityStartDate         *string              `json:"ProductionPlantBatchValidityStartDate"`
	ProductionPlantBatchValidityEndDate           *string              `json:"ProductionPlantBatchValidityEndDate"`
	IssuingPlantBatchValidityStartDate            *string              `json:"IssuingPlantBatchValidityStartDate"`
	IssuingPlantBatchValidityEndDate              *string              `json:"IssuingPlantBatchValidityEndDate"`
	ReceivingPlantBatchValidityStartDate          *string              `json:"ReceivingPlantBatchValidityStartDate"`
	ReceivingPlantBatchValidityEndDate            *string              `json:"ReceivingPlantBatchValidityEndDate"`
	Incoterms                                     string               `json:"Incoterms"`
	BPTaxClassification                           string               `json:"BPTaxClassification"`
	ProductTaxClassification                      string               `json:"ProductTaxClassification"`
	BPAccountAssignmentGroup                      string               `json:"BPAccountAssignmentGroup"`
	ProductAccountAssignmentGroup                 string               `json:"ProductAccountAssignmentGroup"`
	PaymentTerms                                  string               `json:"PaymentTerms"`
	PaymentMethod                                 string               `json:"PaymentMethod"`
	ItemBillingBlockReason                        *bool                `json:"ItemBillingBlockReason"`
	ItemDeliveryBlockStatus                       *bool                `json:"ItemDeliveryBlockStatus"`
	Project                                       string               `json:"Project"`
	AccountingExchangeRate                        *float32             `json:"AccountingExchangeRate"`
	ReferenceDocument                             *int                 `json:"ReferenceDocument"`
	ReferenceDocumentItem                         *int                 `json:"ReferenceDocumentItem"`
	ItemCompleteDeliveryIsDefined                 *bool                `json:"ItemCompleteDeliveryIsDefined"`
	ItemDeliveryStatus                            string               `json:"ItemDeliveryStatus"`
	IssuingStatus                                 string               `json:"IssuingStatus"`
	ReceivingStatus                               string               `json:"ReceivingStatus"`
	BillingStatus                                 string               `json:"BillingStatus"`
	TaxCode                                       string               `json:"TaxCode"`
	TaxRate                                       *float32             `json:"TaxRate"`
	CountryOfOrigin                               string               `json:"CountryOfOrigin"`
	ItemPartner                                   []ItemPartner        `json:"ItemPartner"`
	ItemPricingElement                            []ItemPricingElement `json:"ItemPricingElement"`
	ItemSchedulingLine                            []ItemSchedulingLine `json:"ItemSchedulingLine"`
}

type ItemUpdates struct {
	OrderItemText                    string  `json:"OrderItemText"`
	OrderItemTextByBuyer             string  `json:"OrderItemTextByBuyer"`
	OrderItemTextBySeller            string  `json:"OrderItemTextBySeller"`
	Product                          string  `json:"Product"`
	ProductStandardID                string  `json:"ProductStandardID"`
	ProductGroup                     string  `json:"ProductGroup"`
	RequestedDeliveryDate            *string `json:"RequestedDeliveryDate"`
	DeliverFromParty                 *int    `json:"DeliverFromParty"`
	DeliverToParty                   *int    `json:"DeliverToParty"`
	StockConfirmationPartnerFunction string  `json:"StockConfirmationPartnerFunction"`
	StockConfirmationBusinessPartner *int    `json:"StockConfirmationBusinessPartner"`
	ItemBillingBlockReason           *bool   `json:"ItemBillingBlockReason"`
	ItemDeliveryBlockStatus          *bool   `json:"ItemDeliveryBlockStatus"`
}

type ItemPartner struct {
	PartnerFunction  string           `json:"PartnerFunction"`
	BusinessPartner  *int             `json:"BusinessPartner"`
	ItemPartnerPlant ItemPartnerPlant `json:"ItemPartnerPlant"`
}

type ItemPartnerUpdates struct {
	PartnerFunction string `json:"PartnerFunction"`
}

type ItemPartnerPlant struct {
	Plant string `json:"Plant"`
}

type ItemPartnerPlantUpdates struct {
	Plant string `json:"Plant"`
}

type ItemPricingElement struct {
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

type ItemPricingElementUpdates struct {
	ConditionRateValue         *float32 `json:"ConditionRateValue"`
	ConditionAmount            *float32 `json:"ConditionAmount"`
	ConditionIsManuallyChanged *bool    `json:"ConditionIsManuallyChanged"`
}

type ItemSchedulingLine struct {
	ScheduleLine                                 *int     `json:"ScheduleLine"`
	Product                                      string   `json:"Product"`
	StockConfirmationPartnerFunction             string   `json:"StockConfirmationPartnerFunction"`
	StockConfirmationBusinessPartner             *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                       string   `json:"StockConfirmationPlant"`
	StockConfirmationPlantBatch                  string   `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate *string  `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityEndDate   *string  `json:"StockConfirmationPlantBatchValidityEndDate"`
	ConfirmedDeliveryDate                        *string  `json:"ConfirmedDeliveryDate"`
	RequestedDeliveryDate                        *string  `json:"RequestedDeliveryDate"`
	OrderQuantityInBaseUnit                      *float32 `json:"OrderQuantityInBaseUnit"`
	ConfdOrderQtyByPDTAvailCheck                 *float32 `json:"ConfdOrderQtyByPDTAvailCheck"`
	DeliveredQtyInOrderQtyUnit                   *float32 `json:"DeliveredQtyInOrderQtyUnit"`
	OpenConfdDelivQtyInOrdQtyUnit                *float32 `json:"OpenConfdDelivQtyInOrdQtyUnit"`
	ItemScheduleLineDeliveryBlockStatus          *bool    `json:"DelivBlockReasonForSchedLine"`
	PlusMinusFlag                                string   `json:"PlusMinusFlag"`
}

type ItemSchedulingLineUpdates struct {
	RequestedDeliveryDate               *string  `json:"RequestedDeliveryDate"`
	OpenConfdDelivQtyInOrdQtyUnit       *float32 `json:"OpenConfdDelivQtyInOrdQtyUnit"`
	ItemScheduleLineDeliveryBlockStatus *bool    `json:"ItemScheduleLineDeliveryBlockStatus"`
}
