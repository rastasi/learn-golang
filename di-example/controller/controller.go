package controller

import "github.com/rastasi/learn-golang/di-example/database"

type Controller struct {
	Database database.DatabaseInterface
}

func (c Controller) Index() []string {
	return c.Database.GetAll()
}
