<?php

class Bookmarks {
	function __construct($path) {
		$this->path = $path;
	}

	private function load() {
		if (!$this->tree) {
			$this->tree = json_decode(file_get_contents($this->path));
		}
	}

	private function traversal($tree, &$results, $query) {
		foreach($tree as $item) {
			if ($item->type == 'folder') {
				$this->traversal($item->children, $results, $query);
			} else {
				$name = html_entity_decode(htmlspecialchars($item->name));
				$url = html_entity_decode(htmlspecialchars($item->url));
				$lower_query = strtolower($query);
				$lower_name = strtolower($name);
				$lower_url = strtolower($url);
				$isMatched = (strpos($lower_name, $lower_query) !== false || strpos($lower_url, $lower_query) !== false);

				if ($isMatched) {
					array_push($results, array('name' => $name, 'url' => $url));
				}
			}
		}
	}

	private function sort(&$results) {
		if (count($results) > 0) {
			$sort = array();

			foreach($results as $result => $value) {
				$sort['name'][$result] = $value['name'];
				$sort['url'][$result] = $value['url'];
			}

			array_multisort($sort['name'], SORT_ASC, $sort['url'], SORT_ASC, $results);
		}

		return $results;
	}

	public function exists() {
		return file_exists($this->path);
	}

	public function find($query) {
		$this->load();

		$results = array();

		$this->traversal($this->tree->roots, $results, $query);

		return $this->sort($results);
	}
}