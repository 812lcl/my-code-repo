package main

import (
	"812lcl/my-code-repo/concurrency"
	"812lcl/my-code-repo/panic"
	"812lcl/my-code-repo/sort"
	"812lcl/my-code-repo/structure"
)

func main() {
	structure.DataStructureFunc()
	panic.CrashNotRecover()
	concurrency.ConcurrencyFunc()
	sort.SortDemo()
}
