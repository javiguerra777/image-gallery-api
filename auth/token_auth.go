package auth

import (
    "fmt"
    "net/http"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
    "github.com/gin-gonic/gin"
    "image-gallery-server/models"
)

// GetTokenFromDB retrieves the token from DynamoDB using a secondary index
func GetTokenFromDB(dynamodbClient *dynamodb.DynamoDB, authToken string) (models.DynamoDBAuthToken, error) {
    // Define the query input
    input := &dynamodb.QueryInput{
        TableName: aws.String("tokens"),
        IndexName: aws.String("token-index"),
        KeyConditionExpression: aws.String("#token = :token"),
        ExpressionAttributeNames: map[string]*string{
            "#token": aws.String("token"),
        },
        ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
            ":token": {
                S: aws.String(authToken),
            },
        },
    }

    // Perform the Query operation
    result, err := dynamodbClient.Query(input)
    if err != nil {
        if aerr, ok := err.(awserr.Error); ok {
            switch aerr.Code() {
            case dynamodb.ErrCodeResourceNotFoundException:
                return models.DynamoDBAuthToken{}, fmt.Errorf("table or index not found")
            case dynamodb.ErrCodeProvisionedThroughputExceededException:
                return models.DynamoDBAuthToken{}, fmt.Errorf("throughput exceeded")
            default:
                return models.DynamoDBAuthToken{}, aerr
            }
        }
        return models.DynamoDBAuthToken{}, err
    }

    // Check if the item was found
    if len(result.Items) == 0 {
        return models.DynamoDBAuthToken{}, fmt.Errorf("token not found")
    }

    // Unmarshal the result into a Token struct
    var token models.DynamoDBAuthToken
    err = dynamodbattribute.UnmarshalMap(result.Items[0], &token)
    if err != nil {
        return models.DynamoDBAuthToken{} , fmt.Errorf("failed to unmarshal result: %v", err)
    }

    return token, nil
}

func AuthMiddleware(dynamodbClient *dynamodb.DynamoDB) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
            c.Abort()
            return
        }
        tokenUser, err := GetTokenFromDB(dynamodbClient, authHeader); 
        if err != nil {
            fmt.Println(err)
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }
        c.Set("tokenUser", tokenUser)
        c.Next()
    }
}