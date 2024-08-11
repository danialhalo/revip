# revip
A command-line utility written in Go for performing reverse IP lookups. This tool fetches and processes data from RapidDNS, extracting domain names associated with a given IP address.
![alt text](https://raw.githubusercontent.com/danialhalo/revip/main/banner.png?raw=true)

## Installation
Compiling from source
```
git clone https://github.com/danialhalo/revip.git
cd revip
go build -o revip
```
Or [download a release for your platform](https://github.com/danialhalo/revip/releases) and put the binary in your `$PATH`.
```
wget https://github.com/danialhalo/revip/releases/download/v0.1.0/revip-linux-0.10.tar
tar -xf revip-linux-0.10.tar
sudo mv revip /usr/bin/
```

## Usage Example
Reverse IP lookup on single IP.
```
echo x.x.x.x | ./revip
```
Reverse IP lookup on Multiple IPs.
```
cat ips.txt | ./revip
```

# License
`revip` is distributed under [MIT License](https://github.com/danialhalo/revip/blob/main/LICENSE)

---

<div align="center">

`revip` was created as a learning exercise for getting more fimilare in Go lang.  [@Muhammad Danial](https://twitter.com/DanialHalo).

</div>

