package nadeshiko

type Activity interface {
	Start(connection *Connection)
}

func (connection *Connection) StartActivity(activity Activity) {
	CleanupEventHandlers <- connection
	connection.JQuery("body").Empty()
	activity.Start(connection)
}
