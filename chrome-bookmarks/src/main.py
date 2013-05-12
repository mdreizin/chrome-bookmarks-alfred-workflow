# -*- coding: utf-8 -*-
import time
import lib.alfred as alfred
import lib.bookmarks as bookmarks

args = alfred.args()

if len(args) >= 2:
    vendor = args[0].strip()
    command = args[1].strip()
    query = args[2].strip() if len(args) == 3 else None
    icon = u'icons/%s.png' % vendor
    profile_id = u'%s.profile' % vendor
    profile = alfred.settings[profile_id] if alfred.settings.has_key(profile_id) else u'Default'
    provider = bookmarks.Provider(vendor, profile)

    if command == 'find.bookmarks':
        if query:
            items = provider.find_bookmarks(query)
            items = map(lambda x: alfred.Item(
                attributes={'uid': alfred.uid(time.time()), 'arg': x['url'], 'valid': u'yes'},
                icon=icon,
                title=x['title'],
                subtitle=x['url']
            ), items)

            if len(items) == 0:
                items = [alfred.Item(
                    attributes={'uid': alfred.uid(time.time()), 'valid': u'no'},
                    icon=icon,
                    title=u'No bookmarks found',
                    subtitle=u'No bookmarks matching your query were found'
                )]

            xml = alfred.xml(items, len(items))

            alfred.write(xml)
    elif command == 'get.profiles':
        items = provider.get_profiles(query)
        items = map(lambda x: alfred.Item(
            attributes={'uid': alfred.uid(time.time()), 'arg': x['name'], 'valid': u'yes'},
            icon=icon,
            title=x['name'] if x['name'] != profile else u'* %s' % x['name'],
            subtitle=x['full_path']
        ), items)

        if len(items) == 0:
            items = [alfred.Item(
                attributes={'uid': alfred.uid(time.time()), 'valid': u'no'},
                icon=icon,
                title=u'No profiles found',
                subtitle=u'No profiles were found'
            )]

        xml = alfred.xml(items, len(items))

        alfred.write(xml)
    elif command == 'set.profile':
        alfred.settings[profile_id] = query

        alfred.flush()
        alfred.write(query)