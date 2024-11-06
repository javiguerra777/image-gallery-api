#!/bin/bash
# Stop and remove the existing container if it exists
if [ "$(docker ps -aq -f name=myapp_staging_container)" ]; then
    docker stop myapp_staging_container
    docker rm myapp_staging_container
fi
docker build -t myapp .
docker run --env-file .env.production -p 4000:4000 --name myapp_staging_container myapp

# Exit the script
exit 0