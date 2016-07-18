package scan

type Data struct {
	IP        string `json:"ip,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Error     string `json:"error,omitempty"`
}
