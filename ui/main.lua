local textField = {
	x = 10,
	y = 10,
	width = 400,
	height = 200,
	text = "",
	active = false,
}

function textField:update(dt)
	local click = love.mouse.isDown(1)
	if not click then return end
	local mx, my = love.mouse.getX(), love.mouse.getY()
	local mouseInside = mx > self.x and mx < self.x + self.width and my > self.y and my < self.y + self.height
	self.active = mouseInside
end

function textField:draw()
	love.graphics.setColor(.1, .1, .1)
	love.graphics.rectangle("fill", self.x, self.y, self.width, self.height)
	love.graphics.setColor(1, 1, 1)
	love.graphics.print(self.text, self.x, self.y, 0, 5, 5)
	if self.active then
		love.graphics.rectangle("line", self.x, self.y, self.width, self.height)
	end
end

function textField:textinput(t)
	if not self.active then return end
	self.text = self.text .. t
end

function love.load()

end

function love.update(dt)
	textField:update(dt)
end

function love.draw()
	textField:draw()
end

function love.textinput(t)
	textField:textinput(t)
end

function love.keypressed(key)
	if key == "escape" then
		love.event.quit()
		return
	end
	if key == "r" then
		love.event.quit "restart"
	end
end
