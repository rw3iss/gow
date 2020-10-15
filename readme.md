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


## File structure / explanation:

 - main.go - Simple entry point, runs initial build and starts watcher.
 - lib/builder.go - Manages the process of running 'go build'.
 - lib/watcher.go - Manages the process of watching for file changes, and runs restart routine.
 - lib/server.go  - Manages the process of running the the actual target executable/project. 
 - lib/Log.go - helper to print outputs.
 - lib/Config.go - helper to read in Config file.