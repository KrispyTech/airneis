package main

import (
	"log"

	"github.com/krispyTech/airneis/tooling/db-populator/category"
	httpclient "github.com/krispyTech/airneis/tooling/db-populator/httpClient"
	"github.com/krispyTech/airneis/tooling/db-populator/product"
)

func main() {
	httpClient := httpclient.InitializeHttpApi()
	if err := category.CreateCategories(httpClient); err != nil {
		log.Fatal(err)
	}
	if err := product.CreateProducts(httpClient); err != nil {
		log.Fatal(err)
	}
}
