'use strict';

// Validate the data in form
if (document.getElementById('calculation-form')) {
    var calculationForm = document.getElementById('calculation-form');

    calculationForm.addEventListener('submit', function (event) {
        event.preventDefault();

        var formData = serializeFormObj(calculationForm);
        validateForm(formData);
    });
}

/**
 * This function serialises a Form Element Object into a general Javascript Object
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

/**
 * This function validates form input
 * @param {Object} formData
 * A general Javascript Object
 */
function validateForm(formData) {
    var isValidInput = true;
    for (var key in formData) {
        // removes all whitespace from the form data
        formData[key] = formData[key].replace(/\s/g, '');

        if (formData[key] === '') {
            isValidInput = false;
            displayCalcCard('Please enter a valid input.');
            break;
        } else if (key === 'binary') {
            if (!binaryValidation(formData[key])) {
                isValidInput = false;
                displayCalcCard('Please enter a valid binary number, and is not over 64 characters in length.');
                break;
            }
        } else if (key === 'decimal') {
            if (!decimalValidation(formData[key])) {
                isValidInput = false;
                displayCalcCard('Please enter a valid decimal number, and is not over 999,999,999,999.');
                break;
            } else {
                // data must be an integer string not a float string
                formData[key] = parseInt(formData[key]);
            }
        } else if (key === 'hexadecimal') {
            // hexadecimal to uppercase before it reaches the server
            formData[key] = formData[key].toUpperCase();
            if (!hexadecimalValidation(formData[key])) {
                isValidInput = false;
                displayCalcCard('Please enter a valid hexadecimal number, and is not over 64 characters in length.');
                break;
            }
        } else if (!decimalValidation(formData[key])) {
            isValidInput = false;
            displayCalcCard('Please enter a valid input, and is not over 999,999,999,999.');
            break;
        } else {
            formData[key] = parseInt(formData[key]);
        }
    }
    // if the data is valid, send the Ajax request
    if (isValidInput) {
        document.getElementById("loading-spinner-container").style.opacity = 100;
        submitAjax(formData);
    }
}

/**
 * This function applies the given param to the element #calculation-card and fades in the element.
 * @param {String} content
 */
function displayCalcCard(content) {
    var calculationCard = document.getElementById('calculation-card');
    calculationCard.innerHTML = '<p>' + content + '</p>';
    calculationCard.classList.add('fade-in');
}

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
 * @param {Number} decimal
 * decimal String, must not contain spaces
 * @returns {boolean}
 */
function decimalValidation(decimal) {
    var isValidDecimal = true;

    if (isNaN(decimal) || decimal > 999999999999) {
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
        document.getElementById("loading-spinner-container").style.opacity = 0;

        if (this.status >= 200 && this.status <= 404) {
            var resp = JSON.parse(this.response);

            if (resp.error) {
                displayCalcCard(resp.error);
            } else {
                displayCalcCard(resp.content);
            }
        }
        else if (this.status === 429) {
            displayCalcCard('You sent too many requests to the server, come back tomorrow.')
        } else {
            displayCalcCard('The server has encountered a problem.');
        }
    };
    request.onerror = function () {
        displayCalcCard('There was a connection issue. Check your internet connection or the sever might be down.')
    };
    request.setRequestHeader('Content-Type', 'application/json');
    request.send(JSON.stringify(data));
}
