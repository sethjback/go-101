# Golang's design decisions

> Go's purpose is therefore not to do research into programming language design; it is to improve the working environment for its designers and their coworkers. Go is more about software engineering than programming language research. Or to rephrase, it is about language design in the service of software engineering.

\- Rob pike

Read, learn, love, live [go proverbs](https://go-proverbs.github.io/)

Watch the [talk](https://www.youtube.com/watch?v=PAAkCSZUG1c) (gophercon 2015)

<br/>
<br/>
<br/>

## Favor explicitness over implicitness (clear > clever)

This sometimes leads to more verbosity, but as a whole the language wants you to be explicit


<br/>
<br/>

### Beware the `interface{}` type - use only as a safety valve

```golang
// BOO!
func convert(T interface{}) string {
    switch reflect.TypeOf(T) {
        ...
    }
}
```

```golang
// YAY!
func covertInt(i int) string {

}

func convertFloat(f float32) string {

}

func convertUser(u User) string {

}
```


<br/>

### Handle Errors (vs. `panic`)

Convention is to return errors as the last value from a function. This forces you to handle errors states where they occur.

```golang
err := doSomething()
if err != nil {
    //handle error
}
```

For the love of all that is sacred don't panic in libraries and if you do INFORM POTENTIAL USERS. If a function does not return an error, it will be assumed that the function CAN'T produce an error state.



<br/>

### Only require what you actually need as inputs to functions

```golang

type Services struct {
    DB *sql.DB
    Logger logging.Logger
    Limits limits.Checker
    Emailer email.Sender
}

// BAD
func SaveUser(user *User, servcies *Services) error {
    // puts user in the database
}

// Good
func SaveUser(user *User, db *sql.Db, logger *logging.Logger) error {

}

```

<br/>

### Avoid `init()`

a safety vale that I have yet to run across a valid use for.

<br/>
<br/>
<br/>

## Leverage the Standard Library

Don't import the kitchen sink (this isn't javascript). Re-implement if you need to (keep it local, keep it explicit)

This is a no-no: 

```golang
package toolbelt

func Kebab(s string) string {
	return strings.ToLower(splitAndJoin(s, "-"))
}

func Upper(s string) string {
	return strings.ToUpper(s)
}

func Lower(s string) string {
	return strings.ToLower(s)
}
```

<br/>
<br/>
<br/>

## Useful Zero Values

Built in types all have "zero" values. E.g. `bools == false`, `int == 0`. When considering state, leverage the zero value's meaning.

```golang
type User struct{
    Admin bool // most users aren't admins.
}
```

<br/>
<br/>
<br/>

## Avoid Global state

You can use package vars.

You should _not_ use package vars.

package vars are not constants.

```golang
package Example

var user User // shared by the entire package, and your app if you bring in the package
```

Counter example: the `stdlib` `log` [package](https://cs.opensource.google/go/go/+/refs/tags/go1.26.0:src/log/log.go)

Bad example `cobra`


<br/>
<br/>
<br/>

## Comment your code, but only use GOOD comments

Large topic (See the [Docs](https://go.dev/doc/comment))

The `stdlib` is your friend here - follow it's lead. E.g. [strings](https://pkg.go.dev/strings)


<br/>
<br/>
<br/>

## Leverage built in tooling

`gofumpt`

`gofmt`

`go vet`


<br/>
<br/>
<br/>

## Tests are easy to write and fast to run. You have no excuses.

Any file in a package directly that ends in `_test.go` will be run by:
```shell
go test
```

Tests are as easy as:

```golang
func TestConvertUser(t *testing.T) {
    
}
```

You can run all your tests in one go with:
```shell
go test ./...
```