package requests

type ItemSchedulingLine struct {
	OrderID                                      *int     `json:"OrderID"`
	OrderItem                                    *int     `json:"OrderItem"`
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
