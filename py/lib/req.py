import json

import requests


class Req():
    def __init__(self, conf):
        self.conf = conf.conf
        self.url = self.conf['url']
        self.headers = {
            'Content-type': 'application/json',
            'Authorization': 'TOKEN ' + self.conf['token']
        }

    def get(self):
        print(self.headers)
        req = requests.get(self.url, headers=self.headers)
        print(req.text)

    def put(self, payload):
        resp = requests.post(
            self.url, headers=self.headers, data=json.dumps(payload)
        )
        self.print_resp(payload, resp)

    def print_resp(self, payload, resp):
        print('\n' + str(payload['id']) + ' ' + payload['uri'])
        print('Response code: ' + str(resp.status_code))
        print(resp.json())
