// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribestreaming

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream/eventstreamapi"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribestreaming/types"
	"github.com/aws/smithy-go/middleware"
	smithysync "github.com/aws/smithy-go/sync"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"sync"
	"time"
)

// Starts a bidirectional HTTP/2 stream where audio is streamed to Amazon
// Transcribe and the transcription results are streamed to your application. The
// following are encoded as HTTP/2 headers:
//
// * x-amzn-transcribe-language-code
//
// *
// x-amzn-transcribe-media-encoding
//
// * x-amzn-transcribe-sample-rate
//
// *
// x-amzn-transcribe-session-id
//
// See the  SDK for Go API Reference
// (https://docs.aws.amazon.com/sdk-for-go/api/service/transcribestreamingservice/#TranscribeStreamingService.StartStreamTranscription)
// for more detail.
func (c *Client) StartStreamTranscription(ctx context.Context, params *StartStreamTranscriptionInput, optFns ...func(*Options)) (*StartStreamTranscriptionOutput, error) {
	if params == nil {
		params = &StartStreamTranscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartStreamTranscription", params, optFns, c.addOperationStartStreamTranscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartStreamTranscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartStreamTranscriptionInput struct {

	// Indicates the source language used in the input audio stream.
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// The encoding used for the input audio.
	//
	// This member is required.
	MediaEncoding types.MediaEncoding

	// The sample rate, in Hertz, of the input audio. We suggest that you use 8,000 Hz
	// for low quality audio and 16,000 Hz for high quality audio.
	//
	// This member is required.
	MediaSampleRateHertz *int32

	// Set this field to PII to identify personally identifiable information (PII) in
	// the transcription output. Content identification is performed only upon complete
	// transcription of the audio segments. You can’t set both
	// ContentIdentificationType and ContentRedactionType in the same request. If you
	// set both, your request returns a BadRequestException.
	ContentIdentificationType types.ContentIdentificationType

	// Set this field to PII to redact personally identifiable information (PII) in the
	// transcription output. Content redaction is performed only upon complete
	// transcription of the audio segments. You can’t set both ContentRedactionType and
	// ContentIdentificationType in the same request. If you set both, your request
	// returns a BadRequestException.
	ContentRedactionType types.ContentRedactionType

	// When true, instructs Amazon Transcribe to process each audio channel separately
	// and then merge the transcription output of each channel into a single
	// transcription. Amazon Transcribe also produces a transcription of each item. An
	// item includes the start time, end time, and any alternative transcriptions. You
	// can't set both ShowSpeakerLabel and EnableChannelIdentification in the same
	// request. If you set both, your request returns a BadRequestException.
	EnableChannelIdentification bool

	// When true, instructs Amazon Transcribe to present transcription results that
	// have the partial results stabilized. Normally, any word or phrase from one
	// partial result can change in a subsequent partial result. With partial results
	// stabilization enabled, only the last few words of one partial result can change
	// in another partial result.
	EnablePartialResultsStabilization bool

	// The name of the language model you want to use.
	LanguageModelName *string

	// The number of channels that are in your audio stream.
	NumberOfChannels *int32

	// You can use this field to set the stability level of the transcription results.
	// A higher stability level means that the transcription results are less likely to
	// change. Higher stability levels can come with lower overall transcription
	// accuracy.
	PartialResultsStability types.PartialResultsStability

	// List the PII entity types you want to identify or redact. In order to specify
	// entity types, you must have either ContentIdentificationType or
	// ContentRedactionType enabled. PiiEntityTypes is an optional parameter with a
	// default value of ALL.
	PiiEntityTypes *string

	// A identifier for the transcription session. Use this parameter when you want to
	// retry a session. If you don't provide a session ID, Amazon Transcribe will
	// generate one for you and return it in the response.
	SessionId *string

	// When true, enables speaker identification in your real-time stream.
	ShowSpeakerLabel bool

	// The manner in which you use your vocabulary filter to filter words in your
	// transcript. Remove removes filtered words from your transcription results. Mask
	// masks filtered words with a *** in your transcription results. Tag keeps the
	// filtered words in your transcription results and tags them. The tag appears as
	// VocabularyFilterMatch equal to True
	VocabularyFilterMethod types.VocabularyFilterMethod

	// The name of the vocabulary filter you've created that is unique to your account.
	// Provide the name in this field to successfully use it in a stream.
	VocabularyFilterName *string

	// The name of the vocabulary to use when processing the transcription job.
	VocabularyName *string

	noSmithyDocumentSerde
}

type StartStreamTranscriptionOutput struct {

	// Shows whether content identification was enabled in this stream.
	ContentIdentificationType types.ContentIdentificationType

	// Shows whether content redaction was enabled in this stream.
	ContentRedactionType types.ContentRedactionType

	// Shows whether channel identification has been enabled in the stream.
	EnableChannelIdentification bool

	// Shows whether partial results stabilization has been enabled in the stream.
	EnablePartialResultsStabilization bool

	// The language code for the input audio stream.
	LanguageCode types.LanguageCode

	LanguageModelName *string

	// The encoding used for the input audio stream.
	MediaEncoding types.MediaEncoding

	// The sample rate for the input audio stream. Use 8,000 Hz for low quality audio
	// and 16,000 Hz for high quality audio.
	MediaSampleRateHertz *int32

	// The number of channels identified in the stream.
	NumberOfChannels *int32

	// If partial results stabilization has been enabled in the stream, shows the
	// stability level.
	PartialResultsStability types.PartialResultsStability

	// Lists the PII entity types you specified in your request.
	PiiEntityTypes *string

	// An identifier for the streaming transcription.
	RequestId *string

	// An identifier for a specific transcription session.
	SessionId *string

	// Shows whether speaker identification was enabled in the stream.
	ShowSpeakerLabel bool

	// The vocabulary filtering method used in the real-time stream.
	VocabularyFilterMethod types.VocabularyFilterMethod

	// The name of the vocabulary filter used in your real-time stream.
	VocabularyFilterName *string

	// The name of the vocabulary used when processing the stream.
	VocabularyName *string

	eventStream *StartStreamTranscriptionEventStream

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

// GetStream returns the type to interact with the event stream.
func (o *StartStreamTranscriptionOutput) GetStream() *StartStreamTranscriptionEventStream {
	return o.eventStream
}

func (c *Client) addOperationStartStreamTranscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartStreamTranscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartStreamTranscription{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addEventStreamStartStreamTranscriptionMiddleware(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddRequireMinimumProtocol(stack, 2, 0); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddStreamingEventsPayload(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = eventstreamapi.AddInitializeStreamWriter(stack); err != nil {
		return err
	}
	if err = addOpStartStreamTranscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartStreamTranscription(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opStartStreamTranscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "StartStreamTranscription",
	}
}

// StartStreamTranscriptionEventStream provides the event stream handling for the StartStreamTranscription operation.
//
// For testing and mocking the event stream this type should be initialized via
// the NewStartStreamTranscriptionEventStream constructor function. Using the functional options
// to pass in nested mock behavior.
type StartStreamTranscriptionEventStream struct {
	// AudioStreamWriter is the EventStream writer for the AudioStream events. This
	// value is automatically set by the SDK when the API call is made Use this member
	// when unit testing your code with the SDK to mock out the EventStream Writer.
	//
	// Must not be nil.
	Writer AudioStreamWriter

	// TranscriptResultStreamReader is the EventStream reader for the
	// TranscriptResultStream events. This value is automatically set by the SDK when
	// the API call is made Use this member when unit testing your code with the SDK to
	// mock out the EventStream Reader.
	//
	// Must not be nil.
	Reader TranscriptResultStreamReader

	done      chan struct{}
	closeOnce sync.Once
	err       *smithysync.OnceErr
}

// NewStartStreamTranscriptionEventStream initializes an StartStreamTranscriptionEventStream.
// This function should only be used for testing and mocking the StartStreamTranscriptionEventStream
// stream within your application.
//
// The Writer member must be set before writing events to the stream.
//
// The Reader member must be set before reading events from the stream.
func NewStartStreamTranscriptionEventStream(optFns ...func(*StartStreamTranscriptionEventStream)) *StartStreamTranscriptionEventStream {
	es := &StartStreamTranscriptionEventStream{
		done: make(chan struct{}),
		err:  smithysync.NewOnceErr(),
	}
	for _, fn := range optFns {
		fn(es)
	}
	return es
}

// Send writes the event to the stream blocking until the event is written.
// Returns an error if the event was not written.
func (es *StartStreamTranscriptionEventStream) Send(ctx context.Context, event types.AudioStream) error {
	return es.Writer.Send(ctx, event)
}

// Events returns a channel to read events from.
func (es *StartStreamTranscriptionEventStream) Events() <-chan types.TranscriptResultStream {
	return es.Reader.Events()
}

// Close closes the stream. This will also cause the stream to be closed.
// Close must be called when done using the stream API. Not calling Close
// may result in resource leaks.
//
// Will close the underlying EventStream writer and reader, and no more events can be
// sent or received.
func (es *StartStreamTranscriptionEventStream) Close() error {
	es.closeOnce.Do(es.safeClose)
	return es.Err()
}

func (es *StartStreamTranscriptionEventStream) safeClose() {
	close(es.done)

	t := time.NewTicker(time.Second)
	defer t.Stop()
	writeCloseDone := make(chan error)
	go func() {
		if err := es.Writer.Close(); err != nil {
			es.err.SetError(err)
		}
		close(writeCloseDone)
	}()
	select {
	case <-t.C:
	case <-writeCloseDone:
	}

	es.Reader.Close()
}

// Err returns any error that occurred while reading or writing EventStream Events
// from the service API's response. Returns nil if there were no errors.
func (es *StartStreamTranscriptionEventStream) Err() error {
	if err := es.err.Err(); err != nil {
		return err
	}

	if err := es.Writer.Err(); err != nil {
		return err
	}

	if err := es.Reader.Err(); err != nil {
		return err
	}

	return nil
}

func (es *StartStreamTranscriptionEventStream) waitStreamClose() {
	type errorSet interface {
		ErrorSet() <-chan struct{}
	}

	var inputErrCh <-chan struct{}
	if v, ok := es.Writer.(errorSet); ok {
		inputErrCh = v.ErrorSet()
	}

	var outputErrCh <-chan struct{}
	if v, ok := es.Reader.(errorSet); ok {
		outputErrCh = v.ErrorSet()
	}
	var outputClosedCh <-chan struct{}
	if v, ok := es.Reader.(interface{ Closed() <-chan struct{} }); ok {
		outputClosedCh = v.Closed()
	}

	select {
	case <-es.done:
	case <-inputErrCh:
		es.err.SetError(es.Writer.Err())
		es.Close()

	case <-outputErrCh:
		es.err.SetError(es.Reader.Err())
		es.Close()

	case <-outputClosedCh:
		if err := es.Reader.Err(); err != nil {
			es.err.SetError(es.Reader.Err())
		}
		es.Close()

	}
}
