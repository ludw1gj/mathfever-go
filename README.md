# MathFever
A website built with Go, where users can find mathematical proof and answers to common math problems.

## Project Dependencies 
Golang >=1.8  
gorilla/mux  
oxtoacart/bpool

## Setup Notes:

``` bash
# Get the code and dependencies
go get github.com/gorilla/mux
go get github.com/oxtoacart/bpool
go get github.com/robertjeffs/mathfever-go

# Navigate to the github.com/robertjeffs/mathfever-go directory
cd {GO-PATH}/github.com/robertjeffs/mathfever-go

# To run the project without building (access via localhost:8000)
go run main.go

# To build the project (access via localhost:8000)
go build
./go-mathfever
```
