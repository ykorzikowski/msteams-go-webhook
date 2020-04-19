package msteams

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

type Input struct {
  Type           string   `json:"@type"`
  Id             string   `json:"id"`
  IsMultiline    bool     `json:"isMultiline"`
  Title          string   `json:"title"`
}

type PotentialAction struct {
  Type    string     `json:"@type"`
  Name	  string     `json:"name"`
  Inputs  []*Input    `json:"inputs"`
  Actions []*Action   `json:"inputs"`

}

type Action struct {
	Type	string   `json:"@type"`
	Name	string   `json:"name"`
	Url 	string   `json:"target"`
}

type Fact struct {
  Name	string   `json:"name"`
  Value	string   `json:"value"`
}

type Section struct {
  ActivityTitle   	string   `json:"activityTitle"`
  ActivitySubTitle	string   `json:"activitySubTitle"`
  ActivityImage   	string   `json:"activityImage"`
  Facts             []*Fact   `json:"facts,omitempty"`
  Markdown          bool     `json:"markdown"`
}

type Payload struct {
  Type             string                `json:"@type"`
  Context          string                `json:"@context"`
  ThemeColor       string                `json:"themeColor"`
  Summary          string                `json:"summary"`
  Sections         []*Section             `json:"sections"`
  PotentialActions []*PotentialAction     `json:"potentialAction"`
}

func (potentialAction *PotentialAction) AddInput(input Input) *PotentialAction {
	potentialAction.Inputs = append(potentialAction.Inputs, &input)
	return potentialAction
}

func (potentialAction *PotentialAction) AddAction(action Action) *PotentialAction {
	potentialAction.Actions = append(potentialAction.Actions, &action)
	return potentialAction
}

func (section *Section) AddFact(fact Fact) *Section {
	section.Facts = append(section.Facts, &fact)
	return section
}

func (payload *Payload) AddSection(section Section) *Payload {
	payload.Sections = append(payload.Sections, &section)
	return payload
}

func (payload *Payload) AddPotentialAction(potentialAction PotentialAction) *Payload {
	payload.PotentialActions = append(payload.PotentialActions, &potentialAction)
	return payload
}

func redirectPolicyFunc(req gorequest.Request, via []gorequest.Request) error {
	return fmt.Errorf("Incorrect token (redirection)")
}

func Send(webhookUrl string, proxy string, payload Payload) []error {
	request := gorequest.New().Proxy(proxy)
	resp, _, err := request.
		Post(webhookUrl).
		RedirectPolicy(redirectPolicyFunc).
		Send(payload).
		End()

	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		return []error{fmt.Errorf("Error sending msg. Status: %v", resp.Status)}
	}

	return nil
}
