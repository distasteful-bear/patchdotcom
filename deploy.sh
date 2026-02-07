#!/bin/bash

# Prep Tailwind CSS static files
npx @tailwindcss/cli -i ./src/static/input.css -o ./src/static/output.css

# upload, build and deploy application
gcloud run deploy patchdotcom \
    --source . \
    --region us-east5 \
    --allow-unauthenticated \
    --port 8080
