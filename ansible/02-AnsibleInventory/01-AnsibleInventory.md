# Ansible Inventory

Ansible can work with multiple servers, therefore it needs to establish connectivity to those servers.
This is done via ssh.

Ansible is agent-less, meaning you do not need to install additional software on the target mashines, an ssh connection is sufficient.

Information about the target systems are stored in an inventory file.
The default inventory file is stored at `/etc/ansible/hosts`.
