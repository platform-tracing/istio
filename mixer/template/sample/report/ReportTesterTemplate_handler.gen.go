// Copyright 2017 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// THIS FILE IS AUTOMATICALLY GENERATED.

package istio_mixer_adapter_sample_report

import (
	"context"
	"net"
	"time"

	"istio.io/istio/mixer/pkg/adapter"
)

// Fully qualified name of the template
const TemplateName = "report"

// Instance is constructed by Mixer for the 'report' template.
type Instance struct {
	// Name of the instance as specified in configuration.
	Name string

	Value interface{}

	Dimensions map[string]interface{}

	Int64Primitive int64

	BoolPrimitive bool

	DoublePrimitive float64

	StringPrimitive string

	Int64Map map[string]int64

	TimeStamp time.Time

	Duration time.Duration

	Res1 *Res1
}

type Res1 struct {
	Value interface{}

	Dimensions map[string]interface{}

	Int64Primitive int64

	BoolPrimitive bool

	DoublePrimitive float64

	StringPrimitive string

	Int64Map map[string]int64

	TimeStamp time.Time

	Duration time.Duration

	Res2 *Res2

	Res2Map map[string]*Res2
}

type Res2 struct {
	Value interface{}

	Dimensions map[string]interface{}

	Int64Primitive int64

	TimeStamp time.Time

	Duration time.Duration

	IpAddr net.IP

	DnsName adapter.DNSName

	EmailAddr adapter.EmailAddress

	Uri adapter.URI
}

// HandlerBuilder must be implemented by adapters if they want to
// process data associated with the 'report' template.
//
// Mixer uses this interface to call into the adapter at configuration time to configure
// it with adapter-specific configuration as well as all template-specific type information.
type HandlerBuilder interface {
	adapter.HandlerBuilder

	// SetReportTypes is invoked by Mixer to pass the template-specific Type information for instances that an adapter
	// may receive at runtime. The type information describes the shape of the instance.
	SetReportTypes(map[string]*Type /*Instance name -> Type*/)
}

// Handler must be implemented by adapter code if it wants to
// process data associated with the 'report' template.
//
// Mixer uses this interface to call into the adapter at request time in order to dispatch
// created instances to the adapter. Adapters take the incoming instances and do what they
// need to achieve their primary function.
//
// The name of each instance can be used as a key into the Type map supplied to the adapter
// at configuration time via the method 'SetReportTypes'.
// These Type associated with an instance describes the shape of the instance
type Handler interface {
	adapter.Handler

	// HandleReport is called by Mixer at request time to deliver instances to
	// to an adapter.
	HandleReport(context.Context, []*Instance) error
}