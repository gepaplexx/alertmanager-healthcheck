# builder image
FROM golang:latest as builder
RUN mkdir /build
ADD ./src/ /build/
RUN ls -la /build
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o alertmanager_health .


# generate clean, final image for end users
FROM alpine:3.11.3
COPY --from=builder /build/alertmanager_health .
EXPOSE 2112
# executable
ENTRYPOINT [ "./alertmanager_health" ]
# arguments that can be overridden
CMD [ "3", "300" ]
