```
                    __
    ____  ___  ____/ /
   / __ \/ _ \/ __  / 
  / /_/ /  __/ /_/ /  
 / .___/\___/\__,_/   
/_/                                  
```

# ped - Percent encode decode

URL encode and decode tool written in Go.

## Installation

```
go build main.go && mv main ped
```

The only dependency is Go >1.0. No other dependencies needed.

## Example usage

Encode string

```
❯ ped --encode "/path/you/want/encoded?with=value"
%2Fpath%2Fyou%2Fwant%2Fencoded%3Fwith%3Dvalue
```

Decode string

```
❯ ped --decode "%2Fpath%2Fyou%2Fwant%2Fencoded%3Fwith%3Dvalue"
/path/you/want/encoded?with=value
```

## License

ped is released under MIT license. See [LICENSE](https://github.com/mskri/ped/blob/master/LICENSE).
