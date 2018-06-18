define("util/serialize", ["require", "exports"], function (require, exports) {
    "use strict";
    Object.defineProperty(exports, "__esModule", { value: true });
    exports.serializeForm = function (form) {
        var elements = form.elements;
        var obj = {};
        for (var i = 0; i < elements.length; i += 1) {
            var element = elements[i];
            var type = element.type;
            var name_1 = element.name;
            var value = element.value;
            switch (type) {
                case "hidden":
                case "text":
                    obj[name_1] = value;
                    break;
                default:
                    break;
            }
        }
        return obj;
    };
});
define("util/display", ["require", "exports"], function (require, exports) {
    "use strict";
    Object.defineProperty(exports, "__esModule", { value: true });
    exports.displayCalculationCard = function (content) {
        var calculationCard = document.getElementById("calculation-card");
        calculationCard.innerHTML = "<p>" + content + "</p>";
        calculationCard.classList.add("fade-in");
    };
});
define("util/validate", ["require", "exports", "util/display"], function (require, exports, display_1) {
    "use strict";
    Object.defineProperty(exports, "__esModule", { value: true });
    var isValidBinary = function (binary) {
        if (binary.length > 64) {
            return false;
        }
        for (var i = 0; i < binary.length; i++) {
            if (binary[i] !== "0" && binary[i] !== "1") {
                return false;
            }
        }
        return true;
    };
    var isValidDecimal = function (decimal) {
        return !isNaN(decimal) && decimal < 999999999999;
    };
    var isValidHexadecimal = function (hexadecimal) {
        return /^[A-F0-9]+$/.test(hexadecimal) && hexadecimal.length < 64;
    };
    exports.validateForm = function (form) {
        for (var key in form) {
            form[key] = form[key].replace(/\s/g, "");
            if (form[key] === "") {
                display_1.displayCalculationCard("Please enter a valid input.");
                return [form, false];
            }
            switch (key) {
                case "binary":
                    if (!isValidBinary(form[key])) {
                        display_1.displayCalculationCard("Please enter a valid binary number, and is not over 64 characters in length.");
                        return [form, false];
                    }
                    break;
                case "decimal":
                    if (!isValidDecimal(form[key])) {
                        display_1.displayCalculationCard("Please enter a valid decimal number, and is not over 999,999,999,999.");
                        return [form, false];
                    }
                    else {
                        form[key] = parseInt(form[key]);
                    }
                    break;
                case "hexadecimal":
                    form[key] = form[key].toUpperCase();
                    if (!isValidHexadecimal(form[key])) {
                        display_1.displayCalculationCard("Please enter a valid hexadecimal number, and is not over 64 characters in length.");
                        return [form, false];
                    }
                    break;
                default:
                    if (!isValidDecimal(form[key])) {
                        display_1.displayCalculationCard("Please enter a valid input, and is not over 999,999,999,999.");
                        return [form, false];
                    }
                    var currentPageURL = window.location.pathname;
                    if (currentPageURL.indexOf("percentages") !== -1 ||
                        currentPageURL.indexOf("total-surface-area") !== -1) {
                        form[key] = parseFloat(form[key]);
                    }
                    else {
                        form[key] = parseInt(form[key]);
                    }
            }
        }
        var spinner = document.getElementById("loading-spinner-container");
        spinner.style.opacity = "100";
        return [form, true];
    };
});
define("util/ajax", ["require", "exports", "util/display"], function (require, exports, display_2) {
    "use strict";
    Object.defineProperty(exports, "__esModule", { value: true });
    exports.submitAjax = function (data) {
        var spinner = document.getElementById("loading-spinner-container");
        var currentURL = window.location.pathname;
        var index = currentURL.lastIndexOf("/");
        var calculationSlug = currentURL.substring(index + 1);
        var request = new XMLHttpRequest();
        request.open("POST", "/api/calculation?calculation=" + calculationSlug, true);
        request.onload = function () {
            spinner.style.opacity = "0";
            if (this.status >= 200 && this.status <= 404) {
                var resp = JSON.parse(this.response);
                if (resp.error) {
                    display_2.displayCalculationCard(resp.error);
                }
                else {
                    display_2.displayCalculationCard(resp.content);
                }
            }
            else if (this.status === 429) {
                display_2.displayCalculationCard("You sent too many requests to the server, come back tomorrow.");
            }
            else {
                display_2.displayCalculationCard("The server has encountered a problem.");
            }
        };
        request.onerror = function () {
            spinner.style.opacity = "0";
            display_2.displayCalculationCard("There was a connection issue. Check your internet connection or the sever might be down.");
        };
        request.setRequestHeader("Content-Type", "application/json");
        request.send(JSON.stringify(data));
    };
});
define("App", ["require", "exports", "util/serialize", "util/validate", "util/ajax"], function (require, exports, serialize_1, validate_1, ajax_1) {
    "use strict";
    Object.defineProperty(exports, "__esModule", { value: true });
    var run = function () {
        if (!document.getElementById("calculation-form")) {
            return;
        }
        var calculationForm = document.getElementById("calculation-form");
        calculationForm.addEventListener("submit", function (event) {
            event.preventDefault();
            var form = serialize_1.serializeForm(calculationForm);
            var _a = validate_1.validateForm(form), validatedForm = _a[0], isValid = _a[1];
            if (isValid) {
                ajax_1.submitAjax(validatedForm);
            }
        });
    };
    run();
});
