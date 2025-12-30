# Compilare

Compilare is latin for *to compile*: exactly what compilare does in essence -- compile reads that I find beautiful from multiple sources, and display them in a unified structure.


# Setup

## Required:

- [go](https://go.dev/doc/install) `>=1.25`

## Local

```bash
git clone https://github.com/suynep/compilare
cd compilare/
go run .
```

> The API runs at `localhost:61666` (an arbitrary port choice: perhaps an acknowledgement to the devil within)


### Todo
- [x] implement test agnosticism :P
- [ ] implement expiring sessions 
- [ ] implement post-login checks to redirect user to their presently active session upon multiple login attempts
