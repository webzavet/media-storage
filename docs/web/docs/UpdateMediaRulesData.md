# UpdateMediaRulesData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | media rules id | 
**Type** | **string** |  | 
**Attributes** | [**UpdateMediaRulesDataAttributes**](UpdateMediaRulesDataAttributes.md) |  | 

## Methods

### NewUpdateMediaRulesData

`func NewUpdateMediaRulesData(id string, type_ string, attributes UpdateMediaRulesDataAttributes, ) *UpdateMediaRulesData`

NewUpdateMediaRulesData instantiates a new UpdateMediaRulesData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMediaRulesDataWithDefaults

`func NewUpdateMediaRulesDataWithDefaults() *UpdateMediaRulesData`

NewUpdateMediaRulesDataWithDefaults instantiates a new UpdateMediaRulesData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateMediaRulesData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateMediaRulesData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateMediaRulesData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *UpdateMediaRulesData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateMediaRulesData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateMediaRulesData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *UpdateMediaRulesData) GetAttributes() UpdateMediaRulesDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpdateMediaRulesData) GetAttributesOk() (*UpdateMediaRulesDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UpdateMediaRulesData) SetAttributes(v UpdateMediaRulesDataAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


