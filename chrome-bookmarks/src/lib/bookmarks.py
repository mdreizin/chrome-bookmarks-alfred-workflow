# -*- coding: utf-8 -*-
import os
import re
import json

_APPLICATION_SUPPORT_PATH = '~/Library/Application Support'
_VENDOR_PATHS = {
    'chrome': os.path.join('Google', 'Chrome'),
    'chromium': 'Chromium'
}


def _load_bookmarks(path):
    try:
        with open(os.path.expanduser(path), 'r') as io:
            data = json.load(io)
    except:
        data = None

    return data


def _inspect_bookmarks(data, chain, predicate):
    if data:
        if type(data) == dict:
            if 'type' in data:
                if data['type'] == 'folder':
                    _inspect_bookmarks(data['children'], chain, predicate)
                elif data['type'] == 'url':
                    if predicate(data):
                        chain.append({'title': data['name'], 'url': data['url']})
            else:
                for value in data.itervalues():
                    _inspect_bookmarks(value, chain, predicate)
        elif type(data) == list:
            for value in data:
                _inspect_bookmarks(value, chain, predicate)
    return


class Provider(object):
    def __init__(self, vendor, profile):
        self.vendor = vendor
        self.profile = profile

    @property
    def __vendor_path(self):
        vendor = _VENDOR_PATHS[self.vendor]

        return os.path.join(_APPLICATION_SUPPORT_PATH, vendor)

    def find_bookmarks(self, query):
        items = []

        path = os.path.join(self.__vendor_path, self.profile, 'Bookmarks')
        data = _load_bookmarks(path)

        if data:
            regexp = re.compile(re.escape(query), re.UNICODE | re.IGNORECASE)

            _inspect_bookmarks(data[u'roots'], items, lambda x: bool(regexp.search(x['name'])) or bool(regexp.search(x['url'])))

        return sorted(items, key=lambda x: (x['title'].lower(), x['url'].lower()))

    def get_profiles(self, query):
        try:
            path = self.__vendor_path
            full_path = os.path.expanduser(path)

            profiles = [{'name': x, 'full_path': os.path.join(path, x)} for x in os.listdir(full_path) if x == u'Default' or x.startswith(u'Profile')]
        except:
            profiles = []

        if query:
            regexp = re.compile(re.escape(query), re.UNICODE | re.IGNORECASE)

            profiles = [x for x in profiles if bool(regexp.search(x['name']))]

        return sorted(profiles, key=lambda x: x['name'].lower())