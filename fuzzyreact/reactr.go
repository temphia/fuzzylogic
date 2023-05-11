package fuzzyreact

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/sashabaranov/go-openai"
)

var (
	actionRe = regexp.MustCompile(`^Action: (\w+): (.*)$`)
)

type FuzzyReAct struct {
	client  *openai.Client
	actions map[string]func(qstr string) (string, error)
}

func New(apiKey string) *FuzzyReAct {

	client := openai.NewClient(apiKey)

	return &FuzzyReAct{
		client:  client,
		actions: make(map[string]func(qstr string) (string, error)),
	}
}

func (fr *FuzzyReAct) Execute(question, system string, max int) error {

	next := question

	if system == "" {
		system = buildPrompt()
	}

	fmt.Println(system)

	bot := newChatBot(fr.client, system)

	i := 0
	for max > i {
		fmt.Println("@loop |>", i)
		i = i + 1
		result, err := bot.execute(next)
		if err != nil {
			return err
		}

		actions := []*regexp.Regexp{}
		for _, line := range strings.Split(result, "\n") {
			if match := actionRe.FindStringSubmatch(line); match != nil {
				actions = append(actions, actionRe)
			}
		}

		if len(actions) == 0 {
			return nil
		}

		actionMatch := actions[0].FindStringSubmatch(result)
		action, actionInput := actionMatch[1], actionMatch[2]

		actionFn := fr.actions[action]
		if actionFn == nil {
			return fmt.Errorf("Unknown action: %s: %s", action, actionInput)
		}

		fmt.Printf(" -- running %s %s\n", action, actionInput)

		observation, err := actionFn(actionInput)
		if err != nil {
			return err
		}

		next = fmt.Sprintf("Observation: %s", observation)
		bot.addMessage("assistant", observation)

	}

	return nil
}
