## What this is:

- Go file watcher, builder, and runner. Barebones (fast), cross-platform (uses fsnotify), recursive, and configurable.
- Organized architecture, as an experimental project for learning Go.


## To install:
```
go get github.com/rw3iss/gow
```

### Or from source:
```
git clone https://github.com/rw3iss/gow.git
cd gow
go install
```

## To use:
```
gow
```

Will watch all files in the current directory, recursively, and run 'go build' when there is any change.


## File structure / explanation:

 - main.go - Simple entry point, runs initial build and starts watcher.
 - lib/Watcher.go - Manages the process of watching for file changes, and runs restart routine.
 - lib/Builder.go - Manages the process of running 'go build'.
 - lib/Runner.go  - Manages the process of running the the actual target executable/project. 
 - lib/Config.go  - Helper to read in Config file.
 - lib/StatusWriter.go - A generic Writer which accepts a format string and will print a formatted output.

 ### Utilities:

 - lib/Colors.go - Color definitions.
 - lib/Log.go - Helper to print outputs.

 ## Todo
 - implement Config to accept a custom 'watchDir'.
 - implement Config to accept a custom 'go build' command.
 - implement Config to accept a custom Runner command (currently just runs the package/folder name as the executable).
 - implement Config to specify a regex of file extensions to react to.
 - Update Config to parse incoming key indexes with recursive ability through period notation, ie. Get("Some.Child.Key").s