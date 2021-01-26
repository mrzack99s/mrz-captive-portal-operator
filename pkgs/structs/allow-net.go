package structs

type ZAuthAllowNet struct {
	IPAddress string `json:"IPAddress" binding:"required"`
	DlSpeed   uint32 `json:"DlSpeed" binding:"required"`
	UpSpeed   uint32 `json:"UpSpeed" binding:"required"`
	ShareKey  string `json:"ShareKey" binding:"required"`
}
