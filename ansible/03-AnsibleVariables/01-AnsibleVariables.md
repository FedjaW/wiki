# Ansible Variables

Variables are used to store information that varies with each host.

To see all variables used in a playbook use the `-v` flag

```shell
ansible-playbook -i inventory playbook.yml -v
```

## Variable Types

- string
- number
- boolean

## Global Variable

One option is to pass in the variable when running the playbook:

```shell
ansible-playbook playbook.yml --extra-vars "ntp_server=10.1.1.1"
```

## Magic Variables

useful global variables, even host overarching, see official docs

## Ansible Facts

Ansible Facts are system specific variables.
Facts are information about the host machine when you run the playbook.
To gather the Facts ansible will run automatically a task.

```yml
- name: Print all facts
  hosts: all
  tasks:
  - debug:
      var: ansible_facts
```
