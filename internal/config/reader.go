package config

import (
	"io"
)

// Reader reads configuration for determine http redirects ruls
type Reader interface {
	Read() (*Config, error)
	SetLoger(io.Writer)
	SetReader(io.Reader)
}

// ReadRecords combines the roles of Reader of records from config files and Logger.
type ReadRecords struct {
	Reader    io.Reader
	LogWriter io.Writer
}

func (r *ReadRecords) SetLoger(w io.Writer) {
	r.LogWriter = w
}

func (r *ReadRecords) SetReader(reader io.Reader) {
	r.Reader = reader
}
