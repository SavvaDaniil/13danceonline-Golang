package entity

type User struct {
	Id                int
	Username          string
	Password          *string
	AuthKey           *string
	AccessToken       *string
	Active            int
	Firstname         *string
	DateOfAdd         *string
	ForgetCode        *string
	ForgetTry         int
	ForgetHash        *string
	ForgetCount       int
	DateForgetLastTry *string
}
