package nadeshiko

type Activity interface {
	Start(connection *Connection)
}

func (connection *Connection) StartActivity(activity Activity) {
	//TODO make to be via channel
	CleanupNotification(connection)
	connection.JQuery("body").Empty()
	activity.Start(connection)
}
