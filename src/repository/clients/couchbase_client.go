package clients

import (
	"log"

	"github.com/couchbase/gocb"
)

var (
	bucket *gocb.Bucket
)

func init() {
	cluster, clusterErr := gocb.Connect("couchbase://localhost")
	if clusterErr != nil {
		panic(clusterErr)
	}
	bucketName := "shortener-url"
	cluster.Authenticate(gocb.PasswordAuthenticator{Username: "Administrator", Password: "local1234"})

	var err error
	bucket, err = cluster.OpenBucket(bucketName, "")
	if err != nil {
		panic(err)
	}
	log.Println("db connected")
}

func GetBucket() *gocb.Bucket {
	return bucket
}
