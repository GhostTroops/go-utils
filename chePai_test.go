package go_utils

import (
	"fmt"
	"regexp"
	"testing"
)

func TestWhereCar(t *testing.T) {
	//log.Println(WhereCar("WJ210034"))
	//log.Println(WhereCar("VV10034"))
	//log.Println(WhereCar("川G0034"))
	//log.Println(WhereCar("川A0034"))
	//log.Println(WhereCar("使2170034"))
	//log.Println(WhereCar("台C70034"))
	//log.Println(WhereCar("民航034"))
	//fmt.Println(Fanyi4YoudaoPars(`こんにちは`, "ja"))
	//fmt.Println(Fanyi4YoudaoPars(`안녕하세요`, "ko"))

	var s11 = `HTTP/1.1 500 Internal Server Error
Host: 127.0.0.1:8888
Date: Wed, 07 Feb 2024 05:10:22 GMT
Connection: close
X-Powered-By: PHP/7.3.25
Cache-Control: no-cache, private
Date: Wed, 07 Feb 2024 05:10:22 GMT
Content-Type: application/json

{
    "message": "file_get_contents(DOESNOTEXIST): failed to open stream: No such file or directory",
    "exception": "ErrorException",
    "file": "/src/vendor/facade/ignition/src/Solutions/MakeViewVariableOptionalSolution.php",
    "line": 75,
    "trace": [
        {
            "function": "handleError",
            "class": "Illuminate\\Foundation\\Bootstrap\\HandleExceptions",
            "type": "->"
        },
        {
            "file": "/src/vendor/facade/ignition/src/Solutions/MakeViewVariableOptionalSolution.php",
            "line": 75,
            "function": "file_get_contents"
        },
        {
            "file": "/src/vendor/facade/ignition/src/Solutions/MakeViewVariableOptionalSolution.php",
            "line": 67,
            "function": "makeOptional",
            "class": "Facade\\Ignition\\Solutions\\MakeViewVariableOptionalSolution",
            "type": "->"
        },
        {
            "file": "/src/vendor/facade/ignition/src/Http/Controllers/ExecuteSolutionController.php",
            "line": 19,
            "function": "run",
            "class": "Facade\\Ignition\\Solutions\\MakeViewVariableOptionalSolution",
            "type": "->"
        },
        {
            "file": "/src/vendor/laravel/framework/src/Illuminate/Routing/ControllerDispatcher.php",
            "line": 48,
            "function": "__invoke",
            "class": "Facade\\Ignition\\Http\\Controllers\\ExecuteSolutionController",
            "type": "->"
        },
        {
            "file": "/src/vendor/laravel/framework/src/Illuminate/Routing/Route.php",
            "line": 262,
            "function": "dispatch",
            "class": "Illuminate\\Routing\\ControllerDispatcher",
            "type": "->"
        },
        {
            "file": "/src/vendor/laravel/framework/src/Illuminate/Routing/Route.php",
            "line": 205,
            "function": "runController",
            "class": "Illuminate\\Routing\\Route",
            "type": "->"
        },
        {
            "file": "/src/vendor/laravel/framework/src/Illuminate/Routing/Router.php",
            "line": 721,
            "function": "run",
            "class": "Illuminate\\Routing\\Route",
            "type": "->"
        },
        {
            "file": "/src/vendor/laravel/framework/src/Illuminate/Pipeline/Pipeline.php",
            "line": 128,
            "function": "Illuminate\\Routing\\{closure}",
            "class": "Illuminate\\Routing\\Router",
            "type": "->"
        },
        {
            "file": "/src/vendor/facade/ignition/src/Http/Middleware/IgnitionConfigValueEnabled.php",
            "line": 25,
            "function": "Illuminate\\Pipeline\\{closure}",
            "class": "Illuminate\\Pipeline\\Pipeline",
            "type": "->"
        },
        {
            "file": "/src/vendor/laravel/framework/src/Illuminate/Pipeline/Pipeline.php",
            "line": 167,
            "function": "handle",
            "class": "Facade\\Ignition\\Http\\Middleware\\IgnitionConfigValueEnabled",
            "type": "->"
        },
        {
            "file": "/src/vendor/facade/ignition/src/Http/Middleware/IgnitionEnabled.php",
            "line": 23,
            "function": "Illuminate\\Pipeline\\{closure}",
            "class": "Illuminate\\Pipeline\\Pipeline",
            "type": "->"
        },
        {
            "file": "/src/vendor/laravel/framework/src/Illuminate/Pipeline/Pipeline.php",
            "line": 167,
            "function": "handle",
            "class": "Facade\\Ignition\\Http\\Middleware\\IgnitionEnabled",
            "type": "->"
        },
        {
            "file": "/src/vendor/laravel/framework/src/Illuminate/Pipeline/Pipeline.php",
            "line": 103,
            "function": "Illuminate\\Pipeline\\{closure}",
            "class": "Illuminate\\Pipeline\\Pipeline",
            "type": "->"
        },
        {
            "file": "/src/vendor/laravel/framework/src/Illuminate/Routing/Router.php",
            "line": 723,
            "function": "then",
            "class": "Illuminate\\Pipeline\\Pipeline",
            "type": "->"
        },
        {
            "file": "/src/vendor/laravel/framework/src/Illuminate/Routing/Router.php",
            "line": 698,
            "function": "runRouteWithinStack",
            "class": "Illuminate\\Routing\\Router",
            "type": "->"
        },
        {
            "file": "/src/vendor/laravel/framework/src/Illuminate/Routing/Router.php",
            "line": 662,
            "function": "runRoute",
            "class": "Illuminate\\Routing\\Router",
            "type": "->"
        },
        {
            "file": "/src/vendor/laravel/framework/src/Illuminate/Routing/Router.php",
            "line": 651,
            "function": "dispatchToRoute",
            "class": "Illuminate\\Routing\\Router",
            "type": "->"
        },
        {
            "file": "/src/vendor/laravel/framework/src/Illuminate/Foundation/Http/Kernel.php",
            "line": 167,
            "function": "dispatch",
            "class": "Illuminate\\Routing\\Router",
            "type": "->"
        },
        {
            "file": "/src/vendor/laravel/framework/src/Illuminate/Pipeline/Pipeline.php",
            "line": 128,
            "function": "Illuminate\\Foundation\\Http\\{closure}",
            "class": "Illuminate\\Foundation\\Http\\Kernel",
            "type": "->"
        },
        {
            "file": "/src/vendor/laravel/framework/src/Illuminate/Foundation/Http/Middleware/TransformsRequest.php",
            "line": 21,
            "function": "Illuminate\\Pipeline\\{closure}",
            "class": "Illuminate\\Pipeline\\Pipeline",
            "type": "->"
        },
        {
            "file": "/src/vendor/laravel/framework/src/Illuminate/Foundation/Http/Middleware/ConvertEmptyStringsToNull.php",
            "line": 31,
            "function": "handle",
            "class": "Illuminate\\Foundation\\Http\\Middleware\\TransformsRequest",
            "type": "->"
        },
        {
            "file": "/src/vendor/laravel/framework/src/Illuminate/Pipeline/Pipeline.php",
            "line": 167,
            "function": "handle",
            "class": "Illuminate\\Foundation\\Http\\Middleware\\ConvertEmptyStringsToNull",
            "type": "->"
        },
        {
            "file": "/src/vendor/laravel/framework/src/Illuminate/Foundation/Http/Middleware/TransformsRequest.php",
            "line": 21,
            "function": "Illuminate\\Pipeline\\{closure}",
            "class": "Illuminate\\Pipeline\\Pipeline",
            "type": "->"
        },
        {
            "file": "/src/vendor/laravel/framework/src/Illuminate/Foundation/Http/Middleware/TrimStrings.php",
            "line": 40,
            "function": "handle",
            "class": "Illuminate\\Foundation\\Http\\Middleware\\TransformsRequest",
            "type": "->"
        },
        {
            "file": "/src/vendor/laravel/framework/src/Illuminate/Pipeline/Pipeline.php",
            "line": 167,
            "function": "handle",
            "class": "Illuminate\\Foundation\\Http\\Middleware\\TrimStrings",
            "type": "->"
        },
        {
            "file": "/src/vendor/laravel/framework/src/Illuminate/Foundation/Http/Middleware/ValidatePostSize.php",
            "line": 27,
            "function": "Illuminate\\Pipeline\\{closure}",
            "class": "Illuminate\\Pipeline\\Pipeline",
            "type": "->"
        },
        {
            "file": "/src/vendor/laravel/framework/src/Illuminate/Pipeline/Pipeline.php",
            "line": 167,
            "function": "handle",
            "class": "Illuminate\\Foundation\\Http\\Middleware\\ValidatePostSize",
            "type": "->"
        },
        {
            "file": "/src/vendor/laravel/framework/src/Illuminate/Foundation/Http/Middleware/PreventRequestsDuringMaintenance.php",
            "line": 86,
            "function": "Illuminate\\Pipeline\\{closure}",
            "class": "Illuminate\\Pipeline\\Pipeline",
            "type": "->"
        },
        {
            "file": "/src/vendor/laravel/framework/src/Illuminate/Pipeline/Pipeline.php",
            "line": 167,
            "function": "handle",
            "class": "Illuminate\\Foundation\\Http\\Middleware\\PreventRequestsDuringMaintenance",
            "type": "->"
        },
        {
            "file": "/src/vendor/fruitcake/laravel-cors/src/HandleCors.php",
            "line": 38,
            "function": "Illuminate\\Pipeline\\{closure}",
            "class": "Illuminate\\Pipeline\\Pipeline",
            "type": "->"
        },
        {
            "file": "/src/vendor/laravel/framework/src/Illuminate/Pipeline/Pipeline.php",
            "line": 167,
            "function": "handle",
            "class": "Fruitcake\\Cors\\HandleCors",
            "type": "->"
        },
        {
            "file": "/src/vendor/fideloper/proxy/src/TrustProxies.php",
            "line": 57,
            "function": "Illuminate\\Pipeline\\{closure}",
            "class": "Illuminate\\Pipeline\\Pipeline",
            "type": "->"
        },
        {
            "file": "/src/vendor/laravel/framework/src/Illuminate/Pipeline/Pipeline.php",
            "line": 167,
            "function": "handle",
            "class": "Fideloper\\Proxy\\TrustProxies",
            "type": "->"
        },
        {
            "file": "/src/vendor/laravel/framework/src/Illuminate/Pipeline/Pipeline.php",
            "line": 103,
            "function": "Illuminate\\Pipeline\\{closure}",
            "class": "Illuminate\\Pipeline\\Pipeline",
            "type": "->"
        },
        {
            "file": "/src/vendor/laravel/framework/src/Illuminate/Foundation/Http/Kernel.php",
            "line": 142,
            "function": "then",
            "class": "Illuminate\\Pipeline\\Pipeline",
            "type": "->"
        },
        {
            "file": "/src/vendor/laravel/framework/src/Illuminate/Foundation/Http/Kernel.php",
            "line": 111,
            "function": "sendRequestThroughRouter",
            "class": "Illuminate\\Foundation\\Http\\Kernel",
            "type": "->"
        },
        {
            "file": "/src/public/index.php",
            "line": 52,
            "function": "handle",
            "class": "Illuminate\\Foundation\\Http\\Kernel",
            "type": "->"
        },
        {
            "file": "/src/server.php",
            "line": 21,
            "function": "require_once"
        }
    ]
}`
	var reLogPath = regexp.MustCompile(`"file":\s*"(\/[^"]+?)\/vendor\/[^"]+?"`)
	matches := reLogPath.FindStringSubmatch(s11)
	if len(matches) > 1 {
		fmt.Println(matches[1])
	}
}
