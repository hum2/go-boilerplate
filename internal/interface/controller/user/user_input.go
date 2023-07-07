package user

// postUserInput is a request body parameter of POST /users
type postUserInput struct {
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
}

// putUserInput is a request body parameter of PUT /users
type putUserInput struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
}

// deleteUserInput is a request body parameter of DELETE /users
type deleteUserInput struct {
	ID string `json:"id"`
}
