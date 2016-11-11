#!/usr/bin/env python

import re
import sys

def query_type(path):
    pattern = re.compile('type="([a-zA-Z0-9]+)"')
    types = {}
    with open(path) as f:
        for line in f:
            for match in pattern.finditer(line):
                hit = match.group(1)
                types.setdefault(hit, 0)
                types[hit] += 1
    for k, v in types.iteritems():
        print k, '\t\t\t\t', v
            




if __name__ == '__main__' :
    if len(sys.argv) != 2:
        print('USAGE:python query.py <FILE_NAME>')
        sys.exit(-1)

    query_type(sys.argv[1])
