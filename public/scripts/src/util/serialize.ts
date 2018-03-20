export const serializeForm = (form: HTMLFormElement) => {
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
