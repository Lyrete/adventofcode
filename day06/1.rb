def parse_file(filename)
    fishes = {0 => 0,1 => 0,2 => 0,3 => 0,4 => 0,5 => 0,6 => 0,7 => 0,8 => 0}
    File.readlines(filename).each do |line|
        fish_data = (line.strip.split(',').map(&:to_i))
        for fish in fish_data
            fishes[fish] += 1
        end
    end
    return fishes
end

def pass_day(fishes)
    #Store amount of fish at 0 while we loop through the rest
    new_fish = fishes[0]
    fishes[0] = 0

    fishes.each do |key, value|
        if key != 0 #remove one day from each fish
            fishes[key-1] += value
            fishes[key] -= value        
        end        
    end

    #Return stored fish to hash
    fishes[6] += new_fish
    fishes[8] += new_fish

    #return the hash of fishes
    return fishes
end

def pass_days(days, fishes)
    for i in 1..days do
        fishes = pass_day fishes
    end
    return fishes
end

def count_fishes(fishes)
    count = 0
    fishes.each do |key, value|
        count += value     
    end
    return count
end



