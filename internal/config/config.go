package config

import "fmt"

type Path struct {
	Path string `json:"path" yaml:"path"`
	Url  string `json:"url" yaml:"url"`
}

// Config is a set of all possible configurations for redirecting an http request.
type Config struct {
	Redirections []Path `json:"paths" yaml:"paths,flow"`
}

// PathsToMap converts the slice c.Paths to map[Path]Url.
func (c *Config) PathsToMap() map[string]string {
	if len(c.Redirections) == 0 {
		return nil
	}
	m := make(map[string]string, len(c.Redirections))
	for _, v := range c.Redirections {
		m[v.Path] = v.Url
	}
	return m
}

// AppendPath add a new value to the variable paths.
func (c *Config) AddPath(p string, u string) {
	newPath := Path{
		Path: p,
		Url:  u,
	}
	if c.Redirections == nil {
		c.Redirections = make([]Path, 0)
	}
	c.Redirections = append(c.Redirections, newPath)
}

const SPACE = "   "

// Print displays the contents of the config
func (c *Config) Print() {
	fmt.Printf("\nconfig\n")
	fmt.Printf("%spaths:", SPACE)
	for i := range c.Redirections {
		fmt.Printf("\n\t- path: %v\n\t  url: %v", c.Redirections[i].Path, c.Redirections[i].Url)
	}
	fmt.Println()
}
