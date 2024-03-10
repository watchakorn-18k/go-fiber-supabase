package repositories

import (
	"encoding/json"
	"fmt"
	. "go-fiber-supabase-backend/domain/datasources"
	"go-fiber-supabase-backend/domain/entities"

	"github.com/supabase/postgrest-go"
)

type CoursesRepository struct {
	Client *postgrest.QueryBuilder
}

type ICoursesRepository interface {
	FindCourses() ([]entities.CourseModel, error)
	FindCourse(name string) (entities.CourseModel, error)
}

func NewCoursesRepository(db *SupabaseDB) ICoursesRepository {
	return &CoursesRepository{
		Client: db.Client.From("courses"),
	}
}

func (repo *CoursesRepository) FindCourses() ([]entities.CourseModel, error) {
	data, _, err := repo.Client.Select("*", "exact", false).Execute()
	dataMap := make([]entities.CourseModel, 0)
	json.Unmarshal([]byte(data), &dataMap)
	return dataMap, err
}

func (repo *CoursesRepository) FindCourse(name string) (entities.CourseModel, error) {
	data, _, err := repo.Client.Select("*", "", false).Eq("name", name).ExecuteString()
	if err != nil {
		fmt.Println(err)
		return entities.CourseModel{}, err
	}
	dataMap := []entities.CourseModel{}
	json.Unmarshal([]byte(data), &dataMap)
	return dataMap[0], err
}
