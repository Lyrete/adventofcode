def parse_file(filename)
    fishes = []
    File.readlines(filename).each do |line|
        fishes = (line.strip.split(',').map(&:to_i))
    end
    return fishes
end

def pass_day(fishes)
    for i in 0...fishes.length
        fish = fishes[i]
        if fish == 0
            fishes.append(8)
            fishes[i] = 6
        else
            fishes[i] -= 1
        end
    end
    return fishes
end

def pass_days(days, fishes)
    for i in 1..days do
        fishes = pass_day fishes
    end
    return fishes
end



