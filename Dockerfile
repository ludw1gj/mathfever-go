FROM golang:1.10.0-alpine3.7

# Get git
RUN apk add --no-cache --virtual .build-deps git

# Get dependencies
RUN go get github.com/gorilla/mux
RUN go get github.com/oxtoacart/bpool

# Delete git
RUN apk del .build-deps

# Copy files and install
COPY . /go/src/github.com/ludw1gj/mathfever-go
RUN go install github.com/ludw1gj/mathfever-go

# Set working directory so relative go templates work properly
WORKDIR /go/src/github.com/ludw1gj/mathfever-go

# Run app and expose port
ENTRYPOINT /go/bin/mathfever-go
EXPOSE 8000
