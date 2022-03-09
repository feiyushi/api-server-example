package main

import (
	"apiserver/pkg/model"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v2"
)

var (
	validPayload1 = model.MetadataWithID{
		Data: &model.Metadata{
			Title:   "Valid App 1",
			Version: "0.0.1",
			Maintainers: []model.Maintainer{
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
		},
	}

	invalidPayloadMissingField = model.MetadataWithID{
		Data: &model.Metadata{
			Title: "Valid App 1",
			// Version: "0.0.1",
			Maintainers: []model.Maintainer{
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
		},
	}
)

var _ = Describe("Server", func() {
	var server *httptest.Server

	BeforeEach(func() {
		server = httptest.NewServer(setUpRouter())
	})

	AfterEach(func() {
		server.Close()
	})

	It("can save a metadata with a unique id and get the metadata by id", func() {
		id := "1"
		p, err := yaml.Marshal(validPayload1)
		Expect(err).To(BeNil())
		path := fmt.Sprintf("%s/metadata/%s", server.URL, id)

		// PUT
		put, err := http.NewRequest(http.MethodPut,
			path,
			bytes.NewBuffer(p))
		Expect(err).To(BeNil())

		resp, err := http.DefaultClient.Do(put)
		Expect(err).To(BeNil())

		Expect(resp.StatusCode).To(Equal(http.StatusCreated))

		// GET
		get, err := http.Get(path)
		Expect(err).To(BeNil())
		Expect(get.StatusCode).To(Equal(http.StatusOK))

		defer get.Body.Close()
		body, err := io.ReadAll(get.Body)
		Expect(err).To(BeNil())

		var yml model.MetadataWithID
		Expect(yaml.Unmarshal(body, &yml)).To(Succeed())
		Expect(yml.Data).To(Equal(validPayload1.Data))
		Expect(yml.ID).To(Equal(id))

		// GET: List
		list, err := http.Get(fmt.Sprintf("%s/metadata?company=%s", server.URL, `Random%20Inc.`))
		Expect(err).To(BeNil())
		Expect(get.StatusCode).To(Equal(http.StatusOK))

		defer list.Body.Close()
		listResult, err := io.ReadAll(list.Body)
		Expect(err).To(BeNil())

		var ymll []model.MetadataWithID
		Expect(yaml.Unmarshal(listResult, &ymll)).To(Succeed())
		Expect(len(ymll)).To(Equal(1))
		expected := validPayload1
		expected.ID = id
		Expect(ymll).To(ContainElement(expected))
	})

	It("can reject a PUT request if payload misses a field", func() {
		id := "2"
		p, err := yaml.Marshal(invalidPayloadMissingField)
		Expect(err).To(BeNil())

		put, err := http.NewRequest(http.MethodPut,
			fmt.Sprintf("%s/metadata/%s", server.URL, id),
			bytes.NewBuffer(p))
		Expect(err).To(BeNil())

		resp, err := http.DefaultClient.Do(put)
		Expect(err).To(BeNil())

		Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
	})

})
