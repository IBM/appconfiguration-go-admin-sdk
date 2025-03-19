/**
 * (C) Copyright IBM Corp. 2025.
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

package appconfigurationv1_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`AppConfigurationV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(appConfigurationService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(appConfigurationService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
				URL: "https://appconfigurationv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(appConfigurationService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"APP_CONFIGURATION_URL": "https://appconfigurationv1/api",
				"APP_CONFIGURATION_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1UsingExternalConfig(&appconfigurationv1.AppConfigurationV1Options{
				})
				Expect(appConfigurationService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := appConfigurationService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != appConfigurationService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(appConfigurationService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(appConfigurationService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1UsingExternalConfig(&appconfigurationv1.AppConfigurationV1Options{
					URL: "https://testService/api",
				})
				Expect(appConfigurationService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := appConfigurationService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != appConfigurationService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(appConfigurationService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(appConfigurationService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1UsingExternalConfig(&appconfigurationv1.AppConfigurationV1Options{
				})
				err := appConfigurationService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := appConfigurationService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != appConfigurationService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(appConfigurationService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(appConfigurationService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"APP_CONFIGURATION_URL": "https://appconfigurationv1/api",
				"APP_CONFIGURATION_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1UsingExternalConfig(&appconfigurationv1.AppConfigurationV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(appConfigurationService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"APP_CONFIGURATION_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1UsingExternalConfig(&appconfigurationv1.AppConfigurationV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(appConfigurationService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = appconfigurationv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`Parameterized URL tests`, func() {
		It(`Format parameterized URL with all default values`, func() {
			constructedURL, err := appconfigurationv1.ConstructServiceURL(nil)
			Expect(constructedURL).To(Equal("https://us-south.apprapp.cloud.ibm.com/apprapp/feature/v1/instances/provide-here-your-appconfig-instance-uuid"))
			Expect(constructedURL).ToNot(BeNil())
			Expect(err).To(BeNil())
		})
		It(`Return an error if a provided variable name is invalid`, func() {
			var providedUrlVariables = map[string]string{
				"invalid_variable_name": "value",
			}
			constructedURL, err := appconfigurationv1.ConstructServiceURL(providedUrlVariables)
			Expect(constructedURL).To(Equal(""))
			Expect(err).ToNot(BeNil())
		})
	})
	Describe(`ListEnvironments(listEnvironmentsOptions *ListEnvironmentsOptions) - Operation response error`, func() {
		listEnvironmentsPath := "/environments"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listEnvironmentsPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for expand query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_time"}))
					Expect(req.URL.Query()["tags"]).To(Equal([]string{"version 1.1,pre-release"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["search"]).To(Equal([]string{"test tag"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListEnvironments with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListEnvironmentsOptions model
				listEnvironmentsOptionsModel := new(appconfigurationv1.ListEnvironmentsOptions)
				listEnvironmentsOptionsModel.Expand = core.BoolPtr(true)
				listEnvironmentsOptionsModel.Sort = core.StringPtr("created_time")
				listEnvironmentsOptionsModel.Tags = core.StringPtr("version 1.1,pre-release")
				listEnvironmentsOptionsModel.Include = []string{"features", "properties", "snapshots"}
				listEnvironmentsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listEnvironmentsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listEnvironmentsOptionsModel.Search = core.StringPtr("test tag")
				listEnvironmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.ListEnvironments(listEnvironmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.ListEnvironments(listEnvironmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListEnvironments(listEnvironmentsOptions *ListEnvironmentsOptions)`, func() {
		listEnvironmentsPath := "/environments"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listEnvironmentsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for expand query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_time"}))
					Expect(req.URL.Query()["tags"]).To(Equal([]string{"version 1.1,pre-release"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["search"]).To(Equal([]string{"test tag"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"environments": [{"name": "Name", "environment_id": "EnvironmentID", "description": "Description", "tags": "Tags", "color_code": "#FDD13A", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href", "features": [{"feature_id": "FeatureID", "name": "Name"}], "properties": [{"property_id": "PropertyID", "name": "Name"}], "snapshots": [{"git_config_id": "GitConfigID", "name": "Name"}]}], "limit": 10, "offset": 0, "total_count": 0, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}}`)
				}))
			})
			It(`Invoke ListEnvironments successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the ListEnvironmentsOptions model
				listEnvironmentsOptionsModel := new(appconfigurationv1.ListEnvironmentsOptions)
				listEnvironmentsOptionsModel.Expand = core.BoolPtr(true)
				listEnvironmentsOptionsModel.Sort = core.StringPtr("created_time")
				listEnvironmentsOptionsModel.Tags = core.StringPtr("version 1.1,pre-release")
				listEnvironmentsOptionsModel.Include = []string{"features", "properties", "snapshots"}
				listEnvironmentsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listEnvironmentsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listEnvironmentsOptionsModel.Search = core.StringPtr("test tag")
				listEnvironmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.ListEnvironmentsWithContext(ctx, listEnvironmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.ListEnvironments(listEnvironmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.ListEnvironmentsWithContext(ctx, listEnvironmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listEnvironmentsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for expand query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_time"}))
					Expect(req.URL.Query()["tags"]).To(Equal([]string{"version 1.1,pre-release"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["search"]).To(Equal([]string{"test tag"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"environments": [{"name": "Name", "environment_id": "EnvironmentID", "description": "Description", "tags": "Tags", "color_code": "#FDD13A", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href", "features": [{"feature_id": "FeatureID", "name": "Name"}], "properties": [{"property_id": "PropertyID", "name": "Name"}], "snapshots": [{"git_config_id": "GitConfigID", "name": "Name"}]}], "limit": 10, "offset": 0, "total_count": 0, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}}`)
				}))
			})
			It(`Invoke ListEnvironments successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.ListEnvironments(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListEnvironmentsOptions model
				listEnvironmentsOptionsModel := new(appconfigurationv1.ListEnvironmentsOptions)
				listEnvironmentsOptionsModel.Expand = core.BoolPtr(true)
				listEnvironmentsOptionsModel.Sort = core.StringPtr("created_time")
				listEnvironmentsOptionsModel.Tags = core.StringPtr("version 1.1,pre-release")
				listEnvironmentsOptionsModel.Include = []string{"features", "properties", "snapshots"}
				listEnvironmentsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listEnvironmentsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listEnvironmentsOptionsModel.Search = core.StringPtr("test tag")
				listEnvironmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.ListEnvironments(listEnvironmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListEnvironments with error: Operation request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListEnvironmentsOptions model
				listEnvironmentsOptionsModel := new(appconfigurationv1.ListEnvironmentsOptions)
				listEnvironmentsOptionsModel.Expand = core.BoolPtr(true)
				listEnvironmentsOptionsModel.Sort = core.StringPtr("created_time")
				listEnvironmentsOptionsModel.Tags = core.StringPtr("version 1.1,pre-release")
				listEnvironmentsOptionsModel.Include = []string{"features", "properties", "snapshots"}
				listEnvironmentsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listEnvironmentsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listEnvironmentsOptionsModel.Search = core.StringPtr("test tag")
				listEnvironmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.ListEnvironments(listEnvironmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListEnvironments successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListEnvironmentsOptions model
				listEnvironmentsOptionsModel := new(appconfigurationv1.ListEnvironmentsOptions)
				listEnvironmentsOptionsModel.Expand = core.BoolPtr(true)
				listEnvironmentsOptionsModel.Sort = core.StringPtr("created_time")
				listEnvironmentsOptionsModel.Tags = core.StringPtr("version 1.1,pre-release")
				listEnvironmentsOptionsModel.Include = []string{"features", "properties", "snapshots"}
				listEnvironmentsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listEnvironmentsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listEnvironmentsOptionsModel.Search = core.StringPtr("test tag")
				listEnvironmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.ListEnvironments(listEnvironmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(appconfigurationv1.EnvironmentList)
				nextObject := new(appconfigurationv1.PaginatedListNext)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(appconfigurationv1.EnvironmentList)

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(appconfigurationv1.EnvironmentList)
				nextObject := new(appconfigurationv1.PaginatedListNext)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(appconfigurationv1.EnvironmentList)
				nextObject := new(appconfigurationv1.PaginatedListNext)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listEnvironmentsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"environments":[{"name":"Name","environment_id":"EnvironmentID","description":"Description","tags":"Tags","color_code":"#FDD13A","created_time":"2021-05-12T23:20:50.520Z","updated_time":"2021-05-12T23:20:50.520Z","href":"Href","features":[{"feature_id":"FeatureID","name":"Name"}],"properties":[{"property_id":"PropertyID","name":"Name"}],"snapshots":[{"git_config_id":"GitConfigID","name":"Name"}]}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"environments":[{"name":"Name","environment_id":"EnvironmentID","description":"Description","tags":"Tags","color_code":"#FDD13A","created_time":"2021-05-12T23:20:50.520Z","updated_time":"2021-05-12T23:20:50.520Z","href":"Href","features":[{"feature_id":"FeatureID","name":"Name"}],"properties":[{"property_id":"PropertyID","name":"Name"}],"snapshots":[{"git_config_id":"GitConfigID","name":"Name"}]}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use EnvironmentsPager.GetNext successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				listEnvironmentsOptionsModel := &appconfigurationv1.ListEnvironmentsOptions{
					Expand: core.BoolPtr(true),
					Sort: core.StringPtr("created_time"),
					Tags: core.StringPtr("version 1.1,pre-release"),
					Include: []string{"features", "properties", "snapshots"},
					Limit: core.Int64Ptr(int64(10)),
					Search: core.StringPtr("test tag"),
				}

				pager, err := appConfigurationService.NewEnvironmentsPager(listEnvironmentsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []appconfigurationv1.Environment
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use EnvironmentsPager.GetAll successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				listEnvironmentsOptionsModel := &appconfigurationv1.ListEnvironmentsOptions{
					Expand: core.BoolPtr(true),
					Sort: core.StringPtr("created_time"),
					Tags: core.StringPtr("version 1.1,pre-release"),
					Include: []string{"features", "properties", "snapshots"},
					Limit: core.Int64Ptr(int64(10)),
					Search: core.StringPtr("test tag"),
				}

				pager, err := appConfigurationService.NewEnvironmentsPager(listEnvironmentsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateEnvironment(createEnvironmentOptions *CreateEnvironmentOptions) - Operation response error`, func() {
		createEnvironmentPath := "/environments"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEnvironmentPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateEnvironment with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the CreateEnvironmentOptions model
				createEnvironmentOptionsModel := new(appconfigurationv1.CreateEnvironmentOptions)
				createEnvironmentOptionsModel.Name = core.StringPtr("Dev environment")
				createEnvironmentOptionsModel.EnvironmentID = core.StringPtr("dev-environment")
				createEnvironmentOptionsModel.Description = core.StringPtr("Dev environment description")
				createEnvironmentOptionsModel.Tags = core.StringPtr("development")
				createEnvironmentOptionsModel.ColorCode = core.StringPtr("#FDD13A")
				createEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.CreateEnvironment(createEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.CreateEnvironment(createEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateEnvironment(createEnvironmentOptions *CreateEnvironmentOptions)`, func() {
		createEnvironmentPath := "/environments"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEnvironmentPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Name", "environment_id": "EnvironmentID", "description": "Description", "tags": "Tags", "color_code": "#FDD13A", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href", "features": [{"feature_id": "FeatureID", "name": "Name"}], "properties": [{"property_id": "PropertyID", "name": "Name"}], "snapshots": [{"git_config_id": "GitConfigID", "name": "Name"}]}`)
				}))
			})
			It(`Invoke CreateEnvironment successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the CreateEnvironmentOptions model
				createEnvironmentOptionsModel := new(appconfigurationv1.CreateEnvironmentOptions)
				createEnvironmentOptionsModel.Name = core.StringPtr("Dev environment")
				createEnvironmentOptionsModel.EnvironmentID = core.StringPtr("dev-environment")
				createEnvironmentOptionsModel.Description = core.StringPtr("Dev environment description")
				createEnvironmentOptionsModel.Tags = core.StringPtr("development")
				createEnvironmentOptionsModel.ColorCode = core.StringPtr("#FDD13A")
				createEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.CreateEnvironmentWithContext(ctx, createEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.CreateEnvironment(createEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.CreateEnvironmentWithContext(ctx, createEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEnvironmentPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Name", "environment_id": "EnvironmentID", "description": "Description", "tags": "Tags", "color_code": "#FDD13A", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href", "features": [{"feature_id": "FeatureID", "name": "Name"}], "properties": [{"property_id": "PropertyID", "name": "Name"}], "snapshots": [{"git_config_id": "GitConfigID", "name": "Name"}]}`)
				}))
			})
			It(`Invoke CreateEnvironment successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.CreateEnvironment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateEnvironmentOptions model
				createEnvironmentOptionsModel := new(appconfigurationv1.CreateEnvironmentOptions)
				createEnvironmentOptionsModel.Name = core.StringPtr("Dev environment")
				createEnvironmentOptionsModel.EnvironmentID = core.StringPtr("dev-environment")
				createEnvironmentOptionsModel.Description = core.StringPtr("Dev environment description")
				createEnvironmentOptionsModel.Tags = core.StringPtr("development")
				createEnvironmentOptionsModel.ColorCode = core.StringPtr("#FDD13A")
				createEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.CreateEnvironment(createEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateEnvironment with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the CreateEnvironmentOptions model
				createEnvironmentOptionsModel := new(appconfigurationv1.CreateEnvironmentOptions)
				createEnvironmentOptionsModel.Name = core.StringPtr("Dev environment")
				createEnvironmentOptionsModel.EnvironmentID = core.StringPtr("dev-environment")
				createEnvironmentOptionsModel.Description = core.StringPtr("Dev environment description")
				createEnvironmentOptionsModel.Tags = core.StringPtr("development")
				createEnvironmentOptionsModel.ColorCode = core.StringPtr("#FDD13A")
				createEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.CreateEnvironment(createEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateEnvironmentOptions model with no property values
				createEnvironmentOptionsModelNew := new(appconfigurationv1.CreateEnvironmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.CreateEnvironment(createEnvironmentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateEnvironment successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the CreateEnvironmentOptions model
				createEnvironmentOptionsModel := new(appconfigurationv1.CreateEnvironmentOptions)
				createEnvironmentOptionsModel.Name = core.StringPtr("Dev environment")
				createEnvironmentOptionsModel.EnvironmentID = core.StringPtr("dev-environment")
				createEnvironmentOptionsModel.Description = core.StringPtr("Dev environment description")
				createEnvironmentOptionsModel.Tags = core.StringPtr("development")
				createEnvironmentOptionsModel.ColorCode = core.StringPtr("#FDD13A")
				createEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.CreateEnvironment(createEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateEnvironment(updateEnvironmentOptions *UpdateEnvironmentOptions) - Operation response error`, func() {
		updateEnvironmentPath := "/environments/environment_id"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateEnvironmentPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateEnvironment with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the UpdateEnvironmentOptions model
				updateEnvironmentOptionsModel := new(appconfigurationv1.UpdateEnvironmentOptions)
				updateEnvironmentOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				updateEnvironmentOptionsModel.Name = core.StringPtr("testString")
				updateEnvironmentOptionsModel.Description = core.StringPtr("testString")
				updateEnvironmentOptionsModel.Tags = core.StringPtr("testString")
				updateEnvironmentOptionsModel.ColorCode = core.StringPtr("#FDD13A")
				updateEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.UpdateEnvironment(updateEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.UpdateEnvironment(updateEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateEnvironment(updateEnvironmentOptions *UpdateEnvironmentOptions)`, func() {
		updateEnvironmentPath := "/environments/environment_id"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateEnvironmentPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "environment_id": "EnvironmentID", "description": "Description", "tags": "Tags", "color_code": "#FDD13A", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href", "features": [{"feature_id": "FeatureID", "name": "Name"}], "properties": [{"property_id": "PropertyID", "name": "Name"}], "snapshots": [{"git_config_id": "GitConfigID", "name": "Name"}]}`)
				}))
			})
			It(`Invoke UpdateEnvironment successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the UpdateEnvironmentOptions model
				updateEnvironmentOptionsModel := new(appconfigurationv1.UpdateEnvironmentOptions)
				updateEnvironmentOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				updateEnvironmentOptionsModel.Name = core.StringPtr("testString")
				updateEnvironmentOptionsModel.Description = core.StringPtr("testString")
				updateEnvironmentOptionsModel.Tags = core.StringPtr("testString")
				updateEnvironmentOptionsModel.ColorCode = core.StringPtr("#FDD13A")
				updateEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.UpdateEnvironmentWithContext(ctx, updateEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.UpdateEnvironment(updateEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.UpdateEnvironmentWithContext(ctx, updateEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateEnvironmentPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "environment_id": "EnvironmentID", "description": "Description", "tags": "Tags", "color_code": "#FDD13A", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href", "features": [{"feature_id": "FeatureID", "name": "Name"}], "properties": [{"property_id": "PropertyID", "name": "Name"}], "snapshots": [{"git_config_id": "GitConfigID", "name": "Name"}]}`)
				}))
			})
			It(`Invoke UpdateEnvironment successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.UpdateEnvironment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateEnvironmentOptions model
				updateEnvironmentOptionsModel := new(appconfigurationv1.UpdateEnvironmentOptions)
				updateEnvironmentOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				updateEnvironmentOptionsModel.Name = core.StringPtr("testString")
				updateEnvironmentOptionsModel.Description = core.StringPtr("testString")
				updateEnvironmentOptionsModel.Tags = core.StringPtr("testString")
				updateEnvironmentOptionsModel.ColorCode = core.StringPtr("#FDD13A")
				updateEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.UpdateEnvironment(updateEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateEnvironment with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the UpdateEnvironmentOptions model
				updateEnvironmentOptionsModel := new(appconfigurationv1.UpdateEnvironmentOptions)
				updateEnvironmentOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				updateEnvironmentOptionsModel.Name = core.StringPtr("testString")
				updateEnvironmentOptionsModel.Description = core.StringPtr("testString")
				updateEnvironmentOptionsModel.Tags = core.StringPtr("testString")
				updateEnvironmentOptionsModel.ColorCode = core.StringPtr("#FDD13A")
				updateEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.UpdateEnvironment(updateEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateEnvironmentOptions model with no property values
				updateEnvironmentOptionsModelNew := new(appconfigurationv1.UpdateEnvironmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.UpdateEnvironment(updateEnvironmentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateEnvironment successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the UpdateEnvironmentOptions model
				updateEnvironmentOptionsModel := new(appconfigurationv1.UpdateEnvironmentOptions)
				updateEnvironmentOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				updateEnvironmentOptionsModel.Name = core.StringPtr("testString")
				updateEnvironmentOptionsModel.Description = core.StringPtr("testString")
				updateEnvironmentOptionsModel.Tags = core.StringPtr("testString")
				updateEnvironmentOptionsModel.ColorCode = core.StringPtr("#FDD13A")
				updateEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.UpdateEnvironment(updateEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetEnvironment(getEnvironmentOptions *GetEnvironmentOptions) - Operation response error`, func() {
		getEnvironmentPath := "/environments/environment_id"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEnvironmentPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for expand query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetEnvironment with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetEnvironmentOptions model
				getEnvironmentOptionsModel := new(appconfigurationv1.GetEnvironmentOptions)
				getEnvironmentOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				getEnvironmentOptionsModel.Expand = core.BoolPtr(true)
				getEnvironmentOptionsModel.Include = []string{"features", "properties", "snapshots"}
				getEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.GetEnvironment(getEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.GetEnvironment(getEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetEnvironment(getEnvironmentOptions *GetEnvironmentOptions)`, func() {
		getEnvironmentPath := "/environments/environment_id"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEnvironmentPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for expand query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "environment_id": "EnvironmentID", "description": "Description", "tags": "Tags", "color_code": "#FDD13A", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href", "features": [{"feature_id": "FeatureID", "name": "Name"}], "properties": [{"property_id": "PropertyID", "name": "Name"}], "snapshots": [{"git_config_id": "GitConfigID", "name": "Name"}]}`)
				}))
			})
			It(`Invoke GetEnvironment successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the GetEnvironmentOptions model
				getEnvironmentOptionsModel := new(appconfigurationv1.GetEnvironmentOptions)
				getEnvironmentOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				getEnvironmentOptionsModel.Expand = core.BoolPtr(true)
				getEnvironmentOptionsModel.Include = []string{"features", "properties", "snapshots"}
				getEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.GetEnvironmentWithContext(ctx, getEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.GetEnvironment(getEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.GetEnvironmentWithContext(ctx, getEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEnvironmentPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for expand query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "environment_id": "EnvironmentID", "description": "Description", "tags": "Tags", "color_code": "#FDD13A", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href", "features": [{"feature_id": "FeatureID", "name": "Name"}], "properties": [{"property_id": "PropertyID", "name": "Name"}], "snapshots": [{"git_config_id": "GitConfigID", "name": "Name"}]}`)
				}))
			})
			It(`Invoke GetEnvironment successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.GetEnvironment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetEnvironmentOptions model
				getEnvironmentOptionsModel := new(appconfigurationv1.GetEnvironmentOptions)
				getEnvironmentOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				getEnvironmentOptionsModel.Expand = core.BoolPtr(true)
				getEnvironmentOptionsModel.Include = []string{"features", "properties", "snapshots"}
				getEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.GetEnvironment(getEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetEnvironment with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetEnvironmentOptions model
				getEnvironmentOptionsModel := new(appconfigurationv1.GetEnvironmentOptions)
				getEnvironmentOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				getEnvironmentOptionsModel.Expand = core.BoolPtr(true)
				getEnvironmentOptionsModel.Include = []string{"features", "properties", "snapshots"}
				getEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.GetEnvironment(getEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetEnvironmentOptions model with no property values
				getEnvironmentOptionsModelNew := new(appconfigurationv1.GetEnvironmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.GetEnvironment(getEnvironmentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetEnvironment successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetEnvironmentOptions model
				getEnvironmentOptionsModel := new(appconfigurationv1.GetEnvironmentOptions)
				getEnvironmentOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				getEnvironmentOptionsModel.Expand = core.BoolPtr(true)
				getEnvironmentOptionsModel.Include = []string{"features", "properties", "snapshots"}
				getEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.GetEnvironment(getEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteEnvironment(deleteEnvironmentOptions *DeleteEnvironmentOptions)`, func() {
		deleteEnvironmentPath := "/environments/environment_id"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteEnvironmentPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteEnvironment successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appConfigurationService.DeleteEnvironment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteEnvironmentOptions model
				deleteEnvironmentOptionsModel := new(appconfigurationv1.DeleteEnvironmentOptions)
				deleteEnvironmentOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				deleteEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appConfigurationService.DeleteEnvironment(deleteEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteEnvironment with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the DeleteEnvironmentOptions model
				deleteEnvironmentOptionsModel := new(appconfigurationv1.DeleteEnvironmentOptions)
				deleteEnvironmentOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				deleteEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appConfigurationService.DeleteEnvironment(deleteEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteEnvironmentOptions model with no property values
				deleteEnvironmentOptionsModelNew := new(appconfigurationv1.DeleteEnvironmentOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appConfigurationService.DeleteEnvironment(deleteEnvironmentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListCollections(listCollectionsOptions *ListCollectionsOptions) - Operation response error`, func() {
		listCollectionsPath := "/collections"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCollectionsPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for expand query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_time"}))
					Expect(req.URL.Query()["tags"]).To(Equal([]string{"version 1.1,pre-release"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["search"]).To(Equal([]string{"test tag"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListCollections with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListCollectionsOptions model
				listCollectionsOptionsModel := new(appconfigurationv1.ListCollectionsOptions)
				listCollectionsOptionsModel.Expand = core.BoolPtr(true)
				listCollectionsOptionsModel.Sort = core.StringPtr("created_time")
				listCollectionsOptionsModel.Tags = core.StringPtr("version 1.1,pre-release")
				listCollectionsOptionsModel.Features = []string{"my-feature-id", "cycle-rentals"}
				listCollectionsOptionsModel.Properties = []string{"my-property-id", "email-property"}
				listCollectionsOptionsModel.Include = []string{"features", "properties", "snapshots"}
				listCollectionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listCollectionsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listCollectionsOptionsModel.Search = core.StringPtr("test tag")
				listCollectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.ListCollections(listCollectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.ListCollections(listCollectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListCollections(listCollectionsOptions *ListCollectionsOptions)`, func() {
		listCollectionsPath := "/collections"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCollectionsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for expand query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_time"}))
					Expect(req.URL.Query()["tags"]).To(Equal([]string{"version 1.1,pre-release"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["search"]).To(Equal([]string{"test tag"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"collections": [{"name": "Name", "collection_id": "CollectionID", "description": "Description", "tags": "Tags", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href", "features": [{"feature_id": "FeatureID", "name": "Name"}], "properties": [{"property_id": "PropertyID", "name": "Name"}], "snapshots": [{"git_config_id": "GitConfigID", "name": "Name"}], "features_count": 13, "properties_count": 15, "snapshot_count": 13}], "limit": 10, "offset": 0, "total_count": 0, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}}`)
				}))
			})
			It(`Invoke ListCollections successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the ListCollectionsOptions model
				listCollectionsOptionsModel := new(appconfigurationv1.ListCollectionsOptions)
				listCollectionsOptionsModel.Expand = core.BoolPtr(true)
				listCollectionsOptionsModel.Sort = core.StringPtr("created_time")
				listCollectionsOptionsModel.Tags = core.StringPtr("version 1.1,pre-release")
				listCollectionsOptionsModel.Features = []string{"my-feature-id", "cycle-rentals"}
				listCollectionsOptionsModel.Properties = []string{"my-property-id", "email-property"}
				listCollectionsOptionsModel.Include = []string{"features", "properties", "snapshots"}
				listCollectionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listCollectionsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listCollectionsOptionsModel.Search = core.StringPtr("test tag")
				listCollectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.ListCollectionsWithContext(ctx, listCollectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.ListCollections(listCollectionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.ListCollectionsWithContext(ctx, listCollectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCollectionsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for expand query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_time"}))
					Expect(req.URL.Query()["tags"]).To(Equal([]string{"version 1.1,pre-release"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["search"]).To(Equal([]string{"test tag"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"collections": [{"name": "Name", "collection_id": "CollectionID", "description": "Description", "tags": "Tags", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href", "features": [{"feature_id": "FeatureID", "name": "Name"}], "properties": [{"property_id": "PropertyID", "name": "Name"}], "snapshots": [{"git_config_id": "GitConfigID", "name": "Name"}], "features_count": 13, "properties_count": 15, "snapshot_count": 13}], "limit": 10, "offset": 0, "total_count": 0, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}}`)
				}))
			})
			It(`Invoke ListCollections successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.ListCollections(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListCollectionsOptions model
				listCollectionsOptionsModel := new(appconfigurationv1.ListCollectionsOptions)
				listCollectionsOptionsModel.Expand = core.BoolPtr(true)
				listCollectionsOptionsModel.Sort = core.StringPtr("created_time")
				listCollectionsOptionsModel.Tags = core.StringPtr("version 1.1,pre-release")
				listCollectionsOptionsModel.Features = []string{"my-feature-id", "cycle-rentals"}
				listCollectionsOptionsModel.Properties = []string{"my-property-id", "email-property"}
				listCollectionsOptionsModel.Include = []string{"features", "properties", "snapshots"}
				listCollectionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listCollectionsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listCollectionsOptionsModel.Search = core.StringPtr("test tag")
				listCollectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.ListCollections(listCollectionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListCollections with error: Operation request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListCollectionsOptions model
				listCollectionsOptionsModel := new(appconfigurationv1.ListCollectionsOptions)
				listCollectionsOptionsModel.Expand = core.BoolPtr(true)
				listCollectionsOptionsModel.Sort = core.StringPtr("created_time")
				listCollectionsOptionsModel.Tags = core.StringPtr("version 1.1,pre-release")
				listCollectionsOptionsModel.Features = []string{"my-feature-id", "cycle-rentals"}
				listCollectionsOptionsModel.Properties = []string{"my-property-id", "email-property"}
				listCollectionsOptionsModel.Include = []string{"features", "properties", "snapshots"}
				listCollectionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listCollectionsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listCollectionsOptionsModel.Search = core.StringPtr("test tag")
				listCollectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.ListCollections(listCollectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListCollections successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListCollectionsOptions model
				listCollectionsOptionsModel := new(appconfigurationv1.ListCollectionsOptions)
				listCollectionsOptionsModel.Expand = core.BoolPtr(true)
				listCollectionsOptionsModel.Sort = core.StringPtr("created_time")
				listCollectionsOptionsModel.Tags = core.StringPtr("version 1.1,pre-release")
				listCollectionsOptionsModel.Features = []string{"my-feature-id", "cycle-rentals"}
				listCollectionsOptionsModel.Properties = []string{"my-property-id", "email-property"}
				listCollectionsOptionsModel.Include = []string{"features", "properties", "snapshots"}
				listCollectionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listCollectionsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listCollectionsOptionsModel.Search = core.StringPtr("test tag")
				listCollectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.ListCollections(listCollectionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(appconfigurationv1.CollectionList)
				nextObject := new(appconfigurationv1.PaginatedListNext)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(appconfigurationv1.CollectionList)

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(appconfigurationv1.CollectionList)
				nextObject := new(appconfigurationv1.PaginatedListNext)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(appconfigurationv1.CollectionList)
				nextObject := new(appconfigurationv1.PaginatedListNext)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCollectionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"collections":[{"name":"Name","collection_id":"CollectionID","description":"Description","tags":"Tags","created_time":"2021-05-12T23:20:50.520Z","updated_time":"2021-05-12T23:20:50.520Z","href":"Href","features":[{"feature_id":"FeatureID","name":"Name"}],"properties":[{"property_id":"PropertyID","name":"Name"}],"snapshots":[{"git_config_id":"GitConfigID","name":"Name"}],"features_count":13,"properties_count":15,"snapshot_count":13}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"collections":[{"name":"Name","collection_id":"CollectionID","description":"Description","tags":"Tags","created_time":"2021-05-12T23:20:50.520Z","updated_time":"2021-05-12T23:20:50.520Z","href":"Href","features":[{"feature_id":"FeatureID","name":"Name"}],"properties":[{"property_id":"PropertyID","name":"Name"}],"snapshots":[{"git_config_id":"GitConfigID","name":"Name"}],"features_count":13,"properties_count":15,"snapshot_count":13}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use CollectionsPager.GetNext successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				listCollectionsOptionsModel := &appconfigurationv1.ListCollectionsOptions{
					Expand: core.BoolPtr(true),
					Sort: core.StringPtr("created_time"),
					Tags: core.StringPtr("version 1.1,pre-release"),
					Features: []string{"my-feature-id", "cycle-rentals"},
					Properties: []string{"my-property-id", "email-property"},
					Include: []string{"features", "properties", "snapshots"},
					Limit: core.Int64Ptr(int64(10)),
					Search: core.StringPtr("test tag"),
				}

				pager, err := appConfigurationService.NewCollectionsPager(listCollectionsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []appconfigurationv1.Collection
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use CollectionsPager.GetAll successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				listCollectionsOptionsModel := &appconfigurationv1.ListCollectionsOptions{
					Expand: core.BoolPtr(true),
					Sort: core.StringPtr("created_time"),
					Tags: core.StringPtr("version 1.1,pre-release"),
					Features: []string{"my-feature-id", "cycle-rentals"},
					Properties: []string{"my-property-id", "email-property"},
					Include: []string{"features", "properties", "snapshots"},
					Limit: core.Int64Ptr(int64(10)),
					Search: core.StringPtr("test tag"),
				}

				pager, err := appConfigurationService.NewCollectionsPager(listCollectionsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateCollection(createCollectionOptions *CreateCollectionOptions) - Operation response error`, func() {
		createCollectionPath := "/collections"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCollectionPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateCollection with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the CreateCollectionOptions model
				createCollectionOptionsModel := new(appconfigurationv1.CreateCollectionOptions)
				createCollectionOptionsModel.Name = core.StringPtr("Web App Collection")
				createCollectionOptionsModel.CollectionID = core.StringPtr("web-app-collection")
				createCollectionOptionsModel.Description = core.StringPtr("Collection for Web application")
				createCollectionOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				createCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.CreateCollection(createCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.CreateCollection(createCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCollection(createCollectionOptions *CreateCollectionOptions)`, func() {
		createCollectionPath := "/collections"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCollectionPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Name", "collection_id": "CollectionID", "description": "Description", "tags": "Tags", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href"}`)
				}))
			})
			It(`Invoke CreateCollection successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the CreateCollectionOptions model
				createCollectionOptionsModel := new(appconfigurationv1.CreateCollectionOptions)
				createCollectionOptionsModel.Name = core.StringPtr("Web App Collection")
				createCollectionOptionsModel.CollectionID = core.StringPtr("web-app-collection")
				createCollectionOptionsModel.Description = core.StringPtr("Collection for Web application")
				createCollectionOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				createCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.CreateCollectionWithContext(ctx, createCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.CreateCollection(createCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.CreateCollectionWithContext(ctx, createCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCollectionPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Name", "collection_id": "CollectionID", "description": "Description", "tags": "Tags", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href"}`)
				}))
			})
			It(`Invoke CreateCollection successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.CreateCollection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateCollectionOptions model
				createCollectionOptionsModel := new(appconfigurationv1.CreateCollectionOptions)
				createCollectionOptionsModel.Name = core.StringPtr("Web App Collection")
				createCollectionOptionsModel.CollectionID = core.StringPtr("web-app-collection")
				createCollectionOptionsModel.Description = core.StringPtr("Collection for Web application")
				createCollectionOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				createCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.CreateCollection(createCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateCollection with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the CreateCollectionOptions model
				createCollectionOptionsModel := new(appconfigurationv1.CreateCollectionOptions)
				createCollectionOptionsModel.Name = core.StringPtr("Web App Collection")
				createCollectionOptionsModel.CollectionID = core.StringPtr("web-app-collection")
				createCollectionOptionsModel.Description = core.StringPtr("Collection for Web application")
				createCollectionOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				createCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.CreateCollection(createCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateCollectionOptions model with no property values
				createCollectionOptionsModelNew := new(appconfigurationv1.CreateCollectionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.CreateCollection(createCollectionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateCollection successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the CreateCollectionOptions model
				createCollectionOptionsModel := new(appconfigurationv1.CreateCollectionOptions)
				createCollectionOptionsModel.Name = core.StringPtr("Web App Collection")
				createCollectionOptionsModel.CollectionID = core.StringPtr("web-app-collection")
				createCollectionOptionsModel.Description = core.StringPtr("Collection for Web application")
				createCollectionOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				createCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.CreateCollection(createCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCollection(updateCollectionOptions *UpdateCollectionOptions) - Operation response error`, func() {
		updateCollectionPath := "/collections/collection_id"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCollectionPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateCollection with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the UpdateCollectionOptions model
				updateCollectionOptionsModel := new(appconfigurationv1.UpdateCollectionOptions)
				updateCollectionOptionsModel.CollectionID = core.StringPtr("collection_id")
				updateCollectionOptionsModel.Name = core.StringPtr("testString")
				updateCollectionOptionsModel.Description = core.StringPtr("testString")
				updateCollectionOptionsModel.Tags = core.StringPtr("testString")
				updateCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.UpdateCollection(updateCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.UpdateCollection(updateCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCollection(updateCollectionOptions *UpdateCollectionOptions)`, func() {
		updateCollectionPath := "/collections/collection_id"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCollectionPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "collection_id": "CollectionID", "description": "Description", "tags": "Tags", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href"}`)
				}))
			})
			It(`Invoke UpdateCollection successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the UpdateCollectionOptions model
				updateCollectionOptionsModel := new(appconfigurationv1.UpdateCollectionOptions)
				updateCollectionOptionsModel.CollectionID = core.StringPtr("collection_id")
				updateCollectionOptionsModel.Name = core.StringPtr("testString")
				updateCollectionOptionsModel.Description = core.StringPtr("testString")
				updateCollectionOptionsModel.Tags = core.StringPtr("testString")
				updateCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.UpdateCollectionWithContext(ctx, updateCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.UpdateCollection(updateCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.UpdateCollectionWithContext(ctx, updateCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCollectionPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "collection_id": "CollectionID", "description": "Description", "tags": "Tags", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href"}`)
				}))
			})
			It(`Invoke UpdateCollection successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.UpdateCollection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateCollectionOptions model
				updateCollectionOptionsModel := new(appconfigurationv1.UpdateCollectionOptions)
				updateCollectionOptionsModel.CollectionID = core.StringPtr("collection_id")
				updateCollectionOptionsModel.Name = core.StringPtr("testString")
				updateCollectionOptionsModel.Description = core.StringPtr("testString")
				updateCollectionOptionsModel.Tags = core.StringPtr("testString")
				updateCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.UpdateCollection(updateCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateCollection with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the UpdateCollectionOptions model
				updateCollectionOptionsModel := new(appconfigurationv1.UpdateCollectionOptions)
				updateCollectionOptionsModel.CollectionID = core.StringPtr("collection_id")
				updateCollectionOptionsModel.Name = core.StringPtr("testString")
				updateCollectionOptionsModel.Description = core.StringPtr("testString")
				updateCollectionOptionsModel.Tags = core.StringPtr("testString")
				updateCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.UpdateCollection(updateCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateCollectionOptions model with no property values
				updateCollectionOptionsModelNew := new(appconfigurationv1.UpdateCollectionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.UpdateCollection(updateCollectionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateCollection successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the UpdateCollectionOptions model
				updateCollectionOptionsModel := new(appconfigurationv1.UpdateCollectionOptions)
				updateCollectionOptionsModel.CollectionID = core.StringPtr("collection_id")
				updateCollectionOptionsModel.Name = core.StringPtr("testString")
				updateCollectionOptionsModel.Description = core.StringPtr("testString")
				updateCollectionOptionsModel.Tags = core.StringPtr("testString")
				updateCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.UpdateCollection(updateCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCollection(getCollectionOptions *GetCollectionOptions) - Operation response error`, func() {
		getCollectionPath := "/collections/collection_id"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCollectionPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for expand query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCollection with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetCollectionOptions model
				getCollectionOptionsModel := new(appconfigurationv1.GetCollectionOptions)
				getCollectionOptionsModel.CollectionID = core.StringPtr("collection_id")
				getCollectionOptionsModel.Expand = core.BoolPtr(true)
				getCollectionOptionsModel.Include = []string{"features", "properties", "snapshots"}
				getCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.GetCollection(getCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.GetCollection(getCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCollection(getCollectionOptions *GetCollectionOptions)`, func() {
		getCollectionPath := "/collections/collection_id"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCollectionPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for expand query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "collection_id": "CollectionID", "description": "Description", "tags": "Tags", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href", "features": [{"feature_id": "FeatureID", "name": "Name"}], "properties": [{"property_id": "PropertyID", "name": "Name"}], "snapshots": [{"git_config_id": "GitConfigID", "name": "Name"}], "features_count": 13, "properties_count": 15, "snapshot_count": 13}`)
				}))
			})
			It(`Invoke GetCollection successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the GetCollectionOptions model
				getCollectionOptionsModel := new(appconfigurationv1.GetCollectionOptions)
				getCollectionOptionsModel.CollectionID = core.StringPtr("collection_id")
				getCollectionOptionsModel.Expand = core.BoolPtr(true)
				getCollectionOptionsModel.Include = []string{"features", "properties", "snapshots"}
				getCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.GetCollectionWithContext(ctx, getCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.GetCollection(getCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.GetCollectionWithContext(ctx, getCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCollectionPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for expand query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "collection_id": "CollectionID", "description": "Description", "tags": "Tags", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href", "features": [{"feature_id": "FeatureID", "name": "Name"}], "properties": [{"property_id": "PropertyID", "name": "Name"}], "snapshots": [{"git_config_id": "GitConfigID", "name": "Name"}], "features_count": 13, "properties_count": 15, "snapshot_count": 13}`)
				}))
			})
			It(`Invoke GetCollection successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.GetCollection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCollectionOptions model
				getCollectionOptionsModel := new(appconfigurationv1.GetCollectionOptions)
				getCollectionOptionsModel.CollectionID = core.StringPtr("collection_id")
				getCollectionOptionsModel.Expand = core.BoolPtr(true)
				getCollectionOptionsModel.Include = []string{"features", "properties", "snapshots"}
				getCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.GetCollection(getCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCollection with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetCollectionOptions model
				getCollectionOptionsModel := new(appconfigurationv1.GetCollectionOptions)
				getCollectionOptionsModel.CollectionID = core.StringPtr("collection_id")
				getCollectionOptionsModel.Expand = core.BoolPtr(true)
				getCollectionOptionsModel.Include = []string{"features", "properties", "snapshots"}
				getCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.GetCollection(getCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCollectionOptions model with no property values
				getCollectionOptionsModelNew := new(appconfigurationv1.GetCollectionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.GetCollection(getCollectionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetCollection successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetCollectionOptions model
				getCollectionOptionsModel := new(appconfigurationv1.GetCollectionOptions)
				getCollectionOptionsModel.CollectionID = core.StringPtr("collection_id")
				getCollectionOptionsModel.Expand = core.BoolPtr(true)
				getCollectionOptionsModel.Include = []string{"features", "properties", "snapshots"}
				getCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.GetCollection(getCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteCollection(deleteCollectionOptions *DeleteCollectionOptions)`, func() {
		deleteCollectionPath := "/collections/collection_id"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCollectionPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteCollection successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appConfigurationService.DeleteCollection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteCollectionOptions model
				deleteCollectionOptionsModel := new(appconfigurationv1.DeleteCollectionOptions)
				deleteCollectionOptionsModel.CollectionID = core.StringPtr("collection_id")
				deleteCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appConfigurationService.DeleteCollection(deleteCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteCollection with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the DeleteCollectionOptions model
				deleteCollectionOptionsModel := new(appconfigurationv1.DeleteCollectionOptions)
				deleteCollectionOptionsModel.CollectionID = core.StringPtr("collection_id")
				deleteCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appConfigurationService.DeleteCollection(deleteCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteCollectionOptions model with no property values
				deleteCollectionOptionsModelNew := new(appconfigurationv1.DeleteCollectionOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appConfigurationService.DeleteCollection(deleteCollectionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListFeatures(listFeaturesOptions *ListFeaturesOptions) - Operation response error`, func() {
		listFeaturesPath := "/environments/environment_id/features"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listFeaturesPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for expand query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_time"}))
					Expect(req.URL.Query()["tags"]).To(Equal([]string{"version 1.1,pre-release"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["search"]).To(Equal([]string{"test tag"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListFeatures with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListFeaturesOptions model
				listFeaturesOptionsModel := new(appconfigurationv1.ListFeaturesOptions)
				listFeaturesOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				listFeaturesOptionsModel.Expand = core.BoolPtr(true)
				listFeaturesOptionsModel.Sort = core.StringPtr("created_time")
				listFeaturesOptionsModel.Tags = core.StringPtr("version 1.1,pre-release")
				listFeaturesOptionsModel.Collections = []string{"my-collection-id", "ghzindiapvtltd"}
				listFeaturesOptionsModel.Segments = []string{"my-segment-id", "beta-users"}
				listFeaturesOptionsModel.Include = []string{"collections", "rules", "change_request"}
				listFeaturesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listFeaturesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listFeaturesOptionsModel.Search = core.StringPtr("test tag")
				listFeaturesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.ListFeatures(listFeaturesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.ListFeatures(listFeaturesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListFeatures(listFeaturesOptions *ListFeaturesOptions)`, func() {
		listFeaturesPath := "/environments/environment_id/features"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listFeaturesPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for expand query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_time"}))
					Expect(req.URL.Query()["tags"]).To(Equal([]string{"version 1.1,pre-release"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["search"]).To(Equal([]string{"test tag"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"features": [{"name": "Name", "feature_id": "FeatureID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "enabled_value": "anyValue", "disabled_value": "anyValue", "enabled": false, "rollout_percentage": 100, "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5, "rollout_percentage": 100}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "change_request_number": "ChangeRequestNumber", "change_request_status": "ChangeRequestStatus", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "evaluation_time": "2021-05-12T23:20:50.520Z", "href": "Href"}], "limit": 10, "offset": 0, "total_count": 0, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}}`)
				}))
			})
			It(`Invoke ListFeatures successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the ListFeaturesOptions model
				listFeaturesOptionsModel := new(appconfigurationv1.ListFeaturesOptions)
				listFeaturesOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				listFeaturesOptionsModel.Expand = core.BoolPtr(true)
				listFeaturesOptionsModel.Sort = core.StringPtr("created_time")
				listFeaturesOptionsModel.Tags = core.StringPtr("version 1.1,pre-release")
				listFeaturesOptionsModel.Collections = []string{"my-collection-id", "ghzindiapvtltd"}
				listFeaturesOptionsModel.Segments = []string{"my-segment-id", "beta-users"}
				listFeaturesOptionsModel.Include = []string{"collections", "rules", "change_request"}
				listFeaturesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listFeaturesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listFeaturesOptionsModel.Search = core.StringPtr("test tag")
				listFeaturesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.ListFeaturesWithContext(ctx, listFeaturesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.ListFeatures(listFeaturesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.ListFeaturesWithContext(ctx, listFeaturesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listFeaturesPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for expand query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_time"}))
					Expect(req.URL.Query()["tags"]).To(Equal([]string{"version 1.1,pre-release"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["search"]).To(Equal([]string{"test tag"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"features": [{"name": "Name", "feature_id": "FeatureID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "enabled_value": "anyValue", "disabled_value": "anyValue", "enabled": false, "rollout_percentage": 100, "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5, "rollout_percentage": 100}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "change_request_number": "ChangeRequestNumber", "change_request_status": "ChangeRequestStatus", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "evaluation_time": "2021-05-12T23:20:50.520Z", "href": "Href"}], "limit": 10, "offset": 0, "total_count": 0, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}}`)
				}))
			})
			It(`Invoke ListFeatures successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.ListFeatures(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListFeaturesOptions model
				listFeaturesOptionsModel := new(appconfigurationv1.ListFeaturesOptions)
				listFeaturesOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				listFeaturesOptionsModel.Expand = core.BoolPtr(true)
				listFeaturesOptionsModel.Sort = core.StringPtr("created_time")
				listFeaturesOptionsModel.Tags = core.StringPtr("version 1.1,pre-release")
				listFeaturesOptionsModel.Collections = []string{"my-collection-id", "ghzindiapvtltd"}
				listFeaturesOptionsModel.Segments = []string{"my-segment-id", "beta-users"}
				listFeaturesOptionsModel.Include = []string{"collections", "rules", "change_request"}
				listFeaturesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listFeaturesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listFeaturesOptionsModel.Search = core.StringPtr("test tag")
				listFeaturesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.ListFeatures(listFeaturesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListFeatures with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListFeaturesOptions model
				listFeaturesOptionsModel := new(appconfigurationv1.ListFeaturesOptions)
				listFeaturesOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				listFeaturesOptionsModel.Expand = core.BoolPtr(true)
				listFeaturesOptionsModel.Sort = core.StringPtr("created_time")
				listFeaturesOptionsModel.Tags = core.StringPtr("version 1.1,pre-release")
				listFeaturesOptionsModel.Collections = []string{"my-collection-id", "ghzindiapvtltd"}
				listFeaturesOptionsModel.Segments = []string{"my-segment-id", "beta-users"}
				listFeaturesOptionsModel.Include = []string{"collections", "rules", "change_request"}
				listFeaturesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listFeaturesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listFeaturesOptionsModel.Search = core.StringPtr("test tag")
				listFeaturesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.ListFeatures(listFeaturesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListFeaturesOptions model with no property values
				listFeaturesOptionsModelNew := new(appconfigurationv1.ListFeaturesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.ListFeatures(listFeaturesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListFeatures successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListFeaturesOptions model
				listFeaturesOptionsModel := new(appconfigurationv1.ListFeaturesOptions)
				listFeaturesOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				listFeaturesOptionsModel.Expand = core.BoolPtr(true)
				listFeaturesOptionsModel.Sort = core.StringPtr("created_time")
				listFeaturesOptionsModel.Tags = core.StringPtr("version 1.1,pre-release")
				listFeaturesOptionsModel.Collections = []string{"my-collection-id", "ghzindiapvtltd"}
				listFeaturesOptionsModel.Segments = []string{"my-segment-id", "beta-users"}
				listFeaturesOptionsModel.Include = []string{"collections", "rules", "change_request"}
				listFeaturesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listFeaturesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listFeaturesOptionsModel.Search = core.StringPtr("test tag")
				listFeaturesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.ListFeatures(listFeaturesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(appconfigurationv1.FeaturesList)
				nextObject := new(appconfigurationv1.PaginatedListNext)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(appconfigurationv1.FeaturesList)

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(appconfigurationv1.FeaturesList)
				nextObject := new(appconfigurationv1.PaginatedListNext)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(appconfigurationv1.FeaturesList)
				nextObject := new(appconfigurationv1.PaginatedListNext)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listFeaturesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"features":[{"name":"Name","feature_id":"FeatureID","description":"Description","type":"BOOLEAN","format":"TEXT","enabled_value":"anyValue","disabled_value":"anyValue","enabled":false,"rollout_percentage":100,"tags":"Tags","segment_rules":[{"rules":[{"segments":["Segments"]}],"value":"anyValue","order":5,"rollout_percentage":100}],"segment_exists":false,"collections":[{"collection_id":"CollectionID","name":"Name"}],"change_request_number":"ChangeRequestNumber","change_request_status":"ChangeRequestStatus","created_time":"2021-05-12T23:20:50.520Z","updated_time":"2021-05-12T23:20:50.520Z","evaluation_time":"2021-05-12T23:20:50.520Z","href":"Href"}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"features":[{"name":"Name","feature_id":"FeatureID","description":"Description","type":"BOOLEAN","format":"TEXT","enabled_value":"anyValue","disabled_value":"anyValue","enabled":false,"rollout_percentage":100,"tags":"Tags","segment_rules":[{"rules":[{"segments":["Segments"]}],"value":"anyValue","order":5,"rollout_percentage":100}],"segment_exists":false,"collections":[{"collection_id":"CollectionID","name":"Name"}],"change_request_number":"ChangeRequestNumber","change_request_status":"ChangeRequestStatus","created_time":"2021-05-12T23:20:50.520Z","updated_time":"2021-05-12T23:20:50.520Z","evaluation_time":"2021-05-12T23:20:50.520Z","href":"Href"}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use FeaturesPager.GetNext successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				listFeaturesOptionsModel := &appconfigurationv1.ListFeaturesOptions{
					EnvironmentID: core.StringPtr("environment_id"),
					Expand: core.BoolPtr(true),
					Sort: core.StringPtr("created_time"),
					Tags: core.StringPtr("version 1.1,pre-release"),
					Collections: []string{"my-collection-id", "ghzindiapvtltd"},
					Segments: []string{"my-segment-id", "beta-users"},
					Include: []string{"collections", "rules", "change_request"},
					Limit: core.Int64Ptr(int64(10)),
					Search: core.StringPtr("test tag"),
				}

				pager, err := appConfigurationService.NewFeaturesPager(listFeaturesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []appconfigurationv1.Feature
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use FeaturesPager.GetAll successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				listFeaturesOptionsModel := &appconfigurationv1.ListFeaturesOptions{
					EnvironmentID: core.StringPtr("environment_id"),
					Expand: core.BoolPtr(true),
					Sort: core.StringPtr("created_time"),
					Tags: core.StringPtr("version 1.1,pre-release"),
					Collections: []string{"my-collection-id", "ghzindiapvtltd"},
					Segments: []string{"my-segment-id", "beta-users"},
					Include: []string{"collections", "rules", "change_request"},
					Limit: core.Int64Ptr(int64(10)),
					Search: core.StringPtr("test tag"),
				}

				pager, err := appConfigurationService.NewFeaturesPager(listFeaturesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateFeature(createFeatureOptions *CreateFeatureOptions) - Operation response error`, func() {
		createFeaturePath := "/environments/environment_id/features"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createFeaturePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateFeature with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}

				// Construct an instance of the FeatureSegmentRule model
				featureSegmentRuleModel := new(appconfigurationv1.FeatureSegmentRule)
				featureSegmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				featureSegmentRuleModel.Value = "true"
				featureSegmentRuleModel.Order = core.Int64Ptr(int64(1))
				featureSegmentRuleModel.RolloutPercentage = core.Int64Ptr(int64(50))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				collectionRefModel.CollectionID = core.StringPtr("ghzinc")

				// Construct an instance of the CreateFeatureOptions model
				createFeatureOptionsModel := new(appconfigurationv1.CreateFeatureOptions)
				createFeatureOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				createFeatureOptionsModel.Name = core.StringPtr("Cycle Rentals")
				createFeatureOptionsModel.FeatureID = core.StringPtr("cycle-rentals")
				createFeatureOptionsModel.Type = core.StringPtr("BOOLEAN")
				createFeatureOptionsModel.EnabledValue = "true"
				createFeatureOptionsModel.DisabledValue = "false"
				createFeatureOptionsModel.Description = core.StringPtr("Feature flag to enable Cycle Rentals")
				createFeatureOptionsModel.Format = core.StringPtr("TEXT")
				createFeatureOptionsModel.Enabled = core.BoolPtr(true)
				createFeatureOptionsModel.RolloutPercentage = core.Int64Ptr(int64(100))
				createFeatureOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				createFeatureOptionsModel.SegmentRules = []appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel}
				createFeatureOptionsModel.Collections = []appconfigurationv1.CollectionRef{*collectionRefModel}
				createFeatureOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.CreateFeature(createFeatureOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.CreateFeature(createFeatureOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateFeature(createFeatureOptions *CreateFeatureOptions)`, func() {
		createFeaturePath := "/environments/environment_id/features"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createFeaturePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Name", "feature_id": "FeatureID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "enabled_value": "anyValue", "disabled_value": "anyValue", "enabled": false, "rollout_percentage": 100, "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5, "rollout_percentage": 100}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "change_request_number": "ChangeRequestNumber", "change_request_status": "ChangeRequestStatus", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "evaluation_time": "2021-05-12T23:20:50.520Z", "href": "Href"}`)
				}))
			})
			It(`Invoke CreateFeature successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}

				// Construct an instance of the FeatureSegmentRule model
				featureSegmentRuleModel := new(appconfigurationv1.FeatureSegmentRule)
				featureSegmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				featureSegmentRuleModel.Value = "true"
				featureSegmentRuleModel.Order = core.Int64Ptr(int64(1))
				featureSegmentRuleModel.RolloutPercentage = core.Int64Ptr(int64(50))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				collectionRefModel.CollectionID = core.StringPtr("ghzinc")

				// Construct an instance of the CreateFeatureOptions model
				createFeatureOptionsModel := new(appconfigurationv1.CreateFeatureOptions)
				createFeatureOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				createFeatureOptionsModel.Name = core.StringPtr("Cycle Rentals")
				createFeatureOptionsModel.FeatureID = core.StringPtr("cycle-rentals")
				createFeatureOptionsModel.Type = core.StringPtr("BOOLEAN")
				createFeatureOptionsModel.EnabledValue = "true"
				createFeatureOptionsModel.DisabledValue = "false"
				createFeatureOptionsModel.Description = core.StringPtr("Feature flag to enable Cycle Rentals")
				createFeatureOptionsModel.Format = core.StringPtr("TEXT")
				createFeatureOptionsModel.Enabled = core.BoolPtr(true)
				createFeatureOptionsModel.RolloutPercentage = core.Int64Ptr(int64(100))
				createFeatureOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				createFeatureOptionsModel.SegmentRules = []appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel}
				createFeatureOptionsModel.Collections = []appconfigurationv1.CollectionRef{*collectionRefModel}
				createFeatureOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.CreateFeatureWithContext(ctx, createFeatureOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.CreateFeature(createFeatureOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.CreateFeatureWithContext(ctx, createFeatureOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createFeaturePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Name", "feature_id": "FeatureID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "enabled_value": "anyValue", "disabled_value": "anyValue", "enabled": false, "rollout_percentage": 100, "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5, "rollout_percentage": 100}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "change_request_number": "ChangeRequestNumber", "change_request_status": "ChangeRequestStatus", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "evaluation_time": "2021-05-12T23:20:50.520Z", "href": "Href"}`)
				}))
			})
			It(`Invoke CreateFeature successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.CreateFeature(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}

				// Construct an instance of the FeatureSegmentRule model
				featureSegmentRuleModel := new(appconfigurationv1.FeatureSegmentRule)
				featureSegmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				featureSegmentRuleModel.Value = "true"
				featureSegmentRuleModel.Order = core.Int64Ptr(int64(1))
				featureSegmentRuleModel.RolloutPercentage = core.Int64Ptr(int64(50))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				collectionRefModel.CollectionID = core.StringPtr("ghzinc")

				// Construct an instance of the CreateFeatureOptions model
				createFeatureOptionsModel := new(appconfigurationv1.CreateFeatureOptions)
				createFeatureOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				createFeatureOptionsModel.Name = core.StringPtr("Cycle Rentals")
				createFeatureOptionsModel.FeatureID = core.StringPtr("cycle-rentals")
				createFeatureOptionsModel.Type = core.StringPtr("BOOLEAN")
				createFeatureOptionsModel.EnabledValue = "true"
				createFeatureOptionsModel.DisabledValue = "false"
				createFeatureOptionsModel.Description = core.StringPtr("Feature flag to enable Cycle Rentals")
				createFeatureOptionsModel.Format = core.StringPtr("TEXT")
				createFeatureOptionsModel.Enabled = core.BoolPtr(true)
				createFeatureOptionsModel.RolloutPercentage = core.Int64Ptr(int64(100))
				createFeatureOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				createFeatureOptionsModel.SegmentRules = []appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel}
				createFeatureOptionsModel.Collections = []appconfigurationv1.CollectionRef{*collectionRefModel}
				createFeatureOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.CreateFeature(createFeatureOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateFeature with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}

				// Construct an instance of the FeatureSegmentRule model
				featureSegmentRuleModel := new(appconfigurationv1.FeatureSegmentRule)
				featureSegmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				featureSegmentRuleModel.Value = "true"
				featureSegmentRuleModel.Order = core.Int64Ptr(int64(1))
				featureSegmentRuleModel.RolloutPercentage = core.Int64Ptr(int64(50))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				collectionRefModel.CollectionID = core.StringPtr("ghzinc")

				// Construct an instance of the CreateFeatureOptions model
				createFeatureOptionsModel := new(appconfigurationv1.CreateFeatureOptions)
				createFeatureOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				createFeatureOptionsModel.Name = core.StringPtr("Cycle Rentals")
				createFeatureOptionsModel.FeatureID = core.StringPtr("cycle-rentals")
				createFeatureOptionsModel.Type = core.StringPtr("BOOLEAN")
				createFeatureOptionsModel.EnabledValue = "true"
				createFeatureOptionsModel.DisabledValue = "false"
				createFeatureOptionsModel.Description = core.StringPtr("Feature flag to enable Cycle Rentals")
				createFeatureOptionsModel.Format = core.StringPtr("TEXT")
				createFeatureOptionsModel.Enabled = core.BoolPtr(true)
				createFeatureOptionsModel.RolloutPercentage = core.Int64Ptr(int64(100))
				createFeatureOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				createFeatureOptionsModel.SegmentRules = []appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel}
				createFeatureOptionsModel.Collections = []appconfigurationv1.CollectionRef{*collectionRefModel}
				createFeatureOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.CreateFeature(createFeatureOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateFeatureOptions model with no property values
				createFeatureOptionsModelNew := new(appconfigurationv1.CreateFeatureOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.CreateFeature(createFeatureOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateFeature successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}

				// Construct an instance of the FeatureSegmentRule model
				featureSegmentRuleModel := new(appconfigurationv1.FeatureSegmentRule)
				featureSegmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				featureSegmentRuleModel.Value = "true"
				featureSegmentRuleModel.Order = core.Int64Ptr(int64(1))
				featureSegmentRuleModel.RolloutPercentage = core.Int64Ptr(int64(50))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				collectionRefModel.CollectionID = core.StringPtr("ghzinc")

				// Construct an instance of the CreateFeatureOptions model
				createFeatureOptionsModel := new(appconfigurationv1.CreateFeatureOptions)
				createFeatureOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				createFeatureOptionsModel.Name = core.StringPtr("Cycle Rentals")
				createFeatureOptionsModel.FeatureID = core.StringPtr("cycle-rentals")
				createFeatureOptionsModel.Type = core.StringPtr("BOOLEAN")
				createFeatureOptionsModel.EnabledValue = "true"
				createFeatureOptionsModel.DisabledValue = "false"
				createFeatureOptionsModel.Description = core.StringPtr("Feature flag to enable Cycle Rentals")
				createFeatureOptionsModel.Format = core.StringPtr("TEXT")
				createFeatureOptionsModel.Enabled = core.BoolPtr(true)
				createFeatureOptionsModel.RolloutPercentage = core.Int64Ptr(int64(100))
				createFeatureOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				createFeatureOptionsModel.SegmentRules = []appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel}
				createFeatureOptionsModel.Collections = []appconfigurationv1.CollectionRef{*collectionRefModel}
				createFeatureOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.CreateFeature(createFeatureOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateFeature(updateFeatureOptions *UpdateFeatureOptions) - Operation response error`, func() {
		updateFeaturePath := "/environments/environment_id/features/feature_id"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateFeaturePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateFeature with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}

				// Construct an instance of the FeatureSegmentRule model
				featureSegmentRuleModel := new(appconfigurationv1.FeatureSegmentRule)
				featureSegmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				featureSegmentRuleModel.Value = "true"
				featureSegmentRuleModel.Order = core.Int64Ptr(int64(1))
				featureSegmentRuleModel.RolloutPercentage = core.Int64Ptr(int64(90))

				// Construct an instance of the CollectionUpdateRef model
				collectionUpdateRefModel := new(appconfigurationv1.CollectionUpdateRef)
				collectionUpdateRefModel.CollectionID = core.StringPtr("ghzinc")
				collectionUpdateRefModel.Deleted = core.BoolPtr(true)

				// Construct an instance of the UpdateFeatureOptions model
				updateFeatureOptionsModel := new(appconfigurationv1.UpdateFeatureOptions)
				updateFeatureOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				updateFeatureOptionsModel.FeatureID = core.StringPtr("feature_id")
				updateFeatureOptionsModel.Name = core.StringPtr("Cycle Rentals")
				updateFeatureOptionsModel.Description = core.StringPtr("Feature flags to enable Cycle Rentals")
				updateFeatureOptionsModel.EnabledValue = "true"
				updateFeatureOptionsModel.DisabledValue = "false"
				updateFeatureOptionsModel.Enabled = core.BoolPtr(true)
				updateFeatureOptionsModel.RolloutPercentage = core.Int64Ptr(int64(100))
				updateFeatureOptionsModel.Tags = core.StringPtr("version: 1.1, yet-to-release")
				updateFeatureOptionsModel.SegmentRules = []appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel}
				updateFeatureOptionsModel.Collections = []appconfigurationv1.CollectionUpdateRef{*collectionUpdateRefModel}
				updateFeatureOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.UpdateFeature(updateFeatureOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.UpdateFeature(updateFeatureOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateFeature(updateFeatureOptions *UpdateFeatureOptions)`, func() {
		updateFeaturePath := "/environments/environment_id/features/feature_id"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateFeaturePath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "feature_id": "FeatureID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "enabled_value": "anyValue", "disabled_value": "anyValue", "enabled": false, "rollout_percentage": 100, "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5, "rollout_percentage": 100}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "change_request_number": "ChangeRequestNumber", "change_request_status": "ChangeRequestStatus", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "evaluation_time": "2021-05-12T23:20:50.520Z", "href": "Href"}`)
				}))
			})
			It(`Invoke UpdateFeature successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}

				// Construct an instance of the FeatureSegmentRule model
				featureSegmentRuleModel := new(appconfigurationv1.FeatureSegmentRule)
				featureSegmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				featureSegmentRuleModel.Value = "true"
				featureSegmentRuleModel.Order = core.Int64Ptr(int64(1))
				featureSegmentRuleModel.RolloutPercentage = core.Int64Ptr(int64(90))

				// Construct an instance of the CollectionUpdateRef model
				collectionUpdateRefModel := new(appconfigurationv1.CollectionUpdateRef)
				collectionUpdateRefModel.CollectionID = core.StringPtr("ghzinc")
				collectionUpdateRefModel.Deleted = core.BoolPtr(true)

				// Construct an instance of the UpdateFeatureOptions model
				updateFeatureOptionsModel := new(appconfigurationv1.UpdateFeatureOptions)
				updateFeatureOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				updateFeatureOptionsModel.FeatureID = core.StringPtr("feature_id")
				updateFeatureOptionsModel.Name = core.StringPtr("Cycle Rentals")
				updateFeatureOptionsModel.Description = core.StringPtr("Feature flags to enable Cycle Rentals")
				updateFeatureOptionsModel.EnabledValue = "true"
				updateFeatureOptionsModel.DisabledValue = "false"
				updateFeatureOptionsModel.Enabled = core.BoolPtr(true)
				updateFeatureOptionsModel.RolloutPercentage = core.Int64Ptr(int64(100))
				updateFeatureOptionsModel.Tags = core.StringPtr("version: 1.1, yet-to-release")
				updateFeatureOptionsModel.SegmentRules = []appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel}
				updateFeatureOptionsModel.Collections = []appconfigurationv1.CollectionUpdateRef{*collectionUpdateRefModel}
				updateFeatureOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.UpdateFeatureWithContext(ctx, updateFeatureOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.UpdateFeature(updateFeatureOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.UpdateFeatureWithContext(ctx, updateFeatureOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateFeaturePath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "feature_id": "FeatureID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "enabled_value": "anyValue", "disabled_value": "anyValue", "enabled": false, "rollout_percentage": 100, "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5, "rollout_percentage": 100}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "change_request_number": "ChangeRequestNumber", "change_request_status": "ChangeRequestStatus", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "evaluation_time": "2021-05-12T23:20:50.520Z", "href": "Href"}`)
				}))
			})
			It(`Invoke UpdateFeature successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.UpdateFeature(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}

				// Construct an instance of the FeatureSegmentRule model
				featureSegmentRuleModel := new(appconfigurationv1.FeatureSegmentRule)
				featureSegmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				featureSegmentRuleModel.Value = "true"
				featureSegmentRuleModel.Order = core.Int64Ptr(int64(1))
				featureSegmentRuleModel.RolloutPercentage = core.Int64Ptr(int64(90))

				// Construct an instance of the CollectionUpdateRef model
				collectionUpdateRefModel := new(appconfigurationv1.CollectionUpdateRef)
				collectionUpdateRefModel.CollectionID = core.StringPtr("ghzinc")
				collectionUpdateRefModel.Deleted = core.BoolPtr(true)

				// Construct an instance of the UpdateFeatureOptions model
				updateFeatureOptionsModel := new(appconfigurationv1.UpdateFeatureOptions)
				updateFeatureOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				updateFeatureOptionsModel.FeatureID = core.StringPtr("feature_id")
				updateFeatureOptionsModel.Name = core.StringPtr("Cycle Rentals")
				updateFeatureOptionsModel.Description = core.StringPtr("Feature flags to enable Cycle Rentals")
				updateFeatureOptionsModel.EnabledValue = "true"
				updateFeatureOptionsModel.DisabledValue = "false"
				updateFeatureOptionsModel.Enabled = core.BoolPtr(true)
				updateFeatureOptionsModel.RolloutPercentage = core.Int64Ptr(int64(100))
				updateFeatureOptionsModel.Tags = core.StringPtr("version: 1.1, yet-to-release")
				updateFeatureOptionsModel.SegmentRules = []appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel}
				updateFeatureOptionsModel.Collections = []appconfigurationv1.CollectionUpdateRef{*collectionUpdateRefModel}
				updateFeatureOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.UpdateFeature(updateFeatureOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateFeature with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}

				// Construct an instance of the FeatureSegmentRule model
				featureSegmentRuleModel := new(appconfigurationv1.FeatureSegmentRule)
				featureSegmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				featureSegmentRuleModel.Value = "true"
				featureSegmentRuleModel.Order = core.Int64Ptr(int64(1))
				featureSegmentRuleModel.RolloutPercentage = core.Int64Ptr(int64(90))

				// Construct an instance of the CollectionUpdateRef model
				collectionUpdateRefModel := new(appconfigurationv1.CollectionUpdateRef)
				collectionUpdateRefModel.CollectionID = core.StringPtr("ghzinc")
				collectionUpdateRefModel.Deleted = core.BoolPtr(true)

				// Construct an instance of the UpdateFeatureOptions model
				updateFeatureOptionsModel := new(appconfigurationv1.UpdateFeatureOptions)
				updateFeatureOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				updateFeatureOptionsModel.FeatureID = core.StringPtr("feature_id")
				updateFeatureOptionsModel.Name = core.StringPtr("Cycle Rentals")
				updateFeatureOptionsModel.Description = core.StringPtr("Feature flags to enable Cycle Rentals")
				updateFeatureOptionsModel.EnabledValue = "true"
				updateFeatureOptionsModel.DisabledValue = "false"
				updateFeatureOptionsModel.Enabled = core.BoolPtr(true)
				updateFeatureOptionsModel.RolloutPercentage = core.Int64Ptr(int64(100))
				updateFeatureOptionsModel.Tags = core.StringPtr("version: 1.1, yet-to-release")
				updateFeatureOptionsModel.SegmentRules = []appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel}
				updateFeatureOptionsModel.Collections = []appconfigurationv1.CollectionUpdateRef{*collectionUpdateRefModel}
				updateFeatureOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.UpdateFeature(updateFeatureOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateFeatureOptions model with no property values
				updateFeatureOptionsModelNew := new(appconfigurationv1.UpdateFeatureOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.UpdateFeature(updateFeatureOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateFeature successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}

				// Construct an instance of the FeatureSegmentRule model
				featureSegmentRuleModel := new(appconfigurationv1.FeatureSegmentRule)
				featureSegmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				featureSegmentRuleModel.Value = "true"
				featureSegmentRuleModel.Order = core.Int64Ptr(int64(1))
				featureSegmentRuleModel.RolloutPercentage = core.Int64Ptr(int64(90))

				// Construct an instance of the CollectionUpdateRef model
				collectionUpdateRefModel := new(appconfigurationv1.CollectionUpdateRef)
				collectionUpdateRefModel.CollectionID = core.StringPtr("ghzinc")
				collectionUpdateRefModel.Deleted = core.BoolPtr(true)

				// Construct an instance of the UpdateFeatureOptions model
				updateFeatureOptionsModel := new(appconfigurationv1.UpdateFeatureOptions)
				updateFeatureOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				updateFeatureOptionsModel.FeatureID = core.StringPtr("feature_id")
				updateFeatureOptionsModel.Name = core.StringPtr("Cycle Rentals")
				updateFeatureOptionsModel.Description = core.StringPtr("Feature flags to enable Cycle Rentals")
				updateFeatureOptionsModel.EnabledValue = "true"
				updateFeatureOptionsModel.DisabledValue = "false"
				updateFeatureOptionsModel.Enabled = core.BoolPtr(true)
				updateFeatureOptionsModel.RolloutPercentage = core.Int64Ptr(int64(100))
				updateFeatureOptionsModel.Tags = core.StringPtr("version: 1.1, yet-to-release")
				updateFeatureOptionsModel.SegmentRules = []appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel}
				updateFeatureOptionsModel.Collections = []appconfigurationv1.CollectionUpdateRef{*collectionUpdateRefModel}
				updateFeatureOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.UpdateFeature(updateFeatureOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateFeatureValues(updateFeatureValuesOptions *UpdateFeatureValuesOptions) - Operation response error`, func() {
		updateFeatureValuesPath := "/environments/environment_id/features/feature_id"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateFeatureValuesPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateFeatureValues with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}

				// Construct an instance of the FeatureSegmentRule model
				featureSegmentRuleModel := new(appconfigurationv1.FeatureSegmentRule)
				featureSegmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				featureSegmentRuleModel.Value = "true"
				featureSegmentRuleModel.Order = core.Int64Ptr(int64(1))
				featureSegmentRuleModel.RolloutPercentage = core.Int64Ptr(int64(100))

				// Construct an instance of the UpdateFeatureValuesOptions model
				updateFeatureValuesOptionsModel := new(appconfigurationv1.UpdateFeatureValuesOptions)
				updateFeatureValuesOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				updateFeatureValuesOptionsModel.FeatureID = core.StringPtr("feature_id")
				updateFeatureValuesOptionsModel.Name = core.StringPtr("Cycle Rentals")
				updateFeatureValuesOptionsModel.Description = core.StringPtr("Feature flags to enable Cycle Rentals")
				updateFeatureValuesOptionsModel.Tags = core.StringPtr("version: 1.1, yet-to-release")
				updateFeatureValuesOptionsModel.EnabledValue = "true"
				updateFeatureValuesOptionsModel.DisabledValue = "false"
				updateFeatureValuesOptionsModel.RolloutPercentage = core.Int64Ptr(int64(100))
				updateFeatureValuesOptionsModel.SegmentRules = []appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel}
				updateFeatureValuesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.UpdateFeatureValues(updateFeatureValuesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.UpdateFeatureValues(updateFeatureValuesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateFeatureValues(updateFeatureValuesOptions *UpdateFeatureValuesOptions)`, func() {
		updateFeatureValuesPath := "/environments/environment_id/features/feature_id"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateFeatureValuesPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "feature_id": "FeatureID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "enabled_value": "anyValue", "disabled_value": "anyValue", "enabled": false, "rollout_percentage": 100, "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5, "rollout_percentage": 100}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "change_request_number": "ChangeRequestNumber", "change_request_status": "ChangeRequestStatus", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "evaluation_time": "2021-05-12T23:20:50.520Z", "href": "Href"}`)
				}))
			})
			It(`Invoke UpdateFeatureValues successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}

				// Construct an instance of the FeatureSegmentRule model
				featureSegmentRuleModel := new(appconfigurationv1.FeatureSegmentRule)
				featureSegmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				featureSegmentRuleModel.Value = "true"
				featureSegmentRuleModel.Order = core.Int64Ptr(int64(1))
				featureSegmentRuleModel.RolloutPercentage = core.Int64Ptr(int64(100))

				// Construct an instance of the UpdateFeatureValuesOptions model
				updateFeatureValuesOptionsModel := new(appconfigurationv1.UpdateFeatureValuesOptions)
				updateFeatureValuesOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				updateFeatureValuesOptionsModel.FeatureID = core.StringPtr("feature_id")
				updateFeatureValuesOptionsModel.Name = core.StringPtr("Cycle Rentals")
				updateFeatureValuesOptionsModel.Description = core.StringPtr("Feature flags to enable Cycle Rentals")
				updateFeatureValuesOptionsModel.Tags = core.StringPtr("version: 1.1, yet-to-release")
				updateFeatureValuesOptionsModel.EnabledValue = "true"
				updateFeatureValuesOptionsModel.DisabledValue = "false"
				updateFeatureValuesOptionsModel.RolloutPercentage = core.Int64Ptr(int64(100))
				updateFeatureValuesOptionsModel.SegmentRules = []appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel}
				updateFeatureValuesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.UpdateFeatureValuesWithContext(ctx, updateFeatureValuesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.UpdateFeatureValues(updateFeatureValuesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.UpdateFeatureValuesWithContext(ctx, updateFeatureValuesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateFeatureValuesPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "feature_id": "FeatureID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "enabled_value": "anyValue", "disabled_value": "anyValue", "enabled": false, "rollout_percentage": 100, "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5, "rollout_percentage": 100}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "change_request_number": "ChangeRequestNumber", "change_request_status": "ChangeRequestStatus", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "evaluation_time": "2021-05-12T23:20:50.520Z", "href": "Href"}`)
				}))
			})
			It(`Invoke UpdateFeatureValues successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.UpdateFeatureValues(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}

				// Construct an instance of the FeatureSegmentRule model
				featureSegmentRuleModel := new(appconfigurationv1.FeatureSegmentRule)
				featureSegmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				featureSegmentRuleModel.Value = "true"
				featureSegmentRuleModel.Order = core.Int64Ptr(int64(1))
				featureSegmentRuleModel.RolloutPercentage = core.Int64Ptr(int64(100))

				// Construct an instance of the UpdateFeatureValuesOptions model
				updateFeatureValuesOptionsModel := new(appconfigurationv1.UpdateFeatureValuesOptions)
				updateFeatureValuesOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				updateFeatureValuesOptionsModel.FeatureID = core.StringPtr("feature_id")
				updateFeatureValuesOptionsModel.Name = core.StringPtr("Cycle Rentals")
				updateFeatureValuesOptionsModel.Description = core.StringPtr("Feature flags to enable Cycle Rentals")
				updateFeatureValuesOptionsModel.Tags = core.StringPtr("version: 1.1, yet-to-release")
				updateFeatureValuesOptionsModel.EnabledValue = "true"
				updateFeatureValuesOptionsModel.DisabledValue = "false"
				updateFeatureValuesOptionsModel.RolloutPercentage = core.Int64Ptr(int64(100))
				updateFeatureValuesOptionsModel.SegmentRules = []appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel}
				updateFeatureValuesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.UpdateFeatureValues(updateFeatureValuesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateFeatureValues with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}

				// Construct an instance of the FeatureSegmentRule model
				featureSegmentRuleModel := new(appconfigurationv1.FeatureSegmentRule)
				featureSegmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				featureSegmentRuleModel.Value = "true"
				featureSegmentRuleModel.Order = core.Int64Ptr(int64(1))
				featureSegmentRuleModel.RolloutPercentage = core.Int64Ptr(int64(100))

				// Construct an instance of the UpdateFeatureValuesOptions model
				updateFeatureValuesOptionsModel := new(appconfigurationv1.UpdateFeatureValuesOptions)
				updateFeatureValuesOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				updateFeatureValuesOptionsModel.FeatureID = core.StringPtr("feature_id")
				updateFeatureValuesOptionsModel.Name = core.StringPtr("Cycle Rentals")
				updateFeatureValuesOptionsModel.Description = core.StringPtr("Feature flags to enable Cycle Rentals")
				updateFeatureValuesOptionsModel.Tags = core.StringPtr("version: 1.1, yet-to-release")
				updateFeatureValuesOptionsModel.EnabledValue = "true"
				updateFeatureValuesOptionsModel.DisabledValue = "false"
				updateFeatureValuesOptionsModel.RolloutPercentage = core.Int64Ptr(int64(100))
				updateFeatureValuesOptionsModel.SegmentRules = []appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel}
				updateFeatureValuesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.UpdateFeatureValues(updateFeatureValuesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateFeatureValuesOptions model with no property values
				updateFeatureValuesOptionsModelNew := new(appconfigurationv1.UpdateFeatureValuesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.UpdateFeatureValues(updateFeatureValuesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateFeatureValues successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}

				// Construct an instance of the FeatureSegmentRule model
				featureSegmentRuleModel := new(appconfigurationv1.FeatureSegmentRule)
				featureSegmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				featureSegmentRuleModel.Value = "true"
				featureSegmentRuleModel.Order = core.Int64Ptr(int64(1))
				featureSegmentRuleModel.RolloutPercentage = core.Int64Ptr(int64(100))

				// Construct an instance of the UpdateFeatureValuesOptions model
				updateFeatureValuesOptionsModel := new(appconfigurationv1.UpdateFeatureValuesOptions)
				updateFeatureValuesOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				updateFeatureValuesOptionsModel.FeatureID = core.StringPtr("feature_id")
				updateFeatureValuesOptionsModel.Name = core.StringPtr("Cycle Rentals")
				updateFeatureValuesOptionsModel.Description = core.StringPtr("Feature flags to enable Cycle Rentals")
				updateFeatureValuesOptionsModel.Tags = core.StringPtr("version: 1.1, yet-to-release")
				updateFeatureValuesOptionsModel.EnabledValue = "true"
				updateFeatureValuesOptionsModel.DisabledValue = "false"
				updateFeatureValuesOptionsModel.RolloutPercentage = core.Int64Ptr(int64(100))
				updateFeatureValuesOptionsModel.SegmentRules = []appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel}
				updateFeatureValuesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.UpdateFeatureValues(updateFeatureValuesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetFeature(getFeatureOptions *GetFeatureOptions) - Operation response error`, func() {
		getFeaturePath := "/environments/environment_id/features/feature_id"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getFeaturePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetFeature with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetFeatureOptions model
				getFeatureOptionsModel := new(appconfigurationv1.GetFeatureOptions)
				getFeatureOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				getFeatureOptionsModel.FeatureID = core.StringPtr("feature_id")
				getFeatureOptionsModel.Include = []string{"collections", "rules", "change_request"}
				getFeatureOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.GetFeature(getFeatureOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.GetFeature(getFeatureOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetFeature(getFeatureOptions *GetFeatureOptions)`, func() {
		getFeaturePath := "/environments/environment_id/features/feature_id"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getFeaturePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "feature_id": "FeatureID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "enabled_value": "anyValue", "disabled_value": "anyValue", "enabled": false, "rollout_percentage": 100, "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5, "rollout_percentage": 100}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "change_request_number": "ChangeRequestNumber", "change_request_status": "ChangeRequestStatus", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "evaluation_time": "2021-05-12T23:20:50.520Z", "href": "Href"}`)
				}))
			})
			It(`Invoke GetFeature successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the GetFeatureOptions model
				getFeatureOptionsModel := new(appconfigurationv1.GetFeatureOptions)
				getFeatureOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				getFeatureOptionsModel.FeatureID = core.StringPtr("feature_id")
				getFeatureOptionsModel.Include = []string{"collections", "rules", "change_request"}
				getFeatureOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.GetFeatureWithContext(ctx, getFeatureOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.GetFeature(getFeatureOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.GetFeatureWithContext(ctx, getFeatureOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getFeaturePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "feature_id": "FeatureID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "enabled_value": "anyValue", "disabled_value": "anyValue", "enabled": false, "rollout_percentage": 100, "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5, "rollout_percentage": 100}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "change_request_number": "ChangeRequestNumber", "change_request_status": "ChangeRequestStatus", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "evaluation_time": "2021-05-12T23:20:50.520Z", "href": "Href"}`)
				}))
			})
			It(`Invoke GetFeature successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.GetFeature(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetFeatureOptions model
				getFeatureOptionsModel := new(appconfigurationv1.GetFeatureOptions)
				getFeatureOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				getFeatureOptionsModel.FeatureID = core.StringPtr("feature_id")
				getFeatureOptionsModel.Include = []string{"collections", "rules", "change_request"}
				getFeatureOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.GetFeature(getFeatureOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetFeature with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetFeatureOptions model
				getFeatureOptionsModel := new(appconfigurationv1.GetFeatureOptions)
				getFeatureOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				getFeatureOptionsModel.FeatureID = core.StringPtr("feature_id")
				getFeatureOptionsModel.Include = []string{"collections", "rules", "change_request"}
				getFeatureOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.GetFeature(getFeatureOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetFeatureOptions model with no property values
				getFeatureOptionsModelNew := new(appconfigurationv1.GetFeatureOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.GetFeature(getFeatureOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetFeature successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetFeatureOptions model
				getFeatureOptionsModel := new(appconfigurationv1.GetFeatureOptions)
				getFeatureOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				getFeatureOptionsModel.FeatureID = core.StringPtr("feature_id")
				getFeatureOptionsModel.Include = []string{"collections", "rules", "change_request"}
				getFeatureOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.GetFeature(getFeatureOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteFeature(deleteFeatureOptions *DeleteFeatureOptions)`, func() {
		deleteFeaturePath := "/environments/environment_id/features/feature_id"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteFeaturePath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteFeature successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appConfigurationService.DeleteFeature(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteFeatureOptions model
				deleteFeatureOptionsModel := new(appconfigurationv1.DeleteFeatureOptions)
				deleteFeatureOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				deleteFeatureOptionsModel.FeatureID = core.StringPtr("feature_id")
				deleteFeatureOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appConfigurationService.DeleteFeature(deleteFeatureOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteFeature with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the DeleteFeatureOptions model
				deleteFeatureOptionsModel := new(appconfigurationv1.DeleteFeatureOptions)
				deleteFeatureOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				deleteFeatureOptionsModel.FeatureID = core.StringPtr("feature_id")
				deleteFeatureOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appConfigurationService.DeleteFeature(deleteFeatureOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteFeatureOptions model with no property values
				deleteFeatureOptionsModelNew := new(appconfigurationv1.DeleteFeatureOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appConfigurationService.DeleteFeature(deleteFeatureOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ToggleFeature(toggleFeatureOptions *ToggleFeatureOptions) - Operation response error`, func() {
		toggleFeaturePath := "/environments/environment_id/features/feature_id/toggle"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(toggleFeaturePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ToggleFeature with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ToggleFeatureOptions model
				toggleFeatureOptionsModel := new(appconfigurationv1.ToggleFeatureOptions)
				toggleFeatureOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				toggleFeatureOptionsModel.FeatureID = core.StringPtr("feature_id")
				toggleFeatureOptionsModel.Enabled = core.BoolPtr(true)
				toggleFeatureOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.ToggleFeature(toggleFeatureOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.ToggleFeature(toggleFeatureOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ToggleFeature(toggleFeatureOptions *ToggleFeatureOptions)`, func() {
		toggleFeaturePath := "/environments/environment_id/features/feature_id/toggle"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(toggleFeaturePath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "feature_id": "FeatureID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "enabled_value": "anyValue", "disabled_value": "anyValue", "enabled": false, "rollout_percentage": 100, "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5, "rollout_percentage": 100}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "change_request_number": "ChangeRequestNumber", "change_request_status": "ChangeRequestStatus", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "evaluation_time": "2021-05-12T23:20:50.520Z", "href": "Href"}`)
				}))
			})
			It(`Invoke ToggleFeature successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the ToggleFeatureOptions model
				toggleFeatureOptionsModel := new(appconfigurationv1.ToggleFeatureOptions)
				toggleFeatureOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				toggleFeatureOptionsModel.FeatureID = core.StringPtr("feature_id")
				toggleFeatureOptionsModel.Enabled = core.BoolPtr(true)
				toggleFeatureOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.ToggleFeatureWithContext(ctx, toggleFeatureOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.ToggleFeature(toggleFeatureOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.ToggleFeatureWithContext(ctx, toggleFeatureOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(toggleFeaturePath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "feature_id": "FeatureID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "enabled_value": "anyValue", "disabled_value": "anyValue", "enabled": false, "rollout_percentage": 100, "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5, "rollout_percentage": 100}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "change_request_number": "ChangeRequestNumber", "change_request_status": "ChangeRequestStatus", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "evaluation_time": "2021-05-12T23:20:50.520Z", "href": "Href"}`)
				}))
			})
			It(`Invoke ToggleFeature successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.ToggleFeature(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ToggleFeatureOptions model
				toggleFeatureOptionsModel := new(appconfigurationv1.ToggleFeatureOptions)
				toggleFeatureOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				toggleFeatureOptionsModel.FeatureID = core.StringPtr("feature_id")
				toggleFeatureOptionsModel.Enabled = core.BoolPtr(true)
				toggleFeatureOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.ToggleFeature(toggleFeatureOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ToggleFeature with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ToggleFeatureOptions model
				toggleFeatureOptionsModel := new(appconfigurationv1.ToggleFeatureOptions)
				toggleFeatureOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				toggleFeatureOptionsModel.FeatureID = core.StringPtr("feature_id")
				toggleFeatureOptionsModel.Enabled = core.BoolPtr(true)
				toggleFeatureOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.ToggleFeature(toggleFeatureOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ToggleFeatureOptions model with no property values
				toggleFeatureOptionsModelNew := new(appconfigurationv1.ToggleFeatureOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.ToggleFeature(toggleFeatureOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ToggleFeature successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ToggleFeatureOptions model
				toggleFeatureOptionsModel := new(appconfigurationv1.ToggleFeatureOptions)
				toggleFeatureOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				toggleFeatureOptionsModel.FeatureID = core.StringPtr("feature_id")
				toggleFeatureOptionsModel.Enabled = core.BoolPtr(true)
				toggleFeatureOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.ToggleFeature(toggleFeatureOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListProperties(listPropertiesOptions *ListPropertiesOptions) - Operation response error`, func() {
		listPropertiesPath := "/environments/environment_id/properties"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPropertiesPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for expand query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_time"}))
					Expect(req.URL.Query()["tags"]).To(Equal([]string{"version 1.1,pre-release"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["search"]).To(Equal([]string{"test tag"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListProperties with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListPropertiesOptions model
				listPropertiesOptionsModel := new(appconfigurationv1.ListPropertiesOptions)
				listPropertiesOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				listPropertiesOptionsModel.Expand = core.BoolPtr(true)
				listPropertiesOptionsModel.Sort = core.StringPtr("created_time")
				listPropertiesOptionsModel.Tags = core.StringPtr("version 1.1,pre-release")
				listPropertiesOptionsModel.Collections = []string{"my-collection-id", "ghzindiapvtltd"}
				listPropertiesOptionsModel.Segments = []string{"my-segment-id", "beta-users"}
				listPropertiesOptionsModel.Include = []string{"collections", "rules"}
				listPropertiesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listPropertiesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listPropertiesOptionsModel.Search = core.StringPtr("test tag")
				listPropertiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.ListProperties(listPropertiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.ListProperties(listPropertiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListProperties(listPropertiesOptions *ListPropertiesOptions)`, func() {
		listPropertiesPath := "/environments/environment_id/properties"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPropertiesPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for expand query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_time"}))
					Expect(req.URL.Query()["tags"]).To(Equal([]string{"version 1.1,pre-release"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["search"]).To(Equal([]string{"test tag"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"properties": [{"name": "Name", "property_id": "PropertyID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "value": "anyValue", "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "evaluation_time": "2021-05-12T23:20:50.520Z", "href": "Href"}], "limit": 10, "offset": 0, "total_count": 0, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}}`)
				}))
			})
			It(`Invoke ListProperties successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the ListPropertiesOptions model
				listPropertiesOptionsModel := new(appconfigurationv1.ListPropertiesOptions)
				listPropertiesOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				listPropertiesOptionsModel.Expand = core.BoolPtr(true)
				listPropertiesOptionsModel.Sort = core.StringPtr("created_time")
				listPropertiesOptionsModel.Tags = core.StringPtr("version 1.1,pre-release")
				listPropertiesOptionsModel.Collections = []string{"my-collection-id", "ghzindiapvtltd"}
				listPropertiesOptionsModel.Segments = []string{"my-segment-id", "beta-users"}
				listPropertiesOptionsModel.Include = []string{"collections", "rules"}
				listPropertiesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listPropertiesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listPropertiesOptionsModel.Search = core.StringPtr("test tag")
				listPropertiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.ListPropertiesWithContext(ctx, listPropertiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.ListProperties(listPropertiesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.ListPropertiesWithContext(ctx, listPropertiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPropertiesPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for expand query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_time"}))
					Expect(req.URL.Query()["tags"]).To(Equal([]string{"version 1.1,pre-release"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["search"]).To(Equal([]string{"test tag"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"properties": [{"name": "Name", "property_id": "PropertyID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "value": "anyValue", "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "evaluation_time": "2021-05-12T23:20:50.520Z", "href": "Href"}], "limit": 10, "offset": 0, "total_count": 0, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}}`)
				}))
			})
			It(`Invoke ListProperties successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.ListProperties(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListPropertiesOptions model
				listPropertiesOptionsModel := new(appconfigurationv1.ListPropertiesOptions)
				listPropertiesOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				listPropertiesOptionsModel.Expand = core.BoolPtr(true)
				listPropertiesOptionsModel.Sort = core.StringPtr("created_time")
				listPropertiesOptionsModel.Tags = core.StringPtr("version 1.1,pre-release")
				listPropertiesOptionsModel.Collections = []string{"my-collection-id", "ghzindiapvtltd"}
				listPropertiesOptionsModel.Segments = []string{"my-segment-id", "beta-users"}
				listPropertiesOptionsModel.Include = []string{"collections", "rules"}
				listPropertiesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listPropertiesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listPropertiesOptionsModel.Search = core.StringPtr("test tag")
				listPropertiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.ListProperties(listPropertiesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListProperties with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListPropertiesOptions model
				listPropertiesOptionsModel := new(appconfigurationv1.ListPropertiesOptions)
				listPropertiesOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				listPropertiesOptionsModel.Expand = core.BoolPtr(true)
				listPropertiesOptionsModel.Sort = core.StringPtr("created_time")
				listPropertiesOptionsModel.Tags = core.StringPtr("version 1.1,pre-release")
				listPropertiesOptionsModel.Collections = []string{"my-collection-id", "ghzindiapvtltd"}
				listPropertiesOptionsModel.Segments = []string{"my-segment-id", "beta-users"}
				listPropertiesOptionsModel.Include = []string{"collections", "rules"}
				listPropertiesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listPropertiesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listPropertiesOptionsModel.Search = core.StringPtr("test tag")
				listPropertiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.ListProperties(listPropertiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListPropertiesOptions model with no property values
				listPropertiesOptionsModelNew := new(appconfigurationv1.ListPropertiesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.ListProperties(listPropertiesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListProperties successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListPropertiesOptions model
				listPropertiesOptionsModel := new(appconfigurationv1.ListPropertiesOptions)
				listPropertiesOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				listPropertiesOptionsModel.Expand = core.BoolPtr(true)
				listPropertiesOptionsModel.Sort = core.StringPtr("created_time")
				listPropertiesOptionsModel.Tags = core.StringPtr("version 1.1,pre-release")
				listPropertiesOptionsModel.Collections = []string{"my-collection-id", "ghzindiapvtltd"}
				listPropertiesOptionsModel.Segments = []string{"my-segment-id", "beta-users"}
				listPropertiesOptionsModel.Include = []string{"collections", "rules"}
				listPropertiesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listPropertiesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listPropertiesOptionsModel.Search = core.StringPtr("test tag")
				listPropertiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.ListProperties(listPropertiesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(appconfigurationv1.PropertiesList)
				nextObject := new(appconfigurationv1.PaginatedListNext)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(appconfigurationv1.PropertiesList)

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(appconfigurationv1.PropertiesList)
				nextObject := new(appconfigurationv1.PaginatedListNext)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(appconfigurationv1.PropertiesList)
				nextObject := new(appconfigurationv1.PaginatedListNext)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPropertiesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"total_count":2,"limit":1,"properties":[{"name":"Name","property_id":"PropertyID","description":"Description","type":"BOOLEAN","format":"TEXT","value":"anyValue","tags":"Tags","segment_rules":[{"rules":[{"segments":["Segments"]}],"value":"anyValue","order":5}],"segment_exists":false,"collections":[{"collection_id":"CollectionID","name":"Name"}],"created_time":"2021-05-12T23:20:50.520Z","updated_time":"2021-05-12T23:20:50.520Z","evaluation_time":"2021-05-12T23:20:50.520Z","href":"Href"}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"properties":[{"name":"Name","property_id":"PropertyID","description":"Description","type":"BOOLEAN","format":"TEXT","value":"anyValue","tags":"Tags","segment_rules":[{"rules":[{"segments":["Segments"]}],"value":"anyValue","order":5}],"segment_exists":false,"collections":[{"collection_id":"CollectionID","name":"Name"}],"created_time":"2021-05-12T23:20:50.520Z","updated_time":"2021-05-12T23:20:50.520Z","evaluation_time":"2021-05-12T23:20:50.520Z","href":"Href"}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use PropertiesPager.GetNext successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				listPropertiesOptionsModel := &appconfigurationv1.ListPropertiesOptions{
					EnvironmentID: core.StringPtr("environment_id"),
					Expand: core.BoolPtr(true),
					Sort: core.StringPtr("created_time"),
					Tags: core.StringPtr("version 1.1,pre-release"),
					Collections: []string{"my-collection-id", "ghzindiapvtltd"},
					Segments: []string{"my-segment-id", "beta-users"},
					Include: []string{"collections", "rules"},
					Limit: core.Int64Ptr(int64(10)),
					Search: core.StringPtr("test tag"),
				}

				pager, err := appConfigurationService.NewPropertiesPager(listPropertiesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []appconfigurationv1.Property
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use PropertiesPager.GetAll successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				listPropertiesOptionsModel := &appconfigurationv1.ListPropertiesOptions{
					EnvironmentID: core.StringPtr("environment_id"),
					Expand: core.BoolPtr(true),
					Sort: core.StringPtr("created_time"),
					Tags: core.StringPtr("version 1.1,pre-release"),
					Collections: []string{"my-collection-id", "ghzindiapvtltd"},
					Segments: []string{"my-segment-id", "beta-users"},
					Include: []string{"collections", "rules"},
					Limit: core.Int64Ptr(int64(10)),
					Search: core.StringPtr("test tag"),
				}

				pager, err := appConfigurationService.NewPropertiesPager(listPropertiesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateProperty(createPropertyOptions *CreatePropertyOptions) - Operation response error`, func() {
		createPropertyPath := "/environments/environment_id/properties"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createPropertyPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateProperty with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = "true"
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				collectionRefModel.CollectionID = core.StringPtr("ghzinc")

				// Construct an instance of the CreatePropertyOptions model
				createPropertyOptionsModel := new(appconfigurationv1.CreatePropertyOptions)
				createPropertyOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				createPropertyOptionsModel.Name = core.StringPtr("Email property")
				createPropertyOptionsModel.PropertyID = core.StringPtr("email-property")
				createPropertyOptionsModel.Type = core.StringPtr("BOOLEAN")
				createPropertyOptionsModel.Value = "true"
				createPropertyOptionsModel.Description = core.StringPtr("Property for email")
				createPropertyOptionsModel.Format = core.StringPtr("TEXT")
				createPropertyOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				createPropertyOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				createPropertyOptionsModel.Collections = []appconfigurationv1.CollectionRef{*collectionRefModel}
				createPropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.CreateProperty(createPropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.CreateProperty(createPropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateProperty(createPropertyOptions *CreatePropertyOptions)`, func() {
		createPropertyPath := "/environments/environment_id/properties"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createPropertyPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Name", "property_id": "PropertyID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "value": "anyValue", "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "evaluation_time": "2021-05-12T23:20:50.520Z", "href": "Href"}`)
				}))
			})
			It(`Invoke CreateProperty successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = "true"
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				collectionRefModel.CollectionID = core.StringPtr("ghzinc")

				// Construct an instance of the CreatePropertyOptions model
				createPropertyOptionsModel := new(appconfigurationv1.CreatePropertyOptions)
				createPropertyOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				createPropertyOptionsModel.Name = core.StringPtr("Email property")
				createPropertyOptionsModel.PropertyID = core.StringPtr("email-property")
				createPropertyOptionsModel.Type = core.StringPtr("BOOLEAN")
				createPropertyOptionsModel.Value = "true"
				createPropertyOptionsModel.Description = core.StringPtr("Property for email")
				createPropertyOptionsModel.Format = core.StringPtr("TEXT")
				createPropertyOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				createPropertyOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				createPropertyOptionsModel.Collections = []appconfigurationv1.CollectionRef{*collectionRefModel}
				createPropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.CreatePropertyWithContext(ctx, createPropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.CreateProperty(createPropertyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.CreatePropertyWithContext(ctx, createPropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createPropertyPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Name", "property_id": "PropertyID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "value": "anyValue", "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "evaluation_time": "2021-05-12T23:20:50.520Z", "href": "Href"}`)
				}))
			})
			It(`Invoke CreateProperty successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.CreateProperty(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = "true"
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				collectionRefModel.CollectionID = core.StringPtr("ghzinc")

				// Construct an instance of the CreatePropertyOptions model
				createPropertyOptionsModel := new(appconfigurationv1.CreatePropertyOptions)
				createPropertyOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				createPropertyOptionsModel.Name = core.StringPtr("Email property")
				createPropertyOptionsModel.PropertyID = core.StringPtr("email-property")
				createPropertyOptionsModel.Type = core.StringPtr("BOOLEAN")
				createPropertyOptionsModel.Value = "true"
				createPropertyOptionsModel.Description = core.StringPtr("Property for email")
				createPropertyOptionsModel.Format = core.StringPtr("TEXT")
				createPropertyOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				createPropertyOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				createPropertyOptionsModel.Collections = []appconfigurationv1.CollectionRef{*collectionRefModel}
				createPropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.CreateProperty(createPropertyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateProperty with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = "true"
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				collectionRefModel.CollectionID = core.StringPtr("ghzinc")

				// Construct an instance of the CreatePropertyOptions model
				createPropertyOptionsModel := new(appconfigurationv1.CreatePropertyOptions)
				createPropertyOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				createPropertyOptionsModel.Name = core.StringPtr("Email property")
				createPropertyOptionsModel.PropertyID = core.StringPtr("email-property")
				createPropertyOptionsModel.Type = core.StringPtr("BOOLEAN")
				createPropertyOptionsModel.Value = "true"
				createPropertyOptionsModel.Description = core.StringPtr("Property for email")
				createPropertyOptionsModel.Format = core.StringPtr("TEXT")
				createPropertyOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				createPropertyOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				createPropertyOptionsModel.Collections = []appconfigurationv1.CollectionRef{*collectionRefModel}
				createPropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.CreateProperty(createPropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreatePropertyOptions model with no property values
				createPropertyOptionsModelNew := new(appconfigurationv1.CreatePropertyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.CreateProperty(createPropertyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateProperty successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = "true"
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				collectionRefModel.CollectionID = core.StringPtr("ghzinc")

				// Construct an instance of the CreatePropertyOptions model
				createPropertyOptionsModel := new(appconfigurationv1.CreatePropertyOptions)
				createPropertyOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				createPropertyOptionsModel.Name = core.StringPtr("Email property")
				createPropertyOptionsModel.PropertyID = core.StringPtr("email-property")
				createPropertyOptionsModel.Type = core.StringPtr("BOOLEAN")
				createPropertyOptionsModel.Value = "true"
				createPropertyOptionsModel.Description = core.StringPtr("Property for email")
				createPropertyOptionsModel.Format = core.StringPtr("TEXT")
				createPropertyOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				createPropertyOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				createPropertyOptionsModel.Collections = []appconfigurationv1.CollectionRef{*collectionRefModel}
				createPropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.CreateProperty(createPropertyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateProperty(updatePropertyOptions *UpdatePropertyOptions) - Operation response error`, func() {
		updatePropertyPath := "/environments/environment_id/properties/property_id"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updatePropertyPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateProperty with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = "true"
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the CollectionUpdateRef model
				collectionUpdateRefModel := new(appconfigurationv1.CollectionUpdateRef)
				collectionUpdateRefModel.CollectionID = core.StringPtr("ghzinc")
				collectionUpdateRefModel.Deleted = core.BoolPtr(true)

				// Construct an instance of the UpdatePropertyOptions model
				updatePropertyOptionsModel := new(appconfigurationv1.UpdatePropertyOptions)
				updatePropertyOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				updatePropertyOptionsModel.PropertyID = core.StringPtr("property_id")
				updatePropertyOptionsModel.Name = core.StringPtr("Email property")
				updatePropertyOptionsModel.Description = core.StringPtr("Property for email")
				updatePropertyOptionsModel.Value = "true"
				updatePropertyOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				updatePropertyOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				updatePropertyOptionsModel.Collections = []appconfigurationv1.CollectionUpdateRef{*collectionUpdateRefModel}
				updatePropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.UpdateProperty(updatePropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.UpdateProperty(updatePropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateProperty(updatePropertyOptions *UpdatePropertyOptions)`, func() {
		updatePropertyPath := "/environments/environment_id/properties/property_id"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updatePropertyPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "property_id": "PropertyID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "value": "anyValue", "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "evaluation_time": "2021-05-12T23:20:50.520Z", "href": "Href"}`)
				}))
			})
			It(`Invoke UpdateProperty successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = "true"
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the CollectionUpdateRef model
				collectionUpdateRefModel := new(appconfigurationv1.CollectionUpdateRef)
				collectionUpdateRefModel.CollectionID = core.StringPtr("ghzinc")
				collectionUpdateRefModel.Deleted = core.BoolPtr(true)

				// Construct an instance of the UpdatePropertyOptions model
				updatePropertyOptionsModel := new(appconfigurationv1.UpdatePropertyOptions)
				updatePropertyOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				updatePropertyOptionsModel.PropertyID = core.StringPtr("property_id")
				updatePropertyOptionsModel.Name = core.StringPtr("Email property")
				updatePropertyOptionsModel.Description = core.StringPtr("Property for email")
				updatePropertyOptionsModel.Value = "true"
				updatePropertyOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				updatePropertyOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				updatePropertyOptionsModel.Collections = []appconfigurationv1.CollectionUpdateRef{*collectionUpdateRefModel}
				updatePropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.UpdatePropertyWithContext(ctx, updatePropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.UpdateProperty(updatePropertyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.UpdatePropertyWithContext(ctx, updatePropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updatePropertyPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "property_id": "PropertyID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "value": "anyValue", "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "evaluation_time": "2021-05-12T23:20:50.520Z", "href": "Href"}`)
				}))
			})
			It(`Invoke UpdateProperty successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.UpdateProperty(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = "true"
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the CollectionUpdateRef model
				collectionUpdateRefModel := new(appconfigurationv1.CollectionUpdateRef)
				collectionUpdateRefModel.CollectionID = core.StringPtr("ghzinc")
				collectionUpdateRefModel.Deleted = core.BoolPtr(true)

				// Construct an instance of the UpdatePropertyOptions model
				updatePropertyOptionsModel := new(appconfigurationv1.UpdatePropertyOptions)
				updatePropertyOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				updatePropertyOptionsModel.PropertyID = core.StringPtr("property_id")
				updatePropertyOptionsModel.Name = core.StringPtr("Email property")
				updatePropertyOptionsModel.Description = core.StringPtr("Property for email")
				updatePropertyOptionsModel.Value = "true"
				updatePropertyOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				updatePropertyOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				updatePropertyOptionsModel.Collections = []appconfigurationv1.CollectionUpdateRef{*collectionUpdateRefModel}
				updatePropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.UpdateProperty(updatePropertyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateProperty with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = "true"
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the CollectionUpdateRef model
				collectionUpdateRefModel := new(appconfigurationv1.CollectionUpdateRef)
				collectionUpdateRefModel.CollectionID = core.StringPtr("ghzinc")
				collectionUpdateRefModel.Deleted = core.BoolPtr(true)

				// Construct an instance of the UpdatePropertyOptions model
				updatePropertyOptionsModel := new(appconfigurationv1.UpdatePropertyOptions)
				updatePropertyOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				updatePropertyOptionsModel.PropertyID = core.StringPtr("property_id")
				updatePropertyOptionsModel.Name = core.StringPtr("Email property")
				updatePropertyOptionsModel.Description = core.StringPtr("Property for email")
				updatePropertyOptionsModel.Value = "true"
				updatePropertyOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				updatePropertyOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				updatePropertyOptionsModel.Collections = []appconfigurationv1.CollectionUpdateRef{*collectionUpdateRefModel}
				updatePropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.UpdateProperty(updatePropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdatePropertyOptions model with no property values
				updatePropertyOptionsModelNew := new(appconfigurationv1.UpdatePropertyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.UpdateProperty(updatePropertyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateProperty successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = "true"
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the CollectionUpdateRef model
				collectionUpdateRefModel := new(appconfigurationv1.CollectionUpdateRef)
				collectionUpdateRefModel.CollectionID = core.StringPtr("ghzinc")
				collectionUpdateRefModel.Deleted = core.BoolPtr(true)

				// Construct an instance of the UpdatePropertyOptions model
				updatePropertyOptionsModel := new(appconfigurationv1.UpdatePropertyOptions)
				updatePropertyOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				updatePropertyOptionsModel.PropertyID = core.StringPtr("property_id")
				updatePropertyOptionsModel.Name = core.StringPtr("Email property")
				updatePropertyOptionsModel.Description = core.StringPtr("Property for email")
				updatePropertyOptionsModel.Value = "true"
				updatePropertyOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				updatePropertyOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				updatePropertyOptionsModel.Collections = []appconfigurationv1.CollectionUpdateRef{*collectionUpdateRefModel}
				updatePropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.UpdateProperty(updatePropertyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdatePropertyValues(updatePropertyValuesOptions *UpdatePropertyValuesOptions) - Operation response error`, func() {
		updatePropertyValuesPath := "/environments/environment_id/properties/property_id"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updatePropertyValuesPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdatePropertyValues with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = "true"
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the UpdatePropertyValuesOptions model
				updatePropertyValuesOptionsModel := new(appconfigurationv1.UpdatePropertyValuesOptions)
				updatePropertyValuesOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				updatePropertyValuesOptionsModel.PropertyID = core.StringPtr("property_id")
				updatePropertyValuesOptionsModel.Name = core.StringPtr("Email property")
				updatePropertyValuesOptionsModel.Description = core.StringPtr("Property for email")
				updatePropertyValuesOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				updatePropertyValuesOptionsModel.Value = "true"
				updatePropertyValuesOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				updatePropertyValuesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.UpdatePropertyValues(updatePropertyValuesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.UpdatePropertyValues(updatePropertyValuesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdatePropertyValues(updatePropertyValuesOptions *UpdatePropertyValuesOptions)`, func() {
		updatePropertyValuesPath := "/environments/environment_id/properties/property_id"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updatePropertyValuesPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "property_id": "PropertyID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "value": "anyValue", "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "evaluation_time": "2021-05-12T23:20:50.520Z", "href": "Href"}`)
				}))
			})
			It(`Invoke UpdatePropertyValues successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = "true"
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the UpdatePropertyValuesOptions model
				updatePropertyValuesOptionsModel := new(appconfigurationv1.UpdatePropertyValuesOptions)
				updatePropertyValuesOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				updatePropertyValuesOptionsModel.PropertyID = core.StringPtr("property_id")
				updatePropertyValuesOptionsModel.Name = core.StringPtr("Email property")
				updatePropertyValuesOptionsModel.Description = core.StringPtr("Property for email")
				updatePropertyValuesOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				updatePropertyValuesOptionsModel.Value = "true"
				updatePropertyValuesOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				updatePropertyValuesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.UpdatePropertyValuesWithContext(ctx, updatePropertyValuesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.UpdatePropertyValues(updatePropertyValuesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.UpdatePropertyValuesWithContext(ctx, updatePropertyValuesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updatePropertyValuesPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "property_id": "PropertyID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "value": "anyValue", "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "evaluation_time": "2021-05-12T23:20:50.520Z", "href": "Href"}`)
				}))
			})
			It(`Invoke UpdatePropertyValues successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.UpdatePropertyValues(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = "true"
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the UpdatePropertyValuesOptions model
				updatePropertyValuesOptionsModel := new(appconfigurationv1.UpdatePropertyValuesOptions)
				updatePropertyValuesOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				updatePropertyValuesOptionsModel.PropertyID = core.StringPtr("property_id")
				updatePropertyValuesOptionsModel.Name = core.StringPtr("Email property")
				updatePropertyValuesOptionsModel.Description = core.StringPtr("Property for email")
				updatePropertyValuesOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				updatePropertyValuesOptionsModel.Value = "true"
				updatePropertyValuesOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				updatePropertyValuesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.UpdatePropertyValues(updatePropertyValuesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdatePropertyValues with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = "true"
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the UpdatePropertyValuesOptions model
				updatePropertyValuesOptionsModel := new(appconfigurationv1.UpdatePropertyValuesOptions)
				updatePropertyValuesOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				updatePropertyValuesOptionsModel.PropertyID = core.StringPtr("property_id")
				updatePropertyValuesOptionsModel.Name = core.StringPtr("Email property")
				updatePropertyValuesOptionsModel.Description = core.StringPtr("Property for email")
				updatePropertyValuesOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				updatePropertyValuesOptionsModel.Value = "true"
				updatePropertyValuesOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				updatePropertyValuesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.UpdatePropertyValues(updatePropertyValuesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdatePropertyValuesOptions model with no property values
				updatePropertyValuesOptionsModelNew := new(appconfigurationv1.UpdatePropertyValuesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.UpdatePropertyValues(updatePropertyValuesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdatePropertyValues successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = "true"
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the UpdatePropertyValuesOptions model
				updatePropertyValuesOptionsModel := new(appconfigurationv1.UpdatePropertyValuesOptions)
				updatePropertyValuesOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				updatePropertyValuesOptionsModel.PropertyID = core.StringPtr("property_id")
				updatePropertyValuesOptionsModel.Name = core.StringPtr("Email property")
				updatePropertyValuesOptionsModel.Description = core.StringPtr("Property for email")
				updatePropertyValuesOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				updatePropertyValuesOptionsModel.Value = "true"
				updatePropertyValuesOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				updatePropertyValuesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.UpdatePropertyValues(updatePropertyValuesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProperty(getPropertyOptions *GetPropertyOptions) - Operation response error`, func() {
		getPropertyPath := "/environments/environment_id/properties/property_id"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPropertyPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProperty with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetPropertyOptions model
				getPropertyOptionsModel := new(appconfigurationv1.GetPropertyOptions)
				getPropertyOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				getPropertyOptionsModel.PropertyID = core.StringPtr("property_id")
				getPropertyOptionsModel.Include = []string{"collections", "rules"}
				getPropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.GetProperty(getPropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.GetProperty(getPropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProperty(getPropertyOptions *GetPropertyOptions)`, func() {
		getPropertyPath := "/environments/environment_id/properties/property_id"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPropertyPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "property_id": "PropertyID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "value": "anyValue", "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "evaluation_time": "2021-05-12T23:20:50.520Z", "href": "Href"}`)
				}))
			})
			It(`Invoke GetProperty successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the GetPropertyOptions model
				getPropertyOptionsModel := new(appconfigurationv1.GetPropertyOptions)
				getPropertyOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				getPropertyOptionsModel.PropertyID = core.StringPtr("property_id")
				getPropertyOptionsModel.Include = []string{"collections", "rules"}
				getPropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.GetPropertyWithContext(ctx, getPropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.GetProperty(getPropertyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.GetPropertyWithContext(ctx, getPropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPropertyPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "property_id": "PropertyID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "value": "anyValue", "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "evaluation_time": "2021-05-12T23:20:50.520Z", "href": "Href"}`)
				}))
			})
			It(`Invoke GetProperty successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.GetProperty(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPropertyOptions model
				getPropertyOptionsModel := new(appconfigurationv1.GetPropertyOptions)
				getPropertyOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				getPropertyOptionsModel.PropertyID = core.StringPtr("property_id")
				getPropertyOptionsModel.Include = []string{"collections", "rules"}
				getPropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.GetProperty(getPropertyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProperty with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetPropertyOptions model
				getPropertyOptionsModel := new(appconfigurationv1.GetPropertyOptions)
				getPropertyOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				getPropertyOptionsModel.PropertyID = core.StringPtr("property_id")
				getPropertyOptionsModel.Include = []string{"collections", "rules"}
				getPropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.GetProperty(getPropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPropertyOptions model with no property values
				getPropertyOptionsModelNew := new(appconfigurationv1.GetPropertyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.GetProperty(getPropertyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetProperty successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetPropertyOptions model
				getPropertyOptionsModel := new(appconfigurationv1.GetPropertyOptions)
				getPropertyOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				getPropertyOptionsModel.PropertyID = core.StringPtr("property_id")
				getPropertyOptionsModel.Include = []string{"collections", "rules"}
				getPropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.GetProperty(getPropertyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteProperty(deletePropertyOptions *DeletePropertyOptions)`, func() {
		deletePropertyPath := "/environments/environment_id/properties/property_id"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deletePropertyPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteProperty successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appConfigurationService.DeleteProperty(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeletePropertyOptions model
				deletePropertyOptionsModel := new(appconfigurationv1.DeletePropertyOptions)
				deletePropertyOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				deletePropertyOptionsModel.PropertyID = core.StringPtr("property_id")
				deletePropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appConfigurationService.DeleteProperty(deletePropertyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteProperty with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the DeletePropertyOptions model
				deletePropertyOptionsModel := new(appconfigurationv1.DeletePropertyOptions)
				deletePropertyOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				deletePropertyOptionsModel.PropertyID = core.StringPtr("property_id")
				deletePropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appConfigurationService.DeleteProperty(deletePropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeletePropertyOptions model with no property values
				deletePropertyOptionsModelNew := new(appconfigurationv1.DeletePropertyOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appConfigurationService.DeleteProperty(deletePropertyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListSegments(listSegmentsOptions *ListSegmentsOptions) - Operation response error`, func() {
		listSegmentsPath := "/segments"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSegmentsPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for expand query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_time"}))
					Expect(req.URL.Query()["tags"]).To(Equal([]string{"version 1.1,pre-release"}))
					Expect(req.URL.Query()["include"]).To(Equal([]string{"rules"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["search"]).To(Equal([]string{"test tag"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListSegments with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListSegmentsOptions model
				listSegmentsOptionsModel := new(appconfigurationv1.ListSegmentsOptions)
				listSegmentsOptionsModel.Expand = core.BoolPtr(true)
				listSegmentsOptionsModel.Sort = core.StringPtr("created_time")
				listSegmentsOptionsModel.Tags = core.StringPtr("version 1.1,pre-release")
				listSegmentsOptionsModel.Include = core.StringPtr("rules")
				listSegmentsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listSegmentsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listSegmentsOptionsModel.Search = core.StringPtr("test tag")
				listSegmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.ListSegments(listSegmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.ListSegments(listSegmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListSegments(listSegmentsOptions *ListSegmentsOptions)`, func() {
		listSegmentsPath := "/segments"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSegmentsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for expand query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_time"}))
					Expect(req.URL.Query()["tags"]).To(Equal([]string{"version 1.1,pre-release"}))
					Expect(req.URL.Query()["include"]).To(Equal([]string{"rules"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["search"]).To(Equal([]string{"test tag"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"segments": [{"name": "Name", "segment_id": "SegmentID", "description": "Description", "tags": "Tags", "rules": [{"attribute_name": "AttributeName", "operator": "is", "values": ["Values"]}], "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href", "features": [{"feature_id": "FeatureID", "name": "Name"}], "properties": [{"property_id": "PropertyID", "name": "Name"}]}], "limit": 10, "offset": 0, "total_count": 0, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}}`)
				}))
			})
			It(`Invoke ListSegments successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the ListSegmentsOptions model
				listSegmentsOptionsModel := new(appconfigurationv1.ListSegmentsOptions)
				listSegmentsOptionsModel.Expand = core.BoolPtr(true)
				listSegmentsOptionsModel.Sort = core.StringPtr("created_time")
				listSegmentsOptionsModel.Tags = core.StringPtr("version 1.1,pre-release")
				listSegmentsOptionsModel.Include = core.StringPtr("rules")
				listSegmentsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listSegmentsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listSegmentsOptionsModel.Search = core.StringPtr("test tag")
				listSegmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.ListSegmentsWithContext(ctx, listSegmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.ListSegments(listSegmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.ListSegmentsWithContext(ctx, listSegmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSegmentsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for expand query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_time"}))
					Expect(req.URL.Query()["tags"]).To(Equal([]string{"version 1.1,pre-release"}))
					Expect(req.URL.Query()["include"]).To(Equal([]string{"rules"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["search"]).To(Equal([]string{"test tag"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"segments": [{"name": "Name", "segment_id": "SegmentID", "description": "Description", "tags": "Tags", "rules": [{"attribute_name": "AttributeName", "operator": "is", "values": ["Values"]}], "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href", "features": [{"feature_id": "FeatureID", "name": "Name"}], "properties": [{"property_id": "PropertyID", "name": "Name"}]}], "limit": 10, "offset": 0, "total_count": 0, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}}`)
				}))
			})
			It(`Invoke ListSegments successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.ListSegments(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListSegmentsOptions model
				listSegmentsOptionsModel := new(appconfigurationv1.ListSegmentsOptions)
				listSegmentsOptionsModel.Expand = core.BoolPtr(true)
				listSegmentsOptionsModel.Sort = core.StringPtr("created_time")
				listSegmentsOptionsModel.Tags = core.StringPtr("version 1.1,pre-release")
				listSegmentsOptionsModel.Include = core.StringPtr("rules")
				listSegmentsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listSegmentsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listSegmentsOptionsModel.Search = core.StringPtr("test tag")
				listSegmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.ListSegments(listSegmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListSegments with error: Operation request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListSegmentsOptions model
				listSegmentsOptionsModel := new(appconfigurationv1.ListSegmentsOptions)
				listSegmentsOptionsModel.Expand = core.BoolPtr(true)
				listSegmentsOptionsModel.Sort = core.StringPtr("created_time")
				listSegmentsOptionsModel.Tags = core.StringPtr("version 1.1,pre-release")
				listSegmentsOptionsModel.Include = core.StringPtr("rules")
				listSegmentsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listSegmentsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listSegmentsOptionsModel.Search = core.StringPtr("test tag")
				listSegmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.ListSegments(listSegmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListSegments successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListSegmentsOptions model
				listSegmentsOptionsModel := new(appconfigurationv1.ListSegmentsOptions)
				listSegmentsOptionsModel.Expand = core.BoolPtr(true)
				listSegmentsOptionsModel.Sort = core.StringPtr("created_time")
				listSegmentsOptionsModel.Tags = core.StringPtr("version 1.1,pre-release")
				listSegmentsOptionsModel.Include = core.StringPtr("rules")
				listSegmentsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listSegmentsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listSegmentsOptionsModel.Search = core.StringPtr("test tag")
				listSegmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.ListSegments(listSegmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(appconfigurationv1.SegmentsList)
				nextObject := new(appconfigurationv1.PaginatedListNext)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(appconfigurationv1.SegmentsList)

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(appconfigurationv1.SegmentsList)
				nextObject := new(appconfigurationv1.PaginatedListNext)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(appconfigurationv1.SegmentsList)
				nextObject := new(appconfigurationv1.PaginatedListNext)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSegmentsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"total_count":2,"limit":1,"segments":[{"name":"Name","segment_id":"SegmentID","description":"Description","tags":"Tags","rules":[{"attribute_name":"AttributeName","operator":"is","values":["Values"]}],"created_time":"2021-05-12T23:20:50.520Z","updated_time":"2021-05-12T23:20:50.520Z","href":"Href","features":[{"feature_id":"FeatureID","name":"Name"}],"properties":[{"property_id":"PropertyID","name":"Name"}]}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"segments":[{"name":"Name","segment_id":"SegmentID","description":"Description","tags":"Tags","rules":[{"attribute_name":"AttributeName","operator":"is","values":["Values"]}],"created_time":"2021-05-12T23:20:50.520Z","updated_time":"2021-05-12T23:20:50.520Z","href":"Href","features":[{"feature_id":"FeatureID","name":"Name"}],"properties":[{"property_id":"PropertyID","name":"Name"}]}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use SegmentsPager.GetNext successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				listSegmentsOptionsModel := &appconfigurationv1.ListSegmentsOptions{
					Expand: core.BoolPtr(true),
					Sort: core.StringPtr("created_time"),
					Tags: core.StringPtr("version 1.1,pre-release"),
					Include: core.StringPtr("rules"),
					Limit: core.Int64Ptr(int64(10)),
					Search: core.StringPtr("test tag"),
				}

				pager, err := appConfigurationService.NewSegmentsPager(listSegmentsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []appconfigurationv1.Segment
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use SegmentsPager.GetAll successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				listSegmentsOptionsModel := &appconfigurationv1.ListSegmentsOptions{
					Expand: core.BoolPtr(true),
					Sort: core.StringPtr("created_time"),
					Tags: core.StringPtr("version 1.1,pre-release"),
					Include: core.StringPtr("rules"),
					Limit: core.Int64Ptr(int64(10)),
					Search: core.StringPtr("test tag"),
				}

				pager, err := appConfigurationService.NewSegmentsPager(listSegmentsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateSegment(createSegmentOptions *CreateSegmentOptions) - Operation response error`, func() {
		createSegmentPath := "/segments"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSegmentPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateSegment with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the Rule model
				ruleModel := new(appconfigurationv1.Rule)
				ruleModel.AttributeName = core.StringPtr("email")
				ruleModel.Operator = core.StringPtr("endsWith")
				ruleModel.Values = []string{"@in.mnc.com", "@us.mnc.com"}

				// Construct an instance of the CreateSegmentOptions model
				createSegmentOptionsModel := new(appconfigurationv1.CreateSegmentOptions)
				createSegmentOptionsModel.Name = core.StringPtr("Beta Users")
				createSegmentOptionsModel.SegmentID = core.StringPtr("beta-users")
				createSegmentOptionsModel.Rules = []appconfigurationv1.Rule{*ruleModel}
				createSegmentOptionsModel.Description = core.StringPtr("Segment containing the beta users")
				createSegmentOptionsModel.Tags = core.StringPtr("version: 1.1, stage")
				createSegmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.CreateSegment(createSegmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.CreateSegment(createSegmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateSegment(createSegmentOptions *CreateSegmentOptions)`, func() {
		createSegmentPath := "/segments"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSegmentPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Name", "segment_id": "SegmentID", "description": "Description", "tags": "Tags", "rules": [{"attribute_name": "AttributeName", "operator": "is", "values": ["Values"]}], "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href", "features": [{"feature_id": "FeatureID", "name": "Name"}], "properties": [{"property_id": "PropertyID", "name": "Name"}]}`)
				}))
			})
			It(`Invoke CreateSegment successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the Rule model
				ruleModel := new(appconfigurationv1.Rule)
				ruleModel.AttributeName = core.StringPtr("email")
				ruleModel.Operator = core.StringPtr("endsWith")
				ruleModel.Values = []string{"@in.mnc.com", "@us.mnc.com"}

				// Construct an instance of the CreateSegmentOptions model
				createSegmentOptionsModel := new(appconfigurationv1.CreateSegmentOptions)
				createSegmentOptionsModel.Name = core.StringPtr("Beta Users")
				createSegmentOptionsModel.SegmentID = core.StringPtr("beta-users")
				createSegmentOptionsModel.Rules = []appconfigurationv1.Rule{*ruleModel}
				createSegmentOptionsModel.Description = core.StringPtr("Segment containing the beta users")
				createSegmentOptionsModel.Tags = core.StringPtr("version: 1.1, stage")
				createSegmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.CreateSegmentWithContext(ctx, createSegmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.CreateSegment(createSegmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.CreateSegmentWithContext(ctx, createSegmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSegmentPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Name", "segment_id": "SegmentID", "description": "Description", "tags": "Tags", "rules": [{"attribute_name": "AttributeName", "operator": "is", "values": ["Values"]}], "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href", "features": [{"feature_id": "FeatureID", "name": "Name"}], "properties": [{"property_id": "PropertyID", "name": "Name"}]}`)
				}))
			})
			It(`Invoke CreateSegment successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.CreateSegment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Rule model
				ruleModel := new(appconfigurationv1.Rule)
				ruleModel.AttributeName = core.StringPtr("email")
				ruleModel.Operator = core.StringPtr("endsWith")
				ruleModel.Values = []string{"@in.mnc.com", "@us.mnc.com"}

				// Construct an instance of the CreateSegmentOptions model
				createSegmentOptionsModel := new(appconfigurationv1.CreateSegmentOptions)
				createSegmentOptionsModel.Name = core.StringPtr("Beta Users")
				createSegmentOptionsModel.SegmentID = core.StringPtr("beta-users")
				createSegmentOptionsModel.Rules = []appconfigurationv1.Rule{*ruleModel}
				createSegmentOptionsModel.Description = core.StringPtr("Segment containing the beta users")
				createSegmentOptionsModel.Tags = core.StringPtr("version: 1.1, stage")
				createSegmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.CreateSegment(createSegmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateSegment with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the Rule model
				ruleModel := new(appconfigurationv1.Rule)
				ruleModel.AttributeName = core.StringPtr("email")
				ruleModel.Operator = core.StringPtr("endsWith")
				ruleModel.Values = []string{"@in.mnc.com", "@us.mnc.com"}

				// Construct an instance of the CreateSegmentOptions model
				createSegmentOptionsModel := new(appconfigurationv1.CreateSegmentOptions)
				createSegmentOptionsModel.Name = core.StringPtr("Beta Users")
				createSegmentOptionsModel.SegmentID = core.StringPtr("beta-users")
				createSegmentOptionsModel.Rules = []appconfigurationv1.Rule{*ruleModel}
				createSegmentOptionsModel.Description = core.StringPtr("Segment containing the beta users")
				createSegmentOptionsModel.Tags = core.StringPtr("version: 1.1, stage")
				createSegmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.CreateSegment(createSegmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateSegmentOptions model with no property values
				createSegmentOptionsModelNew := new(appconfigurationv1.CreateSegmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.CreateSegment(createSegmentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateSegment successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the Rule model
				ruleModel := new(appconfigurationv1.Rule)
				ruleModel.AttributeName = core.StringPtr("email")
				ruleModel.Operator = core.StringPtr("endsWith")
				ruleModel.Values = []string{"@in.mnc.com", "@us.mnc.com"}

				// Construct an instance of the CreateSegmentOptions model
				createSegmentOptionsModel := new(appconfigurationv1.CreateSegmentOptions)
				createSegmentOptionsModel.Name = core.StringPtr("Beta Users")
				createSegmentOptionsModel.SegmentID = core.StringPtr("beta-users")
				createSegmentOptionsModel.Rules = []appconfigurationv1.Rule{*ruleModel}
				createSegmentOptionsModel.Description = core.StringPtr("Segment containing the beta users")
				createSegmentOptionsModel.Tags = core.StringPtr("version: 1.1, stage")
				createSegmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.CreateSegment(createSegmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateSegment(updateSegmentOptions *UpdateSegmentOptions) - Operation response error`, func() {
		updateSegmentPath := "/segments/segment_id"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSegmentPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateSegment with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the Rule model
				ruleModel := new(appconfigurationv1.Rule)
				ruleModel.AttributeName = core.StringPtr("testString")
				ruleModel.Operator = core.StringPtr("is")
				ruleModel.Values = []string{"testString"}

				// Construct an instance of the UpdateSegmentOptions model
				updateSegmentOptionsModel := new(appconfigurationv1.UpdateSegmentOptions)
				updateSegmentOptionsModel.SegmentID = core.StringPtr("segment_id")
				updateSegmentOptionsModel.Name = core.StringPtr("testString")
				updateSegmentOptionsModel.Description = core.StringPtr("testString")
				updateSegmentOptionsModel.Tags = core.StringPtr("testString")
				updateSegmentOptionsModel.Rules = []appconfigurationv1.Rule{*ruleModel}
				updateSegmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.UpdateSegment(updateSegmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.UpdateSegment(updateSegmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateSegment(updateSegmentOptions *UpdateSegmentOptions)`, func() {
		updateSegmentPath := "/segments/segment_id"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSegmentPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "segment_id": "SegmentID", "description": "Description", "tags": "Tags", "rules": [{"attribute_name": "AttributeName", "operator": "is", "values": ["Values"]}], "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href", "features": [{"feature_id": "FeatureID", "name": "Name"}], "properties": [{"property_id": "PropertyID", "name": "Name"}]}`)
				}))
			})
			It(`Invoke UpdateSegment successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the Rule model
				ruleModel := new(appconfigurationv1.Rule)
				ruleModel.AttributeName = core.StringPtr("testString")
				ruleModel.Operator = core.StringPtr("is")
				ruleModel.Values = []string{"testString"}

				// Construct an instance of the UpdateSegmentOptions model
				updateSegmentOptionsModel := new(appconfigurationv1.UpdateSegmentOptions)
				updateSegmentOptionsModel.SegmentID = core.StringPtr("segment_id")
				updateSegmentOptionsModel.Name = core.StringPtr("testString")
				updateSegmentOptionsModel.Description = core.StringPtr("testString")
				updateSegmentOptionsModel.Tags = core.StringPtr("testString")
				updateSegmentOptionsModel.Rules = []appconfigurationv1.Rule{*ruleModel}
				updateSegmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.UpdateSegmentWithContext(ctx, updateSegmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.UpdateSegment(updateSegmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.UpdateSegmentWithContext(ctx, updateSegmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSegmentPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "segment_id": "SegmentID", "description": "Description", "tags": "Tags", "rules": [{"attribute_name": "AttributeName", "operator": "is", "values": ["Values"]}], "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href", "features": [{"feature_id": "FeatureID", "name": "Name"}], "properties": [{"property_id": "PropertyID", "name": "Name"}]}`)
				}))
			})
			It(`Invoke UpdateSegment successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.UpdateSegment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Rule model
				ruleModel := new(appconfigurationv1.Rule)
				ruleModel.AttributeName = core.StringPtr("testString")
				ruleModel.Operator = core.StringPtr("is")
				ruleModel.Values = []string{"testString"}

				// Construct an instance of the UpdateSegmentOptions model
				updateSegmentOptionsModel := new(appconfigurationv1.UpdateSegmentOptions)
				updateSegmentOptionsModel.SegmentID = core.StringPtr("segment_id")
				updateSegmentOptionsModel.Name = core.StringPtr("testString")
				updateSegmentOptionsModel.Description = core.StringPtr("testString")
				updateSegmentOptionsModel.Tags = core.StringPtr("testString")
				updateSegmentOptionsModel.Rules = []appconfigurationv1.Rule{*ruleModel}
				updateSegmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.UpdateSegment(updateSegmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateSegment with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the Rule model
				ruleModel := new(appconfigurationv1.Rule)
				ruleModel.AttributeName = core.StringPtr("testString")
				ruleModel.Operator = core.StringPtr("is")
				ruleModel.Values = []string{"testString"}

				// Construct an instance of the UpdateSegmentOptions model
				updateSegmentOptionsModel := new(appconfigurationv1.UpdateSegmentOptions)
				updateSegmentOptionsModel.SegmentID = core.StringPtr("segment_id")
				updateSegmentOptionsModel.Name = core.StringPtr("testString")
				updateSegmentOptionsModel.Description = core.StringPtr("testString")
				updateSegmentOptionsModel.Tags = core.StringPtr("testString")
				updateSegmentOptionsModel.Rules = []appconfigurationv1.Rule{*ruleModel}
				updateSegmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.UpdateSegment(updateSegmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateSegmentOptions model with no property values
				updateSegmentOptionsModelNew := new(appconfigurationv1.UpdateSegmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.UpdateSegment(updateSegmentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateSegment successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the Rule model
				ruleModel := new(appconfigurationv1.Rule)
				ruleModel.AttributeName = core.StringPtr("testString")
				ruleModel.Operator = core.StringPtr("is")
				ruleModel.Values = []string{"testString"}

				// Construct an instance of the UpdateSegmentOptions model
				updateSegmentOptionsModel := new(appconfigurationv1.UpdateSegmentOptions)
				updateSegmentOptionsModel.SegmentID = core.StringPtr("segment_id")
				updateSegmentOptionsModel.Name = core.StringPtr("testString")
				updateSegmentOptionsModel.Description = core.StringPtr("testString")
				updateSegmentOptionsModel.Tags = core.StringPtr("testString")
				updateSegmentOptionsModel.Rules = []appconfigurationv1.Rule{*ruleModel}
				updateSegmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.UpdateSegment(updateSegmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSegment(getSegmentOptions *GetSegmentOptions) - Operation response error`, func() {
		getSegmentPath := "/segments/segment_id"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSegmentPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSegment with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetSegmentOptions model
				getSegmentOptionsModel := new(appconfigurationv1.GetSegmentOptions)
				getSegmentOptionsModel.SegmentID = core.StringPtr("segment_id")
				getSegmentOptionsModel.Include = []string{"features", "properties"}
				getSegmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.GetSegment(getSegmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.GetSegment(getSegmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSegment(getSegmentOptions *GetSegmentOptions)`, func() {
		getSegmentPath := "/segments/segment_id"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSegmentPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "segment_id": "SegmentID", "description": "Description", "tags": "Tags", "rules": [{"attribute_name": "AttributeName", "operator": "is", "values": ["Values"]}], "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href", "features": [{"feature_id": "FeatureID", "name": "Name"}], "properties": [{"property_id": "PropertyID", "name": "Name"}]}`)
				}))
			})
			It(`Invoke GetSegment successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the GetSegmentOptions model
				getSegmentOptionsModel := new(appconfigurationv1.GetSegmentOptions)
				getSegmentOptionsModel.SegmentID = core.StringPtr("segment_id")
				getSegmentOptionsModel.Include = []string{"features", "properties"}
				getSegmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.GetSegmentWithContext(ctx, getSegmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.GetSegment(getSegmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.GetSegmentWithContext(ctx, getSegmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSegmentPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "segment_id": "SegmentID", "description": "Description", "tags": "Tags", "rules": [{"attribute_name": "AttributeName", "operator": "is", "values": ["Values"]}], "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href", "features": [{"feature_id": "FeatureID", "name": "Name"}], "properties": [{"property_id": "PropertyID", "name": "Name"}]}`)
				}))
			})
			It(`Invoke GetSegment successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.GetSegment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSegmentOptions model
				getSegmentOptionsModel := new(appconfigurationv1.GetSegmentOptions)
				getSegmentOptionsModel.SegmentID = core.StringPtr("segment_id")
				getSegmentOptionsModel.Include = []string{"features", "properties"}
				getSegmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.GetSegment(getSegmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSegment with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetSegmentOptions model
				getSegmentOptionsModel := new(appconfigurationv1.GetSegmentOptions)
				getSegmentOptionsModel.SegmentID = core.StringPtr("segment_id")
				getSegmentOptionsModel.Include = []string{"features", "properties"}
				getSegmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.GetSegment(getSegmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSegmentOptions model with no property values
				getSegmentOptionsModelNew := new(appconfigurationv1.GetSegmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.GetSegment(getSegmentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetSegment successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetSegmentOptions model
				getSegmentOptionsModel := new(appconfigurationv1.GetSegmentOptions)
				getSegmentOptionsModel.SegmentID = core.StringPtr("segment_id")
				getSegmentOptionsModel.Include = []string{"features", "properties"}
				getSegmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.GetSegment(getSegmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteSegment(deleteSegmentOptions *DeleteSegmentOptions)`, func() {
		deleteSegmentPath := "/segments/segment_id"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteSegmentPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteSegment successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appConfigurationService.DeleteSegment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteSegmentOptions model
				deleteSegmentOptionsModel := new(appconfigurationv1.DeleteSegmentOptions)
				deleteSegmentOptionsModel.SegmentID = core.StringPtr("segment_id")
				deleteSegmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appConfigurationService.DeleteSegment(deleteSegmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteSegment with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the DeleteSegmentOptions model
				deleteSegmentOptionsModel := new(appconfigurationv1.DeleteSegmentOptions)
				deleteSegmentOptionsModel.SegmentID = core.StringPtr("segment_id")
				deleteSegmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appConfigurationService.DeleteSegment(deleteSegmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteSegmentOptions model with no property values
				deleteSegmentOptionsModelNew := new(appconfigurationv1.DeleteSegmentOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appConfigurationService.DeleteSegment(deleteSegmentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListSnapshots(listSnapshotsOptions *ListSnapshotsOptions) - Operation response error`, func() {
		listSnapshotsPath := "/gitconfigs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSnapshotsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_time"}))
					Expect(req.URL.Query()["collection_id"]).To(Equal([]string{"collection_id"}))
					Expect(req.URL.Query()["environment_id"]).To(Equal([]string{"environment_id"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["search"]).To(Equal([]string{"search_string"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListSnapshots with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListSnapshotsOptions model
				listSnapshotsOptionsModel := new(appconfigurationv1.ListSnapshotsOptions)
				listSnapshotsOptionsModel.Sort = core.StringPtr("created_time")
				listSnapshotsOptionsModel.CollectionID = core.StringPtr("collection_id")
				listSnapshotsOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				listSnapshotsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listSnapshotsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listSnapshotsOptionsModel.Search = core.StringPtr("search_string")
				listSnapshotsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.ListSnapshots(listSnapshotsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.ListSnapshots(listSnapshotsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListSnapshots(listSnapshotsOptions *ListSnapshotsOptions)`, func() {
		listSnapshotsPath := "/gitconfigs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSnapshotsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_time"}))
					Expect(req.URL.Query()["collection_id"]).To(Equal([]string{"collection_id"}))
					Expect(req.URL.Query()["environment_id"]).To(Equal([]string{"environment_id"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["search"]).To(Equal([]string{"search_string"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"git_config": [{"git_config_name": "GitConfigName", "git_config_id": "GitConfigID", "collection": {"name": "Name", "collection_id": "CollectionID"}, "environment": {"name": "Name", "environment_id": "EnvironmentID", "color_code": "ColorCode"}, "git_url": "GitURL", "git_branch": "GitBranch", "git_file_path": "GitFilePath", "last_sync_time": "2022-05-27T23:20:50.520Z", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href"}], "limit": 10, "offset": 0, "total_count": 0, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}}`)
				}))
			})
			It(`Invoke ListSnapshots successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the ListSnapshotsOptions model
				listSnapshotsOptionsModel := new(appconfigurationv1.ListSnapshotsOptions)
				listSnapshotsOptionsModel.Sort = core.StringPtr("created_time")
				listSnapshotsOptionsModel.CollectionID = core.StringPtr("collection_id")
				listSnapshotsOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				listSnapshotsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listSnapshotsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listSnapshotsOptionsModel.Search = core.StringPtr("search_string")
				listSnapshotsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.ListSnapshotsWithContext(ctx, listSnapshotsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.ListSnapshots(listSnapshotsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.ListSnapshotsWithContext(ctx, listSnapshotsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSnapshotsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_time"}))
					Expect(req.URL.Query()["collection_id"]).To(Equal([]string{"collection_id"}))
					Expect(req.URL.Query()["environment_id"]).To(Equal([]string{"environment_id"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["search"]).To(Equal([]string{"search_string"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"git_config": [{"git_config_name": "GitConfigName", "git_config_id": "GitConfigID", "collection": {"name": "Name", "collection_id": "CollectionID"}, "environment": {"name": "Name", "environment_id": "EnvironmentID", "color_code": "ColorCode"}, "git_url": "GitURL", "git_branch": "GitBranch", "git_file_path": "GitFilePath", "last_sync_time": "2022-05-27T23:20:50.520Z", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href"}], "limit": 10, "offset": 0, "total_count": 0, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}}`)
				}))
			})
			It(`Invoke ListSnapshots successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.ListSnapshots(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListSnapshotsOptions model
				listSnapshotsOptionsModel := new(appconfigurationv1.ListSnapshotsOptions)
				listSnapshotsOptionsModel.Sort = core.StringPtr("created_time")
				listSnapshotsOptionsModel.CollectionID = core.StringPtr("collection_id")
				listSnapshotsOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				listSnapshotsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listSnapshotsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listSnapshotsOptionsModel.Search = core.StringPtr("search_string")
				listSnapshotsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.ListSnapshots(listSnapshotsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListSnapshots with error: Operation request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListSnapshotsOptions model
				listSnapshotsOptionsModel := new(appconfigurationv1.ListSnapshotsOptions)
				listSnapshotsOptionsModel.Sort = core.StringPtr("created_time")
				listSnapshotsOptionsModel.CollectionID = core.StringPtr("collection_id")
				listSnapshotsOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				listSnapshotsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listSnapshotsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listSnapshotsOptionsModel.Search = core.StringPtr("search_string")
				listSnapshotsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.ListSnapshots(listSnapshotsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListSnapshots successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListSnapshotsOptions model
				listSnapshotsOptionsModel := new(appconfigurationv1.ListSnapshotsOptions)
				listSnapshotsOptionsModel.Sort = core.StringPtr("created_time")
				listSnapshotsOptionsModel.CollectionID = core.StringPtr("collection_id")
				listSnapshotsOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				listSnapshotsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listSnapshotsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listSnapshotsOptionsModel.Search = core.StringPtr("search_string")
				listSnapshotsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.ListSnapshots(listSnapshotsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(appconfigurationv1.GitConfigList)
				nextObject := new(appconfigurationv1.PaginatedListNext)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(appconfigurationv1.GitConfigList)

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(appconfigurationv1.GitConfigList)
				nextObject := new(appconfigurationv1.PaginatedListNext)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(appconfigurationv1.GitConfigList)
				nextObject := new(appconfigurationv1.PaginatedListNext)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSnapshotsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"total_count":2,"limit":1,"git_config":[{"git_config_name":"GitConfigName","git_config_id":"GitConfigID","collection":{"name":"Name","collection_id":"CollectionID"},"environment":{"name":"Name","environment_id":"EnvironmentID","color_code":"ColorCode"},"git_url":"GitURL","git_branch":"GitBranch","git_file_path":"GitFilePath","last_sync_time":"2022-05-27T23:20:50.520Z","created_time":"2021-05-12T23:20:50.520Z","updated_time":"2021-05-12T23:20:50.520Z","href":"Href"}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"git_config":[{"git_config_name":"GitConfigName","git_config_id":"GitConfigID","collection":{"name":"Name","collection_id":"CollectionID"},"environment":{"name":"Name","environment_id":"EnvironmentID","color_code":"ColorCode"},"git_url":"GitURL","git_branch":"GitBranch","git_file_path":"GitFilePath","last_sync_time":"2022-05-27T23:20:50.520Z","created_time":"2021-05-12T23:20:50.520Z","updated_time":"2021-05-12T23:20:50.520Z","href":"Href"}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use SnapshotsPager.GetNext successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				listSnapshotsOptionsModel := &appconfigurationv1.ListSnapshotsOptions{
					Sort: core.StringPtr("created_time"),
					CollectionID: core.StringPtr("collection_id"),
					EnvironmentID: core.StringPtr("environment_id"),
					Limit: core.Int64Ptr(int64(10)),
					Search: core.StringPtr("search_string"),
				}

				pager, err := appConfigurationService.NewSnapshotsPager(listSnapshotsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []appconfigurationv1.GitConfig
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use SnapshotsPager.GetAll successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				listSnapshotsOptionsModel := &appconfigurationv1.ListSnapshotsOptions{
					Sort: core.StringPtr("created_time"),
					CollectionID: core.StringPtr("collection_id"),
					EnvironmentID: core.StringPtr("environment_id"),
					Limit: core.Int64Ptr(int64(10)),
					Search: core.StringPtr("search_string"),
				}

				pager, err := appConfigurationService.NewSnapshotsPager(listSnapshotsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateGitconfig(createGitconfigOptions *CreateGitconfigOptions) - Operation response error`, func() {
		createGitconfigPath := "/gitconfigs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createGitconfigPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateGitconfig with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the CreateGitconfigOptions model
				createGitconfigOptionsModel := new(appconfigurationv1.CreateGitconfigOptions)
				createGitconfigOptionsModel.GitConfigName = core.StringPtr("boot-strap-configuration")
				createGitconfigOptionsModel.GitConfigID = core.StringPtr("boot-strap-configuration")
				createGitconfigOptionsModel.CollectionID = core.StringPtr("web-app-collection")
				createGitconfigOptionsModel.EnvironmentID = core.StringPtr("dev")
				createGitconfigOptionsModel.GitURL = core.StringPtr("https://github.ibm.com/api/v3/repos/jhondoe-owner/my-test-repo")
				createGitconfigOptionsModel.GitBranch = core.StringPtr("main")
				createGitconfigOptionsModel.GitFilePath = core.StringPtr("code/development/README.json")
				createGitconfigOptionsModel.GitToken = core.StringPtr("61a792eahhGHji223jijb55a6cfdd4d5cde4c8a67esjjhjhHVH")
				createGitconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.CreateGitconfig(createGitconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.CreateGitconfig(createGitconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateGitconfig(createGitconfigOptions *CreateGitconfigOptions)`, func() {
		createGitconfigPath := "/gitconfigs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createGitconfigPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"git_config_name": "GitConfigName", "git_config_id": "GitConfigID", "collection_id": "CollectionID", "environment_id": "EnvironmentID", "git_url": "GitURL", "git_branch": "GitBranch", "git_file_path": "GitFilePath", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href"}`)
				}))
			})
			It(`Invoke CreateGitconfig successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the CreateGitconfigOptions model
				createGitconfigOptionsModel := new(appconfigurationv1.CreateGitconfigOptions)
				createGitconfigOptionsModel.GitConfigName = core.StringPtr("boot-strap-configuration")
				createGitconfigOptionsModel.GitConfigID = core.StringPtr("boot-strap-configuration")
				createGitconfigOptionsModel.CollectionID = core.StringPtr("web-app-collection")
				createGitconfigOptionsModel.EnvironmentID = core.StringPtr("dev")
				createGitconfigOptionsModel.GitURL = core.StringPtr("https://github.ibm.com/api/v3/repos/jhondoe-owner/my-test-repo")
				createGitconfigOptionsModel.GitBranch = core.StringPtr("main")
				createGitconfigOptionsModel.GitFilePath = core.StringPtr("code/development/README.json")
				createGitconfigOptionsModel.GitToken = core.StringPtr("61a792eahhGHji223jijb55a6cfdd4d5cde4c8a67esjjhjhHVH")
				createGitconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.CreateGitconfigWithContext(ctx, createGitconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.CreateGitconfig(createGitconfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.CreateGitconfigWithContext(ctx, createGitconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createGitconfigPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"git_config_name": "GitConfigName", "git_config_id": "GitConfigID", "collection_id": "CollectionID", "environment_id": "EnvironmentID", "git_url": "GitURL", "git_branch": "GitBranch", "git_file_path": "GitFilePath", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href"}`)
				}))
			})
			It(`Invoke CreateGitconfig successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.CreateGitconfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateGitconfigOptions model
				createGitconfigOptionsModel := new(appconfigurationv1.CreateGitconfigOptions)
				createGitconfigOptionsModel.GitConfigName = core.StringPtr("boot-strap-configuration")
				createGitconfigOptionsModel.GitConfigID = core.StringPtr("boot-strap-configuration")
				createGitconfigOptionsModel.CollectionID = core.StringPtr("web-app-collection")
				createGitconfigOptionsModel.EnvironmentID = core.StringPtr("dev")
				createGitconfigOptionsModel.GitURL = core.StringPtr("https://github.ibm.com/api/v3/repos/jhondoe-owner/my-test-repo")
				createGitconfigOptionsModel.GitBranch = core.StringPtr("main")
				createGitconfigOptionsModel.GitFilePath = core.StringPtr("code/development/README.json")
				createGitconfigOptionsModel.GitToken = core.StringPtr("61a792eahhGHji223jijb55a6cfdd4d5cde4c8a67esjjhjhHVH")
				createGitconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.CreateGitconfig(createGitconfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateGitconfig with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the CreateGitconfigOptions model
				createGitconfigOptionsModel := new(appconfigurationv1.CreateGitconfigOptions)
				createGitconfigOptionsModel.GitConfigName = core.StringPtr("boot-strap-configuration")
				createGitconfigOptionsModel.GitConfigID = core.StringPtr("boot-strap-configuration")
				createGitconfigOptionsModel.CollectionID = core.StringPtr("web-app-collection")
				createGitconfigOptionsModel.EnvironmentID = core.StringPtr("dev")
				createGitconfigOptionsModel.GitURL = core.StringPtr("https://github.ibm.com/api/v3/repos/jhondoe-owner/my-test-repo")
				createGitconfigOptionsModel.GitBranch = core.StringPtr("main")
				createGitconfigOptionsModel.GitFilePath = core.StringPtr("code/development/README.json")
				createGitconfigOptionsModel.GitToken = core.StringPtr("61a792eahhGHji223jijb55a6cfdd4d5cde4c8a67esjjhjhHVH")
				createGitconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.CreateGitconfig(createGitconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateGitconfigOptions model with no property values
				createGitconfigOptionsModelNew := new(appconfigurationv1.CreateGitconfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.CreateGitconfig(createGitconfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateGitconfig successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the CreateGitconfigOptions model
				createGitconfigOptionsModel := new(appconfigurationv1.CreateGitconfigOptions)
				createGitconfigOptionsModel.GitConfigName = core.StringPtr("boot-strap-configuration")
				createGitconfigOptionsModel.GitConfigID = core.StringPtr("boot-strap-configuration")
				createGitconfigOptionsModel.CollectionID = core.StringPtr("web-app-collection")
				createGitconfigOptionsModel.EnvironmentID = core.StringPtr("dev")
				createGitconfigOptionsModel.GitURL = core.StringPtr("https://github.ibm.com/api/v3/repos/jhondoe-owner/my-test-repo")
				createGitconfigOptionsModel.GitBranch = core.StringPtr("main")
				createGitconfigOptionsModel.GitFilePath = core.StringPtr("code/development/README.json")
				createGitconfigOptionsModel.GitToken = core.StringPtr("61a792eahhGHji223jijb55a6cfdd4d5cde4c8a67esjjhjhHVH")
				createGitconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.CreateGitconfig(createGitconfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateGitconfig(updateGitconfigOptions *UpdateGitconfigOptions) - Operation response error`, func() {
		updateGitconfigPath := "/gitconfigs/git_config_id"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateGitconfigPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateGitconfig with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the UpdateGitconfigOptions model
				updateGitconfigOptionsModel := new(appconfigurationv1.UpdateGitconfigOptions)
				updateGitconfigOptionsModel.GitConfigID = core.StringPtr("git_config_id")
				updateGitconfigOptionsModel.GitConfigName = core.StringPtr("testString")
				updateGitconfigOptionsModel.CollectionID = core.StringPtr("testString")
				updateGitconfigOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateGitconfigOptionsModel.GitURL = core.StringPtr("testString")
				updateGitconfigOptionsModel.GitBranch = core.StringPtr("testString")
				updateGitconfigOptionsModel.GitFilePath = core.StringPtr("testString")
				updateGitconfigOptionsModel.GitToken = core.StringPtr("testString")
				updateGitconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.UpdateGitconfig(updateGitconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.UpdateGitconfig(updateGitconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateGitconfig(updateGitconfigOptions *UpdateGitconfigOptions)`, func() {
		updateGitconfigPath := "/gitconfigs/git_config_id"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateGitconfigPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"git_config_name": "GitConfigName", "git_config_id": "GitConfigID", "collection": {"name": "Name", "collection_id": "CollectionID"}, "environment": {"name": "Name", "environment_id": "EnvironmentID", "color_code": "ColorCode"}, "git_url": "GitURL", "git_branch": "GitBranch", "git_file_path": "GitFilePath", "last_sync_time": "2022-05-27T23:20:50.520Z", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href"}`)
				}))
			})
			It(`Invoke UpdateGitconfig successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the UpdateGitconfigOptions model
				updateGitconfigOptionsModel := new(appconfigurationv1.UpdateGitconfigOptions)
				updateGitconfigOptionsModel.GitConfigID = core.StringPtr("git_config_id")
				updateGitconfigOptionsModel.GitConfigName = core.StringPtr("testString")
				updateGitconfigOptionsModel.CollectionID = core.StringPtr("testString")
				updateGitconfigOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateGitconfigOptionsModel.GitURL = core.StringPtr("testString")
				updateGitconfigOptionsModel.GitBranch = core.StringPtr("testString")
				updateGitconfigOptionsModel.GitFilePath = core.StringPtr("testString")
				updateGitconfigOptionsModel.GitToken = core.StringPtr("testString")
				updateGitconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.UpdateGitconfigWithContext(ctx, updateGitconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.UpdateGitconfig(updateGitconfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.UpdateGitconfigWithContext(ctx, updateGitconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateGitconfigPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"git_config_name": "GitConfigName", "git_config_id": "GitConfigID", "collection": {"name": "Name", "collection_id": "CollectionID"}, "environment": {"name": "Name", "environment_id": "EnvironmentID", "color_code": "ColorCode"}, "git_url": "GitURL", "git_branch": "GitBranch", "git_file_path": "GitFilePath", "last_sync_time": "2022-05-27T23:20:50.520Z", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href"}`)
				}))
			})
			It(`Invoke UpdateGitconfig successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.UpdateGitconfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateGitconfigOptions model
				updateGitconfigOptionsModel := new(appconfigurationv1.UpdateGitconfigOptions)
				updateGitconfigOptionsModel.GitConfigID = core.StringPtr("git_config_id")
				updateGitconfigOptionsModel.GitConfigName = core.StringPtr("testString")
				updateGitconfigOptionsModel.CollectionID = core.StringPtr("testString")
				updateGitconfigOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateGitconfigOptionsModel.GitURL = core.StringPtr("testString")
				updateGitconfigOptionsModel.GitBranch = core.StringPtr("testString")
				updateGitconfigOptionsModel.GitFilePath = core.StringPtr("testString")
				updateGitconfigOptionsModel.GitToken = core.StringPtr("testString")
				updateGitconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.UpdateGitconfig(updateGitconfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateGitconfig with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the UpdateGitconfigOptions model
				updateGitconfigOptionsModel := new(appconfigurationv1.UpdateGitconfigOptions)
				updateGitconfigOptionsModel.GitConfigID = core.StringPtr("git_config_id")
				updateGitconfigOptionsModel.GitConfigName = core.StringPtr("testString")
				updateGitconfigOptionsModel.CollectionID = core.StringPtr("testString")
				updateGitconfigOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateGitconfigOptionsModel.GitURL = core.StringPtr("testString")
				updateGitconfigOptionsModel.GitBranch = core.StringPtr("testString")
				updateGitconfigOptionsModel.GitFilePath = core.StringPtr("testString")
				updateGitconfigOptionsModel.GitToken = core.StringPtr("testString")
				updateGitconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.UpdateGitconfig(updateGitconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateGitconfigOptions model with no property values
				updateGitconfigOptionsModelNew := new(appconfigurationv1.UpdateGitconfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.UpdateGitconfig(updateGitconfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateGitconfig successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the UpdateGitconfigOptions model
				updateGitconfigOptionsModel := new(appconfigurationv1.UpdateGitconfigOptions)
				updateGitconfigOptionsModel.GitConfigID = core.StringPtr("git_config_id")
				updateGitconfigOptionsModel.GitConfigName = core.StringPtr("testString")
				updateGitconfigOptionsModel.CollectionID = core.StringPtr("testString")
				updateGitconfigOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateGitconfigOptionsModel.GitURL = core.StringPtr("testString")
				updateGitconfigOptionsModel.GitBranch = core.StringPtr("testString")
				updateGitconfigOptionsModel.GitFilePath = core.StringPtr("testString")
				updateGitconfigOptionsModel.GitToken = core.StringPtr("testString")
				updateGitconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.UpdateGitconfig(updateGitconfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetGitconfig(getGitconfigOptions *GetGitconfigOptions) - Operation response error`, func() {
		getGitconfigPath := "/gitconfigs/git_config_id"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGitconfigPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetGitconfig with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetGitconfigOptions model
				getGitconfigOptionsModel := new(appconfigurationv1.GetGitconfigOptions)
				getGitconfigOptionsModel.GitConfigID = core.StringPtr("git_config_id")
				getGitconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.GetGitconfig(getGitconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.GetGitconfig(getGitconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetGitconfig(getGitconfigOptions *GetGitconfigOptions)`, func() {
		getGitconfigPath := "/gitconfigs/git_config_id"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGitconfigPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"git_config_name": "GitConfigName", "git_config_id": "GitConfigID", "collection": {"name": "Name", "collection_id": "CollectionID"}, "environment": {"name": "Name", "environment_id": "EnvironmentID", "color_code": "ColorCode"}, "git_url": "GitURL", "git_branch": "GitBranch", "git_file_path": "GitFilePath", "last_sync_time": "2022-05-27T23:20:50.520Z", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href"}`)
				}))
			})
			It(`Invoke GetGitconfig successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the GetGitconfigOptions model
				getGitconfigOptionsModel := new(appconfigurationv1.GetGitconfigOptions)
				getGitconfigOptionsModel.GitConfigID = core.StringPtr("git_config_id")
				getGitconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.GetGitconfigWithContext(ctx, getGitconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.GetGitconfig(getGitconfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.GetGitconfigWithContext(ctx, getGitconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGitconfigPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"git_config_name": "GitConfigName", "git_config_id": "GitConfigID", "collection": {"name": "Name", "collection_id": "CollectionID"}, "environment": {"name": "Name", "environment_id": "EnvironmentID", "color_code": "ColorCode"}, "git_url": "GitURL", "git_branch": "GitBranch", "git_file_path": "GitFilePath", "last_sync_time": "2022-05-27T23:20:50.520Z", "created_time": "2021-05-12T23:20:50.520Z", "updated_time": "2021-05-12T23:20:50.520Z", "href": "Href"}`)
				}))
			})
			It(`Invoke GetGitconfig successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.GetGitconfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetGitconfigOptions model
				getGitconfigOptionsModel := new(appconfigurationv1.GetGitconfigOptions)
				getGitconfigOptionsModel.GitConfigID = core.StringPtr("git_config_id")
				getGitconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.GetGitconfig(getGitconfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetGitconfig with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetGitconfigOptions model
				getGitconfigOptionsModel := new(appconfigurationv1.GetGitconfigOptions)
				getGitconfigOptionsModel.GitConfigID = core.StringPtr("git_config_id")
				getGitconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.GetGitconfig(getGitconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetGitconfigOptions model with no property values
				getGitconfigOptionsModelNew := new(appconfigurationv1.GetGitconfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.GetGitconfig(getGitconfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetGitconfig successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetGitconfigOptions model
				getGitconfigOptionsModel := new(appconfigurationv1.GetGitconfigOptions)
				getGitconfigOptionsModel.GitConfigID = core.StringPtr("git_config_id")
				getGitconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.GetGitconfig(getGitconfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteGitconfig(deleteGitconfigOptions *DeleteGitconfigOptions)`, func() {
		deleteGitconfigPath := "/gitconfigs/git_config_id"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteGitconfigPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteGitconfig successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appConfigurationService.DeleteGitconfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteGitconfigOptions model
				deleteGitconfigOptionsModel := new(appconfigurationv1.DeleteGitconfigOptions)
				deleteGitconfigOptionsModel.GitConfigID = core.StringPtr("git_config_id")
				deleteGitconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appConfigurationService.DeleteGitconfig(deleteGitconfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteGitconfig with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the DeleteGitconfigOptions model
				deleteGitconfigOptionsModel := new(appconfigurationv1.DeleteGitconfigOptions)
				deleteGitconfigOptionsModel.GitConfigID = core.StringPtr("git_config_id")
				deleteGitconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appConfigurationService.DeleteGitconfig(deleteGitconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteGitconfigOptions model with no property values
				deleteGitconfigOptionsModelNew := new(appconfigurationv1.DeleteGitconfigOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appConfigurationService.DeleteGitconfig(deleteGitconfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PromoteGitconfig(promoteGitconfigOptions *PromoteGitconfigOptions) - Operation response error`, func() {
		promoteGitconfigPath := "/gitconfigs/git_config_id/promote"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(promoteGitconfigPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PromoteGitconfig with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the PromoteGitconfigOptions model
				promoteGitconfigOptionsModel := new(appconfigurationv1.PromoteGitconfigOptions)
				promoteGitconfigOptionsModel.GitConfigID = core.StringPtr("git_config_id")
				promoteGitconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.PromoteGitconfig(promoteGitconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.PromoteGitconfig(promoteGitconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PromoteGitconfig(promoteGitconfigOptions *PromoteGitconfigOptions)`, func() {
		promoteGitconfigPath := "/gitconfigs/git_config_id/promote"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(promoteGitconfigPath))
					Expect(req.Method).To(Equal("PUT"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"git_commit_id": "GitCommitID", "git_commit_message": "GitCommitMessage", "last_sync_time": "2022-05-27T23:20:50.520Z"}`)
				}))
			})
			It(`Invoke PromoteGitconfig successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the PromoteGitconfigOptions model
				promoteGitconfigOptionsModel := new(appconfigurationv1.PromoteGitconfigOptions)
				promoteGitconfigOptionsModel.GitConfigID = core.StringPtr("git_config_id")
				promoteGitconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.PromoteGitconfigWithContext(ctx, promoteGitconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.PromoteGitconfig(promoteGitconfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.PromoteGitconfigWithContext(ctx, promoteGitconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(promoteGitconfigPath))
					Expect(req.Method).To(Equal("PUT"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"git_commit_id": "GitCommitID", "git_commit_message": "GitCommitMessage", "last_sync_time": "2022-05-27T23:20:50.520Z"}`)
				}))
			})
			It(`Invoke PromoteGitconfig successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.PromoteGitconfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PromoteGitconfigOptions model
				promoteGitconfigOptionsModel := new(appconfigurationv1.PromoteGitconfigOptions)
				promoteGitconfigOptionsModel.GitConfigID = core.StringPtr("git_config_id")
				promoteGitconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.PromoteGitconfig(promoteGitconfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PromoteGitconfig with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the PromoteGitconfigOptions model
				promoteGitconfigOptionsModel := new(appconfigurationv1.PromoteGitconfigOptions)
				promoteGitconfigOptionsModel.GitConfigID = core.StringPtr("git_config_id")
				promoteGitconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.PromoteGitconfig(promoteGitconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PromoteGitconfigOptions model with no property values
				promoteGitconfigOptionsModelNew := new(appconfigurationv1.PromoteGitconfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.PromoteGitconfig(promoteGitconfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke PromoteGitconfig successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the PromoteGitconfigOptions model
				promoteGitconfigOptionsModel := new(appconfigurationv1.PromoteGitconfigOptions)
				promoteGitconfigOptionsModel.GitConfigID = core.StringPtr("git_config_id")
				promoteGitconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.PromoteGitconfig(promoteGitconfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RestoreGitconfig(restoreGitconfigOptions *RestoreGitconfigOptions) - Operation response error`, func() {
		restoreGitconfigPath := "/gitconfigs/git_config_id/restore"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(restoreGitconfigPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RestoreGitconfig with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the RestoreGitconfigOptions model
				restoreGitconfigOptionsModel := new(appconfigurationv1.RestoreGitconfigOptions)
				restoreGitconfigOptionsModel.GitConfigID = core.StringPtr("git_config_id")
				restoreGitconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.RestoreGitconfig(restoreGitconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.RestoreGitconfig(restoreGitconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RestoreGitconfig(restoreGitconfigOptions *RestoreGitconfigOptions)`, func() {
		restoreGitconfigPath := "/gitconfigs/git_config_id/restore"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(restoreGitconfigPath))
					Expect(req.Method).To(Equal("PUT"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"environments": [{"name": "Name", "environment_id": "EnvironmentID", "description": "Description", "tags": "Tags", "color_code": "#FDD13A", "features": [{"name": "Name", "feature_id": "FeatureID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "enabled_value": "anyValue", "disabled_value": "anyValue", "enabled": false, "rollout_percentage": 100, "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5, "rollout_percentage": 100}], "collections": [{"collection_id": "CollectionID", "name": "Name"}], "isOverridden": true}], "properties": [{"name": "Name", "property_id": "PropertyID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "value": "anyValue", "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5}], "collections": [{"collection_id": "CollectionID", "name": "Name"}], "isOverridden": true}]}], "segments": [{"name": "Name", "segment_id": "SegmentID", "description": "Description", "tags": "Tags", "rules": [{"attribute_name": "AttributeName", "operator": "is", "values": ["Values"]}]}]}`)
				}))
			})
			It(`Invoke RestoreGitconfig successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the RestoreGitconfigOptions model
				restoreGitconfigOptionsModel := new(appconfigurationv1.RestoreGitconfigOptions)
				restoreGitconfigOptionsModel.GitConfigID = core.StringPtr("git_config_id")
				restoreGitconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.RestoreGitconfigWithContext(ctx, restoreGitconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.RestoreGitconfig(restoreGitconfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.RestoreGitconfigWithContext(ctx, restoreGitconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(restoreGitconfigPath))
					Expect(req.Method).To(Equal("PUT"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"environments": [{"name": "Name", "environment_id": "EnvironmentID", "description": "Description", "tags": "Tags", "color_code": "#FDD13A", "features": [{"name": "Name", "feature_id": "FeatureID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "enabled_value": "anyValue", "disabled_value": "anyValue", "enabled": false, "rollout_percentage": 100, "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5, "rollout_percentage": 100}], "collections": [{"collection_id": "CollectionID", "name": "Name"}], "isOverridden": true}], "properties": [{"name": "Name", "property_id": "PropertyID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "value": "anyValue", "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5}], "collections": [{"collection_id": "CollectionID", "name": "Name"}], "isOverridden": true}]}], "segments": [{"name": "Name", "segment_id": "SegmentID", "description": "Description", "tags": "Tags", "rules": [{"attribute_name": "AttributeName", "operator": "is", "values": ["Values"]}]}]}`)
				}))
			})
			It(`Invoke RestoreGitconfig successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.RestoreGitconfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RestoreGitconfigOptions model
				restoreGitconfigOptionsModel := new(appconfigurationv1.RestoreGitconfigOptions)
				restoreGitconfigOptionsModel.GitConfigID = core.StringPtr("git_config_id")
				restoreGitconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.RestoreGitconfig(restoreGitconfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke RestoreGitconfig with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the RestoreGitconfigOptions model
				restoreGitconfigOptionsModel := new(appconfigurationv1.RestoreGitconfigOptions)
				restoreGitconfigOptionsModel.GitConfigID = core.StringPtr("git_config_id")
				restoreGitconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.RestoreGitconfig(restoreGitconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RestoreGitconfigOptions model with no property values
				restoreGitconfigOptionsModelNew := new(appconfigurationv1.RestoreGitconfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.RestoreGitconfig(restoreGitconfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke RestoreGitconfig successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the RestoreGitconfigOptions model
				restoreGitconfigOptionsModel := new(appconfigurationv1.RestoreGitconfigOptions)
				restoreGitconfigOptionsModel.GitConfigID = core.StringPtr("git_config_id")
				restoreGitconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.RestoreGitconfig(restoreGitconfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListOriginconfigs(listOriginconfigsOptions *ListOriginconfigsOptions) - Operation response error`, func() {
		listOriginconfigsPath := "/originconfigs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listOriginconfigsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListOriginconfigs with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListOriginconfigsOptions model
				listOriginconfigsOptionsModel := new(appconfigurationv1.ListOriginconfigsOptions)
				listOriginconfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.ListOriginconfigs(listOriginconfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.ListOriginconfigs(listOriginconfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListOriginconfigs(listOriginconfigsOptions *ListOriginconfigsOptions)`, func() {
		listOriginconfigsPath := "/originconfigs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listOriginconfigsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"allowed_origins": ["AllowedOrigins"], "created_time": "2022-11-15T23:20:50.000Z", "updated_time": "2022-11-16T21:20:50.000Z", "href": "Href"}`)
				}))
			})
			It(`Invoke ListOriginconfigs successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the ListOriginconfigsOptions model
				listOriginconfigsOptionsModel := new(appconfigurationv1.ListOriginconfigsOptions)
				listOriginconfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.ListOriginconfigsWithContext(ctx, listOriginconfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.ListOriginconfigs(listOriginconfigsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.ListOriginconfigsWithContext(ctx, listOriginconfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listOriginconfigsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"allowed_origins": ["AllowedOrigins"], "created_time": "2022-11-15T23:20:50.000Z", "updated_time": "2022-11-16T21:20:50.000Z", "href": "Href"}`)
				}))
			})
			It(`Invoke ListOriginconfigs successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.ListOriginconfigs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListOriginconfigsOptions model
				listOriginconfigsOptionsModel := new(appconfigurationv1.ListOriginconfigsOptions)
				listOriginconfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.ListOriginconfigs(listOriginconfigsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListOriginconfigs with error: Operation request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListOriginconfigsOptions model
				listOriginconfigsOptionsModel := new(appconfigurationv1.ListOriginconfigsOptions)
				listOriginconfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.ListOriginconfigs(listOriginconfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListOriginconfigs successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListOriginconfigsOptions model
				listOriginconfigsOptionsModel := new(appconfigurationv1.ListOriginconfigsOptions)
				listOriginconfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.ListOriginconfigs(listOriginconfigsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateOriginconfigs(updateOriginconfigsOptions *UpdateOriginconfigsOptions) - Operation response error`, func() {
		updateOriginconfigsPath := "/originconfigs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateOriginconfigsPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateOriginconfigs with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the UpdateOriginconfigsOptions model
				updateOriginconfigsOptionsModel := new(appconfigurationv1.UpdateOriginconfigsOptions)
				updateOriginconfigsOptionsModel.AllowedOrigins = []string{"testString"}
				updateOriginconfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.UpdateOriginconfigs(updateOriginconfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.UpdateOriginconfigs(updateOriginconfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateOriginconfigs(updateOriginconfigsOptions *UpdateOriginconfigsOptions)`, func() {
		updateOriginconfigsPath := "/originconfigs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateOriginconfigsPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"allowed_origins": ["AllowedOrigins"], "created_time": "2022-11-15T23:20:50.000Z", "updated_time": "2022-11-16T21:20:50.000Z", "href": "Href"}`)
				}))
			})
			It(`Invoke UpdateOriginconfigs successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the UpdateOriginconfigsOptions model
				updateOriginconfigsOptionsModel := new(appconfigurationv1.UpdateOriginconfigsOptions)
				updateOriginconfigsOptionsModel.AllowedOrigins = []string{"testString"}
				updateOriginconfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.UpdateOriginconfigsWithContext(ctx, updateOriginconfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.UpdateOriginconfigs(updateOriginconfigsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.UpdateOriginconfigsWithContext(ctx, updateOriginconfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateOriginconfigsPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"allowed_origins": ["AllowedOrigins"], "created_time": "2022-11-15T23:20:50.000Z", "updated_time": "2022-11-16T21:20:50.000Z", "href": "Href"}`)
				}))
			})
			It(`Invoke UpdateOriginconfigs successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.UpdateOriginconfigs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateOriginconfigsOptions model
				updateOriginconfigsOptionsModel := new(appconfigurationv1.UpdateOriginconfigsOptions)
				updateOriginconfigsOptionsModel.AllowedOrigins = []string{"testString"}
				updateOriginconfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.UpdateOriginconfigs(updateOriginconfigsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateOriginconfigs with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the UpdateOriginconfigsOptions model
				updateOriginconfigsOptionsModel := new(appconfigurationv1.UpdateOriginconfigsOptions)
				updateOriginconfigsOptionsModel.AllowedOrigins = []string{"testString"}
				updateOriginconfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.UpdateOriginconfigs(updateOriginconfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateOriginconfigsOptions model with no property values
				updateOriginconfigsOptionsModelNew := new(appconfigurationv1.UpdateOriginconfigsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.UpdateOriginconfigs(updateOriginconfigsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateOriginconfigs successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the UpdateOriginconfigsOptions model
				updateOriginconfigsOptionsModel := new(appconfigurationv1.UpdateOriginconfigsOptions)
				updateOriginconfigsOptionsModel.AllowedOrigins = []string{"testString"}
				updateOriginconfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.UpdateOriginconfigs(updateOriginconfigsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListWorkflowconfig(listWorkflowconfigOptions *ListWorkflowconfigOptions) - Operation response error`, func() {
		listWorkflowconfigPath := "/environments/environment_id/workflowconfigs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listWorkflowconfigPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListWorkflowconfig with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListWorkflowconfigOptions model
				listWorkflowconfigOptionsModel := new(appconfigurationv1.ListWorkflowconfigOptions)
				listWorkflowconfigOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				listWorkflowconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.ListWorkflowconfig(listWorkflowconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.ListWorkflowconfig(listWorkflowconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListWorkflowconfig(listWorkflowconfigOptions *ListWorkflowconfigOptions)`, func() {
		listWorkflowconfigPath := "/environments/environment_id/workflowconfigs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listWorkflowconfigPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"environment_name": "EnvironmentName", "environment_id": "EnvironmentID", "workflow_url": "WorkflowURL", "approval_group_name": "ApprovalGroupName", "approval_expiration": 1, "workflow_credentials": {"username": "admin", "password": "Password", "client_id": "f7b6379b55d08210f8ree233afc7256d", "client_secret": "ClientSecret"}, "enabled": false, "created_time": "2022-11-15T23:20:50.000Z", "updated_time": "2022-11-16T21:20:50.000Z", "href": "Href"}`)
				}))
			})
			It(`Invoke ListWorkflowconfig successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the ListWorkflowconfigOptions model
				listWorkflowconfigOptionsModel := new(appconfigurationv1.ListWorkflowconfigOptions)
				listWorkflowconfigOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				listWorkflowconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.ListWorkflowconfigWithContext(ctx, listWorkflowconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.ListWorkflowconfig(listWorkflowconfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.ListWorkflowconfigWithContext(ctx, listWorkflowconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listWorkflowconfigPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"environment_name": "EnvironmentName", "environment_id": "EnvironmentID", "workflow_url": "WorkflowURL", "approval_group_name": "ApprovalGroupName", "approval_expiration": 1, "workflow_credentials": {"username": "admin", "password": "Password", "client_id": "f7b6379b55d08210f8ree233afc7256d", "client_secret": "ClientSecret"}, "enabled": false, "created_time": "2022-11-15T23:20:50.000Z", "updated_time": "2022-11-16T21:20:50.000Z", "href": "Href"}`)
				}))
			})
			It(`Invoke ListWorkflowconfig successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.ListWorkflowconfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListWorkflowconfigOptions model
				listWorkflowconfigOptionsModel := new(appconfigurationv1.ListWorkflowconfigOptions)
				listWorkflowconfigOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				listWorkflowconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.ListWorkflowconfig(listWorkflowconfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListWorkflowconfig with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListWorkflowconfigOptions model
				listWorkflowconfigOptionsModel := new(appconfigurationv1.ListWorkflowconfigOptions)
				listWorkflowconfigOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				listWorkflowconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.ListWorkflowconfig(listWorkflowconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListWorkflowconfigOptions model with no property values
				listWorkflowconfigOptionsModelNew := new(appconfigurationv1.ListWorkflowconfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.ListWorkflowconfig(listWorkflowconfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListWorkflowconfig successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListWorkflowconfigOptions model
				listWorkflowconfigOptionsModel := new(appconfigurationv1.ListWorkflowconfigOptions)
				listWorkflowconfigOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				listWorkflowconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.ListWorkflowconfig(listWorkflowconfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateWorkflowconfig(createWorkflowconfigOptions *CreateWorkflowconfigOptions) - Operation response error`, func() {
		createWorkflowconfigPath := "/environments/environment_id/workflowconfigs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createWorkflowconfigPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateWorkflowconfig with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ExternalServiceNowCredentials model
				externalServiceNowCredentialsModel := new(appconfigurationv1.ExternalServiceNowCredentials)
				externalServiceNowCredentialsModel.Username = core.StringPtr("user")
				externalServiceNowCredentialsModel.Password = core.StringPtr("pwd")
				externalServiceNowCredentialsModel.ClientID = core.StringPtr("client id value")
				externalServiceNowCredentialsModel.ClientSecret = core.StringPtr("clientsecret")

				// Construct an instance of the CreateWorkflowConfigExternalServiceNow model
				createWorkflowConfigModel := new(appconfigurationv1.CreateWorkflowConfigExternalServiceNow)
				createWorkflowConfigModel.WorkflowURL = core.StringPtr("https://xxxxx.service-now.com")
				createWorkflowConfigModel.ApprovalGroupName = core.StringPtr("WorkflowCRApprovers")
				createWorkflowConfigModel.ApprovalExpiration = core.Int64Ptr(int64(10))
				createWorkflowConfigModel.WorkflowCredentials = externalServiceNowCredentialsModel
				createWorkflowConfigModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the CreateWorkflowconfigOptions model
				createWorkflowconfigOptionsModel := new(appconfigurationv1.CreateWorkflowconfigOptions)
				createWorkflowconfigOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				createWorkflowconfigOptionsModel.WorkflowConfig = createWorkflowConfigModel
				createWorkflowconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.CreateWorkflowconfig(createWorkflowconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.CreateWorkflowconfig(createWorkflowconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateWorkflowconfig(createWorkflowconfigOptions *CreateWorkflowconfigOptions)`, func() {
		createWorkflowconfigPath := "/environments/environment_id/workflowconfigs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createWorkflowconfigPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"environment_name": "EnvironmentName", "environment_id": "EnvironmentID", "workflow_url": "WorkflowURL", "approval_group_name": "ApprovalGroupName", "approval_expiration": 1, "workflow_credentials": {"username": "admin", "password": "Password", "client_id": "f7b6379b55d08210f8ree233afc7256d", "client_secret": "ClientSecret"}, "enabled": false, "created_time": "2022-11-15T23:20:50.000Z", "updated_time": "2022-11-16T21:20:50.000Z", "href": "Href"}`)
				}))
			})
			It(`Invoke CreateWorkflowconfig successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the ExternalServiceNowCredentials model
				externalServiceNowCredentialsModel := new(appconfigurationv1.ExternalServiceNowCredentials)
				externalServiceNowCredentialsModel.Username = core.StringPtr("user")
				externalServiceNowCredentialsModel.Password = core.StringPtr("pwd")
				externalServiceNowCredentialsModel.ClientID = core.StringPtr("client id value")
				externalServiceNowCredentialsModel.ClientSecret = core.StringPtr("clientsecret")

				// Construct an instance of the CreateWorkflowConfigExternalServiceNow model
				createWorkflowConfigModel := new(appconfigurationv1.CreateWorkflowConfigExternalServiceNow)
				createWorkflowConfigModel.WorkflowURL = core.StringPtr("https://xxxxx.service-now.com")
				createWorkflowConfigModel.ApprovalGroupName = core.StringPtr("WorkflowCRApprovers")
				createWorkflowConfigModel.ApprovalExpiration = core.Int64Ptr(int64(10))
				createWorkflowConfigModel.WorkflowCredentials = externalServiceNowCredentialsModel
				createWorkflowConfigModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the CreateWorkflowconfigOptions model
				createWorkflowconfigOptionsModel := new(appconfigurationv1.CreateWorkflowconfigOptions)
				createWorkflowconfigOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				createWorkflowconfigOptionsModel.WorkflowConfig = createWorkflowConfigModel
				createWorkflowconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.CreateWorkflowconfigWithContext(ctx, createWorkflowconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.CreateWorkflowconfig(createWorkflowconfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.CreateWorkflowconfigWithContext(ctx, createWorkflowconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createWorkflowconfigPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"environment_name": "EnvironmentName", "environment_id": "EnvironmentID", "workflow_url": "WorkflowURL", "approval_group_name": "ApprovalGroupName", "approval_expiration": 1, "workflow_credentials": {"username": "admin", "password": "Password", "client_id": "f7b6379b55d08210f8ree233afc7256d", "client_secret": "ClientSecret"}, "enabled": false, "created_time": "2022-11-15T23:20:50.000Z", "updated_time": "2022-11-16T21:20:50.000Z", "href": "Href"}`)
				}))
			})
			It(`Invoke CreateWorkflowconfig successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.CreateWorkflowconfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ExternalServiceNowCredentials model
				externalServiceNowCredentialsModel := new(appconfigurationv1.ExternalServiceNowCredentials)
				externalServiceNowCredentialsModel.Username = core.StringPtr("user")
				externalServiceNowCredentialsModel.Password = core.StringPtr("pwd")
				externalServiceNowCredentialsModel.ClientID = core.StringPtr("client id value")
				externalServiceNowCredentialsModel.ClientSecret = core.StringPtr("clientsecret")

				// Construct an instance of the CreateWorkflowConfigExternalServiceNow model
				createWorkflowConfigModel := new(appconfigurationv1.CreateWorkflowConfigExternalServiceNow)
				createWorkflowConfigModel.WorkflowURL = core.StringPtr("https://xxxxx.service-now.com")
				createWorkflowConfigModel.ApprovalGroupName = core.StringPtr("WorkflowCRApprovers")
				createWorkflowConfigModel.ApprovalExpiration = core.Int64Ptr(int64(10))
				createWorkflowConfigModel.WorkflowCredentials = externalServiceNowCredentialsModel
				createWorkflowConfigModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the CreateWorkflowconfigOptions model
				createWorkflowconfigOptionsModel := new(appconfigurationv1.CreateWorkflowconfigOptions)
				createWorkflowconfigOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				createWorkflowconfigOptionsModel.WorkflowConfig = createWorkflowConfigModel
				createWorkflowconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.CreateWorkflowconfig(createWorkflowconfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateWorkflowconfig with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ExternalServiceNowCredentials model
				externalServiceNowCredentialsModel := new(appconfigurationv1.ExternalServiceNowCredentials)
				externalServiceNowCredentialsModel.Username = core.StringPtr("user")
				externalServiceNowCredentialsModel.Password = core.StringPtr("pwd")
				externalServiceNowCredentialsModel.ClientID = core.StringPtr("client id value")
				externalServiceNowCredentialsModel.ClientSecret = core.StringPtr("clientsecret")

				// Construct an instance of the CreateWorkflowConfigExternalServiceNow model
				createWorkflowConfigModel := new(appconfigurationv1.CreateWorkflowConfigExternalServiceNow)
				createWorkflowConfigModel.WorkflowURL = core.StringPtr("https://xxxxx.service-now.com")
				createWorkflowConfigModel.ApprovalGroupName = core.StringPtr("WorkflowCRApprovers")
				createWorkflowConfigModel.ApprovalExpiration = core.Int64Ptr(int64(10))
				createWorkflowConfigModel.WorkflowCredentials = externalServiceNowCredentialsModel
				createWorkflowConfigModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the CreateWorkflowconfigOptions model
				createWorkflowconfigOptionsModel := new(appconfigurationv1.CreateWorkflowconfigOptions)
				createWorkflowconfigOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				createWorkflowconfigOptionsModel.WorkflowConfig = createWorkflowConfigModel
				createWorkflowconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.CreateWorkflowconfig(createWorkflowconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateWorkflowconfigOptions model with no property values
				createWorkflowconfigOptionsModelNew := new(appconfigurationv1.CreateWorkflowconfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.CreateWorkflowconfig(createWorkflowconfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateWorkflowconfig successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ExternalServiceNowCredentials model
				externalServiceNowCredentialsModel := new(appconfigurationv1.ExternalServiceNowCredentials)
				externalServiceNowCredentialsModel.Username = core.StringPtr("user")
				externalServiceNowCredentialsModel.Password = core.StringPtr("pwd")
				externalServiceNowCredentialsModel.ClientID = core.StringPtr("client id value")
				externalServiceNowCredentialsModel.ClientSecret = core.StringPtr("clientsecret")

				// Construct an instance of the CreateWorkflowConfigExternalServiceNow model
				createWorkflowConfigModel := new(appconfigurationv1.CreateWorkflowConfigExternalServiceNow)
				createWorkflowConfigModel.WorkflowURL = core.StringPtr("https://xxxxx.service-now.com")
				createWorkflowConfigModel.ApprovalGroupName = core.StringPtr("WorkflowCRApprovers")
				createWorkflowConfigModel.ApprovalExpiration = core.Int64Ptr(int64(10))
				createWorkflowConfigModel.WorkflowCredentials = externalServiceNowCredentialsModel
				createWorkflowConfigModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the CreateWorkflowconfigOptions model
				createWorkflowconfigOptionsModel := new(appconfigurationv1.CreateWorkflowconfigOptions)
				createWorkflowconfigOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				createWorkflowconfigOptionsModel.WorkflowConfig = createWorkflowConfigModel
				createWorkflowconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.CreateWorkflowconfig(createWorkflowconfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateWorkflowconfig(updateWorkflowconfigOptions *UpdateWorkflowconfigOptions) - Operation response error`, func() {
		updateWorkflowconfigPath := "/environments/environment_id/workflowconfigs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateWorkflowconfigPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateWorkflowconfig with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ExternalServiceNowCredentials model
				externalServiceNowCredentialsModel := new(appconfigurationv1.ExternalServiceNowCredentials)
				externalServiceNowCredentialsModel.Username = core.StringPtr("user")
				externalServiceNowCredentialsModel.Password = core.StringPtr("updated password")
				externalServiceNowCredentialsModel.ClientID = core.StringPtr("client id value")
				externalServiceNowCredentialsModel.ClientSecret = core.StringPtr("updated client secret")

				// Construct an instance of the UpdateWorkflowConfigUpdateExternalServiceNow model
				updateWorkflowConfigModel := new(appconfigurationv1.UpdateWorkflowConfigUpdateExternalServiceNow)
				updateWorkflowConfigModel.WorkflowURL = core.StringPtr("https://xxxxx.service-now.com")
				updateWorkflowConfigModel.ApprovalGroupName = core.StringPtr("WorkflowCRApprovers")
				updateWorkflowConfigModel.ApprovalExpiration = core.Int64Ptr(int64(5))
				updateWorkflowConfigModel.WorkflowCredentials = externalServiceNowCredentialsModel
				updateWorkflowConfigModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the UpdateWorkflowconfigOptions model
				updateWorkflowconfigOptionsModel := new(appconfigurationv1.UpdateWorkflowconfigOptions)
				updateWorkflowconfigOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				updateWorkflowconfigOptionsModel.UpdateWorkflowConfig = updateWorkflowConfigModel
				updateWorkflowconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.UpdateWorkflowconfig(updateWorkflowconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.UpdateWorkflowconfig(updateWorkflowconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateWorkflowconfig(updateWorkflowconfigOptions *UpdateWorkflowconfigOptions)`, func() {
		updateWorkflowconfigPath := "/environments/environment_id/workflowconfigs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateWorkflowconfigPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"environment_name": "EnvironmentName", "environment_id": "EnvironmentID", "workflow_url": "WorkflowURL", "approval_group_name": "ApprovalGroupName", "approval_expiration": 1, "workflow_credentials": {"username": "admin", "password": "Password", "client_id": "f7b6379b55d08210f8ree233afc7256d", "client_secret": "ClientSecret"}, "enabled": false, "created_time": "2022-11-15T23:20:50.000Z", "updated_time": "2022-11-16T21:20:50.000Z", "href": "Href"}`)
				}))
			})
			It(`Invoke UpdateWorkflowconfig successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the ExternalServiceNowCredentials model
				externalServiceNowCredentialsModel := new(appconfigurationv1.ExternalServiceNowCredentials)
				externalServiceNowCredentialsModel.Username = core.StringPtr("user")
				externalServiceNowCredentialsModel.Password = core.StringPtr("updated password")
				externalServiceNowCredentialsModel.ClientID = core.StringPtr("client id value")
				externalServiceNowCredentialsModel.ClientSecret = core.StringPtr("updated client secret")

				// Construct an instance of the UpdateWorkflowConfigUpdateExternalServiceNow model
				updateWorkflowConfigModel := new(appconfigurationv1.UpdateWorkflowConfigUpdateExternalServiceNow)
				updateWorkflowConfigModel.WorkflowURL = core.StringPtr("https://xxxxx.service-now.com")
				updateWorkflowConfigModel.ApprovalGroupName = core.StringPtr("WorkflowCRApprovers")
				updateWorkflowConfigModel.ApprovalExpiration = core.Int64Ptr(int64(5))
				updateWorkflowConfigModel.WorkflowCredentials = externalServiceNowCredentialsModel
				updateWorkflowConfigModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the UpdateWorkflowconfigOptions model
				updateWorkflowconfigOptionsModel := new(appconfigurationv1.UpdateWorkflowconfigOptions)
				updateWorkflowconfigOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				updateWorkflowconfigOptionsModel.UpdateWorkflowConfig = updateWorkflowConfigModel
				updateWorkflowconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.UpdateWorkflowconfigWithContext(ctx, updateWorkflowconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.UpdateWorkflowconfig(updateWorkflowconfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.UpdateWorkflowconfigWithContext(ctx, updateWorkflowconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateWorkflowconfigPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"environment_name": "EnvironmentName", "environment_id": "EnvironmentID", "workflow_url": "WorkflowURL", "approval_group_name": "ApprovalGroupName", "approval_expiration": 1, "workflow_credentials": {"username": "admin", "password": "Password", "client_id": "f7b6379b55d08210f8ree233afc7256d", "client_secret": "ClientSecret"}, "enabled": false, "created_time": "2022-11-15T23:20:50.000Z", "updated_time": "2022-11-16T21:20:50.000Z", "href": "Href"}`)
				}))
			})
			It(`Invoke UpdateWorkflowconfig successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.UpdateWorkflowconfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ExternalServiceNowCredentials model
				externalServiceNowCredentialsModel := new(appconfigurationv1.ExternalServiceNowCredentials)
				externalServiceNowCredentialsModel.Username = core.StringPtr("user")
				externalServiceNowCredentialsModel.Password = core.StringPtr("updated password")
				externalServiceNowCredentialsModel.ClientID = core.StringPtr("client id value")
				externalServiceNowCredentialsModel.ClientSecret = core.StringPtr("updated client secret")

				// Construct an instance of the UpdateWorkflowConfigUpdateExternalServiceNow model
				updateWorkflowConfigModel := new(appconfigurationv1.UpdateWorkflowConfigUpdateExternalServiceNow)
				updateWorkflowConfigModel.WorkflowURL = core.StringPtr("https://xxxxx.service-now.com")
				updateWorkflowConfigModel.ApprovalGroupName = core.StringPtr("WorkflowCRApprovers")
				updateWorkflowConfigModel.ApprovalExpiration = core.Int64Ptr(int64(5))
				updateWorkflowConfigModel.WorkflowCredentials = externalServiceNowCredentialsModel
				updateWorkflowConfigModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the UpdateWorkflowconfigOptions model
				updateWorkflowconfigOptionsModel := new(appconfigurationv1.UpdateWorkflowconfigOptions)
				updateWorkflowconfigOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				updateWorkflowconfigOptionsModel.UpdateWorkflowConfig = updateWorkflowConfigModel
				updateWorkflowconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.UpdateWorkflowconfig(updateWorkflowconfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateWorkflowconfig with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ExternalServiceNowCredentials model
				externalServiceNowCredentialsModel := new(appconfigurationv1.ExternalServiceNowCredentials)
				externalServiceNowCredentialsModel.Username = core.StringPtr("user")
				externalServiceNowCredentialsModel.Password = core.StringPtr("updated password")
				externalServiceNowCredentialsModel.ClientID = core.StringPtr("client id value")
				externalServiceNowCredentialsModel.ClientSecret = core.StringPtr("updated client secret")

				// Construct an instance of the UpdateWorkflowConfigUpdateExternalServiceNow model
				updateWorkflowConfigModel := new(appconfigurationv1.UpdateWorkflowConfigUpdateExternalServiceNow)
				updateWorkflowConfigModel.WorkflowURL = core.StringPtr("https://xxxxx.service-now.com")
				updateWorkflowConfigModel.ApprovalGroupName = core.StringPtr("WorkflowCRApprovers")
				updateWorkflowConfigModel.ApprovalExpiration = core.Int64Ptr(int64(5))
				updateWorkflowConfigModel.WorkflowCredentials = externalServiceNowCredentialsModel
				updateWorkflowConfigModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the UpdateWorkflowconfigOptions model
				updateWorkflowconfigOptionsModel := new(appconfigurationv1.UpdateWorkflowconfigOptions)
				updateWorkflowconfigOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				updateWorkflowconfigOptionsModel.UpdateWorkflowConfig = updateWorkflowConfigModel
				updateWorkflowconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.UpdateWorkflowconfig(updateWorkflowconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateWorkflowconfigOptions model with no property values
				updateWorkflowconfigOptionsModelNew := new(appconfigurationv1.UpdateWorkflowconfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.UpdateWorkflowconfig(updateWorkflowconfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke UpdateWorkflowconfig successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ExternalServiceNowCredentials model
				externalServiceNowCredentialsModel := new(appconfigurationv1.ExternalServiceNowCredentials)
				externalServiceNowCredentialsModel.Username = core.StringPtr("user")
				externalServiceNowCredentialsModel.Password = core.StringPtr("updated password")
				externalServiceNowCredentialsModel.ClientID = core.StringPtr("client id value")
				externalServiceNowCredentialsModel.ClientSecret = core.StringPtr("updated client secret")

				// Construct an instance of the UpdateWorkflowConfigUpdateExternalServiceNow model
				updateWorkflowConfigModel := new(appconfigurationv1.UpdateWorkflowConfigUpdateExternalServiceNow)
				updateWorkflowConfigModel.WorkflowURL = core.StringPtr("https://xxxxx.service-now.com")
				updateWorkflowConfigModel.ApprovalGroupName = core.StringPtr("WorkflowCRApprovers")
				updateWorkflowConfigModel.ApprovalExpiration = core.Int64Ptr(int64(5))
				updateWorkflowConfigModel.WorkflowCredentials = externalServiceNowCredentialsModel
				updateWorkflowConfigModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the UpdateWorkflowconfigOptions model
				updateWorkflowconfigOptionsModel := new(appconfigurationv1.UpdateWorkflowconfigOptions)
				updateWorkflowconfigOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				updateWorkflowconfigOptionsModel.UpdateWorkflowConfig = updateWorkflowConfigModel
				updateWorkflowconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.UpdateWorkflowconfig(updateWorkflowconfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteWorkflowconfig(deleteWorkflowconfigOptions *DeleteWorkflowconfigOptions)`, func() {
		deleteWorkflowconfigPath := "/environments/environment_id/workflowconfigs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteWorkflowconfigPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteWorkflowconfig successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := appConfigurationService.DeleteWorkflowconfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteWorkflowconfigOptions model
				deleteWorkflowconfigOptionsModel := new(appconfigurationv1.DeleteWorkflowconfigOptions)
				deleteWorkflowconfigOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				deleteWorkflowconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appConfigurationService.DeleteWorkflowconfig(deleteWorkflowconfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteWorkflowconfig with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the DeleteWorkflowconfigOptions model
				deleteWorkflowconfigOptionsModel := new(appconfigurationv1.DeleteWorkflowconfigOptions)
				deleteWorkflowconfigOptionsModel.EnvironmentID = core.StringPtr("environment_id")
				deleteWorkflowconfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appConfigurationService.DeleteWorkflowconfig(deleteWorkflowconfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteWorkflowconfigOptions model with no property values
				deleteWorkflowconfigOptionsModelNew := new(appconfigurationv1.DeleteWorkflowconfigOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appConfigurationService.DeleteWorkflowconfig(deleteWorkflowconfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ImportConfig(importConfigOptions *ImportConfigOptions) - Operation response error`, func() {
		importConfigPath := "/config"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(importConfigPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["clean"]).To(Equal([]string{"true"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ImportConfig with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"testString"}

				// Construct an instance of the FeatureSegmentRule model
				featureSegmentRuleModel := new(appconfigurationv1.FeatureSegmentRule)
				featureSegmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				featureSegmentRuleModel.Value = "testString"
				featureSegmentRuleModel.Order = core.Int64Ptr(int64(38))
				featureSegmentRuleModel.RolloutPercentage = core.Int64Ptr(int64(100))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				collectionRefModel.CollectionID = core.StringPtr("web-app")

				// Construct an instance of the ImportFeatureRequestBody model
				importFeatureRequestBodyModel := new(appconfigurationv1.ImportFeatureRequestBody)
				importFeatureRequestBodyModel.Name = core.StringPtr("Cycle Rentals")
				importFeatureRequestBodyModel.FeatureID = core.StringPtr("cycle-rentals")
				importFeatureRequestBodyModel.Description = core.StringPtr("testString")
				importFeatureRequestBodyModel.Type = core.StringPtr("NUMERIC")
				importFeatureRequestBodyModel.Format = core.StringPtr("TEXT")
				importFeatureRequestBodyModel.EnabledValue = "1"
				importFeatureRequestBodyModel.DisabledValue = "2"
				importFeatureRequestBodyModel.Enabled = core.BoolPtr(true)
				importFeatureRequestBodyModel.RolloutPercentage = core.Int64Ptr(int64(100))
				importFeatureRequestBodyModel.Tags = core.StringPtr("testString")
				importFeatureRequestBodyModel.SegmentRules = []appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel}
				importFeatureRequestBodyModel.Collections = []appconfigurationv1.CollectionRef{*collectionRefModel}
				importFeatureRequestBodyModel.IsOverridden = core.BoolPtr(true)

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = "200"
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the ImportPropertyRequestBody model
				importPropertyRequestBodyModel := new(appconfigurationv1.ImportPropertyRequestBody)
				importPropertyRequestBodyModel.Name = core.StringPtr("Daily Discount")
				importPropertyRequestBodyModel.PropertyID = core.StringPtr("daily_discount")
				importPropertyRequestBodyModel.Description = core.StringPtr("testString")
				importPropertyRequestBodyModel.Type = core.StringPtr("NUMERIC")
				importPropertyRequestBodyModel.Format = core.StringPtr("TEXT")
				importPropertyRequestBodyModel.Value = "100"
				importPropertyRequestBodyModel.Tags = core.StringPtr("pre-release, v1.2")
				importPropertyRequestBodyModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				importPropertyRequestBodyModel.Collections = []appconfigurationv1.CollectionRef{*collectionRefModel}
				importPropertyRequestBodyModel.IsOverridden = core.BoolPtr(true)

				// Construct an instance of the ImportEnvironmentSchema model
				importEnvironmentSchemaModel := new(appconfigurationv1.ImportEnvironmentSchema)
				importEnvironmentSchemaModel.Name = core.StringPtr("Dev")
				importEnvironmentSchemaModel.EnvironmentID = core.StringPtr("dev")
				importEnvironmentSchemaModel.Description = core.StringPtr("Environment created on instance creation")
				importEnvironmentSchemaModel.Tags = core.StringPtr("testString")
				importEnvironmentSchemaModel.ColorCode = core.StringPtr("#FDD13A")
				importEnvironmentSchemaModel.Features = []appconfigurationv1.ImportFeatureRequestBody{*importFeatureRequestBodyModel}
				importEnvironmentSchemaModel.Properties = []appconfigurationv1.ImportPropertyRequestBody{*importPropertyRequestBodyModel}

				// Construct an instance of the ImportCollectionSchema model
				importCollectionSchemaModel := new(appconfigurationv1.ImportCollectionSchema)
				importCollectionSchemaModel.CollectionID = core.StringPtr("web-app")
				importCollectionSchemaModel.Name = core.StringPtr("web-app")
				importCollectionSchemaModel.Description = core.StringPtr("web app collection")
				importCollectionSchemaModel.Tags = core.StringPtr("v1")

				// Construct an instance of the Rule model
				ruleModel := new(appconfigurationv1.Rule)
				ruleModel.AttributeName = core.StringPtr("email")
				ruleModel.Operator = core.StringPtr("is")
				ruleModel.Values = []string{"john@bluecharge.com", "alice@bluecharge.com"}

				// Construct an instance of the ImportSegmentSchema model
				importSegmentSchemaModel := new(appconfigurationv1.ImportSegmentSchema)
				importSegmentSchemaModel.Name = core.StringPtr("Testers")
				importSegmentSchemaModel.SegmentID = core.StringPtr("khpwj68h")
				importSegmentSchemaModel.Description = core.StringPtr("Testers")
				importSegmentSchemaModel.Tags = core.StringPtr("test")
				importSegmentSchemaModel.Rules = []appconfigurationv1.Rule{*ruleModel}

				// Construct an instance of the ImportConfigOptions model
				importConfigOptionsModel := new(appconfigurationv1.ImportConfigOptions)
				importConfigOptionsModel.Environments = []appconfigurationv1.ImportEnvironmentSchema{*importEnvironmentSchemaModel}
				importConfigOptionsModel.Collections = []appconfigurationv1.ImportCollectionSchema{*importCollectionSchemaModel}
				importConfigOptionsModel.Segments = []appconfigurationv1.ImportSegmentSchema{*importSegmentSchemaModel}
				importConfigOptionsModel.Clean = core.StringPtr("true")
				importConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.ImportConfig(importConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.ImportConfig(importConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ImportConfig(importConfigOptions *ImportConfigOptions)`, func() {
		importConfigPath := "/config"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(importConfigPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["clean"]).To(Equal([]string{"true"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"environments": [{"name": "Name", "environment_id": "EnvironmentID", "description": "Description", "tags": "Tags", "color_code": "#FDD13A", "features": [{"name": "Name", "feature_id": "FeatureID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "enabled_value": "anyValue", "disabled_value": "anyValue", "enabled": false, "rollout_percentage": 100, "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5, "rollout_percentage": 100}], "collections": [{"collection_id": "CollectionID", "name": "Name"}], "isOverridden": true}], "properties": [{"name": "Name", "property_id": "PropertyID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "value": "anyValue", "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5}], "collections": [{"collection_id": "CollectionID", "name": "Name"}], "isOverridden": true}]}], "collections": [{"collection_id": "CollectionID", "name": "Name", "description": "Description", "tags": "Tags"}], "segments": [{"name": "Name", "segment_id": "SegmentID", "description": "Description", "tags": "Tags", "rules": [{"attribute_name": "AttributeName", "operator": "is", "values": ["Values"]}]}]}`)
				}))
			})
			It(`Invoke ImportConfig successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"testString"}

				// Construct an instance of the FeatureSegmentRule model
				featureSegmentRuleModel := new(appconfigurationv1.FeatureSegmentRule)
				featureSegmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				featureSegmentRuleModel.Value = "testString"
				featureSegmentRuleModel.Order = core.Int64Ptr(int64(38))
				featureSegmentRuleModel.RolloutPercentage = core.Int64Ptr(int64(100))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				collectionRefModel.CollectionID = core.StringPtr("web-app")

				// Construct an instance of the ImportFeatureRequestBody model
				importFeatureRequestBodyModel := new(appconfigurationv1.ImportFeatureRequestBody)
				importFeatureRequestBodyModel.Name = core.StringPtr("Cycle Rentals")
				importFeatureRequestBodyModel.FeatureID = core.StringPtr("cycle-rentals")
				importFeatureRequestBodyModel.Description = core.StringPtr("testString")
				importFeatureRequestBodyModel.Type = core.StringPtr("NUMERIC")
				importFeatureRequestBodyModel.Format = core.StringPtr("TEXT")
				importFeatureRequestBodyModel.EnabledValue = "1"
				importFeatureRequestBodyModel.DisabledValue = "2"
				importFeatureRequestBodyModel.Enabled = core.BoolPtr(true)
				importFeatureRequestBodyModel.RolloutPercentage = core.Int64Ptr(int64(100))
				importFeatureRequestBodyModel.Tags = core.StringPtr("testString")
				importFeatureRequestBodyModel.SegmentRules = []appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel}
				importFeatureRequestBodyModel.Collections = []appconfigurationv1.CollectionRef{*collectionRefModel}
				importFeatureRequestBodyModel.IsOverridden = core.BoolPtr(true)

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = "200"
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the ImportPropertyRequestBody model
				importPropertyRequestBodyModel := new(appconfigurationv1.ImportPropertyRequestBody)
				importPropertyRequestBodyModel.Name = core.StringPtr("Daily Discount")
				importPropertyRequestBodyModel.PropertyID = core.StringPtr("daily_discount")
				importPropertyRequestBodyModel.Description = core.StringPtr("testString")
				importPropertyRequestBodyModel.Type = core.StringPtr("NUMERIC")
				importPropertyRequestBodyModel.Format = core.StringPtr("TEXT")
				importPropertyRequestBodyModel.Value = "100"
				importPropertyRequestBodyModel.Tags = core.StringPtr("pre-release, v1.2")
				importPropertyRequestBodyModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				importPropertyRequestBodyModel.Collections = []appconfigurationv1.CollectionRef{*collectionRefModel}
				importPropertyRequestBodyModel.IsOverridden = core.BoolPtr(true)

				// Construct an instance of the ImportEnvironmentSchema model
				importEnvironmentSchemaModel := new(appconfigurationv1.ImportEnvironmentSchema)
				importEnvironmentSchemaModel.Name = core.StringPtr("Dev")
				importEnvironmentSchemaModel.EnvironmentID = core.StringPtr("dev")
				importEnvironmentSchemaModel.Description = core.StringPtr("Environment created on instance creation")
				importEnvironmentSchemaModel.Tags = core.StringPtr("testString")
				importEnvironmentSchemaModel.ColorCode = core.StringPtr("#FDD13A")
				importEnvironmentSchemaModel.Features = []appconfigurationv1.ImportFeatureRequestBody{*importFeatureRequestBodyModel}
				importEnvironmentSchemaModel.Properties = []appconfigurationv1.ImportPropertyRequestBody{*importPropertyRequestBodyModel}

				// Construct an instance of the ImportCollectionSchema model
				importCollectionSchemaModel := new(appconfigurationv1.ImportCollectionSchema)
				importCollectionSchemaModel.CollectionID = core.StringPtr("web-app")
				importCollectionSchemaModel.Name = core.StringPtr("web-app")
				importCollectionSchemaModel.Description = core.StringPtr("web app collection")
				importCollectionSchemaModel.Tags = core.StringPtr("v1")

				// Construct an instance of the Rule model
				ruleModel := new(appconfigurationv1.Rule)
				ruleModel.AttributeName = core.StringPtr("email")
				ruleModel.Operator = core.StringPtr("is")
				ruleModel.Values = []string{"john@bluecharge.com", "alice@bluecharge.com"}

				// Construct an instance of the ImportSegmentSchema model
				importSegmentSchemaModel := new(appconfigurationv1.ImportSegmentSchema)
				importSegmentSchemaModel.Name = core.StringPtr("Testers")
				importSegmentSchemaModel.SegmentID = core.StringPtr("khpwj68h")
				importSegmentSchemaModel.Description = core.StringPtr("Testers")
				importSegmentSchemaModel.Tags = core.StringPtr("test")
				importSegmentSchemaModel.Rules = []appconfigurationv1.Rule{*ruleModel}

				// Construct an instance of the ImportConfigOptions model
				importConfigOptionsModel := new(appconfigurationv1.ImportConfigOptions)
				importConfigOptionsModel.Environments = []appconfigurationv1.ImportEnvironmentSchema{*importEnvironmentSchemaModel}
				importConfigOptionsModel.Collections = []appconfigurationv1.ImportCollectionSchema{*importCollectionSchemaModel}
				importConfigOptionsModel.Segments = []appconfigurationv1.ImportSegmentSchema{*importSegmentSchemaModel}
				importConfigOptionsModel.Clean = core.StringPtr("true")
				importConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.ImportConfigWithContext(ctx, importConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.ImportConfig(importConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.ImportConfigWithContext(ctx, importConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(importConfigPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["clean"]).To(Equal([]string{"true"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"environments": [{"name": "Name", "environment_id": "EnvironmentID", "description": "Description", "tags": "Tags", "color_code": "#FDD13A", "features": [{"name": "Name", "feature_id": "FeatureID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "enabled_value": "anyValue", "disabled_value": "anyValue", "enabled": false, "rollout_percentage": 100, "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5, "rollout_percentage": 100}], "collections": [{"collection_id": "CollectionID", "name": "Name"}], "isOverridden": true}], "properties": [{"name": "Name", "property_id": "PropertyID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "value": "anyValue", "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5}], "collections": [{"collection_id": "CollectionID", "name": "Name"}], "isOverridden": true}]}], "collections": [{"collection_id": "CollectionID", "name": "Name", "description": "Description", "tags": "Tags"}], "segments": [{"name": "Name", "segment_id": "SegmentID", "description": "Description", "tags": "Tags", "rules": [{"attribute_name": "AttributeName", "operator": "is", "values": ["Values"]}]}]}`)
				}))
			})
			It(`Invoke ImportConfig successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.ImportConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"testString"}

				// Construct an instance of the FeatureSegmentRule model
				featureSegmentRuleModel := new(appconfigurationv1.FeatureSegmentRule)
				featureSegmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				featureSegmentRuleModel.Value = "testString"
				featureSegmentRuleModel.Order = core.Int64Ptr(int64(38))
				featureSegmentRuleModel.RolloutPercentage = core.Int64Ptr(int64(100))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				collectionRefModel.CollectionID = core.StringPtr("web-app")

				// Construct an instance of the ImportFeatureRequestBody model
				importFeatureRequestBodyModel := new(appconfigurationv1.ImportFeatureRequestBody)
				importFeatureRequestBodyModel.Name = core.StringPtr("Cycle Rentals")
				importFeatureRequestBodyModel.FeatureID = core.StringPtr("cycle-rentals")
				importFeatureRequestBodyModel.Description = core.StringPtr("testString")
				importFeatureRequestBodyModel.Type = core.StringPtr("NUMERIC")
				importFeatureRequestBodyModel.Format = core.StringPtr("TEXT")
				importFeatureRequestBodyModel.EnabledValue = "1"
				importFeatureRequestBodyModel.DisabledValue = "2"
				importFeatureRequestBodyModel.Enabled = core.BoolPtr(true)
				importFeatureRequestBodyModel.RolloutPercentage = core.Int64Ptr(int64(100))
				importFeatureRequestBodyModel.Tags = core.StringPtr("testString")
				importFeatureRequestBodyModel.SegmentRules = []appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel}
				importFeatureRequestBodyModel.Collections = []appconfigurationv1.CollectionRef{*collectionRefModel}
				importFeatureRequestBodyModel.IsOverridden = core.BoolPtr(true)

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = "200"
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the ImportPropertyRequestBody model
				importPropertyRequestBodyModel := new(appconfigurationv1.ImportPropertyRequestBody)
				importPropertyRequestBodyModel.Name = core.StringPtr("Daily Discount")
				importPropertyRequestBodyModel.PropertyID = core.StringPtr("daily_discount")
				importPropertyRequestBodyModel.Description = core.StringPtr("testString")
				importPropertyRequestBodyModel.Type = core.StringPtr("NUMERIC")
				importPropertyRequestBodyModel.Format = core.StringPtr("TEXT")
				importPropertyRequestBodyModel.Value = "100"
				importPropertyRequestBodyModel.Tags = core.StringPtr("pre-release, v1.2")
				importPropertyRequestBodyModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				importPropertyRequestBodyModel.Collections = []appconfigurationv1.CollectionRef{*collectionRefModel}
				importPropertyRequestBodyModel.IsOverridden = core.BoolPtr(true)

				// Construct an instance of the ImportEnvironmentSchema model
				importEnvironmentSchemaModel := new(appconfigurationv1.ImportEnvironmentSchema)
				importEnvironmentSchemaModel.Name = core.StringPtr("Dev")
				importEnvironmentSchemaModel.EnvironmentID = core.StringPtr("dev")
				importEnvironmentSchemaModel.Description = core.StringPtr("Environment created on instance creation")
				importEnvironmentSchemaModel.Tags = core.StringPtr("testString")
				importEnvironmentSchemaModel.ColorCode = core.StringPtr("#FDD13A")
				importEnvironmentSchemaModel.Features = []appconfigurationv1.ImportFeatureRequestBody{*importFeatureRequestBodyModel}
				importEnvironmentSchemaModel.Properties = []appconfigurationv1.ImportPropertyRequestBody{*importPropertyRequestBodyModel}

				// Construct an instance of the ImportCollectionSchema model
				importCollectionSchemaModel := new(appconfigurationv1.ImportCollectionSchema)
				importCollectionSchemaModel.CollectionID = core.StringPtr("web-app")
				importCollectionSchemaModel.Name = core.StringPtr("web-app")
				importCollectionSchemaModel.Description = core.StringPtr("web app collection")
				importCollectionSchemaModel.Tags = core.StringPtr("v1")

				// Construct an instance of the Rule model
				ruleModel := new(appconfigurationv1.Rule)
				ruleModel.AttributeName = core.StringPtr("email")
				ruleModel.Operator = core.StringPtr("is")
				ruleModel.Values = []string{"john@bluecharge.com", "alice@bluecharge.com"}

				// Construct an instance of the ImportSegmentSchema model
				importSegmentSchemaModel := new(appconfigurationv1.ImportSegmentSchema)
				importSegmentSchemaModel.Name = core.StringPtr("Testers")
				importSegmentSchemaModel.SegmentID = core.StringPtr("khpwj68h")
				importSegmentSchemaModel.Description = core.StringPtr("Testers")
				importSegmentSchemaModel.Tags = core.StringPtr("test")
				importSegmentSchemaModel.Rules = []appconfigurationv1.Rule{*ruleModel}

				// Construct an instance of the ImportConfigOptions model
				importConfigOptionsModel := new(appconfigurationv1.ImportConfigOptions)
				importConfigOptionsModel.Environments = []appconfigurationv1.ImportEnvironmentSchema{*importEnvironmentSchemaModel}
				importConfigOptionsModel.Collections = []appconfigurationv1.ImportCollectionSchema{*importCollectionSchemaModel}
				importConfigOptionsModel.Segments = []appconfigurationv1.ImportSegmentSchema{*importSegmentSchemaModel}
				importConfigOptionsModel.Clean = core.StringPtr("true")
				importConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.ImportConfig(importConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ImportConfig with error: Operation request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"testString"}

				// Construct an instance of the FeatureSegmentRule model
				featureSegmentRuleModel := new(appconfigurationv1.FeatureSegmentRule)
				featureSegmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				featureSegmentRuleModel.Value = "testString"
				featureSegmentRuleModel.Order = core.Int64Ptr(int64(38))
				featureSegmentRuleModel.RolloutPercentage = core.Int64Ptr(int64(100))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				collectionRefModel.CollectionID = core.StringPtr("web-app")

				// Construct an instance of the ImportFeatureRequestBody model
				importFeatureRequestBodyModel := new(appconfigurationv1.ImportFeatureRequestBody)
				importFeatureRequestBodyModel.Name = core.StringPtr("Cycle Rentals")
				importFeatureRequestBodyModel.FeatureID = core.StringPtr("cycle-rentals")
				importFeatureRequestBodyModel.Description = core.StringPtr("testString")
				importFeatureRequestBodyModel.Type = core.StringPtr("NUMERIC")
				importFeatureRequestBodyModel.Format = core.StringPtr("TEXT")
				importFeatureRequestBodyModel.EnabledValue = "1"
				importFeatureRequestBodyModel.DisabledValue = "2"
				importFeatureRequestBodyModel.Enabled = core.BoolPtr(true)
				importFeatureRequestBodyModel.RolloutPercentage = core.Int64Ptr(int64(100))
				importFeatureRequestBodyModel.Tags = core.StringPtr("testString")
				importFeatureRequestBodyModel.SegmentRules = []appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel}
				importFeatureRequestBodyModel.Collections = []appconfigurationv1.CollectionRef{*collectionRefModel}
				importFeatureRequestBodyModel.IsOverridden = core.BoolPtr(true)

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = "200"
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the ImportPropertyRequestBody model
				importPropertyRequestBodyModel := new(appconfigurationv1.ImportPropertyRequestBody)
				importPropertyRequestBodyModel.Name = core.StringPtr("Daily Discount")
				importPropertyRequestBodyModel.PropertyID = core.StringPtr("daily_discount")
				importPropertyRequestBodyModel.Description = core.StringPtr("testString")
				importPropertyRequestBodyModel.Type = core.StringPtr("NUMERIC")
				importPropertyRequestBodyModel.Format = core.StringPtr("TEXT")
				importPropertyRequestBodyModel.Value = "100"
				importPropertyRequestBodyModel.Tags = core.StringPtr("pre-release, v1.2")
				importPropertyRequestBodyModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				importPropertyRequestBodyModel.Collections = []appconfigurationv1.CollectionRef{*collectionRefModel}
				importPropertyRequestBodyModel.IsOverridden = core.BoolPtr(true)

				// Construct an instance of the ImportEnvironmentSchema model
				importEnvironmentSchemaModel := new(appconfigurationv1.ImportEnvironmentSchema)
				importEnvironmentSchemaModel.Name = core.StringPtr("Dev")
				importEnvironmentSchemaModel.EnvironmentID = core.StringPtr("dev")
				importEnvironmentSchemaModel.Description = core.StringPtr("Environment created on instance creation")
				importEnvironmentSchemaModel.Tags = core.StringPtr("testString")
				importEnvironmentSchemaModel.ColorCode = core.StringPtr("#FDD13A")
				importEnvironmentSchemaModel.Features = []appconfigurationv1.ImportFeatureRequestBody{*importFeatureRequestBodyModel}
				importEnvironmentSchemaModel.Properties = []appconfigurationv1.ImportPropertyRequestBody{*importPropertyRequestBodyModel}

				// Construct an instance of the ImportCollectionSchema model
				importCollectionSchemaModel := new(appconfigurationv1.ImportCollectionSchema)
				importCollectionSchemaModel.CollectionID = core.StringPtr("web-app")
				importCollectionSchemaModel.Name = core.StringPtr("web-app")
				importCollectionSchemaModel.Description = core.StringPtr("web app collection")
				importCollectionSchemaModel.Tags = core.StringPtr("v1")

				// Construct an instance of the Rule model
				ruleModel := new(appconfigurationv1.Rule)
				ruleModel.AttributeName = core.StringPtr("email")
				ruleModel.Operator = core.StringPtr("is")
				ruleModel.Values = []string{"john@bluecharge.com", "alice@bluecharge.com"}

				// Construct an instance of the ImportSegmentSchema model
				importSegmentSchemaModel := new(appconfigurationv1.ImportSegmentSchema)
				importSegmentSchemaModel.Name = core.StringPtr("Testers")
				importSegmentSchemaModel.SegmentID = core.StringPtr("khpwj68h")
				importSegmentSchemaModel.Description = core.StringPtr("Testers")
				importSegmentSchemaModel.Tags = core.StringPtr("test")
				importSegmentSchemaModel.Rules = []appconfigurationv1.Rule{*ruleModel}

				// Construct an instance of the ImportConfigOptions model
				importConfigOptionsModel := new(appconfigurationv1.ImportConfigOptions)
				importConfigOptionsModel.Environments = []appconfigurationv1.ImportEnvironmentSchema{*importEnvironmentSchemaModel}
				importConfigOptionsModel.Collections = []appconfigurationv1.ImportCollectionSchema{*importCollectionSchemaModel}
				importConfigOptionsModel.Segments = []appconfigurationv1.ImportSegmentSchema{*importSegmentSchemaModel}
				importConfigOptionsModel.Clean = core.StringPtr("true")
				importConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.ImportConfig(importConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke ImportConfig successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"testString"}

				// Construct an instance of the FeatureSegmentRule model
				featureSegmentRuleModel := new(appconfigurationv1.FeatureSegmentRule)
				featureSegmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				featureSegmentRuleModel.Value = "testString"
				featureSegmentRuleModel.Order = core.Int64Ptr(int64(38))
				featureSegmentRuleModel.RolloutPercentage = core.Int64Ptr(int64(100))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				collectionRefModel.CollectionID = core.StringPtr("web-app")

				// Construct an instance of the ImportFeatureRequestBody model
				importFeatureRequestBodyModel := new(appconfigurationv1.ImportFeatureRequestBody)
				importFeatureRequestBodyModel.Name = core.StringPtr("Cycle Rentals")
				importFeatureRequestBodyModel.FeatureID = core.StringPtr("cycle-rentals")
				importFeatureRequestBodyModel.Description = core.StringPtr("testString")
				importFeatureRequestBodyModel.Type = core.StringPtr("NUMERIC")
				importFeatureRequestBodyModel.Format = core.StringPtr("TEXT")
				importFeatureRequestBodyModel.EnabledValue = "1"
				importFeatureRequestBodyModel.DisabledValue = "2"
				importFeatureRequestBodyModel.Enabled = core.BoolPtr(true)
				importFeatureRequestBodyModel.RolloutPercentage = core.Int64Ptr(int64(100))
				importFeatureRequestBodyModel.Tags = core.StringPtr("testString")
				importFeatureRequestBodyModel.SegmentRules = []appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel}
				importFeatureRequestBodyModel.Collections = []appconfigurationv1.CollectionRef{*collectionRefModel}
				importFeatureRequestBodyModel.IsOverridden = core.BoolPtr(true)

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = "200"
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the ImportPropertyRequestBody model
				importPropertyRequestBodyModel := new(appconfigurationv1.ImportPropertyRequestBody)
				importPropertyRequestBodyModel.Name = core.StringPtr("Daily Discount")
				importPropertyRequestBodyModel.PropertyID = core.StringPtr("daily_discount")
				importPropertyRequestBodyModel.Description = core.StringPtr("testString")
				importPropertyRequestBodyModel.Type = core.StringPtr("NUMERIC")
				importPropertyRequestBodyModel.Format = core.StringPtr("TEXT")
				importPropertyRequestBodyModel.Value = "100"
				importPropertyRequestBodyModel.Tags = core.StringPtr("pre-release, v1.2")
				importPropertyRequestBodyModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				importPropertyRequestBodyModel.Collections = []appconfigurationv1.CollectionRef{*collectionRefModel}
				importPropertyRequestBodyModel.IsOverridden = core.BoolPtr(true)

				// Construct an instance of the ImportEnvironmentSchema model
				importEnvironmentSchemaModel := new(appconfigurationv1.ImportEnvironmentSchema)
				importEnvironmentSchemaModel.Name = core.StringPtr("Dev")
				importEnvironmentSchemaModel.EnvironmentID = core.StringPtr("dev")
				importEnvironmentSchemaModel.Description = core.StringPtr("Environment created on instance creation")
				importEnvironmentSchemaModel.Tags = core.StringPtr("testString")
				importEnvironmentSchemaModel.ColorCode = core.StringPtr("#FDD13A")
				importEnvironmentSchemaModel.Features = []appconfigurationv1.ImportFeatureRequestBody{*importFeatureRequestBodyModel}
				importEnvironmentSchemaModel.Properties = []appconfigurationv1.ImportPropertyRequestBody{*importPropertyRequestBodyModel}

				// Construct an instance of the ImportCollectionSchema model
				importCollectionSchemaModel := new(appconfigurationv1.ImportCollectionSchema)
				importCollectionSchemaModel.CollectionID = core.StringPtr("web-app")
				importCollectionSchemaModel.Name = core.StringPtr("web-app")
				importCollectionSchemaModel.Description = core.StringPtr("web app collection")
				importCollectionSchemaModel.Tags = core.StringPtr("v1")

				// Construct an instance of the Rule model
				ruleModel := new(appconfigurationv1.Rule)
				ruleModel.AttributeName = core.StringPtr("email")
				ruleModel.Operator = core.StringPtr("is")
				ruleModel.Values = []string{"john@bluecharge.com", "alice@bluecharge.com"}

				// Construct an instance of the ImportSegmentSchema model
				importSegmentSchemaModel := new(appconfigurationv1.ImportSegmentSchema)
				importSegmentSchemaModel.Name = core.StringPtr("Testers")
				importSegmentSchemaModel.SegmentID = core.StringPtr("khpwj68h")
				importSegmentSchemaModel.Description = core.StringPtr("Testers")
				importSegmentSchemaModel.Tags = core.StringPtr("test")
				importSegmentSchemaModel.Rules = []appconfigurationv1.Rule{*ruleModel}

				// Construct an instance of the ImportConfigOptions model
				importConfigOptionsModel := new(appconfigurationv1.ImportConfigOptions)
				importConfigOptionsModel.Environments = []appconfigurationv1.ImportEnvironmentSchema{*importEnvironmentSchemaModel}
				importConfigOptionsModel.Collections = []appconfigurationv1.ImportCollectionSchema{*importCollectionSchemaModel}
				importConfigOptionsModel.Segments = []appconfigurationv1.ImportSegmentSchema{*importSegmentSchemaModel}
				importConfigOptionsModel.Clean = core.StringPtr("true")
				importConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.ImportConfig(importConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListInstanceConfig(listInstanceConfigOptions *ListInstanceConfigOptions) - Operation response error`, func() {
		listInstanceConfigPath := "/config"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listInstanceConfigPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListInstanceConfig with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListInstanceConfigOptions model
				listInstanceConfigOptionsModel := new(appconfigurationv1.ListInstanceConfigOptions)
				listInstanceConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.ListInstanceConfig(listInstanceConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.ListInstanceConfig(listInstanceConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListInstanceConfig(listInstanceConfigOptions *ListInstanceConfigOptions)`, func() {
		listInstanceConfigPath := "/config"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listInstanceConfigPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"environments": [{"name": "Name", "environment_id": "EnvironmentID", "description": "Description", "tags": "Tags", "color_code": "#FDD13A", "features": [{"name": "Name", "feature_id": "FeatureID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "enabled_value": "anyValue", "disabled_value": "anyValue", "enabled": false, "rollout_percentage": 100, "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5, "rollout_percentage": 100}], "collections": [{"collection_id": "CollectionID", "name": "Name"}], "isOverridden": true}], "properties": [{"name": "Name", "property_id": "PropertyID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "value": "anyValue", "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5}], "collections": [{"collection_id": "CollectionID", "name": "Name"}], "isOverridden": true}]}], "collections": [{"collection_id": "CollectionID", "name": "Name", "description": "Description", "tags": "Tags"}], "segments": [{"name": "Name", "segment_id": "SegmentID", "description": "Description", "tags": "Tags", "rules": [{"attribute_name": "AttributeName", "operator": "is", "values": ["Values"]}]}]}`)
				}))
			})
			It(`Invoke ListInstanceConfig successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the ListInstanceConfigOptions model
				listInstanceConfigOptionsModel := new(appconfigurationv1.ListInstanceConfigOptions)
				listInstanceConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.ListInstanceConfigWithContext(ctx, listInstanceConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.ListInstanceConfig(listInstanceConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.ListInstanceConfigWithContext(ctx, listInstanceConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listInstanceConfigPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"environments": [{"name": "Name", "environment_id": "EnvironmentID", "description": "Description", "tags": "Tags", "color_code": "#FDD13A", "features": [{"name": "Name", "feature_id": "FeatureID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "enabled_value": "anyValue", "disabled_value": "anyValue", "enabled": false, "rollout_percentage": 100, "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5, "rollout_percentage": 100}], "collections": [{"collection_id": "CollectionID", "name": "Name"}], "isOverridden": true}], "properties": [{"name": "Name", "property_id": "PropertyID", "description": "Description", "type": "BOOLEAN", "format": "TEXT", "value": "anyValue", "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5}], "collections": [{"collection_id": "CollectionID", "name": "Name"}], "isOverridden": true}]}], "collections": [{"collection_id": "CollectionID", "name": "Name", "description": "Description", "tags": "Tags"}], "segments": [{"name": "Name", "segment_id": "SegmentID", "description": "Description", "tags": "Tags", "rules": [{"attribute_name": "AttributeName", "operator": "is", "values": ["Values"]}]}]}`)
				}))
			})
			It(`Invoke ListInstanceConfig successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.ListInstanceConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListInstanceConfigOptions model
				listInstanceConfigOptionsModel := new(appconfigurationv1.ListInstanceConfigOptions)
				listInstanceConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.ListInstanceConfig(listInstanceConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListInstanceConfig with error: Operation request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListInstanceConfigOptions model
				listInstanceConfigOptionsModel := new(appconfigurationv1.ListInstanceConfigOptions)
				listInstanceConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.ListInstanceConfig(listInstanceConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListInstanceConfig successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListInstanceConfigOptions model
				listInstanceConfigOptionsModel := new(appconfigurationv1.ListInstanceConfigOptions)
				listInstanceConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.ListInstanceConfig(listInstanceConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PromoteRestoreConfig(promoteRestoreConfigOptions *PromoteRestoreConfigOptions) - Operation response error`, func() {
		promoteRestoreConfigPath := "/config"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(promoteRestoreConfigPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.URL.Query()["git_config_id"]).To(Equal([]string{"git_config_id"}))
					Expect(req.URL.Query()["action"]).To(Equal([]string{"promote"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PromoteRestoreConfig with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the PromoteRestoreConfigOptions model
				promoteRestoreConfigOptionsModel := new(appconfigurationv1.PromoteRestoreConfigOptions)
				promoteRestoreConfigOptionsModel.GitConfigID = core.StringPtr("git_config_id")
				promoteRestoreConfigOptionsModel.Action = core.StringPtr("promote")
				promoteRestoreConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.PromoteRestoreConfig(promoteRestoreConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.PromoteRestoreConfig(promoteRestoreConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PromoteRestoreConfig(promoteRestoreConfigOptions *PromoteRestoreConfigOptions)`, func() {
		promoteRestoreConfigPath := "/config"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(promoteRestoreConfigPath))
					Expect(req.Method).To(Equal("PUT"))

					Expect(req.URL.Query()["git_config_id"]).To(Equal([]string{"git_config_id"}))
					Expect(req.URL.Query()["action"]).To(Equal([]string{"promote"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"git_commit_id": "GitCommitID", "git_commit_message": "GitCommitMessage", "last_sync_time": "2022-05-27T23:20:50.520Z"}`)
				}))
			})
			It(`Invoke PromoteRestoreConfig successfully with retries`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the PromoteRestoreConfigOptions model
				promoteRestoreConfigOptionsModel := new(appconfigurationv1.PromoteRestoreConfigOptions)
				promoteRestoreConfigOptionsModel.GitConfigID = core.StringPtr("git_config_id")
				promoteRestoreConfigOptionsModel.Action = core.StringPtr("promote")
				promoteRestoreConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := appConfigurationService.PromoteRestoreConfigWithContext(ctx, promoteRestoreConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr := appConfigurationService.PromoteRestoreConfig(promoteRestoreConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = appConfigurationService.PromoteRestoreConfigWithContext(ctx, promoteRestoreConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(promoteRestoreConfigPath))
					Expect(req.Method).To(Equal("PUT"))

					Expect(req.URL.Query()["git_config_id"]).To(Equal([]string{"git_config_id"}))
					Expect(req.URL.Query()["action"]).To(Equal([]string{"promote"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"git_commit_id": "GitCommitID", "git_commit_message": "GitCommitMessage", "last_sync_time": "2022-05-27T23:20:50.520Z"}`)
				}))
			})
			It(`Invoke PromoteRestoreConfig successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.PromoteRestoreConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PromoteRestoreConfigOptions model
				promoteRestoreConfigOptionsModel := new(appconfigurationv1.PromoteRestoreConfigOptions)
				promoteRestoreConfigOptionsModel.GitConfigID = core.StringPtr("git_config_id")
				promoteRestoreConfigOptionsModel.Action = core.StringPtr("promote")
				promoteRestoreConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.PromoteRestoreConfig(promoteRestoreConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PromoteRestoreConfig with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the PromoteRestoreConfigOptions model
				promoteRestoreConfigOptionsModel := new(appconfigurationv1.PromoteRestoreConfigOptions)
				promoteRestoreConfigOptionsModel.GitConfigID = core.StringPtr("git_config_id")
				promoteRestoreConfigOptionsModel.Action = core.StringPtr("promote")
				promoteRestoreConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.PromoteRestoreConfig(promoteRestoreConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PromoteRestoreConfigOptions model with no property values
				promoteRestoreConfigOptionsModelNew := new(appconfigurationv1.PromoteRestoreConfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.PromoteRestoreConfig(promoteRestoreConfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke PromoteRestoreConfig successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the PromoteRestoreConfigOptions model
				promoteRestoreConfigOptionsModel := new(appconfigurationv1.PromoteRestoreConfigOptions)
				promoteRestoreConfigOptionsModel.GitConfigID = core.StringPtr("git_config_id")
				promoteRestoreConfigOptionsModel.Action = core.StringPtr("promote")
				promoteRestoreConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := appConfigurationService.PromoteRestoreConfig(promoteRestoreConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			appConfigurationService, _ := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
				URL:           "http://appconfigurationv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCollection successfully`, func() {
				name := "testString"
				collectionID := "testString"
				_model, err := appConfigurationService.NewCollection(name, collectionID)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCollectionRef successfully`, func() {
				collectionID := "testString"
				_model, err := appConfigurationService.NewCollectionRef(collectionID)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCollectionUpdateRef successfully`, func() {
				collectionID := "testString"
				_model, err := appConfigurationService.NewCollectionUpdateRef(collectionID)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateCollectionOptions successfully`, func() {
				// Construct an instance of the CreateCollectionOptions model
				createCollectionOptionsName := "Web App Collection"
				createCollectionOptionsCollectionID := "web-app-collection"
				createCollectionOptionsModel := appConfigurationService.NewCreateCollectionOptions(createCollectionOptionsName, createCollectionOptionsCollectionID)
				createCollectionOptionsModel.SetName("Web App Collection")
				createCollectionOptionsModel.SetCollectionID("web-app-collection")
				createCollectionOptionsModel.SetDescription("Collection for Web application")
				createCollectionOptionsModel.SetTags("version: 1.1, pre-release")
				createCollectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCollectionOptionsModel).ToNot(BeNil())
				Expect(createCollectionOptionsModel.Name).To(Equal(core.StringPtr("Web App Collection")))
				Expect(createCollectionOptionsModel.CollectionID).To(Equal(core.StringPtr("web-app-collection")))
				Expect(createCollectionOptionsModel.Description).To(Equal(core.StringPtr("Collection for Web application")))
				Expect(createCollectionOptionsModel.Tags).To(Equal(core.StringPtr("version: 1.1, pre-release")))
				Expect(createCollectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateEnvironmentOptions successfully`, func() {
				// Construct an instance of the CreateEnvironmentOptions model
				createEnvironmentOptionsName := "Dev environment"
				createEnvironmentOptionsEnvironmentID := "dev-environment"
				createEnvironmentOptionsModel := appConfigurationService.NewCreateEnvironmentOptions(createEnvironmentOptionsName, createEnvironmentOptionsEnvironmentID)
				createEnvironmentOptionsModel.SetName("Dev environment")
				createEnvironmentOptionsModel.SetEnvironmentID("dev-environment")
				createEnvironmentOptionsModel.SetDescription("Dev environment description")
				createEnvironmentOptionsModel.SetTags("development")
				createEnvironmentOptionsModel.SetColorCode("#FDD13A")
				createEnvironmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createEnvironmentOptionsModel).ToNot(BeNil())
				Expect(createEnvironmentOptionsModel.Name).To(Equal(core.StringPtr("Dev environment")))
				Expect(createEnvironmentOptionsModel.EnvironmentID).To(Equal(core.StringPtr("dev-environment")))
				Expect(createEnvironmentOptionsModel.Description).To(Equal(core.StringPtr("Dev environment description")))
				Expect(createEnvironmentOptionsModel.Tags).To(Equal(core.StringPtr("development")))
				Expect(createEnvironmentOptionsModel.ColorCode).To(Equal(core.StringPtr("#FDD13A")))
				Expect(createEnvironmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateFeatureOptions successfully`, func() {
				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				Expect(targetSegmentsModel).ToNot(BeNil())
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}
				Expect(targetSegmentsModel.Segments).To(Equal([]string{"betausers", "premiumusers"}))

				// Construct an instance of the FeatureSegmentRule model
				featureSegmentRuleModel := new(appconfigurationv1.FeatureSegmentRule)
				Expect(featureSegmentRuleModel).ToNot(BeNil())
				featureSegmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				featureSegmentRuleModel.Value = "true"
				featureSegmentRuleModel.Order = core.Int64Ptr(int64(1))
				featureSegmentRuleModel.RolloutPercentage = core.Int64Ptr(int64(50))
				Expect(featureSegmentRuleModel.Rules).To(Equal([]appconfigurationv1.TargetSegments{*targetSegmentsModel}))
				Expect(featureSegmentRuleModel.Value).To(Equal("true"))
				Expect(featureSegmentRuleModel.Order).To(Equal(core.Int64Ptr(int64(1))))
				Expect(featureSegmentRuleModel.RolloutPercentage).To(Equal(core.Int64Ptr(int64(50))))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				Expect(collectionRefModel).ToNot(BeNil())
				collectionRefModel.CollectionID = core.StringPtr("ghzinc")
				Expect(collectionRefModel.CollectionID).To(Equal(core.StringPtr("ghzinc")))

				// Construct an instance of the CreateFeatureOptions model
				environmentID := "environment_id"
				createFeatureOptionsName := "Cycle Rentals"
				createFeatureOptionsFeatureID := "cycle-rentals"
				createFeatureOptionsType := "BOOLEAN"
				createFeatureOptionsEnabledValue := "true"
				createFeatureOptionsDisabledValue := "false"
				createFeatureOptionsModel := appConfigurationService.NewCreateFeatureOptions(environmentID, createFeatureOptionsName, createFeatureOptionsFeatureID, createFeatureOptionsType, createFeatureOptionsEnabledValue, createFeatureOptionsDisabledValue)
				createFeatureOptionsModel.SetEnvironmentID("environment_id")
				createFeatureOptionsModel.SetName("Cycle Rentals")
				createFeatureOptionsModel.SetFeatureID("cycle-rentals")
				createFeatureOptionsModel.SetType("BOOLEAN")
				createFeatureOptionsModel.SetEnabledValue("true")
				createFeatureOptionsModel.SetDisabledValue("false")
				createFeatureOptionsModel.SetDescription("Feature flag to enable Cycle Rentals")
				createFeatureOptionsModel.SetFormat("TEXT")
				createFeatureOptionsModel.SetEnabled(true)
				createFeatureOptionsModel.SetRolloutPercentage(int64(100))
				createFeatureOptionsModel.SetTags("version: 1.1, pre-release")
				createFeatureOptionsModel.SetSegmentRules([]appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel})
				createFeatureOptionsModel.SetCollections([]appconfigurationv1.CollectionRef{*collectionRefModel})
				createFeatureOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createFeatureOptionsModel).ToNot(BeNil())
				Expect(createFeatureOptionsModel.EnvironmentID).To(Equal(core.StringPtr("environment_id")))
				Expect(createFeatureOptionsModel.Name).To(Equal(core.StringPtr("Cycle Rentals")))
				Expect(createFeatureOptionsModel.FeatureID).To(Equal(core.StringPtr("cycle-rentals")))
				Expect(createFeatureOptionsModel.Type).To(Equal(core.StringPtr("BOOLEAN")))
				Expect(createFeatureOptionsModel.EnabledValue).To(Equal("true"))
				Expect(createFeatureOptionsModel.DisabledValue).To(Equal("false"))
				Expect(createFeatureOptionsModel.Description).To(Equal(core.StringPtr("Feature flag to enable Cycle Rentals")))
				Expect(createFeatureOptionsModel.Format).To(Equal(core.StringPtr("TEXT")))
				Expect(createFeatureOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(createFeatureOptionsModel.RolloutPercentage).To(Equal(core.Int64Ptr(int64(100))))
				Expect(createFeatureOptionsModel.Tags).To(Equal(core.StringPtr("version: 1.1, pre-release")))
				Expect(createFeatureOptionsModel.SegmentRules).To(Equal([]appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel}))
				Expect(createFeatureOptionsModel.Collections).To(Equal([]appconfigurationv1.CollectionRef{*collectionRefModel}))
				Expect(createFeatureOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateGitconfigOptions successfully`, func() {
				// Construct an instance of the CreateGitconfigOptions model
				createGitconfigOptionsGitConfigName := "boot-strap-configuration"
				createGitconfigOptionsGitConfigID := "boot-strap-configuration"
				createGitconfigOptionsCollectionID := "web-app-collection"
				createGitconfigOptionsEnvironmentID := "dev"
				createGitconfigOptionsGitURL := "https://github.ibm.com/api/v3/repos/jhondoe-owner/my-test-repo"
				createGitconfigOptionsGitBranch := "main"
				createGitconfigOptionsGitFilePath := "code/development/README.json"
				createGitconfigOptionsGitToken := "61a792eahhGHji223jijb55a6cfdd4d5cde4c8a67esjjhjhHVH"
				createGitconfigOptionsModel := appConfigurationService.NewCreateGitconfigOptions(createGitconfigOptionsGitConfigName, createGitconfigOptionsGitConfigID, createGitconfigOptionsCollectionID, createGitconfigOptionsEnvironmentID, createGitconfigOptionsGitURL, createGitconfigOptionsGitBranch, createGitconfigOptionsGitFilePath, createGitconfigOptionsGitToken)
				createGitconfigOptionsModel.SetGitConfigName("boot-strap-configuration")
				createGitconfigOptionsModel.SetGitConfigID("boot-strap-configuration")
				createGitconfigOptionsModel.SetCollectionID("web-app-collection")
				createGitconfigOptionsModel.SetEnvironmentID("dev")
				createGitconfigOptionsModel.SetGitURL("https://github.ibm.com/api/v3/repos/jhondoe-owner/my-test-repo")
				createGitconfigOptionsModel.SetGitBranch("main")
				createGitconfigOptionsModel.SetGitFilePath("code/development/README.json")
				createGitconfigOptionsModel.SetGitToken("61a792eahhGHji223jijb55a6cfdd4d5cde4c8a67esjjhjhHVH")
				createGitconfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createGitconfigOptionsModel).ToNot(BeNil())
				Expect(createGitconfigOptionsModel.GitConfigName).To(Equal(core.StringPtr("boot-strap-configuration")))
				Expect(createGitconfigOptionsModel.GitConfigID).To(Equal(core.StringPtr("boot-strap-configuration")))
				Expect(createGitconfigOptionsModel.CollectionID).To(Equal(core.StringPtr("web-app-collection")))
				Expect(createGitconfigOptionsModel.EnvironmentID).To(Equal(core.StringPtr("dev")))
				Expect(createGitconfigOptionsModel.GitURL).To(Equal(core.StringPtr("https://github.ibm.com/api/v3/repos/jhondoe-owner/my-test-repo")))
				Expect(createGitconfigOptionsModel.GitBranch).To(Equal(core.StringPtr("main")))
				Expect(createGitconfigOptionsModel.GitFilePath).To(Equal(core.StringPtr("code/development/README.json")))
				Expect(createGitconfigOptionsModel.GitToken).To(Equal(core.StringPtr("61a792eahhGHji223jijb55a6cfdd4d5cde4c8a67esjjhjhHVH")))
				Expect(createGitconfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreatePropertyOptions successfully`, func() {
				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				Expect(targetSegmentsModel).ToNot(BeNil())
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}
				Expect(targetSegmentsModel.Segments).To(Equal([]string{"betausers", "premiumusers"}))

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				Expect(segmentRuleModel).ToNot(BeNil())
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = "true"
				segmentRuleModel.Order = core.Int64Ptr(int64(1))
				Expect(segmentRuleModel.Rules).To(Equal([]appconfigurationv1.TargetSegments{*targetSegmentsModel}))
				Expect(segmentRuleModel.Value).To(Equal("true"))
				Expect(segmentRuleModel.Order).To(Equal(core.Int64Ptr(int64(1))))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				Expect(collectionRefModel).ToNot(BeNil())
				collectionRefModel.CollectionID = core.StringPtr("ghzinc")
				Expect(collectionRefModel.CollectionID).To(Equal(core.StringPtr("ghzinc")))

				// Construct an instance of the CreatePropertyOptions model
				environmentID := "environment_id"
				createPropertyOptionsName := "Email property"
				createPropertyOptionsPropertyID := "email-property"
				createPropertyOptionsType := "BOOLEAN"
				createPropertyOptionsValue := "true"
				createPropertyOptionsModel := appConfigurationService.NewCreatePropertyOptions(environmentID, createPropertyOptionsName, createPropertyOptionsPropertyID, createPropertyOptionsType, createPropertyOptionsValue)
				createPropertyOptionsModel.SetEnvironmentID("environment_id")
				createPropertyOptionsModel.SetName("Email property")
				createPropertyOptionsModel.SetPropertyID("email-property")
				createPropertyOptionsModel.SetType("BOOLEAN")
				createPropertyOptionsModel.SetValue("true")
				createPropertyOptionsModel.SetDescription("Property for email")
				createPropertyOptionsModel.SetFormat("TEXT")
				createPropertyOptionsModel.SetTags("version: 1.1, pre-release")
				createPropertyOptionsModel.SetSegmentRules([]appconfigurationv1.SegmentRule{*segmentRuleModel})
				createPropertyOptionsModel.SetCollections([]appconfigurationv1.CollectionRef{*collectionRefModel})
				createPropertyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createPropertyOptionsModel).ToNot(BeNil())
				Expect(createPropertyOptionsModel.EnvironmentID).To(Equal(core.StringPtr("environment_id")))
				Expect(createPropertyOptionsModel.Name).To(Equal(core.StringPtr("Email property")))
				Expect(createPropertyOptionsModel.PropertyID).To(Equal(core.StringPtr("email-property")))
				Expect(createPropertyOptionsModel.Type).To(Equal(core.StringPtr("BOOLEAN")))
				Expect(createPropertyOptionsModel.Value).To(Equal("true"))
				Expect(createPropertyOptionsModel.Description).To(Equal(core.StringPtr("Property for email")))
				Expect(createPropertyOptionsModel.Format).To(Equal(core.StringPtr("TEXT")))
				Expect(createPropertyOptionsModel.Tags).To(Equal(core.StringPtr("version: 1.1, pre-release")))
				Expect(createPropertyOptionsModel.SegmentRules).To(Equal([]appconfigurationv1.SegmentRule{*segmentRuleModel}))
				Expect(createPropertyOptionsModel.Collections).To(Equal([]appconfigurationv1.CollectionRef{*collectionRefModel}))
				Expect(createPropertyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateSegmentOptions successfully`, func() {
				// Construct an instance of the Rule model
				ruleModel := new(appconfigurationv1.Rule)
				Expect(ruleModel).ToNot(BeNil())
				ruleModel.AttributeName = core.StringPtr("email")
				ruleModel.Operator = core.StringPtr("endsWith")
				ruleModel.Values = []string{"@in.mnc.com", "@us.mnc.com"}
				Expect(ruleModel.AttributeName).To(Equal(core.StringPtr("email")))
				Expect(ruleModel.Operator).To(Equal(core.StringPtr("endsWith")))
				Expect(ruleModel.Values).To(Equal([]string{"@in.mnc.com", "@us.mnc.com"}))

				// Construct an instance of the CreateSegmentOptions model
				createSegmentOptionsName := "Beta Users"
				createSegmentOptionsSegmentID := "beta-users"
				createSegmentOptionsRules := []appconfigurationv1.Rule{}
				createSegmentOptionsModel := appConfigurationService.NewCreateSegmentOptions(createSegmentOptionsName, createSegmentOptionsSegmentID, createSegmentOptionsRules)
				createSegmentOptionsModel.SetName("Beta Users")
				createSegmentOptionsModel.SetSegmentID("beta-users")
				createSegmentOptionsModel.SetRules([]appconfigurationv1.Rule{*ruleModel})
				createSegmentOptionsModel.SetDescription("Segment containing the beta users")
				createSegmentOptionsModel.SetTags("version: 1.1, stage")
				createSegmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createSegmentOptionsModel).ToNot(BeNil())
				Expect(createSegmentOptionsModel.Name).To(Equal(core.StringPtr("Beta Users")))
				Expect(createSegmentOptionsModel.SegmentID).To(Equal(core.StringPtr("beta-users")))
				Expect(createSegmentOptionsModel.Rules).To(Equal([]appconfigurationv1.Rule{*ruleModel}))
				Expect(createSegmentOptionsModel.Description).To(Equal(core.StringPtr("Segment containing the beta users")))
				Expect(createSegmentOptionsModel.Tags).To(Equal(core.StringPtr("version: 1.1, stage")))
				Expect(createSegmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateWorkflowconfigOptions successfully`, func() {
				// Construct an instance of the ExternalServiceNowCredentials model
				externalServiceNowCredentialsModel := new(appconfigurationv1.ExternalServiceNowCredentials)
				Expect(externalServiceNowCredentialsModel).ToNot(BeNil())
				externalServiceNowCredentialsModel.Username = core.StringPtr("admin")
				externalServiceNowCredentialsModel.Password = core.StringPtr("testString")
				externalServiceNowCredentialsModel.ClientID = core.StringPtr("f7b6379b55d08210f8ree233afc7256d")
				externalServiceNowCredentialsModel.ClientSecret = core.StringPtr("testString")
				Expect(externalServiceNowCredentialsModel.Username).To(Equal(core.StringPtr("admin")))
				Expect(externalServiceNowCredentialsModel.Password).To(Equal(core.StringPtr("testString")))
				Expect(externalServiceNowCredentialsModel.ClientID).To(Equal(core.StringPtr("f7b6379b55d08210f8ree233afc7256d")))
				Expect(externalServiceNowCredentialsModel.ClientSecret).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateWorkflowConfigExternalServiceNow model
				createWorkflowConfigModel := new(appconfigurationv1.CreateWorkflowConfigExternalServiceNow)
				Expect(createWorkflowConfigModel).ToNot(BeNil())
				createWorkflowConfigModel.WorkflowURL = core.StringPtr("testString")
				createWorkflowConfigModel.ApprovalGroupName = core.StringPtr("testString")
				createWorkflowConfigModel.ApprovalExpiration = core.Int64Ptr(int64(1))
				createWorkflowConfigModel.WorkflowCredentials = externalServiceNowCredentialsModel
				createWorkflowConfigModel.Enabled = core.BoolPtr(false)
				Expect(createWorkflowConfigModel.WorkflowURL).To(Equal(core.StringPtr("testString")))
				Expect(createWorkflowConfigModel.ApprovalGroupName).To(Equal(core.StringPtr("testString")))
				Expect(createWorkflowConfigModel.ApprovalExpiration).To(Equal(core.Int64Ptr(int64(1))))
				Expect(createWorkflowConfigModel.WorkflowCredentials).To(Equal(externalServiceNowCredentialsModel))
				Expect(createWorkflowConfigModel.Enabled).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the CreateWorkflowconfigOptions model
				environmentID := "environment_id"
				var workflowConfig appconfigurationv1.CreateWorkflowConfigIntf = nil
				createWorkflowconfigOptionsModel := appConfigurationService.NewCreateWorkflowconfigOptions(environmentID, workflowConfig)
				createWorkflowconfigOptionsModel.SetEnvironmentID("environment_id")
				createWorkflowconfigOptionsModel.SetWorkflowConfig(createWorkflowConfigModel)
				createWorkflowconfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createWorkflowconfigOptionsModel).ToNot(BeNil())
				Expect(createWorkflowconfigOptionsModel.EnvironmentID).To(Equal(core.StringPtr("environment_id")))
				Expect(createWorkflowconfigOptionsModel.WorkflowConfig).To(Equal(createWorkflowConfigModel))
				Expect(createWorkflowconfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteCollectionOptions successfully`, func() {
				// Construct an instance of the DeleteCollectionOptions model
				collectionID := "collection_id"
				deleteCollectionOptionsModel := appConfigurationService.NewDeleteCollectionOptions(collectionID)
				deleteCollectionOptionsModel.SetCollectionID("collection_id")
				deleteCollectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteCollectionOptionsModel).ToNot(BeNil())
				Expect(deleteCollectionOptionsModel.CollectionID).To(Equal(core.StringPtr("collection_id")))
				Expect(deleteCollectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteEnvironmentOptions successfully`, func() {
				// Construct an instance of the DeleteEnvironmentOptions model
				environmentID := "environment_id"
				deleteEnvironmentOptionsModel := appConfigurationService.NewDeleteEnvironmentOptions(environmentID)
				deleteEnvironmentOptionsModel.SetEnvironmentID("environment_id")
				deleteEnvironmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteEnvironmentOptionsModel).ToNot(BeNil())
				Expect(deleteEnvironmentOptionsModel.EnvironmentID).To(Equal(core.StringPtr("environment_id")))
				Expect(deleteEnvironmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteFeatureOptions successfully`, func() {
				// Construct an instance of the DeleteFeatureOptions model
				environmentID := "environment_id"
				featureID := "feature_id"
				deleteFeatureOptionsModel := appConfigurationService.NewDeleteFeatureOptions(environmentID, featureID)
				deleteFeatureOptionsModel.SetEnvironmentID("environment_id")
				deleteFeatureOptionsModel.SetFeatureID("feature_id")
				deleteFeatureOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteFeatureOptionsModel).ToNot(BeNil())
				Expect(deleteFeatureOptionsModel.EnvironmentID).To(Equal(core.StringPtr("environment_id")))
				Expect(deleteFeatureOptionsModel.FeatureID).To(Equal(core.StringPtr("feature_id")))
				Expect(deleteFeatureOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteGitconfigOptions successfully`, func() {
				// Construct an instance of the DeleteGitconfigOptions model
				gitConfigID := "git_config_id"
				deleteGitconfigOptionsModel := appConfigurationService.NewDeleteGitconfigOptions(gitConfigID)
				deleteGitconfigOptionsModel.SetGitConfigID("git_config_id")
				deleteGitconfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteGitconfigOptionsModel).ToNot(BeNil())
				Expect(deleteGitconfigOptionsModel.GitConfigID).To(Equal(core.StringPtr("git_config_id")))
				Expect(deleteGitconfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeletePropertyOptions successfully`, func() {
				// Construct an instance of the DeletePropertyOptions model
				environmentID := "environment_id"
				propertyID := "property_id"
				deletePropertyOptionsModel := appConfigurationService.NewDeletePropertyOptions(environmentID, propertyID)
				deletePropertyOptionsModel.SetEnvironmentID("environment_id")
				deletePropertyOptionsModel.SetPropertyID("property_id")
				deletePropertyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deletePropertyOptionsModel).ToNot(BeNil())
				Expect(deletePropertyOptionsModel.EnvironmentID).To(Equal(core.StringPtr("environment_id")))
				Expect(deletePropertyOptionsModel.PropertyID).To(Equal(core.StringPtr("property_id")))
				Expect(deletePropertyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteSegmentOptions successfully`, func() {
				// Construct an instance of the DeleteSegmentOptions model
				segmentID := "segment_id"
				deleteSegmentOptionsModel := appConfigurationService.NewDeleteSegmentOptions(segmentID)
				deleteSegmentOptionsModel.SetSegmentID("segment_id")
				deleteSegmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteSegmentOptionsModel).ToNot(BeNil())
				Expect(deleteSegmentOptionsModel.SegmentID).To(Equal(core.StringPtr("segment_id")))
				Expect(deleteSegmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteWorkflowconfigOptions successfully`, func() {
				// Construct an instance of the DeleteWorkflowconfigOptions model
				environmentID := "environment_id"
				deleteWorkflowconfigOptionsModel := appConfigurationService.NewDeleteWorkflowconfigOptions(environmentID)
				deleteWorkflowconfigOptionsModel.SetEnvironmentID("environment_id")
				deleteWorkflowconfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteWorkflowconfigOptionsModel).ToNot(BeNil())
				Expect(deleteWorkflowconfigOptionsModel.EnvironmentID).To(Equal(core.StringPtr("environment_id")))
				Expect(deleteWorkflowconfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewEnvironment successfully`, func() {
				name := "testString"
				environmentID := "testString"
				_model, err := appConfigurationService.NewEnvironment(name, environmentID)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewExternalServiceNowCredentials successfully`, func() {
				username := "admin"
				password := "testString"
				clientID := "f7b6379b55d08210f8ree233afc7256d"
				clientSecret := "testString"
				_model, err := appConfigurationService.NewExternalServiceNowCredentials(username, password, clientID, clientSecret)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewFeatureOutput successfully`, func() {
				featureID := "testString"
				name := "testString"
				_model, err := appConfigurationService.NewFeatureOutput(featureID, name)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewFeatureSegmentRule successfully`, func() {
				rules := []appconfigurationv1.TargetSegments{}
				value := "testString"
				order := int64(38)
				_model, err := appConfigurationService.NewFeatureSegmentRule(rules, value, order)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGetCollectionOptions successfully`, func() {
				// Construct an instance of the GetCollectionOptions model
				collectionID := "collection_id"
				getCollectionOptionsModel := appConfigurationService.NewGetCollectionOptions(collectionID)
				getCollectionOptionsModel.SetCollectionID("collection_id")
				getCollectionOptionsModel.SetExpand(true)
				getCollectionOptionsModel.SetInclude([]string{"features", "properties", "snapshots"})
				getCollectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCollectionOptionsModel).ToNot(BeNil())
				Expect(getCollectionOptionsModel.CollectionID).To(Equal(core.StringPtr("collection_id")))
				Expect(getCollectionOptionsModel.Expand).To(Equal(core.BoolPtr(true)))
				Expect(getCollectionOptionsModel.Include).To(Equal([]string{"features", "properties", "snapshots"}))
				Expect(getCollectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetEnvironmentOptions successfully`, func() {
				// Construct an instance of the GetEnvironmentOptions model
				environmentID := "environment_id"
				getEnvironmentOptionsModel := appConfigurationService.NewGetEnvironmentOptions(environmentID)
				getEnvironmentOptionsModel.SetEnvironmentID("environment_id")
				getEnvironmentOptionsModel.SetExpand(true)
				getEnvironmentOptionsModel.SetInclude([]string{"features", "properties", "snapshots"})
				getEnvironmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getEnvironmentOptionsModel).ToNot(BeNil())
				Expect(getEnvironmentOptionsModel.EnvironmentID).To(Equal(core.StringPtr("environment_id")))
				Expect(getEnvironmentOptionsModel.Expand).To(Equal(core.BoolPtr(true)))
				Expect(getEnvironmentOptionsModel.Include).To(Equal([]string{"features", "properties", "snapshots"}))
				Expect(getEnvironmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetFeatureOptions successfully`, func() {
				// Construct an instance of the GetFeatureOptions model
				environmentID := "environment_id"
				featureID := "feature_id"
				getFeatureOptionsModel := appConfigurationService.NewGetFeatureOptions(environmentID, featureID)
				getFeatureOptionsModel.SetEnvironmentID("environment_id")
				getFeatureOptionsModel.SetFeatureID("feature_id")
				getFeatureOptionsModel.SetInclude([]string{"collections", "rules", "change_request"})
				getFeatureOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getFeatureOptionsModel).ToNot(BeNil())
				Expect(getFeatureOptionsModel.EnvironmentID).To(Equal(core.StringPtr("environment_id")))
				Expect(getFeatureOptionsModel.FeatureID).To(Equal(core.StringPtr("feature_id")))
				Expect(getFeatureOptionsModel.Include).To(Equal([]string{"collections", "rules", "change_request"}))
				Expect(getFeatureOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetGitconfigOptions successfully`, func() {
				// Construct an instance of the GetGitconfigOptions model
				gitConfigID := "git_config_id"
				getGitconfigOptionsModel := appConfigurationService.NewGetGitconfigOptions(gitConfigID)
				getGitconfigOptionsModel.SetGitConfigID("git_config_id")
				getGitconfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getGitconfigOptionsModel).ToNot(BeNil())
				Expect(getGitconfigOptionsModel.GitConfigID).To(Equal(core.StringPtr("git_config_id")))
				Expect(getGitconfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPropertyOptions successfully`, func() {
				// Construct an instance of the GetPropertyOptions model
				environmentID := "environment_id"
				propertyID := "property_id"
				getPropertyOptionsModel := appConfigurationService.NewGetPropertyOptions(environmentID, propertyID)
				getPropertyOptionsModel.SetEnvironmentID("environment_id")
				getPropertyOptionsModel.SetPropertyID("property_id")
				getPropertyOptionsModel.SetInclude([]string{"collections", "rules"})
				getPropertyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPropertyOptionsModel).ToNot(BeNil())
				Expect(getPropertyOptionsModel.EnvironmentID).To(Equal(core.StringPtr("environment_id")))
				Expect(getPropertyOptionsModel.PropertyID).To(Equal(core.StringPtr("property_id")))
				Expect(getPropertyOptionsModel.Include).To(Equal([]string{"collections", "rules"}))
				Expect(getPropertyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSegmentOptions successfully`, func() {
				// Construct an instance of the GetSegmentOptions model
				segmentID := "segment_id"
				getSegmentOptionsModel := appConfigurationService.NewGetSegmentOptions(segmentID)
				getSegmentOptionsModel.SetSegmentID("segment_id")
				getSegmentOptionsModel.SetInclude([]string{"features", "properties"})
				getSegmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSegmentOptionsModel).ToNot(BeNil())
				Expect(getSegmentOptionsModel.SegmentID).To(Equal(core.StringPtr("segment_id")))
				Expect(getSegmentOptionsModel.Include).To(Equal([]string{"features", "properties"}))
				Expect(getSegmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewImportCollectionSchema successfully`, func() {
				collectionID := "testString"
				name := "testString"
				_model, err := appConfigurationService.NewImportCollectionSchema(collectionID, name)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewImportConfigOptions successfully`, func() {
				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				Expect(targetSegmentsModel).ToNot(BeNil())
				targetSegmentsModel.Segments = []string{"testString"}
				Expect(targetSegmentsModel.Segments).To(Equal([]string{"testString"}))

				// Construct an instance of the FeatureSegmentRule model
				featureSegmentRuleModel := new(appconfigurationv1.FeatureSegmentRule)
				Expect(featureSegmentRuleModel).ToNot(BeNil())
				featureSegmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				featureSegmentRuleModel.Value = "testString"
				featureSegmentRuleModel.Order = core.Int64Ptr(int64(38))
				featureSegmentRuleModel.RolloutPercentage = core.Int64Ptr(int64(100))
				Expect(featureSegmentRuleModel.Rules).To(Equal([]appconfigurationv1.TargetSegments{*targetSegmentsModel}))
				Expect(featureSegmentRuleModel.Value).To(Equal("testString"))
				Expect(featureSegmentRuleModel.Order).To(Equal(core.Int64Ptr(int64(38))))
				Expect(featureSegmentRuleModel.RolloutPercentage).To(Equal(core.Int64Ptr(int64(100))))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				Expect(collectionRefModel).ToNot(BeNil())
				collectionRefModel.CollectionID = core.StringPtr("web-app")
				Expect(collectionRefModel.CollectionID).To(Equal(core.StringPtr("web-app")))

				// Construct an instance of the ImportFeatureRequestBody model
				importFeatureRequestBodyModel := new(appconfigurationv1.ImportFeatureRequestBody)
				Expect(importFeatureRequestBodyModel).ToNot(BeNil())
				importFeatureRequestBodyModel.Name = core.StringPtr("Cycle Rentals")
				importFeatureRequestBodyModel.FeatureID = core.StringPtr("cycle-rentals")
				importFeatureRequestBodyModel.Description = core.StringPtr("testString")
				importFeatureRequestBodyModel.Type = core.StringPtr("NUMERIC")
				importFeatureRequestBodyModel.Format = core.StringPtr("TEXT")
				importFeatureRequestBodyModel.EnabledValue = "1"
				importFeatureRequestBodyModel.DisabledValue = "2"
				importFeatureRequestBodyModel.Enabled = core.BoolPtr(true)
				importFeatureRequestBodyModel.RolloutPercentage = core.Int64Ptr(int64(100))
				importFeatureRequestBodyModel.Tags = core.StringPtr("testString")
				importFeatureRequestBodyModel.SegmentRules = []appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel}
				importFeatureRequestBodyModel.Collections = []appconfigurationv1.CollectionRef{*collectionRefModel}
				importFeatureRequestBodyModel.IsOverridden = core.BoolPtr(true)
				Expect(importFeatureRequestBodyModel.Name).To(Equal(core.StringPtr("Cycle Rentals")))
				Expect(importFeatureRequestBodyModel.FeatureID).To(Equal(core.StringPtr("cycle-rentals")))
				Expect(importFeatureRequestBodyModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(importFeatureRequestBodyModel.Type).To(Equal(core.StringPtr("NUMERIC")))
				Expect(importFeatureRequestBodyModel.Format).To(Equal(core.StringPtr("TEXT")))
				Expect(importFeatureRequestBodyModel.EnabledValue).To(Equal("1"))
				Expect(importFeatureRequestBodyModel.DisabledValue).To(Equal("2"))
				Expect(importFeatureRequestBodyModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(importFeatureRequestBodyModel.RolloutPercentage).To(Equal(core.Int64Ptr(int64(100))))
				Expect(importFeatureRequestBodyModel.Tags).To(Equal(core.StringPtr("testString")))
				Expect(importFeatureRequestBodyModel.SegmentRules).To(Equal([]appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel}))
				Expect(importFeatureRequestBodyModel.Collections).To(Equal([]appconfigurationv1.CollectionRef{*collectionRefModel}))
				Expect(importFeatureRequestBodyModel.IsOverridden).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				Expect(segmentRuleModel).ToNot(BeNil())
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = "200"
				segmentRuleModel.Order = core.Int64Ptr(int64(1))
				Expect(segmentRuleModel.Rules).To(Equal([]appconfigurationv1.TargetSegments{*targetSegmentsModel}))
				Expect(segmentRuleModel.Value).To(Equal("200"))
				Expect(segmentRuleModel.Order).To(Equal(core.Int64Ptr(int64(1))))

				// Construct an instance of the ImportPropertyRequestBody model
				importPropertyRequestBodyModel := new(appconfigurationv1.ImportPropertyRequestBody)
				Expect(importPropertyRequestBodyModel).ToNot(BeNil())
				importPropertyRequestBodyModel.Name = core.StringPtr("Daily Discount")
				importPropertyRequestBodyModel.PropertyID = core.StringPtr("daily_discount")
				importPropertyRequestBodyModel.Description = core.StringPtr("testString")
				importPropertyRequestBodyModel.Type = core.StringPtr("NUMERIC")
				importPropertyRequestBodyModel.Format = core.StringPtr("TEXT")
				importPropertyRequestBodyModel.Value = "100"
				importPropertyRequestBodyModel.Tags = core.StringPtr("pre-release, v1.2")
				importPropertyRequestBodyModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				importPropertyRequestBodyModel.Collections = []appconfigurationv1.CollectionRef{*collectionRefModel}
				importPropertyRequestBodyModel.IsOverridden = core.BoolPtr(true)
				Expect(importPropertyRequestBodyModel.Name).To(Equal(core.StringPtr("Daily Discount")))
				Expect(importPropertyRequestBodyModel.PropertyID).To(Equal(core.StringPtr("daily_discount")))
				Expect(importPropertyRequestBodyModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(importPropertyRequestBodyModel.Type).To(Equal(core.StringPtr("NUMERIC")))
				Expect(importPropertyRequestBodyModel.Format).To(Equal(core.StringPtr("TEXT")))
				Expect(importPropertyRequestBodyModel.Value).To(Equal("100"))
				Expect(importPropertyRequestBodyModel.Tags).To(Equal(core.StringPtr("pre-release, v1.2")))
				Expect(importPropertyRequestBodyModel.SegmentRules).To(Equal([]appconfigurationv1.SegmentRule{*segmentRuleModel}))
				Expect(importPropertyRequestBodyModel.Collections).To(Equal([]appconfigurationv1.CollectionRef{*collectionRefModel}))
				Expect(importPropertyRequestBodyModel.IsOverridden).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the ImportEnvironmentSchema model
				importEnvironmentSchemaModel := new(appconfigurationv1.ImportEnvironmentSchema)
				Expect(importEnvironmentSchemaModel).ToNot(BeNil())
				importEnvironmentSchemaModel.Name = core.StringPtr("Dev")
				importEnvironmentSchemaModel.EnvironmentID = core.StringPtr("dev")
				importEnvironmentSchemaModel.Description = core.StringPtr("Environment created on instance creation")
				importEnvironmentSchemaModel.Tags = core.StringPtr("testString")
				importEnvironmentSchemaModel.ColorCode = core.StringPtr("#FDD13A")
				importEnvironmentSchemaModel.Features = []appconfigurationv1.ImportFeatureRequestBody{*importFeatureRequestBodyModel}
				importEnvironmentSchemaModel.Properties = []appconfigurationv1.ImportPropertyRequestBody{*importPropertyRequestBodyModel}
				Expect(importEnvironmentSchemaModel.Name).To(Equal(core.StringPtr("Dev")))
				Expect(importEnvironmentSchemaModel.EnvironmentID).To(Equal(core.StringPtr("dev")))
				Expect(importEnvironmentSchemaModel.Description).To(Equal(core.StringPtr("Environment created on instance creation")))
				Expect(importEnvironmentSchemaModel.Tags).To(Equal(core.StringPtr("testString")))
				Expect(importEnvironmentSchemaModel.ColorCode).To(Equal(core.StringPtr("#FDD13A")))
				Expect(importEnvironmentSchemaModel.Features).To(Equal([]appconfigurationv1.ImportFeatureRequestBody{*importFeatureRequestBodyModel}))
				Expect(importEnvironmentSchemaModel.Properties).To(Equal([]appconfigurationv1.ImportPropertyRequestBody{*importPropertyRequestBodyModel}))

				// Construct an instance of the ImportCollectionSchema model
				importCollectionSchemaModel := new(appconfigurationv1.ImportCollectionSchema)
				Expect(importCollectionSchemaModel).ToNot(BeNil())
				importCollectionSchemaModel.CollectionID = core.StringPtr("web-app")
				importCollectionSchemaModel.Name = core.StringPtr("web-app")
				importCollectionSchemaModel.Description = core.StringPtr("web app collection")
				importCollectionSchemaModel.Tags = core.StringPtr("v1")
				Expect(importCollectionSchemaModel.CollectionID).To(Equal(core.StringPtr("web-app")))
				Expect(importCollectionSchemaModel.Name).To(Equal(core.StringPtr("web-app")))
				Expect(importCollectionSchemaModel.Description).To(Equal(core.StringPtr("web app collection")))
				Expect(importCollectionSchemaModel.Tags).To(Equal(core.StringPtr("v1")))

				// Construct an instance of the Rule model
				ruleModel := new(appconfigurationv1.Rule)
				Expect(ruleModel).ToNot(BeNil())
				ruleModel.AttributeName = core.StringPtr("email")
				ruleModel.Operator = core.StringPtr("is")
				ruleModel.Values = []string{"john@bluecharge.com", "alice@bluecharge.com"}
				Expect(ruleModel.AttributeName).To(Equal(core.StringPtr("email")))
				Expect(ruleModel.Operator).To(Equal(core.StringPtr("is")))
				Expect(ruleModel.Values).To(Equal([]string{"john@bluecharge.com", "alice@bluecharge.com"}))

				// Construct an instance of the ImportSegmentSchema model
				importSegmentSchemaModel := new(appconfigurationv1.ImportSegmentSchema)
				Expect(importSegmentSchemaModel).ToNot(BeNil())
				importSegmentSchemaModel.Name = core.StringPtr("Testers")
				importSegmentSchemaModel.SegmentID = core.StringPtr("khpwj68h")
				importSegmentSchemaModel.Description = core.StringPtr("Testers")
				importSegmentSchemaModel.Tags = core.StringPtr("test")
				importSegmentSchemaModel.Rules = []appconfigurationv1.Rule{*ruleModel}
				Expect(importSegmentSchemaModel.Name).To(Equal(core.StringPtr("Testers")))
				Expect(importSegmentSchemaModel.SegmentID).To(Equal(core.StringPtr("khpwj68h")))
				Expect(importSegmentSchemaModel.Description).To(Equal(core.StringPtr("Testers")))
				Expect(importSegmentSchemaModel.Tags).To(Equal(core.StringPtr("test")))
				Expect(importSegmentSchemaModel.Rules).To(Equal([]appconfigurationv1.Rule{*ruleModel}))

				// Construct an instance of the ImportConfigOptions model
				importConfigOptionsModel := appConfigurationService.NewImportConfigOptions()
				importConfigOptionsModel.SetEnvironments([]appconfigurationv1.ImportEnvironmentSchema{*importEnvironmentSchemaModel})
				importConfigOptionsModel.SetCollections([]appconfigurationv1.ImportCollectionSchema{*importCollectionSchemaModel})
				importConfigOptionsModel.SetSegments([]appconfigurationv1.ImportSegmentSchema{*importSegmentSchemaModel})
				importConfigOptionsModel.SetClean("true")
				importConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(importConfigOptionsModel).ToNot(BeNil())
				Expect(importConfigOptionsModel.Environments).To(Equal([]appconfigurationv1.ImportEnvironmentSchema{*importEnvironmentSchemaModel}))
				Expect(importConfigOptionsModel.Collections).To(Equal([]appconfigurationv1.ImportCollectionSchema{*importCollectionSchemaModel}))
				Expect(importConfigOptionsModel.Segments).To(Equal([]appconfigurationv1.ImportSegmentSchema{*importSegmentSchemaModel}))
				Expect(importConfigOptionsModel.Clean).To(Equal(core.StringPtr("true")))
				Expect(importConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewImportEnvironmentSchema successfully`, func() {
				name := "testString"
				environmentID := "testString"
				_model, err := appConfigurationService.NewImportEnvironmentSchema(name, environmentID)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewImportFeatureRequestBody successfully`, func() {
				name := "testString"
				featureID := "testString"
				typeVar := "BOOLEAN"
				enabledValue := "testString"
				disabledValue := "testString"
				isOverridden := true
				_model, err := appConfigurationService.NewImportFeatureRequestBody(name, featureID, typeVar, enabledValue, disabledValue, isOverridden)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewImportPropertyRequestBody successfully`, func() {
				name := "testString"
				propertyID := "testString"
				typeVar := "BOOLEAN"
				value := "testString"
				isOverridden := true
				_model, err := appConfigurationService.NewImportPropertyRequestBody(name, propertyID, typeVar, value, isOverridden)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewImportSegmentSchema successfully`, func() {
				name := "testString"
				segmentID := "testString"
				rules := []appconfigurationv1.Rule{}
				_model, err := appConfigurationService.NewImportSegmentSchema(name, segmentID, rules)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewListCollectionsOptions successfully`, func() {
				// Construct an instance of the ListCollectionsOptions model
				listCollectionsOptionsModel := appConfigurationService.NewListCollectionsOptions()
				listCollectionsOptionsModel.SetExpand(true)
				listCollectionsOptionsModel.SetSort("created_time")
				listCollectionsOptionsModel.SetTags("version 1.1,pre-release")
				listCollectionsOptionsModel.SetFeatures([]string{"my-feature-id", "cycle-rentals"})
				listCollectionsOptionsModel.SetProperties([]string{"my-property-id", "email-property"})
				listCollectionsOptionsModel.SetInclude([]string{"features", "properties", "snapshots"})
				listCollectionsOptionsModel.SetLimit(int64(10))
				listCollectionsOptionsModel.SetOffset(int64(0))
				listCollectionsOptionsModel.SetSearch("test tag")
				listCollectionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listCollectionsOptionsModel).ToNot(BeNil())
				Expect(listCollectionsOptionsModel.Expand).To(Equal(core.BoolPtr(true)))
				Expect(listCollectionsOptionsModel.Sort).To(Equal(core.StringPtr("created_time")))
				Expect(listCollectionsOptionsModel.Tags).To(Equal(core.StringPtr("version 1.1,pre-release")))
				Expect(listCollectionsOptionsModel.Features).To(Equal([]string{"my-feature-id", "cycle-rentals"}))
				Expect(listCollectionsOptionsModel.Properties).To(Equal([]string{"my-property-id", "email-property"}))
				Expect(listCollectionsOptionsModel.Include).To(Equal([]string{"features", "properties", "snapshots"}))
				Expect(listCollectionsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listCollectionsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listCollectionsOptionsModel.Search).To(Equal(core.StringPtr("test tag")))
				Expect(listCollectionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListEnvironmentsOptions successfully`, func() {
				// Construct an instance of the ListEnvironmentsOptions model
				listEnvironmentsOptionsModel := appConfigurationService.NewListEnvironmentsOptions()
				listEnvironmentsOptionsModel.SetExpand(true)
				listEnvironmentsOptionsModel.SetSort("created_time")
				listEnvironmentsOptionsModel.SetTags("version 1.1,pre-release")
				listEnvironmentsOptionsModel.SetInclude([]string{"features", "properties", "snapshots"})
				listEnvironmentsOptionsModel.SetLimit(int64(10))
				listEnvironmentsOptionsModel.SetOffset(int64(0))
				listEnvironmentsOptionsModel.SetSearch("test tag")
				listEnvironmentsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listEnvironmentsOptionsModel).ToNot(BeNil())
				Expect(listEnvironmentsOptionsModel.Expand).To(Equal(core.BoolPtr(true)))
				Expect(listEnvironmentsOptionsModel.Sort).To(Equal(core.StringPtr("created_time")))
				Expect(listEnvironmentsOptionsModel.Tags).To(Equal(core.StringPtr("version 1.1,pre-release")))
				Expect(listEnvironmentsOptionsModel.Include).To(Equal([]string{"features", "properties", "snapshots"}))
				Expect(listEnvironmentsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listEnvironmentsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listEnvironmentsOptionsModel.Search).To(Equal(core.StringPtr("test tag")))
				Expect(listEnvironmentsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListFeaturesOptions successfully`, func() {
				// Construct an instance of the ListFeaturesOptions model
				environmentID := "environment_id"
				listFeaturesOptionsModel := appConfigurationService.NewListFeaturesOptions(environmentID)
				listFeaturesOptionsModel.SetEnvironmentID("environment_id")
				listFeaturesOptionsModel.SetExpand(true)
				listFeaturesOptionsModel.SetSort("created_time")
				listFeaturesOptionsModel.SetTags("version 1.1,pre-release")
				listFeaturesOptionsModel.SetCollections([]string{"my-collection-id", "ghzindiapvtltd"})
				listFeaturesOptionsModel.SetSegments([]string{"my-segment-id", "beta-users"})
				listFeaturesOptionsModel.SetInclude([]string{"collections", "rules", "change_request"})
				listFeaturesOptionsModel.SetLimit(int64(10))
				listFeaturesOptionsModel.SetOffset(int64(0))
				listFeaturesOptionsModel.SetSearch("test tag")
				listFeaturesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listFeaturesOptionsModel).ToNot(BeNil())
				Expect(listFeaturesOptionsModel.EnvironmentID).To(Equal(core.StringPtr("environment_id")))
				Expect(listFeaturesOptionsModel.Expand).To(Equal(core.BoolPtr(true)))
				Expect(listFeaturesOptionsModel.Sort).To(Equal(core.StringPtr("created_time")))
				Expect(listFeaturesOptionsModel.Tags).To(Equal(core.StringPtr("version 1.1,pre-release")))
				Expect(listFeaturesOptionsModel.Collections).To(Equal([]string{"my-collection-id", "ghzindiapvtltd"}))
				Expect(listFeaturesOptionsModel.Segments).To(Equal([]string{"my-segment-id", "beta-users"}))
				Expect(listFeaturesOptionsModel.Include).To(Equal([]string{"collections", "rules", "change_request"}))
				Expect(listFeaturesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listFeaturesOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listFeaturesOptionsModel.Search).To(Equal(core.StringPtr("test tag")))
				Expect(listFeaturesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListInstanceConfigOptions successfully`, func() {
				// Construct an instance of the ListInstanceConfigOptions model
				listInstanceConfigOptionsModel := appConfigurationService.NewListInstanceConfigOptions()
				listInstanceConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listInstanceConfigOptionsModel).ToNot(BeNil())
				Expect(listInstanceConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListOriginconfigsOptions successfully`, func() {
				// Construct an instance of the ListOriginconfigsOptions model
				listOriginconfigsOptionsModel := appConfigurationService.NewListOriginconfigsOptions()
				listOriginconfigsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listOriginconfigsOptionsModel).ToNot(BeNil())
				Expect(listOriginconfigsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListPropertiesOptions successfully`, func() {
				// Construct an instance of the ListPropertiesOptions model
				environmentID := "environment_id"
				listPropertiesOptionsModel := appConfigurationService.NewListPropertiesOptions(environmentID)
				listPropertiesOptionsModel.SetEnvironmentID("environment_id")
				listPropertiesOptionsModel.SetExpand(true)
				listPropertiesOptionsModel.SetSort("created_time")
				listPropertiesOptionsModel.SetTags("version 1.1,pre-release")
				listPropertiesOptionsModel.SetCollections([]string{"my-collection-id", "ghzindiapvtltd"})
				listPropertiesOptionsModel.SetSegments([]string{"my-segment-id", "beta-users"})
				listPropertiesOptionsModel.SetInclude([]string{"collections", "rules"})
				listPropertiesOptionsModel.SetLimit(int64(10))
				listPropertiesOptionsModel.SetOffset(int64(0))
				listPropertiesOptionsModel.SetSearch("test tag")
				listPropertiesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listPropertiesOptionsModel).ToNot(BeNil())
				Expect(listPropertiesOptionsModel.EnvironmentID).To(Equal(core.StringPtr("environment_id")))
				Expect(listPropertiesOptionsModel.Expand).To(Equal(core.BoolPtr(true)))
				Expect(listPropertiesOptionsModel.Sort).To(Equal(core.StringPtr("created_time")))
				Expect(listPropertiesOptionsModel.Tags).To(Equal(core.StringPtr("version 1.1,pre-release")))
				Expect(listPropertiesOptionsModel.Collections).To(Equal([]string{"my-collection-id", "ghzindiapvtltd"}))
				Expect(listPropertiesOptionsModel.Segments).To(Equal([]string{"my-segment-id", "beta-users"}))
				Expect(listPropertiesOptionsModel.Include).To(Equal([]string{"collections", "rules"}))
				Expect(listPropertiesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listPropertiesOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listPropertiesOptionsModel.Search).To(Equal(core.StringPtr("test tag")))
				Expect(listPropertiesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListSegmentsOptions successfully`, func() {
				// Construct an instance of the ListSegmentsOptions model
				listSegmentsOptionsModel := appConfigurationService.NewListSegmentsOptions()
				listSegmentsOptionsModel.SetExpand(true)
				listSegmentsOptionsModel.SetSort("created_time")
				listSegmentsOptionsModel.SetTags("version 1.1,pre-release")
				listSegmentsOptionsModel.SetInclude("rules")
				listSegmentsOptionsModel.SetLimit(int64(10))
				listSegmentsOptionsModel.SetOffset(int64(0))
				listSegmentsOptionsModel.SetSearch("test tag")
				listSegmentsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listSegmentsOptionsModel).ToNot(BeNil())
				Expect(listSegmentsOptionsModel.Expand).To(Equal(core.BoolPtr(true)))
				Expect(listSegmentsOptionsModel.Sort).To(Equal(core.StringPtr("created_time")))
				Expect(listSegmentsOptionsModel.Tags).To(Equal(core.StringPtr("version 1.1,pre-release")))
				Expect(listSegmentsOptionsModel.Include).To(Equal(core.StringPtr("rules")))
				Expect(listSegmentsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listSegmentsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listSegmentsOptionsModel.Search).To(Equal(core.StringPtr("test tag")))
				Expect(listSegmentsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListSnapshotsOptions successfully`, func() {
				// Construct an instance of the ListSnapshotsOptions model
				listSnapshotsOptionsModel := appConfigurationService.NewListSnapshotsOptions()
				listSnapshotsOptionsModel.SetSort("created_time")
				listSnapshotsOptionsModel.SetCollectionID("collection_id")
				listSnapshotsOptionsModel.SetEnvironmentID("environment_id")
				listSnapshotsOptionsModel.SetLimit(int64(10))
				listSnapshotsOptionsModel.SetOffset(int64(0))
				listSnapshotsOptionsModel.SetSearch("search_string")
				listSnapshotsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listSnapshotsOptionsModel).ToNot(BeNil())
				Expect(listSnapshotsOptionsModel.Sort).To(Equal(core.StringPtr("created_time")))
				Expect(listSnapshotsOptionsModel.CollectionID).To(Equal(core.StringPtr("collection_id")))
				Expect(listSnapshotsOptionsModel.EnvironmentID).To(Equal(core.StringPtr("environment_id")))
				Expect(listSnapshotsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listSnapshotsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listSnapshotsOptionsModel.Search).To(Equal(core.StringPtr("search_string")))
				Expect(listSnapshotsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListWorkflowconfigOptions successfully`, func() {
				// Construct an instance of the ListWorkflowconfigOptions model
				environmentID := "environment_id"
				listWorkflowconfigOptionsModel := appConfigurationService.NewListWorkflowconfigOptions(environmentID)
				listWorkflowconfigOptionsModel.SetEnvironmentID("environment_id")
				listWorkflowconfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listWorkflowconfigOptionsModel).ToNot(BeNil())
				Expect(listWorkflowconfigOptionsModel.EnvironmentID).To(Equal(core.StringPtr("environment_id")))
				Expect(listWorkflowconfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPromoteGitconfigOptions successfully`, func() {
				// Construct an instance of the PromoteGitconfigOptions model
				gitConfigID := "git_config_id"
				promoteGitconfigOptionsModel := appConfigurationService.NewPromoteGitconfigOptions(gitConfigID)
				promoteGitconfigOptionsModel.SetGitConfigID("git_config_id")
				promoteGitconfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(promoteGitconfigOptionsModel).ToNot(BeNil())
				Expect(promoteGitconfigOptionsModel.GitConfigID).To(Equal(core.StringPtr("git_config_id")))
				Expect(promoteGitconfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPromoteRestoreConfigOptions successfully`, func() {
				// Construct an instance of the PromoteRestoreConfigOptions model
				gitConfigID := "git_config_id"
				action := "promote"
				promoteRestoreConfigOptionsModel := appConfigurationService.NewPromoteRestoreConfigOptions(gitConfigID, action)
				promoteRestoreConfigOptionsModel.SetGitConfigID("git_config_id")
				promoteRestoreConfigOptionsModel.SetAction("promote")
				promoteRestoreConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(promoteRestoreConfigOptionsModel).ToNot(BeNil())
				Expect(promoteRestoreConfigOptionsModel.GitConfigID).To(Equal(core.StringPtr("git_config_id")))
				Expect(promoteRestoreConfigOptionsModel.Action).To(Equal(core.StringPtr("promote")))
				Expect(promoteRestoreConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewProperty successfully`, func() {
				name := "testString"
				propertyID := "testString"
				typeVar := "BOOLEAN"
				value := "testString"
				_model, err := appConfigurationService.NewProperty(name, propertyID, typeVar, value)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPropertyOutput successfully`, func() {
				propertyID := "testString"
				name := "testString"
				_model, err := appConfigurationService.NewPropertyOutput(propertyID, name)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRestoreGitconfigOptions successfully`, func() {
				// Construct an instance of the RestoreGitconfigOptions model
				gitConfigID := "git_config_id"
				restoreGitconfigOptionsModel := appConfigurationService.NewRestoreGitconfigOptions(gitConfigID)
				restoreGitconfigOptionsModel.SetGitConfigID("git_config_id")
				restoreGitconfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(restoreGitconfigOptionsModel).ToNot(BeNil())
				Expect(restoreGitconfigOptionsModel.GitConfigID).To(Equal(core.StringPtr("git_config_id")))
				Expect(restoreGitconfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRule successfully`, func() {
				attributeName := "testString"
				operator := "is"
				values := []string{"testString"}
				_model, err := appConfigurationService.NewRule(attributeName, operator, values)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSegment successfully`, func() {
				name := "testString"
				segmentID := "testString"
				rules := []appconfigurationv1.Rule{}
				_model, err := appConfigurationService.NewSegment(name, segmentID, rules)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSegmentRule successfully`, func() {
				rules := []appconfigurationv1.TargetSegments{}
				value := "testString"
				order := int64(38)
				_model, err := appConfigurationService.NewSegmentRule(rules, value, order)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSnapshotOutput successfully`, func() {
				gitConfigID := "testString"
				name := "testString"
				_model, err := appConfigurationService.NewSnapshotOutput(gitConfigID, name)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewTargetSegments successfully`, func() {
				segments := []string{"testString"}
				_model, err := appConfigurationService.NewTargetSegments(segments)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewToggleFeatureOptions successfully`, func() {
				// Construct an instance of the ToggleFeatureOptions model
				environmentID := "environment_id"
				featureID := "feature_id"
				toggleFeatureOptionsEnabled := true
				toggleFeatureOptionsModel := appConfigurationService.NewToggleFeatureOptions(environmentID, featureID, toggleFeatureOptionsEnabled)
				toggleFeatureOptionsModel.SetEnvironmentID("environment_id")
				toggleFeatureOptionsModel.SetFeatureID("feature_id")
				toggleFeatureOptionsModel.SetEnabled(true)
				toggleFeatureOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(toggleFeatureOptionsModel).ToNot(BeNil())
				Expect(toggleFeatureOptionsModel.EnvironmentID).To(Equal(core.StringPtr("environment_id")))
				Expect(toggleFeatureOptionsModel.FeatureID).To(Equal(core.StringPtr("feature_id")))
				Expect(toggleFeatureOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(toggleFeatureOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateCollectionOptions successfully`, func() {
				// Construct an instance of the UpdateCollectionOptions model
				collectionID := "collection_id"
				updateCollectionOptionsModel := appConfigurationService.NewUpdateCollectionOptions(collectionID)
				updateCollectionOptionsModel.SetCollectionID("collection_id")
				updateCollectionOptionsModel.SetName("testString")
				updateCollectionOptionsModel.SetDescription("testString")
				updateCollectionOptionsModel.SetTags("testString")
				updateCollectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateCollectionOptionsModel).ToNot(BeNil())
				Expect(updateCollectionOptionsModel.CollectionID).To(Equal(core.StringPtr("collection_id")))
				Expect(updateCollectionOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateCollectionOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateCollectionOptionsModel.Tags).To(Equal(core.StringPtr("testString")))
				Expect(updateCollectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateEnvironmentOptions successfully`, func() {
				// Construct an instance of the UpdateEnvironmentOptions model
				environmentID := "environment_id"
				updateEnvironmentOptionsModel := appConfigurationService.NewUpdateEnvironmentOptions(environmentID)
				updateEnvironmentOptionsModel.SetEnvironmentID("environment_id")
				updateEnvironmentOptionsModel.SetName("testString")
				updateEnvironmentOptionsModel.SetDescription("testString")
				updateEnvironmentOptionsModel.SetTags("testString")
				updateEnvironmentOptionsModel.SetColorCode("#FDD13A")
				updateEnvironmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateEnvironmentOptionsModel).ToNot(BeNil())
				Expect(updateEnvironmentOptionsModel.EnvironmentID).To(Equal(core.StringPtr("environment_id")))
				Expect(updateEnvironmentOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateEnvironmentOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateEnvironmentOptionsModel.Tags).To(Equal(core.StringPtr("testString")))
				Expect(updateEnvironmentOptionsModel.ColorCode).To(Equal(core.StringPtr("#FDD13A")))
				Expect(updateEnvironmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateFeatureOptions successfully`, func() {
				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				Expect(targetSegmentsModel).ToNot(BeNil())
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}
				Expect(targetSegmentsModel.Segments).To(Equal([]string{"betausers", "premiumusers"}))

				// Construct an instance of the FeatureSegmentRule model
				featureSegmentRuleModel := new(appconfigurationv1.FeatureSegmentRule)
				Expect(featureSegmentRuleModel).ToNot(BeNil())
				featureSegmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				featureSegmentRuleModel.Value = "true"
				featureSegmentRuleModel.Order = core.Int64Ptr(int64(1))
				featureSegmentRuleModel.RolloutPercentage = core.Int64Ptr(int64(90))
				Expect(featureSegmentRuleModel.Rules).To(Equal([]appconfigurationv1.TargetSegments{*targetSegmentsModel}))
				Expect(featureSegmentRuleModel.Value).To(Equal("true"))
				Expect(featureSegmentRuleModel.Order).To(Equal(core.Int64Ptr(int64(1))))
				Expect(featureSegmentRuleModel.RolloutPercentage).To(Equal(core.Int64Ptr(int64(90))))

				// Construct an instance of the CollectionUpdateRef model
				collectionUpdateRefModel := new(appconfigurationv1.CollectionUpdateRef)
				Expect(collectionUpdateRefModel).ToNot(BeNil())
				collectionUpdateRefModel.CollectionID = core.StringPtr("ghzinc")
				collectionUpdateRefModel.Deleted = core.BoolPtr(true)
				Expect(collectionUpdateRefModel.CollectionID).To(Equal(core.StringPtr("ghzinc")))
				Expect(collectionUpdateRefModel.Deleted).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the UpdateFeatureOptions model
				environmentID := "environment_id"
				featureID := "feature_id"
				updateFeatureOptionsModel := appConfigurationService.NewUpdateFeatureOptions(environmentID, featureID)
				updateFeatureOptionsModel.SetEnvironmentID("environment_id")
				updateFeatureOptionsModel.SetFeatureID("feature_id")
				updateFeatureOptionsModel.SetName("Cycle Rentals")
				updateFeatureOptionsModel.SetDescription("Feature flags to enable Cycle Rentals")
				updateFeatureOptionsModel.SetEnabledValue("true")
				updateFeatureOptionsModel.SetDisabledValue("false")
				updateFeatureOptionsModel.SetEnabled(true)
				updateFeatureOptionsModel.SetRolloutPercentage(int64(100))
				updateFeatureOptionsModel.SetTags("version: 1.1, yet-to-release")
				updateFeatureOptionsModel.SetSegmentRules([]appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel})
				updateFeatureOptionsModel.SetCollections([]appconfigurationv1.CollectionUpdateRef{*collectionUpdateRefModel})
				updateFeatureOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateFeatureOptionsModel).ToNot(BeNil())
				Expect(updateFeatureOptionsModel.EnvironmentID).To(Equal(core.StringPtr("environment_id")))
				Expect(updateFeatureOptionsModel.FeatureID).To(Equal(core.StringPtr("feature_id")))
				Expect(updateFeatureOptionsModel.Name).To(Equal(core.StringPtr("Cycle Rentals")))
				Expect(updateFeatureOptionsModel.Description).To(Equal(core.StringPtr("Feature flags to enable Cycle Rentals")))
				Expect(updateFeatureOptionsModel.EnabledValue).To(Equal("true"))
				Expect(updateFeatureOptionsModel.DisabledValue).To(Equal("false"))
				Expect(updateFeatureOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(updateFeatureOptionsModel.RolloutPercentage).To(Equal(core.Int64Ptr(int64(100))))
				Expect(updateFeatureOptionsModel.Tags).To(Equal(core.StringPtr("version: 1.1, yet-to-release")))
				Expect(updateFeatureOptionsModel.SegmentRules).To(Equal([]appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel}))
				Expect(updateFeatureOptionsModel.Collections).To(Equal([]appconfigurationv1.CollectionUpdateRef{*collectionUpdateRefModel}))
				Expect(updateFeatureOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateFeatureValuesOptions successfully`, func() {
				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				Expect(targetSegmentsModel).ToNot(BeNil())
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}
				Expect(targetSegmentsModel.Segments).To(Equal([]string{"betausers", "premiumusers"}))

				// Construct an instance of the FeatureSegmentRule model
				featureSegmentRuleModel := new(appconfigurationv1.FeatureSegmentRule)
				Expect(featureSegmentRuleModel).ToNot(BeNil())
				featureSegmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				featureSegmentRuleModel.Value = "true"
				featureSegmentRuleModel.Order = core.Int64Ptr(int64(1))
				featureSegmentRuleModel.RolloutPercentage = core.Int64Ptr(int64(100))
				Expect(featureSegmentRuleModel.Rules).To(Equal([]appconfigurationv1.TargetSegments{*targetSegmentsModel}))
				Expect(featureSegmentRuleModel.Value).To(Equal("true"))
				Expect(featureSegmentRuleModel.Order).To(Equal(core.Int64Ptr(int64(1))))
				Expect(featureSegmentRuleModel.RolloutPercentage).To(Equal(core.Int64Ptr(int64(100))))

				// Construct an instance of the UpdateFeatureValuesOptions model
				environmentID := "environment_id"
				featureID := "feature_id"
				updateFeatureValuesOptionsModel := appConfigurationService.NewUpdateFeatureValuesOptions(environmentID, featureID)
				updateFeatureValuesOptionsModel.SetEnvironmentID("environment_id")
				updateFeatureValuesOptionsModel.SetFeatureID("feature_id")
				updateFeatureValuesOptionsModel.SetName("Cycle Rentals")
				updateFeatureValuesOptionsModel.SetDescription("Feature flags to enable Cycle Rentals")
				updateFeatureValuesOptionsModel.SetTags("version: 1.1, yet-to-release")
				updateFeatureValuesOptionsModel.SetEnabledValue("true")
				updateFeatureValuesOptionsModel.SetDisabledValue("false")
				updateFeatureValuesOptionsModel.SetRolloutPercentage(int64(100))
				updateFeatureValuesOptionsModel.SetSegmentRules([]appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel})
				updateFeatureValuesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateFeatureValuesOptionsModel).ToNot(BeNil())
				Expect(updateFeatureValuesOptionsModel.EnvironmentID).To(Equal(core.StringPtr("environment_id")))
				Expect(updateFeatureValuesOptionsModel.FeatureID).To(Equal(core.StringPtr("feature_id")))
				Expect(updateFeatureValuesOptionsModel.Name).To(Equal(core.StringPtr("Cycle Rentals")))
				Expect(updateFeatureValuesOptionsModel.Description).To(Equal(core.StringPtr("Feature flags to enable Cycle Rentals")))
				Expect(updateFeatureValuesOptionsModel.Tags).To(Equal(core.StringPtr("version: 1.1, yet-to-release")))
				Expect(updateFeatureValuesOptionsModel.EnabledValue).To(Equal("true"))
				Expect(updateFeatureValuesOptionsModel.DisabledValue).To(Equal("false"))
				Expect(updateFeatureValuesOptionsModel.RolloutPercentage).To(Equal(core.Int64Ptr(int64(100))))
				Expect(updateFeatureValuesOptionsModel.SegmentRules).To(Equal([]appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel}))
				Expect(updateFeatureValuesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateGitconfigOptions successfully`, func() {
				// Construct an instance of the UpdateGitconfigOptions model
				gitConfigID := "git_config_id"
				updateGitconfigOptionsModel := appConfigurationService.NewUpdateGitconfigOptions(gitConfigID)
				updateGitconfigOptionsModel.SetGitConfigID("git_config_id")
				updateGitconfigOptionsModel.SetGitConfigName("testString")
				updateGitconfigOptionsModel.SetCollectionID("testString")
				updateGitconfigOptionsModel.SetEnvironmentID("testString")
				updateGitconfigOptionsModel.SetGitURL("testString")
				updateGitconfigOptionsModel.SetGitBranch("testString")
				updateGitconfigOptionsModel.SetGitFilePath("testString")
				updateGitconfigOptionsModel.SetGitToken("testString")
				updateGitconfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateGitconfigOptionsModel).ToNot(BeNil())
				Expect(updateGitconfigOptionsModel.GitConfigID).To(Equal(core.StringPtr("git_config_id")))
				Expect(updateGitconfigOptionsModel.GitConfigName).To(Equal(core.StringPtr("testString")))
				Expect(updateGitconfigOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(updateGitconfigOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(updateGitconfigOptionsModel.GitURL).To(Equal(core.StringPtr("testString")))
				Expect(updateGitconfigOptionsModel.GitBranch).To(Equal(core.StringPtr("testString")))
				Expect(updateGitconfigOptionsModel.GitFilePath).To(Equal(core.StringPtr("testString")))
				Expect(updateGitconfigOptionsModel.GitToken).To(Equal(core.StringPtr("testString")))
				Expect(updateGitconfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateOriginconfigsOptions successfully`, func() {
				// Construct an instance of the UpdateOriginconfigsOptions model
				updateOriginconfigsOptionsAllowedOrigins := []string{"testString"}
				updateOriginconfigsOptionsModel := appConfigurationService.NewUpdateOriginconfigsOptions(updateOriginconfigsOptionsAllowedOrigins)
				updateOriginconfigsOptionsModel.SetAllowedOrigins([]string{"testString"})
				updateOriginconfigsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateOriginconfigsOptionsModel).ToNot(BeNil())
				Expect(updateOriginconfigsOptionsModel.AllowedOrigins).To(Equal([]string{"testString"}))
				Expect(updateOriginconfigsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdatePropertyOptions successfully`, func() {
				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				Expect(targetSegmentsModel).ToNot(BeNil())
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}
				Expect(targetSegmentsModel.Segments).To(Equal([]string{"betausers", "premiumusers"}))

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				Expect(segmentRuleModel).ToNot(BeNil())
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = "true"
				segmentRuleModel.Order = core.Int64Ptr(int64(1))
				Expect(segmentRuleModel.Rules).To(Equal([]appconfigurationv1.TargetSegments{*targetSegmentsModel}))
				Expect(segmentRuleModel.Value).To(Equal("true"))
				Expect(segmentRuleModel.Order).To(Equal(core.Int64Ptr(int64(1))))

				// Construct an instance of the CollectionUpdateRef model
				collectionUpdateRefModel := new(appconfigurationv1.CollectionUpdateRef)
				Expect(collectionUpdateRefModel).ToNot(BeNil())
				collectionUpdateRefModel.CollectionID = core.StringPtr("ghzinc")
				collectionUpdateRefModel.Deleted = core.BoolPtr(true)
				Expect(collectionUpdateRefModel.CollectionID).To(Equal(core.StringPtr("ghzinc")))
				Expect(collectionUpdateRefModel.Deleted).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the UpdatePropertyOptions model
				environmentID := "environment_id"
				propertyID := "property_id"
				updatePropertyOptionsModel := appConfigurationService.NewUpdatePropertyOptions(environmentID, propertyID)
				updatePropertyOptionsModel.SetEnvironmentID("environment_id")
				updatePropertyOptionsModel.SetPropertyID("property_id")
				updatePropertyOptionsModel.SetName("Email property")
				updatePropertyOptionsModel.SetDescription("Property for email")
				updatePropertyOptionsModel.SetValue("true")
				updatePropertyOptionsModel.SetTags("version: 1.1, pre-release")
				updatePropertyOptionsModel.SetSegmentRules([]appconfigurationv1.SegmentRule{*segmentRuleModel})
				updatePropertyOptionsModel.SetCollections([]appconfigurationv1.CollectionUpdateRef{*collectionUpdateRefModel})
				updatePropertyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updatePropertyOptionsModel).ToNot(BeNil())
				Expect(updatePropertyOptionsModel.EnvironmentID).To(Equal(core.StringPtr("environment_id")))
				Expect(updatePropertyOptionsModel.PropertyID).To(Equal(core.StringPtr("property_id")))
				Expect(updatePropertyOptionsModel.Name).To(Equal(core.StringPtr("Email property")))
				Expect(updatePropertyOptionsModel.Description).To(Equal(core.StringPtr("Property for email")))
				Expect(updatePropertyOptionsModel.Value).To(Equal("true"))
				Expect(updatePropertyOptionsModel.Tags).To(Equal(core.StringPtr("version: 1.1, pre-release")))
				Expect(updatePropertyOptionsModel.SegmentRules).To(Equal([]appconfigurationv1.SegmentRule{*segmentRuleModel}))
				Expect(updatePropertyOptionsModel.Collections).To(Equal([]appconfigurationv1.CollectionUpdateRef{*collectionUpdateRefModel}))
				Expect(updatePropertyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdatePropertyValuesOptions successfully`, func() {
				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				Expect(targetSegmentsModel).ToNot(BeNil())
				targetSegmentsModel.Segments = []string{"betausers", "premiumusers"}
				Expect(targetSegmentsModel.Segments).To(Equal([]string{"betausers", "premiumusers"}))

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				Expect(segmentRuleModel).ToNot(BeNil())
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = "true"
				segmentRuleModel.Order = core.Int64Ptr(int64(1))
				Expect(segmentRuleModel.Rules).To(Equal([]appconfigurationv1.TargetSegments{*targetSegmentsModel}))
				Expect(segmentRuleModel.Value).To(Equal("true"))
				Expect(segmentRuleModel.Order).To(Equal(core.Int64Ptr(int64(1))))

				// Construct an instance of the UpdatePropertyValuesOptions model
				environmentID := "environment_id"
				propertyID := "property_id"
				updatePropertyValuesOptionsModel := appConfigurationService.NewUpdatePropertyValuesOptions(environmentID, propertyID)
				updatePropertyValuesOptionsModel.SetEnvironmentID("environment_id")
				updatePropertyValuesOptionsModel.SetPropertyID("property_id")
				updatePropertyValuesOptionsModel.SetName("Email property")
				updatePropertyValuesOptionsModel.SetDescription("Property for email")
				updatePropertyValuesOptionsModel.SetTags("version: 1.1, pre-release")
				updatePropertyValuesOptionsModel.SetValue("true")
				updatePropertyValuesOptionsModel.SetSegmentRules([]appconfigurationv1.SegmentRule{*segmentRuleModel})
				updatePropertyValuesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updatePropertyValuesOptionsModel).ToNot(BeNil())
				Expect(updatePropertyValuesOptionsModel.EnvironmentID).To(Equal(core.StringPtr("environment_id")))
				Expect(updatePropertyValuesOptionsModel.PropertyID).To(Equal(core.StringPtr("property_id")))
				Expect(updatePropertyValuesOptionsModel.Name).To(Equal(core.StringPtr("Email property")))
				Expect(updatePropertyValuesOptionsModel.Description).To(Equal(core.StringPtr("Property for email")))
				Expect(updatePropertyValuesOptionsModel.Tags).To(Equal(core.StringPtr("version: 1.1, pre-release")))
				Expect(updatePropertyValuesOptionsModel.Value).To(Equal("true"))
				Expect(updatePropertyValuesOptionsModel.SegmentRules).To(Equal([]appconfigurationv1.SegmentRule{*segmentRuleModel}))
				Expect(updatePropertyValuesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateSegmentOptions successfully`, func() {
				// Construct an instance of the Rule model
				ruleModel := new(appconfigurationv1.Rule)
				Expect(ruleModel).ToNot(BeNil())
				ruleModel.AttributeName = core.StringPtr("testString")
				ruleModel.Operator = core.StringPtr("is")
				ruleModel.Values = []string{"testString"}
				Expect(ruleModel.AttributeName).To(Equal(core.StringPtr("testString")))
				Expect(ruleModel.Operator).To(Equal(core.StringPtr("is")))
				Expect(ruleModel.Values).To(Equal([]string{"testString"}))

				// Construct an instance of the UpdateSegmentOptions model
				segmentID := "segment_id"
				updateSegmentOptionsModel := appConfigurationService.NewUpdateSegmentOptions(segmentID)
				updateSegmentOptionsModel.SetSegmentID("segment_id")
				updateSegmentOptionsModel.SetName("testString")
				updateSegmentOptionsModel.SetDescription("testString")
				updateSegmentOptionsModel.SetTags("testString")
				updateSegmentOptionsModel.SetRules([]appconfigurationv1.Rule{*ruleModel})
				updateSegmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateSegmentOptionsModel).ToNot(BeNil())
				Expect(updateSegmentOptionsModel.SegmentID).To(Equal(core.StringPtr("segment_id")))
				Expect(updateSegmentOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateSegmentOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateSegmentOptionsModel.Tags).To(Equal(core.StringPtr("testString")))
				Expect(updateSegmentOptionsModel.Rules).To(Equal([]appconfigurationv1.Rule{*ruleModel}))
				Expect(updateSegmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateWorkflowconfigOptions successfully`, func() {
				// Construct an instance of the ExternalServiceNowCredentials model
				externalServiceNowCredentialsModel := new(appconfigurationv1.ExternalServiceNowCredentials)
				Expect(externalServiceNowCredentialsModel).ToNot(BeNil())
				externalServiceNowCredentialsModel.Username = core.StringPtr("admin")
				externalServiceNowCredentialsModel.Password = core.StringPtr("testString")
				externalServiceNowCredentialsModel.ClientID = core.StringPtr("f7b6379b55d08210f8ree233afc7256d")
				externalServiceNowCredentialsModel.ClientSecret = core.StringPtr("testString")
				Expect(externalServiceNowCredentialsModel.Username).To(Equal(core.StringPtr("admin")))
				Expect(externalServiceNowCredentialsModel.Password).To(Equal(core.StringPtr("testString")))
				Expect(externalServiceNowCredentialsModel.ClientID).To(Equal(core.StringPtr("f7b6379b55d08210f8ree233afc7256d")))
				Expect(externalServiceNowCredentialsModel.ClientSecret).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateWorkflowConfigUpdateExternalServiceNow model
				updateWorkflowConfigModel := new(appconfigurationv1.UpdateWorkflowConfigUpdateExternalServiceNow)
				Expect(updateWorkflowConfigModel).ToNot(BeNil())
				updateWorkflowConfigModel.WorkflowURL = core.StringPtr("testString")
				updateWorkflowConfigModel.ApprovalGroupName = core.StringPtr("testString")
				updateWorkflowConfigModel.ApprovalExpiration = core.Int64Ptr(int64(1))
				updateWorkflowConfigModel.WorkflowCredentials = externalServiceNowCredentialsModel
				updateWorkflowConfigModel.Enabled = core.BoolPtr(false)
				Expect(updateWorkflowConfigModel.WorkflowURL).To(Equal(core.StringPtr("testString")))
				Expect(updateWorkflowConfigModel.ApprovalGroupName).To(Equal(core.StringPtr("testString")))
				Expect(updateWorkflowConfigModel.ApprovalExpiration).To(Equal(core.Int64Ptr(int64(1))))
				Expect(updateWorkflowConfigModel.WorkflowCredentials).To(Equal(externalServiceNowCredentialsModel))
				Expect(updateWorkflowConfigModel.Enabled).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the UpdateWorkflowconfigOptions model
				environmentID := "environment_id"
				var updateWorkflowConfig appconfigurationv1.UpdateWorkflowConfigIntf = nil
				updateWorkflowconfigOptionsModel := appConfigurationService.NewUpdateWorkflowconfigOptions(environmentID, updateWorkflowConfig)
				updateWorkflowconfigOptionsModel.SetEnvironmentID("environment_id")
				updateWorkflowconfigOptionsModel.SetUpdateWorkflowConfig(updateWorkflowConfigModel)
				updateWorkflowconfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateWorkflowconfigOptionsModel).ToNot(BeNil())
				Expect(updateWorkflowconfigOptionsModel.EnvironmentID).To(Equal(core.StringPtr("environment_id")))
				Expect(updateWorkflowconfigOptionsModel.UpdateWorkflowConfig).To(Equal(updateWorkflowConfigModel))
				Expect(updateWorkflowconfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateWorkflowConfigExternalServiceNow successfully`, func() {
				workflowURL := "testString"
				approvalGroupName := "testString"
				approvalExpiration := int64(1)
				var workflowCredentials *appconfigurationv1.ExternalServiceNowCredentials = nil
				enabled := false
				_, err := appConfigurationService.NewCreateWorkflowConfigExternalServiceNow(workflowURL, approvalGroupName, approvalExpiration, workflowCredentials, enabled)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCreateWorkflowConfigIBMServiceNow successfully`, func() {
				serviceCrn := "testString"
				workflowType := "testString"
				approvalExpiration := int64(1)
				smInstanceCrn := "testString"
				secretID := "testString"
				enabled := false
				_model, err := appConfigurationService.NewCreateWorkflowConfigIBMServiceNow(serviceCrn, workflowType, approvalExpiration, smInstanceCrn, secretID, enabled)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalCollection successfully`, func() {
			// Construct an instance of the model.
			model := new(appconfigurationv1.Collection)
			model.Name = core.StringPtr("testString")
			model.CollectionID = core.StringPtr("testString")
			model.Description = core.StringPtr("testString")
			model.Tags = core.StringPtr("testString")
			model.CreatedTime = CreateMockDateTime("2021-05-12T23:20:50.520Z")
			model.UpdatedTime = CreateMockDateTime("2021-05-12T23:20:50.520Z")
			model.Href = core.StringPtr("testString")
			model.Features = nil
			model.Properties = nil
			model.Snapshots = nil
			model.FeaturesCount = core.Int64Ptr(int64(38))
			model.PropertiesCount = core.Int64Ptr(int64(38))
			model.SnapshotCount = core.Int64Ptr(int64(38))

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *appconfigurationv1.Collection
			err = appconfigurationv1.UnmarshalCollection(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalCollectionRef successfully`, func() {
			// Construct an instance of the model.
			model := new(appconfigurationv1.CollectionRef)
			model.CollectionID = core.StringPtr("testString")
			model.Name = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *appconfigurationv1.CollectionRef
			err = appconfigurationv1.UnmarshalCollectionRef(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalCollectionUpdateRef successfully`, func() {
			// Construct an instance of the model.
			model := new(appconfigurationv1.CollectionUpdateRef)
			model.CollectionID = core.StringPtr("testString")
			model.Deleted = core.BoolPtr(true)

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *appconfigurationv1.CollectionUpdateRef
			err = appconfigurationv1.UnmarshalCollectionUpdateRef(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalCreateWorkflowConfig successfully`, func() {
			// Construct an instance of the model.
			model := new(appconfigurationv1.CreateWorkflowConfig)
			model.EnvironmentName = core.StringPtr("testString")
			model.EnvironmentID = core.StringPtr("testString")
			model.WorkflowURL = core.StringPtr("testString")
			model.ApprovalGroupName = core.StringPtr("testString")
			model.ApprovalExpiration = core.Int64Ptr(int64(1))
			model.WorkflowCredentials = nil
			model.Enabled = core.BoolPtr(false)
			model.CreatedTime = CreateMockDateTime("2022-11-15T23:20:50.000Z")
			model.UpdatedTime = CreateMockDateTime("2022-11-16T21:20:50.000Z")
			model.Href = core.StringPtr("testString")
			model.ServiceCrn = core.StringPtr("testString")
			model.WorkflowType = core.StringPtr("testString")
			model.SmInstanceCrn = core.StringPtr("testString")
			model.SecretID = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *appconfigurationv1.CreateWorkflowConfig
			err = appconfigurationv1.UnmarshalCreateWorkflowConfig(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalEnvironment successfully`, func() {
			// Construct an instance of the model.
			model := new(appconfigurationv1.Environment)
			model.Name = core.StringPtr("testString")
			model.EnvironmentID = core.StringPtr("testString")
			model.Description = core.StringPtr("testString")
			model.Tags = core.StringPtr("testString")
			model.ColorCode = core.StringPtr("#FDD13A")
			model.CreatedTime = CreateMockDateTime("2021-05-12T23:20:50.520Z")
			model.UpdatedTime = CreateMockDateTime("2021-05-12T23:20:50.520Z")
			model.Href = core.StringPtr("testString")
			model.Features = nil
			model.Properties = nil
			model.Snapshots = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *appconfigurationv1.Environment
			err = appconfigurationv1.UnmarshalEnvironment(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalExternalServiceNowCredentials successfully`, func() {
			// Construct an instance of the model.
			model := new(appconfigurationv1.ExternalServiceNowCredentials)
			model.Username = core.StringPtr("admin")
			model.Password = core.StringPtr("testString")
			model.ClientID = core.StringPtr("f7b6379b55d08210f8ree233afc7256d")
			model.ClientSecret = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *appconfigurationv1.ExternalServiceNowCredentials
			err = appconfigurationv1.UnmarshalExternalServiceNowCredentials(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalFeatureOutput successfully`, func() {
			// Construct an instance of the model.
			model := new(appconfigurationv1.FeatureOutput)
			model.FeatureID = core.StringPtr("testString")
			model.Name = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *appconfigurationv1.FeatureOutput
			err = appconfigurationv1.UnmarshalFeatureOutput(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalFeatureSegmentRule successfully`, func() {
			// Construct an instance of the model.
			model := new(appconfigurationv1.FeatureSegmentRule)
			model.Rules = nil
			model.Value = "testString"
			model.Order = core.Int64Ptr(int64(38))
			model.RolloutPercentage = core.Int64Ptr(int64(100))

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *appconfigurationv1.FeatureSegmentRule
			err = appconfigurationv1.UnmarshalFeatureSegmentRule(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalImportCollectionSchema successfully`, func() {
			// Construct an instance of the model.
			model := new(appconfigurationv1.ImportCollectionSchema)
			model.CollectionID = core.StringPtr("testString")
			model.Name = core.StringPtr("testString")
			model.Description = core.StringPtr("testString")
			model.Tags = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *appconfigurationv1.ImportCollectionSchema
			err = appconfigurationv1.UnmarshalImportCollectionSchema(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalImportConfig successfully`, func() {
			// Construct an instance of the model.
			model := new(appconfigurationv1.ImportConfig)
			model.Environments = nil
			model.Collections = nil
			model.Segments = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *appconfigurationv1.ImportConfig
			err = appconfigurationv1.UnmarshalImportConfig(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalImportEnvironmentSchema successfully`, func() {
			// Construct an instance of the model.
			model := new(appconfigurationv1.ImportEnvironmentSchema)
			model.Name = core.StringPtr("testString")
			model.EnvironmentID = core.StringPtr("testString")
			model.Description = core.StringPtr("testString")
			model.Tags = core.StringPtr("testString")
			model.ColorCode = core.StringPtr("#FDD13A")
			model.Features = nil
			model.Properties = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *appconfigurationv1.ImportEnvironmentSchema
			err = appconfigurationv1.UnmarshalImportEnvironmentSchema(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalImportFeatureRequestBody successfully`, func() {
			// Construct an instance of the model.
			model := new(appconfigurationv1.ImportFeatureRequestBody)
			model.Name = core.StringPtr("testString")
			model.FeatureID = core.StringPtr("testString")
			model.Description = core.StringPtr("testString")
			model.Type = core.StringPtr("BOOLEAN")
			model.Format = core.StringPtr("TEXT")
			model.EnabledValue = "testString"
			model.DisabledValue = "testString"
			model.Enabled = core.BoolPtr(true)
			model.RolloutPercentage = core.Int64Ptr(int64(100))
			model.Tags = core.StringPtr("testString")
			model.SegmentRules = nil
			model.Collections = nil
			model.IsOverridden = core.BoolPtr(true)

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *appconfigurationv1.ImportFeatureRequestBody
			err = appconfigurationv1.UnmarshalImportFeatureRequestBody(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalImportPropertyRequestBody successfully`, func() {
			// Construct an instance of the model.
			model := new(appconfigurationv1.ImportPropertyRequestBody)
			model.Name = core.StringPtr("testString")
			model.PropertyID = core.StringPtr("testString")
			model.Description = core.StringPtr("testString")
			model.Type = core.StringPtr("BOOLEAN")
			model.Format = core.StringPtr("TEXT")
			model.Value = "testString"
			model.Tags = core.StringPtr("testString")
			model.SegmentRules = nil
			model.Collections = nil
			model.IsOverridden = core.BoolPtr(true)

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *appconfigurationv1.ImportPropertyRequestBody
			err = appconfigurationv1.UnmarshalImportPropertyRequestBody(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalImportSegmentSchema successfully`, func() {
			// Construct an instance of the model.
			model := new(appconfigurationv1.ImportSegmentSchema)
			model.Name = core.StringPtr("testString")
			model.SegmentID = core.StringPtr("testString")
			model.Description = core.StringPtr("testString")
			model.Tags = core.StringPtr("testString")
			model.Rules = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *appconfigurationv1.ImportSegmentSchema
			err = appconfigurationv1.UnmarshalImportSegmentSchema(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalProperty successfully`, func() {
			// Construct an instance of the model.
			model := new(appconfigurationv1.Property)
			model.Name = core.StringPtr("testString")
			model.PropertyID = core.StringPtr("testString")
			model.Description = core.StringPtr("testString")
			model.Type = core.StringPtr("BOOLEAN")
			model.Format = core.StringPtr("TEXT")
			model.Value = "testString"
			model.Tags = core.StringPtr("testString")
			model.SegmentRules = nil
			model.SegmentExists = core.BoolPtr(true)
			model.Collections = nil
			model.CreatedTime = CreateMockDateTime("2021-05-12T23:20:50.520Z")
			model.UpdatedTime = CreateMockDateTime("2021-05-12T23:20:50.520Z")
			model.EvaluationTime = CreateMockDateTime("2021-05-12T23:20:50.520Z")
			model.Href = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *appconfigurationv1.Property
			err = appconfigurationv1.UnmarshalProperty(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalPropertyOutput successfully`, func() {
			// Construct an instance of the model.
			model := new(appconfigurationv1.PropertyOutput)
			model.PropertyID = core.StringPtr("testString")
			model.Name = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *appconfigurationv1.PropertyOutput
			err = appconfigurationv1.UnmarshalPropertyOutput(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalRule successfully`, func() {
			// Construct an instance of the model.
			model := new(appconfigurationv1.Rule)
			model.AttributeName = core.StringPtr("testString")
			model.Operator = core.StringPtr("is")
			model.Values = []string{"testString"}

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *appconfigurationv1.Rule
			err = appconfigurationv1.UnmarshalRule(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSegment successfully`, func() {
			// Construct an instance of the model.
			model := new(appconfigurationv1.Segment)
			model.Name = core.StringPtr("testString")
			model.SegmentID = core.StringPtr("testString")
			model.Description = core.StringPtr("testString")
			model.Tags = core.StringPtr("testString")
			model.Rules = nil
			model.CreatedTime = CreateMockDateTime("2021-05-12T23:20:50.520Z")
			model.UpdatedTime = CreateMockDateTime("2021-05-12T23:20:50.520Z")
			model.Href = core.StringPtr("testString")
			model.Features = nil
			model.Properties = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *appconfigurationv1.Segment
			err = appconfigurationv1.UnmarshalSegment(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSegmentRule successfully`, func() {
			// Construct an instance of the model.
			model := new(appconfigurationv1.SegmentRule)
			model.Rules = nil
			model.Value = "testString"
			model.Order = core.Int64Ptr(int64(38))

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *appconfigurationv1.SegmentRule
			err = appconfigurationv1.UnmarshalSegmentRule(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalSnapshotOutput successfully`, func() {
			// Construct an instance of the model.
			model := new(appconfigurationv1.SnapshotOutput)
			model.GitConfigID = core.StringPtr("testString")
			model.Name = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *appconfigurationv1.SnapshotOutput
			err = appconfigurationv1.UnmarshalSnapshotOutput(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalTargetSegments successfully`, func() {
			// Construct an instance of the model.
			model := new(appconfigurationv1.TargetSegments)
			model.Segments = []string{"testString"}

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *appconfigurationv1.TargetSegments
			err = appconfigurationv1.UnmarshalTargetSegments(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalUpdateWorkflowConfig successfully`, func() {
			// Construct an instance of the model.
			model := new(appconfigurationv1.UpdateWorkflowConfig)
			model.WorkflowURL = core.StringPtr("testString")
			model.ApprovalGroupName = core.StringPtr("testString")
			model.ApprovalExpiration = core.Int64Ptr(int64(1))
			model.WorkflowCredentials = nil
			model.Enabled = core.BoolPtr(false)
			model.ServiceCrn = core.StringPtr("testString")
			model.SmInstanceCrn = core.StringPtr("testString")
			model.SecretID = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *appconfigurationv1.UpdateWorkflowConfig
			err = appconfigurationv1.UnmarshalUpdateWorkflowConfig(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalCreateWorkflowConfigExternalServiceNow successfully`, func() {
			// Construct an instance of the model.
			model := new(appconfigurationv1.CreateWorkflowConfigExternalServiceNow)
			model.EnvironmentName = core.StringPtr("testString")
			model.EnvironmentID = core.StringPtr("testString")
			model.WorkflowURL = core.StringPtr("testString")
			model.ApprovalGroupName = core.StringPtr("testString")
			model.ApprovalExpiration = core.Int64Ptr(int64(1))
			model.WorkflowCredentials = nil
			model.Enabled = core.BoolPtr(false)
			model.CreatedTime = CreateMockDateTime("2022-11-15T23:20:50.000Z")
			model.UpdatedTime = CreateMockDateTime("2022-11-16T21:20:50.000Z")
			model.Href = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *appconfigurationv1.CreateWorkflowConfigExternalServiceNow
			err = appconfigurationv1.UnmarshalCreateWorkflowConfigExternalServiceNow(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalCreateWorkflowConfigIBMServiceNow successfully`, func() {
			// Construct an instance of the model.
			model := new(appconfigurationv1.CreateWorkflowConfigIBMServiceNow)
			model.EnvironmentName = core.StringPtr("testString")
			model.EnvironmentID = core.StringPtr("testString")
			model.ServiceCrn = core.StringPtr("testString")
			model.WorkflowType = core.StringPtr("testString")
			model.ApprovalExpiration = core.Int64Ptr(int64(1))
			model.SmInstanceCrn = core.StringPtr("testString")
			model.SecretID = core.StringPtr("testString")
			model.Enabled = core.BoolPtr(false)
			model.CreatedTime = CreateMockDateTime("2022-11-15T23:20:50.000Z")
			model.UpdatedTime = CreateMockDateTime("2022-11-16T21:20:50.000Z")
			model.Href = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *appconfigurationv1.CreateWorkflowConfigIBMServiceNow
			err = appconfigurationv1.UnmarshalCreateWorkflowConfigIBMServiceNow(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalUpdateWorkflowConfigUpdateExternalServiceNow successfully`, func() {
			// Construct an instance of the model.
			model := new(appconfigurationv1.UpdateWorkflowConfigUpdateExternalServiceNow)
			model.WorkflowURL = core.StringPtr("testString")
			model.ApprovalGroupName = core.StringPtr("testString")
			model.ApprovalExpiration = core.Int64Ptr(int64(1))
			model.WorkflowCredentials = nil
			model.Enabled = core.BoolPtr(false)

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *appconfigurationv1.UpdateWorkflowConfigUpdateExternalServiceNow
			err = appconfigurationv1.UnmarshalUpdateWorkflowConfigUpdateExternalServiceNow(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalUpdateWorkflowConfigUpdateIBMServiceNow successfully`, func() {
			// Construct an instance of the model.
			model := new(appconfigurationv1.UpdateWorkflowConfigUpdateIBMServiceNow)
			model.ServiceCrn = core.StringPtr("testString")
			model.ApprovalExpiration = core.Int64Ptr(int64(1))
			model.SmInstanceCrn = core.StringPtr("testString")
			model.SecretID = core.StringPtr("testString")
			model.Enabled = core.BoolPtr(false)

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *appconfigurationv1.UpdateWorkflowConfigUpdateIBMServiceNow
			err = appconfigurationv1.UnmarshalUpdateWorkflowConfigUpdateIBMServiceNow(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("VGhpcyBpcyBhIHRlc3Qgb2YgdGhlIGVtZXJnZW5jeSBicm9hZGNhc3Qgc3lzdGVt")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(encodedString string) *[]byte {
	ba, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		panic(err)
	}
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
