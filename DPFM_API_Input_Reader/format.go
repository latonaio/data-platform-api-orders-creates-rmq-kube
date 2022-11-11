package dpfm_api_input_reader

import (
	"data-platform-api-orders-creates-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.Orders
	return &requests.Header{
		OrderID:                         data.OrderID,
		OrderDate:                       data.OrderDate,
		OrderType:                       data.OrderType,
		Buyer:                           data.Buyer,
		Seller:                          data.Seller,
		CreationDate:                    data.CreationDate,
		LastChangeDate:                  data.LastChangeDate,
		ContractType:                    data.ContractType,
		ValidityStartDate:               data.ValidityStartDate,
		ValidityEndDate:                 data.ValidityEndDate,
		InvoiceScheduleStartDate:        data.InvoiceScheduleStartDate,
		InvoiceScheduleEndDate:          data.InvoiceScheduleEndDate,
		TotalNetAmount:                  data.TotalNetAmount,
		TotalTaxAmount:                  data.TotalTaxAmount,
		TotalGrossAmount:                data.TotalGrossAmount,
		OverallDeliveryStatus:           data.OverallDeliveryStatus,
		TotalBlockStatus:                data.TotalBlockStatus,
		OverallOrdReltdBillgStatus:      data.OverallOrdReltdBillgStatus,
		OverallDocReferenceStatus:       data.OverallDocReferenceStatus,
		TransactionCurrency:             data.TransactionCurrency,
		PricingDate:                     data.PricingDate,
		PriceDetnExchangeRate:           data.PriceDetnExchangeRate,
		RequestedDeliveryDate:           data.RequestedDeliveryDate,
		HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
		HeaderBillingBlockReason:        data.HeaderBillingBlockReason,
		DeliveryBlockReason:             data.DeliveryBlockReason,
		Incoterms:                       data.Incoterms,
		PaymentTerms:                    data.PaymentTerms,
		PaymentMethod:                   data.PaymentMethod,
		ReferenceDocument:               data.ReferenceDocument,
		ReferenceDocumentItem:           data.ReferenceDocumentItem,
		BPAccountAssignmentGroup:        data.BPAccountAssignmentGroup,
		AccountingExchangeRate:          data.AccountingExchangeRate,
		BillingDocumentDate:             data.BillingDocumentDate,
		IsExportImportDelivery:          data.IsExportImportDelivery,
		HeaderText:                      data.HeaderText,
	}
}

func (sdc *SDC) ConvertToHeaderPartner(num int) *requests.HeaderPartner {
	dataOrders := sdc.Orders
	data := sdc.Orders.HeaderPartner[num]
	return &requests.HeaderPartner{
		OrderID:                 dataOrders.OrderID,
		PartnerFunction:         data.PartnerFunction,
		BusinessPartner:         data.BusinessPartner,
		BusinessPartnerFullName: data.BusinessPartnerFullName,
		BusinessPartnerName:     data.BusinessPartnerName,
		Organization:            data.Organization,
		Country:                 data.Country,
		Language:                data.Language,
		Currency:                data.Currency,
		ExternalDocumentID:      data.ExternalDocumentID,
		AddressID:               data.AddressID,
	}
}
