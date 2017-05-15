package main

import (
	"cloud.google.com/go/bigquery"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"log"
)

func main() {
	fetchBigQueryData()
}

// BigQueryで実行するためのSQL
const QUERY = `
SELECT  *
FROM [bigquery-public-data:usa_names.usa_1910_2013]
WHERE name = 'Mary'
LIMIT 100
`

func fetchBigQueryData() {
	// 空のcontextを生成
	ctx := context.Background()

	// プロジェクトのID
	projectID := "sandbox-sekky0905"

	// contextとprojectIDを元にBigQuery用のclientを生成
	client, err := bigquery.NewClient(ctx, projectID)

	if err != nil {
		log.Printf("Failed to create client:%v", err)
	}

	// 引数で渡した文字列を元にQueryを生成
	q := client.Query(QUERY)

	// 実行のためのqueryをサービスに送信してIteratorを通じて結果を返す
	// itはIterator
	it, err := q.Read(ctx)

	if err != nil {
		log.Println("Failed to Read Query:%v", err)
	}

	for {
		// BigQueryの結果から、中身を格納するためのBigQuery.Valueのsliceを宣言
		// BigQuery.Valueはinterface{}型
		var values []bigquery.Value

		// 引数に与えたvaluesにnextを格納する
		// Iteratorを返す
		// これ以上結果が存在しない場合には、iterator.Doneを返す
		// iterator.Doneが返ってきたら、forを抜ける
		err := it.Next(&values)
		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Println("Failed to Iterate Query:%v", err)
		}

		fmt.Println(values)
	}

}
