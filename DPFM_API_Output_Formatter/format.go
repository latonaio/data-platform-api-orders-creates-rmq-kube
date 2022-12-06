package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-orders-creates-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-orders-creates-rmq-kube/sub_func_complementer"
)

func ConvertToHeaderCreates(subfuncSDC *sub_func_complementer.SDC) *HeaderCreates {
	data := subfuncSDC.Message.Header

	headerCreate := &HeaderCreates{
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

	return headerCreate
}

func ConvertToHeaderUpdates(headerUpdates dpfm_api_input_reader.HeaderUpdates) *HeaderUpdates {
	data := headerUpdates

	return &HeaderUpdates{
		OrderID:                         data.OrderID,
		TotalNetAmount:                  data.TotalNetAmount,
		TotalTaxAmount:                  data.TotalTaxAmount,
		TotalGrossAmount:                data.TotalGrossAmount,
		TotalBlockStatus:                data.TotalBlockStatus,
		TransactionCurrency:             data.TransactionCurrency,
		PricingDate:                     data.PricingDate,
		PriceDetnExchangeRate:           data.PriceDetnExchangeRate,
		RequestedDeliveryDate:           data.RequestedDeliveryDate,
		HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
		HeaderBillingBlockReason:        data.HeaderBillingBlockReason,
		DeliveryBlockReason:             data.DeliveryBlockReason,
		Incoterms:                       data.Incoterms,
		PaymentTerms:                    data.PaymentTerms,
		BillingDocumentDate:             data.BillingDocumentDate,
		HeaderText:                      data.HeaderText,
	}
}

func ConvertToHeaderPartner(subfuncSDC *sub_func_complementer.SDC) []HeaderPartner {
	var headerPartner []HeaderPartner

	for _, data := range subfuncSDC.Message.HeaderPartner {
		headerPartner = append(headerPartner, HeaderPartner{
			OrderID:                 data.OrderID,
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
		})
	}

	return headerPartner
}

func ConvertToHeaderPartnerPlant(subfuncSDC *sub_func_complementer.SDC) []HeaderPartnerPlant {
	var headerPartnerPlant []HeaderPartnerPlant

	for _, data := range subfuncSDC.Message.HeaderPartnerPlant {
		headerPartnerPlant = append(headerPartnerPlant, HeaderPartnerPlant{
			OrderID:         data.OrderID,
			PartnerFunction: data.PartnerFunction,
			BusinessPartner: data.BusinessPartner,
			Plant:           data.Plant,
		})
	}

	return headerPartnerPlant
}

func ConvertToItem(subfuncSDC *sub_func_complementer.SDC) []Item {
	var item []Item

	for _, data := range subfuncSDC.Message.Item {
		item = append(item, Item{
			OrderID:                          data.OrderID,
			OrderItem:                        data.OrderItem,
			OrderItemCategory:                data.OrderItemCategory,
			OrderItemText:                    data.OrderItemText,
			Product:                          data.Product,
			ProductStandardID:                data.ProductStandardID,
			ProductGroup:                     data.ProductGroup,
			BaseUnit:                         data.BaseUnit,
			PricingDate:                      data.PricingDate,
			PriceDetnExchangeRate:            data.PriceDetnExchangeRate,
			RequestedDeliveryDate:            data.RequestedDeliveryDate,
			StockConfirmationPartnerFunction: data.StockConfirmationPartnerFunction,
			StockConfirmationBusinessPartner: data.StockConfirmationBusinessPartner,
			StockConfirmationPlant:           data.StockConfirmationPlant,
			StockConfirmationPlantBatch:      data.StockConfirmationPlantBatch,
			StockConfirmationPlantBatchValidityStartDate:  data.StockConfirmationPlantBatchValidityStartDate,
			StockConfirmationPlantBatchValidityEndDate:    data.StockConfirmationPlantBatchValidityEndDate,
			ProductIsBatchManagedInStockConfirmationPlant: data.ProductIsBatchManagedInStockConfirmationPlant,
			OrderQuantityInBaseUnit:                       data.OrderQuantityInBaseUnit,
			OrderQuantityInIssuingUnit:                    data.OrderQuantityInIssuingUnit,
			OrderQuantityInReceivingUnit:                  data.OrderQuantityInReceivingUnit,
			OrderIssuingUnit:                              data.OrderIssuingUnit,
			OrderReceivingUnit:                            data.OrderReceivingUnit,
			StockConfirmationPolicy:                       data.StockConfirmationPolicy,
			StockConfirmationStatus:                       data.StockConfirmationStatus,
			ConfdDelivQtyInOrderQtyUnit:                   data.ConfdDelivQtyInOrderQtyUnit,
			ItemWeightUnit:                                data.ItemWeightUnit,
			ProductGrossWeight:                            data.ProductGrossWeight,
			ItemGrossWeight:                               data.ItemGrossWeight,
			ProductNetWeight:                              data.ProductNetWeight,
			ItemNetWeight:                                 data.ItemNetWeight,
			NetAmount:                                     data.NetAmount,
			TaxAmount:                                     data.TaxAmount,
			GrossAmount:                                   data.GrossAmount,
			BillingDocumentDate:                           data.BillingDocumentDate,
			ProductionPlantPartnerFunction:                data.ProductionPlantPartnerFunction,
			ProductionPlantBusinessPartner:                data.ProductionPlantBusinessPartner,
			ProductionPlant:                               data.ProductionPlant,
			ProductionPlantTimeZone:                       data.ProductionPlantTimeZone,
			ProductionPlantStorageLocation:                data.ProductionPlantStorageLocation,
			IssuingPlantPartnerFunction:                   data.IssuingPlantPartnerFunction,
			IssuingPlantBusinessPartner:                   data.IssuingPlantBusinessPartner,
			IssuingPlant:                                  data.IssuingPlant,
			IssuingPlantTimeZone:                          data.IssuingPlantTimeZone,
			IssuingPlantStorageLocation:                   data.IssuingPlantStorageLocation,
			ReceivingPlantPartnerFunction:                 data.ReceivingPlantPartnerFunction,
			ReceivingPlantBusinessPartner:                 data.ReceivingPlantBusinessPartner,
			ReceivingPlant:                                data.ReceivingPlant,
			ReceivingPlantTimeZone:                        data.ReceivingPlantTimeZone,
			ReceivingPlantStorageLocation:                 data.ReceivingPlantStorageLocation,
			ProductIsBatchManagedInProductionPlant:        data.ProductIsBatchManagedInProductionPlant,
			ProductIsBatchManagedInIssuingPlant:           data.ProductIsBatchManagedInIssuingPlant,
			ProductIsBatchManagedInReceivingPlant:         data.ProductIsBatchManagedInReceivingPlant,
			BatchMgmtPolicyInProductionPlant:              data.BatchMgmtPolicyInProductionPlant,
			BatchMgmtPolicyInIssuingPlant:                 data.BatchMgmtPolicyInIssuingPlant,
			BatchMgmtPolicyInReceivingPlant:               data.BatchMgmtPolicyInReceivingPlant,
			ProductionPlantBatch:                          data.ProductionPlantBatch,
			IssuingPlantBatch:                             data.IssuingPlantBatch,
			ReceivingPlantBatch:                           data.ReceivingPlantBatch,
			ProductionPlantBatchValidityStartDate:         data.ProductionPlantBatchValidityStartDate,
			ProductionPlantBatchValidityEndDate:           data.ProductionPlantBatchValidityEndDate,
			IssuingPlantBatchValidityStartDate:            data.IssuingPlantBatchValidityStartDate,
			IssuingPlantBatchValidityEndDate:              data.IssuingPlantBatchValidityEndDate,
			ReceivingPlantBatchValidityStartDate:          data.ReceivingPlantBatchValidityStartDate,
			ReceivingPlantBatchValidityEndDate:            data.ReceivingPlantBatchValidityEndDate,
			Incoterms:                                     data.Incoterms,
			BPTaxClassification:                           data.BPTaxClassification,
			ProductTaxClassification:                      data.ProductTaxClassification,
			BPAccountAssignmentGroup:                      data.BPAccountAssignmentGroup,
			ProductAccountAssignmentGroup:                 data.ProductAccountAssignmentGroup,
			PaymentTerms:                                  data.PaymentTerms,
			PaymentMethod:                                 data.PaymentMethod,
			DocumentRjcnReason:                            data.DocumentRjcnReason,
			ItemBillingBlockReason:                        data.ItemBillingBlockReason,
			Project:                                       data.Project,
			AccountingExchangeRate:                        data.AccountingExchangeRate,
			ReferenceDocument:                             data.ReferenceDocument,
			ReferenceDocumentItem:                         data.ReferenceDocumentItem,
			ItemCompleteDeliveryIsDefined:                 data.ItemCompleteDeliveryIsDefined,
			ItemDeliveryStatus:                            data.ItemDeliveryStatus,
			IssuingStatus:                                 data.IssuingStatus,
			ReceivingStatus:                               data.ReceivingStatus,
			BillingStatus:                                 data.BillingStatus,
			TaxCode:                                       data.TaxCode,
			TaxRate:                                       data.TaxRate,
			CountryOfOrigin:                               data.CountryOfOrigin,
		})
	}

	return item
}
