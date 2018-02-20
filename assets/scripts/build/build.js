define("Handler/SerializeHandler", ["require", "exports"], function (require, exports) {
    "use strict";
    Object.defineProperty(exports, "__esModule", { value: true });
    var SerializeHandler = (function () {
        function SerializeHandler() {
        }
        SerializeHandler.serializeForm = function (form) {
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
        return SerializeHandler;
    }());
    exports.SerializeHandler = SerializeHandler;
});
define("Handler/CalculationCardHandler", ["require", "exports"], function (require, exports) {
    "use strict";
    Object.defineProperty(exports, "__esModule", { value: true });
    var CalculationCardHandler = (function () {
        function CalculationCardHandler() {
        }
        CalculationCardHandler.display = function (content) {
            var calculationCard = document.getElementById("calculation-card");
            calculationCard.innerHTML = "<p>" + content + "</p>";
            calculationCard.classList.add("fade-in");
        };
        return CalculationCardHandler;
    }());
    exports.CalculationCardHandler = CalculationCardHandler;
});
define("Handler/AjaxHandler", ["require", "exports", "Handler/CalculationCardHandler"], function (require, exports, CalculationCardHandler_1) {
    "use strict";
    Object.defineProperty(exports, "__esModule", { value: true });
    var AjaxHandler = (function () {
        function AjaxHandler() {
        }
        AjaxHandler.submit = function (data) {
            var spinner = document.getElementById("loading-spinner-container");
            var url = "/api" + window.location.pathname;
            var request = new XMLHttpRequest();
            request.open("POST", url, true);
            request.onload = function () {
                spinner.style.opacity = "0";
                if (this.status >= 200 && this.status <= 404) {
                    var resp = JSON.parse(this.response);
                    if (resp.error) {
                        CalculationCardHandler_1.CalculationCardHandler.display(resp.error);
                    }
                    else {
                        CalculationCardHandler_1.CalculationCardHandler.display(resp.content);
                    }
                }
                else if (this.status === 429) {
                    CalculationCardHandler_1.CalculationCardHandler.display("You sent too many requests to the server, come back tomorrow.");
                }
                else {
                    CalculationCardHandler_1.CalculationCardHandler.display("The server has encountered a problem.");
                }
            };
            request.onerror = function () {
                spinner.style.opacity = "0";
                CalculationCardHandler_1.CalculationCardHandler.display("There was a connection issue. Check your internet connection or the sever might be down.");
            };
            request.setRequestHeader("Content-Type", "application/json");
            request.send(JSON.stringify(data));
        };
        return AjaxHandler;
    }());
    exports.AjaxHandler = AjaxHandler;
});
define("Handler/ValidationHandler", ["require", "exports", "Handler/CalculationCardHandler", "Handler/AjaxHandler"], function (require, exports, CalculationCardHandler_2, AjaxHandler_1) {
    "use strict";
    Object.defineProperty(exports, "__esModule", { value: true });
    var ValidationHandler = (function () {
        function ValidationHandler() {
        }
        ValidationHandler.validateForm = function (form) {
            var isValidInput = true;
            for (var key in form) {
                form[key] = form[key].replace(/\s/g, "");
                if (form[key] === "") {
                    isValidInput = false;
                    CalculationCardHandler_2.CalculationCardHandler.display("Please enter a valid input.");
                    break;
                }
                else if (key === "binary") {
                    if (!this.binaryValidation(form[key])) {
                        isValidInput = false;
                        CalculationCardHandler_2.CalculationCardHandler.display("Please enter a valid binary number, and is not over 64 characters in length.");
                        break;
                    }
                }
                else if (key === "decimal") {
                    if (!this.decimalValidation(form[key])) {
                        isValidInput = false;
                        CalculationCardHandler_2.CalculationCardHandler.display("Please enter a valid decimal number, and is not over 999,999,999,999.");
                        break;
                    }
                    else {
                        form[key] = parseInt(form[key]);
                    }
                }
                else if (key === "hexadecimal") {
                    form[key] = form[key].toUpperCase();
                    if (!this.hexadecimalValidation(form[key])) {
                        isValidInput = false;
                        CalculationCardHandler_2.CalculationCardHandler.display("Please enter a valid hexadecimal number, and is not over 64 characters in length.");
                        break;
                    }
                }
                else if (!this.decimalValidation(form[key])) {
                    isValidInput = false;
                    CalculationCardHandler_2.CalculationCardHandler.display("Please enter a valid input, and is not over 999,999,999,999.");
                    break;
                }
                else {
                    form[key] = parseInt(form[key]);
                }
            }
            if (isValidInput) {
                var spinner = document.getElementById("loading-spinner-container");
                spinner.style.opacity = "100";
                AjaxHandler_1.AjaxHandler.submit(form);
            }
        };
        ValidationHandler.binaryValidation = function (binary) {
            var isValidBinary = true;
            if (binary.length > 64) {
                isValidBinary = false;
            }
            else {
                for (var i = 0; i < binary.length; i++) {
                    if (binary[i] !== "0" && binary[i] !== "1") {
                        isValidBinary = false;
                        break;
                    }
                }
            }
            return isValidBinary;
        };
        ValidationHandler.decimalValidation = function (decimal) {
            var isValidDecimal = true;
            if (typeof decimal === "number" || decimal > 999999999999) {
                isValidDecimal = false;
            }
            return isValidDecimal;
        };
        ValidationHandler.hexadecimalValidation = function (hexadecimal) {
            var isValidHexadecimal = true;
            if (!/^[A-F0-9]+$/.test(hexadecimal) || hexadecimal.length > 64) {
                isValidHexadecimal = false;
            }
            return isValidHexadecimal;
        };
        return ValidationHandler;
    }());
    exports.ValidationHandler = ValidationHandler;
});
define("HandlerController", ["require", "exports", "Handler/SerializeHandler", "Handler/ValidationHandler"], function (require, exports, SerializeHandler_1, ValidationHandler_1) {
    "use strict";
    Object.defineProperty(exports, "__esModule", { value: true });
    var HandlerController = (function () {
        function HandlerController() {
        }
        HandlerController.run = function () {
            if (document.getElementById("calculation-form")) {
                var calculationForm_1 = document.getElementById("calculation-form");
                calculationForm_1.addEventListener("submit", function (event) {
                    event.preventDefault();
                    var form = SerializeHandler_1.SerializeHandler.serializeForm(calculationForm_1);
                    ValidationHandler_1.ValidationHandler.validateForm(form);
                });
            }
        };
        return HandlerController;
    }());
    HandlerController.run();
});
