# -*- coding: utf-8 -*-
import uuid
import lib.alfred as alfred
import lib.bookmarks as bookmarks

args = alfred.args()

if len(args) > 0:
    query = args[0].strip()

    if query:
        items = bookmarks.find(query)

        items = sorted(items, key=lambda x: (x['title'].lower(), x['url'].lower()))
        items = map((lambda x: alfred.Item(
            attributes={'uid': uuid.uuid1().int >> 64, 'arg': x['url'], 'valid': 'yes'},
            icon=u'icon.png',
            title=x['title'],
            subtitle=x['url']
        )), items)

        if len(items) == 0:
            items = [alfred.Item(
                attributes={'valid': 'no'},
                icon=u'icon.png',
                title=u'No bookmarks found',
                subtitle=u'No bookmarks matching your query were found'
            )]

        xml = alfred.xml(items, len(items))

        alfred.write(xml)