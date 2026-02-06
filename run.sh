#!/bin/bash

while true
do
    npx @tailwindcss/cli -i ./src/static/input.css -o ./src/static/output.css
    go run .
done
