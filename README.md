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
* Go version 1.16 or above.

## Overview

IBM Cloud App Configuration is a centralized feature management and configuration service on [IBM Cloud](https://cloud.ibm.com) for use with web and mobile applications, microservices, and distributed environments.

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

// Use any one of bearer token or API Key
func init() {

	// Using Bearer Token
	authenticator := &core.BearerTokenAuthenticator{
		BearerToken: "<authToken>",
	}
	
	// Using API Key
	authenticator := &core.IamAuthenticator{
		ApiKey: "<apikey>",
	}

	options := &appconfigurationv1.AppConfigurationV1Options{ 
		Authenticator: authenticator, 
  		URL: "https://" + region + ".apprapp.cloud.ibm.com/apprapp/feature/v1/instances/" + guid,
	}

	appConfigurationServiceInstance, err := appconfigurationv1.NewAppConfigurationV1(options)

	if err != nil {
		panic(err)
	}
}

```

- authToken : authToken of the App Configuration service. Get it from the service credentials section of the dashboard. Choose any option from APIKey or Bearer Token to authenticate.
- guid : ID of the App Configuration Instance.
- region : Region of the App Configuration Instance

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
- List*Item*
- Get*Item*
- Update*Item*
- Update*Item*Values
- Delete*Item*

where *Item* can be replaced with Collection, Property, Environment, Feature or Segment.

Refer [this](https://cloud.ibm.com/apidocs/app-configuration) for details on the input parameters for each SDK method.

Note -> You need to have the required access (READER/WRITER/MANAGER/CONFIG_OPERATOR) to the instances for respective APIs.

**For more details on the above points, refer 'sampleProgram.go' in 'examples' directory.**
### Create Collection

```go
createCollectionOptionsModel := appConfigurationServiceInstance.NewCreateCollectionOptions(name, collectionId)
createCollectionOptionsModel.SetDescription(description)
createCollectionOptionsModel.SetTags(tags)
result, response, err := appConfigurationServiceInstance.CreateCollection(createCollectionOptionsModel)
```

### List Collections 
You can list all collections with expand as true
```go
getCollectionsOptionsModel := appConfigurationServiceInstance.NewListCollectionsOptions()
getCollectionsOptionsModel.SetExpand(true)   // setting expand option as "true"
result, response, err := appConfigurationServiceInstance.ListCollections(getCollectionsOptionsModel)
```

### Create segment

```go
ruleArray, _ := appConfigurationServiceInstance.NewRule(attributeName, operator, values)
createSegmentOptionsModel := appConfigurationServiceInstance.NewCreateSegmentOptions()
createSegmentOptionsModel.SetName(name)
createSegmentOptionsModel.SetDescription(description)
createSegmentOptionsModel.SetTags(tags)
createSegmentOptionsModel.SetSegmentID(id)
createSegmentOptionsModel.SetRules([]appconfigurationv1.Rule{*ruleArray})
result, response, err := appConfigurationServiceInstance.CreateSegment(createSegmentOptionsModel)
```

### Create Feature

```go
ruleArray, _ := appConfigurationServiceInstance.NewTargetSegments(segments)
segmentRuleArray, _ := appConfigurationServiceInstance.NewSegmentRule([]appconfigurationv1.TargetSegments{*ruleArray}, value, order)
collectionArray, _ := appConfigurationServiceInstance.NewCollectionRef(collectionId)
createFeatureOptionsModel := appConfigurationServiceInstance.NewCreateFeatureOptions(environmentId, name, id, typeOfFeature, enabledValue, disabledValue)
createFeatureOptionsModel.SetTags(tags)
createFeatureOptionsModel.SetDescription(description)
createFeatureOptionsModel.SetSegmentRules([]appconfigurationv1.SegmentRule{*segmentRuleArray})
createFeatureOptionsModel.SetCollections([]appconfigurationv1.CollectionRef{*collectionArray})
result, response, err := appConfigurationServiceInstance.CreateFeature(createFeatureOptionsModel)
```

### Update Feature
```go
ruleArray, _ := appConfigurationServiceInstance.NewTargetSegments(segments)
segmentRuleArray, _ := appConfigurationServiceInstance.NewSegmentRule([]appconfigurationv1.TargetSegments{*ruleArray}, value, order)
collectionArray, _ := appConfigurationServiceInstance.NewCollectionRef(collectionId)
updateFeatureOptionsModel := appConfigurationServiceInstance.NewUpdateFeatureOptions(environmentId, id)
updateFeatureOptionsModel.SetName(name)
updateFeatureOptionsModel.SetDescription(description)
updateFeatureOptionsModel.SetTags(tags)
updateFeatureOptionsModel.SetDisabledValue(disabledValue)
updateFeatureOptionsModel.SetEnabledValue(enabledValue)
updateFeatureOptionsModel.SetSegmentRules([]appconfigurationv1.SegmentRule{*segmentRuleArray})
updateFeatureOptionsModel.SetCollections([]appconfigurationv1.CollectionRef{*collectionArray})
result, response, err := appConfigurationServiceInstance.UpdateFeature(updateFeatureOptionsModel)
```

### Update Environment
```go
updateEnvironmentOptionsModel := appConfigurationServiceInstance.NewUpdateEnvironmentOptions(environmentId)
updateEnvironmentOptionsModel.SetName(name)
updateEnvironmentOptionsModel.SetDescription(description)
updateEnvironmentOptionsModel.SetTags(tags)
updateEnvironmentOptionsModel.SetColorCode(colorCode)
result, response, err := appConfigurationServiceInstance.UpdateEnvironment(updateEnvironmentOptionsModel)
```

### Get Feature 
```go
getFeatureOptionsModel := appConfigurationServiceInstance.NewGetFeatureOptions(environmentId, featureId)
result, response, err := appConfigurationServiceInstance.GetFeature(getFeatureOptionsModel)
```

### Delete Segment
```go
deleteSegmentOptionsModel := appConfigurationServiceInstance.NewDeleteSegmentOptions(segmentId)
response, err := appConfigurationServiceInstance.DeleteSegment(deleteSegmentOptionsModel)
```

### Toggle Feature
```go
toggleFeatureOptionsModel := appConfigurationServiceInstance.NewToggleFeatureOptions(environmentId, featureId)
toggleFeatureOptionsModel.SetEnabled(enableFlag)
result, response, err := appConfigurationServiceInstance.ToggleFeature(toggleFeatureOptionsModel)
```

### Patch Property
```go
ruleArray, _ := appConfigurationServiceInstance.NewTargetSegments(segments)
segmentRuleArray, _ := appConfigurationServiceInstance.NewSegmentRule([]appconfigurationv1.TargetSegments{*ruleArray}, value, order)
patchPropertyOptionsModel := appConfigurationServiceInstance.NewUpdatePropertyValuesOptions(environmentId, propertyId)
patchPropertyOptionsModel.SetName(name)
patchPropertyOptionsModel.SetDescription(description)
patchPropertyOptionsModel.SetTags(tags)
patchPropertyOptionsModel.SetValue(valueOfProperty)
patchPropertyOptionsModel.SetSegmentRules([]appconfigurationv1.SegmentRule{*segmentRuleArray})
result, response, err := appConfigurationServiceInstance.UpdatePropertyValues(patchPropertyOptionsModel)
```

## License

This project is released under the Apache 2.0 license. The license's full text can be found in [LICENSE](/LICENSE)
