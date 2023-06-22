package interfaces

type CRUDRepository interface {
	Create(entity interface{}) (interface{}, error)
	GetAll(entities []interface{}, skip int, limit int) ([]interface{}, error)
	GetById(id string, entity interface{}) (interface{}, error)
	Update(id string, entity interface{}, new_entity interface{}) (interface{}, error)
	Delete(id string, entity interface{}) error
}

// En Go, una interfaz se define como un tipo con una lista de métodos,
// sin proporcionar la implementación de esos métodos.
// Cualquier tipo que implemente todos los métodos especificados en una
