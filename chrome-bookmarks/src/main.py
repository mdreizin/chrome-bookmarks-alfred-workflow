# -*- coding: utf-8 -*-
import lib.alfred as alfred
import lib.bookmarks as bookmarks

args = alfred.args()

if len(args) >= 2:
    workflow = alfred.Workflow()
    vendor = args[0].strip()
    command = args[1].strip()
    query = args[2].strip() if len(args) == 3 else None
    icon = u'icons/%s.png' % vendor
    profile_id = u'%s.profile' % vendor
    profile = workflow.settings.get(profile_id, u'Default')
    provider = bookmarks.Provider(vendor, profile)

    if command == 'find.bookmarks':
        result = map(lambda x: workflow.feedback.Item(
            attributes={'uid': workflow.uid(), 'arg': x['url'], 'valid': u'yes'},
            icon=icon,
            title=x['title'],
            subtitle=x['url']
        ), provider.find_bookmarks(query))

        if not result:
            result = workflow.feedback.Item(
                attributes={'uid': workflow.uid(), 'valid': u'no'},
                icon=icon,
                title=u'No bookmarks found',
                subtitle=u'No bookmarks matching your query were found'
            )

        workflow.feedback.add(result)

        xml = workflow.feedback.to_xml()

        alfred.write(xml)
    elif command == 'get.profiles':
        result = map(lambda x: workflow.feedback.Item(
            attributes={'uid': workflow.uid(), 'arg': x['name'], 'valid': u'yes'},
            icon=icon,
            title=x['name'] if x['name'] != profile else u'* %s' % x['name'],
            subtitle=x['full_path']
        ), provider.get_profiles(query))

        if not result:
            result = workflow.feedback.Item(
                attributes={'uid': workflow.uid(), 'valid': u'no'},
                icon=icon,
                title=u'No profiles found',
                subtitle=u'No profiles were found'
            )

        workflow.feedback.add(result)

        xml = workflow.feedback.to_xml()

        alfred.write(xml)
    elif command == 'set.profile':
        workflow.settings.set({profile_id: query})
        workflow.settings.save()

        alfred.write(query)