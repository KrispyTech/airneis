package category

import (
	"encoding/json"
	"fmt"
	"net/http"

	httpclient "github.com/krispyTech/airneis/tooling/db-populator/httpClient"
)

type Category struct {
	Name           string `json:"name"`
	ThumbnailUrl   string `json:"thumbnailUrl"`
	Slug           string `json:"slug"`
	OrderOfDisplay *uint  `json:"orderOfDisplay"`
}

func CreateCategories(hc httpclient.HttpApi) error {
	for i := 0; i < 100; i++ {
		categoryID := uint(i)
		category := Category{
			Name:           fmt.Sprintf("Category Name %d", i),
			ThumbnailUrl:   fmt.Sprintf("CategoryURL%d.com", i),
			Slug:           fmt.Sprintf("Category Slug %d", i),
			OrderOfDisplay: &categoryID,
		}

		bodyToSend, err := json.Marshal(category)
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
