Print Number (Golang)
------

Explanation:
---
input is string, so we need sanitize input and only number char only, so I remove all dot chat. 

After that, loop every char on and concat with 0 that repeat to len of input, minus index and plus 1, cause index start from 0.

Requirement:
> Golang, work well with version 1.12

How to Run?
---
```
go run main.go
```