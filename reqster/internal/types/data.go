package types

type DataP struct {
	URL        string
	Headers    map[string]string
	Params     map[string]string
	Query      map[string]string
	Body       string
	BodyPath   string
	OutputPath string
}
