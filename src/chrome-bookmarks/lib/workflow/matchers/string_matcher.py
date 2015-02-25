# -*- coding: utf-8 -*-


class StringMatcher(object):
    def __init__(self, query):
        self.query = query or u''

    def ratio(self, query):
        return 0