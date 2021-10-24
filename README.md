# jobtracker - a self terminating concurrent job queue for indeterminate workloads in golang

This library is primarily useful for technically-recursive work with an unknown size at the start of the work. It was originally created for [goop](https://github.com/deletescape/goop) where i had a need for a concurrent job queue where jobs have the ability to queue new jobs, which terminates itself once all work is done (only indicated by the lack of new jobs). 

## Installation

```bash
go get -u github.com/deletescape/jobtracker@latest
```

## Examples

You can find some examples in the [_examples/](_examples/) directory.

## Projects using jobtracker

* [goop](https://github.com/deletescape/goop)