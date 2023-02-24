package response

import "github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"

type (
	FindAllActiveUserResponse struct {
		Results ActiveUserResults `json:"result"`

		Errors []errors.ErrorInfo
	}

	FindByIdActiveUserResponse struct {
		Result ActiveUserResult `json:"result"`

		Errors []errors.ErrorInfo
	}

	CreateActiveUserResponse struct {
		Result ActiveUserResult `json:"results"`

		Errors []errors.ErrorInfo
	}

	UpdateActiveUserResponse struct {
		Result ActiveUserWithBlogResults `json:"results"`

		Errors []errors.ErrorInfo
	}

	DeleteActiveUserResponse struct {
		Errors []errors.ErrorInfo
	}

	LoginUserResponse struct {
		Result LoginUserResult

		Errors []errors.ErrorInfo
	}
)

type (
	ActiveUserResult struct {
		User ActiveUserEntity `json:"user"`
	}
	ActiveUserResults struct {
		Users []ActiveUserEntity `json:"users"`
	}

	ActiveUserWithBlogResults struct {
		User ActiveUserEntities
	}

	// HistoreUserResult struct {
	// 	Student *HistoryUserEntity `json:"student"`
	// }
	// HistoreUserResults struct {
	// 	Students []*HistoryUserEntity `json:"students"`
	// }

	LoginUserResult struct {
		Status string `json:"status"`
		Token  string `json:"token"`
	}
)

type (
	ActiveUserEntity struct {
		UserId   int    `json:"user_id"`
		Email    string `json:"email"`
		UserName string `json:"user_name"`
		Password string `json:"password"`
		UserRoll string `json:"user_roll"`
		Blog     []ActiveBlogEntity
		Option   Options `json:"option"`
	}
	ActiveUserEntities struct {
		UserId   int    `json:"user_id"`
		Email    string `json:"email"`
		UserName string `json:"user_name"`
		Password string `json:"password"`
		UserRoll string `json:"user_roll"`
		Blogs    []ActiveBlogEntity
		Option   Options `json:"option"`
	}
)
