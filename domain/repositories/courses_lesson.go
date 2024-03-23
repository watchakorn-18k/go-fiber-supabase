package repositories

import (
	"encoding/json"
	. "go-fiber-supabase-backend/domain/datasources"
	"go-fiber-supabase-backend/domain/entities"

	"github.com/supabase/postgrest-go"
)

type CourseLessonsRepository struct {
	Client *postgrest.QueryBuilder
}

type ICourseLessonsRepository interface {
	FindAllCourseLessons() ([]entities.CourseLessonsModel, error)
}

func NewCourseLessonsRepository(db *SupabaseDB) ICourseLessonsRepository {
	return &CourseLessonsRepository{
		Client: db.Client.From("courseLessons"),
	}
}

func (repo *CourseLessonsRepository) FindAllCourseLessons() ([]entities.CourseLessonsModel, error) {
	data, _, err := repo.Client.Select("*", "exact", false).Execute()
	dataMap := make([]entities.CourseLessonsModel, 0)
	json.Unmarshal([]byte(data), &dataMap)
	return dataMap, err
}
