# balajis-script
a script that prints out every clout

Script File: https://github.com/andrewarrow/balajis-script/blob/main/run.sh

```
Goal: a script that prints out every clout, thereby instantly demonstrating just how different Bitclout is from Twitter.

In more detail:

1) Use the docs (docs.bitclout.com/code/dev-setup) to create a script that downloads the Bitclout repos, syncs a full node, and prints out every clout from the genesis block using locally synced data.

2) Make sure this script works from scratch on a clean server. Let's use AWS for simplicity. I think the Bitclout chain is roughly 70GB (~35000 blocks x 2MB per) so you could use a big AWS instance type, maybe d3en.12xlarge.

3) Ensure the script is as turnkey as possible. Ideally you can run it on a Mac from something like aws-cli (aws.amazon.com/cli) using your local AWS credentials. It should automatically create & provision the instance, run the script, give a progress bar, and notify when done.

4) Write this up as a README.md. Put script + README into a GitHub repo and post the link below. Prize remains open for 7 days.
```
