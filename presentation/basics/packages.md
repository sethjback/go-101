# Go Packages

Packages are the building blocks of a go application. Technically they are:

> A _package_ is a collection of source files in the same directory that are compiled together. Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package.

### Packages are tied to a directory

You can have as many `.go` files as you want in a single directory, but they ALL must belong to the same package.

Each `.go` file must declare it's package as the first non-comment line of code in the file:

```golang
package data

// code here
```

Technically package names do NOT have to match their directory name, _but_ this makes importing them awkward. The convention is to have package names match the directory name.

<br/>
<br/>
<br/>

### You can have sub directories with separate packages:

```
/data
    data.go
    access.go
    /sql
        driver.go
    /kv
        driver.go
```

The above example has 3 packages: `data`, `sql`, and `kv`.

<br/>
<br/>
<br/>


### Import cycles!!!

Go does not allow circular imports. Take the following:

```
/server
    api.go
/data
    db.go
```

Server, in it's update user function, imports our `data` package to abstract away the data storage details:

```golang
package server

import "data"

type User struct {
    Name string
}

func UpdateUser(ctx context.Context, u *User) error {
    if user.Name == "joe" {
        //save the user in the database
        err := data.SaveUser(u)
        if err != nil {
            return err
        }
    }
}

```


Data uses the data models defined in `server` for saving and retrieving values:

```golang 
package data

import "server"

func SaveUser(u *server.User) error {
    // code
}

func GetUser() (*server.User, error) {
    //code
}

```

They import each other == circular dependency. This is usually an indication of poorly delineated responsibilities, and rather than try to deal with the cycle at the compiler level go just makes you refactor.

Import cycles will clue you in to either combine packages or to split them up.

```golang
package repo

import "data"

// our models
type User struct {
    Name string
}
```

This can get confusing when your cycles are not direct, e.g., `A` imports `B`, which imports `C`, which imports `A`.

<br/>
<br/>
<br/>

### Libraries vs Binaries

`main` is a special package - it is the entry point for running a go program. If you want to `go build` a binary that you can ship and run you need a `main`.

The package directory does not have to have match the package name:

```
/cmd
    main.go
```

```golang
package main

func main() {
    // main loop
}

```


<br/>
<br/>
<br/>


### Exporting Functionality

You can export `func` signatures and data models from your package by naming them with a capital letter:

```golang
package data

// Exported (available for external packages to use)
func Save(u *User) error {

}

// Not exported. Only available inside the package
func checkAdmin(u *User) bool {
    return u.admin
}

```

Rules apply to data models:

```golang
package data

// Exported
type User struct {
    Name  string
    admin bool // not exported
}

// an importing package can see:
user.Name

// but NOT
user.admin

```

---
[Go Module](gomod.md)