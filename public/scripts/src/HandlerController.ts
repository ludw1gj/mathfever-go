import {serializeForm} from "./Handler/SerializeHandler";
import {validateForm} from "./Handler/ValidationHandler";

function run(): void {
    if (!document.getElementById("calculation-form")) {
        return;
    }

    const calculationForm = document.getElementById("calculation-form") as HTMLFormElement;

    calculationForm.addEventListener("submit", function (event) {
        event.preventDefault();

        const form = serializeForm(calculationForm);
        validateForm(form);
    });
}

run();
