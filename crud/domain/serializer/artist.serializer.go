package serializer

import "github.com/rastasi/learn-golang/crud/domain/model"

type Artist struct {
	ID     uint   `json:"id"`
	Name   string `json:"title"`
	Active int8   `json:"active"`
}

type ArtistSerializer struct{}

func (s ArtistSerializer) SerializeArtist(artist model.ArtistModel) Artist {
	return Artist{
		ID:     artist.ID,
		Name:   artist.Name,
		Active: artist.Active,
	}
}

func (s ArtistSerializer) SerializeArtists(artists []model.ArtistModel) []Artist {
	var serialized []Artist
	for _, artist := range artists {
		serialized = append(serialized, s.SerializeArtist(artist))
	}
	return serialized
}
