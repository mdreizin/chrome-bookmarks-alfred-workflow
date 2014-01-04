# -*- coding: utf-8 -*-
import re
from regexp_matcher import RegexpMatcher


class FuzzyMatcher(RegexpMatcher):
    def to_regexp(self):
        tags = ['(' + re.escape(tag) + ')' for tag in self.query.split()]
        query = '|'.join(tags)

        return re.compile(query, re.UNICODE | re.IGNORECASE)