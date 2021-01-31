package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wiltonlcsj/hgBackendTest/models"
	"io/ioutil"
	"net/http"
	"os"
)

type IApiSearchService interface {
	SearchOnApi(string) string
}

type ApiSearchService struct {
	apiUrl string
	apiKey string
}

func NewApiSearchService() *ApiSearchService {
	return &ApiSearchService{
		apiUrl: os.Getenv("API_CAT_URL"),
		apiKey: os.Getenv("API_CAT_KEY"),
	}
}

func (apiService *ApiSearchService) SearchOnApi(searchTerm string) ([]models.CatBreed, error) {
	if len(searchTerm) == 0 {
		return nil, errors.New("search term cannot be empty")
	}

	searchUrl := fmt.Sprintf("%s/breeds/search", apiService.apiUrl)
	req, err := http.NewRequest("GET", searchUrl, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("x-api-key", apiService.apiKey)
	query := req.URL.Query()
	query.Add("q", searchTerm)
	req.URL.RawQuery = query.Encode()

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var jsonBodyApi SearchBreedReturnApi
	byteValue, _ := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(byteValue, &jsonBodyApi); err != nil {
		return nil, err
	}

	return jsonBodyApi, nil
}
