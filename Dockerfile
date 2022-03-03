###################
# BUILD STAGE #####
###################
FROM golang:1.12-alpine AS builder

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN  go build -o main

###################
# RUN STAGE #######
###################
FROM golang:1.12-alpine

WORKDIR /app

COPY --from=builder /app/main main

EXPOSE 8080

CMD ["./main"]