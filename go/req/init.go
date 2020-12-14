package req

import (
	"../conf"
)

// Req holds req package information and is exported
type Req struct {
	Token   string
	BaseURL string
	// Client  *http.Client
}

// Init is where it all starts
func Init(c conf.Conf) (s Req) {
	s = Req{
		Token:   c.Token,
		BaseURL: c.BaseURL,
	}
	return
}
