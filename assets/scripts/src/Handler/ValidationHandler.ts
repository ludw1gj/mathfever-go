import {CalculationCardHandler} from "./CalculationCardHandler";
import {AjaxHandler} from "./AjaxHandler";

class ValidationHandler {

    public static validateForm(form: Object): void {
        let isValidInput = true;

        for (const key in form) {
            // removes all whitespace from the form data
            form[key] = form[key].replace(/\s/g, "");

            if (form[key] === "") {
                isValidInput = false;
                CalculationCardHandler.display("Please enter a valid input.");
                break;
            } else if (key === "binary") {
                if (!this.binaryValidation(form[key])) {
                    isValidInput = false;
                    CalculationCardHandler.display("Please enter a valid binary number, and is not over 64 characters in length.");
                    break;
                }
            } else if (key === "decimal") {
                if (!this.decimalValidation(form[key])) {
                    isValidInput = false;
                    CalculationCardHandler.display("Please enter a valid decimal number, and is not over 999,999,999,999.");
                    break;
                } else {
                    // data must be an integer string not a float string
                    form[key] = parseInt(form[key]);
                }
            } else if (key === "hexadecimal") {
                // hexadecimal to uppercase before it reaches the server
                form[key] = form[key].toUpperCase();
                if (!this.hexadecimalValidation(form[key])) {
                    isValidInput = false;
                    CalculationCardHandler.display("Please enter a valid hexadecimal number, and is not over 64 characters in length.");
                    break;
                }
            } else if (!this.decimalValidation(form[key])) {
                isValidInput = false;
                CalculationCardHandler.display("Please enter a valid input, and is not over 999,999,999,999.");
                break;
            } else {
                form[key] = parseInt(form[key]);
            }
        }
        // if the data is valid, send the Ajax request
        if (isValidInput) {
            let spinner = document.getElementById("loading-spinner-container") as HTMLElement ;
            spinner.style.opacity = "100";

            AjaxHandler.submit(form);
        }
    }

    private static binaryValidation(binary: string): boolean {
        let isValidBinary = true;

        if (binary.length > 64) {
            isValidBinary = false;
        } else {
            for (let i = 0; i < binary.length; i++) {
                if (binary[i] !== "0" && binary[i] !== "1") {
                    isValidBinary = false;
                    break;
                }
            }
        }
        return isValidBinary;
    }

    private static decimalValidation(decimal: number): boolean {
        let isValidDecimal = true;

        if (typeof decimal === "number" || decimal > 999999999999) {
            isValidDecimal = false;
        }
        return isValidDecimal;
    }

    private static hexadecimalValidation(hexadecimal: string): boolean {
        let isValidHexadecimal = true;

        if (!/^[A-F0-9]+$/.test(hexadecimal) || hexadecimal.length > 64) {
            isValidHexadecimal = false;
        }
        return isValidHexadecimal;
    }

}

export {ValidationHandler}