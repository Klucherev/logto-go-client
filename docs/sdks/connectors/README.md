# Connectors
(*Connectors*)

## Overview

Connectors are the bridge between Logto and other third-party vendors who provide short message service (SMS), email service, or user information on wildly accepted social media.

To learn more about connectors, please see [🪛 Configure connectors](https://docs.logto.io/docs/recipes/configure-connectors/).

### Available Operations

* [CreateConnector](#createconnector) - Create connector
* [ListConnectors](#listconnectors) - Get connectors
* [GetConnector](#getconnector) - Get connector
* [UpdateConnector](#updateconnector) - Update connector
* [DeleteConnector](#deleteconnector) - Delete connector
* [CreateConnectorTest](#createconnectortest) - Test passwordless connector
* [CreateConnectorAuthorizationURI](#createconnectorauthorizationuri) - Get connector's authorization URI

## CreateConnector

Create a connector with the given data.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
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

    res, err := s.Connectors.CreateConnector(ctx, operations.CreateConnectorRequest{
        ConnectorID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateConnectorRequest](../../models/operations/createconnectorrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreateConnectorResponse](../../models/operations/createconnectorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListConnectors

Get all connectors in the current tenant.

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

    res, err := s.Connectors.ListConnectors(ctx, nil)
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
| `target`                                                 | **string*                                                | :heavy_minus_sign:                                       | Filter connectors by target.                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListConnectorsResponse](../../models/operations/listconnectorsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetConnector

Get connector data by ID

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

    res, err := s.Connectors.GetConnector(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the connector.                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetConnectorResponse](../../models/operations/getconnectorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateConnector

Update connector by ID with the given data. This methods performs a partial update.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
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

    res, err := s.Connectors.UpdateConnector(ctx, "<id>", operations.UpdateConnectorRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `id`                                                                                           | *string*                                                                                       | :heavy_check_mark:                                                                             | The unique identifier of the connector.                                                        |
| `requestBody`                                                                                  | [operations.UpdateConnectorRequestBody](../../models/operations/updateconnectorrequestbody.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateConnectorResponse](../../models/operations/updateconnectorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteConnector

Delete connector by ID.

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

    res, err := s.Connectors.DeleteConnector(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the connector.                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteConnectorResponse](../../models/operations/deleteconnectorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateConnectorTest

Test a passwordless (email or SMS) connector by sending a test message to the given phone number or email address.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
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

    res, err := s.Connectors.CreateConnectorTest(ctx, "<id>", operations.CreateConnectorTestRequestBody{
        Config: operations.CreateConnectorTestConfig{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `factoryID`                                                                                            | *string*                                                                                               | :heavy_check_mark:                                                                                     | The unique identifier of the factory.                                                                  |
| `requestBody`                                                                                          | [operations.CreateConnectorTestRequestBody](../../models/operations/createconnectortestrequestbody.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.CreateConnectorTestResponse](../../models/operations/createconnectortestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateConnectorAuthorizationURI

Get authorization URI for specified connector by providing redirect URI and randomly generated state.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto-go-client/models/components"
	logtogoclient "github.com/klucherev/logto-go-client"
	"github.com/klucherev/logto-go-client/models/operations"
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

    res, err := s.Connectors.CreateConnectorAuthorizationURI(ctx, "<id>", operations.CreateConnectorAuthorizationURIRequestBody{
        State: "Virginia",
        RedirectURI: "https://unhappy-label.biz/",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `connectorID`                                                                                                                  | *string*                                                                                                                       | :heavy_check_mark:                                                                                                             | The unique identifier of the connector.                                                                                        |
| `requestBody`                                                                                                                  | [operations.CreateConnectorAuthorizationURIRequestBody](../../models/operations/createconnectorauthorizationurirequestbody.md) | :heavy_check_mark:                                                                                                             | N/A                                                                                                                            |
| `opts`                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.CreateConnectorAuthorizationURIResponse](../../models/operations/createconnectorauthorizationuriresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |