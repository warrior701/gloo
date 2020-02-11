// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gateway/api/v1/route_table.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//
// The **RouteTable** is a child routing object for the Gloo Gateway.
//
// A **RouteTable** gets built into the complete routing configuration when it is referenced by a `delegateAction`,
// either in a parent VirtualService or another RouteTable.
//
// Routes specified in a RouteTable must have their paths start with the prefix provided in the parent's matcher.
//
// For example, the following configuration:
//
// ```
// virtualService: mydomain.com
// match: /a
// delegate: a-routes
// ---
// routeTable: a-routes
// match: /1
//
// ```
//
// would *not be valid*, while
//
// ```
// virtualService: mydomain.com
// match: /a
// delegate: a-routes
// ---
// routeTable: a-routes
// match: /a/1
//
// ```
//
// *would* be valid.
//
//
// A complete configuration might look as follows:
//
// ```yaml
// apiVersion: gateway.solo.io/v1
// kind: VirtualService
// metadata:
//   name: 'any'
//   namespace: 'any'
// spec:
//   virtualHost:
//     domains:
//     - 'any.com'
//     routes:
//     - matchers:
//       - prefix: '/a' # delegate ownership of routes for `any.com/a`
//       delegateAction:
//         ref:
//           name: 'a-routes'
//           namespace: 'a'
//     - matchers:
//       - prefix: '/b' # delegate ownership of routes for `any.com/b`
//       delegateAction:
//         ref:
//           name: 'b-routes'
//           namespace: 'b'
// ```
//
// * A root-level **VirtualService** which delegates routing to to the `a-routes` and `b-routes` **RouteTables**.
// * Routes with `delegateActions` can only use a `prefix` matcher.
//
// ```yaml
// apiVersion: gateway.solo.io/v1
// kind: RouteTable
// metadata:
//   name: 'a-routes'
//   namespace: 'a'
// spec:
//   routes:
//     - matchers:
//       # the path matchers in this RouteTable must begin with the prefix `/a/`
//       - prefix: '/a/1'
//       routeAction:
//         single:
//           upstream:
//             name: 'foo-upstream'
//
//     - matchers:
//       - prefix: '/a/2'
//       routeAction:
//         single:
//           upstream:
//             name: 'bar-upstream'
// ```
//
// * A **RouteTable** which defines two routes.
//
// ```yaml
// apiVersion: gateway.solo.io/v1
// kind: RouteTable
// metadata:
//   name: 'b-routes'
//   namespace: 'b'
// spec:
//   routes:
//     - matchers:
//       # the path matchers in this RouteTable must begin with the prefix `/b/`
//       - regex: '/b/3'
//       routeAction:
//         single:
//           upstream:
//             name: 'bar-upstream'
//     - matchers:
//       - prefix: '/b/c/'
//       # routes in the RouteTable can perform any action, including a delegateAction
//       delegateAction:
//         ref:
//           name: 'c-routes'
//           namespace: 'c'
//
// ```
//
// * A **RouteTable** which both *defines a route* and *delegates to* another **RouteTable**.
//
//
// ```yaml
// apiVersion: gateway.solo.io/v1
// kind: RouteTable
// metadata:
//   name: 'c-routes'
//   namespace: 'c'
// spec:
//   routes:
//     - matchers:
//       - exact: '/b/c/4'
//       routeAction:
//         single:
//           upstream:
//             name: 'qux-upstream'
// ```
//
// * A RouteTable which is a child of another route table.
//
//
// Would produce the following route config for `mydomain.com`:
//
// ```
// /a/1 -> foo-upstream
// /a/2 -> bar-upstream
// /b/3 -> baz-upstream
// /b/c/4 -> qux-upstream
// ```
//
type RouteTable struct {
	// The list of routes for the route table
	Routes []*Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	// When a delegated route defines a `RouteTableSelector` that matches multiple route tables, Gloo will inspect this
	// field to determine the order in which the route tables are to be evaluated. This determines the order in which
	// the routes will appear on the final `Proxy` resource. The field is optional; if no value is specified, the weight
	// defaults to 0 (zero).
	//
	// Gloo will process the route tables matched by a selector in ascending order by weight and collect the routes of
	// each route table in the order they are defined. If multiple route tables define the same weight, Gloo will sort the
	// routes which belong to those tables to avoid short-circuiting (e.g. making sure `/foo/bar` comes before `/foo`).
	// In this scenario, Gloo will also alert the user by adding a warning to the status of the parent resource
	// (the one that specifies the `RouteTableSelector`).
	Weight *types.Int32Value `protobuf:"bytes,2,opt,name=weight,proto3" json:"weight,omitempty"`
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by gloo during validation
	Status core.Status `protobuf:"bytes,6,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RouteTable) Reset()         { *m = RouteTable{} }
func (m *RouteTable) String() string { return proto.CompactTextString(m) }
func (*RouteTable) ProtoMessage()    {}
func (*RouteTable) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d1ea5a66e7f9a13, []int{0}
}
func (m *RouteTable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteTable.Unmarshal(m, b)
}
func (m *RouteTable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteTable.Marshal(b, m, deterministic)
}
func (m *RouteTable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteTable.Merge(m, src)
}
func (m *RouteTable) XXX_Size() int {
	return xxx_messageInfo_RouteTable.Size(m)
}
func (m *RouteTable) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteTable.DiscardUnknown(m)
}

var xxx_messageInfo_RouteTable proto.InternalMessageInfo

func (m *RouteTable) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *RouteTable) GetWeight() *types.Int32Value {
	if m != nil {
		return m.Weight
	}
	return nil
}

func (m *RouteTable) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *RouteTable) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func init() {
	proto.RegisterType((*RouteTable)(nil), "gateway.solo.io.RouteTable")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gateway/api/v1/route_table.proto", fileDescriptor_4d1ea5a66e7f9a13)
}

var fileDescriptor_4d1ea5a66e7f9a13 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0xee, 0xd2, 0x40,
	0x14, 0xc6, 0x2d, 0x36, 0xd5, 0xcc, 0xdf, 0x68, 0x6c, 0x08, 0x69, 0x50, 0x81, 0xb0, 0x62, 0xe3,
	0x4c, 0x6c, 0x37, 0x86, 0xc4, 0x85, 0xec, 0x8c, 0x71, 0x53, 0x8d, 0x0b, 0x37, 0x64, 0x5a, 0x86,
	0x61, 0xa4, 0xf0, 0x9a, 0x99, 0x57, 0xc0, 0x2d, 0x67, 0xf0, 0x10, 0x1e, 0xc1, 0x23, 0x78, 0x0a,
	0x16, 0xde, 0x00, 0x13, 0xf7, 0xa6, 0xd3, 0x29, 0x51, 0x12, 0x8d, 0xbb, 0xce, 0x7c, 0xdf, 0xfb,
	0x5e, 0xbf, 0x5f, 0x86, 0xbc, 0x94, 0x0a, 0x57, 0x55, 0x46, 0x73, 0xd8, 0x30, 0x03, 0x05, 0x3c,
	0x55, 0xc0, 0x64, 0x01, 0xc0, 0x4a, 0x0d, 0x1f, 0x45, 0x8e, 0x86, 0x49, 0x8e, 0x62, 0xcf, 0x3f,
	0x31, 0x5e, 0x2a, 0xb6, 0x7b, 0xc6, 0x34, 0x54, 0x28, 0xe6, 0xc8, 0xb3, 0x42, 0xd0, 0x52, 0x03,
	0x42, 0xf8, 0xc0, 0x39, 0x68, 0x3d, 0x4f, 0x15, 0xf4, 0xbb, 0x12, 0x24, 0x58, 0x8d, 0xd5, 0x5f,
	0x8d, 0xad, 0x1f, 0x8a, 0x03, 0x36, 0x97, 0xe2, 0x80, 0xee, 0x6e, 0x20, 0x01, 0x64, 0x21, 0x98,
	0x3d, 0x65, 0xd5, 0x92, 0xed, 0x35, 0x2f, 0x4b, 0xa1, 0x4d, 0xab, 0xdb, 0x5f, 0x5a, 0x2b, 0x6c,
	0xb7, 0x6f, 0x04, 0xf2, 0x05, 0x47, 0xee, 0xf4, 0xc7, 0xd7, 0xba, 0x41, 0x8e, 0xd5, 0x5f, 0xa7,
	0xdb, 0xb3, 0xd3, 0xe3, 0x7f, 0x16, 0xdd, 0x29, 0x8d, 0x15, 0x2f, 0xe6, 0x46, 0xe8, 0x9d, 0xca,
	0x5d, 0xd9, 0xf1, 0xe7, 0x0e, 0x21, 0x69, 0x8d, 0xe0, 0x5d, 0x4d, 0x20, 0xa4, 0x24, 0xb0, 0x40,
	0x4c, 0xe4, 0x8d, 0x6e, 0x4f, 0x6e, 0xe2, 0x1e, 0xbd, 0x82, 0x41, 0xad, 0x39, 0x75, 0xae, 0x30,
	0x21, 0xc1, 0x5e, 0x28, 0xb9, 0xc2, 0xa8, 0x33, 0xf2, 0x26, 0x37, 0xf1, 0x23, 0xda, 0x10, 0xa0,
	0x2d, 0x01, 0xfa, 0x6a, 0x8b, 0x49, 0xfc, 0x9e, 0x17, 0x95, 0x48, 0x9d, 0x35, 0x7c, 0x4d, 0x82,
	0xa6, 0x57, 0x14, 0xd8, 0xa1, 0x2e, 0xcd, 0x41, 0x8b, 0xcb, 0x86, 0xb7, 0x56, 0x9b, 0x3d, 0xf9,
	0xfa, 0xd3, 0xf7, 0xbe, 0x9d, 0x86, 0xb7, 0x7e, 0x9c, 0x86, 0x0f, 0x51, 0x18, 0x5c, 0xa8, 0xe5,
	0x72, 0x3a, 0x56, 0x72, 0x0b, 0x5a, 0x8c, 0x53, 0x17, 0x11, 0x3e, 0x27, 0x77, 0x5b, 0x88, 0xd1,
	0x1d, 0x1b, 0xd7, 0xfb, 0x33, 0xee, 0x8d, 0x53, 0x67, 0x7e, 0x1d, 0x96, 0x5e, 0xdc, 0xd3, 0xfe,
	0xf1, 0xec, 0xfb, 0xa4, 0xa3, 0xf1, 0x78, 0xf6, 0xef, 0x87, 0xf7, 0x7e, 0x7b, 0x08, 0x66, 0xf6,
	0xa2, 0x5e, 0xfe, 0xe5, 0xfb, 0xc0, 0xfb, 0x90, 0xfc, 0xf7, 0x83, 0x2a, 0xd7, 0xd2, 0xb1, 0xce,
	0x02, 0x5b, 0x3f, 0xf9, 0x15, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x37, 0xc7, 0x35, 0x8e, 0x02, 0x00,
	0x00,
}

func (this *RouteTable) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteTable)
	if !ok {
		that2, ok := that.(RouteTable)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Routes) != len(that1.Routes) {
		return false
	}
	for i := range this.Routes {
		if !this.Routes[i].Equal(that1.Routes[i]) {
			return false
		}
	}
	if !this.Weight.Equal(that1.Weight) {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
