// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// A list of possible transcriptions for the audio.
type Alternative struct {

	// Contains the entities identified as personally identifiable information (PII) in
	// the transcription output.
	Entities []Entity

	// One or more alternative interpretations of the input audio.
	Items []Item

	// The text that was transcribed from the audio.
	Transcript *string

	noSmithyDocumentSerde
}

// Provides a wrapper for the audio chunks that you are sending. For information on
// audio encoding in Amazon Transcribe, see Speech input
// (https://docs.aws.amazon.com/transcribe/latest/dg/input.html). For information
// on audio encoding formats in Amazon Transcribe Medical, see Speech input
// (https://docs.aws.amazon.com/transcribe/latest/dg/input-med.html).
type AudioEvent struct {

	// An audio blob that contains the next part of the audio that you want to
	// transcribe. The maximum audio chunk size is 32 KB.
	AudioChunk []byte

	noSmithyDocumentSerde
}

// Represents the audio stream from your application to Amazon Transcribe.
//
// The following types satisfy this interface:
//  AudioStreamMemberAudioEvent
type AudioStream interface {
	isAudioStream()
}

// A blob of audio from your application. You audio stream consists of one or more
// audio events. For information on audio encoding formats in Amazon Transcribe,
// see Speech input (https://docs.aws.amazon.com/transcribe/latest/dg/input.html).
// For information on audio encoding formats in Amazon Transcribe Medical, see
// Speech input (https://docs.aws.amazon.com/transcribe/latest/dg/input-med.html).
// For more information on stream encoding in Amazon Transcribe, see Event stream
// encoding (https://docs.aws.amazon.com/transcribe/latest/dg/event-stream.html).
// For information on stream encoding in Amazon Transcribe Medical, see Event
// stream encoding
// (https://docs.aws.amazon.com/transcribe/latest/dg/event-stream-med.html).
type AudioStreamMemberAudioEvent struct {
	Value AudioEvent

	noSmithyDocumentSerde
}

func (*AudioStreamMemberAudioEvent) isAudioStream() {}

// The entity identified as personally identifiable information (PII).
type Entity struct {

	// The category of of information identified in this entity; for example, PII.
	Category *string

	// A value between zero and one that Amazon Transcribe assigns to PII identified in
	// the source audio. Larger values indicate a higher confidence in PII
	// identification.
	Confidence *float64

	// The words in the transcription output that have been identified as a PII entity.
	Content *string

	// The end time of speech that was identified as PII.
	EndTime float64

	// The start time of speech that was identified as PII.
	StartTime float64

	// The type of PII identified in this entity; for example, name or credit card
	// number.
	Type *string

	noSmithyDocumentSerde
}

// A word, phrase, or punctuation mark that is transcribed from the input audio.
type Item struct {

	// A value between 0 and 1 for an item that is a confidence score that Amazon
	// Transcribe assigns to each word or phrase that it transcribes.
	Confidence *float64

	// The word or punctuation that was recognized in the input audio.
	Content *string

	// The offset from the beginning of the audio stream to the end of the audio that
	// resulted in the item.
	EndTime float64

	// If speaker identification is enabled, shows the speakers identified in the
	// real-time stream.
	Speaker *string

	// If partial result stabilization has been enabled, indicates whether the word or
	// phrase in the item is stable. If Stable is true, the result is stable.
	Stable *bool

	// The offset from the beginning of the audio stream to the beginning of the audio
	// that resulted in the item.
	StartTime float64

	// The type of the item. PRONUNCIATION indicates that the item is a word that was
	// recognized in the input audio. PUNCTUATION indicates that the item was
	// interpreted as a pause in the input audio.
	Type ItemType

	// Indicates whether a word in the item matches a word in the vocabulary filter
	// you've chosen for your real-time stream. If true then a word in the item matches
	// your vocabulary filter.
	VocabularyFilterMatch bool

	noSmithyDocumentSerde
}

// A list of possible transcriptions for the audio.
type MedicalAlternative struct {

	// Contains the medical entities identified as personal health information in the
	// transcription output.
	Entities []MedicalEntity

	// A list of objects that contains words and punctuation marks that represents one
	// or more interpretations of the input audio.
	Items []MedicalItem

	// The text that was transcribed from the audio.
	Transcript *string

	noSmithyDocumentSerde
}

// The medical entity identified as personal health information.
type MedicalEntity struct {

	// The type of personal health information of the medical entity.
	Category *string

	// A value between zero and one that Amazon Transcribe Medical assigned to the
	// personal health information that it identified in the source audio. Larger
	// values indicate that Amazon Transcribe Medical has higher confidence in the
	// personal health information that it identified.
	Confidence *float64

	// The word or words in the transcription output that have been identified as a
	// medical entity.
	Content *string

	// The end time of the speech that was identified as a medical entity.
	EndTime float64

	// The start time of the speech that was identified as a medical entity.
	StartTime float64

	noSmithyDocumentSerde
}

// A word, phrase, or punctuation mark that is transcribed from the input audio.
type MedicalItem struct {

	// A value between 0 and 1 for an item that is a confidence score that Amazon
	// Transcribe Medical assigns to each word that it transcribes.
	Confidence *float64

	// The word or punctuation mark that was recognized in the input audio.
	Content *string

	// The number of seconds into an audio stream that indicates the creation time of
	// an item.
	EndTime float64

	// If speaker identification is enabled, shows the integer values that correspond
	// to the different speakers identified in the stream. For example, if the value of
	// Speaker in the stream is either a 0 or a 1, that indicates that Amazon
	// Transcribe Medical has identified two speakers in the stream. The value of 0
	// corresponds to one speaker and the value of 1 corresponds to the other speaker.
	Speaker *string

	// The number of seconds into an audio stream that indicates the creation time of
	// an item.
	StartTime float64

	// The type of the item. PRONUNCIATION indicates that the item is a word that was
	// recognized in the input audio. PUNCTUATION indicates that the item was
	// interpreted as a pause in the input audio, such as a period to indicate the end
	// of a sentence.
	Type ItemType

	noSmithyDocumentSerde
}

// The results of transcribing a portion of the input audio stream.
type MedicalResult struct {

	// A list of possible transcriptions of the audio. Each alternative typically
	// contains one Item that contains the result of the transcription.
	Alternatives []MedicalAlternative

	// When channel identification is enabled, Amazon Transcribe Medical transcribes
	// the speech from each audio channel separately. You can use ChannelId to retrieve
	// the transcription results for a single channel in your audio stream.
	ChannelId *string

	// The time, in seconds, from the beginning of the audio stream to the end of the
	// result.
	EndTime float64

	// Amazon Transcribe Medical divides the incoming audio stream into segments at
	// natural points in the audio. Transcription results are returned based on these
	// segments. The IsPartial field is true to indicate that Amazon Transcribe Medical
	// has additional transcription data to send. The IsPartial field is false to
	// indicate that this is the last transcription result for the segment.
	IsPartial bool

	// A unique identifier for the result.
	ResultId *string

	// The time, in seconds, from the beginning of the audio stream to the beginning of
	// the result.
	StartTime float64

	noSmithyDocumentSerde
}

// The medical transcript in a MedicalTranscriptEvent.
type MedicalTranscript struct {

	// MedicalResult objects that contain the results of transcribing a portion of the
	// input audio stream. The array can be empty.
	Results []MedicalResult

	noSmithyDocumentSerde
}

// Represents a set of transcription results from the server to the client. It
// contains one or more segments of the transcription.
type MedicalTranscriptEvent struct {

	// The transcription of the audio stream. The transcription is composed of all of
	// the items in the results list.
	Transcript *MedicalTranscript

	noSmithyDocumentSerde
}

// Represents the transcription result stream from Amazon Transcribe Medical to
// your application.
//
// The following types satisfy this interface:
//  MedicalTranscriptResultStreamMemberTranscriptEvent
type MedicalTranscriptResultStream interface {
	isMedicalTranscriptResultStream()
}

// A portion of the transcription of the audio stream. Events are sent periodically
// from Amazon Transcribe Medical to your application. The event can be a partial
// transcription of a section of the audio stream, or it can be the entire
// transcription of that portion of the audio stream.
type MedicalTranscriptResultStreamMemberTranscriptEvent struct {
	Value MedicalTranscriptEvent

	noSmithyDocumentSerde
}

func (*MedicalTranscriptResultStreamMemberTranscriptEvent) isMedicalTranscriptResultStream() {}

// The result of transcribing a portion of the input audio stream.
type Result struct {

	// A list of possible transcriptions for the audio. Each alternative typically
	// contains one item that contains the result of the transcription.
	Alternatives []Alternative

	// When channel identification is enabled, Amazon Transcribe transcribes the speech
	// from each audio channel separately. You can use ChannelId to retrieve the
	// transcription results for a single channel in your audio stream.
	ChannelId *string

	// The offset in seconds from the beginning of the audio stream to the end of the
	// result.
	EndTime float64

	// Amazon Transcribe divides the incoming audio stream into segments at natural
	// points in the audio. Transcription results are returned based on these segments.
	// The IsPartial field is true to indicate that Amazon Transcribe has additional
	// transcription data to send, false to indicate that this is the last
	// transcription result for the segment.
	IsPartial bool

	// A unique identifier for the result.
	ResultId *string

	// The offset in seconds from the beginning of the audio stream to the beginning of
	// the result.
	StartTime float64

	noSmithyDocumentSerde
}

// The transcription in a TranscriptEvent.
type Transcript struct {

	// Result objects that contain the results of transcribing a portion of the input
	// audio stream. The array can be empty.
	Results []Result

	noSmithyDocumentSerde
}

// Represents a set of transcription results from the server to the client. It
// contains one or more segments of the transcription.
type TranscriptEvent struct {

	// The transcription of the audio stream. The transcription is composed of all of
	// the items in the results list.
	Transcript *Transcript

	noSmithyDocumentSerde
}

// Represents the transcription result stream from Amazon Transcribe to your
// application.
//
// The following types satisfy this interface:
//  TranscriptResultStreamMemberTranscriptEvent
type TranscriptResultStream interface {
	isTranscriptResultStream()
}

// A portion of the transcription of the audio stream. Events are sent periodically
// from Amazon Transcribe to your application. The event can be a partial
// transcription of a section of the audio stream, or it can be the entire
// transcription of that portion of the audio stream.
type TranscriptResultStreamMemberTranscriptEvent struct {
	Value TranscriptEvent

	noSmithyDocumentSerde
}

func (*TranscriptResultStreamMemberTranscriptEvent) isTranscriptResultStream() {}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isAudioStream()                   {}
func (*UnknownUnionMember) isMedicalTranscriptResultStream() {}
func (*UnknownUnionMember) isTranscriptResultStream()        {}
