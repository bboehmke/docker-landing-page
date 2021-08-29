FROM golang:latest
RUN mkdir /src
WORKDIR /src/
COPY . .
RUN CGO_ENABLED=0 go build -v -o /app .

FROM scratch
WORKDIR /
COPY --from=0 /app /app
CMD ["/app"]