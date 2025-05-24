# Verifications
(*Verifications*)

## Overview

Endpoints for creating and validating verification records, which can be used in Profile routes.

### Available Operations

* [CreateVerificationByPassword](#createverificationbypassword) - Create a record by password
* [CreateVerificationByVerificationCode](#createverificationbyverificationcode) - Create a record by verification code
* [VerifyVerificationByVerificationCode](#verifyverificationbyverificationcode) - Verify verification code
* [CreateVerificationBySocial](#createverificationbysocial) - Create a social verification record
* [VerifyVerificationBySocial](#verifyverificationbysocial) - Verify a social verification record

## CreateVerificationByPassword

Create a verification record by verifying the password.

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

    res, err := s.Verifications.CreateVerificationByPassword(ctx, operations.CreateVerificationByPasswordRequest{
        Password: "VHdaYOG8g_Gbmo2",
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.CreateVerificationByPasswordRequest](../../models/operations/createverificationbypasswordrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.CreateVerificationByPasswordResponse](../../models/operations/createverificationbypasswordresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateVerificationByVerificationCode

Create a verification record and send the code to the specified identifier. The code verification can be used to verify the given identifier.

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

    res, err := s.Verifications.CreateVerificationByVerificationCode(ctx, operations.CreateVerificationByVerificationCodeRequest{
        Identifier: operations.CreateVerificationByVerificationCodeIdentifier{
            Type: operations.CreateVerificationByVerificationCodeTypePhone,
            Value: "<value>",
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

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.CreateVerificationByVerificationCodeRequest](../../models/operations/createverificationbyverificationcoderequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.CreateVerificationByVerificationCodeResponse](../../models/operations/createverificationbyverificationcoderesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## VerifyVerificationByVerificationCode

Verify the provided verification code against the identifier. If successful, the verification record will be marked as verified.

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

    res, err := s.Verifications.VerifyVerificationByVerificationCode(ctx, operations.VerifyVerificationByVerificationCodeRequest{
        Identifier: operations.VerifyVerificationByVerificationCodeIdentifier{
            Type: operations.VerifyVerificationByVerificationCodeTypePhone,
            Value: "<value>",
        },
        VerificationID: "<id>",
        Code: "<value>",
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

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.VerifyVerificationByVerificationCodeRequest](../../models/operations/verifyverificationbyverificationcoderequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.VerifyVerificationByVerificationCodeResponse](../../models/operations/verifyverificationbyverificationcoderesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateVerificationBySocial

Create a social verification record and return the authorization URI.

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

    res, err := s.Verifications.CreateVerificationBySocial(ctx, operations.CreateVerificationBySocialRequest{
        State: "Vermont",
        RedirectURI: "https://husky-clavicle.info/",
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.CreateVerificationBySocialRequest](../../models/operations/createverificationbysocialrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.CreateVerificationBySocialResponse](../../models/operations/createverificationbysocialresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## VerifyVerificationBySocial

Verify a social verification record by callback connector data, and save the user information to the record.

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

    res, err := s.Verifications.VerifyVerificationBySocial(ctx, operations.VerifyVerificationBySocialRequest{
        ConnectorData: operations.VerifyVerificationBySocialConnectorData{},
        VerificationRecordID: "<id>",
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.VerifyVerificationBySocialRequest](../../models/operations/verifyverificationbysocialrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.VerifyVerificationBySocialResponse](../../models/operations/verifyverificationbysocialresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |