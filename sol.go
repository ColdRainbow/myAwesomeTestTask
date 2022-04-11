package main

import (
    "fmt"
	"os"
	"strconv"
    "log"
    "net/http"
	"time"
)

var Tb TokenBucket

func goodCase(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("The request is being processed.\n")
	fmt.Fprintf(w, "The request is being processed.")
}

func badCase(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Rate limit was exceeded.\n")
	fmt.Fprintf(w, "Rate limit was exceeded.")
}

func check(w http.ResponseWriter, r *http.Request) {
	if !Tb.IsEmpty() {
		goodCase(w, r)
		return
	}
	badCase(w, r)
}

func main() {
    http.HandleFunc("/task", check)
	if len(os.Args) != 3 {
		fmt.Printf("Something is wrong with arguments\n")
		return
	}
	rate, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Something is wrong with arguments\n")
		return
	}
	Tb = CreateBucket(rate)
	Tb.Start()
	interval, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Something is wrong with arguments\n")
		return
	}
	Tb.T = time.NewTicker(time.Duration(interval) * time.Second)

    fmt.Printf("Starting server at port 8080.\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

type Unit struct {}

type TokenBucket struct {
	Rate	int
	bufChan	chan Unit
	T		*time.Ticker
	aChan	chan Unit
}

func CreateBucket(rate int) TokenBucket {
	bufChan := make(chan Unit, rate)

	return TokenBucket{
		Rate:		rate,
		bufChan:	bufChan,
	}
}

func (tb *TokenBucket) Start() {
	tb.T = time.NewTicker(time.Second)
	tb.aChan = make(chan Unit)

	tb.add()

	go func() {
		defer close(tb.aChan)
		for {
			select {
			case <-tb.aChan:
				tb.T.Stop()
				return
			case <-tb.T.C:
				tb.add()
			}
		}
	}()
}

func (tb *TokenBucket) Stop() {
	tb.aChan <- Unit{}
	<-tb.aChan

	tb.takeOut()
}

func (tb *TokenBucket) IsEmpty() bool {
	select {
	case <-tb.bufChan:
		return false
	default:
		return true
	}
}

func (tb *TokenBucket) add() {
	for i := 0; i < tb.Rate; i++ {
		select {
		case tb.bufChan <- Unit{}:
		default:
		}
	}
}

func (tb *TokenBucket) takeOut() {
	for i := 0; i < tb.Rate; i++ {
		select {
		case <-tb.bufChan:
		default:
		}
	}
}

