import {serializeForm} from "./util/serialize";
import {validateForm} from "./util/validate";
import {submitAjax} from "./util/ajax";

const run = (): void => {
    if (!document.getElementById("calculation-form")) {
        return;
    }

    const calculationForm = document.getElementById("calculation-form") as HTMLFormElement;
    calculationForm.addEventListener("submit", function (event) {
        event.preventDefault();

        const form = serializeForm(calculationForm);
        const [validatedForm, isValid] = validateForm(form);
        if (isValid) {
            submitAjax(validatedForm);
        }
    });
};

run();
