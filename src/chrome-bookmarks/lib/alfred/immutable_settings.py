# -*- coding: utf-8 -*-
from lib.alfred.settings import Settings


class ImmutableSettings(Settings):
    def set(self, **attributes):
        pass

    def unset(self, **keys):
        pass

    def save(self):
        pass