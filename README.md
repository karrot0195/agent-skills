# agent-skills

Agent skills workspace with a sync tool and reusable agent skills.

## Overview

- `cmd/sync-skills` — Go CLI that copies `AGENT_SKILL_PATH` into a target directory (e.g., `.agents/skills`).
- `skills/` — reusable agent skills (e.g., `commit-flow`, `gen-docs`).
- `Makefile` — builds `bin/agent-synch-skills`; targets `synch-local-skills` (uses `AGENT_SKILL_PATH=skills`) and `sync-target-skills` (requires `TARGET=`).

## Directory structure

```
agent-skills/
├── Makefile
├── bin/
│   └── agent-synch-skills
├── cmd/
│   └── sync-skills/
│       └── main.go
├── skills/
│   ├── commit-flow/
│   │   └── SKILL.md
│   └── gen-docs/
│       └── SKILL.md
└── .gitignore
```

## Packages / binaries

- `cmd/sync-skills` — binary `agent-synch-skills`. Usage: `AGENT_SKILL_PATH=skills bin/agent-synch-skills -Target=<dir>`

## Dependencies / requirements

- Go (legacy `GO111MODULE=off` build)
- Standard library only (`flag`, `fmt`, `io`, `os`, `path/filepath`)

## Build

```bash
make build
AGENT_SKILL_PATH=skills make synch-local-skills
make sync-target-skills TARGET=/some/folder
```
