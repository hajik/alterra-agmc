package requests

type UserLogin struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserStoreOne struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserListAll struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserUpdate struct {
	Name     string `json:"name" query:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserDelete struct {
	Name string `json:"name"`
}
