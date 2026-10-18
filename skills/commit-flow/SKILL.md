---
name: commit-flow
description: Commit workflow that reviews current code changes, stages only necessary files, asks before assuming anything should be staged, suggests a commit message based on the diff, runs doc-gen before commit, and waits for user approval.
---

# Commit Flow

Use when the user asks to commit, stage, or prepare changes for a commit.

## Steps

1. **Check current changes** — run `git status` and `git diff` (or `git diff --cached`) to see exactly what changed.
2. **Review code** — read modified/added files to understand what is actually needed.
3. **Stage necessary code only** — add files that contain real, needed changes.
4. **Ask before assuming** — if a file looks like it might not need staging (e.g., temporary edits, unrelated changes, or files you are unsure about), ask the user first. Do not assume something should be staged.
5. **Base on code** — decisions must come from reading the actual diff/content, not guesses.
6. **Suggest commit message** — based on the actual code changes, propose a concise commit message (e.g., `feat: add X`, `fix: correct Y`).
7. **Run doc-gen** — before committing, activate/run the `gen-docs` skill to update documentation if needed.
8. **Confirm with user** — show the proposed message, doc-gen result, and staged files; ask for approval before running `git commit`.

## Rules

- Never stage files without checking them first.
- If unsure whether a change is needed, ask: "Should I stage `<file>`? It contains `<brief summary>`"
- Before committing, run `gen-docs` (or ask to run it) to update docs based on code changes.
- Suggest a commit message derived from the actual diff.
- Wait for explicit user approval (`yes`/`approve`) before running `git commit`.
