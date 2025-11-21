---@class UpdateData
---@field blobs Blob[]
---@field players Player[]
---@field unload integer[]
---@field eat integer[]
---@field size? integer

---@class DataUtils
---@field toUpdateData function
local M = {}

---ensure that table is updateData
---@param data table
---@return UpdateData
function M:toUpdateData(data)
    ---@type UpdateData
    return {
        blobs = data.blobs or {},
        players = data.players or {},
        unload = data.unload or {},
        eat = data.eat or {},
        size = data.size
    }
end

return M