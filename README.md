# autospace

Automatically add spaces between Chinese and English words.

This is a go version of [auto-correct](https://github.com/huacnlee/auto-correct).

## Usage

```go
package main

import "gthub.com/huacnlee/go-autospace"

func main() {
  autospace.Format("长桥LongBridge App下载")
  // => "长桥 LongBridge App 下载"
}
```

## License

This project under MIT license.
