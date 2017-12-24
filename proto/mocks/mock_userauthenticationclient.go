// Automatically generated by pegomock. DO NOT EDIT!
// Source: github.com/chremoas/auth-srv/proto (interfaces: UserAuthenticationClient)

package authsrv_mocks

import (
	context "context"
	proto "github.com/chremoas/auth-srv/proto"
	client "github.com/micro/go-micro/client"
	pegomock "github.com/petergtz/pegomock"
	"reflect"
)

type MockUserAuthenticationClient struct {
	fail func(message string, callerSkip ...int)
}

func NewMockUserAuthenticationClient() *MockUserAuthenticationClient {
	return &MockUserAuthenticationClient{fail: pegomock.GlobalFailHandler}
}

func (mock *MockUserAuthenticationClient) Confirm(_param0 context.Context, _param1 *proto.AuthConfirmRequest, _param2 ...client.CallOption) (*proto.AuthConfirmResponse, error) {
	params := []pegomock.Param{_param0, _param1}
	for _, param := range _param2 {
		params = append(params, param)
	}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Confirm", params, []reflect.Type{reflect.TypeOf((**proto.AuthConfirmResponse)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *proto.AuthConfirmResponse
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*proto.AuthConfirmResponse)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockUserAuthenticationClient) Create(_param0 context.Context, _param1 *proto.AuthCreateRequest, _param2 ...client.CallOption) (*proto.AuthCreateResponse, error) {
	params := []pegomock.Param{_param0, _param1}
	for _, param := range _param2 {
		params = append(params, param)
	}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Create", params, []reflect.Type{reflect.TypeOf((**proto.AuthCreateResponse)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *proto.AuthCreateResponse
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*proto.AuthCreateResponse)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockUserAuthenticationClient) GetRoles(_param0 context.Context, _param1 *proto.GetRolesRequest, _param2 ...client.CallOption) (*proto.AuthConfirmResponse, error) {
	params := []pegomock.Param{_param0, _param1}
	for _, param := range _param2 {
		params = append(params, param)
	}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetRoles", params, []reflect.Type{reflect.TypeOf((**proto.AuthConfirmResponse)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *proto.AuthConfirmResponse
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*proto.AuthConfirmResponse)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockUserAuthenticationClient) VerifyWasCalledOnce() *VerifierUserAuthenticationClient {
	return &VerifierUserAuthenticationClient{mock, pegomock.Times(1), nil}
}

func (mock *MockUserAuthenticationClient) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierUserAuthenticationClient {
	return &VerifierUserAuthenticationClient{mock, invocationCountMatcher, nil}
}

func (mock *MockUserAuthenticationClient) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierUserAuthenticationClient {
	return &VerifierUserAuthenticationClient{mock, invocationCountMatcher, inOrderContext}
}

type VerifierUserAuthenticationClient struct {
	mock                   *MockUserAuthenticationClient
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
}

func (verifier *VerifierUserAuthenticationClient) Confirm(_param0 context.Context, _param1 *proto.AuthConfirmRequest, _param2 ...client.CallOption) *UserAuthenticationClient_Confirm_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	for _, param := range _param2 {
		params = append(params, param)
	}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Confirm", params)
	return &UserAuthenticationClient_Confirm_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type UserAuthenticationClient_Confirm_OngoingVerification struct {
	mock              *MockUserAuthenticationClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *UserAuthenticationClient_Confirm_OngoingVerification) GetCapturedArguments() (context.Context, *proto.AuthConfirmRequest, []client.CallOption) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *UserAuthenticationClient_Confirm_OngoingVerification) GetAllCapturedArguments() (_param0 []context.Context, _param1 []*proto.AuthConfirmRequest, _param2 [][]client.CallOption) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]context.Context, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(context.Context)
		}
		_param1 = make([]*proto.AuthConfirmRequest, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(*proto.AuthConfirmRequest)
		}
		_param2 = make([][]client.CallOption, len(params[2]))
		for u := range params[0] {
			_param2[u] = make([]client.CallOption, len(params)-2)
			for x := 2; x < len(params); x++ {
				if params[x][u] != nil {
					_param2[u][x-2] = params[x][u].(client.CallOption)
				}
			}
		}
	}
	return
}

func (verifier *VerifierUserAuthenticationClient) Create(_param0 context.Context, _param1 *proto.AuthCreateRequest, _param2 ...client.CallOption) *UserAuthenticationClient_Create_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	for _, param := range _param2 {
		params = append(params, param)
	}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Create", params)
	return &UserAuthenticationClient_Create_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type UserAuthenticationClient_Create_OngoingVerification struct {
	mock              *MockUserAuthenticationClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *UserAuthenticationClient_Create_OngoingVerification) GetCapturedArguments() (context.Context, *proto.AuthCreateRequest, []client.CallOption) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *UserAuthenticationClient_Create_OngoingVerification) GetAllCapturedArguments() (_param0 []context.Context, _param1 []*proto.AuthCreateRequest, _param2 [][]client.CallOption) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]context.Context, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(context.Context)
		}
		_param1 = make([]*proto.AuthCreateRequest, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(*proto.AuthCreateRequest)
		}
		_param2 = make([][]client.CallOption, len(params[2]))
		for u := range params[0] {
			_param2[u] = make([]client.CallOption, len(params)-2)
			for x := 2; x < len(params); x++ {
				if params[x][u] != nil {
					_param2[u][x-2] = params[x][u].(client.CallOption)
				}
			}
		}
	}
	return
}

func (verifier *VerifierUserAuthenticationClient) GetRoles(_param0 context.Context, _param1 *proto.GetRolesRequest, _param2 ...client.CallOption) *UserAuthenticationClient_GetRoles_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	for _, param := range _param2 {
		params = append(params, param)
	}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetRoles", params)
	return &UserAuthenticationClient_GetRoles_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type UserAuthenticationClient_GetRoles_OngoingVerification struct {
	mock              *MockUserAuthenticationClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *UserAuthenticationClient_GetRoles_OngoingVerification) GetCapturedArguments() (context.Context, *proto.GetRolesRequest, []client.CallOption) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *UserAuthenticationClient_GetRoles_OngoingVerification) GetAllCapturedArguments() (_param0 []context.Context, _param1 []*proto.GetRolesRequest, _param2 [][]client.CallOption) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]context.Context, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(context.Context)
		}
		_param1 = make([]*proto.GetRolesRequest, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(*proto.GetRolesRequest)
		}
		_param2 = make([][]client.CallOption, len(params[2]))
		for u := range params[0] {
			_param2[u] = make([]client.CallOption, len(params)-2)
			for x := 2; x < len(params); x++ {
				if params[x][u] != nil {
					_param2[u][x-2] = params[x][u].(client.CallOption)
				}
			}
		}
	}
	return
}
