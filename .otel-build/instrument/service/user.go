package service

import (
	"context"
	"otelbuild-demo/model"
)

type User struct {
	Id       int64
	Nickname string
}

func CreateUser(ctx context.Context, id int64, nickname string) (u *model.UserModel, err error) {
	if  callContext30042, _ := OtelOnEnterTrampoline_CreateUser30042(&ctx, &id, &nickname); false { /* NO_NEWWLINE_PLACEHOLDER */ ;	} else {  		defer OtelOnExitTrampoline_CreateUser30042(callContext30042, &u, &err) ;	}  
	user := &model.UserModel{
		Id:       id,
		Nickname: nickname,
	}
	return user, nil
}

// Seeing is not always believing. The following template is a bit tricky, see
// trampoline.go for more details
// Struct Template
type CallContextImpl30042 struct {
	Params     []interface{}
	ReturnVals []interface{}
	SkipCall   bool
	Data       interface{}
}

func (c *CallContextImpl30042) SetSkipCall(skip bool)    { c.SkipCall = skip }
func (c *CallContextImpl30042) IsSkipCall() bool         { return c.SkipCall }
func (c *CallContextImpl30042) SetData(data interface{}) { c.Data = data }
func (c *CallContextImpl30042) GetData() interface{}     { return c.Data }
func (c *CallContextImpl30042) GetKeyData(key string) interface{} {
	if c.Data == nil {
		return nil
	}
	return c.Data.(map[string]interface{})[key]
}
func (c *CallContextImpl30042) SetKeyData(key string, val interface{}) {
	if c.Data == nil {
		c.Data = make(map[string]interface{})
	}
	c.Data.(map[string]interface{})[key] = val
}
func (c *CallContextImpl30042) HasKeyData(key string) bool {
	if c.Data == nil {
		return false
	}
	_, ok := c.Data.(map[string]interface{})[key]
	return ok
}
func (c *CallContextImpl30042) GetParam(idx int) interface{} {
	switch idx {
	case 0:
		return *(c.Params[0].(*context.Context))
	case 1:
		return *(c.Params[1].(*int64))
	case 2:
		return *(c.Params[2].(*string))
	}
	return nil
}
func (c *CallContextImpl30042) SetParam(idx int, val interface{}) {
	if val == nil {
		c.Params[idx] = nil
		return
	}
	switch idx {
	case 0:
		*(c.Params[0].(*context.Context)) = val.(context.Context)
	case 1:
		*(c.Params[1].(*int64)) = val.(int64)
	case 2:
		*(c.Params[2].(*string)) = val.(string)
	}
}
func (c *CallContextImpl30042) GetReturnVal(idx int) interface{} {
	switch idx {
	case 0:
		return *(c.ReturnVals[0].(**model.UserModel))
	case 1:
		return *(c.ReturnVals[1].(*error))
	}
	return nil
}
func (c *CallContextImpl30042) SetReturnVal(idx int, val interface{}) {
	if val == nil {
		c.ReturnVals[idx] = nil
		return
	}
	switch idx {
	case 0:
		*(c.ReturnVals[0].(**model.UserModel)) = val.(*model.UserModel)
	case 1:
		*(c.ReturnVals[1].(*error)) = val.(error)
	}
}

// Trampoline Template
func OtelOnEnterTrampoline_CreateUser30042(ctx *context.Context, id *int64, nickname *string) (CallContext, bool) {
	defer func() {
		if err := recover(); err != nil {
			println("failed to exec onEnter hook", "createUserEnterHook")
			if e, ok := err.(error); ok {
				println(e.Error())
			}
			fetchStack, printStack := OtelGetStackImpl, OtelPrintStackImpl
			if fetchStack != nil && printStack != nil {
				printStack(fetchStack())
			}
		}
	}()
	callContext := &CallContextImpl30042{}
	callContext.Params = []interface{}{ctx, id, nickname}
	if CreateUserEnterHookImpl != nil {
		CreateUserEnterHookImpl(callContext, *ctx, *id, *nickname)
	}
	return callContext, callContext.SkipCall
}
func OtelOnExitTrampoline_CreateUser30042(callContext CallContext, u **model.UserModel, err *error) {
	defer func() {
		if err := recover(); err != nil {
			println("failed to exec onExit hook", "createUserExitHook")
			if e, ok := err.(error); ok {
				println(e.Error())
			}
			fetchStack, printStack := OtelGetStackImpl, OtelPrintStackImpl
			if fetchStack != nil && printStack != nil {
				printStack(fetchStack())
			}
		}
	}()
	callContext.(*CallContextImpl30042).ReturnVals = []interface{}{u, err}
	if CreateUserExitHookImpl != nil {
		CreateUserExitHookImpl(callContext, *u, *err)
	}
}

var CreateUserEnterHookImpl func(callContext CallContext, ctx context.Context, id int64, nickname string)
var CreateUserExitHookImpl func(callContext CallContext, u *model.UserModel, err error)
