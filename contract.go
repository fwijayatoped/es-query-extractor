package esqueryextractor

// Common attribute contract for any client lib extractor
type CommonAttributeContract struct {
	path            string
	querystring     string
	extraAttributes map[string]string
}
