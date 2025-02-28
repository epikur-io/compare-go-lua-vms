local function a(n,u,s) if n<2 then return u+s end return a(n-1,u+s,u) end
local function trfib(i) return a(i-1,1,0) end

trfib(1)