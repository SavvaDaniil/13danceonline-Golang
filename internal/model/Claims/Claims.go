package Claims

type Claims struct {
	jwt.StandartClaims
	Username string `json:"username"`
}
