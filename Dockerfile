FROM golang:1.9.3-alpine3.7

RUN apk --no-cache --update add git

# Get dependencies
RUN go get github.com/gorilla/mux
RUN go get github.com/oxtoacart/bpool

# Copy files and install
COPY . /go/src/github.com/robertjeffs/mathfever-go
RUN go install github.com/robertjeffs/mathfever-go

# Set working directory so relative go templates work properly
WORKDIR /go/src/github.com/robertjeffs/mathfever-go

# Run app and expose port
ENTRYPOINT /go/bin/mathfever-go
EXPOSE 8000