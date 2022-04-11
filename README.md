# Instructions for my awesome solution of the test task 
Go technical test: GatewayFM


## Instructions to run and use the program:

### 1. Initial comand

To run the program, use the following command: <br>
`go run sol.go <rate> <interval>` <br>
where  <br>
*rate* is a number of **allowed requests** during an interval <br>
*interval* is the period of time (in **seconds**) for which the requests are counted <br>
*e.g.* <br>
`go run sol.go 3 10` <br>
This command means that you restrict the number of requests to *3* for every *10* seconds. <br>

If the arguments are provided correctly, you will see the message "*Starting server at port 8080.*" in the terminal window. <br>

If you do not provide the arguments/provide wrong arguments, you will get a message with an explanation of an error. <br>

### 2. Accessibility

To see the web page, open your browser and go to `localhost:8080/task`

### 3. Workflow

In case of success, you will see the message "*The request is being processed.*" on the web page and in the terminal window. <br>

Try to refresh the page. <br>
If you try to refresh it too many times (so that the number of requests is more than the rate you provided) for the given period of time, you will see the message about the error: "*Rate limit was exceeded.*" <br>
*e.g.* <br>
If you refresh the page *4* times within *10* seconds, the output will be like that: <br>
*The request is being processed.* <br>
*The request is being processed.* <br>
*The request is being processed.* <br>
*Rate limit was exceeded.* <br>

Try to wait a little bit (so that the new interval begins) and refresh the page one more time. <br>
You will see the message "*The request is being processed.*" again. <br>

The terminal window saves the history of these messages, so you can see the number of successful and failed requests. <br>


## Description of functions

### 1. goodCase(w http.ResponseWriter, r \*http.Request)
Prints the message *The request is being processed.* in case of success.

### 2. badCase(w http.ResponseWriter, r \*http.Request)
Prints the message *Rate limit was exceeded.* in case of success.

### 3. check(w http.ResponseWriter, r \*http.Request)
Checks that the number of requests is not exceeded and calls goodCase or badCase.

### 4. CreateBucket(rate int) TokenBucket
Creates and returns a new *token bucket*.

### 5. Start()
A method of the structure TokenBucket. <br>
Adds units (tokens) to the bucket with a ticker and select-case.

### 6. Stop()
A method of the structure TokenBucket. <br>
Stops the process of adding units and starts deleting them.

### 7. IsEmpty() bool
A method of the structure TokenBucket. <br>
Returns *true* if there are not units in the bucket (if we can't handle the request). <br>
Returns *false* if there is at least one unit in the bucket. In this case, also decreases the number of units. <br>

### 8. add()
A method of the structure TokenBucket. <br>
Adds units to the buffered channel up to rate.

### 9. takeOut()
A method of the structure TokenBucket. <br>
Deletes units from the buffered channel.

### 10. main()
Prints the messages about errors when parsing the command line arguments. <br>
Calls the functions to create a bucket and to fill it. <br>
Creates a new ticker. <br>
Creates a simple server to check the code. <br>

### Structure ToketBucket
Rate is a number of requests per interval. <br>
bufChan is a buffered channel with units.
T is a channel received from the ticker.
aChan is a channel w/o a buf.

