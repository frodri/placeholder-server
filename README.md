placeholder-server
===================

We all know that web development these days can be a challenging endeavor, especially during breakneck phase and constant iteration that occurs during the concept phase. Thankfully, tools like [Bootstrap](http://getbootstrap.com/), [SASS](http://sass-lang.com/), and even sites like [placehold.it](http://placehold.it) make it pretty easy to get good looking sites up and running quickly.

Here's something most people don't think about, though: What happens when you feel the sudden, irrepressible need to do it all on Comic Sans? I mean, sure, putting "Comic Sans" on bootstrap's @font-family-sans-serif LESS variable might take care of most of the site, but placehold.it sure as hell ain't gonna placate anyone's wishes for whimsical, poorly-kerned captioned images.

That's where *placeholder-server* comes in.

## Overview

*placeholder-server* is a configurable placeholder generator written in Go, with a bevy of cool features:

- **A single executable:** Getting *placeholder-server* to run does not require IIS, Apache, mod_php, a JVM, servlets or any other sort of server setup.
- **Configurable Truetype font-support:** For those times when you want to mix it up with Comic Sans Bold.
- **A familiar interface:** Supports custom background/foreground colors, custom captions, JPEG/GIF/PNG outputs and variable width/height, just like placehold.it does.

## Version support

Tested to compile and run on Go 1.6+ and Windows. The code should be cross-platform compatible, but if any problems do come up, do let me know via the Github Issue tracker.

## Building

*placeholder-server* is currently only available as source code. That means that in order to get *placeholder-server* to run, it has to be built.

To do so, follow these steps:

- Set up a Go development environment. For Windows users, [this link](http://www.wadewegner.com/2014/12/easy-go-programming-setup-for-windows/) provides very good instructions on how to do so. The best link I've found for [setting up Go on Mac](http://todsul.com/tech/setup-golang-on-mac-os-x/) does require some knowledge of the command line. For Linux users, well... let's just say [it](https://developer.fedoraproject.org/tech/languages/go/go-installation.html) [depends](https://wiki.archlinux.org/index.php/Go) [on](https://github.com/golang/go/wiki/Ubuntu) [the](https://en.opensuse.org/SDB:Go) [distro.](https://golang.org/doc/install/source)
- Using a command prompt/terminal, navigate to the `GOPATH` workspace created on the previous step, and execute the *go get* command on the box below. This will create a chain of directories inside the `src/` that will follow a structure similar to the URL in the aforementioned *go get* command.
```
go get github.com/frodri/placeholder-server
``` 
- Using that same command prompt, navigate to the *placeholder-server* directory, and run the go build command on the box below. This will build an executable that can then be used to start running the server proper.
```
go build placeholder-server.go
``` 

## Running the server
To run *placeholder-server*, use a command prompt/terminal to run the executable built on the _**Building a local executable**_ section. By default, *placeholder-server* will use port 8080 and the bundled Roboto Regular font in the `fonts/` folder.
```
placeholder-server
```

*placeholder-server* can be configured by using the `--port` and `--fontfile` command flags. The following command will run *placeholder-server* on port 80 while using the Comic Sans font on a Windows machine. 

```
placeholder-server --port=80 --fontfile="C:\Windows\Fonts\comic.ttf"
```
Linux/OSX operates similarly, after making the usual path adjustments. This example would host *placeholder-server* on the same port, while hosting the Liberation Sans font on a Ubuntu Linux machine.

```
placeholder-server --port=80 --fontfile="/usr/share/fonts/truetype/liberation/LiberationSans-Regular.ttf"
```

A help command is also available for quick reference.

```
placeholder-server --help
```

### Obtaining images

After starting the server, accessing the images is mostly a matter of pointing a web browser to the proper address. The link on the box below, for example, would fetch a 200x100 JPEG image from a local *placeholder-server* instance. 
```
http://localhost:8080/200x100.jpg
``` 
This next example would fetch a 360x360 captioned image with a white background from *placeholder-server* running on a hypothetical Heroku instance on port 80.
```
http://spam-eggs-8408.herokuapp.com/360.png/000000?text=no+scope
```
Further instructions on how to fine-tune the generated images are available by accessing the site root. (http://localhost:8080 for the first example, http://spam-eggs-8408.herokuapp.com for the second)


### Dependencies

*placeholder-server* uses the following dependencies:

- spf13's [*Cobra* CLI library](https://github.com/spf13/cobra) for things like POSIX-style flags and the built-in help command.
- The Gorilla toolkit's [*mux* library](http://www.gorillatoolkit.org/pkg/mux) for URL routing with regexes.
- Golang's [Freetype font rasterizer](https://github.com/golang/freetype) for drawing Truetype fonts into the images.
- lucasb-eyer's [*go-colorful* library](https://github.com/lucasb-eyer/go-colorful) for color proximity calculations and a color.Color constructor that supports CSS-style hex notation.

For convenience purposes, these libraries have been added into the `vendor/` directory, along with a Godep folder for older versions of Go.

### Known Issues

- Fontface sizing on its current state is *very* naive. Doing things like adding long text to a 100x400 image will cause the text to expand past its image bounds.
- Word wrapping for extremely long texts has also not been implemented.


### License 
MIT, with the exception of the Roboto font (Apache license) in the `fonts/` directory.
