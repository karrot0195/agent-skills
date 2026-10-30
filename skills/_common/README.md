# Common skill components

Shared reusable blocks for skills. Skills may load these or be optimized to reference common patterns.

- `user-input.md` - Standard $ARGUMENTS handling block
- `extension-hooks.md` - Generic extension hooks pre-execution check

Special per-skill configuration can be placed in `.special/<skill-name>.yaml` (YAML) to override behavior without duplicating full skill logic.
