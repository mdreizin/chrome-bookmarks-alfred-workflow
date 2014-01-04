# -*- coding: utf-8 -*-
import re
from string_matcher import StringMatcher


class RegexpMatcher(StringMatcher):
    def __init__(self, query):
        super(RegexpMatcher, self).__init__(query)

        self.regexp = self.to_regexp()

    def to_regexp(self):
        return re.compile(re.escape(self.query), re.UNICODE | re.IGNORECASE)

    def ratio(self, query):
        match = self.regexp.search(query)
        matched = bool(match)
        ratio = 1.0 if matched else 0.0

        return ratio