package response

import (
	"capstone/group3/features/favourites"
)

type RestoFav struct {
	RestoID       int          `json:"resto_id"`
	RestoImageUrl []RestoImage `json:"resto_image_url"`
	Rating        float64      `json:"rating"`
	Category      string       `json:"category"`
	RestoName     string       `json:"resto_name"`
	Location      string       `json:"location"`
}

type RestoImage struct {
	RestoImageUrl string `json:"resto_image_url"`
}

func FromCoreAll(data favourites.RestoCore) RestoFav {
	return RestoFav{
		RestoID:       data.ID,
		RestoName:     data.RestoName,
		Location:      data.Location,
		Category:      data.Category,
		Rating:        data.Rating,
		RestoImageUrl: FromRestoImageCoreList(data.RestoImages)}
}

func FromRestoImageCore(data favourites.RestoImage) RestoImage {
	return RestoImage{
		RestoImageUrl: data.RestoImageUrl,
	}
}

func FromRestoImageCoreList(data []favourites.RestoImage) []RestoImage {
	result := []RestoImage{}
	for key := range data {
		result = append(result, FromRestoImageCore(data[key]))
	}
	return result
}

func FromCoreList(data []favourites.RestoCore) []RestoFav {
	result := []RestoFav{}
	for key := range data {
		result = append(result, FromCoreAll(data[key]))
	}
	return result
}
