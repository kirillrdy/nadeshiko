package nadeshiko

type Activity interface {
	Start(connection WebsocketConnection)
}

func (connection WebsocketConnection) SetActivity(activity Activity) {
	//TODO cleanup old activity notifications and callbacks
	connection.JQuery("body").Empty()
	activity.Start(connection)
}
