package model

import (
	"encoding/json"
	"log"
)

// VariationParams a ffc object is used to pass the FFClient
// and flag key name to Server SDK Wrapped API
type VariationParams struct {
	FeatureFlagKeyName string
	User               FFCUser
	NeedAll            bool
}

func OfVariationParams(featureFlagKeyName string, user FFCUser) VariationParams {
	return VariationParams{
		FeatureFlagKeyName: featureFlagKeyName,
		User:               user,
		NeedAll:            len(featureFlagKeyName) > 0,
	}
}

// FromJson build a VariationParams object from json string
// @Param jsonstr a json string
// @Return a VariationParams object
func (v *VariationParams) FromJson(jsonstr string) VariationParams {

	var vp VariationParams
	err := json.Unmarshal([]byte(jsonstr), &vp)

	if err != nil {
		log.Fatalf("convert json string to VariationParams object error, error = %v", err)
		return VariationParams{}
	}
	return vp
}

// Jsonfy serialize a VariationParams to json string
// @return json string
func (v *VariationParams) Jsonfy() string {

	data, err := json.Marshal(v)
	if err != nil {
		log.Fatalf("convert  VariationParams object to string error, error = %v", err)
		return ""
	}
	return string(data)
}
