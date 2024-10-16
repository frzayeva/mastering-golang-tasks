//Define an interface Notifier with a method Notify(message string).
//Implement this interface for two types: EmailNotifier and SMSNotifier.
//EmailNotifier should print "Sending email notification: {message}".
//SMSNotifier should print "Sending SMS notification: {message}".
//Write a function that accepts the Notifier interface and sends a notification.
//

package implementing

import "fmt"

type Notifier interface {
	Notify(message string)
}

type EmailNotifier struct{}

type SMSNotifier struct{}

func (e EmailNotifier) Notify(message string) {

	fmt.Printf(" Sending email notification: %s\n", message)
}

func (s SMSNotifier) Notify(message string) {

	fmt.Printf(" Sending SMS notification: %s\n", message)
}

func Notifieracceptor(n Notifier, message string) {

	n.Notify(message)
}

func Task4() {

	var e Notifier
	var s Notifier

	e = EmailNotifier{}
	s = SMSNotifier{}

	Notifieracceptor(e, "Your item has been shipped ")
	Notifieracceptor(s, "Your item has been shipped")

}
