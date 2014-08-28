package nadeshiko

type Activity interface {
	Start(connection *Connection)
}

func (connection *Connection) StartActivity(activity Activity) {
	connection.JQuery("body").Empty()
	activity.Start(connection)
}

func StartActivity(activity Activity) {
	defaultRoutes.Activity("/", activity)
	startWithPortVerbose(3000, false)
}
