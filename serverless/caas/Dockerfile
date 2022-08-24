FROM golang:1.19-alpine

ENV PORT=8000
ENV GOOGLE_PROJECT_ID=
ENV GOOGLE_APPLICATION_CREDENTIALS=
ENV CRDB_CONN=

WORKDIR /go/src/github.com/adlerhurst/webnesday/serverless/caas
COPY go.* .

RUN go mod download

COPY main.go .
COPY handler ./handler/
COPY storage ./storage/

RUN go build -o caas main.go \
    && ./caas
# ENTRYPOINT [ "caas" ]