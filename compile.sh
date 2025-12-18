#!/bin/bash

# Compile Tailwind CSS
npx tailwindcss -i ./dist/main.css -o ./dist/tailwind.css --minify

# Compile Golang
go build -o . ./main.go
