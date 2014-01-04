# -*- coding: utf-8 -*-
import getopt
import sys
import packages.alfred as alfred
import packages.workflow as core


def main(argv):
    try:
        opts, args = getopt.getopt(argv, 'v:c:q:', ['vendor', 'command', 'query'])
    except getopt.error:
        sys.exit(2)

    vendor = None
    command = None
    query = None

    for opt, arg in opts:
        if opt in ('-v', '--vendor'):
            vendor = arg.strip()
        elif opt in ('-c', '--command'):
            command = arg.strip()
        elif opt in ('-q', '--query'):
            query = arg.strip()

    if vendor and command:
        workflow = alfred.Workflow()
        provider = core.providers.create(vendor, workflow.settings)

        icon = provider.icon

        if command == 'get.profiles':
            profile = provider.profile

            results = map(lambda x: workflow.feedback.Item(
                attributes={'uid': workflow.uid(), 'arg': x['name'], 'valid': u'yes'},
                icon=icon,
                title=x['title'] if x['name'] != profile else u'* %s' % x['title'],
                subtitle=x['full_path']
            ), provider.get_profiles(query))

            if not results:
                results = workflow.feedback.Item(
                    attributes={'uid': workflow.uid(), 'valid': u'no'},
                    icon=icon,
                    title=u'No profiles found',
                    subtitle=u'No profiles were found'
                )

            workflow.feedback.add(results)

            xml = workflow.feedback.to_xml()

            alfred.write(xml)
        elif command == 'set.profile':
            provider.settings.set({'profile': query})
            provider.settings.save()

            alfred.write(u'Now %s uses the "%s" profile' % (provider.name, query))

if __name__ == '__main__':
    main(list(alfred.args()))