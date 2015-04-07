package main

import (
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
	"io/ioutil"
	"log"
)

func main() {

	auth, err := aws.EnvAuth()
	if err != nil {
		log.Fatal(err)
	}
	client := s3.New(auth, aws.USEast)
	resp, err := client.ListBuckets()

	if err != nil {
		log.Fatal(err)
	}

	for _, b := range resp.Buckets {
		m, err := b.GetBucketContents()
		if err != nil {
			log.Fatal(err)
		}
		for k, _ := range *m {
			data, err := b.Get(k)
			if err != nil {
				log.Println(err)
				continue
			}
			if err := ioutil.WriteFile(k, data, 0644); err != nil {
				log.Fatal(err)
			}
			if err := b.Del(k); err != nil {
				log.Println(err)
			}
		}
	}

}
