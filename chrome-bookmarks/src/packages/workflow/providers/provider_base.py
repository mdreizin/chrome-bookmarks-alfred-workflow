# -*- coding: utf-8 -*-
import os
import json
import urllib2
from abc import *
from ..matchers import *
from provider_settings import ProviderSettings

_MINIMUM_RATIO = 50


def load_json(path):
    try:
        with open(os.path.expanduser(path), 'r') as io:
            payload = json.load(io)
    except IOError:
        payload = None

    return payload


def inspect_json(payload):
    if payload:
        if type(payload) == dict:
            if 'type' in payload:
                if payload['type'] == 'folder':
                    for x in inspect_json(payload['children']):
                        yield x
                elif payload['type'] == 'url':
                    yield payload
            else:
                for value in payload.itervalues():
                    for x in inspect_json(value):
                        yield x
        elif type(payload) == list:
            for value in payload:
                for x in inspect_json(value):
                    yield x


def decode_url(url):
    return urllib2.unquote(url.decode('utf-8'))


def make_sortable(query):
    return re.escape(query)


class ProviderBase(object):
    def __init__(self, settings):
        self.__settings = ProviderSettings(self.id, settings)

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
        payload = load_json(self.file_path('Local State')) or {}

        if payload:
            root = (payload['profile'] and payload['profile']['info_cache']) or {}
            profiles = [
                {'name': name,
                 'title': u'%s (%s)' % (profile['name'], profile['user_name']) if profile['user_name'] else profile['name'],
                 'icon': self.icon,
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
            matcher = PartialStringMatcher(query)

            for i, profile in enumerate(profiles):
                name = profile['name']
                title = profile['title']
                full_path = profile['full_path']

                name_ratio = matcher.ratio(name)
                title_ratio = matcher.ratio(title)
                full_path_ratio = matcher.ratio(full_path)

                matched = name_ratio >= _MINIMUM_RATIO or title_ratio >= _MINIMUM_RATIO or full_path_ratio >= _MINIMUM_RATIO

                if not matched:
                    del profiles[i]

        return sorted(profiles, key=lambda x: (make_sortable(x['title']), make_sortable(x['name'])))

    def get_bookmarks(self, query=None):
        bookmarks = []

        if query:
            payload = load_json(self.profile_file_path('Bookmarks'))

            if payload:
                matcher = PartialStringMatcher(query)

                for bookmark in inspect_json(payload['roots']):
                    name = bookmark['name']
                    url = bookmark['url']

                    name_ratio = matcher.ratio(name)
                    url_ratio = matcher.ratio(decode_url(url))

                    matched = name_ratio >= _MINIMUM_RATIO or url_ratio >= _MINIMUM_RATIO

                    if matched:
                        bookmarks.append({'title': name, 'url': url})

        return sorted(bookmarks, key=lambda x: (make_sortable(x['title']), make_sortable(x['url'])))