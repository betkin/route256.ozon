#!/usr/bin/env python3

import json

BEGIN_INDEX = 11
END_INDEX = 35000

f = open('ammo.json', 'w')
methods = ['POST', 'GET', 'PUT', 'DELETE']
for i in range(BEGIN_INDEX, END_INDEX):
    for j in range(0,len(methods)):
        uri_list = ['/api/v1/devices', '/api/v1/devices/' + str(i), '/api/v1/devices/' + str(i), '/api/v1/devices/' + str(i)]
        bodies = ["{\"platform\":\"ThisOS\", \"userId\":\"999\"}", "", "{\"platform\":\"AnotherOS\", \"userId\":\"666\"}", '']
        to_json = {
            "host": "act-device-api",
            "tag": methods[j] + '_',
            "method": methods[j],
            "uri": uri_list[j],
            "body":  bodies[j],
            "headers": {
                "header_1": "",
                "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36",
                "Accept": "*/*",
                "Content-Type": "application/json",
                "Accept-Encoding": "gzip, deflate"
            }
        }
        f.write(json.dumps(to_json))
        f.write('\n')
f.close()
