# highspot
You can run the program with this command:
./mixtapeupdate -mixTape <path/to/mixtape/file> -changes <path/to/changes/file> -output(optional) <path/to/output/file>

# about the code
The idea is to deserialize input json into a MixTape struct and perform the operations in memory before serializing the modified object to disk.
I used a json format for changes as well. I have checked in some samples under data folder (e.g. changes_sample1.json)
Given the sample mixtape.json, I tried to optimize the operations for time which mean't sacrificing space (using sets for O(1) lookup time)
There are some special cases that I have intentionally not handled. e.g. playlists with duplicate songIDs as it can be a subjective call and playlists may play the same song multiple times
The test coverage is also not comprehensive but pretty respectable as the intent was mostly to show the kind of code I would write for production

# scalability
When it comes to scalability, the problem changes dramatically as now we maybe limited by memory. This means we won't be able to load whole files in memory or even store them on disk so we have to rely on reading partial data at a given time and perform the operations.
Numerous questions come to mind like what are the memory constraints? Are we working with a single machine or have a cluster of machines that we can work with? Can we manipulate the mixtape format? Just to name a few.
When I would approach this problem in the real world, I will try to understand the data a little more. Meaning how can we break the data down into manageable chunks and be able to serially apply the changes? If we are in a distributed system setting, we could partition the data by playlist ID or user ID and send the changes to the appropriate machine to be applied on a playlist or a user. Once all changes are applied, data could be serially written back to a file. Use of queues such as Kafka is very common to distribute the data among multiple machines.

