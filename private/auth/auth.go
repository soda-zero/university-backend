package auth

import "time"

func (s Session) isExpired() bool {
	return s.Expiration.Before(time.Now())
}

var sessions = make(map[string]Session)
