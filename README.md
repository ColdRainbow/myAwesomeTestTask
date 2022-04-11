# Instructions for my awesome solution of the test task 
Go technical test: GatewayFM

## Instructions to run and use the program:

### 1. Initial comand

To run the program, use the following command:
`go run sol.go *<rate>* *<interval>*`
where
*<rate>* is a number of **allowed requests** during an interval
*<interval>* is the period of time (in **seconds**) for which the requests are counted
*e.g.*
`go run sol.go 3 10`
this command means that you restrict the number of requests to *3* for every *10* seconds.

If the arguments are provided correctly, you will see the message "`Starting server at port 8080.`" in the terminal window.

If you do not provide the arguments/provide wrong arguments, you will get a messages with an explanation of an error.

### 2. Accessibility

To see the web page, open your browser and go to `localhost:8080/task`

### 3. Workflow

In case of success, you will see the message "`The request is being processed.`" on the web page and in the terminal window.

Try to refresh the page.
If you try to refresh it too many times (so that the number of requests is more than the rate you provided) for the given period of time, you will see the message about the error: "`Rate limit was exceeded.`"
*e.g.*
if you refresh the page *4* times within *10* seconds, the output will be like that:
`The request is being processed.`
`The request is being processed.`
`The request is being processed.`
`Rate limit was exceeded.`

Try to wait a little bit (so that the interval changed) and refresh the page one more time. You will see the message "`The request is being processed.`" again.

The terminal window saves the history of these messages, so you can see the number of successful and failed requests.
