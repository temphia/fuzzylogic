package fuzzyreact

import (
	"fmt"
	"strings"

	"github.com/sashabaranov/go-openai"
)

type FuzzyReAct struct {
	client  *openai.Client
	system  string
	actions map[string]func(qstr string) (string, error)
}

func NewReActWithActions(apiKey string, actions []Action) *FuzzyReAct {
	system, aFuns := BuildPrompt(actions)
	return NewReActWithSystem(apiKey, system, aFuns)

}

func NewReAct(apiKey string) *FuzzyReAct {
	return NewReActWithActions(apiKey, GetActions())
}

func NewReActWithSystem(apiKey, system string,
	actionFuncs map[string]func(string) (string, error)) *FuzzyReAct {

	client := openai.NewClient(apiKey)

	return &FuzzyReAct{
		client:  client,
		system:  system,
		actions: actionFuncs,
	}
}

func (fr *FuzzyReAct) Execute(question string, max int) error {

	next := question

	fmt.Println(fr.system)

	bot := newChatBot(fr.client, fr.system)

	i := 0
	for max > i {
		fmt.Println("@loop |>", i)
		i = i + 1
		result, err := bot.execute(next)
		if err != nil {
			return err
		}

		fmt.Println("@bot_execute", result)
		fmt.Println("@actions", actions)

		observation, err := fr.process(result)
		if err != nil {
			return err
		}

		if observation == "" {
			return nil
		}

		fmt.Println("@OBJ", observation)

		next = fmt.Sprintf("Observation: %s", observation)
		bot.addMessage("assistant", observation)
	}

	return nil
}

func (fr *FuzzyReAct) process(rtext string) (string, error) {
	idx := strings.Index(rtext, "Action:")
	if idx < 0 {
		return "", nil
	}

	rest := rtext[idx+7:]

	actionIndex := strings.Index(rest, ":")
	if actionIndex < 0 {
		return "", fmt.Errorf("not found")
	}

	action := strings.TrimSpace(rest[:actionIndex])
	input := strings.TrimSpace(rest[actionIndex+1:])

	actionFn := fr.actions[action]
	if actionFn == nil {
		return "", fmt.Errorf("Unknown action: %s=>%s", action, input)
	}

	fmt.Printf("Running => %s %s\n", action, input)

	return actionFn(strings.Replace(input, "PAUSE", "", 1))
}
