# -*- coding: utf-8 -*-
import os
import plistlib


class Settings(object):
    def __init__(self, path):
        self.path = os.path.expanduser(path)
        self.__settings = plistlib.readPlist(self.path) if os.path.exists(self.path) else {}

    def set(self, *attributes):
        for attribute in attributes:
            self.__settings = dict(self.__settings, **attribute)

    def unset(self, *keys):
        for key in keys:
            if self.has(key):
                self.__settings.pop(key)

    def has(self, key):
        return key in self.__settings and self.__settings[key] is not None

    def get(self, key, fallback=None):
        return self.__settings[key] if self.has(key) else fallback

    def clear(self):
        self.__settings.clear()

    def save(self):
        path = os.path.dirname(self.path)

        if not os.path.isdir(path):
            os.mkdir(path)

        if not os.access(path, os.W_OK):
            raise IOError('No write access: %s' % path)

        plistlib.writePlist(self.__settings, self.path)