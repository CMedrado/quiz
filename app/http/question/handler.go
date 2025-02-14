package question

// Handler struct for managing question-related HTTP requests.
type Handler struct {
	getQuestionUseCase  GetQuestionUseCase
	submitAnswerUseCase AnswerQuestionUseCase
}

// NewQuestionHandler Constructor for QuestionHandler.
func NewQuestionHandler(
	getQuestionUseCase GetQuestionUseCase,
	submitAnswerUseCase AnswerQuestionUseCase,
) *Handler {
	return &Handler{
		getQuestionUseCase:  getQuestionUseCase,
		submitAnswerUseCase: submitAnswerUseCase,
	}
}
