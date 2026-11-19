# ModifyRumRouteRankingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rank** | **int32** | Target rank for the RUM route, from 1 to the total number of RUM routes. Requests outside that range are rejected. | 

## Methods

### NewModifyRumRouteRankingRequest

`func NewModifyRumRouteRankingRequest(rank int32, ) *ModifyRumRouteRankingRequest`

NewModifyRumRouteRankingRequest instantiates a new ModifyRumRouteRankingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyRumRouteRankingRequestWithDefaults

`func NewModifyRumRouteRankingRequestWithDefaults() *ModifyRumRouteRankingRequest`

NewModifyRumRouteRankingRequestWithDefaults instantiates a new ModifyRumRouteRankingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRank

`func (o *ModifyRumRouteRankingRequest) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *ModifyRumRouteRankingRequest) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *ModifyRumRouteRankingRequest) SetRank(v int32)`

SetRank sets Rank field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


