package nadeshiko

import "log"

var Notifications map[string][]*Connection
var CleanupNotification = make(chan *Connection)


func ListenNotification(notificationType string, connection *Connection) {
	Notifications[notificationType] = append(Notifications[notificationType], connection)
}

func TriggerNotification(notificationType string, notifier func(*Connection)) {
	for _, j := range Notifications[notificationType] {
		notifier(j)
	}
}


func cleanupNotification(connection *Connection) {

	for k, v := range Notifications {
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
		Notifications[k] = new_list
	}
}


func init() {
	go func() {
		for connection := range CleanupNotification {
			cleanupNotification(connection)
		}
	}()
}
