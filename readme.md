# GarlicDB: KV Storage Engine

## Memory Design: BTree

Data in the memory are different in terms of types. We should define a interface to support this.


```go
type Indexer interface {
    Put()
    Get()
    Delete()
}
```


## Disk: IO Operations

Standard IO should be supported at least. We can define an IO manger interface.

```go
type IOManager interface {
    Read()
    Write()
    Sync()
    Close()
}
```

