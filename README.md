# AutoCorrrect

Automatically add spaces between Chinese and English words.

## Other implements

- Ruby - [auto-correct](https://github.com/huacnlee/auto-correct).
- Go - [go-auto-correct](https://github.com/huacnlee/go-auto-correct).
- Rust - [auto-correct.rs](https://github.com/huacnlee/auto-correct.rs).

## Features

- Auto add spacings between Chinese and English words.
- HTML content support.

## Usage

Use `autocorrect.Format` to format plain text.

https://play.golang.org/p/ntVhrGYnxNk

```go
package main

import "gthub.com/huacnlee/go-auto-correct"

func main() {
  autocorrect.Format("长桥LongBridge App下载")
  // => "长桥 LongBridge App 下载"

  autocorrect.Format("Ruby 2.7版本第1次发布")
  // => "Ruby 2.7 版本第 1 次发布"

  autocorrect.Format("于3月10日开始")
  // => "于 3 月 10 日开始"

  autocorrect.Format("包装日期为2013年3月10日")
  // => "包装日期为2013年3月10日"
}
```

Use `autocorrect.FormatHTML` for HTML contents.

https://play.golang.org/p/fTGVlkM-H3W

```go
package main

import "gthub.com/huacnlee/go-auto-correct"

func main() {
  autocorrect.Format(htmlBody)
  // => "<div><p>长桥 LongBridge App 下载</p><p>最新版本 1.0</p></div>"
}
```

## Benchmark

### Format

| Total chars | Duration |
| ----- | ------- |
| 50  | 0.09 ms |
| 100  | 0.14 ms |
| 400  | 0.39 ms |

### FormatHTML

| Total chars | Duration |
| ----- | ------- |
| 2K  | 1.4 ms |

## License

This project under MIT license.
