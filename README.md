go-daycount
===========

Package **daycount** comes with an executable named **go-daycount**.

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
        fmt.Printf(day.MayanLong)
    }

Output:

    13.0.0.0.0

