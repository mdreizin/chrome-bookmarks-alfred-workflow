# encoding: UTF-8
require 'rexml/document'

module Alfred
	class Feedback
		def initialize
			@items = []
		end

		def add_items(items)
			@items = @items + items if items.kind_of? Array
		end

		def add_item(item)
			@items << obj if item.kind_of? Hash
		end

		def to_xml
			doc = REXML::Document.new

			doc << REXML::XMLDecl.new

			root = doc.add_element('items')

			@items.each do |obj|
				item = root.add_element('item')

				obj.each do |key, value|
					if [:uid, :arg, :valid, :autocomplete, :type].include?(key)
						item.attributes[key.to_s] = value.to_s if value
					else
						element = item.add_element(key.to_s)

						if key == 'icon'
							values = value.to_s.split '://'

							if values.length == 2
								text = value.first.to_s

								element.attributes['type'] = values.second.to_s
							else
								text = value
							end
						else
							text = value
						end

						element.add_text text.to_s if text
					end
				end if obj.kind_of? Hash
			end

			doc.to_s
		end

		def self.uid
			Time.now.to_f.to_s.sub(/\./, '')
		end

		def self.item(uid, arg, title, sub_title, autocomplete=nil, valid=true, icon='icon.png', type=nil)
			{
				uid: uid,
				arg: arg,
				title: title,
				subtitle: sub_title,
				autocomplete: autocomplete,
				valid: valid ? 'yes' : 'no',
				icon: icon,
				type: type
			}
		end

		def self.uid_item(arg, title, subtitle, icon='icon.png', valid=true, autocomplete=nil, type=nil)
			self.item self.uid, arg, title, subtitle, autocomplete, valid, icon, type
		end

		def self.empty_item(title, subtitle, icon='icon.png', autocomplete=nil, type=nil)
			self.uid_item nil, title, subtitle, icon, false, autocomplete, type
		end
	end
end