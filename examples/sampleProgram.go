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
	"github.com/IBM/go-sdk-core/v4/core"
)

var appConfigurationServiceInstance *appconfigurationv1.AppConfigurationV1

func initAndReturnSingletonInstance(apiKey string, url string) *appconfigurationv1.AppConfigurationV1 {

	var once sync.Once
	if appConfigurationServiceInstance == nil {
		once.Do(func() {
			if appConfigurationServiceInstance == nil {
				authenticator := &core.IamAuthenticator{
					ApiKey: apiKey,
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
	apiKey := "<apikey>"
	url := "<url>"

	initAndReturnSingletonInstance(apiKey, url)

	createCollection("collectionId", "collectionName", "desc", "tags")
	GetCollections()
	createSegment("segmentName", "segmentId", "desc", "tags", "email", "endsWith", []string{"@in.ibm.com"})
	createFeatureWithBooleanValue("booleanFeatureName", "booleanFeatureId", "desc", "BOOLEAN", "true", "false", "tags", []string{"segmentId"}, 1, "collectionId", false, "true")
	getFeature("booleanFeatureId")
	createFeatureWithNumberValue("numberFeatureName", "numberFeatureId", "desc", "NUMERIC", "1", "2", "tags", []string{"segmentId"}, 1, "collectionId", true, "3")
	updateFeatureWithNumberValue("numberFeatureId", "numberFeatureName", "updatedDesc", "NUMERIC", "1", "1", "tags", []string{}, 1, "collectionId", false, "2", false)
	deleteSegment("segmentId")

}

func createCollection(id string, name string, description string, tags string) {
	createCollectionOptionsModel := appConfigurationServiceInstance.NewCreateCollectionOptions(name)
	createCollectionOptionsModel.SetCollectionID(id)
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

func GetCollections() {
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

func createSegment(name string, id string, description string, tags string, attributeName string, operator string, values []string) {
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

func getFeature(featureId string) {
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

func createFeatureWithBooleanValue(name string, id string, description string, typeOfFeature string, enabledValue string, disabledValue string, tags string, segments []string, order int64, collectionId string, enabledInCollection bool, value string) {
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

func createFeatureWithNumberValue(name string, id string, description string, typeOfFeature string, enabledValue string, disabledValue string, tags string, segments []string, order int64, collectionId string, enabledInCollection bool, value string) {
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

func updateFeatureWithNumberValue(id string, name string, description string, typeOfFeature string, enabledValue string, disabledValue string, tags string, segments []string, order int64, collectionName string, enabledInCollection bool, value string, deletedFlag bool) {
	ruleArray, _ := appConfigurationServiceInstance.NewRule(segments)
	segmentRuleArray, _ := appConfigurationServiceInstance.NewSegmentRule([]appconfigurationv1.Rule{*ruleArray}, value, order)
	collectionArray, _ := appConfigurationServiceInstance.NewCollectionWithDeletedFlag(collectionName, enabledInCollection, deletedFlag)
	updateFeatureOptionsModel := appConfigurationServiceInstance.NewUpdateFeatureOptions(id, name, description, typeOfFeature, enabledValue, disabledValue, tags, []appconfigurationv1.SegmentRule{*segmentRuleArray}, []appconfigurationv1.CollectionWithDeletedFlag{*collectionArray})
	result, response, error := appConfigurationServiceInstance.UpdateFeature(updateFeatureOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Name)
}

func deleteSegment(id string) {
	deleteasegmentOptionsModel := new(appconfigurationv1.DeleteSegmentOptions)
	deleteasegmentOptionsModel.SegmentID = core.StringPtr(id)
	response, error := appConfigurationServiceInstance.DeleteSegment(deleteasegmentOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
}
