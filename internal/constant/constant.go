package constant

const (
	ConfigPort = ":3000"

	USERID        = "user_id"
	Unauthorized  = "Unauthorized"
	Authorization = "Authorization"

	InterestRate = 0.10
	TotalWeeks   = 7
)

var JwtKey = []byte("my_secret_key")
