package nadeshiko

type Activity interface {
	Start(connection *Connection)
}

func (connection *Connection) StartActivity(activity Activity) {
	//TODO consider not doing this on Start
	CleanupEventHandlers <- connection

	connection.JQuery("body").Empty()
	activity.Start(connection)
}
