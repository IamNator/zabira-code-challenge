# zabira-code-challenge

### How to run

First Clone the repo with
```
    $ git clone git@github.com:IamNator/zabira-code-challenge.git
```
Then
```
    $ cd zabira-code-challenge/
```
Next, run the program
```
    $ go run main.go
```

### Testing
To run tests
```
    $ go test ./...
```

#### Running Bench Test
First navigate to the sort dir
```
    $ ch sort
```
Then, run bench test
```
    $ go test -bench=.
```

### Problem Statement

I was asked to create a sorting solution that is extensible so that different teams can easily add their own sorters without affecting existing code. 

<img alt="problem statement" src=".github/images/problem.png">


### My Solution

Using an OOP approach, as demonstrated in my implementation, can make the code more modular and easier to extend. By creating an interface for sorters and implementing that interface in different sorter classes/objects, the team can add new sorting algorithms without changing the existing code. This approach promotes code reuse and separates concerns, which can lead to a more maintainable and extensible codebase.

#### Directory Tree

I created a `sort` packge to implement objects that fulfill the ProductSorter interface

<img alt="directory tree" src=".github/images/dir.png">

#### `sort` Package

The sort package describes the ProductSorter interface, it also contains implementation for the interface

<img alt="product sorter interface" src=".github/images/sorter.png">

##### Sort by `Price`

