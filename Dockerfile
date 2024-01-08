FROM golang:1.21-alpine AS build

WORKDIR /usr/app

COPY . /usr/app

RUN CGO_ENABLED=0 GOOS=linux go build -o pokego-api main.go


FROM scratch

COPY --from=build /usr/app/pokego-api  .

EXPOSE 9090

CMD ["./pokego-api"]