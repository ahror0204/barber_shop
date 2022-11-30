package response

import "time"

type RedisRepositoryStorage interface {
	Set(key, value string) error
	SetWithTTL(key, value string, second time.Duration) error 
	Get(key string) (interface{}, error)
}