package schemas

type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProjectUpdate struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}
