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
	PersonRef string   `yaml:"person"`
	Person    *Person  `yaml:"-"`
	SiteRefs  []string `yaml:"sites"`
	Sites     []*Site  `yaml:"-"`
}

func (w *World) Path(c *Config, file string) string {
	return filepath.Join(c.DataRoot, "world", file)
}

func (w *World) Load(c *Config) error {
	p, err := LoadPerson(c, w.PersonRef)
	if err != nil {
		return err
	}
	w.Person = p
	ss, err := LoadSites(c, w.SiteRefs)
	if err != nil {
		return err
	}
	w.Sites = ss
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

func LoadPerson(c *Config, ref string) (*Person, error) {
	person := &Person{}
	personFile := ref + ".yaml"
	fmt.Println(person.Path(c, personFile))
	bb, err := os.ReadFile(person.Path(c, personFile))
	if err != nil {
		return person, err
	}
	err = yaml.Unmarshal(bb, &person)
	return person, err
}

type Site struct {
	Name         string         `yaml:"name"`
	Environments []*Environment `yaml:"environments"`
}

type Environment struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

func (w *Site) Path(c *Config, file string) string {
	return filepath.Join(c.DataRoot, "site", file)
}
func LoadSites(c *Config, refs []string) ([]*Site, error) {
	fmt.Println(refs)
	sites := []*Site{}
	for _, ref := range refs {
		site := &Site{}
		siteFile := ref + ".yaml"
		fmt.Println(site.Path(c, siteFile))
		bb, err := os.ReadFile(site.Path(c, siteFile))
		if err != nil {
			return sites, err
		}
		err = yaml.Unmarshal(bb, site)
		if err != nil {
			return sites, err
		}
		sites = append(sites, site)
	}
	return sites, nil
}
