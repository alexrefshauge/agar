local M = {
	objects = {}
}

function M:putObjects(objects)
	for _, o in ipairs(objects) do
		self.objects[o.id] = o
	end
end

function M:draw()
	for _, o in pairs(self.objects) do
		local x, y = o.pos.x, o.pos.y
		if o.type == "player" then
			love.graphics.setColor(1, 0, 0)
			love.graphics.print(o.name, 10, 10)
		else
			love.graphics.setColor(1, 1, 1)
		end
		love.graphics.circle("line", x, y, o.size)
	end
end

return M
