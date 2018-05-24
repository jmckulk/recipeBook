package main

import (
	"encoding/json"
	"fmt"
	"log"
	"path"
	"runtime"
	"time"

	"github.com/boltdb/bolt"
)

var db *bolt.DB
var open bool

func Open() error {
	var err error
	_, filename, _, _ := runtime.Caller(0)
	dbfile := path.Join(path.Dir(filename), "RecipeBook.db")
	config := &bolt.Options{Timeout: 1 * time.Second}
	db, err = bolt.Open(dbfile, 0644, config)
	if err != nil {
		log.Println("Open db")
		panic(err)
	}
	open = true
	return nil
}

func Close() {
	db.Close()
	open = false
}

func (r *Recipe) AddRecipe() error {
	if !open {
		return fmt.Errorf("db must be open before adding recipe.")
	}
	err := db.Update(func(tx *bolt.Tx) error {
		book, err := tx.CreateBucketIfNotExists([]byte("book"))
		if err != nil {
			log.Fatal(err)
		}
		enc, err := r.encode()
		if err != nil {
			log.Fatal(err)
		}
		err = book.Put([]byte(r.Name), enc)
		return err
	})
	return err
}

func (r *Recipe) encode() ([]byte, error) {
	enc, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return enc, nil
}

func decode(data []byte) (*Recipe, error) {
	var r *Recipe
	err := json.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetRecipe(id string) (*Recipe, error) {
	if !open {
		return nil, nil
	}
	var r *Recipe
	err := db.View(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket([]byte("book"))
		k := []byte(id)
		r, err = decode(b.Get(k))
		if err != nil {
			return err
		}
		return nil
	})
	return r, err
}

func List() []Recipe {
	var recipes []Recipe
	db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte("book")).Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			recipe, _ := decode(v)
			recipes = append(recipes, *recipe)
		}
		return nil
	})
	return recipes
}

func DeleteRecipe(id string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte("book")).Delete([]byte(id))
	})
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
