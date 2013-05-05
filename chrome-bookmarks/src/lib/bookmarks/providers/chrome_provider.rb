# encoding: UTF-8
require 'json'

module Bookmarks
	module Providers
		class ChromeProvider < Bookmarks::Provider
			ICON = 'icon.png'

			def find(query_text)
				hash = parse

				query = prepare_query query_text

				items = traversal(hash['roots']).find_all do |x|
					query =~ x['name'].to_s or query =~ x['url'].to_s
				end.map do |x|
					item x['url'].to_s, x['name'].to_s, x['url'].to_s, ICON
				end

				if items.empty?
					items << empty_item('No bookmarks found', 'No bookmarks matching your query were found', ICON)
				end

				items || []
				end

			private
			def read
				path = File.expand_path '~/Library/Application Support/Google/Chrome/Default/Bookmarks'

				File.read path if File.exists? path
			end

			private
			def parse
				content = read

				JSON.parse content unless content.empty?
			end

			private
			def traversal(obj, &block)
				return enum_for :traversal, obj unless block

				if obj.kind_of? Array
					obj.each do |x|
						traversal x, &block
					end
				elsif obj.kind_of? Hash
					if obj['type'] == 'folder'
						traversal obj['children'], &block
					elsif obj['type'] == 'url'
						yield obj
					else
						obj.values.each do |x|
							traversal x, &block
						end
					end
				end
			end

			private
			def prepare_query(query)
				Regexp.new(Regexp.escape(query.to_s).encode('UTF-8'), Regexp::IGNORECASE, 'u').freeze
			end
		end
	end
end