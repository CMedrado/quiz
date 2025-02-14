// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"context"
	"quiz/app/http/question"
	"quiz/domain/usecases/questions"
	"sync"
)

// Ensure, that GetQuestionUseCaseMock does implement question.GetQuestionUseCase.
// If this is not the case, regenerate this file with moq.
var _ question.GetQuestionUseCase = &GetQuestionUseCaseMock{}

// GetQuestionUseCaseMock is a mock implementation of question.GetQuestionUseCase.
//
//	func TestSomethingThatUsesGetQuestionUseCase(t *testing.T) {
//
//		// make and configure a mocked question.GetQuestionUseCase
//		mockedGetQuestionUseCase := &GetQuestionUseCaseMock{
//			GetQuestionFunc: func(ctx context.Context, input questions.GetQuestionInput) (questions.GetQuestionOutput, error) {
//				panic("mock out the GetQuestion method")
//			},
//		}
//
//		// use mockedGetQuestionUseCase in code that requires question.GetQuestionUseCase
//		// and then make assertions.
//
//	}
type GetQuestionUseCaseMock struct {
	// GetQuestionFunc mocks the GetQuestion method.
	GetQuestionFunc func(ctx context.Context, input questions.GetQuestionInput) (questions.GetQuestionOutput, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetQuestion holds details about calls to the GetQuestion method.
		GetQuestion []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Input is the input argument value.
			Input questions.GetQuestionInput
		}
	}
	lockGetQuestion sync.RWMutex
}

// GetQuestion calls GetQuestionFunc.
func (mock *GetQuestionUseCaseMock) GetQuestion(ctx context.Context, input questions.GetQuestionInput) (questions.GetQuestionOutput, error) {
	callInfo := struct {
		Ctx   context.Context
		Input questions.GetQuestionInput
	}{
		Ctx:   ctx,
		Input: input,
	}
	mock.lockGetQuestion.Lock()
	mock.calls.GetQuestion = append(mock.calls.GetQuestion, callInfo)
	mock.lockGetQuestion.Unlock()
	if mock.GetQuestionFunc == nil {
		var (
			getQuestionOutputOut questions.GetQuestionOutput
			errOut               error
		)
		return getQuestionOutputOut, errOut
	}
	return mock.GetQuestionFunc(ctx, input)
}

// GetQuestionCalls gets all the calls that were made to GetQuestion.
// Check the length with:
//
//	len(mockedGetQuestionUseCase.GetQuestionCalls())
func (mock *GetQuestionUseCaseMock) GetQuestionCalls() []struct {
	Ctx   context.Context
	Input questions.GetQuestionInput
} {
	var calls []struct {
		Ctx   context.Context
		Input questions.GetQuestionInput
	}
	mock.lockGetQuestion.RLock()
	calls = mock.calls.GetQuestion
	mock.lockGetQuestion.RUnlock()
	return calls
}
