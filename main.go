import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
    client := redis.NewClient(&redis.Options{
        Addr:	  "localhost:6379",
        Password: "", // no password set
        DB:		  0,  // use default DB
    })
}


opt, err := redis.ParseURL("redis://<user>:<pass>@localhost:6379/<db>")
if err != nil {
	panic(err)
}

client := redis.NewClient(opt)

ctx := context.Background()

err := client.Set(ctx, "foo", "bar", 0).Err()
if err != nil {
    panic(err)
}

val, err := client.Get(ctx, "foo").Result()
if err != nil {
    panic(err)
}
fmt.Println("foo", val)

ssion := map[string]string{"name": "John", "surname": "Smith", "company": "Redis", "age": "29"}
for k, v := range session {
    err := client.HSet(ctx, "user-session:123", k, v).Err()
    if err != nil {
        panic(err)
    }
}

userSession := client.HGetAll(ctx, "user-session:123").Val()
fmt.Println(userSession)

