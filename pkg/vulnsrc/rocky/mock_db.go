// Code generated by mockery v1.0.0. DO NOT EDIT.

package rocky

import (
	db "github.com/khulnasoft-lab/vul-db/pkg/db"
	mock "github.com/stretchr/testify/mock"
	bbolt "go.etcd.io/bbolt"

	types "github.com/khulnasoft-lab/vul-db/pkg/types"
)

// MockDB is an autogenerated mock type for the DB type
type MockDB struct {
	mock.Mock
}

type DBBatchUpdateArgs struct {
	Fn         func(*bbolt.Tx) error
	FnAnything bool
}

type DBBatchUpdateReturns struct {
	Err error
}

type DBBatchUpdateExpectation struct {
	Args    DBBatchUpdateArgs
	Returns DBBatchUpdateReturns
}

func (_m *MockDB) ApplyBatchUpdateExpectation(e DBBatchUpdateExpectation) {
	var args []interface{}
	if e.Args.FnAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Fn)
	}
	_m.On("BatchUpdate", args...).Return(e.Returns.Err)
}

func (_m *MockDB) ApplyBatchUpdateExpectations(expectations []DBBatchUpdateExpectation) {
	for _, e := range expectations {
		_m.ApplyBatchUpdateExpectation(e)
	}
}

// BatchUpdate provides a mock function with given fields: fn
func (_m *MockDB) BatchUpdate(fn func(*bbolt.Tx) error) error {
	ret := _m.Called(fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(*bbolt.Tx) error) error); ok {
		r0 = rf(fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type DBDeleteAdvisoryDetailBucketReturns struct {
	_a0 error
}

type DBDeleteAdvisoryDetailBucketExpectation struct {
	Returns DBDeleteAdvisoryDetailBucketReturns
}

func (_m *MockDB) ApplyDeleteAdvisoryDetailBucketExpectation(e DBDeleteAdvisoryDetailBucketExpectation) {
	var args []interface{}
	_m.On("DeleteAdvisoryDetailBucket", args...).Return(e.Returns._a0)
}

func (_m *MockDB) ApplyDeleteAdvisoryDetailBucketExpectations(expectations []DBDeleteAdvisoryDetailBucketExpectation) {
	for _, e := range expectations {
		_m.ApplyDeleteAdvisoryDetailBucketExpectation(e)
	}
}

// DeleteAdvisoryDetailBucket provides a mock function with given fields:
func (_m *MockDB) DeleteAdvisoryDetailBucket() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type DBDeleteVulnerabilityDetailBucketReturns struct {
	Err error
}

type DBDeleteVulnerabilityDetailBucketExpectation struct {
	Returns DBDeleteVulnerabilityDetailBucketReturns
}

func (_m *MockDB) ApplyDeleteVulnerabilityDetailBucketExpectation(e DBDeleteVulnerabilityDetailBucketExpectation) {
	var args []interface{}
	_m.On("DeleteVulnerabilityDetailBucket", args...).Return(e.Returns.Err)
}

func (_m *MockDB) ApplyDeleteVulnerabilityDetailBucketExpectations(expectations []DBDeleteVulnerabilityDetailBucketExpectation) {
	for _, e := range expectations {
		_m.ApplyDeleteVulnerabilityDetailBucketExpectation(e)
	}
}

// DeleteVulnerabilityDetailBucket provides a mock function with given fields:
func (_m *MockDB) DeleteVulnerabilityDetailBucket() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type DBForEachAdvisoryArgs struct {
	Sources         []string
	SourcesAnything bool
	PkgName         string
	PkgNameAnything bool
}

type DBForEachAdvisoryReturns struct {
	Value map[string]db.Value
	Err   error
}

type DBForEachAdvisoryExpectation struct {
	Args    DBForEachAdvisoryArgs
	Returns DBForEachAdvisoryReturns
}

func (_m *MockDB) ApplyForEachAdvisoryExpectation(e DBForEachAdvisoryExpectation) {
	var args []interface{}
	if e.Args.SourcesAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Sources)
	}
	if e.Args.PkgNameAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.PkgName)
	}
	_m.On("ForEachAdvisory", args...).Return(e.Returns.Value, e.Returns.Err)
}

func (_m *MockDB) ApplyForEachAdvisoryExpectations(expectations []DBForEachAdvisoryExpectation) {
	for _, e := range expectations {
		_m.ApplyForEachAdvisoryExpectation(e)
	}
}

// ForEachAdvisory provides a mock function with given fields: sources, pkgName
func (_m *MockDB) ForEachAdvisory(sources []string, pkgName string) (map[string]db.Value, error) {
	ret := _m.Called(sources, pkgName)

	var r0 map[string]db.Value
	if rf, ok := ret.Get(0).(func([]string, string) map[string]db.Value); ok {
		r0 = rf(sources, pkgName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]db.Value)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string, string) error); ok {
		r1 = rf(sources, pkgName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type DBForEachVulnerabilityIDArgs struct {
	Fn         func(*bbolt.Tx, string) error
	FnAnything bool
}

type DBForEachVulnerabilityIDReturns struct {
	Err error
}

type DBForEachVulnerabilityIDExpectation struct {
	Args    DBForEachVulnerabilityIDArgs
	Returns DBForEachVulnerabilityIDReturns
}

func (_m *MockDB) ApplyForEachVulnerabilityIDExpectation(e DBForEachVulnerabilityIDExpectation) {
	var args []interface{}
	if e.Args.FnAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Fn)
	}
	_m.On("ForEachVulnerabilityID", args...).Return(e.Returns.Err)
}

func (_m *MockDB) ApplyForEachVulnerabilityIDExpectations(expectations []DBForEachVulnerabilityIDExpectation) {
	for _, e := range expectations {
		_m.ApplyForEachVulnerabilityIDExpectation(e)
	}
}

// ForEachVulnerabilityID provides a mock function with given fields: fn
func (_m *MockDB) ForEachVulnerabilityID(fn func(*bbolt.Tx, string) error) error {
	ret := _m.Called(fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(*bbolt.Tx, string) error) error); ok {
		r0 = rf(fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type DBGetArgs struct {
	Release         string
	ReleaseAnything bool
	PkgName         string
	PkgNameAnything bool
	Arch            string
	ArchAnything    bool
}

type DBGetReturns struct {
	_a0 []types.Advisory
	_a1 error
}

type DBGetExpectation struct {
	Args    DBGetArgs
	Returns DBGetReturns
}

func (_m *MockDB) ApplyGetExpectation(e DBGetExpectation) {
	var args []interface{}
	if e.Args.ReleaseAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Release)
	}
	if e.Args.PkgNameAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.PkgName)
	}
	if e.Args.ArchAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Arch)
	}
	_m.On("Get", args...).Return(e.Returns._a0, e.Returns._a1)
}

func (_m *MockDB) ApplyGetExpectations(expectations []DBGetExpectation) {
	for _, e := range expectations {
		_m.ApplyGetExpectation(e)
	}
}

// Get provides a mock function with given fields: release, pkgName, arch
func (_m *MockDB) Get(release string, pkgName string, arch string) ([]types.Advisory, error) {
	ret := _m.Called(release, pkgName, arch)

	var r0 []types.Advisory
	if rf, ok := ret.Get(0).(func(string, string, string) []types.Advisory); ok {
		r0 = rf(release, pkgName, arch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Advisory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(release, pkgName, arch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type DBGetAdvisoriesArgs struct {
	Source          string
	SourceAnything  bool
	PkgName         string
	PkgNameAnything bool
}

type DBGetAdvisoriesReturns struct {
	Advisories []types.Advisory
	Err        error
}

type DBGetAdvisoriesExpectation struct {
	Args    DBGetAdvisoriesArgs
	Returns DBGetAdvisoriesReturns
}

func (_m *MockDB) ApplyGetAdvisoriesExpectation(e DBGetAdvisoriesExpectation) {
	var args []interface{}
	if e.Args.SourceAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Source)
	}
	if e.Args.PkgNameAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.PkgName)
	}
	_m.On("GetAdvisories", args...).Return(e.Returns.Advisories, e.Returns.Err)
}

func (_m *MockDB) ApplyGetAdvisoriesExpectations(expectations []DBGetAdvisoriesExpectation) {
	for _, e := range expectations {
		_m.ApplyGetAdvisoriesExpectation(e)
	}
}

// GetAdvisories provides a mock function with given fields: source, pkgName
func (_m *MockDB) GetAdvisories(source string, pkgName string) ([]types.Advisory, error) {
	ret := _m.Called(source, pkgName)

	var r0 []types.Advisory
	if rf, ok := ret.Get(0).(func(string, string) []types.Advisory); ok {
		r0 = rf(source, pkgName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Advisory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(source, pkgName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type DBGetVulnerabilityArgs struct {
	VulnerabilityID         string
	VulnerabilityIDAnything bool
}

type DBGetVulnerabilityReturns struct {
	Vulnerability types.Vulnerability
	Err           error
}

type DBGetVulnerabilityExpectation struct {
	Args    DBGetVulnerabilityArgs
	Returns DBGetVulnerabilityReturns
}

func (_m *MockDB) ApplyGetVulnerabilityExpectation(e DBGetVulnerabilityExpectation) {
	var args []interface{}
	if e.Args.VulnerabilityIDAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.VulnerabilityID)
	}
	_m.On("GetVulnerability", args...).Return(e.Returns.Vulnerability, e.Returns.Err)
}

func (_m *MockDB) ApplyGetVulnerabilityExpectations(expectations []DBGetVulnerabilityExpectation) {
	for _, e := range expectations {
		_m.ApplyGetVulnerabilityExpectation(e)
	}
}

// GetVulnerability provides a mock function with given fields: vulnerabilityID
func (_m *MockDB) GetVulnerability(vulnerabilityID string) (types.Vulnerability, error) {
	ret := _m.Called(vulnerabilityID)

	var r0 types.Vulnerability
	if rf, ok := ret.Get(0).(func(string) types.Vulnerability); ok {
		r0 = rf(vulnerabilityID)
	} else {
		r0 = ret.Get(0).(types.Vulnerability)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(vulnerabilityID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type DBGetVulnerabilityDetailArgs struct {
	CveID         string
	CveIDAnything bool
}

type DBGetVulnerabilityDetailReturns struct {
	Detail map[types.SourceID]types.VulnerabilityDetail
	Err    error
}

type DBGetVulnerabilityDetailExpectation struct {
	Args    DBGetVulnerabilityDetailArgs
	Returns DBGetVulnerabilityDetailReturns
}

func (_m *MockDB) ApplyGetVulnerabilityDetailExpectation(e DBGetVulnerabilityDetailExpectation) {
	var args []interface{}
	if e.Args.CveIDAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.CveID)
	}
	_m.On("GetVulnerabilityDetail", args...).Return(e.Returns.Detail, e.Returns.Err)
}

func (_m *MockDB) ApplyGetVulnerabilityDetailExpectations(expectations []DBGetVulnerabilityDetailExpectation) {
	for _, e := range expectations {
		_m.ApplyGetVulnerabilityDetailExpectation(e)
	}
}

// GetVulnerabilityDetail provides a mock function with given fields: cveID
func (_m *MockDB) GetVulnerabilityDetail(cveID string) (map[types.SourceID]types.VulnerabilityDetail, error) {
	ret := _m.Called(cveID)

	var r0 map[types.SourceID]types.VulnerabilityDetail
	if rf, ok := ret.Get(0).(func(string) map[types.SourceID]types.VulnerabilityDetail); ok {
		r0 = rf(cveID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[types.SourceID]types.VulnerabilityDetail)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(cveID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type DBPutArgs struct {
	_a0         *bbolt.Tx
	_a0Anything bool
	_a1         PutInput
	_a1Anything bool
}

type DBPutReturns struct {
	_a0 error
}

type DBPutExpectation struct {
	Args    DBPutArgs
	Returns DBPutReturns
}

func (_m *MockDB) ApplyPutExpectation(e DBPutExpectation) {
	var args []interface{}
	if e.Args._a0Anything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args._a0)
	}
	if e.Args._a1Anything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args._a1)
	}
	_m.On("Put", args...).Return(e.Returns._a0)
}

func (_m *MockDB) ApplyPutExpectations(expectations []DBPutExpectation) {
	for _, e := range expectations {
		_m.ApplyPutExpectation(e)
	}
}

// Put provides a mock function with given fields: _a0, _a1
func (_m *MockDB) Put(_a0 *bbolt.Tx, _a1 PutInput) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*bbolt.Tx, PutInput) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type DBPutAdvisoryDetailArgs struct {
	Tx                      *bbolt.Tx
	TxAnything              bool
	VulnerabilityID         string
	VulnerabilityIDAnything bool
	PkgName                 string
	PkgNameAnything         bool
	NestedBktNames          []string
	NestedBktNamesAnything  bool
	Advisory                interface{}
	AdvisoryAnything        bool
}

type DBPutAdvisoryDetailReturns struct {
	Err error
}

type DBPutAdvisoryDetailExpectation struct {
	Args    DBPutAdvisoryDetailArgs
	Returns DBPutAdvisoryDetailReturns
}

func (_m *MockDB) ApplyPutAdvisoryDetailExpectation(e DBPutAdvisoryDetailExpectation) {
	var args []interface{}
	if e.Args.TxAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Tx)
	}
	if e.Args.VulnerabilityIDAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.VulnerabilityID)
	}
	if e.Args.PkgNameAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.PkgName)
	}
	if e.Args.NestedBktNamesAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.NestedBktNames)
	}
	if e.Args.AdvisoryAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Advisory)
	}
	_m.On("PutAdvisoryDetail", args...).Return(e.Returns.Err)
}

func (_m *MockDB) ApplyPutAdvisoryDetailExpectations(expectations []DBPutAdvisoryDetailExpectation) {
	for _, e := range expectations {
		_m.ApplyPutAdvisoryDetailExpectation(e)
	}
}

// PutAdvisoryDetail provides a mock function with given fields: tx, vulnerabilityID, pkgName, nestedBktNames, advisory
func (_m *MockDB) PutAdvisoryDetail(tx *bbolt.Tx, vulnerabilityID string, pkgName string, nestedBktNames []string, advisory interface{}) error {
	ret := _m.Called(tx, vulnerabilityID, pkgName, nestedBktNames, advisory)

	var r0 error
	if rf, ok := ret.Get(0).(func(*bbolt.Tx, string, string, []string, interface{}) error); ok {
		r0 = rf(tx, vulnerabilityID, pkgName, nestedBktNames, advisory)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type DBPutDataSourceArgs struct {
	Tx              *bbolt.Tx
	TxAnything      bool
	BktName         string
	BktNameAnything bool
	Source          types.DataSource
	SourceAnything  bool
}

type DBPutDataSourceReturns struct {
	Err error
}

type DBPutDataSourceExpectation struct {
	Args    DBPutDataSourceArgs
	Returns DBPutDataSourceReturns
}

func (_m *MockDB) ApplyPutDataSourceExpectation(e DBPutDataSourceExpectation) {
	var args []interface{}
	if e.Args.TxAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Tx)
	}
	if e.Args.BktNameAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.BktName)
	}
	if e.Args.SourceAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Source)
	}
	_m.On("PutDataSource", args...).Return(e.Returns.Err)
}

func (_m *MockDB) ApplyPutDataSourceExpectations(expectations []DBPutDataSourceExpectation) {
	for _, e := range expectations {
		_m.ApplyPutDataSourceExpectation(e)
	}
}

// PutDataSource provides a mock function with given fields: tx, bktName, source
func (_m *MockDB) PutDataSource(tx *bbolt.Tx, bktName string, source types.DataSource) error {
	ret := _m.Called(tx, bktName, source)

	var r0 error
	if rf, ok := ret.Get(0).(func(*bbolt.Tx, string, types.DataSource) error); ok {
		r0 = rf(tx, bktName, source)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type DBPutRedHatCPEsArgs struct {
	Tx               *bbolt.Tx
	TxAnything       bool
	CpeIndex         int
	CpeIndexAnything bool
	Cpe              string
	CpeAnything      bool
}

type DBPutRedHatCPEsReturns struct {
	Err error
}

type DBPutRedHatCPEsExpectation struct {
	Args    DBPutRedHatCPEsArgs
	Returns DBPutRedHatCPEsReturns
}

func (_m *MockDB) ApplyPutRedHatCPEsExpectation(e DBPutRedHatCPEsExpectation) {
	var args []interface{}
	if e.Args.TxAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Tx)
	}
	if e.Args.CpeIndexAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.CpeIndex)
	}
	if e.Args.CpeAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Cpe)
	}
	_m.On("PutRedHatCPEs", args...).Return(e.Returns.Err)
}

func (_m *MockDB) ApplyPutRedHatCPEsExpectations(expectations []DBPutRedHatCPEsExpectation) {
	for _, e := range expectations {
		_m.ApplyPutRedHatCPEsExpectation(e)
	}
}

// PutRedHatCPEs provides a mock function with given fields: tx, cpeIndex, cpe
func (_m *MockDB) PutRedHatCPEs(tx *bbolt.Tx, cpeIndex int, cpe string) error {
	ret := _m.Called(tx, cpeIndex, cpe)

	var r0 error
	if rf, ok := ret.Get(0).(func(*bbolt.Tx, int, string) error); ok {
		r0 = rf(tx, cpeIndex, cpe)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type DBPutRedHatNVRsArgs struct {
	Tx                 *bbolt.Tx
	TxAnything         bool
	Nvr                string
	NvrAnything        bool
	CpeIndices         []int
	CpeIndicesAnything bool
}

type DBPutRedHatNVRsReturns struct {
	Err error
}

type DBPutRedHatNVRsExpectation struct {
	Args    DBPutRedHatNVRsArgs
	Returns DBPutRedHatNVRsReturns
}

func (_m *MockDB) ApplyPutRedHatNVRsExpectation(e DBPutRedHatNVRsExpectation) {
	var args []interface{}
	if e.Args.TxAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Tx)
	}
	if e.Args.NvrAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Nvr)
	}
	if e.Args.CpeIndicesAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.CpeIndices)
	}
	_m.On("PutRedHatNVRs", args...).Return(e.Returns.Err)
}

func (_m *MockDB) ApplyPutRedHatNVRsExpectations(expectations []DBPutRedHatNVRsExpectation) {
	for _, e := range expectations {
		_m.ApplyPutRedHatNVRsExpectation(e)
	}
}

// PutRedHatNVRs provides a mock function with given fields: tx, nvr, cpeIndices
func (_m *MockDB) PutRedHatNVRs(tx *bbolt.Tx, nvr string, cpeIndices []int) error {
	ret := _m.Called(tx, nvr, cpeIndices)

	var r0 error
	if rf, ok := ret.Get(0).(func(*bbolt.Tx, string, []int) error); ok {
		r0 = rf(tx, nvr, cpeIndices)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type DBPutRedHatRepositoriesArgs struct {
	Tx                 *bbolt.Tx
	TxAnything         bool
	Repository         string
	RepositoryAnything bool
	CpeIndices         []int
	CpeIndicesAnything bool
}

type DBPutRedHatRepositoriesReturns struct {
	Err error
}

type DBPutRedHatRepositoriesExpectation struct {
	Args    DBPutRedHatRepositoriesArgs
	Returns DBPutRedHatRepositoriesReturns
}

func (_m *MockDB) ApplyPutRedHatRepositoriesExpectation(e DBPutRedHatRepositoriesExpectation) {
	var args []interface{}
	if e.Args.TxAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Tx)
	}
	if e.Args.RepositoryAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Repository)
	}
	if e.Args.CpeIndicesAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.CpeIndices)
	}
	_m.On("PutRedHatRepositories", args...).Return(e.Returns.Err)
}

func (_m *MockDB) ApplyPutRedHatRepositoriesExpectations(expectations []DBPutRedHatRepositoriesExpectation) {
	for _, e := range expectations {
		_m.ApplyPutRedHatRepositoriesExpectation(e)
	}
}

// PutRedHatRepositories provides a mock function with given fields: tx, repository, cpeIndices
func (_m *MockDB) PutRedHatRepositories(tx *bbolt.Tx, repository string, cpeIndices []int) error {
	ret := _m.Called(tx, repository, cpeIndices)

	var r0 error
	if rf, ok := ret.Get(0).(func(*bbolt.Tx, string, []int) error); ok {
		r0 = rf(tx, repository, cpeIndices)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type DBPutVulnerabilityArgs struct {
	Tx                      *bbolt.Tx
	TxAnything              bool
	VulnerabilityID         string
	VulnerabilityIDAnything bool
	Vulnerability           types.Vulnerability
	VulnerabilityAnything   bool
}

type DBPutVulnerabilityReturns struct {
	Err error
}

type DBPutVulnerabilityExpectation struct {
	Args    DBPutVulnerabilityArgs
	Returns DBPutVulnerabilityReturns
}

func (_m *MockDB) ApplyPutVulnerabilityExpectation(e DBPutVulnerabilityExpectation) {
	var args []interface{}
	if e.Args.TxAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Tx)
	}
	if e.Args.VulnerabilityIDAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.VulnerabilityID)
	}
	if e.Args.VulnerabilityAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Vulnerability)
	}
	_m.On("PutVulnerability", args...).Return(e.Returns.Err)
}

func (_m *MockDB) ApplyPutVulnerabilityExpectations(expectations []DBPutVulnerabilityExpectation) {
	for _, e := range expectations {
		_m.ApplyPutVulnerabilityExpectation(e)
	}
}

// PutVulnerability provides a mock function with given fields: tx, vulnerabilityID, vulnerability
func (_m *MockDB) PutVulnerability(tx *bbolt.Tx, vulnerabilityID string, vulnerability types.Vulnerability) error {
	ret := _m.Called(tx, vulnerabilityID, vulnerability)

	var r0 error
	if rf, ok := ret.Get(0).(func(*bbolt.Tx, string, types.Vulnerability) error); ok {
		r0 = rf(tx, vulnerabilityID, vulnerability)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type DBPutVulnerabilityDetailArgs struct {
	Tx                      *bbolt.Tx
	TxAnything              bool
	VulnerabilityID         string
	VulnerabilityIDAnything bool
	Source                  types.SourceID
	SourceAnything          bool
	Vulnerability           types.VulnerabilityDetail
	VulnerabilityAnything   bool
}

type DBPutVulnerabilityDetailReturns struct {
	Err error
}

type DBPutVulnerabilityDetailExpectation struct {
	Args    DBPutVulnerabilityDetailArgs
	Returns DBPutVulnerabilityDetailReturns
}

func (_m *MockDB) ApplyPutVulnerabilityDetailExpectation(e DBPutVulnerabilityDetailExpectation) {
	var args []interface{}
	if e.Args.TxAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Tx)
	}
	if e.Args.VulnerabilityIDAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.VulnerabilityID)
	}
	if e.Args.SourceAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Source)
	}
	if e.Args.VulnerabilityAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Vulnerability)
	}
	_m.On("PutVulnerabilityDetail", args...).Return(e.Returns.Err)
}

func (_m *MockDB) ApplyPutVulnerabilityDetailExpectations(expectations []DBPutVulnerabilityDetailExpectation) {
	for _, e := range expectations {
		_m.ApplyPutVulnerabilityDetailExpectation(e)
	}
}

// PutVulnerabilityDetail provides a mock function with given fields: tx, vulnerabilityID, source, vulnerability
func (_m *MockDB) PutVulnerabilityDetail(tx *bbolt.Tx, vulnerabilityID string, source types.SourceID, vulnerability types.VulnerabilityDetail) error {
	ret := _m.Called(tx, vulnerabilityID, source, vulnerability)

	var r0 error
	if rf, ok := ret.Get(0).(func(*bbolt.Tx, string, types.SourceID, types.VulnerabilityDetail) error); ok {
		r0 = rf(tx, vulnerabilityID, source, vulnerability)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type DBPutVulnerabilityIDArgs struct {
	Tx                      *bbolt.Tx
	TxAnything              bool
	VulnerabilityID         string
	VulnerabilityIDAnything bool
}

type DBPutVulnerabilityIDReturns struct {
	Err error
}

type DBPutVulnerabilityIDExpectation struct {
	Args    DBPutVulnerabilityIDArgs
	Returns DBPutVulnerabilityIDReturns
}

func (_m *MockDB) ApplyPutVulnerabilityIDExpectation(e DBPutVulnerabilityIDExpectation) {
	var args []interface{}
	if e.Args.TxAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Tx)
	}
	if e.Args.VulnerabilityIDAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.VulnerabilityID)
	}
	_m.On("PutVulnerabilityID", args...).Return(e.Returns.Err)
}

func (_m *MockDB) ApplyPutVulnerabilityIDExpectations(expectations []DBPutVulnerabilityIDExpectation) {
	for _, e := range expectations {
		_m.ApplyPutVulnerabilityIDExpectation(e)
	}
}

// PutVulnerabilityID provides a mock function with given fields: tx, vulnerabilityID
func (_m *MockDB) PutVulnerabilityID(tx *bbolt.Tx, vulnerabilityID string) error {
	ret := _m.Called(tx, vulnerabilityID)

	var r0 error
	if rf, ok := ret.Get(0).(func(*bbolt.Tx, string) error); ok {
		r0 = rf(tx, vulnerabilityID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type DBRedHatNVRToCPEsArgs struct {
	Nvr         string
	NvrAnything bool
}

type DBRedHatNVRToCPEsReturns struct {
	CpeIndices []int
	Err        error
}

type DBRedHatNVRToCPEsExpectation struct {
	Args    DBRedHatNVRToCPEsArgs
	Returns DBRedHatNVRToCPEsReturns
}

func (_m *MockDB) ApplyRedHatNVRToCPEsExpectation(e DBRedHatNVRToCPEsExpectation) {
	var args []interface{}
	if e.Args.NvrAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Nvr)
	}
	_m.On("RedHatNVRToCPEs", args...).Return(e.Returns.CpeIndices, e.Returns.Err)
}

func (_m *MockDB) ApplyRedHatNVRToCPEsExpectations(expectations []DBRedHatNVRToCPEsExpectation) {
	for _, e := range expectations {
		_m.ApplyRedHatNVRToCPEsExpectation(e)
	}
}

// RedHatNVRToCPEs provides a mock function with given fields: nvr
func (_m *MockDB) RedHatNVRToCPEs(nvr string) ([]int, error) {
	ret := _m.Called(nvr)

	var r0 []int
	if rf, ok := ret.Get(0).(func(string) []int); ok {
		r0 = rf(nvr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(nvr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type DBRedHatRepoToCPEsArgs struct {
	Repository         string
	RepositoryAnything bool
}

type DBRedHatRepoToCPEsReturns struct {
	CpeIndices []int
	Err        error
}

type DBRedHatRepoToCPEsExpectation struct {
	Args    DBRedHatRepoToCPEsArgs
	Returns DBRedHatRepoToCPEsReturns
}

func (_m *MockDB) ApplyRedHatRepoToCPEsExpectation(e DBRedHatRepoToCPEsExpectation) {
	var args []interface{}
	if e.Args.RepositoryAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Repository)
	}
	_m.On("RedHatRepoToCPEs", args...).Return(e.Returns.CpeIndices, e.Returns.Err)
}

func (_m *MockDB) ApplyRedHatRepoToCPEsExpectations(expectations []DBRedHatRepoToCPEsExpectation) {
	for _, e := range expectations {
		_m.ApplyRedHatRepoToCPEsExpectation(e)
	}
}

// RedHatRepoToCPEs provides a mock function with given fields: repository
func (_m *MockDB) RedHatRepoToCPEs(repository string) ([]int, error) {
	ret := _m.Called(repository)

	var r0 []int
	if rf, ok := ret.Get(0).(func(string) []int); ok {
		r0 = rf(repository)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(repository)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type DBSaveAdvisoryDetailsArgs struct {
	Tx            *bbolt.Tx
	TxAnything    bool
	CveID         string
	CveIDAnything bool
}

type DBSaveAdvisoryDetailsReturns struct {
	Err error
}

type DBSaveAdvisoryDetailsExpectation struct {
	Args    DBSaveAdvisoryDetailsArgs
	Returns DBSaveAdvisoryDetailsReturns
}

func (_m *MockDB) ApplySaveAdvisoryDetailsExpectation(e DBSaveAdvisoryDetailsExpectation) {
	var args []interface{}
	if e.Args.TxAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Tx)
	}
	if e.Args.CveIDAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.CveID)
	}
	_m.On("SaveAdvisoryDetails", args...).Return(e.Returns.Err)
}

func (_m *MockDB) ApplySaveAdvisoryDetailsExpectations(expectations []DBSaveAdvisoryDetailsExpectation) {
	for _, e := range expectations {
		_m.ApplySaveAdvisoryDetailsExpectation(e)
	}
}

// SaveAdvisoryDetails provides a mock function with given fields: tx, cveID
func (_m *MockDB) SaveAdvisoryDetails(tx *bbolt.Tx, cveID string) error {
	ret := _m.Called(tx, cveID)

	var r0 error
	if rf, ok := ret.Get(0).(func(*bbolt.Tx, string) error); ok {
		r0 = rf(tx, cveID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
