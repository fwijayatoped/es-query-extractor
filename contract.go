package esqueryextractor

// Common attribute contract for any client lib extractor
type CommonAttributeContract struct {
	service     Service
	fullPath    string
	usecase     string
	rateLimiter *RateLimiter
}

type Contract interface {
	WithFullPath(string) Contract
	WithKeyword(string) Contract
	WithUsecase(string) Contract
}
