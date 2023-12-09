package aoc23

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/require"
	"strconv"
	"strings"
	"testing"
)

const prodInput = `12 12 11 5 -5 -5 40 194 558 1278 2553 4643 7877 12661 19486 28936 41696 58560 80439 108369 143519
6 12 33 79 163 320 638 1301 2644 5220 9879 17859 30889 51304 82172 127433 192050 282172 405309 570519 788607
3 5 5 3 9 60 262 872 2445 6091 13915 29744 60271 116756 217404 390473 678031 1140057 1858241 2938353 4509389
15 21 23 31 67 161 350 681 1224 2114 3674 6758 13659 30354 71646 172084 407585 935627 2065861 4378033 8910071
13 32 75 155 286 477 724 1011 1347 1894 3301 7501 19549 51737 132449 324337 759840 1707391 3689542 7685418 15464987
9 23 59 129 245 419 663 989 1409 1935 2579 3353 4269 5339 6575 7989 9593 11399 13419 15665 18149
29 57 93 134 193 323 653 1442 3171 6715 13675 27029 52429 100813 193667 373531 724690 1413358 2764818 5410241 10564342
26 44 74 139 286 598 1217 2393 4581 8615 15995 29330 52987 94003 163324 277442 460508 747006 1185080 1840613 2802164
26 53 105 196 345 578 945 1572 2782 5357 11103 24058 52986 116265 250950 528745 1082962 2152411 4150601 7771330 14143463
11 17 32 59 96 149 260 550 1277 2909 6212 12353 23018 40545 68072 109700 170671 257561 378488 543335 763988
4 14 38 95 219 477 996 2010 3944 7562 14241 26533 49411 93078 179146 353714 714029 1461171 3001653 6139588 12434184
15 22 23 11 -12 -12 108 544 1629 3870 7981 14909 25850 42252 65802 98394 142075 198966 271155 360559 468752
5 2 16 76 220 504 1038 2072 4176 8596 17934 37408 77123 156092 309358 600875 1147609 2166120 4064281 7621931 14341202
-1 1 8 24 59 128 245 412 603 743 682 164 -1209 -4018 -9073 -17464 -30617 -50355 -78964 -119264 -174685
15 28 45 59 59 30 -39 -128 -117 372 2311 7951 22198 55382 128880 286664 619266 1313842 2758770 5760511 11986520
26 45 68 95 126 161 200 243 290 341 396 455 518 585 656 731 810 893 980 1071 1166
21 28 35 42 49 56 63 70 77 84 91 98 105 112 119 126 133 140 147 154 161
-10 -3 18 55 110 183 282 470 990 2532 6735 17050 40138 88089 182041 358469 680869 1262429 2310889 4218237 7742573
9 25 62 140 283 520 893 1476 2410 3967 6676 11577 20712 38013 70824 132495 247148 459774 860531 1646359 3275530
2 7 33 101 237 478 891 1605 2851 4997 8554 14127 22322 33756 49668 74416 122776 238151 533733 1279139 3072535
-2 0 2 4 6 8 10 12 14 16 18 20 22 24 26 28 30 32 34 36 38
5 14 23 32 41 50 59 68 77 86 95 104 113 122 131 140 149 158 167 176 185
16 17 26 66 184 474 1111 2406 4904 9559 18032 33170 59736 105472 182589 309790 514944 838541 1338070 2093474 3213848
22 42 80 153 295 584 1179 2360 4562 8392 14616 24101 37695 56026 79199 106368 135158 160910 175720 167241 117215
6 21 52 109 211 410 842 1824 4018 8687 18079 35998 68666 126066 224107 388190 659121 1103000 1828405 3018878 5001220
13 19 29 60 142 321 671 1317 2472 4493 7959 13770 23275 38492 62639 101532 167031 284763 509985 957860 1857820
4 16 44 100 215 458 963 1970 3893 7448 13914 25677 47362 88170 166637 320117 623138 1221760 2395656 4666442 8980537
2 2 8 33 108 302 764 1796 3969 8301 16530 31548 58139 104333 184043 322376 566444 1007252 1823320 3365719 6319630
20 35 50 65 80 95 110 125 140 155 170 185 200 215 230 245 260 275 290 305 320
11 6 9 40 133 345 767 1544 2928 5410 10005 18810 36050 70029 136805 267146 519585 1004422 1926633 3660228 6873123
6 19 38 76 156 311 587 1050 1807 3079 5427 10362 21840 49715 117435 278786 653581 1502158 3377441 7427999 15987265
14 14 16 17 17 38 166 642 2053 5705 14305 33159 72242 149761 298273 575110 1079888 1983336 3573686 6329541 11031623
12 26 57 122 266 578 1222 2499 4965 9644 18400 34594 64321 118960 220828 416129 804537 1605237 3299590 6928305 14688071
0 -5 -12 -6 50 226 638 1467 3007 5803 10999 21106 41550 83628 169980 344509 690014 1358827 2624642 4967626 9209846
20 45 79 127 202 336 599 1133 2233 4548 9533 20373 43778 93462 197102 410802 849837 1753978 3623782 7496926 15490289
21 33 41 35 8 -22 24 336 1279 3471 7875 15905 29546 51488 85274 135462 207801 309421 449037 637167 886364
4 -1 3 28 96 246 549 1142 2299 4579 9136 18348 37021 74554 148642 291444 559886 1053375 1945513 3542820 6395190
17 31 59 128 287 617 1241 2334 4133 6947 11167 17276 25859 37613 53357 74042 100761 134759 177443 230392 295367
-2 0 10 35 92 221 504 1106 2380 5121 11131 24392 53376 115420 244780 507138 1025234 2023270 3902180 7364189 13613652
7 28 59 110 201 362 642 1135 2039 3782 7291 14576 30002 63005 133680 283769 597275 1237420 2510163 4969210 9583571
17 45 102 216 429 794 1377 2284 3749 6346 11433 22006 44243 90159 181980 359084 688657 1281579 2315496 4067556 6959897
19 30 47 77 136 268 584 1323 2947 6316 13053 26303 52207 102537 199047 380154 710527 1293974 2289607 3930549 6543326
23 42 77 151 312 652 1344 2723 5462 10940 21981 44288 89155 178518 354335 696191 1355093 2620278 5055188 9780558 19071368
22 43 85 169 330 623 1125 1940 3233 5344 9072 16297 31274 63266 131795 276824 577804 1187897 2392963 4709153 9038134
6 25 52 81 110 156 289 699 1812 4473 10216 21643 42936 80528 143961 246961 408762 655713 1023204 1557949 2320666
15 35 85 188 376 704 1279 2318 4262 8001 15322 29792 58445 114870 224610 434194 825649 1538991 2805987 4999428 8703270
12 26 63 145 318 666 1325 2497 4464 7602 12395 19449 29506 43458 62361 87449 120148 162090 215127 281345 363078
12 24 51 112 233 444 787 1349 2331 4164 7692 14478 27402 52029 99979 197196 403379 855143 1856587 4059742 8808766
5 -3 -17 -26 -9 68 265 696 1577 3299 6531 12358 22459 39330 66557 109144 173901 269897 408983 606390 881407
-7 -14 -27 -34 -3 121 426 1072 2394 5113 10699 21916 43546 83232 152294 266252 444631 709420 1081305 1572490 2174555
15 43 93 185 353 658 1220 2275 4266 7980 14746 26712 47228 81404 137068 226834 373339 623167 1083249 2012077 4037277
0 10 35 75 130 200 285 385 500 630 775 935 1110 1300 1505 1725 1960 2210 2475 2755 3050
10 7 13 38 104 258 580 1190 2267 4102 7216 12583 22007 38711 68205 119509 206816 351689 585895 954988 1522762
14 33 72 154 329 681 1337 2481 4373 7371 11953 18736 28490 42147 60808 85755 118480 160749 214726 283190 369887
10 22 37 58 102 222 555 1408 3394 7630 16009 31558 58894 104790 178863 294396 469306 727270 1099021 1623826 2351158
9 19 35 74 171 395 876 1842 3670 6967 12723 22632 39777 70040 124845 226183 415307 767015 1412039 2570690 4601515
21 44 71 100 145 246 471 903 1598 2501 3333 3537 2528 775 2709 23172 100187 318346 849211 2017954 4409203
12 24 44 89 199 459 1027 2166 4295 8108 14865 27040 49624 92533 174770 331245 623478 1155808 2099218 3725475 6454989
2 -7 -16 -13 19 102 260 509 851 1308 2063 3808 8430 20191 47573 105960 221312 434947 809482 1435889 2441493
9 15 25 31 30 29 45 106 274 746 2170 6491 19008 53026 139755 348272 823883 1858721 4016707 8347127 16738350
8 18 37 63 96 148 259 525 1160 2632 5936 13115 28260 59496 123024 251327 509416 1025822 2049351 4048927 7882776
20 39 77 146 258 425 659 972 1376 1883 2505 3254 4142 5181 6383 7760 9324 11087 13061 15258 17690
6 14 22 30 38 46 54 62 70 78 86 94 102 110 118 126 134 142 150 158 166
3 4 17 55 134 278 530 986 1876 3717 7556 15300 30079 56470 100149 165993 252585 339102 357123 134181 -713180
9 28 61 108 169 244 333 436 553 684 829 988 1161 1348 1549 1764 1993 2236 2493 2764 3049
3 18 54 136 316 696 1460 2922 5603 10356 18564 32442 55480 93070 153366 248432 395739 620078 955962 1450596 2167500
-1 -6 -16 -31 -51 -76 -106 -141 -181 -226 -276 -331 -391 -456 -526 -601 -681 -766 -856 -951 -1051
24 38 72 156 345 744 1544 3065 5802 10470 18044 29790 47283 72408 107340 154499 216476 295926 395424 517280 663309
17 15 6 -4 -4 18 71 157 267 377 444 402 158 -412 -1467 -3205 -5867 -9741 -15166 -22536 -32304
-1 -7 -6 24 123 355 814 1630 2975 5069 8186 12660 18891 27351 38590 53242 72031 95777 125402 161936 206523
-1 -2 9 54 179 484 1175 2641 5555 11000 20650 37140 65020 113239 201182 372347 723723 1466606 3049271 6399370 13393725
10 27 50 80 122 206 429 1035 2562 6112 13856 29985 62474 126253 248691 478710 901370 1660417 2992078 5274334 9097018
-1 11 30 52 73 89 96 90 67 23 -46 -144 -275 -443 -652 -906 -1209 -1565 -1978 -2452 -2991
9 14 14 9 -1 -16 -36 -61 -91 -126 -166 -211 -261 -316 -376 -441 -511 -586 -666 -751 -841
20 41 85 176 348 645 1121 1840 2876 4313 6245 8776 12020 16101 21153 27320 34756 43625 54101 66368 80620
13 32 69 135 242 401 624 948 1523 2853 6354 15511 38128 90599 206059 450263 953223 1973288 4027961 8160817 16481924
10 31 61 97 128 138 112 40 -86 -287 -629 -1277 -2576 -5164 -10122 -19166 -34886 -61037 -102887 -167627 -264848
14 17 31 78 193 442 963 2041 4230 8539 16713 31679 58311 104817 185272 324077 563270 973281 1666119 2805592 4599309
10 29 60 114 216 405 734 1270 2094 3301 5000 7314 10380 14349 19386 25670 33394 42765 54004 67346 83040
10 20 42 84 148 237 375 654 1344 3147 7758 19034 45307 103804 228952 487908 1009562 2038446 4033814 7848568 15045322
-1 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19
14 28 47 75 117 189 352 786 1925 4693 10936 24277 51898 108293 223074 456955 934168 1907958 3892569 7927555 16100571
11 26 61 121 206 311 426 536 621 656 611 451 136 -379 -1144 -2214 -3649 -5514 -7879 -10819 -14414
-5 -12 -19 -8 64 280 780 1778 3577 6580 11295 18332 28390 42232 60646 84390 114119 150292 193057 242112 296540
10 22 55 123 235 396 610 896 1343 2257 4502 10223 24304 57260 130991 290283 625656 1317884 2722249 5522607 11003169
15 34 74 157 328 675 1362 2685 5158 9634 17470 30761 52712 88343 146085 241799 409125 725354 1372886 2777245 5900636
8 13 26 48 75 110 201 519 1486 3952 9411 20259 40163 74771 133302 232076 401852 702023 1246368 2247290 4088395
19 29 54 103 184 305 479 736 1145 1849 3116 5409 9478 16477 28109 46802 75919 120005 185074 278939 411588
7 14 46 113 220 360 512 663 885 1527 3638 9822 25844 63497 145661 315582 655247 1324668 2649638 5319095 10821008
14 38 75 134 234 404 683 1120 1774 2714 4019 5778 8090 11064 14819 19484 25198 32110 40379 50174 61674
6 17 30 45 62 81 102 125 150 177 206 237 270 305 342 381 422 465 510 557 606
17 30 43 56 69 82 95 108 121 134 147 160 173 186 199 212 225 238 251 264 277
20 45 86 162 321 650 1279 2392 4284 7556 13634 25944 52286 109253 229996 477362 965640 1896183 3615554 6710326 12161299
22 43 74 127 231 436 817 1478 2556 4225 6700 10241 15157 21810 30619 42064 56690 75111 98014 126163 160403
-5 -11 -6 33 134 321 597 920 1183 1225 922 435 726 4493 17721 52098 128603 282637 571138 1082197 1947774
-1 13 46 114 256 554 1158 2324 4482 8360 15199 27103 47577 82315 140309 235359 388073 628455 999188 1559728 2391334
3 6 27 78 172 331 603 1091 1998 3696 6836 12532 22677 40485 71402 124593 215293 368410 623889 1044490 1726802
-2 8 39 115 276 592 1198 2367 4640 9025 17262 32133 57783 100020 166590 267488 415478 627162 925165 1342291 1928850
0 2 13 50 146 365 834 1815 3848 7994 16195 31751 59902 108511 188892 316940 514928 814674 1263289 1933440 2941050
19 30 47 87 184 411 913 1963 4074 8241 16458 32768 65278 129837 256481 500365 959825 1805572 3326981 6004209 10618712
13 22 51 111 213 368 587 881 1261 1738 2323 3027 3861 4836 5963 7253 8717 10366 12211 14263 16533
14 36 82 179 369 716 1320 2347 4102 7202 12948 24049 45917 88830 171350 325485 604198 1091992 1919438 3282663 5468977
12 23 54 129 280 556 1040 1871 3280 5671 9810 17227 30988 57056 106532 199149 368484 669455 1188782 2059213 3478448
14 30 65 140 301 628 1246 2345 4234 7488 13297 24192 45422 87455 170546 333427 650686 1268826 2482285 4893464 9745339
9 5 2 0 -1 -1 0 2 5 9 14 20 27 35 44 54 65 77 90 104 119
9 9 22 66 170 378 760 1437 2627 4719 8382 14716 25452 43208 71808 116671 185277 287717 437334 651462 952270
1 10 22 43 86 165 280 402 492 634 1454 5170 17912 54406 146767 360018 818055 1745102 3529210 6815985 12642382
2 1 15 57 140 277 481 765 1142 1625 2227 2961 3840 4877 6085 7477 9066 10865 12887 15145 17652
7 23 42 62 80 106 191 465 1176 2725 5728 11236 21446 41581 84146 177518 381833 816415 1704562 3443358 6708296
-5 -10 -15 -2 67 250 630 1350 2723 5496 11405 24252 51884 109706 226850 457196 899887 1739416 3326840 6346565 12156399
23 48 87 144 239 428 831 1679 3403 6805 13383 25947 49787 94871 179903 339608 637391 1186606 2185143 3969978 7100823
15 29 50 73 93 105 104 85 43 -27 -130 -271 -455 -687 -972 -1315 -1721 -2195 -2742 -3367 -4075
2 3 17 68 194 445 878 1557 2582 4212 7229 13841 29704 68185 159030 365549 817922 1776213 3746497 7688998 15380750
3 11 38 90 174 306 530 949 1763 3309 6119 11083 19977 36985 72588 152664 338527 769219 1747968 3914288 8578188
-6 -6 -9 -24 -64 -146 -291 -524 -874 -1374 -2061 -2976 -4164 -5674 -7559 -9876 -12686 -16054 -20049 -24744 -30216
21 31 41 51 61 71 81 91 101 111 121 131 141 151 161 171 181 191 201 211 221
-1 16 46 89 145 214 296 391 499 620 754 901 1061 1234 1420 1619 1831 2056 2294 2545 2809
18 44 85 152 270 495 943 1831 3530 6630 12017 20962 35222 57153 89835 137209 204226 297008 423021 591260 812446
-5 5 21 44 87 181 381 772 1475 2653 4517 7332 11423 17181 25069 35628 49483 67349 90037 118460 153639
4 18 36 60 104 201 423 941 2174 5103 11855 26687 57522 118241 232128 437464 797808 1422972 2513774 4457062 8021476
8 14 20 26 32 38 44 50 56 62 68 74 80 86 92 98 104 110 116 122 128
11 19 50 132 303 612 1136 2036 3701 7075 14335 30195 64259 135042 276529 548455 1051871 1952019 3511080 6133990 10431247
7 16 37 80 156 284 501 875 1521 2620 4441 7366 11918 18792 28889 43353 63611 91416 128893 178588 243520
0 -3 -9 -20 -27 17 236 880 2381 5424 11075 21076 38530 69368 125226 228674 422142 782391 1442991 2628004 4700939
-1 -3 -11 -20 -2 107 417 1090 2343 4465 7910 13625 23959 44846 90603 193922 426035 936665 2032143 4326166 9026232
25 43 70 108 159 225 308 410 533 679 850 1048 1275 1533 1824 2150 2513 2915 3358 3844 4375
29 56 102 171 260 355 442 553 888 2105 5976 16803 44319 109320 254050 560467 1181039 2388750 4655642 8772595 16027278
14 29 61 114 190 288 416 632 1131 2402 5508 12622 28134 61022 129927 273787 573436 1193922 2464328 5022892 10072708
24 48 89 149 229 341 531 920 1786 3741 8114 17753 38629 82880 174267 357373 712124 1375086 2568051 4631977 8059373
28 53 89 149 256 445 773 1353 2440 4614 9134 18600 38212 78278 159444 323857 658916 1346683 2763335 5672003 11582776
22 36 55 93 188 424 974 2176 4654 9496 18501 34507 61812 106700 178084 288278 453910 696988 1046131 1537977 2218780
9 16 35 66 109 164 231 310 401 504 619 746 885 1036 1199 1374 1561 1760 1971 2194 2429
7 34 79 141 215 300 418 662 1312 3089 7663 18598 43007 94291 196407 390061 740889 1348804 2355843 3946459 6328459
16 43 79 128 206 343 590 1040 1884 3550 7030 14611 31426 68595 149379 321041 677640 1403961 2860225 5744362 11398390
14 29 44 63 101 189 385 796 1618 3205 6184 11641 21413 38533 67889 117174 198222 328845 535308 855603 1343709
22 33 38 38 43 89 268 777 2008 4730 10464 22238 46055 93648 187480 369537 716338 1363847 2547738 4666878 8381123
3 1 -6 -17 -20 16 153 502 1264 2839 6097 12996 27912 60419 131050 283274 608588 1299433 2760990 5846246 12344605
7 4 0 -3 9 73 260 686 1523 3010 5464 9291 14997 23199 34636 50180 70847 97808 132400 176137 230721
15 24 33 42 51 60 69 78 87 96 105 114 123 132 141 150 159 168 177 186 195
10 23 42 67 98 135 178 227 282 343 410 483 562 647 738 835 938 1047 1162 1283 1410
-7 0 30 101 246 534 1118 2324 4794 9689 18941 35521 63683 109189 179666 285540 442457 676712 1035878 1607431 2548622
22 40 60 81 102 122 140 155 166 172 172 165 150 126 92 47 -10 -80 -164 -263 -378
7 10 25 63 148 346 822 1948 4503 10033 21475 44195 87661 168129 313132 569616 1020054 1816286 3251854 5914783 11001561
12 37 73 126 216 388 741 1481 2995 5944 11395 21063 37819 66741 117141 206188 364958 647967 1147469 2014012 3484922
7 18 56 148 348 749 1503 2863 5264 9471 16847 29839 52865 93979 168180 304433 561213 1061164 2066810 4140202 8466114
22 33 40 43 42 37 28 15 -2 -23 -48 -77 -110 -147 -188 -233 -282 -335 -392 -453 -518
-3 6 27 55 83 104 121 170 369 1026 2871 7521 18343 41948 90629 186148 365381 688446 1250067 2195067 3739035
7 9 11 13 15 17 19 21 23 25 27 29 31 33 35 37 39 41 43 45 47
-3 0 5 11 30 112 376 1047 2505 5370 10685 20337 38026 71475 137425 272771 558853 1172912 2493441 5308255 11210950
22 34 55 105 214 430 852 1707 3503 7308 15223 31132 61827 118634 219734 393526 683684 1157110 1916903 3123907 5032564
18 41 79 139 243 454 922 1962 4181 8676 17330 33238 61300 109023 187579 313171 508764 806243 1249065 1895477 2822377
4 5 2 4 33 132 396 1039 2509 5669 12077 24427 47279 88371 161196 290359 522838 951137 1759080 3308474 6296073
0 5 10 13 12 5 -10 -35 -72 -123 -190 -275 -380 -507 -658 -835 -1040 -1275 -1542 -1843 -2180
6 10 24 65 167 388 832 1698 3368 6546 12460 23139 41777 73196 124420 205372 329706 515786 787824 1177189 1723899
2 4 13 37 81 147 234 338 452 566 667 739 763 717 576 312 -106 -712 -1543 -2639 -4043
22 42 80 143 236 364 532 752 1092 1855 4074 10682 29028 75996 188102 441042 986962 2124308 4426060 8969637 17738708
1 21 68 165 348 663 1158 1870 2807 3925 5100 6095 6522 5799 3102 -2688 -13043 -29851 -55484 -92871 -145576
-4 8 30 57 88 138 255 552 1273 2921 6498 13965 29157 59627 120297 240426 476338 933670 1804694 3428641 6386022
14 31 67 132 237 398 644 1030 1659 2722 4572 7857 13748 24311 43087 75961 132420 227321 383313 634082 1028615
6 26 68 145 270 456 716 1063 1510 2070 2756 3581 4558 5700 7020 8531 10246 12178 14340 16745 19406
2 7 34 95 210 412 754 1319 2231 3677 5992 9956 17655 34667 75128 172707 403128 928287 2079158 4505833 9439872
0 9 36 99 226 460 861 1497 2419 3637 5174 7410 12211 25899 66166 176887 457954 1122624 2604098 5749210 12165804
6 6 19 57 139 301 608 1173 2205 4143 7993 16078 33541 71108 149839 308981 618988 1204385 2287965 4284050 8004079
25 39 57 77 99 137 242 542 1306 3037 6586 13243 24686 42535 67041 94112 109407 77579 -76121 -495799 -1441795
-4 -3 -1 2 6 11 17 24 32 41 51 62 74 87 101 116 132 149 167 186 206
9 32 68 125 216 359 577 898 1355 1986 2834 3947 5378 7185 9431 12184 15517 19508 24240 29801 36284
-1 0 15 72 226 570 1250 2490 4644 8329 14783 26775 50723 101217 209976 442483 927245 1904934 3807711 7382962 13880640
16 29 59 109 177 268 417 737 1517 3403 7699 16823 34943 68799 128690 229574 392202 644197 1021015 1566813 2335433
17 37 70 121 195 297 432 605 821 1085 1402 1777 2215 2721 3300 3957 4697 5525 6446 7465 8587
4 19 44 84 148 249 404 634 964 1423 2044 2864 3924 5269 6948 9014 11524 14539 18124 22348 27284
15 25 41 84 188 414 876 1776 3442 6369 11288 19346 32603 55313 97001 179437 351665 720917 1513467 3189564 6652286
13 28 69 145 266 449 729 1176 1913 3133 5136 8461 14284 25402 48336 97374 201749 417618 847087 1667225 3172838
12 26 40 54 68 82 96 110 124 138 152 166 180 194 208 222 236 250 264 278 292
-2 -2 11 44 104 211 418 838 1695 3457 7185 15354 33581 73942 160886 341170 699756 1384240 2641135 4868216 8688166
13 21 44 98 214 454 945 1935 3865 7433 13608 23556 38505 59761 89474 133450 208449 358169 684696 1405842 2953780
2 1 -2 5 50 176 435 874 1521 2408 3715 6180 11993 26488 61106 138414 300596 624055 1244996 2404699 4529468
5 6 4 -2 3 78 363 1128 2842 6273 12631 23765 42422 72571 119788 191689 298387 452936 671710 974648 1385277
-3 13 57 143 285 497 793 1187 1693 2325 3097 4023 5117 6393 7865 9547 11453 13597 15993 18655 21597
21 30 39 51 71 116 236 543 1248 2720 5617 11216 22231 44753 92670 197402 428736 938386 2049431 4437414 9492734
10 17 26 46 90 192 443 1057 2489 5642 12219 25299 50243 96067 177454 317616 552260 934959 1544280 2493076 3940408
17 29 41 53 65 77 89 101 113 125 137 149 161 173 185 197 209 221 233 245 257
4 7 17 34 58 89 127 172 224 283 349 422 502 589 683 784 892 1007 1129 1258 1394
16 32 53 77 97 112 151 310 802 2020 4613 9575 18347 32932 56023 91144 142804 216664 319717 460481 649205
23 31 39 47 55 63 71 79 87 95 103 111 119 127 135 143 151 159 167 175 183
6 13 41 96 177 284 431 667 1117 2065 4123 8590 18260 39300 85591 188476 417909 927056 2044694 4465063 9636756
3 23 61 127 241 443 819 1550 2989 5769 10943 20155 35839 61441 101657 162678 252431 380803 559833 803855 1129573
20 39 77 152 285 492 786 1214 1975 3703 8068 18969 44794 102538 225043 473303 956718 1865447 3520671 6450710 11503627
11 33 69 125 207 321 473 669 915 1217 1581 2013 2519 3105 3777 4541 5403 6369 7445 8637 9951
13 26 51 107 227 475 984 2024 4111 8170 15767 29427 53057 92495 156208 256164 408905 636850 969859 1447091 2119191
10 18 45 115 265 550 1050 1871 3140 5026 7887 12759 22579 44782 96244 211971 459470 961394 1929841 3716619 6884877
9 11 13 15 17 19 21 23 25 27 29 31 33 35 37 39 41 43 45 47 49
-8 -12 -16 -20 -24 -28 -32 -36 -40 -44 -48 -52 -56 -60 -64 -68 -72 -76 -80 -84 -88
8 4 -3 0 52 227 648 1513 3146 6104 11404 20985 38592 71365 132539 245814 452140 819884 1459607 2544982 4341732
11 21 31 40 57 112 265 626 1411 3070 6534 13641 27820 55153 106059 198248 362834 658985 1211477 2306915 4631860
-3 13 43 87 145 217 303 403 517 645 787 943 1113 1297 1495 1707 1933 2173 2427 2695 2977
7 6 22 76 202 464 997 2093 4366 9061 18642 37934 76343 152077 299879 584606 1124137 2126792 3951214 7201596 12880296
-5 -12 -16 -9 23 117 374 1049 2736 6717 15564 34095 70784 139706 263056 474211 821201 1370314 2209376 3450015 5227933
17 21 25 29 33 37 41 45 49 53 57 61 65 69 73 77 81 85 89 93 97
6 13 32 80 180 368 717 1406 2880 6172 13500 29335 62308 128664 258590 507776 978143 1853856 3461417 6365291 11511885
5 23 49 86 154 303 625 1278 2553 5048 10072 20498 42428 88236 181828 367312 722717 1380949 2560835 4611894 8077398
`

// y = 3x
// y = 1x^2 +  + 1
const testInput = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

var inputs = [][2]string{
	{"testInput", testInput},
	{"prodInput", prodInput},
}

func part(t *testing.T, input string) {
	//var reconstructed [][]int
	var total int
	scanLines(t, input, skipBlankLines, func(line string) bool {
		nums := [][]int{mapSlice(func(in string) int {
			return must(t, strconv.Atoi, in)
		}, strings.Split(line, " "))}

		//t.Log(nums)

		for !all(nums[len(nums)-1], func(i int) bool { return i == 0 }) {
			lastRow := nums[len(nums)-1]
			diffs := make([]int, len(lastRow)-1)
			for i := 0; i < len(lastRow)-1; i++ {
				diffs[i] = lastRow[i+1] - lastRow[i]
			}
			nums = append(nums, diffs)
		}
		//t.Log(nums)

		reconstructed := [][]int{append([]int{0}, nums[len(nums)-1]...)}
		//t.Log(reconstructed)
		for i := 0; len(reconstructed) < len(nums); i++ {
			beforeReconstruction := nums[len(nums)-2-i]
			reconstructed = append(reconstructed, append([]int{beforeReconstruction[0] - reconstructed[i][0]}, beforeReconstruction...))
		}
		//t.Log(reconstructed)
		//
		rl := reconstructed[len(reconstructed)-1]
		total += rl[0]
		return true
	})

	t.Log(total)
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func TestPart1(t *testing.T) {
	for _, input := range inputs {
		fmt.Println(input[0])
		part(t, input[1])
		fmt.Println()
	}
}

func scanLines(t *testing.T, input string, processors ...func(string) bool) {
	reader := bufio.NewScanner(strings.NewReader(input))
	for reader.Scan() {
		line := reader.Text()
		var processed bool
		for _, p := range processors {
			if p(line) {
				processed = true
				break
			}
		}
		require.True(t, processed, "not all input was processed for line:", line)
	}
}

func skipBlankLines(line string) bool {
	return line == ""
}

func all[t any](values []t, predicate func(t) bool) bool {
	for _, value := range values {
		if !predicate(value) {
			return false
		}
	}
	return true
}

func mapSlice[in any, out any](mapIt func(in) out, inSlice []in) []out {
	outSlice := make([]out, len(inSlice))
	for i, inV := range inSlice {
		outSlice[i] = mapIt(inV)
	}
	return outSlice
}

func must[I any, O any](t *testing.T, fn func(I) (O, error), in I) O {
	t.Helper()
	o, err := fn(in)
	require.NoError(t, err)
	return o
}
