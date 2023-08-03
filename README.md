# automaed_ssh_fingerprints
A simple program that fingerprints ssh servers using ssh.Dial() and adds to known_hosts.

Needed to add a util to fingerprint ssh servers and add them to known_hosts file to use the other tools without manually adding to known_hosts.

To compile and use:

```
git clone https://github.com/f0rg-02/automaed_ssh_fingerprints
cd automaed_ssh_fingerprints && go build
.\auto_ssh_fingerprints.exe -h ssh_server_ip -p 22 -f path_to_known_hosts
```

Command line options:

```
Usage of .\auto_ssh_fingerprints.exe: 

  -f string
        Known hosts file location. (default ".ssh/known_hosts")
  -h string
        The SSH host to fingerprint.
  -p int
        The SSH port of host. (default 22)
```
#### TODO: Add YAML config file option so fingerprinting can be done in batches. You can also use BASH to do it in batches.
------
<a href="https://www.buymeacoffee.com/alex_f0rg" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/v2/default-red.png" alt="Buy Me A Coffee" style="height: 60px !important;width: 217px !important;" ></a>
