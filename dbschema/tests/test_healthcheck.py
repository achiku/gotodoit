# -*- coding: utf-8 -*-
import json
import requests


def test_healthcheck(base_header, host):
    res = requests.get(
        host+'/v1/healthcheck', headers=base_header)
    j = res.json()
    print(json.dumps(j, indent=2))

    assert res.status_code == 200
