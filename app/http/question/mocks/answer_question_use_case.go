// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"context"
	"quiz/app/http/question"
	"quiz/domain/usecases/questions"
	"sync"
)

// Ensure, that AnswerQuestionUseCaseMock does implement question.AnswerQuestionUseCase.
// If this is not the case, regenerate this file with moq.
var _ question.AnswerQuestionUseCase = &AnswerQuestionUseCaseMock{}

// AnswerQuestionUseCaseMock is a mock implementation of question.AnswerQuestionUseCase.
//
//	func TestSomethingThatUsesAnswerQuestionUseCase(t *testing.T) {
//
//		// make and configure a mocked question.AnswerQuestionUseCase
//		mockedAnswerQuestionUseCase := &AnswerQuestionUseCaseMock{
//			AnswerQuestionFunc: func(ctx context.Context, input questions.AnswerQuestionInput) (questions.AnswerQuestionOutput, error) {
//				panic("mock out the AnswerQuestion method")
//			},
//		}
//
//		// use mockedAnswerQuestionUseCase in code that requires question.AnswerQuestionUseCase
//		// and then make assertions.
//
//	}
type AnswerQuestionUseCaseMock struct {
	// AnswerQuestionFunc mocks the AnswerQuestion method.
	AnswerQuestionFunc func(ctx context.Context, input questions.AnswerQuestionInput) (questions.AnswerQuestionOutput, error)

	// calls tracks calls to the methods.
	calls struct {
		// AnswerQuestion holds details about calls to the AnswerQuestion method.
		AnswerQuestion []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Input is the input argument value.
			Input questions.AnswerQuestionInput
		}
	}
	lockAnswerQuestion sync.RWMutex
}

// AnswerQuestion calls AnswerQuestionFunc.
func (mock *AnswerQuestionUseCaseMock) AnswerQuestion(ctx context.Context, input questions.AnswerQuestionInput) (questions.AnswerQuestionOutput, error) {
	callInfo := struct {
		Ctx   context.Context
		Input questions.AnswerQuestionInput
	}{
		Ctx:   ctx,
		Input: input,
	}
	mock.lockAnswerQuestion.Lock()
	mock.calls.AnswerQuestion = append(mock.calls.AnswerQuestion, callInfo)
	mock.lockAnswerQuestion.Unlock()
	if mock.AnswerQuestionFunc == nil {
		var (
			answerQuestionOutputOut questions.AnswerQuestionOutput
			errOut                  error
		)
		return answerQuestionOutputOut, errOut
	}
	return mock.AnswerQuestionFunc(ctx, input)
}

// AnswerQuestionCalls gets all the calls that were made to AnswerQuestion.
// Check the length with:
//
//	len(mockedAnswerQuestionUseCase.AnswerQuestionCalls())
func (mock *AnswerQuestionUseCaseMock) AnswerQuestionCalls() []struct {
	Ctx   context.Context
	Input questions.AnswerQuestionInput
} {
	var calls []struct {
		Ctx   context.Context
		Input questions.AnswerQuestionInput
	}
	mock.lockAnswerQuestion.RLock()
	calls = mock.calls.AnswerQuestion
	mock.lockAnswerQuestion.RUnlock()
	return calls
}
