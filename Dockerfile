FROM golang:latest

WORKDIR template-go

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN ls -la
RUN go build -o /usr/local/bin/app ./cmd/

CMD ["app"]
