# Go Correlation Id Provider

gorrelation is a Go `net/http` middleware that injects correlation ids into the incoming headers. See 

## Installation

## Usage

### net/http

### Negroni

### Gin

You can create a middleware handler in gin by calling with a gin context.

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thefury/gorrelation"
)

func main() {
	middleware := gorrelation.New()
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		middleware.EnsureContextId(c.Request)
		c.Next()
	})

	// Jobs
	router.GET("/", indexHandler)
	router.Run()
}
```

### HTTPRouter

## Contributing
