function log(msg)
    mc.tellraw("[LOG] " .. msg, "@a")
end

luna.tick(function()
    log("So awesome!")
    log("So sigma!")
end)