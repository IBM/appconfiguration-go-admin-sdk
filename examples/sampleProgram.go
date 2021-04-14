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

func initAndReturnSingletonInstanceWithAPIKey(authToken string, url string) *appconfigurationv1.AppConfigurationV1 {

	var once sync.Once
	if appConfigurationServiceInstance == nil {
		once.Do(func() {
			if appConfigurationServiceInstance == nil {
				authenticator := &core.IamAuthenticator{
					ApiKey: authToken,
				}
				options := &appconfigurationv1.AppConfigurationV1Options{
					Authenticator: authenticator,
					URL:           url,
				}
				var error error
				appConfigurationServiceInstance, error = appconfigurationv1.NewAppConfigurationV1(options)
				if error != nil {
					fmt.Println("Error: " + error.Error())
					return
				}
			}
		})
	}
	return appConfigurationServiceInstance
}

func initAndReturnSingletonInstanceWithBearerToken(authToken string, url string) *appconfigurationv1.AppConfigurationV1 {

	var once sync.Once
	if appConfigurationServiceInstance == nil {
		once.Do(func() {
			if appConfigurationServiceInstance == nil {
				authenticator := &core.BearerTokenAuthenticator{
					BearerToken: authToken,
				}
				options := &appconfigurationv1.AppConfigurationV1Options{
					Authenticator: authenticator,
					URL:           url,
				}
				var error error
				appConfigurationServiceInstance, error = appconfigurationv1.NewAppConfigurationV1(options)
				if error != nil {
					fmt.Println("Error: " + error.Error())
					return
				}
			}
		})
	}
	return appConfigurationServiceInstance
}

func main() {

	authToken := "<authToken>"
	url := "<url>"

	initAndReturnSingletonInstanceWithAPIKey(authToken, url)

	createCollection("collectionId", "collectionName", "desc", "tags")
	createSegment("segmentName", "segmentId", "desc", "tags", "email", "endsWith", []string{"@in.ibm.com"})
	createFeature("booleanFeatureName", "booleanFeatureId", "desc", "BOOLEAN", "true", "false", "tags", []string{"segmentId"}, 1, "collectionId", false, "true")
	createFeature("numberFeatureName", "numberFeatureId", "desc", "NUMERIC", "1", "2", "tags", []string{"segmentId"}, 1, "collectionId", true, "3")
	createProperty("booleanPropertyName", "booleanPropertyId", "desc", "BOOLEAN", "true", "tags", []string{"segmentId"}, "collectionId", 2, "true")
	createProperty("numberPropertyName", "numberPropertyId", "desc", "NUMERIC", "2", "tags", []string{"segmentId"}, "collectionId", 2, "4")

	toggleFeature("numberFeatureId", false)

	getCollections()
	getFeatures()
	getSegments()
	getProperties()

	updateFeature("numberFeatureId", "numberFeatureName", "updatedDesc", "1", "1", "tags", []string{}, 1, "collectionId", false, "2", true)
	updateCollection("collectionId", "collectionName", "updatedDesc", "updatedTags")
	updateSegment("segmentId", "segmentName", "updatedDesc", "updatedTags", "email", "endsWith", []string{"@in.ibm.com"})
	updateProperty("booleanPropertyName", "booleanPropertyId", "updatedDescBoolean", "true", "updatedTags", []string{"segmentId"}, "collectionId", 2, "true")

	getCollection("collectionId")
	getFeature("booleanFeatureId")
	getProperty("booleanPropertyId")
	getSegment("segmentId")

	patchFeature("booleanFeatureId", "booleanFeatureName", "patchedDesc", "1", "12", "tag", []string{}, 1, "2")
	patchProperty("numberPropertyName", "numberPropertyId", "desc", "1", "tags", []string{"segmentId"}, 2, "2")

	deleteSegment("segmentId")
	deleteCollection("collectionId")
	deleteFeature("numberFeatureId")
	deleteFeature("booleanFeatureId")
	deleteProperty("numberPropertyId")
	deleteProperty("booleanPropertyId")
}

// Create examples

func createCollection(collectionId string, name string, description string, tags string) {
	fmt.Println("createCollection")
	createCollectionOptionsModel := appConfigurationServiceInstance.NewCreateCollectionOptions(name)
	createCollectionOptionsModel.SetCollectionID(collectionId)
	createCollectionOptionsModel.SetDescription(description)
	createCollectionOptionsModel.SetTags(tags)
	result, response, error := appConfigurationServiceInstance.CreateCollection(createCollectionOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.CollectionID)
}

func createSegment(name string, id string, description string, tags string, attributeName string, operator string, values []string) {
	fmt.Println("createSegment")
	ruleArray, _ := appConfigurationServiceInstance.NewRuleArray(attributeName, operator, values)
	createSegmentOptionsModel := appConfigurationServiceInstance.NewCreateSegmentOptions(name, id, description, tags, []appconfigurationv1.RuleArray{*ruleArray}, "SDK")
	result, response, error := appConfigurationServiceInstance.CreateSegment(createSegmentOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.SegmentID)
}

func createFeature(name string, id string, description string, typeOfFeature string, enabledValue string, disabledValue string, tags string, segments []string, order int64, collectionId string, enabledInCollection bool, value string) {
	fmt.Println("createFeature")
	ruleArray, _ := appConfigurationServiceInstance.NewRule(segments)
	segmentRuleArray, _ := appConfigurationServiceInstance.NewSegmentRule([]appconfigurationv1.Rule{*ruleArray}, value, order)
	collectionArray, _ := appConfigurationServiceInstance.NewCollection(collectionId, enabledInCollection)
	createFeatureOptionsModel := appConfigurationServiceInstance.NewCreateFeatureOptions(name, id, description, typeOfFeature, enabledValue, disabledValue, tags, []appconfigurationv1.SegmentRule{*segmentRuleArray}, []appconfigurationv1.Collection{*collectionArray}, "SDK")
	result, response, error := appConfigurationServiceInstance.CreateFeature(createFeatureOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.FeatureID)
}


func createProperty(name string, propertyId string, description string, typeOfProperty string, valueOfProperty string, tags string, segments []string, collectionId string, order int64, value string) {
	fmt.Println("createProperty")
	ruleArray2, _ := appConfigurationServiceInstance.NewRule(segments)
	segmentRuleArray, _ := appConfigurationServiceInstance.NewSegmentRule([]appconfigurationv1.Rule{*ruleArray2}, value, order)
	collectionArray, _ := appConfigurationServiceInstance.NewCollectionID(collectionId)
	createPropertyOptionsModel := appConfigurationServiceInstance.NewCreatePropertyOptions(name, propertyId, description, typeOfProperty, valueOfProperty, tags, []appconfigurationv1.SegmentRule{*segmentRuleArray}, []appconfigurationv1.CollectionID{*collectionArray})
	result, response, error := appConfigurationServiceInstance.CreateProperty(createPropertyOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.PropertyID)
}

// Update examples

func updateFeature(id string, name string, description string, enabledValue string, disabledValue string, tags string, segments []string, order int64, collectionName string, enabledInCollection bool, value string, deletedFlag bool) {
	fmt.Println("updateFeatureWithNumberValue")
	ruleArray, _ := appConfigurationServiceInstance.NewRule(segments)
	segmentRuleArray, _ := appConfigurationServiceInstance.NewSegmentRule([]appconfigurationv1.Rule{*ruleArray}, value, order)
	collectionArray, _ := appConfigurationServiceInstance.NewCollectionWithDeletedFlag(collectionName, enabledInCollection, deletedFlag)
	updateFeatureOptionsModel := appConfigurationServiceInstance.NewUpdateFeatureOptions(id, name, description, enabledValue, disabledValue, tags, []appconfigurationv1.SegmentRule{*segmentRuleArray}, []appconfigurationv1.CollectionWithDeletedFlag{*collectionArray})
	result, response, error := appConfigurationServiceInstance.UpdateFeature(updateFeatureOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Name)
}

func updateCollection(collectionId string, name string, description string, tags string) {
	fmt.Println("updateCollection")
	updateCollectionOptionsModel := appConfigurationServiceInstance.NewUpdateCollectionOptions(collectionId, name, description, tags)
	result, response, error := appConfigurationServiceInstance.UpdateCollection(updateCollectionOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Description)
}

func patchFeature(id string, name string, description string, enabledValue string, disabledValue string, tags string, segments []string, order int64, value string) {
	fmt.Println("patchFeatureWithNumberValue")
	ruleArray, _ := appConfigurationServiceInstance.NewRule(segments)
	segmentRuleArray, _ := appConfigurationServiceInstance.NewSegmentRule([]appconfigurationv1.Rule{*ruleArray}, value, order)
	patchFeatureOptionsModel := appConfigurationServiceInstance.NewPatchFeatureOptions(id)
	patchFeatureOptionsModel.SetName(name)
	patchFeatureOptionsModel.SetDescription(description)
	patchFeatureOptionsModel.SetTags(tags)
	patchFeatureOptionsModel.SetDisabledValue(disabledValue)
	patchFeatureOptionsModel.SetEnabledValue(enabledValue)
	patchFeatureOptionsModel.SetSegmentRules([]appconfigurationv1.SegmentRule{*segmentRuleArray})
	result, response, error := appConfigurationServiceInstance.PatchFeature(patchFeatureOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Name)
}

func updateSegment(segmentId string, name string, description string, tags string, attributeName string, operator string, values []string) {
	fmt.Println("updateSegment")
	ruleArray, _ := appConfigurationServiceInstance.NewRuleArray(attributeName, operator, values)
	updateSegmentOptionsModel := appConfigurationServiceInstance.NewUpdateSegmentOptions(segmentId, name, description, tags, []appconfigurationv1.RuleArray{*ruleArray})
	result, response, error := appConfigurationServiceInstance.UpdateSegment(updateSegmentOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Name)
}

func updateProperty(name string, propertyId string, description string, valueOfProperty string, tags string, segments []string, collectionId string, order int64, value string) {
	fmt.Println("updateProperty")
	ruleArray2, _ := appConfigurationServiceInstance.NewRule(segments)
	segmentRuleArray, _ := appConfigurationServiceInstance.NewSegmentRule([]appconfigurationv1.Rule{*ruleArray2}, value, order)
	collectionArray, _ := appConfigurationServiceInstance.NewCollectionID(collectionId)
	updatePropertyOptionsModel := appConfigurationServiceInstance.NewUpdatePropertyOptions(propertyId)
	updatePropertyOptionsModel.SetName(name)
	updatePropertyOptionsModel.SetDescription(description)
	updatePropertyOptionsModel.SetTags(tags)
	updatePropertyOptionsModel.SetValue(valueOfProperty)
	updatePropertyOptionsModel.SetSegmentRules([]appconfigurationv1.SegmentRule{*segmentRuleArray})
	updatePropertyOptionsModel.SetCollections([]appconfigurationv1.CollectionID{*collectionArray})
	result, response, error := appConfigurationServiceInstance.UpdateProperty(updatePropertyOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.PropertyID)
}

func patchProperty(name string, propertyId string, description string, valueOfProperty string, tags string, segments []string, order int64, value string) {
	fmt.Println("patchProperty")
	ruleArray, _ := appConfigurationServiceInstance.NewRule(segments)
	segmentRuleArray, _ := appConfigurationServiceInstance.NewSegmentRule([]appconfigurationv1.Rule{*ruleArray}, value, order)
	patchPropertyOptionsModel := appConfigurationServiceInstance.NewPatchPropertyOptions(propertyId)
	patchPropertyOptionsModel.SetName(name)
	patchPropertyOptionsModel.SetDescription(description)
	patchPropertyOptionsModel.SetTags(tags)
	patchPropertyOptionsModel.SetValue(valueOfProperty)
	patchPropertyOptionsModel.SetSegmentRules([]appconfigurationv1.SegmentRule{*segmentRuleArray})
	result, response, error := appConfigurationServiceInstance.PatchProperty(patchPropertyOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.PropertyID)
}

// Delete examples

func deleteCollection(collectionId string) {
	fmt.Println("deleteCollection")
	deleteCollectionOptionsModel := appConfigurationServiceInstance.NewDeleteCollectionOptions(collectionId)
	response, error := appConfigurationServiceInstance.DeleteCollection(deleteCollectionOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
}

func deleteFeature(featureId string) {
	fmt.Println("deleteFeature")
	deleteFeatureOptionsModel := appConfigurationServiceInstance.NewDeleteafeatureOptions(featureId)
	response, error := appConfigurationServiceInstance.DeleteFeature(deleteFeatureOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
}

func deleteSegment(segmentId string) {
	fmt.Println("deleteSegment")
	deleteSegmentOptionsModel := appConfigurationServiceInstance.NewDeleteSegmentOptions(segmentId)
	response, error := appConfigurationServiceInstance.DeleteSegment(deleteSegmentOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
}

func deleteProperty(propertyId string) {
	fmt.Println("deleteProperty")
	deletePropertyOptionsModel := appConfigurationServiceInstance.NewDeletePropertyOptions(propertyId)
	response, error := appConfigurationServiceInstance.DeleteProperty(deletePropertyOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
}

// List examples

func getCollections() {
	fmt.Println("getCollections")
	getCollectionsOptionsModel := appConfigurationServiceInstance.NewGetCollectionsOptions()
	getCollectionsOptionsModel.SetExpand("true")
	result, response, error := appConfigurationServiceInstance.GetCollections(getCollectionsOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.PageInfo.Count)
}

func getFeatures() {
	fmt.Println("getFeatures")
	getFeaturesOptionsModel := appConfigurationServiceInstance.NewGetFeaturesOptions()
	result, response, error := appConfigurationServiceInstance.GetFeatures(getFeaturesOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.PageInfo.Count)
}

func getSegments() {
	fmt.Println("getSegments")
	getSegmentsOptionsModel := appConfigurationServiceInstance.NewGetSegmentsOptions()
	result, response, error := appConfigurationServiceInstance.GetSegments(getSegmentsOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.PageInfo.Count)
}

func getProperties() {
	fmt.Println("getProperties")
	getPropertiesOptionsModel := appConfigurationServiceInstance.NewGetPropertiesOptions()
	result, response, error := appConfigurationServiceInstance.GetProperties(getPropertiesOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.PageInfo.Count)
}

// Get examples

func getCollection(collectionId string) {
	fmt.Println("getCollection")
	getCollectionOptionsModel := appConfigurationServiceInstance.NewGetCollectionOptions(collectionId)
	result, response, error := appConfigurationServiceInstance.GetCollection(getCollectionOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Name)
}

func getFeature(featureId string) {
	fmt.Println("getFeature")
	getFeatureOptionsModel := appConfigurationServiceInstance.NewGetFeatureOptions(featureId)
	result, response, error := appConfigurationServiceInstance.GetFeature(getFeatureOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Name)        // referenced field, so needs to be de-referenced
	fmt.Println(result.EnabledValue) // non-referenced field, so can be used as it is
}

func getSegment(segmentId string) {
	fmt.Println("getSegment")
	getSegmentOptionsModel := appConfigurationServiceInstance.NewGetSegmentOptions(segmentId)
	result, response, error := appConfigurationServiceInstance.GetSegment(getSegmentOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Name)
}

func getProperty(propertyId string) {
	fmt.Println("getProperty")
	getPropertyOptionsModel := appConfigurationServiceInstance.NewGetPropertyOptions(propertyId)
	result, response, error := appConfigurationServiceInstance.GetProperty(getPropertyOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Name)
}

// Toggle Examples

func toggleFeature(featureId string, enableFlag bool) {
	fmt.Println("toggleFeature")
	toggleFeatureOptionsModel := appConfigurationServiceInstance.NewToggleFeatureOptions(featureId)
	toggleFeatureOptionsModel.SetEnabled(enableFlag)
	result, response, error := appConfigurationServiceInstance.ToggleFeature(toggleFeatureOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Name)
}
