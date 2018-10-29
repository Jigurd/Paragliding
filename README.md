# Paragliding

Choices
- I elected to use the timestamp as an ID. This was a terrible idea and caused me a number of headaches in implementation.

Me Being Bad
- I couldn't for the life of me figure out how to query for a single field from MongoDB, leading to an incredibly inefficient way of getting all timestamps and paging.

Known Issues
- the webhook technically spits out something, but it is not what was specified. The function I had planned is a mess, and I spent the final hours of the time trying to get it to work. 
  This meant I did not have time to implement Clock, which is fairly easy.