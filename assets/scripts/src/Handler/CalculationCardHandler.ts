class CalculationCardHandler {

    public static display(content: string): void {
        let calculationCard = document.getElementById("calculation-card") as HTMLElement;
        calculationCard.innerHTML = "<p>" + content + "</p>";
        calculationCard.classList.add("fade-in");
    }

}

export {CalculationCardHandler}
