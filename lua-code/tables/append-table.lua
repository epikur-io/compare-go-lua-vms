local function append_table()
    local t = {}
    for i=1,500 do 
        table.insert(t, i)
    end
end

append_table()