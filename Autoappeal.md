# Auto Appeal

Auto Appeal System schematics
> Just some basic details on how the auto appeal system should function.


                        +-----------------+
                        |                 |
                        | User starts bot |
                        |                 |
            +-----------+-----------------+----------------------+
            |                                                    |
            |                                                    |
            |                                                    |
            |                                                    |
    +-------v-------+                                +-----------v----+
    |               |                                |                |
    | If not banned |                                |   If banned    |
    |               |                                |                |
    +-------+-------+                                +-----------+----+
            |                                                    |
            |                                                    |
            |                                                    |
            |                                                    |
            |                                        +-----------v---------+
    +-------v-------+                                |                     |
    |               |                                |   Print ban details |
    |  Print scan   |                                |                     |
    |               |                                ++--------------------+-----------------+
    +---------------+                                 |                                      |
                                                      |                                      |
                                                      |                                      |
                                                      |                           +----------v-+
                                               +------v----+                      |            |
                                               |           |                      |  close box |
                                               | unban me  |                      |            |
                                               |           |                      +------------+
                                       +-------+-----------+------------+
                                       |                                |
                                       |                                |
                                       |                                |
                                       |                                |
                                       |                                |
                                       |                                |
                              +--------v--------+             +---------v-----+
                              |                 |             |               |
                              | I get it, unban |             | Go to support |
                              |                 |             |               |
                              +-----------------+             +---------------+
                                |                                |
                                |                                |
                                |                                |
                                |                                |
                                |                                |
                                |                                | 
               +----------------v+            +------------------v----------+
               |                 |            |                             |
               |                 |            |                             |
               |  unban user     |            |  Deny if high coefficient   |
               |                 |            |                             |
               |                 |            |                             |
               +-----------------+            +-----------------------------+


<hr/>

## If user is not banned

All triggers start with the `/start` command, if a user is not banned and sends `/start`, print the following:

### Step 1
Welcome to Sibyl System! \
Please wait while we finish your cymatic scan...\
sleep 5


> _Bot edits the message_

### Step 2
Welcome to Sibyl System! \
Cymatic Scan results:
 - **User**: Dank
 - **ID**: 2039641378
 - **Is banned**: No
 - **Status**: `Civilian`
 - **Crime Coefficient**: Under 100

 <!--
   NOTE: the character that sawada used here was '•', I deleted them, because markdown already
   has a built-in character for that usage.
   But on telegram, in the bot, this '•' character should be used.
-->

[Support group] (goes to psb) \
[What is PsychoPass?] (goes to @PsychoPass) \
[Report Spam] (goes to psb message in group explain how to report, that is pinned) \
[Get API token] 

<hr/>

## If user is banned
### Step 1
Welcome to Sibyl System! \
Please wait while we finish your cymatic scan...
> _sleep 5s_
### Step 2
Welcome to Sibyl System! \
Cymatic Scan results:
 - **User**: Dank
 - **ID**: 2039641378
 - **Is banned**: Yes
 - **Status**: Civilian
 - **Crime Coefficient**: 650
 - **Ban short reason**: `MASSADD`
 - **Ban long reason**: `admin in a group where people were mass adding.`

<!--
   NOTE: the character that sawada used here was '•', I deleted them, because markdown already
   has a built-in character for that usage.
   But on telegram, in the bot, this '•' character should be used.
-->

Since this is your first time we can allow you a one time exception provided that you will not repeat this ever again. \
[I will not do this again! ] \
[Close this message]
### Step 3
If user clicks on the `I will not do this again` button then:

`[username]` You were blacklisted on Sibyl for the reason "SHORT REASON HERE IN LOWERCASE".
 
 _--Detail String for short reason here--_

Such type of actions are often unwanted and unwelcome around Sibyl. \
Please do note that should this ever happen again your ban will be swift and its damage, measurable on the richter scale! \
Click the button below to confirm that you understand this and if you have questions please click the Support button to take your query to the bureau \
[ I read and understand, unban me!] \
[Take me to Support] (take to psb)

### Step 4
If user clicks `I read and understand`:

Check if user was banned for reason `CUSTOM` or if they have a psychopass greater than 600, if it is greater than 600 then reply with:


  Sorry, your coefficient is greater than "value" and cannot be revoked via the auto appeal system, please take your questions to `@PublicSafetyBureau` if you want an unban. \
  [Support group] (goes to psb) \
  [What is PsychoPass?] (goes to @PsychPass) \
  [Report Spam] (goes to psb message in group explain how to report, that is pinned) \
  [Get API token]

<hr/>

## Detail String for short reasons:
| FLAG  | Associated words list  |
| ------------ | ------------ |
| `TROLLING`  | Trolls aren't welcome on sane groups, when you go to groups just to annoy the admins to show much you are in control of mayhem, sibyl steps in. We do not welcome trolls and misbehavers. |
|  `SPAM` | Users who post unwanted content with the aim to promote their own products or links aren't welcome around the communities we protect. |
| `EVADE`  | Users that create more accounts to then evade a previously assigned ban are simply just as guilty, changing your account does not remove your previously caused drama from telegram |
| `CUSTOM`  | xx-- **no string here, cause this is manually appealed** --xx
| ------------ | ------------ |
| `PSYCHOHAZARD`  | You were blacklisted because you were either the owner of a group where someone was spam adding members or cause trouble for other groups and users, when you are in authority to change things and all you do is sit and watch, you are just as guilty as the person who cause the problem. Its about time you have some responsibility for the groups you were admin in. | 
| `MALIMP`  | You were fooling around with intentions to either harm or to either affect the profile of another established user, we do not welcome such users around our communities, you suck! |
| `NSFW` | You were found posting pornographic or suggestively pornographic content in groups that do not welcome such content. |
| `RAID` | You and your pals were engaged in raining a group/bot with the attempt to vandalize, this ban is unappealable and you are never welcome around our communities. 
| `SPAMBOT`  | You were behaving like a scam bot that attempts to ensnare users with falsified data in attempt to scam them. 
| `MASSADD`  | You were spam adding members from other groups to your own, not only this not welcome as platform terms of service this is also unwelcome around Sibyl, your ban will not be appealable. 


