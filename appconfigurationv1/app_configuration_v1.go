/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * IBM OpenAPI SDK Code Generator Version: 3.22.0-937b9a1c-20201211-223043
 */

// Package appconfigurationv1 : Operations and models for the AppConfigurationV1 service
package appconfigurationv1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/IBM/appconfiguration-go-admin-sdk/common"

	"github.com/IBM/go-sdk-core/v4/core"
)

// AppConfigurationV1 : ReST APIs for App Configuration
//
// Version: 1.0
// See: https://{DomainName}/docs/app-configuration/
type AppConfigurationV1 struct {
	Service *core.BaseService
}

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "app_configuration"

func GetDefaultServiceName() string {
	return DefaultServiceName
}

// AppConfigurationV1Options : Service options
type AppConfigurationV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewAppConfigurationV1UsingExternalConfig : constructs an instance of AppConfigurationV1 with passed in options and external configuration.
func NewAppConfigurationV1UsingExternalConfig(options *AppConfigurationV1Options) (appConfiguration *AppConfigurationV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	appConfiguration, err = NewAppConfigurationV1(options)
	if err != nil {
		return
	}

	err = appConfiguration.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = appConfiguration.Service.SetServiceURL(options.URL)
	}
	return
}

// NewAppConfigurationV1 : constructs an instance of AppConfigurationV1 with passed in options.
func NewAppConfigurationV1(options *AppConfigurationV1Options) (service *AppConfigurationV1, err error) {
	serviceOptions := &core.ServiceOptions{
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &AppConfigurationV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "appConfiguration" suitable for processing requests.
func (appConfiguration *AppConfigurationV1) Clone() *AppConfigurationV1 {
	if core.IsNil(appConfiguration) {
		return nil
	}
	clone := *appConfiguration
	clone.Service = appConfiguration.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (appConfiguration *AppConfigurationV1) SetServiceURL(url string) error {
	return appConfiguration.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (appConfiguration *AppConfigurationV1) GetServiceURL() string {
	return appConfiguration.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (appConfiguration *AppConfigurationV1) SetDefaultHeaders(headers http.Header) {
	appConfiguration.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (appConfiguration *AppConfigurationV1) SetEnableGzipCompression(enableGzip bool) {
	appConfiguration.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (appConfiguration *AppConfigurationV1) GetEnableGzipCompression() bool {
	return appConfiguration.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (appConfiguration *AppConfigurationV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	appConfiguration.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (appConfiguration *AppConfigurationV1) DisableRetries() {
	appConfiguration.Service.DisableRetries()
}

// PageInfo : PageInfo struct
type PageInfo struct {
	// total count of the records.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// total page count.
	Count *int64 `json:"count" validate:"required"`
}

// UnmarshalPageInfo unmarshals an instance of PageInfo from the specified map of raw messages.
func UnmarshalPageInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PageInfo)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Rule : Rule struct
type Rule struct {
	// Rules object.
	Segments []string `json:"segments" validate:"required"`
}

// NewRule : Instantiate Rule (Generic Model Constructor)
func (*AppConfigurationV1) NewRule(segments []string) (model *Rule, err error) {
	model = &Rule{
		Segments: segments,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalRule unmarshals an instance of Rule from the specified map of raw messages.
func UnmarshalRule(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Rule)
	err = core.UnmarshalPrimitive(m, "segments", &obj.Segments)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleArray : RuleArray struct
type RuleArray struct {
	// Attribute name.
	AttributeName *string `json:"attribute_name" validate:"required"`

	// Operator to be used for the evaluation.
	Operator *string `json:"operator" validate:"required"`

	// Rule.
	Values []string `json:"values" validate:"required"`
}

// NewRuleArray : Instantiate RuleArray (Generic Model Constructor)
func (*AppConfigurationV1) NewRuleArray(attributeName string, operator string, values []string) (model *RuleArray, err error) {
	model = &RuleArray{
		AttributeName: core.StringPtr(attributeName),
		Operator:      core.StringPtr(operator),
		Values:        values,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalRuleArray unmarshals an instance of RuleArray from the specified map of raw messages.
func UnmarshalRuleArray(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleArray)
	err = core.UnmarshalPrimitive(m, "attribute_name", &obj.AttributeName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operator", &obj.Operator)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "values", &obj.Values)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateCollection : Create Collection
// Create a collection.
func (appConfiguration *AppConfigurationV1) CreateCollection(createCollectionOptions *CreateCollectionOptions) (result *CreateCollection, response *core.DetailedResponse, err error) {
	return appConfiguration.CreateCollectionWithContext(context.Background(), createCollectionOptions)
}

// CreateCollectionWithContext is an alternate form of the CreateCollection method which supports a Context parameter
func (appConfiguration *AppConfigurationV1) CreateCollectionWithContext(ctx context.Context, createCollectionOptions *CreateCollectionOptions) (result *CreateCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createCollectionOptions, "createCollectionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createCollectionOptions, "createCollectionOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appConfiguration.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appConfiguration.Service.Options.URL, `/collections`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createCollectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders(DefaultServiceName, "V1", "CreateCollection")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createCollectionOptions.Name != nil {
		body["name"] = createCollectionOptions.Name
	}
	if createCollectionOptions.CollectionID != nil {
		body["collection_id"] = createCollectionOptions.CollectionID
	}
	if createCollectionOptions.Description != nil {
		body["description"] = createCollectionOptions.Description
	}
	if createCollectionOptions.Tags != nil {
		body["tags"] = createCollectionOptions.Tags
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appConfiguration.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateCollection)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetCollections : Get list of collections
// Get list of collections.
func (appConfiguration *AppConfigurationV1) GetCollections(getCollectionsOptions *GetCollectionsOptions) (result *GetCollections, response *core.DetailedResponse, err error) {
	return appConfiguration.GetCollectionsWithContext(context.Background(), getCollectionsOptions)
}

// GetCollectionsWithContext is an alternate form of the GetCollections method which supports a Context parameter
func (appConfiguration *AppConfigurationV1) GetCollectionsWithContext(ctx context.Context, getCollectionsOptions *GetCollectionsOptions) (result *GetCollections, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getCollectionsOptions, "getCollectionsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appConfiguration.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appConfiguration.Service.Options.URL, `/collections`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCollectionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders(DefaultServiceName, "V1", "GetCollections")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getCollectionsOptions.Size != nil {
		builder.AddQuery("size", fmt.Sprint(*getCollectionsOptions.Size))
	}
	if getCollectionsOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*getCollectionsOptions.Offset))
	}
	if getCollectionsOptions.Features != nil {
		builder.AddQuery("features", fmt.Sprint(*getCollectionsOptions.Features))
	}
	if getCollectionsOptions.Tags != nil {
		builder.AddQuery("tags", fmt.Sprint(*getCollectionsOptions.Tags))
	}
	if getCollectionsOptions.Expand != nil {
		builder.AddQuery("expand", fmt.Sprint(*getCollectionsOptions.Expand))
	}
	if getCollectionsOptions.Include != nil {
		builder.AddQuery("include", fmt.Sprint(*getCollectionsOptions.Include))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appConfiguration.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetCollections)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateCollection : Update Collection
// Update the collection name, tags and description. Collection Id cannot be updated.
func (appConfiguration *AppConfigurationV1) UpdateCollection(updateCollectionOptions *UpdateCollectionOptions) (result *UpdateCollection, response *core.DetailedResponse, err error) {
	return appConfiguration.UpdateCollectionWithContext(context.Background(), updateCollectionOptions)
}

// UpdateCollectionWithContext is an alternate form of the UpdateCollection method which supports a Context parameter
func (appConfiguration *AppConfigurationV1) UpdateCollectionWithContext(ctx context.Context, updateCollectionOptions *UpdateCollectionOptions) (result *UpdateCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateCollectionOptions, "updateCollectionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateCollectionOptions, "updateCollectionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"collection_id": *updateCollectionOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appConfiguration.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appConfiguration.Service.Options.URL, `/collections/{collection_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateCollectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders(DefaultServiceName, "V1", "UpdateCollection")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateCollectionOptions.Name != nil {
		body["name"] = updateCollectionOptions.Name
	}
	if updateCollectionOptions.Description != nil {
		body["description"] = updateCollectionOptions.Description
	}
	if updateCollectionOptions.Tags != nil {
		body["tags"] = updateCollectionOptions.Tags
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appConfiguration.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalUpdateCollection)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetCollection : Get Collection
// Retrieve the details of the collection.
func (appConfiguration *AppConfigurationV1) GetCollection(getCollectionOptions *GetCollectionOptions) (result *GetCollection, response *core.DetailedResponse, err error) {
	return appConfiguration.GetCollectionWithContext(context.Background(), getCollectionOptions)
}

// GetCollectionWithContext is an alternate form of the GetCollection method which supports a Context parameter
func (appConfiguration *AppConfigurationV1) GetCollectionWithContext(ctx context.Context, getCollectionOptions *GetCollectionOptions) (result *GetCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCollectionOptions, "getCollectionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCollectionOptions, "getCollectionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"collection_id": *getCollectionOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appConfiguration.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appConfiguration.Service.Options.URL, `/collections/{collection_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCollectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders(DefaultServiceName, "V1", "GetCollection")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appConfiguration.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetCollection)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteCollection : Delete Collection
// Delete the collection.
func (appConfiguration *AppConfigurationV1) DeleteCollection(deleteCollectionOptions *DeleteCollectionOptions) (response *core.DetailedResponse, err error) {
	return appConfiguration.DeleteCollectionWithContext(context.Background(), deleteCollectionOptions)
}

// DeleteCollectionWithContext is an alternate form of the DeleteCollection method which supports a Context parameter
func (appConfiguration *AppConfigurationV1) DeleteCollectionWithContext(ctx context.Context, deleteCollectionOptions *DeleteCollectionOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteCollectionOptions, "deleteCollectionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteCollectionOptions, "deleteCollectionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"collection_id": *deleteCollectionOptions.CollectionID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appConfiguration.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appConfiguration.Service.Options.URL, `/collections/{collection_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteCollectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders(DefaultServiceName, "V1", "DeleteCollection")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appConfiguration.Service.Request(request, nil)

	return
}

// Collection : Collection struct
type Collection struct {
	// Collection id.
	CollectionID *string `json:"collection_id" validate:"required"`

	// Feature enabled status for the collection.
	Enabled *bool `json:"enabled" validate:"required"`
}

// NewCollection : Instantiate Collection (Generic Model Constructor)
func (*AppConfigurationV1) NewCollection(collectionID string, enabled bool) (model *Collection, err error) {
	model = &Collection{
		CollectionID: core.StringPtr(collectionID),
		Enabled:      core.BoolPtr(enabled),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalCollection unmarshals an instance of Collection from the specified map of raw messages.
func UnmarshalCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Collection)
	err = core.UnmarshalPrimitive(m, "collection_id", &obj.CollectionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CollectionIncludedInFeature : CollectionInFeatureCreate struct
type CollectionIncludedInFeature struct {
	// Collection id.
	CollectionID *string `json:"collection_id" validate:"required"`

	Name *string `json:"name" validate:"required"`

	// Feature enabled status for the collection.
	Enabled *bool `json:"enabled" validate:"required"`
}

// UnmarshalCollectionIncludedInFeature unmarshals an instance of CollectionIncludedInFeature from the specified map of raw messages.
func UnmarshalCollectionIncludedInFeature(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CollectionIncludedInFeature)
	err = core.UnmarshalPrimitive(m, "collection_id", &obj.CollectionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CollectionWithDeletedFlag : CollectionWithDeletedFlag struct
type CollectionWithDeletedFlag struct {
	// Collection id.
	CollectionID *string `json:"collection_id" validate:"required"`

	// Feature enabled status for the collection.
	Enabled *bool `json:"enabled,omitempty"`

	// Don't document this.
	Deleted *bool `json:"deleted,omitempty"`
}

// NewCollectionWithDeletedFlag : Instantiate CollectionWithDeletedFlag (Generic Model Constructor)
func (*AppConfigurationV1) NewCollectionWithDeletedFlag(collectionID string, enabled bool, deleted bool) (model *CollectionWithDeletedFlag, err error) {
	model = &CollectionWithDeletedFlag{
		CollectionID: core.StringPtr(collectionID),
		Enabled:      core.BoolPtr(enabled),
		Deleted:      core.BoolPtr(deleted),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalCollectionWithDeletedFlag unmarshals an instance of CollectionWithDeletedFlag from the specified map of raw messages.
func UnmarshalCollectionWithDeletedFlag(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CollectionWithDeletedFlag)
	err = core.UnmarshalPrimitive(m, "collection_id", &obj.CollectionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deleted", &obj.Deleted)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateCollection : CreateCollection struct
type CreateCollection struct {
	// Collection name.
	Name *string `json:"name" validate:"required"`

	// Guid to which the Collection belongs to.
	Guid *string `json:"guid,omitempty"`

	// Collection Id.
	CollectionID *string `json:"collection_id" validate:"required"`

	// Collection description.
	Description *string `json:"description" validate:"required"`

	// Collection created time.
	CreatedTime *string `json:"created_time" validate:"required"`

	// Collection last updated time.
	UpdatedTime *string `json:"updated_time" validate:"required"`
}

// UnmarshalCreateCollection unmarshals an instance of CreateCollection from the specified map of raw messages.
func UnmarshalCreateCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateCollection)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "guid", &obj.Guid)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "collection_id", &obj.CollectionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_time", &obj.CreatedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_time", &obj.UpdatedTime)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateCollectionOptions : The CreateCollection options.
type CreateCollectionOptions struct {
	// Collection name.
	Name *string `json:"name" validate:"required"`

	// Collection Id.
	CollectionID *string `json:"collection_id," validate:"required"`

	// Description of the collection.
	Description *string `json:"description,omitempty"`

	// Tags associated with the collection.
	Tags *string `json:"tags,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateCollectionOptions : Instantiate CreateCollectionOptions
func (*AppConfigurationV1) NewCreateCollectionOptions(name string) *CreateCollectionOptions {
	return &CreateCollectionOptions{
		Name: core.StringPtr(name),
	}
}

// SetName : Allow user to set Name
func (options *CreateCollectionOptions) SetName(name string) *CreateCollectionOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetCollectionID : Allow user to set CollectionID
func (options *CreateCollectionOptions) SetCollectionID(collectionID string) *CreateCollectionOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateCollectionOptions) SetDescription(description string) *CreateCollectionOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetTags : Allow user to set Tags
func (options *CreateCollectionOptions) SetTags(tags string) *CreateCollectionOptions {
	options.Tags = core.StringPtr(tags)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateCollectionOptions) SetHeaders(param map[string]string) *CreateCollectionOptions {
	options.Headers = param
	return options
}

// DeleteCollectionOptions : The DeleteCollection options.
type DeleteCollectionOptions struct {
	// Collection Id of the collection.
	CollectionID *string `json:"collection_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteCollectionOptions : Instantiate DeleteCollectionOptions
func (*AppConfigurationV1) NewDeleteCollectionOptions(collectionID string) *DeleteCollectionOptions {
	return &DeleteCollectionOptions{
		CollectionID: core.StringPtr(collectionID),
	}
}

// SetCollectionID : Allow user to set CollectionID
func (options *DeleteCollectionOptions) SetCollectionID(collectionID string) *DeleteCollectionOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteCollectionOptions) SetHeaders(param map[string]string) *DeleteCollectionOptions {
	options.Headers = param
	return options
}

// GetSingleCollection : GetSingleCollection struct
type GetSingleCollection struct {
	// Collection name.
	Name *string `json:"name" validate:"required"`

	// Collection id.
	CollectionID *string `json:"collection_id" validate:"required"`

	// Collection description.
	Description *string `json:"description,omitempty"`
}

// UnmarshalGetSingleCollection unmarshals an instance of GetSingleCollection from the specified map of raw messages.
func UnmarshalGetSingleCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetSingleCollection)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "collection_id", &obj.CollectionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateCollection : UpdateCollection struct
type UpdateCollection struct {
	// Collection name.
	Name *string `json:"name" validate:"required"`

	// Collection Id.
	CollectionID *string `json:"collection_id" validate:"required"`

	// Collection description.
	Description *string `json:"description" validate:"required"`

	// Tags associated with the collection.
	CreatedTime *string `json:"created_time" validate:"required"`

	// Collection updated time.
	UpdatedTime *string `json:"updated_time" validate:"required"`
}

// UnmarshalUpdateCollection unmarshals an instance of UpdateCollection from the specified map of raw messages.
func UnmarshalUpdateCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UpdateCollection)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "collection_id", &obj.CollectionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_time", &obj.CreatedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_time", &obj.UpdatedTime)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateCollectionOptions : The UpdateCollection options.
type UpdateCollectionOptions struct {
	// Collection Id of the collection.
	CollectionID *string `json:"collection_id" validate:"required,ne="`

	// Collection name.
	Name *string `json:"name" validate:"required"`

	// Collection description.
	Description *string `json:"description" validate:"required"`

	// Tags associated with the collection.
	Tags *string `json:"tags" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateCollectionOptions : Instantiate UpdateCollectionOptions
func (*AppConfigurationV1) NewUpdateCollectionOptions(collectionID string, name string, description string, tags string) *UpdateCollectionOptions {
	return &UpdateCollectionOptions{
		CollectionID: core.StringPtr(collectionID),
		Name:         core.StringPtr(name),
		Description:  core.StringPtr(description),
		Tags:         core.StringPtr(tags),
	}
}

// SetCollectionID : Allow user to set CollectionID
func (options *UpdateCollectionOptions) SetCollectionID(collectionID string) *UpdateCollectionOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateCollectionOptions) SetName(name string) *UpdateCollectionOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *UpdateCollectionOptions) SetDescription(description string) *UpdateCollectionOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetTags : Allow user to set Tags
func (options *UpdateCollectionOptions) SetTags(tags string) *UpdateCollectionOptions {
	options.Tags = core.StringPtr(tags)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateCollectionOptions) SetHeaders(param map[string]string) *UpdateCollectionOptions {
	options.Headers = param
	return options
}

// CreateFeature : Create Feature
// Create a feature flag.
func (appConfiguration *AppConfigurationV1) CreateFeature(createFeatureOptions *CreateFeatureOptions) (result *CreateFeature, response *core.DetailedResponse, err error) {
	return appConfiguration.CreateFeatureWithContext(context.Background(), createFeatureOptions)
}

// CreateFeatureWithContext is an alternate form of the CreateFeature method which supports a Context parameter
func (appConfiguration *AppConfigurationV1) CreateFeatureWithContext(ctx context.Context, createFeatureOptions *CreateFeatureOptions) (result *CreateFeature, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createFeatureOptions, "createFeatureOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createFeatureOptions, "createFeatureOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appConfiguration.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appConfiguration.Service.Options.URL, `/features`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createFeatureOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders(DefaultServiceName, "V1", "CreateFeature")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createFeatureOptions.Name != nil {
		body["name"] = createFeatureOptions.Name
	}
	if createFeatureOptions.Description != nil {
		body["description"] = createFeatureOptions.Description
	}
	if createFeatureOptions.Type != nil {
		body["type"] = createFeatureOptions.Type
	}
	if createFeatureOptions.EnabledValue != nil {
		if *createFeatureOptions.Type == "BOOLEAN" {
			enabledValue, _ := strconv.ParseBool(*createFeatureOptions.EnabledValue)
			body["enabled_value"] = enabledValue
		} else if *createFeatureOptions.Type == "NUMERIC" {
			enabledValue, _ := strconv.ParseInt(*createFeatureOptions.EnabledValue, 10, 64)
			body["enabled_value"] = enabledValue
		} else {
			body["enabled_value"] = createFeatureOptions.EnabledValue
		}
	}
	if createFeatureOptions.DisabledValue != nil {
		if *createFeatureOptions.Type == "BOOLEAN" {
			disabledValue, _ := strconv.ParseBool(*createFeatureOptions.DisabledValue)
			body["disabled_value"] = disabledValue
		} else if *createFeatureOptions.Type == "NUMERIC" {
			disabledValue, _ := strconv.ParseInt(*createFeatureOptions.DisabledValue, 10, 64)
			body["disabled_value"] = disabledValue
		} else {
			body["disabled_value"] = createFeatureOptions.DisabledValue
		}
	}
	if createFeatureOptions.Tags != nil {
		body["tags"] = createFeatureOptions.Tags
	}
	if createFeatureOptions.SegmentRules != nil {
		body["segment_rules"] = createFeatureOptions.SegmentRules
	}
	if createFeatureOptions.Collections != nil {
		body["collections"] = createFeatureOptions.Collections
	}
	if createFeatureOptions.CreatedMode != nil {
		body["created_mode"] = createFeatureOptions.CreatedMode
	}
	if createFeatureOptions.FeatureID != nil {
		body["feature_id"] = createFeatureOptions.FeatureID
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appConfiguration.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateFeature)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetFeatures : Get list of features
// Get list of features.
func (appConfiguration *AppConfigurationV1) GetFeatures(getFeaturesOptions *GetFeaturesOptions) (result *GetFeatures, response *core.DetailedResponse, err error) {
	return appConfiguration.GetFeaturesWithContext(context.Background(), getFeaturesOptions)
}

// GetFeaturesWithContext is an alternate form of the GetFeatures method which supports a Context parameter
func (appConfiguration *AppConfigurationV1) GetFeaturesWithContext(ctx context.Context, getFeaturesOptions *GetFeaturesOptions) (result *GetFeatures, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getFeaturesOptions, "getFeaturesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appConfiguration.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appConfiguration.Service.Options.URL, `/features`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getFeaturesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders(DefaultServiceName, "V1", "GetFeatures")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getFeaturesOptions.Size != nil {
		builder.AddQuery("size", fmt.Sprint(*getFeaturesOptions.Size))
	}
	if getFeaturesOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*getFeaturesOptions.Offset))
	}
	if getFeaturesOptions.Tags != nil {
		builder.AddQuery("tags", fmt.Sprint(*getFeaturesOptions.Tags))
	}
	if getFeaturesOptions.Collections != nil {
		builder.AddQuery("collections", fmt.Sprint(*getFeaturesOptions.Collections))
	}
	if getFeaturesOptions.Segments != nil {
		builder.AddQuery("segments", fmt.Sprint(*getFeaturesOptions.Segments))
	}
	if getFeaturesOptions.Expand != nil {
		builder.AddQuery("expand", fmt.Sprint(*getFeaturesOptions.Expand))
	}
	if getFeaturesOptions.Include != nil {
		builder.AddQuery("include", fmt.Sprint(*getFeaturesOptions.Include))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appConfiguration.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetFeatures)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateFeature : Update feature properties
// Update a feature flag properties.
func (appConfiguration *AppConfigurationV1) UpdateFeature(updateFeatureOptions *UpdateFeatureOptions) (result *UpdateFeature, response *core.DetailedResponse, err error) {
	return appConfiguration.UpdatefeaturepropertiesWithContext(context.Background(), updateFeatureOptions)
}

// UpdatefeaturepropertiesWithContext is an alternate form of the UpdateFeature method which supports a Context parameter
func (appConfiguration *AppConfigurationV1) UpdatefeaturepropertiesWithContext(ctx context.Context, updateFeatureOptions *UpdateFeatureOptions) (result *UpdateFeature, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateFeatureOptions, "updateFeatureOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateFeatureOptions, "updateFeatureOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"feature_id": *updateFeatureOptions.FeatureID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appConfiguration.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appConfiguration.Service.Options.URL, `/features/{feature_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateFeatureOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders(DefaultServiceName, "V1", "UpdateFeature")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateFeatureOptions.Name != nil {
		body["name"] = updateFeatureOptions.Name
	}
	if updateFeatureOptions.Description != nil {
		body["description"] = updateFeatureOptions.Description
	}
	if updateFeatureOptions.Type != nil {
		body["type"] = updateFeatureOptions.Type
	}
	if updateFeatureOptions.EnabledValue != nil {
		if *updateFeatureOptions.Type == "BOOLEAN" {
			enabledValue, _ := strconv.ParseBool(*updateFeatureOptions.EnabledValue)
			body["enabled_value"] = enabledValue
		} else if *updateFeatureOptions.Type == "NUMERIC" {
			enabledValue, _ := strconv.ParseInt(*updateFeatureOptions.EnabledValue, 10, 64)
			body["enabled_value"] = enabledValue
		} else {
			body["enabled_value"] = updateFeatureOptions.EnabledValue
		}
	}
	if updateFeatureOptions.DisabledValue != nil {
		if *updateFeatureOptions.Type == "BOOLEAN" {
			disabledValue, _ := strconv.ParseBool(*updateFeatureOptions.DisabledValue)
			body["disabled_value"] = disabledValue
		} else if *updateFeatureOptions.Type == "NUMERIC" {
			disabledValue, _ := strconv.ParseInt(*updateFeatureOptions.DisabledValue, 10, 64)
			body["disabled_value"] = disabledValue
		} else {
			body["disabled_value"] = updateFeatureOptions.DisabledValue
		}
	}
	if updateFeatureOptions.Tags != nil {
		body["tags"] = updateFeatureOptions.Tags
	}
	if updateFeatureOptions.SegmentRules != nil {
		body["segment_rules"] = updateFeatureOptions.SegmentRules
	}
	if updateFeatureOptions.Collections != nil {
		body["collections"] = updateFeatureOptions.Collections
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appConfiguration.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalUpdateFeature)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteFeature : Delete a feature
// Delete a feature flag.
func (appConfiguration *AppConfigurationV1) DeleteFeature(deleteFeatureOptions *DeleteafeatureOptions) (response *core.DetailedResponse, err error) {
	return appConfiguration.DeleteafeatureWithContext(context.Background(), deleteFeatureOptions)
}

// DeleteafeatureWithContext is an alternate form of the DeleteFeature method which supports a Context parameter
func (appConfiguration *AppConfigurationV1) DeleteafeatureWithContext(ctx context.Context, deleteFeatureOptions *DeleteafeatureOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteFeatureOptions, "deleteFeatureOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteFeatureOptions, "deleteFeatureOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"feature_id": *deleteFeatureOptions.FeatureID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appConfiguration.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appConfiguration.Service.Options.URL, `/features/{feature_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteFeatureOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders(DefaultServiceName, "V1", "DeleteFeature")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appConfiguration.Service.Request(request, nil)

	return
}

// GetFeature : Get feature details
// Retrieve details of a feature.
func (appConfiguration *AppConfigurationV1) GetFeature(getFeatureOptions *GetFeatureOptions) (result *GetFeature, response *core.DetailedResponse, err error) {
	return appConfiguration.GetFeatureWithContext(context.Background(), getFeatureOptions)
}

// GetFeatureWithContext is an alternate form of the GetFeature method which supports a Context parameter
func (appConfiguration *AppConfigurationV1) GetFeatureWithContext(ctx context.Context, getFeatureOptions *GetFeatureOptions) (result *GetFeature, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getFeatureOptions, "getFeatureOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getFeatureOptions, "getFeatureOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"feature_id": *getFeatureOptions.FeatureID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appConfiguration.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appConfiguration.Service.Options.URL, `/features/{feature_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getFeatureOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders(DefaultServiceName, "V1", "GetFeature")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getFeatureOptions.Include != nil {
		builder.AddQuery("include", fmt.Sprint(*getFeatureOptions.Include))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appConfiguration.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetFeature)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreateFeature : CreateFeature struct
type CreateFeature struct {
	// Feature name.
	Name *string `json:"name" validate:"required"`

	// Feature id.
	FeatureID *string `json:"feature_id" validate:"required"`

	// Feature description.
	Description *string `json:"description" validate:"required"`

	// Type of the feature (Boolean, String, Number).
	Type *string `json:"type" validate:"required"`

	// Value of the feature when it is enabled.
	EnabledValue interface{} `json:"enabled_value" validate:"required"`

	// Value of the feature when it is disabled.
	DisabledValue interface{} `json:"disabled_value" validate:"required"`

	// Segment Rules array.
	SegmentRules []SegmentRule `json:"segment_rules" validate:"required"`

	// Collection array.
	Collections []Collection `json:"collections" validate:"required"`

	// Feature created time.
	CreatedTime *string `json:"created_time" validate:"required"`

	// Feature updated time.
	UpdatedTime *string `json:"updated_time" validate:"required"`
}

// UnmarshalCreateFeature unmarshals an instance of CreateFeature from the specified map of raw messages.
func UnmarshalCreateFeature(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateFeature)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "feature_id", &obj.FeatureID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled_value", &obj.EnabledValue)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "disabled_value", &obj.DisabledValue)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "segment_rules", &obj.SegmentRules, UnmarshalSegmentRule)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "collections", &obj.Collections, UnmarshalCollection)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_time", &obj.CreatedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_time", &obj.UpdatedTime)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateFeatureOptions : The CreateFeature options.
type CreateFeatureOptions struct {
	// Feature name.
	Name *string `json:"name" validate:"required"`

	// Feature id.
	FeatureID *string `json:"feature_id validate:"required"`

	// Feature description.
	Description *string `json:"description"`

	// Type of the feature (Boolean, String, Number).
	Type *string `json:"type" validate:"required"`

	// Value of the feature when it is enabled.
	EnabledValue *string `json:"enabled_value" validate:"required"`

	// Value of the feature when it is disabled.
	DisabledValue *string `json:"disabled_value" validate:"required"`

	// Tags associated with the feature.
	Tags *string `json:"tags"`

	// Segment Rules array.
	SegmentRules []SegmentRule `json:"segment_rules" validate:"required"`

	// Collection array.
	Collections []Collection `json:"collections" validate:"required"`

	// Internal.  Dont document.
	CreatedMode *string `json:"created_mode"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateFeatureOptions : Instantiate CreateFeatureOptions
func (*AppConfigurationV1) NewCreateFeatureOptions(name string, featureID string, description string, typeVar string, enabledValue string, disabledValue string, tags string, segmentRules []SegmentRule, collections []Collection, createdMode string) *CreateFeatureOptions {
	return &CreateFeatureOptions{
		Name:          core.StringPtr(name),
		FeatureID:     core.StringPtr(featureID),
		Description:   core.StringPtr(description),
		Type:          core.StringPtr(typeVar),
		EnabledValue:  core.StringPtr(enabledValue),
		DisabledValue: core.StringPtr(disabledValue),
		Tags:          core.StringPtr(tags),
		SegmentRules:  segmentRules,
		Collections:   collections,
		CreatedMode:   core.StringPtr(createdMode),
	}
}

// SetName : Allow user to set Name
func (options *CreateFeatureOptions) SetName(name string) *CreateFeatureOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateFeatureOptions) SetDescription(description string) *CreateFeatureOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetType : Allow user to set Type
func (options *CreateFeatureOptions) SetType(typeVar string) *CreateFeatureOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetEnabledValue : Allow user to set EnabledValue
func (options *CreateFeatureOptions) SetEnabledValue(enabledValue string) *CreateFeatureOptions {
	options.EnabledValue = core.StringPtr(enabledValue)
	return options
}

// SetDisabledValue : Allow user to set DisabledValue
func (options *CreateFeatureOptions) SetDisabledValue(disabledValue string) *CreateFeatureOptions {
	options.DisabledValue = core.StringPtr(disabledValue)
	return options
}

// SetTags : Allow user to set Tags
func (options *CreateFeatureOptions) SetTags(tags string) *CreateFeatureOptions {
	options.Tags = core.StringPtr(tags)
	return options
}

// SetFeatureID : Allow user to set FeatureID
func (options *CreateFeatureOptions) SetFeatureID(featureID string) *CreateFeatureOptions {
	options.FeatureID = core.StringPtr(featureID)
	return options
}

// SetSegmentRules : Allow user to set SegmentRules
func (options *CreateFeatureOptions) SetSegmentRules(segmentRules []SegmentRule) *CreateFeatureOptions {
	options.SegmentRules = segmentRules
	return options
}

// SetCollections : Allow user to set Collections
func (options *CreateFeatureOptions) SetCollections(collections []Collection) *CreateFeatureOptions {
	options.Collections = collections
	return options
}

// SetCreatedMode : Allow user to set CreatedMode
func (options *CreateFeatureOptions) SetCreatedMode(createdMode string) *CreateFeatureOptions {
	options.CreatedMode = core.StringPtr(createdMode)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateFeatureOptions) SetHeaders(param map[string]string) *CreateFeatureOptions {
	options.Headers = param
	return options
}

// DeleteafeatureOptions : The DeleteFeature options.
type DeleteafeatureOptions struct {
	// Feature Id.
	FeatureID *string `json:"feature_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteafeatureOptions : Instantiate DeleteafeatureOptions
func (*AppConfigurationV1) NewDeleteafeatureOptions(featureID string) *DeleteafeatureOptions {
	return &DeleteafeatureOptions{
		FeatureID: core.StringPtr(featureID),
	}
}

// SetFeatureID : Allow user to set FeatureID
func (options *DeleteafeatureOptions) SetFeatureID(featureID string) *DeleteafeatureOptions {
	options.FeatureID = core.StringPtr(featureID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteafeatureOptions) SetHeaders(param map[string]string) *DeleteafeatureOptions {
	options.Headers = param
	return options
}

// Feature : Feature struct
type Feature struct {
	// Feature id.
	FeatureID *string `json:"feature_id" validate:"required"`

	// Feature name.
	Name *string `json:"name" validate:"required"`

	// Feature enabled status for the collection.
	Enabled *bool `json:"enabled" validate:"required"`
}

// UnmarshalFeature unmarshals an instance of Feature from the specified map of raw messages.
func UnmarshalFeature(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Feature)
	err = core.UnmarshalPrimitive(m, "feature_id", &obj.FeatureID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Feature1 : Feature1 struct
type Feature1 struct {
	// Feature id.
	FeatureID *string `json:"feature_id" validate:"required"`

	// Feature name.
	Name *string `json:"name" validate:"required"`
}

// UnmarshalFeature1 unmarshals an instance of Feature1 from the specified map of raw messages.
func UnmarshalFeature1(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Feature1)
	err = core.UnmarshalPrimitive(m, "feature_id", &obj.FeatureID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetCollection : GetCollection struct
type GetCollection struct {
	// Collection name.
	Name *string `json:"name" validate:"required"`

	// Collection id.
	CollectionID *string `json:"collection_id" validate:"required"`
}

// UnmarshalGetCollection unmarshals an instance of GetCollection from the specified map of raw messages.
func UnmarshalGetCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetCollection)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "collection_id", &obj.CollectionID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetCollectionOptions : The GetCollection options.
type GetCollectionOptions struct {
	// Collection Id for the collection.
	CollectionID *string `json:"collection_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCollectionOptions : Instantiate GetCollectionOptions
func (*AppConfigurationV1) NewGetCollectionOptions(collectionID string) *GetCollectionOptions {
	return &GetCollectionOptions{
		CollectionID: core.StringPtr(collectionID),
	}
}

// SetCollectionID : Allow user to set CollectionID
func (options *GetCollectionOptions) SetCollectionID(collectionID string) *GetCollectionOptions {
	options.CollectionID = core.StringPtr(collectionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetCollectionOptions) SetHeaders(param map[string]string) *GetCollectionOptions {
	options.Headers = param
	return options
}

// GetCollectionsOptions : The GetCollections options.
type GetCollectionsOptions struct {
	// Optional.  Used for pagination.  Size of the number of records retrieved.
	Size *string `json:"size,omitempty"`

	// Optional.  Used for pagination.  Offset used to retrieve records.
	Offset *string `json:"offset,omitempty"`

	// Optional. Filter based on the feature's shortname.
	Features *string `json:"features,omitempty"`

	// Optional.  Filter based on the tags.
	Tags *string `json:"tags,omitempty"`

	// Optional.  Expanded view of the collection details.
	Expand *string `json:"expand,omitempty"`

	// Optional.  Include feature details in the response.
	Include *string `json:"include,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCollectionsOptions : Instantiate GetCollectionsOptions
func (*AppConfigurationV1) NewGetCollectionsOptions() *GetCollectionsOptions {
	return &GetCollectionsOptions{}
}

// SetSize : Allow user to set Size
func (options *GetCollectionsOptions) SetSize(size string) *GetCollectionsOptions {
	options.Size = core.StringPtr(size)
	return options
}

// SetOffset : Allow user to set Offset
func (options *GetCollectionsOptions) SetOffset(offset string) *GetCollectionsOptions {
	options.Offset = core.StringPtr(offset)
	return options
}

// SetFeatures : Allow user to set Features
func (options *GetCollectionsOptions) SetFeatures(features string) *GetCollectionsOptions {
	options.Features = core.StringPtr(features)
	return options
}

// SetTags : Allow user to set Tags
func (options *GetCollectionsOptions) SetTags(tags string) *GetCollectionsOptions {
	options.Tags = core.StringPtr(tags)
	return options
}

// SetExpand : Allow user to set Expand
func (options *GetCollectionsOptions) SetExpand(expand string) *GetCollectionsOptions {
	options.Expand = core.StringPtr(expand)
	return options
}

// SetInclude : Allow user to set Include
func (options *GetCollectionsOptions) SetInclude(include string) *GetCollectionsOptions {
	options.Include = core.StringPtr(include)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetCollectionsOptions) SetHeaders(param map[string]string) *GetCollectionsOptions {
	options.Headers = param
	return options
}

// GetFeatures : GetFeatures struct
type GetFeatures struct {
	// Feature array.
	Features []SingleFeature `json:"features" validate:"required"`

	PageInfo *PageInfo `json:"page_info" validate:"required"`
}

// UnmarshalGetFeatures unmarshals an instance of GetFeatures from the specified map of raw messages.
func UnmarshalGetFeatures(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetFeatures)
	err = core.UnmarshalModel(m, "features", &obj.Features, UnmarshalSingleFeature)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "page_info", &obj.PageInfo, UnmarshalPageInfo)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetFeature : GetFeature struct
type GetFeature struct {
	// Feature name.
	Name *string `json:"name" validate:"required"`

	// Feature id.
	FeatureID *string `json:"feature_id" validate:"required"`

	// Feature description.
	Description *string `json:"description" validate:"required"`

	// Type of the feature (Boolean, String, Number).
	Type *string `json:"type" validate:"required"`

	// Value of the feature when it is enabled.
	EnabledValue interface{} `json:"enabled_value" validate:"required"`

	// Value of the feature when it is disabled.
	DisabledValue interface{} `json:"disabled_value" validate:"required"`

	// Segment rule array.
	SegmentRules []SegmentRule `json:"segment_rules" validate:"required"`

	// Feature created time.
	CreatedTime *string `json:"created_time" validate:"required"`

	// Feature updated time.
	UpdatedTime *string `json:"updated_time" validate:"required"`

	Collections []CollectionIncludedInFeature `json:"collections"`
}

// UnmarshalGetFeature unmarshals an instance of GetFeature from the specified map of raw messages.
func UnmarshalGetFeature(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetFeature)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "feature_id", &obj.FeatureID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled_value", &obj.EnabledValue)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "disabled_value", &obj.DisabledValue)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "segment_rules", &obj.SegmentRules, UnmarshalSegmentRule)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_time", &obj.CreatedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_time", &obj.UpdatedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "collections", &obj.Collections, UnmarshalCollectionIncludedInFeature)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetFeatureOptions : The GetFeature options.
type GetFeatureOptions struct {
	// Feature Id.
	FeatureID *string `json:"feature_id" validate:"required,ne="`

	// Optional.  Feature details to include the associated collections details in the response.
	Include *string `json:"include,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetFeatureOptions : Instantiate GetFeatureOptions
func (*AppConfigurationV1) NewGetFeatureOptions(featureID string) *GetFeatureOptions {
	return &GetFeatureOptions{
		FeatureID: core.StringPtr(featureID),
	}
}

// SetFeatureID : Allow user to set FeatureID
func (options *GetFeatureOptions) SetFeatureID(featureID string) *GetFeatureOptions {
	options.FeatureID = core.StringPtr(featureID)
	return options
}

// SetInclude : Allow user to set Include
func (options *GetFeatureOptions) SetInclude(include string) *GetFeatureOptions {
	options.Include = core.StringPtr(include)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetFeatureOptions) SetHeaders(param map[string]string) *GetFeatureOptions {
	options.Headers = param
	return options
}

// GetCollections : GetCollections struct
type GetCollections struct {
	// Collection array.
	Collections []GetSingleCollection `json:"collections" validate:"required"`

	PageInfo *PageInfo `json:"page_info" validate:"required"`
}

// UnmarshalGetCollections unmarshals an instance of GetCollections from the specified map of raw messages.
func UnmarshalGetCollections(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetCollections)
	err = core.UnmarshalModel(m, "collections", &obj.Collections, UnmarshalGetSingleCollection)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "page_info", &obj.PageInfo, UnmarshalPageInfo)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetFeaturesOptions : The GetFeatures options.
type GetFeaturesOptions struct {
	// Optional.  Used for pagination.  Size of the number of records retrieved.
	Size *string `json:"size,omitempty"`

	// Optional.  Used for pagination.  Offset used to retrieve records.
	Offset *string `json:"offset,omitempty"`

	// Optional.  Filter features by a list of comma separated tags.
	Tags *string `json:"tags,omitempty"`

	// Optional.  Filter features by a list of comma separated collections.
	Collections *string `json:"collections,omitempty"`

	// Optional.  Filter features by a list of comma separated segment Id's.
	Segments *string `json:"segments,omitempty"`

	// Optional.  Expanded view the feature details.
	Expand *string `json:"expand,omitempty"`

	// Optional.  Feature details to include the associated collections or rules details in the response.
	Include *string `json:"include,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetFeaturesOptions : Instantiate GetFeaturesOptions
func (*AppConfigurationV1) NewGetFeaturesOptions() *GetFeaturesOptions {
	return &GetFeaturesOptions{}
}

// SetSize : Allow user to set Size
func (options *GetFeaturesOptions) SetSize(size string) *GetFeaturesOptions {
	options.Size = core.StringPtr(size)
	return options
}

// SetOffset : Allow user to set Offset
func (options *GetFeaturesOptions) SetOffset(offset string) *GetFeaturesOptions {
	options.Offset = core.StringPtr(offset)
	return options
}

// SetTags : Allow user to set Tags
func (options *GetFeaturesOptions) SetTags(tags string) *GetFeaturesOptions {
	options.Tags = core.StringPtr(tags)
	return options
}

// SetCollections : Allow user to set Collections
func (options *GetFeaturesOptions) SetCollections(collections string) *GetFeaturesOptions {
	options.Collections = core.StringPtr(collections)
	return options
}

// SetSegments : Allow user to set Segments
func (options *GetFeaturesOptions) SetSegments(segments string) *GetFeaturesOptions {
	options.Segments = core.StringPtr(segments)
	return options
}

// SetExpand : Allow user to set Expand
func (options *GetFeaturesOptions) SetExpand(expand string) *GetFeaturesOptions {
	options.Expand = core.StringPtr(expand)
	return options
}

// SetInclude : Allow user to set Include
func (options *GetFeaturesOptions) SetInclude(include string) *GetFeaturesOptions {
	options.Include = core.StringPtr(include)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetFeaturesOptions) SetHeaders(param map[string]string) *GetFeaturesOptions {
	options.Headers = param
	return options
}

// UpdateFeature : UpdateFeature struct
type UpdateFeature struct {
	// Feature name.
	Name *string `json:"name" validate:"required"`

	// Feature description.
	Description *string `json:"description" validate:"required"`

	// Type of the feature (Boolean, String, Number).
	Type *string `json:"type" validate:"required"`

	// Value of the feature when it is enabled.
	EnabledValue interface{} `json:"enabled_value" validate:"required"`

	// Value of the feature when it is disabled.
	DisabledValue interface{} `json:"disabled_value" validate:"required"`

	// Segment Rule array.
	SegmentRules []SegmentRule `json:"segment_rules" validate:"required"`

	// Collection array.
	Collections []Collection `json:"collections" validate:"required"`

	// Feature created time.
	CreatedTime *string `json:"created_time" validate:"required"`

	// Feature updated time.
	UpdatedTime *string `json:"updated_time" validate:"required"`
}

// UnmarshalUpdateFeature unmarshals an instance of UpdateFeature from the specified map of raw messages.
func UnmarshalUpdateFeature(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UpdateFeature)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled_value", &obj.EnabledValue)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "disabled_value", &obj.DisabledValue)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "segment_rules", &obj.SegmentRules, UnmarshalSegmentRule)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "collections", &obj.Collections, UnmarshalCollection)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_time", &obj.CreatedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_time", &obj.UpdatedTime)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateFeatureOptions : The UpdateFeature options.
type UpdateFeatureOptions struct {
	// Feature Id.
	FeatureID *string `json:"feature_id" validate:"required,ne="`

	// Feature name.
	Name *string `json:"name" validate:"required"`

	// Feature description.
	Description *string `json:"description" validate:"required"`

	// Type of the feature (Boolean, String, Number).
	Type *string `json:"type" validate:"required"`

	// Value of the feature when it is enabled.
	EnabledValue *string `json:"enabled_value" validate:"required"`

	// Value of the feature when it is disabled.
	DisabledValue *string `json:"disabled_value" validate:"required"`

	// Tags associated with the feature.
	Tags *string `json:"tags" validate:"required"`

	// Segment Rule array.
	SegmentRules []SegmentRule `json:"segment_rules" validate:"required"`

	// Collections array.
	Collections []CollectionWithDeletedFlag `json:"collections" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateFeatureOptions : Instantiate UpdateFeatureOptions
func (*AppConfigurationV1) NewUpdateFeatureOptions(featureID string, name string, description string, typeVar string, enabledValue string, disabledValue string, tags string, segmentRules []SegmentRule, collections []CollectionWithDeletedFlag) *UpdateFeatureOptions {
	return &UpdateFeatureOptions{
		FeatureID:     core.StringPtr(featureID),
		Name:          core.StringPtr(name),
		Description:   core.StringPtr(description),
		Type:          core.StringPtr(typeVar),
		EnabledValue:  core.StringPtr(enabledValue),
		DisabledValue: core.StringPtr(disabledValue),
		Tags:          core.StringPtr(tags),
		SegmentRules:  segmentRules,
		Collections:   collections,
	}
}

// SetFeatureID : Allow user to set FeatureID
func (options *UpdateFeatureOptions) SetFeatureID(featureID string) *UpdateFeatureOptions {
	options.FeatureID = core.StringPtr(featureID)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateFeatureOptions) SetName(name string) *UpdateFeatureOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *UpdateFeatureOptions) SetDescription(description string) *UpdateFeatureOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetType : Allow user to set Type
func (options *UpdateFeatureOptions) SetType(typeVar string) *UpdateFeatureOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetEnabledValue : Allow user to set EnabledValue
func (options *UpdateFeatureOptions) SetEnabledValue(enabledValue string) *UpdateFeatureOptions {
	options.EnabledValue = core.StringPtr(enabledValue)
	return options
}

// SetDisabledValue : Allow user to set DisabledValue
func (options *UpdateFeatureOptions) SetDisabledValue(disabledValue string) *UpdateFeatureOptions {
	options.DisabledValue = core.StringPtr(disabledValue)
	return options
}

// SetTags : Allow user to set Tags
func (options *UpdateFeatureOptions) SetTags(tags string) *UpdateFeatureOptions {
	options.Tags = core.StringPtr(tags)
	return options
}

// SetSegmentRules : Allow user to set SegmentRules
func (options *UpdateFeatureOptions) SetSegmentRules(segmentRules []SegmentRule) *UpdateFeatureOptions {
	options.SegmentRules = segmentRules
	return options
}

// SetCollections : Allow user to set Collections
func (options *UpdateFeatureOptions) SetCollections(collections []CollectionWithDeletedFlag) *UpdateFeatureOptions {
	options.Collections = collections
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateFeatureOptions) SetHeaders(param map[string]string) *UpdateFeatureOptions {
	options.Headers = param
	return options
}

// SingleFeature : SingleFeature struct
type SingleFeature struct {
	// Feature name.
	Name *string `json:"name,omitempty"`

	// Feature id.
	FeatureID *string `json:"feature_id,omitempty"`

	// Feature is associated to any segment or not.
	SegmentExists *bool `json:"segment_exists,omitempty"`

	// Feature description.
	Description *string `json:"description,omitempty"`

	// Type of the feature (Boolean, String, Number).
	Type *string `json:"type,omitempty"`

	// Value of the feature when it is enabled.
	EnabledValue interface{} `json:"enabled_value,omitempty"`

	// Value of the feature when it is disabled.
	DisabledValue interface{} `json:"disabled_value,omitempty"`

	// Feature created time.
	CreatedTime *string `json:"created_time,omitempty"`

	// Feature updated time.
	UpdatedTime *string `json:"updated_time,omitempty"`
}

// UnmarshalSingleFeature unmarshals an instance of SingleFeature from the specified map of raw messages.
func UnmarshalSingleFeature(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SingleFeature)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "feature_id", &obj.FeatureID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "segment_exists", &obj.SegmentExists)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled_value", &obj.EnabledValue)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "disabled_value", &obj.DisabledValue)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_time", &obj.CreatedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_time", &obj.UpdatedTime)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateSegment : Create Segment
// Create a segment of users.
func (appConfiguration *AppConfigurationV1) CreateSegment(createSegmentOptions *CreateSegmentOptions) (result *CreateSegment, response *core.DetailedResponse, err error) {
	return appConfiguration.CreateSegmentWithContext(context.Background(), createSegmentOptions)
}

// CreateSegmentWithContext is an alternate form of the CreateSegment method which supports a Context parameter
func (appConfiguration *AppConfigurationV1) CreateSegmentWithContext(ctx context.Context, createSegmentOptions *CreateSegmentOptions) (result *CreateSegment, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createSegmentOptions, "createSegmentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createSegmentOptions, "createSegmentOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appConfiguration.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appConfiguration.Service.Options.URL, `/segments`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createSegmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders(DefaultServiceName, "V1", "CreateSegment")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createSegmentOptions.Name != nil {
		body["name"] = createSegmentOptions.Name
	}
	if createSegmentOptions.Description != nil {
		body["description"] = createSegmentOptions.Description
	}
	if createSegmentOptions.Tags != nil {
		body["tags"] = createSegmentOptions.Tags
	}
	if createSegmentOptions.Rules != nil {
		body["rules"] = createSegmentOptions.Rules
	}
	if createSegmentOptions.CreatedMode != nil {
		body["created_mode"] = createSegmentOptions.CreatedMode
	}
	if createSegmentOptions.SegmentID != nil {
		body["segment_id"] = createSegmentOptions.SegmentID
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appConfiguration.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateSegment)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetSegments : Get list of segments
// Get list of segments.
func (appConfiguration *AppConfigurationV1) GetSegments(getSegmentsOptions *GetSegmentsOptions) (result *GetSegments, response *core.DetailedResponse, err error) {
	return appConfiguration.GetSegmentsWithContext(context.Background(), getSegmentsOptions)
}

// GetSegmentsWithContext is an alternate form of the GetSegments method which supports a Context parameter
func (appConfiguration *AppConfigurationV1) GetSegmentsWithContext(ctx context.Context, getSegmentsOptions *GetSegmentsOptions) (result *GetSegments, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getSegmentsOptions, "getSegmentsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appConfiguration.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appConfiguration.Service.Options.URL, `/segments`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getSegmentsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders(DefaultServiceName, "V1", "GetSegments")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getSegmentsOptions.Size != nil {
		builder.AddQuery("size", fmt.Sprint(*getSegmentsOptions.Size))
	}
	if getSegmentsOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*getSegmentsOptions.Offset))
	}
	if getSegmentsOptions.Tags != nil {
		builder.AddQuery("tags", fmt.Sprint(*getSegmentsOptions.Tags))
	}
	if getSegmentsOptions.Features != nil {
		builder.AddQuery("features", fmt.Sprint(*getSegmentsOptions.Features))
	}
	if getSegmentsOptions.Expand != nil {
		builder.AddQuery("expand", fmt.Sprint(*getSegmentsOptions.Expand))
	}
	if getSegmentsOptions.Include != nil {
		builder.AddQuery("include", fmt.Sprint(*getSegmentsOptions.Include))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appConfiguration.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetSegments)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateSegment : Update the segment
// Update the segment properties.
func (appConfiguration *AppConfigurationV1) UpdateSegment(updateSegmentOptions *UpdateSegmentOptions) (result *UpdateSegment, response *core.DetailedResponse, err error) {
	return appConfiguration.UpdateSegmentWithContext(context.Background(), updateSegmentOptions)
}

// UpdateSegmentWithContext is an alternate form of the UpdateSegment method which supports a Context parameter
func (appConfiguration *AppConfigurationV1) UpdateSegmentWithContext(ctx context.Context, updateSegmentOptions *UpdateSegmentOptions) (result *UpdateSegment, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateSegmentOptions, "updateSegmentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateSegmentOptions, "updateSegmentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"segment_id": *updateSegmentOptions.SegmentID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appConfiguration.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appConfiguration.Service.Options.URL, `/segments/{segment_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateSegmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders(DefaultServiceName, "V1", "UpdateSegment")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateSegmentOptions.Name != nil {
		body["name"] = updateSegmentOptions.Name
	}
	if updateSegmentOptions.Description != nil {
		body["description"] = updateSegmentOptions.Description
	}
	if updateSegmentOptions.Tags != nil {
		body["tags"] = updateSegmentOptions.Tags
	}
	if updateSegmentOptions.Rules != nil {
		body["rules"] = updateSegmentOptions.Rules
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appConfiguration.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalUpdateSegment)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetSegment : Get segment details
// Retrieve details of a segment.
func (appConfiguration *AppConfigurationV1) GetSegment(getSegmentOptions *GetSegmentOptions) (result *GetSegment, response *core.DetailedResponse, err error) {
	return appConfiguration.GetSegmentWithContext(context.Background(), getSegmentOptions)
}

// GetSegmentWithContext is an alternate form of the GetSegment method which supports a Context parameter
func (appConfiguration *AppConfigurationV1) GetSegmentWithContext(ctx context.Context, getSegmentOptions *GetSegmentOptions) (result *GetSegment, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getSegmentOptions, "getSegmentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getSegmentOptions, "getSegmentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"segment_id": *getSegmentOptions.SegmentID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appConfiguration.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appConfiguration.Service.Options.URL, `/segments/{segment_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getSegmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders(DefaultServiceName, "V1", "GetSegment")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("segment_id", fmt.Sprint(*getSegmentOptions.SegmentID))
	if getSegmentOptions.Include != nil {
		builder.AddQuery("include", fmt.Sprint(*getSegmentOptions.Include))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appConfiguration.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetSegment)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteSegment : Delete a segment
// Delete a segment.
func (appConfiguration *AppConfigurationV1) DeleteSegment(deleteSegmentOptions *DeleteSegmentOptions) (response *core.DetailedResponse, err error) {
	return appConfiguration.DeleteSegmentWithContext(context.Background(), deleteSegmentOptions)
}

// DeleteSegmentWithContext is an alternate form of the DeleteSegment method which supports a Context parameter
func (appConfiguration *AppConfigurationV1) DeleteSegmentWithContext(ctx context.Context, deleteSegmentOptions *DeleteSegmentOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteSegmentOptions, "deleteSegmentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteSegmentOptions, "deleteSegmentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"segment_id": *deleteSegmentOptions.SegmentID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appConfiguration.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appConfiguration.Service.Options.URL, `/segments/{segment_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteSegmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders(DefaultServiceName, "V1", "DeleteSegment")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appConfiguration.Service.Request(request, nil)

	return
}

// CreateSegment : CreateSegment struct
type CreateSegment struct {
	// Segment name.
	Name *string `json:"name" validate:"required"`

	// Segment id.
	SegmentID *string `json:"segment_id" validate:"required"`

	// Segment description.
	Description *string `json:"description" validate:"required"`

	// Segment created time.
	CreatedTime *string `json:"created_time" validate:"required"`

	// Segment updated time.
	UpdatedTime *string `json:"updated_time" validate:"required"`
}

// UnmarshalCreateSegment unmarshals an instance of CreateSegment from the specified map of raw messages.
func UnmarshalCreateSegment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateSegment)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "segment_id", &obj.SegmentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_time", &obj.CreatedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_time", &obj.UpdatedTime)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateSegmentOptions : The CreateSegment options.
type CreateSegmentOptions struct {
	// Segment name.
	Name *string `json:"name" validate:"required"`

	// Segment id.
	SegmentID *string `json:"segment_id" validate:"required"`

	// Segment description.
	Description *string `json:"description"`

	// Tags associated with the segments.
	Tags *string `json:"tags"`

	// Rule array.
	Rules []RuleArray `json:"rules" validate:"required"`

	// Dont document.
	CreatedMode *string `json:"created_mode"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateSegmentOptions : Instantiate CreateSegmentOptions
func (*AppConfigurationV1) NewCreateSegmentOptions(name string, segmentID string, description string, tags string, rules []RuleArray, createdMode string) *CreateSegmentOptions {
	return &CreateSegmentOptions{
		Name:        core.StringPtr(name),
		SegmentID:   core.StringPtr(segmentID),
		Description: core.StringPtr(description),
		Tags:        core.StringPtr(tags),
		Rules:       rules,
		CreatedMode: core.StringPtr(createdMode),
	}
}

// SetName : Allow user to set Name
func (options *CreateSegmentOptions) SetName(name string) *CreateSegmentOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetSegmentID : Allow user to set SegmentID
func (options *CreateSegmentOptions) SetSegmentID(segmentID string) *CreateSegmentOptions {
	options.SegmentID = core.StringPtr(segmentID)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateSegmentOptions) SetDescription(description string) *CreateSegmentOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetTags : Allow user to set Tags
func (options *CreateSegmentOptions) SetTags(tags string) *CreateSegmentOptions {
	options.Tags = core.StringPtr(tags)
	return options
}

// SetRules : Allow user to set Rules
func (options *CreateSegmentOptions) SetRules(rules []RuleArray) *CreateSegmentOptions {
	options.Rules = rules
	return options
}

// SetCreatedMode : Allow user to set CreatedMode
func (options *CreateSegmentOptions) SetCreatedMode(createdMode string) *CreateSegmentOptions {
	options.CreatedMode = core.StringPtr(createdMode)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateSegmentOptions) SetHeaders(param map[string]string) *CreateSegmentOptions {
	options.Headers = param
	return options
}

// DeleteSegmentOptions : The DeleteSegment options.
type DeleteSegmentOptions struct {
	// Segment Id.
	SegmentID *string `json:"segment_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteasegmentOptions : Instantiate DeleteSegmentOptions
func (*AppConfigurationV1) NewDeleteasegmentOptions(segmentID string) *DeleteSegmentOptions {
	return &DeleteSegmentOptions{
		SegmentID: core.StringPtr(segmentID),
	}
}

// SetSegmentID : Allow user to set SegmentID
func (options *DeleteSegmentOptions) SetSegmentID(segmentID string) *DeleteSegmentOptions {
	options.SegmentID = core.StringPtr(segmentID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteSegmentOptions) SetHeaders(param map[string]string) *DeleteSegmentOptions {
	options.Headers = param
	return options
}

// GetAllSegmentSingleSegment : GetAllSegmentSingleSegment struct
type GetAllSegmentSingleSegment struct {
	// Segment name.
	Name *string `json:"name" validate:"required"`

	// Segment id.
	SegmentID *string `json:"segment_id" validate:"required"`

	// Segment description.
	Description *string `json:"description,omitempty"`

	// Feature created time.
	CreatedTime *string `json:"created_time,omitempty"`

	// Feature updated time.
	UpdatedTime *string `json:"updated_time,omitempty"`

	// Rule array.
	Rules []RuleArray `json:"rules,omitempty"`
}

// UnmarshalGetAllSegmentSingleSegment unmarshals an instance of GetAllSegmentSingleSegment from the specified map of raw messages.
func UnmarshalGetAllSegmentSingleSegment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetAllSegmentSingleSegment)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "segment_id", &obj.SegmentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_time", &obj.CreatedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_time", &obj.UpdatedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "rules", &obj.Rules, UnmarshalRuleArray)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetSegments : GetSegments struct
type GetSegments struct {
	// Segment array.
	Segments []GetAllSegmentSingleSegment `json:"segments" validate:"required"`

	PageInfo *PageInfo `json:"page_info" validate:"required"`
}

// UnmarshalGetSegments unmarshals an instance of GetSegments from the specified map of raw messages.
func UnmarshalGetSegments(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetSegments)
	err = core.UnmarshalModel(m, "segments", &obj.Segments, UnmarshalGetAllSegmentSingleSegment)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "page_info", &obj.PageInfo, UnmarshalPageInfo)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetSegmentsOptions : The GetSegments options.
type GetSegmentsOptions struct {
	// Optional.  Used for pagination.  Size of the number of records retrieved.
	Size *string `json:"size,omitempty"`

	// Optional.  Used for pagination.  Offset used to retrieve records.
	Offset *string `json:"offset,omitempty"`

	// Optional.  Filter segments by a list of comma separated tags.
	Tags *string `json:"tags,omitempty"`

	// Optional.  Filter segments by a list of comma separated features.
	Features *string `json:"features,omitempty"`

	// Optional.  Expanded view the segment details.
	Expand *string `json:"expand,omitempty"`

	// Optional.  Segment details to include the associated rules in the response.
	Include *string `json:"include,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetSegmentsOptions : Instantiate GetSegmentsOptions
func (*AppConfigurationV1) NewGetSegmentsOptions() *GetSegmentsOptions {
	return &GetSegmentsOptions{}
}

// SetSize : Allow user to set Size
func (options *GetSegmentsOptions) SetSize(size string) *GetSegmentsOptions {
	options.Size = core.StringPtr(size)
	return options
}

// SetOffset : Allow user to set Offset
func (options *GetSegmentsOptions) SetOffset(offset string) *GetSegmentsOptions {
	options.Offset = core.StringPtr(offset)
	return options
}

// SetTags : Allow user to set Tags
func (options *GetSegmentsOptions) SetTags(tags string) *GetSegmentsOptions {
	options.Tags = core.StringPtr(tags)
	return options
}

// SetFeatures : Allow user to set Features
func (options *GetSegmentsOptions) SetFeatures(features string) *GetSegmentsOptions {
	options.Features = core.StringPtr(features)
	return options
}

// SetExpand : Allow user to set Expand
func (options *GetSegmentsOptions) SetExpand(expand string) *GetSegmentsOptions {
	options.Expand = core.StringPtr(expand)
	return options
}

// SetInclude : Allow user to set Include
func (options *GetSegmentsOptions) SetInclude(include string) *GetSegmentsOptions {
	options.Include = core.StringPtr(include)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetSegmentsOptions) SetHeaders(param map[string]string) *GetSegmentsOptions {
	options.Headers = param
	return options
}

// GetSegment : GetSegment struct
type GetSegment struct {
	// Segment name.
	Name *string `json:"name" validate:"required"`

	// Segment id.
	SegmentID *string `json:"segment_id" validate:"required"`

	// Segment description.
	Description *string `json:"description" validate:"required"`

	// Rule array.
	Rules []RuleArray `json:"rules" validate:"required"`

	// Feature arrary.
	Features []Feature1 `json:"features" validate:"required"`

	// Segment created time.
	CreatedTime *string `json:"created_time" validate:"required"`

	// Segment updated time.
	UpdatedTime *string `json:"updated_time" validate:"required"`
}

// UnmarshalGetSegment unmarshals an instance of GetSegment from the specified map of raw messages.
func UnmarshalGetSegment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetSegment)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "segment_id", &obj.SegmentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "rules", &obj.Rules, UnmarshalRuleArray)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "features", &obj.Features, UnmarshalFeature1)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_time", &obj.CreatedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_time", &obj.UpdatedTime)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetSegmentOptions : The GetSegment options.
type GetSegmentOptions struct {
	// Segment Id.
	SegmentID *string `json:"segment_id" validate:"required,ne="`

	// Optional.  Instructs to include the feature details based on the segments association.
	Include *string `json:"include,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetSegmentOptions : Instantiate GetSegmentOptions
func (*AppConfigurationV1) NewGetSegmentOptions(segmentID string) *GetSegmentOptions {
	return &GetSegmentOptions{
		SegmentID: core.StringPtr(segmentID),
	}
}

// SetSegmentID : Allow user to set SegmentID
func (options *GetSegmentOptions) SetSegmentID(segmentID string) *GetSegmentOptions {
	options.SegmentID = core.StringPtr(segmentID)
	return options
}

// SetInclude : Allow user to set Include
func (options *GetSegmentOptions) SetInclude(include string) *GetSegmentOptions {
	options.Include = core.StringPtr(include)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetSegmentOptions) SetHeaders(param map[string]string) *GetSegmentOptions {
	options.Headers = param
	return options
}

// UpdateSegment : UpdateSegment struct
type UpdateSegment struct {
	// Segment name.
	Name *string `json:"name" validate:"required"`

	// Segment id.
	SegmentID *string `json:"segment_id" validate:"required"`

	// Segment description.
	Description *string `json:"description" validate:"required"`

	// Segment created time.
	CreatedTime *string `json:"created_time" validate:"required"`

	// Segment updated time.
	UpdatedTime *string `json:"updated_time" validate:"required"`
}

// UnmarshalUpdateSegment unmarshals an instance of UpdateSegment from the specified map of raw messages.
func UnmarshalUpdateSegment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UpdateSegment)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "segment_id", &obj.SegmentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_time", &obj.CreatedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_time", &obj.UpdatedTime)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateSegmentOptions : The UpdateSegment options.
type UpdateSegmentOptions struct {
	// Segment Id.
	SegmentID *string `json:"segment_id" validate:"required,ne="`

	// Segment name.
	Name *string `json:"name" validate:"required"`

	// Segment description.
	Description *string `json:"description" validate:"required"`

	// Tags associated with segments.
	Tags *string `json:"tags" validate:"required"`

	// Rule array.
	Rules []RuleArray `json:"rules" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateSegmentOptions : Instantiate UpdateSegmentOptions
func (*AppConfigurationV1) NewUpdateSegmentOptions(segmentID string, name string, description string, tags string, rules []RuleArray) *UpdateSegmentOptions {
	return &UpdateSegmentOptions{
		SegmentID:   core.StringPtr(segmentID),
		Name:        core.StringPtr(name),
		Description: core.StringPtr(description),
		Tags:        core.StringPtr(tags),
		Rules:       rules,
	}
}

// SetSegmentID : Allow user to set SegmentID
func (options *UpdateSegmentOptions) SetSegmentID(segmentID string) *UpdateSegmentOptions {
	options.SegmentID = core.StringPtr(segmentID)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateSegmentOptions) SetName(name string) *UpdateSegmentOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *UpdateSegmentOptions) SetDescription(description string) *UpdateSegmentOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetTags : Allow user to set Tags
func (options *UpdateSegmentOptions) SetTags(tags string) *UpdateSegmentOptions {
	options.Tags = core.StringPtr(tags)
	return options
}

// SetRules : Allow user to set Rules
func (options *UpdateSegmentOptions) SetRules(rules []RuleArray) *UpdateSegmentOptions {
	options.Rules = rules
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateSegmentOptions) SetHeaders(param map[string]string) *UpdateSegmentOptions {
	options.Headers = param
	return options
}

// SegmentRule : SegmentRule struct
type SegmentRule struct {
	// Rules array.
	Rules []Rule `json:"rules" validate:"required"`

	// Value of the segment.
	Value interface{} `json:"value" validate:"required"`

	// Order of the segment, used during evaluation.
	Order *int64 `json:"order" validate:"required"`
}

// NewSegmentRule : Instantiate SegmentRule (Generic Model Constructor)
func (*AppConfigurationV1) NewSegmentRule(rules []Rule, value string, order int64) (model *SegmentRule, err error) {
	model = &SegmentRule{
		Rules: rules,
		Value: value,
		Order: core.Int64Ptr(order),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalSegmentRule unmarshals an instance of SegmentRule from the specified map of raw messages.
func UnmarshalSegmentRule(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SegmentRule)
	err = core.UnmarshalModel(m, "rules", &obj.Rules, UnmarshalRule)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "order", &obj.Order)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
