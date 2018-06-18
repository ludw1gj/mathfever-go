import { displayCalculationCard } from "./display";
import { isValidBinary, isValidDecimal, isValidHexadecimal } from "./validate";

export const createSerializedForm = (form: HTMLFormElement) => {
  const elements = form.elements;
  const obj: { [k: string]: any } = {};

  for (let i = 0; i < elements.length; i += 1) {
    const element = elements[i] as HTMLInputElement;
    const type = element.type;
    const name = element.name;
    const value = element.value;

    switch (type) {
      case "hidden":
      case "text":
        obj[name] = value;
        break;
      default:
        break;
    }
  }
  return obj;
};

export const createValidatedForm = (form: object) => {
  for (const key in form) {
    // removes all whitespace from the form data
    form[key] = form[key].replace(/\s/g, "");

    // check if key's value is valid
    if (!validateAndParseFormKey(key, form)) {
      return null;
    }
  }
  // form is valid
  return form;
};

const validateAndParseFormKey = (key: string, form: object) => {
  if (form[key] === "") {
    displayCalculationCard("Please enter a valid input.");
    return false;
  }

  switch (key) {
    case "binary":
      if (!isValidBinary(form[key])) {
        displayCalculationCard(
          "Please enter a valid binary number, and is not over 64 characters in length."
        );
        return false;
      }
      break;
    case "decimal":
      if (!isValidDecimal(form[key])) {
        displayCalculationCard(
          "Please enter a valid decimal number, and is not over 999,999,999,999."
        );
        return false;
      }

      form[key] = parseInt(form[key]); // data must be an integer string not a float string
      break;
    case "hexadecimal":
      form[key] = form[key].toUpperCase(); // hexadecimal to uppercase before it reaches the server
      if (!isValidHexadecimal(form[key])) {
        displayCalculationCard(
          "Please enter a valid hexadecimal number, and is not over 64 characters in length."
        );
        return false;
      }
      break;
    default:
      if (!isValidDecimal(form[key])) {
        displayCalculationCard(
          "Please enter a valid input, and is not over 999,999,999,999."
        );
        return false;
      }

      // percentages and total-surface-area category require a float input
      const currentPageURL = window.location.pathname;
      if (
        currentPageURL.indexOf("percentages") !== -1 ||
        currentPageURL.indexOf("total-surface-area") !== -1
      ) {
        form[key] = parseFloat(form[key]);
      } else {
        form[key] = parseInt(form[key]); // input is for a int route such as numbers category
      }
  }
  return true;
};
