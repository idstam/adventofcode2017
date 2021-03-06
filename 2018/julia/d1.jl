let
        tot = 0
        prev = Dict()
        foundIt = false
        loopCount = 0
        while !foundIt
            loopCount += 1
            println(loopCount)
            open("data.txt") do file
                for line in eachline(file)
                    operator = line[1]
                    operand = parse(Int, line[2:end])
                    if operator == '+'
                        tot += operand
                    else
                        tot -= operand
                    end
                    if haskey(prev, tot)
                        foundIt = true
                        break;
                    end

                    prev[tot] = 1
                end
            end

        end
        println(tot)

end
