"use strict";

// ********
// * AJAX *
// ********
/**
 * This function submits an ajax request of content-type json to the 'api' route
 * @param {Object} data
 * A Form Element which has been serialized by serializeFormObj function
 */
function submitAjax(data) {
    var url = '/api' + window.location.pathname;
    var request = new XMLHttpRequest();
    request.open('POST', url, true);

    request.onload = function () {
        if (this.status >= 200 && this.status <= 404) {
            var resp = JSON.parse(this.response);

            if (resp.error) {
                displayCalculationCard(resp.error);
            } else {
                displayCalculationCard(resp.content);
            }
        }
        else if (this.status === 429) {
            displayCalculationCard("<p>You sent too many requests to the server, come back tomorrow.</p>")
        } else {
            displayCalculationCard('<p>The server has encountered a problem.</p>');
        }
    };

    request.onerror = function () {
        console.log(request.error);
        displayCalculationCard('<p>There was a connection issue. Check your internet connection or the sever might be down.</p>')
    };

    request.setRequestHeader("X-CSRFToken", data.csrfmiddlewaretoken);
    request.setRequestHeader('Content-Type', 'application/json');
    console.log(JSON.stringify(data), url);
    request.send(JSON.stringify(data));
}

// ***********
// * OUTPUT *
// ***********
var calculationCard = document.getElementById('calculation-card');

/**
 * This function applies the given param to the element #calculation-card and fades in the element.
 * @param {String} output
 */
function displayCalculationCard(output) {
    calculationCard.innerHTML = output;
    calculationCard.classList.add("fade-in");
}

// *****************
// * SERIALIZATION *
// *****************
/**
 * This function serializes a Form Element Object into a general Javascript Object
 * @param {Object} form
 * A DOM Form Element Object
 * @returns {Object}
 * A general Javascript Object
 */
function serializeFormObj(form) {
    var elems = form.elements;
    var obj = {};

    for (var i = 0; i < elems.length; i += 1) {
        var element = elems[i];
        var type = element.type;
        var name = element.name;
        var value = element.value;

        switch (type) {
            case 'hidden':
            case 'text':
                obj[name] = value;
                break;
            default:
                break;
        }
    }
    return obj;
}

// **************
// * VALIDATION *
// **************
/**
 * This function checks if param binary is a binary number
 * @param {String} binary
 * binary String, must not contain spaces
 * @returns {boolean}
 */
function binaryValidation(binary) {
    var isValidBinary = true;

    if (binary.length > 64) {
        isValidBinary = false;
    } else {
        for (var i = 0; i < binary.length; i++) {
            if (binary[i] !== '0' && binary[i] !== '1') {
                isValidBinary = false;
                break;
            }
        }
    }
    return isValidBinary;
}

/**
 * This function checks if param decimal is a decimal number, int and float are accepted
 * @param {String} decimal
 * decimal String, must not contain spaces
 * @returns {boolean}
 */
function decimalValidation(decimal) {
    var isValidDecimal = true;

    if (isNaN(Number(decimal)) || decimal > 999999999999) {
        isValidDecimal = false;
    }
    return isValidDecimal;
}

/**
 * This function checks if param hexadecimal is a hexadecimal number
 * @param {String} hexadecimal
 * hexadecimal String, must not contain spaces and must be uppercase
 * @returns {boolean}
 */
function hexadecimalValidation(hexadecimal) {
    var isValidHexadecimal = true;

    if (!/^[A-F0-9]+$/.test(hexadecimal) || hexadecimal.length > 64) {
        isValidHexadecimal = false;
    }
    return isValidHexadecimal;
}