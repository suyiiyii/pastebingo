package models

func Migrate() {
	db := DB

	err := db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Post{})
	if err != nil {
		panic(err)
	}
}
