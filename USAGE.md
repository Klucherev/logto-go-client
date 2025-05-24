<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"github.com/klucherev/logto"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := logto.New(
		logto.WithSecurity(components.Security{
			ClientID:     logto.String(os.Getenv("LOGTO_CLIENT_ID")),
			ClientSecret: logto.String(os.Getenv("LOGTO_CLIENT_SECRET")),
		}),
	)

	res, err := s.Applications.ListApplications(ctx, operations.ListApplicationsRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.ResponseBodies != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->