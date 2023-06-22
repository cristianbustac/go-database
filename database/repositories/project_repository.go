package repositories

import "gorm.io/gorm"

type ProjectRepository struct {
	*BaseRepository
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{
		BaseRepository: &BaseRepository{dbInstance: db},
	}
}
