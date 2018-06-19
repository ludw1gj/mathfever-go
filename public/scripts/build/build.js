System.register("bin/display", [], function (exports_1, context_1) {
    "use strict";
    var displayCalculationCard;
    var __moduleName = context_1 && context_1.id;
    return {
        setters: [],
        execute: function () {
            exports_1("displayCalculationCard", displayCalculationCard = function (content) {
                var calculationCard = document.getElementById("calculation-card");
                calculationCard.innerHTML = "<p>" + content + "</p>";
                calculationCard.classList.add("fade-in");
            });
        }
    };
});
System.register("bin/validate", [], function (exports_2, context_2) {
    "use strict";
    var isValidBinary, isValidDecimal, isValidHexadecimal;
    var __moduleName = context_2 && context_2.id;
    return {
        setters: [],
        execute: function () {
            exports_2("isValidBinary", isValidBinary = function (binary) {
                if (binary.length > 64) {
                    return false;
                }
                for (var i = 0; i < binary.length; i++) {
                    if (binary[i] !== "0" && binary[i] !== "1") {
                        return false;
                    }
                }
                return true;
            });
            exports_2("isValidDecimal", isValidDecimal = function (decimal) {
                return !isNaN(decimal) && decimal < 999999999999;
            });
            exports_2("isValidHexadecimal", isValidHexadecimal = function (hexadecimal) {
                return /^[A-F0-9]+$/.test(hexadecimal) && hexadecimal.length < 64;
            });
        }
    };
});
System.register("bin/form", ["bin/display", "bin/validate"], function (exports_3, context_3) {
    "use strict";
    var display_1, validate_1, createSerializedForm, createValidatedForm, validateAndParseFormKey;
    var __moduleName = context_3 && context_3.id;
    return {
        setters: [
            function (display_1_1) {
                display_1 = display_1_1;
            },
            function (validate_1_1) {
                validate_1 = validate_1_1;
            }
        ],
        execute: function () {
            exports_3("createSerializedForm", createSerializedForm = function (form) {
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
            });
            exports_3("createValidatedForm", createValidatedForm = function (form) {
                for (var key in form) {
                    form[key] = form[key].replace(/\s/g, "");
                    if (!validateAndParseFormKey(key, form)) {
                        return null;
                    }
                }
                return form;
            });
            validateAndParseFormKey = function (key, form) {
                if (form[key] === "") {
                    display_1.displayCalculationCard("Please enter a valid input.");
                    return false;
                }
                switch (key) {
                    case "binary":
                        if (!validate_1.isValidBinary(form[key])) {
                            display_1.displayCalculationCard("Please enter a valid binary number, and is not over 64 characters in length.");
                            return false;
                        }
                        break;
                    case "decimal":
                        if (!validate_1.isValidDecimal(form[key])) {
                            display_1.displayCalculationCard("Please enter a valid decimal number, and is not over 999,999,999,999.");
                            return false;
                        }
                        form[key] = parseInt(form[key]);
                        break;
                    case "hexadecimal":
                        form[key] = form[key].toUpperCase();
                        if (!validate_1.isValidHexadecimal(form[key])) {
                            display_1.displayCalculationCard("Please enter a valid hexadecimal number, and is not over 64 characters in length.");
                            return false;
                        }
                        break;
                    default:
                        if (!validate_1.isValidDecimal(form[key])) {
                            display_1.displayCalculationCard("Please enter a valid input, and is not over 999,999,999,999.");
                            return false;
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
                return true;
            };
        }
    };
});
System.register("bin/ajax", ["bin/display"], function (exports_4, context_4) {
    "use strict";
    var display_2, submitAjax;
    var __moduleName = context_4 && context_4.id;
    return {
        setters: [
            function (display_2_1) {
                display_2 = display_2_1;
            }
        ],
        execute: function () {
            exports_4("submitAjax", submitAjax = function (data) {
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
            });
        }
    };
});
System.register("index", ["bin/form", "bin/ajax"], function (exports_5, context_5) {
    "use strict";
    var form_1, ajax_1;
    var __moduleName = context_5 && context_5.id;
    return {
        setters: [
            function (form_1_1) {
                form_1 = form_1_1;
            },
            function (ajax_1_1) {
                ajax_1 = ajax_1_1;
            }
        ],
        execute: function () {
            (function () {
                if (!document.getElementById("calculation-form")) {
                    return;
                }
                var calculationForm = document.getElementById("calculation-form");
                calculationForm.addEventListener("submit", function (event) {
                    event.preventDefault();
                    var serializedForm = form_1.createSerializedForm(calculationForm);
                    var validatedForm = form_1.createValidatedForm(serializedForm);
                    if (validatedForm !== null) {
                        var spinner = document.getElementById("loading-spinner-container");
                        spinner.style.opacity = "100";
                        ajax_1.submitAjax(validatedForm);
                    }
                });
            })();
        }
    };
});
