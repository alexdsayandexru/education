FROM golang:1.23-alpine

WORKDIR /src/
COPY . .

RUN go clean --modcache
RUN go mod download
RUN go mod verify

WORKDIR /src

RUN CGO_ENABLED=0 go build -o /bin/app
COPY client /bin

#RUN rm -r /src

ENTRYPOINT ["/bin/app"]
