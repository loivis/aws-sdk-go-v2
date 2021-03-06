// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package mturkiface provides an interface to enable mocking the Amazon Mechanical Turk service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package mturkiface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mturk"
)

// MTurkAPI provides an interface to enable mocking the
// mturk.MTurk service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Mechanical Turk.
//    func myFunc(svc mturkiface.MTurkAPI) bool {
//        // Make svc.AcceptQualificationRequest request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := mturk.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockMTurkClient struct {
//        mturkiface.MTurkAPI
//    }
//    func (m *mockMTurkClient) AcceptQualificationRequest(input *mturk.AcceptQualificationRequestInput) (*mturk.AcceptQualificationRequestOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockMTurkClient{}
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
type MTurkAPI interface {
	AcceptQualificationRequestRequest(*mturk.AcceptQualificationRequestInput) mturk.AcceptQualificationRequestRequest

	ApproveAssignmentRequest(*mturk.ApproveAssignmentInput) mturk.ApproveAssignmentRequest

	AssociateQualificationWithWorkerRequest(*mturk.AssociateQualificationWithWorkerInput) mturk.AssociateQualificationWithWorkerRequest

	CreateAdditionalAssignmentsForHITRequest(*mturk.CreateAdditionalAssignmentsForHITInput) mturk.CreateAdditionalAssignmentsForHITRequest

	CreateHITRequest(*mturk.CreateHITInput) mturk.CreateHITRequest

	CreateHITTypeRequest(*mturk.CreateHITTypeInput) mturk.CreateHITTypeRequest

	CreateHITWithHITTypeRequest(*mturk.CreateHITWithHITTypeInput) mturk.CreateHITWithHITTypeRequest

	CreateQualificationTypeRequest(*mturk.CreateQualificationTypeInput) mturk.CreateQualificationTypeRequest

	CreateWorkerBlockRequest(*mturk.CreateWorkerBlockInput) mturk.CreateWorkerBlockRequest

	DeleteHITRequest(*mturk.DeleteHITInput) mturk.DeleteHITRequest

	DeleteQualificationTypeRequest(*mturk.DeleteQualificationTypeInput) mturk.DeleteQualificationTypeRequest

	DeleteWorkerBlockRequest(*mturk.DeleteWorkerBlockInput) mturk.DeleteWorkerBlockRequest

	DisassociateQualificationFromWorkerRequest(*mturk.DisassociateQualificationFromWorkerInput) mturk.DisassociateQualificationFromWorkerRequest

	GetAccountBalanceRequest(*mturk.GetAccountBalanceInput) mturk.GetAccountBalanceRequest

	GetAssignmentRequest(*mturk.GetAssignmentInput) mturk.GetAssignmentRequest

	GetFileUploadURLRequest(*mturk.GetFileUploadURLInput) mturk.GetFileUploadURLRequest

	GetHITRequest(*mturk.GetHITInput) mturk.GetHITRequest

	GetQualificationScoreRequest(*mturk.GetQualificationScoreInput) mturk.GetQualificationScoreRequest

	GetQualificationTypeRequest(*mturk.GetQualificationTypeInput) mturk.GetQualificationTypeRequest

	ListAssignmentsForHITRequest(*mturk.ListAssignmentsForHITInput) mturk.ListAssignmentsForHITRequest

	ListAssignmentsForHITPages(*mturk.ListAssignmentsForHITInput, func(*mturk.ListAssignmentsForHITOutput, bool) bool) error
	ListAssignmentsForHITPagesWithContext(aws.Context, *mturk.ListAssignmentsForHITInput, func(*mturk.ListAssignmentsForHITOutput, bool) bool, ...aws.Option) error

	ListBonusPaymentsRequest(*mturk.ListBonusPaymentsInput) mturk.ListBonusPaymentsRequest

	ListBonusPaymentsPages(*mturk.ListBonusPaymentsInput, func(*mturk.ListBonusPaymentsOutput, bool) bool) error
	ListBonusPaymentsPagesWithContext(aws.Context, *mturk.ListBonusPaymentsInput, func(*mturk.ListBonusPaymentsOutput, bool) bool, ...aws.Option) error

	ListHITsRequest(*mturk.ListHITsInput) mturk.ListHITsRequest

	ListHITsPages(*mturk.ListHITsInput, func(*mturk.ListHITsOutput, bool) bool) error
	ListHITsPagesWithContext(aws.Context, *mturk.ListHITsInput, func(*mturk.ListHITsOutput, bool) bool, ...aws.Option) error

	ListHITsForQualificationTypeRequest(*mturk.ListHITsForQualificationTypeInput) mturk.ListHITsForQualificationTypeRequest

	ListHITsForQualificationTypePages(*mturk.ListHITsForQualificationTypeInput, func(*mturk.ListHITsForQualificationTypeOutput, bool) bool) error
	ListHITsForQualificationTypePagesWithContext(aws.Context, *mturk.ListHITsForQualificationTypeInput, func(*mturk.ListHITsForQualificationTypeOutput, bool) bool, ...aws.Option) error

	ListQualificationRequestsRequest(*mturk.ListQualificationRequestsInput) mturk.ListQualificationRequestsRequest

	ListQualificationRequestsPages(*mturk.ListQualificationRequestsInput, func(*mturk.ListQualificationRequestsOutput, bool) bool) error
	ListQualificationRequestsPagesWithContext(aws.Context, *mturk.ListQualificationRequestsInput, func(*mturk.ListQualificationRequestsOutput, bool) bool, ...aws.Option) error

	ListQualificationTypesRequest(*mturk.ListQualificationTypesInput) mturk.ListQualificationTypesRequest

	ListQualificationTypesPages(*mturk.ListQualificationTypesInput, func(*mturk.ListQualificationTypesOutput, bool) bool) error
	ListQualificationTypesPagesWithContext(aws.Context, *mturk.ListQualificationTypesInput, func(*mturk.ListQualificationTypesOutput, bool) bool, ...aws.Option) error

	ListReviewPolicyResultsForHITRequest(*mturk.ListReviewPolicyResultsForHITInput) mturk.ListReviewPolicyResultsForHITRequest

	ListReviewPolicyResultsForHITPages(*mturk.ListReviewPolicyResultsForHITInput, func(*mturk.ListReviewPolicyResultsForHITOutput, bool) bool) error
	ListReviewPolicyResultsForHITPagesWithContext(aws.Context, *mturk.ListReviewPolicyResultsForHITInput, func(*mturk.ListReviewPolicyResultsForHITOutput, bool) bool, ...aws.Option) error

	ListReviewableHITsRequest(*mturk.ListReviewableHITsInput) mturk.ListReviewableHITsRequest

	ListReviewableHITsPages(*mturk.ListReviewableHITsInput, func(*mturk.ListReviewableHITsOutput, bool) bool) error
	ListReviewableHITsPagesWithContext(aws.Context, *mturk.ListReviewableHITsInput, func(*mturk.ListReviewableHITsOutput, bool) bool, ...aws.Option) error

	ListWorkerBlocksRequest(*mturk.ListWorkerBlocksInput) mturk.ListWorkerBlocksRequest

	ListWorkerBlocksPages(*mturk.ListWorkerBlocksInput, func(*mturk.ListWorkerBlocksOutput, bool) bool) error
	ListWorkerBlocksPagesWithContext(aws.Context, *mturk.ListWorkerBlocksInput, func(*mturk.ListWorkerBlocksOutput, bool) bool, ...aws.Option) error

	ListWorkersWithQualificationTypeRequest(*mturk.ListWorkersWithQualificationTypeInput) mturk.ListWorkersWithQualificationTypeRequest

	ListWorkersWithQualificationTypePages(*mturk.ListWorkersWithQualificationTypeInput, func(*mturk.ListWorkersWithQualificationTypeOutput, bool) bool) error
	ListWorkersWithQualificationTypePagesWithContext(aws.Context, *mturk.ListWorkersWithQualificationTypeInput, func(*mturk.ListWorkersWithQualificationTypeOutput, bool) bool, ...aws.Option) error

	NotifyWorkersRequest(*mturk.NotifyWorkersInput) mturk.NotifyWorkersRequest

	RejectAssignmentRequest(*mturk.RejectAssignmentInput) mturk.RejectAssignmentRequest

	RejectQualificationRequestRequest(*mturk.RejectQualificationRequestInput) mturk.RejectQualificationRequestRequest

	SendBonusRequest(*mturk.SendBonusInput) mturk.SendBonusRequest

	SendTestEventNotificationRequest(*mturk.SendTestEventNotificationInput) mturk.SendTestEventNotificationRequest

	UpdateExpirationForHITRequest(*mturk.UpdateExpirationForHITInput) mturk.UpdateExpirationForHITRequest

	UpdateHITReviewStatusRequest(*mturk.UpdateHITReviewStatusInput) mturk.UpdateHITReviewStatusRequest

	UpdateHITTypeOfHITRequest(*mturk.UpdateHITTypeOfHITInput) mturk.UpdateHITTypeOfHITRequest

	UpdateNotificationSettingsRequest(*mturk.UpdateNotificationSettingsInput) mturk.UpdateNotificationSettingsRequest

	UpdateQualificationTypeRequest(*mturk.UpdateQualificationTypeInput) mturk.UpdateQualificationTypeRequest
}

var _ MTurkAPI = (*mturk.MTurk)(nil)
