# CypherAccessPolicyTypeConfiguration2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyPattern** | **string** | Pattern to match Cypher keys (e.g., \&quot;secret/_*\&quot;, \&quot;password/_*\&quot;) | 
**Read** | Pointer to **bool** | Allow read access | [optional] 
**Write** | Pointer to **bool** | Allow write access | [optional] 
**Update** | Pointer to **bool** | Allow update access | [optional] 
**Delete** | Pointer to **bool** | Allow delete access | [optional] 
**List** | Pointer to **bool** | Allow list access | [optional] 

## Methods

### NewCypherAccessPolicyTypeConfiguration2

`func NewCypherAccessPolicyTypeConfiguration2(keyPattern string, ) *CypherAccessPolicyTypeConfiguration2`

NewCypherAccessPolicyTypeConfiguration2 instantiates a new CypherAccessPolicyTypeConfiguration2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCypherAccessPolicyTypeConfiguration2WithDefaults

`func NewCypherAccessPolicyTypeConfiguration2WithDefaults() *CypherAccessPolicyTypeConfiguration2`

NewCypherAccessPolicyTypeConfiguration2WithDefaults instantiates a new CypherAccessPolicyTypeConfiguration2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyPattern

`func (o *CypherAccessPolicyTypeConfiguration2) GetKeyPattern() string`

GetKeyPattern returns the KeyPattern field if non-nil, zero value otherwise.

### GetKeyPatternOk

`func (o *CypherAccessPolicyTypeConfiguration2) GetKeyPatternOk() (*string, bool)`

GetKeyPatternOk returns a tuple with the KeyPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPattern

`func (o *CypherAccessPolicyTypeConfiguration2) SetKeyPattern(v string)`

SetKeyPattern sets KeyPattern field to given value.


### GetRead

`func (o *CypherAccessPolicyTypeConfiguration2) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *CypherAccessPolicyTypeConfiguration2) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *CypherAccessPolicyTypeConfiguration2) SetRead(v bool)`

SetRead sets Read field to given value.

### HasRead

`func (o *CypherAccessPolicyTypeConfiguration2) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetWrite

`func (o *CypherAccessPolicyTypeConfiguration2) GetWrite() bool`

GetWrite returns the Write field if non-nil, zero value otherwise.

### GetWriteOk

`func (o *CypherAccessPolicyTypeConfiguration2) GetWriteOk() (*bool, bool)`

GetWriteOk returns a tuple with the Write field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrite

`func (o *CypherAccessPolicyTypeConfiguration2) SetWrite(v bool)`

SetWrite sets Write field to given value.

### HasWrite

`func (o *CypherAccessPolicyTypeConfiguration2) HasWrite() bool`

HasWrite returns a boolean if a field has been set.

### GetUpdate

`func (o *CypherAccessPolicyTypeConfiguration2) GetUpdate() bool`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *CypherAccessPolicyTypeConfiguration2) GetUpdateOk() (*bool, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *CypherAccessPolicyTypeConfiguration2) SetUpdate(v bool)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *CypherAccessPolicyTypeConfiguration2) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetDelete

`func (o *CypherAccessPolicyTypeConfiguration2) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *CypherAccessPolicyTypeConfiguration2) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *CypherAccessPolicyTypeConfiguration2) SetDelete(v bool)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *CypherAccessPolicyTypeConfiguration2) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetList

`func (o *CypherAccessPolicyTypeConfiguration2) GetList() bool`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *CypherAccessPolicyTypeConfiguration2) GetListOk() (*bool, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *CypherAccessPolicyTypeConfiguration2) SetList(v bool)`

SetList sets List field to given value.

### HasList

`func (o *CypherAccessPolicyTypeConfiguration2) HasList() bool`

HasList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


