# encoding: UTF-8
require_relative 'lib/bookmarks'
require_relative 'lib/alfred'

def main(provider_name, query_text)
	provider = Bookmarks::Providers::Factory.create provider_name

	if provider.kind_of? Bookmarks::Provider
		feedback = Alfred::Feedback.new
		items = provider.find(query_text).sort_by { |x| [x[:title], x[:subtitle]] }

		feedback.add_items items

		puts feedback.to_xml
	end
end

main ARGV[0], ARGV[1] unless ARGV.empty?