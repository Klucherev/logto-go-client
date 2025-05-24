# CaptchaProvider
(*CaptchaProvider*)

## Overview

Setup the captcha provider.

### Available Operations

* [GetCaptchaProvider](#getcaptchaprovider) - Get captcha provider
* [UpdateCaptchaProvider](#updatecaptchaprovider) - Update captcha provider
* [DeleteCaptchaProvider](#deletecaptchaprovider) - Delete captcha provider

## GetCaptchaProvider

Get the captcha provider, you can only have one captcha provider.

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

    res, err := s.CaptchaProvider.GetCaptchaProvider(ctx)
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

**[*operations.GetCaptchaProviderResponse](../../models/operations/getcaptchaproviderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateCaptchaProvider

Update the captcha provider with the provided settings.

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

    res, err := s.CaptchaProvider.UpdateCaptchaProvider(ctx, operations.UpdateCaptchaProviderRequest{
        Config: operations.CreateConfigRequestUnionUpdateCaptchaProviderConfigRequest1(
            operations.UpdateCaptchaProviderConfigRequest1{
                Type: "<value>",
                SiteKey: "<value>",
                SecretKey: "<value>",
            },
        ),
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateCaptchaProviderRequest](../../models/operations/updatecaptchaproviderrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateCaptchaProviderResponse](../../models/operations/updatecaptchaproviderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteCaptchaProvider

Delete the captcha provider.

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

    res, err := s.CaptchaProvider.DeleteCaptchaProvider(ctx)
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

**[*operations.DeleteCaptchaProviderResponse](../../models/operations/deletecaptchaproviderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |