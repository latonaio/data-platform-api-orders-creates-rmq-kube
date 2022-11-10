# data-platform-api-orders-creates-rmq-kube

data-platform-api-orders-creates-rmq-kube は、周辺業務システム　を データ連携基盤 と統合することを目的に、API でオーダーデータを取得するマイクロサービスです。  
https://xxx.xxx.io/api/API_ORDERS_SRV/creates/

## 動作環境

data-platform-api-orders-creates-rmq-kube の動作環境は、次の通りです。  
・ OS: LinuxOS （必須）  
・ CPU: ARM/AMD/Intel（いずれか必須）  


## 本レポジトリ が 対応する API サービス
data-platform-api-orders-creates-rmq-kube が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://xxx.xxx.io/api/API_ORDERS_SRV/creates/
* APIサービス名(=baseURL): API_ORDERS_SRV

## 本レポジトリ に 含まれる API名
data-platform-api-orders-creates-rmq-kube には、次の API をコールするためのリソースが含まれています。  

* A_Product（品目マスタ - 一般データ）
* A_ProductPlant（品目マスタ - プラントデータ）
* A_ProductPlantMRPArea（品目マスタ - MRPエリアデータ）
* A_ProductPlantProcurement（品目マスタ - 購買データ）
* A_ProductWorkScheduling（品目マスタ - 作業計画データ）
* A_ProductPlantSales（品目マスタ - 販売プラントデータ）
* A_ProductValuation（品目マスタ - 評価エリアデータ）
* A_ProductSalesDelivery（品目マスタ - 販売組織データ）
* A_ProductPlantQualityMgmt（品目マスタ - 品質管理データ）
* A_ProductSalesTax（品目マスタ - 販売税データ）
* A_ProductDescription（品目マスタ - テキストデータ）
* ToProductDesc（品目マスタ - テキストデータ ※To）

## API への 値入力条件 の 初期値
data-platform-api-orders-creates-rmq-kube において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.Product.Product（品目）
* inoutSDC.Product.Plant.Plant（プラント）
* inoutSDC.Product.Plant.MRPArea.MRPArea（MRPエリア）
* inoutSDC.Product.Accounting.ValuationArea（評価エリア）
* inoutSDC.Product.SalesOrganization.ProductSalesOrg（販売組織）
* inoutSDC.Product.SalesOrganization.ProductDistributionChnl（流通チャネル）
* inoutSDC.Product.ProductDescription.Language（言語キー）
* inoutSDC.Product.ProductDescription.ProductDescription（品目テキスト）
* inoutSDC.Product.SalesTax.Country（国）
* inoutSDC.Product.SalesTax.TaxCategory（税カテゴリ）

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
func (c *DPFMAPICaller) AsyncGetProductMaster(product, plant, mrpArea, valuationArea, productSalesOrg, productDistributionChnl, language, productDescription string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "General":
			func() {
				c.General(product)
				wg.Done()
			}()
		case "Plant":
			func() {
				c.Plant(product, plant)
				wg.Done()
			}()
		case "MRPArea":
			func() {
				c.MRPArea(product, plant, mrpArea)
				wg.Done()
			}()
		case "Procurement":
			func() {
				c.Procurement(product, plant)
				wg.Done()
			}()
		case "WorkScheduling":
			func() {
				c.WorkScheduling(product, plant)
				wg.Done()
			}()
		case "SalesPlant":
			func() {
				c.SalesPlant(product, plant)
				wg.Done()
			}()
		case "Accounting":
			func() {
				c.Accounting(product, valuationArea)
				wg.Done()
			}()
		case "SalesOrganization":
			func() {
				c.SalesOrganization(product, productSalesOrg, productDistributionChnl)
				wg.Done()
			}()
		case "ProductDescByProduct":
			func() {
				c.ProductDescByProduct(product, language)
				wg.Done()
			}()
		case "ProductDescByDesc":
			func() {
				c.ProductDescByDesc(language, productDescription)
				wg.Done()
			}()
		case "Quality":
			func() {
				c.Quality(product, plant)
				wg.Done()
			}()
		case "SalesTax":
			func() {
				c.SalesTax(product, country, taxCategory)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、品目マスタ の 一般データ が取得された結果の JSON の例です。  
以下の項目のうち、"Material" ～ "ProductStandardID" は、/DPFM_API_Output_Formatter/type.go 内 の Type General {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/home/ampamman/go/src/data-platform-api-orders-creates-rmq-kube/DPFM_API_Caller/caller.go#L108",
	"function": "data-platform-api-orders-creates-rmq-kube/DPFM_API_Caller.(*DPFMAPICaller).General",
	"level": "INFO",
	"message": [
		{
			"Product": "21",
			"IndustrySector": "M",
			"ProductType": "FERT",
			"BaseUnit": "PC",
			"ValidityStartDate": "2022-01-25T09:00:00+09:00",
			"ProductGroup": "01",
			"Division": "",
			"GrossWeight": "2.000",
			"WeightUnit": "KG",
			"SizeOrDimensionText": "",
			"ProductStandardID": "",
			"CreationDate": "",
			"LastChangeDate": "2022-09-08T09:00:00+09:00",
			"IsMarkedForDeletion": false,
			"NetWeight": "1.000",
			"ChangeNumber": "",
			"to_Description": "http://XXX.XX.XX.XXX:8080/dpfm/opu/odata/dpfm/API_PRODUCT_SRV/READS/A_Product('21')/to_Description"
		}
	],
	"time": "2022-01-26T14:51:52.138052513+09:00"
}
```
