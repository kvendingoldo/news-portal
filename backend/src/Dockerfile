FROM golang AS builder
ENV GIT_TERMINAL_PROMPT=1
COPY ./main/go/github.com/kvendingoldo/news-portal /go/src/github.com/kvendingoldo/news-portal
WORKDIR /go/src/github.com/kvendingoldo/news-portal
RUN go get ./
RUN go build -o news-portal

FROM golang
WORKDIR /app
COPY --from=builder /go/src/github.com/kvendingoldo/news-portal/news-portal /app/
COPY ./main/resources/config.toml /app/config.toml
EXPOSE 8000
ENTRYPOINT ./news-portal
