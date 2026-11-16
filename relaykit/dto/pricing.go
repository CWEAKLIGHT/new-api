package dto

import "github.com/QuantumNous/new-api/relaykit/types"

// 这里不好动就不动了，本来想独立出来的（
//
// 能力字段（InputModalities 起）是本地补丁：上游 wb-proxy 的 /v1/models 会带
// capabilities / input_modalities 这类元数据，客户端（比如 DSH）靠它判断“这个模型
// 能不能吃图片输入”。new-api 原来只回 OpenAI 的四个基础字段，客户端拿不到能力标记，
// 就会把带图请求直接拦掉（"Model X does not support image input"）。
// 这些字段一律 omitempty，不影响标准 OpenAI 客户端的解析。
type OpenAIModels struct {
	Id                     string               `json:"id"`
	Object                 string               `json:"object"`
	Created                int                  `json:"created"`
	OwnedBy                string               `json:"owned_by"`
	SupportedEndpointTypes []types.EndpointType `json:"supported_endpoint_types"`

	// 本地补丁：模型能力
	Name            string   `json:"name,omitempty"`
	Description     string   `json:"description,omitempty"`
	InputModalities []string `json:"input_modalities,omitempty"`
	OutputModalities []string `json:"output_modalities,omitempty"`
	Modalities      *ModelModalities `json:"modalities,omitempty"`
	Architecture    *ModelArchitecture `json:"architecture,omitempty"`
	Capabilities    *ModelCapabilities `json:"capabilities,omitempty"`
	SupportsVision  *bool    `json:"supports_vision,omitempty"`
	SupportsImages  *bool    `json:"supports_images,omitempty"`
	Vision          *bool    `json:"vision,omitempty"`
	Multimodal      *bool    `json:"multimodal,omitempty"`
	SupportsToolCalls *bool  `json:"supports_tool_calls,omitempty"`
	SupportsReasoning *bool  `json:"supports_reasoning,omitempty"`
	ContextLength   *int     `json:"context_length,omitempty"`
	MaxInputTokens  *int     `json:"max_input_tokens,omitempty"`
	MaxOutputTokens *int     `json:"max_output_tokens,omitempty"`
}

// ModelCapabilities 对应 OpenAI-compatible 生态里常见的 capabilities 结构。
type ModelCapabilities struct {
	Vision     bool `json:"vision"`
	Audio      bool `json:"audio,omitempty"`
	ToolCalls  bool `json:"tool_calls"`
	Reasoning  bool `json:"reasoning"`
}

type ModelModalities struct {
	Input  []string `json:"input"`
	Output []string `json:"output"`
}

type ModelArchitecture struct {
	InputModalities  []string `json:"input_modalities"`
	OutputModalities []string `json:"output_modalities"`
	Modality         string   `json:"modality"`
}

type AnthropicModel struct {
	ID          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	DisplayName string `json:"display_name"`
	Type        string `json:"type"`
}

type GeminiModel struct {
	Name                       any   `json:"name"`
	BaseModelId                any   `json:"baseModelId"`
	Version                    any   `json:"version"`
	DisplayName                any   `json:"displayName"`
	Description                any   `json:"description"`
	InputTokenLimit            any   `json:"inputTokenLimit"`
	OutputTokenLimit           any   `json:"outputTokenLimit"`
	SupportedGenerationMethods []any `json:"supportedGenerationMethods"`
	Thinking                   any   `json:"thinking"`
	Temperature                any   `json:"temperature"`
	MaxTemperature             any   `json:"maxTemperature"`
	TopP                       any   `json:"topP"`
	TopK                       any   `json:"topK"`
}
