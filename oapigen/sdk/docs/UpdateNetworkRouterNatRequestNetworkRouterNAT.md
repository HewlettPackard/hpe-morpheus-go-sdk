# UpdateNetworkRouterNatRequestNetworkRouterNAT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Sets name of NAT | [optional] 
**Action** | Pointer to **string** | The NAT action (e.g. SNAT, DNAT, REFLEXIVE). | [optional] 
**Description** | Pointer to **string** | Description of the NAT rule. | [optional] 
**Enabled** | Pointer to **bool** | Whether the NAT rule is enabled. | [optional] 
**SourceNetwork** | Pointer to **string** | Source network for the NAT rule. | [optional] 
**DestinationNetwork** | Pointer to **string** | Destination network for the NAT rule. | [optional] 
**TranslatedNetwork** | Pointer to **string** | Translated network for the NAT rule. | [optional] 
**Priority** | Pointer to **int64** | Priority of the NAT rule. | [optional] 
**Protocol** | Pointer to **string** | Protocol for the NAT rule. | [optional] 

## Usage

Instantiate with a Go composite literal:

```go
obj := &UpdateNetworkRouterNatRequestNetworkRouterNAT{
    // Set fields directly
}
```


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


