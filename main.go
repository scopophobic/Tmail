package main

import (
	"Tmail/imap"
	"Tmail/tui"
	"fmt"
	"os"
)

func main() {
	fmt.Println("📬 Terminal Email Client")

	// Fetch credentials from environment variables
	username := os.Getenv("EMAIL_ADDRESS")
	password := os.Getenv("EMAIL_PASSWORD")
	imapServer := os.Getenv("IMAP_SERVER")

	if username == "" || password == "" || imapServer == "" {
		fmt.Println("❌ Error: Missing environment variables. Please set EMAIL_ADDRESS, EMAIL_PASSWORD, and IMAP_SERVER.")
		return
	}

	// Fetch emails
	emails, err := imap.FetchEmails(username, password, imapServer)
	if err != nil {
		fmt.Println("❌ Error fetching emails:", err)
		return
	}

	// Show emails in TUI
	tui.DisplayEmails(emails)
}
