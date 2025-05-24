# Authn
(*Authn*)

## Overview

Authentication endpoints for third-party integrations and identity providers.

### Available Operations

* [GetHasuraAuth](#gethasuraauth) - Hasura auth hook endpoint
* [~~AssertSaml~~](#assertsaml) - SAML ACS endpoint (social) :warning: **Deprecated**
* [AssertSingleSignOnSaml](#assertsinglesignonsaml) - SAML ACS endpoint (SSO)

## GetHasuraAuth

The `HASURA_GRAPHQL_AUTH_HOOK` endpoint for Hasura auth. Use this endpoint to integrate Hasura's [webhook authentication flow](https://hasura.io/docs/latest/auth/authentication/webhook/).

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

    res, err := s.Authn.GetHasuraAuth(ctx, "<value>", nil)
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
| `resource`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `unauthorizedRole`                                       | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetHasuraAuthResponse](../../models/operations/gethasuraauthresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ~~AssertSaml~~

The Assertion Consumer Service (ACS) endpoint for Simple Assertion Markup Language (SAML) social connectors.

SAML social connectors are deprecated. Use the SSO SAML connector instead.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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

    res, err := s.Authn.AssertSaml(ctx, "<id>", operations.AssertSamlRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `connectorID`                                                                        | *string*                                                                             | :heavy_check_mark:                                                                   | The unique identifier of the connector.                                              |
| `requestBody`                                                                        | [operations.AssertSamlRequestBody](../../models/operations/assertsamlrequestbody.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.AssertSamlResponse](../../models/operations/assertsamlresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## AssertSingleSignOnSaml

The Assertion Consumer Service (ACS) endpoint for Simple Assertion Markup Language (SAML) single sign-on (SSO) connectors.

This endpoint is used to complete the SAML SSO authentication flow. It receives the SAML assertion response from the identity provider (IdP) and redirects the user to complete the authentication flow.

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

    res, err := s.Authn.AssertSingleSignOnSaml(ctx, "<id>", operations.AssertSingleSignOnSamlRequestBody{
        SAMLResponse: "<value>",
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `connectorID`                                                                                                | *string*                                                                                                     | :heavy_check_mark:                                                                                           | The unique identifier of the connector.                                                                      |
| `requestBody`                                                                                                | [operations.AssertSingleSignOnSamlRequestBody](../../models/operations/assertsinglesignonsamlrequestbody.md) | :heavy_check_mark:                                                                                           | N/A                                                                                                          |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.AssertSingleSignOnSamlResponse](../../models/operations/assertsinglesignonsamlresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |