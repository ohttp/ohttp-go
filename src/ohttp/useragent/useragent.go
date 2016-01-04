// Copyright (c) 2015-2016 Kerem Güneş
//   <http://qeremy.com>
//
// GNU General Public License v3.0
//   <http://www.gnu.org/licenses/gpl-3.0.txt>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

// @package    ohttp
// @subpackage ohttp.useragent
// @author     Kerem Güneş <qeremy[at]gmail[dot]com>
package useragent

// @object ohttp.UserAgent
type UserAgent struct {
    name          string
    platform      string
    version       float32
    versionOrig   string
}

// Oh! HTTP ua details.
// @const string
const (
    OHTTP_NAME    = "Oh! HTTP"
    OHTTP_LINK    = "http://github.com/ohttp/ohttp-go"
    OHTTP_VERSION = "1.0.0"
)

// User agent strings.
// @const string
const (
    FIREFOX       = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:25.0) Gecko/20100101 Firefox/25.0"
    CHROME        = "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) AppleWebKit/525.19 "+
                    "(KHTML, like Gecko) Chrome/1.0.154.53 Safari/525.19"
    CHROMIUM      = "Mozilla/5.0 (X11; Linux i686) AppleWebKit/535.7 (KHTML, like Gecko) "+
                    "Ubuntu/11.10 Chromium/16.0.912.21 Chrome/16.0.912.21 Safari/535.7"
    IE            = "Mozilla/5.0 (IE 11.0; Windows NT 6.3; Trident/7.0; .NET4.0E; .NET4.0C; rv:11.0) like Gecko"

    ANDROID       = "Mozilla/5.0 (Linux; U; Android 2.3; en-us) AppleWebKit/999+ (KHTML, like Gecko) Safari/999.9"
    BLACKBERRY    = "Mozilla/5.0 (BlackBerry; U; BlackBerry 9900; en) AppleWebKit/534.11+ (KHTML, like Gecko) "+
                    "Version/7.1.0.346 Mobile Safari/534.11+"
    IPHONE        = "Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) "+
                    "Version/5.1 Mobile/9A334 Safari/7534.48.3"
    IPAD          = "Mozilla/5.0 (iPad; CPU OS 7_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 "+
                    "Mobile/9A334 Safari/7534.48.3"
)

// @todo
func Parse() {}
