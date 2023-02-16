# builder image
FROM docker.io/library/golang:1.20.1 as builder
RUN mkdir /build
COPY ./src/ /build/
RUN ls -la /build
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o alertmanager_healthcheck .


# generate clean, final image for end users
FROM docker.io/library/alpine:3.17.2
COPY --from=builder /build/alertmanager_healthcheck .
EXPOSE 2112
# executable
ENTRYPOINT [ "./alertmanager_healthcheck" ]
# arguments that can be overridden
CMD [ "3", "300" ]
