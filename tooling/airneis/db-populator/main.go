package main

import (
	"log"

	"github.com/krispyTech/airneis/tooling/db-populator/category"
	httpclient "github.com/krispyTech/airneis/tooling/db-populator/httpClient"
)

func main() {
	httpClient := httpclient.InitializeHttpApi()
	if err := category.CreateCategories(httpClient); err != nil {
		log.Fatal(err)
	}
}
