# Golang command parser [WIP]
Command parser inspired by discord.py system.

# Usage
```go
parser.ParseCommand(*command*)
```

# Examples (WIP)
```go
parser.ParseCommand("!help") // -> [ "!help" ]
```

```go
parser.ParseCommand("!help 1 2 3") // -> [ "!help", "1", "2", "3" ]
```

```go
parser.ParseCommand("!help \"Hello, Mike!\"") // -> [ "!help", "Hello, Mike!" ]
```

```go
parser.ParseCommand("!help    1") // -> [ "!help", "1" ]
```

# License
Apache 2.0
