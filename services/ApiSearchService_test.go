package services

import (
	"testing"
)

func TestApiSearchService_SearchOnApiEmptyTerm(t *testing.T) {
	apisearchService := &ApiSearchService{
		apiUrl: "",
		apiKey: "",
	}

	_, err := apisearchService.SearchOnApi("")
	if err == nil{
		t.Fail()
	}
}

func TestApiSearchService_SearchOnApiOffline(t *testing.T) {
	apisearchWrongService := &ApiSearchService{
		apiUrl: "",
		apiKey: "",
	}
	_, err := apisearchWrongService.SearchOnApi("Sib")
	if err == nil{
		t.Fail()
	}
}

func TestApiSearchService_SearchOnApiOnline(t *testing.T) {
	apiSearchService := &ApiSearchService{
		apiUrl: "https://api.thecatapi.com/v1",
		apiKey: "cd5c4168-1e5f-4db4-8ca3-c77b3da9c117",
	}

	_, err := apiSearchService.SearchOnApi("Sib")
	if err != nil{
		t.Fail()
	}
}
