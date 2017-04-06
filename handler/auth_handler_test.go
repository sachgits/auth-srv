package handler

import (
	"errors"
	"github.com/abaeve/auth-srv/model"
	proto "github.com/abaeve/auth-srv/proto"
	"github.com/abaeve/auth-srv/repository"
	"github.com/golang/mock/gomock"
	"golang.org/x/net/context"
	"testing"
)

//<editor-fold desc="Generated Mock Code">
/* Automatically generated by MockGen. DO NOT EDIT!
 * Source: github.com/abaeve/auth-srv/repository (interfaces: AllianceRepository,CorporationRepository,CharacterRepository,UserRepository,RoleRepository,AuthenticationCodeRepository)
 */

type MockAccessesRepository struct {
	ctrl     *gomock.Controller
	recorder *_MockAccessesRepositoryRecorder
}

type _MockAccessesRepositoryRecorder struct {
	mock *MockAccessesRepository
}

func NewMockAccessesRepository(ctrl *gomock.Controller) *MockAccessesRepository {
	mock := &MockAccessesRepository{ctrl: ctrl}
	mock.recorder = &_MockAccessesRepositoryRecorder{mock}
	return mock
}

func (_m *MockAccessesRepository) EXPECT() *_MockAccessesRepositoryRecorder {
	return _m.recorder
}

func (_m *MockAccessesRepository) FindByChatId(_param0 string) ([]string, error) {
	ret := _m.ctrl.Call(_m, "FindByChatId", _param0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAccessesRepositoryRecorder) FindByChatId(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FindByChatId", arg0)
}

func (_m *MockAccessesRepository) SaveAllianceAndCorpRole(_param0 int64, _param1 int64, _param2 model.Role) error {
	ret := _m.ctrl.Call(_m, "SaveAllianceAndCorpRole", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAccessesRepositoryRecorder) SaveAllianceAndCorpRole(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveAllianceAndCorpRole", arg0, arg1, arg2)
}

func (_m *MockAccessesRepository) SaveAllianceCharacterLeadershipRole(_param0 int64, _param1 model.Role) error {
	ret := _m.ctrl.Call(_m, "SaveAllianceCharacterLeadershipRole", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAccessesRepositoryRecorder) SaveAllianceCharacterLeadershipRole(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveAllianceCharacterLeadershipRole", arg0, arg1)
}

func (_m *MockAccessesRepository) SaveAllianceRole(_param0 int64, _param1 model.Role) error {
	ret := _m.ctrl.Call(_m, "SaveAllianceRole", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAccessesRepositoryRecorder) SaveAllianceRole(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveAllianceRole", arg0, arg1)
}

func (_m *MockAccessesRepository) SaveCharacterRole(_param0 int64, _param1 model.Role) error {
	ret := _m.ctrl.Call(_m, "SaveCharacterRole", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAccessesRepositoryRecorder) SaveCharacterRole(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveCharacterRole", arg0, arg1)
}

func (_m *MockAccessesRepository) SaveCorporationCharacterLeadershipRole(_param0 int64, _param1 model.Role) error {
	ret := _m.ctrl.Call(_m, "SaveCorporationCharacterLeadershipRole", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAccessesRepositoryRecorder) SaveCorporationCharacterLeadershipRole(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveCorporationCharacterLeadershipRole", arg0, arg1)
}

func (_m *MockAccessesRepository) SaveCorporationRole(_param0 int64, _param1 model.Role) error {
	ret := _m.ctrl.Call(_m, "SaveCorporationRole", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAccessesRepositoryRecorder) SaveCorporationRole(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveCorporationRole", arg0, arg1)
}

type MockAllianceRepository struct {
	ctrl     *gomock.Controller
	recorder *_MockAllianceRepositoryRecorder
}

type _MockAllianceRepositoryRecorder struct {
	mock *MockAllianceRepository
}

func NewMockAllianceRepository(ctrl *gomock.Controller) *MockAllianceRepository {
	mock := &MockAllianceRepository{ctrl: ctrl}
	mock.recorder = &_MockAllianceRepositoryRecorder{mock}
	return mock
}

func (_m *MockAllianceRepository) EXPECT() *_MockAllianceRepositoryRecorder {
	return _m.recorder
}

func (_m *MockAllianceRepository) FindByAllianceId(_param0 int64) *model.Alliance {
	ret := _m.ctrl.Call(_m, "FindByAllianceId", _param0)
	ret0, _ := ret[0].(*model.Alliance)
	return ret0
}

func (_mr *_MockAllianceRepositoryRecorder) FindByAllianceId(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FindByAllianceId", arg0)
}

func (_m *MockAllianceRepository) Save(_param0 *model.Alliance) error {
	ret := _m.ctrl.Call(_m, "Save", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAllianceRepositoryRecorder) Save(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Save", arg0)
}

type MockCorporationRepository struct {
	ctrl     *gomock.Controller
	recorder *_MockCorporationRepositoryRecorder
}

type _MockCorporationRepositoryRecorder struct {
	mock *MockCorporationRepository
}

func NewMockCorporationRepository(ctrl *gomock.Controller) *MockCorporationRepository {
	mock := &MockCorporationRepository{ctrl: ctrl}
	mock.recorder = &_MockCorporationRepositoryRecorder{mock}
	return mock
}

func (_m *MockCorporationRepository) EXPECT() *_MockCorporationRepositoryRecorder {
	return _m.recorder
}

func (_m *MockCorporationRepository) FindByCorporationId(_param0 int64) *model.Corporation {
	ret := _m.ctrl.Call(_m, "FindByCorporationId", _param0)
	ret0, _ := ret[0].(*model.Corporation)
	return ret0
}

func (_mr *_MockCorporationRepositoryRecorder) FindByCorporationId(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FindByCorporationId", arg0)
}

func (_m *MockCorporationRepository) Save(_param0 *model.Corporation) error {
	ret := _m.ctrl.Call(_m, "Save", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCorporationRepositoryRecorder) Save(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Save", arg0)
}

type MockCharacterRepository struct {
	ctrl     *gomock.Controller
	recorder *_MockCharacterRepositoryRecorder
}

type _MockCharacterRepositoryRecorder struct {
	mock *MockCharacterRepository
}

func NewMockCharacterRepository(ctrl *gomock.Controller) *MockCharacterRepository {
	mock := &MockCharacterRepository{ctrl: ctrl}
	mock.recorder = &_MockCharacterRepositoryRecorder{mock}
	return mock
}

func (_m *MockCharacterRepository) EXPECT() *_MockCharacterRepositoryRecorder {
	return _m.recorder
}

func (_m *MockCharacterRepository) FindByAutenticationCode(_param0 string) *model.Character {
	ret := _m.ctrl.Call(_m, "FindByAutenticationCode", _param0)
	ret0, _ := ret[0].(*model.Character)
	return ret0
}

func (_mr *_MockCharacterRepositoryRecorder) FindByAutenticationCode(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FindByAutenticationCode", arg0)
}

func (_m *MockCharacterRepository) FindByCharacterId(_param0 int64) *model.Character {
	ret := _m.ctrl.Call(_m, "FindByCharacterId", _param0)
	ret0, _ := ret[0].(*model.Character)
	return ret0
}

func (_mr *_MockCharacterRepositoryRecorder) FindByCharacterId(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FindByCharacterId", arg0)
}

func (_m *MockCharacterRepository) Save(_param0 *model.Character) error {
	ret := _m.ctrl.Call(_m, "Save", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCharacterRepositoryRecorder) Save(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Save", arg0)
}

type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *_MockUserRepositoryRecorder
}

type _MockUserRepositoryRecorder struct {
	mock *MockUserRepository
}

func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &_MockUserRepositoryRecorder{mock}
	return mock
}

func (_m *MockUserRepository) EXPECT() *_MockUserRepositoryRecorder {
	return _m.recorder
}

func (_m *MockUserRepository) FindByChatId(_param0 string) *model.User {
	ret := _m.ctrl.Call(_m, "FindByChatId", _param0)
	ret0, _ := ret[0].(*model.User)
	return ret0
}

func (_mr *_MockUserRepositoryRecorder) FindByChatId(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FindByChatId", arg0)
}

func (_m *MockUserRepository) LinkCharacterToUserByAuthCode(_param0 string, _param1 *model.User) error {
	ret := _m.ctrl.Call(_m, "LinkCharacterToUserByAuthCode", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockUserRepositoryRecorder) LinkCharacterToUserByAuthCode(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LinkCharacterToUserByAuthCode", arg0, arg1)
}

func (_m *MockUserRepository) Save(_param0 *model.User) error {
	ret := _m.ctrl.Call(_m, "Save", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockUserRepositoryRecorder) Save(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Save", arg0)
}

type MockRoleRepository struct {
	ctrl     *gomock.Controller
	recorder *_MockRoleRepositoryRecorder
}

type _MockRoleRepositoryRecorder struct {
	mock *MockRoleRepository
}

func NewMockRoleRepository(ctrl *gomock.Controller) *MockRoleRepository {
	mock := &MockRoleRepository{ctrl: ctrl}
	mock.recorder = &_MockRoleRepositoryRecorder{mock}
	return mock
}

func (_m *MockRoleRepository) EXPECT() *_MockRoleRepositoryRecorder {
	return _m.recorder
}

func (_m *MockRoleRepository) Save(_param0 *model.Role) error {
	ret := _m.ctrl.Call(_m, "Save", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRoleRepositoryRecorder) Save(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Save", arg0)
}

type MockAuthenticationCodeRepository struct {
	ctrl     *gomock.Controller
	recorder *_MockAuthenticationCodeRepositoryRecorder
}

type _MockAuthenticationCodeRepositoryRecorder struct {
	mock *MockAuthenticationCodeRepository
}

func NewMockAuthenticationCodeRepository(ctrl *gomock.Controller) *MockAuthenticationCodeRepository {
	mock := &MockAuthenticationCodeRepository{ctrl: ctrl}
	mock.recorder = &_MockAuthenticationCodeRepositoryRecorder{mock}
	return mock
}

func (_m *MockAuthenticationCodeRepository) EXPECT() *_MockAuthenticationCodeRepositoryRecorder {
	return _m.recorder
}

func (_m *MockAuthenticationCodeRepository) Save(_param0 *model.Character, _param1 string) error {
	ret := _m.ctrl.Call(_m, "Save", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAuthenticationCodeRepositoryRecorder) Save(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Save", arg0, arg1)
}

//</editor-fold>

type testError struct {
	message string
}

func (te *testError) Error() string {
	return te.message
}

func SharedSetup(t *testing.T) (*gomock.Controller,
	*MockAuthenticationCodeRepository,
	*MockUserRepository,
	*MockCharacterRepository,
	*MockCorporationRepository,
	*MockAllianceRepository,
	*MockAccessesRepository) {
	mockCtrl := gomock.NewController(t)

	repository.AuthenticationCodeRepo = NewMockAuthenticationCodeRepository(mockCtrl)
	repository.UserRepo = NewMockUserRepository(mockCtrl)
	repository.CharacterRepo = NewMockCharacterRepository(mockCtrl)
	repository.CorporationRepo = NewMockCorporationRepository(mockCtrl)
	repository.AllianceRepo = NewMockAllianceRepository(mockCtrl)
	repository.AccessRepo = NewMockAccessesRepository(mockCtrl)

	mockAuthRepo := repository.AuthenticationCodeRepo.(*MockAuthenticationCodeRepository)
	mockUserRepo := repository.UserRepo.(*MockUserRepository)
	mockCharRepo := repository.CharacterRepo.(*MockCharacterRepository)
	mockCorpRepo := repository.CorporationRepo.(*MockCorporationRepository)
	mockAlliRepo := repository.AllianceRepo.(*MockAllianceRepository)
	mockAcceRepo := repository.AccessRepo.(*MockAccessesRepository)

	return mockCtrl, mockAuthRepo, mockUserRepo, mockCharRepo, mockCorpRepo, mockAlliRepo, mockAcceRepo
}

func TestCreateEmptyDb(t *testing.T) {
	mockCtrl, mockAuthRepo, _, mockCharRepo, mockCorpRepo, mockAlliRepo, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	authCreateRequest := proto.AuthCreateRequest{
		Token:       "mytoken",
		Alliance:    &proto.Alliance{Name: "Test Alliance", Id: 1, Ticker: "TSTA"},
		Corporation: &proto.Corporation{Name: "Test Corp", Id: 1, Ticker: "TSTC"},
		Character:   &proto.Character{Name: "Test Character", Id: 1},
		AuthScope:   []string{"scope1", "scope2"},
	}
	var authCreateResponse proto.AuthCreateResponse
	var ctx context.Context

	authHandler := AuthHandler{}

	//Set our expectations
	gomock.InOrder(
		mockAlliRepo.EXPECT().FindByAllianceId(authCreateRequest.Alliance.Id).Return(nil),
		mockAlliRepo.EXPECT().Save(
			&model.Alliance{
				AllianceId:     authCreateRequest.Alliance.Id,
				AllianceName:   authCreateRequest.Alliance.Name,
				AllianceTicker: authCreateRequest.Alliance.Ticker,
			},
		).Return(nil),

		mockCorpRepo.EXPECT().FindByCorporationId(authCreateRequest.Corporation.Id).Return(nil),
		mockCorpRepo.EXPECT().Save(
			&model.Corporation{
				CorporationId:     authCreateRequest.Corporation.Id,
				CorporationName:   authCreateRequest.Corporation.Name,
				CorporationTicker: authCreateRequest.Corporation.Ticker,
				AllianceId:        &authCreateRequest.Alliance.Id,
				Alliance: model.Alliance{
					AllianceId:     authCreateRequest.Alliance.Id,
					AllianceName:   authCreateRequest.Alliance.Name,
					AllianceTicker: authCreateRequest.Alliance.Ticker,
				},
			},
		).Return(nil),

		mockCharRepo.EXPECT().FindByCharacterId(authCreateRequest.Character.Id).Return(nil),
		mockCharRepo.EXPECT().Save(
			&model.Character{
				CharacterId:   authCreateRequest.Character.Id,
				CharacterName: authCreateRequest.Character.Name,
				Token:         authCreateRequest.Token,
				CorporationId: authCreateRequest.Corporation.Id,
				Corporation: model.Corporation{
					CorporationId:     authCreateRequest.Corporation.Id,
					CorporationName:   authCreateRequest.Corporation.Name,
					CorporationTicker: authCreateRequest.Corporation.Ticker,
					AllianceId:        &authCreateRequest.Alliance.Id,
					Alliance: model.Alliance{
						AllianceId:     authCreateRequest.Alliance.Id,
						AllianceName:   authCreateRequest.Alliance.Name,
						AllianceTicker: authCreateRequest.Alliance.Ticker,
					},
				},
			},
		).Return(nil),

		mockAuthRepo.EXPECT().Save(
			&model.Character{
				CharacterId:   authCreateRequest.Character.Id,
				CharacterName: authCreateRequest.Character.Name,
				Token:         authCreateRequest.Token,
				CorporationId: authCreateRequest.Corporation.Id,
				Corporation: model.Corporation{
					CorporationId:     authCreateRequest.Corporation.Id,
					CorporationName:   authCreateRequest.Corporation.Name,
					CorporationTicker: authCreateRequest.Corporation.Ticker,
					AllianceId:        &authCreateRequest.Alliance.Id,
					Alliance: model.Alliance{
						AllianceId:     authCreateRequest.Alliance.Id,
						AllianceName:   authCreateRequest.Alliance.Name,
						AllianceTicker: authCreateRequest.Alliance.Ticker,
					},
				},
			},
			gomock.Any(),
		).Return(nil),
	)

	err := authHandler.Create(ctx, &authCreateRequest, &authCreateResponse)

	if err != nil {
		t.Fatalf("Received an error on Create call, expected nothing: %s", err)
	}

	if authCreateResponse.AuthenticationCode == "" {
		t.Fatal("Expected at least something as an authentication code, got nothing")
	}
}

func TestCreateNoAllianceCorporation(t *testing.T) {
	mockCtrl, mockAuthRepo, _, mockCharRepo, mockCorpRepo, mockAlliRepo, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	authCreateRequest := proto.AuthCreateRequest{
		Token:       "mytoken",
		Alliance:    &proto.Alliance{Name: "Test Alliance", Id: 1, Ticker: "TSTA"},
		Corporation: &proto.Corporation{Name: "Test Corp", Id: 1, Ticker: "TSTC"},
		Character:   &proto.Character{Name: "Test Character", Id: 1},
		AuthScope:   []string{"scope1", "scope2"},
	}
	var authCreateResponse proto.AuthCreateResponse
	var ctx context.Context

	authHandler := AuthHandler{}

	//Set our expectations
	gomock.InOrder(
		mockAlliRepo.EXPECT().FindByAllianceId(authCreateRequest.Alliance.Id).Return(nil),
		mockAlliRepo.EXPECT().Save(
			&model.Alliance{
				AllianceId:     authCreateRequest.Alliance.Id,
				AllianceName:   authCreateRequest.Alliance.Name,
				AllianceTicker: authCreateRequest.Alliance.Ticker,
			},
		).Return(nil),

		mockCorpRepo.EXPECT().FindByCorporationId(authCreateRequest.Corporation.Id).Return(nil),
		mockCorpRepo.EXPECT().Save(
			&model.Corporation{
				CorporationId:     authCreateRequest.Corporation.Id,
				CorporationName:   authCreateRequest.Corporation.Name,
				CorporationTicker: authCreateRequest.Corporation.Ticker,
				AllianceId:        &authCreateRequest.Alliance.Id,
				Alliance: model.Alliance{
					AllianceId:     authCreateRequest.Alliance.Id,
					AllianceName:   authCreateRequest.Alliance.Name,
					AllianceTicker: authCreateRequest.Alliance.Ticker,
				},
			},
		).Return(nil),

		mockCharRepo.EXPECT().FindByCharacterId(authCreateRequest.Character.Id).Return(nil),
		mockCharRepo.EXPECT().Save(
			&model.Character{
				CharacterId:   authCreateRequest.Character.Id,
				CharacterName: authCreateRequest.Character.Name,
				Token:         authCreateRequest.Token,
				CorporationId: authCreateRequest.Corporation.Id,
				Corporation: model.Corporation{
					CorporationId:     authCreateRequest.Corporation.Id,
					CorporationName:   authCreateRequest.Corporation.Name,
					CorporationTicker: authCreateRequest.Corporation.Ticker,
					AllianceId:        &authCreateRequest.Alliance.Id,
					Alliance: model.Alliance{
						AllianceId:     authCreateRequest.Alliance.Id,
						AllianceName:   authCreateRequest.Alliance.Name,
						AllianceTicker: authCreateRequest.Alliance.Ticker,
					},
				},
			},
		).Return(nil),

		mockAuthRepo.EXPECT().Save(
			&model.Character{
				CharacterId:   authCreateRequest.Character.Id,
				CharacterName: authCreateRequest.Character.Name,
				Token:         authCreateRequest.Token,
				CorporationId: authCreateRequest.Corporation.Id,
				Corporation: model.Corporation{
					CorporationId:     authCreateRequest.Corporation.Id,
					CorporationName:   authCreateRequest.Corporation.Name,
					CorporationTicker: authCreateRequest.Corporation.Ticker,
					AllianceId:        &authCreateRequest.Alliance.Id,
					Alliance: model.Alliance{
						AllianceId:     authCreateRequest.Alliance.Id,
						AllianceName:   authCreateRequest.Alliance.Name,
						AllianceTicker: authCreateRequest.Alliance.Ticker,
					},
				},
			},
			gomock.Any(),
		).Return(nil),
	)

	err := authHandler.Create(ctx, &authCreateRequest, &authCreateResponse)

	if err != nil {
		t.Fatalf("Received an error on Create call, expected nothing: %s", err)
	}

	if authCreateResponse.AuthenticationCode == "" {
		t.Fatal("Expected at least something as an authentication code, got nothing")
	}
}

func TestAllianceExistsNoCorpOrChar(t *testing.T) {
	mockCtrl, mockAuthRepo, _, mockCharRepo, mockCorpRepo, mockAlliRepo, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	authCreateRequest := proto.AuthCreateRequest{
		Token:       "mytoken",
		Alliance:    &proto.Alliance{Name: "Test Alliance", Id: 1, Ticker: "TSTA"},
		Corporation: &proto.Corporation{Name: "Test Corp", Id: 1, Ticker: "TSTC"},
		Character:   &proto.Character{Name: "Test Character", Id: 1},
		AuthScope:   []string{"scope1", "scope2"},
	}
	var authCreateResponse proto.AuthCreateResponse
	var ctx context.Context

	authHandler := AuthHandler{}

	//Set our expectations
	gomock.InOrder(
		mockAlliRepo.EXPECT().FindByAllianceId(authCreateRequest.Alliance.Id).Return(&model.Alliance{
			AllianceId:     authCreateRequest.Alliance.Id,
			AllianceName:   authCreateRequest.Alliance.Name,
			AllianceTicker: authCreateRequest.Alliance.Ticker,
		}),
		mockAlliRepo.EXPECT().Save(&model.Alliance{}).Times(0),

		mockCorpRepo.EXPECT().FindByCorporationId(authCreateRequest.Corporation.Id).Return(nil),
		mockCorpRepo.EXPECT().Save(
			&model.Corporation{
				CorporationId:     authCreateRequest.Corporation.Id,
				CorporationName:   authCreateRequest.Corporation.Name,
				CorporationTicker: authCreateRequest.Corporation.Ticker,
				AllianceId:        &authCreateRequest.Alliance.Id,
				Alliance: model.Alliance{
					AllianceId:     authCreateRequest.Alliance.Id,
					AllianceName:   authCreateRequest.Alliance.Name,
					AllianceTicker: authCreateRequest.Alliance.Ticker,
				},
			},
		).Return(nil),

		mockCharRepo.EXPECT().FindByCharacterId(authCreateRequest.Character.Id).Return(nil),
		mockCharRepo.EXPECT().Save(
			&model.Character{
				CharacterId:   authCreateRequest.Character.Id,
				CharacterName: authCreateRequest.Character.Name,
				Token:         authCreateRequest.Token,
				CorporationId: authCreateRequest.Corporation.Id,
				Corporation: model.Corporation{
					CorporationId:     authCreateRequest.Corporation.Id,
					CorporationName:   authCreateRequest.Corporation.Name,
					CorporationTicker: authCreateRequest.Corporation.Ticker,
					AllianceId:        &authCreateRequest.Alliance.Id,
					Alliance: model.Alliance{
						AllianceId:     authCreateRequest.Alliance.Id,
						AllianceName:   authCreateRequest.Alliance.Name,
						AllianceTicker: authCreateRequest.Alliance.Ticker,
					},
				},
			},
		).Return(nil),

		mockAuthRepo.EXPECT().Save(
			&model.Character{
				CharacterId:   authCreateRequest.Character.Id,
				CharacterName: authCreateRequest.Character.Name,
				Token:         authCreateRequest.Token,
				CorporationId: authCreateRequest.Corporation.Id,
				Corporation: model.Corporation{
					CorporationId:     authCreateRequest.Corporation.Id,
					CorporationName:   authCreateRequest.Corporation.Name,
					CorporationTicker: authCreateRequest.Corporation.Ticker,
					AllianceId:        &authCreateRequest.Alliance.Id,
					Alliance: model.Alliance{
						AllianceId:     authCreateRequest.Alliance.Id,
						AllianceName:   authCreateRequest.Alliance.Name,
						AllianceTicker: authCreateRequest.Alliance.Ticker,
					},
				},
			},
			gomock.Any(),
		).Return(nil),
	)

	err := authHandler.Create(ctx, &authCreateRequest, &authCreateResponse)

	if err != nil {
		t.Fatalf("Received an error on Create call, expected nothing: %s", err)
	}

	if authCreateResponse.AuthenticationCode == "" {
		t.Fatal("Expected at least something as an authentication code, got nothing")
	}
}

func TestAllianceAndCorpExistButNoChar(t *testing.T) {
	mockCtrl, mockAuthRepo, _, mockCharRepo, mockCorpRepo, mockAlliRepo, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	authCreateRequest := proto.AuthCreateRequest{
		Token:       "mytoken",
		Alliance:    &proto.Alliance{Name: "Test Alliance", Id: 1, Ticker: "TSTA"},
		Corporation: &proto.Corporation{Name: "Test Corp", Id: 1, Ticker: "TSTC"},
		Character:   &proto.Character{Name: "Test Character", Id: 1},
		AuthScope:   []string{"scope1", "scope2"},
	}
	var authCreateResponse proto.AuthCreateResponse
	var ctx context.Context

	authHandler := AuthHandler{}

	//Set our expectations
	gomock.InOrder(
		mockAlliRepo.EXPECT().FindByAllianceId(authCreateRequest.Alliance.Id).Return(&model.Alliance{
			AllianceId:     authCreateRequest.Alliance.Id,
			AllianceName:   authCreateRequest.Alliance.Name,
			AllianceTicker: authCreateRequest.Alliance.Ticker,
		}),
		mockAlliRepo.EXPECT().Save(&model.Alliance{}).Times(0),

		mockCorpRepo.EXPECT().FindByCorporationId(authCreateRequest.Corporation.Id).Return(&model.Corporation{
			CorporationId:     authCreateRequest.Corporation.Id,
			CorporationName:   authCreateRequest.Corporation.Name,
			CorporationTicker: authCreateRequest.Corporation.Ticker,
			AllianceId:        &authCreateRequest.Alliance.Id,
			Alliance: model.Alliance{
				AllianceId:     authCreateRequest.Alliance.Id,
				AllianceName:   authCreateRequest.Alliance.Name,
				AllianceTicker: authCreateRequest.Alliance.Ticker,
			},
		}),
		mockCorpRepo.EXPECT().Save(&model.Corporation{}).Times(0),

		mockCharRepo.EXPECT().FindByCharacterId(authCreateRequest.Character.Id).Return(nil),
		mockCharRepo.EXPECT().Save(
			&model.Character{
				CharacterId:   authCreateRequest.Character.Id,
				CharacterName: authCreateRequest.Character.Name,
				Token:         authCreateRequest.Token,
				CorporationId: authCreateRequest.Corporation.Id,
				Corporation: model.Corporation{
					CorporationId:     authCreateRequest.Corporation.Id,
					CorporationName:   authCreateRequest.Corporation.Name,
					CorporationTicker: authCreateRequest.Corporation.Ticker,
					AllianceId:        &authCreateRequest.Alliance.Id,
					Alliance: model.Alliance{
						AllianceId:     authCreateRequest.Alliance.Id,
						AllianceName:   authCreateRequest.Alliance.Name,
						AllianceTicker: authCreateRequest.Alliance.Ticker,
					},
				},
			},
		).Return(nil),

		mockAuthRepo.EXPECT().Save(
			&model.Character{
				CharacterId:   authCreateRequest.Character.Id,
				CharacterName: authCreateRequest.Character.Name,
				Token:         authCreateRequest.Token,
				CorporationId: authCreateRequest.Corporation.Id,
				Corporation: model.Corporation{
					CorporationId:     authCreateRequest.Corporation.Id,
					CorporationName:   authCreateRequest.Corporation.Name,
					CorporationTicker: authCreateRequest.Corporation.Ticker,
					AllianceId:        &authCreateRequest.Alliance.Id,
					Alliance: model.Alliance{
						AllianceId:     authCreateRequest.Alliance.Id,
						AllianceName:   authCreateRequest.Alliance.Name,
						AllianceTicker: authCreateRequest.Alliance.Ticker,
					},
				},
			},
			gomock.Any(),
		).Return(nil),
	)

	err := authHandler.Create(ctx, &authCreateRequest, &authCreateResponse)

	if err != nil {
		t.Fatalf("Received an error on Create call, expected nothing: %s", err)
	}

	if authCreateResponse.AuthenticationCode == "" {
		t.Fatal("Expected at least something as an authentication code, got nothing")
	}
}

func TestAllianceAndCorpAndCharExist(t *testing.T) {
	mockCtrl, mockAuthRepo, _, mockCharRepo, mockCorpRepo, mockAlliRepo, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	authCreateRequest := proto.AuthCreateRequest{
		Token:       "mytoken",
		Alliance:    &proto.Alliance{Name: "Test Alliance", Id: 1, Ticker: "TSTA"},
		Corporation: &proto.Corporation{Name: "Test Corp", Id: 1, Ticker: "TSTC"},
		Character:   &proto.Character{Name: "Test Character", Id: 1},
		AuthScope:   []string{"scope1", "scope2"},
	}
	var authCreateResponse proto.AuthCreateResponse
	var ctx context.Context

	authHandler := AuthHandler{}

	//Set our expectations
	gomock.InOrder(
		mockAlliRepo.EXPECT().FindByAllianceId(authCreateRequest.Alliance.Id).Return(&model.Alliance{
			AllianceId:     authCreateRequest.Alliance.Id,
			AllianceName:   authCreateRequest.Alliance.Name,
			AllianceTicker: authCreateRequest.Alliance.Ticker,
		}),
		mockAlliRepo.EXPECT().Save(&model.Alliance{}).Times(0),

		mockCorpRepo.EXPECT().FindByCorporationId(authCreateRequest.Corporation.Id).Return(&model.Corporation{
			CorporationId:     authCreateRequest.Corporation.Id,
			CorporationName:   authCreateRequest.Corporation.Name,
			CorporationTicker: authCreateRequest.Corporation.Ticker,
		}),
		mockCorpRepo.EXPECT().Save(&model.Corporation{}).Times(0),

		mockCharRepo.EXPECT().FindByCharacterId(authCreateRequest.Character.Id).Return(&model.Character{
			CharacterId:   authCreateRequest.Character.Id,
			CharacterName: authCreateRequest.Character.Name,
			Token:         authCreateRequest.Token,
			CorporationId: authCreateRequest.Corporation.Id,
			Corporation: model.Corporation{
				CorporationId:     authCreateRequest.Corporation.Id,
				CorporationName:   authCreateRequest.Corporation.Name,
				CorporationTicker: authCreateRequest.Corporation.Ticker,
				AllianceId:        &authCreateRequest.Alliance.Id,
				Alliance: model.Alliance{
					AllianceId:     authCreateRequest.Alliance.Id,
					AllianceName:   authCreateRequest.Alliance.Name,
					AllianceTicker: authCreateRequest.Alliance.Ticker,
				},
			},
		}),
		mockCharRepo.EXPECT().Save(&model.Character{}).Times(0),

		mockAuthRepo.EXPECT().Save(
			&model.Character{
				CharacterId:   authCreateRequest.Character.Id,
				CharacterName: authCreateRequest.Character.Name,
				Token:         authCreateRequest.Token,
				CorporationId: authCreateRequest.Corporation.Id,
				Corporation: model.Corporation{
					CorporationId:     authCreateRequest.Corporation.Id,
					CorporationName:   authCreateRequest.Corporation.Name,
					CorporationTicker: authCreateRequest.Corporation.Ticker,
					AllianceId:        &authCreateRequest.Alliance.Id,
					Alliance: model.Alliance{
						AllianceId:     authCreateRequest.Alliance.Id,
						AllianceName:   authCreateRequest.Alliance.Name,
						AllianceTicker: authCreateRequest.Alliance.Ticker,
					},
				},
			},
			gomock.Any(),
		).Return(nil),
	)

	err := authHandler.Create(ctx, &authCreateRequest, &authCreateResponse)

	if err != nil {
		t.Fatalf("Received an error on Create call, expected nothing: %s", err)
	}

	if authCreateResponse.AuthenticationCode == "" {
		t.Fatal("Expected at least something as an authentication code, got nothing")
	}
}

func TestAllianceErrorCondition(t *testing.T) {
	mockCtrl, _, _, _, _, mockAlliRepo, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	authCreateRequest := proto.AuthCreateRequest{
		Token:       "mytoken",
		Alliance:    &proto.Alliance{Name: "Test Alliance", Id: 1, Ticker: "TSTA"},
		Corporation: &proto.Corporation{Name: "Test Corp", Id: 1, Ticker: "TSTC"},
		Character:   &proto.Character{Name: "Test Character", Id: 1},
		AuthScope:   []string{"scope1", "scope2"},
	}
	var authCreateResponse proto.AuthCreateResponse
	var ctx context.Context

	authHandler := AuthHandler{}

	//Set our expectations
	gomock.InOrder(
		mockAlliRepo.EXPECT().FindByAllianceId(authCreateRequest.Alliance.Id).Return(nil),
		mockAlliRepo.EXPECT().Save(gomock.Any()).Return(&testError{message: "Don't do that alliance"}),
	)

	err := authHandler.Create(ctx, &authCreateRequest, &authCreateResponse)

	if err == nil && err.Error() == "Don't do that alliance" {
		t.Fatal("Expected an error but got nothing")
	}
}

func TestCorporationErrorCondition(t *testing.T) {
	mockCtrl, _, _, _, mockCorpRepo, mockAlliRepo, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	authCreateRequest := proto.AuthCreateRequest{
		Token:       "mytoken",
		Alliance:    &proto.Alliance{Name: "Test Alliance", Id: 1, Ticker: "TSTA"},
		Corporation: &proto.Corporation{Name: "Test Corp", Id: 1, Ticker: "TSTC"},
		Character:   &proto.Character{Name: "Test Character", Id: 1},
		AuthScope:   []string{"scope1", "scope2"},
	}
	var authCreateResponse proto.AuthCreateResponse
	var ctx context.Context

	authHandler := AuthHandler{}

	//Set our expectations
	gomock.InOrder(
		mockAlliRepo.EXPECT().FindByAllianceId(gomock.Any()).Return(nil),
		mockAlliRepo.EXPECT().Save(gomock.Any()).Times(1),

		mockCorpRepo.EXPECT().FindByCorporationId(gomock.Any()).Return(nil),
		mockCorpRepo.EXPECT().Save(gomock.Any()).Return(&testError{message: "Don't do that corp"}),
	)

	err := authHandler.Create(ctx, &authCreateRequest, &authCreateResponse)

	if err == nil && err.Error() == "Don't do that corp" {
		t.Fatal("Expected an error but got nothing")
	}
}

func TestCharacterErrorCondition(t *testing.T) {
	mockCtrl, _, _, mockCharRepo, mockCorpRepo, mockAlliRepo, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	authCreateRequest := proto.AuthCreateRequest{
		Token:       "mytoken",
		Alliance:    &proto.Alliance{Name: "Test Alliance", Id: 1, Ticker: "TSTA"},
		Corporation: &proto.Corporation{Name: "Test Corp", Id: 1, Ticker: "TSTC"},
		Character:   &proto.Character{Name: "Test Character", Id: 1},
		AuthScope:   []string{"scope1", "scope2"},
	}
	var authCreateResponse proto.AuthCreateResponse
	var ctx context.Context

	authHandler := AuthHandler{}

	//Set our expectations
	gomock.InOrder(
		mockAlliRepo.EXPECT().FindByAllianceId(gomock.Any()).Return(nil),
		mockAlliRepo.EXPECT().Save(gomock.Any()).Times(1),

		mockCorpRepo.EXPECT().FindByCorporationId(gomock.Any()).Return(nil),
		mockCorpRepo.EXPECT().Save(gomock.Any()).Times(1),

		mockCharRepo.EXPECT().FindByCharacterId(gomock.Any()).Return(nil),
		mockCharRepo.EXPECT().Save(gomock.Any()).Return(&testError{message: "Don't do that char"}),
	)

	err := authHandler.Create(ctx, &authCreateRequest, &authCreateResponse)

	if err == nil && err.Error() == "Don't do that char" {
		t.Fatal("Expected an error but got nothing")
	}
}

func TestAuthCodeErrorCondition(t *testing.T) {
	mockCtrl, mockAuthRepo, _, mockCharRepo, mockCorpRepo, mockAlliRepo, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	authCreateRequest := proto.AuthCreateRequest{
		Token:       "mytoken",
		Alliance:    &proto.Alliance{Name: "Test Alliance", Id: 1, Ticker: "TSTA"},
		Corporation: &proto.Corporation{Name: "Test Corp", Id: 1, Ticker: "TSTC"},
		Character:   &proto.Character{Name: "Test Character", Id: 1},
		AuthScope:   []string{"scope1", "scope2"},
	}
	var authCreateResponse proto.AuthCreateResponse
	var ctx context.Context

	authHandler := AuthHandler{}

	//Set our expectations
	gomock.InOrder(
		mockAlliRepo.EXPECT().FindByAllianceId(gomock.Any()).Return(nil),
		mockAlliRepo.EXPECT().Save(gomock.Any()).Times(1),

		mockCorpRepo.EXPECT().FindByCorporationId(gomock.Any()).Return(nil),
		mockCorpRepo.EXPECT().Save(gomock.Any()).Times(1),

		mockCharRepo.EXPECT().FindByCharacterId(gomock.Any()).Return(nil),
		mockCharRepo.EXPECT().Save(gomock.Any()).Return(nil),
		mockAuthRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(errors.New("Don't do that auth")),
	)

	err := authHandler.Create(ctx, &authCreateRequest, &authCreateResponse)

	if err == nil && err.Error() == "Don't do that auth" {
		t.Fatal("Expected an error but got nothing")
	}
}

func TestConfirmWithNoChar(t *testing.T) {
	mockCtrl, _, mockUserRepo, mockCharRepo, _, _, mockAcceRepo := SharedSetup(t)
	defer mockCtrl.Finish()

	var ctx context.Context
	var authConfirmResponse proto.AuthConfirmResponse
	authConfirmRequest := proto.AuthConfirmRequest{
		UserId:             "1234567890",
		AuthenticationCode: "123456789012",
	}

	authHandler := AuthHandler{}

	gomock.InOrder(
		mockCharRepo.EXPECT().FindByAutenticationCode(authConfirmRequest.AuthenticationCode).Return(nil),
		mockUserRepo.EXPECT().FindByChatId(authConfirmRequest.UserId).Times(0),
		mockUserRepo.EXPECT().Save(gomock.Any()).Times(0),
		mockUserRepo.EXPECT().LinkCharacterToUserByAuthCode(gomock.Any(), gomock.Any()).Times(0),
		mockAcceRepo.EXPECT().FindByChatId(gomock.Any()).Times(0),
	)

	err := authHandler.Confirm(ctx, &authConfirmRequest, &authConfirmResponse)

	if err == nil || err.Error() != "Invalid Auth Attempt" {
		t.Errorf("Expected a specific error but received: %s", err)
	}
}

func TestConfirmWithAuthNoUser(t *testing.T) {
	mockCtrl, _, mockUserRepo, mockCharRepo, _, _, mockAcceRepo := SharedSetup(t)

	defer mockCtrl.Finish()

	var ctx context.Context
	var authConfirmResponse proto.AuthConfirmResponse
	authConfirmRequest := proto.AuthConfirmRequest{
		UserId:             "1234567890",
		AuthenticationCode: "123456789012",
	}

	authHandler := AuthHandler{}
	expectedCharName := "Test Character Result"

	allianceId := int64(1)

	gomock.InOrder(
		mockCharRepo.EXPECT().FindByAutenticationCode(authConfirmRequest.AuthenticationCode).Return(
			&model.Character{
				CharacterId:   1,
				CharacterName: expectedCharName,
				CorporationId: 1,
				Corporation: model.Corporation{
					CorporationId:     1,
					CorporationName:   "Test Corporation Result",
					CorporationTicker: "TSTC",
					AllianceId:        &allianceId,
					Alliance: model.Alliance{
						AllianceId:     1,
						AllianceName:   "Test Alliance Result",
						AllianceTicker: "TSTA",
					},
				},
			},
		),
		mockUserRepo.EXPECT().FindByChatId(authConfirmRequest.UserId).Return(nil),
		mockUserRepo.EXPECT().Save(
			&model.User{
				ChatId: authConfirmRequest.UserId,
			},
		),
		mockUserRepo.EXPECT().LinkCharacterToUserByAuthCode(authConfirmRequest.AuthenticationCode,
			&model.User{
				ChatId: authConfirmRequest.UserId,
			},
		).Return(nil),
		mockAcceRepo.EXPECT().FindByChatId(authConfirmRequest.UserId).Return([]string{"ROLE1", "ROLE2", "ROLE3"}, nil),
	)

	err := authHandler.Confirm(ctx, &authConfirmRequest, &authConfirmResponse)

	if err != nil {
		t.Errorf("Expected no error but received: %s", err)
	}

	if len(authConfirmResponse.CharacterName) == 0 || authConfirmResponse.CharacterName == "" {
		t.Errorf("Response character name: (%s) doesn't match expected: (%s)", authConfirmResponse.CharacterName, expectedCharName)
	}
}

func TestGetRoles(t *testing.T) {
	mockCtrl, _, mockUserRepo, _, _, _, mockAcceRepo := SharedSetup(t)
	defer mockCtrl.Finish()

	authHandler := AuthHandler{}

	gomock.InOrder(
		mockUserRepo.EXPECT().FindByChatId("1234567890").Return(&model.User{UserId: 1, ChatId: "1234567890"}),
		mockAcceRepo.EXPECT().FindByChatId("1234567890").Return([]string{"ROLE1", "ROLE2"}, nil),
	)

	response := proto.AuthConfirmResponse{}

	err := authHandler.GetRoles(context.Background(), &proto.GetRolesRequest{UserId: "1234567890"}, &response)

	if err != nil {
		t.Fatal("Received an error when none were expected.")
	}

	if !response.Success {
		t.Fatal("Received false success when true was expected")
	}

	if len(response.CharacterName) != 0 {
		t.Fatalf("Received: (%s) for character name when empty was expected", response.CharacterName)
	}

	if len(response.Roles) != 2 {
		t.Fatalf("Expected 2 response but %d were received", len(response.Roles))
	}
}

func TestGetRolesNoUser(t *testing.T) {
	mockCtrl, _, mockUserRepo, _, _, _, mockAcceRepo := SharedSetup(t)
	defer mockCtrl.Finish()

	authHandler := AuthHandler{}

	gomock.InOrder(
		mockUserRepo.EXPECT().FindByChatId("1234567890").Return(nil),
		mockAcceRepo.EXPECT().FindByChatId("1234567890").Return([]string{"ROLE1", "ROLE2"}, nil).Times(0),
	)

	response := proto.AuthConfirmResponse{}

	err := authHandler.GetRoles(context.Background(), &proto.GetRolesRequest{UserId: "1234567890"}, &response)

	if err != nil {
		t.Fatal("Received an error when none were expected.")
	}

	if response.Success {
		t.Fatal("Received true success when false was expected")
	}

	if len(response.Roles) > 0 {
		t.Fatal("Received some roles when we should have errored out")
	}
}

//TODO: Create test case for AuthWithUserAndChar?  Should overwrite or create new auth record?  DB Model supports multiple auths per character... by why?  Was a drunk?
