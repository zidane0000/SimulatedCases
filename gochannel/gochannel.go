package gochannel

import (
	"fmt"
	"sync"
	"time"
)

const (
	maxBuffer   = 100
	flushPeriod = 5 * time.Second
	article     = `Jewish teachings can guide us out of this mess
Forest Hills: The Jewish High Holy Days weighed heavily upon me. On Rosh Hashanah and Yom Kippur, my eyes fixed on the Hebrew text of my prayer book. The words seemed to cry out to me: The world is broken, do something to fix it!

So much is at stake in this Jewish New Year. May God give us all, Jews and non-Jews, the strength of mind and spirit to believe in the inherent goodness of our various religious traditions — and act on it.

I pray for Israel harder than ever, and for all the innocent lives caught up in the terrible Mideast crisis. I implore Israel’s Prime Minister Benjamin Netanyahu to stop the bombing in Gaza. Are more than 42,000 deaths of Palestinians not enough? Until this happens, as so many hostage families and the majority of Israelis know, there will be no peace, no path to a two-state solution. The Mideast will be set ablaze as never before.

Jewish tradition teaches the precept of tikkun olam (fixing the world), of pursuing peace and righteousness through our words and deeds. I say to my fellow Jews: If any hate smolders in your hearts, quell it. Only then will we as a people dare to speak truth to power. It is a painful but essential first step toward pursuing peace. Carmela Ingwer

Butts on fire
Brooklyn: Warning signs should be posted: “Do not throw lit cigarettes or matches from your vehicle onto roads.” I had to stomp out a lit cigarette butt someone threw on my block as they were driving by and I was picking up leaves. Marie Walsh

Show, don’t tell
Brooklyn: We can learn a lot from our British cousins regarding traffic safety and road design. Instead of installing ever more signs on our already confusing Long Island parkways, the New York State Department of Transportation should paint at entrance ramps bold reflective arrows indicating the direction that cars are intended to travel. Pictograms are quicker for the brain to interpret than words in the nanoseconds we have while driving. Eric Wollman

Ups and downs
Ridgewood, N.J.: Every election has winners and losers. Winners: the American public — they will set the agenda; energy production will expand. Losers: money — Democrats outspent Republican two to one; unions — favored one party over the other; elites — out of touch; Barack Obama — preaching gender and race as a reason to vote; media — overall, out of touch with the public; Ukraine — could be forced to compromise; Iran — provocations will be matched with force; racism — enough is enough. Uncertain: NATO, tariffs, Taiwan, alliances, the environment, China. Ed Houlihan

Peaceful transfer
Peters Township, Pa.: I want to thank President Biden and Vice President Kamala Harris for their graciousness in defeat and for demonstrating that they care more about the country than themselves and their egos. In 2020 and to this day, Donald Trump has refused to accept that he lost that election. He did not ever concede, did not invite Biden to the White House and did not attend the inauguration. Instead, he sicced a mob on the U.S. Capitol in an effort to reverse the will of the people. After a bruising and painful loss, Harris called Trump to concede and congratulate him, and Biden invited Trump to the White House to begin what he promises will be an orderly transition in which his administration fully cooperates with the incoming team. Thank you, Biden and Harris, for accepting the will of a narrow majority of the people and for displaying class and dignity! Oren Spiegler

Picked our poison?
Bronx: For the love of God, I hope I’m wrong. But if we are to judge President-elect Trump’s intentions based on his past behavior and the obnoxious tone of his unbelievably and inexplicably successful presidential campaign, our nation is in serious trouble. Weren’t any of those millions who voted for him concerned or even aware of the long list of lies he told during the campaign? Did they also forget about the mayhem and death he fostered on Jan. 6, 2021? Well, now he returns, only this time with the almost unlimited powers of a bonafide dictator. He now owns the tools by which to affect lasting and impactful changes upon our country as detailed in Project 2025. I have no regrets, for I did not vote for him. But for those who did, will they someday soon come to regret it? For our nation’s sake, I sincerely hope not. Carlos Martinez

Group discount
Briarwood: Of course, the president-elect will provide a blanket pardon to all Jan. 6 defendants at his first opportunity. By doing so, he will be pardoning himself without having to pardon himself by name. Mary Elizabeth Ellis

Citing Scripture
Queens Village: To Voicer Melanie Noonan: God forgives only repentant sinners, which is what King David was after his sin with Bath-Sheba. When the prophet Nathan made him aware of his grave sins, David not only immediately and humbly repented, he did not imprison or kill Nathan, as is stated at 2 Samuel 12:7-14. His heartfelt repentance and attitude is beautifully recorded in Psalm 51. Trump, unlike David, is sorely lacking in humility. Can anyone name one time he admitted to or even apologized for a mistake, or not attempted to destroy the livelihood or character of the Nathans who had the courage to confront him — something he said he will now do in earnest? He has never apologized for Jan. 6 and wishes to pardon those who attacked the Capitol and law enforcement, nor has he conceded that he lost in 2020. For these and many other reasons, he is no modern-day David. Tom Aydinian

Talk show low blow
Brooklyn: Jimmy Kimmel teared up on his show a day after Trump won the election. The reason, I believe, is because he realized all his bull against Trump not only had no detrimental effect on the election, but Trump did even better than expected. Then he went on to list everything under the sun that will go wrong under Trump. Where was Kimmel during Trump’s first term? Everything was better than it is now. Does Kimmel really believe his own falsehoods? Millions of voters obviously disagreed with everything he said. How out of touch is this guy? David Balsam

Just going with it?
Whitestone: Where is the outrage and shock? Nobody is calling for manual counts of the ballots. Nobody is questioning the security of the voting machines. The Republi-cons made sure that this election ran like a well-oiled machine. When something is too good to be true, it’s because it isn’t true. I hate to say it, but I guess we are just a bunch of dim, dumb Dems. Something is very wrong here. Robin Mazzia

Overwhelming support
North Massapequa, L.I.: The way the country was going in the last four years had a lot of people scared. I always felt that Trump would win it. I even thought it would be a sweep á la the Ronald Reagan election (49-state sweep). Imagine my surprise when it wasn’t. Seventy-five percent say the country is going in the wrong direction; 69% or better said inflation, the border and immigration were their three top concerns. And yet, the vote was pretty close to 50-50. I understand that Dems vote for Dems and Repubs vote for Repubs. But I have to say, if I was a Dem, I would have voted Republican because the country comes first, not politics. Steven Malichek

Family matters
Glendale: Voicers are starting to comment on the Daniel Penny trial. Until now, I have not seen one letter putting a big part of the blame where it belongs: on Jordan Neely’s family. In numerous interviews, members of his family have acknowledged that he had mental issues. As such, you can not rely on an individual to seek the help they need. Family members should have intervened and got him the help he needed. However, they didn’t. I know personally about the stigma that mental health issues present, however, people with these issues must get help. If they don’t, there are often dire consequences. The family failed to pursue help for him and they allowed him to walk the streets for years untreated. Thomas Murawski`
)

func flush(buffer []byte) {
	// This function should handle the logic to process or store the buffered data.
	// For the purpose of this question, you can assume it prints the buffer content.\
	fmt.Println(string(buffer))
}

func dataStream(msg <-chan byte) {
	// This function should read from the data stream (msg) and buffer the data.
	// When the buffer reaches maxBuffer size, it should call the flush function and reset the buffer.

	listenTimeout := 60 * time.Second
	timeBegin := time.Now()
	timeNow := timeBegin
	lastTimeFlush := timeBegin
	needFlush := false
	for timeNow.Sub(timeBegin) < listenTimeout {
		if len(msg) >= maxBuffer || timeNow.Sub(lastTimeFlush) > flushPeriod {
			needFlush = true
		}

		if needFlush {
			var buffer []byte
			for len(msg) > 0 {
				buffer = append(buffer, <-msg)
			}
			flush(buffer)
			lastTimeFlush = timeNow
			needFlush = false
		}

		timeNow = time.Now()
	}
}

func randomSendData(msg chan<- byte) {
	byteArticle := []byte(article)
	for i := 0; i < len(byteArticle); i++ {
		// fmt.Printf("Send Data: %v\n", byteArticle[i])
		msg <- byte(byteArticle[i])
	}
}

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	// Init channel
	msg := make(chan byte, 100)

	go randomSendData(msg)
	go dataStream(msg)

	fmt.Println("Waiting for goroutines to finish...")
	waitGroup.Wait()
	fmt.Println("Done!")
}
