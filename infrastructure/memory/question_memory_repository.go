package repository

import (
	"context"
	"errors"

	"quiz/domain/entities"
)

var (
	ErrRepositoryNoMoreQuestions = errors.New("could not get question because no more questions")
)

// MemoryQuestionStorage is a repository that provides access to quiz questions stored in memory.
type MemoryQuestionStorage struct {
	questions []entities.Question
}

// NewMemoryQuestionStorage creates and returns a new instance of MemoryRepository.
func NewMemoryQuestionStorage() *MemoryQuestionStorage {
	return &MemoryQuestionStorage{
		[]entities.Question{
			{ID: 0, Question: "What is the national language of Malta?", Answer: []string{"Italian", "Maltese", "English"}, AnswerID: 1},
			{ID: 1, Question: "What is the capital city of Malta?", Answer: []string{"Valletta", "Floriana", "Mdina"}, AnswerID: 0},
			{ID: 2, Question: "Which language is widely spoken in Malta?", Answer: []string{"Italian", "English", "Maltese"}, AnswerID: 2},
			{ID: 3, Question: "Malta is located in which sea?", Answer: []string{"Adriatic Sea", "Mediterranean Sea", "Aegean Sea"}, AnswerID: 1},
			{ID: 4, Question: "What is the official currency of Malta?", Answer: []string{"Euro", "Pound Sterling", "US Dollar"}, AnswerID: 0},
			{ID: 5, Question: "What is the national dish of Malta?", Answer: []string{"Rabbit Stew (Stuffat tal-Fenek)", "Fish and Chips", "Pizza"}, AnswerID: 0},
			{ID: 6, Question: "Who painted the Mona Lisa?", Answer: []string{"Vincent van Gogh", "Leonardo da Vinci", "Pablo Picasso"}, AnswerID: 1},
			{ID: 7, Question: "What is the largest planet in our solar system?", Answer: []string{"Earth", "Jupiter", "Saturn"}, AnswerID: 1},
			{ID: 8, Question: "Which element has the chemical symbol 'O'?", Answer: []string{"Oxygen", "Gold", "Osmium"}, AnswerID: 0},
			{ID: 9, Question: "In which country is the Eiffel Tower located?", Answer: []string{"Italy", "France", "Spain"}, AnswerID: 1},
			{ID: 10, Question: "What is the longest river in the world?", Answer: []string{"Amazon River", "Nile River", "Yangtze River"}, AnswerID: 1},
			{ID: 11, Question: "Which famous structure is located in the city of Valletta, Malta?", Answer: []string{"St. John's Co-Cathedral", "Colosseum", "Pyramids of Giza"}, AnswerID: 0},
			{ID: 12, Question: "Malta became a member of the European Union in which year?", Answer: []string{"2004", "2000", "1999"}, AnswerID: 0},
			{ID: 13, Question: "What is the Maltese term for the national symbol of Malta?", Answer: []string{"Il-Kurċifiss", "Il-Bandiera", "Il-Fjamma"}, AnswerID: 0},
			{ID: 14, Question: "Which island group is part of Malta?", Answer: []string{"Aeolian Islands", "Balearic Islands", "Maltese Archipelago"}, AnswerID: 2},
			{ID: 15, Question: "What was the name of the ancient Roman city located in Malta?", Answer: []string{"Valletta", "Mdina", "Melite"}, AnswerID: 2},
			{ID: 16, Question: "What is the chemical formula for water?", Answer: []string{"H2O", "O2", "CO2"}, AnswerID: 0},
			{ID: 17, Question: "Which city is known as the Big Apple?", Answer: []string{"Los Angeles", "New York", "Chicago"}, AnswerID: 1},
			{ID: 18, Question: "Which country is home to the Great Barrier Reef?", Answer: []string{"Australia", "New Zealand", "South Africa"}, AnswerID: 0},
			{ID: 19, Question: "What is the capital of Japan?", Answer: []string{"Kyoto", "Tokyo", "Osaka"}, AnswerID: 1},
			{ID: 20, Question: "Who wrote '1984'?", Answer: []string{"George Orwell", "Mark Twain", "J.K. Rowling"}, AnswerID: 0},
			{ID: 21, Question: "What is the name of the main airport in Malta?", Answer: []string{"Malta International Airport", "Valletta Airport", "Luqa Airport"}, AnswerID: 0},
			{ID: 22, Question: "In what year did Malta gain independence from the UK?", Answer: []string{"1964", "1959", "1970"}, AnswerID: 0},
			{ID: 23, Question: "Which popular diving spot is located off the coast of Malta?", Answer: []string{"The Blue Hole", "The Great Barrier Reef", "Scapa Flow"}, AnswerID: 0},
			{ID: 24, Question: "What is the largest island of Malta?", Answer: []string{"Gozo", "Comino", "Malta"}, AnswerID: 2},
		},
	}
}

// GetQuestion retrieves a question based on the user's progress.
// It returns an error if the userProgress index is out of range.
func (r *MemoryQuestionStorage) GetQuestion(ctx context.Context, userProgress int) (entities.Question, error) {
	if userProgress < 0 || userProgress >= len(r.questions) {
		return entities.Question{}, ErrRepositoryNoMoreQuestions
	}

	question := r.questions[userProgress]
	return question, nil
}
