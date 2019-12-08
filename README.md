# brainfuck

[![](https://github.com/shilangyu/brainfuck/workflows/ci/badge.svg)](https://github.com/shilangyu/brainfuck/actions)

Extremely primitive brainfuck interpreter

Flavor:

- 8 bitwidth
- 30000 cells
- wrapping cells
- 8 standard commands [+-,.><]

### install

Grab an executable from the [release tab](https://github.com/shilangyu/brainfuck/releases)

Or from source:

```sh
go get github.com/shilangyu/brainfuck
```

### use

```sh
brainfuck ./input.bf
```

```sh
brainfuck -c "+[-[<<[+[--->]-[<<<]]]>>>-]>-.---.>..>.<<<<-.<+.>>>>>.>.<<.<-."
```
