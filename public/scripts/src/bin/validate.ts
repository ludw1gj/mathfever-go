export const isValidBinary = (binary: string): boolean => {
  if (binary.length > 64) {
    return false;
  }

  for (let i = 0; i < binary.length; i++) {
    if (binary[i] !== "0" && binary[i] !== "1") {
      return false;
    }
  }
  return true;
};

export const isValidDecimal = (decimal: number): boolean => {
  return !isNaN(decimal) && decimal < 999999999999;
};

export const isValidHexadecimal = (hexadecimal: string): boolean => {
  return /^[A-F0-9]+$/.test(hexadecimal) && hexadecimal.length < 64;
};
