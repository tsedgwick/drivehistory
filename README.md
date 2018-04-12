# drivehistory
[![Go Report Card](https://goreportcard.com/badge/github.com/tsedgwick/drivehistory)](https://goreportcard.com/report/github.com/tsedgwick/drivehistory)
 

Reads the input file and outputs the drivers sorted by most miles driven to least.

# How to run
if you need to build the binary
`go build`
once you have the binary
`./drivehistory -file data.txt `

# Design
My approach was to break down the problem into three steps:
1. Read and process the input file
2. Sort the drivers
3. Write the output


# Read and Process
I took this approach to allow future requests to be able to be quickly implemented.  

For example, the reading and parsing of the file is abstracted away to easily use any io.Reader. 
```
//process reads the reader and populates the drivers map
func process(reader io.Reader, drivers map[string]Driver) error
```

We will be able to easily pass any input in for future requests if it is decided to turn this into a CLI or http server application.

Secondly, we are passing the map into the function.  This allows the map to have already processed multiple inputs.

# Sort
Go makes sorting easy to implement.  Sort just takes in an interface that has three functions:
```
func (d Drivers) Len() int 
func (d Drivers) Less(i, j int) bool 
func (d Drivers) Swap(i, j int) 
```
By defining these three on a named slice of type drivers, we can easily sort the drivers when needed.

# Write
Similar to the read and process, the write is setup to write out to any output.
```
func write(writer io.Writer, drivers []Driver) error 
```
We pass in an interface to the write in order to allow the consumer of our api to define it.  Currently, this is setup to just run as stout, but it could easily be modified to write to a file for example.

# Future enhancements
If a request is made to make this into an http server, we will have to make the drivers map concurrent.  We can do this easily with a sync map. 
