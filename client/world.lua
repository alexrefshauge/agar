local M = {
	objects = {},
	playerId = -1
}

function M:putObjects(objects)
	for _, o in ipairs(objects) do
		self.objects[o.id] = o
		if o.type == "player" then self.playerId = o.id end
	end
end

function M:unload(unloadIds)
	for _, id in ipairs(unloadIds) do
		M.objects[id] = nil
	end
end

function M:draw()
	for _, o in pairs(self.objects) do
		local x, y = o.pos.x, o.pos.y
		if o.type == "player" then
			love.graphics.setColor(1, 0, 0)
		else
			love.graphics.setColor(1, 1, 1)
		end

		love.graphics.circle("line", x, y, o.size)
	end
end


return M
