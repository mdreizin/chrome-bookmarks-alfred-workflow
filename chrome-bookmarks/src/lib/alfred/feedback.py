# -*- coding: utf-8 -*-
import itertools
from xml.etree import ElementTree as xml


class Feedback(object):
    def __init__(self):
        self.__items = []

    def add(self, *items):
        for item in items:
            if type(item) == list:
                self.__items.extend(item)
            else:
                self.__items.append(item)

    def xml(self, max_results=None):
        if not max_results:
            max_results = len(self.__items)

        root = xml.Element('items')

        for item in itertools.islice(self.__items, max_results):
            root.append(item.xml())

        return root

    def to_xml(self, max_results=None):
        return xml.tostring(self.xml(max_results), encoding='utf-8')

    class Item(object):
        def __init__(self, attributes, title, subtitle, icon=None):
            self.attributes = attributes
            self.title = title
            self.subtitle = subtitle
            self.icon = icon

        def xml(self):
            root = xml.Element('item', self.attributes)

            for attribute in ('title', 'subtitle', 'icon'):
                value = getattr(self, attribute)

                if not value:
                    continue

                attributes = {}

                if attribute == 'icon':
                    if len(value) == 2 and type(value) == tuple:
                        attributes = {'type': unicode(value[0])}
                        value = value[1]

                child = xml.SubElement(root, attribute, attributes)

                child.text = unicode(value)

            return root

        def to_xml(self):
            return xml.tostring(self.xml(), encoding='utf-8')