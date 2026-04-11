FROM golang:latest

WORKDIR /template

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN ls -la
RUN go build -v -o template ./cmd/

CMD ["./template"]