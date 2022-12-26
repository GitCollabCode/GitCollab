# Getting Started with Postgres access library

* Please update this document with useful information when you happen to come across it. 

## pgxpool usage

Sample code demonstrating how pgxpool `Acquire` and `Release` is used inside threads.

```Go
package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// Create a connection pool
	pool, err := pgxpool.Connect(context.Background(), "postgres://username:password@localhost/database")
	if err != nil {
		fmt.Println("Error creating connection pool:", err)
		return
	}
	defer pool.Close()

	// Create a wait group to track the goroutines
	var wg sync.WaitGroup

	// Launch a goroutine for each query we want to execute
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Acquire a connection from the pool
			conn, err := pool.Acquire(context.Background())
			if err != nil {
				fmt.Println("Error acquiring connection:", err)
				return
			}
			defer conn.Release()

			// Execute a query using the connection
			row := conn.QueryRow(context.Background(), "SELECT name FROM users WHERE id=$1", id)

			// Scan the result of the query into a variable
			var name string
			err = row.Scan(&name)
			if err != nil {
				fmt.Println("Error scanning result:", err)
				return
			}

			fmt.Println("Name:", name)
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
```
