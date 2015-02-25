# -*- coding: utf-8 -*-
import re
from string_matcher import StringMatcher


def to_tag(query):
    return query.strip().lower()


def to_tags(query):
    tags = [to_tag(tag) for tag in query.split()]

    return tags


def to_queries(query):
    tags = to_tags(query)
    queries = [re.compile(u'(%s)' % re.escape(tag), re.IGNORECASE | re.UNICODE) for tag in tags]

    return queries


class PartialStringMatcher(StringMatcher):
    def __init__(self, query):
        super(PartialStringMatcher, self).__init__(query)

        self.queries = to_queries(self.query)

    def ratio(self, query):
        query = to_tag(query)
        matched = all([bool(x.search(query)) for x in self.queries])

        return 100 if matched else 0