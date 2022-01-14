package drivers

import (
	_newsDomain "alta/cleanarch/businesses/news"
	_newsDatabase "alta/cleanarch/drivers/databases/news"

	gorm "gorm.io/gorm"
)

//NewNewsRepository Factory with news domain
func NewNewsRepository(conn *gorm.DB) _newsDomain.Repository {
	return _newsDatabase.NewMySQLRepository(conn)
}
