# UpdateWhitelabelSettingsRequestWhitelabelSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Can be used to enable / disable whitelabel feature | [optional] 
**ApplianceName** | Pointer to **string** | Appliance name. Master account only | [optional] 
**DisableSupportMenu** | Pointer to **bool** | Can be used to disable support menu | [optional] 
**ResetHeaderLogo** | Pointer to **bool** | Resets header logo to default header logo | [optional] 
**ResetFooterLogo** | Pointer to **bool** | Resets footer logo to default footer logo | [optional] 
**ResetLoginLogo** | Pointer to **bool** | Resets login logo to default login logo | [optional] 
**ResetFavicon** | Pointer to **bool** | Resets favicon to default favicon | [optional] 
**HeaderBgColor** | Pointer to **string** | Header background color | [optional] 
**HeaderFgColor** | Pointer to **string** | Header foreground color | [optional] 
**NavBgColor** | Pointer to **string** | Nav background color | [optional] 
**NavFgColor** | Pointer to **string** | Nav foreground color | [optional] 
**NavHoverColor** | Pointer to **string** | Nav hover color | [optional] 
**PrimaryButtonBgColor** | Pointer to **string** | Primary button background color | [optional] 
**PrimaryButtonFgColor** | Pointer to **string** | Primary button foreground color | [optional] 
**PrimaryButtonHoverBgColor** | Pointer to **string** | Primary button hover background color | [optional] 
**PrimaryButtonHoverFgColor** | Pointer to **string** | Primary button hover foreground color | [optional] 
**FooterBgColor** | Pointer to **string** | Footer background color | [optional] 
**FooterFgColor** | Pointer to **string** | Footer foreground color | [optional] 
**LoginBgColor** | Pointer to **string** | Login background color | [optional] 
**CopyrightString** | Pointer to **string** | Copyright String | [optional] 
**OverrideCss** | Pointer to **string** | Override CSS | [optional] 
**TermsOfUse** | Pointer to **string** | Terms of use content | [optional] 
**PrivacyPolicy** | Pointer to **string** | Privacy policy content | [optional] 
**SupportMenuLinks** | Pointer to [**[]UpdateWhitelabelSettingsRequestWhitelabelSettingsSupportMenuLinksInner**](UpdateWhitelabelSettingsRequestWhitelabelSettingsSupportMenuLinksInner.md) |  | [optional] 

## Methods

### NewUpdateWhitelabelSettingsRequestWhitelabelSettings

`func NewUpdateWhitelabelSettingsRequestWhitelabelSettings() *UpdateWhitelabelSettingsRequestWhitelabelSettings`

NewUpdateWhitelabelSettingsRequestWhitelabelSettings instantiates a new UpdateWhitelabelSettingsRequestWhitelabelSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWhitelabelSettingsRequestWhitelabelSettingsWithDefaults

`func NewUpdateWhitelabelSettingsRequestWhitelabelSettingsWithDefaults() *UpdateWhitelabelSettingsRequestWhitelabelSettings`

NewUpdateWhitelabelSettingsRequestWhitelabelSettingsWithDefaults instantiates a new UpdateWhitelabelSettingsRequestWhitelabelSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetApplianceName

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetApplianceName() string`

GetApplianceName returns the ApplianceName field if non-nil, zero value otherwise.

### GetApplianceNameOk

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetApplianceNameOk() (*string, bool)`

GetApplianceNameOk returns a tuple with the ApplianceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceName

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) SetApplianceName(v string)`

SetApplianceName sets ApplianceName field to given value.

### HasApplianceName

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) HasApplianceName() bool`

HasApplianceName returns a boolean if a field has been set.

### GetDisableSupportMenu

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetDisableSupportMenu() bool`

GetDisableSupportMenu returns the DisableSupportMenu field if non-nil, zero value otherwise.

### GetDisableSupportMenuOk

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetDisableSupportMenuOk() (*bool, bool)`

GetDisableSupportMenuOk returns a tuple with the DisableSupportMenu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSupportMenu

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) SetDisableSupportMenu(v bool)`

SetDisableSupportMenu sets DisableSupportMenu field to given value.

### HasDisableSupportMenu

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) HasDisableSupportMenu() bool`

HasDisableSupportMenu returns a boolean if a field has been set.

### GetResetHeaderLogo

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetResetHeaderLogo() bool`

GetResetHeaderLogo returns the ResetHeaderLogo field if non-nil, zero value otherwise.

### GetResetHeaderLogoOk

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetResetHeaderLogoOk() (*bool, bool)`

GetResetHeaderLogoOk returns a tuple with the ResetHeaderLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetHeaderLogo

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) SetResetHeaderLogo(v bool)`

SetResetHeaderLogo sets ResetHeaderLogo field to given value.

### HasResetHeaderLogo

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) HasResetHeaderLogo() bool`

HasResetHeaderLogo returns a boolean if a field has been set.

### GetResetFooterLogo

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetResetFooterLogo() bool`

GetResetFooterLogo returns the ResetFooterLogo field if non-nil, zero value otherwise.

### GetResetFooterLogoOk

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetResetFooterLogoOk() (*bool, bool)`

GetResetFooterLogoOk returns a tuple with the ResetFooterLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetFooterLogo

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) SetResetFooterLogo(v bool)`

SetResetFooterLogo sets ResetFooterLogo field to given value.

### HasResetFooterLogo

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) HasResetFooterLogo() bool`

HasResetFooterLogo returns a boolean if a field has been set.

### GetResetLoginLogo

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetResetLoginLogo() bool`

GetResetLoginLogo returns the ResetLoginLogo field if non-nil, zero value otherwise.

### GetResetLoginLogoOk

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetResetLoginLogoOk() (*bool, bool)`

GetResetLoginLogoOk returns a tuple with the ResetLoginLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetLoginLogo

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) SetResetLoginLogo(v bool)`

SetResetLoginLogo sets ResetLoginLogo field to given value.

### HasResetLoginLogo

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) HasResetLoginLogo() bool`

HasResetLoginLogo returns a boolean if a field has been set.

### GetResetFavicon

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetResetFavicon() bool`

GetResetFavicon returns the ResetFavicon field if non-nil, zero value otherwise.

### GetResetFaviconOk

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetResetFaviconOk() (*bool, bool)`

GetResetFaviconOk returns a tuple with the ResetFavicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetFavicon

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) SetResetFavicon(v bool)`

SetResetFavicon sets ResetFavicon field to given value.

### HasResetFavicon

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) HasResetFavicon() bool`

HasResetFavicon returns a boolean if a field has been set.

### GetHeaderBgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetHeaderBgColor() string`

GetHeaderBgColor returns the HeaderBgColor field if non-nil, zero value otherwise.

### GetHeaderBgColorOk

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetHeaderBgColorOk() (*string, bool)`

GetHeaderBgColorOk returns a tuple with the HeaderBgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderBgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) SetHeaderBgColor(v string)`

SetHeaderBgColor sets HeaderBgColor field to given value.

### HasHeaderBgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) HasHeaderBgColor() bool`

HasHeaderBgColor returns a boolean if a field has been set.

### GetHeaderFgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetHeaderFgColor() string`

GetHeaderFgColor returns the HeaderFgColor field if non-nil, zero value otherwise.

### GetHeaderFgColorOk

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetHeaderFgColorOk() (*string, bool)`

GetHeaderFgColorOk returns a tuple with the HeaderFgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderFgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) SetHeaderFgColor(v string)`

SetHeaderFgColor sets HeaderFgColor field to given value.

### HasHeaderFgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) HasHeaderFgColor() bool`

HasHeaderFgColor returns a boolean if a field has been set.

### GetNavBgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetNavBgColor() string`

GetNavBgColor returns the NavBgColor field if non-nil, zero value otherwise.

### GetNavBgColorOk

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetNavBgColorOk() (*string, bool)`

GetNavBgColorOk returns a tuple with the NavBgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavBgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) SetNavBgColor(v string)`

SetNavBgColor sets NavBgColor field to given value.

### HasNavBgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) HasNavBgColor() bool`

HasNavBgColor returns a boolean if a field has been set.

### GetNavFgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetNavFgColor() string`

GetNavFgColor returns the NavFgColor field if non-nil, zero value otherwise.

### GetNavFgColorOk

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetNavFgColorOk() (*string, bool)`

GetNavFgColorOk returns a tuple with the NavFgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavFgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) SetNavFgColor(v string)`

SetNavFgColor sets NavFgColor field to given value.

### HasNavFgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) HasNavFgColor() bool`

HasNavFgColor returns a boolean if a field has been set.

### GetNavHoverColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetNavHoverColor() string`

GetNavHoverColor returns the NavHoverColor field if non-nil, zero value otherwise.

### GetNavHoverColorOk

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetNavHoverColorOk() (*string, bool)`

GetNavHoverColorOk returns a tuple with the NavHoverColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavHoverColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) SetNavHoverColor(v string)`

SetNavHoverColor sets NavHoverColor field to given value.

### HasNavHoverColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) HasNavHoverColor() bool`

HasNavHoverColor returns a boolean if a field has been set.

### GetPrimaryButtonBgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetPrimaryButtonBgColor() string`

GetPrimaryButtonBgColor returns the PrimaryButtonBgColor field if non-nil, zero value otherwise.

### GetPrimaryButtonBgColorOk

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetPrimaryButtonBgColorOk() (*string, bool)`

GetPrimaryButtonBgColorOk returns a tuple with the PrimaryButtonBgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryButtonBgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) SetPrimaryButtonBgColor(v string)`

SetPrimaryButtonBgColor sets PrimaryButtonBgColor field to given value.

### HasPrimaryButtonBgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) HasPrimaryButtonBgColor() bool`

HasPrimaryButtonBgColor returns a boolean if a field has been set.

### GetPrimaryButtonFgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetPrimaryButtonFgColor() string`

GetPrimaryButtonFgColor returns the PrimaryButtonFgColor field if non-nil, zero value otherwise.

### GetPrimaryButtonFgColorOk

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetPrimaryButtonFgColorOk() (*string, bool)`

GetPrimaryButtonFgColorOk returns a tuple with the PrimaryButtonFgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryButtonFgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) SetPrimaryButtonFgColor(v string)`

SetPrimaryButtonFgColor sets PrimaryButtonFgColor field to given value.

### HasPrimaryButtonFgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) HasPrimaryButtonFgColor() bool`

HasPrimaryButtonFgColor returns a boolean if a field has been set.

### GetPrimaryButtonHoverBgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetPrimaryButtonHoverBgColor() string`

GetPrimaryButtonHoverBgColor returns the PrimaryButtonHoverBgColor field if non-nil, zero value otherwise.

### GetPrimaryButtonHoverBgColorOk

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetPrimaryButtonHoverBgColorOk() (*string, bool)`

GetPrimaryButtonHoverBgColorOk returns a tuple with the PrimaryButtonHoverBgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryButtonHoverBgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) SetPrimaryButtonHoverBgColor(v string)`

SetPrimaryButtonHoverBgColor sets PrimaryButtonHoverBgColor field to given value.

### HasPrimaryButtonHoverBgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) HasPrimaryButtonHoverBgColor() bool`

HasPrimaryButtonHoverBgColor returns a boolean if a field has been set.

### GetPrimaryButtonHoverFgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetPrimaryButtonHoverFgColor() string`

GetPrimaryButtonHoverFgColor returns the PrimaryButtonHoverFgColor field if non-nil, zero value otherwise.

### GetPrimaryButtonHoverFgColorOk

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetPrimaryButtonHoverFgColorOk() (*string, bool)`

GetPrimaryButtonHoverFgColorOk returns a tuple with the PrimaryButtonHoverFgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryButtonHoverFgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) SetPrimaryButtonHoverFgColor(v string)`

SetPrimaryButtonHoverFgColor sets PrimaryButtonHoverFgColor field to given value.

### HasPrimaryButtonHoverFgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) HasPrimaryButtonHoverFgColor() bool`

HasPrimaryButtonHoverFgColor returns a boolean if a field has been set.

### GetFooterBgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetFooterBgColor() string`

GetFooterBgColor returns the FooterBgColor field if non-nil, zero value otherwise.

### GetFooterBgColorOk

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetFooterBgColorOk() (*string, bool)`

GetFooterBgColorOk returns a tuple with the FooterBgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterBgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) SetFooterBgColor(v string)`

SetFooterBgColor sets FooterBgColor field to given value.

### HasFooterBgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) HasFooterBgColor() bool`

HasFooterBgColor returns a boolean if a field has been set.

### GetFooterFgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetFooterFgColor() string`

GetFooterFgColor returns the FooterFgColor field if non-nil, zero value otherwise.

### GetFooterFgColorOk

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetFooterFgColorOk() (*string, bool)`

GetFooterFgColorOk returns a tuple with the FooterFgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterFgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) SetFooterFgColor(v string)`

SetFooterFgColor sets FooterFgColor field to given value.

### HasFooterFgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) HasFooterFgColor() bool`

HasFooterFgColor returns a boolean if a field has been set.

### GetLoginBgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetLoginBgColor() string`

GetLoginBgColor returns the LoginBgColor field if non-nil, zero value otherwise.

### GetLoginBgColorOk

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetLoginBgColorOk() (*string, bool)`

GetLoginBgColorOk returns a tuple with the LoginBgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginBgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) SetLoginBgColor(v string)`

SetLoginBgColor sets LoginBgColor field to given value.

### HasLoginBgColor

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) HasLoginBgColor() bool`

HasLoginBgColor returns a boolean if a field has been set.

### GetCopyrightString

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetCopyrightString() string`

GetCopyrightString returns the CopyrightString field if non-nil, zero value otherwise.

### GetCopyrightStringOk

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetCopyrightStringOk() (*string, bool)`

GetCopyrightStringOk returns a tuple with the CopyrightString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrightString

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) SetCopyrightString(v string)`

SetCopyrightString sets CopyrightString field to given value.

### HasCopyrightString

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) HasCopyrightString() bool`

HasCopyrightString returns a boolean if a field has been set.

### GetOverrideCss

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetOverrideCss() string`

GetOverrideCss returns the OverrideCss field if non-nil, zero value otherwise.

### GetOverrideCssOk

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetOverrideCssOk() (*string, bool)`

GetOverrideCssOk returns a tuple with the OverrideCss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideCss

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) SetOverrideCss(v string)`

SetOverrideCss sets OverrideCss field to given value.

### HasOverrideCss

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) HasOverrideCss() bool`

HasOverrideCss returns a boolean if a field has been set.

### GetTermsOfUse

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetTermsOfUse() string`

GetTermsOfUse returns the TermsOfUse field if non-nil, zero value otherwise.

### GetTermsOfUseOk

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetTermsOfUseOk() (*string, bool)`

GetTermsOfUseOk returns a tuple with the TermsOfUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfUse

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) SetTermsOfUse(v string)`

SetTermsOfUse sets TermsOfUse field to given value.

### HasTermsOfUse

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) HasTermsOfUse() bool`

HasTermsOfUse returns a boolean if a field has been set.

### GetPrivacyPolicy

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetPrivacyPolicy() string`

GetPrivacyPolicy returns the PrivacyPolicy field if non-nil, zero value otherwise.

### GetPrivacyPolicyOk

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetPrivacyPolicyOk() (*string, bool)`

GetPrivacyPolicyOk returns a tuple with the PrivacyPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicy

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) SetPrivacyPolicy(v string)`

SetPrivacyPolicy sets PrivacyPolicy field to given value.

### HasPrivacyPolicy

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) HasPrivacyPolicy() bool`

HasPrivacyPolicy returns a boolean if a field has been set.

### GetSupportMenuLinks

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetSupportMenuLinks() []UpdateWhitelabelSettingsRequestWhitelabelSettingsSupportMenuLinksInner`

GetSupportMenuLinks returns the SupportMenuLinks field if non-nil, zero value otherwise.

### GetSupportMenuLinksOk

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) GetSupportMenuLinksOk() (*[]UpdateWhitelabelSettingsRequestWhitelabelSettingsSupportMenuLinksInner, bool)`

GetSupportMenuLinksOk returns a tuple with the SupportMenuLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportMenuLinks

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) SetSupportMenuLinks(v []UpdateWhitelabelSettingsRequestWhitelabelSettingsSupportMenuLinksInner)`

SetSupportMenuLinks sets SupportMenuLinks field to given value.

### HasSupportMenuLinks

`func (o *UpdateWhitelabelSettingsRequestWhitelabelSettings) HasSupportMenuLinks() bool`

HasSupportMenuLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


