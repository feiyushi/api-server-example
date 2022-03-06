package store

import (
	"apiserver/pkg/model"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Store", func() {
	var store MetadataStore
	BeforeEach(func() {
		store = NewInMemoryMetadataStore()
	})
	It("can save and get metadata by id", func() {
		// save metadata
		md := &model.Metadata{Title: "app1"}
		Expect(store.SetMetadata("id1", md)).To(Succeed())
		// get metadata by correct id
		emd, err := store.GetMetadata("id1")
		Expect(err).To(BeNil())
		Expect(emd).To(Equal(md))
		// get metadata by incorrect id
		emd2, err := store.GetMetadata("id2")
		Expect(err).To(BeNil())
		Expect(emd2).To(BeNil())
	})
	It("can list metadata matched by company name", func() {
		// save 3 metadata
		md1 := &model.Metadata{Title: "app1", Company: "msft"}
		md2 := &model.Metadata{Title: "app2", Company: "msft"}
		md3 := &model.Metadata{Title: "app3", Company: "google"}
		Expect(store.SetMetadata("id1", md1)).To(Succeed())
		Expect(store.SetMetadata("id2", md2)).To(Succeed())
		Expect(store.SetMetadata("id3", md3)).To(Succeed())
		// list by company msft
		list, err := (store.ListMedatadaByCompany("msft"))
		Expect(err).To((BeNil()))
		Expect(list).To(ContainElements(md1, md2))
		Expect(list).ToNot(ContainElements(md3))
	})
})
