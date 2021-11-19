# AutoCorrrect for Go

[![Go](https://github.com/longbridgeapp/autocorrect/workflows/Go/badge.svg)](https://github.com/longbridgeapp/autocorrect/actions?query=workflow%3AGo)

Automatically add whitespace between CJK (Chinese, Japanese, Korean) and half-width characters (alphabetical letters, numerical digits and symbols).

## Other implements

- Rust - [autocorrect](https://github.com/huacnlee/autocorrect).
- Ruby - [auto-correct](https://github.com/huacnlee/auto-correct).

## Features

- Auto add spacings between CJK (Chinese, Japanese, Korean) and English words.
- HTML content support.
- Fullwidth -> halfwidth (only for [a-zA-Z0-9], and `：` in time).
- Correct punctuations into Fullwidth near the CJK.
- Cleanup spacings.
- Support options for custom format, unformat.

## Usage

```
go get github.com/longbridgeapp/autocorrect
```

Use `autocorrect.Format` to format plain text.

https://play.golang.org/p/ntVhrGYnxNk

```go
package main

import "github.com/longbridgeapp/autocorrect"

func main() {
  autocorrect.Format("长桥LongBridge App下载")
  // => "长桥 LongBridge App 下载"

  autocorrect.Format("Ruby 2.7版本第1次发布")
  // => "Ruby 2.7 版本第 1 次发布"

  autocorrect.Format("于3月10日开始")
  // => "于 3 月 10 日开始"

  autocorrect.Format("包装日期为2013年3月10日")
  // => "包装日期为 2013 年 3 月 10 日"

  autocorrect.Format("生产环境中使用Go")
  # => "生产环境中使用 Go"

  autocorrect.Format("本番環境でGoを使用する")
  # => "本番環境で Go を使用する"

  autocorrect.Format("프로덕션환경에서Go사용")
  # => "프로덕션환경에서 Go 사용"

  autocorrect.Format("需要符号?自动转换全角字符、数字:我们将在１６：３２分出发去ＣＢＤ中心.")
  # => "需要符号？自动转换全角字符、数字：我们将在 16:32 分出发去 CBD 中心。"
}
```

With custom formatter:

```go
type myFormatter struct {}
func (my myFormatter) Format(text string) string {
  return strings.ReplaceAll(text, "ios", "iOS")
}

autocorrect.Format("新版本ios即将发布", myFormatter{})
// "新版本 iOS 即将发布"
autocorrect.FormatHTML("<p>新版本ios即将发布</p>", myFormatter{})
// "<p>新版本 iOS 即将发布</p>"
```

Use `autocorrect.Unformat` to cleanup spacings in plain text.

```go
package main

import "github.com/longbridgeapp/autocorrect"

func main() {
  autocorrect.Unformat("据港交所最新权益披露资料显示，2019 年 12 月 27 日，三生制药获 JP Morgan Chase & Co.每股均价 9.582 港元，增持 270.3 万股，总价约 2590 万港元。")
  // => "据港交所最新权益披露资料显示，2019年12月27日，三生制药获JP Morgan Chase & Co.每股均价9.582港元，增持270.3万股，总价约2590万港元。"
}
```

Use `autocorrect.FormatHTML` / `autocorrect.UnformatHTML` for HTML contents.

https://play.golang.org/p/pbETBF4OOcj

```go
package main

import "github.com/longbridgeapp/autocorrect"

func main() {
  autocorrect.FormatHTML(htmlBody)
  // => "<div><p>长桥 LongBridge App 下载</p><p>最新版本 1.0</p></div>"
  autocorrect.UnformatHTML(htmlBody)
  // => "<div><p>长桥LongBridge App下载</p><p>最新版本1.0</p></div>"
}
```

## Benchmark

Run `go test -bench=.` to benchmark.

```
pkg: github.com/longbridgeapp/autocorrect
BenchmarkFormat50-12      	   19671	     60175 ns/op
BenchmarkFormat100-12     	   10000	    119076 ns/op
BenchmarkFormat400-12     	    2847	    424984 ns/op
Benchmark_halfwidth-12    	  289411	      4150 ns/op
BenchmarkFormatHTML-12    	    1100	   1097027 ns/op
```

### Format

| Total chars | Duration |
| ----------- | -------- |
| 50          | 0.06 ms  |
| 100         | 0.11 ms  |
| 400         | 0.42 ms  |

### FormatHTML

| Total chars | Duration |
| ----------- | -------- |
| 2K          | 1.09 ms  |

## License

This project under MIT license.
