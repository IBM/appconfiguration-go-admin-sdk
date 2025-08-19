//go:build integration

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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the appconfigurationv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`AppConfigurationV1 Integration Tests`, func() {
	const externalConfigFile = "../app_configuration_v1.env"

	var (
		err          error
		appConfigurationService *appconfigurationv1.AppConfigurationV1
		serviceURL   string
		config       map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(appconfigurationv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			appConfigurationServiceOptions := &appconfigurationv1.AppConfigurationV1Options{}

			appConfigurationService, err = appconfigurationv1.NewAppConfigurationV1UsingExternalConfig(appConfigurationServiceOptions)
			Expect(err).To(BeNil())
			Expect(appConfigurationService).ToNot(BeNil())
			Expect(appConfigurationService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			appConfigurationService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`ListEnvironments - Get list of Environments`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListEnvironments(listEnvironmentsOptions *ListEnvironmentsOptions) with pagination`, func(){
			listEnvironmentsOptions := &appconfigurationv1.ListEnvironmentsOptions{
				Expand: core.BoolPtr(true),
				Sort: core.StringPtr("created_time"),
				Tags: core.StringPtr("version 1.1,pre-release"),
				Include: []string{"features", "properties", "snapshots"},
				Limit: core.Int64Ptr(int64(10)),
				Offset: core.Int64Ptr(int64(0)),
				Search: core.StringPtr("test tag"),
			}

			listEnvironmentsOptions.Offset = nil
			listEnvironmentsOptions.Limit = core.Int64Ptr(1)

			var allResults []appconfigurationv1.Environment
			for {
				environmentList, response, err := appConfigurationService.ListEnvironments(listEnvironmentsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(environmentList).ToNot(BeNil())
				allResults = append(allResults, environmentList.Environments...)

				listEnvironmentsOptions.Offset, err = environmentList.GetNextOffset()
				Expect(err).To(BeNil())

				if listEnvironmentsOptions.Offset == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListEnvironments(listEnvironmentsOptions *ListEnvironmentsOptions) using EnvironmentsPager`, func(){
			listEnvironmentsOptions := &appconfigurationv1.ListEnvironmentsOptions{
				Expand: core.BoolPtr(true),
				Sort: core.StringPtr("created_time"),
				Tags: core.StringPtr("version 1.1,pre-release"),
				Include: []string{"features", "properties", "snapshots"},
				Limit: core.Int64Ptr(int64(10)),
				Search: core.StringPtr("test tag"),
			}

			// Test GetNext().
			pager, err := appConfigurationService.NewEnvironmentsPager(listEnvironmentsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []appconfigurationv1.Environment
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = appConfigurationService.NewEnvironmentsPager(listEnvironmentsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListEnvironments() returned a total of %d item(s) using EnvironmentsPager.\n", len(allResults))
		})
	})

	Describe(`CreateEnvironment - Create Environment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateEnvironment(createEnvironmentOptions *CreateEnvironmentOptions)`, func() {
			createEnvironmentOptions := &appconfigurationv1.CreateEnvironmentOptions{
				Name: core.StringPtr("Dev environment"),
				EnvironmentID: core.StringPtr("dev-environment"),
				Description: core.StringPtr("Dev environment description"),
				Tags: core.StringPtr("development"),
				ColorCode: core.StringPtr("#FDD13A"),
			}

			environment, response, err := appConfigurationService.CreateEnvironment(createEnvironmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(environment).ToNot(BeNil())
		})
	})

	Describe(`UpdateEnvironment - Update Environment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateEnvironment(updateEnvironmentOptions *UpdateEnvironmentOptions)`, func() {
			updateEnvironmentOptions := &appconfigurationv1.UpdateEnvironmentOptions{
				EnvironmentID: core.StringPtr("environment_id"),
				Name: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Tags: core.StringPtr("testString"),
				ColorCode: core.StringPtr("#FDD13A"),
			}

			environment, response, err := appConfigurationService.UpdateEnvironment(updateEnvironmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(environment).ToNot(BeNil())
		})
	})

	Describe(`GetEnvironment - Get Environment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetEnvironment(getEnvironmentOptions *GetEnvironmentOptions)`, func() {
			getEnvironmentOptions := &appconfigurationv1.GetEnvironmentOptions{
				EnvironmentID: core.StringPtr("environment_id"),
				Expand: core.BoolPtr(true),
				Include: []string{"features", "properties", "snapshots"},
			}

			environment, response, err := appConfigurationService.GetEnvironment(getEnvironmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(environment).ToNot(BeNil())
		})
	})

	Describe(`ListCollections - Get list of Collections`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListCollections(listCollectionsOptions *ListCollectionsOptions) with pagination`, func(){
			listCollectionsOptions := &appconfigurationv1.ListCollectionsOptions{
				Expand: core.BoolPtr(true),
				Sort: core.StringPtr("created_time"),
				Tags: core.StringPtr("version 1.1,pre-release"),
				Features: []string{"my-feature-id", "cycle-rentals"},
				Properties: []string{"my-property-id", "email-property"},
				Include: []string{"features", "properties", "snapshots"},
				Limit: core.Int64Ptr(int64(10)),
				Offset: core.Int64Ptr(int64(0)),
				Search: core.StringPtr("test tag"),
			}

			listCollectionsOptions.Offset = nil
			listCollectionsOptions.Limit = core.Int64Ptr(1)

			var allResults []appconfigurationv1.Collection
			for {
				collectionList, response, err := appConfigurationService.ListCollections(listCollectionsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(collectionList).ToNot(BeNil())
				allResults = append(allResults, collectionList.Collections...)

				listCollectionsOptions.Offset, err = collectionList.GetNextOffset()
				Expect(err).To(BeNil())

				if listCollectionsOptions.Offset == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListCollections(listCollectionsOptions *ListCollectionsOptions) using CollectionsPager`, func(){
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

			// Test GetNext().
			pager, err := appConfigurationService.NewCollectionsPager(listCollectionsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []appconfigurationv1.Collection
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = appConfigurationService.NewCollectionsPager(listCollectionsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListCollections() returned a total of %d item(s) using CollectionsPager.\n", len(allResults))
		})
	})

	Describe(`CreateCollection - Create Collection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateCollection(createCollectionOptions *CreateCollectionOptions)`, func() {
			createCollectionOptions := &appconfigurationv1.CreateCollectionOptions{
				Name: core.StringPtr("Web App Collection"),
				CollectionID: core.StringPtr("web-app-collection"),
				Description: core.StringPtr("Collection for Web application"),
				Tags: core.StringPtr("version: 1.1, pre-release"),
			}

			collectionLite, response, err := appConfigurationService.CreateCollection(createCollectionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(collectionLite).ToNot(BeNil())
		})
	})

	Describe(`UpdateCollection - Update Collection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateCollection(updateCollectionOptions *UpdateCollectionOptions)`, func() {
			updateCollectionOptions := &appconfigurationv1.UpdateCollectionOptions{
				CollectionID: core.StringPtr("collection_id"),
				Name: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Tags: core.StringPtr("testString"),
			}

			collectionLite, response, err := appConfigurationService.UpdateCollection(updateCollectionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(collectionLite).ToNot(BeNil())
		})
	})

	Describe(`GetCollection - Get Collection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCollection(getCollectionOptions *GetCollectionOptions)`, func() {
			getCollectionOptions := &appconfigurationv1.GetCollectionOptions{
				CollectionID: core.StringPtr("collection_id"),
				Expand: core.BoolPtr(true),
				Include: []string{"features", "properties", "snapshots"},
			}

			collection, response, err := appConfigurationService.GetCollection(getCollectionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(collection).ToNot(BeNil())
		})
	})

	Describe(`ListFeatures - Get list of Features`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListFeatures(listFeaturesOptions *ListFeaturesOptions) with pagination`, func(){
			listFeaturesOptions := &appconfigurationv1.ListFeaturesOptions{
				EnvironmentID: core.StringPtr("environment_id"),
				Expand: core.BoolPtr(true),
				Sort: core.StringPtr("created_time"),
				Tags: core.StringPtr("version 1.1,pre-release"),
				Collections: []string{"my-collection-id", "ghzindiapvtltd"},
				Segments: []string{"my-segment-id", "beta-users"},
				Include: []string{"collections", "rules", "change_request"},
				Limit: core.Int64Ptr(int64(10)),
				Offset: core.Int64Ptr(int64(0)),
				Search: core.StringPtr("test tag"),
			}

			listFeaturesOptions.Offset = nil
			listFeaturesOptions.Limit = core.Int64Ptr(1)

			var allResults []appconfigurationv1.Feature
			for {
				featuresList, response, err := appConfigurationService.ListFeatures(listFeaturesOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(featuresList).ToNot(BeNil())
				allResults = append(allResults, featuresList.Features...)

				listFeaturesOptions.Offset, err = featuresList.GetNextOffset()
				Expect(err).To(BeNil())

				if listFeaturesOptions.Offset == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListFeatures(listFeaturesOptions *ListFeaturesOptions) using FeaturesPager`, func(){
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

			// Test GetNext().
			pager, err := appConfigurationService.NewFeaturesPager(listFeaturesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []appconfigurationv1.Feature
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = appConfigurationService.NewFeaturesPager(listFeaturesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListFeatures() returned a total of %d item(s) using FeaturesPager.\n", len(allResults))
		})
	})

	Describe(`CreateFeature - Create Feature`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateFeature(createFeatureOptions *CreateFeatureOptions)`, func() {
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

			createFeatureOptions := &appconfigurationv1.CreateFeatureOptions{
				EnvironmentID: core.StringPtr("environment_id"),
				Name: core.StringPtr("Cycle Rentals"),
				FeatureID: core.StringPtr("cycle-rentals"),
				Type: core.StringPtr("BOOLEAN"),
				EnabledValue: core.StringPtr("true"),
				DisabledValue: core.StringPtr("false"),
				Description: core.StringPtr("Feature flag to enable Cycle Rentals"),
				Format: core.StringPtr("TEXT"),
				Enabled: core.BoolPtr(true),
				RolloutPercentage: core.Int64Ptr(int64(100)),
				Tags: core.StringPtr("version: 1.1, pre-release"),
				SegmentRules: []appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel},
				Collections: []appconfigurationv1.CollectionRef{*collectionRefModel},
			}

			feature, response, err := appConfigurationService.CreateFeature(createFeatureOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(feature).ToNot(BeNil())
		})
	})

	Describe(`UpdateFeature - Update Feature`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateFeature(updateFeatureOptions *UpdateFeatureOptions)`, func() {
			targetSegmentsModel := &appconfigurationv1.TargetSegments{
				Segments: []string{"betausers", "premiumusers"},
			}

			featureSegmentRuleModel := &appconfigurationv1.FeatureSegmentRule{
				Rules: []appconfigurationv1.TargetSegments{*targetSegmentsModel},
				Value: core.StringPtr("true"),
				Order: core.Int64Ptr(int64(1)),
				RolloutPercentage: core.Int64Ptr(int64(90)),
			}

			collectionUpdateRefModel := &appconfigurationv1.CollectionUpdateRef{
				CollectionID: core.StringPtr("ghzinc"),
				Deleted: core.BoolPtr(true),
			}

			updateFeatureOptions := &appconfigurationv1.UpdateFeatureOptions{
				EnvironmentID: core.StringPtr("environment_id"),
				FeatureID: core.StringPtr("feature_id"),
				Name: core.StringPtr("Cycle Rentals"),
				Description: core.StringPtr("Feature flags to enable Cycle Rentals"),
				EnabledValue: core.StringPtr("true"),
				DisabledValue: core.StringPtr("false"),
				Enabled: core.BoolPtr(true),
				RolloutPercentage: core.Int64Ptr(int64(100)),
				Tags: core.StringPtr("version: 1.1, yet-to-release"),
				SegmentRules: []appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel},
				Collections: []appconfigurationv1.CollectionUpdateRef{*collectionUpdateRefModel},
			}

			feature, response, err := appConfigurationService.UpdateFeature(updateFeatureOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(feature).ToNot(BeNil())
		})
	})

	Describe(`UpdateFeatureValues - Update Feature Values`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateFeatureValues(updateFeatureValuesOptions *UpdateFeatureValuesOptions)`, func() {
			targetSegmentsModel := &appconfigurationv1.TargetSegments{
				Segments: []string{"betausers", "premiumusers"},
			}

			featureSegmentRuleModel := &appconfigurationv1.FeatureSegmentRule{
				Rules: []appconfigurationv1.TargetSegments{*targetSegmentsModel},
				Value: core.StringPtr("true"),
				Order: core.Int64Ptr(int64(1)),
				RolloutPercentage: core.Int64Ptr(int64(100)),
			}

			updateFeatureValuesOptions := &appconfigurationv1.UpdateFeatureValuesOptions{
				EnvironmentID: core.StringPtr("environment_id"),
				FeatureID: core.StringPtr("feature_id"),
				Name: core.StringPtr("Cycle Rentals"),
				Description: core.StringPtr("Feature flags to enable Cycle Rentals"),
				Tags: core.StringPtr("version: 1.1, yet-to-release"),
				EnabledValue: core.StringPtr("true"),
				DisabledValue: core.StringPtr("false"),
				RolloutPercentage: core.Int64Ptr(int64(100)),
				SegmentRules: []appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel},
			}

			feature, response, err := appConfigurationService.UpdateFeatureValues(updateFeatureValuesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(feature).ToNot(BeNil())
		})
	})

	Describe(`GetFeature - Get Feature`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetFeature(getFeatureOptions *GetFeatureOptions)`, func() {
			getFeatureOptions := &appconfigurationv1.GetFeatureOptions{
				EnvironmentID: core.StringPtr("environment_id"),
				FeatureID: core.StringPtr("feature_id"),
				Include: []string{"collections", "rules", "change_request"},
			}

			feature, response, err := appConfigurationService.GetFeature(getFeatureOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(feature).ToNot(BeNil())
		})
	})

	Describe(`ToggleFeature - Toggle Feature`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ToggleFeature(toggleFeatureOptions *ToggleFeatureOptions)`, func() {
			toggleFeatureOptions := &appconfigurationv1.ToggleFeatureOptions{
				EnvironmentID: core.StringPtr("environment_id"),
				FeatureID: core.StringPtr("feature_id"),
				Enabled: core.BoolPtr(true),
			}

			feature, response, err := appConfigurationService.ToggleFeature(toggleFeatureOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(feature).ToNot(BeNil())
		})
	})

	Describe(`ListProperties - Get list of Properties`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProperties(listPropertiesOptions *ListPropertiesOptions) with pagination`, func(){
			listPropertiesOptions := &appconfigurationv1.ListPropertiesOptions{
				EnvironmentID: core.StringPtr("environment_id"),
				Expand: core.BoolPtr(true),
				Sort: core.StringPtr("created_time"),
				Tags: core.StringPtr("version 1.1,pre-release"),
				Collections: []string{"my-collection-id", "ghzindiapvtltd"},
				Segments: []string{"my-segment-id", "beta-users"},
				Include: []string{"collections", "rules"},
				Limit: core.Int64Ptr(int64(10)),
				Offset: core.Int64Ptr(int64(0)),
				Search: core.StringPtr("test tag"),
			}

			listPropertiesOptions.Offset = nil
			listPropertiesOptions.Limit = core.Int64Ptr(1)

			var allResults []appconfigurationv1.Property
			for {
				propertiesList, response, err := appConfigurationService.ListProperties(listPropertiesOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(propertiesList).ToNot(BeNil())
				allResults = append(allResults, propertiesList.Properties...)

				listPropertiesOptions.Offset, err = propertiesList.GetNextOffset()
				Expect(err).To(BeNil())

				if listPropertiesOptions.Offset == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListProperties(listPropertiesOptions *ListPropertiesOptions) using PropertiesPager`, func(){
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

			// Test GetNext().
			pager, err := appConfigurationService.NewPropertiesPager(listPropertiesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []appconfigurationv1.Property
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = appConfigurationService.NewPropertiesPager(listPropertiesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListProperties() returned a total of %d item(s) using PropertiesPager.\n", len(allResults))
		})
	})

	Describe(`CreateProperty - Create Property`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateProperty(createPropertyOptions *CreatePropertyOptions)`, func() {
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

			createPropertyOptions := &appconfigurationv1.CreatePropertyOptions{
				EnvironmentID: core.StringPtr("environment_id"),
				Name: core.StringPtr("Email property"),
				PropertyID: core.StringPtr("email-property"),
				Type: core.StringPtr("BOOLEAN"),
				Value: core.StringPtr("true"),
				Description: core.StringPtr("Property for email"),
				Format: core.StringPtr("TEXT"),
				Tags: core.StringPtr("version: 1.1, pre-release"),
				SegmentRules: []appconfigurationv1.SegmentRule{*segmentRuleModel},
				Collections: []appconfigurationv1.CollectionRef{*collectionRefModel},
			}

			property, response, err := appConfigurationService.CreateProperty(createPropertyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(property).ToNot(BeNil())
		})
	})

	Describe(`UpdateProperty - Update Property`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateProperty(updatePropertyOptions *UpdatePropertyOptions)`, func() {
			targetSegmentsModel := &appconfigurationv1.TargetSegments{
				Segments: []string{"betausers", "premiumusers"},
			}

			segmentRuleModel := &appconfigurationv1.SegmentRule{
				Rules: []appconfigurationv1.TargetSegments{*targetSegmentsModel},
				Value: core.StringPtr("true"),
				Order: core.Int64Ptr(int64(1)),
			}

			collectionUpdateRefModel := &appconfigurationv1.CollectionUpdateRef{
				CollectionID: core.StringPtr("ghzinc"),
				Deleted: core.BoolPtr(true),
			}

			updatePropertyOptions := &appconfigurationv1.UpdatePropertyOptions{
				EnvironmentID: core.StringPtr("environment_id"),
				PropertyID: core.StringPtr("property_id"),
				Name: core.StringPtr("Email property"),
				Description: core.StringPtr("Property for email"),
				Value: core.StringPtr("true"),
				Tags: core.StringPtr("version: 1.1, pre-release"),
				SegmentRules: []appconfigurationv1.SegmentRule{*segmentRuleModel},
				Collections: []appconfigurationv1.CollectionUpdateRef{*collectionUpdateRefModel},
			}

			property, response, err := appConfigurationService.UpdateProperty(updatePropertyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(property).ToNot(BeNil())
		})
	})

	Describe(`UpdatePropertyValues - Update Property values`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdatePropertyValues(updatePropertyValuesOptions *UpdatePropertyValuesOptions)`, func() {
			targetSegmentsModel := &appconfigurationv1.TargetSegments{
				Segments: []string{"betausers", "premiumusers"},
			}

			segmentRuleModel := &appconfigurationv1.SegmentRule{
				Rules: []appconfigurationv1.TargetSegments{*targetSegmentsModel},
				Value: core.StringPtr("true"),
				Order: core.Int64Ptr(int64(1)),
			}

			updatePropertyValuesOptions := &appconfigurationv1.UpdatePropertyValuesOptions{
				EnvironmentID: core.StringPtr("environment_id"),
				PropertyID: core.StringPtr("property_id"),
				Name: core.StringPtr("Email property"),
				Description: core.StringPtr("Property for email"),
				Tags: core.StringPtr("version: 1.1, pre-release"),
				Value: core.StringPtr("true"),
				SegmentRules: []appconfigurationv1.SegmentRule{*segmentRuleModel},
			}

			property, response, err := appConfigurationService.UpdatePropertyValues(updatePropertyValuesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(property).ToNot(BeNil())
		})
	})

	Describe(`GetProperty - Get Property`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProperty(getPropertyOptions *GetPropertyOptions)`, func() {
			getPropertyOptions := &appconfigurationv1.GetPropertyOptions{
				EnvironmentID: core.StringPtr("environment_id"),
				PropertyID: core.StringPtr("property_id"),
				Include: []string{"collections", "rules"},
			}

			property, response, err := appConfigurationService.GetProperty(getPropertyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(property).ToNot(BeNil())
		})
	})

	Describe(`ListSegments - Get list of Segments`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSegments(listSegmentsOptions *ListSegmentsOptions) with pagination`, func(){
			listSegmentsOptions := &appconfigurationv1.ListSegmentsOptions{
				Expand: core.BoolPtr(true),
				Sort: core.StringPtr("created_time"),
				Tags: core.StringPtr("version 1.1,pre-release"),
				Include: core.StringPtr("rules"),
				Limit: core.Int64Ptr(int64(10)),
				Offset: core.Int64Ptr(int64(0)),
				Search: core.StringPtr("test tag"),
			}

			listSegmentsOptions.Offset = nil
			listSegmentsOptions.Limit = core.Int64Ptr(1)

			var allResults []appconfigurationv1.Segment
			for {
				segmentsList, response, err := appConfigurationService.ListSegments(listSegmentsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(segmentsList).ToNot(BeNil())
				allResults = append(allResults, segmentsList.Segments...)

				listSegmentsOptions.Offset, err = segmentsList.GetNextOffset()
				Expect(err).To(BeNil())

				if listSegmentsOptions.Offset == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListSegments(listSegmentsOptions *ListSegmentsOptions) using SegmentsPager`, func(){
			listSegmentsOptions := &appconfigurationv1.ListSegmentsOptions{
				Expand: core.BoolPtr(true),
				Sort: core.StringPtr("created_time"),
				Tags: core.StringPtr("version 1.1,pre-release"),
				Include: core.StringPtr("rules"),
				Limit: core.Int64Ptr(int64(10)),
				Search: core.StringPtr("test tag"),
			}

			// Test GetNext().
			pager, err := appConfigurationService.NewSegmentsPager(listSegmentsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []appconfigurationv1.Segment
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = appConfigurationService.NewSegmentsPager(listSegmentsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListSegments() returned a total of %d item(s) using SegmentsPager.\n", len(allResults))
		})
	})

	Describe(`CreateSegment - Create Segment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSegment(createSegmentOptions *CreateSegmentOptions)`, func() {
			ruleModel := &appconfigurationv1.Rule{
				AttributeName: core.StringPtr("email"),
				Operator: core.StringPtr("endsWith"),
				Values: []string{"@in.mnc.com", "@us.mnc.com"},
			}

			createSegmentOptions := &appconfigurationv1.CreateSegmentOptions{
				Name: core.StringPtr("Beta Users"),
				SegmentID: core.StringPtr("beta-users"),
				Rules: []appconfigurationv1.Rule{*ruleModel},
				Description: core.StringPtr("Segment containing the beta users"),
				Tags: core.StringPtr("version: 1.1, stage"),
			}

			segment, response, err := appConfigurationService.CreateSegment(createSegmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(segment).ToNot(BeNil())
		})
	})

	Describe(`UpdateSegment - Update Segment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateSegment(updateSegmentOptions *UpdateSegmentOptions)`, func() {
			ruleModel := &appconfigurationv1.Rule{
				AttributeName: core.StringPtr("testString"),
				Operator: core.StringPtr("is"),
				Values: []string{"testString"},
			}

			updateSegmentOptions := &appconfigurationv1.UpdateSegmentOptions{
				SegmentID: core.StringPtr("segment_id"),
				Name: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Tags: core.StringPtr("testString"),
				Rules: []appconfigurationv1.Rule{*ruleModel},
			}

			segment, response, err := appConfigurationService.UpdateSegment(updateSegmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(segment).ToNot(BeNil())
		})
	})

	Describe(`GetSegment - Get Segment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSegment(getSegmentOptions *GetSegmentOptions)`, func() {
			getSegmentOptions := &appconfigurationv1.GetSegmentOptions{
				SegmentID: core.StringPtr("segment_id"),
				Include: []string{"features", "properties"},
			}

			segment, response, err := appConfigurationService.GetSegment(getSegmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(segment).ToNot(BeNil())
		})
	})

	Describe(`ListSnapshots - Get list of Git configs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSnapshots(listSnapshotsOptions *ListSnapshotsOptions) with pagination`, func(){
			listSnapshotsOptions := &appconfigurationv1.ListSnapshotsOptions{
				Sort: core.StringPtr("created_time"),
				CollectionID: core.StringPtr("collection_id"),
				EnvironmentID: core.StringPtr("environment_id"),
				Limit: core.Int64Ptr(int64(10)),
				Offset: core.Int64Ptr(int64(0)),
				Search: core.StringPtr("search_string"),
			}

			listSnapshotsOptions.Offset = nil
			listSnapshotsOptions.Limit = core.Int64Ptr(1)

			var allResults []appconfigurationv1.GitConfig
			for {
				gitConfigList, response, err := appConfigurationService.ListSnapshots(listSnapshotsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(gitConfigList).ToNot(BeNil())
				allResults = append(allResults, gitConfigList.GitConfig...)

				listSnapshotsOptions.Offset, err = gitConfigList.GetNextOffset()
				Expect(err).To(BeNil())

				if listSnapshotsOptions.Offset == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListSnapshots(listSnapshotsOptions *ListSnapshotsOptions) using SnapshotsPager`, func(){
			listSnapshotsOptions := &appconfigurationv1.ListSnapshotsOptions{
				Sort: core.StringPtr("created_time"),
				CollectionID: core.StringPtr("collection_id"),
				EnvironmentID: core.StringPtr("environment_id"),
				Limit: core.Int64Ptr(int64(10)),
				Search: core.StringPtr("search_string"),
			}

			// Test GetNext().
			pager, err := appConfigurationService.NewSnapshotsPager(listSnapshotsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []appconfigurationv1.GitConfig
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = appConfigurationService.NewSnapshotsPager(listSnapshotsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListSnapshots() returned a total of %d item(s) using SnapshotsPager.\n", len(allResults))
		})
	})

	Describe(`CreateGitconfig - Create Git config`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateGitconfig(createGitconfigOptions *CreateGitconfigOptions)`, func() {
			createGitconfigOptions := &appconfigurationv1.CreateGitconfigOptions{
				GitConfigName: core.StringPtr("boot-strap-configuration"),
				GitConfigID: core.StringPtr("boot-strap-configuration"),
				CollectionID: core.StringPtr("web-app-collection"),
				EnvironmentID: core.StringPtr("dev"),
				GitURL: core.StringPtr("https://github.ibm.com/api/v3/repos/jhondoe-owner/my-test-repo"),
				GitBranch: core.StringPtr("main"),
				GitFilePath: core.StringPtr("code/development/README.json"),
				GitToken: core.StringPtr("61a792eahhGHji223jijb55a6cfdd4d5cde4c8a67esjjhjhHVH"),
			}

			createGitConfigResponse, response, err := appConfigurationService.CreateGitconfig(createGitconfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createGitConfigResponse).ToNot(BeNil())
		})
	})

	Describe(`UpdateGitconfig - Update Git Config`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateGitconfig(updateGitconfigOptions *UpdateGitconfigOptions)`, func() {
			updateGitconfigOptions := &appconfigurationv1.UpdateGitconfigOptions{
				GitConfigID: core.StringPtr("git_config_id"),
				GitConfigName: core.StringPtr("testString"),
				CollectionID: core.StringPtr("testString"),
				EnvironmentID: core.StringPtr("testString"),
				GitURL: core.StringPtr("testString"),
				GitBranch: core.StringPtr("testString"),
				GitFilePath: core.StringPtr("testString"),
				GitToken: core.StringPtr("testString"),
			}

			gitConfig, response, err := appConfigurationService.UpdateGitconfig(updateGitconfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(gitConfig).ToNot(BeNil())
		})
	})

	Describe(`GetGitconfig - Get Git Config`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetGitconfig(getGitconfigOptions *GetGitconfigOptions)`, func() {
			getGitconfigOptions := &appconfigurationv1.GetGitconfigOptions{
				GitConfigID: core.StringPtr("git_config_id"),
			}

			gitConfig, response, err := appConfigurationService.GetGitconfig(getGitconfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(gitConfig).ToNot(BeNil())
		})
	})

	Describe(`PromoteGitconfig - Promote configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PromoteGitconfig(promoteGitconfigOptions *PromoteGitconfigOptions)`, func() {
			promoteGitconfigOptions := &appconfigurationv1.PromoteGitconfigOptions{
				GitConfigID: core.StringPtr("git_config_id"),
			}

			gitConfigPromote, response, err := appConfigurationService.PromoteGitconfig(promoteGitconfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(gitConfigPromote).ToNot(BeNil())
		})
	})

	Describe(`RestoreGitconfig - Restore configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RestoreGitconfig(restoreGitconfigOptions *RestoreGitconfigOptions)`, func() {
			restoreGitconfigOptions := &appconfigurationv1.RestoreGitconfigOptions{
				GitConfigID: core.StringPtr("git_config_id"),
			}

			gitConfigRestore, response, err := appConfigurationService.RestoreGitconfig(restoreGitconfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(gitConfigRestore).ToNot(BeNil())
		})
	})

	Describe(`ListIntegrations - Get list of integrations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListIntegrations(listIntegrationsOptions *ListIntegrationsOptions) with pagination`, func(){
			listIntegrationsOptions := &appconfigurationv1.ListIntegrationsOptions{
				Expand: core.BoolPtr(true),
				Limit: core.Int64Ptr(int64(10)),
				Offset: core.Int64Ptr(int64(0)),
			}

			listIntegrationsOptions.Offset = nil
			listIntegrationsOptions.Limit = core.Int64Ptr(1)

			var allResults []appconfigurationv1.Integration
			for {
				integrationList, response, err := appConfigurationService.ListIntegrations(listIntegrationsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(integrationList).ToNot(BeNil())
				allResults = append(allResults, integrationList.Integrations...)

				listIntegrationsOptions.Offset, err = integrationList.GetNextOffset()
				Expect(err).To(BeNil())

				if listIntegrationsOptions.Offset == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListIntegrations(listIntegrationsOptions *ListIntegrationsOptions) using IntegrationsPager`, func(){
			listIntegrationsOptions := &appconfigurationv1.ListIntegrationsOptions{
				Expand: core.BoolPtr(true),
				Limit: core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := appConfigurationService.NewIntegrationsPager(listIntegrationsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []appconfigurationv1.Integration
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = appConfigurationService.NewIntegrationsPager(listIntegrationsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListIntegrations() returned a total of %d item(s) using IntegrationsPager.\n", len(allResults))
		})
	})

	Describe(`CreateIntegration - Create integration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateIntegration(createIntegrationOptions *CreateIntegrationOptions)`, func() {
			createIntegrationMetadataModel := &appconfigurationv1.CreateIntegrationMetadataCreateEnIntegrationMetadata{
				EventNotificationsInstanceCrn: core.StringPtr("crn:v1:bluemix:public:event-notifications:eu-gb:a/4f631ea3b3204b2b878a295604994acf:0eb42def-21aa-4f0a-a975-0812ead6ceee::"),
				EventNotificationsEndpoint: core.StringPtr("https://eu-gb.event-notifications.cloud.ibm.com"),
				EventNotificationsSourceName: core.StringPtr("My App Config"),
				EventNotificationsSourceDescription: core.StringPtr("All the events from App Configuration instance"),
			}

			createIntegrationOptions := &appconfigurationv1.CreateIntegrationOptions{
				IntegrationID: core.StringPtr("lckkhp34t"),
				IntegrationType: core.StringPtr("EVENT_NOTIFICATIONS"),
				Metadata: createIntegrationMetadataModel,
			}

			integration, response, err := appConfigurationService.CreateIntegration(createIntegrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(integration).ToNot(BeNil())
		})
	})

	Describe(`GetIntegration - Get integration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetIntegration(getIntegrationOptions *GetIntegrationOptions)`, func() {
			getIntegrationOptions := &appconfigurationv1.GetIntegrationOptions{
				IntegrationID: core.StringPtr("integration_id"),
			}

			integration, response, err := appConfigurationService.GetIntegration(getIntegrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(integration).ToNot(BeNil())
		})
	})

	Describe(`ListOriginconfigs - Get list of Origin Configs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListOriginconfigs(listOriginconfigsOptions *ListOriginconfigsOptions)`, func() {
			listOriginconfigsOptions := &appconfigurationv1.ListOriginconfigsOptions{
			}

			originConfigList, response, err := appConfigurationService.ListOriginconfigs(listOriginconfigsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(originConfigList).ToNot(BeNil())
		})
	})

	Describe(`UpdateOriginconfigs - Update Origin Configs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateOriginconfigs(updateOriginconfigsOptions *UpdateOriginconfigsOptions)`, func() {
			updateOriginconfigsOptions := &appconfigurationv1.UpdateOriginconfigsOptions{
				AllowedOrigins: []string{"testString"},
			}

			originConfigList, response, err := appConfigurationService.UpdateOriginconfigs(updateOriginconfigsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(originConfigList).ToNot(BeNil())
		})
	})

	Describe(`ListWorkflowconfig - Get Workflow Config`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListWorkflowconfig(listWorkflowconfigOptions *ListWorkflowconfigOptions)`, func() {
			listWorkflowconfigOptions := &appconfigurationv1.ListWorkflowconfigOptions{
				EnvironmentID: core.StringPtr("environment_id"),
			}

			listWorkflowconfigResponse, response, err := appConfigurationService.ListWorkflowconfig(listWorkflowconfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listWorkflowconfigResponse).ToNot(BeNil())
		})
	})

	Describe(`CreateWorkflowconfig - Create Workflow config`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateWorkflowconfig(createWorkflowconfigOptions *CreateWorkflowconfigOptions)`, func() {
			externalServiceNowCredentialsModel := &appconfigurationv1.ExternalServiceNowCredentials{
				Username: core.StringPtr("user"),
				Password: core.StringPtr("pwd"),
				ClientID: core.StringPtr("client id value"),
				ClientSecret: core.StringPtr("clientsecret"),
			}

			createWorkflowConfigModel := &appconfigurationv1.CreateWorkflowConfigExternalServiceNow{
				WorkflowURL: core.StringPtr("https://xxxxx.service-now.com"),
				ApprovalGroupName: core.StringPtr("WorkflowCRApprovers"),
				ApprovalExpiration: core.Int64Ptr(int64(10)),
				WorkflowCredentials: externalServiceNowCredentialsModel,
				Enabled: core.BoolPtr(true),
			}

			createWorkflowconfigOptions := &appconfigurationv1.CreateWorkflowconfigOptions{
				EnvironmentID: core.StringPtr("environment_id"),
				WorkflowConfig: createWorkflowConfigModel,
			}

			createWorkflowconfigResponse, response, err := appConfigurationService.CreateWorkflowconfig(createWorkflowconfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createWorkflowconfigResponse).ToNot(BeNil())
		})
	})

	Describe(`UpdateWorkflowconfig - Update Workflow config`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateWorkflowconfig(updateWorkflowconfigOptions *UpdateWorkflowconfigOptions)`, func() {
			externalServiceNowCredentialsModel := &appconfigurationv1.ExternalServiceNowCredentials{
				Username: core.StringPtr("user"),
				Password: core.StringPtr("updated password"),
				ClientID: core.StringPtr("client id value"),
				ClientSecret: core.StringPtr("updated client secret"),
			}

			updateWorkflowConfigModel := &appconfigurationv1.UpdateWorkflowConfigUpdateExternalServiceNow{
				WorkflowURL: core.StringPtr("https://xxxxx.service-now.com"),
				ApprovalGroupName: core.StringPtr("WorkflowCRApprovers"),
				ApprovalExpiration: core.Int64Ptr(int64(5)),
				WorkflowCredentials: externalServiceNowCredentialsModel,
				Enabled: core.BoolPtr(true),
			}

			updateWorkflowconfigOptions := &appconfigurationv1.UpdateWorkflowconfigOptions{
				EnvironmentID: core.StringPtr("environment_id"),
				UpdateWorkflowConfig: updateWorkflowConfigModel,
			}

			updateWorkflowconfigResponse, response, err := appConfigurationService.UpdateWorkflowconfig(updateWorkflowconfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(updateWorkflowconfigResponse).ToNot(BeNil())
		})
	})

	Describe(`ImportConfig - Import instance configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ImportConfig(importConfigOptions *ImportConfigOptions)`, func() {
			targetSegmentsModel := &appconfigurationv1.TargetSegments{
				Segments: []string{"testString"},
			}

			featureSegmentRuleModel := &appconfigurationv1.FeatureSegmentRule{
				Rules: []appconfigurationv1.TargetSegments{*targetSegmentsModel},
				Value: "testString",
				Order: core.Int64Ptr(int64(38)),
				RolloutPercentage: core.Int64Ptr(int64(100)),
			}

			collectionRefModel := &appconfigurationv1.CollectionRef{
				CollectionID: core.StringPtr("web-app"),
			}

			importFeatureRequestBodyModel := &appconfigurationv1.ImportFeatureRequestBody{
				Name: core.StringPtr("Cycle Rentals"),
				FeatureID: core.StringPtr("cycle-rentals"),
				Description: core.StringPtr("testString"),
				Type: core.StringPtr("NUMERIC"),
				Format: core.StringPtr("TEXT"),
				EnabledValue: core.StringPtr("1"),
				DisabledValue: core.StringPtr("2"),
				Enabled: core.BoolPtr(true),
				RolloutPercentage: core.Int64Ptr(int64(100)),
				Tags: core.StringPtr("testString"),
				SegmentRules: []appconfigurationv1.FeatureSegmentRule{*featureSegmentRuleModel},
				Collections: []appconfigurationv1.CollectionRef{*collectionRefModel},
			}

			segmentRuleModel := &appconfigurationv1.SegmentRule{
				Rules: []appconfigurationv1.TargetSegments{*targetSegmentsModel},
				Value: core.StringPtr("200"),
				Order: core.Int64Ptr(int64(1)),
			}

			importPropertyRequestBodyModel := &appconfigurationv1.ImportPropertyRequestBody{
				Name: core.StringPtr("Daily Discount"),
				PropertyID: core.StringPtr("daily_discount"),
				Description: core.StringPtr("testString"),
				Type: core.StringPtr("NUMERIC"),
				Format: core.StringPtr("TEXT"),
				Value: core.StringPtr("100"),
				Tags: core.StringPtr("pre-release, v1.2"),
				SegmentRules: []appconfigurationv1.SegmentRule{*segmentRuleModel},
				Collections: []appconfigurationv1.CollectionRef{*collectionRefModel},
			}

			importEnvironmentSchemaModel := &appconfigurationv1.ImportEnvironmentSchema{
				Name: core.StringPtr("Dev"),
				EnvironmentID: core.StringPtr("dev"),
				Description: core.StringPtr("Environment created on instance creation"),
				Tags: core.StringPtr("testString"),
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

			importConfigOptions := &appconfigurationv1.ImportConfigOptions{
				Environments: []appconfigurationv1.ImportEnvironmentSchema{*importEnvironmentSchemaModel},
				Collections: []appconfigurationv1.ImportCollectionSchema{*importCollectionSchemaModel},
				Segments: []appconfigurationv1.ImportSegmentSchema{*importSegmentSchemaModel},
				Clean: core.StringPtr("true"),
			}

			instanceConfigAcceptedResponse, response, err := appConfigurationService.ImportConfig(importConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(instanceConfigAcceptedResponse).ToNot(BeNil())
		})
	})

	Describe(`ListInstanceConfig - Export instance configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListInstanceConfig(listInstanceConfigOptions *ListInstanceConfigOptions)`, func() {
			listInstanceConfigOptions := &appconfigurationv1.ListInstanceConfigOptions{
			}

			importConfig, response, err := appConfigurationService.ListInstanceConfig(listInstanceConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(importConfig).ToNot(BeNil())
		})
	})

	Describe(`PromoteRestoreConfig - Promote or Restore snapshot configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PromoteRestoreConfig(promoteRestoreConfigOptions *PromoteRestoreConfigOptions)`, func() {
			promoteRestoreConfigOptions := &appconfigurationv1.PromoteRestoreConfigOptions{
				GitConfigID: core.StringPtr("git_config_id"),
				Action: core.StringPtr("promote"),
			}

			configAction, response, err := appConfigurationService.PromoteRestoreConfig(promoteRestoreConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(configAction).ToNot(BeNil())
		})
	})

	Describe(`InstanceConfigStatus - Get status of instance configuration import / export`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`InstanceConfigStatus(instanceConfigStatusOptions *InstanceConfigStatusOptions)`, func() {
			instanceConfigStatusOptions := &appconfigurationv1.InstanceConfigStatusOptions{
				ReferenceID: core.StringPtr("testString"),
				Action: core.StringPtr("import"),
			}

			instanceConfigStatusResponse, response, err := appConfigurationService.InstanceConfigStatus(instanceConfigStatusOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceConfigStatusResponse).ToNot(BeNil())
		})
	})

	Describe(`DeleteEnvironment - Delete Environment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteEnvironment(deleteEnvironmentOptions *DeleteEnvironmentOptions)`, func() {
			deleteEnvironmentOptions := &appconfigurationv1.DeleteEnvironmentOptions{
				EnvironmentID: core.StringPtr("environment_id"),
			}

			response, err := appConfigurationService.DeleteEnvironment(deleteEnvironmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteCollection - Delete Collection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteCollection(deleteCollectionOptions *DeleteCollectionOptions)`, func() {
			deleteCollectionOptions := &appconfigurationv1.DeleteCollectionOptions{
				CollectionID: core.StringPtr("collection_id"),
			}

			response, err := appConfigurationService.DeleteCollection(deleteCollectionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteFeature - Delete Feature`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteFeature(deleteFeatureOptions *DeleteFeatureOptions)`, func() {
			deleteFeatureOptions := &appconfigurationv1.DeleteFeatureOptions{
				EnvironmentID: core.StringPtr("environment_id"),
				FeatureID: core.StringPtr("feature_id"),
			}

			response, err := appConfigurationService.DeleteFeature(deleteFeatureOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteProperty - Delete Property`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteProperty(deletePropertyOptions *DeletePropertyOptions)`, func() {
			deletePropertyOptions := &appconfigurationv1.DeletePropertyOptions{
				EnvironmentID: core.StringPtr("environment_id"),
				PropertyID: core.StringPtr("property_id"),
			}

			response, err := appConfigurationService.DeleteProperty(deletePropertyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteSegment - Delete Segment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSegment(deleteSegmentOptions *DeleteSegmentOptions)`, func() {
			deleteSegmentOptions := &appconfigurationv1.DeleteSegmentOptions{
				SegmentID: core.StringPtr("segment_id"),
			}

			response, err := appConfigurationService.DeleteSegment(deleteSegmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteGitconfig - Delete Git Config`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteGitconfig(deleteGitconfigOptions *DeleteGitconfigOptions)`, func() {
			deleteGitconfigOptions := &appconfigurationv1.DeleteGitconfigOptions{
				GitConfigID: core.StringPtr("git_config_id"),
			}

			response, err := appConfigurationService.DeleteGitconfig(deleteGitconfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteIntegration - Delete integration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteIntegration(deleteIntegrationOptions *DeleteIntegrationOptions)`, func() {
			deleteIntegrationOptions := &appconfigurationv1.DeleteIntegrationOptions{
				IntegrationID: core.StringPtr("integration_id"),
			}

			response, err := appConfigurationService.DeleteIntegration(deleteIntegrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteWorkflowconfig - Delete  Workflow config`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteWorkflowconfig(deleteWorkflowconfigOptions *DeleteWorkflowconfigOptions)`, func() {
			deleteWorkflowconfigOptions := &appconfigurationv1.DeleteWorkflowconfigOptions{
				EnvironmentID: core.StringPtr("environment_id"),
			}

			response, err := appConfigurationService.DeleteWorkflowconfig(deleteWorkflowconfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
