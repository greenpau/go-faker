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

package govid

var SocialSecurityNumberPrefix = map[string][]string{
	"NH": []string{"001-003"},
	"ME": []string{"004-007"},
	"VT": []string{"008-009"},
	"MA": []string{"010-034"},
	"RI": []string{"035-039"},
	"CT": []string{"040-049"},
	"NY": []string{"050-134"},
	"NJ": []string{"135-158"},
	"PA": []string{"159-211"},
	"MD": []string{"212-220"},
	"DE": []string{"221-222"},
	"VA": []string{"223-231"},
	"NC": []string{"232"},
	"WV": []string{"232-236"},
	"SC": []string{"247-251"},
	"GA": []string{"252-260"},
	"FL": []string{"261-267"},
	"OH": []string{"268-302"},
	"IN": []string{"303-317"},
	"IL": []string{"318-361"},
	"MI": []string{"362-386"},
	"WI": []string{"387-399"},
	"KY": []string{"400-407"},
	"TN": []string{"408-415"},
	"AL": []string{"416-424"},
	"MS": []string{"425-428"},
	"AR": []string{"429-432"},
	"LA": []string{"433-439"},
	"OK": []string{"440-448"},
	"TX": []string{"449-467"},
	"MN": []string{"468-477"},
	"IO": []string{"478-485"},
	"MO": []string{"486-500"},
	"ND": []string{"501-502"},
	"SD": []string{"503-504"},
	"NE": []string{"505-508"},
	"KS": []string{"509-515"},
	"MT": []string{"516-517"},
	"ID": []string{"518-519"},
	"WY": []string{"520"},
	"CO": []string{"521-524"},
	"NM": []string{"525,585"},
	"AZ": []string{"526-527"},
	"UT": []string{"528-529"},
	"NV": []string{"530,680"},
	"WA": []string{"531-539"},
	"OR": []string{"540-544"},
	"CA": []string{"545-573"},
	"AK": []string{"574"},
	"HI": []string{"575-576"},
	"DC": []string{"577-579"},
	"PR": []string{"580-584"},
	"NOT_ISSUED": []string{
		"750-772",
		"587-665",
		"667-679",
		"681-690",
		"691-699",
		"237-246",
	},
}
