export const displayCalculationCard = (content: string) => {
  const calculationCard = document.getElementById(
    "calculation-card"
  ) as HTMLElement;

  calculationCard.innerHTML = "<p>" + content + "</p>";
  calculationCard.classList.add("fade-in");
};
