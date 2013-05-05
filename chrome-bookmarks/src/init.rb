# encoding: UTF-8
def ruby_exec_path(cmd_prefix)
	begin
		ruby_path = `#{cmd_prefix} which ruby`

		if $?.exitstatus == 127
			raise Errno::ENOENT
		end
	rescue Errno::ENOENT
		ruby_path = ''
	end

	ruby_path.to_s.strip
end

def rbenv_ruby_path
	ruby_exec_path '/usr/local/bin/rbenv'
end

def rvm_ruby_path
	ruby_path = ruby_exec_path `/usr/local/.rvm`
	ruby_path = ruby_exec_path `~/.rvm/bin/rvm` if ruby_path.empty?
	ruby_path
end

def current_ruby_path
	ruby_path = '/usr/bin/ruby'

	if RUBY_VERSION < '1.9.3'
		ruby_path = rbenv_ruby_path
		ruby_path = rvm_ruby_path if ruby_path.empty?
	end

	ruby_path
end

def init(*args)
	ruby_path = current_ruby_path

	unless ruby_path.to_s.empty?
		output = `#{ruby_path} -Ku "main.rb" #{args.join(' ')}`

		puts output unless output.empty?
	end
end

init ARGV unless ARGV.empty?