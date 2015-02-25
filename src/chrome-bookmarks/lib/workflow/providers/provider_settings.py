# -*- coding: utf-8 -*-


class ProviderSettings(object):
    def __init__(self, prefix, settings):
        self.__prefix = prefix
        self.__settings = settings

    def __cid(self, *keys):
        key = u'.'.join(keys)

        return u'%s.%s' % (self.__prefix, key)

    def get(self, keys, fallback=None):
        return self.__settings.get(self.__cid(keys), fallback)

    def set(self, *attributes):
        options = []

        for attribute in attributes:
            for key in attribute:
                options.append({self.__cid(key): attribute[key]})

        self.sync(*options)

    def sync(self, *attributes):
        self.__settings.set(*attributes)

    def save(self):
        self.__settings.save()