# A Go HTTP Server to write Home about

### go get it (if you haven't)

`
$ go get github.com/rdallman/goserver
`

### go to it

`
$ cd $GOPATH/src/github.com/rdallman/goserver
`

Right now, this code will not compile. But do not fear:

### [A first task, you have](http://www.empireonline.com/images/features/100greatestcharacters/photos/25.jpg)

Open the `server.go` file in your favorite text editor

```
$ vim server.go
```

You'll notice `main()` is empty except for a `FIXME` and that I've provided an `indexHandler`.
`indexHandler` will work as is. Just worry about `main()` for now.
Use what you learned from your first web server to fill 
in `main()` so that `localhost:8080/` will work with `indexHandler`.
Hint hint, wink wink: This should be 2 lines.

Once you have done that, save the file and enter a terminal

```
$ go install
$ goserver
```

__From now on I'll refer to this instruction simply as "compile and run"__

Please do not move on before you are getting a light blue page at
`localhost:8080/` that has some text on it.

### [Say my name](http://b.vimeocdn.com/ts/400/598/400598295_640.jpg)

Inside of `indexHandler` will be a `//FIXME`, make it work.

That should have been rather trivial, but more important is what the heck
happened? How exactly _did_ your name appear on the page?

Crack open `index.html` and see for yourself. Find:

```
<p>
  {{.Name}}'s First Go Page
</p>
```

If you're not familiar with HTML, that's okay, it won't be very necessary to
finish this tutorial. I just wanted to show y'all how the Go templating works. 

So you'll notice the `{{ }}` syntax. If you'll take a look again at our
``indexHandler`` this should be pretty intuitive. Whenever we declare the `{{
}}` we're telling the HTML that Go code is here. The HTML will
work fine with this blank, as you may have seen previously; it's just an empty
string. You can imagine the power that can be derived from this functionality.
You could, say, fetch some piece of information from elsewhere and then render
it from your server just like so.

There are a million ways to do templating now, and countless arguments over
client (Handlebars.js, Angular.js) vs Server (ours) side templating. Such is
outside the scope of this tutorial. But now you should understand our witch magic
a little better. A good exercise after this tutorial if you're interested in web
development would be to expand on our templating functionality or test out
another one (I use a mix of Go and Handlebars, personally). /endrant

Alright, now that you have your confidence:

### [If you wanna live in a butcher shop, I'm gonna treat you like a piece of meat](http://i.imgur.com/fkeRVvV.png)

Congrats, you've served up your first halfway decent HTML in Go.

Now you'll need to make your own handler, hot shot.

Make a `gopherHandler` with the appropriate parameters and serve
up the file `gophers.html` (it's a thing) at `localhost:8080/views/gophers`.
Make sure to add a `Title` attribute of type `string` to our `Page`, as well.
Feel free to title this page whatever you want, "gophers" is plenty appropriate
if you're not feeling creative. 

If your page still has text on it, something has gone terribly wrong. 
Please don't move on until the text is gone (hopefully not a 404 either).
You'll know if this page has loaded. 

### [Yeahhhhh, if you could just get those pictures working, that'd be greattttt](https://si0.twimg.com/profile_images/1202780279/download.jpg)

So your page may or may not have appeared with a bunch of blank images all over it.
It may already make sense to you why this isn't working and what we need to do
about it. If not, I'll key you in. If we are hosting those images on our
server, it should make sense that we would have to serve them ourself. `ListenAndServe`
can only handle things that we've told it about.

If you like to poke around, you may have noticed the `assets/` directory in our
package. Generally most web developers keep all of their static "assets" inside
of one folder. This is handy for a few reasons. One, it keeps those with anal
retention issues at bay. Two, you only have to "serve" one directory.

Add the following to your main()

```go
http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
```

So there's a lot going on here. This is similar to a lot of our previous
`http.HandleFunc` where it takes a path and a func. If you go to
<http://golang.org/pkg/net/http/#Handle> you can see the signature for
`http.Handle`. It takes a path and a `Handler`. In the above, you can see we
have `http.StripPrefix`, which will return a `Handler`. `http.StripPrefix`
returns a `Handler`, and takes a path and a `Handler`. Wait, wha? It takes a
`Handler` and returns a `Handler`? `http.StripPrefix` can be used in order to
change the current path of the items for the `Handler` in its params. In our
case, we're pretty much using it to add the leading "/". We could serve up the
dir "bob/" at "/jim/" if we wanted by using `http.StripPrefix`. If this doesn't 
make sense, don't sweat it too much, just wanted to show `http.StripPrefix` 
because it could come in handy. 

If you'll take a second, you can go to:

```
localhost:8080/assets/
```

[Woah...](http://i.imgur.com/I4j1rZK.png)

### [Well, what if there is no tomorrow? There wasn't one today](http://www.columbusunderground.com/wp-content/uploads/2012/01/groundhog-day.jpg)

Feel like we're starting to repeat ourselves yet? You may notice the body of
your two handlers being eerily similar. Let's abstract some of that away into
one method and simplify our templates calls before things get too nasty.

First let's move all of our `.html` files into a folder `templates/`:

(these should work on Linux/Mac, if on Windows you are more capable than I)   
```
mkdir templates
mv *.html templates/
```

Then add the following to `server.go`:

```go
var templates = template.Must(template.ParseGlob("templates/*.html"))

func renderTemplate(w http.ResponseWriter, r *http.Request, t string) {
  p := &Page{Name: "YOURNAME", Title: t}
  err := templates.ExecuteTemplate(w, t+".html", p)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
```

Yet again, lots of stuff going on. `template.Must` will make sure that our
templates are all valid before we try to load any of them, and will panic if any
of them are not. `templates.ExecuteTemplate()` can replace our old `t.Execute`,
and the transition there should be pretty intuitive. This time, though, we're
actually going to check for the error. Using the parameter to pick the actual
template (`t+".html`) could actually produce an error, so this is good practice.

With this in place, you can go ahead and replace the body of your `indexHandler`
and `gopherHandler` with something much more concise. Both should be one line.
We'll get to clean that up even more soon, but first let's go ahead and add
another page. 


### taking ideas for third page

### return anonymous func

