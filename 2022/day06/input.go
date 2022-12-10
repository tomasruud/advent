package main

const input = `dfsfmfbbbjnbbpddfcfjcjbjwjqqbtbntnhtnncfnfpnffwpphwppvbvtvztzszfzhhnqnvnpppmzmczcmzczbznnbssbhhvghhzqzvzttjvtvcccdhccmvccgdgttghthppwlwqlwwcswspsfsdfddhwwzlwzwttlglttsjjthhgcgfcfhfhtftrffwcfcsfccnntmmpvpgpfgfsgfflgfgmffbmmqsmqmmcqqmjjzczwczzpjzzgtgnnqqvdqvqllbttftgtztcztttlslrslrlfldljlwjjnvjjmrrhdrhhflhffgnfnjffmvmvjjspsvvrcchhgngwwdwbwwfhwffmggbsgspsjsfjjdwwbtbdtbdbgdbbbzhzvhhwpwlwfwfswwrgggmggssfwwdrwrccssjttltrrnggdbggvzzzrfzfwzfffnfnmnnzmnznpzzdtzdttbvvgffmnfmfmbbqppcgpgtgpgggcwwlpplslbslblhlmlfmlmldmmjhjjgllglhlssswcssprrqrjrhrvrffpfnpfnppzgztgtdgtdggftgftggqrgrnnqpnpgnpnggsjsvscsnswshwswfwfwqffcrrmprrrcgrrggmnnmzmwzmwzzgzzhnzzthtggtmmdrmrbmmqwwjswwphhzvvjqjrrvzrrmssdvdlvvlhvvjggbwbrwwjqjqmmfllttvpttdvtdtddrttpltpllzrlrwlrrvcvpcphcccqwwtpwwbmmtntfntftvfvfqvvwnvwwvhhlccvrcrlclcqcvqqsvvfjjqmmfhfvvctczzlhlsslpspfsswnsnzsspzpccmssmpmbpbmmvnvsnvsszttrppvgpgnncllrvlrvllcvlvwvhvghgvvnhvvmffclfccvwcvccrwrhhrffqrqcrqqhbqhqssmzsmzsmmzrztzllmclmmbgmbbrmrgrwrgwrrgngvvtqqpbqpppgbgccsjjcmmtdmdjmmjcchvchhbccnjngjjnwjnnrvnrrsqshshszzqnngbnnrnggvwwlzwwlqwqmqdqrdqqrhqhffpgffqgqdggfjgglfglgvgjvjmjtthghrhvhgvhvnncnwcnwntwntnjjqjwwrdwddsggfddnhnsnvsstbbjddfhddhbddwbwbnbrnnrhrqhqdhdllvbbnzzbmbttpjjngnqqcffrsrnsnnhzzgllgvlvcchfhzzmzrmrggdvgvgqqwnqqpdqqvjjvnvhnhhgvhhgllhlzlclbcccwppztzhhvmvzzzsbzbppvdvdtdtdvdsdlsscnnqhnhgngghrrzprzzpddmvvhcvvprplpnlppscctzthtptspttmftmftfjjdfdjfjrfrfbfcchmhnnbddwzwvvpvrvnnslnlnhlnhlltslttqpqvvgzvzsstcstsrtrbtbbzmmzrzqrzqqnmqnnpjpttwgtgzgqqgmgnnzgzrggpbggvssqvssmhmshmssvlvlmvvnhnhddwbbllffgbffbbztthbhdhdghhrccfmfrfmmbdbfdbffzfgfrfqqptpgpjjlvvbjjdzdbbszslzlldnngwgddbmbpbwwhphnhmhhlthhgfhfvvpmvmhvmvdvsddsjtzvfmpsrwrrzgcvnnllfjmvfptwncppfmgqbfzrdpnfddghsqfmnqfwfslrsgjmqtfqwhdddsbhtbtpswcbfppcbhzfzbsqljzndcsrlhrrtstgfhhfsqqrwgnncsmstdmjvfjhqnrczlftzzzhqdzjdcdqcgfpmbqntdhzcvbtpssrvmgjwzwfvtpsrsrwrvrsjgrmzqzvbttscldsnnwzvmlztnnpdjrwvhshpdwgvhmlrnhtfccjnldlnhtfncfjjjztjmhrdqpvhggtqzwjsvwdzhdmwhsmgzjcwzqzlwbrlzsmlwhpjvflnppvrbgrsblmjpnqvgpjbpwbjgjqzwvjbgcplccjgbfwlblzfjqpwszbqbcnlbmfmqpmgspscgfdgfwnmcdzcqnjznndjcvlblszcnpflbjqltpfzhffdbwbshtpnwwlspltpcrvbdtflwbjrfnvrflqpgqtjzqwmmsdvtsmgjtrtbrzchwhpfsznjqcbrjcvwqgrcsqpvfzhrdlmnvvhjzpgpnlrmqfvcnlrlcfjblfcgvngdjfdczsrtnnwjndsfcsdlhdnbtplfnhsmmbldmsjwcblghhgqwbnjvqbqhddrmrtncvwnchsfpddzgrrtzntmwnmdwlrvnjgnzjqvptztnqnqmcjmmrhmtstgdvhffbbmphnbtsdmpmscsfdbnfnchrhhjpsfhhswfszgqfcbdbgnrqhrflpfgfgdcrjvrwbvsfmzzhvzvqzgshcqzlfcljnlgshdlhwdchhhvwlshwdrgjfbnptqqglbpcfgrmqjhqvlbzdwgnzfzlpcpjzqwhbfjljszvjdsrmfzntgnjflhnwhpfrlbpvgzmbqwzlgphmbvbfdqfgqqbhzzvrjftnwjzhlrqccwcfzvntscnbfcsrqqnvlvhszpgwtrzjrqbtslctbhtbczwtmsgwczncbjmzqvnthpwjmsbsjnfpsmghnvtqjjnjfwtnmlthrlcpqhjpnvnbbnwrdjfshwhpdwmsbngfhbhsqphlqspcgzwrgfjmqqtlsfgnvqtdgnhmdvvqzjwlhsnvjczbssrnlhwdmdthmtprjjfttfzbbswfwsvvslfnbcvprzhtcqwdrzjrnjjctqfsjrsddlhzcnstqfppjlqhvcbjbfwndwdtdfvnlwgvdvhzzrqthdhdmddfdwschmpwwrnlgsldzhgjrlmtzrnrrtqfctvbncpcjlsvnwjvhgnbshhwlqhtjghvrctlvngjgjrlgshhwscrdvzjqtfrrbssvqlcjjdljpmlzfqqnqmffpsbvgcqzqdjwczbqwjgfvpdgjglnqdshppcsqcmszhrbcpnhjnczlhwsbnfnzsczjvmftngqcvhgpgwlzbmjnmmdvdjfrcwnjrncvfstrvvqsstphmqdpwjqhzgppmgwlgjhgfwqgmjrlsvqpfqznvbqzgtngvpbpttzvngjwtrjgdnvdggzlnmpgbzhtnfnrhgwvhnpqdwfdvvftzllpqblgdglclwtwbchlvwcmmvtchlntlghztfvgfjczcbqmzgnqmrjcmqvjmjhfpztjcvclblmrctzfmsdfvfpwdcbsglgrsjqcgtcblhbtgcgjlwhhqnwhdnwzlhvphtvmlfnbfmgpnnbzqtdtzqrbhfljlwstlbscnrsvhrbrcthvzcngrttddcjqhtmsfpsgldgtsgjtprsttlssmrrfjmrddqvnqcfmphbnjtdsqvptrdzqbqfjqtnrqtjgdpbrlzrlvwcbcqbcmncfmwcpdhgpjdrdcmrqnflsqbllrqslmhsljwghnwcjwvhchcnlgppmphbqtcdfzjpcqcgsjzvmgfjgfsvmvjfqvtpffbpmhnnnrjmqhhhrhrqfqdwzdvzssslzvqhngdpszztgrvjntcpzhbfmhbpvcndsjbtnwgpmztpbrtjmfqrsvndrspdqmlsbldgghfflncszhnsttfslvwhvfnsmmhbbvjqslsjfqplndndwmbmvgchgvhzclrcnhvgbgmpctrggvqpvqvgvncmdwhpmwhzwhlgsnlnwggfbbvdvqrrsmhwzrrpgrjfshzgzjpfwjhpqmqhbjlwhwsfszlshpzprvgprlprrvlcrmbttjrpqsrdcdfwdrzbcfjpvrlrjjdwhbspqmrblvtldqdhtjtjphpqswgvqfftdgqrtjgsmthlhvlcqrwlqtthwjgrcpwcnsqtssqzpzqptrwjjdfchfmmsrsccnlvqbdmbcdjmhpgvnnlttfhggfphvbwqtcztbnsflztcfpbcpjbcmsplhjdbsmzhgnmfrhscmwmfqbljvhgllvvgqzphzbswdzlhmpcvnntczrcnqvlphhjdjjjnhfzzcjjsdlfccwvswvjfgvmlnpvjvcbpglsgtpj`