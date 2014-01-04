# -*- coding: utf-8 -*-
import os
import json
from abc import *
from ..matchers import *


class ProviderBase(object):
    def __init__(self, settings):
        self.__settings = settings

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
        return self.settings.get(u'profile', u'Default')

    @property
    def profile_path(self):
        return self.file_path(self.profile)

    def profile_file_path(self, name):
        return os.path.join(self.profile_path, name)

    def file_path(self, name):
        return os.path.join(self.full_path, name)

    @property
    def settings(self):
        return self.__settings

    def get_profiles(self, query=None):
        profiles = []
        payload = ProviderBase.load_json(self.file_path('Local State')) or {}

        if payload:
            root = (payload['profile'] and payload['profile']['info_cache']) or {}
            profiles = [
                {'name': name, 'title': '{name} ({user_name})'.format(**profile), 'icon': self.icon,
                 'full_path': self.file_path(name)} for (name, profile) in root.items()]
        else:
            try:
                path = os.path.expanduser(self.full_path)

                if os.path.exists(path):
                    profiles = [{'name': profile, 'title': profile, 'full_path': os.path.join(self.full_path, profile),
                                 'icon': self.icon} for
                                profile in os.listdir(path) if profile == u'Default' or profile.startswith(u'Profile')]
            except IOError:
                profiles = []

        if query:
            matcher = RegexpMatcher(query)

            for i, profile in enumerate(profiles):
                name_ratio = matcher.ratio(profile['name'])
                title_ratio = matcher.ratio(profile['title'])
                total_ratio = name_ratio + title_ratio
                matched = total_ratio >= 0.5

                if not matched:
                    del profiles[i]

        return sorted(profiles, key=lambda x: (x['title'], x['name']))

    @staticmethod
    def load_json(path):
        try:
            with open(os.path.expanduser(path), 'r') as io:
                payload = json.load(io)
        except IOError:
            payload = None

        return payload

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

    def get_bookmarks(self, query=None):
        bookmarks = []

        if query:
            payload = ProviderBase.load_json(self.profile_file_path('Bookmarks'))

            if payload:
                matcher = FuzzyMatcher(query)

                for bookmark in ProviderBase.inspect_bookmarks(payload['roots']):
                    name_ratio = matcher.ratio(bookmark['name'])
                    url_ratio = matcher.ratio(bookmark['url'])
                    total_ratio = name_ratio + url_ratio
                    matched = bool(total_ratio >= 0.5)

                    if matched:
                        bookmarks.append({'title': bookmark['name'], 'url': bookmark['url']})

        return sorted(bookmarks, key=lambda x: (x['title'], x['url']))