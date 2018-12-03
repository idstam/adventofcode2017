

let
bounds = 1100

fabric = zeros(Int, bounds, bounds)
claims = zeros(Int, bounds, bounds)

 function unClaim(claimID)
   for x in 1:bounds, y in 1:bounds
     if claims[x, y] == claimID
       claims[x, y] = -1
     end
   end
 end

 open("d3data.txt") do file
    for line in eachline(file)
      line = replace(line, ":" => "")

      tokens = split(line, " ")
      claimID = parse(Int, replace(tokens[1], "#" => ""))
      pos = split(tokens[3], ",")
      size = split(tokens[4], "x")

      for x in parse(Int, pos[1]): parse(Int, pos[1]) + parse(Int, size[1])-1
        for y in parse(Int, pos[2]): parse(Int, pos[2]) + parse(Int, size[2])-1

          claim = claims[x+1, y+1]
          if claim == 0
            claims[x+1, y+1] = claimID
          elseif claim > 0
            unClaim(claim)
            unClaim(claimID)
            claimID = -1
          end
          fabric[x+1, y+1] = fabric[x+1, y+1] +1
        end
      end
    end
  end
  ctr = 0

  for x in 1:bounds, y in 1:bounds
    if fabric[x, y] > 1
      ctr += 1
    end
    if claims[x, y] > 0
      println(claims[x, y])
      break
    end
  end

  println(ctr)
end
