task :test do
  puts "Running Go Tests..."
  system('go test ./tests/... -v')
end

task :gosec do
  puts "Running GoSec..."
  system('gosec ./...')
end

task :all => [:test, :gosec] do
  puts "All tasks completed."
end
