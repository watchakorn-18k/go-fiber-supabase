package gateways

import (
	"go-fiber-supabase-backend/domain/entities"

	"github.com/gofiber/fiber/v2"
)

func (h HTTPGateway) GetAllCourses(ctx *fiber.Ctx) error {

	data, err := h.CoursesService.GetAllCourses()
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: "cannot get all courses data"})
	}
	return ctx.Status(fiber.StatusOK).JSON(data)
}

func (h HTTPGateway) GetCourse(ctx *fiber.Ctx) error {
	params := ctx.Queries()
	name := params["name"]
	data, err := h.CoursesService.GetCourse(name)
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: "cannot get course data"})
	}
	return ctx.Status(fiber.StatusOK).JSON(data)
}
