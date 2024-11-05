package models

type Config struct {
	AwsConfig AwsConfig
	PostgresConfig PostgresConfig
}

type AwsConfig struct {
	Region string
	AWS_SECRET_ACCESS_KEY string
	AWS_ACCESS_KEY_ID string
}
type PostgresConfig struct {
	DB_NAME string
	DB_USER string
	DB_PASSWORD string
	DB_PORT string
	DB_HOST string
	SSL_MODE string
}