package response

import "github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"

type (
	FindAllActive{}Response struct {
		Result Active{}Results `json:"result"`

		Errors []errors.ErrorInfo
	}

	FindByIDActive{}Response struct {
		Result ActiveUserWith{}Result `json:"result"`

		Errors []errors.ErrorInfo
	}

	CreateActive{}Response struct {
		Result Active{}Result `json:"results"`

		Errors []errors.ErrorInfo
	}

	UpdateActive{}Response struct {
		Result Active{}Result `json:"results"`

		Errors []errors.ErrorInfo
	}
	DeleteActive{}Response struct {
		Errors []errors.ErrorInfo
	}
)

type (
	Active{}Result struct {
		{} Active{}Entity `json:"{}"`
	}
	ActiveUserWith{}Result struct {
		{} ActiveUserEntities `json:"user_with_{}"`
	}
	Active{}Results struct {
		{}s []Active{}Entity `json:"{}s"`
	}

	// History{}Result struct {
	// 	Student *HistoryUserEntity `json:"student"`
	// }
	// History{}Results struct {
	// 	Students []*History{}Entity `json:"students"`
	// }

)

type (
	Active{}Entity struct {
		{}Id      int                     `json:""`
		Option      Options                 `json:"option"`
	}
)
