package serializer

import "github.com/rastasi/learn-golang/crud/app/model"

type AlbumSerializer struct {
	ID     uint    `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func SerializeAlbum(album model.AlbumModel) AlbumSerializer {
	return AlbumSerializer{
		ID:     album.ID,
		Title:  album.Title,
		Artist: album.Artist,
		Price:  album.Price,
	}
}

func SerializeAlbums(albums []model.AlbumModel) []AlbumSerializer {
	var serialized []AlbumSerializer
	for _, album := range albums {
		serialized = append(serialized, SerializeAlbum(album))
	}
	return serialized
}
