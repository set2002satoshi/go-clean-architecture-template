package request

type (
	UserFindByIdRequest struct {
		ID int `json:"id"`
	}
	UserDeleteRequest struct {
		ID int `json:"id"`
	}
	UserCreateRequest struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Password string `json:"pass"`
		Roll     string `json:"roll"`
	}

	UserUpdateRequest struct {
		ID       int    `json:"id"`
		Email    string `json:"email"`
		Name     string `json:"name"`
		Password string `json:"password"`
		Roll     string `json:"roll"`
		Revision int    `json:"revision"`
	}
)

type (
	UserLoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"pass"`
	}
)
