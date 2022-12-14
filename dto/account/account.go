package accountdto

type RequestAccount struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ResponseAccount struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
