go-daycount
===========

Package **daycounts** provides functions for counting days.

Command **daycount** produces information about days in a variety of formats.

Installation
------------

    go get github.com/jtacoma/go-daycount

Usage
-----

Code:

    import (
        "github.com/jtacoma/go-daycount/daycount"
        "fmt"
    )
    func main()
    {
        day := daycount.Parse("2012-12-21")
        fmt.Println(day.MayanLong().String())
    }

Output:

    13.0.0.0.0

