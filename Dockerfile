FROM golang:1.17.7
ENV CGO_ENABLED 0
ADD . /src
WORKDIR /src
RUN go build -a --installsuffix cgo --ldflags="-s" -o synonym

FROM debian:11-slim
ENV GIN_MODE=release
RUN apt-get update && \
	apt-get upgrade -y && \
	apt-get install -y ca-certificates && \
	apt-get clean
COPY --from=0 /src/synonym /usr/bin/synonym
ENTRYPOINT ["/usr/bin/synonym"]
