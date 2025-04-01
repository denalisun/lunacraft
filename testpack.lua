-- execute as @a if score @s test.sigma matches 1.. run say hi
-- this is just a dumb test btw ignore

execute()
    :as(Players.ALL)
    :ifScoreBetween(Players.SELF, "test.sigma", 1, nil)
    :run(function()
        say("So sigma!", Players.SELF)
end)