package constants

// CountryCode represents Tunisia's international calling code
const CountryCode = "+216"

// Carrier represents a mobile carrier in Tunisia
type Carrier struct {
	Name     string
	Prefixes []string
}

// Bank represents a Tunisian bank
type Bank struct {
	Name string
	Code string
}

// Governorate represents a Tunisian governorate
type Governorate struct {
	Name       string
	PostalCode string
	Region     string
}

// Carriers contains all Tunisian mobile carriers and their prefixes
var Carriers = map[string]Carrier{
	"OOREDOO": {
		Name:     "Ooredoo Tunisia",
		Prefixes: []string{"2", "5"},
	},
	"ORANGE": {
		Name:     "Orange Tunisia",
		Prefixes: []string{"4"},
	},
	"TELECOM": {
		Name:     "Tunisie Telecom",
		Prefixes: []string{"9"},
	},
}

// ValidPrefixes contains all valid mobile prefixes
var ValidPrefixes = []string{"2", "4", "5", "9"}

// Banks contains major Tunisian banks
var Banks = map[string]Bank{
	"01": {Name: "Banque Centrale de Tunisie", Code: "01"},
	"02": {Name: "Banque de Tunisie", Code: "02"},
	"03": {Name: "Banque Internationale Arabe de Tunisie", Code: "03"},
	"04": {Name: "Banque de l'Habitat", Code: "04"},
	"05": {Name: "Banque Nationale Agricole", Code: "05"},
	"07": {Name: "Société Tunisienne de Banque", Code: "07"},
	"08": {Name: "Union Bancaire pour le Commerce et l'Industrie", Code: "08"},
	"10": {Name: "Stusid Bank", Code: "10"},
	"11": {Name: "Banque Zitouna", Code: "11"},
	"12": {Name: "Banque de Financement des Petites et Moyennes Entreprises", Code: "12"},
	"14": {Name: "Citibank", Code: "14"},
	"16": {Name: "Arab Tunisian Bank", Code: "16"},
	"17": {Name: "Banque Attijari de Tunisie", Code: "17"},
	"20": {Name: "Amen Bank", Code: "20"},
	"21": {Name: "Banque Franco-Tunisienne", Code: "21"},
	"23": {Name: "Banque Tuniso-Libyenne", Code: "23"},
	"24": {Name: "Union Internationale de Banques", Code: "24"},
	"25": {Name: "Banque Tuniso-Koweïtienne", Code: "25"},
	"26": {Name: "Banque Arabe Tuniso-Saoudienne", Code: "26"},
	"28": {Name: "Banque Tunisienne de Solidarité", Code: "28"},
	"29": {Name: "Banque Tuniso-Qatarie d'Investissement", Code: "29"},
	"32": {Name: "Banque Internationale de Commerce et d'Industrie de Tunisie", Code: "32"},
	"34": {Name: "Banque Européenne pour la Tunisie", Code: "34"},
	"35": {Name: "Banque de Coopération du Maghreb Arabe", Code: "35"},
	"38": {Name: "Banque Tunisienne des Participations", Code: "38"},
	"39": {Name: "Banque Al Baraka d'Investissement", Code: "39"},
	"47": {Name: "Banque Tuniso-Emiratie d'Investissement", Code: "47"},
	"81": {Name: "Poste Tunisienne", Code: "81"},
}

// Governorates contains all Tunisian governorates
var Governorates = map[string]Governorate{
	"TUNIS": {
		Name:       "Tunis",
		PostalCode: "1000",
		Region:     "North",
	},
	"ARIANA": {
		Name:       "Ariana",
		PostalCode: "2000",
		Region:     "North",
	},
	"BEN_AROUS": {
		Name:       "Ben Arous",
		PostalCode: "2013",
		Region:     "North",
	},
	"MANOUBA": {
		Name:       "Manouba",
		PostalCode: "2010",
		Region:     "North",
	},
	"NABEUL": {
		Name:       "Nabeul",
		PostalCode: "8000",
		Region:     "North",
	},
	"ZAGHOUAN": {
		Name:       "Zaghouan",
		PostalCode: "1100",
		Region:     "North",
	},
	"BIZERTE": {
		Name:       "Bizerte",
		PostalCode: "7000",
		Region:     "North",
	},
	"BEJA": {
		Name:       "Béja",
		PostalCode: "9000",
		Region:     "North",
	},
	"JENDOUBA": {
		Name:       "Jendouba",
		PostalCode: "8100",
		Region:     "North",
	},
	"KEF": {
		Name:       "Le Kef",
		PostalCode: "7100",
		Region:     "North",
	},
	"SILIANA": {
		Name:       "Siliana",
		PostalCode: "6100",
		Region:     "North",
	},
	"SOUSSE": {
		Name:       "Sousse",
		PostalCode: "4000",
		Region:     "Center",
	},
	"MONASTIR": {
		Name:       "Monastir",
		PostalCode: "5000",
		Region:     "Center",
	},
	"MAHDIA": {
		Name:       "Mahdia",
		PostalCode: "5100",
		Region:     "Center",
	},
	"SFAX": {
		Name:       "Sfax",
		PostalCode: "3000",
		Region:     "Center",
	},
	"KAIROUAN": {
		Name:       "Kairouan",
		PostalCode: "3100",
		Region:     "Center",
	},
	"KASSERINE": {
		Name:       "Kasserine",
		PostalCode: "1200",
		Region:     "Center",
	},
	"SIDI_BOUZID": {
		Name:       "Sidi Bouzid",
		PostalCode: "9100",
		Region:     "Center",
	},
	"GABES": {
		Name:       "Gabès",
		PostalCode: "6000",
		Region:     "South",
	},
	"MEDENINE": {
		Name:       "Médenine",
		PostalCode: "4100",
		Region:     "South",
	},
	"TATAOUINE": {
		Name:       "Tataouine",
		PostalCode: "3200",
		Region:     "South",
	},
	"GAFSA": {
		Name:       "Gafsa",
		PostalCode: "2100",
		Region:     "South",
	},
	"TOZEUR": {
		Name:       "Tozeur",
		PostalCode: "2200",
		Region:     "South",
	},
	"KEBILI": {
		Name:       "Kébili",
		PostalCode: "4200",
		Region:     "South",
	},
}
