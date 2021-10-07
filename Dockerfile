FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o webconsole . 
FROM alpine
RUN mkdir -p /app/confs
RUN mkdir -p /app/docs
COPY --from=builder /build/webconsole /app/ 
COPY --from=builder /build/confs/ /app/confs
COPY --from=builder /build/docs/ /app/docs
WORKDIR /app 
CMD ["./webconsole"]
