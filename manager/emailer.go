package manager

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const dev = true

// SendPaymentEmail is used to send an email notification to stakers
func (m *Manager) SendPaymentEmail(emailAddress string, coin string, reward *big.Int) (int, error) {
	content := fmt.Sprintf("<br>Payment Received<br>Coin %s<br>Amount %v<br>", coin, reward)
	from := mail.NewEmail("stake-sendgrid-api", "sgapi@rtradetechnologies.com")
	subject := "Stake Payment Receive"
	to := mail.NewEmail("Mining Stake", emailAddress)

	mContent := mail.NewContent("text/html", content)
	mail := mail.NewV3MailInit(from, subject, to, mContent)

	response, err := m.SendGridClient.Send(mail)
	if err != nil {
		return 0, err
	}
	return response.StatusCode, nil
}

// SendNotificationEmail is used to send us an email when we detect a stake in the system
func (m *Manager) SendNotificationEmail(depositer common.Address, amountStaked *big.Int, duration *big.Int, khSec *big.Int, id *big.Int) (int, error) {
	content := fmt.Sprintf("<br>Staker: 0x%x<br><br>RTC Staked: %v<br><br>Weeks Staked: %v<br><br>KhSec: %v<br><br>Stake Id: %v<br>", depositer, amountStaked, duration, khSec, id)
	from := mail.NewEmail("stake-sendgrid-api", "sgapi@rtradetechnologies.com")
	subject := "New Stake Deposit Detected In Staking Contract"
	to := mail.NewEmail("Mining Stake", "stake@rtradetechnologies.com")

	mContent := mail.NewContent("text/html", content)
	mail := mail.NewV3MailInit(from, subject, to, mContent)

	response, err := m.SendGridClient.Send(mail)
	if err != nil {
		return 0, err
	}
	return response.StatusCode, nil

}

// SendEmailsForStakePayout is used to send emalis to stakers when they have been paid out for a stake
func (m *Manager) SendEmailsForStakePayout(stakers map[common.Address]*big.Int, coin string) error {
	for k, v := range stakers {
		stakeID := m.Bolt.RetrieveStakeIDInformationForAddress(k)
		encryptedEmail, err := m.ContractHandler.GetStakerEmailForStakeId(nil, k, stakeID)
		if err != nil {
			return err
		}
		/* email decryption */
		if dev != true {
			_, err = m.SendPaymentEmail(encryptedEmail, coin, v)
			if err != nil {
				return err
			}
		} else {
			_, err = m.SendPaymentEmail("alext@rtradetechnologies.com", coin, v)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
