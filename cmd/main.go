package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bketelsen/server"
	"gopkg.in/yaml.v2"
)

func main() {
	config := &server.Config{
		DataRoot: "/Users/bjk/src/github.com/bketelsen/driab/cue-experiments/server/data/",
	}
	world := server.World{}
	worldPath := world.Path(config, "world.yaml")
	bb, err := os.ReadFile(worldPath)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(bb, &world)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(world)
	err = world.Load(config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(world.Person.FirstName)

}
