# Systems
(*Systems*)

## Overview

Endpoints for system constants and information.

### Available Operations

* [GetSystemApplicationConfig](#getsystemapplicationconfig) - Get the application constants.

## GetSystemApplicationConfig

Get the application constants.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New(
        logto.WithSecurity(components.Security{
            ClientID: logto.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logto.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.Systems.GetSystemApplicationConfig(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetSystemApplicationConfigResponse](../../models/operations/getsystemapplicationconfigresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |