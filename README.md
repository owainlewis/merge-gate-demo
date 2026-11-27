# Merge Gate Demo

A small Go service used to demonstrate automated code review, merge policies, and deterministic tripwires in [Machinist](https://github.com/owainlewis/machinist).

## Structure

- `pkg/stats/` — Pure utility helpers (two-way doors: easy to test and undo with a single `git revert`)
- `pkg/pricing/` — Subscription plans and billing calculations (one-way doors: money and pricing)
- `pkg/auth/` — Bearer token parsing and HMAC verification (one-way doors: security and identity)
- `db/migrations/` — SQL schema migrations (one-way doors: data safety and deterministic tripwire)
- `.machinist/merge-policy.md` — Written merge policy read from `origin/HEAD` by Machinist's `merge-gate` step

## Running checks

```sh
go test -v ./...
go vet ./...
```
