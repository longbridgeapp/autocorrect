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

  autospace.Format("Ruby 2.7版本第1次发布")
  // => "Ruby 2.7 版本第 1 次发布"
}
```

## License

This project under MIT license.
