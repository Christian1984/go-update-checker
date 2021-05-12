# Version Checker for Go

go-update-checker is a go library for checking the version of a currently installed application or package against its latest release on github. It also enables caching and setting a minimum interval of days after which a updatecheck against the github API should be performed to prevent spamming the API.

Versions used with go-update-checker must follow [SemVer](http://semver.org/).

## Installation and Usage

Installation can be done with a normal `go get`:

```
$ go get github.com/Christian1984/go-update-checker
```

#### Update Check Example

```go
import (
    "fmt"

    updatechecker "github.com/Christian1984/go-update-checker"
)

func main() {
    uc := updatechecker.New("Christian1984", "go-update-checker", "Go Update Checker", "", 0, true, false)
    uc.CheckForUpdate("0.0.1")

    /*
    =============================================================
    === INFO: A new update is available for Go Update Checker ===

    Version: 0.1.0

    Title: Go Update Checker - 0.1.0 - Initial Release

    Description:
    Initial Release


    Download the latest version here:
    https://github.com/Christian1984/vfrmap-for-vr/releases
    =============================================================
    */

    uc.CheckForUpdate("1.0.0")
    /*
    ========================================================================
    === INFO: You are running the latestest Version of Go Update Checker ===
    ========================================================================
    */
}
```

## Issues and Contributing

If you find an issue with this library, please report an issue. If you'd like, we welcome any contributions. Fork this library and submit a pull request.
