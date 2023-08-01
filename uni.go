package main

import (
	"flag"
	"fmt"
	"context"
	"log"
	"os"
	"os/exec"
	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
)



func main() {
	commandFlag := flag.String("c", "", "The command to execute (e.g., \"create a new folder named new\")")
	execFlag := flag.Bool("exec", false, "Execute the specified command")
	flag.Parse()
	godotenv.Load()
	if *commandFlag == "" {
		fmt.Println("Error: No command specified. Please provide a command using the -c flag.")
		return
	}
	var gptResponse = executeCommand(*commandFlag)
	if *execFlag {
		executeUnixCommand(gptResponse)
	} else {
		fmt.Println("Generated command:")
		fmt.Println(*commandFlag)
	}
}

func executeCommand(command string) string {
	// Parse the user input and extract the folder name
	fmt.Printf("Command:  %s",command)
	// Pass this through GPT 

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatalln("Missing API KEY")
	}

	ctx := context.Background()
	client := gpt3.NewClient(apiKey)

	resp, err := client.ChatCompletion(ctx, gpt3.ChatCompletionRequest{
		Messages:  []gpt3.ChatCompletionRequestMessage{
			{
				Role: "system",
				Content: "YOU ARE A LINUX BASED ASSISTANT WHO HELPS PEOPLE TO WRITE IN LINUX. YOU ONLY SPEAK IN UNIX LANGUAGE. YOU DON'T HALLUCINATE AND REPLY ONLY IN ONE LINE WHICH IS THE UNIX COMMAND",
			},
			{
				Role: "user",
				Content: command,
			},
		},
		Stop:[]string{"."},
	})
	if err != nil {
		log.Fatalln(err)
	}
	var response string = resp.Choices[0].Message.Content
	fmt.Println()
	fmt.Printf("Unix Command: %s",response)
	return response
}

func executeUnixCommand(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)

	// Collect the standard output of the command
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

