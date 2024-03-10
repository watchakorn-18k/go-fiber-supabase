package gateways

import (
	service "go-fiber-supabase-backend/src/services"

	"github.com/gofiber/fiber/v2"
)

type HTTPGateway struct {
	CoursesService service.ICoursesService
}

func NewHTTPGateway(app *fiber.App, courses service.ICoursesService) {
	gateway := &HTTPGateway{
		CoursesService: courses,
	}

	GatewayCourses(*gateway, app)
}
