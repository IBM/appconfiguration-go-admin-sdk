# IBM Cloud App Configuration Go Admin SDK

The IBM Cloud App Configuration Go Admin SDK allows developers to programmatically manage the [App Configuration](https://cloud.ibm.com/apidocs/app-configuration) service

## Table of Contents

  - [Prerequisites](#prerequisites)
  - [Overview](#overview)
  - [Installation](#installation)
  - [Import the SDK](#import-the-sdk)
  - [Initialize SDK](#initialize-sdk)
  - [Using the SDK](#using-the-sdk)
  - [License](#license)


## Prerequisites

* An [IBM Cloud](https://cloud.ibm.com/registration) account.
* An [App Configuration](https://cloud.ibm.com/docs/app-configuration) instance.
* Go version 1.15 or above.

## Overview

IBM Cloud App Configuration is a centralized feature management and configuration service on [IBM Cloud](https://www.cloud.ibm.com) for use with web and mobile applications, microservices, and distributed environments.

Use the Go Admin SDK to manage the App Configuration service instance. The Go Admin SDK provides APIs to define and manage feature flags, collections and segments. Alternately, you can also use the IBM Cloud App Configuration CLI to manage the App Configuration service instance. You can find more information about the CLI [here.](https://cloud.ibm.com/docs/app-configuration?topic=app-configuration-cli-plugin-app-configuration-cli) 

## Installation

Install using the command.

```bash
go get -u github.com/IBM/appconfiguration-go-admin-sdk
```

## Import the SDK

To import the module 

```go	
import "github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
```

then run `go mod tidy` to download and install the new dependency and update your Go application's
`go.mod` file.

## Initialize SDK
Initialize the sdk to connect with your App Configuration service instance.
```go
var appConfigurationServiceInstance *appconfigurationv1.AppConfigurationV1


func init() {
	authenticator := &core.IamAuthenticator{
		ApiKey: "apikey",
	}

	options := &appconfigurationv1.AppConfigurationV1Options{ 
		Authenticator: authenticator, 
  		URL: "url"
	}

	appconfigurationServiceInstance, err := appconfigurationv1.NewAppConfigurationV1(options)

	if err != nil {
		panic(err)
	}
}

```

- apikey : apikey of the App Configuration service. Get it from the service credentials section of the dashboard.
- url : url of the App Configuration Instance. URL instance can found from [here](https://cloud.ibm.com/apidocs/app-configuration#endpoint-url)

## Using the SDK

### Note 
Every (non-delete) SDK API call gives 3 items in response. 1st item is the result which is the actual item returned from the server available for consumption. 2nd item is the response along with the metadata recieved from the server(including response code). 3rd item is the error (if any). For delete SDK API calls, response and error is received from the server as the result is empty. It is advisable to check if there are no errors and response code from the server is 2XX before using the result for further operations.

Following is how one can access properties of the result object from SDK API call -
- if the property is a reference, dereference and then use it.
- if the property is not a reference, use it as it is.


Steps to use the SDK method's -
- Create the input request object with all needed parameters either from Constructors or from struct
- call the API.
- Check for error from 'error' and status code from 'response' & consume 'result' accordingly.

SDK Methods to consume ->
- Create*Item*
- Get*Item*s
- Get*Item*
- Update*Item*
- Delete*Item*

where *Item* can be replaced with Collection, Feature or Segment.

Refer [this](https://cloud.ibm.com/apidocs/app-configuration) for details on the input parameters for each SDK method.

**For more details on the above points, refer 'sampleProgram.go' in 'examples' directory.**
### Create Collection

```go
createCollectionOptionsModel := appConfigurationServiceInstance.NewCreateCollectionOptions(name)
createCollectionOptionsModel.SetCollectionID(id)
createCollectionOptionsModel.SetDescription(description)
createCollectionOptionsModel.SetTags(tags)
result, response, error := appConfigurationServiceInstance.CreateCollection(createCollectionOptionsModel)
```

### Get Collections 
You can list all collections with expand as true
```go
getCollectionsOptionsModel := appConfigurationServiceInstance.NewGetCollectionsOptions()
getCollectionsOptionsModel.SetExpand("true")   // setting expand option as "true"
result, response, error := appConfigurationServiceInstance.GetCollections(getCollectionsOptionsModel)
```

### Create segment

```go
ruleArray, _ := appConfigurationServiceInstance.NewRuleArray(attributeName, operator, values)
createSegmentOptionsModel := appConfigurationServiceInstance.NewCreateSegmentOptions(name, id, description, tags, []appconfigurationv1.RuleArray{*ruleArray})
result, response, error := appConfigurationServiceInstance.CreateSegment(createSegmentOptionsModel)
```

### Create Feature

```go
ruleArray, _ := appConfigurationServiceInstance.NewRule(segments)
segmentRuleArray, _ := appConfigurationServiceInstance.NewSegmentRule([]appconfigurationv1.Rule{*ruleArray}, value, order)
collectionArray, _ := appConfigurationServiceInstance.NewCollection(collectionId, enabledInCollection)
createFeatureOptionsModel := appConfigurationServiceInstance.NewCreateFeatureOptions(name, id, description, typeOfFeature, enabledValue, disabledValue, tags, []appconfigurationv1.SegmentRule{*segmentRuleArray}, []appconfigurationv1.Collection{*collectionArray})
result, response, error := appConfigurationServiceInstance.CreateFeature(createFeatureOptionsModel)
```

### Update Feature
```go
ruleArray, _ := appConfigurationServiceInstance.NewRule(segments)
segmentRuleArray, _ := appConfigurationServiceInstance.NewSegmentRule([]appconfigurationv1.Rule{*ruleArray}, value, order)
collectionArray, _ := appConfigurationServiceInstance.NewCollectionWithDeletedFlag(collectionName, enabledInCollection, deletedFlag)
updateFeatureOptionsModel := appConfigurationServiceInstance.NewUpdateFeatureOptions(id, name, description, typeOfFeature, enabledValue, disabledValue, tags, []appconfigurationv1.SegmentRule{*segmentRuleArray}, []appconfigurationv1.CollectionWithDeletedFlag{*collectionArray})
result, response, error := appConfigurationServiceInstance.UpdateFeature(updateFeatureOptionsModel)
```

### Get Feature 
```go
getFeatureOptionsModel := appConfigurationServiceInstance.NewGetFeatureOptions(featureId)
result, response, error := appConfigurationServiceInstance.GetFeature(getFeatureOptionsModel)
```

### Delete Segment
```go
deleteasegmentOptionsModel := new(appconfigurationv1.DeleteSegmentOptions)
deleteasegmentOptionsModel.SegmentID = core.StringPtr(id)
response, error := appConfigurationServiceInstance.DeleteSegment(deleteasegmentOptionsModel)
```

## License

This project is released under the Apache 2.0 license. The license's full text can be found in [LICENSE](/LICENSE)