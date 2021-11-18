package database

type Database struct {
}

func (db Database) GetAll() []string {
	return []string{"one", "two", "three"}
}

type DatabaseMock struct {
}

func (db DatabaseMock) GetAll() []string {
	return []string{"mocked", "mocked", "mocked"}
}

type DatabaseInterface interface {
	GetAll() []string
}
