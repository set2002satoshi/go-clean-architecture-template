package response

import (
	"time"
)

type (
	Options struct {
		Revision  int       `json:"revision"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
