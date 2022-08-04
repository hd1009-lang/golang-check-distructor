package main

import (
	"director/configs"
	"director/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	fmt.Println("Go MySQL Tutorial")
	app := fiber.New(fiber.Config{BodyLimit: 130 * 1024 * 1024})
	routes.Routes(app)
	port := configs.Config("PORT")
	if port == "" {
		port = "1334"
	}
	log.Fatal(app.Listen(":" + port))
	//db := database.MariaDB()
	//defer db.Close()
	//distributor_news_id := helpers.GenerateId()
	//loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	//now := time.Now().In(loc)
	//fmt.Println(now)
	//err1, err := db.Exec("INSERT INTO distributor_news (supplier_id, create_at, update_at, distributor_news_id, categories_id, status, expiry_date, content,title) VALUES (?, ?, ?, ?, ?, ?, ?, ?,? ) ", 589, now, now, distributor_news_id, 10, "wait", now, "Hello 2", "Check")
	//fmt.Println(err1)
	//if err != nil {
	//	panic(err.Error())
	//}

	// be careful deferring Queries if you are using transactions
}
