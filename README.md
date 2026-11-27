# Merge Gate Demo

A small Go service used to demonstrate automated code review, merge policies, and deterministic tripwires in [Machinist](https://github.com/owainlewis/machinist).

## Structure

- `pkg/stats/` — Pure utility helpers (two-way doors: easy to test and undo with a single `git revert`)
- `pkg/pricing/` — Subscription plans and billing calculations (one-way doors: money and pricing)
- `pkg/auth/` — Bearer token parsing and HMAC verification (one-way doors: security and identity)
- `db/migrations/` — SQL schema migrations (one-way doors: data safety and deterministic tripwire)
- `.machinist/merge-policy.md` — Written merge policy read from `origin/HEAD` by Machinist's `merge-gate` step

## Quick examples

Using `pkg/stats` helpers:

```go
package main

import (
	"fmt"

	"github.com/owainlewis/merge-gate-demo/pkg/stats"
)

func main() {
	sum := stats.Add(2, 3)                    // 5
	avg := stats.Average([]float64{2, 4, 6})  // 4
	empty := stats.Average(nil)               // 0

	fmt.Println(sum, avg, empty)
}
```

## Observe mode

When `github.auto_merge: false` is set in Machinist, the `merge-gate` step runs in **observe mode**: it still evaluates each pull request against `.machinist/merge-policy.md` and records its verdict (applying `merge-gate:observe` instead of `merge-gate:auto-merge` on passing low-risk changes), without enabling GitHub auto-merge. This lets you review gate decisions before allowing automatic merges.

## Running checks

```sh
go test -v ./...
go vet ./...
```
