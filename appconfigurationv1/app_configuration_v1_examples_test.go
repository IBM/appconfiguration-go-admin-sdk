//go:build examples

/**
 * (C) Copyright IBM Corp. 2024.
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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the app-configuration service.
//
// The following configuration properties are assumed to be defined:
// APP_CONFIGURATION_URL=<service base url>
// APP_CONFIGURATION_AUTH_TYPE=iam
// APP_CONFIGURATION_APIKEY=<IAM apikey>
// APP_CONFIGURATION_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
var _ = Describe(`AppConfigurationV1 Examples Tests`, func() {

	const externalConfigFile = "../app_configuration_v1.env"

	var (
		appConfigurationService *appconfigurationv1.AppConfigurationV1
		config       map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping examples...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(appconfigurationv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			appConfigurationServiceOptions := &appconfigurationv1.AppConfigurationV1Options{}

			appConfigurationService, err = appconfigurationv1.NewAppConfigurationV1UsingExternalConfig(appConfigurationServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(appConfigurationService).ToNot(BeNil())
		})
	})

	Describe(`AppConfigurationV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListEnvironments request example`, func() {
			fmt.Println("\nListEnvironments() result:")
			// begin-list_environments
			listEnvironmentsOptions := &appconfigurationv1.ListEnvironmentsOptions{
				Expand: core.BoolPtr(true),
				Sort: core.StringPtr("created_time"),
				Tags: core.StringPtr("version 1.1,pre-release"),
				Include: []string{"features", "properties", "snapshots"},
				Limit: core.Int64Ptr(int64(10)),
				Search: core.StringPtr("test tag"),
			}

			pager, err := appConfigurationService.NewEnvironmentsPager(listEnvironmentsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []appconfigurationv1.Environment
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_environments
		})
		It(`CreateEnvironment request example`, func() {
			fmt.Println("\nCreateEnvironment() result:")
			// begin-create_environment

			createEnvironmentOptions := appConfigurationService.NewCreateEnvironmentOptions(
				"Dev environment",
				"dev-environment",
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

			// end-create_environment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(environment).ToNot(BeNil())
		})
		It(`UpdateEnvironment request example`, func() {
			fmt.Println("\nUpdateEnvironment() result:")
			// begin-update_environment

			updateEnvironmentOptions := appConfigurationService.NewUpdateEnvironmentOptions(
				"environment_id",
			)

			environment, response, err := appConfigurationService.UpdateEnvironment(updateEnvironmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(environment, "", "  ")
			fmt.Println(string(b))

			// end-update_environment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(environment).ToNot(BeNil())
		})
		It(`GetEnvironment request example`, func() {
			fmt.Println("\nGetEnvironment() result:")
			// begin-get_environment

			getEnvironmentOptions := appConfigurationService.NewGetEnvironmentOptions(
				"environment_id",
			)
			getEnvironmentOptions.SetExpand(true)
			getEnvironmentOptions.SetInclude([]string{"features", "properties", "snapshots"})

			environment, response, err := appConfigurationService.GetEnvironment(getEnvironmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(environment, "", "  ")
			fmt.Println(string(b))

			// end-get_environment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(environment).ToNot(BeNil())
		})
		It(`ListCollections request example`, func() {
			fmt.Println("\nListCollections() result:")
			// begin-list_collections
			listCollectionsOptions := &appconfigurationv1.ListCollectionsOptions{
				Expand: core.BoolPtr(true),
				Sort: core.StringPtr("created_time"),
				Tags: core.StringPtr("version 1.1,pre-release"),
				Features: []string{"my-feature-id", "cycle-rentals"},
				Properties: []string{"my-property-id", "email-property"},
				Include: []string{"features", "properties", "snapshots"},
				Limit: core.Int64Ptr(int64(10)),
				Search: core.StringPtr("test tag"),
			}

			pager, err := appConfigurationService.NewCollectionsPager(listCollectionsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []appconfigurationv1.Collection
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_collections
		})
		It(`CreateCollection request example`, func() {
			fmt.Println("\nCreateCollection() result:")
			// begin-create_collection

			createCollectionOptions := appConfigurationService.NewCreateCollectionOptions(
				"Web App Collection",
				"web-app-collection",
			)
			createCollectionOptions.SetDescription("Collection for Web application")
			createCollectionOptions.SetTags("version: 1.1, pre-release")

			collectionLite, response, err := appConfigurationService.CreateCollection(createCollectionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(collectionLite, "", "  ")
			fmt.Println(string(b))

			// end-create_collection

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(collectionLite).ToNot(BeNil())
		})
		It(`UpdateCollection request example`, func() {
			fmt.Println("\nUpdateCollection() result:")
			// begin-update_collection

			updateCollectionOptions := appConfigurationService.NewUpdateCollectionOptions(
				"collection_id",
			)

			collectionLite, response, err := appConfigurationService.UpdateCollection(updateCollectionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(collectionLite, "", "  ")
			fmt.Println(string(b))

			// end-update_collection

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(collectionLite).ToNot(BeNil())
		})
		It(`GetCollection request example`, func() {
			fmt.Println("\nGetCollection() result:")
			// begin-get_collection

			getCollectionOptions := appConfigurationService.NewGetCollectionOptions(
				"collection_id",
			)
			getCollectionOptions.SetExpand(true)
			getCollectionOptions.SetInclude([]string{"features", "properties", "snapshots"})

			collection, response, err := appConfigurationService.GetCollection(getCollectionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(collection, "", "  ")
			fmt.Println(string(b))

			// end-get_collection

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(collection).ToNot(BeNil())
		})
		It(`ListFeatures request example`, func() {
			fmt.Println("\nListFeatures() result:")
			// begin-list_features
			listFeaturesOptions := &appconfigurationv1.ListFeaturesOptions{
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

			pager, err := appConfigurationService.NewFeaturesPager(listFeaturesOptions)
			if err != nil {
				panic(err)
			}

			var allResults []appconfigurationv1.Feature
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_features
		})
		It(`CreateFeature request example`, func() {
			fmt.Println("\nCreateFeature() result:")
			// begin-create_feature

			targetSegmentsModel := &appconfigurationv1.TargetSegments{
				Segments: []string{"betausers", "premiumusers"},
			}

			featureSegmentRuleModel := &appconfigurationv1.FeatureSegmentRule{
				Rules: []appconfigurationv1.TargetSegments{*targetSegmentsModel},
				Value: core.StringPtr("true"),
				Order: core.Int64Ptr(int64(1)),
				RolloutPercentage: core.Int64Ptr(int64(50)),
			}

			collectionRefModel := &appconfigurationv1.CollectionRef{
				CollectionID: core.StringPtr("ghzinc"),
			}

			createFeatureOptions := appConfigurationService.NewCreateFeatureOptions(
				"environment_id",
				"Cycle Rentals",
				"cycle-rentals",
				"BOOLEAN",
				"true",
				"false",
			)
			createFeatureOptions.SetDescription("Feature flag to enable Cycle Rentals")
			createFeatureOptions.SetEnabled(true)
			createFeatureOptions.SetRolloutPercentage(int64(100))
			createFeatureOptions.SetTags("version: 1.1, pre-release")
			createFeatureOptions.SetSegmentRules([]appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel})
			createFeatureOptions.SetCollections([]appconfigurationv1.CollectionRef{*collectionRefModel})

			feature, response, err := appConfigurationService.CreateFeature(createFeatureOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(feature, "", "  ")
			fmt.Println(string(b))

			// end-create_feature

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(feature).ToNot(BeNil())
		})
		It(`UpdateFeature request example`, func() {
			fmt.Println("\nUpdateFeature() result:")
			// begin-update_feature

			targetSegmentsModel := &appconfigurationv1.TargetSegments{
				Segments: []string{"betausers", "premiumusers"},
			}

			featureSegmentRuleModel := &appconfigurationv1.FeatureSegmentRule{
				Rules: []appconfigurationv1.TargetSegments{*targetSegmentsModel},
				Value: core.StringPtr("true"),
				Order: core.Int64Ptr(int64(1)),
				RolloutPercentage: core.Int64Ptr(int64(90)),
			}

			collectionRefModel := &appconfigurationv1.CollectionRef{
				CollectionID: core.StringPtr("ghzinc"),
			}

			updateFeatureOptions := appConfigurationService.NewUpdateFeatureOptions(
				"environment_id",
				"feature_id",
			)
			updateFeatureOptions.SetName("Cycle Rentals")
			updateFeatureOptions.SetDescription("Feature flags to enable Cycle Rentals")
			updateFeatureOptions.SetEnabledValue("true")
			updateFeatureOptions.SetDisabledValue("false")
			updateFeatureOptions.SetEnabled(true)
			updateFeatureOptions.SetRolloutPercentage(int64(100))
			updateFeatureOptions.SetTags("version: 1.1, yet-to-release")
			updateFeatureOptions.SetSegmentRules([]appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel})
			updateFeatureOptions.SetCollections([]appconfigurationv1.CollectionRef{*collectionRefModel})

			feature, response, err := appConfigurationService.UpdateFeature(updateFeatureOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(feature, "", "  ")
			fmt.Println(string(b))

			// end-update_feature

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(feature).ToNot(BeNil())
		})
		It(`UpdateFeatureValues request example`, func() {
			fmt.Println("\nUpdateFeatureValues() result:")
			// begin-update_feature_values

			targetSegmentsModel := &appconfigurationv1.TargetSegments{
				Segments: []string{"betausers", "premiumusers"},
			}

			featureSegmentRuleModel := &appconfigurationv1.FeatureSegmentRule{
				Rules: []appconfigurationv1.TargetSegments{*targetSegmentsModel},
				Value: core.StringPtr("true"),
				Order: core.Int64Ptr(int64(1)),
				RolloutPercentage: core.Int64Ptr(int64(100)),
			}

			updateFeatureValuesOptions := appConfigurationService.NewUpdateFeatureValuesOptions(
				"environment_id",
				"feature_id",
			)
			updateFeatureValuesOptions.SetName("Cycle Rentals")
			updateFeatureValuesOptions.SetDescription("Feature flags to enable Cycle Rentals")
			updateFeatureValuesOptions.SetTags("version: 1.1, yet-to-release")
			updateFeatureValuesOptions.SetEnabledValue("true")
			updateFeatureValuesOptions.SetDisabledValue("false")
			updateFeatureValuesOptions.SetRolloutPercentage(int64(100))
			updateFeatureValuesOptions.SetSegmentRules([]appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel})

			feature, response, err := appConfigurationService.UpdateFeatureValues(updateFeatureValuesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(feature, "", "  ")
			fmt.Println(string(b))

			// end-update_feature_values

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(feature).ToNot(BeNil())
		})
		It(`GetFeature request example`, func() {
			fmt.Println("\nGetFeature() result:")
			// begin-get_feature

			getFeatureOptions := appConfigurationService.NewGetFeatureOptions(
				"environment_id",
				"feature_id",
			)
			getFeatureOptions.SetInclude([]string{"collections", "rules", "change_request"})

			feature, response, err := appConfigurationService.GetFeature(getFeatureOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(feature, "", "  ")
			fmt.Println(string(b))

			// end-get_feature

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(feature).ToNot(BeNil())
		})
		It(`ToggleFeature request example`, func() {
			fmt.Println("\nToggleFeature() result:")
			// begin-toggle_feature

			toggleFeatureOptions := appConfigurationService.NewToggleFeatureOptions(
				"environment_id",
				"feature_id",
				true,
			)

			feature, response, err := appConfigurationService.ToggleFeature(toggleFeatureOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(feature, "", "  ")
			fmt.Println(string(b))

			// end-toggle_feature

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(feature).ToNot(BeNil())
		})
		It(`ListProperties request example`, func() {
			fmt.Println("\nListProperties() result:")
			// begin-list_properties
			listPropertiesOptions := &appconfigurationv1.ListPropertiesOptions{
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

			pager, err := appConfigurationService.NewPropertiesPager(listPropertiesOptions)
			if err != nil {
				panic(err)
			}

			var allResults []appconfigurationv1.Property
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_properties
		})
		It(`CreateProperty request example`, func() {
			fmt.Println("\nCreateProperty() result:")
			// begin-create_property

			targetSegmentsModel := &appconfigurationv1.TargetSegments{
				Segments: []string{"betausers", "premiumusers"},
			}

			segmentRuleModel := &appconfigurationv1.SegmentRule{
				Rules: []appconfigurationv1.TargetSegments{*targetSegmentsModel},
				Value: core.StringPtr("true"),
				Order: core.Int64Ptr(int64(1)),
			}

			collectionRefModel := &appconfigurationv1.CollectionRef{
				CollectionID: core.StringPtr("ghzinc"),
			}

			createPropertyOptions := appConfigurationService.NewCreatePropertyOptions(
				"environment_id",
				"Email property",
				"email-property",
				"BOOLEAN",
				"true",
			)
			createPropertyOptions.SetDescription("Property for email")
			createPropertyOptions.SetTags("version: 1.1, pre-release")
			createPropertyOptions.SetSegmentRules([]appconfigurationv1.SegmentRule{*segmentRuleModel})
			createPropertyOptions.SetCollections([]appconfigurationv1.CollectionRef{*collectionRefModel})

			property, response, err := appConfigurationService.CreateProperty(createPropertyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(property, "", "  ")
			fmt.Println(string(b))

			// end-create_property

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(property).ToNot(BeNil())
		})
		It(`UpdateProperty request example`, func() {
			fmt.Println("\nUpdateProperty() result:")
			// begin-update_property

			targetSegmentsModel := &appconfigurationv1.TargetSegments{
				Segments: []string{"betausers", "premiumusers"},
			}

			segmentRuleModel := &appconfigurationv1.SegmentRule{
				Rules: []appconfigurationv1.TargetSegments{*targetSegmentsModel},
				Value: core.StringPtr("true"),
				Order: core.Int64Ptr(int64(1)),
			}

			collectionRefModel := &appconfigurationv1.CollectionRef{
				CollectionID: core.StringPtr("ghzinc"),
			}

			updatePropertyOptions := appConfigurationService.NewUpdatePropertyOptions(
				"environment_id",
				"property_id",
			)
			updatePropertyOptions.SetName("Email property")
			updatePropertyOptions.SetDescription("Property for email")
			updatePropertyOptions.SetValue("true")
			updatePropertyOptions.SetTags("version: 1.1, pre-release")
			updatePropertyOptions.SetSegmentRules([]appconfigurationv1.SegmentRule{*segmentRuleModel})
			updatePropertyOptions.SetCollections([]appconfigurationv1.CollectionRef{*collectionRefModel})

			property, response, err := appConfigurationService.UpdateProperty(updatePropertyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(property, "", "  ")
			fmt.Println(string(b))

			// end-update_property

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(property).ToNot(BeNil())
		})
		It(`UpdatePropertyValues request example`, func() {
			fmt.Println("\nUpdatePropertyValues() result:")
			// begin-update_property_values

			targetSegmentsModel := &appconfigurationv1.TargetSegments{
				Segments: []string{"betausers", "premiumusers"},
			}

			segmentRuleModel := &appconfigurationv1.SegmentRule{
				Rules: []appconfigurationv1.TargetSegments{*targetSegmentsModel},
				Value: core.StringPtr("true"),
				Order: core.Int64Ptr(int64(1)),
			}

			updatePropertyValuesOptions := appConfigurationService.NewUpdatePropertyValuesOptions(
				"environment_id",
				"property_id",
			)
			updatePropertyValuesOptions.SetName("Email property")
			updatePropertyValuesOptions.SetDescription("Property for email")
			updatePropertyValuesOptions.SetTags("version: 1.1, pre-release")
			updatePropertyValuesOptions.SetValue("true")
			updatePropertyValuesOptions.SetSegmentRules([]appconfigurationv1.SegmentRule{*segmentRuleModel})

			property, response, err := appConfigurationService.UpdatePropertyValues(updatePropertyValuesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(property, "", "  ")
			fmt.Println(string(b))

			// end-update_property_values

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(property).ToNot(BeNil())
		})
		It(`GetProperty request example`, func() {
			fmt.Println("\nGetProperty() result:")
			// begin-get_property

			getPropertyOptions := appConfigurationService.NewGetPropertyOptions(
				"environment_id",
				"property_id",
			)
			getPropertyOptions.SetInclude([]string{"collections", "rules"})

			property, response, err := appConfigurationService.GetProperty(getPropertyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(property, "", "  ")
			fmt.Println(string(b))

			// end-get_property

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(property).ToNot(BeNil())
		})
		It(`ListSegments request example`, func() {
			fmt.Println("\nListSegments() result:")
			// begin-list_segments
			listSegmentsOptions := &appconfigurationv1.ListSegmentsOptions{
				Expand: core.BoolPtr(true),
				Sort: core.StringPtr("created_time"),
				Tags: core.StringPtr("version 1.1,pre-release"),
				Include: core.StringPtr("rules"),
				Limit: core.Int64Ptr(int64(10)),
				Search: core.StringPtr("test tag"),
			}

			pager, err := appConfigurationService.NewSegmentsPager(listSegmentsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []appconfigurationv1.Segment
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_segments
		})
		It(`CreateSegment request example`, func() {
			fmt.Println("\nCreateSegment() result:")
			// begin-create_segment

			ruleModel := &appconfigurationv1.Rule{
				AttributeName: core.StringPtr("email"),
				Operator: core.StringPtr("endsWith"),
				Values: []string{"@in.mnc.com", "@us.mnc.com"},
			}

			createSegmentOptions := appConfigurationService.NewCreateSegmentOptions(
				"Beta Users",
				"beta-users",
				[]appconfigurationv1.Rule{*ruleModel},
			)
			createSegmentOptions.SetDescription("Segment containing the beta users")
			createSegmentOptions.SetTags("version: 1.1, stage")

			segment, response, err := appConfigurationService.CreateSegment(createSegmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(segment, "", "  ")
			fmt.Println(string(b))

			// end-create_segment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(segment).ToNot(BeNil())
		})
		It(`UpdateSegment request example`, func() {
			fmt.Println("\nUpdateSegment() result:")
			// begin-update_segment

			updateSegmentOptions := appConfigurationService.NewUpdateSegmentOptions(
				"segment_id",
			)

			segment, response, err := appConfigurationService.UpdateSegment(updateSegmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(segment, "", "  ")
			fmt.Println(string(b))

			// end-update_segment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(segment).ToNot(BeNil())
		})
		It(`GetSegment request example`, func() {
			fmt.Println("\nGetSegment() result:")
			// begin-get_segment

			getSegmentOptions := appConfigurationService.NewGetSegmentOptions(
				"segment_id",
			)
			getSegmentOptions.SetInclude([]string{"features", "properties"})

			segment, response, err := appConfigurationService.GetSegment(getSegmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(segment, "", "  ")
			fmt.Println(string(b))

			// end-get_segment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(segment).ToNot(BeNil())
		})
		It(`ListSnapshots request example`, func() {
			fmt.Println("\nListSnapshots() result:")
			// begin-list_snapshots
			listSnapshotsOptions := &appconfigurationv1.ListSnapshotsOptions{
				Sort: core.StringPtr("created_time"),
				CollectionID: core.StringPtr("collection_id"),
				EnvironmentID: core.StringPtr("environment_id"),
				Limit: core.Int64Ptr(int64(10)),
				Search: core.StringPtr("search_string"),
			}

			pager, err := appConfigurationService.NewSnapshotsPager(listSnapshotsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []appconfigurationv1.GitConfig
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_snapshots
		})
		It(`CreateGitconfig request example`, func() {
			fmt.Println("\nCreateGitconfig() result:")
			// begin-create_gitconfig

			createGitconfigOptions := appConfigurationService.NewCreateGitconfigOptions(
				"boot-strap-configuration",
				"boot-strap-configuration",
				"web-app-collection",
				"dev",
				"https://github.ibm.com/api/v3/repos/jhondoe-owner/my-test-repo",
				"main",
				"code/development/README.json",
				"61a792eahhGHji223jijb55a6cfdd4d5cde4c8a67esjjhjhHVH",
			)

			createGitConfigResponse, response, err := appConfigurationService.CreateGitconfig(createGitconfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createGitConfigResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_gitconfig

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createGitConfigResponse).ToNot(BeNil())
		})
		It(`UpdateGitconfig request example`, func() {
			fmt.Println("\nUpdateGitconfig() result:")
			// begin-update_gitconfig

			updateGitconfigOptions := appConfigurationService.NewUpdateGitconfigOptions(
				"git_config_id",
			)

			gitConfig, response, err := appConfigurationService.UpdateGitconfig(updateGitconfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(gitConfig, "", "  ")
			fmt.Println(string(b))

			// end-update_gitconfig

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(gitConfig).ToNot(BeNil())
		})
		It(`GetGitconfig request example`, func() {
			fmt.Println("\nGetGitconfig() result:")
			// begin-get_gitconfig

			getGitconfigOptions := appConfigurationService.NewGetGitconfigOptions(
				"git_config_id",
			)

			gitConfig, response, err := appConfigurationService.GetGitconfig(getGitconfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(gitConfig, "", "  ")
			fmt.Println(string(b))

			// end-get_gitconfig

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(gitConfig).ToNot(BeNil())
		})
		It(`PromoteGitconfig request example`, func() {
			fmt.Println("\nPromoteGitconfig() result:")
			// begin-promote_gitconfig

			promoteGitconfigOptions := appConfigurationService.NewPromoteGitconfigOptions(
				"git_config_id",
			)

			gitConfigPromote, response, err := appConfigurationService.PromoteGitconfig(promoteGitconfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(gitConfigPromote, "", "  ")
			fmt.Println(string(b))

			// end-promote_gitconfig

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(gitConfigPromote).ToNot(BeNil())
		})
		It(`RestoreGitconfig request example`, func() {
			fmt.Println("\nRestoreGitconfig() result:")
			// begin-restore_gitconfig

			restoreGitconfigOptions := appConfigurationService.NewRestoreGitconfigOptions(
				"git_config_id",
			)

			gitConfigRestore, response, err := appConfigurationService.RestoreGitconfig(restoreGitconfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(gitConfigRestore, "", "  ")
			fmt.Println(string(b))

			// end-restore_gitconfig

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(gitConfigRestore).ToNot(BeNil())
		})
		It(`ListOriginconfigs request example`, func() {
			fmt.Println("\nListOriginconfigs() result:")
			// begin-list_originconfigs

			listOriginconfigsOptions := appConfigurationService.NewListOriginconfigsOptions()

			originConfigList, response, err := appConfigurationService.ListOriginconfigs(listOriginconfigsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(originConfigList, "", "  ")
			fmt.Println(string(b))

			// end-list_originconfigs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(originConfigList).ToNot(BeNil())
		})
		It(`UpdateOriginconfigs request example`, func() {
			fmt.Println("\nUpdateOriginconfigs() result:")
			// begin-update_originconfigs

			updateOriginconfigsOptions := appConfigurationService.NewUpdateOriginconfigsOptions(
				[]string{"testString"},
			)

			originConfigList, response, err := appConfigurationService.UpdateOriginconfigs(updateOriginconfigsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(originConfigList, "", "  ")
			fmt.Println(string(b))

			// end-update_originconfigs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(originConfigList).ToNot(BeNil())
		})
		It(`ListWorkflowconfig request example`, func() {
			fmt.Println("\nListWorkflowconfig() result:")
			// begin-list_workflowconfig

			listWorkflowconfigOptions := appConfigurationService.NewListWorkflowconfigOptions(
				"environment_id",
			)

			listWorkflowconfigResponse, response, err := appConfigurationService.ListWorkflowconfig(listWorkflowconfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(listWorkflowconfigResponse, "", "  ")
			fmt.Println(string(b))

			// end-list_workflowconfig

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listWorkflowconfigResponse).ToNot(BeNil())
		})
		It(`CreateWorkflowconfig request example`, func() {
			fmt.Println("\nCreateWorkflowconfig() result:")
			// begin-create_Workflowconfig

			workflowCredentialsModel := &appconfigurationv1.WorkflowCredentials{
				Username: core.StringPtr("user"),
				Password: core.StringPtr("pwd"),
				ClientID: core.StringPtr("client id value"),
				ClientSecret: core.StringPtr("clientsecret"),
			}

			createWorkflowconfigRequestModel := &appconfigurationv1.CreateWorkflowconfigRequestWorkflowConfig{
				WorkflowURL: core.StringPtr("https://xxxxx.service-now.com"),
				ApprovalGroupName: core.StringPtr("WorkflowCRApprovers"),
				ApprovalExpiration: core.Int64Ptr(int64(10)),
				WorkflowCredentials: workflowCredentialsModel,
				Enabled: core.BoolPtr(true),
			}

			createWorkflowconfigOptions := appConfigurationService.NewCreateWorkflowconfigOptions(
				"environment_id",
				createWorkflowconfigRequestModel,
			)

			createWorkflowconfigResponse, response, err := appConfigurationService.CreateWorkflowconfig(createWorkflowconfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createWorkflowconfigResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_Workflowconfig

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createWorkflowconfigResponse).ToNot(BeNil())
		})
		It(`UpdateWorkflowconfig request example`, func() {
			fmt.Println("\nUpdateWorkflowconfig() result:")
			// begin-update_Workflowconfig

			workflowCredentialsModel := &appconfigurationv1.WorkflowCredentials{
				Username: core.StringPtr("user"),
				Password: core.StringPtr("updated password"),
				ClientID: core.StringPtr("client id value"),
				ClientSecret: core.StringPtr("updated client secret"),
			}

			updateWorkflowconfigRequestModel := &appconfigurationv1.UpdateWorkflowconfigRequestUpdateWorkflowConfig{
				WorkflowURL: core.StringPtr("https://xxxxx.service-now.com"),
				ApprovalGroupName: core.StringPtr("WorkflowCRApprovers"),
				ApprovalExpiration: core.Int64Ptr(int64(5)),
				WorkflowCredentials: workflowCredentialsModel,
				Enabled: core.BoolPtr(true),
			}

			updateWorkflowconfigOptions := appConfigurationService.NewUpdateWorkflowconfigOptions(
				"environment_id",
				updateWorkflowconfigRequestModel,
			)

			updateWorkflowconfigResponse, response, err := appConfigurationService.UpdateWorkflowconfig(updateWorkflowconfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(updateWorkflowconfigResponse, "", "  ")
			fmt.Println(string(b))

			// end-update_Workflowconfig

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(updateWorkflowconfigResponse).ToNot(BeNil())
		})
		It(`ImportConfig request example`, func() {
			fmt.Println("\nImportConfig() result:")
			// begin-import_config

			targetSegmentsModel := &appconfigurationv1.TargetSegments{
				Segments: []string{"testString"},
			}

			featureSegmentRuleModel := &appconfigurationv1.FeatureSegmentRule{
				Rules: []appconfigurationv1.TargetSegments{*targetSegmentsModel},
				Value: "testString",
				Order: core.Int64Ptr(int64(38)),
			}

			collectionRefModel := &appconfigurationv1.CollectionRef{
				CollectionID: core.StringPtr("web-app"),
			}

			importFeatureRequestBodyModel := &appconfigurationv1.ImportFeatureRequestBody{
				Name: core.StringPtr("Cycle Rentals"),
				FeatureID: core.StringPtr("cycle-rentals"),
				Type: core.StringPtr("NUMERIC"),
				EnabledValue: core.StringPtr("1"),
				DisabledValue: core.StringPtr("2"),
				Enabled: core.BoolPtr(true),
				RolloutPercentage: core.Int64Ptr(int64(100)),
				SegmentRules: []appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel},
				Collections: []appconfigurationv1.CollectionRef{*collectionRefModel},
				IsOverridden: core.BoolPtr(true),
			}

			segmentRuleModel := &appconfigurationv1.SegmentRule{
				Rules: []appconfigurationv1.TargetSegments{*targetSegmentsModel},
				Value: core.StringPtr("200"),
				Order: core.Int64Ptr(int64(1)),
			}

			importPropertyRequestBodyModel := &appconfigurationv1.ImportPropertyRequestBody{
				Name: core.StringPtr("Daily Discount"),
				PropertyID: core.StringPtr("daily_discount"),
				Type: core.StringPtr("NUMERIC"),
				Value: core.StringPtr("100"),
				Tags: core.StringPtr("pre-release, v1.2"),
				SegmentRules: []appconfigurationv1.SegmentRule{*segmentRuleModel},
				Collections: []appconfigurationv1.CollectionRef{*collectionRefModel},
				IsOverridden: core.BoolPtr(true),
			}

			importEnvironmentSchemaModel := &appconfigurationv1.ImportEnvironmentSchema{
				Name: core.StringPtr("Dev"),
				EnvironmentID: core.StringPtr("dev"),
				Description: core.StringPtr("Environment created on instance creation"),
				ColorCode: core.StringPtr("#FDD13A"),
				Features: []appconfigurationv1.ImportFeatureRequestBody{*importFeatureRequestBodyModel},
				Properties: []appconfigurationv1.ImportPropertyRequestBody{*importPropertyRequestBodyModel},
			}

			importCollectionSchemaModel := &appconfigurationv1.ImportCollectionSchema{
				CollectionID: core.StringPtr("web-app"),
				Name: core.StringPtr("web-app"),
				Description: core.StringPtr("web app collection"),
				Tags: core.StringPtr("v1"),
			}

			ruleModel := &appconfigurationv1.Rule{
				AttributeName: core.StringPtr("email"),
				Operator: core.StringPtr("is"),
				Values: []string{"john@bluecharge.com", "alice@bluecharge.com"},
			}

			importSegmentSchemaModel := &appconfigurationv1.ImportSegmentSchema{
				Name: core.StringPtr("Testers"),
				SegmentID: core.StringPtr("khpwj68h"),
				Description: core.StringPtr("Testers"),
				Tags: core.StringPtr("test"),
				Rules: []appconfigurationv1.Rule{*ruleModel},
			}

			importConfigOptions := appConfigurationService.NewImportConfigOptions()
			importConfigOptions.SetEnvironments([]appconfigurationv1.ImportEnvironmentSchema{*importEnvironmentSchemaModel})
			importConfigOptions.SetCollections([]appconfigurationv1.ImportCollectionSchema{*importCollectionSchemaModel})
			importConfigOptions.SetSegments([]appconfigurationv1.ImportSegmentSchema{*importSegmentSchemaModel})
			importConfigOptions.SetClean("true")

			importConfig, response, err := appConfigurationService.ImportConfig(importConfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(importConfig, "", "  ")
			fmt.Println(string(b))

			// end-import_config

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(importConfig).ToNot(BeNil())
		})
		It(`ListInstanceConfig request example`, func() {
			fmt.Println("\nListInstanceConfig() result:")
			// begin-list_instance_config

			listInstanceConfigOptions := appConfigurationService.NewListInstanceConfigOptions()

			importConfig, response, err := appConfigurationService.ListInstanceConfig(listInstanceConfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(importConfig, "", "  ")
			fmt.Println(string(b))

			// end-list_instance_config

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(importConfig).ToNot(BeNil())
		})
		It(`PromoteRestoreConfig request example`, func() {
			fmt.Println("\nPromoteRestoreConfig() result:")
			// begin-promote_restore_config

			promoteRestoreConfigOptions := appConfigurationService.NewPromoteRestoreConfigOptions(
				"git_config_id",
				"promote",
			)

			configAction, response, err := appConfigurationService.PromoteRestoreConfig(promoteRestoreConfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(configAction, "", "  ")
			fmt.Println(string(b))

			// end-promote_restore_config

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(configAction).ToNot(BeNil())
		})
		It(`DeleteEnvironment request example`, func() {
			// begin-delete_environment

			deleteEnvironmentOptions := appConfigurationService.NewDeleteEnvironmentOptions(
				"environment_id",
			)

			response, err := appConfigurationService.DeleteEnvironment(deleteEnvironmentOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteEnvironment(): %d\n", response.StatusCode)
			}

			// end-delete_environment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteCollection request example`, func() {
			// begin-delete_collection

			deleteCollectionOptions := appConfigurationService.NewDeleteCollectionOptions(
				"collection_id",
			)

			response, err := appConfigurationService.DeleteCollection(deleteCollectionOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteCollection(): %d\n", response.StatusCode)
			}

			// end-delete_collection

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteFeature request example`, func() {
			// begin-delete_feature

			deleteFeatureOptions := appConfigurationService.NewDeleteFeatureOptions(
				"environment_id",
				"feature_id",
			)

			response, err := appConfigurationService.DeleteFeature(deleteFeatureOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteFeature(): %d\n", response.StatusCode)
			}

			// end-delete_feature

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteProperty request example`, func() {
			// begin-delete_property

			deletePropertyOptions := appConfigurationService.NewDeletePropertyOptions(
				"environment_id",
				"property_id",
			)

			response, err := appConfigurationService.DeleteProperty(deletePropertyOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteProperty(): %d\n", response.StatusCode)
			}

			// end-delete_property

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteSegment request example`, func() {
			// begin-delete_segment

			deleteSegmentOptions := appConfigurationService.NewDeleteSegmentOptions(
				"segment_id",
			)

			response, err := appConfigurationService.DeleteSegment(deleteSegmentOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteSegment(): %d\n", response.StatusCode)
			}

			// end-delete_segment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteGitconfig request example`, func() {
			// begin-delete_gitconfig

			deleteGitconfigOptions := appConfigurationService.NewDeleteGitconfigOptions(
				"git_config_id",
			)

			response, err := appConfigurationService.DeleteGitconfig(deleteGitconfigOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteGitconfig(): %d\n", response.StatusCode)
			}

			// end-delete_gitconfig

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteWorkflowconfig request example`, func() {
			// begin-delete_workflowconfig

			deleteWorkflowconfigOptions := appConfigurationService.NewDeleteWorkflowconfigOptions(
				"environment_id",
			)

			response, err := appConfigurationService.DeleteWorkflowconfig(deleteWorkflowconfigOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteWorkflowconfig(): %d\n", response.StatusCode)
			}

			// end-delete_workflowconfig

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})
