package main

const wordSearch =
`SSSXMASAMSSSSSSXMASXMASXXMXMAXSSMSSXXSMMSXMMSMMMMMAXSSMMXMAMAMSMXXSMSSMXXMAMMXSMMSXMAMAMXSMMMSMSAMXXMSMXMXSAAXMSMMXSSMSASMXMSAMXXXMMASXXMSSM
AMMAMXXAXAXAAAXAMASXMAXMSMSSSMAXMASMMMAAMMSSMAAMASXXXAXXXMAXAXAMMXMXAAMSASXMSXMAASMMSSXMASAAXMAMXSAXMASAMMMASXMAAXMXAASAMXAMMMMSSSMSAMXAAAAS
MMSSMMSSMXSMSMSMMASAMXSXAAAAAMMMMAXXAXMMSASASMMSAXMASXMMSMMSMSAMSAMMMAMSAXMASAMMMSAAAMAMASMMXMAMAXAASXSMSXXSMASMSMASMMMMMSMSAAMAAAMMAMSSMSSM
XAAAAXAMXMAMAAXAMXSXMASMMSMSMMMAMSSSSSMAMXMAMXMMMSXAMAXAAXMAXXXXSASASAAMXMMAMASMXSMMSSXMASXAMMMMMSSMMXMASAXXMXMAXAMXASMSMAXSMSSMMMMMMMAXAMXM
MMSSMMMSMMMMMMSAMXMAXAMAMXXAMASXMAMMXAMASMMSMMSAMXMASXMSASMSSSMXSAMXSMMSAXMASMMMAMXMXMAMXMMMSAAAAMAAXXMAMMMSXMSSMXXSAMASMMMSXMAXSASASMAMXAAX
XAMAMSAAMAAXXAXAXAMSMXSSMMSXSASMMXSMSAMXSXXMASMXMAMXXXAXMXMMAAAXXSXXXAAXXSXAXMAXMAMMAMXMAAAASXSMSSXMMMASMXAXAAMASAXMSMAMAMASMSMMSASASMMSMSSS
SXMAMMMSXSSSMSSSMSSXAXAAAMMMMASMSXAASAXAMXSSMMSSSMSSMASMXAMMSMMMAMSMMMMSMSMXSMSMSMSMXSAXMSMMSAMAAAAXASAMMMSSMMSAMXXAXSXSXMMMAXSAMXMXMAXMAAAX
XAMSMMXMAXAAAXMXMXAASMSSMMAAMAMAXSMMMMMSSMMAMAMXAXAAMXXAXXSAAXXMAMAMAAAAAXAAXAAMAAXXMAXMAAXXMAMMMSMSAMXSAAAXAAMASMXMASMMMSMMSMMXXSMMXMMAMMSS
MSMMAXAMSMSMMMMASAMXXXXAASMMSAMXMMMAAAAMAMSAMASXXMMMMAMSSMMMMSMSASMSSMSMSMSMMXMSMMMXMASMSMSXSXMMXMXAXMMSMMSSMMSMMXAMAXAAASAAMAXSXXAAAXXAXAAA
AAAXAMMXMAXXMAMMSAMAMASXMMAAXASMMAAXSMMSAMMAMMSAXSMSMXSAAMASMAASASXMAAMXXXMXAAMAMASASAMXMASAAXXXAXAMSXMMMMXMAMXXAXXMSSXMSSMMMSAMAMSSMMAMMMSS
SSSMXAXMMAMAXSXMSAMAMXXSASMSSSMMSXSMXMASASXSMSMMMXAAASMMSMASMMMMMMMMMMMAXASMSMSASXSASAMXMAMXMASMMSMXAASASMMMSSSMSXMAXAXSAMXAXMAMXMAAMXSMMMAM
XAMAMSSXMSSSMXAAMASXSMMMXMAAMAAXSAXAMMASAMAAXMASASMMSMXMAMMMXAMMMSMSSMMXMSMAXASXMAMXMASXMMMAXMMAXAMMSMMASAAMAAAAAMSSMAMXAXMMSSSMSSMSMXMAXXAM
MAMAMAXMXMAMAMMMMAAXAAXMAMMMMMMMSAMXXMAMAMMMMSAMXASMXAXSASASXMMAXXAAMAMXXMASMMMAAXMASMMXXXSSMMAAXXSAXXSASXMMMSMMMXAAXMAXMMSAXAAAAASAMXMXMSXM
SSMXMXSXSMMMSMSSMXXSSMMXASASXXMAMAMSSMSSSMMXAAXXMAXXMXMXASASAASMSMMMSSMSMSMMMASMMXMXMMMMMMMAASMSSXMMSAMASXXXMXXAMMSSMMMMSAXMMSMMSSMMSASAMMMS
XMASXMMMMAMAXSMXMMMMAMXSASASMMMMSAMMAAAMAAXMASMXSASMMMSMMMAMMMMMAAAAAMAXAAAXSAMXXXMASAAAAAASMMAAMMSAMAMAMMMXMAXXSAXAXAAAMXMXMXAAXAAASAMAMAAA
XXMXMAMASAMMSMSAAASXMXMAAMASMXAAMAXSMMMSSMMSXMXXMXXAAAXMAMAMASASXMMSSMSMSMMMMAMMAXSXSSSSXSXMXMMMMAMASXMAMXMAMSSMMMSMSMMSSSSSSSMMXSMMMMMSMMSS
SSSMSSSMSAXMAASXSMSAMXMMXMAMAMMSXMMMASMAMAMAASMXMMSSMMSSMSAXXMAMXXMAXAAAASXSSXMASAMXMAMAXXMSAXAAMSSXMASXSXXXXXAMXAAXAAAAAASAAMAMAXAMMSAXAAMX
AAAAAXMXMASMMMMXMAMAMAMXSMMSMMAMAMAMXMMMSXMMXMAAAAXXAMMXASXSSMSMXMMMXSMSMSXXAMXXMASAMXMSSMASMSSSSMMMXMAMSAASMSMMSSMSSSMMSMMSMSXMASAMXMASMMMM
MSMMMXXXMMMMAMAXXMMSMMSAAAAMXMAMMMMMAXXMAXXXSMXMXSSSMMMMMMMAXAASXXASMMMXMXMSAMXSXSMMXXSXAMAMMAAMMAAXAXSAMXMMAAXMAMAAXAXMXXXAXMXMASMMMMAMXMAM
XAMSASMXMXAXSMMMXSAMAAMMMMMSSMSMSAASXMMMMSMAAXMSXMMMXAXXMASMMSMSXMASAAXMMMAXAXASMMAXSSMSMMSSMMSMSSMSXSMMAAXMSMMMSSMMXMMMMMSMSAXMAMAAAXMMMMMX
SASMASAASXMSXAAAAMASMMSSSXMXAAAAMSMXAXXAAAMMMMMSAAMSSSSXSAMXMMMMMSMMXAMXMASMMMXSAMAMXAAXXAMAXMMXMAXXMXAMXXXAXAMXAAXSXXSAAAAAMAMXMXSASASASASM
MMMMMMMMMAMAXMMSXMASXAAAMMXMMMMSMAMSMMSMSAXMXMASXMMAAAMAMMMMXMAAAAAAXXMASMXAXSMMMMXXSMMMSXSMMMMMSXMMASMMSMMSMMMMSSMSAASXSSMSMAMAXXMAXASASASA
SSXMMASMMSMAMXXMMSASAMMSMMSMMXAMXMMAAMAAXAXSAMXSSMMMSMMSXSASAMSMXSMMMMSASAMAMMAXASXAMXMXXAMAXSAAXMAAMAXAXAAAAAAMAAAMMMMAXMAAMASXSAMAMMMAMAMX
MAAAMASAAAMXSMSAAMASXMAXAAMAMMXSAXMMSMMSMXMSASMXMXMXAXAXASAMMXMMXXMAAMMMMAMMSXSMASMMMAMXMMMMAAMSSXSSMMMMMMSSSMSXSMMMXAMSMMSMSMAMXAMXSASXMXMX
SSMMMAMMMMMAAAMMMSMMAMXMMMMAMMMMMSXXAAMMXXAXAXAMMASXMSSMAMAMSAMXASMSMSAXMAMASAMMXMAASASAAAAXMMXAMAMAMXAAAAXAXAMAXXMASMMAMAAXAMAMSMMMMASAMXXX
MXMXMXSASAMSXSXSAXASMMSXSMSMSAAAAMXSSSMAXMMMXMMXSASAXAXMXMAXSASMXSAMASASXXMASAXXXSXMSAMMSMSMXXMMSASAMSSSSSMMMAMAMASAMASMMSSMXSAXAAAAMAMMXXMS
SAXXAAAAMAXXAXAMXSMMMAMAXAAMXMMMAXXAXAMXSAXSXMXAMASAMXSXXMAMMMMMAMAMAMAMMSMMSMMAMSAMMAMXAMMMMMMASXSAMMAMAMAMSXMXSAMASAMXAAMMMSMSSSSSMMSXSAAX
XSMMMMXXXXMMAMAMXSMAAXMAMMMSAMXSXMXSSSMMSAMSAMXMMAMASXSMSMMSASAXMMSMMMAMAAAASAMXASAMMMMMAXSAMAMXSASMMMAMAMMMMASAMASXMXSMMXMAMXXAAAXAAXAASMMM
XMXAAMSMSMSMAMMMAXAXSXMXMMAXASAMASMMAXMXMMMXAMAXMASXMAXAAAXMASXSXMXAAXAMSSSMXAMAXSAMAMSSMMSASMSMSMSMSSSMXMSASAMASXMXAMSXMASMSSMMSMSSMMMMMAMX
XMSSMSAAAAAXMMAMXSMSXAMAMSMSMMASMMAMAMMAAAAMXMMXSASXMXMSMSMMAMMXAXSSMXXAMXXAMSMMXSAMXSAAAASAMXAASAMAMAXMAXMAMMXAMAMMSAMAMMSAAAMAMXMAAAMXSXMS
XMAMXSMSMSMMMXXXXXMAMMMMMAAAXAAMXSSMAMSSMMXSAAMAMAMAAAMAAAMMASMSSMXAXAMSSMMSMMAXASAMMMXSMMMSMSMMMAMMMAMSSSMSSSMASAMAXXMAMMMMMAMAMAXMMMXASAMS
XMASXSMMAMXAMSMSSMMMSMAMSMSMSSMAMXMSSMMXAAASMSMXMASXMMSMSMXMASAAXMMSMMXAAAAXAXAMXSXMXAXXXMAXXMXXXSMMMAMXAXAAAXSXSAMXMMMMXMAAXAMXSASMSXMMSAMX
MMXSASXMXMMXXAAAXAAXAAAMXXMAMASXAAXAMAASMMMSAMMSSXMAMAMXAXMMAMMMSSMMASMMMSMSAMSMXSMMMSSXMASXAXXMAMAMSASMSMMMSMXMMMMSMSASASMSSSSMMASAASXXSAMX
ASMMXMASXXAXSMMMSMMSSSMSAMMMMMAXSSMMSMMSXMAMAMAAMMSAMASMMSSMXSAAAAXSAMAMAMMAMXAASXMAAXMMMMMSMMXMASAMSASAAXAAAAXMAAAAASASAXAAMMAMMMMMMMMAXAMS
SXAXASAMAMMMAXAXAMMAMAXMAMAAMXMMXMAXXXAXAMXSXMMMSXSASASAAAXAAXMMSSMMASAMSMSASXMSMASMSSXXAAXAXSAMMSAMMMMMMSMMSSMSSSXMMMXMAMSMMMSMSXMASXMMMXMX
XAXSAMXSXMXMMSMSMMMASMMMMSAMSAXSAXSMSMSSMMAXMXMAXMSXMXSXMMSMMSAXAMASAMMXMAMASAMAMXMAXXASMXSASXAXAMAMXSAAMAXMAAAAAAXSMMSMSMXAXAMMSAXMSMAAAXMX
MAXSAMMAMXXXXXMAMASMMMAXMAMMSMSSMMXAAMXAAMMSAMMMSAXMXXXXXAXXXAMMSSMMMMSAMAMAMASAXMMMMMMAAMXMAMMMSMSMMAAXXAXMSMMMSMMSAAXAMASAMMSASAMSMMSMSAMX
AMMMMMMAMMSASXMASMSAMSXMAMXXXSAMXAMXMXSXMMXXASAMXMXAMMMXMASXMASAAMAMXAXMXAMAMAMXMSAMASXMMMMXSAMAMXXASMSMMXMXXMXMAMASMMSSMAMAMXMXMXMAAAMMXMMX
SXAAAAXMMMSAMAMASXSAMAXMSXMASMMSMXSMSASAXMASAMXSAMMSAAAMMMMASAMXSXMMMXSXSMSAMSSMAMASMMAXXXAAXMMSSXSAMXAAXMASASASMSMXAMAAMSSMMXSXMASMSMXSSMSS
ASMSSSSSXAMAMAMASMSAMXXXMASMXAXMAXXXMASAMSAMXSASAXAXSSXSASMMMASAMXSAMMMXAMMAMAAMASXMMXSMSMMMMXXAXXMAMSXSMMMXASASXAXSXMMXMAAAMAXAMXAMXMASXAAX
MAMMAMAMMASXSASXSASXMSXSXXMXMAMXMXSSMAMXMMXXAMXSXMSMXMASMMAMXSAAMASASASAXMMAMSSMAMAAAAMSSMAMSMSMMSSSMMMMASXSMMMMXSMSXSSXMSSMMMSAMXMAAMXSMMMS
SASMSMAMAXXAXASAMAMMMAAAASMMMAXMAMXAMXSXMMMMAMAMMMXXAXMMXSXMXAMAMASAMASMMSXSAXAMMMSAMXSAXAASAAASAAMMAMXMAMAXMASXXMAXAAXMAXXAAAMAMXXSXSAMAMAM
SASAMXAMSMMSMXMASMSAXMXMMMAASMMMSXMXMSXMAAXSSMAXAMMMMMAAXSMMXMSAMXMXMXMXAXAXSAMXAAMMXMMMMSMSMSMMXSSSSMSMSMMMSASAAMAMXMSAMXSMMMSAMMXMXMASAMAS
MMMAMSSSXSAAAXSMMMAMAMXXXMMMMXAXAAASMMMSXSXMASMSXMAASXMMXMASXXXMMAXMXMSXSMMMAMXMMSSMAXAXXXAXMAXXSXXAAXMMMAMXMXSMXMXMAXAXMAMMAMSMMMAMASAMXSMS
SAXMMMXAAMSSSXAAASMMMMSXXSAXXSXMMMMMAAAMXMASAMXAAMSMMAASASXMXSSSSMMSAAMAXMASMSXSSMAMASAMXAMMSAMMAMMSMMSASAMSMAXXAMASXSMAMMSMMMXAASASASMMMSXS
SMXSAXXMMMMMMXSMMSXSAAMMXXASXMAXXXAMSMMSASXMAMMMSAXASXMMAMAMXMAXXAAMSSMAMSASXMMMASMMMAASMSAAMXSAMAXAMXMMXMMAMAXMAMAXMSXMMMAAMSMSMSAMASXSAMMX
MAASXSMSAXAAMXMAAXASMMMAMSMMASMMSSMXXSAMXSMSSMMAMMSAMMASXSAMMMSMMMMMAXXXMMASMMAMAMXAMMMMAAMSMSAMXXAASAXXSSMMMASXSMXSAMXMAXMSMAAAAMXMSMMMASXM
MMMSMXAAAMMXSAMMAMXMASXSMAXSAMXAAAXSAMMSASXAXAMAMXMAMSMMMSASXAAXXXXMMMMXXMXMMXAMSXSMSAMMAMSAXAMXSXMMSXMAMXAXMXSAMAAMAMMSSSMAXMSMSMMXXAXSAMAA
MAMXMMSMSMAMMAXAMXAMAMXASAXMMSMMMSMMAMXMAMMMMXMSSSSXMASAXSAMMSMMSMMXSAAMSMMMAMMXXASASMSSSXSMAMAAXXXMMMSXMAMMSMMMMMMSSMXMXAMXSXAAXXMAXMMMXSMS
SASAMMMAAMASMSAASMSSSMSAMXSXAAAAAXAXAMAMAXXSASXMAMXMSXMMXMSMMASXMASAMAXXXAAMASMAMAMAMXMXMAMXMMSSMMMMAAXMASXMXAAXMAMAXMAXSMMSXXMSMXMASMSMMSXX
SXSASASXMXXMAAMMSAMXAAAMSMMMSSSMSSSMASXMXXAAASXMAMMMMSSSMAXMSASXSXMMSSSMSSMSAXMAMSMXMXAMMASMXAAAAXAMMMMSAMXSXSMMSAMXSXSXSXMSMSMMXSXMAXAAAXMM
MASAMXSAMXSMMMXSMAMMMMMMAASAMAAAXAMSMMMMSSMMXMASASXSAAMAMXSAMASMSAMXSAAAAAAXXXXAXAXAXXAXMAMXMMMMMSMSMSXMASASMMAASXMASXMXMAAXAMSAASMXMSMMMSMA
MMMMMMSAMASAXAMXMAXSAMXMXSMXSXMMMXMAXAMXAAMXMXAMXXAMMMSXMAMAMAMASMMXMXMMMXXXMASXSMXMMSAMXSSMMSSMAAAAAAXSXMASAMMMSAMXSMSASAMMSMMMSSMAXAXAAAAS
XMAMSAMMMMXXMMMASAMXASAMSMMMSXXXXXXMSSSMSMMMAMSSSMAMXASAMXSAMXMXMASAMASXMMSSSMAAAAASAMAAAMAAAAAMSMSMSMMAXMMSAMXASAMXSASXAASAXAXXMXXXSMSMXSXM
ASAMMAXXASMMMMSAXMASAMASAAAAMMMMMSMMAXAMXSAMXMAAMXMXMAMAMXMASXMMMXMAMAMAXAAAAXMSMSMMAXAMXMXMMMMMMXAAMAMMXMXXXMMMSXMMMMMMMAMAXAMXSAMXMXAMXMAM
XSASXMMSMSASAAMMSSXAXSXMMSMMSAXAAAXMSSMMAAMXAMMSMMMMSMSMMAXMXXAAXSSSMMSXMMSSMMXXXAMSSMASXSSMMSAAAXMMSAMSAMASAMAMXMMAXMXAXSMSMSMAMAMMXAXMASAM
XSMMAMXAASXMXSSXAMMMMSAMXMMMSMSMSSSXAAAMMMMSSXXAMAAAAMAMSSSXAXMSMAAAAAAXXMAMMXMMSXMAMSMMMAAAAMXMMXMXSXMSASAMXSMSMAMSXAXSAMAXMASXSAMXSMSXXSMS
AMXSXXSMXMAMXMAMXXAXMXAMASAAMXMAAAAMSSMMSAASMXSASMMMXXAXAAMMMMXXXMAMSMSSSMAMXMAXAXMAXASAMSMMMSSMMXMAMSASXMASASXAXSMXMMMXAMAMMAMXSMSMSAAXXMMS
XMASAXAMXXAMXSXAMSSSSSXMAMMMMAMMMMMAXAMASMSMAASAMMSSMSSSMMXAAASMMXSAMXAAXMASAMMMMMMXXAXXMAAXXAAAAAXAMMXMASXSASXSMMSXXMASXMMSMMSXSXAAMMMMXAAX
SMASMMXMMSMSMMASAAAAAAXMXXSSSXSAMSSSSMMMMMXMMMMAAXAAXAMXMXSSSMSASAMXMMMSMXMAMSSMSASMMSMMMSMMMMSXMASMMSAMXMASAMXMASAMXSXMAAAAAMSASXMSMAXASMMS
MMAMMSXMAASAASAMSMMMXMXXMMAAAXSMMAAXXXAASMSMSMXSMMMSMXSAMXXAAASAMXMSAMXAMXSAMXAASXSAAAAAAAXMSAMMSXXMASMSAMMMAMAMXMMSAMXMSMMMXMXAMXXAMXXXMSAX
XXMAAMAMXMMSMMAXXXMXXMMAAXMMMMMSMMSMMMXMSMAAAXMXAXXMAAAAASMSMMMMMXSAMXSASXMXXMMMMASMMSSMSSSXMAXAXSAMXMASXSXSXMASMXXMASMXXAAXAMMXMXMSMMMSAMXS
XMSMMXSMMSMMMSSMMSMXSASMSMXXXAAAMXMASMSMMMMSMMASAMMMMMSSMMAXMAAAMMSXSXMXMAMXMMSSMMMAMXMMXXMAMAMMMSMMMMXMASAMXMMSAMSSMMXAXSAMXMASMMMMAAASAMSX
MAMAXAAXAAAXAAAAXAAXSASAAASMSMSMSAMMMAAAXAXXXXXAMXXAAXAXMMMMMSSXSAMMMMMASAMMXAAAAMSMMAXSXMXAMXSMASMAXSXSXMMMXMAMXMMASAMXMMXMSMXMXAASMMXXAXXS
ASXAMSSMSSXMMSSMSMSMMAMAMSAAMAXAMMSSMMMMXSSMAMSAMMSXSMASXSMMAXXAMXXAAAXAMASAMSSSMMXAMASXMASXMAXMXMMXSMAMXXMSMMAMMXSAMXASAMXSAMMSMMMSASAXMMAX
SXMAXXAXXAASAAMMAMAAXMAXMMMSMAMMSMAAASAXSAMMAMMASAMSXMAMAXAMSSMXMAMXSSSXMMMXAMXXXSMSMMMMAMXMSMSMMXMXAMXMAMMMAMSSMMMAMMSAAXSMASAAMSMSAMAASMSX
XMAXMSMMMSMMASXMAMMMMMMSAXAAMSMAAMSMMAAXMASMAXSAMAMMXMAXMMSMXAMAMMSAAAAXSMMMSMMMXMAMXXXMXMXMAMAAMAMSXMXMASAAMSAAAXSAMMAMXMAAAMMSMAAMAMSMXAMA
XSSSMAXXAAASAMXSSSSXSAAMAMSMSMMMMMXMXMSXAAAXMXMMMSMSMMSAMXXXSAMXAAXMMMMMAAXAXAMSAMSMMMAMASXSMSMSMXXAASMMASASMMASMMMXSAMXSSXXMXXXMMSMMXXXMXMA
MAAMXMMMXSXMASXMAAXAMMXSMMMMSAMXSMSAXXMXMMSSSMXSAMMAAAMMSMSMMMMSMMMXAAAXSMMXSASMMXAAASXXXXAMXMAAXMMSSMAMMSAMMXMXMXMAMAMAMAAXMXXXAXMAMAMAMMMS
SMMMAMASMXAMXXXMMMMSMMAMAXXAXAXXAAXMASXASAMAAAAMASXMMMXAAMAAAAAXSXAMMSAMXASAMXMAXSSSMAAXSMSMAMSMXAXMXMXMMMXAXMMMSAMAMAMASMSXMASXMMMAMASAMAXM
MAMSXSAMAMMMMSMMAXXAAMSSSMMSSSMSMMMXMAMAMAMSMMXSAMAXMAMSMSSSSMSSSMSXAXMAXXMASASXMAMXXMXMAAAMXMAASMXMAMSSXMXAMAAASXSSSMSASXXAMAXAAAMASXXMMSSM
XAMMAMMMSAAAAAMSMMMMMMXAXXXMAXXXXXMASXMASXMXSMAMAMXMMAXAAXMAAAXAXAXMMMSSMXSAMAXSMSSSXXXSMSMSMSMMMAASAMAMASXAAXMMXXXXAXSAMXSMMSMSMMXXSMSMAMAX
SMSSMSMAMSSMXSMSASXMXSMSMSMMAMXMXSMMMXSASXMXMMMSSMSMSMSASMMSMMMMMMMMXAAAAMMSMXMAXAXMASXAAMXMXXXAMMMAAMXSAXSAMXSAMXSSMMXAXAMSXMAXXSXMXAAMXSAM
ASAAAAMXXAMMMXAMXMAAASAXAAAMASXMASAAMXMASXSAXAAAAAAAAASAMXXMAXXMAXASMMSXXXAAMMMSMMSXAASMMMASXSMSSXXMXMAMXMMMMAMASAAAXMSXMMASMMSMAXAMMXMSMMMM
MMSSSMSXMAXSAMAMSSMMMMAMSMXMAMAMASMMAAMAMASXXMSXMXMSMMMAMXMSMMMSMSASAMXMSMSMXAAAXXMMMXAXAXSMAAXASAMXMSSXMXAAMXSAMMSMMMMSAMAMAAMXXSAMAASAMAXM
XXAXXMXMAMXMSXXXXXXXAMXMMMXMSMMMASASXMSAMXMSXXMASAMAXMMMMMXAAAMAAMASAMAASAMXSMMMXAXAXAMXSSXMSMMASMXMMAAASXSMSAMASXMXXAASXMMSMMSAXSXMSMSASMSM
MMMMAXAMXXXAXMMMSMMSXSAXMASMMASMMSMMAXXAMXAAMXAXMASASMAMAAMSSMSMXMAMXSSMMAMAAAXSSMSMSASAMMMAASMMMXAXMMSMMAAXMAMMMMMASMMMAMXMASMXXMXXAXSAMAAX
AAAASXSAMXMAMAAAAMAAASMXSMMAXAMAAXAMXMMSMSSMSMSMXMMXMMAXXXAAMMXMXMASXMAMMSMMMSMAAAAAXAMMSMMSAMXAAMMSSMXAMSMMSSMSAMXAXMASAMXSMMSXXMAXMXMAMSMS
SMSMMAXASXMAXSMSSSSSMMXXMASAMXSMMSMMSSMAAXAXSAMAASMMMSMSAASAMMAMXXMAMSAMXMAXAAMMMMMSMXMASXAMASMMXSAMAASMMAAAAXXAMXMSSSMSASXXMAXXMASMMAMAMMXX
MMMAMSMMXMMXXXAXXAAMAMXMXAMAMAAMAMAAAASMSMMMMAMSMXAAXAXMXMMAMMASAMXSMSMSASMMSXSASAAXMAXSMMXSAMMMMMAXMMMMSSMMSSSMMMAXAMXSAMAXMMSMSAXXMXSXXMSX
MAXMXXXMASMMAMMMMMMMMXASMAMAMSSMSSMMSXMXAAASMSMMMSSMMMSMMXXXMXXXAMAMXMASAXAAAAAMSMSASMMAXSMMMSSSSSSMSAXXAAAXAAXMAMXMAMXMAMXMASAAAAMSSXAXSSMM
MMSXAXAMMSXMAXAASAMXXSASXSMXMMAMXAXAXXMSSSMSAAXXMAXAAXXAASXSSMSSSMXMAMMMSMMMMMMXMAXMASXMXSAMXXAAAAAAXSSMSSMMMSMSXSMSMMSSSMMAMMMSMSAASMAMXAAX
AXAMSXSAMXXSXSSSSMMSAMXMAMMAXMASMMMAMXAAAAAMMMMSMSXSMSMMMSAAAAMAXAMSMSMAAXAAXASAMMMSAMXXAMAMSMMMMMMSMXMAMAAXMAMXMAASMMMAMMXSXAMXAMMMMMMSXSMM
SAMXMXAAXXMAMXXMXXXSXMSMMMXMXXXMAXMAAMMMSMMMAXXAAAAMSMASMMMMMMMAMXMAAAMSSSSSSMSAXSAMAXMMMSSMXAXXXXAXAAMAMXSMSASASMXMAAMAMAAMMSMMXMXXXSAMAMAM
MMXSMMSSMMMSAXMASMXMMXXAAASXMSASXMSMSASAXASMXSSMSMSMAMSASAMXMXMMMXMMSMXMXMXXMASAMMMMXXAAMAMAMSXMMMSSSMSASXMASASAMMMSSMSSSMXSAMMSAMXMXMASAMAM
XAAMAAXAMAAAMAASASAASXSSMMSAASAMXXXAMMXAMASXAXXAAXAMAMXAMXSMMMSAMXMXXMASASMMXXMXSMAAMXXMMAXMMMXXAAXAAAXASAMMMMMXMAAAAXAAAAAMASAXASAMXMMMMSMM
MMSSMSSSSSSSMXMMAMMSMAAMXXMMMMAMMMMMMMMMMAMMSSMSMSMSSMXXSAMAAXMASAXXAMAXASAMMXMASAMMSAMSSSXSAXMMXSMMMMMMMXMXAMAXSMMSMMMMMMMMMMXSAXMXSXMAXAMX
AAXAAMMAAAAAAASMAMXXMMMMSSMSSSXMXSAAAAAASAXAMAAAMXXAMXSXMASMSSSMMMMSAMSSMSAMASMASXMAMXSAAXAXASAMXMMXAMMAAXSSMXAMXMAMAXXSSSMSXAXMMMSXMASMMXSX
SXSAMXMMMMMMSMXMASXAXAAXMAAAXMAMXMASXXSAXMMSSMSMASMSSMMASAMAMAAAAAMMAMMAASASMXMXSXMXSSMMSMSMASMSAAMSMSMMASMXMASXMSASXMMXAAASMXMAMAXAXMMSAMAX
XMXMMMSMSAMXAXXXAXXXXSASMMMMMSAMAXAMMXMMXSAAAAXAAXMAAASXMXSMMXMSMSSSMMSMMMAXXXSAMXMAMAAAMAXMMMAXMMMXXAAMMMMASMXAAMMXMASMMMMMAMSSMMSSMMAXXMAM
SAMXXMAMSAMAAMXMSSMSMMASASXMXSMSSMASXXXMAMASMMMMXXMSSMMASAXXASXMAXAAXXAMXMMMASXMMAMASMMMSMSASMSMSASXSSSMMMSASXMMXMMXMASMMSSMXMAMAMXMAMMSAMXS
XMAMXMMXMAMMXMMMXAAXAXASAMXSAMXAAMSMXSAMXMMXXAAMXAMAXAXAMXMASAMMSMSMMSXSMXMAXMAMXASAMAAMXMSAMAAASASXMXXAAMMMSXMMSMMSMMSAXAAAXMASXMSSSMAMASAX
MSSXMAXSSMMAAMAMMMMMXMMMMMAMMSMSMMASASMMASMMSSSSSMMASMMSSSMXXMAXAAAMMSXAXAXXAXAMXXMMSSMMAMMSMMMMMMMAMMXMMMAMXXSAXAAAAAXMMSSMXSMMMXMAMMSSMMXS
AAAMXMXMAMMXMSASAXMMMXMASMASXSAXXSXMMSAMAAAAMXMAMXMMMAAMAAMAMXSSMSMSAMSAMMSMAMSSSMAXAXASMMSXSXSXMASAMAMSMSMMSXMASMMSXMSXAXAMAMAAXMMAMMXAAASX
SMMAMMMSAMMSMSASMSAXSMSAXSASMSAMAMXXASXMXSMMMAMAMMSAMXSMAMMXSAMXAMXSAASXAAXMAMMAMSSMSSXMAXXAAMMMSASASXSAAAXMSMMAMAMAMXAMXSAMXSSMSSSMSXSMMMSS
XXXMASXMASXAAMAMAMAMAAMMMMMXAMMMXMMMASXMMXXASMMMSXSAMAMAXMXXAASMXMAMMAMXSAAMXSMAMAXAMXAMAMMAMMAXMXMXSXMMSMSXMASMSXASXSXSASAMMAAXAXAMXAXXAXMA
SSXSSXXSAMMMMMAMXMAAMXMMSSSMSMSXSASMSXXSAMSMAAAASAMAMSSMMAXSSMMMXMXSSSXXMASAAMMXSAMXMSSMXXAAXMMSMMMMMMXAMAMXSXMMXMMMAXAMMXAMMSXMSMXMMAMSXMAM
XAXXMAMMXSXXXSASASMMSAAAAAAAXAMXSAXXXAMMSAAMSAMXSASXMMAMXMMMASAXMAXMAMXXMAMXXMAXMXMXXAMXMSSXXMMAXAAAAXMMSSSMMMASMXXMMMMSSSSMMAXAMXSSMXMXAXXX
XMMAMSMMASXMXXASASAASMXMMMMMMAMAMAMMMXMAMSXXXXXAMAMMXSMMMSMSAMXSAMXMAMMXMASMSMXXMAMXMSMMAAAMSSSSSMSSXSAAAXAAMSAMXSMSXMAAXAXAXMMMSAMXMAMSMMXX
XXAXMMAMAXMSSMXMMMMMMSXASAMXXXMMSSMSAMMAXMMSMMMSSMMSXMAAXAAMXSXXMAXMAMSAMMAAAMAMSASXMMAMXMSMAAAXAAMAMSMMMSSMMMXSMMAAAMMSSMXSMXAAMMSASMMAMMXX
SSSSMSSMSSMAXSAXMASAAXMAMXSMMSAAAAMXMXMSXSAAMXAAMAAXASMMSMSMXMAXASXSMMMAMMMMMMAMXASXASAMAXMMMMMMMMMAXMMSAMXMSMMXASASXMAMAMAAMMMMMXSASMSAAMMS
AAASXAAMMAMXSMXXSAMMXSXSMASAASMMSMMAMSMXAMSMSMSSSMXSMMAAAMAAMMSMAXMAXASMMSSXMSSMMAMMXMASMSXMSAMXMSSMMSAMMMMMAAXMXMAXAMXXAMSSMAMASXMXMASMMAAS
MMMMMMXMMMMMMXMXMAXSMSAASAMMMMAAAXMMSAAMMMXMAXXAMXXMASMSMSSSMAAMXMXMSMSXAAAXMAAXSXMXMMXAXAAAMSAMXAAAAMAXMAAMSSMXMMXMMSSMAXXAMXSASAMASXXXSMMS
SASAXSXSASXAAAMAMXMSAMMMMXMXMXMSSSMMSMSMAMMMMXMAMMSMAMXAAAAAMSXSAMXXAXSMMMMMMASAMASMSXAMXSAMSXSXMMSMMXSMSSSXXXAASMSASAMXXMXAMXMMSAMXSAXXMSXM
SASAXSASAMSMSSSSSMMMAMXSASXXXAXMAMAMMXXXASMSAAMXMXAAASXMMMSMMMMXAAMSSXSAMXXSXAMASXMAMXMSAXMXAAMXSAXAAXXAMAMXMMSMSAMXMASAXSMXMXAMSAMXMAMSASXM
SXMMMMAMXMXMXAXXAAXSXXXMAXAAMSSMXSSMMSMSMSAXSSMXXSXMMXMXSMXMASMSMMMXMASXMSMMMSSMMMMAMMMMMSMSMMMAMSSMMMMAMAXMAAAXMAMASAMXSXAXSSMMXSAXMSMMAMSA
SASXSMSMSMSAMMMSSMMMXMMMSMMSMMAXAAAMXMAAAMMMMMAMMMXMAXSASMXSAMXAAASAMXMAXXXMAAAAAXMMMAXAAMMXAAMXMASXAAXXMXSAMSASMMMASAMXSMMMAAASXMXMMMMMAMXS
SAMAXAAXAAAMAXAAMAAMAMAAXMAMASAMMSMMAMSMMMAAAAMAAAASXMMASAMXASXSSMSXSXSMMSMMMSMMMSAXXXSMSXMSSMSMMMXMMMSASXXXXXXAAAMXSAMAMAAXMAMXAMAXAAAXXMAM
MAMSMMMMMSMASMMMXMAMMSMSSMAMXMXXXMXSXXAMXSSSXSXSMSMSAMMMMAMSXMAXXXXAMXMXMAAMAMXSAXMSSMAAAXMAMXAAAMSMSAAAMSAMXXSSSMMASMMASMMXSAMSAMMSSSSSSMAS
SAMXAMXXAAAMMASXMMSAMXXMAXSSMMXSAXMAMXAMXMAXAMMXAMASMMAASAMXSMXMMMMAMAMAXSSMSSXMASXAAMMMMXMASXSSMXAMXAMMMXSXXXAXMAMAMXMAMMXASAXSXSMAXMAAASAM
SXSSSMSMMSSSMASMXAAMMSMSXMXAAAAMSMXMASXMAMAMSMSMAMAMASXXSAMAMXAAAASAMMSMXMXAMSAMASMSMMXMAMSAMXAAXMSMSAMXXAAMSMMXSAMXSSMXSAMASMMSASMXSMMMMMXS
SXMXMXMAMAAAMSMAMSMMAAXAXMSSMMXXAXAMXAASXMMSMASMAMXSXMMXXAMASMMMMMXAMAAMXSMSMSMMASAXAXAXXAMMSMXMXMAASAMXMXSAAAXXMAMSAMMXAMMASAAXMXMAMXXMSAMX
MASASMMSMMMMMAAAAAXMSSSMSMAXMASXSMMXMSXMXXMAMAMSXMMMMSXSSXSASAASMSMMMSMSAMXXAMXMMSASXMAXAAMXXMAMAMMMXMAXAMXXSXMXMAMMAMSSSXMAMMSSMAMMSAMXMXMS
SAMXSAAXAAXSXMSMSXSMAAAXMAXSMAMAMAMSXMAMASMMMAMSMXMAAMMMAAXAMMMAASMMXAAMASAMAMXXXMXMASASXMMXAMAMXSXMAXAMSMSAMASXMAMSAMAAMXMMXAXAMXXMXMXAXAXM
MASXSMMSMMXXAMXMXMAMMSMSMSMSMSXAMAMSASXMASAASMSMAASMMSMMMMMSXSMMSMAXMMMMMMMSAMMSMSAXMMMMAAXMMMXMAXMASMMXMAMAMAMASAMAAMMSMXMXMASMSSMMASMMSSMM
MASAMAAXMSMSMMAXAAAMXMXAAXAXAMSSXXMSAMXMASXMMXAMSMSAAMXSAXXMAXXXXXMMSXASXSXSAMXMASAXMAMMSMMSAASMXSAAXMSAMXMSMXSAMXSMXMXMMSMXSAXAAAASAMAMAMAA
MAMAXMMMMAMAASXSXSXSAXSMSMMMXMAMMMMSAMXSMSMMXSXMMMSMMSAMSSSMSMMMMAXAMSMXAMASXMAMMMSASXSAAAAXMASAAXMXMAMXSAAXAAMMSMXAAXXXAAXMXSMXMSMMASXMASMM
MSSSMSXXSASMMMMMMAXMAMXAXXMAXMAMXAAXXMMMAMMMAMMMAMXSMMXSMMXSAAASMMMXXASMXMAMMXAXAMXXXAMXXMMMSMMMMXMMMAMASASAMMSAMAMSMMSMMMXMAMMSXXAXXXASXSXS
MXAAAXAAMXXXXAAXMASXSAMXMXXAXSMMSMSSMXMMAMAMMXMMASAMXAMMAMXXMXMMAXAXSAMXSMSSSSMSSSMXMXMASXSAAAXSSMAASAMASAMXMAXMMSMAASMSMXAMXMASMSMMMSXMSMMM
XMAMMMMMXSSMSSMXMASXMAASMSMSMMAAXXMAXXASASXXSAXSAMMMMSMSAMSMSSMSMMSXMXMASXAAXXAAAAASMXMMAAMSSMMAAMSMSASMSMXXMMXMAMMSMSAAXMXSAMXMMAAMXSAMXMAX
MXXXAXMMAAAXMAMXMAXXMAXXAAXXASMMSXSAMMMSAMAAMAMMMMSAAAMXMMMAAAMSXMAAMXMASMMSMMMMSMMMAASXMXMXMASMMMAXSXMXMASMSXMSASAXAMMMSSMMXMSASXSSMSAMAXSS
MSMMMSAMMSSMXAXMSSMMSXXMSMXSMMAAAMMMMAXMXSMMMMSAMASMSMMXSAMMMSMMASXSMAMSXMMMXMAXMAAXXMXAXASXSAMXXSXXMAMXMMMAAAAMAMXMAMXXAAAMAXXXMAXAXSMMAMXA
MAXAAXXXAMAMMSMXAMAXSAXMASXMXSMMMSAAMSXAAXAXAAAAMAXXXXAAXXMXMAXSMMAXXXMAMXAXXMXSSSMSMSXSMMXAMXMMMXMSMMMAMAMMMSMMAMASXMMMSSMMMSSSSSSMMXAMMXSA
SMSXSSMMSSSMAMXMAMSMMXMMASXSAXMMXSMSAMMMASMMMSXXMASMSMMSSSXMSMXXMMMMMMSSXXMSSXAMAMASAAAXAXMMMSMAXMAAAASMSAXAMAMSMSASAMXMAMXAXAAXAXAXXXXMSASX
MXMXMAMMAAMMXMAMAMXASMSMAMAMMSAXAMXMAXXXMAMMAXMMMXMAAAXMAAXXAMSMSAAAXMAMXASAMMSMAMMMXMMXMMAAAASXMSSMSMSASASMMSMAAMMSAMXMASXSMMSMMSSMMSMAMASM
AAMAXAMMMMMAMXMMAXSAMAAMAMXMXMXMASAMAMXMMAAMASXSAMMSMMMMMXMMAXSASXSXSMASXMMASAXMXMAMMSSSXMXMMMSXAAMAAAMAMAAAAASXMMMSMMSMASAMAMXAXAMSAAMAMMMM
SMSMSSSSSSMMXASXMMSAMXMSMSMMSMSMMMXXAXASXSSMAMMMSMAAAXXSXMMSSMMXMAXAAMAXXXSAMXSMMSMXAAAMMSASXAMMMASMMMMAMXMMSMMMASXMMAAMXMXMMMSSMMMMMXSAXSAS
XXAXMAAAMMASMMSAAASAMXXAXAAMAMAAXASXMSXSAAXMXSMAAMMSMMMMMSAAMMMSMAMSMMMSSXMAXXXMAASMMMSMASASAMXMMMAMAXSXSAAXMASXMMAXMXSSXMMSAAXXAMASAMMMSMAS
MSMSMMMMMSAMAAMMMMSAMXMAXXMMASMSMMMAAAMMMMMSAAMSSMXAXAXAAMMMSAAMMXMAXMSAMXMMXSSXMASMAMAXXMAMMXAXXMASXMXASMSMSAMXASMMSAMXAXASMSMSMMAMASMMMMMM
XAXAMXSXMMXSMMXSXMSAMXSMMSSSMSAAMXMSMMMSAMXMASXXAXSAXSSMXXAMSMSSSMSMSXMASXMAAMAAXMMMAXMSSMAMASMSXSASAMMMMMXMMASMMMSAMXMSAMXSAXXXAMXMAMAMAAAS
SSSMSAMXMXAXAXMMAAMAMMAMXAAAMMAMMSAAAXAXAXASXMMSMXMSMMAXAMMXSXMAAASXSXMASAMMXSASMSXSSXSAMMAXMMMXMMASXSAMAMXAMSMXAAMSSXMMMMMMAMMMMMSMMSSSXSSS
XXAMMMXSMAXXAMXMAMXSMMMMMMSMMSXSAXSMMMMSSMMSASXAXXAAXXSMMXSAMXMMMMMASAMASAXMXMAXAXAAAXMASXMSXASAMMMMAXXSAMSSMMASMMMAMMAXAAAMMMMMAAMAXXXMAXAM
MSMMXSASXASAMXSAAXAMAXXSAAAXXAAMMXMSAMXAMAASAMMMMMSMSAAAAAMMMASAMXMXMAMAXXMSAMSMMMMMMMMXMASAXSSXXAAMAMAMMMMAAMAMXAMAMXMSXSMSAAASMMSSMMMMXMAS
XAASAMASMASMSAASMMSSXMASMSSXMMSMSSMASAMXSSMMAMMXSXAAMMXMMXMXSXSASASMSAMSMMMMASAAMMXAMXMXSSMSSXMASXSSMAXXMASMMMMMSMSXSAXMXXASMSMSAAAAMAASXMAM
SSSMSMAMXAMXXXMASAMMMMMXAXAMSAMXAAAAMSAMXAASXSAAAXMXMASXXAXAMMSXMXMAMXSXASAMSMMAMXSSMAMAXMAXMASASMAAMXMMSMSXMAAMXXSAMXXSAMXMXAMSMMMSMSXXAXAS
XXMAMMSSMMXMMSXMMMMAAAMMXMMMMASMMSMSSXMASMMMAXMXMMSXSXAMMSMASAMXSAMXMXMXXMASXAXMMMAASAMASMSMSXMASMMXSASAAXMASMSXSAMXSAMMMMASXMXXXSAXXXMSSMAA
SAMXMMAMASXSAAMXAAXMMMSSXXMXSSMAXXAXXXAXMAAMMMMAXAMAMXMAMAAAMASXSASXMMMSMSMMMSMMAAMASXSAXAMMXXMXMXAAMXSSSMSAMXAASASAMASXXMAMXSXMMXAXXXXAMMSM
MAMASMMSAMMMSSSSSSSSSMAMMMMAXASXMMMMMSMSMSSSMAMMMSMAMAMMSMSXSASMSMMAAAAAAXASAMASMSXMXMAMMSMMMMSXSMMXSAMAMXMASMMMMAMASAMASAMXAMAAXMXMSSMMSAMX
SAMASAXMXSMAAMMXAAAASMAMAAMXSAMXAAMAAAXMAAAAMXSAAXSASMSAAMAAMXMASASMXMASXSMMASMMXMAMSAMXXMAMSASAMXAAMMMSMMSXMASAMXMXMASAAAAMASXSXAAAAAAAMASA
SXSXSXMASMMMXSAMMMMMMSASXSXMAXMXSXSMSSSMMMSMMMSMXMAMXXMMSSMXMXMSMXAAXMAMAXXSMMMMXSSMMAXXSSXMASMXMASASXXMAXSAMXMASMXXSXMMSSMSXAMMASMSSMMMSSMA`