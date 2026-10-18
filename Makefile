.PHONY: build synch-local-skills sync-target-skills

build:
	mkdir -p bin
	GO111MODULE=off go build -o bin/agent-synch-skills ./cmd/sync-skills

synch-local-skills: build
	AGENT_SKILL_PATH=skills bin/agent-synch-skills -Target=.agents/skills

sync-target-skills: build
	@if [ -z "$(TARGET)" ]; then echo "Error: TARGET is required (e.g., make sync-target-skills TARGET=foo)"; exit 1; fi
	bin/agent-synch-skills -Target=$(TARGET)
