// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/assisted-service/internal/provider/registry (interfaces: ProviderRegistry)

// Package registry is a generated GoMock package.
package registry

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	common "github.com/openshift/assisted-service/internal/common"
	installcfg "github.com/openshift/assisted-service/internal/installcfg"
	provider "github.com/openshift/assisted-service/internal/provider"
	usage "github.com/openshift/assisted-service/internal/usage"
	models "github.com/openshift/assisted-service/models"
	logrus "github.com/sirupsen/logrus"
)

// MockProviderRegistry is a mock of ProviderRegistry interface.
type MockProviderRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockProviderRegistryMockRecorder
}

// MockProviderRegistryMockRecorder is the mock recorder for MockProviderRegistry.
type MockProviderRegistryMockRecorder struct {
	mock *MockProviderRegistry
}

// NewMockProviderRegistry creates a new mock instance.
func NewMockProviderRegistry(ctrl *gomock.Controller) *MockProviderRegistry {
	mock := &MockProviderRegistry{ctrl: ctrl}
	mock.recorder = &MockProviderRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProviderRegistry) EXPECT() *MockProviderRegistryMockRecorder {
	return m.recorder
}

// AddPlatformToInstallConfig mocks base method.
func (m *MockProviderRegistry) AddPlatformToInstallConfig(arg0 models.PlatformType, arg1 *installcfg.InstallerConfigBaremetal, arg2 *common.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPlatformToInstallConfig", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPlatformToInstallConfig indicates an expected call of AddPlatformToInstallConfig.
func (mr *MockProviderRegistryMockRecorder) AddPlatformToInstallConfig(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPlatformToInstallConfig", reflect.TypeOf((*MockProviderRegistry)(nil).AddPlatformToInstallConfig), arg0, arg1, arg2)
}

// AreHostsSupported mocks base method.
func (m *MockProviderRegistry) AreHostsSupported(arg0 models.PlatformType, arg1 []*models.Host) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AreHostsSupported", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AreHostsSupported indicates an expected call of AreHostsSupported.
func (mr *MockProviderRegistryMockRecorder) AreHostsSupported(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AreHostsSupported", reflect.TypeOf((*MockProviderRegistry)(nil).AreHostsSupported), arg0, arg1)
}

// GenerateProviderManifests mocks base method.
func (m *MockProviderRegistry) GenerateProviderManifests(arg0 context.Context, arg1 *common.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateProviderManifests", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateProviderManifests indicates an expected call of GenerateProviderManifests.
func (mr *MockProviderRegistryMockRecorder) GenerateProviderManifests(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateProviderManifests", reflect.TypeOf((*MockProviderRegistry)(nil).GenerateProviderManifests), arg0, arg1)
}

// Get mocks base method.
func (m *MockProviderRegistry) Get(arg0 string) (provider.Provider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(provider.Provider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockProviderRegistryMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProviderRegistry)(nil).Get), arg0)
}

// GetActualSchedulableMasters mocks base method.
func (m *MockProviderRegistry) GetActualSchedulableMasters(arg0 *common.Cluster) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActualSchedulableMasters", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActualSchedulableMasters indicates an expected call of GetActualSchedulableMasters.
func (mr *MockProviderRegistryMockRecorder) GetActualSchedulableMasters(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActualSchedulableMasters", reflect.TypeOf((*MockProviderRegistry)(nil).GetActualSchedulableMasters), arg0)
}

// GetSupportedProvidersByHosts mocks base method.
func (m *MockProviderRegistry) GetSupportedProvidersByHosts(arg0 []*models.Host) ([]models.PlatformType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSupportedProvidersByHosts", arg0)
	ret0, _ := ret[0].([]models.PlatformType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSupportedProvidersByHosts indicates an expected call of GetSupportedProvidersByHosts.
func (mr *MockProviderRegistryMockRecorder) GetSupportedProvidersByHosts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSupportedProvidersByHosts", reflect.TypeOf((*MockProviderRegistry)(nil).GetSupportedProvidersByHosts), arg0)
}

// InitProviders mocks base method.
func (m *MockProviderRegistry) InitProviders(arg0 logrus.FieldLogger) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InitProviders", arg0)
}

// InitProviders indicates an expected call of InitProviders.
func (mr *MockProviderRegistryMockRecorder) InitProviders(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitProviders", reflect.TypeOf((*MockProviderRegistry)(nil).InitProviders), arg0)
}

// IsHostSupported mocks base method.
func (m *MockProviderRegistry) IsHostSupported(arg0 models.PlatformType, arg1 *models.Host) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsHostSupported", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsHostSupported indicates an expected call of IsHostSupported.
func (mr *MockProviderRegistryMockRecorder) IsHostSupported(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsHostSupported", reflect.TypeOf((*MockProviderRegistry)(nil).IsHostSupported), arg0, arg1)
}

// PostCreateManifestsHook mocks base method.
func (m *MockProviderRegistry) PostCreateManifestsHook(arg0 *common.Cluster, arg1 *[]string, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostCreateManifestsHook", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostCreateManifestsHook indicates an expected call of PostCreateManifestsHook.
func (mr *MockProviderRegistryMockRecorder) PostCreateManifestsHook(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostCreateManifestsHook", reflect.TypeOf((*MockProviderRegistry)(nil).PostCreateManifestsHook), arg0, arg1, arg2)
}

// PreCreateManifestsHook mocks base method.
func (m *MockProviderRegistry) PreCreateManifestsHook(arg0 *common.Cluster, arg1 *[]string, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PreCreateManifestsHook", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PreCreateManifestsHook indicates an expected call of PreCreateManifestsHook.
func (mr *MockProviderRegistryMockRecorder) PreCreateManifestsHook(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreCreateManifestsHook", reflect.TypeOf((*MockProviderRegistry)(nil).PreCreateManifestsHook), arg0, arg1, arg2)
}

// Register mocks base method.
func (m *MockProviderRegistry) Register(arg0 provider.Provider) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Register", arg0)
}

// Register indicates an expected call of Register.
func (mr *MockProviderRegistryMockRecorder) Register(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockProviderRegistry)(nil).Register), arg0)
}

// SetPlatformUsages mocks base method.
func (m *MockProviderRegistry) SetPlatformUsages(arg0 models.PlatformType, arg1 *models.Platform, arg2 map[string]models.Usage, arg3 usage.API) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPlatformUsages", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPlatformUsages indicates an expected call of SetPlatformUsages.
func (mr *MockProviderRegistryMockRecorder) SetPlatformUsages(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPlatformUsages", reflect.TypeOf((*MockProviderRegistry)(nil).SetPlatformUsages), arg0, arg1, arg2, arg3)
}

// SetPlatformValuesInDBUpdates mocks base method.
func (m *MockProviderRegistry) SetPlatformValuesInDBUpdates(arg0 models.PlatformType, arg1 *models.Platform, arg2 map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPlatformValuesInDBUpdates", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPlatformValuesInDBUpdates indicates an expected call of SetPlatformValuesInDBUpdates.
func (mr *MockProviderRegistryMockRecorder) SetPlatformValuesInDBUpdates(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPlatformValuesInDBUpdates", reflect.TypeOf((*MockProviderRegistry)(nil).SetPlatformValuesInDBUpdates), arg0, arg1, arg2)
}
