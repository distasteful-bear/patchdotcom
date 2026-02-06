FROM alpine:3.19

COPY patchdotcom patchdotcom
COPY ./src ./src

ENV GIN_MODE=release
ENV PORT=8080

EXPOSE 8080

CMD ["./patchdotcom"]
