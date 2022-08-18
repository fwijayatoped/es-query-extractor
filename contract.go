package esqueryextractor

type Contract interface {
	WithPath(path string) Contract
	WithQueryString(querystring string) Contract
	WithExtraAttributes(attributes map[string]string) Contract
}

type CommonAttributeContract struct {
	path            string
	querystring     string
	extraAttributes map[string]string
}
