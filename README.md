# balajis-script
a script that prints out every clout

# Step 1

Using aws competitor <a href="https://www.vultr.com/?ref=8507322">Vultr</a> we made this pricing choice:

![pricing](https://i.imgur.com/vlOuX5Z.png)

There are cheaper than $40 a month servers but that's the bare minimum to have enough hard drive space and ram to get through the full download. At the time of this writing it will use about 65 GB to download everything.

If you get the $80 or $160 a month ones your processing might go faster but we were able to complete the task with the 160 GB drive one.

Once you have your price selected, create a new instance and select the image:

![image](https://i.imgur.com/fFDIP14.png)

ubuntu 18.04 is a nice stable version and this script has been tested with it:

https://github.com/andrewarrow/balajis-script/blob/main/run.sh

# Step 2

ssh into your new instance as root:

```
ssh root@ip_address_of_new_vultr_server
```

(Note: you need to setup your sshkeys with vultr if you haven't already.)

Once you are in run the commands from the script and notice the sleep 86400 line.

When you get to that one, it'll wait a full 24 hours for the blockchain to get current.

# Step 3

After the long wait run the rest of the commands and you should see posts!


```
Goal: a script that prints out every clout, thereby instantly demonstrating just how different Bitclout is from Twitter.

In more detail:

1) Use the docs (docs.bitclout.com/code/dev-setup) to create a script that downloads the Bitclout repos, syncs a full node, and prints out every clout from the genesis block using locally synced data.

2) Make sure this script works from scratch on a clean server. Let's use AWS for simplicity. I think the Bitclout chain is roughly 70GB (~35000 blocks x 2MB per) so you could use a big AWS instance type, maybe d3en.12xlarge.

3) Ensure the script is as turnkey as possible. Ideally you can run it on a Mac from something like aws-cli (aws.amazon.com/cli) using your local AWS credentials. It should automatically create & provision the instance, run the script, give a progress bar, and notify when done.

4) Write this up as a README.md. Put script + README into a GitHub repo and post the link below. Prize remains open for 7 days.
```

Sample Output:

```
{"Body":"Yes \"storm\" seems perfect!","ImageURLs":null}
{"Body":"18% and yes Raj stole this mode, but now we know it works!\n\nGet in before the greedy whales !","ImageURLs":[]}
{"Body":"Yes! Self-comparison is the only way.","ImageURLs":null}
{"Body":"Anything Vietnamese- bun bo nam bo for example 👍🤤","ImageURLs":null}
{"Body":"I’ve been on my way :( ","ImageURLs":null}
{"Body":"I wasn't going to post it\nbut who'd like to see a pic of the tooth (warning bloody) that was giving me such pain for the last few days\n?","ImageURLs":[]}
{"Body":"You must be set at 99% percent because I want to give you everything.  ","ImageURLs":null}
{"Body":"Nice! wanted to let you know that I DMed you just in case you wasn't notified, I noticed that I don't get notified by email. - DJ Paul w/ 36 Mafia","ImageURLs":null}
{"Body":"","ImageURLs":null}
{"Body":"Canadians are hilarious, what could go wrong?","ImageURLs":null}
{"Body":"😃","ImageURLs":null}
{"Body":"I respect the fuck out of everyone here who has a shitty username 💯 like bro you knew you could make it without being named @sex 😂","ImageURLs":[]}
{"Body":"Breakfast Mukbang \n\n\nCredits: Pambansang Kolokoy 💗💗","ImageURLs":[]}
{"Body":"","ImageURLs":null}
{"Body":"🙏 Thank you very much @RajLahoti 🏆","ImageURLs":[]}
{"Body":"That’s a good one thanks 🙏 ","ImageURLs":null}
```

