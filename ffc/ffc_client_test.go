package ffc

import (
	"fmt"
	"github.com/feature-flags-co/ffc-go-sdk/data"
	"github.com/feature-flags-co/ffc-go-sdk/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

var client Client
var ffcUser model.FFCUser

var testJson = "{\"messageType\":\"data-sync\",\"data\":{\"eventType\":\"full\",\"featureFlags\":[{\"_id\":\"FF__177__206__416__testflag\",\"id\":\"FF__177__206__416__testflag\",\"environmentId\":416,\"isArchived\":false,\"ff\":{\"id\":\"FF__177__206__416__testflag\",\"name\":\"testflag\",\"type\":1,\"keyName\":\"testflag\",\"environmentId\":416,\"creatorUserId\":\"62c6ef20b7b5bd6e556e62ea\",\"status\":\"Enabled\",\"isDefaultRulePercentageRolloutsIncludedInExpt\":null,\"lastUpdatedTime\":\"2022-09-13T12:16:55.2879106Z\",\"defaultRulePercentageRollouts\":[{\"exptRollout\":null,\"rolloutPercentage\":[0,1],\"valueOption\":{\"localId\":2,\"displayOrder\":2,\"variationValue\":\"false\"}}],\"variationOptionWhenDisabled\":{\"localId\":2,\"displayOrder\":2,\"variationValue\":\"false\"}},\"ffp\":[],\"fftuwmtr\":[{\"ruleId\":\"660f5428-c446-4340-8c5d-72f4d4f88ec1\",\"ruleName\":\"\\u89C4\\u52191\",\"isIncludedInExpt\":null,\"ruleJsonContent\":[{\"property\":\"User is in segment\",\"operation\":null,\"value\":\"[\\u00226319b0609d4600dec6dfca7c\\u0022]\"}],\"valueOptionsVariationRuleValues\":[{\"exptRollout\":null,\"rolloutPercentage\":[0,1],\"valueOption\":{\"localId\":1,\"displayOrder\":1,\"variationValue\":\"true\"}}]}],\"targetIndividuals\":[{\"individuals\":[{\"id\":\"WU__416__zttt\",\"name\":\"zttt\",\"keyId\":\"zttt\",\"email\":null}],\"valueOption\":{\"localId\":1,\"displayOrder\":1,\"variationValue\":\"true\"}},{\"individuals\":[{\"id\":\"WU__416__oll\",\"name\":\"oll\",\"keyId\":\"oll\",\"email\":null}],\"valueOption\":{\"localId\":2,\"displayOrder\":2,\"variationValue\":\"false\"}}],\"variationOptions\":[{\"localId\":1,\"displayOrder\":1,\"variationValue\":\"true\"},{\"localId\":2,\"displayOrder\":2,\"variationValue\":\"false\"}],\"exptIncludeAllRules\":true,\"timestamp\":1663071415287},{\"_id\":\"FF__177__206__416__\\u7F13\\u5B58\\u65F6\\u95F4\",\"id\":\"FF__177__206__416__\\u7F13\\u5B58\\u65F6\\u95F4\",\"environmentId\":416,\"isArchived\":false,\"ff\":{\"id\":\"FF__177__206__416__\\u7F13\\u5B58\\u65F6\\u95F4\",\"name\":\"\\u7F13\\u5B58\\u65F6\\u95F4\",\"type\":1,\"keyName\":\"\\u7F13\\u5B58\\u65F6\\u95F4\",\"environmentId\":416,\"creatorUserId\":\"62c6ef20b7b5bd6e556e62ea\",\"status\":\"Enabled\",\"isDefaultRulePercentageRolloutsIncludedInExpt\":null,\"lastUpdatedTime\":\"2022-09-08T09:26:06.7903472Z\",\"defaultRulePercentageRollouts\":[{\"exptRollout\":null,\"rolloutPercentage\":[0,1],\"valueOption\":{\"localId\":1,\"displayOrder\":1,\"variationValue\":\"true\"}}],\"variationOptionWhenDisabled\":{\"localId\":2,\"displayOrder\":2,\"variationValue\":\"false\"}},\"ffp\":[],\"fftuwmtr\":[],\"targetIndividuals\":[{\"individuals\":[{\"id\":\"WU__416__oll\",\"name\":\"oll\",\"keyId\":\"oll\",\"email\":null}],\"valueOption\":{\"localId\":1,\"displayOrder\":1,\"variationValue\":\"true\"}},{\"individuals\":[],\"valueOption\":{\"localId\":2,\"displayOrder\":2,\"variationValue\":\"false\"}}],\"variationOptions\":[{\"localId\":1,\"displayOrder\":1,\"variationValue\":\"true\"},{\"localId\":2,\"displayOrder\":2,\"variationValue\":\"false\"}],\"exptIncludeAllRules\":true,\"timestamp\":1662629166790},{\"_id\":\"FF__177__206__416__\\u6807\\u9898\",\"id\":\"FF__177__206__416__\\u6807\\u9898\",\"environmentId\":416,\"isArchived\":false,\"ff\":{\"id\":\"FF__177__206__416__\\u6807\\u9898\",\"name\":\"\\u6807\\u9898\",\"type\":1,\"keyName\":\"\\u6807\\u9898\",\"environmentId\":416,\"creatorUserId\":\"62c6ef20b7b5bd6e556e62ea\",\"status\":\"Enabled\",\"isDefaultRulePercentageRolloutsIncludedInExpt\":null,\"lastUpdatedTime\":\"2022-07-07T14:35:37.322Z\",\"defaultRulePercentageRollouts\":[{\"exptRollout\":null,\"rolloutPercentage\":[0,1],\"valueOption\":{\"localId\":2,\"displayOrder\":2,\"variationValue\":\"false\"}}],\"variationOptionWhenDisabled\":{\"localId\":2,\"displayOrder\":2,\"variationValue\":\"false\"}},\"ffp\":[],\"fftuwmtr\":[],\"targetIndividuals\":[],\"variationOptions\":[{\"localId\":1,\"displayOrder\":1,\"variationValue\":\"true\"},{\"localId\":2,\"displayOrder\":2,\"variationValue\":\"false\"}],\"exptIncludeAllRules\":null,\"timestamp\":1657204537322}],\"segments\":[{\"id\":\"6319b0609d4600dec6dfca7c\",\"included\":[\"zttt\"],\"excluded\":[\"oll\"],\"rules\":[],\"isArchived\":false,\"timestamp\":1662628949954}],\"userTags\":[]}}\n"

func newClient() {
	//envSecret := "ZDMzLTY3NDEtNCUyMDIxMTAxNzIxNTYyNV9fMzZfXzQ2X185OF9fZGVmYXVsdF80ODEwNA=="
	envSecret := "NWM4LTAzODgtNCUyMDIyMDcwNzE0MzUzN19fMTc3X18yMDZfXzQxNl9fZGVmYXVsdF8zNDY2Yw=="
	//streamingBuilder := ffc.NewStreamingBuilder().NewStreamingURI("wss://api-dev.featureflag.co")

	config := NewConfigBuilder().
		SetOffline(true).
		//SetUpdateProcessorFactory(streamingBuilder).
		Build()
	client = NewClient(envSecret, config)
	fmt.Println(client)

	ffcUser = model.NewFFUserBuilder().
		UserName("zttt").
		Key("zttt").
		Country("country").
		Email("email").
		Custom("key", "value").Build()

	client.InitializeFromExternalJson(testJson)
}

func TestMain(m *testing.M) {
	newClient()
	m.Run()
}

func Test_DataStorage(t *testing.T) {

	var isOk bool
	isOk = client.InitializeFromExternalJson(testJson)
	assert.EqualValues(t, true, isOk)

	isOk = client.dataStorage.IsInitialized()
	assert.EqualValues(t, true, isOk)

	featureFlags := client.dataStorage.GetAll(data.FeaturesCat)
	assert.EqualValues(t, 3, len(featureFlags))

	tags := client.GetAllUserTags()
	assert.EqualValues(t, 0, len(tags))
}

func Test_ClientApis(t *testing.T) {

	client.GetAllLatestFlagsVariations(ffcUser)

	tags := client.GetAllUserTags()
	assert.EqualValues(t, 2, len(tags))
}
