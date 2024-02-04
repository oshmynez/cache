package main

import (
	"errors"
	"fmt"
)

type Methods interface {
	New()
	Set()
	Get()
	Delete()
}

type Elem struct {
	key   string
	value interface{}
}

type Cache struct {
	array []Elem
}

func (c Cache) New() Cache {
	return c
}

func (c *Cache) Set(key string, value interface{}) {
	for _, item := range c.array {
		if item.key == key {
			item.value = value
			return
		}
	}
	elem := Elem{key, value}
	c.array = append(c.array, elem)
}

func (c *Cache) Get(key string) interface{} {
	for _, item := range c.array {
		if item.key == key {
			return item.value
		}
	}
	return nil
}

func (c *Cache) Delete(key string) (string, error) {
	for _, item := range c.array {
		if item.key == key {
			item = c.array[len(c.array)-1]
			c.array = c.array[:len(c.array)-1]
			return key, nil
		}
	}
	return key, errors.New("there are not any numbers in array")

}

func main() {
	cache := Cache{}.New()

	cache.Set("1", "apple")
	fmt.Printf("%+v\n", cache)

	cache.Set("2", 34)
	fmt.Printf("%+v\n", cache)

	v := cache.Get("2")
	fmt.Printf("%+v\n", v)

	key, status := cache.Delete("1")
	fmt.Println(key, status)
}
