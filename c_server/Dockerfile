FROM alpine:3.21
RUN apk add --no-cache --virtual .build-deps bash gcc musl-dev openssl go make
RUN wget https://go.dev/dl/go1.24.2.linux-amd64.tar.gz \
  && tar -C /usr/local -xzf go1.24.2.linux-amd64.tar.gz \
  && export PATH=$PATH:/usr/local/go/bin \
  && rm go1.24.2.linux-amd64.tar.gz

RUN mkdir /app
WORKDIR /app

COPY . .

ENTRYPOINT ["go", "run", "src/main.go"]