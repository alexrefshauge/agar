M = {
	size = 42
}

M.helloWorld = function()
	love.graphics.print("hello world", 400, 400, 0, 10, 10)
end

function M:initialize(size, seed)
	love.graphics.print(self.size, 400, 400, 0, 10, 10)
end

return M
