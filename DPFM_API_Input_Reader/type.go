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
	ConnectionKey       string   `json:"connection_key"`
	Result              bool     `json:"result"`
	RedisKey            string   `json:"redis_key"`
	Filepath            string   `json:"filepath"`
	APIStatusCode       int      `json:"api_status_code"`
	RuntimeSessionID    string   `json:"runtime_session_id"`
	BusinessPartner     *int     `json:"business_partner"`
	ServiceLabel        string   `json:"service_label"`
	Orders              Orders   `json:"Orders"`
	APISchema           string   `json:"api_schema"`
	Accepter            []string `json:"accepter"`
	OrderID             *int     `json:"order_id"`
	Deleted             bool     `json:"deleted"`
	SQLUpdateResult     *bool    `json:"sql_update_result"`
	SQLUpdateError      string   `json:"sql_update_error"`
	SubfuncResult       *bool    `json:"subfunc_result"`
	SubfuncError        string   `json:"subfunc_error"`
	ExconfResult        *bool    `json:"exconf_result"`
	ExconfError         string   `json:"exconf_error"`
	APIProcessingResult *bool    `json:"api_processing_result"`
	APIProcessingError  string   `json:"api_processing_error"`
}

type Orders struct {
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
	InvoiceScheduleStartDate        *string         `json:"InvoiceScheduleStartDate"`
	InvoiceScheduleEndDate          *string         `json:"InvoiceScheduleEndDate"`
	TotalNetAmount                  *float32        `json:"TotalNetAmount"`
	TotalTaxAmount                  *float32        `json:"TotalTaxAmount"`
	TotalGrossAmount                *float32        `json:"TotalGrossAmount"`
	OverallDeliveryStatus           string          `json:"OverallDeliveryStatus"`
	TotalBlockStatus                *bool           `json:"TotalBlockStatus"`
	OverallOrdReltdBillgStatus      string          `json:"OverallOrdReltdBillgStatus"`
	OverallDocReferenceStatus       string          `json:"OverallDocReferenceStatus"`
	TransactionCurrency             string          `json:"TransactionCurrency"`
	PricingDate                     *string         `json:"PricingDate"`
	PriceDetnExchangeRate           *float32        `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate           *string         `json:"RequestedDeliveryDate"`
	HeaderCompleteDeliveryIsDefined *bool           `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderBillingBlockReason        *bool           `json:"HeaderBillingBlockReason"`
	DeliveryBlockReason             *bool           `json:"DeliveryBlockReason"`
	Incoterms                       string          `json:"Incoterms"`
	PaymentTerms                    string          `json:"PaymentTerms"`
	PaymentMethod                   string          `json:"PaymentMethod"`
	ReferenceDocument               *int            `json:"ReferenceDocument"`
	ReferenceDocumentItem           *int            `json:"ReferenceDocumentItem"`
	BPAccountAssignmentGroup        string          `json:"BPAccountAssignmentGroup"`
	AccountingExchangeRate          *float32        `json:"AccountingExchangeRate"`
	BillingDocumentDate             *string         `json:"BillingDocumentDate"`
	IsExportImportDelivery          *bool           `json:"IsExportImportDelivery"`
	HeaderText                      string          `json:"HeaderText"`
	HeaderPartner                   []HeaderPartner `json:"HeaderPartner"`
	Address                         []Address       `json:"Address"`
	HeaderPDF                       []HeaderPDF     `json:"HeaderPDF"`
	Item                            []Item          `json:"Item"`
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

type HeaderPartnerPlant struct {
	Plant string `json:"Plant"`
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
	Product                                       string               `json:"Product"`
	ProductStandardID                             string               `json:"ProductStandardID"`
	ProductGroup                                  string               `json:"ProductGroup"`
	BaseUnit                                      string               `json:"BaseUnit"`
	PricingDate                                   *string              `json:"PricingDate"`
	PriceDetnExchangeRate                         *float32             `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate                         *string              `json:"RequestedDeliveryDate"`
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
	BillingDocumentDate                           *string              `json:"BillingDocumentDate"`
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
	DocumentRjcnReason                            *bool                `json:"DocumentRjcnReason"`
	ItemBillingBlockReason                        *bool                `json:"ItemBillingBlockReason"`
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

type ItemPartner struct {
	PartnerFunction  string           `json:"PartnerFunction"`
	BusinessPartner  *int             `json:"BusinessPartner"`
	ItemPartnerPlant ItemPartnerPlant `json:"ItemPartnerPlant"`
}

type ItemPartnerPlant struct {
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
	DelivBlockReasonForSchedLine                 *bool    `json:"DelivBlockReasonForSchedLine"`
	PlusMinusFlag                                string   `json:"PlusMinusFlag"`
}
