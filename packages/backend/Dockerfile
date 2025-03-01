FROM golang:1-alpine AS build

WORKDIR /source

COPY . .

RUN go build .

FROM alpine AS run

WORKDIR /app

COPY --from=build /source/backend .

RUN chmod +x ./backend

EXPOSE 8080

CMD [ "./backend" ]
