package model

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"

	"github.com/blang/semver/v4"
)

var emailPattern = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type MetadataWithID struct {
	ID   string    `yaml:"id"`
	Data *Metadata `yaml:"metadata" binding:"required"`
}

func NewMetadataWithID(id string, data *Metadata) *MetadataWithID {
	return &MetadataWithID{
		ID:   id,
		Data: data,
	}
}

type Metadata struct {
	Title       string       `yaml:"title"`
	Version     string       `yaml:"version"`
	Maintainers []Maintainer `yaml:"maintainers"`
	Company     string       `yaml:"company"`
	Website     string       `yaml:"website"`
	Source      string       `yaml:"source"`
	License     string       `yaml:"license"`
	Description string       `yaml:"description"`
}

func (md *Metadata) Validate() error {
	if md.Title == "" {
		return errors.New("missing field 'title'")
	}
	if md.Version == "" {
		return errors.New("missing field 'version'")
	}
	if _, err := semver.Parse(md.Version); err != nil {
		return fmt.Errorf("field 'version' is an invalid semantic version: %v", err)
	}
	if len(md.Maintainers) == 0 {
		return errors.New("missing field 'maintainers'")
	}
	for _, m := range md.Maintainers {
		if err := m.Validate(); err != nil {
			return err
		}
	}
	if md.Company == "" {
		return errors.New("missing field 'company'")
	}
	if md.Website == "" {
		return errors.New("missing field 'website'")
	}
	if _, err := url.ParseRequestURI(md.Website); err != nil {
		return fmt.Errorf("field 'website' is an invalid url: %v", err)
	}
	if md.Source == "" {
		return errors.New("missing field 'source'")
	}
	if _, err := url.ParseRequestURI(md.Source); err != nil {
		return fmt.Errorf("field 'source' is an invalid url: %v", err)
	}
	if md.License == "" {
		return errors.New("missing field 'license'")
	}
	if md.Description == "" {
		return errors.New("missing field 'description'")
	}
	return nil
}

type Maintainer struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}

func (mt *Maintainer) Validate() error {
	if mt.Name == "" {
		return errors.New("missing field 'maintainers.name'")
	}
	if mt.Email == "" {
		return errors.New("missing field 'maintainers.email'")
	}
	//https://docs.isitarealemail.com/how-to-validate-email-addresses-in-golang
	if !emailPattern.MatchString(mt.Email) {
		return errors.New("incorrect email format")
	}
	return nil
}
