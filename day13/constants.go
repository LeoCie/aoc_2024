package main

const example = `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=10000000008400, Y=10000000005400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=10000000012748, Y=10000000012176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=10000000007870, Y=10000000006450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=10000000018641, Y=10000000010279`

const input = `Button A: X+44, Y+17
Button B: X+28, Y+68
Prize: X=13736, Y=11226

Button A: X+45, Y+82
Button B: X+79, Y+11
Prize: X=8906, Y=6390

Button A: X+81, Y+11
Button B: X+84, Y+89
Prize: X=12933, Y=9438

Button A: X+13, Y+88
Button B: X+94, Y+72
Prize: X=6101, Y=6312

Button A: X+48, Y+24
Button B: X+14, Y+24
Prize: X=9182, Y=13400

Button A: X+95, Y+95
Button B: X+18, Y+82
Prize: X=3067, Y=6203

Button A: X+35, Y+16
Button B: X+19, Y+47
Prize: X=6861, Y=7675

Button A: X+38, Y+52
Button B: X+72, Y+25
Prize: X=2714, Y=2611

Button A: X+11, Y+84
Button B: X+84, Y+12
Prize: X=14128, Y=8120

Button A: X+80, Y+33
Button B: X+24, Y+79
Prize: X=4328, Y=8488

Button A: X+18, Y+48
Button B: X+23, Y+11
Prize: X=17816, Y=16130

Button A: X+23, Y+37
Button B: X+90, Y+28
Prize: X=1831, Y=1077

Button A: X+23, Y+59
Button B: X+57, Y+22
Prize: X=3673, Y=9424

Button A: X+43, Y+11
Button B: X+27, Y+61
Prize: X=11999, Y=2559

Button A: X+15, Y+52
Button B: X+78, Y+38
Prize: X=12116, Y=1704

Button A: X+13, Y+51
Button B: X+74, Y+34
Prize: X=18504, Y=1672

Button A: X+25, Y+52
Button B: X+43, Y+27
Prize: X=16970, Y=19095

Button A: X+11, Y+73
Button B: X+84, Y+87
Prize: X=2766, Y=4713

Button A: X+16, Y+50
Button B: X+69, Y+16
Prize: X=18250, Y=4500

Button A: X+57, Y+36
Button B: X+11, Y+30
Prize: X=3609, Y=5180

Button A: X+83, Y+57
Button B: X+12, Y+31
Prize: X=12528, Y=17116

Button A: X+51, Y+26
Button B: X+38, Y+66
Prize: X=12977, Y=13356

Button A: X+58, Y+32
Button B: X+21, Y+91
Prize: X=5071, Y=8595

Button A: X+72, Y+52
Button B: X+16, Y+86
Prize: X=6912, Y=7672

Button A: X+11, Y+50
Button B: X+87, Y+21
Prize: X=1783, Y=4360

Button A: X+28, Y+95
Button B: X+68, Y+29
Prize: X=3344, Y=4084

Button A: X+70, Y+29
Button B: X+14, Y+47
Prize: X=2958, Y=15887

Button A: X+15, Y+66
Button B: X+39, Y+13
Prize: X=2972, Y=7253

Button A: X+80, Y+72
Button B: X+14, Y+61
Prize: X=2892, Y=7346

Button A: X+22, Y+70
Button B: X+59, Y+12
Prize: X=1341, Y=16544

Button A: X+56, Y+14
Button B: X+58, Y+68
Prize: X=7046, Y=5988

Button A: X+48, Y+19
Button B: X+30, Y+69
Prize: X=3134, Y=9547

Button A: X+13, Y+69
Button B: X+84, Y+21
Prize: X=17327, Y=9221

Button A: X+61, Y+18
Button B: X+23, Y+58
Prize: X=5481, Y=7366

Button A: X+19, Y+75
Button B: X+64, Y+17
Prize: X=7118, Y=2036

Button A: X+77, Y+53
Button B: X+17, Y+42
Prize: X=4993, Y=6463

Button A: X+22, Y+91
Button B: X+83, Y+22
Prize: X=1304, Y=1538

Button A: X+72, Y+14
Button B: X+18, Y+82
Prize: X=16730, Y=7650

Button A: X+22, Y+69
Button B: X+58, Y+18
Prize: X=8266, Y=15413

Button A: X+91, Y+30
Button B: X+11, Y+98
Prize: X=4148, Y=8068

Button A: X+41, Y+97
Button B: X+95, Y+49
Prize: X=5233, Y=4823

Button A: X+27, Y+71
Button B: X+49, Y+38
Prize: X=5512, Y=6227

Button A: X+38, Y+59
Button B: X+48, Y+21
Prize: X=1508, Y=13802

Button A: X+20, Y+66
Button B: X+39, Y+16
Prize: X=5759, Y=9094

Button A: X+28, Y+49
Button B: X+61, Y+34
Prize: X=2331, Y=10929

Button A: X+36, Y+12
Button B: X+32, Y+48
Prize: X=6676, Y=380

Button A: X+29, Y+99
Button B: X+64, Y+15
Prize: X=5364, Y=7527

Button A: X+28, Y+76
Button B: X+66, Y+20
Prize: X=4834, Y=11484

Button A: X+16, Y+48
Button B: X+57, Y+33
Prize: X=13558, Y=16646

Button A: X+14, Y+67
Button B: X+97, Y+64
Prize: X=2118, Y=6134

Button A: X+56, Y+19
Button B: X+39, Y+77
Prize: X=15954, Y=948

Button A: X+78, Y+36
Button B: X+32, Y+87
Prize: X=4814, Y=7278

Button A: X+35, Y+82
Button B: X+92, Y+51
Prize: X=5417, Y=5945

Button A: X+22, Y+56
Button B: X+68, Y+37
Prize: X=11748, Y=13276

Button A: X+45, Y+90
Button B: X+97, Y+36
Prize: X=6416, Y=9198

Button A: X+24, Y+47
Button B: X+67, Y+18
Prize: X=4275, Y=4636

Button A: X+59, Y+83
Button B: X+27, Y+11
Prize: X=5178, Y=18866

Button A: X+29, Y+14
Button B: X+38, Y+59
Prize: X=3422, Y=4010

Button A: X+13, Y+67
Button B: X+74, Y+27
Prize: X=10476, Y=11200

Button A: X+16, Y+41
Button B: X+65, Y+32
Prize: X=4467, Y=17280

Button A: X+70, Y+29
Button B: X+16, Y+63
Prize: X=10028, Y=11676

Button A: X+43, Y+15
Button B: X+27, Y+41
Prize: X=16621, Y=14031

Button A: X+65, Y+37
Button B: X+18, Y+88
Prize: X=5842, Y=9468

Button A: X+95, Y+47
Button B: X+15, Y+76
Prize: X=8645, Y=9489

Button A: X+44, Y+87
Button B: X+84, Y+23
Prize: X=1776, Y=2510

Button A: X+63, Y+20
Button B: X+28, Y+69
Prize: X=5978, Y=3100

Button A: X+25, Y+53
Button B: X+32, Y+11
Prize: X=17829, Y=19516

Button A: X+22, Y+71
Button B: X+50, Y+12
Prize: X=10934, Y=18746

Button A: X+13, Y+27
Button B: X+61, Y+24
Prize: X=6987, Y=18278

Button A: X+21, Y+71
Button B: X+56, Y+14
Prize: X=14109, Y=17227

Button A: X+13, Y+47
Button B: X+75, Y+23
Prize: X=8672, Y=2174

Button A: X+12, Y+31
Button B: X+84, Y+62
Prize: X=3620, Y=16855

Button A: X+22, Y+76
Button B: X+54, Y+13
Prize: X=5094, Y=2499

Button A: X+17, Y+81
Button B: X+89, Y+84
Prize: X=3592, Y=6573

Button A: X+17, Y+55
Button B: X+71, Y+31
Prize: X=14465, Y=8867

Button A: X+68, Y+73
Button B: X+70, Y+11
Prize: X=5658, Y=3444

Button A: X+34, Y+65
Button B: X+55, Y+22
Prize: X=7050, Y=11058

Button A: X+42, Y+17
Button B: X+12, Y+57
Prize: X=4340, Y=3970

Button A: X+27, Y+47
Button B: X+37, Y+18
Prize: X=7500, Y=18692

Button A: X+13, Y+39
Button B: X+64, Y+34
Prize: X=7671, Y=1527

Button A: X+56, Y+20
Button B: X+11, Y+32
Prize: X=11587, Y=916

Button A: X+63, Y+15
Button B: X+32, Y+82
Prize: X=11332, Y=12954

Button A: X+51, Y+24
Button B: X+14, Y+38
Prize: X=5645, Y=18686

Button A: X+11, Y+37
Button B: X+49, Y+29
Prize: X=921, Y=789

Button A: X+36, Y+81
Button B: X+62, Y+15
Prize: X=428, Y=9890

Button A: X+11, Y+37
Button B: X+65, Y+15
Prize: X=10375, Y=1385

Button A: X+44, Y+31
Button B: X+14, Y+89
Prize: X=4416, Y=4694

Button A: X+45, Y+83
Button B: X+83, Y+48
Prize: X=10607, Y=11262

Button A: X+79, Y+12
Button B: X+12, Y+81
Prize: X=14104, Y=5297

Button A: X+66, Y+36
Button B: X+13, Y+47
Prize: X=190, Y=2382

Button A: X+25, Y+53
Button B: X+31, Y+15
Prize: X=18365, Y=3909

Button A: X+87, Y+38
Button B: X+12, Y+41
Prize: X=5127, Y=3777

Button A: X+75, Y+19
Button B: X+56, Y+77
Prize: X=6844, Y=6382

Button A: X+86, Y+44
Button B: X+16, Y+36
Prize: X=5230, Y=5040

Button A: X+41, Y+12
Button B: X+16, Y+70
Prize: X=17465, Y=18832

Button A: X+49, Y+68
Button B: X+62, Y+27
Prize: X=6998, Y=4634

Button A: X+84, Y+63
Button B: X+15, Y+99
Prize: X=7203, Y=7245

Button A: X+32, Y+63
Button B: X+41, Y+16
Prize: X=17232, Y=12724

Button A: X+28, Y+53
Button B: X+47, Y+16
Prize: X=17378, Y=18764

Button A: X+33, Y+15
Button B: X+12, Y+68
Prize: X=1773, Y=2307

Button A: X+37, Y+53
Button B: X+24, Y+11
Prize: X=11740, Y=11110

Button A: X+15, Y+28
Button B: X+57, Y+30
Prize: X=1571, Y=9214

Button A: X+22, Y+96
Button B: X+90, Y+64
Prize: X=2236, Y=3840

Button A: X+86, Y+16
Button B: X+13, Y+13
Prize: X=5952, Y=1192

Button A: X+61, Y+18
Button B: X+32, Y+37
Prize: X=7304, Y=4663

Button A: X+29, Y+60
Button B: X+64, Y+26
Prize: X=7221, Y=786

Button A: X+62, Y+27
Button B: X+17, Y+61
Prize: X=6605, Y=16984

Button A: X+98, Y+49
Button B: X+15, Y+98
Prize: X=7095, Y=10878

Button A: X+62, Y+34
Button B: X+22, Y+45
Prize: X=12276, Y=14000

Button A: X+17, Y+43
Button B: X+77, Y+49
Prize: X=2036, Y=16856

Button A: X+38, Y+33
Button B: X+20, Y+94
Prize: X=2626, Y=9867

Button A: X+24, Y+67
Button B: X+66, Y+28
Prize: X=5388, Y=6604

Button A: X+74, Y+89
Button B: X+11, Y+76
Prize: X=4761, Y=10936

Button A: X+15, Y+47
Button B: X+44, Y+17
Prize: X=12840, Y=6497

Button A: X+99, Y+78
Button B: X+34, Y+89
Prize: X=7254, Y=7395

Button A: X+62, Y+17
Button B: X+30, Y+80
Prize: X=4608, Y=14928

Button A: X+82, Y+46
Button B: X+36, Y+93
Prize: X=2766, Y=3663

Button A: X+13, Y+95
Button B: X+75, Y+51
Prize: X=6297, Y=10227

Button A: X+78, Y+21
Button B: X+43, Y+64
Prize: X=8779, Y=6505

Button A: X+71, Y+11
Button B: X+55, Y+87
Prize: X=1471, Y=1719

Button A: X+13, Y+43
Button B: X+27, Y+16
Prize: X=9344, Y=1441

Button A: X+22, Y+39
Button B: X+38, Y+22
Prize: X=3346, Y=4117

Button A: X+74, Y+46
Button B: X+14, Y+61
Prize: X=3812, Y=6658

Button A: X+12, Y+51
Button B: X+77, Y+23
Prize: X=9400, Y=13651

Button A: X+57, Y+16
Button B: X+17, Y+68
Prize: X=19280, Y=2976

Button A: X+51, Y+25
Button B: X+26, Y+63
Prize: X=3067, Y=17639

Button A: X+19, Y+75
Button B: X+52, Y+13
Prize: X=6759, Y=16632

Button A: X+11, Y+22
Button B: X+39, Y+21
Prize: X=3694, Y=3227

Button A: X+18, Y+25
Button B: X+97, Y+43
Prize: X=5555, Y=2854

Button A: X+76, Y+21
Button B: X+12, Y+72
Prize: X=2560, Y=3020

Button A: X+23, Y+78
Button B: X+71, Y+18
Prize: X=7519, Y=5690

Button A: X+47, Y+20
Button B: X+44, Y+70
Prize: X=15882, Y=1900

Button A: X+92, Y+58
Button B: X+23, Y+88
Prize: X=6049, Y=9914

Button A: X+14, Y+78
Button B: X+69, Y+60
Prize: X=3920, Y=8214

Button A: X+83, Y+53
Button B: X+13, Y+37
Prize: X=9099, Y=849

Button A: X+45, Y+93
Button B: X+97, Y+63
Prize: X=3639, Y=5871

Button A: X+25, Y+96
Button B: X+43, Y+26
Prize: X=1429, Y=5070

Button A: X+55, Y+32
Button B: X+14, Y+46
Prize: X=16108, Y=17064

Button A: X+37, Y+15
Button B: X+42, Y+57
Prize: X=9465, Y=2744

Button A: X+20, Y+65
Button B: X+87, Y+20
Prize: X=7990, Y=2320

Button A: X+11, Y+52
Button B: X+56, Y+22
Prize: X=19807, Y=5634

Button A: X+24, Y+40
Button B: X+67, Y+30
Prize: X=6816, Y=3520

Button A: X+99, Y+81
Button B: X+97, Y+11
Prize: X=14522, Y=7438

Button A: X+44, Y+29
Button B: X+16, Y+35
Prize: X=4532, Y=11498

Button A: X+54, Y+27
Button B: X+23, Y+42
Prize: X=14547, Y=10613

Button A: X+70, Y+43
Button B: X+13, Y+33
Prize: X=16727, Y=1479

Button A: X+19, Y+53
Button B: X+76, Y+35
Prize: X=6162, Y=12142

Button A: X+44, Y+70
Button B: X+47, Y+17
Prize: X=16784, Y=15030

Button A: X+37, Y+21
Button B: X+16, Y+72
Prize: X=2564, Y=1644

Button A: X+55, Y+15
Button B: X+36, Y+63
Prize: X=7936, Y=6738

Button A: X+44, Y+17
Button B: X+24, Y+52
Prize: X=14152, Y=2836

Button A: X+31, Y+18
Button B: X+23, Y+43
Prize: X=538, Y=18851

Button A: X+40, Y+13
Button B: X+13, Y+53
Prize: X=2482, Y=7391

Button A: X+84, Y+29
Button B: X+12, Y+61
Prize: X=8060, Y=14653

Button A: X+13, Y+30
Button B: X+49, Y+24
Prize: X=5132, Y=12200

Button A: X+18, Y+44
Button B: X+68, Y+33
Prize: X=3808, Y=4246

Button A: X+32, Y+52
Button B: X+53, Y+24
Prize: X=7078, Y=10780

Button A: X+11, Y+52
Button B: X+82, Y+38
Prize: X=5349, Y=2758

Button A: X+50, Y+16
Button B: X+31, Y+41
Prize: X=4751, Y=2173

Button A: X+36, Y+13
Button B: X+11, Y+56
Prize: X=5727, Y=16087

Button A: X+18, Y+90
Button B: X+93, Y+76
Prize: X=6861, Y=8242

Button A: X+42, Y+28
Button B: X+40, Y+99
Prize: X=6156, Y=7793

Button A: X+32, Y+95
Button B: X+84, Y+12
Prize: X=4040, Y=7721

Button A: X+15, Y+53
Button B: X+67, Y+22
Prize: X=3852, Y=4385

Button A: X+73, Y+22
Button B: X+17, Y+60
Prize: X=12463, Y=7634

Button A: X+47, Y+21
Button B: X+42, Y+92
Prize: X=3473, Y=5067

Button A: X+25, Y+71
Button B: X+58, Y+22
Prize: X=13586, Y=2782

Button A: X+16, Y+42
Button B: X+68, Y+42
Prize: X=5500, Y=1418

Button A: X+97, Y+90
Button B: X+19, Y+91
Prize: X=5033, Y=10246

Button A: X+57, Y+31
Button B: X+22, Y+63
Prize: X=3855, Y=2709

Button A: X+33, Y+71
Button B: X+60, Y+25
Prize: X=5786, Y=11777

Button A: X+66, Y+55
Button B: X+34, Y+99
Prize: X=7346, Y=9867

Button A: X+11, Y+48
Button B: X+44, Y+20
Prize: X=12299, Y=14760

Button A: X+33, Y+16
Button B: X+37, Y+60
Prize: X=4858, Y=5552

Button A: X+32, Y+46
Button B: X+31, Y+14
Prize: X=454, Y=8586

Button A: X+29, Y+12
Button B: X+24, Y+49
Prize: X=17339, Y=17815

Button A: X+63, Y+48
Button B: X+22, Y+70
Prize: X=2724, Y=6228

Button A: X+37, Y+81
Button B: X+66, Y+46
Prize: X=6413, Y=5077

Button A: X+95, Y+63
Button B: X+12, Y+56
Prize: X=5338, Y=5894

Button A: X+67, Y+43
Button B: X+12, Y+43
Prize: X=14331, Y=10454

Button A: X+95, Y+46
Button B: X+33, Y+51
Prize: X=1280, Y=970

Button A: X+41, Y+13
Button B: X+37, Y+66
Prize: X=14946, Y=2828

Button A: X+38, Y+87
Button B: X+52, Y+36
Prize: X=6692, Y=8760

Button A: X+68, Y+30
Button B: X+17, Y+54
Prize: X=4171, Y=13544

Button A: X+98, Y+53
Button B: X+31, Y+81
Prize: X=8023, Y=7358

Button A: X+25, Y+39
Button B: X+50, Y+25
Prize: X=11950, Y=16760

Button A: X+29, Y+57
Button B: X+47, Y+17
Prize: X=16869, Y=8959

Button A: X+14, Y+60
Button B: X+74, Y+16
Prize: X=4420, Y=11652

Button A: X+73, Y+33
Button B: X+11, Y+28
Prize: X=11353, Y=1006

Button A: X+17, Y+93
Button B: X+97, Y+44
Prize: X=4434, Y=5764

Button A: X+65, Y+38
Button B: X+14, Y+27
Prize: X=8171, Y=15034

Button A: X+48, Y+15
Button B: X+75, Y+88
Prize: X=7785, Y=8308

Button A: X+60, Y+21
Button B: X+26, Y+73
Prize: X=14588, Y=15227

Button A: X+16, Y+41
Button B: X+25, Y+17
Prize: X=16174, Y=8643

Button A: X+78, Y+83
Button B: X+18, Y+95
Prize: X=2334, Y=8703

Button A: X+89, Y+34
Button B: X+77, Y+98
Prize: X=4954, Y=5116

Button A: X+13, Y+91
Button B: X+89, Y+24
Prize: X=5768, Y=8030

Button A: X+20, Y+91
Button B: X+46, Y+37
Prize: X=4492, Y=6310

Button A: X+41, Y+26
Button B: X+12, Y+34
Prize: X=14523, Y=18506

Button A: X+61, Y+17
Button B: X+22, Y+31
Prize: X=2813, Y=2102

Button A: X+25, Y+66
Button B: X+65, Y+29
Prize: X=5490, Y=13313

Button A: X+72, Y+15
Button B: X+62, Y+79
Prize: X=7812, Y=5196

Button A: X+50, Y+24
Button B: X+21, Y+50
Prize: X=11770, Y=10756

Button A: X+21, Y+61
Button B: X+59, Y+25
Prize: X=14485, Y=1757

Button A: X+57, Y+27
Button B: X+17, Y+98
Prize: X=4721, Y=6104

Button A: X+52, Y+28
Button B: X+34, Y+59
Prize: X=18112, Y=2000

Button A: X+26, Y+15
Button B: X+21, Y+43
Prize: X=11182, Y=15659

Button A: X+64, Y+19
Button B: X+15, Y+66
Prize: X=4656, Y=11592

Button A: X+11, Y+28
Button B: X+68, Y+20
Prize: X=14591, Y=6652

Button A: X+72, Y+34
Button B: X+48, Y+80
Prize: X=5568, Y=7216

Button A: X+12, Y+65
Button B: X+51, Y+46
Prize: X=1113, Y=2575

Button A: X+89, Y+83
Button B: X+16, Y+84
Prize: X=8718, Y=9650

Button A: X+22, Y+48
Button B: X+60, Y+20
Prize: X=5972, Y=12828

Button A: X+41, Y+12
Button B: X+24, Y+80
Prize: X=3172, Y=8080

Button A: X+26, Y+47
Button B: X+38, Y+22
Prize: X=13684, Y=9513

Button A: X+85, Y+49
Button B: X+16, Y+79
Prize: X=1448, Y=4184

Button A: X+14, Y+29
Button B: X+76, Y+47
Prize: X=2894, Y=2461

Button A: X+32, Y+65
Button B: X+43, Y+17
Prize: X=957, Y=14381

Button A: X+13, Y+39
Button B: X+91, Y+27
Prize: X=9100, Y=3192

Button A: X+75, Y+44
Button B: X+13, Y+34
Prize: X=19238, Y=6708

Button A: X+42, Y+11
Button B: X+33, Y+80
Prize: X=7223, Y=4144

Button A: X+82, Y+24
Button B: X+20, Y+21
Prize: X=6458, Y=2496

Button A: X+20, Y+92
Button B: X+46, Y+14
Prize: X=5692, Y=8004

Button A: X+44, Y+21
Button B: X+37, Y+64
Prize: X=6377, Y=10820

Button A: X+13, Y+42
Button B: X+44, Y+13
Prize: X=8684, Y=9969

Button A: X+64, Y+27
Button B: X+16, Y+56
Prize: X=12592, Y=15407

Button A: X+94, Y+13
Button B: X+19, Y+62
Prize: X=2049, Y=5805

Button A: X+28, Y+51
Button B: X+59, Y+15
Prize: X=4005, Y=2949

Button A: X+44, Y+74
Button B: X+28, Y+12
Prize: X=18272, Y=6332

Button A: X+25, Y+99
Button B: X+72, Y+21
Prize: X=6586, Y=2838

Button A: X+50, Y+12
Button B: X+25, Y+57
Prize: X=650, Y=4076

Button A: X+11, Y+58
Button B: X+88, Y+23
Prize: X=1441, Y=542

Button A: X+81, Y+17
Button B: X+11, Y+74
Prize: X=5196, Y=12508

Button A: X+64, Y+26
Button B: X+13, Y+63
Prize: X=12867, Y=2773

Button A: X+40, Y+68
Button B: X+54, Y+22
Prize: X=19338, Y=12234

Button A: X+97, Y+31
Button B: X+20, Y+30
Prize: X=10064, Y=4562

Button A: X+22, Y+65
Button B: X+90, Y+76
Prize: X=5604, Y=6682

Button A: X+75, Y+80
Button B: X+92, Y+15
Prize: X=6476, Y=4580

Button A: X+45, Y+12
Button B: X+21, Y+58
Prize: X=2196, Y=4044

Button A: X+56, Y+78
Button B: X+88, Y+13
Prize: X=6000, Y=2002

Button A: X+17, Y+28
Button B: X+43, Y+19
Prize: X=9932, Y=4000

Button A: X+61, Y+17
Button B: X+69, Y+94
Prize: X=11610, Y=10264

Button A: X+12, Y+37
Button B: X+64, Y+22
Prize: X=6588, Y=2721

Button A: X+33, Y+84
Button B: X+43, Y+25
Prize: X=5751, Y=7038

Button A: X+42, Y+13
Button B: X+32, Y+53
Prize: X=402, Y=15288

Button A: X+80, Y+36
Button B: X+35, Y+64
Prize: X=4240, Y=6540

Button A: X+11, Y+75
Button B: X+88, Y+75
Prize: X=6094, Y=7950

Button A: X+24, Y+65
Button B: X+38, Y+23
Prize: X=2426, Y=4093

Button A: X+60, Y+26
Button B: X+44, Y+84
Prize: X=3528, Y=4256

Button A: X+21, Y+48
Button B: X+45, Y+24
Prize: X=18347, Y=4784

Button A: X+59, Y+61
Button B: X+11, Y+47
Prize: X=6475, Y=8939

Button A: X+82, Y+76
Button B: X+18, Y+61
Prize: X=6948, Y=9010

Button A: X+50, Y+32
Button B: X+21, Y+45
Prize: X=4510, Y=13786

Button A: X+35, Y+18
Button B: X+21, Y+61
Prize: X=1960, Y=4773

Button A: X+13, Y+67
Button B: X+75, Y+25
Prize: X=851, Y=509

Button A: X+17, Y+60
Button B: X+60, Y+16
Prize: X=11777, Y=14204

Button A: X+95, Y+41
Button B: X+29, Y+68
Prize: X=1467, Y=1077

Button A: X+37, Y+15
Button B: X+21, Y+37
Prize: X=11942, Y=17050

Button A: X+50, Y+87
Button B: X+78, Y+28
Prize: X=5870, Y=8598

Button A: X+66, Y+25
Button B: X+21, Y+52
Prize: X=15023, Y=10182

Button A: X+47, Y+16
Button B: X+25, Y+51
Prize: X=11410, Y=4058

Button A: X+13, Y+28
Button B: X+29, Y+17
Prize: X=7100, Y=9518

Button A: X+56, Y+73
Button B: X+50, Y+19
Prize: X=5558, Y=5629

Button A: X+20, Y+11
Button B: X+12, Y+49
Prize: X=6664, Y=5586

Button A: X+52, Y+22
Button B: X+35, Y+66
Prize: X=15518, Y=1178

Button A: X+64, Y+26
Button B: X+14, Y+39
Prize: X=19550, Y=12353

Button A: X+32, Y+76
Button B: X+44, Y+14
Prize: X=3472, Y=12360

Button A: X+18, Y+54
Button B: X+45, Y+21
Prize: X=6758, Y=12890

Button A: X+24, Y+65
Button B: X+70, Y+28
Prize: X=12308, Y=3512

Button A: X+12, Y+52
Button B: X+47, Y+31
Prize: X=1826, Y=2042

Button A: X+52, Y+27
Button B: X+14, Y+79
Prize: X=3276, Y=3566

Button A: X+14, Y+37
Button B: X+52, Y+28
Prize: X=984, Y=1930

Button A: X+28, Y+59
Button B: X+65, Y+28
Prize: X=5328, Y=4689

Button A: X+49, Y+76
Button B: X+28, Y+13
Prize: X=1652, Y=2471

Button A: X+60, Y+15
Button B: X+34, Y+81
Prize: X=8882, Y=13163

Button A: X+23, Y+36
Button B: X+69, Y+16
Prize: X=7590, Y=4060

Button A: X+52, Y+23
Button B: X+21, Y+56
Prize: X=13723, Y=17067

Button A: X+69, Y+49
Button B: X+12, Y+29
Prize: X=6017, Y=17808

Button A: X+70, Y+34
Button B: X+27, Y+60
Prize: X=12167, Y=1082

Button A: X+48, Y+20
Button B: X+14, Y+57
Prize: X=15880, Y=3036

Button A: X+48, Y+27
Button B: X+22, Y+48
Prize: X=17050, Y=18680

Button A: X+79, Y+26
Button B: X+14, Y+69
Prize: X=17574, Y=11732

Button A: X+69, Y+31
Button B: X+23, Y+56
Prize: X=12060, Y=5392

Button A: X+72, Y+30
Button B: X+12, Y+52
Prize: X=17708, Y=8116

Button A: X+50, Y+58
Button B: X+92, Y+12
Prize: X=6106, Y=3010

Button A: X+34, Y+12
Button B: X+30, Y+60
Prize: X=4324, Y=12032

Button A: X+42, Y+25
Button B: X+19, Y+83
Prize: X=1705, Y=5388

Button A: X+55, Y+31
Button B: X+28, Y+53
Prize: X=5059, Y=10879

Button A: X+68, Y+39
Button B: X+16, Y+32
Prize: X=9704, Y=5358

Button A: X+19, Y+56
Button B: X+58, Y+20
Prize: X=2026, Y=19704

Button A: X+11, Y+36
Button B: X+67, Y+23
Prize: X=10970, Y=4058

Button A: X+80, Y+57
Button B: X+14, Y+34
Prize: X=8256, Y=9160

Button A: X+27, Y+59
Button B: X+55, Y+17
Prize: X=17060, Y=8336

Button A: X+23, Y+55
Button B: X+63, Y+35
Prize: X=8700, Y=11980

Button A: X+39, Y+19
Button B: X+13, Y+37
Prize: X=3760, Y=9512

Button A: X+13, Y+46
Button B: X+53, Y+15
Prize: X=12692, Y=15703

Button A: X+35, Y+21
Button B: X+12, Y+23
Prize: X=17330, Y=16607

Button A: X+92, Y+31
Button B: X+35, Y+53
Prize: X=9507, Y=6706

Button A: X+63, Y+38
Button B: X+12, Y+39
Prize: X=11420, Y=6622

Button A: X+19, Y+89
Button B: X+88, Y+66
Prize: X=9975, Y=13835

Button A: X+11, Y+17
Button B: X+34, Y+15
Prize: X=12659, Y=13151

Button A: X+15, Y+81
Button B: X+35, Y+32
Prize: X=4145, Y=9509

Button A: X+75, Y+17
Button B: X+14, Y+76
Prize: X=14391, Y=8499

Button A: X+35, Y+14
Button B: X+43, Y+63
Prize: X=6138, Y=12933

Button A: X+42, Y+21
Button B: X+24, Y+58
Prize: X=6458, Y=16895

Button A: X+15, Y+73
Button B: X+73, Y+13
Prize: X=10485, Y=16589

Button A: X+74, Y+25
Button B: X+32, Y+67
Prize: X=5598, Y=7173

Button A: X+16, Y+29
Button B: X+29, Y+17
Prize: X=3733, Y=4649

Button A: X+25, Y+67
Button B: X+56, Y+24
Prize: X=10163, Y=7249

Button A: X+11, Y+22
Button B: X+41, Y+21
Prize: X=17426, Y=18474

Button A: X+16, Y+68
Button B: X+87, Y+57
Prize: X=3042, Y=3546

Button A: X+23, Y+56
Button B: X+46, Y+25
Prize: X=11163, Y=15816

Button A: X+38, Y+22
Button B: X+11, Y+39
Prize: X=1010, Y=18450

Button A: X+26, Y+79
Button B: X+59, Y+15
Prize: X=16657, Y=993

Button A: X+43, Y+26
Button B: X+15, Y+45
Prize: X=11157, Y=9529

Button A: X+25, Y+18
Button B: X+36, Y+83
Prize: X=1738, Y=3135

Button A: X+90, Y+13
Button B: X+54, Y+86
Prize: X=7182, Y=3618

Button A: X+95, Y+57
Button B: X+28, Y+83
Prize: X=10654, Y=9570

Button A: X+40, Y+25
Button B: X+21, Y+40
Prize: X=8555, Y=18625

Button A: X+32, Y+13
Button B: X+11, Y+46
Prize: X=9870, Y=2101`
