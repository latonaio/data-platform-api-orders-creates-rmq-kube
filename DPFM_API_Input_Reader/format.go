package dpfm_api_input_reader

import (
	"data-platform-api-orders-creates-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToBpExistenceConf() {

}

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.Header
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
	dataOrders := sdc.Header
	data := sdc.Header.HeaderPartner[num]
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

func (sdc *SDC) ConvertToHeaderPartnerPlant(hpNum, hppNum int) *requests.HeaderPartnerPlant {
	dataOrders := sdc.Header
	dataHeaderPartner := sdc.Header.HeaderPartner[hpNum]
	data := dataHeaderPartner.HeaderPartnerPlant[hppNum]
	return &requests.HeaderPartnerPlant{
		OrderID:         dataOrders.OrderID,
		PartnerFunction: dataHeaderPartner.PartnerFunction,
		BusinessPartner: dataHeaderPartner.BusinessPartner,
		Plant:           data.Plant,
	}
}

// func (sdc *SDC) ConvertToHeaderPartnerContact(hpNum, hpcNum int) *requests.HeaderPartnerContact {
// 	dataOrders := sdc.Orders
// 	dataHeaderPartner := sdc.Orders.HeaderPartner[hpNum]
// 	data := dataHeaderPartner.HeaderPartnerContact[hpcNum]
// 	return &requests.HeaderPartnerContact{
// 		OrderID:           dataOrders.OrderID,
// 		PartnerFunction:   dataHeaderPartner.PartnerFunction,
// 		BusinessPartner:   dataHeaderPartner.BusinessPartner,
// 		ContactID:         data.ContactID,
// 		ContactPersonName: data.ContactPersonName,
// 		EmailAddress:      data.EmailAddress,
// 		PhoneNumber:       data.PhoneNumber,
// 		MobilePhoneNumber: data.MobilePhoneNumber,
// 		FaxNumber:         data.FaxNumber,
// 		ContactTag1:       data.ContactTag1,
// 		ContactTag2:       data.ContactTag2,
// 		ContactTag3:       data.ContactTag3,
// 		ContactTag4:       data.ContactTag4,
// 	}
// }
