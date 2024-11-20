// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !js

package http

import _ "unsafe" // for linkname

// RoundTrip should be an internal detail,
// but widely used packages access it using linkname.
// Notable members of the hall of shame include:
//   - github.com/erda-project/erda-infra
//
// Do not remove or change the type signature.
// See go.dev/issue/67401.
//
//go:linkname badRoundTrip net/http.(*Transport).RoundTrip
func badRoundTrip(*Transport, *Request) (*Response, error)

// RoundTrip implements the [RoundTripper] interface.
//
// For higher-level HTTP client support (such as handling of cookies
// and redirects), see [Get], [Post], and the [Client] type.
//
// Like the RoundTripper interface, the error types returned
// by RoundTrip are unspecified.
func (t *Transport) RoundTrip(req *Request) (retVal0 *Response, retVal1 error) {
	if  callContext84013, _ := OtelOnEnterTrampoline_RoundTrip84013(&t, &req); false { /* NO_NEWWLINE_PLACEHOLDER */ ;	} else {  		defer OtelOnExitTrampoline_RoundTrip84013(callContext84013, &retVal0, &retVal1) ;	}  
	return t.roundTrip(req)
}

// Seeing is not always believing. The following template is a bit tricky, see
// trampoline.go for more details
// Struct Template
type CallContextImpl84013 struct {
	Params     []interface{}
	ReturnVals []interface{}
	SkipCall   bool
	Data       interface{}
}

func (c *CallContextImpl84013) SetSkipCall(skip bool)    { c.SkipCall = skip }
func (c *CallContextImpl84013) IsSkipCall() bool         { return c.SkipCall }
func (c *CallContextImpl84013) SetData(data interface{}) { c.Data = data }
func (c *CallContextImpl84013) GetData() interface{}     { return c.Data }
func (c *CallContextImpl84013) GetKeyData(key string) interface{} {
	if c.Data == nil {
		return nil
	}
	return c.Data.(map[string]interface{})[key]
}
func (c *CallContextImpl84013) SetKeyData(key string, val interface{}) {
	if c.Data == nil {
		c.Data = make(map[string]interface{})
	}
	c.Data.(map[string]interface{})[key] = val
}
func (c *CallContextImpl84013) HasKeyData(key string) bool {
	if c.Data == nil {
		return false
	}
	_, ok := c.Data.(map[string]interface{})[key]
	return ok
}
func (c *CallContextImpl84013) GetParam(idx int) interface{} {
	switch idx {
	case 0:
		return *(c.Params[0].(**Transport))
	case 1:
		return *(c.Params[1].(**Request))
	}
	return nil
}
func (c *CallContextImpl84013) SetParam(idx int, val interface{}) {
	if val == nil {
		c.Params[idx] = nil
		return
	}
	switch idx {
	case 0:
		*(c.Params[0].(**Transport)) = val.(*Transport)
	case 1:
		*(c.Params[1].(**Request)) = val.(*Request)
	}
}
func (c *CallContextImpl84013) GetReturnVal(idx int) interface{} {
	switch idx {
	case 0:
		return *(c.ReturnVals[0].(**Response))
	case 1:
		return *(c.ReturnVals[1].(*error))
	}
	return nil
}
func (c *CallContextImpl84013) SetReturnVal(idx int, val interface{}) {
	if val == nil {
		c.ReturnVals[idx] = nil
		return
	}
	switch idx {
	case 0:
		*(c.ReturnVals[0].(**Response)) = val.(*Response)
	case 1:
		*(c.ReturnVals[1].(*error)) = val.(error)
	}
}

// Trampoline Template
func OtelOnEnterTrampoline_RoundTrip84013(t **Transport, req **Request) (CallContext, bool) {
	defer func() {
		if err := recover(); err != nil {
			println("failed to exec onEnter hook", "httpClientEnterHook")
			if e, ok := err.(error); ok {
				println(e.Error())
			}
			fetchStack, printStack := OtelGetStackImpl, OtelPrintStackImpl
			if fetchStack != nil && printStack != nil {
				printStack(fetchStack())
			}
		}
	}()
	callContext := &CallContextImpl84013{}
	callContext.Params = []interface{}{t, req}
	if HttpClientEnterHookImpl != nil {
		HttpClientEnterHookImpl(callContext, *t, *req)
	}
	return callContext, callContext.SkipCall
}
func OtelOnExitTrampoline_RoundTrip84013(callContext CallContext, retVal0 **Response, retVal1 *error) {
	defer func() {
		if err := recover(); err != nil {
			println("failed to exec onExit hook", "httpClientExitHook")
			if e, ok := err.(error); ok {
				println(e.Error())
			}
			fetchStack, printStack := OtelGetStackImpl, OtelPrintStackImpl
			if fetchStack != nil && printStack != nil {
				printStack(fetchStack())
			}
		}
	}()
	callContext.(*CallContextImpl84013).ReturnVals = []interface{}{retVal0, retVal1}
	if HttpClientExitHookImpl != nil {
		HttpClientExitHookImpl(callContext, *retVal0, *retVal1)
	}
}

var HttpClientEnterHookImpl func(callContext CallContext, t *Transport, req *Request)
var HttpClientExitHookImpl func(callContext CallContext, retVal0 *Response, retVal1 error)
