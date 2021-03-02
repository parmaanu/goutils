package tuiutils

import "github.com/AlecAivazis/survey/v2"

// TakeUserInput provides a tui for selecting an input on terminal and returns the selected output and error
func TakeUserInput(message string, options []string) (string, error) {
	commandsMenu := []*survey.Question{
		{
			Name:     "UserCommand",
			Validate: survey.Required,
			Prompt: &survey.Select{
				Message: message,
				Options: options,
				VimMode: true,
			},
		},
	}

	answer := struct {
		UserCommand string
	}{}

	err := survey.Ask(commandsMenu, &answer)
	return answer.UserCommand, err
}
