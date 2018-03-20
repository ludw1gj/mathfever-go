import {displayCalculationCard} from "./display";

const isValidBinary = (binary: string): boolean => {
    if (binary.length > 64) {
        return false;
    }

    for (let i = 0; i < binary.length; i++) {
        if (binary[i] !== "0" && binary[i] !== "1") {
            return false;
        }
    }
    return true;
};

const isValidDecimal = (decimal: number): boolean => {
    return !isNaN(decimal) && decimal < 999999999999;
};

const isValidHexadecimal = (hexadecimal: string): boolean => {
    return /^[A-F0-9]+$/.test(hexadecimal) && hexadecimal.length < 64;
};

export const validateForm = (form: Object): Array<Object> => {
    for (const key in form) {
        form[key] = form[key].replace(/\s/g, ""); // removes all whitespace from the form data
        if (form[key] === "") {
            displayCalculationCard("Please enter a valid input.");
            return [form, false];
        }

        switch (key) {
            case "binary":
                if (!isValidBinary(form[key])) {
                    displayCalculationCard("Please enter a valid binary number, and is not over 64 characters in length.");
                    return [form, false];
                }
                break;
            case "decimal":
                if (!isValidDecimal(form[key])) {
                    displayCalculationCard("Please enter a valid decimal number, and is not over 999,999,999,999.");
                    return [form, false];
                } else {
                    form[key] = parseInt(form[key]); // data must be an integer string not a float string
                }
                break;
            case "hexadecimal":
                form[key] = form[key].toUpperCase(); // hexadecimal to uppercase before it reaches the server
                if (!isValidHexadecimal(form[key])) {
                    displayCalculationCard("Please enter a valid hexadecimal number, and is not over 64 characters in length.");
                    return [form, false];
                }
                break;
            default:
                if (!isValidDecimal(form[key])) {
                    displayCalculationCard("Please enter a valid input, and is not over 999,999,999,999.");
                    return [form, false];
                }
                // percentages and total-surface-area category require a float input
                const currentPageURL = window.location.pathname;
                if (currentPageURL.indexOf("percentages") !== -1 || currentPageURL.indexOf("total-surface-area") !== -1) {
                    form[key] = parseFloat(form[key]);
                } else {
                    form[key] = parseInt(form[key]); // input is for a int route such as numbers category

                }
        }
    }
    // the function has not returned, therefore data is valid
    const spinner = document.getElementById("loading-spinner-container") as HTMLElement;
    spinner.style.opacity = "100";
    return [form, true]
};
