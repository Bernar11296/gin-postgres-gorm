package database

import (
	"errors"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

type Article struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Rate        int    `json:"rate"`
}

// func getEnvVariable(key string) string {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatal("Error loading .env file", err)
// 	}
// 	return os.Getenv(key)
// }

func NewPostgreSQLClient() {
	// conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
	// 	"localhost",
	// 	"5432",
	// 	"postgres",
	// 	"postgres",
	// 	"12345678",
	// )

	// db, err = gorm.Open("postgres", conn)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(Article{})
}

func CreateArticle(a *Article) (*Article, error) {
	res := db.Create(a)
	if res.RowsAffected == 0 {
		return &Article{}, errors.New("article not created")
	}
	return a, nil
}

func ReadArticle(id string) (*Article, error) {
	var article Article
	res := db.First(&article, id)
	if res.RowsAffected == 0 {
		return nil, errors.New("article not found")
	}
	return &article, nil
}

func ReadArticles() ([]*Article, error) {
	var articles []*Article
	res := db.Find(&articles)
	if res.Error != nil {
		return nil, errors.New("authors not found")
	}
	return articles, nil
}

func UpdateArticle(article *Article) (*Article, error) {
	var updateArticle Article
	result := db.Model(&updateArticle).Where(article.ID).Updates(article)
	if result.RowsAffected == 0 {
		return &Article{}, errors.New("artcile not updated")
	}
	return &updateArticle, nil
}

func DeleteArticle(id string) error {
	var deleteArticle Article
	result := db.Where(id).Delete(&deleteArticle)
	if result.RowsAffected == 0 {
		return errors.New("article data not deleted")
	}
	return nil
}
