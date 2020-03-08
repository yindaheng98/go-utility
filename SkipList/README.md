# github.com/yindaheng98/go-utility/Skiplist

## Introduction

The skip list is a probabilisitc data structure that is built upon the general idea of a linked list. Here is [Wiki about skip list](https://brilliant.org/wiki/skip-lists/). The random index level generation algorithm in this skip list is based on [this article](https://yindaheng98.github.io/%E6%95%B0%E5%AD%A6/SkipList.html).

## Usage

There are only two important struct in this package. The struct `Skiplist` stands for a skip list, and the struct `Node` stands for the pointer to a node in the `Skiplist`.

The struct `RandLevel` is the random index level generator. It is used by `Skiplist` as an unexported member. It is not necessary for you to care about it.

### Create a `Skiplist`

You can create a `Skiplist` by assigning the size of the skip list with the maximum index level or the index decade factor ([What's this?](https://yindaheng98.github.io/%E6%95%B0%E5%AD%A6/SkipList.html)).

If you want to assign the maximum index level, you can use:

```go
skiplist := NewWithLevel(listSize, indexLevel)
```

Or if want to assign the index decade factor, you can use:

```go
skiplist := NewWithC(listSize, C)
```

### Insert a node

If you want to insert a node with the assigned value (for example, 4.24) into the `skiplist`, just:

```go
node := skiplist.Insert(4.24)
```

This function will return a `*Node`, pointing to where the value exists.

### Find a node

The skiplist can find the pointer of the largest node whose value is smaller than a specified value. For example, if there are 3 node [1.0 , 2.0 , 3.0] in the `skiplist`, you can find the largest node smaller than 2.1:

```go
node1 := skiplist.Insert(2.1)
```

`node1` will be the pointer of the node 2.0.

If a node larger than your specific value is not exists, `nil` will return. For example:

```go
node2 := skiplist.Insert(1.0)
```

Since all the node are no more than 1.0, `node2` will be `nil`.

### Delete a node

Before delete, you should know the pointer of the node you want to delete. Delete a node is easy:

```go
skiplist.Delete(node_pointer_you_want_to_delete)
```

Be careful: This function will not check if the node is in the `skiplist`, so you must be confident that the node that is going to be deleted is in the `skiplist`.

### Change the value in a node

If you want to increase the value of a node:

```go
node_after := skiplist.Delta(node_before, 1.0)//increase by 1.0
```

Or decrease:

```go
node_after := skiplist.Delta(node_before, -1.0)//decrease by 1.0
```

Increase or decrease will effect the order of the nodes, so the mechanism of the methods `Delta` is just delete the node, insert a new value and return the new pointer.

If you want to change all the node's value:

```go
skiplist.Delta(1.0)//increase by 1.0
skiplist.Delta(-1.0)//decrease by 1.0
```

Increase or decrease by same value can not effect the order of the nodes.

### Count the nodes

If you want to know how much nodes are in your skip list:

```go
how_much_nodes_here := skiplist.Count()
```

### Traversal

This methods will return a list of all the node in `skiplist` in ascending order.

```go
nodes := skiplist.TraversalAll()
```

Or if you do not want all the nodes, you can get the smallest `n`(`uint64`) values in the list:

```go
nodes := skiplist.Traversal(n)
```