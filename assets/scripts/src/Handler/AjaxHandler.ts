import {CalculationCardHandler} from "./CalculationCardHandler";

class AjaxHandler {

    public static submit(data: Object) {
        let spinner = document.getElementById("loading-spinner-container") as HTMLElement;

        const url = "/api" + window.location.pathname;
        let request = new XMLHttpRequest();
        request.open("POST", url, true);

        request.onload = function () {
            spinner.style.opacity = "0";

            if (this.status >= 200 && this.status <= 404) {
                const resp = JSON.parse(this.response);

                if (resp.error) {
                    CalculationCardHandler.display(resp.error);
                } else {
                    CalculationCardHandler.display(resp.content);
                }
            }
            else if (this.status === 429) {
                CalculationCardHandler.display("You sent too many requests to the server, come back tomorrow.")
            } else {
                CalculationCardHandler.display("The server has encountered a problem.");
            }
        };
        request.onerror = function () {
            spinner.style.opacity = "0";
            CalculationCardHandler.display("There was a connection issue. Check your internet connection or the sever might be down.")
        };
        request.setRequestHeader("Content-Type", "application/json");
        request.send(JSON.stringify(data));
    }

}

export {AjaxHandler}
