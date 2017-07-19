package main

import (
	"log"
	"os"

	goraph "github.com/gyuho/goraph"
)

func main() {
	// https://github.com/gyuho/goraph/blob/master/testdata/graph_00.png
	g, err := CreateGraphFromJSON("./graph.json", "graph_00")
	if err != nil {
		log.Println(err)
		return
	}
	//log.Println(g)
	var from goraph.StringID = "A"
	var to goraph.StringID = "E"
	ids, _, err := goraph.Dijkstra(g, from, to)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(ids)
}

func CreateGraphFromJSON(file string, name string) (goraph.Graph, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	g, err := goraph.NewGraphFromJSON(f, name)
	if err != nil {
		return nil, err
	}
	return g, nil
}
