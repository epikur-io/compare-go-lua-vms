--calculates the nth fibonacci number. Breaks for negative or non-integer n.
local function fibs(n) 
    return n < 2 and n or fibs(n - 1) + fibs(n - 2) 
end

fibs(1)