package schemas

type User struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
}

type UserUpdate struct {
	Name     *string `json:"name"`
	LastName *string `json:"lastname"`
}
