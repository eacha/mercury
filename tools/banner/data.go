package banner

type BannerData struct {
	IP        string `json:"ip,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Error     string `json:"error,omitempty"`
	Banner    string `json:"banner,omitempty"`
}
