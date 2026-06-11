# strings
Carve ASCII and Unicode strings from files.

```console
go install go.foxforensics.eu/strings@latest
```

## Usage
```console
$ strings [-nmtao] FILE
```

### Options
* `-n` Minimum string length (default 3)
* `-m` Maximum string length
* `-t` Trim spaces from ends
* `-a` Show ASCII strings only
* `-o` Show file offset

## Acknowledgements
The carving algorithm is based on the original [Strings](https://github.com/robpike/strings) by [Rob Pike](https://github.com/robpike).

## License
Released under the [MIT License](LICENSE.md).