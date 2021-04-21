FROM golang:alpine
RUN apk add git
RUN apk add chromium
RUN apk add tor 
#RUN nohup tor &
COPY code.go /go/code.go
RUN go build code.go
