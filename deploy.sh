#!/bin/bash

# Prep Tailwind CSS static files
npx @tailwindcss/cli -i ./src/static/input.css -o ./src/static/output.css

# upload, build and deploy application
gcloud run deploy patchdotcom \
    --source . \
    --project patchdotcom \
    --region us-east4 \
    --allow-unauthenticated \
    --port 8080
