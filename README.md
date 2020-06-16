# AutoCorrrect for Go

[![Build Status](https://travis-ci.org/huacnlee/go-auto-correct.svg?branch=master)](https://travis-ci.org/huacnlee/go-auto-correct)

Automatically add whitespace between CJK (Chinese, Japanese, Korean) and half-width characters (alphabetical letters, numerical digits and symbols).

## Other implements

- Ruby - [auto-correct](https://github.com/huacnlee/auto-correct).
- Go - [go-auto-correct](https://github.com/huacnlee/go-auto-correct).
- Rust - [auto-correct.rs](https://github.com/huacnlee/auto-correct.rs).

## Features

- Auto add spacings between CJK (Chinese, Japanese, Korean) and English words.
- HTML content support.
- FullWidth -> HalfWidth (only for [a-zA-Z0-9], and `：` in time).

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

  autocorrect.Format("生产环境中使用Go")
  # => "生产环境中使用 Go"

  autocorrect.Format("本番環境でGoを使用する")
  # => "本番環境で Go を使用する"

  autocorrect.Format("프로덕션환경에서Go사용")
  # => "프로덕션환경에서 Go 사용"

  autocorrect.Format("自动转换全角字符、数字：我们将在１６：３２分出发去ＣＢＤ中心。")
  # => "自动转换全角字符、数字：我们将在 16:32 分出发去 CBD 中心。"
}
```

Use `autocorrect.FormatHTML` for HTML contents.

https://play.golang.org/p/fTGVlkM-H3W

```go
package main

import "github.com/huacnlee/go-auto-correct"

func main() {
  autocorrect.FormatHTML(htmlBody)
  // => "<div><p>长桥 LongBridge App 下载</p><p>最新版本 1.0</p></div>"
}
```

## Benchmark

Run `go test -bench=.` to benchmark.

```
pkg: github.com/huacnlee/go-auto-correct
BenchmarkFormat50-12      	   19671	     60175 ns/op
BenchmarkFormat100-12     	   10000	    119076 ns/op
BenchmarkFormat400-12     	    2847	    424984 ns/op
BenchmarkFormatHTML-12    	    1100	   1097027 ns/op
```

### Format

| Total chars | Duration |
| ----- | ------- |
| 50  | 0.06 ms |
| 100  | 0.11 ms |
| 400  | 0.42 ms |

### FormatHTML

| Total chars | Duration |
| ----- | ------- |
| 2K  | 1.09 ms |

## License

This project under MIT license.
