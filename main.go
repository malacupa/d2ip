package main

import (
  "os"
  "strings"
  "fmt"
  "io/ioutil"
  "net"
  "context"
  "time"
  "flag"
  "sort"
  "bytes"
)

func resolve(sem <-chan bool, dom string, out chan<- map[string]string, res []*net.Resolver, useIp6 bool) {
  defer func() { <-sem }()

  ipd := make(map[string]string)
  for _, r := range res {
    ips, _ := r.LookupIPAddr(context.Background(), dom)
    for _, ipa := range ips {
      if useIp6 || ipa.IP.To4() != nil {
        ip := ipa.IP.String()
        if _, ok := ipd[ip]; !ok {
          ipd[ip] = dom
        }
      }
    }
  }

  out <- ipd
}

func stringInSlice(a string, list []string) bool {
  // kudos https://stackoverflow.com/questions/15323767/does-go-have-if-x-in-construct-similar-to-python
  for _, b := range list {
    if b == a {
      return true
    }
  }
  return false
}

type scope struct {
  iplist []string
}

func (s *scope) includes(ip string) bool {
  if s.iplist == nil {
    return true
  } else {
    return stringInSlice(ip, s.iplist)
  }
}

type byIP [][]string

func (s byIP) Len() int {
  return len(s)
}
func (s byIP) Swap(i, j int) {
  s[i], s[j] = s[j], s[i]
}
func (s byIP) Less(i, j int) bool {
  return bytes.Compare(net.ParseIP(s[i][0]), net.ParseIP(s[j][0])) < 0
}

func myUsage() {
  fmt.Printf("Usage: %s [-6] <domain_list.txt> [<ip_scope.txt>]\n", os.Args[0])
  flag.PrintDefaults()
  fmt.Printf("  <domain_list.txt>\tFile containing 1 domain per Unix line\n")
  fmt.Printf("  <ip_scope.txt>\tFile containg 1 IP address per Unix line to filter out only domains resolving to IPs in scope\n")
}

func main() {

  flag.Usage = myUsage
  ip6Ptr := flag.Bool("6", false, "include IPv6 addresses in output")
  flag.Parse()

  if flag.NArg() < 1 || flag.NArg() > 2 {
    flag.Usage()
    os.Exit(1)
  }

  // read list of domains
  dat, err := ioutil.ReadFile(flag.Arg(0))
  if err != nil {
    panic(err)
  }
  doms := strings.Split(string(dat), "\n")

  // read list of scope IPs to error soon
  var s scope
  if flag.NArg() == 2 {
    dat, err := ioutil.ReadFile(flag.Arg(1))
    if err != nil {
      panic(err)
    }
    scopeList := strings.Split(string(dat), "\n")
    s = scope{scopeList}
  } else {
    s = scope{nil}
  }

  // init resolvers
  var res []*net.Resolver
  for _, dnsIP := range []string{"1.1.1.1","8.8.8.8","9.9.9.9"} {
    res = append(res, &net.Resolver {
        PreferGo: true,
        Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
          d := net.Dialer{
            Timeout: time.Millisecond * time.Duration(2000),
          }
          return d.DialContext(ctx, network, dnsIP + ":53")
        },
      },
    )
  }

  out := make(chan map[string]string)
  sem := make(chan bool, 9999) // hardcoded concurrency limit

  for _, dom := range doms {
    sem <- true
    go resolve(sem, dom, out, res, *ip6Ptr)
  }

  ipd := make(map[string][]string)

  for i := 0; i < len(doms) ; i++ {
    for ip, dom := range <-out {
      if s.includes(ip) {
        _, ok := ipd[ip]
        if ok && !stringInSlice(dom, ipd[ip]) {
          ipd[ip] = append(ipd[ip], dom)
        } else {
          ipd[ip] = []string{dom}
        }
      }
    }
  }

  var ipdArr [][]string
  for ip, doms := range ipd {
    ipdArr = append(ipdArr, append([]string{ip}, doms...)) 
  }

  sort.Sort(byIP(ipdArr))
  for _, l := range ipdArr {
    print(strings.Join(l, ","), "\n")
  }
}
