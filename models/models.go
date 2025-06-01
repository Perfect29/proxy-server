package models

type ProxyRequest struct {
	Method  string            `json:"method"`
	URL     string            `json:"URL"`
	Headers map[string]string `json:"header"`
}

type ProxyResponse struct {
	ID      string              `json:"id"`
	Status  int                 `json:"status"`
	Headers map[string][]string `json:"headers"`
	Length  int64               `json:"length"`
}

type ProxyLog struct {
	Request  ProxyRequest  `json:"request"`
	Response ProxyResponse `json:"response"`
}
