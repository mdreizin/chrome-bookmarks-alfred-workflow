# -*- coding: utf-8 -*-
from provider_base import ProviderBase


class ChromiumProvider(ProviderBase):
    @property
    def id(self):
        return self.name.lower()

    @property
    def name(self):
        return u'Chromium'

    @property
    def full_path(self):
        return '~/Library/Application Support/Chromium'