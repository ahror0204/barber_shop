package repo

type RedisRepositoryStorage interface {
	Set(key, value string) error
	SetWithTTL(key, value string, second int) error
	Get(key string) (interface{}, error)
}
