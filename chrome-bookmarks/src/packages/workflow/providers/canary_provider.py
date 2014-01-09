# -*- coding: utf-8 -*-
from provider_base import ProviderBase


class CanaryProvider(ProviderBase):
    @property
    def id(self):
        return u'canary'

    @property
    def name(self):
        return u'Google Chrome Canary'

    @property
    def full_path(self):
        return '~/Library/Application Support/Google/Chrome Canary'