# Go Server Side SDK

## Introduction

This is the Go Server Side SDK for the feature management platform [featureflag.co](https://featureflag.co/). It is
intended for use in a multi-user Go server applications.

This SDK has two main purposes:

- Store the available feature flags and evaluate the feature flag variation for a given user
- Send feature flag usage and custom events for the insights and A/B/n testing.

## Data synchonization

We use websocket to make the local data synchronized with the server, and then store them in memory by default. Whenever there is any change to a feature flag or its related data, this change will be pushed to the SDK, the average synchronization time is less than **100** ms. Be aware the websocket connection may be interrupted due to internet outage, but it will be resumed automatically once the problem is gone.

## Offline mode support

In the offline mode, SDK DOES NOT exchange any data with [featureflag.co](https://featureflag.co/)

In the following situation, the SDK would work when there is no internet connection: it has been initialized in
using `ffc.InitializeFromExternalJson(json)` function

To open the offline mode:
```go
config := ffc.DefaultFFCConfigBuilder().
		SetOffline(false).
		Build()
client = ffc.NewClient(envSecret, config)
```
## Evaluation of a feature flag

SDK will initialize all the related data(feature flags, segments etc.) in the bootstrapping and receive the data updates
in real time, as mentioned in the above.

After initialization, the SDK has all the feature flags in the memory and all evaluation is done locally and
synchronously, the average evaluation time is < **10** ms.

## SDK

### FFCClient

Applications SHOULD instantiate a single instance for the lifetime of the application. In the case where an application
needs to evaluate feature flags from different environments, you may create multiple clients, but they should still be
retained for the lifetime of the application rather than created per request or per thread.

### Bootstrapping

The bootstrapping is in fact call this function `ffc.NewClient(envSecret, config)`, in which the SDK will be 
initialized, using streaming from [featureflag.co](https://featureflag.co/).


```go
client = ffc.NewClient(envSecret, config)
if(client.IsInitialized()){
  // do whatever is appropriate
}
```

Note that the _**sdkKey(envSecret)**_ is mandatory.

### FFCConfig and Components

`FFCConfig` exposes advanced configuration options for the `FFCClient`.

`startWaitTime`: how long the constructor will block awaiting a successful data sync. Setting this to a zero or negative
duration will not block and cause the constructor to return immediately.

`offline`: Set whether SDK is offline. when set to true no connection to feature-flag.co anymore

We strongly recommend to use the default configuration or just set `startWaitTime` or `offline` if necessary.

```go
// default configuration
config = ffc.DefaultFFCConfig()
client = ffc.NewClient(envSecret, config)
```

### Evaluation

SDK calculates the value of a feature flag for a given user, and returns a flag vlaue/an object that describes the way 
that the value was determined.

`FFUser`: A collection of attributes that can affect flag evaluation, usually corresponding to a user of your application.
This object contains built-in properties(`key`, `userName`, `email` and `country`). The `key` and `userName` are required.
The `key` must uniquely identify each user; this could be a username or email address for authenticated users, or a ID for anonymous users.
The `userName` is used to search your user quickly. All other built-in properties are optional, you may also define custom properties with arbitrary names and values.

```go
client = ffc.NewClient(envSecret, config)

// FFUser creation
ffcUser := model.NewFFUserBuilder().
		UserName("userName").
		Country("country").
		Email("email").
		Custom("key", "value").Build()

// be sure that SDK is initialized
// this is not required
if(client.isInitialized()){
// Evaluation details
flagtStatue :=client.VariationDetail("featureFlagKey",ffcUser,"defaultValue")
// Flag value
res := client.Variation("flag key", ffcUser, "Not Found");
// get all variations for a given user
userTags :=client.GetAllLatestFlagsVariations(ffcUser)
```

If evaluation called before Go SDK client initialized or you set the wrong flag key or user for the evaluation, SDK will return 
the default value you set. The `FlagState` and `AllFlagStates` will all details of later evaluation including the error reason.

SDK supports String, Boolean, and Number as the return type of flag values.
