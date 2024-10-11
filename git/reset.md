# git reset

siehe auch: [[git/boot.dev/reset|reset]] und [[revert_vs_reset]]

- Reset is another command that performs different actions depending on the arguments
  - with a path
  - without a path
    - By default, git performs a `git reset -mixed`

- For commits:
  - Moves the HEAD pointer, optionally modifies files
- For file path:
  - Does not move HEAD pointer, modifies files
