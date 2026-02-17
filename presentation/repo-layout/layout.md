# Project Layout

There is really no "official" way to layout a large project, but over the years there are some best practices that have been memorialized here:

[https://github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout)

A typical layout will look like:

```
/cmd
/pkg
/internal
/vendor
/web
Taskfile
Makefile
go.mod
go.sum
README.md
LICENSE.md
```

<br/>
<br/>
<br/>

### /cmd

This is typically where you will build the different shippable binaries in your project. For example, if you have both a client and a cli, you might have:

```
/cmd
    /server
        main.go
    /cli
        main.go
```

Both the client and server packages will be `main` with the `main()` function entrypoint. When you `go build` in either of those directories, it will output a `server` and `cli` binary.

<br/>
<br/>
<br/>

### /pkg

This is where all your packages will live that are used by your `server` and `cli` binaries. NOTE: any packages here are available for 3rd parties to import and use.

### /internal

These packages are available for your module to use, but CANNOT be imported by 3rd parties that import your module.

<br/>
<br/>
<br/>

### /vendor

If you want to vendor your go dependencies (`go mod vendor`) this is there all the imported external packages will be downloaded and stored. When building your binaries, code from these repos will be used. If there is no `vendor` directory, go will download the dependencies before building

<br/>
<br/>
<br/>

### /web

For mono-repos that have a web UI / component, this is where the javascript lives. Go has a great feature that allows you to embed file systems into the binary itself, and a common pattern is to build you UI and embed it, mounting it as a website served directly from your binary.

----
[Go Ethos](../ethos.md)