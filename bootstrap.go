package nadeshiko

func (document *Document) IncludeBootstrapCdn() {
	document.JQuery("body").Append(`<link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.0.0/css/bootstrap.min.css">`)
}

func (document *Document) IncludeBootstrapJsCdn() {
	document.JQuery("body").Append(`<script src="//netdna.bootstrapcdn.com/bootstrap/3.0.0/js/bootstrap.min.js"></script>`)
}

func (element JQuerySelectedElements) Tooltip() {
	element.zeroArgumentMethod("tooltip")
}
