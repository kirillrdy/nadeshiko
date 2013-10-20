package nadeshiko

func (connection *Connection) IncludeBootstrapCdn() {
	connection.JQuery("body").Append(`<link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.0.0/css/bootstrap.min.css">`)
}

func (connection *Connection) IncludeBootstrapJsCdn() {
	connection.JQuery("body").Append(`<script src="//netdna.bootstrapcdn.com/bootstrap/3.0.0/js/bootstrap.min.js"></script>`)
}

func (element JQuerySelectedElements) Tooltip() {
	element.zeroArgumentMethod("tooltip")
}

