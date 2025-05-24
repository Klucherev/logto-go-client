# WellKnown
(*WellKnown*)

## Overview

Well-Known routes provide information and resources that can be discovered by clients without the need for authentication.

### Available Operations

* [GetWellKnownEndpoint](#getwellknownendpoint)
* [~~GetSignInExperienceConfig~~](#getsigninexperienceconfig) - Get full sign-in experience :warning: **Deprecated**
* [GetSignInExperiencePhrases](#getsigninexperiencephrases) - Get localized phrases
* [GetWellKnownExperience](#getwellknownexperience) - Get full sign-in experience
* [GetWellKnownManagementOpenapiJSON](#getwellknownmanagementopenapijson) - Get Management API swagger JSON
* [GetWellKnownExperienceOpenapiJSON](#getwellknownexperienceopenapijson) - Get Experience API swagger JSON
* [GetWellKnownUserOpenapiJSON](#getwellknownuseropenapijson) - Get User API swagger JSON

## GetWellKnownEndpoint

### Example Usage

```go
package main

import(
	"context"
	"github.com/klucherev/logto"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New()

    res, err := s.WellKnown.GetWellKnownEndpoint(ctx, "<id>")
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
| `tenantID`                                               | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the tenant.                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetWellKnownEndpointResponse](../../models/operations/getwellknownendpointresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ~~GetSignInExperienceConfig~~

Get the full sign-in experience configuration.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"context"
	"github.com/klucherev/logto"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New()

    res, err := s.WellKnown.GetSignInExperienceConfig(ctx, nil, nil)
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
| `organizationID`                                         | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `appID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetSignInExperienceConfigResponse](../../models/operations/getsigninexperienceconfigresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetSignInExperiencePhrases

Get localized phrases based on the specified language.

### Example Usage

```go
package main

import(
	"context"
	"github.com/klucherev/logto"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New()

    res, err := s.WellKnown.GetSignInExperiencePhrases(ctx, nil)
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
| `lng`                                                    | **string*                                                | :heavy_minus_sign:                                       | The language tag for localization.                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetSignInExperiencePhrasesResponse](../../models/operations/getsigninexperiencephrasesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetWellKnownExperience

Get the full sign-in experience configuration.

### Example Usage

```go
package main

import(
	"context"
	"github.com/klucherev/logto"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New()

    res, err := s.WellKnown.GetWellKnownExperience(ctx, nil, nil)
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
| `organizationID`                                         | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `appID`                                                  | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetWellKnownExperienceResponse](../../models/operations/getwellknownexperienceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetWellKnownManagementOpenapiJSON

The endpoint for the Management API JSON document. The JSON conforms to the [OpenAPI v3.0.1](https://spec.openapis.org/oas/v3.0.1) (a.k.a. Swagger) specification.

### Example Usage

```go
package main

import(
	"context"
	"github.com/klucherev/logto"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New()

    res, err := s.WellKnown.GetWellKnownManagementOpenapiJSON(ctx)
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

**[*operations.GetWellKnownManagementOpenapiJSONResponse](../../models/operations/getwellknownmanagementopenapijsonresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetWellKnownExperienceOpenapiJSON

The endpoint for the Experience API JSON document. The JSON conforms to the [OpenAPI v3.0.1](https://spec.openapis.org/oas/v3.0.1) (a.k.a. Swagger) specification.

### Example Usage

```go
package main

import(
	"context"
	"github.com/klucherev/logto"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New()

    res, err := s.WellKnown.GetWellKnownExperienceOpenapiJSON(ctx)
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

**[*operations.GetWellKnownExperienceOpenapiJSONResponse](../../models/operations/getwellknownexperienceopenapijsonresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetWellKnownUserOpenapiJSON

The endpoint for the User API JSON document. The JSON conforms to the [OpenAPI v3.0.1](https://spec.openapis.org/oas/v3.0.1) (a.k.a. Swagger) specification.

### Example Usage

```go
package main

import(
	"context"
	"github.com/klucherev/logto"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New()

    res, err := s.WellKnown.GetWellKnownUserOpenapiJSON(ctx)
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

**[*operations.GetWellKnownUserOpenapiJSONResponse](../../models/operations/getwellknownuseropenapijsonresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |