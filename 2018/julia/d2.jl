let
 tot2 = 0
 tot3 = 0
 open("d2data.txt") do file
    for line in eachline(file)
           check = Dict()
                        has2 = false
                        has3 = false
                        for c in line
                                val = get(check, c, 0)
                                check[c] = val +1
                        end
                println(line)
                    for k in keys(check)
                            if check[k] == 2 && !has2
                                    has2 = true
                                    tot2 += 1
                                    print(k)
                                    print(check[k])
                            end
                            if check[k] == 3 && !has3
                                    has3 = true
                                    tot3 += 1
                                    print(k)
                                    print(check[k])

                            end
                    end
println("")
            end
    end

        println("-----")
        println(tot2)
        println(tot3)
        println(tot2 * tot3)

end
