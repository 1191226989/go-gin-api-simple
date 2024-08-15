package captcha

import (
	"time"

	"go-gin-api-simple/internal/repository/redis"
)

type StoreRedis struct {
	store redis.Repo
}

// Set sets the digits for the captcha id.
func (sr *StoreRedis) Set(id string, value string) error {
	return sr.store.Set(id, value, time.Second*60*5)
}

// Get returns stored digits for the captcha id. Clear indicates
// whether the captcha must be deleted from the store.
func (sr *StoreRedis) Get(id string, clear bool) string {
	value, err := sr.store.Get(id)
	if err != nil {
		return ""
	}
	if clear {
		sr.store.Del(id)
	}
	return value
}

// Verify captcha's answer directly
func (sr *StoreRedis) Verify(id, answer string, clear bool) bool {
	value, err := sr.store.Get(id)
	if err != nil {
		return false
	}
	if value != answer {
		return false
	}
	if clear {
		sr.store.Del(id)
	}
	return true
}

func NewStoreRedis(rr redis.Repo) *StoreRedis {
	return &StoreRedis{
		store: rr,
	}
}
