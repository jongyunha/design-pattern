package pets

import "go-breeders/models"

func NetPet(species string) *models.Pet {
	return &models.Pet{
		Species:     species,
		Breed:       "",
		MinWeight:   0,
		MaxWeight:   0,
		Description: "no description entered yet",
		LifeSpan:    0,
	}
}
