package models

// Token represents the structure of the token item in DynamoDB
type DynamoDBAuthToken struct {
	TokenId string `json:"token_id"`
	Token string `json:"token"`
	UserId int `json:"user_id"`
}