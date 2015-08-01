#!/usr/bin/python3

import time
import socket

hostnames = { "www.google.com",
              "www.intrarts.com",
              "www.apcera.com",
              "www.cisco.com",
              "www.juniper.net",
              "www.huanle.com"
}


for host in hostnames:
    try:
        ip = socket.gethostbyname(host)
        print(host, ip)
        ipex = socket.gethostbyname_ex(host)
        print(host, ipex)
        if socket.has_ipv6:
            addr = socket.getaddrinfo(host, 80, 0, 0, socket.IPPROTO_TCP)
            print(host, addr)
        except socket.gaierror:
            sys.exit()
