# -*- coding: utf-8 -*-
from provider_base import ProviderBase


class ChromeProvider(ProviderBase):
    @property
    def id(self):
        return u'chrome'

    @property
    def name(self):
        return u'Google Chrome'

    @property
    def full_path(self):
        return '~/Library/Application Support/Google/Chrome'