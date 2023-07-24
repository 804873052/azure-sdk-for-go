//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azopenai

// azureCoreFoundationsError - The error object.
type azureCoreFoundationsError struct {
	// REQUIRED; One of a server-defined set of error codes.
	Code *string

	// REQUIRED; A human-readable representation of the error.
	Message *string

	// An array of details about specific errors that led to this reported error.
	Details []azureCoreFoundationsError

	// An object containing more specific information than the current object about the error.
	Innererror *azureCoreFoundationsErrorInnererror

	// The target of the error.
	Target *string
}

// azureCoreFoundationsErrorInnererror - An object containing more specific information than the current object about the
// error.
type azureCoreFoundationsErrorInnererror struct {
	// One of a server-defined set of error codes.
	Code *string

	// Inner error.
	Innererror *azureCoreFoundationsInnerErrorInnererror
}

// azureCoreFoundationsErrorResponse - A response containing error details.
type azureCoreFoundationsErrorResponse struct {
	// REQUIRED; The error object.
	Error *azureCoreFoundationsErrorResponseError
}

// azureCoreFoundationsErrorResponseError - The error object.
type azureCoreFoundationsErrorResponseError struct {
	// REQUIRED; One of a server-defined set of error codes.
	Code *string

	// REQUIRED; A human-readable representation of the error.
	Message *string

	// An array of details about specific errors that led to this reported error.
	Details []azureCoreFoundationsError

	// An object containing more specific information than the current object about the error.
	Innererror *azureCoreFoundationsErrorInnererror

	// The target of the error.
	Target *string
}

// azureCoreFoundationsInnerError - An object containing more specific information about the error. As per Microsoft One API
// guidelines -
// https://github.com/Microsoft/api-guidelines/blob/vNext/Guidelines.md#7102-error-condition-responses.
type azureCoreFoundationsInnerError struct {
	// One of a server-defined set of error codes.
	Code *string

	// Inner error.
	Innererror *azureCoreFoundationsInnerErrorInnererror
}

// azureCoreFoundationsInnerErrorInnererror - Inner error.
type azureCoreFoundationsInnerErrorInnererror struct {
	// One of a server-defined set of error codes.
	Code *string

	// Inner error.
	Innererror *azureCoreFoundationsInnerErrorInnererror
}

// batchImageGenerationOperationResponse - A polling status update or final response payload for an image operation.
type batchImageGenerationOperationResponse struct {
	// REQUIRED; A timestamp when this job or item was created (in unix epochs).
	Created *int64

	// REQUIRED; The ID of the operation.
	ID *string

	// REQUIRED; The status of the operation
	Status *azureOpenAIOperationState

	// The error if the operation failed.
	Error *azureCoreFoundationsError

	// A timestamp when this operation and its associated images expire and will be deleted (in unix epochs).
	Expires *int64

	// The result of the operation if the operation succeeded.
	Result *ImageGenerations
}

// ChatChoice - The representation of a single prompt completion as part of an overall chat completions request. Generally,
// n choices are generated per provided prompt with a default value of 1. Token limits and
// other settings may limit the number of choices generated.
type ChatChoice struct {
	// REQUIRED; The reason that this chat completions choice completed its generated.
	FinishReason *CompletionsFinishReason

	// REQUIRED; The ordered index associated with this chat completions choice.
	Index *int32

	// Information about the content filtering category (hate, sexual, violence, selfharm), if it has been detected, as well as
	// the severity level (verylow, low, medium, high-scale that determines the
	// intensity and risk level of harmful content) and if it has been filtered or not.
	ContentFilterResults *ChatChoiceContentFilterResults

	// The delta message content for a streaming response.
	Delta *ChatChoiceDelta

	// The chat message for a given chat completions prompt.
	Message *ChatChoiceMessage
}

// ChatChoiceContentFilterResults - Information about the content filtering category (hate, sexual, violence, selfharm), if
// it has been detected, as well as the severity level (verylow, low, medium, high-scale that determines the
// intensity and risk level of harmful content) and if it has been filtered or not.
type ChatChoiceContentFilterResults struct {
	// REQUIRED; Describes language attacks or uses that include pejorative or discriminatory language with reference to a person
	// or identity group on the basis of certain differentiating attributes of these groups
	// including but not limited to race, ethnicity, nationality, gender identity and expression, sexual orientation, religion,
	// immigration status, ability status, personal appearance, and body size.
	Hate *ContentFilterResultsHate

	// REQUIRED; Describes language related to physical actions intended to purposely hurt, injure, or damage one’s body, or kill
	// oneself.
	SelfHarm *ContentFilterResultsSelfHarm

	// REQUIRED; Describes language related to anatomical organs and genitals, romantic relationships, acts portrayed in erotic
	// or affectionate terms, physical sexual acts, including those portrayed as an assault or a
	// forced sexual violent act against one’s will, prostitution, pornography, and abuse.
	Sexual *ContentFilterResultsSexual

	// REQUIRED; Describes language related to physical actions intended to hurt, injure, damage, or kill someone or something;
	// describes weapons, etc.
	Violence *ContentFilterResultsViolence
}

// ChatChoiceDelta - The delta message content for a streaming response.
type ChatChoiceDelta struct {
	// REQUIRED; The role associated with this message payload.
	Role *ChatRole

	// The text associated with this message payload.
	Content *string

	// The name and arguments of a function that should be called, as generated by the model.
	FunctionCall *ChatMessageFunctionCall

	// The name of the author of this message. name is required if role is function, and it should be the name of the function
	// whose response is in the content. May contain a-z, A-Z, 0-9, and underscores,
	// with a maximum length of 64 characters.
	Name *string
}

// ChatChoiceMessage - The chat message for a given chat completions prompt.
type ChatChoiceMessage struct {
	// REQUIRED; The role associated with this message payload.
	Role *ChatRole

	// The text associated with this message payload.
	Content *string

	// The name and arguments of a function that should be called, as generated by the model.
	FunctionCall *ChatMessageFunctionCall

	// The name of the author of this message. name is required if role is function, and it should be the name of the function
	// whose response is in the content. May contain a-z, A-Z, 0-9, and underscores,
	// with a maximum length of 64 characters.
	Name *string
}

// ChatCompletions - Representation of the response data from a chat completions request. Completions support a wide variety
// of tasks and generate text that continues from or "completes" provided prompt data.
type ChatCompletions struct {
	// REQUIRED; The collection of completions choices associated with this completions response. Generally, n choices are generated
	// per provided prompt with a default value of 1. Token limits and other settings may
	// limit the number of choices generated.
	Choices []ChatChoice

	// REQUIRED; The first timestamp associated with generation activity for this completions response, represented as seconds
	// since the beginning of the Unix epoch of 00:00 on 1 Jan 1970.
	Created *int32

	// REQUIRED; A unique identifier associated with this chat completions response.
	ID *string

	// REQUIRED; Usage information for tokens processed and generated as part of this completions operation.
	Usage *CompletionsUsage

	// Content filtering results for zero or more prompts in the request. In a streaming request, results for different prompts
	// may arrive at different times or in different orders.
	PromptAnnotations []PromptFilterResult
}

// ChatCompletionsOptions - The configuration information for a chat completions request. Completions support a wide variety
// of tasks and generate text that continues from or "completes" provided prompt data.
type ChatCompletionsOptions struct {
	// REQUIRED; The collection of context messages associated with this chat completions request. Typical usage begins with a
	// chat message for the System role that provides instructions for the behavior of the
	// assistant, followed by alternating messages between the User and Assistant roles.
	Messages []ChatMessage

	// A value that influences the probability of generated tokens appearing based on their cumulative frequency in generated
	// text. Positive values will make tokens less likely to appear as their frequency
	// increases and decrease the likelihood of the model repeating the same statements verbatim.
	FrequencyPenalty *float32

	// Controls how the model responds to function calls. "none" means the model does not call a function, and responds to the
	// end-user. "auto" means the model can pick between an end-user or calling a
	// function. Specifying a particular function via {"name": "my_function"} forces the model to call that function. "none" is
	// the default when no functions are present. "auto" is the default if functions
	// are present.
	FunctionCall *ChatCompletionsOptionsFunctionCall

	// A list of functions the model may generate JSON inputs for.
	Functions []FunctionDefinition

	// A map between GPT token IDs and bias scores that influences the probability of specific tokens appearing in a completions
	// response. Token IDs are computed via external tokenizer tools, while bias
	// scores reside in the range of -100 to 100 with minimum and maximum values corresponding to a full ban or exclusive selection
	// of a token, respectively. The exact behavior of a given bias score varies
	// by model.
	LogitBias map[string]*int32

	// The maximum number of tokens to generate.
	MaxTokens *int32

	// REQUIRED: DeploymentID specifies the name of the deployment (for Azure OpenAI) or model (for OpenAI) to use for this request.
	DeploymentID string

	// The number of chat completions choices that should be generated for a chat completions response. Because this setting can
	// generate many completions, it may quickly consume your token quota. Use
	// carefully and ensure reasonable settings for max_tokens and stop.
	N *int32

	// A value that influences the probability of generated tokens appearing based on their existing presence in generated text.
	// Positive values will make tokens less likely to appear when they already exist
	// and increase the model's likelihood to output new topics.
	PresencePenalty *float32

	// A collection of textual sequences that will end completions generation.
	Stop []string

	// The sampling temperature to use that controls the apparent creativity of generated completions. Higher values will make
	// output more random while lower values will make results more focused and
	// deterministic. It is not recommended to modify temperature and top_p for the same completions request as the interaction
	// of these two settings is difficult to predict.
	Temperature *float32

	// An alternative to sampling with temperature called nucleus sampling. This value causes the model to consider the results
	// of tokens with the provided probability mass. As an example, a value of 0.15
	// will cause only the tokens comprising the top 15% of probability mass to be considered. It is not recommended to modify
	// temperature and top_p for the same completions request as the interaction of
	// these two settings is difficult to predict.
	TopP *float32

	// An identifier for the caller or end user of the operation. This may be used for tracking or rate-limiting purposes.
	User *string
}

// ChatMessage - A single, role-attributed message within a chat completion interaction.
type ChatMessage struct {
	// REQUIRED; The role associated with this message payload.
	Role *ChatRole

	// The text associated with this message payload.
	Content *string

	// The name and arguments of a function that should be called, as generated by the model.
	FunctionCall *ChatMessageFunctionCall

	// The name of the author of this message. name is required if role is function, and it should be the name of the function
	// whose response is in the content. May contain a-z, A-Z, 0-9, and underscores,
	// with a maximum length of 64 characters.
	Name *string
}

// ChatMessageFunctionCall - The name and arguments of a function that should be called, as generated by the model.
type ChatMessageFunctionCall struct {
	// REQUIRED; The arguments to call the function with, as generated by the model in JSON format. Note that the model does not
	// always generate valid JSON, and may hallucinate parameters not defined by your function
	// schema. Validate the arguments in your code before calling your function.
	Arguments *string

	// REQUIRED; The name of the function to call.
	Name *string
}

// Choice - The representation of a single prompt completion as part of an overall completions request. Generally, n choices
// are generated per provided prompt with a default value of 1. Token limits and other
// settings may limit the number of choices generated.
type Choice struct {
	// REQUIRED; Reason for finishing
	FinishReason *CompletionsFinishReason

	// REQUIRED; The ordered index associated with this completions choice.
	Index *int32

	// REQUIRED; The log probabilities model for tokens associated with this completions choice.
	LogProbs *ChoiceLogProbs

	// REQUIRED; The generated text for a given completions prompt.
	Text *string

	// Information about the content filtering category (hate, sexual, violence, selfharm), if it has been detected, as well as
	// the severity level (verylow, low, medium, high-scale that determines the
	// intensity and risk level of harmful content) and if it has been filtered or not.
	ContentFilterResults *ChoiceContentFilterResults
}

// ChoiceContentFilterResults - Information about the content filtering category (hate, sexual, violence, selfharm), if it
// has been detected, as well as the severity level (verylow, low, medium, high-scale that determines the
// intensity and risk level of harmful content) and if it has been filtered or not.
type ChoiceContentFilterResults struct {
	// REQUIRED; Describes language attacks or uses that include pejorative or discriminatory language with reference to a person
	// or identity group on the basis of certain differentiating attributes of these groups
	// including but not limited to race, ethnicity, nationality, gender identity and expression, sexual orientation, religion,
	// immigration status, ability status, personal appearance, and body size.
	Hate *ContentFilterResultsHate

	// REQUIRED; Describes language related to physical actions intended to purposely hurt, injure, or damage one’s body, or kill
	// oneself.
	SelfHarm *ContentFilterResultsSelfHarm

	// REQUIRED; Describes language related to anatomical organs and genitals, romantic relationships, acts portrayed in erotic
	// or affectionate terms, physical sexual acts, including those portrayed as an assault or a
	// forced sexual violent act against one’s will, prostitution, pornography, and abuse.
	Sexual *ContentFilterResultsSexual

	// REQUIRED; Describes language related to physical actions intended to hurt, injure, damage, or kill someone or something;
	// describes weapons, etc.
	Violence *ContentFilterResultsViolence
}

// ChoiceLogProbs - The log probabilities model for tokens associated with this completions choice.
type ChoiceLogProbs struct {
	// REQUIRED; The text offsets associated with tokens in this completions data.
	TextOffset []int32

	// REQUIRED; A collection of log probability values for the tokens in this completions data.
	TokenLogProbs []float32

	// REQUIRED; The textual forms of tokens evaluated in this probability model.
	Tokens []string

	// REQUIRED; A mapping of tokens to maximum log probability values in this completions data.
	TopLogProbs []any
}

// Completions - Representation of the response data from a completions request. Completions support a wide variety of tasks
// and generate text that continues from or "completes" provided prompt data.
type Completions struct {
	// REQUIRED; The collection of completions choices associated with this completions response. Generally, n choices are generated
	// per provided prompt with a default value of 1. Token limits and other settings may
	// limit the number of choices generated.
	Choices []Choice

	// REQUIRED; The first timestamp associated with generation activity for this completions response, represented as seconds
	// since the beginning of the Unix epoch of 00:00 on 1 Jan 1970.
	Created *int32

	// REQUIRED; A unique identifier associated with this completions response.
	ID *string

	// REQUIRED; Usage information for tokens processed and generated as part of this completions operation.
	Usage *CompletionsUsage

	// Content filtering results for zero or more prompts in the request. In a streaming request, results for different prompts
	// may arrive at different times or in different orders.
	PromptAnnotations []PromptFilterResult
}

// CompletionsLogProbabilityModel - Representation of a log probabilities model for a completions generation.
type CompletionsLogProbabilityModel struct {
	// REQUIRED; The text offsets associated with tokens in this completions data.
	TextOffset []int32

	// REQUIRED; A collection of log probability values for the tokens in this completions data.
	TokenLogProbs []float32

	// REQUIRED; The textual forms of tokens evaluated in this probability model.
	Tokens []string

	// REQUIRED; A mapping of tokens to maximum log probability values in this completions data.
	TopLogProbs []any
}

// CompletionsOptions - The configuration information for a completions request. Completions support a wide variety of tasks
// and generate text that continues from or "completes" provided prompt data.
type CompletionsOptions struct {
	// REQUIRED; The prompts to generate completions from.
	Prompt []string

	// A value that controls how many completions will be internally generated prior to response formulation. When used together
	// with n, bestof controls the number of candidate completions and must be
	// greater than n. Because this setting can generate many completions, it may quickly consume your token quota. Use carefully
	// and ensure reasonable settings for maxtokens and stop.
	BestOf *int32

	// A value specifying whether completions responses should include input prompts as prefixes to their generated output.
	Echo *bool

	// A value that influences the probability of generated tokens appearing based on their cumulative frequency in generated
	// text. Positive values will make tokens less likely to appear as their frequency
	// increases and decrease the likelihood of the model repeating the same statements verbatim.
	FrequencyPenalty *float32

	// A map between GPT token IDs and bias scores that influences the probability of specific tokens appearing in a completions
	// response. Token IDs are computed via external tokenizer tools, while bias
	// scores reside in the range of -100 to 100 with minimum and maximum values corresponding to a full ban or exclusive selection
	// of a token, respectively. The exact behavior of a given bias score varies
	// by model.
	LogitBias map[string]*int32

	// A value that controls the emission of log probabilities for the provided number of most likely tokens within a completions
	// response.
	LogProbs *int32

	// The maximum number of tokens to generate.
	MaxTokens *int32

	// REQUIRED: DeploymentID specifies the name of the deployment (for Azure OpenAI) or model (for OpenAI) to use for this request.
	DeploymentID string

	// The number of completions choices that should be generated per provided prompt as part of an overall completions response.
	// Because this setting can generate many completions, it may quickly consume
	// your token quota. Use carefully and ensure reasonable settings for max_tokens and stop.
	N *int32

	// A value that influences the probability of generated tokens appearing based on their existing presence in generated text.
	// Positive values will make tokens less likely to appear when they already exist
	// and increase the model's likelihood to output new topics.
	PresencePenalty *float32

	// A collection of textual sequences that will end completions generation.
	Stop []string

	// The sampling temperature to use that controls the apparent creativity of generated completions. Higher values will make
	// output more random while lower values will make results more focused and
	// deterministic. It is not recommended to modify temperature and top_p for the same completions request as the interaction
	// of these two settings is difficult to predict.
	Temperature *float32

	// An alternative to sampling with temperature called nucleus sampling. This value causes the model to consider the results
	// of tokens with the provided probability mass. As an example, a value of 0.15
	// will cause only the tokens comprising the top 15% of probability mass to be considered. It is not recommended to modify
	// temperature and top_p for the same completions request as the interaction of
	// these two settings is difficult to predict.
	TopP *float32

	// An identifier for the caller or end user of the operation. This may be used for tracking or rate-limiting purposes.
	User *string
}

// CompletionsUsage - Representation of the token counts processed for a completions request. Counts consider all tokens across
// prompts, choices, choice alternates, best_of generations, and other consumers.
type CompletionsUsage struct {
	// REQUIRED; The number of tokens generated across all completions emissions.
	CompletionTokens *int32

	// REQUIRED; The number of tokens in the provided prompts for the completions request.
	PromptTokens *int32

	// REQUIRED; The total number of tokens processed for the completions request and response.
	TotalTokens *int32
}

// ContentFilterResults - Information about the content filtering category, if it has been detected.
type ContentFilterResults struct {
	// REQUIRED; Describes language attacks or uses that include pejorative or discriminatory language with reference to a person
	// or identity group on the basis of certain differentiating attributes of these groups
	// including but not limited to race, ethnicity, nationality, gender identity and expression, sexual orientation, religion,
	// immigration status, ability status, personal appearance, and body size.
	Hate *ContentFilterResultsHate

	// REQUIRED; Describes language related to physical actions intended to purposely hurt, injure, or damage one’s body, or kill
	// oneself.
	SelfHarm *ContentFilterResultsSelfHarm

	// REQUIRED; Describes language related to anatomical organs and genitals, romantic relationships, acts portrayed in erotic
	// or affectionate terms, physical sexual acts, including those portrayed as an assault or a
	// forced sexual violent act against one’s will, prostitution, pornography, and abuse.
	Sexual *ContentFilterResultsSexual

	// REQUIRED; Describes language related to physical actions intended to hurt, injure, damage, or kill someone or something;
	// describes weapons, etc.
	Violence *ContentFilterResultsViolence
}

// ContentFilterResultsHate - Describes language attacks or uses that include pejorative or discriminatory language with reference
// to a person or identity group on the basis of certain differentiating attributes of these groups
// including but not limited to race, ethnicity, nationality, gender identity and expression, sexual orientation, religion,
// immigration status, ability status, personal appearance, and body size.
type ContentFilterResultsHate struct {
	// REQUIRED; A value indicating whether or not the content has been filtered.
	Filtered *bool

	// REQUIRED; Ratings for the intensity and risk level of filtered content.
	Severity *ContentFilterSeverity
}

// ContentFilterResultsSelfHarm - Describes language related to physical actions intended to purposely hurt, injure, or damage
// one’s body, or kill oneself.
type ContentFilterResultsSelfHarm struct {
	// REQUIRED; A value indicating whether or not the content has been filtered.
	Filtered *bool

	// REQUIRED; Ratings for the intensity and risk level of filtered content.
	Severity *ContentFilterSeverity
}

// ContentFilterResultsSexual - Describes language related to anatomical organs and genitals, romantic relationships, acts
// portrayed in erotic or affectionate terms, physical sexual acts, including those portrayed as an assault or a
// forced sexual violent act against one’s will, prostitution, pornography, and abuse.
type ContentFilterResultsSexual struct {
	// REQUIRED; A value indicating whether or not the content has been filtered.
	Filtered *bool

	// REQUIRED; Ratings for the intensity and risk level of filtered content.
	Severity *ContentFilterSeverity
}

// ContentFilterResultsViolence - Describes language related to physical actions intended to hurt, injure, damage, or kill
// someone or something; describes weapons, etc.
type ContentFilterResultsViolence struct {
	// REQUIRED; A value indicating whether or not the content has been filtered.
	Filtered *bool

	// REQUIRED; Ratings for the intensity and risk level of filtered content.
	Severity *ContentFilterSeverity
}

// Deployment - A specific deployment
type Deployment struct {
	// READ-ONLY; Specifies either the model deployment name (when using Azure OpenAI) or model name (when using non-Azure OpenAI)
	// to use for this request.
	DeploymentID *string
}

// EmbeddingItem - Representation of a single embeddings relatedness comparison.
type EmbeddingItem struct {
	// REQUIRED; List of embeddings value for the input prompt. These represent a measurement of the vector-based relatedness
	// of the provided input.
	Embedding []float32

	// REQUIRED; Index of the prompt to which the EmbeddingItem corresponds.
	Index *int32
}

// Embeddings - Representation of the response data from an embeddings request. Embeddings measure the relatedness of text
// strings and are commonly used for search, clustering, recommendations, and other similar
// scenarios.
type Embeddings struct {
	// REQUIRED; Embedding values for the prompts submitted in the request.
	Data []EmbeddingItem

	// REQUIRED; Usage counts for tokens input using the embeddings API.
	Usage *EmbeddingsUsage
}

// EmbeddingsOptions - The configuration information for an embeddings request. Embeddings measure the relatedness of text
// strings and are commonly used for search, clustering, recommendations, and other similar scenarios.
type EmbeddingsOptions struct {
	// REQUIRED; Input texts to get embeddings for, encoded as a an array of strings. Each input must not exceed 2048 tokens in
	// length.
	// Unless you are embedding code, we suggest replacing newlines (\n) in your input with a single space, as we have observed
	// inferior results when newlines are present.
	Input []string

	// REQUIRED: DeploymentID specifies the name of the deployment (for Azure OpenAI) or model (for OpenAI) to use for this request.
	DeploymentID string

	// An identifier for the caller or end user of the operation. This may be used for tracking or rate-limiting purposes.
	User *string
}

// EmbeddingsUsage - Usage counts for tokens input using the embeddings API.
type EmbeddingsUsage struct {
	// REQUIRED; Number of tokens sent in the original request.
	PromptTokens *int32

	// REQUIRED; Total number of tokens transacted in this request/response.
	TotalTokens *int32
}

// EmbeddingsUsageAutoGenerated - Measurement of the amount of tokens used in this request and response.
type EmbeddingsUsageAutoGenerated struct {
	// REQUIRED; Number of tokens sent in the original request.
	PromptTokens *int32

	// REQUIRED; Total number of tokens transacted in this request/response.
	TotalTokens *int32
}

// FunctionCall - The name and arguments of a function that should be called, as generated by the model.
type FunctionCall struct {
	// REQUIRED; The arguments to call the function with, as generated by the model in JSON format. Note that the model does not
	// always generate valid JSON, and may hallucinate parameters not defined by your function
	// schema. Validate the arguments in your code before calling your function.
	Arguments *string

	// REQUIRED; The name of the function to call.
	Name *string
}

// FunctionDefinition - The definition of a caller-specified function that chat completions may invoke in response to matching
// user input.
type FunctionDefinition struct {
	// REQUIRED; The name of the function to be called.
	Name *string

	// A description of what the function does. The model will use this description when selecting the function and interpreting
	// its parameters.
	Description *string

	// The parameters the functions accepts, described as a JSON Schema object.
	Parameters any
}

// FunctionName - A structure that specifies the exact name of a specific, request-provided function to use when processing
// a chat completions operation.
type FunctionName struct {
	// REQUIRED; The name of the function to call.
	Name *string
}

// ImageGenerationOptions - Represents the request data used to generate images.
type ImageGenerationOptions struct {
	// REQUIRED; A description of the desired images.
	Prompt *string

	// The number of images to generate (defaults to 1).
	N *int32

	// The format in which image generation response items should be presented. Azure OpenAI only supports URL response items.
	ResponseFormat *ImageGenerationResponseFormat

	// The desired size of the generated images. Must be one of 256x256, 512x512, or 1024x1024 (defaults to 1024x1024).
	Size *ImageSize

	// A unique identifier representing your end-user, which can help to monitor and detect abuse.
	User *string
}

// ImageGenerations - The result of the operation if the operation succeeded.
type ImageGenerations struct {
	// REQUIRED; A timestamp when this job or item was created (in unix epochs).
	Created *int64

	// REQUIRED; The images generated by the operator.
	Data []ImageGenerationsDataItem
}

// ImageLocation - An image response item that provides a URL from which an image may be accessed.
type ImageLocation struct {
	// REQUIRED; The URL that provides temporary access to download the generated image.
	URL *string
}

// ImagePayload - An image response item that directly represents the image data as a base64-encoded string.
type ImagePayload struct {
	// REQUIRED; The complete data for an image represented as a base64-encoded string.
	B64JSON *string
}

// PromptFilterResult - Content filtering results for a single prompt in the request.
type PromptFilterResult struct {
	// REQUIRED; The index of this prompt in the set of prompt results
	PromptIndex *int32

	// Content filtering results for this prompt
	ContentFilterResults *PromptFilterResultContentFilterResults
}

// PromptFilterResultContentFilterResults - Content filtering results for this prompt
type PromptFilterResultContentFilterResults struct {
	// REQUIRED; Describes language attacks or uses that include pejorative or discriminatory language with reference to a person
	// or identity group on the basis of certain differentiating attributes of these groups
	// including but not limited to race, ethnicity, nationality, gender identity and expression, sexual orientation, religion,
	// immigration status, ability status, personal appearance, and body size.
	Hate *ContentFilterResultsHate

	// REQUIRED; Describes language related to physical actions intended to purposely hurt, injure, or damage one’s body, or kill
	// oneself.
	SelfHarm *ContentFilterResultsSelfHarm

	// REQUIRED; Describes language related to anatomical organs and genitals, romantic relationships, acts portrayed in erotic
	// or affectionate terms, physical sexual acts, including those portrayed as an assault or a
	// forced sexual violent act against one’s will, prostitution, pornography, and abuse.
	Sexual *ContentFilterResultsSexual

	// REQUIRED; Describes language related to physical actions intended to hurt, injure, damage, or kill someone or something;
	// describes weapons, etc.
	Violence *ContentFilterResultsViolence
}
