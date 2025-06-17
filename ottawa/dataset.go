package ottawa

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type Dataset struct {
	Type        string   `json:"@type"`
	Identifier  string   `json:"indentifier"`
	LandingPage string   `json:"landingPage"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Keyword     []string `json:"keyword"`
	Issued      string   `json:"issued"`
	Modified    string   `json:"modified"`
	Publisher   struct {
		Name string `json:"name"`
	} `json:"publisher"`
	ContactPoint struct {
		Type     string `json:"@type"`
		Fn       string `json:"fn"`
		HasEmail string `json:"hasEmail"`
	} `json:"contactPoint"`
	AccessLevel  string `json:"accessLevel"`
	Spatial      string `json:"spatial"`
	License      string `json:"license"`
	Distribution []struct {
		Type      string `json:"@type"`
		Title     string `json:"title"`
		Format    string `json:"format"`
		MediaType string `json:"mediaType"`
		AccessURL string `json:"accessURL"`
	} `json:"distribution"`
	Theme []string `json:"theme"`
}

type Payload struct {
	Context     string `json:"@context"`
	Type        string `json:"@type"`
	ConformsTo  string `json:"conformsTo"`
	DescribedBy string `json:"describedBy"`
	Dataset     []Dataset
}

const DATASET_URL = "https://open.ottawa.ca/api/feed/dcat-us/1.1.json"

// Grabs all the datasets made available to the public by the City of Ottawa.
//
// It doesn’t actually download the data in the dataset(s), just the information
// on the dataset itself. A separate program should make use of this method to
// download the datasets.
//
// For more information, visit https://open.ottawa.ca/pages/developer-resources
func ListAllDataset() ([]Dataset, error) {
	client := http.Client{Timeout: 30 * time.Second}

	resp, err := client.Get(DATASET_URL)
	if err != nil {
		log.Fatalf("Error grabbing data from source: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var payload Payload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	return payload.Dataset, nil
}
