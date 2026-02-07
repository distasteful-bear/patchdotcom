FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
RUN CGO_ENABLED=0 go build -o patchdotcom

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/patchdotcom .
COPY ./src ./src

ENV GIN_MODE=release
ENV PORT=8080

EXPOSE 8080

# SENDGRID_API_KEY is expected at runtime via: docker run --env-file sendgrid.docker.env
CMD ["./patchdotcom"]
