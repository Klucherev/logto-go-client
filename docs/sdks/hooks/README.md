# Hooks
(*Hooks*)

## Overview

Hook enables you to effortlessly receive real-time updates regarding specific events, such as user registration, sign-in, or password reset. See [🪝 Webhooks] to get started and learn more.

### Available Operations

* [ListHooks](#listhooks) - Get hooks
* [CreateHook](#createhook) - Create a hook
* [GetHook](#gethook) - Get hook
* [UpdateHook](#updatehook) - Update hook
* [DeleteHook](#deletehook) - Delete hook
* [ListHookRecentLogs](#listhookrecentlogs) - Get recent logs for a hook
* [CreateHookTest](#createhooktest) - Test hook
* [UpdateHookSigningKey](#updatehooksigningkey) - Update signing key for a hook

## ListHooks

Get a list of hooks with optional pagination.

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

    res, err := s.Hooks.ListHooks(ctx, nil, nil, nil)
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
| `includeExecutionStats`                                  | **string*                                                | :heavy_minus_sign:                                       | Whether to include execution stats in the response.      |
| `page`                                                   | **int64*                                                 | :heavy_minus_sign:                                       | Page number (starts from 1).                             |
| `pageSize`                                               | **int64*                                                 | :heavy_minus_sign:                                       | Entries per page.                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListHooksResponse](../../models/operations/listhooksresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateHook

Create a new hook with the given data.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"github.com/klucherev/logto/models/operations"
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

    res, err := s.Hooks.CreateHook(ctx, operations.CreateHookRequest{
        Config: operations.CreateHookConfigRequest{
            URL: "https://exotic-hoof.net/",
        },
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.CreateHookRequest](../../models/operations/createhookrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.CreateHookResponse](../../models/operations/createhookresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetHook

Get hook details by ID.

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

    res, err := s.Hooks.GetHook(ctx, "<id>", nil)
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the hook.                       |
| `includeExecutionStats`                                  | **string*                                                | :heavy_minus_sign:                                       | Whether to include execution stats in the response.      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetHookResponse](../../models/operations/gethookresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateHook

Update hook details by ID with the given data.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"github.com/klucherev/logto/models/operations"
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

    res, err := s.Hooks.UpdateHook(ctx, "<id>", operations.UpdateHookRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `id`                                                                                 | *string*                                                                             | :heavy_check_mark:                                                                   | The unique identifier of the hook.                                                   |
| `requestBody`                                                                        | [operations.UpdateHookRequestBody](../../models/operations/updatehookrequestbody.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.UpdateHookResponse](../../models/operations/updatehookresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteHook

Delete hook by ID.

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

    res, err := s.Hooks.DeleteHook(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the hook.                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteHookResponse](../../models/operations/deletehookresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListHookRecentLogs

Get recent logs that match the given query for the specified hook with pagination.

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

    res, err := s.Hooks.ListHookRecentLogs(ctx, "<id>", nil, nil, nil)
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the hook.                       |
| `logKey`                                                 | **string*                                                | :heavy_minus_sign:                                       | The log key to filter logs.                              |
| `page`                                                   | **int64*                                                 | :heavy_minus_sign:                                       | Page number (starts from 1).                             |
| `pageSize`                                               | **int64*                                                 | :heavy_minus_sign:                                       | Entries per page.                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListHookRecentLogsResponse](../../models/operations/listhookrecentlogsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateHookTest

Test the specified hook with the given events and config.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/klucherev/logto/models/components"
	"github.com/klucherev/logto"
	"github.com/klucherev/logto/models/operations"
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

    res, err := s.Hooks.CreateHookTest(ctx, "<id>", operations.CreateHookTestRequestBody{
        Events: []operations.CreateHookTestEvent{},
        Config: operations.CreateHookTestConfig{
            URL: "https://orderly-institute.biz",
        },
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `id`                                                                                         | *string*                                                                                     | :heavy_check_mark:                                                                           | The unique identifier of the hook.                                                           |
| `requestBody`                                                                                | [operations.CreateHookTestRequestBody](../../models/operations/createhooktestrequestbody.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateHookTestResponse](../../models/operations/createhooktestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateHookSigningKey

Update the signing key for the specified hook.

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

    res, err := s.Hooks.UpdateHookSigningKey(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the hook.                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.UpdateHookSigningKeyResponse](../../models/operations/updatehooksigningkeyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |