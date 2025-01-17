package nats

import (
	"log"

	"github.com/nats-io/nats.go"
)

type NatServer struct {
	Url    string
	Client *nats.Conn
}

func (n *NatServer) Connect() error {
	nc, err := nats.Connect(n.Url)
	if err != nil {
		log.Fatalf("Error connecting to NATS: %v", err)
	}

	n.Client = nc

	return err
}

func (n *NatServer) Subscribe(subjects []string) {
	for _, subject := range subjects {
		_, err := n.Client.Subscribe(subject, func(msg *nats.Msg) {
			log.Printf("Received message on [%s]: %s", msg.Subject, string(msg.Data))
			// Process the message here (e.g., log, store in database, trigger an action, etc.)
		})

		if err != nil {
			log.Fatalf("Error subscribing to subject [%s]: %v", subject, err)
		}

		log.Printf("Subscribed to [%s]", subject)
	}
}
