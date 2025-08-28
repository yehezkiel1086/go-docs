package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func contextTimeout(ctx context.Context) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	done := make(chan struct{})

	// think of the following func as an api
	go func() {
		time.Sleep(3 * time.Second)
		close(done)
	}()

	select {
	case <- done:
		fmt.Println("Called the API")
	case <- ctxWithTimeout.Done():
		fmt.Println("Oh no my timeout expired", ctxWithTimeout.Err())
		// do some logic to handle this
	}
}

func exampleWithValues() {
	type key int
	const UserKey key = 0

	ctx := context.Background()

	ctxWithValue := context.WithValue(ctx, UserKey, "123")

	if userID, ok := ctxWithValue.Value(UserKey).(string); ok {
		fmt.Println("This is the userID", userID)
	} else {
		fmt.Println("This is a protected route - no userID found.")
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 4 * time.Second)
	defer cancel()

	select {
	case <- time.After(3 * time.Second):
		fmt.Println("API Response!")
	case <- ctx.Done():
		fmt.Println("Oh no, the context expired")
		http.Error(w, "Request context timeout", http.StatusRequestTimeout)
		return
	}
}

func main() {
	// ctx := context.Background()
	// contextTimeout(ctx)
	// exampleWithValues()

	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}
