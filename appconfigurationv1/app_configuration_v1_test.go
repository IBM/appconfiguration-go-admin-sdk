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

package appconfigurationv1_test

import (
	"bytes"
	"context"
	"fmt"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
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
	Describe(`ListEnvironments(listEnvironmentsOptions *ListEnvironmentsOptions) - Operation response error`, func() {
		listEnvironmentsPath := "/environments"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listEnvironmentsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for expand query parameter

					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_time"}))

					Expect(req.URL.Query()["tags"]).To(Equal([]string{"version 1.1, pre-release"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				listEnvironmentsOptionsModel.Tags = core.StringPtr("version 1.1, pre-release")
				listEnvironmentsOptionsModel.Include = []string{"features"}
				listEnvironmentsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listEnvironmentsOptionsModel.Offset = core.Int64Ptr(int64(0))
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
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listEnvironmentsPath))
					Expect(req.Method).To(Equal("GET"))


					// TODO: Add check for expand query parameter

					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_time"}))

					Expect(req.URL.Query()["tags"]).To(Equal([]string{"version 1.1, pre-release"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"environments": [{"name": "Name", "environment_id": "EnvironmentID", "description": "Description", "tags": "Tags", "color_code": "#FDD13A", "created_time": "2019-01-01T12:00:00", "updated_time": "2019-01-01T12:00:00", "href": "Href", "features": [{"feature_id": "FeatureID", "name": "Name"}], "properties": [{"property_id": "PropertyID", "name": "Name"}]}], "limit": 5, "offset": 6, "total_count": 10, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}}`)
				}))
			})
			It(`Invoke ListEnvironments successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.ListEnvironments(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListEnvironmentsOptions model
				listEnvironmentsOptionsModel := new(appconfigurationv1.ListEnvironmentsOptions)
				listEnvironmentsOptionsModel.Expand = core.BoolPtr(true)
				listEnvironmentsOptionsModel.Sort = core.StringPtr("created_time")
				listEnvironmentsOptionsModel.Tags = core.StringPtr("version 1.1, pre-release")
				listEnvironmentsOptionsModel.Include = []string{"features"}
				listEnvironmentsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listEnvironmentsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listEnvironmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.ListEnvironments(listEnvironmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.ListEnvironmentsWithContext(ctx, listEnvironmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr = appConfigurationService.ListEnvironments(listEnvironmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.ListEnvironmentsWithContext(ctx, listEnvironmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
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
				listEnvironmentsOptionsModel.Tags = core.StringPtr("version 1.1, pre-release")
				listEnvironmentsOptionsModel.Include = []string{"features"}
				listEnvironmentsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listEnvironmentsOptionsModel.Offset = core.Int64Ptr(int64(0))
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
	})
	Describe(`CreateEnvironment(createEnvironmentOptions *CreateEnvironmentOptions) - Operation response error`, func() {
		createEnvironmentPath := "/environments"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEnvironmentPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
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
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Name", "environment_id": "EnvironmentID", "description": "Description", "tags": "Tags", "color_code": "#FDD13A", "created_time": "2019-01-01T12:00:00", "updated_time": "2019-01-01T12:00:00", "href": "Href", "features": [{"feature_id": "FeatureID", "name": "Name"}], "properties": [{"property_id": "PropertyID", "name": "Name"}]}`)
				}))
			})
			It(`Invoke CreateEnvironment successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

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

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.CreateEnvironmentWithContext(ctx, createEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr = appConfigurationService.CreateEnvironment(createEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.CreateEnvironmentWithContext(ctx, createEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
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
	})
	Describe(`UpdateEnvironment(updateEnvironmentOptions *UpdateEnvironmentOptions) - Operation response error`, func() {
		updateEnvironmentPath := "/environments/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateEnvironmentPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				updateEnvironmentOptionsModel.EnvironmentID = core.StringPtr("testString")
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
		updateEnvironmentPath := "/environments/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "environment_id": "EnvironmentID", "description": "Description", "tags": "Tags", "color_code": "#FDD13A", "created_time": "2019-01-01T12:00:00", "updated_time": "2019-01-01T12:00:00", "href": "Href", "features": [{"feature_id": "FeatureID", "name": "Name"}], "properties": [{"property_id": "PropertyID", "name": "Name"}]}`)
				}))
			})
			It(`Invoke UpdateEnvironment successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.UpdateEnvironment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateEnvironmentOptions model
				updateEnvironmentOptionsModel := new(appconfigurationv1.UpdateEnvironmentOptions)
				updateEnvironmentOptionsModel.EnvironmentID = core.StringPtr("testString")
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

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.UpdateEnvironmentWithContext(ctx, updateEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr = appConfigurationService.UpdateEnvironment(updateEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.UpdateEnvironmentWithContext(ctx, updateEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
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
				updateEnvironmentOptionsModel.EnvironmentID = core.StringPtr("testString")
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
	})
	Describe(`GetEnvironment(getEnvironmentOptions *GetEnvironmentOptions) - Operation response error`, func() {
		getEnvironmentPath := "/environments/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEnvironmentPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for expand query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				getEnvironmentOptionsModel.EnvironmentID = core.StringPtr("testString")
				getEnvironmentOptionsModel.Expand = core.BoolPtr(true)
				getEnvironmentOptionsModel.Include = []string{"features"}
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
		getEnvironmentPath := "/environments/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEnvironmentPath))
					Expect(req.Method).To(Equal("GET"))


					// TODO: Add check for expand query parameter

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "environment_id": "EnvironmentID", "description": "Description", "tags": "Tags", "color_code": "#FDD13A", "created_time": "2019-01-01T12:00:00", "updated_time": "2019-01-01T12:00:00", "href": "Href", "features": [{"feature_id": "FeatureID", "name": "Name"}], "properties": [{"property_id": "PropertyID", "name": "Name"}]}`)
				}))
			})
			It(`Invoke GetEnvironment successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.GetEnvironment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetEnvironmentOptions model
				getEnvironmentOptionsModel := new(appconfigurationv1.GetEnvironmentOptions)
				getEnvironmentOptionsModel.EnvironmentID = core.StringPtr("testString")
				getEnvironmentOptionsModel.Expand = core.BoolPtr(true)
				getEnvironmentOptionsModel.Include = []string{"features"}
				getEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.GetEnvironment(getEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.GetEnvironmentWithContext(ctx, getEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr = appConfigurationService.GetEnvironment(getEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.GetEnvironmentWithContext(ctx, getEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
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
				getEnvironmentOptionsModel.EnvironmentID = core.StringPtr("testString")
				getEnvironmentOptionsModel.Expand = core.BoolPtr(true)
				getEnvironmentOptionsModel.Include = []string{"features"}
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
	})

	Describe(`DeleteEnvironment(deleteEnvironmentOptions *DeleteEnvironmentOptions)`, func() {
		deleteEnvironmentPath := "/environments/testString"
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
				appConfigurationService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := appConfigurationService.DeleteEnvironment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteEnvironmentOptions model
				deleteEnvironmentOptionsModel := new(appconfigurationv1.DeleteEnvironmentOptions)
				deleteEnvironmentOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appConfigurationService.DeleteEnvironment(deleteEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				appConfigurationService.DisableRetries()
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
				deleteEnvironmentOptionsModel.EnvironmentID = core.StringPtr("testString")
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
	Describe(`ListCollections(listCollectionsOptions *ListCollectionsOptions) - Operation response error`, func() {
		listCollectionsPath := "/collections"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCollectionsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for expand query parameter

					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_time"}))

					Expect(req.URL.Query()["tags"]).To(Equal([]string{"version 1.1, pre-release"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				listCollectionsOptionsModel.Tags = core.StringPtr("version 1.1, pre-release")
				listCollectionsOptionsModel.Features = []string{"testString"}
				listCollectionsOptionsModel.Properties = []string{"testString"}
				listCollectionsOptionsModel.Include = []string{"features"}
				listCollectionsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listCollectionsOptionsModel.Offset = core.Int64Ptr(int64(0))
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
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCollectionsPath))
					Expect(req.Method).To(Equal("GET"))


					// TODO: Add check for expand query parameter

					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_time"}))

					Expect(req.URL.Query()["tags"]).To(Equal([]string{"version 1.1, pre-release"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"collections": [{"name": "Name", "collection_id": "CollectionID", "description": "Description", "tags": "Tags", "created_time": "2019-01-01T12:00:00", "updated_time": "2019-01-01T12:00:00", "href": "Href", "features": [{"feature_id": "FeatureID", "name": "Name"}], "properties": [{"property_id": "PropertyID", "name": "Name"}], "features_count": 13, "properties_count": 15}], "limit": 5, "offset": 6, "total_count": 10, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}}`)
				}))
			})
			It(`Invoke ListCollections successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.ListCollections(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListCollectionsOptions model
				listCollectionsOptionsModel := new(appconfigurationv1.ListCollectionsOptions)
				listCollectionsOptionsModel.Expand = core.BoolPtr(true)
				listCollectionsOptionsModel.Sort = core.StringPtr("created_time")
				listCollectionsOptionsModel.Tags = core.StringPtr("version 1.1, pre-release")
				listCollectionsOptionsModel.Features = []string{"testString"}
				listCollectionsOptionsModel.Properties = []string{"testString"}
				listCollectionsOptionsModel.Include = []string{"features"}
				listCollectionsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listCollectionsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listCollectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.ListCollections(listCollectionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.ListCollectionsWithContext(ctx, listCollectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr = appConfigurationService.ListCollections(listCollectionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.ListCollectionsWithContext(ctx, listCollectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
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
				listCollectionsOptionsModel.Tags = core.StringPtr("version 1.1, pre-release")
				listCollectionsOptionsModel.Features = []string{"testString"}
				listCollectionsOptionsModel.Properties = []string{"testString"}
				listCollectionsOptionsModel.Include = []string{"features"}
				listCollectionsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listCollectionsOptionsModel.Offset = core.Int64Ptr(int64(0))
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
	})
	Describe(`CreateCollection(createCollectionOptions *CreateCollectionOptions) - Operation response error`, func() {
		createCollectionPath := "/collections"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCollectionPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
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
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Name", "collection_id": "CollectionID", "description": "Description", "tags": "Tags", "created_time": "2019-01-01T12:00:00", "updated_time": "2019-01-01T12:00:00", "href": "Href"}`)
				}))
			})
			It(`Invoke CreateCollection successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

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

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.CreateCollectionWithContext(ctx, createCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr = appConfigurationService.CreateCollection(createCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.CreateCollectionWithContext(ctx, createCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
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
	})
	Describe(`UpdateCollection(updateCollectionOptions *UpdateCollectionOptions) - Operation response error`, func() {
		updateCollectionPath := "/collections/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCollectionPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				updateCollectionOptionsModel.CollectionID = core.StringPtr("testString")
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
		updateCollectionPath := "/collections/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "collection_id": "CollectionID", "description": "Description", "tags": "Tags", "created_time": "2019-01-01T12:00:00", "updated_time": "2019-01-01T12:00:00", "href": "Href"}`)
				}))
			})
			It(`Invoke UpdateCollection successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.UpdateCollection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateCollectionOptions model
				updateCollectionOptionsModel := new(appconfigurationv1.UpdateCollectionOptions)
				updateCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				updateCollectionOptionsModel.Name = core.StringPtr("testString")
				updateCollectionOptionsModel.Description = core.StringPtr("testString")
				updateCollectionOptionsModel.Tags = core.StringPtr("testString")
				updateCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.UpdateCollection(updateCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.UpdateCollectionWithContext(ctx, updateCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr = appConfigurationService.UpdateCollection(updateCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.UpdateCollectionWithContext(ctx, updateCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
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
				updateCollectionOptionsModel.CollectionID = core.StringPtr("testString")
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
	})
	Describe(`GetCollection(getCollectionOptions *GetCollectionOptions) - Operation response error`, func() {
		getCollectionPath := "/collections/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCollectionPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for expand query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				getCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				getCollectionOptionsModel.Expand = core.BoolPtr(true)
				getCollectionOptionsModel.Include = []string{"features"}
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
		getCollectionPath := "/collections/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCollectionPath))
					Expect(req.Method).To(Equal("GET"))


					// TODO: Add check for expand query parameter

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "collection_id": "CollectionID", "description": "Description", "tags": "Tags", "created_time": "2019-01-01T12:00:00", "updated_time": "2019-01-01T12:00:00", "href": "Href", "features": [{"feature_id": "FeatureID", "name": "Name"}], "properties": [{"property_id": "PropertyID", "name": "Name"}], "features_count": 13, "properties_count": 15}`)
				}))
			})
			It(`Invoke GetCollection successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.GetCollection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCollectionOptions model
				getCollectionOptionsModel := new(appconfigurationv1.GetCollectionOptions)
				getCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				getCollectionOptionsModel.Expand = core.BoolPtr(true)
				getCollectionOptionsModel.Include = []string{"features"}
				getCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.GetCollection(getCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.GetCollectionWithContext(ctx, getCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr = appConfigurationService.GetCollection(getCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.GetCollectionWithContext(ctx, getCollectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
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
				getCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				getCollectionOptionsModel.Expand = core.BoolPtr(true)
				getCollectionOptionsModel.Include = []string{"features"}
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
	})

	Describe(`DeleteCollection(deleteCollectionOptions *DeleteCollectionOptions)`, func() {
		deleteCollectionPath := "/collections/testString"
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
				appConfigurationService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := appConfigurationService.DeleteCollection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteCollectionOptions model
				deleteCollectionOptionsModel := new(appconfigurationv1.DeleteCollectionOptions)
				deleteCollectionOptionsModel.CollectionID = core.StringPtr("testString")
				deleteCollectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appConfigurationService.DeleteCollection(deleteCollectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				appConfigurationService.DisableRetries()
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
				deleteCollectionOptionsModel.CollectionID = core.StringPtr("testString")
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
	Describe(`ListFeatures(listFeaturesOptions *ListFeaturesOptions) - Operation response error`, func() {
		listFeaturesPath := "/environments/testString/features"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listFeaturesPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for expand query parameter

					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_time"}))

					Expect(req.URL.Query()["tags"]).To(Equal([]string{"version 1.1, pre-release"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				listFeaturesOptionsModel.EnvironmentID = core.StringPtr("testString")
				listFeaturesOptionsModel.Expand = core.BoolPtr(true)
				listFeaturesOptionsModel.Sort = core.StringPtr("created_time")
				listFeaturesOptionsModel.Tags = core.StringPtr("version 1.1, pre-release")
				listFeaturesOptionsModel.Collections = []string{"testString"}
				listFeaturesOptionsModel.Segments = []string{"testString"}
				listFeaturesOptionsModel.Include = []string{"collections"}
				listFeaturesOptionsModel.Limit = core.Int64Ptr(int64(1))
				listFeaturesOptionsModel.Offset = core.Int64Ptr(int64(0))
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
		listFeaturesPath := "/environments/testString/features"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listFeaturesPath))
					Expect(req.Method).To(Equal("GET"))


					// TODO: Add check for expand query parameter

					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_time"}))

					Expect(req.URL.Query()["tags"]).To(Equal([]string{"version 1.1, pre-release"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"features": [{"name": "Name", "feature_id": "FeatureID", "description": "Description", "type": "BOOLEAN", "enabled_value": "anyValue", "disabled_value": "anyValue", "enabled": false, "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "created_time": "2019-01-01T12:00:00", "updated_time": "2019-01-01T12:00:00", "evaluation_time": "2019-01-01T12:00:00", "href": "Href"}], "limit": 5, "offset": 6, "total_count": 10, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}}`)
				}))
			})
			It(`Invoke ListFeatures successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.ListFeatures(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListFeaturesOptions model
				listFeaturesOptionsModel := new(appconfigurationv1.ListFeaturesOptions)
				listFeaturesOptionsModel.EnvironmentID = core.StringPtr("testString")
				listFeaturesOptionsModel.Expand = core.BoolPtr(true)
				listFeaturesOptionsModel.Sort = core.StringPtr("created_time")
				listFeaturesOptionsModel.Tags = core.StringPtr("version 1.1, pre-release")
				listFeaturesOptionsModel.Collections = []string{"testString"}
				listFeaturesOptionsModel.Segments = []string{"testString"}
				listFeaturesOptionsModel.Include = []string{"collections"}
				listFeaturesOptionsModel.Limit = core.Int64Ptr(int64(1))
				listFeaturesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listFeaturesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.ListFeatures(listFeaturesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.ListFeaturesWithContext(ctx, listFeaturesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr = appConfigurationService.ListFeatures(listFeaturesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.ListFeaturesWithContext(ctx, listFeaturesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
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
				listFeaturesOptionsModel.EnvironmentID = core.StringPtr("testString")
				listFeaturesOptionsModel.Expand = core.BoolPtr(true)
				listFeaturesOptionsModel.Sort = core.StringPtr("created_time")
				listFeaturesOptionsModel.Tags = core.StringPtr("version 1.1, pre-release")
				listFeaturesOptionsModel.Collections = []string{"testString"}
				listFeaturesOptionsModel.Segments = []string{"testString"}
				listFeaturesOptionsModel.Include = []string{"collections"}
				listFeaturesOptionsModel.Limit = core.Int64Ptr(int64(1))
				listFeaturesOptionsModel.Offset = core.Int64Ptr(int64(0))
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
	})
	Describe(`CreateFeature(createFeatureOptions *CreateFeatureOptions) - Operation response error`, func() {
		createFeaturePath := "/environments/testString/features"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createFeaturePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				targetSegmentsModel.Segments = []string{"testString"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = core.StringPtr("true")
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				collectionRefModel.CollectionID = core.StringPtr("ghzinc")

				// Construct an instance of the CreateFeatureOptions model
				createFeatureOptionsModel := new(appconfigurationv1.CreateFeatureOptions)
				createFeatureOptionsModel.EnvironmentID = core.StringPtr("testString")
				createFeatureOptionsModel.Name = core.StringPtr("Cycle Rentals")
				createFeatureOptionsModel.FeatureID = core.StringPtr("cycle-rentals")
				createFeatureOptionsModel.Type = core.StringPtr("BOOLEAN")
				createFeatureOptionsModel.EnabledValue = core.StringPtr("true")
				createFeatureOptionsModel.DisabledValue = core.StringPtr("false")
				createFeatureOptionsModel.Description = core.StringPtr("Feature flag to enable Cycle Rentals")
				createFeatureOptionsModel.Enabled = core.BoolPtr(true)
				createFeatureOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				createFeatureOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
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
		createFeaturePath := "/environments/testString/features"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Name", "feature_id": "FeatureID", "description": "Description", "type": "BOOLEAN", "enabled_value": "anyValue", "disabled_value": "anyValue", "enabled": false, "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "created_time": "2019-01-01T12:00:00", "updated_time": "2019-01-01T12:00:00", "evaluation_time": "2019-01-01T12:00:00", "href": "Href"}`)
				}))
			})
			It(`Invoke CreateFeature successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.CreateFeature(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"testString"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = core.StringPtr("true")
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				collectionRefModel.CollectionID = core.StringPtr("ghzinc")

				// Construct an instance of the CreateFeatureOptions model
				createFeatureOptionsModel := new(appconfigurationv1.CreateFeatureOptions)
				createFeatureOptionsModel.EnvironmentID = core.StringPtr("testString")
				createFeatureOptionsModel.Name = core.StringPtr("Cycle Rentals")
				createFeatureOptionsModel.FeatureID = core.StringPtr("cycle-rentals")
				createFeatureOptionsModel.Type = core.StringPtr("BOOLEAN")
				createFeatureOptionsModel.EnabledValue = core.StringPtr("true")
				createFeatureOptionsModel.DisabledValue = core.StringPtr("false")
				createFeatureOptionsModel.Description = core.StringPtr("Feature flag to enable Cycle Rentals")
				createFeatureOptionsModel.Enabled = core.BoolPtr(true)
				createFeatureOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				createFeatureOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				createFeatureOptionsModel.Collections = []appconfigurationv1.CollectionRef{*collectionRefModel}
				createFeatureOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.CreateFeature(createFeatureOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.CreateFeatureWithContext(ctx, createFeatureOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr = appConfigurationService.CreateFeature(createFeatureOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.CreateFeatureWithContext(ctx, createFeatureOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
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
				targetSegmentsModel.Segments = []string{"testString"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = core.StringPtr("true")
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				collectionRefModel.CollectionID = core.StringPtr("ghzinc")

				// Construct an instance of the CreateFeatureOptions model
				createFeatureOptionsModel := new(appconfigurationv1.CreateFeatureOptions)
				createFeatureOptionsModel.EnvironmentID = core.StringPtr("testString")
				createFeatureOptionsModel.Name = core.StringPtr("Cycle Rentals")
				createFeatureOptionsModel.FeatureID = core.StringPtr("cycle-rentals")
				createFeatureOptionsModel.Type = core.StringPtr("BOOLEAN")
				createFeatureOptionsModel.EnabledValue = core.StringPtr("true")
				createFeatureOptionsModel.DisabledValue = core.StringPtr("false")
				createFeatureOptionsModel.Description = core.StringPtr("Feature flag to enable Cycle Rentals")
				createFeatureOptionsModel.Enabled = core.BoolPtr(true)
				createFeatureOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				createFeatureOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
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
	})
	Describe(`UpdateFeature(updateFeatureOptions *UpdateFeatureOptions) - Operation response error`, func() {
		updateFeaturePath := "/environments/testString/features/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateFeaturePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				targetSegmentsModel.Segments = []string{"testString"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = core.StringPtr("testString")
				segmentRuleModel.Order = core.Int64Ptr(int64(38))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				collectionRefModel.CollectionID = core.StringPtr("testString")

				// Construct an instance of the UpdateFeatureOptions model
				updateFeatureOptionsModel := new(appconfigurationv1.UpdateFeatureOptions)
				updateFeatureOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateFeatureOptionsModel.FeatureID = core.StringPtr("testString")
				updateFeatureOptionsModel.Name = core.StringPtr("testString")
				updateFeatureOptionsModel.Description = core.StringPtr("testString")
				updateFeatureOptionsModel.EnabledValue = core.StringPtr("testString")
				updateFeatureOptionsModel.DisabledValue = core.StringPtr("testString")
				updateFeatureOptionsModel.Enabled = core.BoolPtr(true)
				updateFeatureOptionsModel.Tags = core.StringPtr("testString")
				updateFeatureOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				updateFeatureOptionsModel.Collections = []appconfigurationv1.CollectionRef{*collectionRefModel}
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
		updateFeaturePath := "/environments/testString/features/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "feature_id": "FeatureID", "description": "Description", "type": "BOOLEAN", "enabled_value": "anyValue", "disabled_value": "anyValue", "enabled": false, "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "created_time": "2019-01-01T12:00:00", "updated_time": "2019-01-01T12:00:00", "evaluation_time": "2019-01-01T12:00:00", "href": "Href"}`)
				}))
			})
			It(`Invoke UpdateFeature successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.UpdateFeature(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"testString"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = core.StringPtr("testString")
				segmentRuleModel.Order = core.Int64Ptr(int64(38))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				collectionRefModel.CollectionID = core.StringPtr("testString")

				// Construct an instance of the UpdateFeatureOptions model
				updateFeatureOptionsModel := new(appconfigurationv1.UpdateFeatureOptions)
				updateFeatureOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateFeatureOptionsModel.FeatureID = core.StringPtr("testString")
				updateFeatureOptionsModel.Name = core.StringPtr("testString")
				updateFeatureOptionsModel.Description = core.StringPtr("testString")
				updateFeatureOptionsModel.EnabledValue = core.StringPtr("testString")
				updateFeatureOptionsModel.DisabledValue = core.StringPtr("testString")
				updateFeatureOptionsModel.Enabled = core.BoolPtr(true)
				updateFeatureOptionsModel.Tags = core.StringPtr("testString")
				updateFeatureOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				updateFeatureOptionsModel.Collections = []appconfigurationv1.CollectionRef{*collectionRefModel}
				updateFeatureOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.UpdateFeature(updateFeatureOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.UpdateFeatureWithContext(ctx, updateFeatureOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr = appConfigurationService.UpdateFeature(updateFeatureOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.UpdateFeatureWithContext(ctx, updateFeatureOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
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
				targetSegmentsModel.Segments = []string{"testString"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = core.StringPtr("testString")
				segmentRuleModel.Order = core.Int64Ptr(int64(38))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				collectionRefModel.CollectionID = core.StringPtr("testString")

				// Construct an instance of the UpdateFeatureOptions model
				updateFeatureOptionsModel := new(appconfigurationv1.UpdateFeatureOptions)
				updateFeatureOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateFeatureOptionsModel.FeatureID = core.StringPtr("testString")
				updateFeatureOptionsModel.Name = core.StringPtr("testString")
				updateFeatureOptionsModel.Description = core.StringPtr("testString")
				updateFeatureOptionsModel.EnabledValue = core.StringPtr("testString")
				updateFeatureOptionsModel.DisabledValue = core.StringPtr("testString")
				updateFeatureOptionsModel.Enabled = core.BoolPtr(true)
				updateFeatureOptionsModel.Tags = core.StringPtr("testString")
				updateFeatureOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				updateFeatureOptionsModel.Collections = []appconfigurationv1.CollectionRef{*collectionRefModel}
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
	})
	Describe(`UpdateFeatureValues(updateFeatureValuesOptions *UpdateFeatureValuesOptions) - Operation response error`, func() {
		updateFeatureValuesPath := "/environments/testString/features/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateFeatureValuesPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				targetSegmentsModel.Segments = []string{"testString"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = core.StringPtr("testString")
				segmentRuleModel.Order = core.Int64Ptr(int64(38))

				// Construct an instance of the UpdateFeatureValuesOptions model
				updateFeatureValuesOptionsModel := new(appconfigurationv1.UpdateFeatureValuesOptions)
				updateFeatureValuesOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateFeatureValuesOptionsModel.FeatureID = core.StringPtr("testString")
				updateFeatureValuesOptionsModel.Name = core.StringPtr("testString")
				updateFeatureValuesOptionsModel.Description = core.StringPtr("testString")
				updateFeatureValuesOptionsModel.Tags = core.StringPtr("testString")
				updateFeatureValuesOptionsModel.EnabledValue = core.StringPtr("testString")
				updateFeatureValuesOptionsModel.DisabledValue = core.StringPtr("testString")
				updateFeatureValuesOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
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
		updateFeatureValuesPath := "/environments/testString/features/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "feature_id": "FeatureID", "description": "Description", "type": "BOOLEAN", "enabled_value": "anyValue", "disabled_value": "anyValue", "enabled": false, "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "created_time": "2019-01-01T12:00:00", "updated_time": "2019-01-01T12:00:00", "evaluation_time": "2019-01-01T12:00:00", "href": "Href"}`)
				}))
			})
			It(`Invoke UpdateFeatureValues successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.UpdateFeatureValues(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"testString"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = core.StringPtr("testString")
				segmentRuleModel.Order = core.Int64Ptr(int64(38))

				// Construct an instance of the UpdateFeatureValuesOptions model
				updateFeatureValuesOptionsModel := new(appconfigurationv1.UpdateFeatureValuesOptions)
				updateFeatureValuesOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateFeatureValuesOptionsModel.FeatureID = core.StringPtr("testString")
				updateFeatureValuesOptionsModel.Name = core.StringPtr("testString")
				updateFeatureValuesOptionsModel.Description = core.StringPtr("testString")
				updateFeatureValuesOptionsModel.Tags = core.StringPtr("testString")
				updateFeatureValuesOptionsModel.EnabledValue = core.StringPtr("testString")
				updateFeatureValuesOptionsModel.DisabledValue = core.StringPtr("testString")
				updateFeatureValuesOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				updateFeatureValuesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.UpdateFeatureValues(updateFeatureValuesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.UpdateFeatureValuesWithContext(ctx, updateFeatureValuesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr = appConfigurationService.UpdateFeatureValues(updateFeatureValuesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.UpdateFeatureValuesWithContext(ctx, updateFeatureValuesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
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
				targetSegmentsModel.Segments = []string{"testString"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = core.StringPtr("testString")
				segmentRuleModel.Order = core.Int64Ptr(int64(38))

				// Construct an instance of the UpdateFeatureValuesOptions model
				updateFeatureValuesOptionsModel := new(appconfigurationv1.UpdateFeatureValuesOptions)
				updateFeatureValuesOptionsModel.EnvironmentID = core.StringPtr("testString")
				updateFeatureValuesOptionsModel.FeatureID = core.StringPtr("testString")
				updateFeatureValuesOptionsModel.Name = core.StringPtr("testString")
				updateFeatureValuesOptionsModel.Description = core.StringPtr("testString")
				updateFeatureValuesOptionsModel.Tags = core.StringPtr("testString")
				updateFeatureValuesOptionsModel.EnabledValue = core.StringPtr("testString")
				updateFeatureValuesOptionsModel.DisabledValue = core.StringPtr("testString")
				updateFeatureValuesOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
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
	})
	Describe(`GetFeature(getFeatureOptions *GetFeatureOptions) - Operation response error`, func() {
		getFeaturePath := "/environments/testString/features/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getFeaturePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["include"]).To(Equal([]string{"collections"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				getFeatureOptionsModel.EnvironmentID = core.StringPtr("testString")
				getFeatureOptionsModel.FeatureID = core.StringPtr("testString")
				getFeatureOptionsModel.Include = core.StringPtr("collections")
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
		getFeaturePath := "/environments/testString/features/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getFeaturePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["include"]).To(Equal([]string{"collections"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "feature_id": "FeatureID", "description": "Description", "type": "BOOLEAN", "enabled_value": "anyValue", "disabled_value": "anyValue", "enabled": false, "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "created_time": "2019-01-01T12:00:00", "updated_time": "2019-01-01T12:00:00", "evaluation_time": "2019-01-01T12:00:00", "href": "Href"}`)
				}))
			})
			It(`Invoke GetFeature successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.GetFeature(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetFeatureOptions model
				getFeatureOptionsModel := new(appconfigurationv1.GetFeatureOptions)
				getFeatureOptionsModel.EnvironmentID = core.StringPtr("testString")
				getFeatureOptionsModel.FeatureID = core.StringPtr("testString")
				getFeatureOptionsModel.Include = core.StringPtr("collections")
				getFeatureOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.GetFeature(getFeatureOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.GetFeatureWithContext(ctx, getFeatureOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr = appConfigurationService.GetFeature(getFeatureOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.GetFeatureWithContext(ctx, getFeatureOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
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
				getFeatureOptionsModel.EnvironmentID = core.StringPtr("testString")
				getFeatureOptionsModel.FeatureID = core.StringPtr("testString")
				getFeatureOptionsModel.Include = core.StringPtr("collections")
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
	})

	Describe(`DeleteFeature(deleteFeatureOptions *DeleteFeatureOptions)`, func() {
		deleteFeaturePath := "/environments/testString/features/testString"
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
				appConfigurationService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := appConfigurationService.DeleteFeature(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteFeatureOptions model
				deleteFeatureOptionsModel := new(appconfigurationv1.DeleteFeatureOptions)
				deleteFeatureOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteFeatureOptionsModel.FeatureID = core.StringPtr("testString")
				deleteFeatureOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appConfigurationService.DeleteFeature(deleteFeatureOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				appConfigurationService.DisableRetries()
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
				deleteFeatureOptionsModel.EnvironmentID = core.StringPtr("testString")
				deleteFeatureOptionsModel.FeatureID = core.StringPtr("testString")
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
		toggleFeaturePath := "/environments/testString/features/testString/toggle"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(toggleFeaturePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				toggleFeatureOptionsModel.EnvironmentID = core.StringPtr("testString")
				toggleFeatureOptionsModel.FeatureID = core.StringPtr("testString")
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
		toggleFeaturePath := "/environments/testString/features/testString/toggle"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "feature_id": "FeatureID", "description": "Description", "type": "BOOLEAN", "enabled_value": "anyValue", "disabled_value": "anyValue", "enabled": false, "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "created_time": "2019-01-01T12:00:00", "updated_time": "2019-01-01T12:00:00", "evaluation_time": "2019-01-01T12:00:00", "href": "Href"}`)
				}))
			})
			It(`Invoke ToggleFeature successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.ToggleFeature(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ToggleFeatureOptions model
				toggleFeatureOptionsModel := new(appconfigurationv1.ToggleFeatureOptions)
				toggleFeatureOptionsModel.EnvironmentID = core.StringPtr("testString")
				toggleFeatureOptionsModel.FeatureID = core.StringPtr("testString")
				toggleFeatureOptionsModel.Enabled = core.BoolPtr(true)
				toggleFeatureOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.ToggleFeature(toggleFeatureOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.ToggleFeatureWithContext(ctx, toggleFeatureOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr = appConfigurationService.ToggleFeature(toggleFeatureOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.ToggleFeatureWithContext(ctx, toggleFeatureOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
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
				toggleFeatureOptionsModel.EnvironmentID = core.StringPtr("testString")
				toggleFeatureOptionsModel.FeatureID = core.StringPtr("testString")
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
	})
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
	Describe(`ListProperties(listPropertiesOptions *ListPropertiesOptions) - Operation response error`, func() {
		listPropertiesPath := "/environments/testString/properties"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPropertiesPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for expand query parameter

					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_time"}))

					Expect(req.URL.Query()["tags"]).To(Equal([]string{"version 1.1, pre-release"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				listPropertiesOptionsModel.EnvironmentID = core.StringPtr("testString")
				listPropertiesOptionsModel.Expand = core.BoolPtr(true)
				listPropertiesOptionsModel.Sort = core.StringPtr("created_time")
				listPropertiesOptionsModel.Tags = core.StringPtr("version 1.1, pre-release")
				listPropertiesOptionsModel.Collections = []string{"testString"}
				listPropertiesOptionsModel.Segments = []string{"testString"}
				listPropertiesOptionsModel.Include = []string{"collections"}
				listPropertiesOptionsModel.Limit = core.Int64Ptr(int64(1))
				listPropertiesOptionsModel.Offset = core.Int64Ptr(int64(0))
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
		listPropertiesPath := "/environments/testString/properties"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPropertiesPath))
					Expect(req.Method).To(Equal("GET"))


					// TODO: Add check for expand query parameter

					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_time"}))

					Expect(req.URL.Query()["tags"]).To(Equal([]string{"version 1.1, pre-release"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"properties": [{"name": "Name", "property_id": "PropertyID", "description": "Description", "type": "BOOLEAN", "value": "anyValue", "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "created_time": "2019-01-01T12:00:00", "updated_time": "2019-01-01T12:00:00", "evaluation_time": "2019-01-01T12:00:00", "href": "Href"}], "limit": 5, "offset": 6, "total_count": 10, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}}`)
				}))
			})
			It(`Invoke ListProperties successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.ListProperties(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListPropertiesOptions model
				listPropertiesOptionsModel := new(appconfigurationv1.ListPropertiesOptions)
				listPropertiesOptionsModel.EnvironmentID = core.StringPtr("testString")
				listPropertiesOptionsModel.Expand = core.BoolPtr(true)
				listPropertiesOptionsModel.Sort = core.StringPtr("created_time")
				listPropertiesOptionsModel.Tags = core.StringPtr("version 1.1, pre-release")
				listPropertiesOptionsModel.Collections = []string{"testString"}
				listPropertiesOptionsModel.Segments = []string{"testString"}
				listPropertiesOptionsModel.Include = []string{"collections"}
				listPropertiesOptionsModel.Limit = core.Int64Ptr(int64(1))
				listPropertiesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listPropertiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.ListProperties(listPropertiesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.ListPropertiesWithContext(ctx, listPropertiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr = appConfigurationService.ListProperties(listPropertiesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.ListPropertiesWithContext(ctx, listPropertiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
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
				listPropertiesOptionsModel.EnvironmentID = core.StringPtr("testString")
				listPropertiesOptionsModel.Expand = core.BoolPtr(true)
				listPropertiesOptionsModel.Sort = core.StringPtr("created_time")
				listPropertiesOptionsModel.Tags = core.StringPtr("version 1.1, pre-release")
				listPropertiesOptionsModel.Collections = []string{"testString"}
				listPropertiesOptionsModel.Segments = []string{"testString"}
				listPropertiesOptionsModel.Include = []string{"collections"}
				listPropertiesOptionsModel.Limit = core.Int64Ptr(int64(1))
				listPropertiesOptionsModel.Offset = core.Int64Ptr(int64(0))
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
	})
	Describe(`CreateProperty(createPropertyOptions *CreatePropertyOptions) - Operation response error`, func() {
		createPropertyPath := "/environments/testString/properties"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createPropertyPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				targetSegmentsModel.Segments = []string{"testString"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = core.StringPtr("true")
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				collectionRefModel.CollectionID = core.StringPtr("ghzinc")

				// Construct an instance of the CreatePropertyOptions model
				createPropertyOptionsModel := new(appconfigurationv1.CreatePropertyOptions)
				createPropertyOptionsModel.EnvironmentID = core.StringPtr("testString")
				createPropertyOptionsModel.Name = core.StringPtr("Email property")
				createPropertyOptionsModel.PropertyID = core.StringPtr("email-property")
				createPropertyOptionsModel.Type = core.StringPtr("BOOLEAN")
				createPropertyOptionsModel.Value = core.StringPtr("true")
				createPropertyOptionsModel.Description = core.StringPtr("Property for email")
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
		createPropertyPath := "/environments/testString/properties"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Name", "property_id": "PropertyID", "description": "Description", "type": "BOOLEAN", "value": "anyValue", "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "created_time": "2019-01-01T12:00:00", "updated_time": "2019-01-01T12:00:00", "evaluation_time": "2019-01-01T12:00:00", "href": "Href"}`)
				}))
			})
			It(`Invoke CreateProperty successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.CreateProperty(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"testString"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = core.StringPtr("true")
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				collectionRefModel.CollectionID = core.StringPtr("ghzinc")

				// Construct an instance of the CreatePropertyOptions model
				createPropertyOptionsModel := new(appconfigurationv1.CreatePropertyOptions)
				createPropertyOptionsModel.EnvironmentID = core.StringPtr("testString")
				createPropertyOptionsModel.Name = core.StringPtr("Email property")
				createPropertyOptionsModel.PropertyID = core.StringPtr("email-property")
				createPropertyOptionsModel.Type = core.StringPtr("BOOLEAN")
				createPropertyOptionsModel.Value = core.StringPtr("true")
				createPropertyOptionsModel.Description = core.StringPtr("Property for email")
				createPropertyOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				createPropertyOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				createPropertyOptionsModel.Collections = []appconfigurationv1.CollectionRef{*collectionRefModel}
				createPropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.CreateProperty(createPropertyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.CreatePropertyWithContext(ctx, createPropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr = appConfigurationService.CreateProperty(createPropertyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.CreatePropertyWithContext(ctx, createPropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
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
				targetSegmentsModel.Segments = []string{"testString"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = core.StringPtr("true")
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				collectionRefModel.CollectionID = core.StringPtr("ghzinc")

				// Construct an instance of the CreatePropertyOptions model
				createPropertyOptionsModel := new(appconfigurationv1.CreatePropertyOptions)
				createPropertyOptionsModel.EnvironmentID = core.StringPtr("testString")
				createPropertyOptionsModel.Name = core.StringPtr("Email property")
				createPropertyOptionsModel.PropertyID = core.StringPtr("email-property")
				createPropertyOptionsModel.Type = core.StringPtr("BOOLEAN")
				createPropertyOptionsModel.Value = core.StringPtr("true")
				createPropertyOptionsModel.Description = core.StringPtr("Property for email")
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
	})
	Describe(`UpdateProperty(updatePropertyOptions *UpdatePropertyOptions) - Operation response error`, func() {
		updatePropertyPath := "/environments/testString/properties/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updatePropertyPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				targetSegmentsModel.Segments = []string{"testString"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = core.StringPtr("testString")
				segmentRuleModel.Order = core.Int64Ptr(int64(38))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				collectionRefModel.CollectionID = core.StringPtr("testString")

				// Construct an instance of the UpdatePropertyOptions model
				updatePropertyOptionsModel := new(appconfigurationv1.UpdatePropertyOptions)
				updatePropertyOptionsModel.EnvironmentID = core.StringPtr("testString")
				updatePropertyOptionsModel.PropertyID = core.StringPtr("testString")
				updatePropertyOptionsModel.Name = core.StringPtr("testString")
				updatePropertyOptionsModel.Description = core.StringPtr("testString")
				updatePropertyOptionsModel.Value = core.StringPtr("testString")
				updatePropertyOptionsModel.Tags = core.StringPtr("testString")
				updatePropertyOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				updatePropertyOptionsModel.Collections = []appconfigurationv1.CollectionRef{*collectionRefModel}
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
		updatePropertyPath := "/environments/testString/properties/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "property_id": "PropertyID", "description": "Description", "type": "BOOLEAN", "value": "anyValue", "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "created_time": "2019-01-01T12:00:00", "updated_time": "2019-01-01T12:00:00", "evaluation_time": "2019-01-01T12:00:00", "href": "Href"}`)
				}))
			})
			It(`Invoke UpdateProperty successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.UpdateProperty(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"testString"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = core.StringPtr("testString")
				segmentRuleModel.Order = core.Int64Ptr(int64(38))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				collectionRefModel.CollectionID = core.StringPtr("testString")

				// Construct an instance of the UpdatePropertyOptions model
				updatePropertyOptionsModel := new(appconfigurationv1.UpdatePropertyOptions)
				updatePropertyOptionsModel.EnvironmentID = core.StringPtr("testString")
				updatePropertyOptionsModel.PropertyID = core.StringPtr("testString")
				updatePropertyOptionsModel.Name = core.StringPtr("testString")
				updatePropertyOptionsModel.Description = core.StringPtr("testString")
				updatePropertyOptionsModel.Value = core.StringPtr("testString")
				updatePropertyOptionsModel.Tags = core.StringPtr("testString")
				updatePropertyOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				updatePropertyOptionsModel.Collections = []appconfigurationv1.CollectionRef{*collectionRefModel}
				updatePropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.UpdateProperty(updatePropertyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.UpdatePropertyWithContext(ctx, updatePropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr = appConfigurationService.UpdateProperty(updatePropertyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.UpdatePropertyWithContext(ctx, updatePropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
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
				targetSegmentsModel.Segments = []string{"testString"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = core.StringPtr("testString")
				segmentRuleModel.Order = core.Int64Ptr(int64(38))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				collectionRefModel.CollectionID = core.StringPtr("testString")

				// Construct an instance of the UpdatePropertyOptions model
				updatePropertyOptionsModel := new(appconfigurationv1.UpdatePropertyOptions)
				updatePropertyOptionsModel.EnvironmentID = core.StringPtr("testString")
				updatePropertyOptionsModel.PropertyID = core.StringPtr("testString")
				updatePropertyOptionsModel.Name = core.StringPtr("testString")
				updatePropertyOptionsModel.Description = core.StringPtr("testString")
				updatePropertyOptionsModel.Value = core.StringPtr("testString")
				updatePropertyOptionsModel.Tags = core.StringPtr("testString")
				updatePropertyOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				updatePropertyOptionsModel.Collections = []appconfigurationv1.CollectionRef{*collectionRefModel}
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
	})
	Describe(`UpdatePropertyValues(updatePropertyValuesOptions *UpdatePropertyValuesOptions) - Operation response error`, func() {
		updatePropertyValuesPath := "/environments/testString/properties/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updatePropertyValuesPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				targetSegmentsModel.Segments = []string{"testString"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = core.StringPtr("testString")
				segmentRuleModel.Order = core.Int64Ptr(int64(38))

				// Construct an instance of the UpdatePropertyValuesOptions model
				updatePropertyValuesOptionsModel := new(appconfigurationv1.UpdatePropertyValuesOptions)
				updatePropertyValuesOptionsModel.EnvironmentID = core.StringPtr("testString")
				updatePropertyValuesOptionsModel.PropertyID = core.StringPtr("testString")
				updatePropertyValuesOptionsModel.Name = core.StringPtr("testString")
				updatePropertyValuesOptionsModel.Description = core.StringPtr("testString")
				updatePropertyValuesOptionsModel.Tags = core.StringPtr("testString")
				updatePropertyValuesOptionsModel.Value = core.StringPtr("testString")
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
		updatePropertyValuesPath := "/environments/testString/properties/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "property_id": "PropertyID", "description": "Description", "type": "BOOLEAN", "value": "anyValue", "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "created_time": "2019-01-01T12:00:00", "updated_time": "2019-01-01T12:00:00", "evaluation_time": "2019-01-01T12:00:00", "href": "Href"}`)
				}))
			})
			It(`Invoke UpdatePropertyValues successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.UpdatePropertyValues(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				targetSegmentsModel.Segments = []string{"testString"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = core.StringPtr("testString")
				segmentRuleModel.Order = core.Int64Ptr(int64(38))

				// Construct an instance of the UpdatePropertyValuesOptions model
				updatePropertyValuesOptionsModel := new(appconfigurationv1.UpdatePropertyValuesOptions)
				updatePropertyValuesOptionsModel.EnvironmentID = core.StringPtr("testString")
				updatePropertyValuesOptionsModel.PropertyID = core.StringPtr("testString")
				updatePropertyValuesOptionsModel.Name = core.StringPtr("testString")
				updatePropertyValuesOptionsModel.Description = core.StringPtr("testString")
				updatePropertyValuesOptionsModel.Tags = core.StringPtr("testString")
				updatePropertyValuesOptionsModel.Value = core.StringPtr("testString")
				updatePropertyValuesOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				updatePropertyValuesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.UpdatePropertyValues(updatePropertyValuesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.UpdatePropertyValuesWithContext(ctx, updatePropertyValuesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr = appConfigurationService.UpdatePropertyValues(updatePropertyValuesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.UpdatePropertyValuesWithContext(ctx, updatePropertyValuesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
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
				targetSegmentsModel.Segments = []string{"testString"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = core.StringPtr("testString")
				segmentRuleModel.Order = core.Int64Ptr(int64(38))

				// Construct an instance of the UpdatePropertyValuesOptions model
				updatePropertyValuesOptionsModel := new(appconfigurationv1.UpdatePropertyValuesOptions)
				updatePropertyValuesOptionsModel.EnvironmentID = core.StringPtr("testString")
				updatePropertyValuesOptionsModel.PropertyID = core.StringPtr("testString")
				updatePropertyValuesOptionsModel.Name = core.StringPtr("testString")
				updatePropertyValuesOptionsModel.Description = core.StringPtr("testString")
				updatePropertyValuesOptionsModel.Tags = core.StringPtr("testString")
				updatePropertyValuesOptionsModel.Value = core.StringPtr("testString")
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
	})
	Describe(`GetProperty(getPropertyOptions *GetPropertyOptions) - Operation response error`, func() {
		getPropertyPath := "/environments/testString/properties/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPropertyPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["include"]).To(Equal([]string{"collections"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				getPropertyOptionsModel.EnvironmentID = core.StringPtr("testString")
				getPropertyOptionsModel.PropertyID = core.StringPtr("testString")
				getPropertyOptionsModel.Include = core.StringPtr("collections")
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
		getPropertyPath := "/environments/testString/properties/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPropertyPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["include"]).To(Equal([]string{"collections"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "property_id": "PropertyID", "description": "Description", "type": "BOOLEAN", "value": "anyValue", "tags": "Tags", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": "anyValue", "order": 5}], "segment_exists": false, "collections": [{"collection_id": "CollectionID", "name": "Name"}], "created_time": "2019-01-01T12:00:00", "updated_time": "2019-01-01T12:00:00", "evaluation_time": "2019-01-01T12:00:00", "href": "Href"}`)
				}))
			})
			It(`Invoke GetProperty successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.GetProperty(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPropertyOptions model
				getPropertyOptionsModel := new(appconfigurationv1.GetPropertyOptions)
				getPropertyOptionsModel.EnvironmentID = core.StringPtr("testString")
				getPropertyOptionsModel.PropertyID = core.StringPtr("testString")
				getPropertyOptionsModel.Include = core.StringPtr("collections")
				getPropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.GetProperty(getPropertyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.GetPropertyWithContext(ctx, getPropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr = appConfigurationService.GetProperty(getPropertyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.GetPropertyWithContext(ctx, getPropertyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
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
				getPropertyOptionsModel.EnvironmentID = core.StringPtr("testString")
				getPropertyOptionsModel.PropertyID = core.StringPtr("testString")
				getPropertyOptionsModel.Include = core.StringPtr("collections")
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
	})

	Describe(`DeleteProperty(deletePropertyOptions *DeletePropertyOptions)`, func() {
		deletePropertyPath := "/environments/testString/properties/testString"
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
				appConfigurationService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := appConfigurationService.DeleteProperty(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeletePropertyOptions model
				deletePropertyOptionsModel := new(appconfigurationv1.DeletePropertyOptions)
				deletePropertyOptionsModel.EnvironmentID = core.StringPtr("testString")
				deletePropertyOptionsModel.PropertyID = core.StringPtr("testString")
				deletePropertyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appConfigurationService.DeleteProperty(deletePropertyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				appConfigurationService.DisableRetries()
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
				deletePropertyOptionsModel.EnvironmentID = core.StringPtr("testString")
				deletePropertyOptionsModel.PropertyID = core.StringPtr("testString")
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
	Describe(`ListSegments(listSegmentsOptions *ListSegmentsOptions) - Operation response error`, func() {
		listSegmentsPath := "/segments"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSegmentsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for expand query parameter

					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_time"}))

					Expect(req.URL.Query()["tags"]).To(Equal([]string{"version 1.1, pre-release"}))

					Expect(req.URL.Query()["include"]).To(Equal([]string{"rules"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				listSegmentsOptionsModel.Tags = core.StringPtr("version 1.1, pre-release")
				listSegmentsOptionsModel.Include = core.StringPtr("rules")
				listSegmentsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listSegmentsOptionsModel.Offset = core.Int64Ptr(int64(0))
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
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSegmentsPath))
					Expect(req.Method).To(Equal("GET"))


					// TODO: Add check for expand query parameter

					Expect(req.URL.Query()["sort"]).To(Equal([]string{"created_time"}))

					Expect(req.URL.Query()["tags"]).To(Equal([]string{"version 1.1, pre-release"}))

					Expect(req.URL.Query()["include"]).To(Equal([]string{"rules"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"segments": [{"name": "Name", "segment_id": "SegmentID", "description": "Description", "tags": "Tags", "rules": [{"attribute_name": "AttributeName", "operator": "is", "values": ["Values"]}], "created_time": "2019-01-01T12:00:00", "updated_time": "2019-01-01T12:00:00", "href": "Href", "features": [{"feature_id": "FeatureID", "name": "Name"}], "properties": [{"property_id": "PropertyID", "name": "Name"}]}], "limit": 5, "offset": 6, "total_count": 10, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}}`)
				}))
			})
			It(`Invoke ListSegments successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.ListSegments(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListSegmentsOptions model
				listSegmentsOptionsModel := new(appconfigurationv1.ListSegmentsOptions)
				listSegmentsOptionsModel.Expand = core.BoolPtr(true)
				listSegmentsOptionsModel.Sort = core.StringPtr("created_time")
				listSegmentsOptionsModel.Tags = core.StringPtr("version 1.1, pre-release")
				listSegmentsOptionsModel.Include = core.StringPtr("rules")
				listSegmentsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listSegmentsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listSegmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.ListSegments(listSegmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.ListSegmentsWithContext(ctx, listSegmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr = appConfigurationService.ListSegments(listSegmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.ListSegmentsWithContext(ctx, listSegmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
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
				listSegmentsOptionsModel.Tags = core.StringPtr("version 1.1, pre-release")
				listSegmentsOptionsModel.Include = core.StringPtr("rules")
				listSegmentsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listSegmentsOptionsModel.Offset = core.Int64Ptr(int64(0))
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
	})
	Describe(`CreateSegment(createSegmentOptions *CreateSegmentOptions) - Operation response error`, func() {
		createSegmentPath := "/segments"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSegmentPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				ruleModel.Values = []string{"testString"}

				// Construct an instance of the CreateSegmentOptions model
				createSegmentOptionsModel := new(appconfigurationv1.CreateSegmentOptions)
				createSegmentOptionsModel.Name = core.StringPtr("Beta Users")
				createSegmentOptionsModel.SegmentID = core.StringPtr("beta-users")
				createSegmentOptionsModel.Description = core.StringPtr("Segment containing the beta users")
				createSegmentOptionsModel.Tags = core.StringPtr("version: 1.1, stage")
				createSegmentOptionsModel.Rules = []appconfigurationv1.Rule{*ruleModel}
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
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Name", "segment_id": "SegmentID", "description": "Description", "tags": "Tags", "rules": [{"attribute_name": "AttributeName", "operator": "is", "values": ["Values"]}], "created_time": "2019-01-01T12:00:00", "updated_time": "2019-01-01T12:00:00", "href": "Href", "features": [{"feature_id": "FeatureID", "name": "Name"}], "properties": [{"property_id": "PropertyID", "name": "Name"}]}`)
				}))
			})
			It(`Invoke CreateSegment successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.CreateSegment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Rule model
				ruleModel := new(appconfigurationv1.Rule)
				ruleModel.AttributeName = core.StringPtr("email")
				ruleModel.Operator = core.StringPtr("endsWith")
				ruleModel.Values = []string{"testString"}

				// Construct an instance of the CreateSegmentOptions model
				createSegmentOptionsModel := new(appconfigurationv1.CreateSegmentOptions)
				createSegmentOptionsModel.Name = core.StringPtr("Beta Users")
				createSegmentOptionsModel.SegmentID = core.StringPtr("beta-users")
				createSegmentOptionsModel.Description = core.StringPtr("Segment containing the beta users")
				createSegmentOptionsModel.Tags = core.StringPtr("version: 1.1, stage")
				createSegmentOptionsModel.Rules = []appconfigurationv1.Rule{*ruleModel}
				createSegmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.CreateSegment(createSegmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.CreateSegmentWithContext(ctx, createSegmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr = appConfigurationService.CreateSegment(createSegmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.CreateSegmentWithContext(ctx, createSegmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateSegment with error: Operation request error`, func() {
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
				ruleModel.Values = []string{"testString"}

				// Construct an instance of the CreateSegmentOptions model
				createSegmentOptionsModel := new(appconfigurationv1.CreateSegmentOptions)
				createSegmentOptionsModel.Name = core.StringPtr("Beta Users")
				createSegmentOptionsModel.SegmentID = core.StringPtr("beta-users")
				createSegmentOptionsModel.Description = core.StringPtr("Segment containing the beta users")
				createSegmentOptionsModel.Tags = core.StringPtr("version: 1.1, stage")
				createSegmentOptionsModel.Rules = []appconfigurationv1.Rule{*ruleModel}
				createSegmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.CreateSegment(createSegmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateSegment(updateSegmentOptions *UpdateSegmentOptions) - Operation response error`, func() {
		updateSegmentPath := "/segments/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSegmentPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				updateSegmentOptionsModel.SegmentID = core.StringPtr("testString")
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
		updateSegmentPath := "/segments/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "segment_id": "SegmentID", "description": "Description", "tags": "Tags", "rules": [{"attribute_name": "AttributeName", "operator": "is", "values": ["Values"]}], "created_time": "2019-01-01T12:00:00", "updated_time": "2019-01-01T12:00:00", "href": "Href", "features": [{"feature_id": "FeatureID", "name": "Name"}], "properties": [{"property_id": "PropertyID", "name": "Name"}]}`)
				}))
			})
			It(`Invoke UpdateSegment successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

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
				updateSegmentOptionsModel.SegmentID = core.StringPtr("testString")
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

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.UpdateSegmentWithContext(ctx, updateSegmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr = appConfigurationService.UpdateSegment(updateSegmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.UpdateSegmentWithContext(ctx, updateSegmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
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
				updateSegmentOptionsModel.SegmentID = core.StringPtr("testString")
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
	})
	Describe(`GetSegment(getSegmentOptions *GetSegmentOptions) - Operation response error`, func() {
		getSegmentPath := "/segments/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSegmentPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				getSegmentOptionsModel.SegmentID = core.StringPtr("testString")
				getSegmentOptionsModel.Include = []string{"features"}
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
		getSegmentPath := "/segments/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSegmentPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "segment_id": "SegmentID", "description": "Description", "tags": "Tags", "rules": [{"attribute_name": "AttributeName", "operator": "is", "values": ["Values"]}], "created_time": "2019-01-01T12:00:00", "updated_time": "2019-01-01T12:00:00", "href": "Href", "features": [{"feature_id": "FeatureID", "name": "Name"}], "properties": [{"property_id": "PropertyID", "name": "Name"}]}`)
				}))
			})
			It(`Invoke GetSegment successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.GetSegment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSegmentOptions model
				getSegmentOptionsModel := new(appconfigurationv1.GetSegmentOptions)
				getSegmentOptionsModel.SegmentID = core.StringPtr("testString")
				getSegmentOptionsModel.Include = []string{"features"}
				getSegmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.GetSegment(getSegmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.GetSegmentWithContext(ctx, getSegmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr = appConfigurationService.GetSegment(getSegmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.GetSegmentWithContext(ctx, getSegmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
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
				getSegmentOptionsModel.SegmentID = core.StringPtr("testString")
				getSegmentOptionsModel.Include = []string{"features"}
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
	})

	Describe(`DeleteSegment(deleteSegmentOptions *DeleteSegmentOptions)`, func() {
		deleteSegmentPath := "/segments/testString"
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
				appConfigurationService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := appConfigurationService.DeleteSegment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteSegmentOptions model
				deleteSegmentOptionsModel := new(appconfigurationv1.DeleteSegmentOptions)
				deleteSegmentOptionsModel.SegmentID = core.StringPtr("testString")
				deleteSegmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appConfigurationService.DeleteSegment(deleteSegmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				appConfigurationService.DisableRetries()
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
				deleteSegmentOptionsModel.SegmentID = core.StringPtr("testString")
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
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			appConfigurationService, _ := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
				URL:           "http://appconfigurationv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCollection successfully`, func() {
				name := "testString"
				collectionID := "testString"
				model, err := appConfigurationService.NewCollection(name, collectionID)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCollectionRef successfully`, func() {
				collectionID := "testString"
				model, err := appConfigurationService.NewCollectionRef(collectionID)
				Expect(model).ToNot(BeNil())
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
				targetSegmentsModel.Segments = []string{"testString"}
				Expect(targetSegmentsModel.Segments).To(Equal([]string{"testString"}))

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				Expect(segmentRuleModel).ToNot(BeNil())
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = core.StringPtr("true")
				segmentRuleModel.Order = core.Int64Ptr(int64(1))
				Expect(segmentRuleModel.Rules).To(Equal([]appconfigurationv1.TargetSegments{*targetSegmentsModel}))
				Expect(segmentRuleModel.Value).To(Equal(core.StringPtr("true")))
				Expect(segmentRuleModel.Order).To(Equal(core.Int64Ptr(int64(1))))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				Expect(collectionRefModel).ToNot(BeNil())
				collectionRefModel.CollectionID = core.StringPtr("ghzinc")
				Expect(collectionRefModel.CollectionID).To(Equal(core.StringPtr("ghzinc")))

				// Construct an instance of the CreateFeatureOptions model
				environmentID := "testString"
				createFeatureOptionsName := "Cycle Rentals"
				createFeatureOptionsFeatureID := "cycle-rentals"
				createFeatureOptionsType := "BOOLEAN"
				createFeatureOptionsEnabledValue := core.StringPtr("true")
				createFeatureOptionsDisabledValue := core.StringPtr("false")
				createFeatureOptionsModel := appConfigurationService.NewCreateFeatureOptions(environmentID, createFeatureOptionsName, createFeatureOptionsFeatureID, createFeatureOptionsType, createFeatureOptionsEnabledValue, createFeatureOptionsDisabledValue)
				createFeatureOptionsModel.SetEnvironmentID("testString")
				createFeatureOptionsModel.SetName("Cycle Rentals")
				createFeatureOptionsModel.SetFeatureID("cycle-rentals")
				createFeatureOptionsModel.SetType("BOOLEAN")
				createFeatureOptionsModel.SetEnabledValue(core.StringPtr("true"))
				createFeatureOptionsModel.SetDisabledValue(core.StringPtr("false"))
				createFeatureOptionsModel.SetDescription("Feature flag to enable Cycle Rentals")
				createFeatureOptionsModel.SetEnabled(true)
				createFeatureOptionsModel.SetTags("version: 1.1, pre-release")
				createFeatureOptionsModel.SetSegmentRules([]appconfigurationv1.SegmentRule{*segmentRuleModel})
				createFeatureOptionsModel.SetCollections([]appconfigurationv1.CollectionRef{*collectionRefModel})
				createFeatureOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createFeatureOptionsModel).ToNot(BeNil())
				Expect(createFeatureOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(createFeatureOptionsModel.Name).To(Equal(core.StringPtr("Cycle Rentals")))
				Expect(createFeatureOptionsModel.FeatureID).To(Equal(core.StringPtr("cycle-rentals")))
				Expect(createFeatureOptionsModel.Type).To(Equal(core.StringPtr("BOOLEAN")))
				Expect(createFeatureOptionsModel.EnabledValue).To(Equal(core.StringPtr("true")))
				Expect(createFeatureOptionsModel.DisabledValue).To(Equal(core.StringPtr("false")))
				Expect(createFeatureOptionsModel.Description).To(Equal(core.StringPtr("Feature flag to enable Cycle Rentals")))
				Expect(createFeatureOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(createFeatureOptionsModel.Tags).To(Equal(core.StringPtr("version: 1.1, pre-release")))
				Expect(createFeatureOptionsModel.SegmentRules).To(Equal([]appconfigurationv1.SegmentRule{*segmentRuleModel}))
				Expect(createFeatureOptionsModel.Collections).To(Equal([]appconfigurationv1.CollectionRef{*collectionRefModel}))
				Expect(createFeatureOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreatePropertyOptions successfully`, func() {
				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				Expect(targetSegmentsModel).ToNot(BeNil())
				targetSegmentsModel.Segments = []string{"testString"}
				Expect(targetSegmentsModel.Segments).To(Equal([]string{"testString"}))

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				Expect(segmentRuleModel).ToNot(BeNil())
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = core.StringPtr("true")
				segmentRuleModel.Order = core.Int64Ptr(int64(1))
				Expect(segmentRuleModel.Rules).To(Equal([]appconfigurationv1.TargetSegments{*targetSegmentsModel}))
				Expect(segmentRuleModel.Value).To(Equal(core.StringPtr("true")))
				Expect(segmentRuleModel.Order).To(Equal(core.Int64Ptr(int64(1))))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				Expect(collectionRefModel).ToNot(BeNil())
				collectionRefModel.CollectionID = core.StringPtr("ghzinc")
				Expect(collectionRefModel.CollectionID).To(Equal(core.StringPtr("ghzinc")))

				// Construct an instance of the CreatePropertyOptions model
				environmentID := "testString"
				createPropertyOptionsName := "Email property"
				createPropertyOptionsPropertyID := "email-property"
				createPropertyOptionsType := "BOOLEAN"
				createPropertyOptionsValue := core.StringPtr("true")
				createPropertyOptionsModel := appConfigurationService.NewCreatePropertyOptions(environmentID, createPropertyOptionsName, createPropertyOptionsPropertyID, createPropertyOptionsType, createPropertyOptionsValue)
				createPropertyOptionsModel.SetEnvironmentID("testString")
				createPropertyOptionsModel.SetName("Email property")
				createPropertyOptionsModel.SetPropertyID("email-property")
				createPropertyOptionsModel.SetType("BOOLEAN")
				createPropertyOptionsModel.SetValue(core.StringPtr("true"))
				createPropertyOptionsModel.SetDescription("Property for email")
				createPropertyOptionsModel.SetTags("version: 1.1, pre-release")
				createPropertyOptionsModel.SetSegmentRules([]appconfigurationv1.SegmentRule{*segmentRuleModel})
				createPropertyOptionsModel.SetCollections([]appconfigurationv1.CollectionRef{*collectionRefModel})
				createPropertyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createPropertyOptionsModel).ToNot(BeNil())
				Expect(createPropertyOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(createPropertyOptionsModel.Name).To(Equal(core.StringPtr("Email property")))
				Expect(createPropertyOptionsModel.PropertyID).To(Equal(core.StringPtr("email-property")))
				Expect(createPropertyOptionsModel.Type).To(Equal(core.StringPtr("BOOLEAN")))
				Expect(createPropertyOptionsModel.Value).To(Equal(core.StringPtr("true")))
				Expect(createPropertyOptionsModel.Description).To(Equal(core.StringPtr("Property for email")))
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
				ruleModel.Values = []string{"testString"}
				Expect(ruleModel.AttributeName).To(Equal(core.StringPtr("email")))
				Expect(ruleModel.Operator).To(Equal(core.StringPtr("endsWith")))
				Expect(ruleModel.Values).To(Equal([]string{"testString"}))

				// Construct an instance of the CreateSegmentOptions model
				createSegmentOptionsModel := appConfigurationService.NewCreateSegmentOptions()
				createSegmentOptionsModel.SetName("Beta Users")
				createSegmentOptionsModel.SetSegmentID("beta-users")
				createSegmentOptionsModel.SetDescription("Segment containing the beta users")
				createSegmentOptionsModel.SetTags("version: 1.1, stage")
				createSegmentOptionsModel.SetRules([]appconfigurationv1.Rule{*ruleModel})
				createSegmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createSegmentOptionsModel).ToNot(BeNil())
				Expect(createSegmentOptionsModel.Name).To(Equal(core.StringPtr("Beta Users")))
				Expect(createSegmentOptionsModel.SegmentID).To(Equal(core.StringPtr("beta-users")))
				Expect(createSegmentOptionsModel.Description).To(Equal(core.StringPtr("Segment containing the beta users")))
				Expect(createSegmentOptionsModel.Tags).To(Equal(core.StringPtr("version: 1.1, stage")))
				Expect(createSegmentOptionsModel.Rules).To(Equal([]appconfigurationv1.Rule{*ruleModel}))
				Expect(createSegmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteCollectionOptions successfully`, func() {
				// Construct an instance of the DeleteCollectionOptions model
				collectionID := "testString"
				deleteCollectionOptionsModel := appConfigurationService.NewDeleteCollectionOptions(collectionID)
				deleteCollectionOptionsModel.SetCollectionID("testString")
				deleteCollectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteCollectionOptionsModel).ToNot(BeNil())
				Expect(deleteCollectionOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCollectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteEnvironmentOptions successfully`, func() {
				// Construct an instance of the DeleteEnvironmentOptions model
				environmentID := "testString"
				deleteEnvironmentOptionsModel := appConfigurationService.NewDeleteEnvironmentOptions(environmentID)
				deleteEnvironmentOptionsModel.SetEnvironmentID("testString")
				deleteEnvironmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteEnvironmentOptionsModel).ToNot(BeNil())
				Expect(deleteEnvironmentOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteEnvironmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteFeatureOptions successfully`, func() {
				// Construct an instance of the DeleteFeatureOptions model
				environmentID := "testString"
				featureID := "testString"
				deleteFeatureOptionsModel := appConfigurationService.NewDeleteFeatureOptions(environmentID, featureID)
				deleteFeatureOptionsModel.SetEnvironmentID("testString")
				deleteFeatureOptionsModel.SetFeatureID("testString")
				deleteFeatureOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteFeatureOptionsModel).ToNot(BeNil())
				Expect(deleteFeatureOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteFeatureOptionsModel.FeatureID).To(Equal(core.StringPtr("testString")))
				Expect(deleteFeatureOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeletePropertyOptions successfully`, func() {
				// Construct an instance of the DeletePropertyOptions model
				environmentID := "testString"
				propertyID := "testString"
				deletePropertyOptionsModel := appConfigurationService.NewDeletePropertyOptions(environmentID, propertyID)
				deletePropertyOptionsModel.SetEnvironmentID("testString")
				deletePropertyOptionsModel.SetPropertyID("testString")
				deletePropertyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deletePropertyOptionsModel).ToNot(BeNil())
				Expect(deletePropertyOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(deletePropertyOptionsModel.PropertyID).To(Equal(core.StringPtr("testString")))
				Expect(deletePropertyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteSegmentOptions successfully`, func() {
				// Construct an instance of the DeleteSegmentOptions model
				segmentID := "testString"
				deleteSegmentOptionsModel := appConfigurationService.NewDeleteSegmentOptions(segmentID)
				deleteSegmentOptionsModel.SetSegmentID("testString")
				deleteSegmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteSegmentOptionsModel).ToNot(BeNil())
				Expect(deleteSegmentOptionsModel.SegmentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteSegmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewEnvironment successfully`, func() {
				name := "testString"
				environmentID := "testString"
				model, err := appConfigurationService.NewEnvironment(name, environmentID)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewFeature successfully`, func() {
				name := "testString"
				featureID := "testString"
				typeVar := "BOOLEAN"
				enabledValue := core.StringPtr("testString")
				disabledValue := core.StringPtr("testString")
				model, err := appConfigurationService.NewFeature(name, featureID, typeVar, enabledValue, disabledValue)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewFeatureOutput successfully`, func() {
				featureID := "testString"
				name := "testString"
				model, err := appConfigurationService.NewFeatureOutput(featureID, name)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGetCollectionOptions successfully`, func() {
				// Construct an instance of the GetCollectionOptions model
				collectionID := "testString"
				getCollectionOptionsModel := appConfigurationService.NewGetCollectionOptions(collectionID)
				getCollectionOptionsModel.SetCollectionID("testString")
				getCollectionOptionsModel.SetExpand(true)
				getCollectionOptionsModel.SetInclude([]string{"features"})
				getCollectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCollectionOptionsModel).ToNot(BeNil())
				Expect(getCollectionOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(getCollectionOptionsModel.Expand).To(Equal(core.BoolPtr(true)))
				Expect(getCollectionOptionsModel.Include).To(Equal([]string{"features"}))
				Expect(getCollectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetEnvironmentOptions successfully`, func() {
				// Construct an instance of the GetEnvironmentOptions model
				environmentID := "testString"
				getEnvironmentOptionsModel := appConfigurationService.NewGetEnvironmentOptions(environmentID)
				getEnvironmentOptionsModel.SetEnvironmentID("testString")
				getEnvironmentOptionsModel.SetExpand(true)
				getEnvironmentOptionsModel.SetInclude([]string{"features"})
				getEnvironmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getEnvironmentOptionsModel).ToNot(BeNil())
				Expect(getEnvironmentOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(getEnvironmentOptionsModel.Expand).To(Equal(core.BoolPtr(true)))
				Expect(getEnvironmentOptionsModel.Include).To(Equal([]string{"features"}))
				Expect(getEnvironmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetFeatureOptions successfully`, func() {
				// Construct an instance of the GetFeatureOptions model
				environmentID := "testString"
				featureID := "testString"
				getFeatureOptionsModel := appConfigurationService.NewGetFeatureOptions(environmentID, featureID)
				getFeatureOptionsModel.SetEnvironmentID("testString")
				getFeatureOptionsModel.SetFeatureID("testString")
				getFeatureOptionsModel.SetInclude("collections")
				getFeatureOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getFeatureOptionsModel).ToNot(BeNil())
				Expect(getFeatureOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(getFeatureOptionsModel.FeatureID).To(Equal(core.StringPtr("testString")))
				Expect(getFeatureOptionsModel.Include).To(Equal(core.StringPtr("collections")))
				Expect(getFeatureOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPropertyOptions successfully`, func() {
				// Construct an instance of the GetPropertyOptions model
				environmentID := "testString"
				propertyID := "testString"
				getPropertyOptionsModel := appConfigurationService.NewGetPropertyOptions(environmentID, propertyID)
				getPropertyOptionsModel.SetEnvironmentID("testString")
				getPropertyOptionsModel.SetPropertyID("testString")
				getPropertyOptionsModel.SetInclude("collections")
				getPropertyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPropertyOptionsModel).ToNot(BeNil())
				Expect(getPropertyOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(getPropertyOptionsModel.PropertyID).To(Equal(core.StringPtr("testString")))
				Expect(getPropertyOptionsModel.Include).To(Equal(core.StringPtr("collections")))
				Expect(getPropertyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSegmentOptions successfully`, func() {
				// Construct an instance of the GetSegmentOptions model
				segmentID := "testString"
				getSegmentOptionsModel := appConfigurationService.NewGetSegmentOptions(segmentID)
				getSegmentOptionsModel.SetSegmentID("testString")
				getSegmentOptionsModel.SetInclude([]string{"features"})
				getSegmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSegmentOptionsModel).ToNot(BeNil())
				Expect(getSegmentOptionsModel.SegmentID).To(Equal(core.StringPtr("testString")))
				Expect(getSegmentOptionsModel.Include).To(Equal([]string{"features"}))
				Expect(getSegmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListCollectionsOptions successfully`, func() {
				// Construct an instance of the ListCollectionsOptions model
				listCollectionsOptionsModel := appConfigurationService.NewListCollectionsOptions()
				listCollectionsOptionsModel.SetExpand(true)
				listCollectionsOptionsModel.SetSort("created_time")
				listCollectionsOptionsModel.SetTags("version 1.1, pre-release")
				listCollectionsOptionsModel.SetFeatures([]string{"testString"})
				listCollectionsOptionsModel.SetProperties([]string{"testString"})
				listCollectionsOptionsModel.SetInclude([]string{"features"})
				listCollectionsOptionsModel.SetLimit(int64(1))
				listCollectionsOptionsModel.SetOffset(int64(0))
				listCollectionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listCollectionsOptionsModel).ToNot(BeNil())
				Expect(listCollectionsOptionsModel.Expand).To(Equal(core.BoolPtr(true)))
				Expect(listCollectionsOptionsModel.Sort).To(Equal(core.StringPtr("created_time")))
				Expect(listCollectionsOptionsModel.Tags).To(Equal(core.StringPtr("version 1.1, pre-release")))
				Expect(listCollectionsOptionsModel.Features).To(Equal([]string{"testString"}))
				Expect(listCollectionsOptionsModel.Properties).To(Equal([]string{"testString"}))
				Expect(listCollectionsOptionsModel.Include).To(Equal([]string{"features"}))
				Expect(listCollectionsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listCollectionsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listCollectionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListEnvironmentsOptions successfully`, func() {
				// Construct an instance of the ListEnvironmentsOptions model
				listEnvironmentsOptionsModel := appConfigurationService.NewListEnvironmentsOptions()
				listEnvironmentsOptionsModel.SetExpand(true)
				listEnvironmentsOptionsModel.SetSort("created_time")
				listEnvironmentsOptionsModel.SetTags("version 1.1, pre-release")
				listEnvironmentsOptionsModel.SetInclude([]string{"features"})
				listEnvironmentsOptionsModel.SetLimit(int64(1))
				listEnvironmentsOptionsModel.SetOffset(int64(0))
				listEnvironmentsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listEnvironmentsOptionsModel).ToNot(BeNil())
				Expect(listEnvironmentsOptionsModel.Expand).To(Equal(core.BoolPtr(true)))
				Expect(listEnvironmentsOptionsModel.Sort).To(Equal(core.StringPtr("created_time")))
				Expect(listEnvironmentsOptionsModel.Tags).To(Equal(core.StringPtr("version 1.1, pre-release")))
				Expect(listEnvironmentsOptionsModel.Include).To(Equal([]string{"features"}))
				Expect(listEnvironmentsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listEnvironmentsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listEnvironmentsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListFeaturesOptions successfully`, func() {
				// Construct an instance of the ListFeaturesOptions model
				environmentID := "testString"
				listFeaturesOptionsModel := appConfigurationService.NewListFeaturesOptions(environmentID)
				listFeaturesOptionsModel.SetEnvironmentID("testString")
				listFeaturesOptionsModel.SetExpand(true)
				listFeaturesOptionsModel.SetSort("created_time")
				listFeaturesOptionsModel.SetTags("version 1.1, pre-release")
				listFeaturesOptionsModel.SetCollections([]string{"testString"})
				listFeaturesOptionsModel.SetSegments([]string{"testString"})
				listFeaturesOptionsModel.SetInclude([]string{"collections"})
				listFeaturesOptionsModel.SetLimit(int64(1))
				listFeaturesOptionsModel.SetOffset(int64(0))
				listFeaturesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listFeaturesOptionsModel).ToNot(BeNil())
				Expect(listFeaturesOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(listFeaturesOptionsModel.Expand).To(Equal(core.BoolPtr(true)))
				Expect(listFeaturesOptionsModel.Sort).To(Equal(core.StringPtr("created_time")))
				Expect(listFeaturesOptionsModel.Tags).To(Equal(core.StringPtr("version 1.1, pre-release")))
				Expect(listFeaturesOptionsModel.Collections).To(Equal([]string{"testString"}))
				Expect(listFeaturesOptionsModel.Segments).To(Equal([]string{"testString"}))
				Expect(listFeaturesOptionsModel.Include).To(Equal([]string{"collections"}))
				Expect(listFeaturesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listFeaturesOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listFeaturesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListPropertiesOptions successfully`, func() {
				// Construct an instance of the ListPropertiesOptions model
				environmentID := "testString"
				listPropertiesOptionsModel := appConfigurationService.NewListPropertiesOptions(environmentID)
				listPropertiesOptionsModel.SetEnvironmentID("testString")
				listPropertiesOptionsModel.SetExpand(true)
				listPropertiesOptionsModel.SetSort("created_time")
				listPropertiesOptionsModel.SetTags("version 1.1, pre-release")
				listPropertiesOptionsModel.SetCollections([]string{"testString"})
				listPropertiesOptionsModel.SetSegments([]string{"testString"})
				listPropertiesOptionsModel.SetInclude([]string{"collections"})
				listPropertiesOptionsModel.SetLimit(int64(1))
				listPropertiesOptionsModel.SetOffset(int64(0))
				listPropertiesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listPropertiesOptionsModel).ToNot(BeNil())
				Expect(listPropertiesOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(listPropertiesOptionsModel.Expand).To(Equal(core.BoolPtr(true)))
				Expect(listPropertiesOptionsModel.Sort).To(Equal(core.StringPtr("created_time")))
				Expect(listPropertiesOptionsModel.Tags).To(Equal(core.StringPtr("version 1.1, pre-release")))
				Expect(listPropertiesOptionsModel.Collections).To(Equal([]string{"testString"}))
				Expect(listPropertiesOptionsModel.Segments).To(Equal([]string{"testString"}))
				Expect(listPropertiesOptionsModel.Include).To(Equal([]string{"collections"}))
				Expect(listPropertiesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listPropertiesOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listPropertiesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListSegmentsOptions successfully`, func() {
				// Construct an instance of the ListSegmentsOptions model
				listSegmentsOptionsModel := appConfigurationService.NewListSegmentsOptions()
				listSegmentsOptionsModel.SetExpand(true)
				listSegmentsOptionsModel.SetSort("created_time")
				listSegmentsOptionsModel.SetTags("version 1.1, pre-release")
				listSegmentsOptionsModel.SetInclude("rules")
				listSegmentsOptionsModel.SetLimit(int64(1))
				listSegmentsOptionsModel.SetOffset(int64(0))
				listSegmentsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listSegmentsOptionsModel).ToNot(BeNil())
				Expect(listSegmentsOptionsModel.Expand).To(Equal(core.BoolPtr(true)))
				Expect(listSegmentsOptionsModel.Sort).To(Equal(core.StringPtr("created_time")))
				Expect(listSegmentsOptionsModel.Tags).To(Equal(core.StringPtr("version 1.1, pre-release")))
				Expect(listSegmentsOptionsModel.Include).To(Equal(core.StringPtr("rules")))
				Expect(listSegmentsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listSegmentsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listSegmentsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewProperty successfully`, func() {
				name := "testString"
				propertyID := "testString"
				typeVar := "BOOLEAN"
				value := core.StringPtr("testString")
				model, err := appConfigurationService.NewProperty(name, propertyID, typeVar, value)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPropertyOutput successfully`, func() {
				propertyID := "testString"
				name := "testString"
				model, err := appConfigurationService.NewPropertyOutput(propertyID, name)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRule successfully`, func() {
				attributeName := "testString"
				operator := "is"
				values := []string{"testString"}
				model, err := appConfigurationService.NewRule(attributeName, operator, values)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSegment successfully`, func() {
				name := "testString"
				segmentID := "testString"
				rules := []appconfigurationv1.Rule{}
				model, err := appConfigurationService.NewSegment(name, segmentID, rules)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSegmentRule successfully`, func() {
				rules := []appconfigurationv1.TargetSegments{}
				value := core.StringPtr("testString")
				order := int64(38)
				model, err := appConfigurationService.NewSegmentRule(rules, value, order)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewTargetSegments successfully`, func() {
				segments := []string{"testString"}
				model, err := appConfigurationService.NewTargetSegments(segments)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewToggleFeatureOptions successfully`, func() {
				// Construct an instance of the ToggleFeatureOptions model
				environmentID := "testString"
				featureID := "testString"
				toggleFeatureOptionsModel := appConfigurationService.NewToggleFeatureOptions(environmentID, featureID)
				toggleFeatureOptionsModel.SetEnvironmentID("testString")
				toggleFeatureOptionsModel.SetFeatureID("testString")
				toggleFeatureOptionsModel.SetEnabled(true)
				toggleFeatureOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(toggleFeatureOptionsModel).ToNot(BeNil())
				Expect(toggleFeatureOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(toggleFeatureOptionsModel.FeatureID).To(Equal(core.StringPtr("testString")))
				Expect(toggleFeatureOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(toggleFeatureOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateCollectionOptions successfully`, func() {
				// Construct an instance of the UpdateCollectionOptions model
				collectionID := "testString"
				updateCollectionOptionsModel := appConfigurationService.NewUpdateCollectionOptions(collectionID)
				updateCollectionOptionsModel.SetCollectionID("testString")
				updateCollectionOptionsModel.SetName("testString")
				updateCollectionOptionsModel.SetDescription("testString")
				updateCollectionOptionsModel.SetTags("testString")
				updateCollectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateCollectionOptionsModel).ToNot(BeNil())
				Expect(updateCollectionOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(updateCollectionOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateCollectionOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateCollectionOptionsModel.Tags).To(Equal(core.StringPtr("testString")))
				Expect(updateCollectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateEnvironmentOptions successfully`, func() {
				// Construct an instance of the UpdateEnvironmentOptions model
				environmentID := "testString"
				updateEnvironmentOptionsModel := appConfigurationService.NewUpdateEnvironmentOptions(environmentID)
				updateEnvironmentOptionsModel.SetEnvironmentID("testString")
				updateEnvironmentOptionsModel.SetName("testString")
				updateEnvironmentOptionsModel.SetDescription("testString")
				updateEnvironmentOptionsModel.SetTags("testString")
				updateEnvironmentOptionsModel.SetColorCode("#FDD13A")
				updateEnvironmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateEnvironmentOptionsModel).ToNot(BeNil())
				Expect(updateEnvironmentOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
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
				targetSegmentsModel.Segments = []string{"testString"}
				Expect(targetSegmentsModel.Segments).To(Equal([]string{"testString"}))

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				Expect(segmentRuleModel).ToNot(BeNil())
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = core.StringPtr("testString")
				segmentRuleModel.Order = core.Int64Ptr(int64(38))
				Expect(segmentRuleModel.Rules).To(Equal([]appconfigurationv1.TargetSegments{*targetSegmentsModel}))
				Expect(segmentRuleModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(segmentRuleModel.Order).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				Expect(collectionRefModel).ToNot(BeNil())
				collectionRefModel.CollectionID = core.StringPtr("testString")
				Expect(collectionRefModel.CollectionID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateFeatureOptions model
				environmentID := "testString"
				featureID := "testString"
				updateFeatureOptionsModel := appConfigurationService.NewUpdateFeatureOptions(environmentID, featureID)
				updateFeatureOptionsModel.SetEnvironmentID("testString")
				updateFeatureOptionsModel.SetFeatureID("testString")
				updateFeatureOptionsModel.SetName("testString")
				updateFeatureOptionsModel.SetDescription("testString")
				updateFeatureOptionsModel.SetEnabledValue(core.StringPtr("testString"))
				updateFeatureOptionsModel.SetDisabledValue(core.StringPtr("testString"))
				updateFeatureOptionsModel.SetEnabled(true)
				updateFeatureOptionsModel.SetTags("testString")
				updateFeatureOptionsModel.SetSegmentRules([]appconfigurationv1.SegmentRule{*segmentRuleModel})
				updateFeatureOptionsModel.SetCollections([]appconfigurationv1.CollectionRef{*collectionRefModel})
				updateFeatureOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateFeatureOptionsModel).ToNot(BeNil())
				Expect(updateFeatureOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(updateFeatureOptionsModel.FeatureID).To(Equal(core.StringPtr("testString")))
				Expect(updateFeatureOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateFeatureOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateFeatureOptionsModel.EnabledValue).To(Equal(core.StringPtr("testString")))
				Expect(updateFeatureOptionsModel.DisabledValue).To(Equal(core.StringPtr("testString")))
				Expect(updateFeatureOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(updateFeatureOptionsModel.Tags).To(Equal(core.StringPtr("testString")))
				Expect(updateFeatureOptionsModel.SegmentRules).To(Equal([]appconfigurationv1.SegmentRule{*segmentRuleModel}))
				Expect(updateFeatureOptionsModel.Collections).To(Equal([]appconfigurationv1.CollectionRef{*collectionRefModel}))
				Expect(updateFeatureOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateFeatureValuesOptions successfully`, func() {
				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				Expect(targetSegmentsModel).ToNot(BeNil())
				targetSegmentsModel.Segments = []string{"testString"}
				Expect(targetSegmentsModel.Segments).To(Equal([]string{"testString"}))

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				Expect(segmentRuleModel).ToNot(BeNil())
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = core.StringPtr("testString")
				segmentRuleModel.Order = core.Int64Ptr(int64(38))
				Expect(segmentRuleModel.Rules).To(Equal([]appconfigurationv1.TargetSegments{*targetSegmentsModel}))
				Expect(segmentRuleModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(segmentRuleModel.Order).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the UpdateFeatureValuesOptions model
				environmentID := "testString"
				featureID := "testString"
				updateFeatureValuesOptionsModel := appConfigurationService.NewUpdateFeatureValuesOptions(environmentID, featureID)
				updateFeatureValuesOptionsModel.SetEnvironmentID("testString")
				updateFeatureValuesOptionsModel.SetFeatureID("testString")
				updateFeatureValuesOptionsModel.SetName("testString")
				updateFeatureValuesOptionsModel.SetDescription("testString")
				updateFeatureValuesOptionsModel.SetTags("testString")
				updateFeatureValuesOptionsModel.SetEnabledValue(core.StringPtr("testString"))
				updateFeatureValuesOptionsModel.SetDisabledValue(core.StringPtr("testString"))
				updateFeatureValuesOptionsModel.SetSegmentRules([]appconfigurationv1.SegmentRule{*segmentRuleModel})
				updateFeatureValuesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateFeatureValuesOptionsModel).ToNot(BeNil())
				Expect(updateFeatureValuesOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(updateFeatureValuesOptionsModel.FeatureID).To(Equal(core.StringPtr("testString")))
				Expect(updateFeatureValuesOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateFeatureValuesOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateFeatureValuesOptionsModel.Tags).To(Equal(core.StringPtr("testString")))
				Expect(updateFeatureValuesOptionsModel.EnabledValue).To(Equal(core.StringPtr("testString")))
				Expect(updateFeatureValuesOptionsModel.DisabledValue).To(Equal(core.StringPtr("testString")))
				Expect(updateFeatureValuesOptionsModel.SegmentRules).To(Equal([]appconfigurationv1.SegmentRule{*segmentRuleModel}))
				Expect(updateFeatureValuesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdatePropertyOptions successfully`, func() {
				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				Expect(targetSegmentsModel).ToNot(BeNil())
				targetSegmentsModel.Segments = []string{"testString"}
				Expect(targetSegmentsModel.Segments).To(Equal([]string{"testString"}))

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				Expect(segmentRuleModel).ToNot(BeNil())
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = core.StringPtr("testString")
				segmentRuleModel.Order = core.Int64Ptr(int64(38))
				Expect(segmentRuleModel.Rules).To(Equal([]appconfigurationv1.TargetSegments{*targetSegmentsModel}))
				Expect(segmentRuleModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(segmentRuleModel.Order).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the CollectionRef model
				collectionRefModel := new(appconfigurationv1.CollectionRef)
				Expect(collectionRefModel).ToNot(BeNil())
				collectionRefModel.CollectionID = core.StringPtr("testString")
				Expect(collectionRefModel.CollectionID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdatePropertyOptions model
				environmentID := "testString"
				propertyID := "testString"
				updatePropertyOptionsModel := appConfigurationService.NewUpdatePropertyOptions(environmentID, propertyID)
				updatePropertyOptionsModel.SetEnvironmentID("testString")
				updatePropertyOptionsModel.SetPropertyID("testString")
				updatePropertyOptionsModel.SetName("testString")
				updatePropertyOptionsModel.SetDescription("testString")
				updatePropertyOptionsModel.SetValue(core.StringPtr("testString"))
				updatePropertyOptionsModel.SetTags("testString")
				updatePropertyOptionsModel.SetSegmentRules([]appconfigurationv1.SegmentRule{*segmentRuleModel})
				updatePropertyOptionsModel.SetCollections([]appconfigurationv1.CollectionRef{*collectionRefModel})
				updatePropertyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updatePropertyOptionsModel).ToNot(BeNil())
				Expect(updatePropertyOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(updatePropertyOptionsModel.PropertyID).To(Equal(core.StringPtr("testString")))
				Expect(updatePropertyOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updatePropertyOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updatePropertyOptionsModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(updatePropertyOptionsModel.Tags).To(Equal(core.StringPtr("testString")))
				Expect(updatePropertyOptionsModel.SegmentRules).To(Equal([]appconfigurationv1.SegmentRule{*segmentRuleModel}))
				Expect(updatePropertyOptionsModel.Collections).To(Equal([]appconfigurationv1.CollectionRef{*collectionRefModel}))
				Expect(updatePropertyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdatePropertyValuesOptions successfully`, func() {
				// Construct an instance of the TargetSegments model
				targetSegmentsModel := new(appconfigurationv1.TargetSegments)
				Expect(targetSegmentsModel).ToNot(BeNil())
				targetSegmentsModel.Segments = []string{"testString"}
				Expect(targetSegmentsModel.Segments).To(Equal([]string{"testString"}))

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				Expect(segmentRuleModel).ToNot(BeNil())
				segmentRuleModel.Rules = []appconfigurationv1.TargetSegments{*targetSegmentsModel}
				segmentRuleModel.Value = core.StringPtr("testString")
				segmentRuleModel.Order = core.Int64Ptr(int64(38))
				Expect(segmentRuleModel.Rules).To(Equal([]appconfigurationv1.TargetSegments{*targetSegmentsModel}))
				Expect(segmentRuleModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(segmentRuleModel.Order).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the UpdatePropertyValuesOptions model
				environmentID := "testString"
				propertyID := "testString"
				updatePropertyValuesOptionsModel := appConfigurationService.NewUpdatePropertyValuesOptions(environmentID, propertyID)
				updatePropertyValuesOptionsModel.SetEnvironmentID("testString")
				updatePropertyValuesOptionsModel.SetPropertyID("testString")
				updatePropertyValuesOptionsModel.SetName("testString")
				updatePropertyValuesOptionsModel.SetDescription("testString")
				updatePropertyValuesOptionsModel.SetTags("testString")
				updatePropertyValuesOptionsModel.SetValue(core.StringPtr("testString"))
				updatePropertyValuesOptionsModel.SetSegmentRules([]appconfigurationv1.SegmentRule{*segmentRuleModel})
				updatePropertyValuesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updatePropertyValuesOptionsModel).ToNot(BeNil())
				Expect(updatePropertyValuesOptionsModel.EnvironmentID).To(Equal(core.StringPtr("testString")))
				Expect(updatePropertyValuesOptionsModel.PropertyID).To(Equal(core.StringPtr("testString")))
				Expect(updatePropertyValuesOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updatePropertyValuesOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updatePropertyValuesOptionsModel.Tags).To(Equal(core.StringPtr("testString")))
				Expect(updatePropertyValuesOptionsModel.Value).To(Equal(core.StringPtr("testString")))
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
				segmentID := "testString"
				updateSegmentOptionsModel := appConfigurationService.NewUpdateSegmentOptions(segmentID)
				updateSegmentOptionsModel.SetSegmentID("testString")
				updateSegmentOptionsModel.SetName("testString")
				updateSegmentOptionsModel.SetDescription("testString")
				updateSegmentOptionsModel.SetTags("testString")
				updateSegmentOptionsModel.SetRules([]appconfigurationv1.Rule{*ruleModel})
				updateSegmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateSegmentOptionsModel).ToNot(BeNil())
				Expect(updateSegmentOptionsModel.SegmentID).To(Equal(core.StringPtr("testString")))
				Expect(updateSegmentOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateSegmentOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateSegmentOptionsModel.Tags).To(Equal(core.StringPtr("testString")))
				Expect(updateSegmentOptionsModel.Rules).To(Equal([]appconfigurationv1.Rule{*ruleModel}))
				Expect(updateSegmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
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
			mockDate := CreateMockDate()
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime()
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate() *strfmt.Date {
	d := strfmt.Date(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
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
