# Compilare

Compilare is latin for *to compile*: exactly what compilare does in essence -- compile reads that I find beautiful from multiple sources, and display them in a unified structure.



### Todo
- `manager/manager.go` handles routine fetches from the three available sources *viz Aeon, Psyche, and Hackernews*
    - presently, the `CheckAndSaveLastRunTime` function uses the functions from the `test` package *(which were written for unit-tests)*
    - TODO: implement test agnosticism :P
