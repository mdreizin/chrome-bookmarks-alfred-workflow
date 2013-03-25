<?php

require_once('workflows.php');
require_once('bookmarks.php');

$query = $argv[1];
$workflow = new Workflows();
$bookmarks = new Bookmarks($workflow->home().'/Library/Application Support/Google/Chrome/Default/Bookmarks');

if (!$bookmarks->exists()) {
	$workflow->result(null, null, 'Bookmarks file not found', 'Unable to locate your bookmarks file', 'icon.png', 'no');
} else {
	$results = $bookmarks->find($query);

	foreach($results as $result) {
		$workflow->result($result['url'], $result['url'], $result['name'], $result['url'], 'icon.png');
	}

	if (count($results) == 0) {
		$workflow->result(null, null, 'No bookmarks found', 'No bookmarks matching your querey were found', 'icon.png', 'no');
	}
}

echo $workflow->toxml();