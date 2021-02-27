package main

import (
	"hasGoriraAPI/app/model"
	"hasGoriraAPI/db"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	connection := db.DBConnect()
	defer connection.Close()

	createPrefectureSeedData(connection)
	createZooSeedData(connection)
}

func createPrefectureSeedData(connection *gorm.DB){
	prefs := []model.Prefecture{
		{ Name: "北海道"},
		{ Name: "青森県" },
		{ Name: "岩手県" },
		{ Name: "宮城県" },
		{ Name: "秋田県" },
		{ Name: "山形県" },
		{ Name: "福島県" },
		{ Name: "茨城県" },
		{ Name: "栃木県" },
		{ Name: "群馬県" },
		{ Name: "埼玉県" },
		{ Name: "千葉県" },
		{ Name: "東京都" },
		{ Name: "神奈川県" },
		{ Name: "新潟県" },
		{ Name: "富山県" },
		{ Name: "石川県" },
		{ Name: "福井県" },
		{ Name: "山梨県" },
		{ Name: "長野県" },
		{ Name: "岐阜県" },
		{ Name: "静岡県" },
		{ Name: "愛知県" },
		{ Name: "三重県" },
		{ Name: "滋賀県" },
		{ Name: "京都府" },
		{ Name: "大阪府" },
		{ Name: "兵庫県" },
		{ Name: "奈良県" },
		{ Name: "和歌山県" },
		{ Name: "鳥取県" },
		{ Name: "島根県" },
		{ Name: "岡山県" },
		{ Name: "広島県" },
		{ Name: "山口県" },
		{ Name: "徳島県" },
		{ Name: "香川県" },
		{ Name: "愛媛県" },
		{ Name: "高知県" },
		{ Name: "福岡県" },
		{ Name: "佐賀県" },
		{ Name: "長崎県" },
		{ Name: "熊本県" },
		{ Name: "大分県" },
		{ Name: "宮崎県" },
		{ Name: "鹿児島県" },
		{ Name: "沖縄県" }, 
	}
	
	for _, pref := range prefs {
		connection.Create(&pref)
	}
}

func createZooSeedData(connection *gorm.DB){
	zoos := []model.Zoo{
		{ PrefId: 1,  Name: "北海道ゴリラ大学", HasGorira:true },
		{ PrefId: 12, Name: "ぽんぽこ動物園",   HasGorira:false },
		{ PrefId: 3,  Name: "わにわに博物館",   HasGorira:false },
		{ PrefId: 33, Name: "ゴリラの里",      HasGorira:true },
		{ PrefId: 22, Name: "白鳥の森",        HasGorira:false },
	}
	
	for _, zoo := range zoos {
		connection.Create(&zoo)
	}
}