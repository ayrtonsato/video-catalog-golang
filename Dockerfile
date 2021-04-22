FROM golang:1.16-alpine

WORKDIR /usr/src/
COPY . .

RUN apk add --no-cache build-base shadow openssl bash postgresql-client nodejs npm git make

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest \
	&& go install github.com/golang/mock/mockgen@v1.5.0 && go get -u github.com/kyoh86/richgo

RUN go get -u -v ./... && go install -v ./...

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

EXPOSE 9000
