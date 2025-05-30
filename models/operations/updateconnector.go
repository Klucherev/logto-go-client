// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/klucherev/logto-go-client/internal/utils"
	"github.com/klucherev/logto-go-client/models/components"
)

// UpdateConnectorConfigRequest - The connector config object that will be passed to the connector. The config object should be compatible with the connector factory.
type UpdateConnectorConfigRequest struct {
}

// UpdateConnectorNameRequest - Validator function
type UpdateConnectorNameRequest struct {
}

// UpdateConnectorMetadataRequest - Custom connector metadata, will be used to overwrite the default connector metadata.
type UpdateConnectorMetadataRequest struct {
	Target *string `json:"target,omitempty"`
	// Validator function
	Name     *UpdateConnectorNameRequest `json:"name,omitempty"`
	Logo     *string                     `json:"logo,omitempty"`
	LogoDark *string                     `json:"logoDark,omitempty"`
}

func (o *UpdateConnectorMetadataRequest) GetTarget() *string {
	if o == nil {
		return nil
	}
	return o.Target
}

func (o *UpdateConnectorMetadataRequest) GetName() *UpdateConnectorNameRequest {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateConnectorMetadataRequest) GetLogo() *string {
	if o == nil {
		return nil
	}
	return o.Logo
}

func (o *UpdateConnectorMetadataRequest) GetLogoDark() *string {
	if o == nil {
		return nil
	}
	return o.LogoDark
}

type UpdateConnectorRequestBody struct {
	// The connector config object that will be passed to the connector. The config object should be compatible with the connector factory.
	Config *UpdateConnectorConfigRequest `json:"config,omitempty"`
	// Custom connector metadata, will be used to overwrite the default connector metadata.
	Metadata *UpdateConnectorMetadataRequest `json:"metadata,omitempty"`
	// Whether to sync user profile from the identity provider to Logto at each sign-in. If `false`, the user profile will only be synced when the user is created.
	SyncProfile *bool `json:"syncProfile,omitempty"`
}

func (o *UpdateConnectorRequestBody) GetConfig() *UpdateConnectorConfigRequest {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *UpdateConnectorRequestBody) GetMetadata() *UpdateConnectorMetadataRequest {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *UpdateConnectorRequestBody) GetSyncProfile() *bool {
	if o == nil {
		return nil
	}
	return o.SyncProfile
}

type UpdateConnectorRequest struct {
	// The unique identifier of the connector.
	ID          string                     `pathParam:"style=simple,explode=false,name=id"`
	RequestBody UpdateConnectorRequestBody `request:"mediaType=application/json"`
}

func (o *UpdateConnectorRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateConnectorRequest) GetRequestBody() UpdateConnectorRequestBody {
	if o == nil {
		return UpdateConnectorRequestBody{}
	}
	return o.RequestBody
}

// UpdateConnectorConfigResponse - arbitrary
type UpdateConnectorConfigResponse struct {
}

// UpdateConnectorMetadataNameResponse - Validator function
type UpdateConnectorMetadataNameResponse struct {
}

type UpdateConnectorMetadataResponse struct {
	Target *string `json:"target,omitempty"`
	// Validator function
	Name     *UpdateConnectorMetadataNameResponse `json:"name,omitempty"`
	Logo     *string                              `json:"logo,omitempty"`
	LogoDark *string                              `json:"logoDark,omitempty"`
}

func (o *UpdateConnectorMetadataResponse) GetTarget() *string {
	if o == nil {
		return nil
	}
	return o.Target
}

func (o *UpdateConnectorMetadataResponse) GetName() *UpdateConnectorMetadataNameResponse {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateConnectorMetadataResponse) GetLogo() *string {
	if o == nil {
		return nil
	}
	return o.Logo
}

func (o *UpdateConnectorMetadataResponse) GetLogoDark() *string {
	if o == nil {
		return nil
	}
	return o.LogoDark
}

// UpdateConnectorNameResponse - Validator function
type UpdateConnectorNameResponse struct {
}

// UpdateConnectorDescription - Validator function
type UpdateConnectorDescription struct {
}

type UpdateConnectorFormItemType string

const (
	UpdateConnectorFormItemTypeText          UpdateConnectorFormItemType = "Text"
	UpdateConnectorFormItemTypeNumber        UpdateConnectorFormItemType = "Number"
	UpdateConnectorFormItemTypeMultilineText UpdateConnectorFormItemType = "MultilineText"
	UpdateConnectorFormItemTypeSwitch        UpdateConnectorFormItemType = "Switch"
	UpdateConnectorFormItemTypeJSON          UpdateConnectorFormItemType = "Json"
)

func (e UpdateConnectorFormItemType) ToPointer() *UpdateConnectorFormItemType {
	return &e
}
func (e *UpdateConnectorFormItemType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Text":
		fallthrough
	case "Number":
		fallthrough
	case "MultilineText":
		fallthrough
	case "Switch":
		fallthrough
	case "Json":
		*e = UpdateConnectorFormItemType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateConnectorFormItemType: %v", v)
	}
}

type UpdateConnectorShowCondition3 struct {
	TargetKey   string `json:"targetKey"`
	ExpectValue any    `json:"expectValue,omitempty"`
}

func (o *UpdateConnectorShowCondition3) GetTargetKey() string {
	if o == nil {
		return ""
	}
	return o.TargetKey
}

func (o *UpdateConnectorShowCondition3) GetExpectValue() any {
	if o == nil {
		return nil
	}
	return o.ExpectValue
}

type UpdateConnectorFormItem3 struct {
	Type           UpdateConnectorFormItemType     `json:"type"`
	Key            string                          `json:"key"`
	Label          string                          `json:"label"`
	Placeholder    *string                         `json:"placeholder,omitempty"`
	Required       *bool                           `json:"required,omitempty"`
	DefaultValue   any                             `json:"defaultValue,omitempty"`
	ShowConditions []UpdateConnectorShowCondition3 `json:"showConditions,omitempty"`
	Description    *string                         `json:"description,omitempty"`
	Tooltip        *string                         `json:"tooltip,omitempty"`
	IsConfidential *bool                           `json:"isConfidential,omitempty"`
}

func (o *UpdateConnectorFormItem3) GetType() UpdateConnectorFormItemType {
	if o == nil {
		return UpdateConnectorFormItemType("")
	}
	return o.Type
}

func (o *UpdateConnectorFormItem3) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *UpdateConnectorFormItem3) GetLabel() string {
	if o == nil {
		return ""
	}
	return o.Label
}

func (o *UpdateConnectorFormItem3) GetPlaceholder() *string {
	if o == nil {
		return nil
	}
	return o.Placeholder
}

func (o *UpdateConnectorFormItem3) GetRequired() *bool {
	if o == nil {
		return nil
	}
	return o.Required
}

func (o *UpdateConnectorFormItem3) GetDefaultValue() any {
	if o == nil {
		return nil
	}
	return o.DefaultValue
}

func (o *UpdateConnectorFormItem3) GetShowConditions() []UpdateConnectorShowCondition3 {
	if o == nil {
		return nil
	}
	return o.ShowConditions
}

func (o *UpdateConnectorFormItem3) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdateConnectorFormItem3) GetTooltip() *string {
	if o == nil {
		return nil
	}
	return o.Tooltip
}

func (o *UpdateConnectorFormItem3) GetIsConfidential() *bool {
	if o == nil {
		return nil
	}
	return o.IsConfidential
}

type UpdateConnectorSelectItem2 struct {
	Value string `json:"value"`
}

func (o *UpdateConnectorSelectItem2) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type UpdateConnectorShowCondition2 struct {
	TargetKey   string `json:"targetKey"`
	ExpectValue any    `json:"expectValue,omitempty"`
}

func (o *UpdateConnectorShowCondition2) GetTargetKey() string {
	if o == nil {
		return ""
	}
	return o.TargetKey
}

func (o *UpdateConnectorShowCondition2) GetExpectValue() any {
	if o == nil {
		return nil
	}
	return o.ExpectValue
}

type UpdateConnectorFormItem2 struct {
	Type           string                          `json:"type"`
	SelectItems    []UpdateConnectorSelectItem2    `json:"selectItems"`
	Key            string                          `json:"key"`
	Label          string                          `json:"label"`
	Placeholder    *string                         `json:"placeholder,omitempty"`
	Required       *bool                           `json:"required,omitempty"`
	DefaultValue   any                             `json:"defaultValue,omitempty"`
	ShowConditions []UpdateConnectorShowCondition2 `json:"showConditions,omitempty"`
	Description    *string                         `json:"description,omitempty"`
	Tooltip        *string                         `json:"tooltip,omitempty"`
	IsConfidential *bool                           `json:"isConfidential,omitempty"`
}

func (o *UpdateConnectorFormItem2) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *UpdateConnectorFormItem2) GetSelectItems() []UpdateConnectorSelectItem2 {
	if o == nil {
		return []UpdateConnectorSelectItem2{}
	}
	return o.SelectItems
}

func (o *UpdateConnectorFormItem2) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *UpdateConnectorFormItem2) GetLabel() string {
	if o == nil {
		return ""
	}
	return o.Label
}

func (o *UpdateConnectorFormItem2) GetPlaceholder() *string {
	if o == nil {
		return nil
	}
	return o.Placeholder
}

func (o *UpdateConnectorFormItem2) GetRequired() *bool {
	if o == nil {
		return nil
	}
	return o.Required
}

func (o *UpdateConnectorFormItem2) GetDefaultValue() any {
	if o == nil {
		return nil
	}
	return o.DefaultValue
}

func (o *UpdateConnectorFormItem2) GetShowConditions() []UpdateConnectorShowCondition2 {
	if o == nil {
		return nil
	}
	return o.ShowConditions
}

func (o *UpdateConnectorFormItem2) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdateConnectorFormItem2) GetTooltip() *string {
	if o == nil {
		return nil
	}
	return o.Tooltip
}

func (o *UpdateConnectorFormItem2) GetIsConfidential() *bool {
	if o == nil {
		return nil
	}
	return o.IsConfidential
}

type UpdateConnectorSelectItem1 struct {
	Value string `json:"value"`
	Title string `json:"title"`
}

func (o *UpdateConnectorSelectItem1) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

func (o *UpdateConnectorSelectItem1) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

type UpdateConnectorShowCondition1 struct {
	TargetKey   string `json:"targetKey"`
	ExpectValue any    `json:"expectValue,omitempty"`
}

func (o *UpdateConnectorShowCondition1) GetTargetKey() string {
	if o == nil {
		return ""
	}
	return o.TargetKey
}

func (o *UpdateConnectorShowCondition1) GetExpectValue() any {
	if o == nil {
		return nil
	}
	return o.ExpectValue
}

type UpdateConnectorFormItem1 struct {
	Type           string                          `json:"type"`
	SelectItems    []UpdateConnectorSelectItem1    `json:"selectItems"`
	Key            string                          `json:"key"`
	Label          string                          `json:"label"`
	Placeholder    *string                         `json:"placeholder,omitempty"`
	Required       *bool                           `json:"required,omitempty"`
	DefaultValue   any                             `json:"defaultValue,omitempty"`
	ShowConditions []UpdateConnectorShowCondition1 `json:"showConditions,omitempty"`
	Description    *string                         `json:"description,omitempty"`
	Tooltip        *string                         `json:"tooltip,omitempty"`
	IsConfidential *bool                           `json:"isConfidential,omitempty"`
}

func (o *UpdateConnectorFormItem1) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *UpdateConnectorFormItem1) GetSelectItems() []UpdateConnectorSelectItem1 {
	if o == nil {
		return []UpdateConnectorSelectItem1{}
	}
	return o.SelectItems
}

func (o *UpdateConnectorFormItem1) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *UpdateConnectorFormItem1) GetLabel() string {
	if o == nil {
		return ""
	}
	return o.Label
}

func (o *UpdateConnectorFormItem1) GetPlaceholder() *string {
	if o == nil {
		return nil
	}
	return o.Placeholder
}

func (o *UpdateConnectorFormItem1) GetRequired() *bool {
	if o == nil {
		return nil
	}
	return o.Required
}

func (o *UpdateConnectorFormItem1) GetDefaultValue() any {
	if o == nil {
		return nil
	}
	return o.DefaultValue
}

func (o *UpdateConnectorFormItem1) GetShowConditions() []UpdateConnectorShowCondition1 {
	if o == nil {
		return nil
	}
	return o.ShowConditions
}

func (o *UpdateConnectorFormItem1) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdateConnectorFormItem1) GetTooltip() *string {
	if o == nil {
		return nil
	}
	return o.Tooltip
}

func (o *UpdateConnectorFormItem1) GetIsConfidential() *bool {
	if o == nil {
		return nil
	}
	return o.IsConfidential
}

type UpdateConnectorFormItemUnionType string

const (
	UpdateConnectorFormItemUnionTypeUpdateConnectorFormItem1 UpdateConnectorFormItemUnionType = "UpdateConnector_formItem_1"
	UpdateConnectorFormItemUnionTypeUpdateConnectorFormItem2 UpdateConnectorFormItemUnionType = "UpdateConnector_formItem_2"
	UpdateConnectorFormItemUnionTypeUpdateConnectorFormItem3 UpdateConnectorFormItemUnionType = "UpdateConnector_formItem_3"
)

type UpdateConnectorFormItemUnion struct {
	UpdateConnectorFormItem1 *UpdateConnectorFormItem1 `queryParam:"inline"`
	UpdateConnectorFormItem2 *UpdateConnectorFormItem2 `queryParam:"inline"`
	UpdateConnectorFormItem3 *UpdateConnectorFormItem3 `queryParam:"inline"`

	Type UpdateConnectorFormItemUnionType
}

func CreateUpdateConnectorFormItemUnionUpdateConnectorFormItem1(updateConnectorFormItem1 UpdateConnectorFormItem1) UpdateConnectorFormItemUnion {
	typ := UpdateConnectorFormItemUnionTypeUpdateConnectorFormItem1

	return UpdateConnectorFormItemUnion{
		UpdateConnectorFormItem1: &updateConnectorFormItem1,
		Type:                     typ,
	}
}

func CreateUpdateConnectorFormItemUnionUpdateConnectorFormItem2(updateConnectorFormItem2 UpdateConnectorFormItem2) UpdateConnectorFormItemUnion {
	typ := UpdateConnectorFormItemUnionTypeUpdateConnectorFormItem2

	return UpdateConnectorFormItemUnion{
		UpdateConnectorFormItem2: &updateConnectorFormItem2,
		Type:                     typ,
	}
}

func CreateUpdateConnectorFormItemUnionUpdateConnectorFormItem3(updateConnectorFormItem3 UpdateConnectorFormItem3) UpdateConnectorFormItemUnion {
	typ := UpdateConnectorFormItemUnionTypeUpdateConnectorFormItem3

	return UpdateConnectorFormItemUnion{
		UpdateConnectorFormItem3: &updateConnectorFormItem3,
		Type:                     typ,
	}
}

func (u *UpdateConnectorFormItemUnion) UnmarshalJSON(data []byte) error {

	var updateConnectorFormItem3 UpdateConnectorFormItem3 = UpdateConnectorFormItem3{}
	if err := utils.UnmarshalJSON(data, &updateConnectorFormItem3, "", true, true); err == nil {
		u.UpdateConnectorFormItem3 = &updateConnectorFormItem3
		u.Type = UpdateConnectorFormItemUnionTypeUpdateConnectorFormItem3
		return nil
	}

	var updateConnectorFormItem1 UpdateConnectorFormItem1 = UpdateConnectorFormItem1{}
	if err := utils.UnmarshalJSON(data, &updateConnectorFormItem1, "", true, true); err == nil {
		u.UpdateConnectorFormItem1 = &updateConnectorFormItem1
		u.Type = UpdateConnectorFormItemUnionTypeUpdateConnectorFormItem1
		return nil
	}

	var updateConnectorFormItem2 UpdateConnectorFormItem2 = UpdateConnectorFormItem2{}
	if err := utils.UnmarshalJSON(data, &updateConnectorFormItem2, "", true, true); err == nil {
		u.UpdateConnectorFormItem2 = &updateConnectorFormItem2
		u.Type = UpdateConnectorFormItemUnionTypeUpdateConnectorFormItem2
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for UpdateConnectorFormItemUnion", string(data))
}

func (u UpdateConnectorFormItemUnion) MarshalJSON() ([]byte, error) {
	if u.UpdateConnectorFormItem1 != nil {
		return utils.MarshalJSON(u.UpdateConnectorFormItem1, "", true)
	}

	if u.UpdateConnectorFormItem2 != nil {
		return utils.MarshalJSON(u.UpdateConnectorFormItem2, "", true)
	}

	if u.UpdateConnectorFormItem3 != nil {
		return utils.MarshalJSON(u.UpdateConnectorFormItem3, "", true)
	}

	return nil, errors.New("could not marshal union type UpdateConnectorFormItemUnion: all fields are null")
}

type UpdateConnectorPlatform string

const (
	UpdateConnectorPlatformNative    UpdateConnectorPlatform = "Native"
	UpdateConnectorPlatformUniversal UpdateConnectorPlatform = "Universal"
	UpdateConnectorPlatformWeb       UpdateConnectorPlatform = "Web"
)

func (e UpdateConnectorPlatform) ToPointer() *UpdateConnectorPlatform {
	return &e
}
func (e *UpdateConnectorPlatform) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Native":
		fallthrough
	case "Universal":
		fallthrough
	case "Web":
		*e = UpdateConnectorPlatform(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateConnectorPlatform: %v", v)
	}
}

type UpdateConnectorType string

const (
	UpdateConnectorTypeEmail  UpdateConnectorType = "Email"
	UpdateConnectorTypeSms    UpdateConnectorType = "Sms"
	UpdateConnectorTypeSocial UpdateConnectorType = "Social"
)

func (e UpdateConnectorType) ToPointer() *UpdateConnectorType {
	return &e
}
func (e *UpdateConnectorType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Email":
		fallthrough
	case "Sms":
		fallthrough
	case "Social":
		*e = UpdateConnectorType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateConnectorType: %v", v)
	}
}

// UpdateConnectorResponseBody - The updated connector.
type UpdateConnectorResponseBody struct {
	ID          string `json:"id"`
	SyncProfile bool   `json:"syncProfile"`
	// arbitrary
	Config      UpdateConnectorConfigResponse   `json:"config"`
	Metadata    UpdateConnectorMetadataResponse `json:"metadata"`
	ConnectorID string                          `json:"connectorId"`
	Target      string                          `json:"target"`
	// Validator function
	Name UpdateConnectorNameResponse `json:"name"`
	// Validator function
	Description    UpdateConnectorDescription     `json:"description"`
	Logo           string                         `json:"logo"`
	LogoDark       *string                        `json:"logoDark"`
	Readme         string                         `json:"readme"`
	ConfigTemplate *string                        `json:"configTemplate,omitempty"`
	FormItems      []UpdateConnectorFormItemUnion `json:"formItems,omitempty"`
	CustomData     map[string]any                 `json:"customData,omitempty"`
	FromEmail      *string                        `json:"fromEmail,omitempty"`
	Platform       *UpdateConnectorPlatform       `json:"platform"`
	IsStandard     *bool                          `json:"isStandard,omitempty"`
	Type           UpdateConnectorType            `json:"type"`
	IsDemo         *bool                          `json:"isDemo,omitempty"`
	ExtraInfo      map[string]any                 `json:"extraInfo,omitempty"`
	Usage          *float64                       `json:"usage,omitempty"`
}

func (o *UpdateConnectorResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateConnectorResponseBody) GetSyncProfile() bool {
	if o == nil {
		return false
	}
	return o.SyncProfile
}

func (o *UpdateConnectorResponseBody) GetConfig() UpdateConnectorConfigResponse {
	if o == nil {
		return UpdateConnectorConfigResponse{}
	}
	return o.Config
}

func (o *UpdateConnectorResponseBody) GetMetadata() UpdateConnectorMetadataResponse {
	if o == nil {
		return UpdateConnectorMetadataResponse{}
	}
	return o.Metadata
}

func (o *UpdateConnectorResponseBody) GetConnectorID() string {
	if o == nil {
		return ""
	}
	return o.ConnectorID
}

func (o *UpdateConnectorResponseBody) GetTarget() string {
	if o == nil {
		return ""
	}
	return o.Target
}

func (o *UpdateConnectorResponseBody) GetName() UpdateConnectorNameResponse {
	if o == nil {
		return UpdateConnectorNameResponse{}
	}
	return o.Name
}

func (o *UpdateConnectorResponseBody) GetDescription() UpdateConnectorDescription {
	if o == nil {
		return UpdateConnectorDescription{}
	}
	return o.Description
}

func (o *UpdateConnectorResponseBody) GetLogo() string {
	if o == nil {
		return ""
	}
	return o.Logo
}

func (o *UpdateConnectorResponseBody) GetLogoDark() *string {
	if o == nil {
		return nil
	}
	return o.LogoDark
}

func (o *UpdateConnectorResponseBody) GetReadme() string {
	if o == nil {
		return ""
	}
	return o.Readme
}

func (o *UpdateConnectorResponseBody) GetConfigTemplate() *string {
	if o == nil {
		return nil
	}
	return o.ConfigTemplate
}

func (o *UpdateConnectorResponseBody) GetFormItems() []UpdateConnectorFormItemUnion {
	if o == nil {
		return nil
	}
	return o.FormItems
}

func (o *UpdateConnectorResponseBody) GetCustomData() map[string]any {
	if o == nil {
		return nil
	}
	return o.CustomData
}

func (o *UpdateConnectorResponseBody) GetFromEmail() *string {
	if o == nil {
		return nil
	}
	return o.FromEmail
}

func (o *UpdateConnectorResponseBody) GetPlatform() *UpdateConnectorPlatform {
	if o == nil {
		return nil
	}
	return o.Platform
}

func (o *UpdateConnectorResponseBody) GetIsStandard() *bool {
	if o == nil {
		return nil
	}
	return o.IsStandard
}

func (o *UpdateConnectorResponseBody) GetType() UpdateConnectorType {
	if o == nil {
		return UpdateConnectorType("")
	}
	return o.Type
}

func (o *UpdateConnectorResponseBody) GetIsDemo() *bool {
	if o == nil {
		return nil
	}
	return o.IsDemo
}

func (o *UpdateConnectorResponseBody) GetExtraInfo() map[string]any {
	if o == nil {
		return nil
	}
	return o.ExtraInfo
}

func (o *UpdateConnectorResponseBody) GetUsage() *float64 {
	if o == nil {
		return nil
	}
	return o.Usage
}

type UpdateConnectorResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The updated connector.
	Object *UpdateConnectorResponseBody
}

func (o *UpdateConnectorResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateConnectorResponse) GetObject() *UpdateConnectorResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
