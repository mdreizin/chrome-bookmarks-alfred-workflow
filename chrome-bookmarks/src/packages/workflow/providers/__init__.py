# -*- coding: utf-8 -*-
from chrome_provider import *
from chromium_provider import *

def create(vendor, settings=None):
    """
    :rtype : packages.workflow.providers.ProviderBase
    """
    name = '%sProvider' % vendor.title()
    ctor = globals()[name]
    provider = ctor()

    if settings:
        provider.settings = settings

    return provider