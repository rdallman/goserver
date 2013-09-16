# A Go HTTP Server to write Home about

If you haven't already and you're even somewhat unfamiliar with Go, I'd
recommend to first tackle my first tutorial at
<http://github.com/rdallman/gofirst>. A lot of setting up Go is covered there,
and is actually needed to finish this tutorial. 

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
Use the code from `firstserver.go` in the first tutorial as a reference to fill in `server.go`'s
`main()` so that `localhost:8080/` will work with `indexHandler`.
Hint hint, wink wink: This should be 2 lines. I should also mention the
following two web pages that you might want to keep open in a tab:

* <http://golang.org/pkg/net/http/>
* <http://golang.org/pkg/html/template/>

The Go docs are pretty nice, and should prove to be a valuable resource in all
of your trials with Go. Anyway, back on point:

Once you have filled in main, save the file and enter a terminal ( __make sure
you're in `$GOPATH/src/github.com/rdallman/goserver` each time you do this__ ):

```
$ go install
$ goserver
```

__From now on I'll refer to this instruction simply as "compile and run"__

Please do not move on before you are getting a light blue page at
`localhost:8080/` that has some text on it.

### [Say my name](http://b.vimeocdn.com/ts/400/598/400598295_640.jpg)

Now take a look at `indexHandler`. The first line is instantiation of a Page
struct, using what should be the now familiar `:=` syntax.

The second line is a
little less straight forward. `template.ParseFiles()` will return a Template
object and an Error object. Errors are quite common as a multiple return value
in Go. In a lot of cases though, or simply to test if things are working first,
it can be useful to just get a `_`, which means you do not intend to handle
that return value: in our case, the Error. 

This is necessary because Go's
compiler will yell at you if you get a value but do nothing with it. This is
nice, because it will keep your code clean. Go will also yell at you if you try
to compile with a library import that you do not use. Both these caveats can be
annoying at first, and can encourage people to simply comment things out, but I
would encourage you to simply use version control and not commit commented code
(you can roll back to it).

The third line, `t.Execute(w,p)` will simply write
the contents of the template in conjunction with our Page `p` as a response to
the http request. This might make more sense after you get a chance to look at
the template in another minute. Basically, it will take the items we give it
from a `Page` p and put them in the specified `.html` file. 
First let's get our server working, though.

Inside of `indexHandler` you should have seen a `//FIXME`, make it work.

__compile and run__ and repeat until your name appears on the page

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

Make a `gopherHandler` in `server.go` with the appropriate parameters and serve
up the file `gophers.html` (it's a thing) at `localhost:8080/views/gophers`.
Make sure to add a `Title` attribute of type `string` to our `Page` struct, as well.
Feel free to title this page whatever you want for now, "gophers" is plenty appropriate
if you're not feeling creative. Don't forget to add a new `http.HandleFunc` in
main for your new Handler.

If your page still has text on it, something has gone terribly wrong. 

__compile and run__ until you at least see a black bar across the top of the
page at `localhost:8080/views/gophers`

### [Yeahhhhh, if you could just get those pictures working, that'd be greattttt](https://si0.twimg.com/profile_images/1202780279/download.jpg)

So your page may or may not have appeared with a bunch of blank images all over it.
It may already make sense to you why this isn't working and what we need to do
about it. If not, I'll key you in. If we are hosting those images on our
server, it should make sense that we would have to serve them ourself. `ListenAndServe`
can only handle things that we've told it about, and right now it only knows
about "/" and "/views/gophers"

If you like to poke around, you may have noticed the `assets/` directory in our
package. Generally, most web developers keep all of their static "assets" inside
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
because it could come in handy (there's also an exercise at the end that uses
this)

Now you should __compile and run__ and check out `localhost:8080/views/gophers`

After that, if you'd like, you can see our file server at:

```
localhost:8080/assets/
```

If none of the above is working, call somebody over because something has gone
terribly wrong and it's probably not your fault. 

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

Then add the following to `server.go`, replacing "YOURNAME" with your real name:

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
templates are all valid before we try to load any of them, and will "panic" if any
of them are not. "panic" is basically throwing an error, but if you're curious, 
you can find "panic" well-documented in the Go docs. 

Inside of `.Must`, you'll see `template.ParseGlob`, which can take a pattern and parse all of the files, instead of
one at a time like we were doing previously. `*.html` simply means all `.html`.
Pre-compiling the templates has some (huge) speed benefits, as well.

Inside of `renderTemplate`, we have some new stuff.
`templates.ExecuteTemplate()` can replace our old `t.Execute`,
and the transition there should be pretty intuitive. This time, though, we're
actually going to check for the error. Using the parameter to pick the actual
template (`t+".html`) could actually produce an error, so this is good practice.

With this in place, you can go ahead and replace the body of your `indexHandler`
and `gopherHandler` with something much more concise. Both should be one line.
I've left this as an exercise for you, if you will carefully look at our
`renderTemplate` method, the Handler methods should just contain a call to `renderTemplate`.
Make sure to add `Title: "index"` to your `Page` in `indexHandler`, and to be
very careful with your titles, since we're using that to render `title+".html"`
in `renderTemplate()`.

__compile and run__ and make sure that both of your pages still work first, then move
to the next step.

We'll get to clean that up even more soon, but first let's go ahead and add
another page. 


### [I'm a real page](http://blog.lib.umn.edu/graz0029/wednesdaywanderings/aaa.jpg)

Okay, so enough with making really crappy looking web pages. I get it. You want
something that will actually be useful. Let's simplify our go server just a
little bit and then we can serve up a nice web page. 

Add the following to your `server.go`:

```go
const viewsPath = "/views/"

func homeHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len(viewsPath):]
  renderTemplate(w, r, title)
}
```

Now we are getting the title of the page from the actual URL. This should make
sense to do, since we were specifying it twice and rendering the template based off
the name of the page anyway. To do this, we are using a slice. Go slices are
very powerful and very popular in Go, so take a second to make sure that this
makes sense. If you'll look at `[len(viewPath):]`, the colon syntax may make some
degree of sense if you're familiar with Python lists / strings. Go's slices are
very similar to this, but are a little bit lower level and have a better degree
of control. I'm not going to get into the semantics, but there's a really good article at
<http://blog.golang.org/go-slices-usage-and-internals> that will go more in
depth if you'd like to read it later. Back to the colon, you can see that we are
specifying the value of the length of the constant "/views/" (7) and then 'slicing'
it. What this will do, effectively, is get rid of the first 7 bytes ("/views/") 
from `r.URL.PATH` and return whatever is left. So, in the case of
"/views/gophers", title := "gophers". 

Now you should __compile and run__ and point your browser to
`localhost:8080/views/home`. Make sure this is working before moving on (you
should see Jordan-Hare). You don't need to replace all of your handlers with
this new slice stuff, we'll take care of that shortly.

### [Remember, remember the 5th of November](https://si0.twimg.com/profile_images/1769643466/258844_104131489680984_104118713015595_32268_721285_o__1_.jpeg)

Go has first class or anonymous functions (or closures) baked in, which we can use in our web
server to cut down on how much we are currently repeating ourselves. You may
have noticed we have a lot of this now in main():

```go
http.HandleFunc("/views/xxxx", someHandler)
```

We also have some pretty unnecessary handlers now, since they simply include one
pretty simple line (that comes from the parameter). So let's do something about
that. I've wrote up a function that returns a function for you, all that you
need to do is copy the following into your `server.go`:

```go
func handle(t string) (string, http.HandlerFunc) {
  return viewsPath+t, func(w http.ResponseWriter, r *http.Request) {
    p := &Page{ Name: "Reed", Title: t }
    err := templates.ExecuteTemplate(w, t+".html", p)
    if err != nil {
      http.NotFound(w, r)
    }
  }
}
```

This function utilizes Go's multiple returns to go ahead and pretty much
eliminate our redundant specification of the path when we're rendering the 
template based on the name of the page, anyway. We'll concatenate the desired
path that we specify (t) onto our constant `viewsPath` to keep our urls the
same. Now, instead of making a call to `renderTemplate`, we've basically put the
body of that into our anonymous function. This should seem a pretty natural
progression, since all we were doing inside of each Handler was calling
`renderTemplate` with the title, anyway. Now, if we cannot find the template for
the given URL, we can return a 404 to the user. This is nice for security
reasons (previously, any URL worked)

Now, you can go ahead and delete the following methods:

```go
gopherHandler()
homeHandler()
indexHandler()
renderTemplate()
```

I'd encourage you to just forget about `index` for now, since it's really a
matter of taste how to handle that from here on out, I'll leave it as an
exercise for the end. For now, just worry about rendering "gophers" and "home".
To get accustomed to our new method, I'd like for you to fix up `main()`
yourself. Basically all that needs to be changed (re: simplified) are the parameters inside of
each `http.HandleFunc`, and you should go ahead and delete the line for handling
"/" so that we can return a 404 if someone tries to go to any url that isn't
gophers or home. 

__compile and run__ until "/views/gophers" and "/views/home" work as they did
previously. 

### [GOOD JOB MAN](http://cdn.rsvlts.com/wp-content/uploads/2013/01/wwb0d.gif)

Now your entire web server should be two methods, `handle` and `main`. Pretty
cool, huh? That's all it takes to serve up a few web pages. What I've given you
is hopefully a great base to build web apps and/or sites in Go. Obviously what
we have built is rather simple; only 2 pages, but you should now be pretty familiar
with Go's `net/http` package, as well as some of the functionality of the core
library. Go encourages concise, clear code and what I've given you should be
considered "idiomatic" Go code, so it is a good spring board into other things. 
I really hope this was a helpful foray into Go, and that I've piqued your
interest in it.  

Below I've listed a few ideas for where to go from here. These should be useful
from a learning Go standpoint just as much as for learning web development stuff.
What you have now is a great base if you're curious in web development. In the
web 2.0 era, it can't hurt. Plus, there's some really cool stuff going on right
now with javascript, bootstrap, markdown and all sorts of other cool tools. Some
of the stuff I've listed below are a few things that are trending now or that
I've had positive experiences with, but I don't know everything. 
Enough of me ranting though, __go__ hack around! 

### Exercises and Ideas for Expansion

* rename index.html and make it a worthwhile page (toy around with jquery,
  Bootstrap, fonts; I gave you the CDN's in the header)
* simplify all of the HTML, it's super repetitive at the moment. e.g. Learn about
  partials, try using `template.HTML`
* put the CSS in a style sheet and make the pages look how you want
* render "home.html" as your "/" using `http.StripPrefix` (handle() doesn't have
  to return 2 values)
* Listen for changes on the file system to recompile your templates, so no more
  restarting your server
* play around with Handlebars.js or Angular.js or make an app using Ember.js
* make a blog engine in Markdown using "BlackFriday" markdown library
* Turn this into your own personal page
* Push it to Heroku
* Let me know what else you can think up!
