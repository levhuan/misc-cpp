#! /usr/bin/python3

import os
import sys
import socket
import time

initd_server_addr = "/var/lib/continuum/instance_manager/f6d04c3b/socket"
proto_version = '1\n'
inner_len = '4\n'
wait = 'WAIT'

count = 0

s = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)

try:
    s.connect(initd_server_addr)
except socket.error:
    print("encountered error!")
    sys.exit(1)

s.setblocking(0)
s.sendall(proto_version.encode('utf-8'))
s.sendall(proto_version.encode('utf-8'))
s.sendall(proto_version.encode('utf-8'))
s.sendall(inner_len.encode('utf-8'))
s.sendall(wait.encode('utf-8'))
s.sendall(wait.encode('utf-8'))

while count < 5:
    time.sleep(1)
    count = count + 1

s.close()
