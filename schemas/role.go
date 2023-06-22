package schemas

type Role struct {
	Name string `json:"name"`
}

type RoleUpdate struct {
	Name *string `json:"name"`
}
