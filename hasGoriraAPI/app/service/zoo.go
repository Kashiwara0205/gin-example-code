package service

import (
	"hasGoriraAPI/db"
)

type zooRecord struct {
	Name  string
	PrefName string
	HasGorira  bool
}

type ZooService struct{}
func (s *ZooService) GetZoos() []zooRecord{
	connection := db.DBConnect()
	defer connection.Close()

	var records []zooRecord

	connection.
		Table("zoos").
		Select("zoos.name, prefectures.name as pref_name, zoos.has_gorira").
		Joins("left outer join prefectures on prefectures.id = zoos.pref_id").
		Scan(&records)

	return records
}