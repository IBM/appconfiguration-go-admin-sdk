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

package main

import (
	"fmt"
	"sync"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/IBM/go-sdk-core/v5/core"
)

var appConfigurationServiceInstance *appconfigurationv1.AppConfigurationV1

func initAndReturnSingletonInstanceWithAPIKey(authToken string, guid string, region string) *appconfigurationv1.AppConfigurationV1 {

	var once sync.Once
	if appConfigurationServiceInstance == nil {
		once.Do(func() {
			if appConfigurationServiceInstance == nil {
				authenticator := &core.IamAuthenticator{
					ApiKey: authToken,
				}
				options := &appconfigurationv1.AppConfigurationV1Options{
					Authenticator: authenticator,
					URL:           "https://" + region + ".apprapp.cloud.ibm.com/apprapp/feature/v1/instances/" + guid,
				}
				var err error
				appConfigurationServiceInstance, err = appconfigurationv1.NewAppConfigurationV1(options)
				if err != nil {
					fmt.Println("Error: " + err.Error())
					return
				}
			}
		})
	}
	return appConfigurationServiceInstance
}

func initAndReturnSingletonInstanceWithBearerToken(authToken string, guid string, region string) *appconfigurationv1.AppConfigurationV1 {

	var once sync.Once
	if appConfigurationServiceInstance == nil {
		once.Do(func() {
			if appConfigurationServiceInstance == nil {
				authenticator := &core.BearerTokenAuthenticator{
					BearerToken: authToken,
				}
				options := &appconfigurationv1.AppConfigurationV1Options{
					Authenticator: authenticator,
					URL:           "https://" + region + ".apprapp.cloud.ibm.com/apprapp/feature/v1/instances/" + guid,
				}
				var err error
				appConfigurationServiceInstance, err = appconfigurationv1.NewAppConfigurationV1(options)
				if err != nil {
					fmt.Println("Error: " + err.Error())
					return
				}
			}
		})
	}
	return appConfigurationServiceInstance
}

func main() {

	authToken := "<authToken>"
	guid := "<guid>"
	region := "<region>"

	initAndReturnSingletonInstanceWithAPIKey(authToken, guid, region)

	createEnvironment("environmentId", "environmentName", "desc", "tags", "#FDD13A")
	createCollection("collectionId", "collectionName", "desc", "tags")
	createSegment("segmentName", "segmentId", "desc", "tags", "email", "endsWith", []string{"@in.ibm.com"})
	createFeature("environmentId", "booleanFeatureName", "booleanFeatureId", "desc", "BOOLEAN", "true", "false", "tags", []string{"segmentId"}, 1, "collectionId", "true")
	createFeature("environmentId", "numberFeatureName", "numberFeatureId", "desc", "NUMERIC", "1", "2", "tags", []string{"segmentId"}, 1, "collectionId", "3")
	createProperty("environmentId", "booleanPropertyName", "booleanPropertyId", "desc", "BOOLEAN", "true", "tags", []string{"segmentId"}, "collectionId", 2, "true")
	createProperty("environmentId", "numberPropertyName", "numberPropertyId", "desc", "NUMERIC", "2", "tags", []string{"segmentId"}, "collectionId", 2, "4")

	toggleFeature("environmentId", "booleanFeatureId", true)

	listEnvironments()
	listCollections()
	listFeatures("environmentId")
	listSegments()
	listProperties("environmentId")

	updateFeature("environmentId", "numberFeatureId", "numberFeatureName", "updatedDesc", "1", "1", "tags", []string{}, 1, "collectionId", "2", true)
	updateCollection("collectionId", "collectionName", "updatedDesc", "updatedTags")
	updateSegment("segmentId", "segmentName", "updatedDesc", "updatedTags", "email", "endsWith", []string{"@in.ibm.com"})
	updateProperty("environmentId", "booleanPropertyName", "booleanPropertyId", "updatedDescBoolean", "true", "updatedTags", []string{"segmentId"}, "collectionId", 2, "true")
	updateEnvironment("environmentId", "environmentName", "updatedDesc", "tags", "#FDD13A")

	getEnvironment("environmentId")
	getCollection("collectionId")
	getFeature("environmentId", "booleanFeatureId")
	getProperty("environmentId", "booleanPropertyId")
	getSegment("segmentId")

	updateFeatureValues("environmentId", "booleanFeatureId", "booleanFeatureName", "patchedDesc", "1", "12", "tag", []string{}, 1, "2")
	updatePropertyValues("environmentId", "numberPropertyName", "numberPropertyId", "desc", "1", "tags", []string{"segmentId"}, 2, "2")

	deleteFeature("environmentId", "numberFeatureId")
	deleteFeature("environmentId", "booleanFeatureId")
	deleteProperty("environmentId", "numberPropertyId")
	deleteProperty("environmentId", "booleanPropertyId")
	deleteSegment("segmentId")
	deleteCollection("collectionId")
	deleteEnvironment("environmentId")
}

// Create examples

func createEnvironment(environmentId string, name string, description string, tags string, colorCode string) {
	fmt.Println("createEnvironment")
	createEnvironmentOptionsModel := appConfigurationServiceInstance.NewCreateEnvironmentOptions(name, environmentId)
	createEnvironmentOptionsModel.SetDescription(description)
	createEnvironmentOptionsModel.SetTags(tags)
	createEnvironmentOptionsModel.SetColorCode(colorCode)
	result, response, err := appConfigurationServiceInstance.CreateEnvironment(createEnvironmentOptionsModel)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.EnvironmentID)
}

func createCollection(collectionId string, name string, description string, tags string) {
	fmt.Println("createCollection")
	createCollectionOptionsModel := appConfigurationServiceInstance.NewCreateCollectionOptions(name, collectionId)
	createCollectionOptionsModel.SetDescription(description)
	createCollectionOptionsModel.SetTags(tags)
	result, response, err := appConfigurationServiceInstance.CreateCollection(createCollectionOptionsModel)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.CollectionID)
}

func createSegment(name string, id string, description string, tags string, attributeName string, operator string, values []string) {
	fmt.Println("createSegment")
	ruleArray, _ := appConfigurationServiceInstance.NewRule(attributeName, operator, values)
	createSegmentOptionsModel := appConfigurationServiceInstance.NewCreateSegmentOptions()
	createSegmentOptionsModel.SetName(name)
	createSegmentOptionsModel.SetDescription(description)
	createSegmentOptionsModel.SetTags(tags)
	createSegmentOptionsModel.SetSegmentID(id)
	createSegmentOptionsModel.SetRules([]appconfigurationv1.Rule{*ruleArray})
	result, response, err := appConfigurationServiceInstance.CreateSegment(createSegmentOptionsModel)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.SegmentID)
}

func createFeature(environmentId string, name string, id string, description string, typeOfFeature string, enabledValue string, disabledValue string, tags string, segments []string, order int64, collectionId string, value string) {
	fmt.Println("createFeature")
	ruleArray, _ := appConfigurationServiceInstance.NewTargetSegments(segments)
	segmentRuleArray, _ := appConfigurationServiceInstance.NewSegmentRule([]appconfigurationv1.TargetSegments{*ruleArray}, value, order)
	collectionArray, _ := appConfigurationServiceInstance.NewCollectionRef(collectionId)
	createFeatureOptionsModel := appConfigurationServiceInstance.NewCreateFeatureOptions(environmentId, name, id, typeOfFeature, enabledValue, disabledValue)
	createFeatureOptionsModel.SetTags(tags)
	createFeatureOptionsModel.SetDescription(description)
	createFeatureOptionsModel.SetSegmentRules([]appconfigurationv1.SegmentRule{*segmentRuleArray})
	createFeatureOptionsModel.SetCollections([]appconfigurationv1.CollectionRef{*collectionArray})
	result, response, err := appConfigurationServiceInstance.CreateFeature(createFeatureOptionsModel)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.FeatureID)
}

func createProperty(environmentId string, name string, propertyId string, description string, typeOfProperty string, valueOfProperty string, tags string, segments []string, collectionId string, order int64, value string) {
	fmt.Println("createProperty")
	ruleArray, _ := appConfigurationServiceInstance.NewTargetSegments(segments)
	segmentRuleArray, _ := appConfigurationServiceInstance.NewSegmentRule([]appconfigurationv1.TargetSegments{*ruleArray}, value, order)
	collectionArray, _ := appConfigurationServiceInstance.NewCollectionRef(collectionId)
	createPropertyOptionsModel := appConfigurationServiceInstance.NewCreatePropertyOptions(environmentId, name, propertyId, typeOfProperty, valueOfProperty)
	createPropertyOptionsModel.SetTags(tags)
	createPropertyOptionsModel.SetDescription(description)
	createPropertyOptionsModel.SetSegmentRules([]appconfigurationv1.SegmentRule{*segmentRuleArray})
	createPropertyOptionsModel.SetCollections([]appconfigurationv1.CollectionRef{*collectionArray})
	result, response, err := appConfigurationServiceInstance.CreateProperty(createPropertyOptionsModel)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.PropertyID)
}

// Update examples

func updateEnvironment(environmentId string, name string, description string, tags string, colorCode string) {
	fmt.Println("updateEnvironment")
	updateEnvironmentOptionsModel := appConfigurationServiceInstance.NewUpdateEnvironmentOptions(environmentId)
	updateEnvironmentOptionsModel.SetName(name)
	updateEnvironmentOptionsModel.SetDescription(description)
	updateEnvironmentOptionsModel.SetTags(tags)
	updateEnvironmentOptionsModel.SetColorCode(colorCode)
	result, response, err := appConfigurationServiceInstance.UpdateEnvironment(updateEnvironmentOptionsModel)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Description)
}

func updateFeature(environmentId string, id string, name string, description string, enabledValue string, disabledValue string, tags string, segments []string, order int64, collectionId string, value string, deletedFlag bool) {
	fmt.Println("updateFeatureWithNumberValue")
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
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Name)
}

func updateCollection(collectionId string, name string, description string, tags string) {
	fmt.Println("updateCollection")
	updateCollectionOptionsModel := appConfigurationServiceInstance.NewUpdateCollectionOptions(collectionId)
	updateCollectionOptionsModel.SetName(name)
	updateCollectionOptionsModel.SetTags(tags)
	updateCollectionOptionsModel.SetDescription(description)
	result, response, err := appConfigurationServiceInstance.UpdateCollection(updateCollectionOptionsModel)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Description)
}

func updateSegment(segmentId string, name string, description string, tags string, attributeName string, operator string, values []string) {
	fmt.Println("updateSegment")
	ruleArray, _ := appConfigurationServiceInstance.NewRule(attributeName, operator, values)
	updateSegmentOptionsModel := appConfigurationServiceInstance.NewUpdateSegmentOptions(segmentId)
	updateSegmentOptionsModel.SetName(name)
	updateSegmentOptionsModel.SetDescription(description)
	updateSegmentOptionsModel.SetTags(tags)
	updateSegmentOptionsModel.SetRules([]appconfigurationv1.Rule{*ruleArray})
	result, response, err := appConfigurationServiceInstance.UpdateSegment(updateSegmentOptionsModel)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Name)
}

func updateProperty(environmentId string, name string, propertyId string, description string, valueOfProperty string, tags string, segments []string, collectionId string, order int64, value string) {
	fmt.Println("updateProperty")
	ruleArray, _ := appConfigurationServiceInstance.NewTargetSegments(segments)
	segmentRuleArray, _ := appConfigurationServiceInstance.NewSegmentRule([]appconfigurationv1.TargetSegments{*ruleArray}, value, order)
	collectionArray, _ := appConfigurationServiceInstance.NewCollectionRef(collectionId)
	updatePropertyOptionsModel := appConfigurationServiceInstance.NewUpdatePropertyOptions(environmentId, propertyId)
	updatePropertyOptionsModel.SetName(name)
	updatePropertyOptionsModel.SetDescription(description)
	updatePropertyOptionsModel.SetTags(tags)
	updatePropertyOptionsModel.SetValue(valueOfProperty)
	updatePropertyOptionsModel.SetSegmentRules([]appconfigurationv1.SegmentRule{*segmentRuleArray})
	updatePropertyOptionsModel.SetCollections([]appconfigurationv1.CollectionRef{*collectionArray})
	result, response, err := appConfigurationServiceInstance.UpdateProperty(updatePropertyOptionsModel)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.PropertyID)
}

// Update Values Examples

func updatePropertyValues(environmentId string, name string, propertyId string, description string, valueOfProperty string, tags string, segments []string, order int64, value string) {
	fmt.Println("patchProperty")
	ruleArray, _ := appConfigurationServiceInstance.NewTargetSegments(segments)
	segmentRuleArray, _ := appConfigurationServiceInstance.NewSegmentRule([]appconfigurationv1.TargetSegments{*ruleArray}, value, order)
	patchPropertyOptionsModel := appConfigurationServiceInstance.NewUpdatePropertyValuesOptions(environmentId, propertyId)
	patchPropertyOptionsModel.SetName(name)
	patchPropertyOptionsModel.SetDescription(description)
	patchPropertyOptionsModel.SetTags(tags)
	patchPropertyOptionsModel.SetValue(valueOfProperty)
	patchPropertyOptionsModel.SetSegmentRules([]appconfigurationv1.SegmentRule{*segmentRuleArray})
	result, response, err := appConfigurationServiceInstance.UpdatePropertyValues(patchPropertyOptionsModel)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.PropertyID)
}

func updateFeatureValues(environmentId string, id string, name string, description string, enabledValue string, disabledValue string, tags string, segments []string, order int64, value string) {
	fmt.Println("patchFeatureWithNumberValue")
	ruleArray, _ := appConfigurationServiceInstance.NewTargetSegments(segments)
	segmentRuleArray, _ := appConfigurationServiceInstance.NewSegmentRule([]appconfigurationv1.TargetSegments{*ruleArray}, value, order)
	patchFeatureOptionsModel := appConfigurationServiceInstance.NewUpdateFeatureValuesOptions(environmentId, id)
	patchFeatureOptionsModel.SetName(name)
	patchFeatureOptionsModel.SetDescription(description)
	patchFeatureOptionsModel.SetTags(tags)
	patchFeatureOptionsModel.SetDisabledValue(disabledValue)
	patchFeatureOptionsModel.SetEnabledValue(enabledValue)
	patchFeatureOptionsModel.SetSegmentRules([]appconfigurationv1.SegmentRule{*segmentRuleArray})
	result, response, err := appConfigurationServiceInstance.UpdateFeatureValues(patchFeatureOptionsModel)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Name)
}

// Delete examples

func deleteCollection(collectionId string) {
	fmt.Println("deleteCollection")
	deleteCollectionOptionsModel := appConfigurationServiceInstance.NewDeleteCollectionOptions(collectionId)
	response, err := appConfigurationServiceInstance.DeleteCollection(deleteCollectionOptionsModel)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println(response.StatusCode)
}

func deleteFeature(environmentId string, featureId string) {
	fmt.Println("deleteFeature")
	deleteFeatureOptionsModel := appConfigurationServiceInstance.NewDeleteFeatureOptions(environmentId, featureId)
	response, err := appConfigurationServiceInstance.DeleteFeature(deleteFeatureOptionsModel)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println(response.StatusCode)
}

func deleteSegment(segmentId string) {
	fmt.Println("deleteSegment")
	deleteSegmentOptionsModel := appConfigurationServiceInstance.NewDeleteSegmentOptions(segmentId)
	response, err := appConfigurationServiceInstance.DeleteSegment(deleteSegmentOptionsModel)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println(response.StatusCode)
}

func deleteProperty(environmentId string, propertyId string) {
	fmt.Println("deleteProperty")
	deletePropertyOptionsModel := appConfigurationServiceInstance.NewDeletePropertyOptions(environmentId, propertyId)
	response, err := appConfigurationServiceInstance.DeleteProperty(deletePropertyOptionsModel)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println(response.StatusCode)
}

func deleteEnvironment(environmentId string) {
	fmt.Println("deleteEnvironment")
	deleteEnvironmentOptionsModel := appConfigurationServiceInstance.NewDeleteEnvironmentOptions(environmentId)
	response, err := appConfigurationServiceInstance.DeleteEnvironment(deleteEnvironmentOptionsModel)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println(response.StatusCode)
}

// List examples

func listCollections() {
	fmt.Println("listCollections")
	listCollectionsOptionsModel := appConfigurationServiceInstance.NewListCollectionsOptions()
	listCollectionsOptionsModel.SetExpand(true)
	result, response, err := appConfigurationServiceInstance.ListCollections(listCollectionsOptionsModel)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(len(result.Collections))
}

func listFeatures(environmentId string) {
	fmt.Println("listFeatures")
	listFeaturesOptionsModel := appConfigurationServiceInstance.NewListFeaturesOptions(environmentId)
	result, response, err := appConfigurationServiceInstance.ListFeatures(listFeaturesOptionsModel)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(len(result.Features))
}

func listSegments() {
	fmt.Println("listSegments")
	listSegmentsOptionsModel := appConfigurationServiceInstance.NewListSegmentsOptions()
	result, response, err := appConfigurationServiceInstance.ListSegments(listSegmentsOptionsModel)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(len(result.Segments))
}

func listProperties(environmentId string) {
	fmt.Println("listProperties")
	listPropertiesOptionsModel := appConfigurationServiceInstance.NewListPropertiesOptions(environmentId)
	result, response, err := appConfigurationServiceInstance.ListProperties(listPropertiesOptionsModel)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(len(result.Properties))
}

func listEnvironments() {
	fmt.Println("listEnvironments")
	listEnvironmentsOptionsModel := appConfigurationServiceInstance.NewListEnvironmentsOptions()
	result, response, err := appConfigurationServiceInstance.ListEnvironments(listEnvironmentsOptionsModel)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(len(result.Environments))
}

// Get examples

func getCollection(collectionId string) {
	fmt.Println("getCollection")
	getCollectionOptionsModel := appConfigurationServiceInstance.NewGetCollectionOptions(collectionId)
	result, response, err := appConfigurationServiceInstance.GetCollection(getCollectionOptionsModel)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Name)
}

func getFeature(environmentId string, featureId string) {
	fmt.Println("getFeature")
	getFeatureOptionsModel := appConfigurationServiceInstance.NewGetFeatureOptions(environmentId, featureId)
	result, response, err := appConfigurationServiceInstance.GetFeature(getFeatureOptionsModel)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Name)        // referenced field, so needs to be de-referenced
	fmt.Println(result.EnabledValue) // non-referenced field, so can be used as it is
}

func getSegment(segmentId string) {
	fmt.Println("getSegment")
	getSegmentOptionsModel := appConfigurationServiceInstance.NewGetSegmentOptions(segmentId)
	result, response, err := appConfigurationServiceInstance.GetSegment(getSegmentOptionsModel)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Name)
}

func getProperty(environmentId string, propertyId string) {
	fmt.Println("getProperty")
	getPropertyOptionsModel := appConfigurationServiceInstance.NewGetPropertyOptions(environmentId, propertyId)
	result, response, err := appConfigurationServiceInstance.GetProperty(getPropertyOptionsModel)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Name)
}

func getEnvironment(environmentId string) {
	fmt.Println("getEnvironment")
	getEnvironmentOptionsModel := appConfigurationServiceInstance.NewGetEnvironmentOptions(environmentId)
	result, response, err := appConfigurationServiceInstance.GetEnvironment(getEnvironmentOptionsModel)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Name)
}

// Toggle Examples

func toggleFeature(environmentId string, featureId string, enableFlag bool) {
	fmt.Println("toggleFeature")
	toggleFeatureOptionsModel := appConfigurationServiceInstance.NewToggleFeatureOptions(environmentId, featureId)
	toggleFeatureOptionsModel.SetEnabled(enableFlag)
	result, response, err := appConfigurationServiceInstance.ToggleFeature(toggleFeatureOptionsModel)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Name)
}
