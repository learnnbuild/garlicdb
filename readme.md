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