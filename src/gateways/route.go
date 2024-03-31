package gateways

import "github.com/gofiber/fiber/v2"

func GatewayCourses(gateway HTTPGateway, app *fiber.App) {
	api := app.Group("/api/course")

	api.Get("/course-all", gateway.GetAllCourses)
	api.Get("/get-course", gateway.GetCourse)
	api.Get("/lesson-all", gateway.GetAllCourseLessons)
	api.Get("/get-lesson", gateway.GetCourseLesson)
}
