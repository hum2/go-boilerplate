package user

// getUserOutput is a response of GET /users
type getUserOutput struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
}

// postUserOutput is a response of POST /users
type postUserOutput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
