# Status
(*Status*)

## Overview

Endpoints for health check.

### Available Operations

* [GetStatus](#getstatus) - Health check

## GetStatus

The traditional health check API. No authentication needed.

> **Note**
> Even if 204 is returned, it does not guarantee all the APIs are working properly since they may depend on additional resources or external services.

### Example Usage

```go
package main

import(
	"context"
	logtogoclient "github.com/klucherev/logto-go-client"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New()

    res, err := s.Status.GetStatus(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
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

**[*operations.GetStatusResponse](../../models/operations/getstatusresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |