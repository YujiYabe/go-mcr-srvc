// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: backend/internal/1_framework/parameter/grpc/person.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type V1ImmutableParameter struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	TraceID          string                 `protobuf:"bytes,1,opt,name=traceID,proto3" json:"traceID,omitempty"`                    // リクエスト情報: システム全体で一意な追跡ID。ログ追跡やデバッグに使用
	RequestStartTime int64                  `protobuf:"varint,2,opt,name=requestStartTime,proto3" json:"requestStartTime,omitempty"` // リクエスト情報: リクエストの開始時刻（Unix timestamp）
	ClientIP         string                 `protobuf:"bytes,3,opt,name=clientIP,proto3" json:"clientIP,omitempty"`                  // リクエスト情報: クライアントのIPアドレス
	UserAgent        string                 `protobuf:"bytes,4,opt,name=userAgent,proto3" json:"userAgent,omitempty"`                // リクエスト情報: クライアントのユーザーエージェント情報
	UserID           string                 `protobuf:"bytes,5,opt,name=userID,proto3" json:"userID,omitempty"`                      // ユーザー認証情報: ユーザーを一意に識別するID
	Permissions      []string               `protobuf:"bytes,6,rep,name=permissions,proto3" json:"permissions,omitempty"`            // ユーザー認証情報: ユーザーに付与された権限リスト
	AccessToken      string                 `protobuf:"bytes,7,opt,name=accessToken,proto3" json:"accessToken,omitempty"`            // ユーザー認証情報: ユーザー認証用のJWTトークン
	TenantID         string                 `protobuf:"bytes,8,opt,name=tenantID,proto3" json:"tenantID,omitempty"`                  // ビジネスコンテキスト: マルチテナント環境での所属テナントID
	Locale           string                 `protobuf:"bytes,9,opt,name=locale,proto3" json:"locale,omitempty"`                      // ビジネスコンテキスト: ユーザーの言語設定（例：ja-JP）
	Timezone         string                 `protobuf:"bytes,10,opt,name=timezone,proto3" json:"timezone,omitempty"`                 // ビジネスコンテキスト: ユーザーのタイムゾーン（例：Asia/Tokyo）
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *V1ImmutableParameter) Reset() {
	*x = V1ImmutableParameter{}
	mi := &file_backend_internal_1_framework_parameter_grpc_person_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *V1ImmutableParameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V1ImmutableParameter) ProtoMessage() {}

func (x *V1ImmutableParameter) ProtoReflect() protoreflect.Message {
	mi := &file_backend_internal_1_framework_parameter_grpc_person_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V1ImmutableParameter.ProtoReflect.Descriptor instead.
func (*V1ImmutableParameter) Descriptor() ([]byte, []int) {
	return file_backend_internal_1_framework_parameter_grpc_person_proto_rawDescGZIP(), []int{0}
}

func (x *V1ImmutableParameter) GetTraceID() string {
	if x != nil {
		return x.TraceID
	}
	return ""
}

func (x *V1ImmutableParameter) GetRequestStartTime() int64 {
	if x != nil {
		return x.RequestStartTime
	}
	return 0
}

func (x *V1ImmutableParameter) GetClientIP() string {
	if x != nil {
		return x.ClientIP
	}
	return ""
}

func (x *V1ImmutableParameter) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *V1ImmutableParameter) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *V1ImmutableParameter) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *V1ImmutableParameter) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *V1ImmutableParameter) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *V1ImmutableParameter) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *V1ImmutableParameter) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

type V1MutableParameter struct {
	state                 protoimpl.MessageState `protogen:"open.v1"`
	Timestamp             string                 `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                          // 分散システム制御情報: 処理の各段階での時刻記録
	RetryCount            int32                  `protobuf:"varint,2,opt,name=retryCount,proto3" json:"retryCount,omitempty"`                       // 分散システム制御情報: リトライ試行回数の現在値
	CircuitBreakerEnabled bool                   `protobuf:"varint,3,opt,name=circuitBreakerEnabled,proto3" json:"circuitBreakerEnabled,omitempty"` // 分散システム制御情報: サーキットブレーカーの現在の有効状態
	RateLimit             int32                  `protobuf:"varint,4,opt,name=rateLimit,proto3" json:"rateLimit,omitempty"`                         // 分散システム制御情報: 現在の速度制限値（リクエスト/秒）
	TimeoutSeconds        int32                  `protobuf:"varint,5,opt,name=timeoutSeconds,proto3" json:"timeoutSeconds,omitempty"`               // リクエスト情報: リクエスト処理の残り時間（秒）
	ResourceQuota         int32                  `protobuf:"varint,6,opt,name=resourceQuota,proto3" json:"resourceQuota,omitempty"`                 // デバッグ/モニタリング情報: 割り当てられたリソース使用量の制限値
	SamplingEnabled       bool                   `protobuf:"varint,7,opt,name=samplingEnabled,proto3" json:"samplingEnabled,omitempty"`             // デバッグ/モニタリング情報: パフォーマンスサンプリングの有効状態
	DebugMode             bool                   `protobuf:"varint,8,opt,name=debugMode,proto3" json:"debugMode,omitempty"`                         // デバッグ/モニタリング情報: デバッグモードの有効状態
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *V1MutableParameter) Reset() {
	*x = V1MutableParameter{}
	mi := &file_backend_internal_1_framework_parameter_grpc_person_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *V1MutableParameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V1MutableParameter) ProtoMessage() {}

func (x *V1MutableParameter) ProtoReflect() protoreflect.Message {
	mi := &file_backend_internal_1_framework_parameter_grpc_person_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V1MutableParameter.ProtoReflect.Descriptor instead.
func (*V1MutableParameter) Descriptor() ([]byte, []int) {
	return file_backend_internal_1_framework_parameter_grpc_person_proto_rawDescGZIP(), []int{1}
}

func (x *V1MutableParameter) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *V1MutableParameter) GetRetryCount() int32 {
	if x != nil {
		return x.RetryCount
	}
	return 0
}

func (x *V1MutableParameter) GetCircuitBreakerEnabled() bool {
	if x != nil {
		return x.CircuitBreakerEnabled
	}
	return false
}

func (x *V1MutableParameter) GetRateLimit() int32 {
	if x != nil {
		return x.RateLimit
	}
	return 0
}

func (x *V1MutableParameter) GetTimeoutSeconds() int32 {
	if x != nil {
		return x.TimeoutSeconds
	}
	return 0
}

func (x *V1MutableParameter) GetResourceQuota() int32 {
	if x != nil {
		return x.ResourceQuota
	}
	return 0
}

func (x *V1MutableParameter) GetSamplingEnabled() bool {
	if x != nil {
		return x.SamplingEnabled
	}
	return false
}

func (x *V1MutableParameter) GetDebugMode() bool {
	if x != nil {
		return x.DebugMode
	}
	return false
}

type V1CommonParameter struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Immutable     *V1ImmutableParameter  `protobuf:"bytes,1,opt,name=immutable,proto3" json:"immutable,omitempty"`
	Mutable       *V1MutableParameter    `protobuf:"bytes,2,opt,name=mutable,proto3" json:"mutable,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *V1CommonParameter) Reset() {
	*x = V1CommonParameter{}
	mi := &file_backend_internal_1_framework_parameter_grpc_person_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *V1CommonParameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V1CommonParameter) ProtoMessage() {}

func (x *V1CommonParameter) ProtoReflect() protoreflect.Message {
	mi := &file_backend_internal_1_framework_parameter_grpc_person_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V1CommonParameter.ProtoReflect.Descriptor instead.
func (*V1CommonParameter) Descriptor() ([]byte, []int) {
	return file_backend_internal_1_framework_parameter_grpc_person_proto_rawDescGZIP(), []int{2}
}

func (x *V1CommonParameter) GetImmutable() *V1ImmutableParameter {
	if x != nil {
		return x.Immutable
	}
	return nil
}

func (x *V1CommonParameter) GetMutable() *V1MutableParameter {
	if x != nil {
		return x.Mutable
	}
	return nil
}

type V1GetPersonByConditionRequest struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	V1CommonParameter *V1CommonParameter     `protobuf:"bytes,1,opt,name=v1CommonParameter,proto3,oneof" json:"v1CommonParameter,omitempty"` // V1CommonParameter
	V1PersonParameter *V1PersonParameter     `protobuf:"bytes,2,opt,name=v1PersonParameter,proto3" json:"v1PersonParameter,omitempty"`       // V1PersonParameterArray
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *V1GetPersonByConditionRequest) Reset() {
	*x = V1GetPersonByConditionRequest{}
	mi := &file_backend_internal_1_framework_parameter_grpc_person_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *V1GetPersonByConditionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V1GetPersonByConditionRequest) ProtoMessage() {}

func (x *V1GetPersonByConditionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_internal_1_framework_parameter_grpc_person_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V1GetPersonByConditionRequest.ProtoReflect.Descriptor instead.
func (*V1GetPersonByConditionRequest) Descriptor() ([]byte, []int) {
	return file_backend_internal_1_framework_parameter_grpc_person_proto_rawDescGZIP(), []int{3}
}

func (x *V1GetPersonByConditionRequest) GetV1CommonParameter() *V1CommonParameter {
	if x != nil {
		return x.V1CommonParameter
	}
	return nil
}

func (x *V1GetPersonByConditionRequest) GetV1PersonParameter() *V1PersonParameter {
	if x != nil {
		return x.V1PersonParameter
	}
	return nil
}

type V1GetPersonByConditionResponse struct {
	state                  protoimpl.MessageState  `protogen:"open.v1"`
	V1CommonParameter      *V1CommonParameter      `protobuf:"bytes,1,opt,name=v1CommonParameter,proto3" json:"v1CommonParameter,omitempty"`           // V1CommonParameter
	V1PersonParameterArray *V1PersonParameterArray `protobuf:"bytes,2,opt,name=v1PersonParameterArray,proto3" json:"v1PersonParameterArray,omitempty"` // V1PersonParameterArray
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *V1GetPersonByConditionResponse) Reset() {
	*x = V1GetPersonByConditionResponse{}
	mi := &file_backend_internal_1_framework_parameter_grpc_person_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *V1GetPersonByConditionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V1GetPersonByConditionResponse) ProtoMessage() {}

func (x *V1GetPersonByConditionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_internal_1_framework_parameter_grpc_person_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V1GetPersonByConditionResponse.ProtoReflect.Descriptor instead.
func (*V1GetPersonByConditionResponse) Descriptor() ([]byte, []int) {
	return file_backend_internal_1_framework_parameter_grpc_person_proto_rawDescGZIP(), []int{4}
}

func (x *V1GetPersonByConditionResponse) GetV1CommonParameter() *V1CommonParameter {
	if x != nil {
		return x.V1CommonParameter
	}
	return nil
}

func (x *V1GetPersonByConditionResponse) GetV1PersonParameterArray() *V1PersonParameterArray {
	if x != nil {
		return x.V1PersonParameterArray
	}
	return nil
}

type V1PersonParameter struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            *uint32                `protobuf:"varint,2,opt,name=id,proto3,oneof" json:"id,omitempty"`                                     // ID
	Name          *string                `protobuf:"bytes,3,opt,name=name,proto3,oneof" json:"name,omitempty"`                                  // Name
	MailAddress   *string                `protobuf:"bytes,4,opt,name=mail_address,json=mailAddress,proto3,oneof" json:"mail_address,omitempty"` // Mail Address
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *V1PersonParameter) Reset() {
	*x = V1PersonParameter{}
	mi := &file_backend_internal_1_framework_parameter_grpc_person_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *V1PersonParameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V1PersonParameter) ProtoMessage() {}

func (x *V1PersonParameter) ProtoReflect() protoreflect.Message {
	mi := &file_backend_internal_1_framework_parameter_grpc_person_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V1PersonParameter.ProtoReflect.Descriptor instead.
func (*V1PersonParameter) Descriptor() ([]byte, []int) {
	return file_backend_internal_1_framework_parameter_grpc_person_proto_rawDescGZIP(), []int{5}
}

func (x *V1PersonParameter) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *V1PersonParameter) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *V1PersonParameter) GetMailAddress() string {
	if x != nil && x.MailAddress != nil {
		return *x.MailAddress
	}
	return ""
}

type V1PersonParameterArray struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Persons       []*V1PersonParameter   `protobuf:"bytes,2,rep,name=persons,proto3" json:"persons,omitempty"` // Array of V1PersonParameter
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *V1PersonParameterArray) Reset() {
	*x = V1PersonParameterArray{}
	mi := &file_backend_internal_1_framework_parameter_grpc_person_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *V1PersonParameterArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V1PersonParameterArray) ProtoMessage() {}

func (x *V1PersonParameterArray) ProtoReflect() protoreflect.Message {
	mi := &file_backend_internal_1_framework_parameter_grpc_person_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V1PersonParameterArray.ProtoReflect.Descriptor instead.
func (*V1PersonParameterArray) Descriptor() ([]byte, []int) {
	return file_backend_internal_1_framework_parameter_grpc_person_proto_rawDescGZIP(), []int{6}
}

func (x *V1PersonParameterArray) GetPersons() []*V1PersonParameter {
	if x != nil {
		return x.Persons
	}
	return nil
}

var File_backend_internal_1_framework_parameter_grpc_person_proto protoreflect.FileDescriptor

var file_backend_internal_1_framework_parameter_grpc_person_proto_rawDesc = []byte{
	0x0a, 0x38, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x31, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x22, 0xc2, 0x02, 0x0a, 0x14, 0x56,
	0x31, 0x49, 0x6d, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x44, 0x12, 0x2a, 0x0a,
	0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x50, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x50, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x22,
	0xbc, 0x02, 0x0a, 0x12, 0x56, 0x31, 0x4d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x72, 0x79, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x15, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x42,
	0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x15, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x42, 0x72, 0x65, 0x61,
	0x6b, 0x65, 0x72, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x61,
	0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72,
	0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73,
	0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x51, 0x75, 0x6f, 0x74,
	0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69,
	0x6e, 0x67, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x62, 0x75, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x64, 0x65, 0x62, 0x75, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x22, 0x95,
	0x01, 0x0a, 0x11, 0x56, 0x31, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x09, 0x69, 0x6d, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0x49, 0x6d, 0x6d, 0x75, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x09, 0x69,
	0x6d, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x6d, 0x75, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0x4d, 0x75, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x07, 0x6d,
	0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x22, 0xdc, 0x01, 0x0a, 0x1d, 0x56, 0x31, 0x47, 0x65, 0x74,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x54, 0x0a, 0x11, 0x76, 0x31, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x11, 0x76, 0x31, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x4f,
	0x0a, 0x11, 0x76, 0x31, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x11, 0x76, 0x31,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x42,
	0x14, 0x0a, 0x12, 0x5f, 0x76, 0x31, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x22, 0xd1, 0x01, 0x0a, 0x1e, 0x56, 0x31, 0x47, 0x65, 0x74, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x11, 0x76, 0x31, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x11, 0x76, 0x31, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x5e, 0x0a, 0x16, 0x76, 0x31, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x41, 0x72,
	0x72, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x41, 0x72, 0x72, 0x61,
	0x79, 0x52, 0x16, 0x76, 0x31, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x41, 0x72, 0x72, 0x61, 0x79, 0x22, 0x8a, 0x01, 0x0a, 0x11, 0x56, 0x31,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12,
	0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x02, 0x69,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a,
	0x0c, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0b, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x55, 0x0a, 0x16, 0x56, 0x31, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x41, 0x72, 0x72, 0x61, 0x79,
	0x12, 0x3b, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x2e, 0x56, 0x31, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x52, 0x07, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x32, 0x7f, 0x0a,
	0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x75, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x2e, 0x56, 0x31, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x42, 0x79, 0x43, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x2e,
	0x56, 0x31, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x42, 0x79, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x54,
	0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x75, 0x6a,
	0x69, 0x59, 0x61, 0x62, 0x65, 0x2f, 0x67, 0x6f, 0x2d, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2d,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x31, 0x5f, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_backend_internal_1_framework_parameter_grpc_person_proto_rawDescOnce sync.Once
	file_backend_internal_1_framework_parameter_grpc_person_proto_rawDescData = file_backend_internal_1_framework_parameter_grpc_person_proto_rawDesc
)

func file_backend_internal_1_framework_parameter_grpc_person_proto_rawDescGZIP() []byte {
	file_backend_internal_1_framework_parameter_grpc_person_proto_rawDescOnce.Do(func() {
		file_backend_internal_1_framework_parameter_grpc_person_proto_rawDescData = protoimpl.X.CompressGZIP(file_backend_internal_1_framework_parameter_grpc_person_proto_rawDescData)
	})
	return file_backend_internal_1_framework_parameter_grpc_person_proto_rawDescData
}

var file_backend_internal_1_framework_parameter_grpc_person_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_backend_internal_1_framework_parameter_grpc_person_proto_goTypes = []any{
	(*V1ImmutableParameter)(nil),           // 0: grpc_parameter.V1ImmutableParameter
	(*V1MutableParameter)(nil),             // 1: grpc_parameter.V1MutableParameter
	(*V1CommonParameter)(nil),              // 2: grpc_parameter.V1CommonParameter
	(*V1GetPersonByConditionRequest)(nil),  // 3: grpc_parameter.V1GetPersonByConditionRequest
	(*V1GetPersonByConditionResponse)(nil), // 4: grpc_parameter.V1GetPersonByConditionResponse
	(*V1PersonParameter)(nil),              // 5: grpc_parameter.V1PersonParameter
	(*V1PersonParameterArray)(nil),         // 6: grpc_parameter.V1PersonParameterArray
}
var file_backend_internal_1_framework_parameter_grpc_person_proto_depIdxs = []int32{
	0, // 0: grpc_parameter.V1CommonParameter.immutable:type_name -> grpc_parameter.V1ImmutableParameter
	1, // 1: grpc_parameter.V1CommonParameter.mutable:type_name -> grpc_parameter.V1MutableParameter
	2, // 2: grpc_parameter.V1GetPersonByConditionRequest.v1CommonParameter:type_name -> grpc_parameter.V1CommonParameter
	5, // 3: grpc_parameter.V1GetPersonByConditionRequest.v1PersonParameter:type_name -> grpc_parameter.V1PersonParameter
	2, // 4: grpc_parameter.V1GetPersonByConditionResponse.v1CommonParameter:type_name -> grpc_parameter.V1CommonParameter
	6, // 5: grpc_parameter.V1GetPersonByConditionResponse.v1PersonParameterArray:type_name -> grpc_parameter.V1PersonParameterArray
	5, // 6: grpc_parameter.V1PersonParameterArray.persons:type_name -> grpc_parameter.V1PersonParameter
	3, // 7: grpc_parameter.Person.GetPersonByCondition:input_type -> grpc_parameter.V1GetPersonByConditionRequest
	4, // 8: grpc_parameter.Person.GetPersonByCondition:output_type -> grpc_parameter.V1GetPersonByConditionResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_backend_internal_1_framework_parameter_grpc_person_proto_init() }
func file_backend_internal_1_framework_parameter_grpc_person_proto_init() {
	if File_backend_internal_1_framework_parameter_grpc_person_proto != nil {
		return
	}
	file_backend_internal_1_framework_parameter_grpc_person_proto_msgTypes[3].OneofWrappers = []any{}
	file_backend_internal_1_framework_parameter_grpc_person_proto_msgTypes[5].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_backend_internal_1_framework_parameter_grpc_person_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_backend_internal_1_framework_parameter_grpc_person_proto_goTypes,
		DependencyIndexes: file_backend_internal_1_framework_parameter_grpc_person_proto_depIdxs,
		MessageInfos:      file_backend_internal_1_framework_parameter_grpc_person_proto_msgTypes,
	}.Build()
	File_backend_internal_1_framework_parameter_grpc_person_proto = out.File
	file_backend_internal_1_framework_parameter_grpc_person_proto_rawDesc = nil
	file_backend_internal_1_framework_parameter_grpc_person_proto_goTypes = nil
	file_backend_internal_1_framework_parameter_grpc_person_proto_depIdxs = nil
}
