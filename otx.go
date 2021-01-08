package ioc

// Provider cho otx.alienvault.com
type OtxProvider struct {
	APIKey string
	URL string
}

type OtxlResult struct {
	Data []OtxInfo `json:"data"`
}

type OtxInfo struct {
	Indicator string `json:"indicator"`
	TypeIndicator string `json:"type_indicator"`
}
