package genetic

// YstrMarkerTranslation contains the different names for
// a specific Y-STR marker that are used by different companies.
// It also specifies the name and index of the marker
// inside the Phylofriend program.
type YstrMarkerTranslation struct {
	// Index is the index of this marker inside Phylofriend.
	Index int
	// InternalName is the name that is used by this program.
	InternalName string
	// FTDNAName is the name that is used by Family Tree DNA.
	FTDNAName string
	//  YFullName is the name that is used by YFull.
	YFullName string
}

// YstrMarkerTable contains the Phylofriend index
// and the names for specific markers that are used by Phylofriend,
// FamilyTreeDNA and YFull.
//
// Format: index, Internal name, FTDNA name, YFull name.
//
// The first 111 entries contain Y-STR markers in Family Tree DNA
// order (used for 12, 37, 67, 111 marker tests, 2016-02-14).
//
// The following markers are additionally reported by YFull.
//
// The last four entries are extra space for DYS464.
// Most man have 4 markers at DYS464, but in some cases
// 6 have been found. I added an extra two fields just
// to be safe if some more are to be found.
//
// For more information on markers take a look at
// http://en.wikipedia.org/wiki/List_of_Y-STR_markers
// or http://isogg.org/wiki/Y-STR.
//
// General rules for translation:
// FTDNA uses a, b,... suffixes for palindromic markers.
// YFull uses .1, .2,... suffixes.
// FTDNA uses _ as a separator, YFull uses -.
//
// DYF406S1	at Family Tree DNA seems to be identical to DYF406 at YFull,
// and DYF395S1 to DYF395, but I am not totally sure.
var YstrMarkerTable []YstrMarkerTranslation = []YstrMarkerTranslation{
	{0, "DYS393", "DYS393", "DYS393"},
	{1, "DYS390", "DYS390", "DYS390"},
	{2, "DYS19", "DYS19", "DYS19"},
	{3, "DYS391", "DYS391", "DYS391"},
	{4, "DYS385a", "DYS385a", "DYS385.1"},
	{5, "DYS385b", "DYS385b", "DYS385.2"},
	{6, "DYS426", "DYS426", "DYS426"},
	{7, "DYS388", "DYS388", "DYS388"},
	{8, "DYS439", "DYS439", "DYS439"},
	{9, "DYS389i", "DYS389i", "DYS389I"},
	{10, "DYS392", "DYS392", "DYS392"},
	{11, "DYS389ii", "DYS389ii", "DYS389II"},
	{12, "DYS458", "DYS458", "DYS458"},
	{13, "DYS459a", "DYS459a", "DYS459.1"},
	{14, "DYS459b", "DYS459b", "DYS459.2"},
	{15, "DYS455", "DYS455", "DYS455"},
	{16, "DYS454", "DYS454", "DYS454"},
	{17, "DYS447", "DYS447", "DYS447"},
	{18, "DYS437", "DYS437", "DYS437"},
	{19, "DYS448", "DYS448", "DYS448"},
	{20, "DYS449", "DYS449", "DYS449"},
	{21, "DYS464a", "DYS464a", "DYS464.1"},
	{22, "DYS464b", "DYS464b", "DYS464.2"},
	{23, "DYS464c", "DYS464c", "DYS464.3"},
	{24, "DYS464d", "DYS464d", "DYS464.4"},
	{25, "DYS460", "DYS460", "DYS460"},
	{26, "Y_GATA_H4", "Y_GATA_H4", "Y-GATA-H4"},
	{27, "YCAIIa", "YCAIIa", "YCAII.1"},
	{28, "YCAIIb", "YCAIIb", "YCAII.2"},
	{29, "DYS456", "DYS456", "DYS456"},
	{30, "DYS607", "DYS607", "DYS607"},
	{31, "DYS576", "DYS576", "DYS576"},
	{32, "DYS570", "DYS570", "DYS570"},
	{33, "CDYa", "CDYa", "CDY.1"},
	{34, "CDYb", "CDYb", "CDY.2"},
	{35, "DYS442", "DYS442", "DYS442"},
	{36, "DYS438", "DYS438", "DYS438"},
	{37, "DYS531", "DYS531", "DYS531"},
	{38, "DYS578", "DYS578", "DYS578"},
	{39, "DYF395S1a", "DYF395S1a", "DYF395.1"},
	{40, "DYF395S1b", "DYF395S1b", "DYF395.2"},
	{41, "DYS590", "DYS590", "DYS590"},
	{42, "DYS537", "DYS537", "DYS537"},
	{43, "DYS641", "DYS641", "DYS641"},
	{44, "DYS472", "DYS472", "DYS472"},
	{45, "DYF406S1", "DYF406S1", "DYF406"},
	{46, "DYS511", "DYS511", "DYS511"},
	{47, "DYS425", "DYS425", "DYS425"},
	{48, "DYS413a", "DYS413a", "DYS413.1"},
	{49, "DYS413b", "DYS413b", "DYS413.2"},
	{50, "DYS557", "DYS557", "DYS557"},
	{51, "DYS594", "DYS594", "DYS594"},
	{52, "DYS436", "DYS436", "DYS436"},
	{53, "DYS490", "DYS490", "DYS490"},
	{54, "DYS534", "DYS534", "DYS534"},
	{55, "DYS450", "DYS450", "DYS450"},
	{56, "DYS444", "DYS444", "DYS444"},
	{57, "DYS481", "DYS481", "DYS481"},
	{58, "DYS520", "DYS520", "DYS520"},
	{59, "DYS446", "DYS446", "DYS446"},
	{60, "DYS617", "DYS617", "DYS617"},
	{61, "DYS568", "DYS568", "DYS568"},
	{62, "DYS487", "DYS487", "DYS487"},
	{63, "DYS572", "DYS572", "DYS572"},
	{64, "DYS640", "DYS640", "DYS640"},
	{65, "DYS492", "DYS492", "DYS492"},
	{66, "DYS565", "DYS565", "DYS565"},
	{67, "DYS710", "DYS710", "DYS710"},
	{68, "DYS485", "DYS485", "DYS485"},
	{69, "DYS632", "DYS632", "DYS632"},
	{70, "DYS495", "DYS495", "DYS495"},
	{71, "DYS540", "DYS540", "DYS540"},
	{72, "DYS714", "DYS714", "DYS714"},
	{73, "DYS716", "DYS716", "DYS716"},
	{74, "DYS717", "DYS717", "DYS717"},
	{75, "DYS505", "DYS505", "DYS505"},
	{76, "DYS556", "DYS556", "DYS556"},
	{77, "DYS549", "DYS549", "DYS549"},
	{78, "DYS589", "DYS589", "DYS589"},
	{79, "DYS522", "DYS522", "DYS522"},
	{80, "DYS494", "DYS494", "DYS494"},
	{81, "DYS533", "DYS533", "DYS533"},
	{82, "DYS636", "DYS636", "DYS636"},
	{83, "DYS575", "DYS575", "DYS575"},
	{84, "DYS638", "DYS638", "DYS638"},
	{85, "DYS462", "DYS462", "DYS462"},
	{86, "DYS452", "DYS452", "DYS452"},
	{87, "DYS445", "DYS445", "DYS445"},
	{88, "Y_GATA_A10", "Y_GATA_A10", "Y-GATA-A10"},
	{89, "DYS463", "DYS463", "DYS463"},
	{90, "DYS441", "DYS441", "DYS441"},
	{91, "Y_GGAAT_1B07", "Y_GGAAT_1B07", "Y-GGAAT-1B07"},
	{92, "DYS525", "DYS525", "DYS525"},
	{93, "DYS712", "DYS712", "DYS712"},
	{94, "DYS593", "DYS593", "DYS593"},
	{95, "DYS650", "DYS650", "DYS650"},
	{96, "DYS532", "DYS532", "DYS532"},
	{97, "DYS715", "DYS715", "DYS715"},
	{98, "DYS504", "DYS504", "DYS504"},
	{99, "DYS513", "DYS513", "DYS513"},
	{100, "DYS561", "DYS561", "DYS561"},
	{101, "DYS552", "DYS552", "DYS552"},
	{102, "DYS726", "DYS726", "DYS726"},
	{103, "DYS635", "DYS635", "DYS635"},
	{104, "DYS587", "DYS587", "DYS587"},
	{105, "DYS643", "DYS643", "DYS643"},
	{106, "DYS497", "DYS497", "DYS497"},
	{107, "DYS510", "DYS510", "DYS510"},
	{108, "DYS434", "DYS434", "DYS434"},
	{109, "DYS461", "DYS461", "DYS461"},
	{110, "DYS435", "DYS435", "DYS435"}, // End Family Tree DNA 111 markers.
	{111, "L1313", "L1313", "L1313"},    // Start YFull extra markers.
	{112, "L14", "L14", "L14"},
	{113, "ATA71D03", "ATA71D03", "ATA71D03"},
	{114, "DXYS156", "DXYS156", "DXYS156"},
	{115, "G09411", "G09411", "G09411"},
	{116, "DYS723", "DYS723", "DYS723"},
	{117, "DYS722", "DYS722", "DYS722"},
	{118, "DYS721", "DYS721", "DYS721"},
	{119, "DYS720", "DYS720", "DYS720"},
	{120, "DYS719", "DYS719", "DYS719"},
	{121, "DYS718", "DYS718", "DYS718"},
	{122, "DYS713", "DYS713", "DYS713"},
	{123, "DYS711", "DYS711", "DYS711"},
	{124, "DYS709", "DYS709", "DYS709"},
	{125, "DYS708", "DYS708", "DYS708"},
	{126, "DYS707", "DYS707", "DYS707"},
	{127, "DYS706", "DYS706", "DYS706"},
	{128, "DYS705", "DYS705", "DYS705"},
	{129, "DYS703", "DYS703", "DYS703"},
	{130, "DYS702", "DYS702", "DYS702"},
	{131, "DYS701", "DYS701", "DYS701"},
	{132, "DYS696", "DYS696", "DYS696"},
	{133, "DYS695", "DYS695", "DYS695"},
	{134, "DYS694", "DYS694", "DYS694"},
	{135, "DYS692", "DYS692", "DYS692"},
	{136, "DYS688", "DYS688", "DYS688"},
	{137, "DYS687", "DYS687", "DYS687"},
	{138, "DYS686", "DYS686", "DYS686"},
	{139, "DYS685", "DYS685", "DYS685"},
	{140, "DYS684", "DYS684", "DYS684"},
	{141, "DYS683", "DYS683", "DYS683"},
	{142, "DYS681", "DYS681", "DYS681"},
	{143, "DYS679", "DYS679", "DYS679"},
	{144, "DYS678", "DYS678", "DYS678"},
	{145, "DYS677", "DYS677", "DYS677"},
	{146, "DYS676", "DYS676", "DYS676"},
	{147, "DYS675", "DYS675", "DYS675"},
	{148, "DYS673", "DYS673", "DYS673"},
	{149, "DYS672", "DYS672", "DYS672"},
	{150, "DYS668", "DYS668", "DYS668"},
	{151, "DYS667", "DYS667", "DYS667"},
	{152, "DYS666", "DYS666", "DYS666"},
	{153, "DYS664", "DYS664", "DYS664"},
	{154, "DYS662", "DYS662", "DYS662"},
	{155, "DYS656", "DYS656", "DYS656"},
	{156, "DYS655", "DYS655", "DYS655"},
	{157, "DYS651", "DYS651", "DYS651"},
	{158, "DYS649", "DYS649", "DYS649"},
	{159, "DYS645", "DYS645", "DYS645"},
	{160, "DYS644", "DYS644", "DYS644"},
	{161, "DYS642", "DYS642", "DYS642"},
	{162, "DYS639", "DYS639", "DYS639"},
	{163, "DYS637", "DYS637", "DYS637"},
	{164, "DYS634", "DYS634", "DYS634"},
	{165, "DYS633", "DYS633", "DYS633"},
	{166, "DYS631", "DYS631", "DYS631"},
	{167, "DYS630", "DYS630", "DYS630"},
	{168, "DYS629", "DYS629", "DYS629"},
	{169, "DYS627", "DYS627", "DYS627"},
	{170, "DYS626", "DYS626", "DYS626"},
	{171, "DYS625", "DYS625", "DYS625"},
	{172, "DYS624", "DYS624", "DYS624"},
	{173, "DYS623", "DYS623", "DYS623"},
	{174, "DYS622", "DYS622", "DYS622"},
	{175, "DYS621", "DYS621", "DYS621"},
	{176, "DYS620", "DYS620", "DYS620"},
	{177, "DYS619", "DYS619", "DYS619"},
	{178, "DYS618", "DYS618", "DYS618"},
	{179, "DYS616", "DYS616", "DYS616"},
	{180, "DYS615", "DYS615", "DYS615"},
	{181, "DYS614", "DYS614", "DYS614"},
	{182, "DYS613", "DYS613", "DYS613"},
	{183, "DYS612", "DYS612", "DYS612"},
	{184, "DYS611", "DYS611", "DYS611"},
	{185, "DYS609", "DYS609", "DYS609"},
	{186, "DYS608", "DYS608", "DYS608"},
	{187, "DYS600", "DYS600", "DYS600"},
	{188, "DYS599", "DYS599", "DYS599"},
	{189, "DYS598", "DYS598", "DYS598"},
	{190, "DYS596", "DYS596", "DYS596"},
	{191, "DYS595", "DYS595", "DYS595"},
	{192, "DYS592", "DYS592", "DYS592"},
	{193, "DYS588", "DYS588", "DYS588"},
	{194, "DYS585", "DYS585", "DYS585"},
	{195, "DYS584", "DYS584", "DYS584"},
	{196, "DYS583", "DYS583", "DYS583"},
	{197, "DYS582", "DYS582", "DYS582"},
	{198, "DYS581", "DYS581", "DYS581"},
	{199, "DYS580", "DYS580", "DYS580"},
	{200, "DYS579", "DYS579", "DYS579"},
	{201, "DYS577", "DYS577", "DYS577"},
	{202, "DYS574", "DYS574", "DYS574"},
	{203, "DYS573", "DYS573", "DYS573"},
	{204, "DYS571", "DYS571", "DYS571"},
	{205, "DYS569", "DYS569", "DYS569"},
	{206, "DYS567", "DYS567", "DYS567"},
	{207, "DYS562", "DYS562", "DYS562"},
	{208, "DYS559", "DYS559", "DYS559"},
	{209, "DYS558", "DYS558", "DYS558"},
	{210, "DYS554", "DYS554", "DYS554"},
	{211, "DYS551", "DYS551", "DYS551"},
	{212, "DYS550", "DYS550", "DYS550"},
	{213, "DYS548", "DYS548", "DYS548"},
	{214, "DYS547", "DYS547", "DYS547"},
	{215, "DYS546", "DYS546", "DYS546"},
	{216, "DYS545", "DYS545", "DYS545"},
	{217, "DYS544", "DYS544", "DYS544"},
	{218, "DYS543", "DYS543", "DYS543"},
	{219, "DYS542", "DYS542", "DYS542"},
	{220, "DYS541", "DYS541", "DYS541"},
	{221, "DYS539", "DYS539", "DYS539"},
	{222, "DYS538", "DYS538", "DYS538"},
	{223, "DYS536", "DYS536", "DYS536"},
	{224, "DYS530", "DYS530", "DYS530"},
	{225, "DYS523", "DYS523", "DYS523"},
	{226, "DYS521", "DYS521", "DYS521"},
	{227, "DYS518", "DYS518", "DYS518"},
	{228, "DYS517", "DYS517", "DYS517"},
	{229, "DYS516", "DYS516", "DYS516"},
	{230, "DYS514", "DYS514", "DYS514"},
	{231, "DYS512", "DYS512", "DYS512"},
	{232, "DYS509", "DYS509", "DYS509"},
	{233, "DYS508", "DYS508", "DYS508"},
	{234, "DYS507", "DYS507", "DYS507"},
	{235, "DYS506", "DYS506", "DYS506"},
	{236, "DYS502", "DYS502", "DYS502"},
	{237, "DYS501", "DYS501", "DYS501"},
	{238, "DYS500", "DYS500", "DYS500"},
	{239, "DYS499", "DYS499", "DYS499"},
	{240, "DYS496", "DYS496", "DYS496"},
	{241, "DYS493", "DYS493", "DYS493"},
	{242, "DYS491", "DYS491", "DYS491"},
	{243, "DYS489", "DYS489", "DYS489"},
	{244, "DYS488", "DYS488", "DYS488"},
	{245, "DYS484", "DYS484", "DYS484"},
	{246, "DYS480", "DYS480", "DYS480"},
	{247, "DYS478", "DYS478", "DYS478"},
	{248, "DYS477", "DYS477", "DYS477"},
	{249, "DYS476", "DYS476", "DYS476"},
	{250, "DYS475", "DYS475", "DYS475"},
	{251, "DYS474", "DYS474", "DYS474"},
	{252, "DYS473", "DYS473", "DYS473"},
	{253, "DYS471", "DYS471", "DYS471"},
	{254, "DYS470", "DYS470", "DYS470"},
	{255, "DYS469", "DYS469", "DYS469"},
	{256, "DYS468", "DYS468", "DYS468"},
	{257, "DYS467", "DYS467", "DYS467"},
	{258, "DYS466", "DYS466", "DYS466"},
	{259, "DYS453", "DYS453", "DYS453"},
	{260, "DYS443", "DYS443", "DYS443"},
	{261, "DYF394", "DYF394", "DYF394"},
	{262, "DYF393", "DYF393", "DYF393"},
	{263, "DYF392", "DYF392", "DYF392"},
	{264, "DYF389", "DYF389", "DYF389"},
	{265, "DYF382", "DYF382", "DYF382"},
	{266, "DYR89", "DYR89", "DYR89"},
	{267, "DYR87", "DYR87", "DYR87"},
	{268, "DYR85", "DYR85", "DYR85"},
	{269, "DYR84", "DYR84", "DYR84"},
	{270, "DYR83", "DYR83", "DYR83"},
	{271, "DYR82", "DYR82", "DYR82"},
	{272, "DYR81", "DYR81", "DYR81"},
	{273, "DYR80", "DYR80", "DYR80"},
	{274, "DYR79", "DYR79", "DYR79"},
	{275, "DYR78", "DYR78", "DYR78"},
	{276, "DYR77", "DYR77", "DYR77"},
	{277, "DYR76", "DYR76", "DYR76"},
	{278, "DYR75", "DYR75", "DYR75"},
	{279, "DYR74", "DYR74", "DYR74"},
	{280, "DYR73", "DYR73", "DYR73"},
	{281, "DYR71", "DYR71", "DYR71"},
	{282, "DYR70", "DYR70", "DYR70"},
	{283, "DYR69", "DYR69", "DYR69"},
	{284, "DYR175", "DYR175", "DYR175"},
	{285, "DYR174", "DYR174", "DYR174"},
	{286, "DYR173", "DYR173", "DYR173"},
	{287, "DYR172", "DYR172", "DYR172"},
	{288, "DYR171", "DYR171", "DYR171"},
	{289, "DYR170", "DYR170", "DYR170"},
	{290, "DYR169", "DYR169", "DYR169"},
	{291, "DYR168", "DYR168", "DYR168"},
	{292, "DYR167", "DYR167", "DYR167"},
	{293, "DYR166", "DYR166", "DYR166"},
	{294, "DYR163", "DYR163", "DYR163"},
	{295, "DYR165", "DYR165", "DYR165"},
	{296, "DYR164", "DYR164", "DYR164"},
	{297, "DYR162", "DYR162", "DYR162"},
	{298, "DYR161", "DYR161", "DYR161"},
	{299, "DYR160", "DYR160", "DYR160"},
	{300, "DYR159", "DYR159", "DYR159"},
	{301, "DYR158", "DYR158", "DYR158"},
	{302, "DYR157", "DYR157", "DYR157"},
	{303, "DYR156", "DYR156", "DYR156"},
	{304, "DYR154", "DYR154", "DYR154"},
	{305, "DYR152", "DYR152", "DYR152"},
	{306, "DYR146", "DYR146", "DYR146"},
	{307, "DYR150", "DYR150", "DYR150"},
	{308, "DYR144", "DYR144", "DYR144"},
	{309, "DYR143", "DYR143", "DYR143"},
	{310, "DYR139", "DYR139", "DYR139"},
	{311, "DYR138", "DYR138", "DYR138"},
	{312, "DYR137", "DYR137", "DYR137"},
	{313, "DYR136", "DYR136", "DYR136"},
	{314, "DYR135", "DYR135", "DYR135"},
	{315, "DYR131", "DYR131", "DYR131"},
	{316, "DYR130", "DYR130", "DYR130"},
	{317, "DYR127", "DYR127", "DYR127"},
	{318, "DYR126", "DYR126", "DYR126"},
	{319, "DYR123", "DYR123", "DYR123"},
	{320, "DYR120", "DYR120", "DYR120"},
	{321, "DYR119", "DYR119", "DYR119"},
	{322, "DYR118", "DYR118", "DYR118"},
	{323, "DYR117", "DYR117", "DYR117"},
	{324, "DYR116", "DYR116", "DYR116"},
	{325, "DYR115", "DYR115", "DYR115"},
	{326, "DYR113", "DYR113", "DYR113"},
	{327, "DYR112", "DYR112", "DYR112"},
	{328, "DYR111", "DYR111", "DYR111"},
	{329, "DYR110", "DYR110", "DYR110"},
	{330, "DYR108", "DYR108", "DYR108"},
	{331, "DYR107", "DYR107", "DYR107"},
	{332, "DYR106", "DYR106", "DYR106"},
	{333, "DYR105", "DYR105", "DYR105"},
	{334, "DYR104", "DYR104", "DYR104"},
	{335, "DYR103", "DYR103", "DYR103"},
	{336, "DYR102", "DYR102", "DYR102"},
	{337, "DYR101", "DYR101", "DYR101"},
	{338, "DYR100", "DYR100", "DYR100"},
	{339, "DYR99", "DYR99", "DYR99"},
	{340, "DYR97", "DYR97", "DYR97"},
	{341, "DYR96", "DYR96", "DYR96"},
	{342, "DYR95", "DYR95", "DYR95"},
	{343, "DYR94", "DYR94", "DYR94"},
	{344, "DYR93", "DYR93", "DYR93"},
	{345, "DYR92", "DYR92", "DYR92"},
	{346, "DYR91", "DYR91", "DYR91"},
	{347, "DYR90", "DYR90", "DYR90"},
	{348, "DYR65", "DYR65", "DYR65"},
	{349, "DYR62", "DYR62", "DYR62"},
	{350, "DYR61", "DYR61", "DYR61"},
	{351, "DYR60", "DYR60", "DYR60"},
	{352, "DYR59", "DYR59", "DYR59"},
	{353, "DYR57", "DYR57", "DYR57"},
	{354, "DYR56", "DYR56", "DYR56"},
	{355, "DYR55", "DYR55", "DYR55"},
	{356, "DYR54", "DYR54", "DYR54"},
	{357, "DYR52", "DYR52", "DYR52"},
	{358, "DYR51", "DYR51", "DYR51"},
	{359, "DYR49", "DYR49", "DYR49"},
	{360, "DYR48", "DYR48", "DYR48"},
	{361, "DYR47", "DYR47", "DYR47"},
	{362, "DYR46", "DYR46", "DYR46"},
	{363, "DYR44", "DYR44", "DYR44"},
	{364, "DYR43", "DYR43", "DYR43"},
	{365, "DYR41", "DYR41", "DYR41"},
	{366, "DYR40", "DYR40", "DYR40"},
	{367, "DYR39", "DYR39", "DYR39"},
	{368, "DYR33", "DYR33", "DYR33"},
	{369, "DYR32", "DYR32", "DYR32"},
	{370, "DYR31", "DYR31", "DYR31"},
	{371, "DYR30", "DYR30", "DYR30"},
	{372, "DYR29", "DYR29", "DYR29"},
	{373, "DYR28", "DYR28", "DYR28"},
	{374, "DYR27", "DYR27", "DYR27"},
	{375, "DYR26", "DYR26", "DYR26"},
	{376, "DYR23", "DYR23", "DYR23"},
	{377, "DYR20", "DYR20", "DYR20"},
	{378, "DYR19", "DYR19", "DYR19"},
	{379, "DYR15", "DYR15", "DYR15"},
	{380, "DYR14", "DYR14", "DYR14"},
	{381, "DYR13", "DYR13", "DYR13"},
	{382, "DYR12", "DYR12", "DYR12"},
	{383, "DYR10", "DYR10", "DYR10"},
	{384, "DYR8", "DYR8", "DYR8"},
	{385, "DYR7", "DYR7", "DYR7"},
	{386, "DYR6", "DYR6", "DYR6"},
	{387, "DYR5", "DYR5", "DYR5"},
	{388, "DYR3", "DYR3", "DYR3"},
	{389, "DYR2", "DYR2", "DYR2"},
	{390, "DYR1", "DYR1", "DYR1"},
	{391, "DYS526A", "DYS526A", "DYS526A"},
	{392, "DYS526B", "DYS526B", "DYS526B"},
	{393, "DYS527.1", "DYS527.1", "DYS527.1"},
	{394, "DYS527.2", "DYS527.2", "DYS527.2"},
	{395, "DYS528.1", "DYS528.1", "DYS528.1"},
	{396, "DYS528.2", "DYS528.2", "DYS528.2"},
	{397, "DYS725.1", "DYS725.1", "DYS725.1"},
	{398, "DYS725.2", "DYS725.2", "DYS725.2"},
	{399, "DYS725.3", "DYS725.3", "DYS725.3"},
	{400, "DYS725.4", "DYS725.4", "DYS725.4"},
	{401, "DYF371.1", "DYF371.1", "DYF371.1"},
	{402, "DYF371.2", "DYF371.2", "DYF371.2"},
	{403, "DYF371.3", "DYF371.3", "DYF371.3"},
	{404, "DYF371.4", "DYF371.4", "DYF371.4"},
	{405, "DYF380.1", "DYF380.1", "DYF380.1"},
	{406, "DYF380.2", "DYF380.2", "DYF380.2"},
	{407, "DYF381.1", "DYF381.1", "DYF381.1"},
	{408, "DYF381.2", "DYF381.2", "DYF381.2"},
	{409, "DYF383.1", "DYF383.1", "DYF383.1"},
	{410, "DYF383.2", "DYF383.2", "DYF383.2"},
	{411, "DYF384.1", "DYF384.1", "DYF384.1"},
	{412, "DYF384.2", "DYF384.2", "DYF384.2"},
	{413, "DYF385.1", "DYF385.1", "DYF385.1"},
	{414, "DYF385.2", "DYF385.2", "DYF385.2"},
	{415, "DYF386.1", "DYF386.1", "DYF386.1"},
	{416, "DYF386.2", "DYF386.2", "DYF386.2"},
	{417, "DYF386.3", "DYF386.3", "DYF386.3"},
	{418, "DYF386.4", "DYF386.4", "DYF386.4"},
	{419, "DYF387.1", "DYF387.1", "DYF387.1"},
	{420, "DYF387.2", "DYF387.2", "DYF387.2"},
	{421, "DYF390.1", "DYF390.1", "DYF390.1"},
	{422, "DYF390.2", "DYF390.2", "DYF390.2"},
	{423, "DYF391.1", "DYF391.1", "DYF391.1"},
	{424, "DYF391.2", "DYF391.2", "DYF391.2"},
	{425, "DYF396.1", "DYF396.1", "DYF396.1"},
	{426, "DYF396.2", "DYF396.2", "DYF396.2"},
	{427, "DYF398.1", "DYF398.1", "DYF398.1"},
	{428, "DYF398.2", "DYF398.2", "DYF398.2"},
	{429, "DYF399.1", "DYF399.1", "DYF399.1"},
	{430, "DYF399.2", "DYF399.2", "DYF399.2"},
	{431, "DYF399.3", "DYF399.3", "DYF399.3"},
	{432, "DYF400.1", "DYF400.1", "DYF400.1"},
	{433, "DYF400.2", "DYF400.2", "DYF400.2"},
	{434, "DYF401.1", "DYF401.1", "DYF401.1"},
	{435, "DYF401.2", "DYF401.2", "DYF401.2"},
	{436, "DYF403.1", "DYF403.1", "DYF403.1"},
	{437, "DYF403.2", "DYF403.2", "DYF403.2"},
	{438, "DYF404.1", "DYF404.1", "DYF404.1"},
	{439, "DYF404.2", "DYF404.2", "DYF404.2"},
	{440, "DYF405.1", "DYF405.1", "DYF405.1"},
	{441, "DYF405.2", "DYF405.2", "DYF405.2"},
	{442, "DYF407.1", "DYF407.1", "DYF407.1"},
	{443, "DYF407.2", "DYF407.2", "DYF407.2"},
	{444, "DYF408.1", "DYF408.1", "DYF408.1"},
	{445, "DYF408.2", "DYF408.2", "DYF408.2"},
	{446, "DYF409.1", "DYF409.1", "DYF409.1"},
	{447, "DYF409.2", "DYF409.2", "DYF409.2"},
	{448, "DYF410.1", "DYF410.1", "DYF410.1"},
	{449, "DYF410.2", "DYF410.2", "DYF410.2"},
	{450, "DYF411.1", "DYF411.1", "DYF411.1"},
	{451, "DYF411.2", "DYF411.2", "DYF411.2"},
	{452, "DYF412.1", "DYF412.1", "DYF412.1"},
	{453, "DYF412.2", "DYF412.2", "DYF412.2"},
	{454, "DYR9.1", "DYR9.1", "DYR9.1"},
	{455, "DYR9.2", "DYR9.2", "DYR9.2"},
	{456, "DYR17.1", "DYR17.1", "DYR17.1"},
	{457, "DYR17.2", "DYR17.2", "DYR17.2"},
	{458, "DYR17.3", "DYR17.3", "DYR17.3"},
	{459, "DYR18.1", "DYR18.1", "DYR18.1"},
	{460, "DYR18.2", "DYR18.2", "DYR18.2"},
	{461, "DYR35.1", "DYR35.1", "DYR35.1"},
	{462, "DYR35.2", "DYR35.2", "DYR35.2"},
	{463, "DYR36.1", "DYR36.1", "DYR36.1"},
	{464, "DYR36.2", "DYR36.2", "DYR36.2"},
	{465, "DYR38.1", "DYR38.1", "DYR38.1"},
	{466, "DYR38.2", "DYR38.2", "DYR38.2"},
	{467, "DYR45.1", "DYR45.1", "DYR45.1"},
	{468, "DYR45.2", "DYR45.2", "DYR45.2"},
	{469, "DYR45.3", "DYR45.3", "DYR45.3"},
	{470, "DYR58.1", "DYR58.1", "DYR58.1"},
	{471, "DYR58.2", "DYR58.2", "DYR58.2"},
	{472, "DYR63.1", "DYR63.1", "DYR63.1"},
	{473, "DYR63.2", "DYR63.2", "DYR63.2"},
	{474, "DYR64.1", "DYR64.1", "DYR64.1"},
	{475, "DYR64.2", "DYR64.2", "DYR64.2"},
	{476, "DYR66.1", "DYR66.1", "DYR66.1"},
	{477, "DYR66.2", "DYR66.2", "DYR66.2"},
	{478, "DYR67.1", "DYR67.1", "DYR67.1"},
	{479, "DYR67.2", "DYR67.2", "DYR67.2"},
	{480, "DYR67.3", "DYR67.3", "DYR67.3"},
	{481, "DYR67.4", "DYR67.4", "DYR67.4"},
	{482, "DYR68.1", "DYR68.1", "DYR68.1"},
	{483, "DYR68.2", "DYR68.2", "DYR68.2"},
	{484, "DYR68.3", "DYR68.3", "DYR68.3"},
	{485, "DYR68.4", "DYR68.4", "DYR68.4"},
	{486, "DYR88.1", "DYR88.1", "DYR88.1"},
	{487, "DYR88.2", "DYR88.2", "DYR88.2"},
	{488, "DYR121.1", "DYR121.1", "DYR121.1"},
	{489, "DYR121.2", "DYR121.2", "DYR121.2"},
	{490, "DYR122.1", "DYR122.1", "DYR122.1"},
	{491, "DYR122.2", "DYR122.2", "DYR122.2"},
	{492, "DYR124.1", "DYR124.1", "DYR124.1"},
	{493, "DYR124.2", "DYR124.2", "DYR124.2"},
	{494, "DYR124.3", "DYR124.3", "DYR124.3"},
	{495, "DYR125.1", "DYR125.1", "DYR125.1"},
	{496, "DYR125.2", "DYR125.2", "DYR125.2"},
	{497, "DYR128.1", "DYR128.1", "DYR128.1"},
	{498, "DYR128.2", "DYR128.2", "DYR128.2"},
	{499, "DYR132.1", "DYR132.1", "DYR132.1"},
	{500, "DYR132.2", "DYR132.2", "DYR132.2"}, // End YFull extra markers.
	{501, "DYS464e", "DYS464e", "DYS464.6"},   // Start extra space for DYS464 (very rare).
	{502, "DYS464f", "DYS464f", "DYS464.7"},
	{503, "DYS464g", "DYS464g", "DYS464.8"},
	{504, "DYS464h", "DYS464h", "DYS464.9"},
}
