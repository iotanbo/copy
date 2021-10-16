# copy

Go package for advanced copying files and directories.
This is a modified version of https://github.com/otiai10/copy.

Overwrite modes are changed to following:
```go
// NoOverwrite does nothing if destination exists (default behavior).
NoOverwrite DestExistsAction = iota
// Merge leaves existing items intact and only copies items that do not exist in dest.
Merge
// OverwriteIntersection overwrites existing common items but leaves
// those unique to destination intact.
OverwriteIntersection
// OverwriteFull deletes destination and then copies source items.
OverwriteFull
```

# Example Usage

```go
err := Copy("your/directory", "your/directory.copy")
```

# Advanced Usage

```go
// Options specifies optional actions on copying.
type Options struct {

	// OnSymlink can specify what to do on symlink
	OnSymlink func(src string) SymlinkAction

	// OnDestExists can specify what to do when there is a directory already existing in destination.
	OnDestExists func(src, dest string) DestExistsAction

	// Skip can specify which files should be skipped
	Skip func(src string) (bool, error)

	// AddPermission to every entities,
	// NO MORE THAN 0777
	AddPermission os.FileMode

	// Sync file after copy.
	// Useful in case when file must be on the disk
	// (in case crash happens, for example),
	// at the expense of some performance penalty
	Sync bool

	// Preserve the atime and the mtime of the entries
	// On linux we can preserve only up to 1 millisecond accuracy
	PreserveTimes bool

}
```

```go
// For example...
import (
	// go get github.com/iotanbo/copy@v1.6.0-advcopy
	advcopy "github.com/iotanbo/copy"
)
opt := advcopy.Options{
	OnDestExists: func(src, dest string) advcopy.DestExistsAction {
		return advcopy.Merge
	},
}
err := advcopy.Copy("your/directory", "your/directory.copy", opt)
```
