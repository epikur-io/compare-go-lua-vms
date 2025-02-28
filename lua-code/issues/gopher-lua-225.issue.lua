local _set = {
	__add = function(k, x)
		local sum = {}
		if #k >= #x then
			for i, v in ipairs(k) do
				sum[i] = v+(x[i] or 0)
			end
		else
			for i, v in ipairs(x) do
				sum[i] = v+(k[i] or 0)
			end		
		end
		return sum
	end;
}

local function set(t)
	local o = t or {}
	setmetatable(o, _set)	return o
end

local t = set({1})
local x = (t + t + t + t + t + t + t + t + t + t + t + t + t + t + t + t + t + t + t + t + t +t + t + t + t + t + t)[1]
