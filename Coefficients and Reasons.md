
This readme contains the basics and understanding of Coefficients and ban levels on PsychoPass

**Types of Levels **
- Under 100: Suspect is not a target for enforcement action. The trigger of the Dominator will be locked.
- 100-299: Suspect is classified as a latent criminal and is a target for enforcement action. The Dominator is set to Non-Lethal Paralyzer mode.
- Over 300 ‒ Suspect poses a serious threat to the society. Lethal force is authorized. The Dominator will automatically switch to Lethal Eliminator.

**Special Types of Blacklisting**
- PSYCHOHAZARD - A group of users banned because of some users being banned
- CUSTOM - Any reason not in above is automatically custom, custom has no pattern, its a ban written and explained on why, anything that does not fit above bans is called custom

### Coefficients and Flags

==== Flags     -  ========

| FLAG  |Coefficient Range   | Recommended action  |
| ------------ | ------------ | ------------ |
| Civilian  | 0-80  |  None |
| Restored  | 81-100  | None, but user has a history of past ban  |

#### Actionable Flags

##### Suggested Action is auto mute the user
| FLAG  | Coefficient Range | Explanation  |
| ------------ | ------------ | ------------ |
| TROLLING  | 101-150  | Trolling  |
|  SPAM | 151-200  | Spam/Unwanted Promotion  |
| EVADE  | 201-250  | Ban Evade using alts  |
| CUSTOM  | 251-300  |  Any Custom reason |

##### Suggested Action is ban the user
| FLAG  | Coefficient Range | Explanation  |
| ------------ | ------------ | ------------ |
| PSYCHOHAZARD  |  301-350 |  Bulk banned due to some bad users |
| MALIMP  | 351-400  | Malicious Impersonation  |
| NSFW | 401-450  |  Sending NSFW Content in SFW |
| RAID | 451-500  | Bulk join raid to vandalize   |
| SPAMBOT  | 501-550  |  Crypto, btc, forex trading scams, thotbot |
| MASSADD  | 551-600  |  Mass adding to group/channel |

### Trigger word aliases

| FLAG  | Associated words list  |
| ------------ | ------------ |
| EVADE   |  evade, banevade| 
| MALIMP  |  impersonation, malimp, fake profile| 
| NSFW    |  porn, pornography, nsfw, cp| 
| Crypto  |  btc, crypto, forex, trading, binary| 
| MASSADD |  spam add, kidnapping, member scraping, member adding, mass adding, spam adding, bulk adding| 

### Formula ideas needed for civilian bs 
We can use a weight formula that starts from 80 and deducts points for every good thing done. This same logic can also be given to the server itself but not saved into the DB, just for the assign command beautification.

**Weights - Start from 80, decrease for every thing user has done by some weight 
**
Common groups (int)
USER ID (int)
Username (bool) 10
Profile pic visibility (bool)10
setbio (bool)15
set me (bool)15
has a first name (bool)10
has a last name (bool)5
[========|=]
80-5-10-15-10-10=30
xxxxxx=================================

??????
