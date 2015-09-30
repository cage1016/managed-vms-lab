// Copyright 2014 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

var initTime = time.Now()

func main() {
	http.HandleFunc("/module1/sayhi", sayHiHandle)
	http.HandleFunc("/", handle)
	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	log.Infof(c, "Serving the module1 page.")
	tmpl.Execute(w, time.Since(initTime))
}

func sayHiHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hi!")
}

var tmpl = template.Must(template.New("front").Parse(`
<html><body>
<p>
Hello, this is Module 1!
</p>
<p>
This instance has been running for <em>{{.}}</em>.
</p>
</body></html>
`))
