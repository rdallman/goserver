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

Crack open `index.html` and see for yourself! Find:

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
another one (I used a mix of Go and Handlebars, personally). /endrant

Alright, now that you have your confidence:

### [If you wanna live in a butcher shop, I'm gonna treat you like a piece of meat](http://i.imgur.com/fkeRVvV.png)

Congrats, you've served up your first halfway decent HTML in Go.

Now you'll need to make your own handler, hot shot.

Make a `gopherHandler` with the appropriate parameters and serve
up the file `gophers.html` (it's a thing) at `localhost:8080/views/gophers`. 
An empty `Page` will be sufficient.

If your page still has text on it, something has gone terribly wrong. 
Please don't move on until the text is gone (hopefully not a 404 either).

### [Yeahhhhh, if you could just get those pictures working, that'd be greattttt](https://si0.twimg.com/profile_images/1202780279/download.jpg)

It may already make sense to you why this isn't working and what we need to do
about it. If not, I'll key you in. So your page may or may not have appeared
with a bunch of blank images all over it. If we are hosting those images on our
server, it should make sense that we would have to serve them ourself. `ListenAndServe`
isn't nice enough to do this much for us, but I mean, c'mon, give him a break. 
