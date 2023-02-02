# builder image
FROM golang:latest as builder
RUN mkdir /build
ADD ./src/ /build/
RUN ls -la /build
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o alertmanager_healthcheck .


# generate clean, final image for end users
FROM alpine:3.17.1
COPY --from=builder /build/alertmanager_healthcheck .
EXPOSE 2112
# executable
ENTRYPOINT [ "./alertmanager_healthcheck" ]
# arguments that can be overridden
CMD [ "3", "300" ]
