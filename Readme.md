# About

This project is a ready to deploy service for generating Twitter snoflake IDs and expose that as web-api in a kubernetes cluster. This contains following.

 * application code in golang
 * script to build the docker image
 * kubernetes deploymemt configuration

# quick start

Just do `docker compose up` in the project root directory. To get a snowflake id, open https://localhost:8080/generate in browser.

# Production Deployment

Build and push the image to some registry. Apply the kubernetes config.

## creating docker image
create docker image by following command.

    ./build.sh

## deploying to kubernetes
Assuming you have a kubernetes cluster running and `kubectl` is configured, run following to deploy the service. This assumes you alreadt have a namespace `prod` created for deploying the service.

    kubectl apply -f k8s/prod
