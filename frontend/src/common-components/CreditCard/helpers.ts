export const formatCreditCardExpiryDate = (
    expiryDate: string
) => {
    if (expiryDate === "") return "##/##";
    else return expiryDate;
}

export const formatCreditCardHolderName = (
    cardHolderName: string
) => {
    if (cardHolderName === "") return "#### ####";
    else return cardHolderName;
}

export const formatCreditCardNumber = (
    cardNumber: string
) => {
    if (cardNumber === "") return "#### #### #### ####";
    else return cardNumber;
}

export const formatCreditCardCVV= (
    securityCode: string
) => {
    if (securityCode === "") return "###";
    else return securityCode;
}