# JSON
(*Swagger.JSON*)

## Overview

Endpoints for the Swagger JSON document.

### Available Operations

* [GetSwaggerJSON](#getswaggerjson) - Get Swagger JSON

## GetSwaggerJSON

The endpoint for the current JSON document. The JSON conforms to the [OpenAPI v3.0.1](https://spec.openapis.org/oas/v3.0.1) (a.k.a. Swagger) specification.

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

    res, err := s.Swagger.JSON.GetSwaggerJSON(ctx)
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

**[*operations.GetSwaggerJSONResponse](../../models/operations/getswaggerjsonresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |