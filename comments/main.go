package comments

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	newrelic "github.com/newrelic/go-agent"
	"github.com/thinkerou/favicon"
)

func init() {
	router := gin.Default()
	cfg := newrelic.NewConfig(os.Getenv("APP_NAME"), os.Getenv("NEW_RELIC_API_KEY"))
	app, err := newrelic.NewApplication(cfg)
	if err != nil {
		log.Printf("failed to make new_relic app: %v", err)
	} else {
		router.Use(adapters.NewRelicMonitoring(app))
	}
}

func main() {
	/* 	app := fiber.New()

	   	app.Get("/", func(c *fiber.Ctx) error {
	   		return c.SendString("Hello, World 👋!")
	   	}) */

	app := gin.Default()
	app.Use(favicon.New("./favicon.ico"))
	app.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello favicon.")
	})
	app.Run(":8000")
}
