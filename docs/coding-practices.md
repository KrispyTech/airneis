# Coding Practices

Author : Cédric Gautier
Date : 2024-01-01

## SUMMARY:

Have a standard on coding when we build our application.

## CONTEXT:

There are multiple layers when creating a software in which there's a team of multiple developers.

Everyone has a coding style, their own way to write software but we would like to have some practices that we should respect to be able to have omni-present default practice.

## SOLUTION:

### Linting

First of all, there is a linting standard that we'd like to respect. We'll have one for both parts : backend, frontend and mobile.

#### Back-End

In the back-end we'll be using golang-ci-lint that has a respected standard. We'll be using the default linting that it proposes but this may be subject to be modified in the futur.

- [errcheck](https://github.com/kisielk/errcheck) : Errcheck is a program for checking for unchecked errors in Go code. These unchecked errors can be critical bugs in some cases.
- [gosimple ⚙️](https://github.com/dominikh/go-tools/tree/master/simple) : Linter for Go source code that specializes in simplifying code.
- [govet ⚙️](https://pkg.go.dev/cmd/vet) : Vet examines Go source code and reports suspicious constructs, such as Printf calls whose arguments do not align with the format string.
- [ineffassign](https://golangci-lint.run/usage/linters/#:~:text=v1.0.0-,ineffassign,-Detects%20when%20assignments) : Detects when assignments to existing variables are not used.
- [staticcheck ⚙️](https://golangci-lint.run/usage/linters/#:~:text=v1.0.0-,staticcheck,-%E2%9A%99%EF%B8%8F) : It's a set of rules from staticcheck. It's not the same thing as the staticcheck binary. The author of staticcheck doesn't support or approve the use of staticcheck as a library inside golangci-lint.
- [unused ⚙️](https://github.com/dominikh/go-tools/tree/master/unused) : Checks Go code for unused constants, variables, functions and types.

### Case type & practices

#### Global

- Follow `camelCase` coding standard.
- Always re-use, upgrade, or improve the existing code. The core can be modified but make sure to refactor affected code and always think how the code can be scalable.
- Always start the function name with an action. (e.g: `getSecretKey`, `initializeGithubClient`)

#### Back-End

- If needed for constants, declare them like and where it's used :

```go
const (
	stagingEnv    string = "staging"
	productionEnv string = "production"
)
```

- If there are more than 2 possibilities of a result, use a switch case. Having a default value is also recommended if you're awaiting a default task.

```go
switch song.Name {
	case "Talk 2 Me":
		fmt.Sprintf("Best moddy song is %s", song.Name)

	case "Like Ice Spice":
		fmt.Sprintf("Best afrobeats song is %s", song.Name)
	}

  default:
    fmt.Sprintf("Best song ever is %s", song.Name)
```

- When there are more than 3 data types to return. Make a struct to return your data, you should do something like this:

```go
package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type Song struct {
	Name         string
	Artist       string
	Album        string
	Genre        string
	AmountPlayed int
}

func GetSong(name string) (Song, error) {
	return Song{
		Name: "Pink + White",
		Artist: "Frank Ocean",
		Album: "Blond",
		Genre: "R&B",
		AmountPlayed: 1000000000,
	}, nil
}

func main() {
	song, err := getSong()
	if err != nil {
		return errors.Wrapf(err, "Main, Unable to getSong")
	}
	fmt.Println(song.Name)

}
```

### Code reusability

#### Back-End

**When coding :**

- Make public functions in Go only when needed. (Start of function should be lowercased to make it private)

### Error Handling

#### Back-End

We'll be using the errors package in Go to be able to get clearer information when debugging our code.

There are two types of errors that we should use.

- When you're inside a child function, always return

```go
return errors.Errorf("Function name, unable to do whatever blabla ... %s", err)
```

- When you're the parent functions, always return wrap errors with a context.

Context can be : "Unable to marshall, unable to connect, unable to find, ..."

```go
return errors.Wrapf(err, "parentFunction, context %s", childVariable)
```
