package models

// Path represents the path that should be redirected to the URL.
type Redirect struct {
	Path string `json:"path" yaml:"path"`
	Url  string `json:"url" yaml:"url"`
}

// PathEntity represents a path record in the db.
type RedirectEntity struct {
	Redirect
	Entity
}
