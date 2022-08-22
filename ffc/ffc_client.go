package ffc

import (
	"github.com/feature-flags-co/ffc-go-sdk/data"
	"github.com/feature-flags-co/ffc-go-sdk/model"
	"github.com/feature-flags-co/ffc-go-sdk/utils"
	"log"
)

type Client struct {
	Offline     bool
	Evaluator   Evaluator
	EnvSecret   string
	dataStorage data.DataStorage
}

// NewClient create a new client instance.
func NewClient(envSecret string, config *Config) Client {

	basicConfig := BasicConfig{OffLine: config.OffLine, EnvSecret: envSecret}
	contextConfig := Context{BasicConfig: basicConfig, HttpConfig: config.HttpConfig}
	stream := NewStreaming(contextConfig, config.StreamingBuilder.StreamingURI)
	go stream.Connect()

	var evaluator Evaluator
	evaluator = NewEvaluator()
	return Client{
		Offline:     config.OffLine,
		Evaluator:   evaluator,
		EnvSecret:   envSecret,
		dataStorage: data.GetDataStorage(),
	}
}

// IsInitialized Tests whether the client is ready to be used.
// @Return true if the client is ready, or false if it is still initializing
func (c *Client) IsInitialized() bool {
	return c.dataStorage.IsInitialized()
}

// VariationWithUser Calculates the value of a feature flag for a given user.
// @Param featureFlagKey the unique key for the feature flag
// @Param user the end user requesting the flag
// @Param defaultValue the default value of the flag
// @Return  the variation for the given user, or defaultValue if the flag is disabled or an error occurs
func (c *Client) VariationWithUser(featureFlagKey string, user model.FFCUser, defaultValue string) string {
	evalResult := c.evaluateInternal(featureFlagKey, user, defaultValue, false)
	return evalResult.Value
}

// BoolVariationWithUser Calculates the value of a feature flag for a given user.
// @Param featureFlagKey the unique key for the feature flag
// @Param user the end user requesting the flag
// @Param defaultValue the default value of the flag
// @Return  the variation for the given user, or defaultValue if the flag is disabled or an error occurs
func (c *Client) BoolVariationWithUser(featureFlagKey string, user model.FFCUser, defaultValue bool) bool {
	evalResult := c.evaluateInternal(featureFlagKey, user, defaultValue, true)
	return utils.ToBool(evalResult.Value)
}

// IsEnableWithUser alias of boolVariation for a given user
// @Param featureFlagKey the unique key for the feature flag
// @Param user the end user requesting the flag
// @Return if the flag should be enabled, or false if the flag is disabled, or an error occurs
func (c *Client) IsEnableWithUser(featureFlagKey string, user model.FFCUser) bool {
	evalResult := c.evaluateInternal(featureFlagKey, user, false, true)
	return utils.ToBool(evalResult.Value)
}

// FloatVariationWithUser Calculates the value of a feature flag for a given user.
// @Param featureFlagKey the unique key for the feature flag
// @Param user the end user requesting the flag
// @Param defaultValue the default value of the flag
// @Return  the variation for the given user, or defaultValue if the flag is disabled or an error occurs
func (c *Client) FloatVariationWithUser(featureFlagKey string, user model.FFCUser, defaultValue float64) float64 {
	evalResult := c.evaluateInternal(featureFlagKey, user, defaultValue, true)
	return utils.GetFloat64(evalResult.Value)
}

// IntVariationWithUser Calculates the value of a feature flag for a given user.
// @Param featureFlagKey the unique key for the feature flag
// @Param user the end user requesting the flag
// @Param defaultValue the default value of the flag
// @Return  the variation for the given user, or defaultValue if the flag is disabled or an error occurs
func (c *Client) IntVariationWithUser(featureFlagKey string, user model.FFCUser, defaultValue int) int {
	evalResult := c.evaluateInternal(featureFlagKey, user, defaultValue, true)
	return utils.GetInt(evalResult.Value)
}

// Int64VariationWithUser Calculates the value of a feature flag for a given user.
// @Param featureFlagKey the unique key for the feature flag
// @Param user the end user requesting the flag
// @Param defaultValue the default value of the flag
// @Return  the variation for the given user, or defaultValue if the flag is disabled or an error occurs
func (c *Client) Int64VariationWithUser(featureFlagKey string, user model.FFCUser, defaultValue int64) int64 {
	evalResult := c.evaluateInternal(featureFlagKey, user, defaultValue, true)
	return utils.GetInt64(evalResult.Value)
}

// IsFlagKnown Returns true if the specified feature flag currently exists.
// @Param featureFlagKey the unique key for the feature flag
// @Return true if the flag exists
func (c *Client) IsFlagKnown(featureFlagKey string) bool {

	if !c.IsInitialized() {
		log.Printf("FFC GO SDK: isFlagKnown is called before Java SDK client is initialized for feature flag")
		return false
	}
	flag := c.getFlagInternal(featureFlagKey)
	return len(flag.Id) == 0
}

// InitializeFromExternalJson initialization in the offline mode
// @Param featureFlagKey the unique key for the feature flag
func (c *Client) InitializeFromExternalJson(featureFlagKey string) {

}

// GetAllLatestFlagsVariations  Returns a list of all feature flags value with details for a given user, including the reason
// that describes the way the value was determined, that can be used on the client side sdk or a front end .
// @Param user the end user requesting the flag
// @Return
func (c *Client) GetAllLatestFlagsVariations(user model.FFCUser) []model.AllFlagState {
	stats := make([]model.AllFlagState, 0)
	return stats
}

// GetAllUserTags return a list of user tags used to instantiate a {@link FFCUser}
// @Return a list of user tags
func (c *Client) GetAllUserTags() []model.UserTag {
	return []model.UserTag{}
}

// VariationDetailWithUser Calculates the value of a feature flag for a given user.
// @Param featureFlagKey the unique key for the feature flag
// @Param user the end user requesting the flag
// @Param defaultValue the default value of the flag
// @Return  the variation for the given user, or defaultValue if the flag is disabled or an error occurs
func (c *Client) VariationDetailWithUser(featureFlagKey string, user model.FFCUser, defaultValue string) model.FlagState {
	return model.FlagState{}
}

// BoolVariationDetailWithUser Calculates the value of a feature flag for a given user.
// @Param featureFlagKey the unique key for the feature flag
// @Param user the end user requesting the flag
// @Param defaultValue the default value of the flag
// @Return  the variation for the given user, or defaultValue if the flag is disabled or an error occurs
func (c *Client) BoolVariationDetailWithUser(featureFlagKey string, user model.FFCUser,
	defaultValue bool) model.FlagState {
	return model.FlagState{}
}

// FloatVariationDetailWithUser Calculates the value of a feature flag for a given user.
// @Param featureFlagKey the unique key for the feature flag
// @Param user the end user requesting the flag
// @Param defaultValue the default value of the flag
// @Return  the variation for the given user, or defaultValue if the flag is disabled or an error occurs
func (c *Client) FloatVariationDetailWithUser(featureFlagKey string, user model.FFCUser,
	defaultValue float64) model.FlagState {
	return model.FlagState{}
}

// IntVariationDetailWithUser Calculates the value of a feature flag for a given user.
// @Param featureFlagKey the unique key for the feature flag
// @Param user the end user requesting the flag
// @Param defaultValue the default value of the flag
// @Return  the variation for the given user, or defaultValue if the flag is disabled or an error occurs
func (c *Client) IntVariationDetailWithUser(featureFlagKey string, user model.FFCUser,
	defaultValue int) model.FlagState {
	return model.FlagState{}
}

// Int64VariationDetailWithUser Calculates the value of a feature flag for a given user.
// @Param featureFlagKey the unique key for the feature flag
// @Param user the end user requesting the flag
// @Param defaultValue the default value of the flag
// @Return  the variation for the given user, or defaultValue if the flag is disabled or an error occurs
func (c *Client) Int64VariationDetailWithUser(featureFlagKey string, user model.FFCUser,
	defaultValue int64) model.FlagState {
	return model.FlagState{}
}

// Flush  Flushes all pending events.
func (c *Client) Flush() {
}

// TrackMetricWithUser tracks that a user performed an event and provides a default numeric value for custom metrics
// @Param user  the user that performed the event
// @Param eventName the name of the event
func (c *Client) TrackMetricWithUser(user model.FFCUser, eventName string) {
}

// TrackMetric tracks that a user performed an event and provides a default numeric value for custom metrics
// @Param eventName the name of the event
func (c *Client) TrackMetric(eventName string) {
}

// TrackMetricWithUserAndValue tracks that a user performed an event and provides a default numeric value for custom metrics
// @Param user  the user that performed the event
// @Param eventName the name of the event
// @Param metricValue a numeric value used by the experimentation feature in numeric custom metrics.
func (c *Client) TrackMetricWithUserAndValue(user model.FFCUser, eventName string, metricValue float64) {
}

// TrackMetricWithValue tracks that a user performed an event and provides a default numeric value for custom metrics
// @Param eventName the name of the event
// @Param metricValue a numeric value used by the experimentation feature in numeric custom metrics.
func (c *Client) TrackMetricWithValue(eventName string, metricValue float64) {
}

// TrackMetricsWithUser tracks that a user performed an event and provides a default numeric value for custom metrics
// @Param user  the user that performed the event
// @Param eventName the name of the events
func (c *Client) TrackMetricsWithUser(user model.FFCUser, eventName ...string) {
}

// TrackMetrics tracks that a user performed an event and provides a default numeric value for custom metrics
// @Param user  the user that performed the event
// @Param eventName the name of the events
func (c *Client) TrackMetrics(eventName ...string) {
}

// TrackMetricSeriesWithUser tracks that a user performed an event and provides a default numeric value for custom metrics
// @Param user  the user that performed the event
// @Param metrics event name and numeric value in K/V
func (c *Client) TrackMetricSeriesWithUser(user model.FFCUser, metrics map[string]float64) {
}

// TrackMetricSeries tracks that a user performed an event and provides a default numeric value for custom metrics
// @Param metrics event name and numeric value in K/V
func (c *Client) TrackMetricSeries(metrics map[string]float64) {
}

func (c *Client) evaluateInternal(featureFlagKey string, user model.FFCUser, defaultValue interface{},
	checkType bool) data.EvalResult {

	// not finish init data
	if !c.IsInitialized() {
		log.Print("FFC GO SDK: evaluation is called before GO SDK client is initialized for feature flag, " +
			"well using the default value")
		return data.ErrorWithDefaultValue(defaultValue.(string),
			model.EvaReasonClientNotReady,
			featureFlagKey,
			model.EvaFlagNameUnknown)
	}

	// featureFlagKey is blank
	if len(featureFlagKey) == 0 {
		log.Print("FFC GO SDK:null feature flag key; returning default value")
		data.ErrorWithDefaultValue(defaultValue.(string),
			model.EvaReasonFlagNotFound,
			featureFlagKey,
			model.EvaFlagNameUnknown)
	}
	featureFlag := c.getFlagInternal(featureFlagKey)
	if len(featureFlag.Id) == 0 {
		log.Printf("FFC GO SDK:unknown feature flag %s; returning default value", featureFlagKey)
		data.ErrorWithDefaultValue(defaultValue.(string),
			model.EvaReasonFlagNotFound,
			featureFlagKey,
			model.EvaFlagNameUnknown)
	}

	if len(user.UserName) == 0 {
		log.Printf("FFC GO SDK:null user for feature flag  %s; returning default value", featureFlagKey)
		data.ErrorWithDefaultValue(defaultValue.(string),
			model.EvaReasonUserNotSpecified,
			featureFlagKey,
			model.EvaFlagNameUnknown)
	}

	event := data.OfFlagEvent(user)
	evaResult := c.Evaluator.Evaluate(featureFlag, user, &event)

	if checkType {
		log.Printf("FFC GO SDK:evaluation result %s didn't matched expected type ", evaResult.Value)
		data.ErrorWithDefaultValue(defaultValue.(string),
			model.EvaReasonWrongType,
			evaResult.KeyName,
			evaResult.Name)
	}

	// TODO
	//eventHandler.accept(event);

	er := data.EvalResult{
		Index:            evaResult.Index,
		Value:            evaResult.Value,
		Reason:           evaResult.Reason,
		SendToExperiment: evaResult.SendToExperiment,
		KeyName:          evaResult.KeyName,
		Name:             evaResult.Name,
	}
	return er
}

func (c *Client) getFlagInternal(featureFlagKey string) data.FeatureFlag {
	flagId := data.NewFeatureFlagIdByEnvSecret(c.EnvSecret, featureFlagKey).GetFeatureFlagId()
	item := c.dataStorage.Get(data.FeaturesCat, flagId)
	if item == (data.Item{}) {
		return data.FeatureFlag{}
	}
	return item.Item.(data.FeatureFlag)

}
