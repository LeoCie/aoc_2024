package main

const disk =
`7042192786322879901576993891377590511736741767315192964332868097913680327911408049473948726628836920307340669188113753803153659468398841667256318837213497449812197859317387836613881252917950945128252426535112243445296728458774405960537866519974382685282661658775832645577641732573851074191137879610308435763694951566486076573058749746613426198593281045313634719567241442126860694723599098697826148376874125702787376363265323766543824792873142662169183991703810773745107532276027108376363537164866817548351178251589694021291144651980222647255198191153253359734116931687349047711421487820703873385042893877648469412767577467451065221387511949255339645739381794648273123066595994521829159739317738882552843757379979261660634894256435106449654385945249487032715058404021837676625743957560165521201358185667671380199414747726318627904878234383462572204082274161436799493748733239446061758855547841179842406282829698616425875880813169984647833023116044137765246590552925656944913797679478342237306174223339621116758070506669643893213346395864384061555266851264881787654158179480958552443212945822487563107963539564539073559688415843931078563492709122392454457155892730717888312145574390607095641343608815476511799581908798947014334329777386922079399899519159835514937760546816625571429875276785842592635545373045993242205064305949505186668626634262878358743241422627215795462111733049347364549124276623341847875586289371761616796027248910433752942185802142607393464649486989827987322418968542885380991874111379747464349016316925736062264298801815163894187962663272602987873259603615275812914814174044525229664431316631783941335898202122921255314117827596792715934629794371416346818960821383324624882074714443299462497675753234123211733786826860247966544812117841564058708574352212965711151050593395797797719041272122147936225284703379596448507288545135292624164013546825685436844243814831555746619337118458918369827538263661173175544163727289917176308232854033322237816841667538541110495632717669777555888369671126373378753557102854216014481959837513923865583640976534277710135123338335939516204925371224185361971420321974295923292771382277259834706862272362817643687549492043482282122488898546662349232015252926795289102247278888344769366524387892328268867054835092475280518171442771955964892998724192925683991071148090197918398041607150932629862488907592892489706486904429358958375080949370992020794625629035238082414166408913811520293393556999566538336612352152733616263837812136441954859446526137829647848478862552272384789177469871283820169190286870455458408577888275231527319689985872904416975052971845536887315173644624202246827094523242491294372296929116834663672481578667502082814484488587582483205143991570772646281528838757427963929670858548722123183675474629229345829834297453942040861367816769642554117538855341846528193155959232819880901594165211517997781921648048716982126993201187854392203632349647317283275450399081162784409610174463714987764788315834482198979561275651375484952320604169277456967176262049713540913846949995437010298190715653802983886856923234805766872052952059498160119890478655948924297357883477783892638118982256637452375921115874298144502114245114754539371072276218437071212132355326186611886952533842217921528114741499597294737014656661103115176080823552516792264658595914763683908690488392665173572193142352141722771117602240251448471576536186358877406875263111136621982046833247417482857084924448994199995412691283863629206523923782121562307234991762341660584631118028937318626963481992709482154877233426822032862170726019123321394780963337561551974835384335537959897848435486923255728666456148333354105394754734484514889827675781859566842188675313114178885835162621236884928839389712422840672387814648809917449528957683455195444656803322404843227080567527552956911966585970652025273470571398484282236253741733769513714468695960145381954595838270183332465626591359348198297682334137384789219557434036326271901886849415191449233534105787649848294874764132588220873270808415259019708654909398559466509868122588785228414768336799301174844960674551705125838999821181669778539666169999726810979892841152476836553715438230944344208541704929753684258836346062373162899895853173461860909785328411161225698978927868118732649967742476611780844916733573307996607282146363814875835574803434722049869010445114241516952713573626433833449533484136279714902839378937565315904343434259473868576774988316915328737219245048871436938889933710364793409568777756499111895389898223724857197246912481545986685335777017489886721476246984862515709936996274178150466118276926768611387525437545268693108897121062264244175952385689827994644781815711949633226090777159635080193133346435327758623296826461725045415598187753615092497325525627761181894885582630732188665840847995196266212941109584732985542627962694493153437230528023165097275566237645899666471371497320205530873044566281123111811267765963507342805334994959516525372031416132836650494558216771726394381284351739187822267368948175698940582213211122522382479032974892109069684221605029299358754482207065958041377392144786555632378898217897286354226645309892648043146334499140847031183663844384463748351821834297699297164162854236178940764178187524719512596392502182903849539299246478807732109385922913152237303769124536838270878134771967742987304592494916901935711781774718661776695125151365709423602618558417275159324998771818301537732980399226795234819824508792997192482738737471766075948824794470952763331451304662131822638088919183148881715380158799738825966329223614296010115759488875595144906472148898403913643848968787982759276683331934722220946483693145885080991143228754622521282299416526851949939316484170794155829090103975233534585035107280931047682677224856213372351736214914976636237344634676854448621354299335112580246810168394363916805035961335571454659899806467821766212570214520329158288497791674378563554648814045415847417694558049644223791337528521677064535322609750735357938082144059802430738270864650197812839783542161792223612639514275243218283586382292663195418016249032886842318817543154888837226939191973295376441033548844393065741376829257888224289530107633373712725728325437193271871068712526846030483344724191462137824114929498361856942456199053223573426079612319671913545212549883104654624743319423744753542626475776327545107145484113587878118250977772173194471494621073917158645489571168248190391484402384828748651372648258332115836820639995985231312467348258598652757337684724212086283159734675898176468688429933167212742944314948606826608585338313164237362160205251775762333741137527649598262221936917537856625284548621748941788179448724282393862092466793603771591370977377114986952666132311437795107712637129696913982873154720137495915653358828608992503195174782307555319096158167577960615738636490631142713577493697975770681681698860434288866613686477139280492647728056693980377749923465956388615569536293979131953852646661314954431222749851914763656674558296475546431268637055992642841092976314387234879854473331345731133919248397642516464158536169885959377034593928834851666261499337686953217341218133165680642296263316936857155022289135267865761743649175718091709628457147762054616492852876651179953964727029635278654189303468962512743740448271239976897665716070388465624820909665753096136243313110484482464986591043967999978933794541787312392033725443135921864179202672686311275332364845123162871754456266349868122988645713111137711181559247343739134515323540199761852016701372134682952520713469991771794512615996704014917786898119463111934699158757732995735995919436654711385856242629629583265254186327686975391553381988485226752647116899178547151626434866403965921944792047151199107783312193629683385748799863683144704079418255969129625326499817811939957059891269328572823356511630777771674877914131242040395150954335223531733469717628809450483541732140457275244664805519827766489420287328171912352491158584291240506228184651957526562918685674715757267551278112174116334269306363404696311647279414689172949711703020853558336659842391826965265818972831279185201441565575259358512858714126121923363117831314907128784528276277274242387553225811377238458293727272215577344429855499367420817685217861398077493377126119147870957334715811562364128411644176917282491067449324401922895058486792442932217357312355519887125563786697858971977987777299629697751849434760225335869160685742192911306081884895591021237258687418449210472263854648522934525267243219901266424972244551277365814655165375666612108278977220793371173677155333557083452324816497283131955826884089827269345996974088656165625871387652428053332152312966514054656539689210942070969360515618714163914343419399367153465357153045129990599311605718571353969863673263428268338020594922565071797547778739733387574221286440527291702956155381286880898626697687587892264342959837212649817577862056895624746884122246147712285565387593364131673788676978239350915720151397102132506160498826459778147911279752746963313581794665472198713784883273444352664588147059289642958748692920179427815463464929388240394595605589412651288565402945714971533323645344913126478510405173263884785887354021605498747880337369307971839192628910697063914864375175264510824087788318492734419874229916253447905756507361706929896764319188997179124420113987758527287533847127291581327524935112848885938118758987195933908368983556325570471669781286649473589994976465104058616266729045802527774866796465353977527986672478243459222275891288728592925037899848811376802539706045461965763015695197798239331586866863183936146450377673726171389731134997513271814631586220848763905627859169897417809727174158248381841943448546469089806557631930288245146892129665589572522494118952292626866396528653253631333767697472769682776516457833408296979370848484352076885698683260377214691963402086401876451750363785272649499260977775879841138778997733436062762392629286322817822446307122963827232192369657764323171532553060807022305255993977552740621990266788309757955321187248596068988719892466223278759270514787864838337818512669787245488138197167241169383191786074389083551370892257655241835615653783476044956499435486479883292913924227995228871825465758959878873026361011534830449888208994472626344983705017359917187911867019556491161683267675867580947899908396494518988758514858612490214496565996308934433079799985634150277933638392196832515597482875901738238161461676135258694342854081598124614512273292677684347426235991776339947855783722366061545483366439806538449276577780302696135229207098595427853566201313544978728228792135243474195790382256992983902266841535838771421062981761947830796842854186404359563719394235768621942682579739678584801141965460486288854610759393972595728133457823939624911320988643753187703649648898382864206675444925156590549440358097882074129934492487202269918824349637357612885016494915324146501219113491862079633082916897733672921067965043251066658886805987447381419851871215695946939783559257404166918270597999587355292262649740435357166083276280678317165731251910722133416296127699336747864390234596812526663484357162509825501482217348949940234632127840252177361495171591508957885629853532646528966020584615326970143016306857706042964592206638382752313470357528889827449943661267435724493377715611216912944162334152345864959893794132928290703746336920786765112212467049254747208968648391881838699321688622282873152286138788917438185779109021694137763917585391406651693439641016526641166718495010106565504950749482148963487826133982271431587752464153105721467737606739896910852738526775125988539164766919751618818888623070622061108525141237442539606887394172482017302797988219714396719585545762763194614243547549551598634663601245633794273160762022585082167141749088331121212815185392601350669698712877424298318921858153774370761899858886466752699715636676401640782657564975189025959166676699462237367661583148238768601467844427539654214495412786399463837927869733132770626044256196245650746097513489449414909917839230209857737257374256702231553068993363494146556524129910364927546356209246164859619410658339713013607546863922539940871764854939331759117665898445223882664775471151161665264273206738396073606692808055393568149039773364827425328194284691668391525480571579162690818426332570308980818755321075484220232315346373419877962784836849465952261167248860587827458870639449162020594027522398355044539342169078742975254740462226424249594331463837402841583880556517257276917928722163154067895081954831599699272770283855788963139384226577428241706754559721459286431563981031276932262985319382159645922986785338558185757671324660897628637597443614425686756697413672977662638654248443456849549335252258409253611792609096208536458870118551175078985844449156281088491923386941401948807374753318678998111463584796878286537172603812897152558161218254663699252638475168913180602140339816402399306913665181776439899172821991688953754165261931348465148257526891997319774979708127681174451880385312374584944951482938128521845657178554196820974922443715498810177041367076743735244626739510641725717884642215688387956150743111954070324631998348724344527641901462813222544468129136623587256950235630638277651021303566167312779166166319338646818433865761614114924482632311971246202911423080742684554317585563524355479310339065991772529547581323165093822067312052197630619148658631539049663629766324667638379267465829684917596620549718294644469738476660593658276170646068892755361958875377456631358639714088773788635393837291757425931733652511769830709895711673483781465940746284494539319474845124498170783359622979188831914841476793995573128290203176439629641733213237415166329549677321214214317461165616674094359319871368282510739545706933142987413741867913839985701585367515624398526565313696514020784949304963835642609047599074717946499671271574303744451939595448549059331198383023505764514078531858686982503855318037117870184134891757469271233117638036881352777310678694552987753635924931911458445932829154783974589164866915842082843023431185769689224399986547146966511666188030208347184231564394842887204273787176303961119863927072173891613771496567878652262582974613864021253339342630649787289677442946481114836797393949522876117365674125981991562227934862933531217961263984195659323498921796704845875854635730774387864363123826438662158249621181292125861997226197415380284258463837139645253958271117156196918381289542452235219113199915367850543537731236776179147598512353251524917644995849206323513634983150805343512847322054484493341733521599443249103417355996417720207841332291762135885679216526781254494939643038245817472540751996927114197036582270102528866711351757147868885650319225703673872025855194745182889488921224296710298570723511682948573162456959195380583952261058529726805250314928144886643395892154633216914926122785783412188087843215241161517645851449443449647793589144604957899338185598276635242666651661334489687451666174347042386212486154282152982360755199328678776417531071285268301443154692354834827896879659614278344338718258847899651457854612223986828069835622192645734439535133448631729715876195503641545733894077121048664759347442189696741786106043865613813376549879198485295989981193942440821328542369964836293313665111419753886950124476491944881958222614614762406830228499664565891443941534381439191460578891933868872980813853268324389626505042812210778461717694984643569039427961123440825483895654387093157973728756773710179480858151458412593633381711252450633095352873427453158712943941497482333281143452844789767597656627342763912899988659607921202682576987733996105367683858412935827519374965129358199176898763823062803946211848468442321032663385924046251874765850212136665321602799438497134027785862265467925489573887699191357726938212622211808073242114716772938228538088938966991353537638153613943342682730979188385391951573785279681265543216954344614414538747843419773321343388628452823024188887373068227979676149347091928079463648987763738315813074463859653329633239919727435770249360499437532930348056485620116338401271141373482756421595311892535532758884258161195359577629705764171224877948706258582925848271599361525516721559835713575781274364193135311742932623529834314199535940851597532247383626642967526279213356381461407692671131448692182846667074661397252314846030833197505042844973767856513597469410959061547352895029344192568037489860891670483337554568128368617087599764314243741957621577249196138998794990718368889287474463211840419836579964246386807656684579363654525390342829142099637332511582995327357827624353549733187346622968136610385395353984179275995288341853966735473158894325904589963376627982469054806634271625704312976735683832254124111861849135347021944731731565261191831738965388147360873041967912549942621827941221303860424441308437716347706098831677392347442918125343212628382091762657701888283463896538581137689228774335568448108228235140526646718177853744941158495866734511611971244579505821611136137543305797159424507091957215177141394542615147139320485034899393905465177512155625708059796613943384802071651563965589199673302552439562979931631987194775628080682542212681758287704666816167576790839914944396733549423390412920393561607957993220528634833937654199804322712495945871526739856319662868749736831281694856835886901020291387312617407845732222161598644440954184352436402529217051865550476676503621671652983515371030906238167481877558194251759876887067624538304463969038933693388969944434363388349476924823678540679538859518439879862628947358937550522280176054654086205522363887527455377958602313449741739989477594158915218488996554808686552374908392197633326961288545384990228080343672972370148757997816908739483422132780605062299287201597551732907379995230615693523512661461918435193624118674324074574245407442782368622940967560414935569446615264995577459683331766737689867161972449584947927221602437662868205334868121359221223240836183701919779176979144685435601444783857595960234631829964748015315746549837335420499145807621848846199628432385273341855225501594405050856624105974706960708415406587231186901795741336593588196634119985456531503162676254394559232854624778799943779885224897146778743698333875599235119785511122212922855586836196929787487422259398589996305512355712504269897627893328641656489554598249704366536644326821802719143160969982529814492011896837982738514018255733282426386718525340219510628587761331992531636356181543879914707171329461823480342787767237118247207813175410321961895329871588279246776999873541141996923789902017926428488086641393551936725450371731278618682891118326276622179587756250855885234580254994504442273945699516203680414717741151246922205336519386538878371875117120941962221344882753428416713628686656956835174364134223724941747154995975536324833729327629968676188643768867602014654462801036647517841615287887295596845366792159928242815822616021391929532319425290498693962415994047382114814784444816982492176249747383743033862713473194891897764136134174213972378353528580713727754251609728227569182452525692564690259889678768696034622427859855978186581383405972777120601195986452717448916620153858504872531090884390733548203717889933137425498240132913351021624871312896148546696775773678269537406321716146984710492254708245554891195442638253828077534016912148791395495380621762364257684710428097405166524220799068584943977653955465511490912228278188115711108627551439686641985736263049936829729119112467853394805655601817893419195344623627989645128621836273692915564188272490881685391293338894649047141853765435615788396616435474289942834448337319348315865740189333206681983497409791243927717724468635494674212163853069264259568756733718156421285180365779405492362973687987337170431590764920768421465895564977343768991671185472583865805230932825923877615479747848549710843686249935767359387246122256776957764531521129659263748556435181354171299071565287175531624441667485293996719337209485646736843977641597171440363019194456722363558013598387574636177275957110537047438177389168302724516750542638402696544849346450778776772143717019515072285`