# -*- coding: utf-8 -*-
import os
import re
import json
from abc import *
from provider_settings import *

class ProviderBase(object):
    @abstractproperty
    def id(self):
        pass

    @abstractproperty
    def name(self):
        pass

    @abstractproperty
    def full_path(self):
        pass

    @property
    def icon(self):
        return u'icons/%s.png' % self.id

    @property
    def profile(self):
        return self.settings.get('profile', u'Default')

    @property
    def profile_path(self):
        return os.path.join(self.full_path, self.profile)

    def file_path(self, name):
        return os.path.join(self.profile_path, name)

    @property
    def settings(self):
        return self.__settings

    @settings.setter
    def settings(self, value):
        self.__settings = ProviderSettings(self.id, value)

    def get_profiles(self, query=None):
        try:
            path = os.path.expanduser(self.full_path)

            profiles = [{'name': x, 'full_path': os.path.join(self.full_path, x)} for x in os.listdir(path) if x == u'Default' or x.startswith(u'Profile')]
        except IOError:
            profiles = []

        if query:
            regexp = re.compile(re.escape(query), re.UNICODE | re.IGNORECASE)

            profiles = [x for x in profiles if bool(regexp.search(x['name']))]

        return sorted(profiles, key=lambda x: x['name'].lower())

    @staticmethod
    def inspect_bookmarks(payload):
        if payload:
            if type(payload) == dict:
                if 'type' in payload:
                    if payload['type'] == 'folder':
                        for x in ProviderBase.inspect_bookmarks(payload['children']):
                            yield x
                    elif payload['type'] == 'url':
                        yield payload
                else:
                    for value in payload.itervalues():
                        for x in ProviderBase.inspect_bookmarks(value):
                            yield x
            elif type(payload) == list:
                for value in payload:
                    for x in ProviderBase.inspect_bookmarks(value):
                        yield x

    def get_bookmarks(self, query):
        bookmarks = []

        if query:
            path = os.path.expanduser(self.file_path('Bookmarks'))

            try:
                with open(path, 'r') as io:
                    payload = json.load(io)
            except IOError:
                payload = None

            if payload:
                regexp = re.compile(re.escape(query), re.UNICODE | re.IGNORECASE)

                for x in ProviderBase.inspect_bookmarks(payload['roots']):
                    matched = bool(regexp.search(x['name'])) or bool(regexp.search(x['url']))

                    if matched:
                        bookmarks.append({'title': x['name'], 'url': x['url']})

        return sorted(bookmarks, key=lambda x: (x['title'].lower(), x['url'].lower()))