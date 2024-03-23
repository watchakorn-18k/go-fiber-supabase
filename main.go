package main

import (
	ds "go-fiber-supabase-backend/domain/datasources"
	repo "go-fiber-supabase-backend/domain/repositories"
	gw "go-fiber-supabase-backend/src/gateways"
	"go-fiber-supabase-backend/src/middlewares"
	sv "go-fiber-supabase-backend/src/services"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {

	// // // remove this before deploy ###################
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// /// ############################################

	app := fiber.New()
	middlewares.Logger(app)
	app.Use(recover.New())
	app.Use(cors.New())

	supabasedb := ds.NewSupabaseDB()

	courseRepo := repo.NewCoursesRepository(supabasedb)
	courseLessonRepo := repo.NewCourseLessonsRepository(supabasedb)
	sv0 := sv.NewCoursesService(courseRepo, courseLessonRepo)

	gw.NewHTTPGateway(app, sv0)

	PORT := os.Getenv("DB_PORT_LOGIN")

	if PORT == "" {
		PORT = "8080"
	}

	app.Listen(":" + PORT)
}
