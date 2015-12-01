package main

type notificationCenter struct {
	messages []string
}

func NewNotificationCenter() *notificationCenter {
	n := new(notificationCenter)
	n.messages = make([]string, 0)
	return n
}

func (n *notificationCenter) notify(message string) {
	n.messages = append(n.messages, message)
}

func (n *notificationCenter) each(callback func(int, string)) {
	for i := range n.messages {
		callback(i, n.messages[i])
	}
	n.messages = make([]string, 0)
}
