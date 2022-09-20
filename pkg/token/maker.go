package token

import "time"

type Maker interface {
	// create a new token for a spacific user id, user name and duration
	CreateToken(username string, duration time.Duration) (string, *Payload, error)

	VerifyToken(token string) (*Payload, error)
}
