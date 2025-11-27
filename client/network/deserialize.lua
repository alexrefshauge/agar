function type(data)
	return string.byte(data, 1)
end

---unpack state packet
---@param data string
function state(data)
	local playerCount = 0
	local blobCount = 0
	playerCount = string.unpack(">i", data:sub(2, 7))
end
