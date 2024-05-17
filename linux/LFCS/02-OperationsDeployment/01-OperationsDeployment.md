# Operations Deployment

## Boot, Reboot and shutdown a system safely

reboot

```shell
# root user
systemctl reboot
# regular user
sudo systemctl reboot
# force it
sudo systemctl reboot --force
# like pressing reset button
sudo systemctl reboot --force --force
```

shutdown

```shell
systemctl poweroff
# regular user
sudo systemctl reboot
# force it
systemctl poweroff --force
# like unplugging from power source
systemctl poweroff --force --force
# shutdown 2 hours from now
sudo shutdown +120
# to cancel the planned shutdown
sudo shutdown -c
```

## Boot change system into different operating modes

default

```shell
# oprating system is configured to boot into a graphical environment
systemctl get-default
# graphical.target
```

set it to multi-user target

```shell
sudo systemctl set-default multi-user.target
```

In this files the startup programms are defined and in what order they start.

## Install, configure and troubeshoot bootloader

topic `grub`, skipped.

## Manage the startup process and services

Linux has an `init = initialization system`, that controlls the boot up of the programms, even if one crashes it restarts it.

There are various `units`:

- service
- socket
- device
- timer

In a nutshell the service unit tells the init system all it needs to know about how to manage the entire lifecycle of a certain application.

```shell
# man page of the service unit
man systemd.service
```

For example the service unit needs to boot the ssh-daemon to let users connect to it via ssh.
There is a service unit that instructs the init system how to do that.

```shell
# to look at that file:
systemctl cat sshd.service
```

```shell
# man page of systemd services
man systemd.services
```

```shell
# look at a list of system wide services (file from here can be used as a template)
ls /lib/systemd/system
```

With an own systemd file like e.g. `/etc/systemd/system/myapp.service` we can set up restart instructions, startup instructions etc. for our own app that must be controlled by the system.
After adding such an own service we would need to reload the daemon: `sudo systemctl daemon-reload`.

To start the service run `sudo systemctl start myapp.service`.
To see logs run `sudo journalctl -f`.
