package repository

import (
	"flag"
	"io/ioutil"
	"os"
)

type Repository struct {
	MethodF  *string
	ServiceF *string
	FileF    *string
	IdF      *int64
}

func NewRepository() *Repository {
	return &Repository{MethodF: flag.String("method", "", "a method to call"),
		ServiceF: flag.String("service", "", "a service"),
		FileF:    flag.String("file", "", "a file with config of service"),
		IdF:      flag.Int64("id", 0, "an id of config")}
}

func (repo *Repository) ReadConfigFromFile() ([]byte, error) {
	f, err := os.Open(*repo.FileF)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	config, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (repo *Repository) GetID() int64 {
	return *repo.IdF
}

func (repo *Repository) GetService() string {
	return *repo.ServiceF
}

func (repo *Repository) GetMethod() string {
	return *repo.MethodF
}
