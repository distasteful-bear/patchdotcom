#!/bin/bash

# Prep Tailwind CSS static files
npx @tailwindcss/cli -i ./src/static/input.css -o ./src/static/output.css

# build go executable for server
go build -o patchdotcom

# upload, build and deploy application
gcloud run deploy patchdotcom \
    --source . \
    --region us-east5 \
    --allow-unauthenticated \
    --port 8080
