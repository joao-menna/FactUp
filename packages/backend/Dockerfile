FROM golang:1-bookworm AS build

ENV CGO_ENABLED=1
ENV GOARCH=amd64
ENV GOOS=linux

RUN apt-get update && apt-get install -y libwebp-dev

WORKDIR /source

COPY . .

RUN go build -ldflags="-s -w" -o backend .

FROM debian:bookworm-slim AS run

RUN apt-get update && \
    apt-get install -y --no-install-recommends libwebp-dev && \
    apt-get clean

WORKDIR /app

COPY --from=build /source/backend .

RUN chmod +x ./backend

EXPOSE 8080

CMD [ "./backend" ]
