package nadeshiko

import "log"

var events = make(map[string][]*Connection)

type EventSubscription struct {
	EventName	string
	Connection	*Connection
}

var CleanupEventHandlers = make(chan *Connection)
var SubscribeToEvent = make(chan EventSubscription)


func subscribeToEvent(eventName string, connection *Connection) {
	events[eventName] = append(events[eventName], connection)
}

func TriggerEvent(eventName string, notifier func(*Connection)) {
	for _, callback := range events[eventName] {
		notifier(callback)
	}
}


func cleanupEventHandlers(connection *Connection) {

	for k, v := range events{
		var new_list []*Connection
		for _, a_connection := range v {
			if a_connection != connection {
				new_list = append(new_list, a_connection)
			} else {
				if Verbose {
					log.Printf("Removing Notification '%s' for client %v\n", k, connection)
				}
			}
		}
		events[k] = new_list
	}
}


func init() {
	go func() {
		for connection := range CleanupEventHandlers {
			cleanupEventHandlers(connection)
		}
	}()

	go func() {
		for eventSubscriber := range SubscribeToEvent {
			subscribeToEvent(eventSubscriber.EventName, eventSubscriber.Connection)
		}
	}()
}
