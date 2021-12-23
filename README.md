# gktool

Cli tool to generate secure key

⚠ **_It's new version of [gpwd](https://github.com/Ulrichinho/gpwd) except that It doesn't call API because it integrates an algorithme of generate of keys and some functions to ckeck strength of generation_**

## Install

**_Golang 1.17 must be install to add gktool cli_**

Go ahead compile it yourself :

```sh
go install github.com/Ulrichinho/gktool@latest
```

**_It's compatible with macOs, linux and Windows_**

## Usage

```text
gktool [global options]
```

### Global options

```text
GLOBAL OPTIONS:
   --length value, -l value    define length of key (default: 16)
   --quantity value, -q value  define quantity of key (default: 1)
   --no-upper                  define if you don't want upper chars (default: false)
   --no-lower                  define if you don't want lower chars (default: false)
   --no-number                 define if you don't want number chars (default: false)
   --no-symbol                 define if you don't want symbol chars (default: false)
   --export, -e                export generate key(s) in file (default: false)
   --help, -h                  show help (default: false)
   --version, -V               print only the version (default: false)

```

## Version

To see version :

```text
gktool [-V | --version]
```

see [`CHANGELOG`](./CHANGELOG.md) for more details about versions

## Help

To have more help about this command type :

```text
gktool [-h | --help]
```

## License

The Apache License (APACHE) - see [`LICENSE`](./LICENSE.md) for more details
