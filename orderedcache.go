package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type orderedCache struct {
	path string
	size int

	dirty   bool
	items   []string
	seekmap map[string]bool
}

func newOrderedCache(path string, size int) *orderedCache {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("Creating new cache")
		data = []byte("[]")
	}

	var items []string
	err = json.Unmarshal(data, &items)
	if err != nil {
		log.Fatalln("Invalid cache content: ", err)
	}

	if len(items) > size {
		data = data[:size]
	}

	seekmap := make(map[string]bool, size)
	for _, item := range items {
		seekmap[item] = true
	}

	return &orderedCache{
		path:    path,
		size:    size,
		items:   items,
		seekmap: seekmap,
	}
}

func (o *orderedCache) add(item string) {
	o.dirty = true
	o.seekmap[item] = true
	o.items = append(o.items, item)
	if len(o.items) > o.size {
		delete(o.seekmap, o.items[0])
		o.items = o.items[1:]
	}
}

func (o *orderedCache) exists(item string) bool {
	return o.seekmap[item]
}

func (o *orderedCache) commit() {
	if !o.dirty {
		return
	}
	o.dirty = false

	data, err := json.Marshal(o.items)
	if err != nil {
		log.Fatalln("Failed marshaling items: ", err)
	}

	err = ioutil.WriteFile(o.path, data, 0644)
	if err != nil {
		log.Fatalln("Failed writing cache to disk: ", err)
	}
}
