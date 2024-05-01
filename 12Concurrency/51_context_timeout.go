package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningOperation(ctx context.Context) error {

	// Simulate a long-running operation

	select {

	case <-ctx.Done():

		// Operation canceled due to timeout or manual cancellation

		return ctx.Err()

	case <-time.After(10 * time.Second):

		// Operation completed successfully after 10 seconds

		fmt.Println("Long-running operation completed successfully")

		return nil

	}

}

func main() {

	// Create a context with a timeout of 5 seconds

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel() // Cancel the context to avoid leaking resources

	// Invoke the long-running operation and handle its result

	err := longRunningOperation(ctx)

	if err != nil {

		fmt.Println("Error:", err)

	}
}
