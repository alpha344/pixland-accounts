package gateway

import "github.com/alpha344/pixlandx/mailer"

type Container struct {
	Authenticator Authenticator
	Mailer        mailer.Mailer
	Storage       Storage
}


