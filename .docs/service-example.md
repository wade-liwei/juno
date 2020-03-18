# Juno service example
Following you will find how to create a Linux service to run Juno in the background.
  
It assumes you have properly cloned the repo and installed `juno` using the `ubuntu` user. If your user is different, please change it accordingly:

```bash
tee /etc/systemd/system/juno.service > /dev/null <<EOF  
[Unit]
Description=Juno service
After=network-online.target

[Service]
User=ubuntu
ExecStart=/home/ubuntu/go/bin/djuno parse <path/to/config.toml>
Restart=always
RestartSec=3
LimitNOFILE=4096 # To compensate the "Too many open files" issue.

[Install]
WantedBy=multi-user.target
EOF
```

Once created, you can start it running 

```bash
systemctl enable juno
systemctl start juno
```

You can verify the status by running 

```
systemctl status juno
```

Or see the logs using 

```
tail -100f /var/log/syslog
```
