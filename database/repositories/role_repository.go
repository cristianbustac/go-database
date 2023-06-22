package repositories

import "gorm.io/gorm"

type RoleRepository struct {
	*BaseRepository
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{
		BaseRepository: &BaseRepository{dbInstance: db},
	}
}
