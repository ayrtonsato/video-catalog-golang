FROM golang:1.16-alpine

RUN apk add --no-cache build-base
RUN apk add --no-cache shadow openssl bash postgresql-client nodejs npm git make

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest \
	&& go install github.com/golang/mock/mockgen@v1.5.0 && go get -u github.com/kyoh86/richgo

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

WORKDIR /usr/src/

COPY . .

RUN go get -u -v ./...
RUN go install -v ./...

EXPOSE 9000
