# Image Gallery Server API
Description coming soon


## Getting Started

To set up a local development environment run thi

## Staging

Before you set up a staging environment you will need to have docker installed on your computer.

To set up staging environment to test out docker container locally, first make sure you have a `.env.production` file, in this file make sure it is set up like this:
```.env
REGION=aws-region
SECRET_ACCESS_KEY=aws-iam-user-secret-key
ACCESS_KEY_ID=aws-iam-user-access-key
DB_NAME=postgres-database-name
DB_USER=postgres-database-username
DB_PASSWORD=postgres-database-password
DB_PORT=postgres-database-port
DB_HOST=postgres-database-host
SSL_MODE=postgres-database-ssl-mode
```

Once you have your environment variables set up, run the `start_staging.sh` file to set up a docker container for staging purposes. To execute the code you will need to have Git Bash installed on your computer and then run the command:
```sh
bash start_staging.sh
```
## Deployment

in progress