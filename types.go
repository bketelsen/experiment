package server

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DataRoot string
}

type Pather interface {
	Path(*Config) string
}

type World struct {
	Name      string
	PersonRef string  `yaml:"person"`
	Person    *Person `yaml:"-"`
}

func (w *World) Path(c *Config, file string) string {
	return filepath.Join(c.DataRoot, file)
}

func (w *World) Load(c *Config) error {
	person := Person{}
	personFile := w.PersonRef + ".yaml"
	fmt.Println(person.Path(c, personFile))
	bb, err := os.ReadFile(person.Path(c, personFile))
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(bb, &person)
	if err != nil {
		return err
	}
	fmt.Println(person)
	w.Person = &person
	return nil
}

type Person struct {
	FirstName string `yaml:"first_name"`
	LastName  string `yaml:"last_name"`
	Age       int
}

func (w *Person) Path(c *Config, file string) string {
	return filepath.Join(c.DataRoot, "person", file)
}
