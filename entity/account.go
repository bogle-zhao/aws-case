package entity

type Account struct {
	UserName string `dynamodbav:"username"`
	Password string `dynamodbav:"password"`
	Avatar   string `dynamodbav:"avatar"`
}
