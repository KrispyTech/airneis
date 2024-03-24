package product

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	httpclient "github.com/krispyTech/airneis/tooling/db-populator/httpClient"
	model "github.com/krispyTech/airneis/tooling/db-populator/models"
)

func CreateProducts(hc httpclient.HttpApi) error {
	amountToAdd := 30

	for count := 0; count < amountToAdd; count++ {
		productID := uint(count)
		product := model.Product{
			Name:           				fmt.Sprintf("Product Name %d", count),
			Slug:           				fmt.Sprintf("Product Slug %d", count),
			CategoryID:							uint(rand.Intn(29) + 1),
			Category:								model.Category{
															Name: fmt.Sprintf("Category name %d", count),
															ThumbnailURL: fmt.Sprintf("CategoryURL%d.com", count),
															Slug: fmt.Sprintf("Category Slug %d", count),
															OrderOfDisplay: &productID,
															},
			PriceWithoutTaxes: 			rand.Intn(2500),
			PriceWithTaxes: 				rand.Intn(2500),
			Description:    				fmt.Sprintf("Product Description %d", count),
			ThumbnailURL:   				fmt.Sprintf("ProductURL%d.com", count),
			OutOfStock:         		false,
			Materials:              []model.Material{},
			HighlanderOrder: 				0,
			OrderOfPriority: &productID,
			ProductImages:          []model.ProductImage{},
		}

		bodyToSend, err := json.Marshal(product)
		if err != nil {
			return fmt.Errorf("CreateProducts, Unable to marshall")
		}

		headers := make(map[string]string)
		headers["content-type"] = "application/json"

		_, status, err := hc.Post("http://localhost:3001/api/products", headers, bodyToSend)
		if err != nil || status != http.StatusOK {
			fmt.Printf("CreateProducts, request impossible %s\n", err)
			continue
		}
	}

	return nil
}
