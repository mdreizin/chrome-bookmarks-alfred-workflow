# -*- coding: utf-8 -*-
import uuid
import lib.alfred as alfred
import lib.bookmarks as bookmarks

args = alfred.args()

if len(args) > 0:
    query = args[0].strip()
    is_chromium = len(args) > 1;

    if query:
        items = bookmarks.find(query, is_chromium)
        icon = u'icon.png' if not is_chromium else u'icon_chromium.png'

        items = sorted(items, key=lambda x: (x['title'].lower(), x['url'].lower()))
        items = map((lambda x: alfred.Item(
            attributes={'uid': uuid.uuid1().int >> 64, 'arg': x['url'], 'valid': u'yes'},
            icon=icon,
            title=x['title'],
            subtitle=x['url']
        )), items)

        if len(items) == 0:
            items = [alfred.Item(
                attributes={'valid': u'no'},
                icon=icon,
                title=u'No bookmarks found',
                subtitle=u'No bookmarks matching your query were found'
            )]

        xml = alfred.xml(items, len(items))

        alfred.write(xml)