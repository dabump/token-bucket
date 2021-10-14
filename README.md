# Simple Token Bucket
Simple token bucket implementation used for simple rate limiting

The algorithm will fill the bucket back to it's initial size every second

Based on Token Bucket wiki: [wikipedia](https://en.wikipedia.org/wiki/Token_bucket)

Minimum requirements:
- Go 1.17

## How to use in client code
### Create a new bucket configuration based on 5 tokens per second
#### Usage example:
```go
bucket := tokenbucket.NewBucket("MyBucket", 5)
```

### Create and start daemon to manage the lifecycle of the bucket
#### Usage example:
```go
daemon := tokenbucket.NewDaemon(b, flags)
daemon.Start()
```

### Request a token
Invoking the `Hit()` function on the daemon will return either `true`/`false` 
indicating if a token was redeemed or not (ie, the bucket was empty will return false)

### To stop the daemon from filling and managing the bucket
#### Usage example:
```go
daemon := tokenbucket.NewDaemon(b, flags)
daemon.Start()
// code ...
daemon.Stop()
```

## Daemon flags
`Retriable` This flag will enable a random wait between 0-5 seconds before retying a failed token fetch

`Forgiving` If failed to retrieve a token, forgiving will return true if the last token was successful

#### Usage example:
```go
flags := Forgiving | Retryable
w := NewDaemon(b, flags)
w.Start()
```