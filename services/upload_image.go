package services

import (
    "bytes"
    "fmt"
    "mime/multipart"
    "net/http"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "image-gallery-server/models"
)

var cfg *models.Config = LoadConfig()
// UploadFileToS3 uploads a file to an S3 bucket
func UploadFileToS3(file multipart.File, fileHeader *multipart.FileHeader, bucketName string, filePath string) (string, error) {
    // Initialize a session that the SDK uses to load configuration, credentials, and region from the environment
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String(cfg.AwsConfig.Region),
    })
    if err != nil {
        return "", fmt.Errorf("failed to create session: %v", err)
    }

    // Create S3 service client
    svc := s3.New(sess)

    // Read the file content into a buffer
    buffer := new(bytes.Buffer)
    _, err = buffer.ReadFrom(file)
    if err != nil {
        return "", fmt.Errorf("failed to read file: %v", err)
    }

    // Create the S3 object key
    key := fmt.Sprint(filePath, "/", fileHeader.Filename)

    // Upload the file to S3
    _, err = svc.PutObject(&s3.PutObjectInput{
        Bucket:        aws.String(bucketName),
        Key:           aws.String(key),
        Body:          bytes.NewReader(buffer.Bytes()),
        ContentLength: aws.Int64(fileHeader.Size),
        ContentType:   aws.String(http.DetectContentType(buffer.Bytes())),
    })
    if err != nil {
        return "", fmt.Errorf("failed to upload file to S3: %v", err)
    }

   // Construct the full URL
   url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, key)

   return url, nil
}

func DeleteFileFromS3(bucketName, key string) error {
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String(cfg.AwsConfig.Region),
    })
    if err != nil {
        return fmt.Errorf("failed to create session: %v", err)
    }
    svc := s3.New(sess)

    _, err = svc.DeleteObject(&s3.DeleteObjectInput{
        Bucket: aws.String(bucketName),
        Key: aws.String(key),
    })
    if err != nil {
        return fmt.Errorf("failed to delete file from S3: %v", err)
    }
    err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
        Bucket: aws.String(bucketName),
        Key: aws.String(key),
    })
    if err != nil {
        return fmt.Errorf("failed to wait until object not exists: %v", err)
    }
    return nil
}