# Configuration & Basic Concepts

When you install `ansible` it creates a default configuration file at `/etc/ansible/ansible.cfg`.

Playbooks will use this global `ansible` configuration.

To use a customized version of a config put a modified `ansible.cfg` file into the playbook.

Another option is to use environment variables to override specific variables.

```shell
# To list all configuration methods run
ansible-config list
```

```shell
# Shows the current config file
ansible-config view
```
