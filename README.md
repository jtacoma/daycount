daycount
========

Command **daycount** produces information about days in a variety of formats.

Package **daycount** provides types and functions for counting days.

Package **daycount/engines** supports a variety of presentation formats for the results of daycount queries.

Note that as this project is in early development, anything might change from version to version.

Installation
------------

    go install github.com/jtacoma/daycount

Usage
-----

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

