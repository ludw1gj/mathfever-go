import {displayCalculationCard} from "./CalculationCardHandler";

export function submitAjax(data: Object) {
    const spinner = document.getElementById("loading-spinner-container") as HTMLElement;

    const url = "/api" + window.location.pathname;
    const request = new XMLHttpRequest();
    request.open("POST", url, true);

    request.onload = function () {
        spinner.style.opacity = "0";

        if (this.status >= 200 && this.status <= 404) {
            const resp = JSON.parse(this.response);

            if (resp.error) {
                displayCalculationCard(resp.error);
            } else {
                displayCalculationCard(resp.content);
            }
        }
        else if (this.status === 429) {
            displayCalculationCard("You sent too many requests to the server, come back tomorrow.")
        } else {
            displayCalculationCard("The server has encountered a problem.");
        }
    };
    request.onerror = function () {
        spinner.style.opacity = "0";
        displayCalculationCard("There was a connection issue. Check your internet connection or the sever might be down.")
    };
    request.setRequestHeader("Content-Type", "application/json");
    request.send(JSON.stringify(data));
}
