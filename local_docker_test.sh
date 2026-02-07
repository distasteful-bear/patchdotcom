#!/bin/bash

# Prep Tailwind CSS static files
npx @tailwindcss/cli -i ./src/static/input.css -o ./src/static/output.css

# build go executable for server (statically linked for Alpine)
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o patchdotcom

# build docker image
sudo docker build . --tag patchdotcom

sudo docker run -p 8080:8080 patchdotcom
