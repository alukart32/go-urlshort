package config

import (
	"fmt"

	"alukart32.com/urlshort/internal/models"
)

// Config is a set of all possible configurations for redirecting an http request.
type Config struct {
	Redirections []models.Redirect `json:"paths" yaml:"paths,flow"`
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
func (c *Config) AddRedirect(path string, url string) {
	newPath := models.Redirect{
		Path: path,
		Url:  url,
	}
	if c.Redirections == nil {
		c.Redirections = make([]models.Redirect, 1)
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
