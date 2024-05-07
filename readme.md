# Battle Algorithm:
Follows rules below:
    * Highest speed monster goes first but if speed is equal then higher attack among two goes first
    * Damage is equal to (attack-defense) but the difference is zero or less than zero then damage is 1.
    * Remaining HP equals (HP-damage)
    * Monsters battles in turns until one wins.
    * The winner will the one who make the opponent hp to zero.


NOTE: This was the part of assignment that I got to work on and all built it from scratch and I tried to utilize concurrency feature as well in the project in the endpoint "FightAll". I added delay of 2 seconds to prove that concurrency is working fine. 
There are 3 endpoints that I built 
    * /monsters - to get monster information 
    * /monster - to create new monster
    * /FightAll - Fight all the combination of monster and give the result of who fought and who won as a result.

This is a simple api. Next thing to add is adding authentication and proper validation and more unit test