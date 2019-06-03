package db

import (
	"github.com/go-redis/redis"
	"os"
)

var (
	client = redis.NewClient(&redis.Options{
		Addr:	os.Getenv("REDIS_ADDR"),
		Password:	os.Getenv("REDIS_PW"),
		DB:	0,
	})
)


// SetFile sets fileid -> filename in redis
func SetFile(filename *string,fileID *string) (er error) {
    err := client.HSet(*fileID,"filename",*filename).Err()
    if err !=nil{
        print(err)
        return err
    }
    return nil
}


// GetFileName returns original filename of a fileID
func GetFileName(fileID *string)(filename string,er error){
    val,err := client.HGet(*fileID,"filename").Result()
    // error
    if err == nil{
        // no key
        if err == redis.Nil{
            return "",nil
        }
        // other error
        print(err)
        return "",err
    }
    return val,nil
}



// DelEntry removes a fileID record in redis
func DelEntry(fileID *string)(er error){
    err := client.Del(*fileID).Err()
    if err != nil {
        print(err)
        return err
    }
    return nil
}
