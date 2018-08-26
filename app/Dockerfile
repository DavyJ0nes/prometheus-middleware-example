# Multistage Build

### CREATE DOCKERMASTER USER
FROM alpine:3.8 AS alpine
RUN adduser -D -u 10001 dockmaster

## MAIN IMAGE
FROM scratch
LABEL Name=prometheus-middleware
LABEL Author=davyj0nes

COPY --from=alpine /etc/passwd /etc/passwd

ADD prometheus-middleware /
USER dockmaster

EXPOSE 8080
CMD ["./prometheus-middleware"]
