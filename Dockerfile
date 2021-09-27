FROM golang:alpine 
WORKDIR /simple-cli
COPY . .
RUN apk add build-base
RUN go install .
