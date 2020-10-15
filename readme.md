## To install:
```
go build
go install
```

## To use:
```
gow
```

Will watch all files in the current directory, and run 'go build' when they change.
Uses coroutines for the watching, and spawning of restarts. 
Stops server before build as separate coroutine.