# Encoding JSON with Go Exercise
Let's solve a real world problem with Go to learn a bit more about the language.

The rules:
- no external packages
- make no changes to `programmer_test.go`
- you can change anything in `programmer.go` (including the Programmer struct), but you do need a Programmer struct
- after cloning down the repo, just run `go test` to run tests

## The Problem
You need to send and receive JSON data representing a programmer.  It looks like this:

```JSON
{
    "name": "Jon Dancy",
    "birthday": "7/25/1986",
    "favoriteLanguage": "Go",
    "lastScore": 5
}
```

- name: It's just a simple string.  Don't overthink this one.
- birthday: The day this programmer was born.  Must be in this format when sent/received.
- favoriteLanguage: Valid options are "Go", "JavaScript", and "CSharp".  Sorry functional programmers.  When receiving data, case should not mater (ie. "JAvAScrIPT" is valid), but it should normalize when being sent.
- lastScore: Integer, 0-5.  If a programmer hasn't done anyhting yet, `null` is valid.

## Additional Info
- the output of the failing tests should be helpful 
- feel free to look at the code in the test file (although I'm not sure how usefull it will be)
- this is not meant to be super tricky, ask for help if you get stuck
- a solution to this that doesn't go out of it's way to be clever will be about 90 lines of code