## To install:
```
go build
go install
```

## To use:
```
gow
```

Will watch all files in the current directory, recursively, and run 'go build' when there is any change.


## File structure / explanation:

 - main.go - Simple entry point, runs initial build and starts watcher.
 - lib/Builder.go - Manages the process of running 'go build'.
 - lib/Watcher.go - Manages the process of watching for file changes, and runs restart routine.
 - lib/Server.go  - Manages the process of running the the actual target executable/project. 

 ### Utilities:

 - lib/Colors.go - Color definitions.
 - lib/Log.go - Helper to print outputs.
 - lib/Config.go - Helper to read in Config file.