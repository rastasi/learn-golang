package serializer

import "github.com/rastasi/learn-golang/crud/domain/model"

type Album struct {
	ID     uint    `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type AlbumSerializer struct{}

func (s AlbumSerializer) SerializeAlbum(album model.AlbumModel) Album {
	return Album{
		ID:     album.ID,
		Title:  album.Title,
		Artist: album.Artist,
		Price:  album.Price,
	}
}

func (s AlbumSerializer) SerializeAlbums(albums []model.AlbumModel) []Album {
	var serialized []Album
	for _, album := range albums {
		serialized = append(serialized, s.SerializeAlbum(album))
	}
	return serialized
}
