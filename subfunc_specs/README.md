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
	| Table Searched                           | Name of the Table                            | 
	| Number Range SQL Number Range            | Data_platform_number_range_number_range_data | 
	| Field Searched                           | Data Type / Number of Digits                 | 
	| ServiceLabel <br> FieldNameWithNumberRange <br> NumberRangeFrom  <br> NumberRangeTo                            | string(varchar) / 50 <br> string(varchar) / 100 <br> int / 10 <br> int / 10                                 | 
	| Single Record or Array                   | Memo                                         | 
	| Array                                    | 対象テーブルの全レコードを取得               |   

	### [参考]
	
	以下の例のようなデータが取得される。  
		
	| NumberRangeID | ServiceLabel   | FieldNameWithNumberRange | NumberRangeFrom | NumberRangeTo | 
	| ------------- | -------------- | ------------------------ | --------------- | ------------- | 
	| “01”        | “ORDERS”     | OrderID                  | 10000000        | 19999999      | 
	| “02”        | “QUOTATIONS” | Quotation                | 20000000        | 29999999      | 
	| “08”        | “INQUIRIES”  | Inquiry                  | 80000000        | 89999999      |   
		
	0-2-2. 入力ファイルのReferenceDocumentが取得したテーブルのどの範囲に当てはまるかを判定し、その判定されたレコードのServiceLabelをセットする。
