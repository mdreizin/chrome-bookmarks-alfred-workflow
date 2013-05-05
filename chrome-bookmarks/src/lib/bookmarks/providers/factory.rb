# encoding: UTF-8
Dir.glob(File.dirname(File.absolute_path(__FILE__)) + '/*_provider.rb', &method(:require))

module Bookmarks
	module Providers
		module Factory
			def self.create(provider_name)
				eval("Bookmarks::Providers::#{provider_name.to_s.strip.capitalize}Provider").new rescue nil
			end
		end
	end
end