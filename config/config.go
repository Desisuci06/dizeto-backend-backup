package config

import (
	model_about "dizeto-backend/app/model/about"
	model_client "dizeto-backend/app/model/client"
	model_counting "dizeto-backend/app/model/counting"
	model_highlight "dizeto-backend/app/model/highlight_porto"
	model_item "dizeto-backend/app/model/item"
	model_page "dizeto-backend/app/model/page"
	model_pricing "dizeto-backend/app/model/pricing"
	model_testimoni "dizeto-backend/app/model/testimoni"
	model_title "dizeto-backend/app/model/title"
	model_user "dizeto-backend/app/model/user"
	"dizeto-backend/utils"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitDB() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPassword)
	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&model_user.User{},
		&model_about.About{},
		&model_highlight.HighlightPortofolio{},
		&model_pricing.Pricing{},
		&model_testimoni.Testimoni{},
		&model_counting.Counting{},
		&model_client.Client{},
		&model_page.Page{},
		&model_item.ItemList{},
		&model_title.Title{},
	).Error

	if err != nil {
		return nil, err
	}

	// database user seeding
	var users = []model_user.User{}
	db.Where("role = ?", "admin").Find(&users)
	fmt.Println(len(users))
	if len(users) == 0 {
		err = SeedUsers(db)
		if err != nil {
			return nil, err
		}
	}

	// database page seeding
	var page = []model_page.Page{}
	db.Where("id = ?", 1).Find(&page)
	fmt.Println(len(page))
	if len(page) == 0 {
		err = SeedPage(db)
		if err != nil {
			return nil, err
		}
	}

	// database title seeding
	var title = []model_title.Title{}
	db.Where("id = ?", 1).Find(&title)
	fmt.Println(len(title))
	if len(title) == 0 {
		err = SeedTitle(db)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

func SeedUsers(db *gorm.DB) error {
	userID, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword("admin")
	if err != nil {
		return err
	}
	userAdmin := model_user.User{ID: userID, Username: "admin", Password: hashedPassword, FirstName: "Admin", LastName: "Dizeto", Email: "admin@gmail.com", Role: "admin"}
	db.Create(&userAdmin)

	return nil
}

func SeedPage(db *gorm.DB) error {
	pageID := 1

	page := model_page.Page{ID: uint(pageID), Title: "Page 1"}
	db.Create(&page)
	return nil
}

func SeedTitle(db *gorm.DB) error {
	about := model_title.Title{ID: 1, KdTitle: "ABT", NmTitle: "ABOUT"}
	portofolio := model_title.Title{ID: 2, KdTitle: "PORT", NmTitle: "PORTOFOLIO"}
	pricing := model_title.Title{ID: 3, KdTitle: "PRIC", NmTitle: "PRICING"}
	testimoni := model_title.Title{ID: 4, KdTitle: "TEST", NmTitle: "TESTIMONI"}
	counting := model_title.Title{ID: 5, KdTitle: "COUN", NmTitle: "COUNTING"}
	client := model_title.Title{ID: 6, KdTitle: "CLIE", NmTitle: "CLIENT"}
	db.Create(&about)
	db.Create(&portofolio)
	db.Create(&pricing)
	db.Create(&testimoni)
	db.Create(&counting)
	db.Create(&client)
	return nil
}
