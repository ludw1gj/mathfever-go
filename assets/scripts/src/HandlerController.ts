import {SerializeHandler} from "./Handler/SerializeHandler";
import {ValidationHandler} from "./Handler/ValidationHandler";

class HandlerController {

    public static run(): void {
        if (document.getElementById("calculation-form")) {
            const calculationForm = document.getElementById("calculation-form") as HTMLFormElement;

            calculationForm.addEventListener("submit", function (event) {
                event.preventDefault();

                const form = SerializeHandler.serializeForm(calculationForm);
                ValidationHandler.validateForm(form);
            });
        }
    }
}

HandlerController.run();
