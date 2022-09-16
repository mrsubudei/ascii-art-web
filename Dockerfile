#build
FROM golang:alpine3.16 AS build

LABEL stage=build

WORKDIR /app

COPY . ./

RUN go build cmd/main.go

#copy all needed files into second container
FROM alpine:3.16 AS runner

WORKDIR /app

LABEL authors="@Subudei"

COPY --from=build /app/main /app/main

COPY /templates /app/templates

CMD ["/app/main"]

RUN echo localhost:8087
