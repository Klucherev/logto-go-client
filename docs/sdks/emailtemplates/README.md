# EmailTemplates
(*EmailTemplates*)

## Overview

Manage custom i18n email templates for various types of emails, such as sign-in verification codes and password resets.

### Available Operations

* [ReplaceEmailTemplates](#replaceemailtemplates) - Replace email templates
* [ListEmailTemplates](#listemailtemplates) - Get email templates
* [DeleteEmailTemplates](#deleteemailtemplates) - Delete email templates
* [GetEmailTemplate](#getemailtemplate) - Get email template by ID
* [DeleteEmailTemplate](#deleteemailtemplate) - Delete an email template
* [UpdateEmailTemplateDetails](#updateemailtemplatedetails) - Update email template details

## ReplaceEmailTemplates

Create or replace a list of email templates. If an email template with the same language tag and template type already exists, its details will be updated.

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

    res, err := s.EmailTemplates.ReplaceEmailTemplates(ctx, operations.ReplaceEmailTemplatesRequest{
        Templates: []operations.Template{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseBodies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ReplaceEmailTemplatesRequest](../../models/operations/replaceemailtemplatesrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ReplaceEmailTemplatesResponse](../../models/operations/replaceemailtemplatesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListEmailTemplates

Get the list of email templates.

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

    res, err := s.EmailTemplates.ListEmailTemplates(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseBodies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                   | Type                                                                                                                        | Required                                                                                                                    | Description                                                                                                                 |
| --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                                       | :heavy_check_mark:                                                                                                          | The context to use for the request.                                                                                         |
| `languageTag`                                                                                                               | **string*                                                                                                                   | :heavy_minus_sign:                                                                                                          | The language tag of the email template, e.g., `en` or `fr`.                                                                 |
| `templateType`                                                                                                              | [*operations.ListEmailTemplatesQueryParamTemplateType](../../models/operations/listemailtemplatesqueryparamtemplatetype.md) | :heavy_minus_sign:                                                                                                          | The type of the email template, e.g. `SignIn` or `ForgotPassword`                                                           |
| `opts`                                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                                    | :heavy_minus_sign:                                                                                                          | The options for this request.                                                                                               |

### Response

**[*operations.ListEmailTemplatesResponse](../../models/operations/listemailtemplatesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteEmailTemplates

Bulk delete email templates by their language tag and template type.

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

    res, err := s.EmailTemplates.DeleteEmailTemplates(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                   | Type                                                                                                        | Required                                                                                                    | Description                                                                                                 |
| ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                       | :heavy_check_mark:                                                                                          | The context to use for the request.                                                                         |
| `languageTag`                                                                                               | **string*                                                                                                   | :heavy_minus_sign:                                                                                          | The language tag of the email template, e.g., `en` or `fr`.                                                 |
| `templateType`                                                                                              | [*operations.DeleteEmailTemplatesTemplateType](../../models/operations/deleteemailtemplatestemplatetype.md) | :heavy_minus_sign:                                                                                          | The type of the email template, e.g. `SignIn` or `ForgotPassword`                                           |
| `opts`                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                    | :heavy_minus_sign:                                                                                          | The options for this request.                                                                               |

### Response

**[*operations.DeleteEmailTemplatesResponse](../../models/operations/deleteemailtemplatesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetEmailTemplate

Get the email template by its ID.

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

    res, err := s.EmailTemplates.GetEmailTemplate(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the email template.             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetEmailTemplateResponse](../../models/operations/getemailtemplateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteEmailTemplate

Delete an email template by its ID.

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

    res, err := s.EmailTemplates.DeleteEmailTemplate(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the email template.             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteEmailTemplateResponse](../../models/operations/deleteemailtemplateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateEmailTemplateDetails

Update the details of an email template by its ID.

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

    res, err := s.EmailTemplates.UpdateEmailTemplateDetails(ctx, "<id>", operations.UpdateEmailTemplateDetailsRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `id`                                                                                                                 | *string*                                                                                                             | :heavy_check_mark:                                                                                                   | The unique identifier of the email template.                                                                         |
| `requestBody`                                                                                                        | [operations.UpdateEmailTemplateDetailsRequestBody](../../models/operations/updateemailtemplatedetailsrequestbody.md) | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.UpdateEmailTemplateDetailsResponse](../../models/operations/updateemailtemplatedetailsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |