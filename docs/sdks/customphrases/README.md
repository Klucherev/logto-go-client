# CustomPhrases
(*CustomPhrases*)

## Overview

Endpoints for managing custom phrases that allow you to customize the phrases displayed in the sign-in experience.

See [Localized language](https://docs.logto.io/docs/recipes/customize-sie/localized-language/) to learn more about custom phrases for localization.

### Available Operations

* [ListCustomPhrases](#listcustomphrases) - Get all custom phrases
* [GetCustomPhrase](#getcustomphrase) - Get custom phrases
* [ReplaceCustomPhrase](#replacecustomphrase) - Upsert custom phrases
* [DeleteCustomPhrase](#deletecustomphrase) - Delete custom phrase

## ListCustomPhrases

Get all custom phrases for all languages.

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

    res, err := s.CustomPhrases.ListCustomPhrases(ctx)
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

**[*operations.ListCustomPhrasesResponse](../../models/operations/listcustomphrasesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetCustomPhrase

Get custom phrases for the specified language tag.

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

    res, err := s.CustomPhrases.GetCustomPhrase(ctx, operations.GetCustomPhraseLanguageTagGaIe)
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
| `languageTag`                                                                                  | [operations.GetCustomPhraseLanguageTag](../../models/operations/getcustomphraselanguagetag.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetCustomPhraseResponse](../../models/operations/getcustomphraseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ReplaceCustomPhrase

Upsert custom phrases for the specified language tag. Upsert means that if the custom phrases already exist, they will be updated. Otherwise, they will be created.

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

    res, err := s.CustomPhrases.ReplaceCustomPhrase(ctx, operations.ReplaceCustomPhraseLanguageTagEuEs, components.TranslationObject{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                   | Type                                                                                                                        | Required                                                                                                                    | Description                                                                                                                 | Example                                                                                                                     |
| --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                                       | :heavy_check_mark:                                                                                                          | The context to use for the request.                                                                                         |                                                                                                                             |
| `languageTag`                                                                                                               | [operations.ReplaceCustomPhraseLanguageTag](../../models/operations/replacecustomphraselanguagetag.md)                      | :heavy_check_mark:                                                                                                          | N/A                                                                                                                         |                                                                                                                             |
| `translationObject`                                                                                                         | [components.TranslationObject](../../models/components/translationobject.md)                                                | :heavy_check_mark:                                                                                                          | N/A                                                                                                                         | {<br/>"input": {<br/>"username": "Username",<br/>"password": "Password"<br/>},<br/>"action": {<br/>"sign_in": "Sign In",<br/>"continue": "Continue"<br/>}<br/>} |
| `opts`                                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                                    | :heavy_minus_sign:                                                                                                          | The options for this request.                                                                                               |                                                                                                                             |

### Response

**[*operations.ReplaceCustomPhraseResponse](../../models/operations/replacecustomphraseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteCustomPhrase

Delete custom phrases for the specified language tag.

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

    res, err := s.CustomPhrases.DeleteCustomPhrase(ctx, operations.DeleteCustomPhraseLanguageTagSySy)
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
| `languageTag`                                                                                        | [operations.DeleteCustomPhraseLanguageTag](../../models/operations/deletecustomphraselanguagetag.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.DeleteCustomPhraseResponse](../../models/operations/deletecustomphraseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |