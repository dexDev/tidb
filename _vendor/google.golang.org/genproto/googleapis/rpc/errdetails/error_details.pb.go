// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/rpc/error_details.proto

/*
Package errdetails is a generated protocol buffer package.

It is generated from these files:
	google/rpc/error_details.proto

It has these top-level messages:
	RetryInfo
	DebugInfo
	QuotaFailure
	BadRequest
	RequestInfo
	ResourceInfo
	Help
	LocalizedMessage
*/
package errdetails

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Describes when the clients can retry a failed request. Clients could ignore
// the recommendation here or retry when this information is missing from error
// responses.
//
// It's always recommended that clients should use exponential backoff when
// retrying.
//
// Clients should wait until `retry_delay` amount of time has passed since
// receiving the error response before retrying.  If retrying requests also
// fail, clients should use an exponential backoff scheme to gradually increase
// the delay between retries based on `retry_delay`, until either a maximum
// number of retires have been reached or a maximum retry delay cap has been
// reached.
type RetryInfo struct {
	// Clients should wait at least this long between retrying the same request.
	RetryDelay *google_protobuf.Duration `protobuf:"bytes,1,opt,name=retry_delay,json=retryDelay" json:"retry_delay,omitempty"`
}

func (m *RetryInfo) Reset()                    { *m = RetryInfo{} }
func (m *RetryInfo) String() string            { return proto.CompactTextString(m) }
func (*RetryInfo) ProtoMessage()               {}
func (*RetryInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RetryInfo) GetRetryDelay() *google_protobuf.Duration {
	if m != nil {
		return m.RetryDelay
	}
	return nil
}

// Describes additional debugging info.
type DebugInfo struct {
	// The stack trace entries indicating where the error occurred.
	StackEntries []string `protobuf:"bytes,1,rep,name=stack_entries,json=stackEntries" json:"stack_entries,omitempty"`
	// Additional debugging information provided by the server.
	Detail string `protobuf:"bytes,2,opt,name=detail" json:"detail,omitempty"`
}

func (m *DebugInfo) Reset()                    { *m = DebugInfo{} }
func (m *DebugInfo) String() string            { return proto.CompactTextString(m) }
func (*DebugInfo) ProtoMessage()               {}
func (*DebugInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DebugInfo) GetStackEntries() []string {
	if m != nil {
		return m.StackEntries
	}
	return nil
}

func (m *DebugInfo) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

// Describes how a quota check failed.
//
// For example if a daily limit was exceeded for the calling project,
// a service could respond with a QuotaFailure detail containing the project
// id and the description of the quota limit that was exceeded.  If the
// calling project hasn't enabled the service in the developer console, then
// a service could respond with the project id and set `service_disabled`
// to true.
//
// Also see RetryDetail and Help types for other details about handling a
// quota failure.
type QuotaFailure struct {
	// Describes all quota violations.
	Violations []*QuotaFailure_Violation `protobuf:"bytes,1,rep,name=violations" json:"violations,omitempty"`
}

func (m *QuotaFailure) Reset()                    { *m = QuotaFailure{} }
func (m *QuotaFailure) String() string            { return proto.CompactTextString(m) }
func (*QuotaFailure) ProtoMessage()               {}
func (*QuotaFailure) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *QuotaFailure) GetViolations() []*QuotaFailure_Violation {
	if m != nil {
		return m.Violations
	}
	return nil
}

// A message type used to describe a single quota violation.  For example, a
// daily quota or a custom quota that was exceeded.
type QuotaFailure_Violation struct {
	// The subject on which the quota check failed.
	// For example, "clientip:<ip address of client>" or "project:<Google
	// developer project id>".
	Subject string `protobuf:"bytes,1,opt,name=subject" json:"subject,omitempty"`
	// A description of how the quota check failed. Clients can use this
	// description to find more about the quota configuration in the service's
	// public documentation, or find the relevant quota limit to adjust through
	// developer console.
	//
	// For example: "Service disabled" or "Daily Limit for read operations
	// exceeded".
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *QuotaFailure_Violation) Reset()                    { *m = QuotaFailure_Violation{} }
func (m *QuotaFailure_Violation) String() string            { return proto.CompactTextString(m) }
func (*QuotaFailure_Violation) ProtoMessage()               {}
func (*QuotaFailure_Violation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func (m *QuotaFailure_Violation) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *QuotaFailure_Violation) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Describes violations in a client request. This error type focuses on the
// syntactic aspects of the request.
type BadRequest struct {
	// Describes all violations in a client request.
	FieldViolations []*BadRequest_FieldViolation `protobuf:"bytes,1,rep,name=field_violations,json=fieldViolations" json:"field_violations,omitempty"`
}

func (m *BadRequest) Reset()                    { *m = BadRequest{} }
func (m *BadRequest) String() string            { return proto.CompactTextString(m) }
func (*BadRequest) ProtoMessage()               {}
func (*BadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BadRequest) GetFieldViolations() []*BadRequest_FieldViolation {
	if m != nil {
		return m.FieldViolations
	}
	return nil
}

// A message type used to describe a single bad request field.
type BadRequest_FieldViolation struct {
	// A path leading to a field in the request body. The value will be a
	// sequence of dot-separated identifiers that identify a protocol buffer
	// field. E.g., "field_violations.field" would identify this field.
	Field string `protobuf:"bytes,1,opt,name=field" json:"field,omitempty"`
	// A description of why the request element is bad.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *BadRequest_FieldViolation) Reset()                    { *m = BadRequest_FieldViolation{} }
func (m *BadRequest_FieldViolation) String() string            { return proto.CompactTextString(m) }
func (*BadRequest_FieldViolation) ProtoMessage()               {}
func (*BadRequest_FieldViolation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

func (m *BadRequest_FieldViolation) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *BadRequest_FieldViolation) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Contains metadata about the request that clients can attach when filing a bug
// or providing other forms of feedback.
type RequestInfo struct {
	// An opaque string that should only be interpreted by the service generating
	// it. For example, it can be used to identify requests in the service's logs.
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId" json:"request_id,omitempty"`
	// Any data that was used to serve this request. For example, an encrypted
	// stack trace that can be sent back to the service provider for debugging.
	ServingData string `protobuf:"bytes,2,opt,name=serving_data,json=servingData" json:"serving_data,omitempty"`
}

func (m *RequestInfo) Reset()                    { *m = RequestInfo{} }
func (m *RequestInfo) String() string            { return proto.CompactTextString(m) }
func (*RequestInfo) ProtoMessage()               {}
func (*RequestInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RequestInfo) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *RequestInfo) GetServingData() string {
	if m != nil {
		return m.ServingData
	}
	return ""
}

// Describes the resource that is being accessed.
type ResourceInfo struct {
	// A name for the type of resource being accessed, e.g. "sql table",
	// "cloud storage bucket", "file", "Google calendar"; or the type URL
	// of the resource: e.g. "type.googleapis.com/google.pubsub.v1.Topic".
	ResourceType string `protobuf:"bytes,1,opt,name=resource_type,json=resourceType" json:"resource_type,omitempty"`
	// The name of the resource being accessed.  For example, a shared calendar
	// name: "example.com_4fghdhgsrgh@group.calendar.google.com", if the current
	// error is [google.rpc.Code.PERMISSION_DENIED][google.rpc.Code.PERMISSION_DENIED].
	ResourceName string `protobuf:"bytes,2,opt,name=resource_name,json=resourceName" json:"resource_name,omitempty"`
	// The owner of the resource (optional).
	// For example, "user:<owner email>" or "project:<Google developer project
	// id>".
	Owner string `protobuf:"bytes,3,opt,name=owner" json:"owner,omitempty"`
	// Describes what error is encountered when accessing this resource.
	// For example, updating a cloud project may require the `writer` permission
	// on the developer console project.
	Description string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
}

func (m *ResourceInfo) Reset()                    { *m = ResourceInfo{} }
func (m *ResourceInfo) String() string            { return proto.CompactTextString(m) }
func (*ResourceInfo) ProtoMessage()               {}
func (*ResourceInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ResourceInfo) GetResourceType() string {
	if m != nil {
		return m.ResourceType
	}
	return ""
}

func (m *ResourceInfo) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *ResourceInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ResourceInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Provides links to documentation or for performing an out of band action.
//
// For example, if a quota check failed with an error indicating the calling
// project hasn't enabled the accessed service, this can contain a URL pointing
// directly to the right place in the developer console to flip the bit.
type Help struct {
	// URL(s) pointing to additional information on handling the current error.
	Links []*Help_Link `protobuf:"bytes,1,rep,name=links" json:"links,omitempty"`
}

func (m *Help) Reset()                    { *m = Help{} }
func (m *Help) String() string            { return proto.CompactTextString(m) }
func (*Help) ProtoMessage()               {}
func (*Help) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Help) GetLinks() []*Help_Link {
	if m != nil {
		return m.Links
	}
	return nil
}

// Describes a URL link.
type Help_Link struct {
	// Describes what the link offers.
	Description string `protobuf:"bytes,1,opt,name=description" json:"description,omitempty"`
	// The URL of the link.
	Url string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
}

func (m *Help_Link) Reset()                    { *m = Help_Link{} }
func (m *Help_Link) String() string            { return proto.CompactTextString(m) }
func (*Help_Link) ProtoMessage()               {}
func (*Help_Link) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

func (m *Help_Link) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Help_Link) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// Provides a localized error message that is safe to return to the user
// which can be attached to an RPC error.
type LocalizedMessage struct {
	// The locale used following the specification defined at
	// http://www.rfc-editor.org/rfc/bcp/bcp47.txt.
	// Examples are: "en-US", "fr-CH", "es-MX"
	Locale string `protobuf:"bytes,1,opt,name=locale" json:"locale,omitempty"`
	// The localized error message in the above locale.
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *LocalizedMessage) Reset()                    { *m = LocalizedMessage{} }
func (m *LocalizedMessage) String() string            { return proto.CompactTextString(m) }
func (*LocalizedMessage) ProtoMessage()               {}
func (*LocalizedMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *LocalizedMessage) GetLocale() string {
	if m != nil {
		return m.Locale
	}
	return ""
}

func (m *LocalizedMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RetryInfo)(nil), "google.rpc.RetryInfo")
	proto.RegisterType((*DebugInfo)(nil), "google.rpc.DebugInfo")
	proto.RegisterType((*QuotaFailure)(nil), "google.rpc.QuotaFailure")
	proto.RegisterType((*QuotaFailure_Violation)(nil), "google.rpc.QuotaFailure.Violation")
	proto.RegisterType((*BadRequest)(nil), "google.rpc.BadRequest")
	proto.RegisterType((*BadRequest_FieldViolation)(nil), "google.rpc.BadRequest.FieldViolation")
	proto.RegisterType((*RequestInfo)(nil), "google.rpc.RequestInfo")
	proto.RegisterType((*ResourceInfo)(nil), "google.rpc.ResourceInfo")
	proto.RegisterType((*Help)(nil), "google.rpc.Help")
	proto.RegisterType((*Help_Link)(nil), "google.rpc.Help.Link")
	proto.RegisterType((*LocalizedMessage)(nil), "google.rpc.LocalizedMessage")
}

func init() { proto.RegisterFile("google/rpc/error_details.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 551 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xc1, 0x6f, 0xd3, 0x3e,
	0x14, 0xc7, 0x95, 0x75, 0xdb, 0x4f, 0x79, 0xed, 0x6f, 0x94, 0x08, 0x50, 0xa9, 0x04, 0x2a, 0x41,
	0x48, 0x95, 0x90, 0x52, 0x69, 0xdc, 0xc6, 0x01, 0xa9, 0x64, 0x5b, 0x27, 0x0d, 0x28, 0x11, 0xe2,
	0xc0, 0x25, 0x72, 0x93, 0xd7, 0xc8, 0xd4, 0x8d, 0x83, 0xed, 0x0c, 0x95, 0xbf, 0x82, 0x3b, 0x37,
	0x4e, 0xfc, 0x99, 0xc8, 0xb1, 0xbd, 0xa6, 0xeb, 0x85, 0x5b, 0xbe, 0xcf, 0x1f, 0x7f, 0xf3, 0x7d,
	0x89, 0x9f, 0xe1, 0x69, 0xc1, 0x79, 0xc1, 0x70, 0x22, 0xaa, 0x6c, 0x82, 0x42, 0x70, 0x91, 0xe6,
	0xa8, 0x08, 0x65, 0x32, 0xaa, 0x04, 0x57, 0x3c, 0x00, 0xb3, 0x1e, 0x89, 0x2a, 0x1b, 0x3a, 0xb6,
	0x59, 0x59, 0xd4, 0xcb, 0x49, 0x5e, 0x0b, 0xa2, 0x28, 0x2f, 0x0d, 0x1b, 0x5e, 0x82, 0x9f, 0xa0,
	0x12, 0x9b, 0xab, 0x72, 0xc9, 0x83, 0x33, 0xe8, 0x0a, 0x2d, 0xd2, 0x1c, 0x19, 0xd9, 0x0c, 0xbc,
	0x91, 0x37, 0xee, 0x9e, 0x3e, 0x8e, 0xac, 0x9d, 0xb3, 0x88, 0x62, 0x6b, 0x91, 0x40, 0x43, 0xc7,
	0x1a, 0x0e, 0x67, 0xe0, 0xc7, 0xb8, 0xa8, 0x8b, 0xc6, 0xe8, 0x39, 0xfc, 0x2f, 0x15, 0xc9, 0x56,
	0x29, 0x96, 0x4a, 0x50, 0x94, 0x03, 0x6f, 0xd4, 0x19, 0xfb, 0x49, 0xaf, 0x29, 0x9e, 0x9b, 0x5a,
	0xf0, 0x08, 0x8e, 0x4d, 0xee, 0xc1, 0xc1, 0xc8, 0x1b, 0xfb, 0x89, 0x55, 0xe1, 0x2f, 0x0f, 0x7a,
	0x1f, 0x6b, 0xae, 0xc8, 0x05, 0xa1, 0xac, 0x16, 0x18, 0x4c, 0x01, 0x6e, 0x28, 0x67, 0xcd, 0x3b,
	0x8d, 0x55, 0xf7, 0x34, 0x8c, 0xb6, 0x4d, 0x46, 0x6d, 0x3a, 0xfa, 0xec, 0xd0, 0xa4, 0xb5, 0x6b,
	0x78, 0x09, 0xfe, 0xed, 0x42, 0x30, 0x80, 0xff, 0x64, 0xbd, 0xf8, 0x8a, 0x99, 0x6a, 0x7a, 0xf4,
	0x13, 0x27, 0x83, 0x11, 0x74, 0x73, 0x94, 0x99, 0xa0, 0x95, 0x06, 0x6d, 0xb0, 0x76, 0x29, 0xfc,
	0xe3, 0x01, 0x4c, 0x49, 0x9e, 0xe0, 0xb7, 0x1a, 0xa5, 0x0a, 0xe6, 0xd0, 0x5f, 0x52, 0x64, 0x79,
	0xba, 0x97, 0xf0, 0x45, 0x3b, 0xe1, 0x76, 0x47, 0x74, 0xa1, 0xf1, 0x6d, 0xc8, 0x7b, 0xcb, 0x1d,
	0x2d, 0x87, 0x33, 0x38, 0xd9, 0x45, 0x82, 0x07, 0x70, 0xd4, 0x40, 0x36, 0xac, 0x11, 0xff, 0x10,
	0xf5, 0x03, 0x74, 0xed, 0x4b, 0x9b, 0x9f, 0xf2, 0x04, 0x40, 0x18, 0x99, 0x52, 0xe7, 0xe5, 0xdb,
	0xca, 0x55, 0x1e, 0x3c, 0x83, 0x9e, 0x44, 0x71, 0x43, 0xcb, 0x22, 0xcd, 0x89, 0x22, 0xce, 0xd0,
	0xd6, 0x62, 0xa2, 0x48, 0xf8, 0xd3, 0x83, 0x5e, 0x82, 0x92, 0xd7, 0x22, 0x43, 0xf7, 0x9f, 0x85,
	0xd5, 0xa9, 0xda, 0x54, 0x68, 0x5d, 0x7b, 0xae, 0xf8, 0x69, 0x53, 0xe1, 0x0e, 0x54, 0x92, 0x35,
	0x5a, 0xe7, 0x5b, 0xe8, 0x3d, 0x59, 0xa3, 0xee, 0x91, 0x7f, 0x2f, 0x51, 0x0c, 0x3a, 0xa6, 0xc7,
	0x46, 0xdc, 0xed, 0xf1, 0x70, 0xbf, 0x47, 0x0e, 0x87, 0x33, 0x64, 0x55, 0xf0, 0x12, 0x8e, 0x18,
	0x2d, 0x57, 0xee, 0xe3, 0x3f, 0x6c, 0x7f, 0x7c, 0x0d, 0x44, 0xd7, 0xb4, 0x5c, 0x25, 0x86, 0x19,
	0x9e, 0xc1, 0xa1, 0x96, 0x77, 0xed, 0xbd, 0x3d, 0xfb, 0xa0, 0x0f, 0x9d, 0x5a, 0xb8, 0x03, 0xaa,
	0x1f, 0xc3, 0x18, 0xfa, 0xd7, 0x3c, 0x23, 0x8c, 0xfe, 0xc0, 0xfc, 0x1d, 0x4a, 0x49, 0x0a, 0xd4,
	0x27, 0x99, 0xe9, 0x9a, 0xeb, 0xdf, 0x2a, 0x7d, 0xce, 0xd6, 0x06, 0xb1, 0x0e, 0x4e, 0x4e, 0x19,
	0x9c, 0x64, 0x7c, 0xdd, 0x0a, 0x39, 0xbd, 0x7f, 0xae, 0x27, 0x39, 0x36, 0x83, 0x3c, 0xd7, 0xa3,
	0x36, 0xf7, 0xbe, 0xbc, 0xb1, 0x40, 0xc1, 0x19, 0x29, 0x8b, 0x88, 0x8b, 0x62, 0x52, 0x60, 0xd9,
	0x0c, 0xe2, 0xc4, 0x2c, 0x91, 0x8a, 0x4a, 0x77, 0x11, 0xd8, 0x5b, 0xe0, 0xf5, 0xf6, 0xf1, 0xf7,
	0x41, 0x27, 0x99, 0xbf, 0x5d, 0x1c, 0x37, 0x3b, 0x5e, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xc0,
	0x5e, 0xc6, 0x6f, 0x39, 0x04, 0x00, 0x00,
}
