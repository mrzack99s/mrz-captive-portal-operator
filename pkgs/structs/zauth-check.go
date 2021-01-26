package structs

type ZAuthCheck struct {
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
}

func (b *ZAuthCheck) TableName() string {
	return "zauth_check"
}
