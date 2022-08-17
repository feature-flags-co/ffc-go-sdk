package model

import (
	"fmt"
	"github.com/feature-flags-co/ffc-go-sdk/utils"
	"strings"
)

type FeatureFlagIdByEnvSecret struct {
	FeatureFlagId string
	EnvId         string
	AccountId     string
	ProjectId     string
}

func NewFeatureFlagIdByEnvSecret(envSecret string, featureFlagKeyName string) *FeatureFlagIdByEnvSecret {

	keyOriginText, _ := utils.Base64Decode(envSecret)
	keys := strings.Split(keyOriginText, "__")
	accountId := keys[1]
	projectId := keys[2]
	envId := keys[3]
	featureFlagId := buildFeatureFlagId(featureFlagKeyName, envId, accountId, projectId)
	return &FeatureFlagIdByEnvSecret{
		FeatureFlagId: featureFlagId,
		EnvId:         envId,
		AccountId:     accountId,
		ProjectId:     projectId,
	}
}

func (f *FeatureFlagIdByEnvSecret) GetFeatureFlagId() string {
	return f.FeatureFlagId
}

func buildFeatureFlagId(featureFlagKeyName string, envId string, accountId string, projectId string) string {
	return fmt.Sprintf("FF__%s__%s__%s__%s", accountId, projectId, envId, featureFlagKeyName)

}
