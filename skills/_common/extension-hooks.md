## Pre-Execution Checks

**Check for extension hooks (before execution)**:
- Check if `.specify/extensions.yml` exists in the project root.
- If it exists, read it and look for entries under the `hooks.before_<HOOK>` key where `<HOOK>` matches the current command context.
- If the YAML cannot be parsed or is invalid, do not skip silently: tell the user that `.specify/extensions.yml` could not be read (include the parser error) and that no hooks were checked, including any mandatory (`optional: false`) hooks registered there, then continue normally
- Filter out hooks where `enabled` is explicitly `false`. Treat hooks without an `enabled` field as enabled by default.
- For each remaining hook, do **not** attempt to interpret or evaluate hook `condition` expressions:
  - If the hook has no `condition` field, or it is null/empty, treat the hook as executable
  - If the hook defines a non-empty `condition`, skip the hook and leave condition evaluation to the HookExecutor implementation
- When constructing command invocations from hook command names, replace dots (`.`) with hyphens (`-`). For example, `speckit.git.commit` → `/speckit-git-commit`.
- For each executable hook, output the following based on its `optional` flag:
  - **Optional hook** (`optional: true`): show hook info with `To execute: /{command}`
  - **Mandatory hook** (`optional: false`): show `EXECUTE_COMMAND: {command}` and **WAIT for the result before proceeding**. Run it the same way as the command itself (e.g. `/skill:speckit-...` or `$speckit-...`).
- If no hooks are registered or `.specify/extensions.yml` does not exist, skip silently
