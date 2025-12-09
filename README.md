# IBM Cloud App Configuration Go Admin SDK

Go client library to interact with the
various [IBM Cloud® App Configuration APIs](https://cloud.ibm.com/apidocs/app-configuration).

## Table of Contents

<!--
  The TOC below is generated using the `markdown-toc` node package.

      https://github.com/jonschlinkert/markdown-toc

  You should regenerate the TOC after making changes to this file.

      npx markdown-toc -i README.md
  -->

<!-- toc -->

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
    * [Go modules](#go-modules)
    * [`go get` command](#go-get-command)
- [Using the SDK](#using-the-sdk)
- [Questions](#questions)
- [Issues](#issues)
- [Open source @ IBM](#open-source--ibm)
- [Contributing](#contributing)
- [License](#license)

<!-- tocstop -->

## Overview

The IBM Cloud App Configuration Go SDK allows developers to programmatically manage
the [App Configuration](https://cloud.ibm.com/apidocs/app-configuration) service. Alternately, you can also use the IBM
Cloud App Configuration CLI to manage the App Configuration service instance. You can find more information about the
CLI [here.](https://cloud.ibm.com/docs/app-configuration?topic=app-configuration-app-configuration-cli)

## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration

* An [IBM Cloud][ibm-cloud-onboarding] account.
* An [App Configuration service instance](https://cloud.ibm.com/catalog/services/app-configuration).
* An IAM API key to allow the SDK to access your account. Create one [here](https://cloud.ibm.com/iam/apikeys).
* Go version 1.23 or above.

## Installation

The current version of this SDK: 0.5.3

**Note: The v1.x.x versions of the App Configuration Go Admin SDK have been retracted. Use the latest available version
of the SDK.**

### Go modules

If your application uses Go modules for dependency management (recommended), just add the import.
Here is an example:

```go
import (
"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
)
```

Next, run `go build` or `go mod tidy` to download and install the new dependencies and update your application's
`go.mod` file.

In the example above, the `appconfigurationv1` part of the import path is the package name
associated with the App Configuration service.

### `go get` command

Alternatively, you can use the `go get` command to download and install the appropriate packages needed by your
application:

```
go get -u github.com/IBM/appconfiguration-go-admin-sdk
```

## Using the SDK

Some capabilities such as Percentage rollout & Git configs are applicable only for specific plans (Lite, Standard &
enterprise). [See here](https://cloud.ibm.com/docs/app-configuration?topic=app-configuration-ac-faqs-usage#faq-ac-capabilities)
for full list of capabilities that are plan wise.

### Basic usage

- All methods return a response and an error. The response contains the body, the headers, the status code, and the
  status text.
- Use the `URL` parameter to set
  the [Endpoint URL](https://test.cloud.ibm.com/apidocs/app-configuration?code=go#endpoint-url) that is specific to your
  App Configuration service instance.

#### Examples

Construct a service client and use it to create, retrieve and manage resources from your App Configuration instance.

Here's an example `main.go` file:

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/IBM/go-sdk-core/v5/core"
)

func main() {

	authenticator := &core.IamAuthenticator{
		ApiKey: "<IBM_CLOUD_API_KEY>",
	}
	options := &appconfigurationv1.AppConfigurationV1Options{
		Authenticator: authenticator,
		URL:           "https://" + region + ".apprapp.cloud.ibm.com/apprapp/feature/v1/instances/" + guid,
	}
	appConfigurationService, err := appconfigurationv1.NewAppConfigurationV1(options)
	if err != nil {
		panic(err)
	}

	createEnvironmentOptions := appConfigurationService.NewCreateEnvironmentOptions(
		"Dev environment",
		"dev",
	)
	createEnvironmentOptions.SetDescription("Dev environment description")
	createEnvironmentOptions.SetTags("development")
	createEnvironmentOptions.SetColorCode("#FDD13A")

	environment, response, err := appConfigurationService.CreateEnvironment(createEnvironmentOptions)
	if err != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(environment, "", "  ")
	fmt.Println(string(b))

```

- guid : Instance ID of the App Configuration instance.
- region : Region of the App Configuration instance.

Replace the `URL` and `ApiKey` values. Then run the `go run main.go` command to compile and run your Go program.

Examples for rest APIs can be found [here](https://cloud.ibm.com/apidocs/app-configuration?code=go).

Also, more examples are documented in [sampleProgram.go](examples/sampleProgram.go) file of this repo.

For more information and IBM Cloud SDK usage examples for Go, see
the [IBM Cloud SDK Common documentation](https://github.com/IBM/ibm-cloud-sdk-common/blob/master/README.md).

### Using private endpoints

If you
enable [service endpoints](https://cloud.ibm.com/docs/account?topic=account-vrf-service-endpoint&interface=ui#service-endpoint)
in your account, you can send API requests over the IBM Cloud private network. While constructing the service client the
endpoint URLs of the IAM(authenticator) & App Configuration(service) should be modified to
point to private endpoints. See below

```go
    authenticator := &core.IamAuthenticator{
        ApiKey: "<IBM_CLOUD_API_KEY>",
        URL: "https://private.iam.cloud.ibm.com",
    }
    options := &appconfigurationv1.AppConfigurationV1Options{
        Authenticator: authenticator,
        URL:           "https://private." + region + ".apprapp.cloud.ibm.com/apprapp/feature/v1/instances/" + guid,
	}
```

## Questions

If you are having difficulties using this SDK or have a question about the IBM Cloud services,
please ask a question at
[Stack Overflow](http://stackoverflow.com/questions/ask?tags=ibm-cloud).

## Issues

If you encounter an issue with the project, you are welcome to submit a
[bug report](https://github.com/IBM/appconfiguration-go-admin-sdk/issues).
Before that, please search for similar issues. It's possible that someone has already reported the problem.

## Open source @ IBM

Find more open source projects on the [IBM Github Page](http://ibm.github.io/)

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).

## License

This SDK project is released under the Apache 2.0 license.
The license's full text can be found in [LICENSE](LICENSE).
