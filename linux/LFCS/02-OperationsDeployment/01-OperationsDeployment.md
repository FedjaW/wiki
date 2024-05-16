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
```
