package banner

import "github.com/eacha/mercury/lib/scan"

type BannerData struct {
	scan.Data
	Banner string `json:"banner,omitempty"`
}
