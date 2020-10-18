## What this is:

- Go file watcher, builder, and runner. 
- sBarebones (fast), cross-platform (uses fsnotify), recursive, and configurable.
- Organized architecture, as an experimental project for practicing Go.


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
```
main.go                 - Entry point, runs initial build and starts watcher.
lib/                    - (Underlying application library)
    Watcher.go          - Manages process of watching files, and restart routine.
    Builder.go          - Manages process of running build command.
    Runner.go           - Manages process of running actual target executable. 
    Config.go           - Helper to read in configuration file.
    utils/              - (Library utilites)
        Colors.go       - Color definitions.
        FormatWriter.go - Generic Writer, accepts format string and prints output.
        Log.go          - Helper to print outputs.
```

 ## Todo
 - implement Config to accept a custom 'watchDir'.
 - implement Config to accept a custom 'go build' command.
 - implement Config to accept a custom Runner command (currently just runs the package/folder name as the executable).
 - implement Config to specify a regex of file extensions to react to.
 - Update Config to parse incoming key indexes with recursive ability through period notation, ie. Get("Some.Child.Key").