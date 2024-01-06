# automaed_ssh_fingerprints
A simple program that fingerprints ssh servers using ssh.Dial() and adds to known_hosts.

Needed to add a util to fingerprint ssh servers and add them to known_hosts file to use the other tools without manually adding to known_hosts.

Added the ability to use a very simple config file via yaml. Might add an option in the future to just specify a text file of ip addresses with default settings for port and known_hosts file. The idea behind the yaml is to give a bit of flexibility with options and configuring the ssh port per server and yada yada yada. It also looks clean and I have too much free time right now.

Added the ability to specify a file with a list of ips. IP/Hostnames can either be just as is or a colon added with port (e.g. 192.168.1.1 or 192.168.1.1:22 or ilikepie or ilikepie:22 or whatever). See example_file.txt for how it can be used.

To compile and use:

```
git clone https://github.com/f0rg-02/automaed_ssh_fingerprints
cd automaed_ssh_fingerprints && go build

For help: ./auto_ssh_fingerprints -h

Basic usage:

.\auto_ssh_fingerprints.exe -s ssh_server_ip -p 22 -o path_to_known_hosts

or

.\auto_ssh_fingerprints.exe -c config.yaml (See example_config.yaml for the very simple config file)

or

./auto_ssh_fingerprints -i test_ips -o $HOME/.ssh/known_hosts (I've switched to Linux as my main desktop on windows it would be .\auto_ssh_fingerprints.exe)
```

Command line options:

```
Usage of ./auto_ssh_fingerprints: 

  -c string
        The yaml config file with addresses and ports of hosts.
  -h    Show this help (default true)
  -i string
        A file that contains a list of ips
  -o string
        Known hosts file location to write public keys to. (default ".ssh/known_hosts")
  -p int
        The SSH port of host. (default 22)
  -s string
        The SSH host to fingerprint.
```

Required flags are either `-c` or `-s` or `-i`


TODO: Test all the new functions (Done. Everything works beautifully like I knew it would.)
------
<a href="https://www.buymeacoffee.com/alex_f0rg" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/v2/default-red.png" alt="Buy Me A Coffee" style="height: 60px !important;width: 217px !important;" ></a>
