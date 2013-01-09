package nadeshiko

import "log"

var notifications = make(map[string][]*Connection)


var CleanupNotification = make(chan *Connection)


func ListenNotification(notificationType string, connection *Connection) {
	notifications[notificationType] = append(notifications[notificationType], connection)
}

func TriggerNotification(notificationType string, notifier func(*Connection)) {
	for _, j := range notifications[notificationType] {
		notifier(j)
	}
}


func cleanupNotification(connection *Connection) {

	for k, v := range notifications {
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
		notifications[k] = new_list
	}
}


func init() {
	go func() {
		for connection := range CleanupNotification {
			cleanupNotification(connection)
		}
	}()
}
