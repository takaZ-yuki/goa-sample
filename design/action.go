package design

import . "goa.design/goa/v3/dsl"

// Service describes a service
var _ = Service("goa_sample", func() {
    Description("Echo your request")
    // サービスで提供されるメソッドを複数定義できます
    Method("echo", func() {
        // Payload はメソッドのペイロードを記述します
        // ここでは、ペイロードは2つのフィールドからなるオブジェクトです
        Payload(func() {
            // Attribute はフィールドインデックス、フィールド名、タイプ、および説明が与えられたオブジェクトを記述します
            Attribute("name", String, "Your name")
            Attribute("age", Int, "Your age")
            // Both attributes must be provided when invoking "add"
            Required("name", "age")
        })

				// Result はメソッドの結果を記述します
        // ここでは、結果は単純に1つの文字列値です
        Result(String)

				// HTTP は HTTP トランスポートへの対応を記述します
        HTTP(func() {
            // サービスへのリクエストは HTTP GET リクエストで構成されています
            // ペイロードはパスパラメータとしてエンコードされます
            GET("/echo/{name}")
						// クエリパラメーター
						Param("age")
            // レスポンスは HTTP ステータス "200 OK" を使用します
            // 結果はレスポンスボディにエンコードされます（デフォルト）
            Response(StatusOK)
        })
    })
})