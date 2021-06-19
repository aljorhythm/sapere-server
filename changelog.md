# Interfaces, Code Structure, Modelling and Service Tests

## Interfaces

- In Go, callers instead of callees specify interfaces
- Clear separation of Data, Business layers
    - Data has it's view of a `User`, Business has it's view of a `User`

### Cons

- Super explicit, from a non-Go style perspective there's a lot of duplicate code

### Pros

- Somewhat like domain driven design but applied vertically, this design forces each layer to be explicit about what it assumes.

## Service Tests

- Decided to go with `golang` for integrations
  
### Cons

- Static typing is not very useful for asserting service level json responses
- Sort of in the same location as source - bad practices can contaminate test code with assumptions from source code
- verbose

### Pros

- Strongly typing is useful in the future for testing gRPC/GraphQL endpoints
- Although strongly and statically typed, liberal use of maps is allowed
- Easy to write, maybe some strong typing can help
- No need to manage another framework/language in CI, can use same tools for linting etc..
- http is already a package
- also supports concurrency which would be a problem in `python`

## Conservative use of pointers

- todo link

## References

- https://stackoverflow.com/questions/44864790/why-cant-i-use-a-pointer-to-a-specific-type-where-interface-is-expected
- https://stackoverflow.com/questions/30652577/go-doing-a-get-request-and-building-the-querystring

# Readme

add writing and docs
sneak 
- minor change for github action branch matcher to include all branches except `main`
- checkout step for main branch

# Dockerize

- add this changelog file
- add start server log
- add makefile build, test, run and containerize steps
- add github action
- publish image to dockerhub

## Practices, Principles and Philosophy

see ppp.md

## External

- create aljorhythm/sapere-server docker repository

## Not included

- service test framework
- database
- deployment
- release github action steps

## References

- https://blog.golang.org/docker
- https://docs.docker.com/docker-hub/builds/
- https://stackoverflow.com/questions/2826029/passing-additional-variables-from-command-line-to-make
- https://www.shellscript.sh/exitcodes.html
- https://github.com/mvdan/github-actions-golang
- https://docs.docker.com/docker-hub/repos/