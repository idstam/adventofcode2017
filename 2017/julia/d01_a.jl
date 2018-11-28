
input = [1, 1, 2, 2]

push!(input, input[1])

println(input)

lastVal = 0
sum = 0
for val in input
    
    if val == lastVal
        sum += val
        lastVal = val
    end
end
println(sum)
