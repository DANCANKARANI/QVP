package database
import(
	"github.com/go-redis/redis/v8"
)
//connecting to RedisClient
func RedisClient()*redis.Client{
	rdb :=redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return rdb
}