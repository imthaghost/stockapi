<p align="center">
  <!-- <a href="#">
    <img alt="jedi" src="docs/media/luke.png"> 
  </a> -->
</p>

<br>

<p align="center">
A Go Restful API that scrapes real-time stock data from Yahoo Finance.
</p>

<br>

<p align="center">
   <a href="https://goreportcard.com/report/github.com/imthaghost/stockapi"><img src="https://goreportcard.com/badge/github.com/imthaghost/stockapi"></a>
   <a href="https://github.com/imthaghost/gitmoji-changelog">
    <img src="https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg"alt="gitmoji-changelog">
  </a>
</p>
<br>

<!-- ![Example](/docs/media/teslaexample.gif) -->

## Table of Contents

-   [Installation](#installation)
-   [Examples](#examples)
-   [License](#license)
-   [Contributors](#contributors)

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

## Contributors

Contributions are welcome! Please see [Contributing Guide](https://imthaghost/zeus) for more details.

<table>
  <tr>
    <td align="center"><a href="https://github.com/imthaghost"><img src="https://avatars3.githubusercontent.com/u/46610773?s=460&v=4" width="75px;" alt="Gary Frederick"/><br /><sub><b>Tha Ghost</b></sub></a><br /><a href="https://github.com/imthaghost/stockapi/commits?author=imthaghost" title="Code">💻</a></td>
  </tr>
