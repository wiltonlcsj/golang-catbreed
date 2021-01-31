package services

import "github.com/wiltonlcsj/hgBackendTest/models"

type SearchBreedService struct {
	Name string `json:"name" form:"name" binding:"required"`
}

type SearchBreedReturnApi []models.CatBreed
