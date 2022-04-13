package userDTO

type UserLoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegistrationDTO struct {
	Firstname string `json:"firstname"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type UserForgetDTO struct {
	Step     int    `json:"step"`
	ForgetId int    `json:"forget_id"`
	Username string `json:"username"`
	Code     string `json:"code"`
}

type UserProfileDTO struct {
	Firstname        string `json:"firstname"`
	Username         string `json:"username"`
	PasswordNew      string `json:"passwordNew"`
	PasswordNewAgain string `json:"passwordNewAgain"`
	PasswordCurrent  string `json:"passwordCurrent"`
}
