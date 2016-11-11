#!/usr/bin/env python
import sys
import bs4
import jinja2
import os

path = os.path.split(os.path.realpath(__file__))[0]
const_path = os.path.join(path, '../../game_server/tlog/const.go')

const_tpl = """package tlog
{%for name,group in groups |dictsort%}
//{{name}}
const (
{%for const in group%} 
    {{const.sign}} = {{const.value}} //{{const.desc}} {% endfor %}
)
{% endfor %}
"""

def gen_code(fd):
    xmlStr = fd.read().decode('utf8')
    bs = bs4.BeautifulSoup(xmlStr)
    code_tpl = jinja2.Environment().from_string(const_tpl)
    const_groups = {}
    for group in bs.find_all('macrosgroup'):
        consts = []
        entries = group.find_all('macro')
        for entry in entries:
            sign = entry['name'] 
            value = entry['value']
            desc = entry['desc']
            consts.append(dict(sign=sign, value=value, desc=desc))
        group_name = group['name']
        if group.has_attr('desc'):
            group_name += group['desc']
        const_groups[group_name] = consts
    data = code_tpl.render(groups = const_groups).encode('utf8')
    print data
    file = open(const_path,'wb')
    file.write(data)
    file.close()





if __name__ == '__main__':
    if len(sys.argv) != 2:
        print('USAGE: ./go_const.py <FILE_NAME>')
        sys.exit(-1)
    with open(sys.argv[1], 'r') as fd:
        gen_code(fd)
