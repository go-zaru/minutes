package internal

import (
	"errors"

	"github.com/manifoldco/promptui"
)

// PromptString ...
func PromptString(label string) (string, error) {
	validate := func(input string) error {
		if len(input) < 2 {
			return errors.New("Entries must be 2 chars or more")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}

	result, err := prompt.Run()
	return result, err
}
