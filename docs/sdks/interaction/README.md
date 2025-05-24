# Interaction
(*Interaction*)

## Overview

Interaction endpoints are used to manage and process interactions for end-users, such as sign-in experience. Interaction endpoints are legacy endpoints that are used internally, will be replaced with Experience endpoints instead.

### Available Operations

* [PutAPIInteraction](#putapiinteraction)
* [DeleteAPIInteraction](#deleteapiinteraction)
* [PutAPIInteractionEvent](#putapiinteractionevent)
* [PatchAPIInteractionIdentifiers](#patchapiinteractionidentifiers)
* [PutAPIInteractionProfile](#putapiinteractionprofile)
* [PatchAPIInteractionProfile](#patchapiinteractionprofile)
* [DeleteAPIInteractionProfile](#deleteapiinteractionprofile)
* [PostAPIInteractionSubmit](#postapiinteractionsubmit)
* [PostAPIInteractionConsent](#postapiinteractionconsent)
* [GetAPIInteractionConsent](#getapiinteractionconsent)
* [PostAPIInteractionVerificationSocialAuthorizationURI](#postapiinteractionverificationsocialauthorizationuri)
* [PostAPIInteractionVerificationVerificationCode](#postapiinteractionverificationverificationcode)
* [PostAPIInteractionVerificationTotp](#postapiinteractionverificationtotp)
* [PostAPIInteractionVerificationWebauthnRegistration](#postapiinteractionverificationwebauthnregistration)
* [PostAPIInteractionVerificationWebauthnAuthentication](#postapiinteractionverificationwebauthnauthentication)
* [PostAPIInteractionBindMfa](#postapiinteractionbindmfa)
* [PutAPIInteractionMfa](#putapiinteractionmfa)
* [PutAPIInteractionMfaSkipped](#putapiinteractionmfaskipped)
* [PostAPIInteractionSingleSignOnConnectorIDAuthorizationURL](#postapiinteractionsinglesignonconnectoridauthorizationurl)
* [PostAPIInteractionSingleSignOnConnectorIDAuthentication](#postapiinteractionsinglesignonconnectoridauthentication)
* [PostAPIInteractionSingleSignOnConnectorIDRegistration](#postapiinteractionsinglesignonconnectoridregistration)
* [GetAPIInteractionSingleSignOnConnectors](#getapiinteractionsinglesignonconnectors)

## PutAPIInteraction

### Example Usage

```go
package main

import(
	"context"
	"github.com/klucherev/logto"
	"github.com/klucherev/logto/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New()

    res, err := s.Interaction.PutAPIInteraction(ctx, operations.PutAPIInteractionRequest{
        Event: operations.PutAPIInteractionEventRegister,
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.PutAPIInteractionRequest](../../models/operations/putapiinteractionrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.PutAPIInteractionResponse](../../models/operations/putapiinteractionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteAPIInteraction

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

    res, err := s.Interaction.DeleteAPIInteraction(ctx)
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

**[*operations.DeleteAPIInteractionResponse](../../models/operations/deleteapiinteractionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## PutAPIInteractionEvent

### Example Usage

```go
package main

import(
	"context"
	"github.com/klucherev/logto"
	"github.com/klucherev/logto/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New()

    res, err := s.Interaction.PutAPIInteractionEvent(ctx, operations.PutAPIInteractionEventRequest{
        Event: operations.PutAPIInteractionEventEventRegister,
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PutAPIInteractionEventRequest](../../models/operations/putapiinteractioneventrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.PutAPIInteractionEventResponse](../../models/operations/putapiinteractioneventresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## PatchAPIInteractionIdentifiers

### Example Usage

```go
package main

import(
	"context"
	"github.com/klucherev/logto"
	"github.com/klucherev/logto/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New()

    res, err := s.Interaction.PatchAPIInteractionIdentifiers(ctx, operations.CreatePatchAPIInteractionIdentifiersRequestPatchAPIInteractionIdentifiersRequestBody7(
        operations.PatchAPIInteractionIdentifiersRequestBody7{
            ConnectorID: "<id>",
            Email: "Gaetano7@hotmail.com",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.PatchAPIInteractionIdentifiersRequest](../../models/operations/patchapiinteractionidentifiersrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.PatchAPIInteractionIdentifiersResponse](../../models/operations/patchapiinteractionidentifiersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## PutAPIInteractionProfile

### Example Usage

```go
package main

import(
	"context"
	"github.com/klucherev/logto"
	"github.com/klucherev/logto/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New()

    res, err := s.Interaction.PutAPIInteractionProfile(ctx, operations.PutAPIInteractionProfileRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PutAPIInteractionProfileRequest](../../models/operations/putapiinteractionprofilerequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.PutAPIInteractionProfileResponse](../../models/operations/putapiinteractionprofileresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## PatchAPIInteractionProfile

### Example Usage

```go
package main

import(
	"context"
	"github.com/klucherev/logto"
	"github.com/klucherev/logto/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New()

    res, err := s.Interaction.PatchAPIInteractionProfile(ctx, operations.PatchAPIInteractionProfileRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PatchAPIInteractionProfileRequest](../../models/operations/patchapiinteractionprofilerequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.PatchAPIInteractionProfileResponse](../../models/operations/patchapiinteractionprofileresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteAPIInteractionProfile

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

    res, err := s.Interaction.DeleteAPIInteractionProfile(ctx)
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

**[*operations.DeleteAPIInteractionProfileResponse](../../models/operations/deleteapiinteractionprofileresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## PostAPIInteractionSubmit

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

    res, err := s.Interaction.PostAPIInteractionSubmit(ctx)
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

**[*operations.PostAPIInteractionSubmitResponse](../../models/operations/postapiinteractionsubmitresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## PostAPIInteractionConsent

### Example Usage

```go
package main

import(
	"context"
	"github.com/klucherev/logto"
	"github.com/klucherev/logto/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New()

    res, err := s.Interaction.PostAPIInteractionConsent(ctx, operations.PostAPIInteractionConsentRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.PostAPIInteractionConsentRequest](../../models/operations/postapiinteractionconsentrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.PostAPIInteractionConsentResponse](../../models/operations/postapiinteractionconsentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetAPIInteractionConsent

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

    res, err := s.Interaction.GetAPIInteractionConsent(ctx)
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

**[*operations.GetAPIInteractionConsentResponse](../../models/operations/getapiinteractionconsentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## PostAPIInteractionVerificationSocialAuthorizationURI

### Example Usage

```go
package main

import(
	"context"
	"github.com/klucherev/logto"
	"github.com/klucherev/logto/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New()

    res, err := s.Interaction.PostAPIInteractionVerificationSocialAuthorizationURI(ctx, operations.PostAPIInteractionVerificationSocialAuthorizationURIRequest{
        ConnectorID: "<id>",
        State: "Texas",
        RedirectURI: operations.PostAPIInteractionVerificationSocialAuthorizationURIRedirectURI{},
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

| Parameter                                                                                                                                                        | Type                                                                                                                                                             | Required                                                                                                                                                         | Description                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                            | :heavy_check_mark:                                                                                                                                               | The context to use for the request.                                                                                                                              |
| `request`                                                                                                                                                        | [operations.PostAPIInteractionVerificationSocialAuthorizationURIRequest](../../models/operations/postapiinteractionverificationsocialauthorizationurirequest.md) | :heavy_check_mark:                                                                                                                                               | The request object to use for the request.                                                                                                                       |
| `opts`                                                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                                                         | :heavy_minus_sign:                                                                                                                                               | The options for this request.                                                                                                                                    |

### Response

**[*operations.PostAPIInteractionVerificationSocialAuthorizationURIResponse](../../models/operations/postapiinteractionverificationsocialauthorizationuriresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## PostAPIInteractionVerificationVerificationCode

### Example Usage

```go
package main

import(
	"context"
	"github.com/klucherev/logto"
	"github.com/klucherev/logto/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New()

    res, err := s.Interaction.PostAPIInteractionVerificationVerificationCode(ctx, operations.CreatePostAPIInteractionVerificationVerificationCodeRequestPostAPIInteractionVerificationVerificationCodeRequestBody1(
        operations.PostAPIInteractionVerificationVerificationCodeRequestBody1{
            Email: "<value>",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `request`                                                                                                                                            | [operations.PostAPIInteractionVerificationVerificationCodeRequest](../../models/operations/postapiinteractionverificationverificationcoderequest.md) | :heavy_check_mark:                                                                                                                                   | The request object to use for the request.                                                                                                           |
| `opts`                                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                                             | :heavy_minus_sign:                                                                                                                                   | The options for this request.                                                                                                                        |

### Response

**[*operations.PostAPIInteractionVerificationVerificationCodeResponse](../../models/operations/postapiinteractionverificationverificationcoderesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## PostAPIInteractionVerificationTotp

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

    res, err := s.Interaction.PostAPIInteractionVerificationTotp(ctx)
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

**[*operations.PostAPIInteractionVerificationTotpResponse](../../models/operations/postapiinteractionverificationtotpresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## PostAPIInteractionVerificationWebauthnRegistration

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

    res, err := s.Interaction.PostAPIInteractionVerificationWebauthnRegistration(ctx)
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

**[*operations.PostAPIInteractionVerificationWebauthnRegistrationResponse](../../models/operations/postapiinteractionverificationwebauthnregistrationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## PostAPIInteractionVerificationWebauthnAuthentication

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

    res, err := s.Interaction.PostAPIInteractionVerificationWebauthnAuthentication(ctx)
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

**[*operations.PostAPIInteractionVerificationWebauthnAuthenticationResponse](../../models/operations/postapiinteractionverificationwebauthnauthenticationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## PostAPIInteractionBindMfa

### Example Usage

```go
package main

import(
	"context"
	"github.com/klucherev/logto"
	"github.com/klucherev/logto/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New()

    res, err := s.Interaction.PostAPIInteractionBindMfa(ctx, operations.CreatePostAPIInteractionBindMfaRequestPostAPIInteractionBindMfaRequestBody1(
        operations.PostAPIInteractionBindMfaRequestBody1{
            Type: "<value>",
            Code: "<value>",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.PostAPIInteractionBindMfaRequest](../../models/operations/postapiinteractionbindmfarequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.PostAPIInteractionBindMfaResponseResponse](../../models/operations/postapiinteractionbindmfaresponseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## PutAPIInteractionMfa

### Example Usage

```go
package main

import(
	"context"
	"github.com/klucherev/logto"
	"github.com/klucherev/logto/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New()

    res, err := s.Interaction.PutAPIInteractionMfa(ctx, operations.CreatePutAPIInteractionMfaRequestPutAPIInteractionMfaRequestBody2(
        operations.PutAPIInteractionMfaRequestBody2{
            Type: "<value>",
            ID: "<id>",
            RawID: "<id>",
            ClientExtensionResults: operations.PutAPIInteractionMfaClientExtensionResults{},
            Response: operations.InteractionPutAPIInteractionMfaResponse{
                ClientDataJSON: "<value>",
                AuthenticatorData: "<value>",
                Signature: "<value>",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PutAPIInteractionMfaRequest](../../models/operations/putapiinteractionmfarequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.PutAPIInteractionMfaResponseResponse](../../models/operations/putapiinteractionmfaresponseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## PutAPIInteractionMfaSkipped

### Example Usage

```go
package main

import(
	"context"
	"github.com/klucherev/logto"
	"github.com/klucherev/logto/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New()

    res, err := s.Interaction.PutAPIInteractionMfaSkipped(ctx, operations.PutAPIInteractionMfaSkippedRequest{
        MfaSkipped: true,
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.PutAPIInteractionMfaSkippedRequest](../../models/operations/putapiinteractionmfaskippedrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.PutAPIInteractionMfaSkippedResponse](../../models/operations/putapiinteractionmfaskippedresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## PostAPIInteractionSingleSignOnConnectorIDAuthorizationURL

### Example Usage

```go
package main

import(
	"context"
	"github.com/klucherev/logto"
	"github.com/klucherev/logto/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logto.New()

    res, err := s.Interaction.PostAPIInteractionSingleSignOnConnectorIDAuthorizationURL(ctx, "<id>", operations.PostAPIInteractionSingleSignOnConnectorIDAuthorizationURLRequestBody{
        State: "Georgia",
        RedirectURI: operations.PostAPIInteractionSingleSignOnConnectorIDAuthorizationURLRedirectURI{},
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

| Parameter                                                                                                                                                                          | Type                                                                                                                                                                               | Required                                                                                                                                                                           | Description                                                                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                              | :heavy_check_mark:                                                                                                                                                                 | The context to use for the request.                                                                                                                                                |
| `connectorID`                                                                                                                                                                      | *string*                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                 | The unique identifier of the connector.                                                                                                                                            |
| `requestBody`                                                                                                                                                                      | [operations.PostAPIInteractionSingleSignOnConnectorIDAuthorizationURLRequestBody](../../models/operations/postapiinteractionsinglesignonconnectoridauthorizationurlrequestbody.md) | :heavy_check_mark:                                                                                                                                                                 | N/A                                                                                                                                                                                |
| `opts`                                                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                 | The options for this request.                                                                                                                                                      |

### Response

**[*operations.PostAPIInteractionSingleSignOnConnectorIDAuthorizationURLResponse](../../models/operations/postapiinteractionsinglesignonconnectoridauthorizationurlresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## PostAPIInteractionSingleSignOnConnectorIDAuthentication

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

    res, err := s.Interaction.PostAPIInteractionSingleSignOnConnectorIDAuthentication(ctx, "<id>", map[string]any{

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

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `connectorID`                                            | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the connector.                  |
| `requestBody`                                            | map[string]*any*                                         | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PostAPIInteractionSingleSignOnConnectorIDAuthenticationResponse](../../models/operations/postapiinteractionsinglesignonconnectoridauthenticationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## PostAPIInteractionSingleSignOnConnectorIDRegistration

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

    res, err := s.Interaction.PostAPIInteractionSingleSignOnConnectorIDRegistration(ctx, "<id>")
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
| `connectorID`                                            | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the connector.                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PostAPIInteractionSingleSignOnConnectorIDRegistrationResponse](../../models/operations/postapiinteractionsinglesignonconnectoridregistrationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetAPIInteractionSingleSignOnConnectors

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

    res, err := s.Interaction.GetAPIInteractionSingleSignOnConnectors(ctx, "Olaf_Metz@gmail.com")
    if err != nil {
        log.Fatal(err)
    }
    if res.Strings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `email`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetAPIInteractionSingleSignOnConnectorsResponse](../../models/operations/getapiinteractionsinglesignonconnectorsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |