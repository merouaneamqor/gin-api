FROM golang:latest\nWORKDIR /app\nCOPY . .\nRUN go build -o main ./cmd/server\nCMD ["./main"]
