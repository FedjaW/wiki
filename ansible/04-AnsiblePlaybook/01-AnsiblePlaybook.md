# Ansible Playbook

```shell
# Run a playbook
ansible-playbook playbook.yml
```

```shell
# Help
ansible-playbook --help
```

## Verifying Playbooks

- Check Mode
  - Ansible's "dry run" where no actual changes are made on hosts
  - Allows preview of playbook changes without applying them
  - Use the `--check` option to rujn a playbook in check mode
- Diff Mode
  - Provides a before-and-after comparison of playbook changes
  - Understand and verify the impact of playbook changes before applyhing them
  - Utilize the `--diff` option to run a playbook in diff mode
- Syntax Check
  - Ensures playbook syntax is error-free
  - Use the `--syntax-check` option

## Ansible lint

A command line tool to lint your playbook.

```shell
# run
ansible-lint playbook.yml
```
