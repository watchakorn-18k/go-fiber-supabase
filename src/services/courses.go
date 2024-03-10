package services

import (
	"go-fiber-supabase-backend/domain/entities"
	"go-fiber-supabase-backend/domain/repositories"
)

type coursesService struct {
	CoursesRepository repositories.ICoursesRepository
}

type ICoursesService interface {
	GetAllCourses() (entities.ResponseModel, error)
	GetCourse(name string) (entities.ResponseModel, error)
}

func NewCoursesService(repo0 repositories.ICoursesRepository) ICoursesService {
	return &coursesService{
		CoursesRepository: repo0,
	}
}

func (sv coursesService) GetAllCourses() (entities.ResponseModel, error) {
	courseData, err := sv.CoursesRepository.FindCourses()
	if err != nil {
		return entities.ResponseModel{}, err
	}

	return entities.ResponseModel{Data: courseData, Message: "ok"}, nil
}

func (sv coursesService) GetCourse(name string) (entities.ResponseModel, error) {
	data, err := sv.CoursesRepository.FindCourse(name)
	if err != nil {
		return entities.ResponseModel{Message: "not found"}, err
	}
	return entities.ResponseModel{Message: "ok", Data: data}, nil
}