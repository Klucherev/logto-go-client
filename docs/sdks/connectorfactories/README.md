# ConnectorFactories
(*ConnectorFactories*)

## Overview

Connector factories are used to create connectors. They can be treated as preconfigured templates for connectors.

### Available Operations

* [ListConnectorFactories](#listconnectorfactories) - Get connector factories
* [GetConnectorFactory](#getconnectorfactory) - Get connector factory

## ListConnectorFactories

Get all connector factories data available in Logto.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.ConnectorFactories.ListConnectorFactories(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseBodies != nil {
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

**[*operations.ListConnectorFactoriesResponse](../../models/operations/listconnectorfactoriesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetConnectorFactory

Get connector factory by the given ID.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtogoclient.New(
        logtogoclient.WithSecurity(components.Security{
            ClientID: logtogoclient.String(os.Getenv("LOGTO_CLIENT_ID")),
            ClientSecret: logtogoclient.String(os.Getenv("LOGTO_CLIENT_SECRET")),
        }),
    )

    res, err := s.ConnectorFactories.GetConnectorFactory(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the connector factory.          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetConnectorFactoryResponse](../../models/operations/getconnectorfactoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |