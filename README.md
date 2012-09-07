daycount
========

Command **daycount** produces information about days in a variety of
formats.

Package **daycount** provides types and functions for counting days.

Package **daycount/engines** supports a variety of presentation formats
for the results of daycount queries.

Note that as this project is in early development, anything might change
from version to version.

Installation
------------

To get fresh code and build your own binaries, you have to install:

* [Go](http://golang.org)
* [Git](http://git-scm.com)
* [Mercurial](http://mercurial.selenic.com)

Once that's done, this simple command should be enough to get daycount
up and running:

    go get github.com/jtacoma/daycount

Alternatively, binary distributions without the above prerequisites may
be(come) available.

Usage
-----

These examples don't actually work or do anything useful yet, but show
more or less what is intended.

Command:

    daycount -d 2011-01-01 -c 356 -f pdf > 2011.pdf

Code:

    import (
        "github.com/jtacoma/daycount/daycount"
        "fmt"
    )
    func main()
    {
        day := daycount.Parse("2012-12-21")
        fmt.Println(day.MayanLong().String())
    }

Output:

    13.0.0.0.0

