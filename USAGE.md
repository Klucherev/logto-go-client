<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/components"
	"github.com/klucherev/logto-go-client/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := logtogoclient.New(
		logtogoclient.WithSecurity(components.Security{
			ClientID:     logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
			ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
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