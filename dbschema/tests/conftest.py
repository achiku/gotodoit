# -*- coding: utf-8 -*-
import calendar
from datetime import datetime

import pytest


def pytest_addoption(parser):
    parser.addoption(
        "--token", action="store", default="f34c4ec86600bbfb8ed1aff7d6d48217",
        help="valid access token")
    parser.addoption(
        "--host", action="store", default="http://localhost:8508",
        help="host")


@pytest.yield_fixture(scope='function')
def access_token(request):
    return request.config.getoption("--token")


@pytest.yield_fixture(scope='function')
def host(request):
    return request.config.getoption("--host")


@pytest.yield_fixture(scope='function')
def base_header(access_token):
    return {
        'Content-Type': 'application/json',
        'Authorization': 'bearer {0}'.format(access_token),
        'Gotodoit-UUID': 'xxx',
    }


def get_unixtimestamp():
    return calendar.timegm(datetime.utcnow().timetuple())
