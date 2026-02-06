FROM alpine:3.19

COPY patchdotcom .
COPY .src .

ENV GIN_MODE=release
ENV PORT=8080

EXPOSE 8080

CMD ["./patchdotcom"]
