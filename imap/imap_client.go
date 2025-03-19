package imap

import (
	"fmt"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

// FetchEmails connects to the IMAP server and retrieves emails
func FetchEmails(username, password, server string) ([]string, error) {
	// Connect to server
	c, err := client.DialTLS(server, nil)
	if err != nil {
		return nil, err
	}
	defer c.Logout()

	// Login
	if err := c.Login(username, password); err != nil {
		return nil, err
	}

	// Select Inbox
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		return nil, err
	}

	// Get latest emails
	from := uint32(1)
	to := mbox.Messages
	if mbox.Messages > 10 {
		from = mbox.Messages - 9
	}
	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)

	messages := make(chan *imap.Message, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope}, messages)
	}()

	var emailList []string
	for msg := range messages {
		emailList = append(emailList, fmt.Sprintf("📩 %s - %s", msg.Envelope.From[0].Address(), msg.Envelope.Subject))
	}

	return emailList, nil
}
