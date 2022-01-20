package main

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Comment struct {
	Id     uint   `json:"id"`
	PostId uint   `json:"post_id"`
	Text   string `json:"text"`
}

func main() {

	dsn := "root:comments@tcp(127.0.0.1:3307)/comments?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)

	}

	db.AutoMigrate(Comment{})

	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/api/posts/:id/comments", func(c *fiber.Ctx) error {
		var comments []Comment

		db.Find(&comments, "post_id = ?", c.Params("id"))
		return c.JSON(comments)
	})

	app.Post("/api/comment", func(c *fiber.Ctx) error {
		var comment Comment

		if err := c.BodyParser(&comment); err != nil {
			return err
		}

		db.Create(&comment)
		return c.JSON(comment)
	})

	app.Listen(":4000")
}
