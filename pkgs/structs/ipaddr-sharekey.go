package structs

type ZAuthIPAddressAndShareKey struct {
	IPAddress string `json:"IPAddress" binding:"required"`
	ShareKey  string `json:"ShareKey" binding:"required"`
}
