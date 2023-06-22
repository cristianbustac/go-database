package repositories

import (
	"gorm.io/gorm"
)

// En Go, una interfaz se define como un tipo con una lista de métodos,
// sin proporcionar la implementación de esos métodos.
// Cualquier tipo que implemente todos los métodos especificados en una

type BaseRepository struct {
	dbInstance *gorm.DB
}

func (rep *BaseRepository) Create(entity interface{}) (interface{}, error) {
	if err := rep.dbInstance.Create(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (rep *BaseRepository) GetAll(entities interface{}, skip int, limit int) (*gorm.DB, error) {
	entities_new := rep.dbInstance.Find(entities).Offset(skip).Limit(limit)
	if entities_new.Error != nil {
		return nil, entities_new.Error
	}
	return entities_new, nil
}

func (rep *BaseRepository) GetById(id string, entity interface{}) (interface{}, error) {
	if err := rep.dbInstance.First(entity, id).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (rep *BaseRepository) Update(id string, entity interface{}, new_entity interface{}) (interface{}, error) {
	if err := rep.dbInstance.First(entity, id).Error; err != nil {
		return nil, err
	}
	if err := rep.dbInstance.Model(entity).Updates(new_entity).Error; err != nil {
		return nil, err
	}
	if err := rep.dbInstance.First(entity, id).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (rep *BaseRepository) Delete(id string, entity interface{}) error {
	if err := rep.dbInstance.First(entity, id).Error; err != nil {
		return err
	}
	if err := rep.dbInstance.Delete(entity).Error; err != nil {
		return err
	}
	return nil
}
