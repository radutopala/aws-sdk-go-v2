// Code generated by smithy-go-codegen DO NOT EDIT.
package lexruntimeservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexruntimeservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sends user input to Amazon Lex. Client applications can use this API to send
// requests to Amazon Lex at runtime. Amazon Lex then interprets the user input
// using the machine learning model it built for the bot. In response, Amazon Lex
// returns the next message to convey to the user an optional responseCard to
// display. Consider the following example messages:
//
//     * For a user input "I
// would like a pizza", Amazon Lex might return a response with a message eliciting
// slot data (for example, PizzaSize): "What size pizza would you like?"
//
//     *
// After the user provides all of the pizza order information, Amazon Lex might
// return a response with a message to obtain user confirmation "Proceed with the
// pizza order?".
//
//     * After the user replies to a confirmation prompt with a
// "yes", Amazon Lex might return a conclusion statement: "Thank you, your cheese
// pizza has been ordered.".
//
//     <p> Not all Amazon Lex messages require a user
// response. For example, a conclusion statement does not require a response. Some
// messages require only a "yes" or "no" user response. In addition to the
// <code>message</code>, Amazon Lex provides additional context about the message
// in the response that you might use to enhance client behavior, for example, to
// display the appropriate client user interface. These are the
// <code>slotToElicit</code>, <code>dialogState</code>, <code>intentName</code>,
// and <code>slots</code> fields in the response. Consider the following examples:
// </p> <ul> <li> <p>If the message is to elicit slot data, Amazon Lex returns the
// following context information:</p> <ul> <li> <p> <code>dialogState</code> set to
// ElicitSlot </p> </li> <li> <p> <code>intentName</code> set to the intent name in
// the current context </p> </li> <li> <p> <code>slotToElicit</code> set to the
// slot name for which the <code>message</code> is eliciting information </p> </li>
// <li> <p> <code>slots</code> set to a map of slots, configured for the intent,
// with currently known values </p> </li> </ul> </li> <li> <p> If the message is a
// confirmation prompt, the <code>dialogState</code> is set to ConfirmIntent and
// <code>SlotToElicit</code> is set to null. </p> </li> <li> <p>If the message is a
// clarification prompt (configured for the intent) that indicates that user intent
// is not understood, the <code>dialogState</code> is set to ElicitIntent and
// <code>slotToElicit</code> is set to null. </p> </li> </ul> <p> In addition,
// Amazon Lex also returns your application-specific
// <code>sessionAttributes</code>. For more information, see <a
// href="https://docs.aws.amazon.com/lex/latest/dg/context-mgmt.html">Managing
// Conversation Context</a>. </p>
func (c *Client) PostText(ctx context.Context, params *PostTextInput, optFns ...func(*Options)) (*PostTextOutput, error) {
	stack := middleware.NewStack("PostText", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	awsmiddleware.AddResolveServiceEndpointMiddleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	v4.AddHTTPSignerMiddlewares(stack, options)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPostText(options.Region), middleware.Before)
	addawsRestjson1_serdeOpPostTextMiddlewares(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "PostText",
			Err:           err,
		}
	}
	out := result.(*PostTextOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PostTextInput struct {
	// The alias of the Amazon Lex bot.
	BotAlias *string
	// The name of the Amazon Lex bot.
	BotName *string
	// The text that the user entered (Amazon Lex interprets this text).
	InputText *string
	// Request-specific information passed between Amazon Lex and a client application.
	// The namespace x-amz-lex: is reserved for special attributes. Don't create any
	// request attributes with the prefix x-amz-lex:. For more information, see Setting
	// Request Attributes
	// (https://docs.aws.amazon.com/lex/latest/dg/context-mgmt.html#context-mgmt-request-attribs).
	RequestAttributes map[string]*string
	// Application-specific information passed between Amazon Lex and a client
	// application. For more information, see Setting Session Attributes
	// (https://docs.aws.amazon.com/lex/latest/dg/context-mgmt.html#context-mgmt-session-attribs).
	SessionAttributes map[string]*string
	// The ID of the client application user. Amazon Lex uses this to identify a user's
	// conversation with your bot. At runtime, each request must contain the userID
	// field. To decide the user ID to use for your application, consider the following
	// factors.
	//
	//     * The userID field must not contain any personally identifiable
	// information of the user, for example, name, personal identification numbers, or
	// other end user personal information.
	//
	//     * If you want a user to start a
	// conversation on one device and continue on another device, use a user-specific
	// identifier.
	//
	//     * If you want the same user to be able to have two independent
	// conversations on two different devices, choose a device-specific identifier.
	//
	//
	// * A user can't have two independent conversations with two different versions of
	// the same bot. For example, a user can't have a conversation with the PROD and
	// BETA versions of the same bot. If you anticipate that a user will need to have
	// conversation with two different versions, for example, while testing, include
	// the bot alias in the user ID to separate the two conversations.
	UserId *string
}

type PostTextOutput struct {
	// Identifies the current state of the user interaction. Amazon Lex returns one of
	// the following values as dialogState. The client can optionally use this
	// information to customize the user interface.
	//
	//     * ElicitIntent - Amazon Lex
	// wants to elicit user intent. For example, a user might utter an intent ("I want
	// to order a pizza"). If Amazon Lex cannot infer the user intent from this
	// utterance, it will return this dialogState.
	//
	//     * ConfirmIntent - Amazon Lex is
	// expecting a "yes" or "no" response. For example, Amazon Lex wants user
	// confirmation before fulfilling an intent. Instead of a simple "yes" or "no," a
	// user might respond with additional information. For example, "yes, but make it
	// thick crust pizza" or "no, I want to order a drink". Amazon Lex can process such
	// additional information (in these examples, update the crust type slot value, or
	// change intent from OrderPizza to OrderDrink).
	//
	//     * ElicitSlot - Amazon Lex is
	// expecting a slot value for the current intent. For example, suppose that in the
	// response Amazon Lex sends this message: "What size pizza would you like?". A
	// user might reply with the slot value (e.g., "medium"). The user might also
	// provide additional information in the response (e.g., "medium thick crust
	// pizza"). Amazon Lex can process such additional information appropriately.
	//
	//
	// * Fulfilled - Conveys that the Lambda function configured for the intent has
	// successfully fulfilled the intent.  </li> <li> <p>
	// <code>ReadyForFulfillment</code> - Conveys that the client has to fulfill the
	// intent. </p> </li> <li> <p> <code>Failed</code> - Conveys that the conversation
	// with the user failed. </p> <p> This can happen for various reasons including
	// that the user did not provide an appropriate response to prompts from the
	// service (you can configure how many times Amazon Lex can prompt a user for
	// specific information), or the Lambda function failed to fulfill the intent. </p>
	// </li> </ul>
	DialogState types.DialogState
	// The current user intent that Amazon Lex is aware of.
	IntentName *string
	// The message to convey to the user. The message can come from the bot's
	// configuration or from a Lambda function. If the intent is not configured with a
	// Lambda function, or if the Lambda function returned Delegate as the
	// dialogAction.type its response, Amazon Lex decides on the next course of action
	// and selects an appropriate message from the bot's configuration based on the
	// current interaction context. For example, if Amazon Lex isn't able to understand
	// user input, it uses a clarification prompt message. When you create an intent
	// you can assign messages to groups. When messages are assigned to groups Amazon
	// Lex returns one message from each group in the response. The message field is an
	// escaped JSON string containing the messages. For more information about the
	// structure of the JSON string returned, see msg-prompts-formats (). If the Lambda
	// function returns a message, Amazon Lex passes it to the client in its response.
	Message *string
	// The format of the response message. One of the following values:
	//
	//     *
	// PlainText - The message contains plain UTF-8 text.
	//
	//     * CustomPayload - The
	// message is a custom format defined by the Lambda function.
	//
	//     * SSML - The
	// message contains text formatted for voice output.
	//
	//     * Composite - The message
	// contains an escaped JSON object containing one or more messages from the groups
	// that messages were assigned to when the intent was created.
	MessageFormat types.MessageFormatType
	// Represents the options that the user has to respond to the current prompt.
	// Response Card can come from the bot configuration (in the Amazon Lex console,
	// choose the settings button next to a slot) or from a code hook (Lambda
	// function).
	ResponseCard *types.ResponseCard
	// The sentiment expressed in and utterance. When the bot is configured to send
	// utterances to Amazon Comprehend for sentiment analysis, this field contains the
	// result of the analysis.
	SentimentResponse *types.SentimentResponse
	// A map of key-value pairs representing the session-specific context information.
	SessionAttributes map[string]*string
	// A unique identifier for the session.
	SessionId *string
	// If the dialogState value is ElicitSlot, returns the name of the slot for which
	// Amazon Lex is eliciting a value.
	SlotToElicit *string
	// The intent slots that Amazon Lex detected from the user input in the
	// conversation. Amazon Lex creates a resolution list containing likely values for
	// a slot. The value that it returns is determined by the valueSelectionStrategy
	// selected when the slot type was created or updated. If valueSelectionStrategy is
	// set to ORIGINAL_VALUE, the value provided by the user is returned, if the user
	// value is similar to the slot values. If valueSelectionStrategy is set to
	// TOP_RESOLUTION Amazon Lex returns the first value in the resolution list or, if
	// there is no resolution list, null. If you don't specify a
	// valueSelectionStrategy, the default is ORIGINAL_VALUE.
	Slots map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPostTextMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add("&awsRestjson1_serializeOpPostText{}", middleware.After)
	stack.Deserialize.Add("&awsRestjson1_deserializeOpPostText{}", middleware.After)
}

func newServiceMetadataMiddleware_opPostText(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "Lex Runtime Service",
		ServiceID:      "lexruntimeservice",
		EndpointPrefix: "lexruntimeservice",
		SigningName:    "lex",
		OperationName:  "PostText",
	}
}
