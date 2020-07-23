<p align="center">
A Go Restful API that scrapes real-time stock data from Yahoo Finance.
</p>

<br>

<p align="center">
   <a href="https://goreportcard.com/report/github.com/imthaghost/stockapi"><img src="https://goreportcard.com/badge/github.com/imthaghost/stockapi"></a>
</p>
<br>


## Table of Contents

-   [Installation](#installation)
-   [Examples](#examples)
-   [License](#license)


## 🚀 Quick Start

### Installation

```sh
$ go get https://github.com/imthaghost/stockapi
```

### Instantiate API

Create `server.go`

```go

package main

import (

"github.com/imthaghost/stockapi/server"

)

func main() {

    s := server.NewServer()
    s.Start(":8000")

}

```

Start server

```sh
$ go run server.go
```

## Examples

![Postman](/docs/media/postman.gif)

## 📝 License

By contributing, you agree that your contributions will be licensed under its MIT License.

In short, when you submit code changes, your submissions are understood to be under the same [MIT License](http://choosealicense.com/licenses/mit/) that covers the project. Feel free to contact the maintainers if that's a concern.

