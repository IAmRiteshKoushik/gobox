# AMS Sketch

The AMS sketch estimates frequency moments of a data stream.
Random signs turn updates into compact projections whose squared values estimate the second moment.
The second frequency moment measures concentration, including collision-heavy distributions.
Independent repetitions can be combined to reduce estimation error.
It processes updates online with memory independent of the stream length.
Use it for approximate statistics such as the number of equal-item pairs.
