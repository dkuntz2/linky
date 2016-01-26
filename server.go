package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"net/url"
	"time"
)

type link struct {
	uri        string
	created_at time.Time
}

func main() {
	var l link
	l.created_at = time.Now()
	uri, err := url.Parse("https://don.kuntz.co")
	if err != nil {
		panic(err)
	}
	l.uri = uri.String()

	db, err := bolt.Open("links.db", 0600, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("links"))
		if err != nil {
			panic(err)
		}

		b.Put([]byte(l.created_at.String()), []byte(l.uri))

		return nil
	})
	if err != nil {
		panic(err)
	}

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("links"))

		b.ForEach(func(key, value []byte) error {
			fmt.Printf("%s @ %v\n", value, string(key))

			return nil
		})

		return nil
	})
}
