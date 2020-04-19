[![GoDoc](https://godoc.org/github.com/ykorzikowski/msteams-go-webhook?status.svg)](https://godoc.org/github.com/ykorzikowski/msteams-go-webhook)

# msteams-go-webhook

Go Lang library to send messages to Slack via Incoming Webhooks.

## Usage
```go
package main

import "github.com/ykorzikowski/msteams-go-webhook"
import "fmt"

func main() {
    webhookUrl := "https://hooks.slack.com/services/foo/bar/baz"

    section1 := msteams.Section {
      ActivityTitle: "![TestImage](https://47a92947.ngrok.io/Content/Images/default.png)Larry Bryant created a new task",
      ActivitySubTitle: "On Project Tango",
      ActivityImage: "https://teamsnodesample.azurewebsites.net/static/img/image5.png",
      Markdown: true
    }
    section1.AddFact(msteams.Fact { Name: "Assigned To", Value: "Unassigned" })

    potentialAction1 := msteams.PotentialAction {
      Type: "ActionCard",
      Name: "Add a comment",
    }
    potentialAction1.AddInput ( msteams.Input {Type: "TextInput", Id: "comment", IsMultiline: false, Title: "Add a comment here for this task"})
    potentialAction1.AddAction( msteams.Action {Type: "HttpPOST", Name: "Add comment", Target: "http://..."})

    payload := msteams.Payload {
      Type: "MessageCard",
      Context: "http://schema.org/extensions",
      ThemeColor: "0076D7",
      Summary: "Larry Bryant created a new task",
      Sections: []msteams.Section{section1},
      PotentialActions: []msteams.PotentialAction{potentialAction1},
    }

    err := msteams.Send(webhookUrl, "", payload)
    if len(err) > 0 {
      fmt.Printf("error: %s\n", err)
    }
}
```

## License
Licensed under the Apache License, Version 2.0: http://www.apache.org/licenses/LICENSE-2.0

## Reference used
https://github.com/ashwanthkumar/slack-go-webhook
