// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/docker/dhe-deploy/pkg/jobrunner-framework/schema (interfaces: JobrunnerManager)

package schema

import (
	rethinkutil "github.com/docker/dhe-deploy/rethinkutil"
	gomock "github.com/vikstrous/mock/gomock"
	gorethink_v2 "gopkg.in/dancannon/gorethink.v2"
	io "io"
	time "time"
)

// Mock of JobrunnerManager interface
type MockJobrunnerManager struct {
	ctrl     *gomock.Controller
	recorder *_MockJobrunnerManagerRecorder
}

// Recorder for MockJobrunnerManager (not exported)
type _MockJobrunnerManagerRecorder struct {
	mock *MockJobrunnerManager
}

func NewMockJobrunnerManager(ctrl *gomock.Controller) *MockJobrunnerManager {
	mock := &MockJobrunnerManager{ctrl: ctrl}
	mock.recorder = &_MockJobrunnerManagerRecorder{mock}
	return mock
}

func (_m *MockJobrunnerManager) EXPECT() *_MockJobrunnerManagerRecorder {
	return _m.recorder
}

func (_m *MockJobrunnerManager) CancelJob(_param0 string) error {
	ret := _m.ctrl.Call(_m, "CancelJob", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockJobrunnerManagerRecorder) CancelJob(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CancelJob", arg0)
}

func (_m *MockJobrunnerManager) ClaimJob(_param0 string, _param1 string, _param2 time.Time) (*Job, error) {
	ret := _m.ctrl.Call(_m, "ClaimJob", _param0, _param1, _param2)
	ret0, _ := ret[0].(*Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockJobrunnerManagerRecorder) ClaimJob(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ClaimJob", arg0, arg1, arg2)
}

func (_m *MockJobrunnerManager) CountJobsWithActionStatus(_param0 string, _param1 string) (uint, error) {
	ret := _m.ctrl.Call(_m, "CountJobsWithActionStatus", _param0, _param1)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockJobrunnerManagerRecorder) CountJobsWithActionStatus(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CountJobsWithActionStatus", arg0, arg1)
}

func (_m *MockJobrunnerManager) CreateJob(_param0 *Job) error {
	ret := _m.ctrl.Call(_m, "CreateJob", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockJobrunnerManagerRecorder) CreateJob(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateJob", arg0)
}

func (_m *MockJobrunnerManager) DB() string {
	ret := _m.ctrl.Call(_m, "DB")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockJobrunnerManagerRecorder) DB() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DB")
}

func (_m *MockJobrunnerManager) DeleteActionConfig(_param0 string) error {
	ret := _m.ctrl.Call(_m, "DeleteActionConfig", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockJobrunnerManagerRecorder) DeleteActionConfig(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteActionConfig", arg0)
}

func (_m *MockJobrunnerManager) DeleteCron(_param0 string) error {
	ret := _m.ctrl.Call(_m, "DeleteCron", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockJobrunnerManagerRecorder) DeleteCron(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteCron", arg0)
}

func (_m *MockJobrunnerManager) DeleteJob(_param0 string) error {
	ret := _m.ctrl.Call(_m, "DeleteJob", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockJobrunnerManagerRecorder) DeleteJob(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteJob", arg0)
}

func (_m *MockJobrunnerManager) GetActionConfig(_param0 string) (*ActionConfig, error) {
	ret := _m.ctrl.Call(_m, "GetActionConfig", _param0)
	ret0, _ := ret[0].(*ActionConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockJobrunnerManagerRecorder) GetActionConfig(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetActionConfig", arg0)
}

func (_m *MockJobrunnerManager) GetCron(_param0 string) (*Cron, error) {
	ret := _m.ctrl.Call(_m, "GetCron", _param0)
	ret0, _ := ret[0].(*Cron)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockJobrunnerManagerRecorder) GetCron(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetCron", arg0)
}

func (_m *MockJobrunnerManager) GetCronChanges() (<-chan CronChange, io.Closer, error) {
	ret := _m.ctrl.Call(_m, "GetCronChanges")
	ret0, _ := ret[0].(<-chan CronChange)
	ret1, _ := ret[1].(io.Closer)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockJobrunnerManagerRecorder) GetCronChanges() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetCronChanges")
}

func (_m *MockJobrunnerManager) GetJob(_param0 string) (*Job, error) {
	ret := _m.ctrl.Call(_m, "GetJob", _param0)
	ret0, _ := ret[0].(*Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockJobrunnerManagerRecorder) GetJob(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetJob", arg0)
}

func (_m *MockJobrunnerManager) GetJobLogs(_param0 string, _param1 uint, _param2 uint) ([]JobLog, error) {
	ret := _m.ctrl.Call(_m, "GetJobLogs", _param0, _param1, _param2)
	ret0, _ := ret[0].([]JobLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockJobrunnerManagerRecorder) GetJobLogs(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetJobLogs", arg0, arg1, arg2)
}

func (_m *MockJobrunnerManager) GetLeastRecentlyScheduledJobsForWorkerWithAction(_param0 string, _param1 string, _param2 uint, _param3 uint) ([]Job, error) {
	ret := _m.ctrl.Call(_m, "GetLeastRecentlyScheduledJobsForWorkerWithAction", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].([]Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockJobrunnerManagerRecorder) GetLeastRecentlyScheduledJobsForWorkerWithAction(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetLeastRecentlyScheduledJobsForWorkerWithAction", arg0, arg1, arg2, arg3)
}

func (_m *MockJobrunnerManager) GetMostRecentlyScheduledJobs(_param0 uint, _param1 uint) ([]Job, error) {
	ret := _m.ctrl.Call(_m, "GetMostRecentlyScheduledJobs", _param0, _param1)
	ret0, _ := ret[0].([]Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockJobrunnerManagerRecorder) GetMostRecentlyScheduledJobs(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMostRecentlyScheduledJobs", arg0, arg1)
}

func (_m *MockJobrunnerManager) GetMostRecentlyScheduledJobsForWorker(_param0 string, _param1 uint, _param2 uint) ([]Job, error) {
	ret := _m.ctrl.Call(_m, "GetMostRecentlyScheduledJobsForWorker", _param0, _param1, _param2)
	ret0, _ := ret[0].([]Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockJobrunnerManagerRecorder) GetMostRecentlyScheduledJobsForWorker(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMostRecentlyScheduledJobsForWorker", arg0, arg1, arg2)
}

func (_m *MockJobrunnerManager) GetMostRecentlyScheduledJobsWithAction(_param0 string, _param1 uint, _param2 uint) ([]Job, error) {
	ret := _m.ctrl.Call(_m, "GetMostRecentlyScheduledJobsWithAction", _param0, _param1, _param2)
	ret0, _ := ret[0].([]Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockJobrunnerManagerRecorder) GetMostRecentlyScheduledJobsWithAction(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMostRecentlyScheduledJobsWithAction", arg0, arg1, arg2)
}

func (_m *MockJobrunnerManager) GetNewJobChanges() (<-chan JobChange, io.Closer, error) {
	ret := _m.ctrl.Call(_m, "GetNewJobChanges")
	ret0, _ := ret[0].(<-chan JobChange)
	ret1, _ := ret[1].(io.Closer)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockJobrunnerManagerRecorder) GetNewJobChanges() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetNewJobChanges")
}

func (_m *MockJobrunnerManager) GetNextUnclaimedJob(_param0 string) (*Job, error) {
	ret := _m.ctrl.Call(_m, "GetNextUnclaimedJob", _param0)
	ret0, _ := ret[0].(*Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockJobrunnerManagerRecorder) GetNextUnclaimedJob(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetNextUnclaimedJob", arg0)
}

func (_m *MockJobrunnerManager) GetOwnJobCancellations(_param0 string) (<-chan JobChange, io.Closer, error) {
	ret := _m.ctrl.Call(_m, "GetOwnJobCancellations", _param0)
	ret0, _ := ret[0].(<-chan JobChange)
	ret1, _ := ret[1].(io.Closer)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockJobrunnerManagerRecorder) GetOwnJobCancellations(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetOwnJobCancellations", arg0)
}

func (_m *MockJobrunnerManager) GetRecoverableJobChanges(_param0 []string) (<-chan JobChange, io.Closer, error) {
	ret := _m.ctrl.Call(_m, "GetRecoverableJobChanges", _param0)
	ret0, _ := ret[0].(<-chan JobChange)
	ret1, _ := ret[1].(io.Closer)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockJobrunnerManagerRecorder) GetRecoverableJobChanges(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRecoverableJobChanges", arg0)
}

func (_m *MockJobrunnerManager) HeartbeatJob(_param0 string, _param1 time.Time) error {
	ret := _m.ctrl.Call(_m, "HeartbeatJob", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockJobrunnerManagerRecorder) HeartbeatJob(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HeartbeatJob", arg0, arg1)
}

func (_m *MockJobrunnerManager) InsertJobLog(_param0 string, _param1 string, _param2 int) error {
	ret := _m.ctrl.Call(_m, "InsertJobLog", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockJobrunnerManagerRecorder) InsertJobLog(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InsertJobLog", arg0, arg1, arg2)
}

func (_m *MockJobrunnerManager) ListActionConfigs() ([]ActionConfig, error) {
	ret := _m.ctrl.Call(_m, "ListActionConfigs")
	ret0, _ := ret[0].([]ActionConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockJobrunnerManagerRecorder) ListActionConfigs() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListActionConfigs")
}

func (_m *MockJobrunnerManager) ListCrons() ([]Cron, error) {
	ret := _m.ctrl.Call(_m, "ListCrons")
	ret0, _ := ret[0].([]Cron)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockJobrunnerManagerRecorder) ListCrons() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListCrons")
}

func (_m *MockJobrunnerManager) SafeGetActionConfig(_param0 string) (*ActionConfig, error) {
	ret := _m.ctrl.Call(_m, "SafeGetActionConfig", _param0)
	ret0, _ := ret[0].(*ActionConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockJobrunnerManagerRecorder) SafeGetActionConfig(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SafeGetActionConfig", arg0)
}

func (_m *MockJobrunnerManager) ScaleDB(_param0 uint, _param1 bool) error {
	ret := _m.ctrl.Call(_m, "ScaleDB", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockJobrunnerManagerRecorder) ScaleDB(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ScaleDB", arg0, arg1)
}

func (_m *MockJobrunnerManager) Session() *gorethink_v2.Session {
	ret := _m.ctrl.Call(_m, "Session")
	ret0, _ := ret[0].(*gorethink_v2.Session)
	return ret0
}

func (_mr *_MockJobrunnerManagerRecorder) Session() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Session")
}

func (_m *MockJobrunnerManager) SetupDB(_param0 uint) error {
	ret := _m.ctrl.Call(_m, "SetupDB", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockJobrunnerManagerRecorder) SetupDB(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetupDB", arg0)
}

func (_m *MockJobrunnerManager) Tables() []rethinkutil.Table {
	ret := _m.ctrl.Call(_m, "Tables")
	ret0, _ := ret[0].([]rethinkutil.Table)
	return ret0
}

func (_mr *_MockJobrunnerManagerRecorder) Tables() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Tables")
}

func (_m *MockJobrunnerManager) UpdateActionConfig(_param0 *ActionConfig) (*ActionConfig, error) {
	ret := _m.ctrl.Call(_m, "UpdateActionConfig", _param0)
	ret0, _ := ret[0].(*ActionConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockJobrunnerManagerRecorder) UpdateActionConfig(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateActionConfig", arg0)
}

func (_m *MockJobrunnerManager) UpdateCron(_param0 *Cron) (*Cron, error) {
	ret := _m.ctrl.Call(_m, "UpdateCron", _param0)
	ret0, _ := ret[0].(*Cron)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockJobrunnerManagerRecorder) UpdateCron(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateCron", arg0)
}

func (_m *MockJobrunnerManager) UpdateJobStatus(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "UpdateJobStatus", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockJobrunnerManagerRecorder) UpdateJobStatus(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateJobStatus", arg0, arg1)
}