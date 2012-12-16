package nadeshiko

var Notifications map[string] []*WebsocketConnection

func ListenNotification(notificationType string, connection *WebsocketConnection){
	Notifications[notificationType] = append(Notifications[notificationType],connection)
}

func TriggerNotification(notificationType string, notifier func(*WebsocketConnection)) {
	for _, j := range Notifications[notificationType] {
		notifier(j)
	}
}
