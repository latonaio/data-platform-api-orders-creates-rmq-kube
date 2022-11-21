# subfunc_specs
## 機能サービス名:【Function Orders > DPFM_FUNCTION_ORDERS_SRV】

## 機能名：【data_platform_function_orders_creates_subfunc】

## ＜機能概要＞
オーダー登録の補助機能。

## ＜仕様詳細＞
0. ReferenceDocumentおよびReferenceDocumentItemの値による登録種別 / 参照先伝票種別の判定

	0-1. ReferenceDocumentおよびReferenceDocumentItemの値による登録種別の判定  
	
	入力ファイルにおけるReferenceDocumentおよびReferenceDocumentItemの値の存在の有無によって、「参照登録」なのか、「直接登録」なのかを判別する。ReferenceDocumentの値がnull以外またはブランク以外、かつ、ReferenceDocumentItemの値がnull以外またはブランク以外、である場合、「参照登録」と判定する。それ以外の場合、「直接登録」と判定する。  
	
	0-2. ReferenceDocumentの値による参照先伝票種別の判定  　　
	
	同入力ファイルにおけるReferenceDocumentおよびReferenceDocumentItemの値によって、ユーザがどのサービスラベルの伝票を参照してオーダーを登録しようとしているのかを判別する。  

	0-2-1.番号範囲テーブルの全レコード/値を取得し、配列として保持する。  
	
	| Target / Processing Type                 | Key(s)                                       | 
	| ---------------------------------------- | -------------------------------------------- | 
	| ServiceLabel <br> FieldNameWithNumberRange <br> NumberRangeFrom <br> NumberRangeTo <br> / Get and Hold                           | None                                         | 
	|   <b>Table Searched</b>                  |   <b>Name of the Table</b>                   | 
	| Number Range SQL Number Range            | Data_platform_number_range_number_range_data | 
	| <b>Field Searched</b>                    | <b>Data Type / Number of Digits</b>          | 
	| ServiceLabel <br> FieldNameWithNumberRange <br> NumberRangeFrom  <br> NumberRangeTo                            | string(varchar) / 50 <br> string(varchar) / 100 <br> int / 10 <br> int / 10                                 | 
	| <b>Single Record or Array</b>            | <b>Memo</b>                                 | 
	| Array                                    | 対象テーブルの全レコードを取得               |   

	### [参考]
	
	以下の例のようなデータが取得される。  
		
	| NumberRangeID | ServiceLabel   | FieldNameWithNumberRange | NumberRangeFrom | NumberRangeTo | 
	| ------------- | -------------- | ------------------------ | --------------- | ------------- | 
	| “01”        | “ORDERS”     | OrderID                  | 10000000        | 19999999      | 
	| “02”        | “QUOTATIONS” | Quotation                | 20000000        | 29999999      | 
	| “08”        | “INQUIRIES”  | Inquiry                  | 80000000        | 89999999      |   
		
	0-2-2. 入力ファイルのReferenceDocumentが取得したテーブルのどの範囲に当てはまるかを判定し、その判定されたレコードのServiceLabelをセットする。
	
1. Orders Header に関する補助機能
Orders Header の項目別または項目群別に、次の補助機能の開発を行う。

	1-0. 入力ファイルのbusiness_partnerがBuyerであるかSellerであるかの判断

	入力ファイルのbusiness_partnerが、入力ファイルのBuyer, Sellerのいずれと一致するかを判断する。Buyerと一致した場合は、内部テーブル項目のBuyerOrSellerに”Buyer”をセットし、Sellerと一致した場合は、内部テーブル項目のBuyerOrSellerに”Seller”をセットする。BuyerにもSellerにも一致しない場合、もしくは、BuyerにもSellerにも一致する場合、エラーメッセージを出力して終了する。

	＜ロジックまとめ＞

	| ケース | 入力ファイルのbusiness_partnerがBuyerと一致するかどうか | 入力ファイルのbusiness_partnerがSellerと一致するかどうか | 内部テーブル項目のBuyerOrSellerにセットする値 or エラー | 
	| ------ | ------------------------------------------------------- | -------------------------------------------------------- | 	------------------------------------------------------- | 
	| A      | true                                                    | false                                                    | 	“Buyer”                                               | 
	| B      | false                                                   | true                                                     | 	“Seller”                                              | 
	| C      | false                                                   | false                                                    | エラ	ー                                                  | 
	| D      | true                                                    | true                                                     | エラ	ー                                                  | 

	1-1. ビジネスパートナ 得意先データ / 仕入先データ の取得

	| Property                 | Description                                                                                                                                                                                                                                          | EC  | 
	| ------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
	| Incoterms                | インコタームズ。貿易取引を行う際、輸送タイプに応じてインコタームズを指定する。ビジネスパートナマスタの得意先データまたは仕入先データのインコタームズが提案される。必要に応じて、インコタームズマスタから選択して指定する。参照の場合は変更できない。 | ✔  | 
	| PaymentTerms             | 支払条件。ビジネスパートナマスタの得意先データまたは仕入先データの支払条件が提案される。変更の必要があれば、支払条件マスタから選択して指定する。参照の場合は変更できない。                                                                           | ✔  | 
	| PaymentMethod            | 支払方法。ビジネスパートナマスタの得意先データまたは仕入先データの支払方法が提案される。変更の必要があれば、支払方法マスタから選択して指定する。参照の場合は変更できない。                                                                           | ✔  | 
	| BPAccountAssignmentGroup | ビジネスパートナ勘定設定グループ。ビジネスパートナ得意先データまたは仕入先データの勘定設定グループがコピーされる。変更不可。                                                                                                                         |     | 
	
	1-1-1. BusinessPartner=入力ファイルのbusiness_partner, Customer=入力ファイルのBuyer(1-0の処理結果が”Seller”の場合)、または、Supplier=入力ファイルのSeller(1-0の処理結果が”Buyer”の場合)をキーとして、対象のビジネスパートナ得意先データまたは仕入先データのIncoterms / PaymentTerms / PaymentMethod / BPAccountAssignmentGroupを検索して値をセットする。
	
	| Target / Processing Type                     | Key(s)                                              | 
	| -------------------------------------------- | --------------------------------------------------- | 
	| Incoterms <br> PaymentTerms <br> PaymentMethod <br> BPAccountAssignmentGroup<br> / Get and Set             | BusinessPartner=business_partner(入力ファイル)  <br> Customer=Buyer(入力ファイル)または、<br>Supplier= Seller(入力ファイル)               | 
	| <b>Table Searched</b>                        | <b>Name of the Table</b>                            | 
	| Business Partner SQL Customer　Dataまたは <br> Business Partner SQL Supplier　Data          | data_platform_business_partner_customer_data または <br> data_platform_business_partner_supplier_data | 
	| <b>Field Searched</b>                        | <b>Data Type / Number of Digits</b>                 | 
	| Incoterms <br> PaymentTerms <br> PaymentMethod <br> BPAccountAssignmentGroup                     | string(varchar) / 3 <br> string(varchar) / 4 <br> string(varchar) / 1 <br> string(varchar) / 2                          | 
	| <b>Single Record or Array</b>                | <b>Memo</b>                                        | 
	| Single Record                                |                                                     | 
	
	1-2. OrderID
	
	| Property                 | Description                                                                                                                                                                                                                                          | EC  | 
	| ------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
	| OrderID                | オーダー番号。自動採番される。 |    | 
	
	1-2-1. ServiceLabel=入力ファイルのservice_label, FieldNameWithNumberRange=当処理のProperty(=”OrderID”)をキーとして、対象のNumber Range Latest NumberのLatestNumberを検索し保持する。
	
	| Target / Processing Type                                | Key(s)                                        | 
	| ------------------------------------------------------- | --------------------------------------------- | 
	| ServiceLabel FieldNameWithNumberRange <br> LatestNumber <br> / Get and Hold                                          | ServiceLabel=service_label <br> FieldNameWithNumberRange =“OrderID”(当処理のProperty) | 
	| <b>Table Searched</b>                                   | <b>Name of the Table</b>                      | 
	| Number Range SQL Latest Number Data                     | data_platform_number_range_latest_number_data | 
	| <b>Field Searched</b>                                   | <b>Data Type / Number of Digits</b>           | 
	| ServiceLabel <br> FieldNameWithNumberRange <br> LatestNumber                                            | string(varchar) / 50 <br> string(varchar) / 100 <br> int / 16                                                | 
	| <b>Single Record or Array</b>                           | <b>Memo</b>                                  |  | 
	| Single Record                                           |                                               | 
	
	1-2-2. 保持されたLatestNumberに1を足したものをOrderIDにセットする。
	
	1-3. BillingDocumentDate
	
	| Property                 | Description                                                                                                                                                                                                                                          | EC  | 
	| ------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
	| BillingDocumentDate                | 請求書日付。支払条件の値により自動提案される。通常、請求書日付と同じ日付になる。任意の日付の場合は、その日付を入力する。 |    | 
	
	1-3-0. 入力ファイルのBillingDocumentDateがブランクである場合、次の処理を行う。そうでない場合、BillingDocumentDateの値をそのままセットする。
	
	1-3-1. BillingDocumentDate[請求書日付（＝請求書の締め日）]を計算するために、入力ファイルのRequestedDeliveryDateを保持する。
	
	| Target / Processing Type | Key(s)                       | 
	| ------------------------ | ---------------------------- | 
	| RequestedDeliveryDate <br> / Hold                   | None                         | 
	| <b>Table Searched</b>    | <b>Name of the Table</b>     | 
	| None                     | None                         | 
	| <b>Field Searched</b>    | <b>Data Type / Number of Digits</b> | 
	|None                     | string(varchar) / 80         | 
	| <b>Single Record or Array</b> | <b>Memo</b>                         | 
	| Single Record            |                              | 
	
	1-3-2. BillingDocumentDate[請求書日付（＝請求書の締め日）]を計算するために、1-1で取得したPaymentTermsをキーとして、対象の支払条件のDueDate, BaseDateCalcFixedDate, BaseDateCalcAddMonth、を検索して保持する。
	
	| Target / Processing Type             | Key(s)                                         |
	| ------------------------------------ | ---------------------------------------------- |
	| DueDate <br> BaseDateCalcFixedDate <br> BaseDateCalcAddMonth <br> / Get and Hold                       | 1-1で取得したPaymentTerms                      | 
	| <b>Table Searched</b>                | <b>Name of the Table</b>                       |
	| Payment Terms SQL Payment Terms Data | data_platform_payment_terms_payment_terms_data |
	| <b>Field Searched</b>                | <b>Data Type / Number of Digits</b>            |
	| DueDate <br> BaseDateCalcFixedDate <br> BaseDateCalcAddMonth                 | int / 2 <br> int / 2 <br> int / 2                              |                                                | 
	| <b>Single Record or Array</b>        | <b>Memo</b>                                   |
	| Single Record                        |                                                |
	
	1-3-3. RequestedDeliveryDate, DueDate, BaseDateCalcFixedDate, BaseDateCalcAddMonthをもとに、請求書日付（＝請求書の締め日）を計算し、BillingDocumentDateにセットする。
	e.g.) RequestedDeliveryDateが2022-03-08、DueDateが31、BaseDateCalcFixedDateが31, BaseDateCalcAddMonthが0、だった場合、請求書日付は2022-03-31と計算される。
	
	1-5. CompleteDeliveryIsDefined
	
	| Property                 | Description                                                                                                                                                                                                                                          | EC  | 
	| ------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
	| CompleteDeliveryIsDefined                | 入出荷完了ステータス。以下のステータスが設定される。変更不可。 <br> true : 完全入出荷完了 <br> false : 完全入出荷未完了
	 |    |
	
	 1-5-1. 1-4-1のデータを使って、次の判定とデータのセットを行う。
	検索結果が0件であった場合、CompleteDeliveryIsDefinedにfalseをセットする。検索結果が1件以上であり、かつ、全ての検索結果値がfalseの場合はfalseをセットする。検索結果が1件以上であり、かつ、 全ての検索結果値がtrueの場合はtrueをセットする。それ以外の場合は部分入出荷完了falseをセットする。
	
	<ロジックまとめ>
	
	| 検索結果件数 | 検索結果値(CompleteDeliveryIsDefined) | CompleteDeliveryIsDefinedにセットする値 | 
	| ------------ | ------------------------------------- | --------------------------------------- | 
	| 0件          | -                                     | false                                   | 
	| 1件以上      | 全てfalse                             | false                                   | 
	| 1件以上      | 全てtrue                              | true                                    | 
	| 1件以上      | 上記以外                              | false                                   | 
	
	1-7. OverallDocReferenceStatus
	
	| Property                 | Description                                                                                                                                                                                                                                          | EC  | 
	| ------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
	| OverallDocReferenceStatus                | 伝票参照ステータス。見積または引合もしくは購買依頼参照の場合、以下の参照ステータスが設定される。変更不可。 <br> 見積参照：QT <br> 引合参照：IN <br> 購買依頼参照：PR
	 |    | 
	
	1-7-1. 0-2-2でセットされたServiceLabelが”QUOTATIONS”である場合、”QT”をセットする。
	同ServiceLabelが、”INQUIRIES”である場合、”IN”をセットする。
	同ServiceLabelが、”PURCHASE_REQUISITION”である場合、”PR“をセットする。
	
	1-8. PricingDate
	
	| Property                 | Description                                                                                                                                                                                                                                          | EC  | 
	| ------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
	| PricingDate                | 価格設定日付。初期値としてはシステム日付が提案される。通常、オーダー日付と同じ日付を入力する。価格設定日付を任意で決めたい場合、その日付を入力する。参照の場合は変更できない。
	 |    | 
	
	1-8-1. [入力ファイルのPricingDate]がnullの場合、PricingDateに[システム日付]をセットする。
	
	1-8-2. [入力ファイルのPricingDate]がnullでない場合、PricingDateに[入力ファイルのPricingDate]をセットする。
	
	1-9. PriceDetnExchangeRate
	
	| Property              | Description                                                                              | EC  | 
	| --------------------- | ---------------------------------------------------------------------------------------- | --- | 
	| PriceDetnExchangeRate | 価格決定のための為替レート。必要な場合、為替レートを入力する。参照の場合は変更できない。 |     | 
	
	1-9-1. [入力ファイルのPriceDetnExchangeRate]がnullでない場合、[入力ファイルのPriceDetnExchangeRate]を[PriceDetnExchangeRate]にセットする。
	
	1-10. AccountingExchangeRate  

	|Property | Description | EC |
    | ------- | ----------- | ---- |
	|AccountingExchangeRate | 会計為替レート。外貨建て取引の場合、為替レートを入力する。参照の場合は変更できない。 |    |	
	
	1-10-1. [入力ファイルのAccountingExchangeRate]がnullでない場合、[入力ファイルのAccountingExchangeRate]を[AccountingExchangeRate]にセットする。
	
	1-11. TransactionCurrency  

	|Property | Description | EC |
    | ------- | ----------- | ---- |
	| TransactionCurrency | 取引通貨。ビジネスパートナマスタの得意先データまたは仕入先データの通貨が提案される。変更の必要があれば、通貨コードマスタから選択して指定する。参照の場合は変更できない。 | ✔  |
	
	1-11-1. 2-2-1で取得したCurrencyを[TransactionCurrency]にセットする。

2. Orders Header Partner
	次の補助機能の開発を行う。

2-1. ビジネスパートナマスタの取引先機能データの取得

| Property                                    | Description                                                                                                                                                                                                                                                                                                                                                      | EC  | 
| ------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
| PartnerFunction                             | 取引先機能。BUYER:買い手には、ヘッダのBUYERの値が設定される。SELLER:売り手には、ヘッダのSELLERの値が設定される。それ以外の取引先機能として、ビジネスパートナマスタの取引先機能が提案される。必要に応じて変更する。参照の場合は変更できない。なお、オーダーヘッダにおいて、それぞれの取引先機能レコードはオーダーに対して下記の()内の数だけ設定することができる。 <br> 提案される取引先機能: <br> BUYER:買い手(一つ) <br> SELLER:売り手(一つ) <br> CUSTOMER:受注先(一つ) <br> SUPPLIER:仕入先(一つ) <br> MANUFACTURER:製造者(複数) <br> DELIVERFROM:入出荷元(複数) <br> DELIVERTO:入出荷先(複数) <br> LOGI:物流業者(複数) <br> BILLTO:請求先(一つ) <br> BILLFROM:請求元(一つ) <br> PAYEE:支払人(一つ) <br> RECEIVER:受取人(一つ) <br> PSPROVIDER:支払決済サービスプロバイダ(一つ) | ✔        |

2-1-1. BusinessPartner=入力ファイルのbusiness_partner、Customer=入力ファイルのBuyer(1-0の処理結果が”Seller”の場合)、または、Supplier=入力ファイルのSeller(1-0の処理結果が”Buyer”の場合)、をキーとして、対象のBusiness Partner Customer Partner FunctionまたはBusiness Partner Supplier Partner Functionの次のデータのレコード・値を取得し保持する。

| Target / Processing Type                                      | Key(s)                                                                 | 
| ------------------------------------------------------------- | ---------------------------------------------------------------------- | 
| business_partner <br> PartnerCounter <br> PartnerFunction <br> BusinessPartner <br> DefaultPartner <br> / Get and Hold                                                | BusinessPartner=business_partner(入力ファイル) <br> Customer=Buyer(入力ファイル)または、 <br> Supplier=Seller(入力ファイル)                                 | 
| <b>Table Searched</b>                                         | <b>Name of the Table</b>                                               | 
| Business Partner SQL Customer Partner Function または、 <br> Business Partner SQL Supplier Partner Function                | data_platform_business_partner_customer_partner_function_data または、 <br> data_platform_business_partner_supplier_partner_function_data | 
| <b>Field Searched</b>                                         | <b>Data Type / Number of Digits</b>                                    | 
| PartnerCounter <br> PartnerFunction <br> PartnerFunctionBusinessPartner <br> DefaultPartner                                                | int / 3 <br> string(varchar) / 40 <br> int / 12 <br> bool (tinyint) / 1                                            | 
| <b>Single Record or Array</b>                                 | <b>Memo</b>                                                           | 
| Array                                                         |                                                                        | 

[参考①]
以下の例のようなデータが取得される。

| business_partner | Buyer | PartnerCounter | PartnerFunction | PartnerFunctionBusinessPartner | 
| ---------------- | ----- | -------------- | --------------- | ------------------------------ | 
| 201              | 101   | 1              | BUYER           | 101                            | 
| 201              | 101   | 2              | SELLER          | 201                            | 
| 201              | 101   | 3              | CUSTOMER        | 101                            | 
| 201              | 101   | 4              | SUPPLIER        | 201                            | 
| 201              | 101   | 5              | MANUFACTURER    | 202                            | 
| 201              | 101   | 6              | DELIVERFROM     | 203                            | 
| 201              | 101   | 7              | DELIVERTO       | 102                            | 
| 201              | 101   | 8              | LOGI            | 301                            | 
| 201              | 101   | 9              | BILLTO          | 101                            | 
| 201              | 101   | 10             | BILLFROM        | 201                            | 
| 201              | 101   | 11             | PAYEE           | 101                            | 
| 201              | 101   | 12             | RECEIVER        | 201                            | 
| 201              | 101   | 13             | PSPROVIDER      | 401                            | 

2-2. ビジネスパートナの一般データの取得

| Property                | Description                                                                                                                        | EC  | 
| ----------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | --- | 
| BusinessPartnerFullName | ビジネスパートナフルネーム。ビジネスパートナマスタから設定される。変更不可。                                                       |     | 
| BusinessPartnerName     | ビジネスパートナ名称。ビジネスパートナマスタから設定される。変更不可。                                                             |     | 
| Language                | 言語コード。取引先機能に対応するビジネスパートナマスタの言語コードが設定される。必要に応じて変更する。参照の場合は変更できない。   | ✔  | 
| Currency                | 通貨コード。取引先機能に対応するビジネスパートナマスタの通貨コードが設定される。必要に応じて変更する。参照の場合は変更できない。   | ✔  | 
| AddressID               | 住所ID。ビジネスパートナマスタまたは見積の住所IDが設定される。マニュアルで住所を更新する場合、住所IDが新たに設定される。変更不可。 |     | 

2-2-1. 2-1-1で取得した[PartnerFunctionBusinessPartner]をキーとして、対象のビジネスパートナのBusinessPartnerFullName, BusinessPartnerName, Country, Language, Currency, AddressIDの値を取得し保持する。
	
| Target / Processing Type     | Key(s)                                       | 
| ---------------------------- | -------------------------------------------- | 
| 下記Field Searched <br> / Get and Hold               | PartnerFunctionBusinessPartner(2-1で取得)    | 
| <b>Table Searched</b>        | <b>Name of the Table</b>                     | 
| Business Partner SQL General | data_platform_business_partner _general_data | 
| <b>Field Searched</b>        | <b>Data Type / Number of Digits</b>          | 
| BusinessPartnerFullName <br> BusinessPartnerName <br> Country <br> Language <br> Currency <br> AddressID                    | string(varchar) / 200 <br> string(varchar) / 100 <br> string(varchar) / 3 <br> string(varchar) / 2 <br> string(varchar) / 5 <br> int / 12                     | 
| <b>Single Record or Array</b>| <b>Memo</b>                                 | 
| Array                        |                                              | 

[参考②]
2-1で取得されたデータと、2-2で取得されたデータを並べると、以下の例のようになる。

| business_partner                                        | Buyer | PartnerCounter | PartnerFunction | PartnerFunctionBusinessPartner | BusinessPartnerName (※FullName, Language, Currency, AddressIDの例示は省略)  | 
| ------------------------------------------------------- | ----- | -------------- | --------------- | ------------------------------ | -------------------- | 
| 201                                                     | 101   | 1              | BUYER           | 101                            | ローソン本社         | 
| 201                                                     | 101   | 2              | SELLER          | 201                            | 山崎パン販売         | 
| 201                                                     | 101   | 3              | CUSTOMER        | 101                            | ローソン本社         | 
| 201                                                     | 101   | 4              | SUPPLIER        | 201                            | 山崎パン販売         | 
| 201                                                     | 101   | 5              | MANUFACTURER    | 202                            | 山崎パン             | 
|                                                         |       |                |                 |                                |                      | 
| 201                                                     | 101   | 6              | DELIVERFROM     | 203                            | 山崎パン松戸第２工場 | 
| 201                                                     | 101   | 7              | DELIVERTO       | 102                            | ローソン虎ノ門店     | 
| 201                                                     | 101   | 8              | LOGI            | 301                            | 日立物流             | 
| 201                                                     | 101   | 9              | BILLTO          | 101                            | ローソン本社         | 
| 201                                                     | 101   | 10             | BILLFROM        | 201                            | 山崎パン販売         | 
| 201                                                     | 101   | 11             | PAYEE           | 101                            | ローソン本社         | 
| 201                                                     | 101   | 12             | RECEIVER        | 201                            | 山崎パン販売         | 
| 201                                                     | 101   | 13             | PSPROVIDER      | 401                            | GMOあおぞら銀行      | 

2-1で取得されたServiceLabelが”QUOTATIONS”の場合はOverallDocReferenceStatus に“QT”、”INQUIRIES”の場合はOverallDocReferenceStatusに”IN”をセットする。それ以外の場合は、何もセットしない。  

<ロジックまとめ>　　

| ServiceLabel   | OverallDocReferenceStatusにセットする値 | 
| -------------- | --------------------------------------- | 
| “QUOTATIONS” | “QT”                                  | 
| “INQUIRIES”  | “IN”                                  | 

2-1-1. 参照登録(101)が発動した場合に、変更が不可となる制御をかける。つまり、inputファイルに値があっても、その値は採用されない。

3. Orders Header Partner Contact
次の補助機能の開発を行う。

3-1. ビジネスパートナマスタの取引先機能コンタクトデータの取得  

| Property          | Description                                                                                                                  | EC  | 
| ----------------- | ---------------------------------------------------------------------------------------------------------------------------- | --- | 
| ContactID         | コンタクトID。ビジネスパートナ取引先機能コンタクトデータから設定される。必要に応じて変更する。参照の場合は変更できない。     |     | 
| ContactPersonName | コンタクト担当者名。ビジネスパートナ取引先機能コンタクトから設定される。必要に応じて変更する。参照の場合は変更できない。     |     | 
| EmailAddress      | Eメールアドレス。ビジネスパートナ取引先機能コンタクトデータから設定される。必要に応じて変更する。参照の場合は変更できない。  |     | 
| PhoneNumber       | 電話番号。ビジネスパートナ取引先機能コンタクトデータから設定される。必要に応じて変更する。参照の場合は変更できない。         |     | 
| MobilePhoneNumber | モバイル電話番号。ビジネスパートナ取引先機能コンタクトデータから設定される。必要に応じて変更する。参照の場合は変更できない。 |     | 
| FaxNumber         | ファクス番号。ビジネスパートナ取引先機能コンタクトデータから設定される。必要に応じて変更する。参照の場合は変更できない。     |     | 
| ContactTag1       | コンタクトタグ1。ビジネスパートナ取引先機能コンタクトデータから設定される。必要に応じて変更する。参照の場合は変更できない。  |     | 
| ContactTag2       | コンタクトタグ2。ビジネスパートナ取引先機能コンタクトデータから設定される。必要に応じて変更する。参照の場合は変更できない。  |     | 
| ContactTag3       | コンタクトタグ3。ビジネスパートナ取引先機能コンタクトデータから設定される。必要に応じて変更する。参照の場合は変更できない。  |     | 
| ContactTag4       | コンタクトタグ4。ビジネスパートナ取引先機能コンタクトデータから設定される。必要に応じて変更する。参照の場合は変更できない。  |     | 

3-1-1. [BusinessPartner=入力ファイルのbusiness_partner、Customer=入力ファイルのBuyer(1-0の処理結果が”Seller”の場合)、または、Supplier=入力ファイルのSeller(1-0の処理結果が”Buyer”の場合)、2-1-1で取得したPartnerCounter]、をキーとして、対象のBusiness Partner Customer Partner Function Contact または Business Partner Supplier Partner Function Contactの次のデータのレコード・値を取得し保持する。

| Target / Processing Type                                               | Key(s)                                                                        | 
| ---------------------------------------------------------------------- | ----------------------------------------------------------------------------- | 
| 下記Field Searched <br> / Get and Hold                                                         | BusinessPartner=business_partner(入力ファイル) <br> Customer=Buyer(入力ファイル)または <br> Supplier=Seller(入力ファイル) <br> PartnerCounter(2-1-1で取得)                                            | 
| <b>Table Searched</b>                                                  | <b>Name of the Table</b>                                                      | 
| Business Partner SQL Customer Partner Function Contact Dataまたは <br>  Business Partner SQL Supplier Partner Function Contact Data            | data_platform_business_partner_customer_partner_function_contact_dataまたは <br>  data_platform_business_partner_supplier_partner_ function_contact_data | 
| <b>Field Searched</b>                                                  | <b>Data Type / Number of Digits</b>                                           | 
| ContactID <br> ContactPersonName <br> EmailAddress <br> PhoneNumber <br>  MobilePhoneNumber <br> FaxNumber <br> ContactTag1 <br> ContactTag2 <br> ContactTag3 <br> ContactTag4 <br> DefaultContact                                                         | int / 4 <br> string(varchar) / 100 <br> string(varchar) / 200 <br> string(varchar) / 100 <br> string(varchar) / 100 <br> string(varchar) / 100 <br> string(varchar) / 40 <br> string(varchar) / 40 <br> string(varchar) / 40 <br> string(varchar) / 40 <br> bool(tinyint) / 1                                                      | 
| <b>Single Record or Array</b>                                          | <b>Memo</b>                                                                  | 
| Array                                                                  |                                                                               | 

4. Orders Header Partner Plant
次の補助機能の開発を行う。

4-1. ビジネスパートナマスタの取引先プラントデータの取得  

| Property             | Description                                                                  | EC  | 
| -------------------- | ---------------------------------------------------------------------------- | --- | 
| PartnerFunction      | 取引先機能。ヘッダパートナの取引先機能のうち次のものが設定される。変更不可。 <br> MANUFACTURER:製造者 <br> DELIVERFROM:入出荷元 <br>  DELIVERTO:入出荷先   |                                                                              | 


4-1-1. [BusinessPartner=入力ファイルのbusiness_partner、Customer=入力ファイルのBuyer(1-0の処理結果が”Seller”の場合)、または、Supplier=入力ファイルのSeller(1-0の処理結果が”Buyer”の場合)、2-1-1で取得したPartnerCounter、2-1-1で取得したPartnerFunction、2-1-1で取得したPartnerFunctionBusinessPartner]、をキーとして、対象のBusiness Partner Customer Partner Plant(1-0の処理結果が”Seller”の場合)、または、 Business Partner Supplier Partner Plant(1-0の処理結果が”Buyer”の場合)、の次のデータのレコード・値を取得し保持する。

| Target / Processing Type                                   | Key(s)                                                             | 
| ---------------------------------------------------------- | ------------------------------------------------------------------ | 
| 下記Field Searched <br>/ Get and Hold                                             | BusinessPartner=business_partner(入力ファイル) <br> Customer=Buyer(入力ファイル)または、 <br> Supplier=Seller(入力ファイル)  <br> PartnerCounter(2-1-1で取得) <br>  PartnerFunction(2-1-1で取得) <br> PartnerFunctionBusinessPartner(2-1-1で取得)                | 
| <b>Table Searched</b>                                      | <b>Name of the Table</b>                                           | 
| Business Partner SQL Customer Partner Plant Dataまたは、 <br> Business Partner SQL Supplier Partner Plant Data           | data_platform_business_partner_customer_partner_plant_dataまたは、 <br> data_platform_business_partner_supplier_partner_plant_data | 
| <b>Field Searched</b>                                      | <b>Data Type / Number of Digits</b>                                | 
| PlantCounter <br> Plant <br> DefaultPlant                                               | int / 3 <br> string(varchar) / 4 <br> bool (tinyint) / 1                                         | 
| <b>Single Record or Array</b>                              | <b>Memo</b>                                                       | 
| Array                                                      |                                                                    | 


5. Orders Item
次の補助機能の開発を行う。  

5-0. OrderItem  

| Property  | Description                                                             | EC  | 
| --------- | ----------------------------------------------------------------------- | --- | 
| OrderItem | オーダー明細番号。1～999999の範囲で入力する。参照の場合は変更できない。 |     |


5-0-1. 入力ファイルのItemの明細の数だけ、1～Nまで整数の番号を付与して配列に保持する。

5-1. BPTaxClassification  

| Property            | Description                                                                                                                    | EC  | 
| ------------------- | ------------------------------------------------------------------------------------------------------------------------------ | --- | 
| BPTaxClassification | ビジネスパートナ税分類。ビジネスパートナマスタの税分類が以下の値が提案される。必要に応じて変更する。参照の場合は変更できない。 <br> 0:非課税 <br> 1:納税義務          |                                                                                                                                | 


5-1-1. BusinessPartner=入力ファイルのbusiness_partner, Customer=入力ファイルのBuyer(1-0の処理結果が”Seller”の場合)、または、Supplier=入力ファイルのSeller(1-0の処理結果が”Buyer”の場合), DepartureCountry=2.2.1で取得したPartnerFunction=”BUYER”のCountry(1-0の処理結果が”Seller”の場合)、または、PartnerFunction=”SELLER”のCountry(1-0の処理結果が”Buyer”の場合)、をキーとして、対象のビジネスパートナのBPTaxClassificationを検索する。

| Target / Processing Type                                           | Key(s)                                                   | 
| ------------------------------------------------------------------ | -------------------------------------------------------- | 
| BPTaxClassification <br>  / Get and Hold                                                     | BusinessPartner=入力ファイルの”business_partner” <br> Customer=入力ファイルのBuyerまたは、 <br> Supplier=Seller(入力ファイル) <br> DepartureCountry=2.2.1で取得したPartnerFunction=”BUYER”のCountry | 
| <b>Table Searched</b>                                              | <b>Name of the Table</b>                                 | 
| Business Partner SQL Customer Tax Dataまたは、 <br> Business Partner SQL Supplier Tax Data                             | data_platform_business_partner_customer_tax_dataまたは、 <br> data_platform_business_partner_supplier_tax_data                   | 
| <b>Field Searched</b>                                              | <b>Data Type / Number of Digits</b>                      | 
| BPTaxClassification                                                | string(varchar) / 1                                      | 
| <b>Single Record or Array</b>                                      | <b>Memo</b>                                             | 
| Single Record                                                      |                                                          | 


5-1-2. 入力ファイルのProduct, BusinessPartner=入力ファイルのbusiness_partner, Country=2.2.1で取得したPartnerFunction=”BUYER”のCountry(1-0の処理結果が”Seller”の場合)、または、PartnerFunction=”SELLER”のCountry(1-0の処理結果が”Buyer”の場合), TaxCategory=”MWST”, をキーとして、対象の品目のProductTaxClassificationを検索する。

| Target / Processing Type                                                  | Key(s)                                | 
| ------------------------------------------------------------------------- | ------------------------------------- | 
| ProductTaxClassification <br> / Get and Hold                                                            | Product=入力ファイルのProduct <br> BusinessPartner=入力ファイルの”business_partner” <br>  Country=2.2.1で取得したPartnerFunction=”BUYER”または”SELLER”のCountry <br> TaxCategory=”MWST”                                                      | 
| <b>Table Searched</b>                                                     | <b>Name of the Table</b>              | 
| Product Master SQL Tax Data                                               | data_platform_product_master_tax_data | 
| <b>Field Searched</b>                                                     | <b>Data Type / Number of Digits</b>   | 
| ProductTaxClassification                                                  | string(varchar) / 1                   | 
| <b>Single Record or Array</b>                                             | <b>Memo</b>                          | 
| Single Record                                                             |                                       | 


5-1-3. 5-1-1で保持したCustomerのBPTaxClassification(1-0の処理結果が”Seller”の場合)、または、SupplierのBPTaxClassification(1-0の処理結果が”Buyer”の場合)、と5-1-2で保持したProductのProductTaxClassification、をもとに、税コード(TaxCode)をセットする。セットのロジックとしては、以下の表を参照する。  
＜税コードセットのロジック＞  

| ケース | Customer Tax Classification | Product Tax Classification | Order Type | Tax Code |
| ------ | --------------------------- | -------------------------- | ---------- | -------- |
| ケース1 |          1                 |            1               |   販売系    |    A1   |
| ケース2 |          0                 |            0               |   販売系    |    A0   |
| ケース3 |          0                 |            1               |   販売系    |    A0   |
| ケース4 |          1                 |            0               |   販売系    |    A0   |
| ケース5 |          1                 |            1               |   購買系    |    V1   |
| ケース6 |          0                 |            0               |   購買系    |    V0   |
| ケース7 |          0                 |            1               |   購買系    |    V0   |
| ケース8 |          1                 |            0               |   購買系    |    V0   |
 

上記の8のケースがあり、Business Partner Customer TaxのBPTaxClassification(1-0の処理結果が”Seller”の場合)、または、Business Partner Supplier TaxのBPTaxClassification(1-0の処理結果が”Buyer”の場合)、および、Product Taxの ProductTaxClassification、において、1は課税、０は非課税、を表す。どちらも”１”の場合のみ消費税が課税となり、その場合、税コードは販売系/購買系でそれぞれA1/V1とセットされる。それ以外の場合は消費税が非課税となり、販売系/購買系でそれぞれA0/V0とセットされる。

5-2. 品目マスタ一般データの取得  

| Property                      | Description                                                                                                      | EC  | 
| ----------------------------- | ---------------------------------------------------------------------------------------------------------------- | --- | 
| ProductStandardID             | 品目標準ID。品目マスタから設定される。変更不可。                                                                 |     | 
| ProductGroup                  | 品目グループ。品目マスタから設定される。変更不可。                                                               |     | 
| BaseUnit                      | 基本数量単位。品目マスタから設定される。変更不可。                                                               |     | 
| ItemWeightUnit                | 重量単位。品目マスタの重量単位が提案される。必要に応じて変更する。参照の場合は変更できない。                     | ✔  | 
| ProductGrossWeight            | 品目総重量。品目マスタの総重量が設定される。必要に応じて任意の総重量を入力する。参照の場合は変更できない。       |     | 
| ProductItemWeight             | 品目正味重量。品目マスタの正味重量が設定される。必要に応じて任意の正味重量を入力する。参照の場合は変更できない。 |     | 
| ProductAccountAssignmentGroup | 品目勘定設定グループ。品目マスタの勘定設定グループがコピーされる。変更不可。                                     |     | 
| CountryOfOrigin               | 原産国。品目マスタから参照して設定される。変更不可。                                                             |     | 


5-2-1. 入力ファイルの[Product]をキーとして、対象の品目マスタ一般データの[ProductStandardID, ProductGroup, BaseUnit, WeightUnit, NetWeight, ProductAccountAssignmentGroup, CountryOfOrigin]、を検索して値をセットする。

| Target / Processing Type         | Key(s)                                    | 
| -------------------------------- | ----------------------------------------- | 
| ProductStandardID <br> ProductGroup <br> BaseUnit <br> WeightUnit <br> GrossWeight <br> NetWeight <br> ProductAccountAssignmentGroup <br> CountryOfOrigin                  | [Product] (入力ファイル)                   | 
| <b>Table Searched</b>            | <b>Name of the Table</b>                  | 
| Product Master SQL General　Data | data_platform_product_master_general_data | 
| <b>Field Searched</b>            | <b>Data Type / Number of Digits</b>       | 
| ProductStandardID <br> ProductGroup <br> BaseUnit <br> ItemWeightUnit <br> ProductGrossWeight <br> ProductNetWeight <br> ProductAccountAssignmentGroup <br> CountryOfOrigin                  | string(varchar) / 18 <br> string(varchar) / 9 <br> string(varchar) / 3 <br> string(varchar) / 3 <br>  float / 15 <br> float / 15 <br> string(varchar) / 2 <br>  string(varchar) / 3              | 
| <b>Single Record or Array</b>    | <b>Memo</b>                              | 
| Array                            |                                           | 


5-3. OrderItemText  

| Property      | Description                                                                                                                                        | EC  | 
| ------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
| OrderItemText | オーダー明細テキスト。品目マスタテキストデータから初期値が提案される。必要に応じて任意のオーダー明細テキストを入力する。参照の場合は変更できない。 |     | 

5-3-1. [Product=入力ファイルのProduct, Language=2.2.1で取得したPartnerFunction
=”1-0でセットした値”のLanguage]をキーとして、対象の品目マスタ品目説明データの[ProductDescription]を検索して値をセットする。  

| Target / Processing Type                                                 | Key(s)                                                 | 
| ------------------------------------------------------------------------ | ------------------------------------------------------ | 
| OrderItemText <br>  / Get and Set                                                            | [Product(入力ファイル) <br> Language=2.2.1で取得したPartnerFunction=”1-0でセットした値”のLanguage] | 
| <b>Table Searched</b>                                                    | <b>Name of the Table</b>                               | 
| Product Master SQL Product Description Data                              | data_platform_product_master_ product_description_data | 
| <b>Field Searched</b>                                                    | <b>Data Type / Number of Digits</b>                    | 
| ProductDescription                                                       | string(varchar) / 200                                  | 
| <b>Single Record or Array</b>                                            | <b>Memo</b>                                           | 
| Array                                                                    | 


5-4. StockConfirmationPartnerFunction / StockConfirmationBusinessPartner / StockConfirmationPlant
入力ファイルの[OrderItemCategory=”INVP”の明細]に対して、本処理を実行する。  

| Property                         | Description                                                                                                                                                  | EC  | 
| -------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | --- | 
| StockConfirmationPartnerFunction | 在庫確認を行う取引先機能。ヘッダパートナプラントから設定される。必要に応じて変更する。（PRODUCT_STOCK_AVAIRABILITY_CHECK_SRVを利用、機能はオプション） 参照の場合は変更できない。       | ✔    | 
| StockConfirmationBusinessPartner | 在庫確認を行うビジネスパートナ。ヘッダパートナプラントから設定される。必要に応じて変更する。（PRODUCT_STOCK_AVAIRABILITY_CHECK_SRVを利用、機能はオプション 参照の場合は変更できない。       | ✔  | 
| StockConfirmationPlant           | 在庫確認プラント。ヘッダパートナプラントから設定される。必要に応じて変更する。参照の場合は変更できない。                                                     | ✔  | 

5-4-1. 4-1-1で取得した[ヘッダパートナプラントのデータ]から、DefaultStockConfrimationPlant=trueをキーとして、[PartnerFunction, BusinessPartner, Plant]の値を検索してセットする。  

<4-1-1で取得した[ヘッダパートナプラントのデータ]に対しての処理>  

| Target / Processing Type         | Key(s)                                                                                                       | 
| -------------------------------- | ------------------------------------------------------------------------------------------------------------ | 
| StockConfirmationPartnerFunction <br> StockConfirmationBusinessPartner <br>  StockConfirmationPlant <br> / Get and Set                    | DefaultStockConfriamtionPlant=true                                                                           | 
| <b>Table Searched</b>            | <b>Name of the Table</b>                                                                                     | 
| [Header Partner Plant] <br> (4-1-1で取得したデータ)          | None                                                                                                         | 
| <b>Field Searched</b>            | <b>Data Type / Number of Digits</b>                                                                          | 
| PartnerFunction <br> BusinessPartner <br> Plant                            | string(varchar) / 40 <br> int / 12 <br> string(varchar) / 4              | 
| <b>Single Record or Array</b>    | <b>Memo</b>                                                                                                 | 
| Single Record                    | 前提として、ヘッダパートナプラントにDefaultStockConfriamtionPlant=trueのレコードは必ず一つ以下しか存在しない | 


5-5. OrderIssuingUnit / OrderReceivingUnit / IssuingStorageLocation / ReceivingStorageLocation
入力ファイルの[OrderItemCategory=”INVP”の明細]に対して、本処理を実行する。  

| Property           | Description                                                                                                                                      | EC  | 
| ------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | --- | 
| OrderIssuingUnit   | オーダー出荷数量単位。品目マスタビジネスパートナプラントデータの出荷数量単位または入荷数量単位が提案される。変更不可。参照の場合は変更できない。 |     | 
| OrderReceivingUnit | オーダー入荷数量単位。品目マスタビジネスパートナプラントデータの出荷数量単位または入荷数量単位が提案される。変更不可。参照の場合は変更できない。 |     | 

5-5-1. [Product=入力ファイルのProduct、BusinessPartner=5-4-1で値セットされたStockConfirmationBusinessPartner、Plant=5-4-1で値セットされたStockConfrimationPlant]、をキーとして、対象の品目マスタBPプラントデータの[IssuingDeliveryUnit, ReceivingDeliveryUnit, IsBatchManagementRequired]
を検索して[OrderIssuingUnit, OrderReceivingUnit, ProductIsBatchManagedInStockConfirmationPlant
]として値をセットする。  

| Target / Processing Type                                              | Key(s)                                     | 
| --------------------------------------------------------------------- | ------------------------------------------ | 
| OrderIssuingUnit <br> OrderReceivingUnit <br>  ProductIsBatchManagedInStockConfirmationPlant <br> / Get and Set                                                         | [Product(入力ファイル) <br> BusinessPartner=5-4-1で値セットされたStockConfrimationBusinessPartner <br> Plant=5-4-1で値セットされたStockConfrimationPlant]                    | 
| <b>Table Searched</b>                                                 | <b>Name of the Table</b>                   | 
| Product Master SQL BP Plant　Data                                     | data_platform_product_master_bp_plant_data | 
| <b>Field Searched</b>                                                 | <b>Data Type / Number of Digits</b>        | 
| IssuingDeliveryUnit <br> ReceivingDeliveryUnit <br> IsBatchManagementRequired                                             | string(varchar) / 3 <br> string(varchar) / 3 <br> bool(tinyint) / 1                                                     | 
| <b>Single Record or Array</b>                                         | <b>Memo</b>                               | 
| Single Record                                                         |                                            | 


5-6. ProductionPlantPartnerFunction / ProductionPlantBusinessPartner / ProductionPlant
入力ファイルの[OrderItemCategory=”INVP”の明細]に対して、本処理を実行する。  

| Property                       | Description                                                                                                                                                                                    | EC  | 
| ------------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
| ProductionPlantPartnerFunction | 製造プラント取引先機能。ヘッダパートナプラントの該当取引先機能から提案される。必要に応じてヘッダから取引先機能・ビジネスパートナ・プラントの組合せを入力する。参照の場合は変更できない。       | ✔  | 
| ProductionPlantBusinessPartner | 製造プラントビジネスパートナ。ヘッダパートナプラントの該当取引先機能から提案される。必要に応じてヘッダから取引先機能・ビジネスパートナ・プラントの組合せを入力する。参照の場合は変更できない。 | ✔  | 
| ProductionPlant                | 製造プラント。ヘッダパートナプラントの該当取引先機能から提案される。必要に応じてヘッダから取引先機能・ビジネスパートナ・プラントの組合せを入力する。参照の場合は変更できない。                 | ✔  | 


5-6-1. 4-1-1で取得した[ヘッダパートナプラントのデータ]から、
[PartnerFunction=”MANUFACTURER”, DefaultPlant=true]をキーとして、[BusinessPartner, Plant]の値を検索してセットする。  

| Target / Processing Type       | Key(s)                            | 
| ------------------------------ | --------------------------------- | 
| ProductionPlantPartnerFunction <br> ProductionPlantBusinessPartner <br> ProductionPlant <br> / Get and Set                  | [PartnerFunction=”MANUFACTURER” <br> DefaultPlant=true]             | 
| <b>Table Searched</b>          | <b>Name of the Table</b>          | 
| None                           | None                              | 
| <b>Field Searched</b>          | <b>Data Type / Number of Digits</b> 
| BusinessPartner <br> Plant                          | int / 12 <br> string(varchar) / 4            | 
| <b>Single Record or Array</b>  | <b>Memo</b>                      | 
| Array                          |                                   | 



5-7. IssuingPlantPartnerFunction / IssuingPlantBusinessPartner / IssuingPlant
入力ファイルの[OrderItemCategory=”INVP”の明細]に対して、本処理を実行する。  

| Property                    | Description                                                                                                                                                                                    | EC  | 
| --------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
| IssuingPlantPartnerFunction | 出荷プラント取引先機能。ヘッダパートナプラントの該当取引先機能から提案される。必要に応じてヘッダから取引先機能・ビジネスパートナ・プラントの組合せを入力する。参照の場合は変更できない。       | ✔  | 
| IssuingPlantBusinessPartner | 出荷プラントビジネスパートナ。ヘッダパートナプラントの該当取引先機能から提案される。必要に応じてヘッダから取引先機能・ビジネスパートナ・プラントの組合せを入力する。参照の場合は変更できない。 | ✔  | 
| IssuingPlant                | 出荷プラント。ヘッダパートナプラントの該当取引先機能から提案される。必要に応じてヘッダから取引先機能・ビジネスパートナ・プラントの組合せを入力する。参照の場合は変更できない。                 | ✔  | 


5-7-1. 4-1-1で取得した[ヘッダパートナプラントのデータ]から、
[PartnerFunction=”DELIVERFROM”, DefaultPlant=true]をキーとして、[BusinessPartner, Plant]の値を検索してセットする。  

| Target / Processing Type    | Key(s)                           | 
| --------------------------- | -------------------------------- | 
| IssuingPlantPartnerFunction <br> IssuingPlantBusinessPartner <br> IssuingPlant <br> / Get and Set               | [PartnerFunction=”DELIVERFROM” <br> DefaultPlant=true]          | 
| <b>Table Searched</b>       | <b>Name of the Table</b>         | 
| None                        | None                             | 
| <b>Field Searched</b>       | <b>Data Type / Number of Digits</b>
| BusinessPartner <br> Plant                       | int / 12 <br> string(varchar) / 4         | 
| <b>Single Record or Array</b> | <b>Memo</b>                     | 
| Array                       |                                  | 


5-8. ReceivingPlantPartnerFunction / ReceivingPlantBusinessPartner / ReceivingPlant
入力ファイルの[OrderItemCategory=”INVP”の明細]に対して、本処理を実行する。  

| Property                      | Description                                                                                                                                                                                    | EC  | 
| ----------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
| ReceivingPlantPartnerFunction | 入荷プラント取引先機能。ヘッダパートナプラントの該当取引先機能から提案される。必要に応じてヘッダから取引先機能・ビジネスパートナ・プラントの組合せを入力する。参照の場合は変更できない。       | ✔  | 
| ReceivingPlantBusinessPartner | 入荷プラントビジネスパートナ。ヘッダパートナプラントの該当取引先機能から提案される。必要に応じてヘッダから取引先機能・ビジネスパートナ・プラントの組合せを入力する。参照の場合は変更できない。 | ✔  | 
| ReceivingPlant                | 入荷プラント。ヘッダパートナプラントの該当取引先機能から提案される。必要に応じてヘッダから取引先機能・ビジネスパートナ・プラントの組合せを入力する。参照の場合は変更できない。                 | ✔  | 


5-8-1. 4-1-1で取得した[ヘッダパートナプラントのデータ]から、
[PartnerFunction=”DELIVERTO”, DefaultPlant=true]をキーとして、[BusinessPartner, Plant]の値を検索してセットする。  

| Target / Processing Type      | Key(s)                         | 
| ----------------------------- | ------------------------------ | 
| ReceivingPlantPartnerFunction <br> ReceivingPlantBusinessPartner <br> ReceivingPlant <br> / Get and Set                 | [PartnerFunction=”DELIVERTO” <br> DefaultPlant=true]            | 
| <b>Table Searched</b>         | <b>Name of the Table</b>       | 
| None                          | None                           | 
| <b>Field Searched</b>         | <b>Data Type / Number of Digits</b> |
|BusinessPartner <br> Plant                         | int / 12 <br> string(varchar) / 4           | 
| <b>Single Record or Array</b> | <b>Memo</b>                   | 
| Array                         |                                | 

5-9. ProductionPlantTimeZone
入力ファイルの[OrderItemCategory=”INVP”の明細]に対して、本処理を実行する。  

| Property                | Description                                                          | EC  | 
| ----------------------- | -------------------------------------------------------------------- | --- | 
| ProductionPlantTimeZone | 製造プラントのタイムゾーン。プラントマスタから設定される。変更不可。 |     | 

5-9-1. [5-6-1で取得した製造プラントのデータ]から、
BusinessPartner, Plantをキーとして、対象のプラントマスタ一般データの[TimeZone]を検索して値をセットする。  

| Target / Processing Type | Key(s)                            | 
| ------------------------ | --------------------------------- | 
| ProductionPlantTimeZone <br> / Get and Set            | BusinessPartner <br> Plant                    | 
| <b>Table Searched</b>    | <b>Name of the Table</b>          | 
| Plant SQL General Data   | data_platform_plant_general_data  |
| <b>Field Searched</b>    | <b>Data Type / Number of Digits</b> 
| TimeZone                 | string(varchar) / 3               | 
| <b>Single Record or Array</b> | <b>Memo</b>                              | 
| Array                    |                                   | 


5-10. IssuingPlantTimeZone
入力ファイルの[OrderItemCategory=”INVP”の明細]に対して、本処理を実行する。  

| Property             | Description                                                          | EC  | 
| -------------------- | -------------------------------------------------------------------- | --- | 
| IssuingPlantTimeZone | 出荷プラントのタイムゾーン。プラントマスタから設定される。変更不可。 |     | 

5-10-1. [5-7-1で取得した出荷プラントのデータ]から、
[BusinessPartner, Plant]をキーとして、対象のプラントマスタ一般データの[TimeZone]を検索して値をセットする。  

| Target / Processing Type | Key(s)                            | 
| ------------------------ | --------------------------------- | 
| IssuingPlantTimeZone <br> / Get and Set            | BusinessPartner <br>  Plant                    | 
| <b>Table Searched</b>    | <b>Name of the Table</b>          | 
| Plant SQL General Data   | data_platform_plant_general_data | 
| <b>Field Searched</b>    | <b>Data Type / Number of Digits</b> 
| TimeZone                 | string(varchar) / 3               | 
| <b>Single Record or Array</b> | <b>Memo</b>                              | 
| Array                    |                                   | 

	

5-11. ReceivingPlantTimeZone
入力ファイルの[OrderItemCategory=”INVP”の明細]に対して、本処理を実行する。  

| Property               | Description                                                          | EC  | 
| ---------------------- | -------------------------------------------------------------------- | --- | 
| ReceivingPlantTimeZone | 入荷プラントのタイムゾーン。プラントマスタから設定される。変更不可。 |     | 

5-11-1. [5-8-1で取得した入荷プラントのデータ]から、
[BusinessPartner, Plant]をキーとして、対象のプラントマスタ一般データの[TimeZone]を検索して値をセットする。  

| Target / Processing Type | Key(s)                            | 
| ------------------------ | --------------------------------- | 
| ReceivingPlantTimeZone <br> / Get and Set            | BusinessPartner <br> Plant                    | 
| <b>Table Searched</b>    | <b>Name of the Table</b>          | 
| Plant SQL General Data   | data_platform_plant_general_data | 
| <b>Field Searched</b>    | <b>Data Type / Number of Digits</b> 
| TimeZone                 | string(varchar) / 3               | 
| <b>Single Record or Array</b> | <b>Memo</b>                              | 
| Single Record            |                                   | 



5-12. ProductionPlantStorageLocation  
入力ファイルの[OrderItemCategory=”INVP”の明細]に対して、本処理を実行する。    

| Property                       | Description                                                                                | EC  | 
| ------------------------------ | ------------------------------------------------------------------------------------------ | --- | 
| ProductionPlantStorageLocation | 製造プラント在庫保管場所。保管場所をマスタから指定して入力する。参照の場合は変更できない。 | ✔  | 

5-12-1. [入力ファイルのProduct, 5-6-1でセットされたProductionPlant]をキーとして、
対象の品目マスタ作業計画データの[ProductionInvtryManagedLoc]を検索して値をセットする。  

| Target / Processing Type                     | Key(s)                                            | 
| -------------------------------------------- | ------------------------------------------------- | 
| ProductionPlantStorageLocation / Get and Set | [Product(入力ファイル) <br>  ProductionPlant(4-6-1でセット)]              | 
| <b>Table Searched</b>                        | <b>Name of the Table</b>                          | 
| Product Master SQL Work Scheduling Data      | data_platform_product_master_work_scheduling_data | 
| <b>Field Searched</b>                        | <b>Data Type / Number of Digits</b>               | 
| ProductionInvtryManagedLoc                   | string(varchar) / 4                               | 
| <b>Single Record or Array</b>                | <b>Memo</b>                                      | 
| Array                                        |                                                   | 


5-13. IssuingPlantStorageLocation
入力ファイルの[OrderItemCategory=”INVP”の明細]に対して、本処理を実行する。  

| Property                    | Description                                                                                                                                                  | EC  | 
| --------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | --- | 
| IssuingPlantStorageLocation | 出荷プラントの在庫保管場所。品目マスタまたはプラントマスタから初期値が設定される。必要に応じて保管場所マスタから選択して入力する。参照の場合は変更できない。 | ✔  | 

5-13-1. [入力ファイルのProduct, 5-7-1でセットされたIssuingPlant]をキーとして、
対象の品目マスタBPプラントデータの[IssuingStorageLocation]を検索して値をセットする。  

| Target / Processing Type          | Key(s)                                     | 
| --------------------------------- | ------------------------------------------ | 
| IssuingPlantStorageLocation <br> / Get and Set                     | [Product(入力ファイル) <br> IssuingPlant(5-7-1でセット)]      | 
| <b>Table Searched</b>             | <b>Name of the Table</b>                   | 
| Product Master SQL BP Plant　Data | data_platform_product_master_bp_plant_data | 
| <b>Field Searched</b>             | <b>Data Type / Number of Digits</b>        | 
| IssuingStorageLocation            | string(varchar) / 4                        | 
| <b>Single Record or Array</b>     | <b>Memo</b>                               | 
| Array                             |                                            | 


5-14. ReceivingPlantStorageLocation
入力ファイルの[OrderItemCategory=”INVP”の明細]に対して、本処理を実行する。  

| Property                      | Description                                                                                                                                                  | EC  | 
| ----------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | --- | 
| ReceivingPlantStorageLocation | 入荷プラントの在庫保管場所。品目マスタまたはプラントマスタから初期値が設定される。必要に応じて保管場所マスタから選択して入力する。参照の場合は変更できない。 | ✔  | 

5-14-1. [入力ファイルのProduct, 5-8-1でセットされたReceivingPlant]をキーとして、
対象の品目マスタBPプラントデータの[ReceivingStorageLocation]を検索して値をセットする。  

| Target / Processing Type          | Key(s)                                     | 
| --------------------------------- | ------------------------------------------ | 
| ReceivingPlantStorageLocation <br> / Get and Set                     | [Product(入力ファイル) <br> ReceivingPlant(5-7-1でセット)]    | 
| <b>Table Searched</b>             | <b>Name of the Table</b>                   | 
| Product Master SQL BP Plant　Data | data_platform_product_master_bp_plant_data | 
| <b>Field Searched</b>             | <b>Data Type / Number of Digits</b>        | 
| ReceivingStorageLocation          | string(varchar) / 4                        | 
| <b>Single Record or Array</b>     | <b>Memo</b>                               | 
| Array                             |                                            | 


5-15. Incoterms  

| Property  | Description                                                                                                                              | EC  | 
| --------- | ---------------------------------------------------------------------------------------------------------------------------------------- | --- | 
| Incoterms | インコタームズ。ヘッダのインコタームズがコピーされる。必要に応じて、インコタームズマスタから選択して指定する。参照の場合は変更できない。 | ✔  | 

5-15-1. HeaderのIncotermsから、[明細の配列保持データ]に対して[Incoterms]をコピーする。  

5-16. PaymentTerms  

| Property     | Description                                                                                                                | EC  | 
| ------------ | -------------------------------------------------------------------------------------------------------------------------- | --- | 
| PaymentTerms | 支払条件。ヘッダの支払条件が提案される。変更の必要があれば、支払条件マスタから選択して指定する。参照の場合は変更できない。 | ✔  | 

5-16-1. HeaderのPaymentTermsから、[明細の配列保持データ]に対して[PaymentTermsをコピー]する。  

5-17. PaymentMethod  

| Property      | Description                                                                                                                | EC  | 
| ------------- | -------------------------------------------------------------------------------------------------------------------------- | --- | 
| PaymentMethod | 支払方法。ヘッダの支払方法が提案される。変更の必要があれば、支払方法マスタから選択して指定する。参照の場合は変更できない。 | ✔  | 

5-17-1. HeaderのPaymentMethodから、[明細の配列保持データ]に対して[PaymentMethodをコピー]する。  

5-18. ItemGrossWeight  

| Property        | Description                                                                                                                      | EC  | 
| --------------- | -------------------------------------------------------------------------------------------------------------------------------- | --- | 
| ItemGrossWeight | 合計総重量。総重量とオーダー数量により合計総重量が計算される。必要に応じて任意の合計総重量を入力する。参照の場合は変更できない。 | 


5-18-1. 5-2-1で保持した[ProductGrossWeight]の値と、[入力ファイルのOrderQuantityInBaseUnit]の値を掛け算することにより、[ItemGrossWeight]の値を求め、セットする。  

5-19. ItemNetWeight  

| Property                | Description                                                                                                                              | EC  | 
| ----------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | --- | 
| ItemNetWeight           | 合計正味重量。正味重量とオーダー数量により合計正味重量が計算される。必要に応じて任意の合計正味重量を入力する。参照の場合は変更できない。 |     | 
| OrderQuantityInBaseUnit | オーダー数量(基本数量単位)。参照の場合は変更できない。                                                                                   |     |

5-19-1. 5-2-1で保持した[ProductNetWeight]の値と、[入力ファイルのOrderQuantityInBaseUnit]の値を掛け算することにより、[ItemNetWeight]の値を求め、セットする。  

5-20. TaxRateの計算  

| Property | Description                                                                                                 | EC  | 
| -------- | ----------------------------------------------------------------------------------------------------------- | --- | 
| TaxCode  | 消費税コード。オーダー明細のBPTaxClassification、ProductTaxClassificationの組合せから設定される。変更不可。 |     | 
| TaxRate  | 消費税率。設定されたTaxCodeから決定される。変更不可。                                                       |     | 

5-20-1. [5-1-3でセットしたTaxCode]を[コピーして値をセット]する。  

5-20-2. [Country=”JP”, 5-1-3でセットしたTaxCode, ValidityEndDate≧システム日付, ValidityStartDate≦システム日付]をキーとして、消費税率データの[TaxRate]を検索して値をセットする。  

| Target / Processing Type         | Key(s)                               | 
| -------------------------------- | ------------------------------------ | 
| TaxRate <br> / Get and Hold                   | [Country=”JP” <br> TaxCode(5-1-3でセット) <br> ValidityEndDate≧システム日付 <br> ValidityStartDate≦システム日付] | 
| <b>Table Searched</b>            | <b>Name of the Table</b>             | 
| Tax Code SQL Tax Rate Data       | data_platform_tax_code_tax_rate_data | 
| <b>Field Searched</b>            | <b>Data Type / Number of Digits</b>  | 
| TaxRate                          | float / 6                            | 
| <b>Single Record or Array</b>    | <b>Memo</b>                         | 
| Single Record                    |                                      |

5-21. NetAmount  

| Property  | Description                                                                                                      | EC  | 
| --------- | ---------------------------------------------------------------------------------------------------------------- | --- | 
| NetAmount | 正味金額。オーダー数量と条件価格とを乗じた金額が計算される。条件価格は明細価格決定要素にて計算される。変更不可。 |     | 

5-21-1. 8-2で計算した[ConditionAmount]の値を[NetAmount]にセットする。  

5-21-2. [NetAmount]の小数点以下の部分の桁数をカウントする。  

5-22. TaxAmount  

| Property  | Description                                                                                                                                                                                                                                                                                                             | EC  | 
| --------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
| TaxAmount | 税金額。明細の税分類が「1:納税義務」の場合、適切な税率にもとづいて消費税金額が自動計算される。明細の税分類が「0:非課税」の場合、税金額はゼロとなる。明細の税分類が「1:納税義務」の場合、ユーザが税金額を入力することはできるが、理論値の消費税額と2通貨単位以上の差が出る場合、エラーとなる。参照の場合は変更できない。 |     | 

5-22-1. 5-1-3でセットされた[TaxCode]が”A1”か”V1”である場合、5-20で保持した[TaxRate]に、5-21-1で求めたNetAmountを乗じる。5-22-1で求めた値の小数点以下部分を、5-21-2でカウントした桁数に四捨五入して、[TaxAmount]の値を求め、セットする。  

5-22-2. 5-1-3でセットされた[TaxCode]が”A1”と”V1”以外である場合、[TaxAmount]に[０]をセットする。  

5-22-3. 入力ファイルの[TaxAmount]がnullの場合、[5-22-1または5-22-2で求めたTaxAmount]を[TaxAmount]にセットする。  

5-22-4. [入力ファイルのTaxAmount]がnullでない場合、[入力ファイルのTaxAmount]を保持する。  

5-22-5. [入力ファイルのTaxAmount]がnullでない場合、[5-22-3でセットした値]から[5-22-4で保持した値]を引き算し、[引き算した値の絶対値]が2以上の明細が一つでもある場合、エラーメッセージを出力して処理を終了する。[引き算した値の絶対値]が2以上の明細が一つも無い場合、[5-22-3または5-22-5でセットまたは保持したTaxAmount]を[TaxAmount]にセットする。  

5-23. GrossAmount  

| Property    | Description                                                                                                                                                                                                         | EC  | 
| ----------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
| GrossAmount | 総額。消費税を含んだ総額が自動計算される。必要に応じて変更する。総額と正味金額との差額が消費税額として正しいかどうかチェックされる。理論値と2通貨単位以上の差額がある場合はエラーとなる。参照の場合は変更できない。 |     | 


5-23-1. [5-21のNetAmount]に、[5-22のTaxAmount]を加算し、[GrossAmount]として保持する。  

5-23-2. [入力ファイルのGrossAmount]がnullでない場合、[5-23-1で求めたGrossAmount]を[GrossAmount]にセットする。  

5-23-3. [入力ファイルのGrossAmount]がnullである場合、[入力ファイルのGrossAmount]を保持する。  

5-23-4. 5-23-2の値から5-23-3の値を引き算し、[引き算した値の絶対値]が2以上の明細が一つでもある場合、エラーメッセージを出力して処理を終了する。[引き算した値の絶対値]が2以上の明細が一つも無い場合、[5-23-3で求めたGrossAmount]を[GrossAmount]にセットする。  

5-24. ProductionPlantBatch  

入力ファイルの[OrderItemCategory=”INVP”の明細]に対して、本処理を実行する。  

| Property             | Description                                                                                                                                                                                                                                                         | EC  | 
| -------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
| ProductionPlantBatch | 製造プラントロット番号。ロット管理フラグとロット管理方針に基づき、必要な場合、オーダー時点で製造プラントのロット番号を入力する、または、ロットが自動決定される。（※ロットマスタに無い場合新規ロットとして自動登録される）参照[見積/引合以外]の場合は変更できない。 | ✔  | 

5-24-1. ロットマスタの存在性チェック(ProductionPlantBatch)  
入力ファイルの[Product, ProductionPlantBusinessPartner, ProductionPlant, ProductionPlantBatch, ValidityStartDate≦システム日付, ValidityEndDate≧システム日付]が、ロットマスタに存在するかどうかのチェックを行う。当該処理はdata_platform_batch_master_record_exconf にて行う。data_platform_batch_master_record_exconfの結果がtrueの場合、ロットマスタに存在する。同結果がfalseの場合、ロットマスタに存在しない。内部的に[ProductionPlantBatchExConf(bool / 1)]の項目を保持し、[結果がtrue]の場合、[同項目の値をtrue]とし、[結果がfalse]の場合、[同項目の値をfalse]とする。  

< data_platform_batch_master_record_exconfで処理する内容>(ProductionPlantBatch)  

| Target / Processing Type           | Key(s)                                             | 
| ---------------------------------- | -------------------------------------------------- | 
| ProductionPlantBatch <br> / API ExConf                       | [Product <br>  ProductionPlantBusinessPartner <br> ProductionPlant <br> ProductionPlantBatch <br>  ValidityStartDate≦システム日付 <br> ValidityEndDate≧システム日付]     | 
| <b>Table Searched</b>              | <b>Name of the Table</b>                           | 
| Batch Master Record SQL Batch Data | data_platform_batch_master_record_batch_data       | 
| <b>Field Searched</b>              | <b>Data Type / Number of Digits</b>                | 
| Product <br> BusinessPartner <br> Plant <br> Batch                              | string(varchar) / 40 <br> int / 12 <br> string(varchar) / 4 <br> string(varchar) / 10               | 
| <b>Single Record or Array</b>      | <b>Memo</b>                                       | 
| Array                              | data_platform_batch_master_record_exconfを用いる。 | 


5-24-2. ロットマスタの登録(ProductionPlantBatch)
5-24-1の結果がfalseの場合、ロットマスタの登録を行う。  

<ロットマスタ登録の処理内容(ProductionPlantBatch)＞  

| Target / Processing Type                                                                                 | Key(s)                                              | 
| -------------------------------------------------------------------------------------------------------- | --------------------------------------------------- | 
| ProductionPlantBatch <br> / API Creates                                                                                            | [Product <br> ProductionPlantBusinessPartner <br> ProductionPlant <br>  ProductionPlantBatch <br> ValidityStartDate=システム日付 <br> ValidityEndDate=9999-12-31]                                                                              | 
| <b>Table Searched</b>                                                                                    | <b>Name of the Table</b>                            | 
| Batch Master Record SQL Batch Data                                                                       | data_platform_batch_master_record_batch_data        | 
| <b>Field Created</b>                                                                                            | <b>Data Type / Number of Digits</b>                 | 
| Product <br> BusinessPartner <br> Plant <br> Batch <br> CountryOfOrigin(4-2-1で取得) <br>  ValidityStartDate(入力ファイルのProductionPlantBatchValidityStartDateまたは、ブランクの場合システム日付) <br> ManufactureDate(“”※ブランク) <br> CreationDateTime(システム日付時刻) <br> LastChangeDateTime(システム日付時刻) <br>  IsMarkedForDeletion(false)                                                                               | string(varchar) / 40 <br> int / 12 <br> string(varchar) / 4 <br> string(varchar) / 10 <br> string(varchar) / 3 <br> date <br /> <br /> date <br> datetime <br> datetime <br> bool (tinyint) / 1                                                                                       | 
| <b>Single Record or Array</b>                                                                            | <b>Memo</b>                                        | 
| Array                                                                                                    | data_platform_batch_master_record_createsを用いる。 | 

5-25. IssuingPlantBatch  
入力ファイルの[OrderItemCategory=”INVP”の明細]に対して、本処理を実行する。  

| Property          | Description                                                                                                                                                                                                                                                         | EC  | 
| ----------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
| IssuingPlantBatch | 出荷プラントロット番号。ロット管理フラグとロット管理方針に基づき、必要な場合、オーダー時点で出荷プラントのロット番号を入力する、または、ロットが自動決定される。（※ロットマスタに無い場合新規ロットとして自動登録される）参照[見積/引合以外]の場合は変更できない。 | ✔  | 

5-25-1. ロットマスタの存在性チェック(IssuingPlantBatch)
入力ファイルの[Product, IssuingPlantBusinessPartner, IssuingPlant, IssuingPlantBatch, ValidityStartDate≦システム日付, ValidityEndDate≧システム日付]が、ロットマスタに存在するかどうかのチェックを行う。当該処理はdata_platform_batch_master_record_exconf にて行う。data_platform_batch_master_record_exconfの結果がtrueの場合、ロットマスタに存在する。同結果がfalseの場合、ロットマスタに存在しない。内部的に[IssuingPlantBatchExConf(bool / 1)]の項目を保持し、[結果がtrue]の場合、[同項目の値をtrue]とし、[結果がfalse]の場合、[同項目の値をfalse]とする。  

< data_platform_batch_master_record_exconfで処理する内容>(IssuingPlantBatch)  

| Target / Processing Type           | Key(s)                                             | 
| ---------------------------------- | -------------------------------------------------- | 
| IssuingPlantBatch <br> / API ExConf                       | [Product <br>  IssuingPlantBusinessPartner <br> IssuingPlant <br> IssuingPlantBatch <br>  ValidityStartDate≦システム日付 <br> ValidityEndDate≧システム日付]     | 
| <b>Table Searched</b>              | <b>Name of the Table</b>                           | 
| Batch Master Record SQL Batch Data | data_platform_batch_master_record_batch_data       | 
| <b>Field Searched</b>              | <b>Data Type / Number of Digits</b>                | 
| Product <br> BusinessPartner <br> Plant <br> Batch                              | string(varchar) / 40 <br> int / 12 <br> string(varchar) / 4 <br> string(varchar) / 10               | 
| <b>Single Record or Array</b>      | <b>Memo</b>                                       | 
| Array                              | data_platform_batch_master_record_exconfを用いる。 | 


5-25-2. ロットマスタの自動登録(IssuingPlantBatch)
5-25-1の結果がfalseの場合、ロットマスタの登録を行う。

<ロットマスタ登録の処理内容(IssuingPlantBatch)＞  

| Target / Processing Type                                                                                 | Key(s)                                              | 
| -------------------------------------------------------------------------------------------------------- | --------------------------------------------------- | 
| IssuingPlantBatch <br> / API Creates                                                                                            | [Product <br> IssuingPlantBusinessPartner <br> IssuingPlant <br> IssuingPlantBatch <br> ValidityStartDate=システム日付 <br> ValidityEndDate=9999-12-31]                                                                              | 
| <b>Table Searched</b>                                                                                    | <b>Name of the Table</b>                            | 
| Batch Master Record SQL Batch Data                                                                       | data_platform_batch_master_record_batch_data        | 
| <b>Field Created</b>                                                                                            | <b>Data Type / Number of Digits</b>                 | 
| Product <br> BusinessPartner  <br> Plant <br> Batch <br> CountryOfOrigin(4-2-1で取得) <br> ValidityStartDate(入力ファイルのProductionPlantBatchValidityStartDateまたは、ブランクの場合システム日付) <br> ManufactureDate(ブランク)  <br> CreationDateTime(システム日付時刻)  <br> LastChangeDateTime(システム日付時刻)  <br> IsMarkedForDeletion(false)                                                                               | string(varchar) / 40 <br> int / 12 <br> string(varchar) / 4  <br> string(varchar) / 10 <br> string(varchar) / 3  <br> date <br /> <br /> date <br> datetime <br> datetime <br> bool (tinyint) / 1                                                                                       | 
| <b>Single Record or Array</b>                                                                            | <b>Memo</b>                                        | 
| Single Record                                                                                            | data_platform_batch_master_record_createsを用いる。 | 


5-26. ReceivingPlantBatch
入力ファイルの[OrderItemCategory=”INVP”の明細]に対して、本処理を実行する。  

| Property            | Description                                                                                                                                                                                                                                                         | EC  | 
| ------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
| ReceivingPlantBatch | 入荷プラントロット番号。ロット管理フラグとロット管理方針に基づき、必要な場合、オーダー時点で入荷プラントのロット番号を入力する、または、ロットが自動決定される。（※ロットマスタに無い場合新規ロットとして自動登録される）参照[見積/引合以外]の場合は変更できない。 | ✔  | 

5-26-1. ロットマスタの存在性チェック(ReceivingPlantBatch)
入力ファイルの[Product, ReceivingPlantBusinessPartner, ReceivingPlant, ReceivingPlantBatch, ValidityStartDate≦システム日付, ValidityEndDate≧システム日付]が、ロットマスタに存在するかどうかのチェックを行う。当該処理はdata_platform_batch_master_record_exconf にて行う。data_platform_batch_master_record_exconfの結果がtrueの場合、ロットマスタに存在する。同結果がfalseの場合、ロットマスタに存在しない。内部的に[ReceivingPlantBatchExConf(bool / 1)]の項目を保持し、[結果がtrue]の場合、[同項目の値をtrue]とし、[結果がfalse]の場合、[同項目の値をfalse]とする。  

< data_platform_batch_master_record_exconfで処理する内容>(ReceivingPlantBatch)  

| Target / Processing Type           | Key(s)                                             | 
| ---------------------------------- | -------------------------------------------------- | 
| ReceivingPlantBatch  / API ExConf                       | [Product <br> ReceivingPlantBusinessPartner  <br> ReceivingPlant  <br> ReceivingPlantBatch <br> ValidityStartDate≦システム日付 <br> ValidityEndDate≧システム日付]     | 
| <b>Table Searched</b>              | <b>Name of the Table</b>                           | 
| Batch Master Record SQL Batch Data | data_platform_batch_master_record_batch_data       | 
| <b>Field Searched</b>              | <b>Data Type / Number of Digits</b>                | 
| Product <br> BusinessPartner <br> Plant <br> Batch                              | string(varchar) / 40  <br> int / 12 <br> string(varchar) / 4 <br> string(varchar) / 10               | 
| <b>Single Record or Array</b>      | <b>Memo</b>                                       | 
| Array                              | data_platform_batch_master_record_exconfを用いる。 | 


5-26-2. ロットマスタの自動登録(ReceivingPlantBatch)  
5-26-1の結果がfalseの場合、ロットマスタの登録を行う。  

<ロットマスタ登録の処理内容(ReceivingPlantBatch)＞  

| Target / Processing Type                                                                                 | Key(s)                                              | 
| -------------------------------------------------------------------------------------------------------- | --------------------------------------------------- | 
| ReceivingPlantBatch  <br> / API Creates                                                                                            | [Product  <br> ReceivingPlantBusinessPartner  <br> ReceivingPlant  <br> ReceivingPlantBatch  <br> ValidityStartDate=システム日付 <br> ValidityEndDate=9999-12-31]                                                                              | 
| <b>Table Searched</b>                                                                                    | <b>Name of the Table</b>                            | 
| Batch Master Record SQL Batch Data                                                                       | data_platform_batch_master_record_batch_data        | 
| <b>Field Created</b>                                                                                            | <b>Data Type / Number of Digits</b>                 | 
| Product <br> BusinessPartner  <br> Plant  <br> Batch <br> CountryOfOrigin(4-2-1で取得) <br> ValidityStartDate(入力ファイルのProductionPlantBatchValidityStartDateまたは、ブランクの場合システム日付) <br> ManufactureDate(ブランク) <br> CreationDateTime(システム日付時刻) <br> LastChangeDateTime(システム日付時刻) <br> IsMarkedForDeletion(false)                                                                               | string(varchar) / 40 <br> int / 12 <br> string(varchar) / 4 <br> string(varchar) / 10 <br> string(varchar) / 3  <br> date <br /> <br /> date <br> datetime <br> datetime <br> bool (tinyint) / 1                                                                                       | 
| <b>Single Record or Array</b>                                                                            | <b>Memo</b>                                        | 
| Array                                                                                                    | data_platform_batch_master_record_createsを用いる。 | 


5-27. ItemCompleteDeliveryIsDefined  
入力ファイルの[OrderItemCategory=”INVP”の明細]に対して、本処理を実行する。  

| Property                      | Description                                                        | EC  | 
| ----------------------------- | ------------------------------------------------------------------ | --- | 
| ItemCompleteDeliveryIsDefined | 明細入出荷完了ステータス。以下のステータスが設定される。変更不可。 <br> true : 入出荷完了 <br> false : 入出荷未完了          |                                                                    | 


5-27-1.入力ファイルの[OrderID、OrderItem]をキーとして、対象の全ての入出荷伝票明細の[ItemCompleteDeliveryIsDefined]を検索して保持する。  

| Target / Processing Type                     | Key(s)                                    | 
| -------------------------------------------- | ----------------------------------------- | 
| ItemCompleteDeliveryIsDefined / Get and Hold | [OrderID <br> OrderItem]                                   | 
| <b>Table Searched</b>                        | <b>Name of the Table</b>                  | 
| Delivery Document SQL Item                   | data_platform_delivery_document_item_data | 
| <b>Field Searched</b>                        | <b>Data Type / Number of Digits</b>       | 
| ItemCompleteDeliveryIsDefined                | bool (tinyint) / 1                        | 
| <b>Single Record or Array</b>                | <b>Memo</b>                              | 
| Array                                        |                                           | 


5-27-2. 5-27-1の検索結果が0件であった場合、ItemCompleteDeliveryIsDefinedにfalse(入出荷未完了)をセットする。検索結果が1件以上であり、かつ、全ての検索結果値がfalseの場合はfalse(入出荷未完了)をセットする。検索結果が1件以上であり、かつ、 全ての検索結果値がtrueの場合は完全出荷true(入出荷完了)をセットする。それ以外の場合はfalse(入出荷未完了)をセットする。  

<5-27-2のロジックまとめ>  

| 検索結果件数 | 検索結果値(ItemCompleteDeliveryIsDefined) | ItemCompleteDeliveryIsDefinedにセットする値 | 
| ------------ | ----------------------------------------- | ------------------------------------------- | 
| 0件          | -                                         | false                                       | 
| 1件以上      | 全てfalse                                 | false                                       | 
| 1件以上      | 全てtrue                                  | true                                        | 
| 1件以上      | 上記以外                                  | false                                       | 


5-28. 在庫確認
入力ファイルの[OrderItemCategory=”INVP”の明細]に対して、本処理を実行する。    
5-28-1.在庫確認是非の判断  

| Property                 | Description                                                                                                                                                                                                                                                                 | EC  | 
| ------------------------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
| OrderItemCategory        | オーダー明細カテゴリ。品目の明細カテゴリからINVPまたはSRVPが設定される。必要に応じてSDTP:仕入先直送明細に変更する。品目の明細カテゴリと合わない明細カテゴリは設定できない。参照の場合は変更できない。INVPまたはSRVPが設定または入力された場合、利用可能在庫確認が行われる。 <br>INVP:在庫明細  <br> SRVP:サービス/非在庫明細 <br>SDTP:仕入先直送明細      |                                                                                                                                                                                                                                                                             | 


5-28-1-1. [OrderItemCategory=”INVP”]の時、本在庫確認の処理を実行する。  

5-28-2. 在庫確認①(ロット単位での在庫確認)  
5-28-2-1. [5-5-1の品目在庫確認プラントにおけるIsBatchManagementRequired]が[true]だった場合、[対象のOrderID, OrderItem, Product]についてロット単位での在庫確認を行う。  

5-28-2-2. [OrderID, OrderItem, OrderQuantityUnit, OrderQuantityInBaseUnit, StockConfirmationPartnerFunction]に基づき、[入力ファイルのProduct, 5-4で取得したStockConfirmationBusinessPartner, 5-4で取得したStockConfirmationPlant, 入力ファイルのStockConfirmationPlantBatch, 入力ファイルのStockConfirmationPlantBatchValidityStartDate, 入力ファイルのStockConfirmationPlantBatchValidityEndDate, 入力ファイルのRequestedDeliveryDate]をinputとして、data_platform_function_product_availability_check(利用可能在庫確認)を実行し、品目利用可能在庫データの[Product, BusinessPartner, Plant, Batch, BatchValidityEndDate, ProductStockAvailabilityDate, AvailableProductStock]、を取得して保持する。プラントやロット、日付等のキー違いなどでエラーが発生した場合、エラーメッセージを出力して終了する。  

| Target / Processing Type                                   | Key(s)                                            | 
| ---------------------------------------------------------- | ------------------------------------------------- | 
| Product <br> BusinessPartner <br> Plant <br> Batch <br> BatchValidityEndDate <br> ProductStockAvailabilityDate <br> AvailableProductStock <br> / Get and Hold                                             | Product(入力ファイル)  <br> StockConfirmationBusinessPartner(5-4で取得) <br> StockConfirmationPlant(5-4で取得)  <br> RequestedDeliveryDate(入力ファイル) <br> StockConfirmationPlantBatch(入力ファイル)  <br> StockConfirmationPlantBatchValidityStartDate(入力ファイル) <br>StockConfirmationPlantBatchValidityEndDate(入力ファイル)   | 
| <b>Table Searched</b>                                      | <b>Name of the Function</b>                              | 
| None                                                       | data_platform_function_product_availability_check | 
| <b>Field Searched</b>                                      | <b>Data Type / Number of Digits</b>               | 
| Product  <br> BusinessPartner  <br> Plant <br> Batch <br> BatchValidityEndDate <br> ProductStockAvailabilityDate  <br> AvailableProductStock                                      | string(varchar) / 40 <br> int / 12 <br>  string(varchar) / 4 <br> string(varchar) / 10  <br> date <br> date  <br> float / 15                                                 | 
| <b>Single Record or Array</b>                              | <b>Memo</b>                                      | 
| Array                                                      |                                                   | 


5-28-2-3. 納入日程行(Orders Item Schedule Line)の生成(ロット在庫)    
[5-28-2-2で在庫確認が行われたOrderID, OrdeItemの明細]について、同じ明細数だけ、次の処理の通り、[納入日程行]を生成する。    
＜納入日程行の生成処理＞※在庫確認の結果の在庫確認数量がゼロである場合も納入日程行が生成される。  

| 生成する項目                                                                                     | 生成する値等のロジック                                                                                         | 
| ------------------------------------------------------------------------------------------------ | -------------------------------------------------------------------------------------------------------------- | 
| OrderID(※Key)                                                                                   | 本処理(5-28-2)にて保持済のOrderID。                                                                            | 
| OrderItem(※Key)                                                                                 | 本処理(5-28-2)にて保持済のOrderItem。                                                                          | 
| ScheduleLine(※Key)                                                                              | 本処理(5-28-2)内にて1～999の連番で生成する。                                                                   | 
| Product                                                                                          | 本処理(5-28-2)にて保持済みのProduct。                                                                          | 
| StockConfirmationPartnerFunction                                                                 | 本処理(5-28-2)にて保持済みのStockConfirmationPartnerFunction。                                                 | 
| StockConfirmationBusinessPartner                                                                 | 本処理(5-28-2)にて保持済みのStockConfirmationBusinessPartner。                                                 | 
| StockConfirmationPlant                                                                           | 本処理(5-28-2)にて保持済みのStockConfirmationPlant。                                                           | 
| StockConfirmationPlantBatch                                                                      | 本処理(5-28-2)にて保持済みのStockConfirmationPlantBatch。                                                      | 
| StockConfirmationPlantBatchValidityStartDate                                                     | 本処理(5-28-2)にて保持済みのStockConfirmationPlantBatchValidityStartDate                                       | 
| StockConfirmationPlantBatchValidityEndDate                                                       | 本処理(5-28-2)にて保持済みのStockConfirmationPlantBatchValidityEndDate                                         | 
| RequestedDeliveryDate                                                                            | 本処理(5-28-2)にて保持済みのRequestedDeliveryDate。                                                            | 
| ConfirmedDeliveryDate                                                                            | 5-28-2-2で取得したProductStockAvailabilityDate <br>※確認された在庫が無い場合、API処理の結果により、RequestedDeliveryDateと同じ日付がセットされる。 | 
| OrderQuantityInBaseUnit                                                                          | 本処理(5-28-2)にて保持済みのOrderQuantityInBaseUnit。                                                          | 
| ConfdOrderQtyByPDTAvailCheck                                                                     | 5-28-2-2で取得したAvailableProductStock <br> ※確認された在庫が無い場合、API処理の結果により、ゼロがセットされる。                            | 
| DeliveredQtyInOrderQtyUnit                                                                       | nullをセット。                                                                                                 | 
| OpenConfdDelivQtyInOrdQtyUnit                                                                    | OrderQuantityInBaseUnitからConfdOrderQtyByPDTAvailCheckを減算した値を計算してセット。                          | 
| StockIsFullyConfirmed                                                                            | OrderQuantityInBaseUnit= ConfdOrderQtyByPDTAvailCheckである場合、trueをセット。それ以外の場合、falseをセット。 | 
| DelivBlockReasonForSchedLine                                                                     | Falseをセット。                                                                                                | 
| PlusMinusFlag                                                                                    | “-“(マイナス)をセット。                                                                                      | 


5-28-2-4. オーダー明細への在庫確認済フラグ/在庫確認ステータスのセット  

| Property                | Description                                                            | EC  | 
| ----------------------- | ---------------------------------------------------------------------- | --- | 
| StockConfirmationStatus | 在庫確認ステータス。次の在庫確認ステータスが常時更新される。変更不可。<br> CL:全部完了済 <br> PP:部分完了済 <br> NP:確認済数量無し       |     | 


5-28-2-3において、[StockIsFullyConfirmed]がtrueの場合、[StockConfirmationStatus]に”CL”をセットする。[StockIsFullyConfirmed]がfalseでかつ[ConfdOrderQtyByPDTAvailCheck]がゼロである場合、”NP”をセットする。それ以外の場合、”PP”をセットする。  

5-28-3. 在庫確認②(通常の在庫確認)  
5-28-3-1. 5-5-1の品目在庫確認プラントにおけるIsBatchManagementRequiredがfalseだった場合、次の処理を行う。  

5-28-3-2. [OrderID, OrderItem, OrderQuantityUnit, OrderQuantityInBaseUnit, StockConfirmationPartnerFunction]に基づき、[入力ファイルのProduct, 5-4で取得したStockConfirmationBusinessPartner, 5-4で取得したStockConfirmationPlant, 入力ファイルのRequestedDeliveryDate]をinputとして、data_platform_function_product_availability_check(利用可能在庫確認)を実行し、品目利用可能在庫データの[Product, BusinessPartner, Plant, ProductStockAvailabilityDate, AvailableProductStock]、を取得して保持する。プラント、日付等のキー違いなどでエラーが発生した場合、エラーメッセージを出力して終了する。  

| Target / Processing Type                    | Key(s)                                            | 
| ------------------------------------------- | ------------------------------------------------- | 
| Product <br> BusinessPartner <br> Plant <br> ProductStockAvailabilityDate <br> AvailableProductStock <br> / Get and Hold                              | Product(入力ファイル) <br> StockConfirmationBusinessPartner(5-4で取得) <br> StockConfirmationPlant(5-4で取得) <br> RequestedDeliveryDate(入力ファイル)         | 
| <b>Table Searched</b>                       | <b>Name of the Function</b>                              | 
| None                                        | data_platform_function_product_availability_check | 
| <b>Field Searched</b>                       | <b>Data Type / Number of Digits</b>               | 
| Product <br> BusinessPartner <br> Plant <br> ProductStockAvailabilityDate <br> AvailableProductStock   | string(varchar) / 40 <br> int / 12 <br> string(varchar) / 4 <br> date <br> float / 15                                  | 
| <b>Single Record or Array</b>               | <b>Memo</b>                                      | 
| Array                                       |                                                   | 


5-28-3-3. 納入日程行(Orders Item Schedule Line)の生成(通常の在庫確認)  
[5-28-3-2で在庫確認が行われたOrderID, OrdeItemの明細]について、同じ明細数だけ、次の処理の通り、[納入日程行]を生成する。  
＜納入日程行の生成処理＞※在庫確認の結果の在庫確認数量がゼロである場合も納入日程行が生成される。  

| 生成する項目                                                                                     | 生成する値等のロジック                                                                                         | 
| ------------------------------------------------------------------------------------------------ | -------------------------------------------------------------------------------------------------------------- | 
| OrderID(※Key)                                                                                   | 本処理(5-28-3)にて保持済のOrderID。                                                                            | 
| OrderItem(※Key)                                                                                 | 本処理(5-28-3)にて保持済のOrderItem。                                                                          | 
| ScheduleLine(※Key)                                                                              | 本処理(5-28-3)内にて1～999の連番で生成する。                                                                   | 
| Product                                                                                          | 本処理(5-28-3)にて保持済みのProduct。                                                                          | 
| StockConfirmationPartnerFunction                                                                 | 本処理(5-28-3)にて保持済みのStockConfirmationPartnerFunction。                                                 | 
| StockConfirmationBusinessPartner                                                                 | 本処理(5-28-3)にて保持済みのStockConfirmationBusinessPartner。                                                 | 
| StockConfirmationPlant                                                                           | 本処理(5-28-3)にて保持済みのStockConfirmationPlant。                                                           | 
| StockConfirmationPlantBatch                                                                      | “”(ブランク)をセット。                                                                                       | 
| StockConfirmationPlantBatchValidityStartDate                                                     | nullをセット。                                                                                                 | 
| StockConfirmationPlantBatchValidityEndDate                                                       | nullをセット。                                                                                                 | 
| RequestedDeliveryDate                                                                            | 本処理(5-28-3)にて保持済みのRequestedDeliveryDate。                                                            | 
| ConfirmedDeliveryDate                                                                            | 5-28-3-2で取得したProductStockAvailabilityDate <br> ※確認された在庫が無い場合、API処理の結果により、RequestedDeliveryDateと同じ日付がセットされる。 | 
| OrderQuantityInBaseUnit                                                                          | 本処理(5-28-3)にて保持済みのOrderQuantityInBaseUnit。                                                          | 
| ConfdOrderQtyByPDTAvailCheck                                                                     | 5-28-3-2で取得したAvailableProductStock  <br> ※確認された在庫が無い場合、API処理の結果により、ゼロがセットされる。                            | 
| DeliveredQtyInOrderQtyUnit                                                                       | nullをセット。                                                                                                 | 
| OpenConfdDelivQtyInOrdQtyUnit                                                                    | OrderQuantityInBaseUnitからConfdOrderQtyByPDTAvailCheckを減算した値を計算してセット。                          | 
| StockIsFullyConfirmed                                                                            | OrderQuantityInBaseUnit= ConfdOrderQtyByPDTAvailCheckである場合、trueをセット。それ以外の場合、falseをセット。 | 
| DelivBlockReasonForSchedLine                                                                     | Falseをセット。                                                                                                | 
| PlusMinusFlag                                                                                    | “-“(マイナス)をセット。                                                                                      | 


5-28-3-4. オーダー明細への在庫確認済フラグ/在庫確認ステータスのセット

| Property                | Description                                                            | EC  | 
| ----------------------- | ---------------------------------------------------------------------- | --- | 
| StockConfirmationStatus | 在庫確認ステータス。次の在庫確認ステータスが常時更新される。変更不可。 <br> CL:全部完了済 <br> PP:部分完了済 <br> NP:確認済数量無し  |    |	

5-28-3-3において、[StockIsFullyConfirmed]がtrueの場合、[StockConfirmationStatus]に”CL”をセットする。[StockIsFullyConfirmed]がfalseでかつ[ConfdOrderQtyByPDTAvailCheck]がゼロである場合、”NP”をセットする。それ以外の場合、”PP”をセットする。  

5-29. PricingDate  

| Property    | Description                                                  | EC  | 
| ----------- | ------------------------------------------------------------ | --- | 
| PricingDate | 価格設定日付。ヘッダの価格設定日付がコピーされる。変更不可。 | ✔  | 

5-29-1. HeaderのPricingDateから、[明細の配列保持データ]に対して[PricingDateをコピー]する。  

5-30. StorkConfirmationPartnerFunction, StockConfirmationBusinessPartner, StockConfirmationPlant  
入力ファイルの[OrderItemCategory=”INVP”の明細]に対して、本処理を実行する。  

| Property                         | Description                                                                                                                                                  | EC  | 
| -------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | --- | 
| StockConfirmationPartnerFunction | 在庫確認を行う取引先機能。ヘッダパートナプラントから設定される。必要に応じて変更する。（PRODUCT_STOCK_AVAIRABILITY_CHECK_SRVを利用、機能はオプション）<br> 参照の場合は変更できない。       | ✔                                                                                                                                                           | 
| StockConfirmationBusinessPartner | 在庫確認を行うビジネスパートナ。ヘッダパートナプラントから設定される。必要に応じて変更する。（PRODUCT_STOCK_AVAIRABILITY_CHECK_SRVを利用、機能はオプション）<br> 参照の場合は変更できない。       | ✔                                                                                                                                                           | 
| StockConfirmationPlant           | 在庫確認プラント。ヘッダパートナプラントから設定される。必要に応じて変更する。<br> 参照の場合は変更できない。       | ✔                                                                                                                                                           | 

5-30-1. Headerの[StorkConfirmationPartnerFunction, StockConfirmationBusinessPartner, StockConfirmationPlant]から、[明細の配列保持データ]に対して[StorkConfirmationPartnerFunction, StockConfirmationBusinessPartner, StockConfirmationPlantをコピー]する。  

5-30-2. [入力ファイルのStorkConfirmationPartnerFunction, StockConfirmationBusinessPartner, StockConfirmationPlant]が全てnull以外の場合、[入力ファイルのStorkConfirmationPartnerFunction, StockConfirmationBusinessPartner, StockConfirmationPlant]をセットする。  

5-31. ConfdDelivQtyInOrderQtyUnit  
入力ファイルの[OrderItemCategory=”INVP”の明細]に対して、本処理を実行する。  

| Property                    | Description                                                                  | EC  | 
| --------------------------- | ---------------------------------------------------------------------------- | --- | 
| ConfdDelivQtyInOrderQtyUnit | 在庫確認済数量。（機能はオプション）在庫確認済の数量が更新される。変更不可。 |    |	

5-31-1. 5-28-2-3, または, 5-28-3-3の[ConfdOrderQtyByPDTAvailCheckの値]をセットする。  

5-32. BillingDocumentDate  

| Property            | Description                                                                                        | EC  | 
| ------------------- | -------------------------------------------------------------------------------------------------- | --- | 
| BillingDocumentDate | 請求伝票日付。ヘッダの請求伝票日付がコピーされる。必要に応じて変更する。参照の場合は変更できない。 |    |	

5-32-1. HeaderのBillingDocumentDateから、[明細の配列保持データ]に対して[BillingDocumentDateをコピー]する。  

5-32-2. [入力ファイルのBillingDocumentDate]がnull以外の場合、[入力ファイルのBillingDocumentDate]をセットする。  

5-33. PriceDetnExchangeRate  

| Property              | Description                                                                                  | EC  | 
| --------------------- | -------------------------------------------------------------------------------------------- | --- | 
| PriceDetnExchangeRate | 価格決定のための為替レート。ヘッダの、価格決定のための為替レート、がコピーされる。変更不可。 |     |	

5-33-1. HeaderのPriceDetnExchangeRateから、[明細の配列保持データ]に対して[PriceDetnExchangeRateをコピー]する。  

5-34. AccountingExchangeRate  

| Property               | Description                                                        | EC  | 
| ---------------------- | ------------------------------------------------------------------ | --- | 
| AccountingExchangeRate | 会計為替レート。ヘッダから会計為替レートがコピーされる。変更不可。 |      |	

5-34-1. HeaderのAccountingExchangeRateから、[明細の配列保持データ]に対して[AccountingExchangeRateをコピー]する。  

5-35. ReferenceDocument / ReferenceDocumentItem  

| Property                                                             | Description                                                                                                                                                                                                                                          | EC  | 
| -------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
| ReferenceDocument                                                    | 参照伝票。見積参照の場合は見積番号、引合参照の場合は引合番号、購買依頼参照の場合は購買依頼番号、オーダー参照の場合はオーダー番号、入出荷伝票参照の場合は入出荷伝票番号、がヘッダからコピーされる。変更不可。                                         |     | 
| ReferenceDocumentItem                                                | 参照伝票明細。購買依頼参照の場合は購買依頼明細番号、オーダー参照の場合はオーダー明細番号、購買依頼参照の場合は購買依頼明細番号、オーダー参照の場合はオーダー明細番号、入出荷伝票参照の場合は入出荷伝票明細番号、がヘッダからコピーされる。変更不可。 <br> ※見積参照および引合参照の場合はヘッダレベルで明細単位で参照できない |     |

5-35-1. [入力ファイルのHeaderのReferenceDocument, ReferenceDocumentItem]から、[明細の配列保持データ]に対して[ReferenceDocument, ReferenceDocumentItemをコピー]する。  

5-36. ItemDeliveryStatus  

入力ファイルの[OrderItemCategory=”INVP”の明細]に対して、本処理を実行する。  

| Property             | Description                                                    | EC  | 
| -------------------- | -------------------------------------------------------------- | --- | 
| ItemDeliveryStatus   | 明細入出荷ステータス。以下のステータスが設定される。変更不可。 <br> 未入出荷：NP <br> 部分入出荷完了済：PP <br> 入出荷完了済：CL     |     |

5-36-1.入力ファイルの[OrderID、OrderItem]をキーとして、対象の全ての入出荷伝票明細の[ItemCompleteDeliveryIsDefined]を検索して保持する。  

| Target / Processing Type      | Key(s)                                    | 
| ----------------------------- | ----------------------------------------- | 
| ItemDeliveryStatus <br> / Get and Set   | [OrderID <br> OrderItem]                    | 
| <b>Table Searched</b>         | <b>Name of the Table</b>                  | 
| Delivery Document SQL Item    | data_platform_delivery_document_item_data | 
| <b>Field Searched</b>         | <b>Data Type / Number of Digits</b>       | 
| ItemCompleteDeliveryIsDefined | bool (tinyint) / 1                        | 
| <b>Single Record or Array</b> | <b>Memo</b>                              | 
| Array                         | 	

5-36-2. 5-36-1の検索結果が0件であった場合、ItemDeliveryStatus に”NP”(未入出荷)をセットする。検索結果が1件以上であり、かつ、全ての検索結果値がfalseの場合は”NP”(未入出荷)をセットする。検索結果が1件以上であり、かつ、 全ての検索結果値がtrueの場合は”CL”(入出荷完了済)をセットする。それ以外の場合は”PP”(部分出荷完了済)をセットする。  

<5-36-2のロジックまとめ>  

| 検索結果件数 | 検索結果値(ItemCompleteDeliveryIsDefined) | ItemDeliveryStatusにセットする値 | 
| ------------ | ----------------------------------------- | -------------------------------- | 
| 0件          | -                                         | “NP”                           | 
| 1件以上      | 全てfalse                                 | “NP”                           | 
| 1件以上      | 全てtrue                                  | “CL”                           | 
| 1件以上      | 上記以外                                  | “PP”                           | 

5-37. IssuingStatus  

入力ファイルの[OrderItemCategory=”INVP”の明細]に対して、本処理を実行する。  

| Property           | Description                                              | EC  | 
| ------------------ | -------------------------------------------------------- | --- | 
| IsssuingStatus     | 出荷ステータス。以下のステータスが設定される。変更不可。<br> 未出荷：NP  <br> 部分出荷完了済：PP <br> 出荷完了済：CL     |    |

5-37-1.入力ファイルの[OrderID、OrderItem]をキーとして、対象の全ての入出荷伝票明細の[ActualGoodsIssueDate]を検索して保持する。  

| Target / Processing Type   | Key(s)                                    | 
| -------------------------- | ----------------------------------------- | 
| IsssuingStatus <br> / Get and Hold             | [OrderID <br> OrderItem]                 | 
| <b>Table Searched</b>      | <b>Name of the Table</b>                  | 
| Delivery Document SQL Item | data_platform_delivery_document_item_data | 
| <b>Field Searched</b>      | <b>Data Type / Number of Digits</b>       | 
| ActualGoodsIssueDate       | date                                      | 
| <b>Single Record or Array</b> | <b>Memo</b>                              | 
| Array                      | 	

5-37-2. 5-37-1の検索結果が0件であった場合、IsssuingStatusに”NP”(未入出荷)をセットする。検索結果が1件以上であり、かつ、全ての検索結果値がnullの場合は”NP”(未入出荷)をセットする。検索結果が1件以上であり、かつ、 全ての検索結果値がnull以外の場合は”CL”(入出荷完了済)をセットする。それ以外の場合は”PP”(部分出荷完了済)をセットする。  

<5-37-2のロジックまとめ>  

| 検索結果件数 | 検索結果値(IsssuingStatus) | IsssuingStatusにセットする値 | 
| ------------ | -------------------------- | ---------------------------- | 
| 0件          | -                          | “NP”                       | 
| 1件以上      | 全てnull                   | “NP”                       | 
| 1件以上      | 全てnull以外               | “CL”                       | 
| 1件以上      | 上記以外                   | “PP”                       | 

5-38. ReceivingStatus  

入力ファイルの[OrderItemCategory=”INVP”の明細]に対して、本処理を実行する。  

| Property           | Description                                              | EC  | 
| ------------------ | -------------------------------------------------------- | --- | 
| ReceivingStatus    | 入荷ステータス。以下のステータスが設定される。変更不可。<br> 未入荷：NP <br> 部分入荷完了済：PP <br> 入荷完了済：CL     |     |

5-38-1.入力ファイルの[OrderID、OrderItem]をキーとして、対象の全ての入出荷伝票明細の[ActualGoodsReceiptDate]を検索して保持する。  

| Target / Processing Type   | Key(s)                                    | 
| -------------------------- | ----------------------------------------- | 
| ReceivingStatus <br>/ Get and Hold   | [OrderID <br> OrderItem]        | 
| <b>Table Searched</b>      | <b>Name of the Table</b>                  | 
| Delivery Document SQL Item | data_platform_delivery_document_item_data | 
| <b>Field Searched</b>      | <b>Data Type / Number of Digits</b>       | 
| ActualGoodsReceiptDate     | date                                      | 
| <b>Single Record or Array</b> | <b>Memo</b>                              | 
| Array                      |                                           |

5-38-2. 5-38-1の検索結果が0件であった場合、ReceivingStatusに”NP”(未入出荷)をセットする。検索結果が1件以上であり、かつ、全ての検索結果値がnullの場合は”NP”(未入出荷)をセットする。検索結果が1件以上であり、かつ、 全ての検索結果値がnull以外の場合は”CL”(入出荷完了済)をセットする。それ以外の場合は”PP”(部分出荷完了済)をセットする。  

<5-38-2のロジックまとめ>  

| 検索結果件数 | 検索結果値(ActualGoodsReceiptDate) | ReceivingStatusにセットする値 | 
| ------------ | ---------------------------------- | ----------------------------- | 
| 0件          | -                                  | “NP”                        | 
| 1件以上      | 全てnull                           | “NP”                        | 
| 1件以上      | 全てnull以外                       | “CL”                        | 
| 1件以上      | 上記以外                           | “PP”                        | 

5-39. BillingStatus  

| Property           | Description                                              | EC  | 
| ------------------ | -------------------------------------------------------- | --- | 
| BillingStatus      | 請求ステータス。以下のステータスが設定される。変更不可。 <br> 未請求：NP <br> 一部請求完了済：PP <br> 請求完了済：CL     |       | 

5-39-1.入力ファイルのOrderID , OrderItemをキーとして、対象の全ての請求伝票のOrderID , OrderItemを検索して値を保持する。  

| Target / Processing Type  | Key(s)                                   | 
| ------------------------- | ---------------------------------------- | 
| OrderID <br> OrderItem <br> / Get and Hold | OrderID <br> OrderItem  | 
| <b>Table Searched</b>     | <b>Name of the Table</b>                 | 
| Invoice Document SQL Item | data_platform_invoice_document_item_data | 
| <b>Field Searched</b>     | <b>Data Type / Number of Digits</b>      | 
| OrderID<br> OrderItem     | int / 16 <br> int / 6                   | 
| <b>Single Record or Array</b> | <b>Memo</b>                                     | 
| Array                     |                                          |

5-39-2. 5-39-1の検索結果が0件であった場合、BillingStatusに未請求”NP”をセットする。検索結果が全件であった場合、請求完了”CL”をセットする。それ以外の場合、一部請求完了済”PP”をセットする。  

<ロジックまとめ>  

| 検索結果件数 | BillingStatusにセットする値 | 
| ------------ | --------------------------- | 
| 0件          | “NP”                      | 
| 全件         | “CL”                      | 
| 上記以外     | “PP”                      | 

5-40. StockConfirmationStatus  
入力ファイルの[OrderItemCategory=”INVP”の明細]に対して、本処理を実行する。  

| Property                | Description                                                        | EC  | 
| ----------------------- | ------------------------------------------------------------------ | --- | 
| StockConfirmationStatus | 在庫確認ステータス。次の在庫確認ステータスが設定される。変更不可。 <br> CL:全部完了済  <br> PP:部分完了済 <br> ZR:確認済数量ゼロ       |                                                                    | 

5-40-1. 5-28.の在庫確認のデータにおいて、[ConfdOrderQtyByPDTAvailCheck]がゼロである場合、[StockConfirmationStatus]に”ZR”をセットする。それ以外の場合で、[OrderQuantityInBaseUnit]と [ConfdOrderQtyByPDTAvailCheck]の値が同じである場合、[StockConfirmationStatus]に”CL”をセットする。それ以外の場合、[StockConfirmationStatus]に”PP”をセットする。  

5-41. OrderItemTextByBuyer  

| Property             | Description                                                                                                                                                                                     | EC  | 
| -------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
| OrderItemTextByBuyer | Buyerによるオーダー明細テキスト。品目マスタビジネスパートナテキストデータ(Buyer)から初期値が提案される。必要に応じて任意のBuyerによるオーダー明細テキストを入力する。参照の場合は変更できない。 |     | 

5-41-1. [Product=入力ファイルのProduct, BusinessPartner=入力ファイルのBuyer, Language=2.2.1で取得したPartnerFunction=”Buyer”のLanguage]をキーとして、対象の品目マスタビジネスパートナテキストデータ(Buyer)の[ProductDescription]を検索して値をセットする。  

| Target / Processing Type                                        | Key(s)                                                                    | 
| --------------------------------------------------------------- | ------------------------------------------------------------------------- | 
| OrderItemTextByBuyer <br> / Get and Set                                                   | [Product(入力ファイル) <br> BusinessPartner=入力ファイルのBuyer <br> Language=2.2.1で取得したPartnerFunction=”Buyer”のLanguage]    | 
| <b>Table Searched</b>                                           | <b>Name of the Table</b>                                                  | 
| Product Master SQL Product Description By Business Partner Data | data_platform_product_master_product_description_by_business_partner_data | 
| <b>Field Searched</b>                                           | <b>Data Type / Number of Digits</b>                                       | 
| ProductDescription                                              | string(varchar) / 200                                                     | 
| <b>Single Record or Array</b>                                   | <b>Memo</b>                                                              | 
| Array                                                           |                                                                           | 

5-42. OrderItemTextBySeller  

| Property              | Description                                                                                                                                                                                        | EC  | 
| --------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
| OrderItemTextBySeller | Sellerによるオーダー明細テキスト。品目マスタビジネスパートナテキストデータ(Seller)から初期値が提案される。必要に応じて任意のSellerによるオーダー明細テキストを入力する。参照の場合は変更できない。 |     |	

5-42-1. [Product=入力ファイルのProduct, BusinessPartner=入力ファイルのSeller, Language=2.2.1で取得したPartnerFunction
=”Seller”のLanguage]をキーとして、対象の品目マスタビジネスパートナテキストデータ(Seller)の[ProductDescription]を検索して値をセットする。  

| Target / Processing Type                                        | Key(s)                                                                    | 
| --------------------------------------------------------------- | ------------------------------------------------------------------------- | 
| OrderItemTextBySeller <br> / Get and Set                        | [Product(入力ファイル) <br> BusinessPartner=入力ファイルのSeller <br> Language=2.2.1で取得したPartnerFunction=”Seller”のLanguage]   | 
| <b>Table Searched</b>                                           | <b>Name of the Table</b>                                                  | 
| Product Master SQL Product Description By Business Partner Data | data_platform_product_master_product_description_by_business_partner_data | 
| <b>Field Searched</b>                                           | <b>Data Type / Number of Digits</b>                                       | 
| ProductDescription                                              | string(varchar) / 200                                                     | 
| <b>Single Record or Array</b>                                   | <b>Memo</b>                                                              | 
| Array                                                           |                                                                           |   


| Property                               | Description                                                                                                                                                                                                         | EC  | 
| -------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
| StockConfirmationStatus                | 在庫確認ステータス。次の在庫確認ステータスが常時更新される。変更不可。<br> CL:全部完了済 <br> PP:部分完了済 <br> NP:未確認  |     | 
| ProductIsBatchManagedInProductionPlant | ロット管理フラグ。品目マスタBPプラントデータから参照される。変更不可。                                                                                                                                              |     | 
| ProductIsBatchManagedInIssuingPlant    | ロット管理フラグ。品目マスタBPプラントデータから参照される。変更不可。                                                                                                                                              |     | 
| ProductIsBatchManagedInReceivingPlant  | ロット管理フラグ。品目マスタBPプラントデータから参照される。変更不可。                                                                                                                                              |     | 
| BatchMgmtPolicyInProductionPlant       | ロット管理方針。品目マスタBPプラントデータから参照される。変更不可。                                                                                                                                                |     | 
| BatchMgmtPolicyInIssuingPlant          | ロット管理方針。品目マスタBPプラントデータから参照される。変更不可。                                                                                                                                                |     | 
| BatchMgmtPolicyInReceivingPlant        | ロット管理方針。品目マスタBPプラントデータから参照される。変更不可。                                                                                                                                                |     | 
| ProductionPlantBatchValidityStartDate  | 製造プラントロット有効開始日付。ロット管理フラグとロット管理方針に基づき、必要な場合、製造プラントのロットの有効開始日付を入力する、または、有効開始日付が自動決定される。参照[見積/引合以外]の場合は変更できない。 |     | 
| ProductionPlantBatchValidityEndDate    | 製造プラントロット有効終了日付。ロット管理フラグとロット管理方針に基づき、必要な場合、製造プラントのロットの有効終了日付を入力する、または、有効終了日付が自動決定される。参照[見積/引合以外]の場合は変更できない。 |     | 
| IssuingPlantBatchValidityStartDate     | 出荷プラントロット有効開始日付。ロット管理フラグとロット管理方針に基づき、必要な場合、出荷プラントのロットの有効開始日付を入力する、または、有効開始日付が自動決定される。参照[見積/引合以外]の場合は変更できない。 |     | 
| IssuingPlantBatchValidityEndDate       | 出荷プラントロット有効終了日付。ロット管理フラグとロット管理方針に基づき、必要な場合、出荷プラントのロットの有効終了日付を入力する、または、有効終了日付が自動決定される。参照[見積/引合以外]の場合は変更できない。 |     | 
| ReceivingPlantBatchValidityStartDate   | 入荷プラントロット有効開始日付。ロット管理フラグとロット管理方針に基づき、必要な場合、入荷プラントのロットの有効開始日付を入力する、または、有効開始日付が自動決定される。参照[見積/引合以外]の場合は変更できない。 |     | 
| ReceivingPlantBatchValidityEndDate     | 入荷プラントロット有効終了日付。ロット管理フラグとロット管理方針に基づき、必要な場合、入荷プラントのロットの有効終了日付を入力する、または、有効終了日付が自動決定される。参照[見積/引合以外]の場合は変更できない。 |     | 

6. Orders Item Partner
次の補助機能の開発を行う。

6-1. ヘッダ取引先機能データのコピー  

| Property                                    | Description                                                                                                                                                                                                                                                                                                                                                                                      | EC  | 
| ------------------------------------------- | ----------------------------------------------- | --- | 
| PartnerFunction                             | 取引先機能。オーダーヘッダの取引先機能と、それぞれのビジネスパートナが提案される。必要に応じて変更する（BUYERとSELLERのレコードは変更できない）。参照の場合は変更できない。なお、オーダー明細において、それぞれの取引先機能レコードはオーダーに対して下記の()内の数だけ設定することができる。 <br> 提案される取引先機能: <br> BUYER:買い手(一つ) <br> SELLER:売り手(一つ) <br> CUSTOMER:受注先(一つ)<br> SUPPLIER:仕入先(一つ) <br> MANUFACTURER:製造者(一つ) <br> DELIVERFROM:入出荷元(一つ) <br> DELIVERTO:入出荷先(一つ) <br> LOGI:物流業者(一つ) <br> BILLTO:請求先(一つ)  <br> BILLFROM:請求元(一つ) <br> PAYEE:支払人(一つ) <br> RECEIVER:受取人(一つ) <br> PSPROVIDER:支払決済サービスプロバイダ(一つ) | ✔     | 
| BusinessPartner                             | ビジネスパートナコード。ヘッダの取引先機能に対応するビジネスパートナコードが設定される。BUYER買い手、SELLER売り手、CUSTOMER受注先、SUPPLIER仕入先、BILLTO請求先、BILLFROM請求元、PAYEE支払人、RECEIVER受取人、PSPROVIDER支払決済サービスプロバイダ、はヘッダのレコード・値が設定され変更不可。その他の取引先機能に対するビジネスパートナコードは必要に応じて更新する。参照の場合は変更できない。 | ✔  | 

6-1-1. 2-1-1で保持した[ヘッダの取引先機能データ]を、[DefaultPartner=true]で絞り込み、[Orders Item Partner Data]として、[OrderID, OrderItem, PartnerCounter, PartnerFunction, BusinessPartner, DefaultPartner]の形式で保持する。  

| Target / Processing Type | Key(s)                       | 
| ------------------------ | ---------------------------- | 
| OrderID <br> OrderItem <br> business_partner <br> PartnerCounter <br> PartnerFunction <br> BusinessPartner <br> DefaultPartner <br> / Get and Hold           | DefaultPartner=true          | 
| <b>Table Searched</b>    | <b>Name of the Table</b>     | 
| None                     | None                         | 
| <b>Field Arranged</b>           | <b>Data Type / Number of Digits</b> | 
| OrderID <br> OrderItem <br> PartnerCounter <br> PartnerFunction <br> BusinessPartner <br> DefaultPartner   | int / 16 <br> int / 6 <br> int / 3 <br> string(varchar) / 40 <br>int / 12 <br> bool (tinyint) / 1      | 
| <b>Single Record or Array</b> | <b>Memo</b>                         | 
| Array                    |                              |

7. Orders Item Partner Plant
7-1. ヘッダ取引先プラントデータのコピー  

| Property             | Description                                                                                                                                                            | EC  | 
| -------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
| PartnerFunction      | 取引先機能。ヘッダパートナの取引先機能のうち次のものが設定される。変更不可。<br> MANUFACTURER:製造者 <br> DELIVERFROM:入出荷元 <br> DELIVERTO:入出荷先   | ✔  | 
| BusinessPartner      | ビジネスパートナコード。ヘッダパートナの上記取引先機能に対してヘッダパートナのビジネスパートナが設定される。変更不可。                                                 | ✔  | 
| Plant                | プラント。ビジネスパートナマスタの得意先取引先プラントデータまたは仕入先取引先プラントデータからプラントが設定される。必要に応じて変更する。参照の場合は変更できない。 | ✔  | 

7-1-1. 4-1-1で保持した[ヘッダの取引先プラントデータ]と、6-1-1で保持した[ヘッダの取引先機能データ]を、[OrderID, OrderItem, PartnerCounter]をキーとして結合する。  

7-1-2. [7-1-1で結合したデータ]を、[DefaultPlant=true]で絞り込み、[Orders Item Partner Plant Data]として、[OrderID, OrderItem, PartnerCounter, PartnerFunction, BusinessPartner, DefaultPartner, PlantCounter, Plant, DefaultPlant]の形式で保持する。  

| Target / Processing Type | Key(s)                       | 
| ------------------------ | ---------------------------- | 
| OrderID <br> OrderItem <br> business_partner <br> PartnerCounter <br> PartnerFunction <br> BusinessPartner <br> DefaultPartner <br> PlantCounter <br> Plant <br> DefaultPlant <br> / Get and Hold           | DefaultPlant=true            | 
| <b>Table Searched</b>    | <b>Name of the Table</b>     | 
| None                     | None                         | 
| <b>Field Arranged</b>           | <b>Data Type / Number of Digits</b> | 
| OrderID <br> OrderItem <br> PartnerCounter <br> PartnerFunction <br> BusinessPartner <br> DefaultPartner <br> PlantCounter <br> Plant <br> DefaultPlant    | int / 16 <br> int / 6 <br> int / 3 <br> string(varchar) / 40 <br> int / 12 <br> bool (tinyint) / 1 <br> int / 3 <br> string(varchar) / 4 <br> bool (tinyint) / 1   | 
| <b>Single Record or Array</b> | <b>Memo</b>                         | 
| Array                    |                              | 

8. Orders Item Pricing Element 
次の補助機能の開発を行う。
8-1. 価格マスタデータの取得(入力ファイルの[ConditionAmount]がnullである場合)  

| Property                  | Description                                                                                                                                                           | EC  | 
| ------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
| ConditionRecord           | 条件レコード。価格マスタのキー。品目・受注先/仕入先・流通チャネルの組合せにより、価格マスタマスタから自動設定される。変更不可。参照の場合は変更できない。             |     | 
| ConditionSequentialNumber | 条件連続番号。価格マスタのキー。品目・受注先/仕入先・流通チャネルの組合せにより、価格マスタマスタから自動設定される。変更不可。参照の場合は変更できない。             |     | 
| ConditionType             | 条件タイプ。価格マスタのキー。品目・受注先/仕入先・流通チャネルの組合せにより、価格マスタマスタから自動設定される。変更不可。                                         |     | 
| ConditionRateValue        | 条件レート値。価格マスタのキー。品目・受注先/仕入先・流通チャネルの組合せにより、価格マスタマスタから自動設定される。必要に応じて変更する。参照の場合は変更できない。 |     | 

8-1-1. 入力ファイルのItem Pricing Elementの[ConditionAmount]がnullである場合、次の処理の通り、価格計算を行う。  

8-1-2. [BusinessPartner=入力ファイルのbusiness_partner, 入力ファイルのProduct, Customer=入力ファイルのBuyer(1-0の処理結果が”Seller”の場合)、または、Supplier=入力ファイルのSeller(1-0の処理結果が”Buyer”の場合)]をキーとして、 ConditionValidityEndDate≧1-8でセットしたPricingDate、ConditionValidityStartDate≦1-8でセットしたPricingDate、をキーとして、価格マスタの[ConditionRecord, ConditionSequentialNumber, ConditionType, ConditionRateValue]を検索して値を保持する。  

| Target / Processing Type                           | Key(s)                                        | 
| -------------------------------------------------- | --------------------------------------------- | 
| ConditionRecord <br> ConditionSequentialNumber <br> ConditionType <br> ConditionRateValue <br> / Get and Hold  | business_partner(入力ファイル) <br> Product(入力ファイル) <br> Customer入力ファイルのBuyer、または、 <br> Supplier=入力ファイルのSeller <br> ConditionValidityEndDate≧PricingDate(1-8でセット)<br>ConditionValidityEndDate≦PricingDate(1-8でセット) | 
| <b>Table Searched</b>                              | <b>Name of the Table</b>                      | 
| Price Master SQL Price Master Data                 | data_platform_price_master_ price_master_data | 
| <b>Field Searched</b>                              | <b>Data Type / Number of Digits</b>           | 
| ConditionRecord <br> ConditionSequentialNumber<br> ConditionType <br> ConditionRateValue   | int / 12 <br> int / 2 <br> string(varchar) / 4 <br> float / 13                                         | 
| <b>Single Record or Array</b>                      | <b>Memo</b>                                  | 
| Array                                              |                                               | 

8-2. 価格の計算(入力ファイルの[ConditionAmount]がnullである場合)  

| Property                | Description                                                                                                  | EC  | 
| ----------------------- | ------------------------------------------------------------------------------------------------------------ | --- | 
| OrderQuantityInBaseUnit | オーダー数量(基本数量単位)。参照の場合は変更できない。                                                       |     | 
| ConditionQuantity       | 条件数量。明細の数量がコピーされる。変更不可。                                                               |     | 
| ConditionQuantityUnit   | 条件数量単位。明細の数量単位がコピーされる。変更不可。                                                       |     | 
| ConditionAmount         | 条件金額。条件レート値と条件数量を乗じた値として計算される。必要に応じて変更する。参照の場合は変更できない。 |      |	

8-2-1. 入力ファイルのItemの[OrderQuantityInBaseUnit]をItem Pricing Elementの[ConditionQuantity]として保持する。  

8-2-2. 8-1で取得した[ConditionRateValue]の小数点以下の桁数をカウントする。  

8-2-3. 8-1で取得した[ConditionRateValue]と、8-2-1で保持した[ConditionQuantity]を掛け算する。  

8-2-4. 8-2-3で求めた値の小数点以下部分を、8-2-2でカウントした桁数に四捨五入して、[ConditionAmount]の値を求め、セットする。  
8-2-5. [ConditionIsManuallyChanged]にfalseをセットする。  

8-3. 価格の保持(入力ファイルの[ConditionAmount]がnullでない場合)  

| Property                   | Description                                                                                      | EC  | 
| -------------------------- | ------------------------------------------------------------------------------------------------ | --- | 
| ConditionIsManuallyChanged | 条件のマニュアル変更の有無。マニュアル変更した場合、trueが設定される。参照の場合は変更できない。 |      |	

8-3-1. 入力ファイルの[ConditionAmount]がnullでない場合、[ConditionIsManuallyChanged]にtrueをセットする。  

8-4. PricingProcedureCounter  

| Property                | Description                                                | EC  | 
| ----------------------- | ---------------------------------------------------------- | --- | 
| PricingProcedureCounter | 価格手続カウンタ。1～999の範囲で自動採番される。変更不可。 |     |	

8-4-1. 8-1-2で取得した明細の数だけ、1～Nまで整数の番号を付与して配列に保持する。  

8-5.消費税率の取得  

8-5-1. [5-1で取得したTaxCode, Country=”JP”, ValidityEndDate≧システム日付, ValidityStartDate≦システム日付]、をキーとして、消費税率データのTaxRateを検索して値を保持する。  

| Target / Processing Type        | Key(s)                               | 
| ------------------------------- | ------------------------------------ | 
| TaxRate / Get and Hold          | TaxCode <br> Country=”JP” <br> ValidityEndDate≧システム日付 <br> ValidityStartDate≦システム日付 | 
| <b>Table Searched</b>           | <b>Name of the Table</b>             | 
| Tax Code SQL Tax Rate Data      | data_platform_tax_code_tax_rate_data | 
| <b>Field Searched</b>           | <b>Data Type / Number of Digits</b>  | 
| TaxRate                         | float / 6                            | 
| <b>Single Record or Array</b>   | <b>Memo</b>                         | 
| Single Record                   |                                      | 

9. Orders Address
9-1. 住所マスタからの住所データの取得([PostalCode, LocalRegion, Country, District, StreetName, CityName, Building, Floor, Room])  

| Property    | Description                                                                                                                                 | EC  | 
| ----------- | ------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
| AddressID   | 住所ID。ビジネスパートナマスタまたは見積/引合から設定される。オーダーでマニュアルで住所を更新する場合、住所IDが新たに設定される。変更不可。 |     | 
| PostalCode  | 郵便番号。住所IDまたは見積/引合から設定される。必要に応じて変更する。                                                                       | ✔  | 
| LocalRegion | ローカル地域。住所IDまたは見積/引合から設定される。必要に応じて変更する。                                                                   | ✔  | 
| Country     | 国。住所IDまたは見積/引合から設定される。必要に応じて変更する。                                                                             | ✔  | 
| District    | ディストリクト。住所IDまたは見積/引合から設定される。必要に応じて変更する。                                                                 | ✔  | 
| StreetName  | 地名・番地。住所IDまたは見積/引合から設定される。必要に応じて変更する。                                                                     |     | 
| CityName    | 市区町村。住所IDまたは見積/引合から設定される。必要に応じて変更する。                                                                       |     | 
| Building    | 建物名。住所IDまたは見積/引合から設定される。必要に応じて変更する。                                                                         |     | 
| Floor       | 階数。住所IDまたは見積/引合から設定される。必要に応じて変更する。                                                                           |     | 
| Room        | 部屋番号。住所IDまたは見積/引合から設定される。必要に応じて変更する。                                                                       |     | 

9-1-1. 2-2-2で取得した[AddressID]をコピーしてセットする。  

9-1-2. [9-1-1でセットした AddressID]と、[ValidityEndDate≧システム日付]をキーとして、対象の住所データの[PostalCode, LocalRegion, Country, District, StreetName, CityName, Building, Floor, Room]、を検索しセットする。  

| Target / Processing Type | Key(s) |
| ------------------------ | ---------------------------------- | 
| 下記Field Searched <br>/ Get and Set |	[AddressID] <br> [ValidityEndDate≧システム日付] |
| <b>Table Searched</b>    | <b>Name of the Table</b>           | 
| Address SQL Address Data | data_platform_address_address_data | 
| <b>Field Searched</b>    | <b>Data Type / Number of Digits</b>| 
| PostalCode <br> LocalRegion <br> Country <br> District <br> StreetName <br> CityName <br> Building <br>Floor <br> Room                     | string(varchar) / 10 <br> string(varchar) / 3  <br> string(varchar) / 3 <br> string(varchar) / 6  <br> string(varchar) / 200 <br> string(varchar) / 200 <br> string(varchar) / 100 <br> int / 4 <br> int / 8                  | 
| <b>Single Record or Array</b> | <b>Memo</b>                               | 
| Array                    |                                    | 

9-2. AddressIDの登録(ユーザーが任意の住所を入力ファイルで指定した場合)  

| Property  | Description                                                                                                                                 | EC  | 
| --------- | ------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
| AddressID | 住所ID。ビジネスパートナマスタまたは見積/引合から設定される。オーダーでマニュアルで住所を更新する場合、住所IDが新たに設定される。変更不可。 |     | 


9-2-1. 入力ファイルの[ValidityEndDate, ValidityStartDate, PostalCode, LocalRegion, Country, GlobalRegion, TimeZone, StreetName, CityName]が、ブランクでなく、かつnullでない場合、住所IDの登録を行う。  

9-2-2. AddressIDについては、ServiceLabel=”ADDRESS_ID”, FieldNameWithNumberRange=当処理のProperty(=”AddressID”)をキーとして、対象のNumber Range Latest NumberのLatestNumberを検索し保持する。  

| Target / Processing Type                                  | Key(s)                                        | 
| --------------------------------------------------------- | --------------------------------------------- | 
| ServiceLabel FieldNameWithNumberRange <br> LatestNumber <br> / Get and Hold   | ServiceLabel=”ADDRESS_ID”  <br> FieldNameWithNumberRange =”AddressID”(当処理のProperty) | 
| <b>Table Searched</b>                                     | <b>Name of the Table</b>                      | 
| Number Range SQL Latest Number Data                       | data_platform_number_range_latest_number_data | 
| <b>Field Searched</b>                                     | <b>Data Type / Number of Digits</b>           | 
| ServiceLabel <br> FieldNameWithNumberRange <br> LatestNumber  | string(varchar) / 50 <br> string(varchar) / 100 <br> int / 10                                                  | 
| <b>Single Record or Array</b>                             | <b>Memo</b>                                  | 
| Single Record                                             |                                               | 

9-2-3. 保持されたLatestNumberに1を足したものをAddressIDにセットする。  

9-2-4. アドレスマスタレコードの生成  
次の処理の通り、アドレスマスタレコードを生成し、アドレスマスタに登録する。  

| 生成する項目      | 生成する値等のロジック            | 
| ----------------- | --------------------------------- | 
| AddressID         | 9-2-3でセットしたAddressID。      | 
| ValidityEndDate   | 入力ファイルのValidityEndDate。   | 
| ValidityStartDate | 入力ファイルのValidityStartDate。 | 
| PostalCode        | 入力ファイルのPostalCode。        | 
| LocalRegion       | 入力ファイルのLocalRegion。       | 
| Country           | 入力ファイルのCountry。           | 
| GlobalRegion      | 入力ファイルのGlobalRegion。      | 
| TimeZone          | 入力ファイルのTimeZone。          | 
| District          | 入力ファイルのDistrict。          | 
| StreetName        | 入力ファイルのStreetName。        | 
| CityName          | 入力ファイルのCityName。          | 
| Builiding         | 入力ファイルのBuiliding。         | 
| Floor             | 入力ファイルのFloor。             | 
| Room              | 入力ファイルのRoom。              |   


| Target / Processing Type                                                               | Key(s)                                | 
| -------------------------------------------------------------------------------------- | ------------------------------------- | 
| アドレスマスタの各項目(上記表の通り) / Creates                                         | AddressID(9-2-3でセットしたAddressID) | 
| <b>Table Searched</b>                                                                  | <b>Name of the Table</b>              | 
| None                                                                                   | None                                  | 
| <b>Field Searched</b>                                                                  | <b>Data Type / Number of Digits</b>   | 
| アドレスマスタの各項目(上記表の通り)                                                   | 上記表の通り                          | 
| <b>Single Record or Array</b>                                                          | <b>Memo</b>                          | 
| Single Record                                                                          | data_platform_api_address_creates <br> ※本処理を行ってからOrders Createsの最終処理であるDBへのレコード登録を行う必要がある。 |   


10. Headerの集計項目等の計算とセット
10-1. TotalNetAmount  
<計算元:Orders Item(5-20)>  

| Property                | Description                                                            | EC  | 
| ----------------------- | ---------------------------------------------------------------------- | --- | 
| NetAmount               | 正味金額。オーダー数量と条件価格とを乗じた金額が計算される。変更不可。 |     | 
| OrderQuantityInBaseUnit | オーダー数量(基本数量単位)。参照の場合は変更できない。                 |      |	

<計算:Orders Header>  

| Property       | Description                                                                                                                      | EC  | 
| -------------- | -------------------------------------------------------------------------------------------------------------------------------- | --- | 
| TotalNetAmount | 合計正味金額。消費税を除いた合計正味金額が自動計算される。必要に応じて変更する。明細の正味金額の合計との整合性がチェックされる。 |    |	

10-1-1. 5-20で計算・セットされた[OrderItem, NetAmount]の金額を、全てのOrderItemに対して合計計算する。計算された金額を、[TotalNetAmount]に保持する。  

10-1-2. 入力ファイルの[TotalNetAmount]がnullでない場合、10-1-1の[TotalNetAmount]と入力ファイルの[TotalNetAmount ]を比較して、一致しない場合、エラーメッセージを出力して終了する。一致する場合、入力ファイルの[TotalNetAmount]を[TotalNetAmount]に保持する。  

10-2. TotalTaxAmount
<計算元:Orders Item(5-21)>  

| Property  | Description                                                                                                                                                                                                                                                                                                             | EC  | 
| --------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
| TaxAmount | 税金額。明細の税分類が「1:納税義務」の場合、適切な税率にもとづいて消費税金額が自動計算される。明細の税分類が「0:非課税」の場合、税金額はゼロとなる。明細の税分類が「1:納税義務」の場合、ユーザが税金額を入力することはできるが、理論値の消費税額と2通貨単位以上の差が出る場合、エラーとなる。参照の場合は変更できない。 |     | 

<計算:Orders Header>  

| Property       | Description                                                                                                        | EC  | 
| -------------- | ------------------------------------------------------------------------------------------------------------------ | --- | 
| TotalTaxAmount | 合計消費税額。合計消費税額が自動計算される。必要に応じて変更する。明細の消費税額の合計との整合性がチェックされる。 |   |	

10-2-1. 5-21で計算・セットされた[OrderItem, TaxAmount]の金額を、全てのOrderItemに対して合計計算する。計算された金額を、[TotalTaxAmount]に保持する。  

10-2-2. 入力ファイルの[TotalTaxAmount]がnullでない場合、10-2-1の[TotalTaxAmount]と入力ファイルの[TotalTaxAmount ]を比較して、一致しない場合、エラーメッセージを出力して終了する。一致する場合、入力ファイルの[TotalTaxAmount]を[TotalTaxAmount]に保持する。  

10-3. TotalGrossAmount
<計算元:Orders Item(5-22)>  

| Property    | Description                                                                                                                                                                                   | EC  | 
| ----------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
| GrossAmount | 総額。消費税を含んだ総額が自動計算される。必要に応じて変更する。変更された金額が正しいかどうかチェックされる。理論値と2通貨単位以上の差額がある場合はエラーとなる。参照の場合は変更できない。 |     | 

<計算:Orders Header>  

| Property         | Description                                                                                                          | EC  | 
| ---------------- | -------------------------------------------------------------------------------------------------------------------- | --- | 
| TotalGrossAmount | 合計総額。消費税を含んだ合計総額が自動計算される。必要に応じて変更する。明細の総額の合計との整合性がチェックされる。 |     | 

10-3-1. 5-22で計算・セットされた[OrderItem, GrossAmount]の金額を、全てのOrderItemに対して合計計算する。計算された金額を、[TotalGrossAmount]に保持する。  

10-3-2. 入力ファイルの[TotalGrossAmount]がnullでない場合、10-3-1の[TotalGrossAmount]と入力ファイルの[TotalGrossAmount ]を比較して、一致しない場合、エラーメッセージを出力して終了する。一致する場合、入力ファイルの[TotalGrossAmount]を[TotalGrossAmount]に保持する。  

10. 参照登録
10-0. ServiceLabelのセット  
0-2-2の通り、入力ファイルのReferenceDocumentをキーとして、ServiceLabelを取得・セットし、各ServiceLabelによって、対象のレコード・値を検索する。また、参照登録が発動した場合には、変更が不可となる制御をかける。つまり、inputファイルに値があっても、その値は採用されない。  

10-1.見積参照  
ServiceLabelが”QUOTATIONS”である場合、見積参照となる。  

10-1-1. 見積ヘッダの取得  
ReferenceDocumentをキーとして、以下の項目を取得する。  
[見積ヘッダ]  

| Target / Processing Type | Key(s)                               | 
| ------------------------ | ------------------------------------ | 
| 下記Field Searched <br>/ Get and Set            | ReferenceDocument                    | 
| <b>Table Searched</b>    | <b>Name of the Table</b>             | 
| Quotations SQL Header    | data_platform_quotations_header_data | 
| <b>Field Searched</b>    | <b>Data Type / Number of Digits</b>  | 
| Quotation <br> QuotationType <br> Buyer  <br> Seller <br> ContractType <br> VaridityStartDate <br> ValidityEndDate <br> InvoiceScheduleStartDate <br> InvoiceScheduleEndDate <br> TotalNetAmount <br> TotalTaxAmount <br> TotalGrossAmount <br>TransactionCurrency <br> PricingDate <br> PriceDetnExchangeRate <br> RequestedDeliveryDate <br> Incoterms  <br> PaymentTerms <br> PaymentMethod <br> BillingDocumentDate <br> HeaderText               | int / 16 <br> string(varchar) / 3 <br> int / 12 <br> int / 12 <br> string(varchar) / 4 <br> date <br> date <br> date <br> date <br> float / 13 <br>float / 13 <br> float / 13 <br> string(varchar) / 5 <br> date <br> float / 8  <br> date <br> string(varchar) / 4 <br> string(varchar) / 4 <br> string(varchar) / 1 <br> date <br> string(varchar) / 100    | 
| <b>Single Record or Array</b> | <b>Memo</b>                                 | 
| Single Record            | ReferenceDocument = Quotation        |   

[見積ヘッダパートナ]  

| Target / Processing Type          | Key(s)                                       | 
| --------------------------------- | -------------------------------------------- | 
| 下記Field Searched<br> / Get and Set   | ReferenceDocument                            | 
| <b>Table Searched</b>             | <b>Name of the Table</b>                     | 
| Quotations SQL Header Partner     | data_platform_quotations_header_partner_data | 
| <b>Field Searched</b>             | <b>Data Type / Number of Digits</b>          | 
| Quotation <br> PartnerFunction <br> BusinessPartner <br> BusinessPartnerFullName <br> BusinessPartnerName <br> Counrtry <br> Language <br> Currency <br> ExternalDocumentID <br> AddressID                         | int / 16 <br> string(varchar) / 40 <br> int / 12 <br> string(varchar) / 200 <br> string(varchar) / 100 <br> string(varchar) / 3 <br> string(varchar) / 2 <br> string(varchar) / 5 <br> string(varchar) / 40 <br> int / 12                          | 
| <b>Single Record or Array</b>     | <b>Memo</b>                                 | 
| Array                             | ReferenceDocument = Quotation <br> ReferenceDocumentItem = Quotation |   

[見積明細]  

| Target / Processing Type              | Key(s)                                    | 
| ------------------------------------- | ----------------------------------------- | 
| 下記Field Searched <br>/ Get and Set       | ReferenceDocument / ReferenceDocumentItem | 
| <b>Table Searched</b>                 | <b>Name of the Table</b>                  | 
| Quotations SQL Header <br> Quotations SQL Item   | data_platform_quotations_header_data  <br> data_platform_quotations_item_data    | 
| <b>Field Searched</b>                 | <b>Data Type / Number of Digits</b>       | 
| Quotation <br> QuotationItem          | int / 10  <br> int / 6                    | 
| <b>Single Record or Array</b>         | <b>Memo</b>                              | 
| Single Record                         | ReferenceDocument = Quotation  <br> ReferenceDocumentItem = QuotationItem |   

上記の表から分かる通り、入力ファイルのReferenceDocument または ReferenceDocumentItemをキーとして、各ターゲットを参照しセットする。また、参照登録(101)が発動した場合には、変更が不可となる制御をかける。つまり、inputファイルに値があっても、その値は採用されない。  

以下は参照登録の対象である。  

| Property       | Description                                                                                                                                                                    | EC  | 
| -------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | --- | 
| TotalNetAmount | オーダー金額。参照の場合は変更できない。                                                                                                                                       |     | 
| PricingDate    | 価格設定日付。初期値としてはシステム日付が提案される。通常、オーダー日付と同じ日付を入力する。価格設定日付を任意で決めたい場合、その日付を入力する。参照の場合は変更できない。 | 	 |


[オーダー参照]
SRT:販売返品リクエスト
SCR:販売クレジットメモリクエスト
SDR:販売デビットメモリクエスト
PRT:購買返品リクエスト
PDR:購買デビットメモリクエスト
PCR:購買クレジットメモリクエスト

[オーダーヘッダ]  

| Target / Processing Type | Key(s)                           | 
| ------------------------ | -------------------------------- | 
| OrderID / Get and Set    | ReferenceDocument                | 
| <b>Table Searched</b>    | <b>Name of the Table</b>         | 
| Orders SQL Header        | data_platform_orders_header_data | 
| <b>Field Searched</b>    | <b>Data Type / Number of Digits</b>
| OrderID <br> PriceDetnExchangeRate <br> DistributionChannel <br> Division <br> AccountingExchangeRate <br> BPTaxClassification      | int / 10 <br> string(varchar) / 80 <br> string(varchar) / 2 <br> string(varchar) / 2 <br> string(varchar) / 16 <br> string(varchar) / 1      | 
| <b>Single Record or Array</b> | <b>Memo</b>                             | 
| Single Record            | ReferenceDocument = OrderID      | 

[オーダー明細]  

| Target / Processing Type          | Key(s)                                    | 
| --------------------------------- | ----------------------------------------- | 
| OrderID <br> OrderItem <br>/ Get and Set  | ReferenceDocument / ReferenceDocumentItem | 
| <b>Table Searched</b>             | <b>Name of the Table</b>                  | 
| Orders SQL Header <br> Orders SQL Item | data_platform_orders_header_data <br> data_platform_orders_item_data    | 
| <b>Field Searched</b>             | <b>Data Type / Number of Digits</b>       | 
| OrderID <br> OrderItem             | int / 16 <br> int / 6                           | 
| <b>Single Record or Array</b>     | <b>Memo</b>                              | 
| Single Record                     | ReferenceDocument = OrderID <br> ReferenceDocumentItem = OrderItem | 

上記の表から分かる通り、入力ファイルのReferenceDocument または ReferenceDocumentItemをキーとして、各ターゲットを参照しセットする。また、参照登録(101)が発動した場合には、変更が不可となる制御をかける。つまり、inputファイルに値があっても、その値は採用されない。  

以下は参照登録の対象である。  

| Property               | Description                                                                                                                    | EC  | 
| ---------------------- | ------------------------------------------------------------------------------------------------------------------------------ | --- | 
| PriceDetnExchangeRate  | 価格決定のための為替レート。必要な場合、為替レートを入力する。参照の場合は変更できない。                                       |     | 
| DistributionChannel    | 流通チャネル。次から選択する。または参照されて設定される。参照の場合は変更できない。<br> DS:直接販売、EC:EC販売 | ✔                                                                                                                             | 
| Division               | 製品部門。次の事業領域から選択する。または参照されて設定される。参照の場合は変更できない。 <br> MF:製造 <br> RT:小売 <br> TR:卸売                | ✔                                                                                                                             | 
| AccountingExchangeRate | 会計為替レート。外貨建て取引の場合、為替レートを入力する。参照の場合は変更できない。                                           |     | 
| BPTaxClassification    | ビジネスパートナ税分類。ビジネスパートナマスタの税分類が以下の値が提案される。必要に応じて変更する。参照の場合は変更できない。<br> 0:非課税 <br> 1:納税義務             |                                                                                                                                | 

10-2. 引合参照  
ServiceLabelが”INQUIRIES”である場合、引合参照となる。  

10-3. 購買依頼参照  
ServiceLabelが”PURCHASE_REQUISITION”である場合、引合参照となる。  

10-4. オーダー参照  
ServiceLabelが”ORDERS”である場合、オーダー参照となる。  

11. 数量単位変換(OrderQuantityInIssuingUnit / OrderQuantityInReceivingUnit)
入力ファイルの[OrderItemCategory=”INVP”の明細]に対して、本処理を実行する。  
11-1. 数量単位変換実行の是非の判定  
入力ファイルの[基本数量単位(BaseUnit)]と、[5-5-1でセットしたOrderIssuingUnit]または(および)[5-5-1でセットしたOrderReceivingUnit]が異なる場合に、数量単位変換実行の次の一連の処理を行う。  

11-2. [BaseUnit]と、[5-5-1でセットしたOrderIssuingUnit]または(および)[ 5-5-1でセットしたOrderReceivingUnit]をキーとして、数量単位変換マスタの[QuantityUnitFrom, QuantityUnitTo, ConversionCoefficient]を検索して保持する。  

| Target / Processing Type                              | Key(s)                                                                | 
| ----------------------------------------------------- | --------------------------------------------------------------------- | 
| ConversionCoefficient <br> / Get and Hold             | BaseUnit <br> OrderIssuingUnitまたは(および) <br> OrderReceivingUnit   | 
| <b>Table Searched</b>                                 | <b>Name of the Table</b>                                              | 
| Quantity Unit Conversion SQL Quantity Unit Conversion | data_platform_quantity_unit_conversion_ quantity_unit_conversion_data | 
| <b>Field Searched</b>                                 | <b>Data Type / Number of Digits</b>                                   | 
| QuantityUnitFrom <br> QuantityUnitTo <br> ConversionCoefficient | string(varchar) / 3 <br> string(varchar) / 3 <br> float /24                                            | 
| <b>Single Record or Array</b>                         | <b>Memo</b>                                                          | 
| Array                                                 |                                                                       | 

11-3. 11-2で保持した[OrderIssuingUnitに対応するConversionCoefficient]、または(および) [OrderReceivingUnitに対応するConversionCoefficient]を、入力ファイルのOrderQuantityInBaseUnitに乗算して、[OrderQuantityInIssuingUnit]と[OrderQuantityInReceivingUnit]を計算する。  

99-1. CreationDate  

| Property     | Description              | EC  | 
| ------------ | ------------------------ | --- | 
| CreationDate | 作成日。自動生成される。 |   |	

99-1-1. CreationDateにシステム日付を設定する。  

99-2. LastChangeDate  

| Property       | Description                    | EC  | 
| -------------- | ------------------------------ | --- | 
| LastChangeDate | 最終更新日時。自動生成される。 |     | 

99-2-1. LastChangeDateにシステム日付を設定する。  

100. 仕入先データから取得した値と得意先データから取得した値の比較
2-2-1で保持したCurrency、1-1-1で保持したIncoterms、PaymentTerms、PaymentMethodについて、得意先データから取得した値と仕入先データから取得した値を比較する。  

100-1. Currency  

| Property | Description                                                                                                                      | EC  | 
| -------- | -------------------------------------------------------------------------------------------------------------------------------- | --- | 
| Currency | 通貨コード。取引先機能に対応するビジネスパートナマスタの通貨コードが設定される。必要に応じて変更する。参照の場合は変更できない。 |     | 

100-1-1. Currencyの得意先データから取得した値と仕入先データから取得した値を比較し、一致した場合は処理をそのまま実行し、一致しなかった場合はエラーメッセージを出力し終了する。  

<100-1-1まとめ>  

| Data From Customer | Data From Supplier | 比較結果 | 比較後の処理 | 
| ------------------ | ------------------ | -------- | ------------ | 
| Currency           | Currency           | ＝       | そのまま実行 | 
| Currency           | Currency           | ≠       | エラー       | 


100-2. Incoterms、PaymentTerms、PaymentMethod  

| Property      | Description                                                                                                                                                                                                                                          | EC  | 
| ------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
| Incoterms     | インコタームズ。貿易取引を行う際、輸送タイプに応じてインコタームズを指定する。ビジネスパートナマスタの得意先データまたは仕入先データのインコタームズが提案される。必要に応じて、インコタームズマスタから選択して指定する。参照の場合は変更できない。 | ✔  | 
| PaymentTerms  | 支払条件。ビジネスパートナマスタの得意先データまたは仕入先データの支払条件が提案される。変更の必要があれば、支払条件マスタから選択して指定する。参照の場合は変更できない。                                                                           | ✔  | 
| PaymentMethod | 支払方法。ビジネスパートナマスタの得意先データまたは仕入先データの支払方法が提案される。変更の必要があれば、支払方法マスタから選択して指定する。参照の場合は変更できない。                                                                           | ✔  | 

100-2-1. Incoterms、PaymentTerms、PaymentMethodの得意先データから取得した値と仕入先データから取得した値を比較し、一致した場合は処理をそのまま実行し、一致しなかった場合はエラーメッセージを出力して終了する。  

<100-2-1まとめ>  

| Data From Customer | Data From Supplier | 比較結果 | 比較後の処理 | 
| ------------------ | ------------------ | -------- | ------------ | 
| Incoterms          | Incoterms          | ＝       | そのまま実行 | 
| Incoterms          | Incoterms          | ≠       | エラー       | 
| PaymentTerms       | PaymentTerms       | ＝       | そのまま実行 | 
| PaymentTerms       | PaymentTerms       | ≠       | エラー       | 
| PaymentMethod      | PaymentMethod      | ＝       | そのまま実行 | 
| PaymentMethod      | PaymentMethod      | ≠       | エラー       | 

200. 登録する値の整列とセット
200-1. Orders Headerデータの整列とセット  
次のOrders Headerデータを整列しDB/Tableへの登録値としてセットする。      
<対象テーブル: data_platform_orders_header_data>  

| <b>Field Created</b>                   | <b>Data Type / Number of Digits</b> | <b>Specification</b>                            | 
| ------------------------------- | ---------------------------- | ---------------------------------------- | 
| OrderID(※Key)                  | int / 16                     | 1-2-2でセットされたOrderID               | 
| OrderDate                       | date                         | 入力ファイルのOrderDate                  | 
| OrderType                       | string(varchar) / 3          | 入力ファイルのOrderType                  | 
| Buyer                           | int / 12                     | 入力ファイルのBuyer                      | 
| Seller                          | int / 12                     | 入力ファイルのSeller                     | 
| CreationDate                    | date                         | 99-1でセットしたCreationDate             | 
| LastChangeDate                  | date                         | 99-2でセットしたLastChangeDate           | 
| ContractType                    | string(varchar) / 4          | 入力ファイルのContractType               | 
| VaridityStartDate               | Date                         | 入力ファイルのVaridityStartDate          | 
| VaridityEndDate                 | Date                         | 入力ファイルのVaridityEndDate            | 
| InvoiceScheduleStartDate        | Date                         | 入力ファイルのInvoiceScheduleStartDate   | 
| InvoiceScheduleEndDate          | Date                         | 入力ファイルのInvoiceScheduleEndDate     | 
| TotalNetAmount                  | float / 13                   | 10-1で計算・セットしたTotalNetAmount     | 
| TotalTaxAmount                  | float / 13                   | 10-2で計算・セットしたTotalTaxAmount     | 
| TotalGrossAmount                | float / 13                   | 10-3で計算・セットしたTotalGrossAmount   | 
| OverallDeliveryStatus           | string(varchar) / 2          | “PP”をセットする。                       | 
| TotalBlockStatus                | bool(tinyint) / 1            | falseをセットする。                        | 
| OverallOrdReltdBillgStatus      | string(varchar) / 2          | “PP”をセットする。                       | 
| OverallDocReferenceStatus       | string(varchar) / 2          | 1-7でセットしたOverallDocReferenceStatus | 
| TransactionCurrency             | string(varchar) / 5          | 1-11で保持したTransactionCurrency        | 
| PricingDate                     | Date                         | 1-8でセットしたPricingDate               | 
| PriceDetnExchangeRate           | Date                         | 1-9でセットしたPriceDetnExchangeRate     | 
| RequestedDeliveryDate           | Date                         | 入力ファイルのRequestedDeliveryDate      | 
| HeaderCompleteDeliveryIsDefined | bool(tinyint) / 1            | falseをセットする。                        | 
| HeaderBillingBlockReason        | bool(tinyint) / 1            | falseをセットする。                        | 
| DeliveryBlockReason             | bool(tinyint) / 1            | falseをセットする。                        | 
| Incoterms                       | string(varchar) / 3          | 1-1でセットしたIncoterms                 | 
| PaymentTerms                    | string(varchar) / 4          | 1-1でセットしたPaymentTerms              | 
| PaymentMethod                   | string(varchar) / 1          | 1-1でセットしたPaymentMethod             | 
| ReferenceDocument               | int / 16                     | 入力ファイルのReferenceDocument          | 
| ReferenceDocumentItem           | int / 6                      | 入力ファイルのReferenceDocumentItem      | 
| BPAccountAssignmentGroup        | string(varchar) / 2          | 1-1でセットしたBPAccountAssignmentGroup  | 
| AccountingExchangeRate          | float / 8                    | 入力ファイルのAccountingExchangeRate     | 
| BillingDocumentDate             | Date                         | 1-3でセットしたBillingDocumentDate       | 
| IsExportImportDelivery          | bool(tinyint) / 1            |                                          | 
| HeaderText                      | string(varchar) / 200        | 入力ファイルのHeaderText                 | 

200-2. Orders Header Partnerデータの整列とセット
次のOrders Header Partnerデータを整列しDB/Tableへの登録値としてセットする。  
<対象テーブル: data_platform_orders_header_partner_data>  

| <b>Field Created</b>           | <b>Data Type / Number of Digits</b> | <b>Specification</b>                                            | 
| ----------------------- | ---------------------------- | -------------------------------------------------------- | 
| OrderID(※Key)          | int / 16                     | 1-2-2でセットされたOrderID                               | 
| PartnerFunction(※Key)  | string(varchar) / 40         | 2-1-1で取得したPartnerFunction                           | 
| BusinessPartner(※Key)  | int / 12                     | 2-1-1で取得したPartnerFunctionBusinessPartner            | 
| BusinessPartnerFullName | string(varchar) / 200        | 2-2-1で取得したBusinessPartnerFullName                   | 
| BusinessPartnerName     | string(varchar) / 100        | 2-2-1で取得したBusinessPartnerName                       | 
| Organization            | string(varchar) / 6          |                                                          | 
| Country                 | string(varchar) / 3          | 2-2-1で取得したCountry                                   | 
| Language                | string(varchar) / 2          | 2-2-1で取得したLanguage                                  | 
| Currency                | string(varchar) / 5          | 2-2-1で取得したCurrency                                  | 
| AddressID               | int / 12                     | 2-2-1で取得したAddressIDまたは9-2-3でセットしたAddressID | 

200-3. Orders Header Partner Plantデータの整列とセット
次のOrders Header Partner Plantデータを整列しDB/Tableへの登録値としてセットする。  
<対象テーブル: data_platform_orders_header_partner_plant_data>  

| <b>Field Created</b>          | <b>Data Type / Number of Digits</b> | <b>Specification</b>                                 | 
| ---------------------- | ---------------------------- | --------------------------------------------- | 
| OrderID(※Key)         | int / 16                     | 1-2-2でセットされたOrderID                    | 
| PartnerFunction(※Key) | string(varchar) / 40         | 2-1-1で取得したPartnerFunction                | 
| BusinessPartner(※Key) | int / 12                     | 2-1-1で取得したPartnerFunctionBusinessPartner | 
| Plant(※Key)           | string(varchar) / 4          | 4-1-1で取得したPlant                          | 

200-4. Orders Header Partner Contactデータの整列とセット
次のOrders Header Partner Contactデータを整列しDB/Tableへの登録値としてセットする。  
<対象テーブル: data_platform_orders_header_partner_contact_data>  

| <b>Field Created</b>          | <b>Data Type / Number of Digits</b> | <b>Specification</b>                                 | 
| ---------------------- | ---------------------------- | --------------------------------------------- | 
| OrderID(※Key)         | int / 16                     | 1-2-2でセットされたOrderID                    | 
| PartnerFunction(※Key) | string(varchar) / 40         | 2-1-1で取得したPartnerFunction                | 
| BusinessPartner(※Key) | int / 12                     | 2-1-1で取得したPartnerFunctionBusinessPartner | 
| ContactID(※Key)       | int / 4                      | 3-1-1で取得したContactID                      | 
| ContactPersonName      | string(varchar) / 100        | 3-1-1で取得したContactPersonName              | 
| EmailAddress           | string(varchar) / 200        | 3-1-1で取得したEmailAddress                   | 
| PhoneNumber            | string(varchar) / 100        | 3-1-1で取得したPhoneNumber                    | 
| MobilePhoneNumber      | string(varchar) / 100        | 3-1-1で取得したMobilePhoneNumber              | 
| FaxNumber              | string(varchar) / 100        | 3-1-1で取得したFaxNumber                      | 
| ContactTag1            | string(varchar) / 40         | 3-1-1で取得したContactTag1                    | 
| ContactTag2            | string(varchar) / 40         | 3-1-1で取得したContactTag2                    | 
| ContactTag3            | string(varchar) / 40         | 3-1-1で取得したContactTag3                    | 
| ContactTag4            | string(varchar) / 40         | 3-1-1で取得したContactTag4                    | 

200-11. Orders Item Partnerデータの整列とセット
次のOrders Item Partnerデータを整列しDB/Tableへの登録値としてセットする。    
<対象テーブル: data_platform_orders_item_partner_data>    

| <b>Field Created</b>          | <b>Data Type / Number of Digits</b> | <b>Specification</b>                  | 
| ---------------------- | ---------------------------- | ------------------------------ | 
| OrderID(※Key)         | int / 16                     | 6-1-1で保持したOrderID         | 
| OrderItem(※Key)       | int / 6                      | 6-1-1で保持したOrderItem       | 
| PartnerFunction(※Key) | string(varchar) / 40         | 6-1-1で保持したPartnerFunction | 
| BusinessPartner(※Key) | int / 12                     | 6-1-1で保持したBusinessPartner | 

200-12. Orders Item Partner Plantデータの整列とセット
次のOrders Item Partner Plantデータを整列しDB/Tableへの登録値としてセットする。    
<対象テーブル: data_platform_orders_item_partner_plant_data>  

| <b>Field Created</b>          | <b>Data Type / Number of Digits</b> | <b>Specification</b>                  | 
| ---------------------- | ---------------------------- | ------------------------------ | 
| OrderID(※Key)         | int / 16                     | 7-1-2で保持したOrderID         | 
| OrderItem(※Key)       | int / 6                      | 7-1-2で保持したOrderItem       | 
| PartnerFunction(※Key) | string(varchar) / 40         | 7-1-2で保持したPartnerFunction | 
| BusinessPartner(※Key) | int / 12                     | 7-1-2で保持したBusinessPartner | 
| Plant(※Key)           | string(varchar) / 4          | 7-1-2で保持したPlant           | 

200-13. Orders Item Pricing Elementデータの整列とセット
次のOrders Item Pricing Elementデータを整列しDB/Tableへの登録値としてセットする。  

<対象テーブル: data_platform_item_pricing_element_data>  

| <b>Field Created</b>                   | <b>Data Type / Number of Digits</b> | <b>Specification</b>                                      | 
| ------------------------------- | ---------------------------- | -------------------------------------------------- | 
| OrderID(※Key)                  | int / 16                     | 7-1-2で保持したOrderID                             | 
| OrderItem(※Key)                | int / 6                      | 7-1-2で保持したOrderItem                           | 
| PricingProcedureCounter(※Key)  | int / 3                      | 8-4で保持したPricingProcedureCounter               | 
| ConditionRecord                 | int / 12                     | 8-1で保持したConditionRecord                       | 
| ConditionRecordSequentialNumber | int / 3                      | 8-1で保持したConditionRecordSequentialNumber       | 
| ConditionType                   | string(varchar) / 4          | 8-1で保持したConditionType                         | 
| ConditionRateValue              |                              | 8-1で保持したConditionRateValue                    | 
| PricingDate                     | date                         | 5-28で保持したPricingDate                          | 
| ConditionCurrency               | string(varchar) / 5          | 1-11で保持したTransactionCurrency                  | 
| ConditionQuantity               | float / 15                   | 8-2で保持したConditionQuantity <br>または <br> 入力ファイルのConditionQuantity | 
| ConditionQuantityUnit           | string(varchar) / 3          | 入力ファイルのBaseUnit                             | 
| TaxCode                         | string(varchar) / 2          | 5-26でセットTaxCode                                | 
| ConditionAmount                 | float / 13                   | 8-2でセットしたConditionAmount <br>または <br>入力ファイルのConditionAmount   | 
| TransactionCurrency             | string(varchar) / 5          | 1-11で保持したTransactionCurrency                  | 
| ConditionIsManuallyChanged      | bool(tinyint) / 1            | 8-2または8-3でセットしたConditionIsManuallyChanged | 

200-14. Orders Item Schedule Lineデータの整列とセット
次のOrders Item Schedule Lineデータを整列しDB/Tableへの登録値としてセットする。  
<対象テーブル: data_platform_orders_item_schedule_line_data>  

| <b>Field Created</b>                                | <b>Data Type / Number of Digits</b> | <b>Specification</b>                                              | 
| -------------------------------------------- | ---------------------------- | ---------------------------------------------------------- | 
| OrderID(※Key)                               | int / 16                     | 5-28-2-3または5-28-3-3(以下同様)で生成したOrderID          | 
| OrderItem(※Key)                             | int / 6                      | 同様に生成したOrderItem                                    | 
| ScheduleLine(※Key)                          | int / 3                      | 同様に生成したScheduleLine                                 | 
| Product                                      | string(varchar) / 40         | 同様に生成したProduct                                      | 
| StockConfirmationPartnerFunction             | string(varchar) / 40         | 同様に生成したStockConfirmationPartnerFunction             | 
| StockConfirmationBusinessPartner             | int / 12                     | 同様に生成したStockConfirmationBusinessPartner             | 
| StockConfirmationPlant                       | string(varchar) / 4          | 同様に生成したStockConfirmationPlant                       | 
| StockConfirmationPlantBatch                  | string(varchar) / 10         | 同様に生成したStockConfirmationPlantBatch                  | 
| StockConfirmationPlantBatchValidityStartDate | date                         | 同様に生成したStockConfirmationPlantBatchValidityStartDate | 
| StockConfirmationPlantBatchValidityEndDate   | date                         | 同様に生成したStockConfirmationPlantBatchValidityEndDate   | 
| RequestedDeliveryDate                        | date                         | 同様に生成したRequestedDeliveryDate                        | 
| ConfirmedDeliveryDate                        | date                         | 同様に生成したConfirmedDeliveryDate                        | 
| OrderQuantityInBaseUnit                      | string(varchar) / 3          | 同様に生成したOrderQuantityInBaseUnit                      | 
| ConfdOrderQtyByPDTAvailCheck                 | float / 15                   | 同様に生成したConfdOrderQtyByPDTAvailCheck                 | 
| DeliveredQtyInOrderQtyUnit                   | float / 15                   | 同様に生成したDeliveredQtyInOrderQtyUnit                   | 
| OpenConfdDelivQtyInOrdQtyUnit                | float / 15                   | 同様に生成したOpenConfdDelivQtyInOrdQtyUnit                | 
| StockIsFullyConfirmed                        | bool(tinyint) / 1            | 同様に生成したStockIsFullyConfirmed                        | 
| DelivBlockReasonForSchedLine                 | bool(tinyint) / 1            | 同様に生成したDelivBlockReasonForSchedLine                 | 
| PlusMinusFlag                                | string(varchar) / 1          | 同様に生成したPlusMinusFlag                                | 

200-20. Orders Addressデータの整列とセット
次のOrders Addressデータを整列しDB/Tableへの登録値としてセットする。  
<対象テーブル: data_platform_orders_address_data>  

| <b>Field Created</b>    | <b>Data Type / Number of Digits</b> | <b>Specification</b>                                            | 
| ---------------- | ---------------------------- | -------------------------------------------------------- | 
| OrderID(※Key)   | int / 16                     | 1-2-2でセットされたOrderID                               | 
| AddressID(※Key) | int / 12                     | 2-2-1で取得したAddressIDまたは9-2-3でセットしたAddressID | 
| PostalCode       | string(varchar) / 10         | 9-1で取得したPostalCodeまたは入力ファイルのPostalCode    | 
| LocalRegion      | string(varchar) / 3          | 9-1で取得したLocalRegionまたは入力ファイルのLocalRegion  | 
| Country          | string(varchar) / 3          | 9-1で取得したCountryまたは入力ファイルのCountry          | 
| District         | string(varchar) / 6          | 9-1で取得したDistrictまたは入力ファイルのDistrict        | 
| StreetName       | string(varchar) / 200        | 9-1で取得したStreetNameまたは入力ファイルのStreetName    | 
| CityName         | string(varchar) / 200        | 9-1で取得したCityNameまたは入力ファイルのCityName        | 
| Builiding        | string(varchar) / 100        | 9-1で取得したBuilidingまたは入力ファイルのBuiliding      | 
| Floor            | int / 4                      | 9-1で取得したFloorまたは入力ファイルのFloor              | 
| Room             | int / 8                      | 9-1で取得したFloorまたは入力ファイルのFloor              | 