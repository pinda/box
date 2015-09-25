# box
Golang client library for [Box View API](http://developers.box.com/view/)

Installation
------------

  go get github.com/pinda/box

Usage
-----
For example creating a new document is as simple as:

```go
boxClient := box.NewClient(<YOUR_API_KEY>)
newDoc := box.DocumentInput{
	URL:        <YOUR_URL_HERE>,
	Thumbnails: "1024x768",
}
doc, err := boxClient.Documents.NewURL(newDoc)
if err != nil {
	return err
}
```

License
-------

The MIT License (MIT)

Contributed by [Joeri Djojosoeparto](https://github.com/pinda),
