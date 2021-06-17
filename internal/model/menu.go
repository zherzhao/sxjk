package model

type Menus struct {
	TableName string `eorm:"tableName"`
	Year      string `eorm:"year"`
	Levels    int    `eorm:"levelCount"`
}
