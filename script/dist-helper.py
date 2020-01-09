#!/usr/bin/env python3
# -*- coding: UTF-8 -*-

import subprocess

compilePath = "../src/"

def runcmd(cmd,parameter):
    subprocess.Popen(cmd+" "+parameter)


def compile():
    runcmd('set','GOOS=linux')
    runcmd('go build',compilePath+'main.go')

if __name__ == '__main__':


