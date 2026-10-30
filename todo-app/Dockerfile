FROM alpine:3.21

COPY --from=hashicorp/http-echo:latest /http-echo /http-echo

ENV PORT=8080

EXPOSE 8080

ENTRYPOINT ["/bin/sh", "-c", \
  "exec /http-echo -listen=:${PORT} -text=\"Server started in port ${PORT}\""]