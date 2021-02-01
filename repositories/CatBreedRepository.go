package repositories

import (
	"github.com/wiltonlcsj/hgBackendTest/adapters"
	"github.com/wiltonlcsj/hgBackendTest/models"
	"log"
)

type ICatBreedRepository interface {
	FindByName(name string) []models.CatBreed
	InsertNew(breed models.CatBreed)
}

type CatBreedRepository struct {
	DbAdapter *adapters.DatabaseAdapter
}

func NewCatBreedRepository() *CatBreedRepository {
	return &CatBreedRepository{
		DbAdapter: adapters.NewDatabaseAdapter(),
	}
}

func (repository *CatBreedRepository) FindByName(name string) []models.CatBreed {
	rows, err := repository.DbAdapter.QueryMultiple("SELECT * from cat_breed WHERE name like ?", "%"+name+"%")
	if err == nil {
		var catBreeds []models.CatBreed

		for rows.Next() {
			catbreed := new(models.CatBreed)
			if err := rows.Scan(
				&catbreed.Id,
				&catbreed.Adaptability,
				&catbreed.AffectionLevel,
				&catbreed.AltNames,
				&catbreed.CfaURL,
				&catbreed.ChildFriendly,
				&catbreed.CountryCode,
				&catbreed.CountryCodes,
				&catbreed.Description,
				&catbreed.DogFriendly,
				&catbreed.EnergyLevel,
				&catbreed.Experimental,
				&catbreed.Grooming,
				&catbreed.Hairless,
				&catbreed.HealthIssues,
				&catbreed.Hypoallergenic,
				&catbreed.Indoor,
				&catbreed.Intelligence,
				&catbreed.Lap,
				&catbreed.LifeSpan,
				&catbreed.Name,
				&catbreed.Natural,
				&catbreed.Origin,
				&catbreed.Rare,
				&catbreed.ReferenceImageId,
				&catbreed.Rex,
				&catbreed.SheddingLevel,
				&catbreed.ShortLegs,
				&catbreed.SocialNeeds,
				&catbreed.StrangerFriendly,
				&catbreed.SuppressedTail,
				&catbreed.Temperament,
				&catbreed.VcahospitalsURL,
				&catbreed.VetstreetURL,
				&catbreed.Vocalisation,
				&catbreed.WikipediaURL,
			); err != nil {
				log.Println(err)
				continue
			}
			catBreeds = append(catBreeds, *catbreed)
		}

		defer rows.Close()
		return catBreeds
	}

	return nil
}

func (repository *CatBreedRepository) VerifyIfExists(name string) (bool, error) {
	row, err := repository.DbAdapter.QueryRow("SELECT id FROM cat_breed where name = ?", name)
	if err != nil {
		return false, err
	}

	catbreed := &models.CatBreed{}
	err = row.Scan(&catbreed.Id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (repository *CatBreedRepository) InsertNew(breed models.CatBreed) {
	_, err := repository.DbAdapter.ExecuteCommand("INSERT INTO cat_breed VALUES "+
		"(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		breed.Id, breed.Adaptability, breed.AffectionLevel, breed.AltNames, breed.CfaURL, breed.ChildFriendly,
		breed.CountryCode, breed.CountryCodes, breed.Description, breed.DogFriendly, breed.EnergyLevel,
		breed.Experimental, breed.Grooming, breed.Hairless, breed.HealthIssues, breed.Hypoallergenic,
		breed.Indoor, breed.Intelligence, breed.Lap, breed.LifeSpan, breed.Name, breed.Natural,
		breed.Origin, breed.Rare, breed.ReferenceImageId, breed.Rex, breed.SheddingLevel, breed.ShortLegs,
		breed.SocialNeeds, breed.StrangerFriendly, breed.SuppressedTail, breed.Temperament, breed.VcahospitalsURL,
		breed.VetstreetURL, breed.Vocalisation, breed.WikipediaURL)
	if err != nil {
		log.Println("cannot insert new breed row")
	}
}
