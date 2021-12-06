require 'set'

def parse_file_unique(filename)
    sum = 0
    answers = Set[]
    File.readlines(filename).each do |line|
        data = (line.strip.split(''))
        if data.length == 0
            sum += answers.length
            answers = Set[]
            next
        end
        answers = answers.merge(data)
    end
    sum += answers.length
    return sum
end

def parse_file_all(filename)
    sum = 0
    answers = Set[]
    new_group = true
    File.readlines(filename).each do |line|
        data = (line.strip.split(''))
        if data.length == 0
            sum += answers.length
            answers = Set[]
            new_group = true
            next
        end
        if answers.length == 0 && new_group
            answers = answers.merge(data)
            new_group = false
        end
        answers = answers & data
    end
    sum += answers.length
    return sum
end

puts 'Task 1:'
puts parse_file_unique 'data.txt'
puts 'Task 2:'
puts parse_file_all 'data.txt'
