package redis
//
//import (
//	"github.com/go-redis/redis"
//)
//
//var redisDb *redis.Client

//func initRedis() {
//	 redisDb = redis.NewClient(&redis.Options{
//		Addr:     "127.0.0.1:6379",
//		Password: "",
//		DB:       0,
//	})
//}
//
//func main() {
//
//	initRedis()
//	fmt.Println("redis初始化成功...")
//	redisZset()
//}
//
//func redisZset() {
//
//	key := "language_rank"
//	languages := []*redis.Z{
//		&redis.Z{Score: 90.0, Member: "Golang"},
//		&redis.Z{Score: 98.0, Member: "Java"},
//		&redis.Z{Score: 95.0, Member: "Python"},
//		&redis.Z{Score: 97.0, Member: "C/C++"},
//	}
//	//redisDb.ZAdd(key, languages...)
//	fmt.Println(key, languages)
//}