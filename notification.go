package nadeshiko

var Notifications map[string] []*Connection

func ListenNotification(notificationType string, connection *Connection){
	Notifications[notificationType] = append(Notifications[notificationType],connection)
}

func TriggerNotification(notificationType string, notifier func(*Connection)) {
	for _, j := range Notifications[notificationType] {
		notifier(j)
	}
}
