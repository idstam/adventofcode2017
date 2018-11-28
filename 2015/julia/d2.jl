let
        tot = 0
        totR = 0
        open("data.txt") do file
            for line in eachline(file)
                    tokens = split(line, "x")
                    ints = sort(map(x -> parse(Int, x), tokens))
                    a = 3*ints[1]*ints[2] + 2*ints[2]*ints[3] + 2*ints[1]*ints[3]
                    r = 2*ints[1] + 2*ints[2] + ints[1]*ints[2]*ints[3]
                tot += a
                totR += r
            end
        end
        println(tot)
        println(totR)
end
