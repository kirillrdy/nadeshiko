package nadeshiko

import "log"

type Events map[string][]*Connection

var events = make(Events)

type EventSubscription struct {
	EventName	string
	Connection	*Connection
}

var CleanupEventHandlers = make(chan *Connection)
var SubscribeToEvent = make(chan EventSubscription)
var MutateEvents = make(chan func(Events) Events)


func subscribeToEvent(eventName string, connection *Connection) {
	mutator := func(original_events Events) Events {
		original_events[eventName] = append(original_events[eventName], connection)
		return original_events
	}
	MutateEvents <- mutator

}

func TriggerEvent(eventName string, notifier func(*Connection)) {
	for _, callback := range events[eventName] {
		notifier(callback)
	}
}


//func cleanupEventHandlers(connection *Connection) {
//
//	for k, v := range events{
//		var new_list []*Connection
//		for _, a_connection := range v {
//			if a_connection != connection {
//				new_list = append(new_list, a_connection)
//			} else {
//				if Verbose {
//					log.Printf("Removing Notification '%s' for client %v\n", k, connection)
//				}
//			}
//		}
//		events[k] = new_list
//	}
//}


func init() {
	go func() {
		for connection := range CleanupEventHandlers {
			mutator := func(orignal Events) Events {
				var new_events = make(Events)
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
					new_events[k] = new_list
				}
			return new_events
			}
			MutateEvents <- mutator
		}
	}()

	go func() {
		for eventSubscriber := range SubscribeToEvent {
			subscribeToEvent(eventSubscriber.EventName, eventSubscriber.Connection)
		}
	}()

	go func() {
		for mutato := range MutateEvents {
			events = mutato(events)
		}
	}()

}
