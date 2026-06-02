# CreateLoadBalancerPoolRequestLoadBalancerPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**VipBalance** | Pointer to **string** | Balance Algorithm | [optional] 
**MinActive** | Pointer to **int64** | Min Active Members | [optional] 
**Config** | Pointer to **map[string]interface{}** | Configuration object with parameters that vary by type. | [optional] 

## Usage

Instantiate with a Go composite literal:

```go
obj := &CreateLoadBalancerPoolRequestLoadBalancerPool{
    // Set fields directly
}
```


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


