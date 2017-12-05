// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package xrayiface provides an interface to enable mocking the AWS X-Ray service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package xrayiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/xray"
)

// XRayAPI provides an interface to enable mocking the
// xray.XRay service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS X-Ray.
//    func myFunc(svc xrayiface.XRayAPI) bool {
//        // Make svc.BatchGetTraces request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := xray.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockXRayClient struct {
//        xrayiface.XRayAPI
//    }
//    func (m *mockXRayClient) BatchGetTraces(input *xray.BatchGetTracesInput) (*xray.BatchGetTracesOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockXRayClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type XRayAPI interface {
	BatchGetTraces(*xray.BatchGetTracesInput) (*xray.BatchGetTracesOutput, error)
	BatchGetTracesWithContext(aws.Context, *xray.BatchGetTracesInput, ...request.Option) (*xray.BatchGetTracesOutput, error)
	BatchGetTracesRequest(*xray.BatchGetTracesInput) (*request.Request, *xray.BatchGetTracesOutput)

	BatchGetTracesPages(*xray.BatchGetTracesInput, func(*xray.BatchGetTracesOutput, bool) bool) error
	BatchGetTracesPagesWithContext(aws.Context, *xray.BatchGetTracesInput, func(*xray.BatchGetTracesOutput, bool) bool, ...request.Option) error

	GetServiceGraph(*xray.GetServiceGraphInput) (*xray.GetServiceGraphOutput, error)
	GetServiceGraphWithContext(aws.Context, *xray.GetServiceGraphInput, ...request.Option) (*xray.GetServiceGraphOutput, error)
	GetServiceGraphRequest(*xray.GetServiceGraphInput) (*request.Request, *xray.GetServiceGraphOutput)

	GetServiceGraphPages(*xray.GetServiceGraphInput, func(*xray.GetServiceGraphOutput, bool) bool) error
	GetServiceGraphPagesWithContext(aws.Context, *xray.GetServiceGraphInput, func(*xray.GetServiceGraphOutput, bool) bool, ...request.Option) error

	GetTraceGraph(*xray.GetTraceGraphInput) (*xray.GetTraceGraphOutput, error)
	GetTraceGraphWithContext(aws.Context, *xray.GetTraceGraphInput, ...request.Option) (*xray.GetTraceGraphOutput, error)
	GetTraceGraphRequest(*xray.GetTraceGraphInput) (*request.Request, *xray.GetTraceGraphOutput)

	GetTraceGraphPages(*xray.GetTraceGraphInput, func(*xray.GetTraceGraphOutput, bool) bool) error
	GetTraceGraphPagesWithContext(aws.Context, *xray.GetTraceGraphInput, func(*xray.GetTraceGraphOutput, bool) bool, ...request.Option) error

	GetTraceSummaries(*xray.GetTraceSummariesInput) (*xray.GetTraceSummariesOutput, error)
	GetTraceSummariesWithContext(aws.Context, *xray.GetTraceSummariesInput, ...request.Option) (*xray.GetTraceSummariesOutput, error)
	GetTraceSummariesRequest(*xray.GetTraceSummariesInput) (*request.Request, *xray.GetTraceSummariesOutput)

	GetTraceSummariesPages(*xray.GetTraceSummariesInput, func(*xray.GetTraceSummariesOutput, bool) bool) error
	GetTraceSummariesPagesWithContext(aws.Context, *xray.GetTraceSummariesInput, func(*xray.GetTraceSummariesOutput, bool) bool, ...request.Option) error

	PutTelemetryRecords(*xray.PutTelemetryRecordsInput) (*xray.PutTelemetryRecordsOutput, error)
	PutTelemetryRecordsWithContext(aws.Context, *xray.PutTelemetryRecordsInput, ...request.Option) (*xray.PutTelemetryRecordsOutput, error)
	PutTelemetryRecordsRequest(*xray.PutTelemetryRecordsInput) (*request.Request, *xray.PutTelemetryRecordsOutput)

	PutTraceSegments(*xray.PutTraceSegmentsInput) (*xray.PutTraceSegmentsOutput, error)
	PutTraceSegmentsWithContext(aws.Context, *xray.PutTraceSegmentsInput, ...request.Option) (*xray.PutTraceSegmentsOutput, error)
	PutTraceSegmentsRequest(*xray.PutTraceSegmentsInput) (*request.Request, *xray.PutTraceSegmentsOutput)
}

var _ XRayAPI = (*xray.XRay)(nil)
