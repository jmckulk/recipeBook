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
	check(err)
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
		check(err)
		enc, err := json.Marshal(r)
		check(err)
		err = book.Put([]byte(r.Name), enc)
		return err
	})
	return err
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
		err = json.Unmarshal(b.Get(k), &r)
		// r, err = decode(b.Get(k))
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

func DeleteRecipe(id string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte("book")).Delete([]byte(id))
	})
	check(err)
	return nil
}

func UpdateTime(id, time string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		var err error
		var recipe *Recipe
		book := tx.Bucket([]byte("book"))
		err = json.Unmarshal(book.Get([]byte(id)), &recipe)
		// recipe, err := decode(book.Get([]byte(id)))
		if recipe == nil || err != nil {
			log.Println("Unable to Update Time. Check to make sure recipe is in book.")
			return err
		} else {
			recipe.CookTime = time
			newRecipe, err := json.Marshal(recipe)
			check(err)
			err = book.Put([]byte(recipe.Name), newRecipe)
		}
		return err
	})
	return err
}

func UpdateIngredientList(id string, ingredient Ingredient) error {
	err := db.Update(func(tx *bolt.Tx) error {
		var err error
		var recipe *Recipe
		book := tx.Bucket([]byte("book"))
		err = json.Unmarshal(book.Get([]byte(id)), &recipe)
		// recipe, err := decode(book.Get([]byte(id)))
		if recipe == nil || err != nil {
			log.Println("Unable to update ingredient list. Check to make sure recipe is in book.")
			return err
		} else {
			recipe.IngredientList = append(recipe.IngredientList, ingredient)
			newRecipe, err := json.Marshal(recipe)
			check(err)
			err = book.Put([]byte(recipe.Name), newRecipe)
		}
		return err
	})
	return err
}
