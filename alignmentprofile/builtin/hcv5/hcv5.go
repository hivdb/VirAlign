package hcv5

import (
	ap "github.com/hivdb/nucamino/alignmentprofile"
	a "github.com/hivdb/nucamino/types/amino"
)

var hcv5PositionalIndelScores = ap.GenePositionalIndelScores{}

// HCV5 Reference Sequences; Genebank accession ID: AF064490.1

var (
	HCV5SEQ_NS3 = a.ReadString(`
		APITAYAQQTRGVLGAIIVSLTGRDKNEAEGEVQVLSTATQTFLGTCINGVMWTVFHGAGAKTLAGPKGP
		VVQMYTNVDKDLVGWPTPPGTRSLTPCTCGSADLYLVTRHADVVPARRRGDTRASLLSPRPISYLKGSSG
		GPVMCPSGHVVGVFRAAVCTRGVAKALDFIPVENLETTMRSPVFTDNSTPPAVPHEFQVGHLHAPTGSGK
		STKVPAAYAAQGYKVLVLNPSVAATLGFGAYMSRAYGVDPNIRTGVRTVTTGAAITYSTYGKFLADGGCS
		GGAYDVIICDECHSQDATTILGIGTVLDQAETAGARLVVLATATPPGSVTTPHPNIEEVALPSEGEIPFY
		GRAIPLALIKGGRHLIFCHSKKKCDELAKQLTSQGVNAVAYYRGLDVAVIPATGDVVVCSTDALMTGFTG
		DFDSVIDCNTTVTQTVDFSLDPTFTIETTTVPQDAVSRSQRRGRTGRGRHGIYRYVSSGERPSGIFDSVV
		LCECYDAGCAWYDLTPAETTVRLRAYLNTPGLPVCQDHLEFWEGVFTGLTNIDAHMLSQTKQGGENFPYL
		VAYQATVCVRAKAPPPSWDTMWKCMLRLKPTLTGPTPLLYRLGAVQNEITLTHPITKYIMACMSADLEVI
		T
	`)
	HCV5SEQ_NS5A = a.ReadString(`
		DGTWLRAIWDWVCTALTDFKAWLQAKLLPQLPGVPFLSCQRGYRGVWRGDGVNSTKCPCGATISGHVKNG
		TMRIVGPKLCSNTWHGTFPINATTTGPSVPAPAPNYKFALWRVGAADYAEVRRVGDYHYITGVTQDNLKC
		PCQVPSPEFFTELDGVRIHRYAPPCNPLLREEVCFSVGLHSFVVGSQLPCEPEPDVTVLTSMLSDPAHIT
		AETAKRRLDRGSPPSLASSSASQLSAPSLKATCTTQGHHPDADLIEANLLWRQCMGGNITRVEAENKVVI
		LDSFEPLKADDDDREISVSADCFRRGPAFPPALPIWARPGYDPPLLETWKQPDYDPPQVSGCPLPPAGLP
		PVPPPRRKRKPVVLSDSNVSQVLADLAHARFKADTQSIEGQDSAVGTSSQPDSGPEEKRDDDSDAASYSS
		MPPLEGEPGDPDLSSGSWSTVSDEDSVVCC
	`)
	HCV5SEQ_NS5B = a.ReadString(`
		SMSYSWTGALITPCSAEEEKLPINPLSNTLLRHHNLVYSTSSRSAGQRQKKVTFDRLQVLDDHYREVVDE
		MKRLASKVKARLLPLEEACGLTPPHSARSKYGYGAKEVRSLDKKALNHIKGVWQDLLDDSDTPLPTTIMA
		KNEVFAVEPSKGGKKPARLIVYPDLGVRVCEKRALYDIAQKLPTALMGPSYGFQYSPAQRVEFLLKTWRS
		KKTPMAFSYDTRCFDSTVTEHDIMTEESIYQSCDLQPEARAAIRSLTQRLYCGGPMYNSKGQQCGYRRCR
		ASGVFTTSMGNTMTCYIKALASCRAAKLRDCTLLVCGDDLVAICESQGTHEDEASLRAFTEAMTRYSAPP
		GDPPVPAYDLELVTSCSSNVSVAHDASGNRVYYLTRDPQVPLARAAWETAKHSPVNSWLGNIIMYAPTLW
		ARIVLMTHFFSVLQSQEQLEKALAFEMYGSVYSVTPLDLPAIIQRLHGLSAFTLHSYSPSEINRVSSCLR
		KLGVPPLRAWRHRARAVRAKLIAQGGKAAICGIYLFNWAVKTKRKLTPLADADRLDLSSWFTVGAGGGDI
		YHSMSRARPRCILLCLLLLTVGVGIFLLPAR
	`)
)

var HCV5RefLookup = ap.ReferenceSeqs{
	"NS3":  HCV5SEQ_NS3,
	"NS5A": HCV5SEQ_NS5A,
	"NS5B": HCV5SEQ_NS5B,
}

var Profile = ap.AlignmentProfile{
	StopCodonPenalty:         4,
	GapOpeningPenalty:        10,
	GapExtensionPenalty:      2,
	IndelCodonOpeningBonus:   0,
	IndelCodonExtensionBonus: 2,
	GeneIndelScores:          hcv5PositionalIndelScores,
	ReferenceSequences:       HCV5RefLookup,
}
