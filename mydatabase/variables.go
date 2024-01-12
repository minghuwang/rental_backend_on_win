package mydatabase

import "database/sql"

var Db *sql.DB
var TbName = "property_info"

type PropertyInfoSchema struct {
	// Variables to hold column values
	PropertyID  int    `json:"PropertyID"`
	Address     string `json:"Address"`
	PictureLink string `json:"PictureLink"`
	OpenTime1   int64  `json:"OpenTime1"`
	OpenTime2   int64  `json:"OpenTime2"`
}