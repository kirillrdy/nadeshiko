package nadeshiko

type Activity interface {
	Start(connection *Connection)
}

func (connection *Connection) SetActivity(activity Activity) {

	RemoveNotification(connection)

	connection.JQuery("body").Empty()
	activity.Start(connection)
}

