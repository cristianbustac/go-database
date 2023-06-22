package repositories

import "gorm.io/gorm"

type UsersRolesRepository struct {
	*RelationRepository
}

func NewUsersRolesRepository(db *gorm.DB) *UsersRolesRepository {
	return &UsersRolesRepository{
		RelationRepository: &RelationRepository{dbInstance: db},
	}
}
