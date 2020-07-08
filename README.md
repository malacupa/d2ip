# d2ip

Supply IP list and get IPs paired with domains in CSV like in usage example below.

## Install

```bash
go get https://github.com/malacupa/d2ip
```

## Usage

```bash
./d2ip
Usage: ./d2ip [-6] <domain_list.txt> [<ip_scope.txt>]
  -6	include IPv6 addresses in output
  <domain_list.txt>	File containing 1 domain per Unix line
  <ip_scope.txt>	File containg 1 IP address per Unix line to filter out only domains resolving to IPs in scope
```

Example output:
```
./d2ip domains2.txt
74.125.0.9,r3---sn-ab5l6nsy.c.mail.google.com
74.125.0.39,r1---sn-ab5l6nzs.c.mail.google.com
74.125.0.198,r1---sn-ab5sznly.c.mail.google.com
74.125.0.199,r2---sn-ab5sznly.c.mail.google.com
74.125.15.214,r4---sn-3c27sn7z.c.mail.google.com
74.125.96.232,r3---sn-oguelner.c.mail.google.com
74.125.102.39,r1---sn-oguesnsy.c.mail.google.com
74.125.106.25,r3---sn-i3b7kn7r.c.mail.google.com,r3.sn-i3b7kn7r.c.mail.google.com
74.125.110.105,r4---sn-4g5ednsr.c.mail.google.com,r4.sn-4g5ednsr.c.mail.google.com
74.125.153.10,r4---sn-hpa7zne6.c.mail.google.com
...continues...
172.217.23.197,alt21.mail.google.com,alt18.mail.google.com,alt7.mail.google.com,alt8.mail.google.com,alt23.mail.google.com,alt1.mail.google.com,alt9.mail.google.com,alt28.mail.google.com,alt5.mail.google.com,alt0.mail.google.com,alt45.mail.google.com,alt47.mail.google.com,alt25.mail.google.com,alt4.mail.google.com,alt2.mail.google.com,alt38.mail.google.com,alt31.mail.google.com,alt49.mail.google.com
172.217.23.229,mail.google.com,alt21.mail.google.com,alt18.mail.google.com,alt8.mail.google.com,alt23.mail.google.com,alt1.mail.google.com,alt9.mail.google.com,alt0.mail.google.com,alt45.mail.google.com,alt38.mail.google.com,alt31.mail.google.com,alt3.mail.google.com,alt49.mail.google.com
172.217.23.231,filetransferenabled.mail.google.com,isolated.mail.google.com,chatenabled.mail.google.com
```

## License
[MIT](https://choosealicense.com/licenses/mit/)
