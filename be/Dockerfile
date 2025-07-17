FROM golang:1.24.5-bookworm AS development

WORKDIR /usr/src/ctx_fndr_api
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -o server
CMD [ "./server" ]
