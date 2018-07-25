# Go AWS

Is a simple library to aid in AWS account monitoring, responses from
AWS are parsed into individual items (instances, volumes, etc.),
however the original data from AWS is preserved as a decoded JSON
object. 

The data we are interested in is extracted from the original AWS data
by a recievers function.  The function typically just knows the
correct index to dereference the desired value.

Methods will be created that satisfy the "Cloud" interface / model.
