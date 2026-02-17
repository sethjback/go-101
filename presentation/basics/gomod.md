# The `go.mod` file

> A module is a collection of related Go packages that are released together.

Modules are the way go manages project dependencies. Any package in your own module can be accessed directly, but if you want to use an external package NOT in the standard library, it needs to be in your `go.mod` file.

Fortunately, `go` will handle this for you pretty well via a few commands

<br/>
<br/>
<br/>

### Initialize your project

The first thing you need to do


```shell
$> go mod init <module name>
```

Module names can be anything, but typically you will want them to be the part of a repo and use the access to URL as your module path


```shell
$> go mod init github.com/sethjback/go-101
```

<br/>
<br/>
<br/>


### Accessing internal packages

Any package you create that lives "under" the module, i.e. is a path inside your module, can be accessed directly.

This implies that the `go.mod` will be in the "root" of your repo, but not necessarily. It has to be in the root of your module.

This:
```
/radstuff
    /radserver
        go.mod
    /radclient
        go.mod
```

Or:
```
/radstuff
go.mod
    /radserver
    /radclient
```

<br/>
<br/>
<br/>


### Brining in external packages

In the root of you module, run `go get`:

```shell
go get github.com/go-chi/chi/v5
```

The go mod now looks like:

```yaml
module github.com/sethjback/go-101

go 1.25.7

require github.com/go-chi/chi/v5 v5.2.5
```

you can now add any package from `github.com/go-chi/chi/v5` in your `import` statements. If you import something and it isn't in your `go.mod` you will get a warning. Fortunately, running `go mod tidy` will auto add any imports into you `go.mod` file

<br/>
<br/>
<br/>


### go.sum

Handles versioning for dependencies

```
github.com/go-chi/chi/v5 v5.2.5 h1:Eg4myHZBjyvJmAFjFvWgrqDTXFyOzjj7YIm3L3mu6Ug=
github.com/go-chi/chi/v5 v5.2.5/go.mod h1:X7Gx4mteadT3eDOMTsXzmI4/rwUpOwBHLpAfupzFJP0=
```

---
[Repo Layout](../repo-layout/layout.md)