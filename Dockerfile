FROM golang:1.17.6 as builder
ENV GOPROXY=https://goproxy.cn
WORKDIR /build
COPY ./src .
RUN go mod tidy
RUN go build

FROM  alpine:latest
RUN mkdir -p /cmd
WORKDIR  /cmd
COPY  --from=builder /build/servertest  .
EXPOSE 8000
CMD ["./servertest"]