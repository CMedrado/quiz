package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// quizCmd defines the "quiz" command, which starts the interactive quiz
var quizCmd = &cobra.Command{
	Use:   "quiz",
	Short: "Start the interactive quiz",
	Run: func(cmd *cobra.Command, args []string) {
		runQuiz()
	},
}

func init() {
	// Add the quiz command to the root command
	rootCmd.AddCommand(quizCmd)
}

// runQuiz is the core logic of the interactive quiz
func runQuiz() {
	ctx := context.Background()         // Create a background context
	reader := bufio.NewReader(os.Stdin) // Set up input reader

	// Welcome message and prompt for username
	fmt.Println("Welcome to the Quiz!")
	fmt.Print("Enter your username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username) // Trim any extra spaces or newline

	for {
		// Fetch the next question from the API
		question, err := fetchNextQuestion(ctx, username)
		if err != nil {
			// Handle error fetching the question
			fmt.Println("Error fetching the next question!")
			break
		}

		// If no more questions are available, end the quiz
		if question.Question == "" {
			fmt.Println("No more questions. Quiz finished!")
			break
		}

		// Display the question and its possible answers
		fmt.Printf("\nQuestion %d: %s\n", question.NumberQuestion+1, question.Question)
		for i, ans := range question.Answers {
			fmt.Printf("  %d. %s\n", i+1, ans)
		}

		// Prompt the user for their answer
		fmt.Print("Your answer (enter the number): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input) // Clean the input
		answerIndex, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error parse int:", err)
			continue
		}

		// Submit the answer to the API
		correct, err := answerQuestion(ctx, username, question.NumberQuestion, answerIndex-1)
		if err != nil {
			// Handle error submitting the answer
			fmt.Println("Error submitting the answer:", err)
			continue
		}

		// Provide feedback on whether the answer was correct
		if correct {
			fmt.Println("✅ Correct!")
		} else {
			fmt.Println("❌ Incorrect!")
		}

		// Fetch the player's current position and score
		position, err := getPlayerPosition(ctx, username)
		if err != nil {
			// Handle error fetching player position
			fmt.Println("Error getting player position:", err)
			continue
		}

		// Display player's position and score
		fmt.Printf("%s:  Position:%d Score:%d\n", position.UserID, position.Position, position.Score)

		// Ask the user if they want to continue
		fmt.Print("\nDo you want to continue? (yes/no): ")
		cont, _ := reader.ReadString('\n')
		cont = strings.TrimSpace(strings.ToLower(cont)) // Normalize the response

		// If user chooses not to continue, show the leaderboard and exit
		if cont != "yes" {
			rank, err := getLeaderboard(ctx)
			if err != nil {
				// Handle error fetching the leaderboard
				fmt.Println("Error getting rank:", err)
				fmt.Println("Thanks for playing! Your progress is saved.")
				break
			}

			// Display leaderboard
			fmt.Println("Player UserID  Score")
			for _, player := range rank.Players {
				fmt.Printf("%d - %s - %d\n", player.Position, player.UserID, player.Score)
			}

			// End the quiz and thank the player
			fmt.Println("Thanks for playing! Your progress is saved.")
			break
		}
	}
}
