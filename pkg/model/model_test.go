package model

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Maintainer Validate", func() {
	var mt Maintainer
	BeforeEach(func() {
		mt = Maintainer{
			Name:  "maintainer 1",
			Email: "m1@gmail.com",
		}
	})
	It("can't have empty name", func() {
		mt.Name = ""
		Expect(mt.Validate()).NotTo(Succeed())
	})
	It("can't have no email", func() {
		mt.Email = ""
		Expect(mt.Validate()).NotTo(Succeed())
	})
	It("can't have invalid email", func() {
		mt.Email = "email"
		Expect(mt.Validate()).NotTo(Succeed())
	})
})

var _ = Describe("Metadata Validate", func() {
	var md Metadata
	BeforeEach(func() {
		md = Metadata{
			Title:   "Valid App 1",
			Version: "0.0.1",
			Maintainers: []Maintainer{
				{
					Name:  "firstmaintainer app1",
					Email: "firstmaintainer@hotmail.com",
				},
				{
					Name:  "secondmaintainer app1",
					Email: "secondmaintainer@gmail.com",
				}},
			Company:     "Random Inc.",
			Website:     "https://website.com",
			Source:      "https://github.com/random/repo",
			License:     "Apache-2.0",
			Description: "### Interesting Title\nSome application content, and description",
		}
	})
	It("can't have empty title", func() {
		md.Title = ""
		Expect(md.Validate()).NotTo(Succeed())
	})
	It("can't have empty version", func() {
		md.Version = ""
		Expect(md.Validate()).NotTo(Succeed())
	})
	It("can't have invalid version", func() {
		md.Version = "100"
		Expect(md.Validate()).NotTo(Succeed())
	})
	It("can't have no maintainers", func() {
		md.Maintainers = make([]Maintainer, 0)
		Expect(md.Validate()).NotTo(Succeed())
	})
	It("can't have empty company", func() {
		md.Company = ""
		Expect(md.Validate()).NotTo(Succeed())
	})
	It("can't have empty website", func() {
		md.Website = ""
		Expect(md.Validate()).NotTo(Succeed())
	})
	It("can't have invalid website", func() {
		md.Website = "html"
		Expect(md.Validate()).NotTo(Succeed())
	})
	It("can't have empty source", func() {
		md.Source = ""
		Expect(md.Validate()).NotTo(Succeed())
	})
	It("can't have invalid source", func() {
		md.Source = "github"
		Expect(md.Validate()).NotTo(Succeed())
	})
	It("can't have empty license", func() {
		md.License = ""
		Expect(md.Validate()).NotTo(Succeed())
	})
	It("can't have empty description", func() {
		md.Description = ""
		Expect(md.Validate()).NotTo(Succeed())
	})
})
