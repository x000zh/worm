package utils

import (
	"bytes"
	"compress/gzip"
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)


// 压缩存储
func RedisGzipSet(ctx context.Context, redisClient *redis.Client, key string, data []byte, expired time.Duration) error {

	buf := new(bytes.Buffer)
	gzipW := gzip.NewWriter(buf)
	gzipW.Write(data)
	gzipW.Close()

	zipData := buf.Bytes()
	
	_, err := redisClient.Set(context.Background(), key, zipData, time.Minute).Result()
	return err
}


//读取并解压
func RedisGzipGet(ctx context.Context, redisClient *redis.Client, key string) ([]byte, error) {
	s, err := redisClient.Get(context.Background(), key).Bytes()
	if nil != err {
		return nil, err
	}
	
	gzipR, err := gzip.NewReader(bytes.NewBuffer(s))
	
	if nil != err {
		return nil, err
	}

	out := new(bytes.Buffer);

	bs := make([]byte, len(s))
	for  {
		n,err :=  gzipR.Read(bs)
		if 0 < n {
			out.Write(bs[0:n])
		}
		if nil != err {
			return nil, err
		}
	}
	return out.Bytes(), nil
}
