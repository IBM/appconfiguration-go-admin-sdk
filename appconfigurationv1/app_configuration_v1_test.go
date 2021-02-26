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
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"time"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var serviceName = strings.ToUpper(appconfigurationv1.GetDefaultServiceName())

var _ = Describe(`AppConfigurationV1`, func() {
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
				serviceName + "_URL":       "https://appconfigurationv1/api",
				serviceName + "_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1UsingExternalConfig(&appconfigurationv1.AppConfigurationV1Options{})
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
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1UsingExternalConfig(&appconfigurationv1.AppConfigurationV1Options{})
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
				serviceName + "_URL":       "https://appconfigurationv1/api",
				serviceName + "_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1UsingExternalConfig(&appconfigurationv1.AppConfigurationV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(appConfigurationService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				serviceName + "_AUTH_TYPE": "NOAuth",
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
				serviceName + "_URL":       "https://appconfigurationv1/api",
				serviceName + "_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1UsingExternalConfig(&appconfigurationv1.AppConfigurationV1Options{})
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
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1UsingExternalConfig(&appconfigurationv1.AppConfigurationV1Options{})
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
				serviceName + "_URL":       "https://appconfigurationv1/api",
				serviceName + "_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1UsingExternalConfig(&appconfigurationv1.AppConfigurationV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(appConfigurationService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				serviceName + "_AUTH_TYPE": "NOAuth",
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
				serviceName + "_URL":       "https://appconfigurationv1/api",
				serviceName + "_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1UsingExternalConfig(&appconfigurationv1.AppConfigurationV1Options{})
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
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1UsingExternalConfig(&appconfigurationv1.AppConfigurationV1Options{})
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
				serviceName + "_URL":       "https://appconfigurationv1/api",
				serviceName + "_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1UsingExternalConfig(&appconfigurationv1.AppConfigurationV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(appConfigurationService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				serviceName + "_AUTH_TYPE": "NOAuth",
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

	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			appConfigurationService, _ := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
				URL:           "http://appconfigurationv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCollection successfully`, func() {
				collectionID := "testString"
				enabled := true
				model, err := appConfigurationService.NewCollection(collectionID, enabled)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCollectionWithDeletedFlag successfully`, func() {
				collectionID := "testString"
				model, err := appConfigurationService.NewCollectionWithDeletedFlag(collectionID, true, true)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateCollectionOptions successfully`, func() {
				// Construct an instance of the CreateCollectionOptions model
				createCollectionOptionsName := "GHz Inc"
				createCollectionOptionsModel := appConfigurationService.NewCreateCollectionOptions(createCollectionOptionsName)
				createCollectionOptionsModel.SetName("GHz Inc")
				createCollectionOptionsModel.SetCollectionID("ghzinc")
				createCollectionOptionsModel.SetDescription("Collection for GHz Inc")
				createCollectionOptionsModel.SetTags("version: 1.1, pre-release")
				createCollectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCollectionOptionsModel).ToNot(BeNil())
				Expect(createCollectionOptionsModel.Name).To(Equal(core.StringPtr("GHz Inc")))
				Expect(createCollectionOptionsModel.CollectionID).To(Equal(core.StringPtr("ghzinc")))
				Expect(createCollectionOptionsModel.Description).To(Equal(core.StringPtr("Collection for GHz Inc")))
				Expect(createCollectionOptionsModel.Tags).To(Equal(core.StringPtr("version: 1.1, pre-release")))
				Expect(createCollectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateFeatureOptions successfully`, func() {
				// Construct an instance of the Rule model
				ruleModel := new(appconfigurationv1.Rule)
				Expect(ruleModel).ToNot(BeNil())
				ruleModel.Segments = []string{"testString"}
				Expect(ruleModel.Segments).To(Equal([]string{"testString"}))

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				Expect(segmentRuleModel).ToNot(BeNil())
				segmentRuleModel.Rules = []appconfigurationv1.Rule{*ruleModel}
				segmentRuleModel.Value = core.BoolPtr(true)
				segmentRuleModel.Order = core.Int64Ptr(int64(1))
				Expect(segmentRuleModel.Rules).To(Equal([]appconfigurationv1.Rule{*ruleModel}))
				Expect(segmentRuleModel.Value).To(Equal(core.BoolPtr(true)))
				Expect(segmentRuleModel.Order).To(Equal(core.Int64Ptr(int64(1))))

				// Construct an instance of the Collection model
				collectionModel := new(appconfigurationv1.Collection)
				Expect(collectionModel).ToNot(BeNil())
				collectionModel.CollectionID = core.StringPtr("ghzinc")
				collectionModel.Enabled = core.BoolPtr(true)
				Expect(collectionModel.CollectionID).To(Equal(core.StringPtr("ghzinc")))
				Expect(collectionModel.Enabled).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the CreateFeatureOptions model
				createFeatureOptionsName := "Cycle Rentals"
				createFeatureOptionsId := "cycle-rentals"
				createFeatureOptionsDescription := "Feature flags to enable Cycle Rentals"
				createFeatureOptionsType := "Boolean"
				createFeatureOptionsEnabledValue := "true"
				createFeatureOptionsDisabledValue := "false"
				createFeatureOptionsTags := "version: 1.1, pre-release"
				createFeatureOptionsSegmentRules := []appconfigurationv1.SegmentRule{}
				createFeatureOptionsCollections := []appconfigurationv1.Collection{}
				createFeatureOptionsCreatedMode := "API"
				createFeatureOptionsModel := appConfigurationService.NewCreateFeatureOptions(createFeatureOptionsName, createFeatureOptionsId, createFeatureOptionsDescription, createFeatureOptionsType, createFeatureOptionsEnabledValue, createFeatureOptionsDisabledValue, createFeatureOptionsTags, createFeatureOptionsSegmentRules, createFeatureOptionsCollections, createFeatureOptionsCreatedMode)
				createFeatureOptionsModel.SetName("Cycle Rentals")
				createFeatureOptionsModel.SetDescription("Feature flags to enable Cycle Rentals")
				createFeatureOptionsModel.SetType("Boolean")
				createFeatureOptionsModel.SetEnabledValue("true")
				createFeatureOptionsModel.SetDisabledValue("false")
				createFeatureOptionsModel.SetTags("version: 1.1, pre-release")
				createFeatureOptionsModel.SetSegmentRules([]appconfigurationv1.SegmentRule{*segmentRuleModel})
				createFeatureOptionsModel.SetCollections([]appconfigurationv1.Collection{*collectionModel})
				createFeatureOptionsModel.SetCreatedMode("API")
				createFeatureOptionsModel.SetFeatureID("cycle-rentals")
				createFeatureOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createFeatureOptionsModel).ToNot(BeNil())
				Expect(createFeatureOptionsModel.Name).To(Equal(core.StringPtr("Cycle Rentals")))
				Expect(createFeatureOptionsModel.Description).To(Equal(core.StringPtr("Feature flags to enable Cycle Rentals")))
				Expect(createFeatureOptionsModel.Type).To(Equal(core.StringPtr("Boolean")))
				Expect(createFeatureOptionsModel.EnabledValue).To(Equal(core.StringPtr("true")))
				Expect(createFeatureOptionsModel.DisabledValue).To(Equal(core.StringPtr("false")))
				Expect(createFeatureOptionsModel.Tags).To(Equal(core.StringPtr("version: 1.1, pre-release")))
				Expect(createFeatureOptionsModel.SegmentRules).To(Equal([]appconfigurationv1.SegmentRule{*segmentRuleModel}))
				Expect(createFeatureOptionsModel.Collections).To(Equal([]appconfigurationv1.Collection{*collectionModel}))
				Expect(createFeatureOptionsModel.CreatedMode).To(Equal(core.StringPtr("API")))
				Expect(createFeatureOptionsModel.FeatureID).To(Equal(core.StringPtr("cycle-rentals")))
				Expect(createFeatureOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateSegmentOptions successfully`, func() {
				// Construct an instance of the RuleArray model
				ruleArrayModel := new(appconfigurationv1.RuleArray)
				Expect(ruleArrayModel).ToNot(BeNil())
				ruleArrayModel.AttributeName = core.StringPtr("email")
				ruleArrayModel.Operator = core.StringPtr("endsWith")
				ruleArrayModel.Values = []string{"testString"}
				Expect(ruleArrayModel.AttributeName).To(Equal(core.StringPtr("email")))
				Expect(ruleArrayModel.Operator).To(Equal(core.StringPtr("endsWith")))
				Expect(ruleArrayModel.Values).To(Equal([]string{"testString"}))

				// Construct an instance of the CreateSegmentOptions model
				createSegmentOptionsName := "Beta Users"
				createSegmentOptionsId := "beta-users"
				createSegmentOptionsDescription := "Segment containing the beta users"
				createSegmentOptionsTags := "version: 1.1, stage"
				createSegmentOptionsRules := []appconfigurationv1.RuleArray{}
				createSegmentOptionsCreatedMode := "API"
				createSegmentOptionsModel := appConfigurationService.NewCreateSegmentOptions(createSegmentOptionsName, createSegmentOptionsId, createSegmentOptionsDescription, createSegmentOptionsTags, createSegmentOptionsRules, createSegmentOptionsCreatedMode)
				createSegmentOptionsModel.SetName("Beta Users")
				createSegmentOptionsModel.SetDescription("Segment containing the beta users")
				createSegmentOptionsModel.SetTags("version: 1.1, stage")
				createSegmentOptionsModel.SetRules([]appconfigurationv1.RuleArray{*ruleArrayModel})
				createSegmentOptionsModel.SetCreatedMode("API")
				createSegmentOptionsModel.SetSegmentID("beta-users")
				createSegmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createSegmentOptionsModel).ToNot(BeNil())
				Expect(createSegmentOptionsModel.Name).To(Equal(core.StringPtr("Beta Users")))
				Expect(createSegmentOptionsModel.Description).To(Equal(core.StringPtr("Segment containing the beta users")))
				Expect(createSegmentOptionsModel.Tags).To(Equal(core.StringPtr("version: 1.1, stage")))
				Expect(createSegmentOptionsModel.Rules).To(Equal([]appconfigurationv1.RuleArray{*ruleArrayModel}))
				Expect(createSegmentOptionsModel.CreatedMode).To(Equal(core.StringPtr("API")))
				Expect(createSegmentOptionsModel.SegmentID).To(Equal(core.StringPtr("beta-users")))
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
			It(`Invoke NewDeleteafeatureOptions successfully`, func() {
				// Construct an instance of the DeleteafeatureOptions model
				featureID := "testString"
				deleteafeatureOptionsModel := appConfigurationService.NewDeleteafeatureOptions(featureID)
				deleteafeatureOptionsModel.SetFeatureID("testString")
				deleteafeatureOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteafeatureOptionsModel).ToNot(BeNil())
				Expect(deleteafeatureOptionsModel.FeatureID).To(Equal(core.StringPtr("testString")))
				Expect(deleteafeatureOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteasegmentOptions successfully`, func() {
				// Construct an instance of the DeleteSegmentOptions model
				segmentID := "testString"
				deleteasegmentOptionsModel := appConfigurationService.NewDeleteasegmentOptions(segmentID)
				deleteasegmentOptionsModel.SetSegmentID("testString")
				deleteasegmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteasegmentOptionsModel).ToNot(BeNil())
				Expect(deleteasegmentOptionsModel.SegmentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteasegmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCollectionOptions successfully`, func() {
				// Construct an instance of the GetCollectionOptions model
				collectionID := "testString"
				getCollectionOptionsModel := appConfigurationService.NewGetCollectionOptions(collectionID)
				getCollectionOptionsModel.SetCollectionID("testString")
				getCollectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCollectionOptionsModel).ToNot(BeNil())
				Expect(getCollectionOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(getCollectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCollectionsOptions successfully`, func() {
				// Construct an instance of the GetCollectionsOptions model
				getCollectionsOptionsModel := appConfigurationService.NewGetCollectionsOptions()
				getCollectionsOptionsModel.SetSize("{{size}}")
				getCollectionsOptionsModel.SetOffset("{{offset}}")
				getCollectionsOptionsModel.SetFeatures("{{feature_id}}")
				getCollectionsOptionsModel.SetTags("{{tag1,tag2}}")
				getCollectionsOptionsModel.SetExpand("{{booleanvalue}}")
				getCollectionsOptionsModel.SetInclude("features")
				getCollectionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCollectionsOptionsModel).ToNot(BeNil())
				Expect(getCollectionsOptionsModel.Size).To(Equal(core.StringPtr("{{size}}")))
				Expect(getCollectionsOptionsModel.Offset).To(Equal(core.StringPtr("{{offset}}")))
				Expect(getCollectionsOptionsModel.Features).To(Equal(core.StringPtr("{{feature_id}}")))
				Expect(getCollectionsOptionsModel.Tags).To(Equal(core.StringPtr("{{tag1,tag2}}")))
				Expect(getCollectionsOptionsModel.Expand).To(Equal(core.StringPtr("{{booleanvalue}}")))
				Expect(getCollectionsOptionsModel.Include).To(Equal(core.StringPtr("features")))
				Expect(getCollectionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetFeatureOptions successfully`, func() {
				// Construct an instance of the GetFeatureOptions model
				featureID := "testString"
				getfeaturedetailsOptionsModel := appConfigurationService.NewGetFeatureOptions(featureID)
				getfeaturedetailsOptionsModel.SetFeatureID("testString")
				getfeaturedetailsOptionsModel.SetInclude("collections")
				getfeaturedetailsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getfeaturedetailsOptionsModel).ToNot(BeNil())
				Expect(getfeaturedetailsOptionsModel.FeatureID).To(Equal(core.StringPtr("testString")))
				Expect(getfeaturedetailsOptionsModel.Include).To(Equal(core.StringPtr("collections")))
				Expect(getfeaturedetailsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetFeaturesOptions successfully`, func() {
				// Construct an instance of the GetFeaturesOptions model
				getlistoffeaturesOptionsModel := appConfigurationService.NewGetFeaturesOptions()
				getlistoffeaturesOptionsModel.SetSize("{{size}}")
				getlistoffeaturesOptionsModel.SetOffset("{{offset}}")
				getlistoffeaturesOptionsModel.SetTags("{{commaseparatedtags}}")
				getlistoffeaturesOptionsModel.SetCollections("{{commaseparatedcollections}}")
				getlistoffeaturesOptionsModel.SetSegments("{{commaseparatedsegments}}")
				getlistoffeaturesOptionsModel.SetExpand("{{boolean}}")
				getlistoffeaturesOptionsModel.SetInclude("collections/rules")
				getlistoffeaturesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getlistoffeaturesOptionsModel).ToNot(BeNil())
				Expect(getlistoffeaturesOptionsModel.Size).To(Equal(core.StringPtr("{{size}}")))
				Expect(getlistoffeaturesOptionsModel.Offset).To(Equal(core.StringPtr("{{offset}}")))
				Expect(getlistoffeaturesOptionsModel.Tags).To(Equal(core.StringPtr("{{commaseparatedtags}}")))
				Expect(getlistoffeaturesOptionsModel.Collections).To(Equal(core.StringPtr("{{commaseparatedcollections}}")))
				Expect(getlistoffeaturesOptionsModel.Segments).To(Equal(core.StringPtr("{{commaseparatedsegments}}")))
				Expect(getlistoffeaturesOptionsModel.Expand).To(Equal(core.StringPtr("{{boolean}}")))
				Expect(getlistoffeaturesOptionsModel.Include).To(Equal(core.StringPtr("collections/rules")))
				Expect(getlistoffeaturesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSegmentsOptions successfully`, func() {
				// Construct an instance of the GetSegmentsOptions model
				getlistofsegmentsOptionsModel := appConfigurationService.NewGetSegmentsOptions()
				getlistofsegmentsOptionsModel.SetSize("{{size}}")
				getlistofsegmentsOptionsModel.SetOffset("{{offset}}")
				getlistofsegmentsOptionsModel.SetTags("{{commaseparatedtags}}")
				getlistofsegmentsOptionsModel.SetFeatures("{{commaseparatedfeatures}}")
				getlistofsegmentsOptionsModel.SetExpand("{{$randomBoolean}}")
				getlistofsegmentsOptionsModel.SetInclude("rules")
				getlistofsegmentsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getlistofsegmentsOptionsModel).ToNot(BeNil())
				Expect(getlistofsegmentsOptionsModel.Size).To(Equal(core.StringPtr("{{size}}")))
				Expect(getlistofsegmentsOptionsModel.Offset).To(Equal(core.StringPtr("{{offset}}")))
				Expect(getlistofsegmentsOptionsModel.Tags).To(Equal(core.StringPtr("{{commaseparatedtags}}")))
				Expect(getlistofsegmentsOptionsModel.Features).To(Equal(core.StringPtr("{{commaseparatedfeatures}}")))
				Expect(getlistofsegmentsOptionsModel.Expand).To(Equal(core.StringPtr("{{$randomBoolean}}")))
				Expect(getlistofsegmentsOptionsModel.Include).To(Equal(core.StringPtr("rules")))
				Expect(getlistofsegmentsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSegmentOptions successfully`, func() {
				// Construct an instance of the GetSegmentOptions model
				segmentID := "/segments/testString"
				getsegmentdetailsOptionsModel := appConfigurationService.NewGetSegmentOptions(segmentID)
				getsegmentdetailsOptionsModel.SetSegmentID(segmentID)
				getsegmentdetailsOptionsModel.SetInclude("features")
				getsegmentdetailsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getsegmentdetailsOptionsModel).ToNot(BeNil())
				Expect(getsegmentdetailsOptionsModel.SegmentID).To(Equal(core.StringPtr("/segments/testString")))
				Expect(getsegmentdetailsOptionsModel.Include).To(Equal(core.StringPtr("features")))
				Expect(getsegmentdetailsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRule successfully`, func() {
				segments := []string{"testString"}
				model, err := appConfigurationService.NewRule(segments)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuleArray successfully`, func() {
				attributeName := "testString"
				operator := "testString"
				values := []string{"testString"}
				model, err := appConfigurationService.NewRuleArray(attributeName, operator, values)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSegmentRule successfully`, func() {
				rules := []appconfigurationv1.Rule{}
				value := "true"
				order := int64(38)
				model, err := appConfigurationService.NewSegmentRule(rules, value, order)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateCollectionOptions successfully`, func() {
				// Construct an instance of the UpdateCollectionOptions model
				collectionID := "testString"
				updateCollectionOptionsName := "GHz Inc Updated"
				updateCollectionOptionsDescription := "Collection for GHz Inc updated"
				updateCollectionOptionsTags := "version: 1.1, pre-release, new tag addition"
				updateCollectionOptionsModel := appConfigurationService.NewUpdateCollectionOptions(collectionID, updateCollectionOptionsName, updateCollectionOptionsDescription, updateCollectionOptionsTags)
				updateCollectionOptionsModel.SetCollectionID("testString")
				updateCollectionOptionsModel.SetName("GHz Inc Updated")
				updateCollectionOptionsModel.SetDescription("Collection for GHz Inc updated")
				updateCollectionOptionsModel.SetTags("version: 1.1, pre-release, new tag addition")
				updateCollectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateCollectionOptionsModel).ToNot(BeNil())
				Expect(updateCollectionOptionsModel.CollectionID).To(Equal(core.StringPtr("testString")))
				Expect(updateCollectionOptionsModel.Name).To(Equal(core.StringPtr("GHz Inc Updated")))
				Expect(updateCollectionOptionsModel.Description).To(Equal(core.StringPtr("Collection for GHz Inc updated")))
				Expect(updateCollectionOptionsModel.Tags).To(Equal(core.StringPtr("version: 1.1, pre-release, new tag addition")))
				Expect(updateCollectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateFeatureOptions successfully`, func() {
				// Construct an instance of the Rule model
				ruleModel := new(appconfigurationv1.Rule)
				Expect(ruleModel).ToNot(BeNil())
				ruleModel.Segments = []string{"testString"}
				Expect(ruleModel.Segments).To(Equal([]string{"testString"}))

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				Expect(segmentRuleModel).ToNot(BeNil())
				segmentRuleModel.Rules = []appconfigurationv1.Rule{*ruleModel}
				segmentRuleModel.Value = core.BoolPtr(true)
				segmentRuleModel.Order = core.Int64Ptr(int64(1))
				Expect(segmentRuleModel.Rules).To(Equal([]appconfigurationv1.Rule{*ruleModel}))
				Expect(segmentRuleModel.Value).To(Equal(core.BoolPtr(true)))
				Expect(segmentRuleModel.Order).To(Equal(core.Int64Ptr(int64(1))))

				// Construct an instance of the CollectionWithDeletedFlag model
				collection2Model := new(appconfigurationv1.CollectionWithDeletedFlag)
				Expect(collection2Model).ToNot(BeNil())
				collection2Model.CollectionID = core.StringPtr("ghzinc")
				collection2Model.Enabled = core.BoolPtr(true)
				collection2Model.Deleted = core.BoolPtr(true)
				Expect(collection2Model.CollectionID).To(Equal(core.StringPtr("ghzinc")))
				Expect(collection2Model.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(collection2Model.Deleted).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the UpdateFeatureOptions model
				featureID := "testString"
				updatefeaturepropertiesOptionsName := "Cycle Rentals"
				updatefeaturepropertiesOptionsDescription := "Feature flags to enable Cycle Rentals"
				updatefeaturepropertiesOptionsType := "Boolean"
				updatefeaturepropertiesOptionsEnabledValue := "true"
				updatefeaturepropertiesOptionsDisabledValue := "false"
				updatefeaturepropertiesOptionsTags := "version: 1.1, yet-to-release"
				updatefeaturepropertiesOptionsSegmentRules := []appconfigurationv1.SegmentRule{}
				updatefeaturepropertiesOptionsCollections := []appconfigurationv1.CollectionWithDeletedFlag{}
				updatefeaturepropertiesOptionsModel := appConfigurationService.NewUpdateFeatureOptions(featureID, updatefeaturepropertiesOptionsName, updatefeaturepropertiesOptionsDescription, updatefeaturepropertiesOptionsType, updatefeaturepropertiesOptionsEnabledValue, updatefeaturepropertiesOptionsDisabledValue, updatefeaturepropertiesOptionsTags, updatefeaturepropertiesOptionsSegmentRules, updatefeaturepropertiesOptionsCollections)
				updatefeaturepropertiesOptionsModel.SetFeatureID("testString")
				updatefeaturepropertiesOptionsModel.SetName("Cycle Rentals")
				updatefeaturepropertiesOptionsModel.SetDescription("Feature flags to enable Cycle Rentals")
				updatefeaturepropertiesOptionsModel.SetType("Boolean")
				updatefeaturepropertiesOptionsModel.SetEnabledValue("true")
				updatefeaturepropertiesOptionsModel.SetDisabledValue("false")
				updatefeaturepropertiesOptionsModel.SetTags("version: 1.1, yet-to-release")
				updatefeaturepropertiesOptionsModel.SetSegmentRules([]appconfigurationv1.SegmentRule{*segmentRuleModel})
				updatefeaturepropertiesOptionsModel.SetCollections([]appconfigurationv1.CollectionWithDeletedFlag{*collection2Model})
				updatefeaturepropertiesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updatefeaturepropertiesOptionsModel).ToNot(BeNil())
				Expect(updatefeaturepropertiesOptionsModel.FeatureID).To(Equal(core.StringPtr("testString")))
				Expect(updatefeaturepropertiesOptionsModel.Name).To(Equal(core.StringPtr("Cycle Rentals")))
				Expect(updatefeaturepropertiesOptionsModel.Description).To(Equal(core.StringPtr("Feature flags to enable Cycle Rentals")))
				Expect(updatefeaturepropertiesOptionsModel.Type).To(Equal(core.StringPtr("Boolean")))
				Expect(updatefeaturepropertiesOptionsModel.EnabledValue).To(Equal(core.StringPtr("true")))
				Expect(updatefeaturepropertiesOptionsModel.DisabledValue).To(Equal(core.StringPtr("false")))
				Expect(updatefeaturepropertiesOptionsModel.Tags).To(Equal(core.StringPtr("version: 1.1, yet-to-release")))
				Expect(updatefeaturepropertiesOptionsModel.SegmentRules).To(Equal([]appconfigurationv1.SegmentRule{*segmentRuleModel}))
				Expect(updatefeaturepropertiesOptionsModel.Collections).To(Equal([]appconfigurationv1.CollectionWithDeletedFlag{*collection2Model}))
				Expect(updatefeaturepropertiesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateSegmentOptions successfully`, func() {
				// Construct an instance of the RuleArray model
				ruleArrayModel := new(appconfigurationv1.RuleArray)
				Expect(ruleArrayModel).ToNot(BeNil())
				ruleArrayModel.AttributeName = core.StringPtr("email")
				ruleArrayModel.Operator = core.StringPtr("endsWith")
				ruleArrayModel.Values = []string{"testString"}
				Expect(ruleArrayModel.AttributeName).To(Equal(core.StringPtr("email")))
				Expect(ruleArrayModel.Operator).To(Equal(core.StringPtr("endsWith")))
				Expect(ruleArrayModel.Values).To(Equal([]string{"testString"}))

				// Construct an instance of the UpdateSegmentOptions model
				segmentID := "testString"
				updatethesegmentOptionsName := "Beta Users"
				updatethesegmentOptionsDescription := "Segment containing the beta users"
				updatethesegmentOptionsTags := "version: 1.1, pre-release"
				updatethesegmentOptionsRules := []appconfigurationv1.RuleArray{}
				updatethesegmentOptionsModel := appConfigurationService.NewUpdateSegmentOptions(segmentID, updatethesegmentOptionsName, updatethesegmentOptionsDescription, updatethesegmentOptionsTags, updatethesegmentOptionsRules)
				updatethesegmentOptionsModel.SetSegmentID("testString")
				updatethesegmentOptionsModel.SetName("Beta Users")
				updatethesegmentOptionsModel.SetDescription("Segment containing the beta users")
				updatethesegmentOptionsModel.SetTags("version: 1.1, pre-release")
				updatethesegmentOptionsModel.SetRules([]appconfigurationv1.RuleArray{*ruleArrayModel})
				updatethesegmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updatethesegmentOptionsModel).ToNot(BeNil())
				Expect(updatethesegmentOptionsModel.SegmentID).To(Equal(core.StringPtr("testString")))
				Expect(updatethesegmentOptionsModel.Name).To(Equal(core.StringPtr("Beta Users")))
				Expect(updatethesegmentOptionsModel.Description).To(Equal(core.StringPtr("Segment containing the beta users")))
				Expect(updatethesegmentOptionsModel.Tags).To(Equal(core.StringPtr("version: 1.1, pre-release")))
				Expect(updatethesegmentOptionsModel.Rules).To(Equal([]appconfigurationv1.RuleArray{*ruleArrayModel}))
				Expect(updatethesegmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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

var _ = Describe(`AppConfigurationV1`, func() {
	var testServer *httptest.Server
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
				createCollectionOptionsModel.Name = core.StringPtr("GHz Inc")
				createCollectionOptionsModel.CollectionID = core.StringPtr("ghzinc")
				createCollectionOptionsModel.Description = core.StringPtr("Collection for GHz Inc")
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "guid": "Guid", "collection_id": "CollectionID", "description": "Description", "created_time": "CreatedTime", "updated_time": "UpdatedTime"}`)
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
				createCollectionOptionsModel.Name = core.StringPtr("GHz Inc")
				createCollectionOptionsModel.CollectionID = core.StringPtr("ghzinc")
				createCollectionOptionsModel.Description = core.StringPtr("Collection for GHz Inc")
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
				createCollectionOptionsModel.Name = core.StringPtr("GHz Inc")
				createCollectionOptionsModel.CollectionID = core.StringPtr("ghzinc")
				createCollectionOptionsModel.Description = core.StringPtr("Collection for GHz Inc")
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
	Describe(`GetCollections(getCollectionsOptions *GetCollectionsOptions) - Operation response error`, func() {
		getCollectionsPath := "/collections"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCollectionsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["size"]).To(Equal([]string{"{{size}}"}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{"{{offset}}"}))

					Expect(req.URL.Query()["features"]).To(Equal([]string{"{{feature_id}}"}))

					Expect(req.URL.Query()["tags"]).To(Equal([]string{"{{tag1,tag2}}"}))

					Expect(req.URL.Query()["expand"]).To(Equal([]string{"{{booleanvalue}}"}))

					Expect(req.URL.Query()["include"]).To(Equal([]string{"features"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCollections with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetCollectionsOptions model
				getCollectionsOptionsModel := new(appconfigurationv1.GetCollectionsOptions)
				getCollectionsOptionsModel.Size = core.StringPtr("{{size}}")
				getCollectionsOptionsModel.Offset = core.StringPtr("{{offset}}")
				getCollectionsOptionsModel.Features = core.StringPtr("{{feature_id}}")
				getCollectionsOptionsModel.Tags = core.StringPtr("{{tag1,tag2}}")
				getCollectionsOptionsModel.Expand = core.StringPtr("{{booleanvalue}}")
				getCollectionsOptionsModel.Include = core.StringPtr("features")
				getCollectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.GetCollections(getCollectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.GetCollections(getCollectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetCollections(getCollectionsOptions *GetCollectionsOptions)`, func() {
		getCollectionsPath := "/collections"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCollectionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["size"]).To(Equal([]string{"{{size}}"}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{"{{offset}}"}))

					Expect(req.URL.Query()["features"]).To(Equal([]string{"{{feature_id}}"}))

					Expect(req.URL.Query()["tags"]).To(Equal([]string{"{{tag1,tag2}}"}))

					Expect(req.URL.Query()["expand"]).To(Equal([]string{"{{booleanvalue}}"}))

					Expect(req.URL.Query()["include"]).To(Equal([]string{"features"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"collections": [{"name": "Name", "collection_id": "CollectionID", "description": "Description"}], "page_info": {"total_count": 10, "count": 5}}`)
				}))
			})
			It(`Invoke GetCollections successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.GetCollections(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCollectionsOptions model
				getCollectionsOptionsModel := new(appconfigurationv1.GetCollectionsOptions)
				getCollectionsOptionsModel.Size = core.StringPtr("{{size}}")
				getCollectionsOptionsModel.Offset = core.StringPtr("{{offset}}")
				getCollectionsOptionsModel.Features = core.StringPtr("{{feature_id}}")
				getCollectionsOptionsModel.Tags = core.StringPtr("{{tag1,tag2}}")
				getCollectionsOptionsModel.Expand = core.StringPtr("{{booleanvalue}}")
				getCollectionsOptionsModel.Include = core.StringPtr("features")
				getCollectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.GetCollections(getCollectionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.GetCollectionsWithContext(ctx, getCollectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr = appConfigurationService.GetCollections(getCollectionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.GetCollectionsWithContext(ctx, getCollectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetCollections with error: Operation request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetCollectionsOptions model
				getCollectionsOptionsModel := new(appconfigurationv1.GetCollectionsOptions)
				getCollectionsOptionsModel.Size = core.StringPtr("{{size}}")
				getCollectionsOptionsModel.Offset = core.StringPtr("{{offset}}")
				getCollectionsOptionsModel.Features = core.StringPtr("{{feature_id}}")
				getCollectionsOptionsModel.Tags = core.StringPtr("{{tag1,tag2}}")
				getCollectionsOptionsModel.Expand = core.StringPtr("{{booleanvalue}}")
				getCollectionsOptionsModel.Include = core.StringPtr("features")
				getCollectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.GetCollections(getCollectionsOptionsModel)
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
				updateCollectionOptionsModel.Name = core.StringPtr("GHz Inc Updated")
				updateCollectionOptionsModel.Description = core.StringPtr("Collection for GHz Inc updated")
				updateCollectionOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release, new tag addition")
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "collection_id": "CollectionID", "description": "Description", "created_time": "CreatedTime", "updated_time": "UpdatedTime"}`)
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
				updateCollectionOptionsModel.Name = core.StringPtr("GHz Inc Updated")
				updateCollectionOptionsModel.Description = core.StringPtr("Collection for GHz Inc updated")
				updateCollectionOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release, new tag addition")
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
				updateCollectionOptionsModel.Name = core.StringPtr("GHz Inc Updated")
				updateCollectionOptionsModel.Description = core.StringPtr("Collection for GHz Inc updated")
				updateCollectionOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release, new tag addition")
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

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "collection_id": "CollectionID"}`)
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
})

var _ = Describe(`AppConfigurationV1`, func() {
	var testServer *httptest.Server

	Describe(`CreateFeature(createFeatureOptions *CreateFeatureOptions) - Operation response error`, func() {
		createFeaturePath := "/features"
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

				// Construct an instance of the Rule model
				ruleModel := new(appconfigurationv1.Rule)
				ruleModel.Segments = []string{"testString"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.Rule{*ruleModel}
				segmentRuleModel.Value = core.BoolPtr(true)
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the Collection model
				collectionModel := new(appconfigurationv1.Collection)
				collectionModel.CollectionID = core.StringPtr("ghzinc")
				collectionModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the CreateFeatureOptions model
				createFeatureOptionsModel := new(appconfigurationv1.CreateFeatureOptions)
				createFeatureOptionsModel.Name = core.StringPtr("Cycle Rentals")
				createFeatureOptionsModel.Description = core.StringPtr("Feature flags to enable Cycle Rentals")
				createFeatureOptionsModel.Type = core.StringPtr("Boolean")
				createFeatureOptionsModel.EnabledValue = core.StringPtr("true")
				createFeatureOptionsModel.DisabledValue = core.StringPtr("false")
				createFeatureOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				createFeatureOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				createFeatureOptionsModel.Collections = []appconfigurationv1.Collection{*collectionModel}
				createFeatureOptionsModel.CreatedMode = core.StringPtr("API")
				createFeatureOptionsModel.FeatureID = core.StringPtr("cycle-rentals")
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
		createFeaturePath := "/features"
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "feature_id": "FeatureID", "description": "Description", "type": "Type", "enabled_value": "true", "disabled_value": "false", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": false, "order": 5}], "collections": [{"collection_id": "CollectionID", "enabled": false}], "created_time": "CreatedTime", "updated_time": "UpdatedTime"}`)
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

				// Construct an instance of the Rule model
				ruleModel := new(appconfigurationv1.Rule)
				ruleModel.Segments = []string{"testString"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.Rule{*ruleModel}
				segmentRuleModel.Value = core.BoolPtr(true)
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the Collection model
				collectionModel := new(appconfigurationv1.Collection)
				collectionModel.CollectionID = core.StringPtr("ghzinc")
				collectionModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the CreateFeatureOptions model
				createFeatureOptionsModel := new(appconfigurationv1.CreateFeatureOptions)
				createFeatureOptionsModel.Name = core.StringPtr("Cycle Rentals")
				createFeatureOptionsModel.Description = core.StringPtr("Feature flags to enable Cycle Rentals")
				createFeatureOptionsModel.Type = core.StringPtr("Boolean")
				createFeatureOptionsModel.EnabledValue = core.StringPtr("true")
				createFeatureOptionsModel.DisabledValue = core.StringPtr("false")
				createFeatureOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				createFeatureOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				createFeatureOptionsModel.Collections = []appconfigurationv1.Collection{*collectionModel}
				createFeatureOptionsModel.CreatedMode = core.StringPtr("API")
				createFeatureOptionsModel.FeatureID = core.StringPtr("cycle-rentals")
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

				// Construct an instance of the Rule model
				ruleModel := new(appconfigurationv1.Rule)
				ruleModel.Segments = []string{"testString"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.Rule{*ruleModel}
				segmentRuleModel.Value = core.BoolPtr(true)
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the Collection model
				collectionModel := new(appconfigurationv1.Collection)
				collectionModel.CollectionID = core.StringPtr("ghzinc")
				collectionModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the CreateFeatureOptions model
				createFeatureOptionsModel := new(appconfigurationv1.CreateFeatureOptions)
				createFeatureOptionsModel.Name = core.StringPtr("Cycle Rentals")
				createFeatureOptionsModel.Description = core.StringPtr("Feature flags to enable Cycle Rentals")
				createFeatureOptionsModel.Type = core.StringPtr("Boolean")
				createFeatureOptionsModel.EnabledValue = core.StringPtr("true")
				createFeatureOptionsModel.DisabledValue = core.StringPtr("false")
				createFeatureOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				createFeatureOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				createFeatureOptionsModel.Collections = []appconfigurationv1.Collection{*collectionModel}
				createFeatureOptionsModel.CreatedMode = core.StringPtr("API")
				createFeatureOptionsModel.FeatureID = core.StringPtr("cycle-rentals")
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
	Describe(`GetFeatures(getlistoffeaturesOptions *GetFeaturesOptions) - Operation response error`, func() {
		getlistoffeaturesPath := "/features"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getlistoffeaturesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["size"]).To(Equal([]string{"{{size}}"}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{"{{offset}}"}))

					Expect(req.URL.Query()["tags"]).To(Equal([]string{"{{commaseparatedtags}}"}))

					Expect(req.URL.Query()["collections"]).To(Equal([]string{"{{commaseparatedcollections}}"}))

					Expect(req.URL.Query()["segments"]).To(Equal([]string{"{{commaseparatedsegments}}"}))

					Expect(req.URL.Query()["expand"]).To(Equal([]string{"{{boolean}}"}))

					Expect(req.URL.Query()["include"]).To(Equal([]string{"collections/rules"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetFeatures with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetFeaturesOptions model
				getlistoffeaturesOptionsModel := new(appconfigurationv1.GetFeaturesOptions)
				getlistoffeaturesOptionsModel.Size = core.StringPtr("{{size}}")
				getlistoffeaturesOptionsModel.Offset = core.StringPtr("{{offset}}")
				getlistoffeaturesOptionsModel.Tags = core.StringPtr("{{commaseparatedtags}}")
				getlistoffeaturesOptionsModel.Collections = core.StringPtr("{{commaseparatedcollections}}")
				getlistoffeaturesOptionsModel.Segments = core.StringPtr("{{commaseparatedsegments}}")
				getlistoffeaturesOptionsModel.Expand = core.StringPtr("{{boolean}}")
				getlistoffeaturesOptionsModel.Include = core.StringPtr("collections/rules")
				getlistoffeaturesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.GetFeatures(getlistoffeaturesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.GetFeatures(getlistoffeaturesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetFeatures(getlistoffeaturesOptions *GetFeaturesOptions)`, func() {
		getlistoffeaturesPath := "/features"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getlistoffeaturesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["size"]).To(Equal([]string{"{{size}}"}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{"{{offset}}"}))

					Expect(req.URL.Query()["tags"]).To(Equal([]string{"{{commaseparatedtags}}"}))

					Expect(req.URL.Query()["collections"]).To(Equal([]string{"{{commaseparatedcollections}}"}))

					Expect(req.URL.Query()["segments"]).To(Equal([]string{"{{commaseparatedsegments}}"}))

					Expect(req.URL.Query()["expand"]).To(Equal([]string{"{{boolean}}"}))

					Expect(req.URL.Query()["include"]).To(Equal([]string{"collections/rules"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"features": [{"name": "Name", "feature_id": "FeatureID", "segment_exists": false, "description": "Description", "type": "Type", "enabled_value": "true", "disabled_value": "false", "created_time": "CreatedTime", "updated_time": "UpdatedTime"}], "page_info": {"total_count": 10, "count": 5}}`)
				}))
			})
			It(`Invoke GetFeatures successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.GetFeatures(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetFeaturesOptions model
				getlistoffeaturesOptionsModel := new(appconfigurationv1.GetFeaturesOptions)
				getlistoffeaturesOptionsModel.Size = core.StringPtr("{{size}}")
				getlistoffeaturesOptionsModel.Offset = core.StringPtr("{{offset}}")
				getlistoffeaturesOptionsModel.Tags = core.StringPtr("{{commaseparatedtags}}")
				getlistoffeaturesOptionsModel.Collections = core.StringPtr("{{commaseparatedcollections}}")
				getlistoffeaturesOptionsModel.Segments = core.StringPtr("{{commaseparatedsegments}}")
				getlistoffeaturesOptionsModel.Expand = core.StringPtr("{{boolean}}")
				getlistoffeaturesOptionsModel.Include = core.StringPtr("collections/rules")
				getlistoffeaturesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.GetFeatures(getlistoffeaturesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.GetFeaturesWithContext(ctx, getlistoffeaturesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr = appConfigurationService.GetFeatures(getlistoffeaturesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.GetFeaturesWithContext(ctx, getlistoffeaturesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetFeatures with error: Operation request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetFeaturesOptions model
				getlistoffeaturesOptionsModel := new(appconfigurationv1.GetFeaturesOptions)
				getlistoffeaturesOptionsModel.Size = core.StringPtr("{{size}}")
				getlistoffeaturesOptionsModel.Offset = core.StringPtr("{{offset}}")
				getlistoffeaturesOptionsModel.Tags = core.StringPtr("{{commaseparatedtags}}")
				getlistoffeaturesOptionsModel.Collections = core.StringPtr("{{commaseparatedcollections}}")
				getlistoffeaturesOptionsModel.Segments = core.StringPtr("{{commaseparatedsegments}}")
				getlistoffeaturesOptionsModel.Expand = core.StringPtr("{{boolean}}")
				getlistoffeaturesOptionsModel.Include = core.StringPtr("collections/rules")
				getlistoffeaturesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.GetFeatures(getlistoffeaturesOptionsModel)
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
	Describe(`UpdateFeature(updatefeaturepropertiesOptions *UpdateFeatureOptions) - Operation response error`, func() {
		updatefeaturepropertiesPath := "/features/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updatefeaturepropertiesPath))
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

				// Construct an instance of the Rule model
				ruleModel := new(appconfigurationv1.Rule)
				ruleModel.Segments = []string{"testString"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.Rule{*ruleModel}
				segmentRuleModel.Value = core.BoolPtr(true)
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the CollectionWithDeletedFlag model
				collection2Model := new(appconfigurationv1.CollectionWithDeletedFlag)
				collection2Model.CollectionID = core.StringPtr("ghzinc")
				collection2Model.Enabled = core.BoolPtr(true)
				collection2Model.Deleted = core.BoolPtr(true)

				// Construct an instance of the UpdateFeatureOptions model
				updatefeaturepropertiesOptionsModel := new(appconfigurationv1.UpdateFeatureOptions)
				updatefeaturepropertiesOptionsModel.FeatureID = core.StringPtr("testString")
				updatefeaturepropertiesOptionsModel.Name = core.StringPtr("Cycle Rentals")
				updatefeaturepropertiesOptionsModel.Description = core.StringPtr("Feature flags to enable Cycle Rentals")
				updatefeaturepropertiesOptionsModel.Type = core.StringPtr("Boolean")
				updatefeaturepropertiesOptionsModel.EnabledValue = core.StringPtr("true")
				updatefeaturepropertiesOptionsModel.DisabledValue = core.StringPtr("false")
				updatefeaturepropertiesOptionsModel.Tags = core.StringPtr("version: 1.1, yet-to-release")
				updatefeaturepropertiesOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				updatefeaturepropertiesOptionsModel.Collections = []appconfigurationv1.CollectionWithDeletedFlag{*collection2Model}
				updatefeaturepropertiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.UpdateFeature(updatefeaturepropertiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.UpdateFeature(updatefeaturepropertiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateFeature(updatefeaturepropertiesOptions *UpdateFeatureOptions)`, func() {
		updatefeaturepropertiesPath := "/features/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updatefeaturepropertiesPath))
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "type": "Type", "enabled_value": "true", "disabled_value": "false", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": false, "order": 5}], "collections": [{"collection_id": "CollectionID", "enabled": false}], "created_time": "CreatedTime", "updated_time": "UpdatedTime"}`)
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

				// Construct an instance of the Rule model
				ruleModel := new(appconfigurationv1.Rule)
				ruleModel.Segments = []string{"testString"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.Rule{*ruleModel}
				segmentRuleModel.Value = core.BoolPtr(true)
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the CollectionWithDeletedFlag model
				collection2Model := new(appconfigurationv1.CollectionWithDeletedFlag)
				collection2Model.CollectionID = core.StringPtr("ghzinc")
				collection2Model.Enabled = core.BoolPtr(true)
				collection2Model.Deleted = core.BoolPtr(true)

				// Construct an instance of the UpdateFeatureOptions model
				updatefeaturepropertiesOptionsModel := new(appconfigurationv1.UpdateFeatureOptions)
				updatefeaturepropertiesOptionsModel.FeatureID = core.StringPtr("testString")
				updatefeaturepropertiesOptionsModel.Name = core.StringPtr("Cycle Rentals")
				updatefeaturepropertiesOptionsModel.Description = core.StringPtr("Feature flags to enable Cycle Rentals")
				updatefeaturepropertiesOptionsModel.Type = core.StringPtr("Boolean")
				updatefeaturepropertiesOptionsModel.EnabledValue = core.StringPtr("true")
				updatefeaturepropertiesOptionsModel.DisabledValue = core.StringPtr("false")
				updatefeaturepropertiesOptionsModel.Tags = core.StringPtr("version: 1.1, yet-to-release")
				updatefeaturepropertiesOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				updatefeaturepropertiesOptionsModel.Collections = []appconfigurationv1.CollectionWithDeletedFlag{*collection2Model}
				updatefeaturepropertiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.UpdateFeature(updatefeaturepropertiesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.UpdatefeaturepropertiesWithContext(ctx, updatefeaturepropertiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr = appConfigurationService.UpdateFeature(updatefeaturepropertiesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.UpdatefeaturepropertiesWithContext(ctx, updatefeaturepropertiesOptionsModel)
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

				// Construct an instance of the Rule model
				ruleModel := new(appconfigurationv1.Rule)
				ruleModel.Segments = []string{"testString"}

				// Construct an instance of the SegmentRule model
				segmentRuleModel := new(appconfigurationv1.SegmentRule)
				segmentRuleModel.Rules = []appconfigurationv1.Rule{*ruleModel}
				segmentRuleModel.Value = core.BoolPtr(true)
				segmentRuleModel.Order = core.Int64Ptr(int64(1))

				// Construct an instance of the CollectionWithDeletedFlag model
				collection2Model := new(appconfigurationv1.CollectionWithDeletedFlag)
				collection2Model.CollectionID = core.StringPtr("ghzinc")
				collection2Model.Enabled = core.BoolPtr(true)
				collection2Model.Deleted = core.BoolPtr(true)

				// Construct an instance of the UpdateFeatureOptions model
				updatefeaturepropertiesOptionsModel := new(appconfigurationv1.UpdateFeatureOptions)
				updatefeaturepropertiesOptionsModel.FeatureID = core.StringPtr("testString")
				updatefeaturepropertiesOptionsModel.Name = core.StringPtr("Cycle Rentals")
				updatefeaturepropertiesOptionsModel.Description = core.StringPtr("Feature flags to enable Cycle Rentals")
				updatefeaturepropertiesOptionsModel.Type = core.StringPtr("Boolean")
				updatefeaturepropertiesOptionsModel.EnabledValue = core.StringPtr("true")
				updatefeaturepropertiesOptionsModel.DisabledValue = core.StringPtr("false")
				updatefeaturepropertiesOptionsModel.Tags = core.StringPtr("version: 1.1, yet-to-release")
				updatefeaturepropertiesOptionsModel.SegmentRules = []appconfigurationv1.SegmentRule{*segmentRuleModel}
				updatefeaturepropertiesOptionsModel.Collections = []appconfigurationv1.CollectionWithDeletedFlag{*collection2Model}
				updatefeaturepropertiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.UpdateFeature(updatefeaturepropertiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateFeatureOptions model with no property values
				updatefeaturepropertiesOptionsModelNew := new(appconfigurationv1.UpdateFeatureOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.UpdateFeature(updatefeaturepropertiesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteFeature(deleteafeatureOptions *DeleteafeatureOptions)`, func() {
		deleteafeaturePath := "/features/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteafeaturePath))
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

				// Construct an instance of the DeleteafeatureOptions model
				deleteafeatureOptionsModel := new(appconfigurationv1.DeleteafeatureOptions)
				deleteafeatureOptionsModel.FeatureID = core.StringPtr("testString")
				deleteafeatureOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appConfigurationService.DeleteFeature(deleteafeatureOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				response, operationErr = appConfigurationService.DeleteFeature(deleteafeatureOptionsModel)
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

				// Construct an instance of the DeleteafeatureOptions model
				deleteafeatureOptionsModel := new(appconfigurationv1.DeleteafeatureOptions)
				deleteafeatureOptionsModel.FeatureID = core.StringPtr("testString")
				deleteafeatureOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appConfigurationService.DeleteFeature(deleteafeatureOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteafeatureOptions model with no property values
				deleteafeatureOptionsModelNew := new(appconfigurationv1.DeleteafeatureOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appConfigurationService.DeleteFeature(deleteafeatureOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetFeature(getfeaturedetailsOptions *GetFeatureOptions) - Operation response error`, func() {
		getfeaturedetailsPath := "/features/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getfeaturedetailsPath))
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
				getfeaturedetailsOptionsModel := new(appconfigurationv1.GetFeatureOptions)
				getfeaturedetailsOptionsModel.FeatureID = core.StringPtr("testString")
				getfeaturedetailsOptionsModel.Include = core.StringPtr("collections")
				getfeaturedetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.GetFeature(getfeaturedetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.GetFeature(getfeaturedetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetFeature(getfeaturedetailsOptions *GetFeatureOptions)`, func() {
		getfeaturedetailsPath := "/features/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getfeaturedetailsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["include"]).To(Equal([]string{"collections"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "feature_id": "FeatureID", "description": "Description", "type": "Type", "enabled_value": "true", "disabled_value": "false", "segment_rules": [{"rules": [{"segments": ["Segments"]}], "value": false, "order": 5}], "created_time": "CreatedTime", "updated_time": "UpdatedTime"}`)
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
				getfeaturedetailsOptionsModel := new(appconfigurationv1.GetFeatureOptions)
				getfeaturedetailsOptionsModel.FeatureID = core.StringPtr("testString")
				getfeaturedetailsOptionsModel.Include = core.StringPtr("collections")
				getfeaturedetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.GetFeature(getfeaturedetailsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.GetFeatureWithContext(ctx, getfeaturedetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr = appConfigurationService.GetFeature(getfeaturedetailsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.GetFeatureWithContext(ctx, getfeaturedetailsOptionsModel)
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
				getfeaturedetailsOptionsModel := new(appconfigurationv1.GetFeatureOptions)
				getfeaturedetailsOptionsModel.FeatureID = core.StringPtr("testString")
				getfeaturedetailsOptionsModel.Include = core.StringPtr("collections")
				getfeaturedetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.GetFeature(getfeaturedetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetFeatureOptions model with no property values
				getfeaturedetailsOptionsModelNew := new(appconfigurationv1.GetFeatureOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.GetFeature(getfeaturedetailsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
})

var _ = Describe(`AppConfigurationV1`, func() {
	var testServer *httptest.Server

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

				// Construct an instance of the RuleArray model
				ruleArrayModel := new(appconfigurationv1.RuleArray)
				ruleArrayModel.AttributeName = core.StringPtr("email")
				ruleArrayModel.Operator = core.StringPtr("endsWith")
				ruleArrayModel.Values = []string{"testString"}

				// Construct an instance of the CreateSegmentOptions model
				createSegmentOptionsModel := new(appconfigurationv1.CreateSegmentOptions)
				createSegmentOptionsModel.Name = core.StringPtr("Beta Users")
				createSegmentOptionsModel.Description = core.StringPtr("Segment containing the beta users")
				createSegmentOptionsModel.Tags = core.StringPtr("version: 1.1, stage")
				createSegmentOptionsModel.Rules = []appconfigurationv1.RuleArray{*ruleArrayModel}
				createSegmentOptionsModel.CreatedMode = core.StringPtr("API")
				createSegmentOptionsModel.SegmentID = core.StringPtr("beta-users")
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "segment_id": "SegmentID", "description": "Description", "rules": [{"attribute_name": "AttributeName", "operator": "Operator", "values": ["Values"]}], "created_time": "CreatedTime", "updated_time": "UpdatedTime"}`)
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

				// Construct an instance of the RuleArray model
				ruleArrayModel := new(appconfigurationv1.RuleArray)
				ruleArrayModel.AttributeName = core.StringPtr("email")
				ruleArrayModel.Operator = core.StringPtr("endsWith")
				ruleArrayModel.Values = []string{"testString"}

				// Construct an instance of the CreateSegmentOptions model
				createSegmentOptionsModel := new(appconfigurationv1.CreateSegmentOptions)
				createSegmentOptionsModel.Name = core.StringPtr("Beta Users")
				createSegmentOptionsModel.Description = core.StringPtr("Segment containing the beta users")
				createSegmentOptionsModel.Tags = core.StringPtr("version: 1.1, stage")
				createSegmentOptionsModel.Rules = []appconfigurationv1.RuleArray{*ruleArrayModel}
				createSegmentOptionsModel.CreatedMode = core.StringPtr("API")
				createSegmentOptionsModel.SegmentID = core.StringPtr("beta-users")
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
			It(`Invoke CreateSegment with error: Operation validation and request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the RuleArray model
				ruleArrayModel := new(appconfigurationv1.RuleArray)
				ruleArrayModel.AttributeName = core.StringPtr("email")
				ruleArrayModel.Operator = core.StringPtr("endsWith")
				ruleArrayModel.Values = []string{"testString"}

				// Construct an instance of the CreateSegmentOptions model
				createSegmentOptionsModel := new(appconfigurationv1.CreateSegmentOptions)
				createSegmentOptionsModel.Name = core.StringPtr("Beta Users")
				createSegmentOptionsModel.Description = core.StringPtr("Segment containing the beta users")
				createSegmentOptionsModel.Tags = core.StringPtr("version: 1.1, stage")
				createSegmentOptionsModel.Rules = []appconfigurationv1.RuleArray{*ruleArrayModel}
				createSegmentOptionsModel.CreatedMode = core.StringPtr("API")
				createSegmentOptionsModel.SegmentID = core.StringPtr("beta-users")
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
	})
	Describe(`GetSegments(getlistofsegmentsOptions *GetSegmentsOptions) - Operation response error`, func() {
		getlistofsegmentsPath := "/segments"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getlistofsegmentsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["size"]).To(Equal([]string{"{{size}}"}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{"{{offset}}"}))

					Expect(req.URL.Query()["tags"]).To(Equal([]string{"{{commaseparatedtags}}"}))

					Expect(req.URL.Query()["features"]).To(Equal([]string{"{{commaseparatedfeatures}}"}))

					Expect(req.URL.Query()["expand"]).To(Equal([]string{"{{$randomBoolean}}"}))

					Expect(req.URL.Query()["include"]).To(Equal([]string{"rules"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSegments with error: Operation response processing error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetSegmentsOptions model
				getlistofsegmentsOptionsModel := new(appconfigurationv1.GetSegmentsOptions)
				getlistofsegmentsOptionsModel.Size = core.StringPtr("{{size}}")
				getlistofsegmentsOptionsModel.Offset = core.StringPtr("{{offset}}")
				getlistofsegmentsOptionsModel.Tags = core.StringPtr("{{commaseparatedtags}}")
				getlistofsegmentsOptionsModel.Features = core.StringPtr("{{commaseparatedfeatures}}")
				getlistofsegmentsOptionsModel.Expand = core.StringPtr("{{$randomBoolean}}")
				getlistofsegmentsOptionsModel.Include = core.StringPtr("rules")
				getlistofsegmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.GetSegments(getlistofsegmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.GetSegments(getlistofsegmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetSegments(getlistofsegmentsOptions *GetSegmentsOptions)`, func() {
		getlistofsegmentsPath := "/segments"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getlistofsegmentsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["size"]).To(Equal([]string{"{{size}}"}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{"{{offset}}"}))

					Expect(req.URL.Query()["tags"]).To(Equal([]string{"{{commaseparatedtags}}"}))

					Expect(req.URL.Query()["features"]).To(Equal([]string{"{{commaseparatedfeatures}}"}))

					Expect(req.URL.Query()["expand"]).To(Equal([]string{"{{$randomBoolean}}"}))

					Expect(req.URL.Query()["include"]).To(Equal([]string{"rules"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"segments": [{"name": "Name", "segment_id": "SegmentID", "description": "Description", "created_time": "CreatedTime", "updated_time": "UpdatedTime", "rules": [{"attribute_name": "AttributeName", "operator": "Operator", "values": ["Values"], "attribut_name": "AttributName"}]}], "page_info": {"total_count": 10, "count": 5}}`)
				}))
			})
			It(`Invoke GetSegments successfully`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())
				appConfigurationService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := appConfigurationService.GetSegments(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSegmentsOptions model
				getlistofsegmentsOptionsModel := new(appconfigurationv1.GetSegmentsOptions)
				getlistofsegmentsOptionsModel.Size = core.StringPtr("{{size}}")
				getlistofsegmentsOptionsModel.Offset = core.StringPtr("{{offset}}")
				getlistofsegmentsOptionsModel.Tags = core.StringPtr("{{commaseparatedtags}}")
				getlistofsegmentsOptionsModel.Features = core.StringPtr("{{commaseparatedfeatures}}")
				getlistofsegmentsOptionsModel.Expand = core.StringPtr("{{$randomBoolean}}")
				getlistofsegmentsOptionsModel.Include = core.StringPtr("rules")
				getlistofsegmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.GetSegments(getlistofsegmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.GetSegmentsWithContext(ctx, getlistofsegmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr = appConfigurationService.GetSegments(getlistofsegmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.GetSegmentsWithContext(ctx, getlistofsegmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetSegments with error: Operation request error`, func() {
				appConfigurationService, serviceErr := appconfigurationv1.NewAppConfigurationV1(&appconfigurationv1.AppConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(appConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetSegmentsOptions model
				getlistofsegmentsOptionsModel := new(appconfigurationv1.GetSegmentsOptions)
				getlistofsegmentsOptionsModel.Size = core.StringPtr("{{size}}")
				getlistofsegmentsOptionsModel.Offset = core.StringPtr("{{offset}}")
				getlistofsegmentsOptionsModel.Tags = core.StringPtr("{{commaseparatedtags}}")
				getlistofsegmentsOptionsModel.Features = core.StringPtr("{{commaseparatedfeatures}}")
				getlistofsegmentsOptionsModel.Expand = core.StringPtr("{{$randomBoolean}}")
				getlistofsegmentsOptionsModel.Include = core.StringPtr("rules")
				getlistofsegmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.GetSegments(getlistofsegmentsOptionsModel)
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
	Describe(`UpdateSegment(updatethesegmentOptions *UpdateSegmentOptions) - Operation response error`, func() {
		updatethesegmentPath := "/segments/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updatethesegmentPath))
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

				// Construct an instance of the RuleArray model
				ruleArrayModel := new(appconfigurationv1.RuleArray)
				ruleArrayModel.AttributeName = core.StringPtr("email")
				ruleArrayModel.Operator = core.StringPtr("endsWith")
				ruleArrayModel.Values = []string{"testString"}

				// Construct an instance of the UpdateSegmentOptions model
				updatethesegmentOptionsModel := new(appconfigurationv1.UpdateSegmentOptions)
				updatethesegmentOptionsModel.SegmentID = core.StringPtr("testString")
				updatethesegmentOptionsModel.Name = core.StringPtr("Beta Users")
				updatethesegmentOptionsModel.Description = core.StringPtr("Segment containing the beta users")
				updatethesegmentOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				updatethesegmentOptionsModel.Rules = []appconfigurationv1.RuleArray{*ruleArrayModel}
				updatethesegmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.UpdateSegment(updatethesegmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.UpdateSegment(updatethesegmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateSegment(updatethesegmentOptions *UpdateSegmentOptions)`, func() {
		updatethesegmentPath := "/segments/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updatethesegmentPath))
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "segment_id": "SegmentID", "description": "Description", "rules": [{"attribute_name": "AttributeName", "operator": "Operator", "values": ["Values"]}], "created_time": "CreatedTime", "updated_time": "UpdatedTime"}`)
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

				// Construct an instance of the RuleArray model
				ruleArrayModel := new(appconfigurationv1.RuleArray)
				ruleArrayModel.AttributeName = core.StringPtr("email")
				ruleArrayModel.Operator = core.StringPtr("endsWith")
				ruleArrayModel.Values = []string{"testString"}

				// Construct an instance of the UpdateSegmentOptions model
				updatethesegmentOptionsModel := new(appconfigurationv1.UpdateSegmentOptions)
				updatethesegmentOptionsModel.SegmentID = core.StringPtr("testString")
				updatethesegmentOptionsModel.Name = core.StringPtr("Beta Users")
				updatethesegmentOptionsModel.Description = core.StringPtr("Segment containing the beta users")
				updatethesegmentOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				updatethesegmentOptionsModel.Rules = []appconfigurationv1.RuleArray{*ruleArrayModel}
				updatethesegmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.UpdateSegment(updatethesegmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.UpdateSegmentWithContext(ctx, updatethesegmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr = appConfigurationService.UpdateSegment(updatethesegmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.UpdateSegmentWithContext(ctx, updatethesegmentOptionsModel)
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

				// Construct an instance of the RuleArray model
				ruleArrayModel := new(appconfigurationv1.RuleArray)
				ruleArrayModel.AttributeName = core.StringPtr("email")
				ruleArrayModel.Operator = core.StringPtr("endsWith")
				ruleArrayModel.Values = []string{"testString"}

				// Construct an instance of the UpdateSegmentOptions model
				updatethesegmentOptionsModel := new(appconfigurationv1.UpdateSegmentOptions)
				updatethesegmentOptionsModel.SegmentID = core.StringPtr("testString")
				updatethesegmentOptionsModel.Name = core.StringPtr("Beta Users")
				updatethesegmentOptionsModel.Description = core.StringPtr("Segment containing the beta users")
				updatethesegmentOptionsModel.Tags = core.StringPtr("version: 1.1, pre-release")
				updatethesegmentOptionsModel.Rules = []appconfigurationv1.RuleArray{*ruleArrayModel}
				updatethesegmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.UpdateSegment(updatethesegmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateSegmentOptions model with no property values
				updatethesegmentOptionsModelNew := new(appconfigurationv1.UpdateSegmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.UpdateSegment(updatethesegmentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSegment(getsegmentdetailsOptions *GetSegmentOptions) - Operation response error`, func() {
		getsegmentdetailsPath := "/segments/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getsegmentdetailsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["segment_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["include"]).To(Equal([]string{"features"}))

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
				getsegmentdetailsOptionsModel := new(appconfigurationv1.GetSegmentOptions)
				getsegmentdetailsOptionsModel.SegmentID = core.StringPtr("testString")
				getsegmentdetailsOptionsModel.Include = core.StringPtr("features")
				getsegmentdetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := appConfigurationService.GetSegment(getsegmentdetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				appConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = appConfigurationService.GetSegment(getsegmentdetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetSegment(getsegmentdetailsOptions *GetSegmentOptions)`, func() {
		getsegmentdetailsPath := "/segments/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getsegmentdetailsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["segment_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["include"]).To(Equal([]string{"features"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "segment_id": "SegmentID", "description": "Description", "rules": [{"attribute_name": "AttributeName", "operator": "Operator", "values": ["Values"]}], "features": [{"feature_id": "FeatureID", "name": "Name"}], "created_time": "CreatedTime", "updated_time": "UpdatedTime"}`)
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
				getsegmentdetailsOptionsModel := new(appconfigurationv1.GetSegmentOptions)
				getsegmentdetailsOptionsModel.SegmentID = core.StringPtr("testString")
				getsegmentdetailsOptionsModel.Include = core.StringPtr("features")
				getsegmentdetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = appConfigurationService.GetSegment(getsegmentdetailsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.GetSegmentWithContext(ctx, getsegmentdetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				result, response, operationErr = appConfigurationService.GetSegment(getsegmentdetailsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = appConfigurationService.GetSegmentWithContext(ctx, getsegmentdetailsOptionsModel)
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
				getsegmentdetailsOptionsModel := new(appconfigurationv1.GetSegmentOptions)
				getsegmentdetailsOptionsModel.SegmentID = core.StringPtr("/segments/testString")
				getsegmentdetailsOptionsModel.Include = core.StringPtr("features")
				getsegmentdetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := appConfigurationService.GetSegment(getsegmentdetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSegmentOptions model with no property values
				getsegmentdetailsOptionsModelNew := new(appconfigurationv1.GetSegmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = appConfigurationService.GetSegment(getsegmentdetailsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteSegment(deleteasegmentOptions *DeleteSegmentOptions)`, func() {
		deleteasegmentPath := "/segments/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteasegmentPath))
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
				deleteasegmentOptionsModel := new(appconfigurationv1.DeleteSegmentOptions)
				deleteasegmentOptionsModel.SegmentID = core.StringPtr("testString")
				deleteasegmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = appConfigurationService.DeleteSegment(deleteasegmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				appConfigurationService.DisableRetries()
				response, operationErr = appConfigurationService.DeleteSegment(deleteasegmentOptionsModel)
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
				deleteasegmentOptionsModel := new(appconfigurationv1.DeleteSegmentOptions)
				deleteasegmentOptionsModel.SegmentID = core.StringPtr("testString")
				deleteasegmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := appConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := appConfigurationService.DeleteSegment(deleteasegmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteSegmentOptions model with no property values
				deleteasegmentOptionsModelNew := new(appconfigurationv1.DeleteSegmentOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = appConfigurationService.DeleteSegment(deleteasegmentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
})
