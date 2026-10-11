Well becuase i got a internship my routine got bit messed up and on top of that extremely
lazy and picky guy specially my routines gets compltedly re hauled. So even tho it was really
simple it took me some time to cope with it -  again all blames on me

So am making a crawler and scraper named crawby so i asked my clanker if it suggest some 
articles typp shttt for me to read, it suggested me this -> https://go.dev/blog/pipelines

Well the fist thing the article talks about is making a transformer channel, a channel that is in
middle of  2 channel but its only job is to transform (like hash or multiply) the data from the 
inbound channel  (recieve only channel) and pipe it to another inbound channel. Well the
articke doesnt explicitly names that channel like how i named it a transformer channel. But.
yeah essentially thats what the it really does. 

But there is a huge problem with this kinda architecture and that is everything is happening
in a linear stream. So i thought for while..... even though my answer was present in the article 
itself right after where i had left reading....... But in the end I came up with the idea myself and 
it was to run concurrent transformer channel. Asked my clanker and it said that yeah its 
possible, but note that it cmoes with cost of being inconsistent. so lets say you have  a x1, x2, 
x3, ....... xN data piped to the inbound channel. But you want these  xN data to be piped in a linear 
and consitant pattern such that x2 should come right after x1 than  concurrent transformer channels 
are not what you want as data is piped in a inconsistant manner by different transformer channels.

After this my clanker also suggested me to use buffered channel buffer = faster. According
to the clanker buffer channel are faster becuase a un-buffered channel  copies data 
from one goroutine stack to another goroutines stack via direct memory move (memmove)

well u may wonder why copying is wrong, it is essential what every programme does to share
data or to pass data to each other, but in go the sender channel has to be sleep till the reciever
is born or comes back from sleep. In a way a un-buffered channel is liable to look over the data 
till the reciever comes!

internally a channel is called  `hchan` having a  `buf unsafe.pointer` - This buffer (basically 
ram) can be thought of as mailbox, which can be easily just be picked up ( copy )

well after this it suggested me to use context which i didnt found much ammusing, it was
mostly related to make sure there is no resource leak. After thi i again asked my clanker
as to weather really the programme we just wrote was concurrent, it said that yes it is
 -  but not in the way i thought, i actually didnt think enoug and ofc didnt read properly
 
The whole sauce here was a that we were running  having "worker pool", note that this
isn't a vanilla Go concept, just a way for me to describe it in two word. Though the 
concept of fan-out and fan-in is exactly the same here.

Next am thinking to use these whole text for the clanker to scan and  make a article for
modium

Really great article, as i got to know quite alot. Also i really like Go, its an simple yet complex 
but a realy beautifull language. For me rust and c++ language are kinda scary, yes i used to 
code once in c++ - in my second year for leetcode. I just like how elegantly simple it is. Low 
level feature with simple syntax. It fcking awesome like i have never felt so much great and 
happy programming. I have had fun programming in rust too, but the ecstatic fealing i get
whilst programming in Go is just something else. I will prolly miss this days when i grow more
older in few years
