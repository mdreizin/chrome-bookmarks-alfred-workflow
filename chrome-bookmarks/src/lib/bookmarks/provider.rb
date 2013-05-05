# encoding: UTF-8
require_relative '../../lib/alfred/feedback'

module Bookmarks
	class Provider
		def find(query_text)
		end

		protected
		def item *args
			Alfred::Feedback.uid_item *args
		end

		protected
		def empty_item *args
			Alfred::Feedback.empty_item *args
		end
	end
end