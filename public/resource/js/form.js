'use strict';

if (document.getElementById('calculation-form') !== null) {
    var calculationForm = document.getElementById('calculation-form');

    calculationForm.addEventListener('submit', function (event) {
        event.preventDefault();

        var formData = serializeFormObj(calculationForm);
        var isValidInput = true;

        // validate the data
        for (var key in formData) {
            // removes all whitespace from the form data
            formData[key] = formData[key].replace(/\s/g, '');

            if (formData[key] === '') {
                isValidInput = false;
                displayCalculationCard('<p>Please enter a valid input.</p>');
                break;
            }
            else if (key === 'binary') {
                if (!binaryValidation(formData[key])) {
                    isValidInput = false;
                    displayCalculationCard('<p>Please enter a valid binary number, and is not over 64 characters in length.</p>');
                    break;
                }
            } else if (key === 'decimal') {
                if (!decimalValidation(formData[key])) {
                    isValidInput = false;
                    displayCalculationCard('<p>Please enter a valid decimal number, and is not over 999,999,999,999.</p>');
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
                    displayCalculationCard('<p>Please enter a valid hexadecimal number, and is not over 64 characters in length.</p>');
                    break;
                }
            } else if (!decimalValidation(formData[key])) {
                isValidInput = false;
                displayCalculationCard('<p>Please enter a valid input, and is not over 999,999,999,999.</p>');
                break;
            } else {
                formData[key] = parseInt(formData[key]);
            }
        }
        // if the data is valid, send the Ajax request
        if (isValidInput) {
            submitAjax(formData);
        }
    });
}
