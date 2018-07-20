# Go AWS

Is a simple library to aid in AWS account monitoring and
management. First it will scan multiple regions (or all regions, if
desired) for running assets.

The running assets can be compared with a set of assets we _expect_
have running.

```bash
% goa [--region region]			
% goa regions
% goa region us-west-1
% goa inst 
% goa inst terminate [--region region] [i-inst1 i-inst2 ... ]
% goa vols 
% goa vols terminate [--region region] [v-vol1 v-vol2 ... ]
% goa bkts 
% goa bkts <bucket-name> objs 
% goa clear bucket
% goa clear intances 
% goa clear volumes
```

It is also capable of cleaning up instances and volumes.  This utility
was used to clean up some hacked accounts that incurred an overnight
attack and spin up of 431 odd Virtual Machines.

## List and Delete Instances

List and delete instances.

## List and Delete Volumes

List and delete volumes.
