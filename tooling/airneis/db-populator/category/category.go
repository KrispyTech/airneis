package category

import (
	"encoding/json"
	"fmt"
	"net/http"

	httpclient "github.com/krispyTech/airneis/tooling/db-populator/httpClient"
	model "github.com/krispyTech/airneis/tooling/db-populator/models"
)

func CreateCategories(hc httpclient.HttpApi) error {
	amountToAdd := 30

	for count := 0; count < amountToAdd; count++ {
		categoryID := uint(count)
		category := model.Category{
			Name:           fmt.Sprintf("Category Name %d", count),
			ThumbnailURL:   fmt.Sprintf("CategoryURL%d.com", count),
			Slug:           fmt.Sprintf("Category Slug %d", count),
			OrderOfDisplay: &categoryID,
		}

		bodyToSend, err := json.Marshal(category)
		if err != nil {
			return fmt.Errorf("CreateCategories, Unable to marshall")
		}

		headers := make(map[string]string)
		headers["content-type"] = "application/json"

		_, status, err := hc.Post("http://localhost:3001/api/categories", headers, bodyToSend)
		if err != nil || status != http.StatusOK {
			fmt.Printf("CreateCategories, request impossible %s\n", err)
			continue
		}
	}

	return nil
}
