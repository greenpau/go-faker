// Copyright 2022 Paul Greenberg greenpau@outlook.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package phones

// North American Numbering Plan (NANP)

var Classifications = []string{
	"Landline",
	"Cellular",
	"VoIP",
	"Home",
	"Cell",
	"Business",
	"Fax",
}

var NanpAreaCodes = map[string][]string{
	"AK": []string{"907"},
	"AL": []string{"205", "251", "256", "334", "659", "938"},
	"AR": []string{"327", "479", "501", "870"},
	"AS": []string{"684"},
	"AZ": []string{"480", "520", "602", "623", "928"},
	"CA": []string{
		"209", "213", "310", "323", "408", "415", "424", "442", "510", "530",
		"559", "562", "619", "626", "650", "657", "661", "669", "707", "714",
		"747", "760", "805", "818", "831", "858", "909", "916", "925", "949",
		"951",
	},
	"CO": []string{"303", "719", "720", "970"},
	"CT": []string{"203", "475", "860", "959"},
	"DC": []string{"202"},
	"DE": []string{"302"},
	"FL": []string{
		"239", "305", "321", "352", "386", "407", "561", "689", "727",
		"754", "772", "786", "813", "850", "863", "904", "941", "954",
	},
	"GA": []string{"229", "404", "470", "478", "678", "706", "762", "770", "912"},
	"GU": []string{"671"},
	"HI": []string{"808"},
	"IA": []string{"319", "515", "563", "641", "712"},
	"ID": []string{"208"},
	"IL": []string{
		"217", "224", "309", "312", "331", "447", "464", "618",
		"630", "708", "730", "773", "779", "815", "847", "872",
	},
	"IN": []string{"219", "260", "317", "574", "765", "812"},
	"KS": []string{"316", "620", "785", "913"},
	"KY": []string{"270", "364", "502", "606", "859"},
	"LA": []string{"225", "318", "337", "504", "985"},
	"MA": []string{"339", "351", "413", "508", "617", "774", "781", "857", "978"},
	"MD": []string{"227", "240", "301", "410", "443", "667"},
	"ME": []string{"207"},
	"MI": []string{
		"231", "248", "269", "313", "517", "586",
		"616", "679", "734", "810", "906", "947", "989",
	},
	"MN": []string{"218", "320", "507", "612", "651", "763", "952"},
	"MO": []string{"314", "417", "557", "573", "636", "660", "816", "975"},
	"MS": []string{"228", "601", "662", "769"},
	"MT": []string{"406"},
	"NC": []string{"252", "336", "704", "828", "910", "919", "980", "984"},
	"ND": []string{"701"},
	"NE": []string{"308", "402", "531"},
	"NH": []string{"603"},
	"NJ": []string{"201", "551", "609", "732", "848", "856", "862", "908", "973"},
	"NM": []string{"505", "575"},
	"NV": []string{"702", "725", "775"},
	"NY": []string{
		"212", "315", "347", "516", "518", "585", "607", "631",
		"646", "716", "718", "845", "914", "917", "929",
	},
	"OH": []string{
		"216", "234", "283", "330", "380", "419", "440",
		"513", "567", "614", "740", "937",
	},
	"OK": []string{"405", "539", "580", "918"},
	"OR": []string{"458", "503", "541", "971"},
	"PA": []string{
		"215", "267", "272", "412", "484", "570", "610",
		"717", "724", "814", "878",
	},
	"RI": []string{"401"},
	"SC": []string{"803", "843", "864"},
	"SD": []string{"605"},
	"TN": []string{"423", "615", "731", "865", "901", "931"},
	"TX": []string{
		"210", "214", "254", "281", "325", "361", "409", "430", "432", "469",
		"512", "682", "713", "737", "806", "817", "830", "832", "903", "915",
		"936", "940", "956", "972", "979",
	},
	"UT": []string{"385", "435", "801"},
	"VA": []string{"276", "434", "540", "571", "703", "757", "804"},
	"VT": []string{"802"},
	"WA": []string{"206", "253", "360", "425", "509", "564"},
	"WI": []string{"262", "274", "414", "534", "608", "715", "920"},
	"WV": []string{"304", "681"},
	"WY": []string{"307"},
}
