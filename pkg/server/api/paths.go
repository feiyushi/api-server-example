package api

// api server routes
const (
	// "/metadata"
	MetadataRoute = "/metadata"

	// "/metadata/{id}"
	MetadataByIdRoute = MetadataRoute + "/:" + ParameterID
)

const (
	ParameterID = "id"
)

const (
	QueryCompany = "company"
)
