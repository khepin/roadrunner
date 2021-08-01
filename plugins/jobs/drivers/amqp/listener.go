package amqp

import amqp "github.com/rabbitmq/amqp091-go"

func (j *JobConsumer) listener(deliv <-chan amqp.Delivery) {
	go func() {
		for { //nolint:gosimple
			select {
			case msg, ok := <-deliv:
				if !ok {
					j.log.Info("delivery channel closed, leaving the rabbit listener")
					return
				}

				d, err := j.fromDelivery(msg)
				if err != nil {
					j.log.Error("amqp delivery convert", "error", err)
					continue
				}
				// insert job into the main priority queue
				j.pq.Insert(d)
			}
		}
	}()
}
