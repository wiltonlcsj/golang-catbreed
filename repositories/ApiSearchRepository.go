package repositories

import (
	"github.com/wiltonlcsj/hgBackendTest/adapters"
	"github.com/wiltonlcsj/hgBackendTest/models"
	"log"
)

type IApiSearchRepository interface {
	FindByQueryName(name string) bool
	InsertNew(apiSearch models.ApiSearch)
}

type ApiSearchRepository struct {
	DbAdapter *adapters.DatabaseAdapter
}

func NewApiSearchRepository() *ApiSearchRepository {
	return &ApiSearchRepository{
		DbAdapter: adapters.NewDatabaseAdapter(),
	}
}

func (repository *ApiSearchRepository) FindByQueryName(name string) bool {
	row, err := repository.DbAdapter.QueryRow("SELECT id FROM api_search where query_param = ?", name)
	if err != nil {
		return false
	}

	apisearch := &models.ApiSearch{}
	err = row.Scan(&apisearch.Id)
	if err != nil {
		return false
	}

	return true
}

func (repository *ApiSearchRepository) InsertNew(apiSearch models.ApiSearch) {
	_, err := repository.DbAdapter.ExecuteCommand("INSERT INTO api_search (query_param, api_status, response) "+
		"VALUES (?, ?, ?)", apiSearch.QueryParam, apiSearch.ApiStatus, apiSearch.Response)
	if err != nil {
		log.Println("cannot insert new api search row")
	}
}
