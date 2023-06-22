package repositories

import (
	"gorm.io/gorm"
)

type RelationRepository struct {
	dbInstance *gorm.DB
}

func (rep *RelationRepository) CreateRelation(
	primary_id string,
	secondary_id string,
	entity_primary interface{},
	entity_secondary interface{},
	asociation string,
) error {
	if err := rep.dbInstance.First(entity_primary, primary_id).Error; err != nil {
		return err
	}
	if err := rep.dbInstance.First(entity_secondary, secondary_id).Error; err != nil {
		return err
	}
	rep.dbInstance.Model(entity_primary).Association(asociation).Append(entity_secondary)
	return nil
}

func (rep *RelationRepository) GetAllRelations(
	entityPrimary interface{},
	skip int,
	limit int,
	association string,
) (*gorm.DB, error) {
	db := rep.dbInstance.Preload(association).Offset(skip).Limit(limit).Find(entityPrimary)
	if db.Error != nil {
		return nil, db.Error
	}
	return db, nil
}

func (rep *RelationRepository) DeleteRelation(
	primary_id string,
	secondary_id string,
	entity_primary interface{},
	entity_secondary interface{},
	asociation string,
) error {
	if err := rep.dbInstance.First(entity_primary, primary_id).Error; err != nil {
		return err
	}
	if err := rep.dbInstance.First(entity_secondary, secondary_id).Error; err != nil {
		return err
	}
	if err := rep.dbInstance.Model(entity_primary).Association(asociation).Delete(entity_secondary); err != nil {
		return err
	}
	return nil
}
