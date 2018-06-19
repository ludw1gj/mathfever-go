import { createSerializedForm, createValidatedForm } from "./bin/form";
import { submitAjax } from "./bin/ajax";

(() => {
  if (!document.getElementById("calculation-form")) {
    return;
  }

  const calculationForm = document.getElementById(
    "calculation-form"
  ) as HTMLFormElement;

  calculationForm.addEventListener("submit", event => {
    event.preventDefault();

    const serializedForm = createSerializedForm(calculationForm);
    const validatedForm = createValidatedForm(serializedForm);

    if (validatedForm !== null) {
      const spinner = document.getElementById(
        "loading-spinner-container"
      ) as HTMLElement;
      spinner.style.opacity = "100";

      submitAjax(validatedForm);
    }
  });
})();
