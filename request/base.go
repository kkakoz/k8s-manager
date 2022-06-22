package request

type ListReq struct {
	Namespace string `json:"namespace"`
	Label     string `query:"label"`
	Field     string `query:"field"`
	Limit     int64  `query:"limit"`
	Continue  string `query:"continue"`
}

type DeleteReq struct {
	Name      string `url:"name"`
	Namespace string `json:"namespace"`
}

type ApplyReq struct {
	Content   string `json:"content"`
	Namespace string `json:"namespace"`
}