### vita

> *n.* a brief biographical sketch

vita is a simple command-line tool that helps you organize your thoughts in
simple UNIX fashion. All it does is manage text files inside your own folder
based on the concepts of notebooks. vita will point you to your latest journal
entry for the day, without you having to manage naming, tagging or thinking
about the newest file you need to create. Simply pipe vita's output as
a filename into your favorite editor!

vita is built with golang, but check out the releases for pre-compiled binaries
for your system and architecture. If you don't see yours there, please let me
know!


#### usage

vita's API is small and low touch, all it aims to do is help organize your
streams of consciousness.

```
# get the filename of today's journal entry for your default notebook
$ vita today
/Users/JohnDoe/journal/ideas/01-22-18.txt

# open today's journal entry in vim to the last line
$ vim + `vita today`

# get the filename of today's journal entry for the 'ideas' notebook
$ vita today ideas
/Users/JohnDoe/journal/ideas/01-22-18.txt

# list all of your existing notebooks
$ vita notebooks
```

#### vita is not...

- a text editor - send the filenames that vita outputs to vim, emacs, etc...
- a version control tool - use git, mercurial, etc to keep a history of your
    text files.
- a markdown renderer/blog - vita only works with text files and keeps track of
    them for you. Hook it up to a static site renderer if you want more.


#### development

vita uses [go-bindata][https://github.com/kevinburke/go-bindata] as a tool to
package its template files so that a single binary can be built and transported
with 0 dependencies.

If you add/modify templates, make sure they are in the `data/` directory, and
run `go-bindata data/...` to convert it to a go file that can be built. See its
README for more advanced usage.
