package main

import (
	"context"
	"fmt"
	"log"
	"zrDispatch/common/utils"

	"github.com/olivere/elastic/v7"
)

func main() {
	client, err := elastic.NewClient(elastic.SetURL("http://172.16.130.138:9200"))
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %v", err)
	}

	indexName := "cyberspace-resources_ads" // 替换为您的索引名

	// 构建查询
	query := elastic.NewMatchAllQuery()

	// 构建搜索服务
	searchService := client.Search().Index(indexName).Query(query)

	// 设置每页大小
	searchService.Size(10)

	// 执行查询
	searchResult, err := searchService.Do(context.Background())
	if err != nil {
		log.Fatalf("Error executing search query: %v", err)
	}

	// 处理搜索结果
	if searchResult.Hits.TotalHits.Value > 0 {
		for _, hit := range searchResult.Hits.Hits {
			source := hit.Source
			// 处理 source 数据
			utils.PrinfI("s", source)
		}
	} else {
		fmt.Println("No documents found")
	}
}
