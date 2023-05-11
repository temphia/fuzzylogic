package fuzzyreact

import (
	"fmt"
	"strings"
)

type Action struct {
	Name    string
	Info    string
	Example string

	ActionFunc func(string) (string, error)
}

var (
	actions = []Action{{
		Name:    "wikipedia",
		Example: "Django",
		Info:    "Returns a summary from searching Wikipedia",
	}}
)

func buildPrompt() string {

	var buf strings.Builder

	buf.WriteString(`You run in a loop of Thought, Action, PAUSE, Observation.
At the end of the loop you output an Answer
Use Thought to describe your thoughts about the question you have been asked.
Use Action to run one of the actions available to you - then return PAUSE.
Observation will be the result of running those actions.

Your available actions are:
	
`)

	for _, action := range actions {
		buf.WriteString(fmt.Sprintf("%s:\ne.g. %s: %s\n%s\n", action.Name, action.Name, action.Example, action.Info))
	}

	buf.WriteString("\n")

	buf.WriteString(`Always look things up on Wikipedia if you have the opportunity to do so.
	
Example session:
	
Question: What is the capital of France?
Thought: I should look up France on Wikipedia
Action: wikipedia: France
PAUSE
	
You will be called again with this:
	
Observation: France is a country. The capital is Paris.
	
You then output:
	
Answer: The capital of France is Paris`)

	return buf.String()

}
