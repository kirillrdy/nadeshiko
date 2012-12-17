package nadeshiko

import "fmt"

type Activity interface {
	Start(connection *Connection)
}

func (connection *Connection) SetActivity(activity Activity) {

	for k, v := range Notifications {
		var new_list []*Connection
		for _, a_connection := range v {
			if a_connection != connection {
				new_list = append(new_list,a_connection)
			} else {
				fmt.Printf("Removing Notification '%s' for client that changed activity %v\n", k, connection)
			}
		}
		Notifications[k] = new_list
	}

	connection.JQuery("body").Empty()
	activity.Start(connection)
}

