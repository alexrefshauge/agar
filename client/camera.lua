local M = {
	pos = { x = 0, y = 0 },
	center = { x = 0, y = 0 },
	target = { x = 0, y = 0 }
}

function M:update()
	local w, h = love.graphics.getWidth(), love.graphics.getHeight()

	-- smooth follow
	self.pos.x = self.pos.x + (self.target.x - self.pos.x) * 0.1
	self.pos.y = self.pos.y + (self.target.y - self.pos.y) * 0.1

	-- offset so target stays centered in screen
	self.center.x = self.pos.x - w / 2
	self.center.y = self.pos.y - h / 2
end

function M:applyTransform()
	love.graphics.push()
	love.graphics.translate(-self.center.x, -self.center.y)
end

function M:clearTransform()
	love.graphics.pop()
end

return M
