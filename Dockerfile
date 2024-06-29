# syntax=docker/dockerfile:1

FROM alpine:latest

# Set destination for COPY
WORKDIR /gokvstore


# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY gokvstore ./

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080/tcp

# Run
CMD ["./gokvstore"]