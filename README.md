# data-platform-api-orders-creates-rmq-kube

data-platform-api-orders-creates-rmq-kube は、周辺業務システム　を データ連携基盤 と統合することを目的に、API でオーダーデータを取得するマイクロサービスです。  
https://xxx.xxx.io/api/API_ORDERS_SRV/creates/

## 動作環境

data-platform-api-orders-creates-rmq-kube の動作環境は、次の通りです。  
・ OS: LinuxOS （必須）  
・ CPU: ARM/AMD/Intel（いずれか必須）  


## 本レポジトリ が 対応する API サービス
data-platform-api-orders-creates-rmq-kube が対応する APIサービス は、次のものです。

APIサービス URL: https://xxx.xxx.io/api/API_ORDERS_SRV/creates/

## 本レポジトリ に 含まれる API名
data-platform-api-orders-creates-rmq-kube には、次の API をコールするためのリソースが含まれています。  

* A_Header（オーダー - ヘッダデータ）
* A_HeaderPartner（オーダー - ヘッダ取引先データ）
* A_HeaderPartnerPlant（オーダー - ヘッダ取引先プラントデータ）
* A_HeaderPartnerContact（オーダー - ヘッダ取引先コンタクトデータ）
* A_Item（オーダー - 明細データ）
* A_ItemPartner（オーダー - 明細取引先データ）
* A_ItemPartnerPlant（オーダー - 明細取引先プラントデータ）
* A_ItemPricingElement（オーダー - 明細取引先プラントデータ）
* A_ItemScheduleLine（オーダー - 明細納入日程行データ）
* A_Address（オーダー - 住所データ）

## API への 値入力条件 の 初期値
data-platform-api-orders-creates-rmq-kube において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

## データ連携基盤のAPIの選択的コール

Latona および AION の データ連携基盤 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。    
  
```
	"api_schema": "DPFMOrdersCreates",
	"accepter": ["Header"],
	"order_id": null,
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "DPFMOrdersCreates",
	"accepter": ["All"],
	"order_id": null,
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて DPFM_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *DPFMAPICaller) AsyncOrderCreates(
	accepter []string,
	input *dpfm_api_input_reader.SDC,

	log *logger.Logger,
) []error {
	wg := sync.WaitGroup{}
	mtx := sync.Mutex{}
	errs := make([]error, 0, 5)
	exconfAllExist := false

	subFuncFin := make(chan error)
	exconfFin := make(chan error)

	wg.Add(1)
	go func() {
		defer wg.Done()
		var e []error
		exconfAllExist, e = c.confirmor.Conf(input, log)
		if len(e) != 0 {
			mtx.Lock()
			errs = append(errs, e...)
			mtx.Unlock()
			exconfFin <- xerrors.Errorf("exconf error")
			return
		}
		exconfFin <- nil
	}()

	for _, fn := range accepter {
		wg.Add(1)
		switch fn {
		case "Header":
			go c.headerCreate(&wg, &mtx, subFuncFin, log, errs, input)
		case "Item":
			errs = append(errs, xerrors.Errorf("accepter Item is not implement yet"))
		default:
			wg.Done()
		}
	}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-data-platform](https://github.com/latonaio/golang-logging-library-for-data-platform) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は オーダー の ヘッダデータ が取得された結果の JSON の例です。  
以下の項目のうち、"OrderID" ～ "PlusMinusFlag" は、/DPFM_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
{
  "cursor": "/go/src/github.com/latonaio/main.go#L101",
  "function": "main.callProcess",
  "level": "INFO",
  "message": {
  	"connection_key": "request",
	"result": true,
	"redis_key": "abcdefg",
	"filepath": "/var/lib/aion/Data/rededge_sdc/abcdef.json",
	"runtime_session_id":"boi9ar543dg91ipdnspi099u231280ab0v8af0ew",
	"business_partner": 201,
	"service_label": "ORDERS",
    "Orders": {
      "OrderID": 11,
      "OrderDate": "",
      "OrderType": "",
      "Buyer": 101,
      "Seller": 201,
      "CreationDate": "",
      "LastChangeDate": "",
      "ContractType": "",
      "ValidityStartDate": "",
      "ValidityEndDate": "",
      "InvoiceScheduleStartDate": "",
      "InvoiceScheduleEndDate": "",
      "TotalNetAmount": null,
      "TotalTaxAmount": null,
      "TotalGrossAmount": null,
      "OverallDeliveryStatus": "",
      "TotalBlockStatus": null,
      "OverallOrdReltdBillgStatus": "",
      "OverallDocReferenceStatus": "",
      "TransactionCurrency": "",
      "PricingDate": "",
      "PriceDetnExchangeRate": null,
      "RequestedDeliveryDate": "",
      "HeaderCompleteDeliveryIsDefined": null,
      "HeaderBillingBlockReason": null,
      "DeliveryBlockReason": null,
      "Incoterms": "CIF",
      "PaymentTerms": "0001",
      "PaymentMethod": "T",
      "ReferenceDocument": null,
      "ReferenceDocumentItem": null,
      "BPAccountAssignmentGroup": "01",
      "AccountingExchangeRate": null,
      "BillingDocumentDate": "",
      "IsExportImportDelivery": null,
      "HeaderText": "",
      "HeaderPartner": [
        {
          "PartnerFunction": "DELIVERTO",
          "BusinessPartner": 102,
          "BusinessPartnerFullName": "株式会社ABC虎ノ門店",
          "BusinessPartnerName": "ABC虎ノ門店",
          "Organization": "",
          "Country": "JP",
          "Language": "JA",
          "Currency": "JPY",
          "ExternalDocumentID": "",
          "AddressID": 200000,
          "HeaderPartnerContact": null,
          "HeaderPartnerPlant": [
            {
              "Plant": "AB02"
            }
          ]
        },
        {
          "PartnerFunction": "BUYER",
          "BusinessPartner": 101,
          "BusinessPartnerFullName": "株式会社ABC本社",
          "BusinessPartnerName": "ABC本社",
          "Organization": "",
          "Country": "JP",
          "Language": "JA",
          "Currency": "JPY",
          "ExternalDocumentID": "",
          "AddressID": 100000,
          "HeaderPartnerContact": [
            {
              "ContactID": null,
              "ContactPersonName": "",
              "EmailAddress": "",
              "PhoneNumber": "",
              "MobilePhoneNumber": "",
              "FaxNumber": "",
              "ContactTag1": "",
              "ContactTag2": "",
              "ContactTag3": "",
              "ContactTag4": ""
            }
          ],
          "HeaderPartnerPlant": [
            {
              "Plant": "AB01"
            }
          ]
        },
        {
          "PartnerFunction": "SELLER",
          "BusinessPartner": 201,
          "BusinessPartnerFullName": "パン販売株式会社",
          "BusinessPartnerName": "パン販売",
          "Organization": "",
          "Country": "JP",
          "Language": "JA",
          "Currency": "JPY",
          "ExternalDocumentID": "",
          "AddressID": 300000,
          "HeaderPartnerContact": [
            {
              "ContactID": null,
              "ContactPersonName": "",
              "EmailAddress": "",
              "PhoneNumber": "",
              "MobilePhoneNumber": "",
              "FaxNumber": "",
              "ContactTag1": "",
              "ContactTag2": "",
              "ContactTag3": "",
              "ContactTag4": ""
            }
          ],
          "HeaderPartnerPlant": [
            {
              "Plant": "TE01"
            }
          ]
        }
      ],
      "Address": [
        {
          "AddressID": null,
          "PostalCode": "",
          "LocalRegion": "",
          "Country": "",
          "District": "",
          "StreetName": "",
          "CityName": "",
          "Building": "",
          "Floor": null,
          "Room": null
        }
      ],
      "HeaderPDF": [
        {
          "DocType": "",
          "DocVersionID": null,
          "DocID": "",
          "DocIssuerBusinessPartner": null,
          "FileName": ""
        }
      ],
      "Item": [
        {
          "OrderItem": 1,
          "OrderItemCategory": "",
          "OrderItemText": "RobotA",
          "Product": "A3750",
          "ProductStandardID": "CDC",
          "ProductGroup": "01",
          "BaseUnit": "PC",
          "PricingDate": "",
          "PriceDetnExchangeRate": null,
          "RequestedDeliveryDate": "",
          "StockConfirmationPartnerFunction": "",
          "StockConfirmationBusinessPartner": null,
          "StockConfirmationPlant": "",
          "StockConfirmationPlantBatch": "",
          "StockConfirmationPlantBatchValidityStartDate": "",
          "StockConfirmationPlantBatchValidityEndDate": "",
          "ProductIsBatchManagedInStockConfirmationPlant": null,
          "OrderQuantityInBaseUnit": null,
          "OrderQuantityInIssuingUnit": null,
          "OrderQuantityInReceivingUnit": null,
          "OrderIssuingUnit": "",
          "OrderReceivingUnit": "",
          "StockConfirmationPolicy": "",
          "StockConfirmationStatus": "",
          "ConfdDelivQtyInOrderQtyUnit": null,
          "ItemWeightUnit": "G",
          "ProductGrossWeight": 300.05,
          "ItemGrossWeight": null,
          "ProductNetWeight": 300.05,
          "ItemNetWeight": null,
          "NetAmount": null,
          "TaxAmount": null,
          "GrossAmount": null,
          "BillingDocumentDate": "",
          "ProductionPlantPartnerFunction": "",
          "ProductionPlantBusinessPartner": null,
          "ProductionPlant": "",
          "ProductionPlantTimeZone": "",
          "ProductionPlantStorageLocation": "",
          "IssuingPlantPartnerFunction": "",
          "IssuingPlantBusinessPartner": null,
          "IssuingPlant": "",
          "IssuingPlantTimeZone": "",
          "IssuingPlantStorageLocation": "",
          "ReceivingPlantPartnerFunction": "",
          "ReceivingPlantBusinessPartner": null,
          "ReceivingPlant": "",
          "ReceivingPlantTimeZone": "",
          "ReceivingPlantStorageLocation": "",
          "ProductIsBatchManagedInProductionPlant": null,
          "ProductIsBatchManagedInIssuingPlant": null,
          "ProductIsBatchManagedInReceivingPlant": null,
          "BatchMgmtPolicyInProductionPlant": "",
          "BatchMgmtPolicyInIssuingPlant": "",
          "BatchMgmtPolicyInReceivingPlant": "",
          "ProductionPlantBatch": "",
          "IssuingPlantBatch": "",
          "ReceivingPlantBatch": "",
          "ProductionPlantBatchValidityStartDate": "",
          "ProductionPlantBatchValidityEndDate": "",
          "IssuingPlantBatchValidityStartDate": "",
          "IssuingPlantBatchValidityEndDate": "",
          "ReceivingPlantBatchValidityStartDate": "",
          "ReceivingPlantBatchValidityEndDate": "",
          "Incoterms": "",
          "BPTaxClassification": "1",
          "ProductTaxClassification": "",
          "BPAccountAssignmentGroup": "",
          "ProductAccountAssignmentGroup": "01",
          "PaymentTerms": "",
          "PaymentMethod": "",
          "DocumentRjcnReason": null,
          "ItemBillingBlockReason": null,
          "Project": "",
          "AccountingExchangeRate": null,
          "ReferenceDocument": null,
          "ReferenceDocumentItem": null,
          "ItemCompleteDeliveryIsDefined": null,
          "ItemDeliveryStatus": "",
          "IssuingStatus": "",
          "ReceivingStatus": "",
          "BillingStatus": "",
          "TaxCode": "",
          "TaxRate": null,
          "CountryOfOrigin": "JPY",
          "ItemPartner": [
            {
              "PartnerFunction": "",
              "BusinessPartner": null,
              "ItemPartnerPlant": {
                "Plant": ""
              }
            }
          ],
          "ItemPricingElement": [
            {
              "PricingProcedureStep": null,
              "PricingProcedureCounter": null,
              "ConditionType": "",
              "PricingDate": "",
              "ConditionRateValue": null,
              "ConditionCurrency": "",
              "ConditionQuantity": null,
              "ConditionQuantityUnit": "",
              "ConditionRecord": null,
              "ConditionSequentialNumber": null,
              "TaxCode": "",
              "ConditionAmount": null,
              "TransactionCurrency": "",
              "ConditionIsManuallyChanged": null
            }
          ],
          "ItemSchedulingLine": [
            {
              "ScheduleLine": null,
              "Product": "",
              "StockConfirmationPartnerFunction": "",
              "StockConfirmationBusinessPartner": null,
              "StockConfirmationPlant": "",
              "StockConfirmationPlantBatch": "",
              "StockConfirmationPlantBatchValidityStartDate": "",
              "StockConfirmationPlantBatchValidityEndDate": "",
              "ConfirmedDeliveryDate": "",
              "RequestedDeliveryDate": "",
              "OrderQuantityInBaseUnit": null,
              "ConfdOrderQtyByPDTAvailCheck": null,
              "DeliveredQtyInOrderQtyUnit": null,
              "OpenConfdDelivQtyInOrdQtyUnit": null,
              "DelivBlockReasonForSchedLine": null,
              "PlusMinusFlag": ""
            }
          ]
        },
        {
          "OrderItem": 2,
          "OrderItemCategory": "",
          "OrderItemText": "RobotB",
          "Product": "W999",
          "ProductStandardID": "CAC",
          "ProductGroup": "01",
          "BaseUnit": "PC",
          "PricingDate": "",
          "PriceDetnExchangeRate": null,
          "RequestedDeliveryDate": "",
          "StockConfirmationPartnerFunction": "",
          "StockConfirmationBusinessPartner": null,
          "StockConfirmationPlant": "",
          "StockConfirmationPlantBatch": "",
          "StockConfirmationPlantBatchValidityStartDate": "",
          "StockConfirmationPlantBatchValidityEndDate": "",
          "ProductIsBatchManagedInStockConfirmationPlant": null,
          "OrderQuantityInBaseUnit": null,
          "OrderQuantityInIssuingUnit": null,
          "OrderQuantityInReceivingUnit": null,
          "OrderIssuingUnit": "",
          "OrderReceivingUnit": "",
          "StockConfirmationPolicy": "",
          "StockConfirmationStatus": "",
          "ConfdDelivQtyInOrderQtyUnit": null,
          "ItemWeightUnit": "G",
          "ProductGrossWeight": 100.05,
          "ItemGrossWeight": null,
          "ProductNetWeight": 100.05,
          "ItemNetWeight": null,
          "NetAmount": null,
          "TaxAmount": null,
          "GrossAmount": null,
          "BillingDocumentDate": "",
          "ProductionPlantPartnerFunction": "",
          "ProductionPlantBusinessPartner": null,
          "ProductionPlant": "",
          "ProductionPlantTimeZone": "",
          "ProductionPlantStorageLocation": "",
          "IssuingPlantPartnerFunction": "",
          "IssuingPlantBusinessPartner": null,
          "IssuingPlant": "",
          "IssuingPlantTimeZone": "",
          "IssuingPlantStorageLocation": "",
          "ReceivingPlantPartnerFunction": "",
          "ReceivingPlantBusinessPartner": null,
          "ReceivingPlant": "",
          "ReceivingPlantTimeZone": "",
          "ReceivingPlantStorageLocation": "",
          "ProductIsBatchManagedInProductionPlant": null,
          "ProductIsBatchManagedInIssuingPlant": null,
          "ProductIsBatchManagedInReceivingPlant": null,
          "BatchMgmtPolicyInProductionPlant": "",
          "BatchMgmtPolicyInIssuingPlant": "",
          "BatchMgmtPolicyInReceivingPlant": "",
          "ProductionPlantBatch": "",
          "IssuingPlantBatch": "",
          "ReceivingPlantBatch": "",
          "ProductionPlantBatchValidityStartDate": "",
          "ProductionPlantBatchValidityEndDate": "",
          "IssuingPlantBatchValidityStartDate": "",
          "IssuingPlantBatchValidityEndDate": "",
          "ReceivingPlantBatchValidityStartDate": "",
          "ReceivingPlantBatchValidityEndDate": "",
          "Incoterms": "",
          "BPTaxClassification": "1",
          "ProductTaxClassification": "",
          "BPAccountAssignmentGroup": "",
          "ProductAccountAssignmentGroup": "01",
          "PaymentTerms": "",
          "PaymentMethod": "",
          "DocumentRjcnReason": null,
          "ItemBillingBlockReason": null,
          "Project": "",
          "AccountingExchangeRate": null,
          "ReferenceDocument": null,
          "ReferenceDocumentItem": null,
          "ItemCompleteDeliveryIsDefined": null,
          "ItemDeliveryStatus": "",
          "IssuingStatus": "",
          "ReceivingStatus": "",
          "BillingStatus": "",
          "TaxCode": "",
          "TaxRate": null,
          "CountryOfOrigin": "JPY",
          "ItemPartner": [
            {
              "PartnerFunction": "",
              "BusinessPartner": null,
              "ItemPartnerPlant": {
                "Plant": ""
              }
            }
          ],
          "ItemPricingElement": [
            {
              "PricingProcedureStep": null,
              "PricingProcedureCounter": null,
              "ConditionType": "",
              "PricingDate": "",
              "ConditionRateValue": null,
              "ConditionCurrency": "",
              "ConditionQuantity": null,
              "ConditionQuantityUnit": "",
              "ConditionRecord": null,
              "ConditionSequentialNumber": null,
              "TaxCode": "",
              "ConditionAmount": null,
              "TransactionCurrency": "",
              "ConditionIsManuallyChanged": null
            }
          ],
          "ItemSchedulingLine": [
            {
              "ScheduleLine": null,
              "Product": "",
              "StockConfirmationPartnerFunction": "",
              "StockConfirmationBusinessPartner": null,
              "StockConfirmationPlant": "",
              "StockConfirmationPlantBatch": "",
              "StockConfirmationPlantBatchValidityStartDate": "",
              "StockConfirmationPlantBatchValidityEndDate": "",
              "ConfirmedDeliveryDate": "",
              "RequestedDeliveryDate": "",
              "OrderQuantityInBaseUnit": null,
              "ConfdOrderQtyByPDTAvailCheck": null,
              "DeliveredQtyInOrderQtyUnit": null,
              "OpenConfdDelivQtyInOrdQtyUnit": null,
              "DelivBlockReasonForSchedLine": null,
              "PlusMinusFlag": ""
            }
          ]
        }
      ]
    },
    "api_schema": "DPFMOrdersCreates",
    "accepter": ["All"],
    "order_id": null,
    "deleted": false
  },
  "runtime_session_id": "boi9ar543dg91ipdnspi099u231280ab0v8af0ew",
  "time": "2022-11-08T11:21:23+09:00"
}
```
