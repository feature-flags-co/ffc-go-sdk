package ffc

import (
	"github.com/feature-flags-co/ffc-go-sdk/common"
	"github.com/feature-flags-co/ffc-go-sdk/datamodel"
	"log"
)

type Client struct {
	Offline   bool
	Evaluator Evaluator
}

// NewClient create a new client instance.
func NewClient(envSecret string, config *Config) Client {

	basicConfig := BasicConfig{OffLine: config.OffLine, EnvSecret: envSecret}
	contextConfig := Context{BasicConfig: basicConfig, HttpConfig: config.HttpConfig}
	stream := NewStreaming(contextConfig, config.StreamingBuilder.StreamingURI)
	go stream.Connect()

	// TODO init this Evaluator Object
	var evaluator Evaluator
	return Client{
		Offline:   config.OffLine,
		Evaluator: evaluator,
	}
}

// IsInitialized Tests whether the client is ready to be used.
// @Return true if the client is ready, or false if it is still initializing
func (c *Client) IsInitialized() bool {
	return datamodel.GetDataStorage().IsInitialized()
}

// VariationWithUser Calculates the value of a feature flag for a given user.
// @Param featureFlagKey the unique key for the feature flag
// @Param user the end user requesting the flag
// @Param defaultValue the default value of the flag
// @Return  the variation for the given user, or defaultValue if the flag is disabled or an error occurs
func (c *Client) VariationWithUser(featureFlagKey string, user common.FFCUser, defaultValue string) string {

	return ""
}

// Variation Calculates the value of a feature flag for a given user.
// @Param featureFlagKey the unique key for the feature flag
// @Param defaultValue the default value of the flag
// @Return  the variation for the given user, or defaultValue if the flag is disabled or an error occurs
func (c *Client) Variation(featureFlagKey string, defaultValue string) string {
	return ""
}

// BoolVariationWithUser Calculates the value of a feature flag for a given user.
// @Param featureFlagKey the unique key for the feature flag
// @Param user the end user requesting the flag
// @Param defaultValue the default value of the flag
// @Return  the variation for the given user, or defaultValue if the flag is disabled or an error occurs
func (c *Client) BoolVariationWithUser(featureFlagKey string, user common.FFCUser, defaultValue bool) string {

	return ""
}

// BoolVariation Calculates the value of a feature flag for a given user.
// @Param featureFlagKey the unique key for the feature flag
// @Param defaultValue the default value of the flag
// @Return  the variation for the given user, or defaultValue if the flag is disabled or an error occurs
func (c *Client) BoolVariation(featureFlagKey string, defaultValue bool) string {

	return ""
}

// IsEnableWithUser alias of boolVariation for a given user
// @Param featureFlagKey the unique key for the feature flag
// @Param user the end user requesting the flag
// @Return if the flag should be enabled, or false if the flag is disabled, or an error occurs
func (c *Client) IsEnableWithUser(featureFlagKey string, user common.FFCUser) string {

	return ""
}

// IsEnable alias of boolVariation for a given user
// @Param featureFlagKey the unique key for the feature flag
// @Return if the flag should be enabled, or false if the flag is disabled, or an error occurs
func (c *Client) IsEnable(featureFlagKey string) string {

	return ""
}

// FloatVariationWithUser Calculates the value of a feature flag for a given user.
// @Param featureFlagKey the unique key for the feature flag
// @Param user the end user requesting the flag
// @Param defaultValue the default value of the flag
// @Return  the variation for the given user, or defaultValue if the flag is disabled or an error occurs
func (c *Client) FloatVariationWithUser(featureFlagKey string, user common.FFCUser, defaultValue float64) string {

	return ""
}

// FloatVariation Calculates the value of a feature flag for a given user.
// @Param featureFlagKey the unique key for the feature flag
// @Param defaultValue the default value of the flag
// @Return  the variation for the given user, or defaultValue if the flag is disabled or an error occurs
func (c *Client) FloatVariation(featureFlagKey string, defaultValue float64) string {
	return ""
}

// IntVariationWithUser Calculates the value of a feature flag for a given user.
// @Param featureFlagKey the unique key for the feature flag
// @Param user the end user requesting the flag
// @Param defaultValue the default value of the flag
// @Return  the variation for the given user, or defaultValue if the flag is disabled or an error occurs
func (c *Client) IntVariationWithUser(featureFlagKey string, user common.FFCUser, defaultValue int) string {

	return ""
}

// IntVariation Calculates the value of a feature flag for a given user.
// @Param featureFlagKey the unique key for the feature flag
// @Param defaultValue the default value of the flag
// @Return  the variation for the given user, or defaultValue if the flag is disabled or an error occurs
func (c *Client) IntVariation(featureFlagKey string, defaultValue int) string {

	return ""
}

// Int64VariationWithUser Calculates the value of a feature flag for a given user.
// @Param featureFlagKey the unique key for the feature flag
// @Param user the end user requesting the flag
// @Param defaultValue the default value of the flag
// @Return  the variation for the given user, or defaultValue if the flag is disabled or an error occurs
func (c *Client) Int64VariationWithUser(featureFlagKey string, user common.FFCUser, defaultValue int64) string {

	return ""
}

// Int64Variation Calculates the value of a feature flag for a given user.
// @Param featureFlagKey the unique key for the feature flag
// @Param defaultValue the default value of the flag
// @Return  the variation for the given user, or defaultValue if the flag is disabled or an error occurs
func (c *Client) Int64Variation(featureFlagKey string, defaultValue int64) string {

	return ""
}

// IsFlagKnown Returns true if the specified feature flag currently exists.
// @Param featureFlagKey the unique key for the feature flag
// @Return true if the flag exists
func (c *Client) IsFlagKnown(featureFlagKey string) bool {

	return false
}

// InitializeFromExternalJson initialization in the offline mode
// @Param featureFlagKey the unique key for the feature flag
func (c *Client) InitializeFromExternalJson(featureFlagKey string) {

}

// GetAllLatestFlagsVariations  Returns a list of all feature flags value with details for a given user, including the reason
// that describes the way the value was determined, that can be used on the client side sdk or a front end .
// @Param user the end user requesting the flag
// @Return
func (c *Client) GetAllLatestFlagsVariations(user common.FFCUser) []common.AllFlagState {
	stats := make([]common.AllFlagState, 0)
	return stats
}

// GetAllUserTags return a list of user tags used to instantiate a {@link FFCUser}
// @Return a list of user tags
func (c *Client) GetAllUserTags() []common.UserTag {
	return []common.UserTag{}
}

// VariationDetailWithUser Calculates the value of a feature flag for a given user.
// @Param featureFlagKey the unique key for the feature flag
// @Param user the end user requesting the flag
// @Param defaultValue the default value of the flag
// @Return  the variation for the given user, or defaultValue if the flag is disabled or an error occurs
func (c *Client) VariationDetailWithUser(featureFlagKey string, user common.FFCUser, defaultValue string) common.FlagState {
	return common.FlagState{}
}

// VariationDetail Calculates the value of a feature flag for a given user.
// @Param featureFlagKey the unique key for the feature flag
// @Param defaultValue the default value of the flag
// @Return  the variation for the given user, or defaultValue if the flag is disabled or an error occurs
func (c *Client) VariationDetail(featureFlagKey string, defaultValue string) common.FlagState {
	return common.FlagState{}
}

// BoolVariationDetailWithUser Calculates the value of a feature flag for a given user.
// @Param featureFlagKey the unique key for the feature flag
// @Param user the end user requesting the flag
// @Param defaultValue the default value of the flag
// @Return  the variation for the given user, or defaultValue if the flag is disabled or an error occurs
func (c *Client) BoolVariationDetailWithUser(featureFlagKey string, user common.FFCUser,
	defaultValue bool) common.FlagState {
	return common.FlagState{}
}

// BoolVariationDetail Calculates the value of a feature flag for a given user.
// @Param featureFlagKey the unique key for the feature flag
// @Param defaultValue the default value of the flag
// @Return  the variation for the given user, or defaultValue if the flag is disabled or an error occurs
func (c *Client) BoolVariationDetail(featureFlagKey string, defaultValue bool) common.FlagState {
	return common.FlagState{}
}

// FloatVariationDetailWithUser Calculates the value of a feature flag for a given user.
// @Param featureFlagKey the unique key for the feature flag
// @Param user the end user requesting the flag
// @Param defaultValue the default value of the flag
// @Return  the variation for the given user, or defaultValue if the flag is disabled or an error occurs
func (c *Client) FloatVariationDetailWithUser(featureFlagKey string, user common.FFCUser,
	defaultValue float64) common.FlagState {
	return common.FlagState{}
}

// FloatVariationDetail Calculates the value of a feature flag for a given user.
// @Param featureFlagKey the unique key for the feature flag
// @Param defaultValue the default value of the flag
// @Return  the variation for the given user, or defaultValue if the flag is disabled or an error occurs
func (c *Client) FloatVariationDetail(featureFlagKey string, defaultValue float64) common.FlagState {
	return common.FlagState{}
}

// IntVariationDetailWithUser Calculates the value of a feature flag for a given user.
// @Param featureFlagKey the unique key for the feature flag
// @Param user the end user requesting the flag
// @Param defaultValue the default value of the flag
// @Return  the variation for the given user, or defaultValue if the flag is disabled or an error occurs
func (c *Client) IntVariationDetailWithUser(featureFlagKey string, user common.FFCUser,
	defaultValue int) common.FlagState {
	return common.FlagState{}
}

// IntVariationDetail Calculates the value of a feature flag for a given user.
// @Param featureFlagKey the unique key for the feature flag
// @Param defaultValue the default value of the flag
// @Return  the variation for the given user, or defaultValue if the flag is disabled or an error occurs
func (c *Client) IntVariationDetail(featureFlagKey string, defaultValue int) common.FlagState {
	return common.FlagState{}
}

// Int64VariationDetailWithUser Calculates the value of a feature flag for a given user.
// @Param featureFlagKey the unique key for the feature flag
// @Param user the end user requesting the flag
// @Param defaultValue the default value of the flag
// @Return  the variation for the given user, or defaultValue if the flag is disabled or an error occurs
func (c *Client) Int64VariationDetailWithUser(featureFlagKey string, user common.FFCUser,
	defaultValue int64) common.FlagState {
	return common.FlagState{}
}

// Int64VariationDetail Calculates the value of a feature flag for a given user.
// @Param featureFlagKey the unique key for the feature flag
// @Param defaultValue the default value of the flag
// @Return  the variation for the given user, or defaultValue if the flag is disabled or an error occurs
func (c *Client) Int64VariationDetail(featureFlagKey string, defaultValue int64) common.FlagState {
	return common.FlagState{}
}

// Flush  Flushes all pending events.
func (c *Client) Flush() {
}

// TrackMetricWithUser tracks that a user performed an event and provides a default numeric value for custom metrics
// @Param user  the user that performed the event
// @Param eventName the name of the event
func (c *Client) TrackMetricWithUser(user common.FFCUser, eventName string) {
}

// TrackMetric tracks that a user performed an event and provides a default numeric value for custom metrics
// @Param eventName the name of the event
func (c *Client) TrackMetric(eventName string) {
}

// TrackMetricWithUserAndValue tracks that a user performed an event and provides a default numeric value for custom metrics
// @Param user  the user that performed the event
// @Param eventName the name of the event
// @Param metricValue a numeric value used by the experimentation feature in numeric custom metrics.
func (c *Client) TrackMetricWithUserAndValue(user common.FFCUser, eventName string, metricValue float64) {
}

// TrackMetricWithValue tracks that a user performed an event and provides a default numeric value for custom metrics
// @Param eventName the name of the event
// @Param metricValue a numeric value used by the experimentation feature in numeric custom metrics.
func (c *Client) TrackMetricWithValue(eventName string, metricValue float64) {
}

// TrackMetricsWithUser tracks that a user performed an event and provides a default numeric value for custom metrics
// @Param user  the user that performed the event
// @Param eventName the name of the events
func (c *Client) TrackMetricsWithUser(user common.FFCUser, eventName ...string) {
}

// TrackMetrics tracks that a user performed an event and provides a default numeric value for custom metrics
// @Param user  the user that performed the event
// @Param eventName the name of the events
func (c *Client) TrackMetrics(eventName ...string) {
}

// TrackMetricSeriesWithUser tracks that a user performed an event and provides a default numeric value for custom metrics
// @Param user  the user that performed the event
// @Param metrics event name and numeric value in K/V
func (c *Client) TrackMetricSeriesWithUser(user common.FFCUser, metrics map[string]float64) {
}

// TrackMetricSeries tracks that a user performed an event and provides a default numeric value for custom metrics
// @Param metrics event name and numeric value in K/V
func (c *Client) TrackMetricSeries(metrics map[string]float64) {
}

func (c *Client) evaluateInternal(featureFlagKey string, user common.FFCUser, defaultValue interface{},
	checkType bool) datamodel.EvalResult {

	// not finish init data
	if !c.IsInitialized() {
		log.Print("FFC GO SDK: evaluation is called before GO SDK client is initialized for feature flag, " +
			"well using the default value")
		return datamodel.ErrorWithDefaultValue(defaultValue.(string),
			common.EvaReasonClientNotReady,
			featureFlagKey,
			common.EvaFlagNameUnknown)
	}

	// TODO 
	var featureFlag datamodel.FeatureFlag
	event := datamodel.OfFlagEvent(user)

	evaResult := c.Evaluator.Evaluate(featureFlag, user, &event)
	return evaResult
}
