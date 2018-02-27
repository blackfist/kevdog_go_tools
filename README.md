# Kevin's Go Tools
These are simple utilities Kevin wrote to make his life easier

# Installation

```
go install github.com/blackfist/kevdog_go_tools
```


# The utilities
## random_sample
Takes a newline separated list from standard input and returns the list in a random 
order. The option -n option can be applied to select a random subset of what was input.

```
pbpaste | random_sample
words
this
is
of
list
a
```

```
pbpaste | random_sample -n 3
this
a
list
```

## splunk_list
Takes a newline separated list from standard input and changes it to a list formatted
for splunk. You can provide a prefix so that each item in your list is prefixed.

```
pbpaste | splunk_list -prefix word=
(word=this OR word=is OR word=a OR word=list OR word=of OR word=words)
```

## sql_list
Takes a newline separated list from standard input and changes it into a list formatted
for a SQL query

```
pbpaste | sql_list
('this','is','a','list','of','words')
```