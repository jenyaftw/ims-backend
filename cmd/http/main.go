package main

import (
	"fmt"

	"github.com/jenyaftw/scaffold-go/internal/adapters/delivery/http"
)

func main() {
	r := http.NewRouter().WithHost("localhost").WithPort(3333).Build()

	fmt.Printf("Listening on http://%s:%d\n", r.Host, r.Port)
	if err := r.ListenAndServe(); err != nil {
		panic(err)
	}
}
