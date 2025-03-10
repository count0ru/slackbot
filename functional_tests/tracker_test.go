// functional_tests/tracker_test.go
package functional_tests

import (
	"context"
	"flag"
	"fmt"
	"os"
	"testing"

	"slackbot/internal/config"
	"slackbot/internal/httpclient"
	"slackbot/pkg/tracker"
)

var (
	jiraURL     string
	jiraToken   string
	project     string
	issuetype   string
	title       string
	description string
	ticketKey   string
)

func init() {
	jiraURL = os.Getenv("JIRA_URL")
	jiraToken = os.Getenv("JIRA_TOKEN")
	project = os.Getenv("JIRA_PROJECT")
	issuetype = os.Getenv("JIRA_ISSUETYPE")
	title = os.Getenv("JIRA_TITLE")
	description = os.Getenv("JIRA_DESCRIPTION")
	ticketKey = os.Getenv("JIRA_TICKET_KEY")

	flag.StringVar(&jiraURL, "jira-url", jiraURL, "Jira URL")
	flag.StringVar(&jiraToken, "jira-token", jiraToken, "Jira Token")
	flag.StringVar(&project, "project", project, "Jira Project")
	flag.StringVar(&issuetype, "issuetype", issuetype, "Jira Issuetype")
	flag.StringVar(&title, "title", title, "Ticket Title")
	flag.StringVar(&description, "description", description, "Ticket Description")
	flag.StringVar(&ticketKey, "ticket-key", ticketKey, "Ticket Key")
	flag.Parse()
}

func newJiraTracker(t *testing.T) tracker.Tracker {
	if jiraURL == "" || jiraToken == "" || project == "" || issuetype == "" {
		t.Skip("Jira URL, Token, Project, and Issuetype are required")
	}

	cfg := &config.Config{
		Tracker: config.TrackerConfig{
			URL:   jiraURL,
			Token: jiraToken,
		},
		Ticket: config.TicketConfig{
			Project: config.ProjectConfig{
				Key: project,
			},
			Issuetype: config.IssuetypeConfig{
				Name: issuetype,
			},
		},
	}

	client, err := httpclient.NewClient()
	if err != nil {
		t.Fatalf("Failed to create HTTP client: %v", err)
	}
	jiraTracker, err := tracker.NewJiraTracker(cfg, client)
	if err != nil {
		t.Fatalf("Failed to create JiraTracker: %v", err)
	}
	return jiraTracker
}

func TestCreateTicket(t *testing.T) {
	if title == "" || description == "" {
		t.Skip("Title and Description are required")
	}

	jiraTracker := newJiraTracker(t)
	ticketKey, err := jiraTracker.CreateTicket(context.Background(), title, description)
	if err != nil {
		t.Fatalf("Failed to create ticket: %v", err)
	}
	fmt.Printf("Created ticket with key: %s\n", ticketKey)
}

func TestGetTicketStatus(t *testing.T) {
	if ticketKey == "" {
		t.Skip("Ticket Key is required")
	}

	jiraTracker := newJiraTracker(t)
	status, err := jiraTracker.GetTicketStatus(context.Background(), ticketKey)
	if err != nil {
		t.Fatalf("Failed to get ticket status: %v", err)
	}
	fmt.Printf("Ticket status: %s\n", status)
}

// Примеры запуска тестов:
// go test -v functional_tests/tracker_test.go -run TestCreateTicket -jira-url="your_jira_url" -jira-token="your_jira_token" -project="your_project" -issuetype="your_issuetype" -title="Test Ticket" -description="Test Description"
// go test -v functional_tests/tracker_test.go -run TestGetTicketStatus -jira-url="your_jira_url" -jira-token="your_jira_token" -project="your_project" -issuetype="your_issuetype" -ticket-key="your_ticket_key"

// export JIRA_URL="your_jira_url"
// export JIRA_TOKEN="your_jira_token"
// export JIRA_PROJECT="your_project"
// export JIRA_ISSUETYPE="your_issuetype"
// export JIRA_TITLE="Test Ticket"
// export JIRA_DESCRIPTION="Test Description"
// export JIRA_TICKET_KEY="your_ticket_key"

// go test -v functional_tests/tracker_test.go -run TestCreateTicket
// go test -v functional_tests/tracker_test.go -run TestGetTicketStatus
