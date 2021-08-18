package diaries

// Request Parameters

type GetListInput struct {
	Order string `json:"order" form:"order"`
	Page  int    `json:"page" form:"page"`
}

type AddInput struct {
	User  int    // domain.User
	title string `json:"title" form:"title"`
	Body  string `json:"body" form:"body"`
}

type GetInput struct {
	id int `json:"id"`
}

type UpdateInput struct {
	User  int    // domain.User
	ID    int    //`validate:"gt=0"`
	title string `json:"title" form:"title"`
	Body  string `json:"body" form:"body"`
	// IsDeleted       bool `json:"isDeleted" form:"isDeleted"`
}

type DeleteInput struct {
	ID   int `json:"id"`
	User int
}

type Like struct {
	ID   int `json:"id"`
	User int
}

type DeLike struct {
	ID   int `json:"id"`
	User int
}
