## What this is:

- Go file watcher, builder, and runner. 
- Barebones (fast), cross-platform (uses fsnotify), recursive, and configurable.
- Organized architecture, as an experimental project for practicing Go.

![Alt text](https://i.imgur.com/xsDePqc.png "Output")

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
    Application.go      - Wrapping context for entire Application.
    Builder.go          - Manages process of running build command.
    Runner.go           - Manages process of running actual target executable. 
    Watcher.go          - Manages process of watching files, and restart routine.
    utils/              - (Library utilites)
        Colors.go       - Color definitions.
        Config.go       - Helper to read in configuration file.
        FormatWriter.go - Generic Writer, accepts format string and prints output.
        Log.go          - Helper to print outputs.
test/                   - (Test package, run 'gow' here and it will run test.go)
    test.go             - Simple test program to print/test automatic changes.
```

 ## Todo
 - implement Config to specify a regex of file extensions to react to.