# Creating Go Library/Module
```bash
# path should be the same as the github repo path
go mod init github.com/tlylt/golib
# package files can have different names, may not need to be the same as the repo name
touch golib.go

# quick release management

## tidy stuff
go mod tidy
## start with v0.1.0
git tag v0.1.0
git push origin v0.1.0
## check if the module is now available to the public
go list -m github.com/tlylt/golib@v0.1.0
```

# Refer to a local module
In the caller module:
```bash
# assume caller is in another folder besides golib folder
go mod edit -replace github.com/tlylt/golib=../golib

```
# References
- https://go.dev/blog/publishing-go-modules