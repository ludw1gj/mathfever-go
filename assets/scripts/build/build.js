define("Handler/SerializeHandler", ["require", "exports"], function (require, exports) {
    "use strict";
    Object.defineProperty(exports, "__esModule", { value: true });

    function serializeForm(form) {
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
    }

    exports.serializeForm = serializeForm;
});
define("Handler/CalculationCardHandler", ["require", "exports"], function (require, exports) {
    "use strict";
    Object.defineProperty(exports, "__esModule", { value: true });

    function displayCalculationCard(content) {
        var calculationCard = document.getElementById("calculation-card");
        calculationCard.innerHTML = "<p>" + content + "</p>";
        calculationCard.classList.add("fade-in");
    }

    exports.displayCalculationCard = displayCalculationCard;
});
define("Handler/AjaxHandler", ["require", "exports", "Handler/CalculationCardHandler"], function (require, exports, CalculationCardHandler_1) {
    "use strict";
    Object.defineProperty(exports, "__esModule", { value: true });

    function submitAjax(data) {
        var spinner = document.getElementById("loading-spinner-container");
        var url = "/api" + window.location.pathname;
        var request = new XMLHttpRequest();
        request.open("POST", url, true);
        request.onload = function () {
            spinner.style.opacity = "0";
            if (this.status >= 200 && this.status <= 404) {
                var resp = JSON.parse(this.response);
                if (resp.error) {
                    CalculationCardHandler_1.displayCalculationCard(resp.error);
                }
                else {
                    CalculationCardHandler_1.displayCalculationCard(resp.content);
                }
            }
            else if (this.status === 429) {
                CalculationCardHandler_1.displayCalculationCard("You sent too many requests to the server, come back tomorrow.");
            }
            else {
                CalculationCardHandler_1.displayCalculationCard("The server has encountered a problem.");
            }
        };
        request.onerror = function () {
            spinner.style.opacity = "0";
            CalculationCardHandler_1.displayCalculationCard("There was a connection issue. Check your internet connection or the sever might be down.");
        };
        request.setRequestHeader("Content-Type", "application/json");
        request.send(JSON.stringify(data));
    }

    exports.submitAjax = submitAjax;
});
define("Handler/ValidationHandler", ["require", "exports", "Handler/CalculationCardHandler", "Handler/AjaxHandler"], function (require, exports, CalculationCardHandler_2, AjaxHandler_1) {
    "use strict";
    Object.defineProperty(exports, "__esModule", { value: true });

    function validateForm(form) {
        for (var key in form) {
            form[key] = form[key].replace(/\s/g, "");
            if (form[key] === "") {
                CalculationCardHandler_2.displayCalculationCard("Please enter a valid input.");
                return;
            }
            switch (key) {
                case "binary":
                    if (!isValidBinary(form[key])) {
                        CalculationCardHandler_2.displayCalculationCard("Please enter a valid binary number, and is not over 64 characters in length.");
                        return;
                    }
                    break;
                case "decimal":
                    if (!isValidDecimal(form[key])) {
                        CalculationCardHandler_2.displayCalculationCard("Please enter a valid decimal number, and is not over 999,999,999,999.");
                        return;
                    }
                    else {
                        form[key] = parseInt(form[key]);
                    }
                    break;
                case "hexadecimal":
                    form[key] = form[key].toUpperCase();
                    if (!isValidHexadecimal(form[key])) {
                        CalculationCardHandler_2.displayCalculationCard("Please enter a valid hexadecimal number, and is not over 64 characters in length.");
                        return;
                    }
                    break;
                default:
                    if (!isValidDecimal(form[key])) {
                        CalculationCardHandler_2.displayCalculationCard("Please enter a valid input, and is not over 999,999,999,999.");
                        return;
                    }
                    else {
                        form[key] = parseInt(form[key]);
                    }
            }
        }
        var spinner = document.getElementById("loading-spinner-container");
        spinner.style.opacity = "100";
        AjaxHandler_1.submitAjax(form);
    }

    exports.validateForm = validateForm;

    function isValidBinary(binary) {
        if (binary.length > 64) {
            return false;
        }
        for (var i = 0; i < binary.length; i++) {
            if (binary[i] !== "0" && binary[i] !== "1") {
                return false;
            }
        }
        return true;
    }

    function isValidDecimal(decimal) {
        return !isNaN(decimal) && decimal < 999999999999;
    }

    function isValidHexadecimal(hexadecimal) {
        return /^[A-F0-9]+$/.test(hexadecimal) && hexadecimal.length < 64;
    }
});
define("HandlerController", ["require", "exports", "Handler/SerializeHandler", "Handler/ValidationHandler"], function (require, exports, SerializeHandler_1, ValidationHandler_1) {
    "use strict";
    Object.defineProperty(exports, "__esModule", { value: true });

    function run() {
        if (!document.getElementById("calculation-form")) {
            return;
        }
        var calculationForm = document.getElementById("calculation-form");
        calculationForm.addEventListener("submit", function (event) {
            event.preventDefault();
            var form = SerializeHandler_1.serializeForm(calculationForm);
            ValidationHandler_1.validateForm(form);
        });
    }

    run();
});
