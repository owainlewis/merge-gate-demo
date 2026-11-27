# Balanced Merge Policy

Code quality has already been checked by the review step, and CI runs before merge.
Judge **blast radius and reversibility**: can a single `git revert` undo this cleanly?
When in doubt, choose `human_review`.

Auto-merge (`decision: "approve"`, `risk: "low"`) normal, reversible work:
- Scoped features, utility helpers, bug fixes, and internal refactors with unit tests
- Unit or integration tests, documentation, copy tweaks, and routine maintenance
- Pull requests under ~400 changed lines that match the requested task

Require human review (`decision: "human_review"`, `risk: "medium"` or `"high"`) for:
- **Data & schema:** database migrations, schema changes, or destructive data operations (`db/migrations/`)
- **Security & identity:** authentication, authorization, sessions, tokens, crypto, or secrets (`pkg/auth/`)
- **Money & external contracts:** billing, payments, pricing, or breaking public APIs (`pkg/pricing/`)
- **Release & supply chain:** CI/CD workflows (`.github/workflows/`), merge policy changes, deploy configs, or new/major dependencies
- **Size & integrity:** large diffs (over ~400 lines), disabling/skipping/deleting tests, or sweeping unrelated edits
