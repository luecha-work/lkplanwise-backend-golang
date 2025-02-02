package mail

import (
	"testing"

	"github.com/lkplanwise-api/utils"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	config, err := utils.LoadConfig("..")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "A test email"
	content := `
	<h1>Hello Gin Framework</h1>
	<p>This is a test email sent using <b>Gin Framework</b>.</p>
	`
	to := []string{"luecha.kanm@gmail.com"}
	attachFiles := []string{"../README.md"}

	err = sender.SendMail(subject, content, to, nil, nil, attachFiles)

	require.NoError(t, err)
}
