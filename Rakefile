# frozen_string_literal: true

# Color codes for output
module OutputColors
  RED = "\e[31m"
  GREEN = "\e[32m"
  CYAN = "\e[36m"
  YELLOW = "\e[33m"
  RESET = "\e[0m"
end

# Helper method to log task execution with enhanced formatting
def log_task(task_name, command, success_message, failure_message)
  puts "\n#{OutputColors::CYAN}╔═#{'═' * 62}═╗#{OutputColors::RESET}"
  puts "#{OutputColors::YELLOW}🔶 Running #{task_name}...#{OutputColors::RESET}"
  result = if system(command)
             "\n#{OutputColors::GREEN}✔ #{success_message}#{OutputColors::RESET}"
           else
             "\n#{OutputColors::RED}✘ #{failure_message}#{OutputColors::RESET}"
           end
  puts result
  puts "#{OutputColors::CYAN}╚═#{'═' * 62}═╝#{OutputColors::RESET}"
end

# Tasks
task :fmt do
  log_task('Go Fmt', 'go fmt ./...', 'Go Fmt completed', 'Go Fmt failed')
end
task :test do
  log_task('Go Tests', 'go test ./tests...', 'Go Tests passed', 'Go Tests failed')
end
task :gosec do
  log_task('GoSec', 'gosec -quiet ./...', 'GoSec completed', 'GoSec found issues')
end

# Default task runs all tasks
task :default do
  puts "\n#{OutputColors::YELLOW}🔶 Rakefile started. Running all tasks...#{OutputColors::RESET}"
  %i[fmt test gosec].each { |task| Rake::Task[task].invoke }
  puts "#{OutputColors::GREEN}✔ All tasks completed!#{OutputColors::RESET}"
end
