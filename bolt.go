package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

var boltdb *bolt.DB
var open bool

func OpenBolt() error {
	var err error
	// _, filename, _, _ := runtime.Caller(0)
	// dbfile := path.Join(path.Dir(filename), "RecipeBook.db")
	config := &bolt.Options{Timeout: 1 * time.Second}
	// db, err = bolt.Open(dbfile, 0644, config)
	boltdb, err = bolt.Open("/tmp/RecipeBook.db", 0644, config)
	check(err)
	open = true
	return nil
}

func CloseBolt() {
	boltdb.Close()
	open = false
}

func (r *Recipe) AddRecipeBolt() error {
	if !open {
		return fmt.Errorf("db must be open before adding recipe.")
	}
	err := boltdb.Update(func(tx *bolt.Tx) error {
		book, err := tx.CreateBucketIfNotExists([]byte("book"))
		check(err)
		enc, err := json.Marshal(r)
		check(err)
		err = book.Put([]byte(r.Name), enc)
		return err
	})
	return err
}

func GetRecipeBolt(id string) (*Recipe, error) {
	if !open {
		return nil, nil
	}
	var r *Recipe
	err := boltdb.View(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket([]byte("book"))
		k := []byte(id)
		err = json.Unmarshal(b.Get(k), &r)
		if err != nil {
			return err
		}
		return nil
	})
	return r, err
}

func ListBolt() []Recipe {
	var recipes []Recipe
	boltdb.View(func(tx *bolt.Tx) error {
		book := tx.Bucket([]byte("book"))
		if book == nil {
			return nil
		}
		var err error
		err = book.ForEach(func(k []byte, v []byte) error {
			var recipe *Recipe
			err := json.Unmarshal(v, &recipe)
			check(err)
			recipes = append(recipes, *recipe)
			return nil
		})
		check(err)
		return nil
	})
	return recipes
}

func DeleteRecipeBolt(id string) error {
	err := boltdb.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte("book")).Delete([]byte(id))
	})
	check(err)
	return nil
}
