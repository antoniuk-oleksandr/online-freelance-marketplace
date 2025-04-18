import { CreditCardType } from "@/types/CreditCardType";
import { z } from "zod";

export const getCreditCardType = (cardNumber: string): CreditCardType => {
    if (!cardNumber) return CreditCardType.Unknown;

    const sanitizedNumber = cardNumber.replace(/\s+/g, "");

    if (/^4/.test(sanitizedNumber)) {
        return CreditCardType.Visa;
    }

    if (/^5[1-5]/.test(sanitizedNumber)) {
        return CreditCardType.MasterCard;
    }

    if (/^3[47]/.test(sanitizedNumber)) {
        return CreditCardType.AmericanExpress;
    }

    if (/^6(?:011|5[0-9]{1}|22[2-9][0-9]{1}|[2-5][0-9]{2})/.test(sanitizedNumber)) {
        return CreditCardType.Discover;
    }

    return CreditCardType.Unknown;
};

export const getCreditCardMaxLength = (cardNumber: string): number => {
    const cardType = getCreditCardType(cardNumber);

    switch (cardType) {
        case CreditCardType.Visa:
        case CreditCardType.MasterCard:
        case CreditCardType.Discover:
            return 19;
        case CreditCardType.AmericanExpress:
            return 15;
        default:
            return 19;
    }
};


const testCardNumber = (cardNumber?: string) => {
    if (!cardNumber) return false;

    cardNumber = cardNumber.replace(/\s+/g, "");

    let arr = [];

    for (let i = 0; i < cardNumber.length; i++) {
        if (i % 2 == 1) {
            arr.push(parseInt(cardNumber[i]));
            continue;
        }

        const num = parseInt(cardNumber[i]) * 2;
        if (num > 9) arr.push(num - 9);
        else arr.push(num);
    }

    return arr.reduce((a, b) => a + b) % 10 == 0;
}

const testCardNumberLength = (cardNumber?: string) => {
    if (!cardNumber) return false;

    cardNumber = cardNumber.replace(/\s+/g, "");

    const maxLength = getCreditCardMaxLength(cardNumber)
    return cardNumber.length <= maxLength && cardNumber.length >= 16;
}

const cardExpirationRegex = /^(0[1-9]|1[0-2])\/([0-9]{2})$/;
const cvvRegex = /^[0-9]{3,4}$/;
const cardNameRegex = /^[a-zA-Z\s'-]+$/;

export const paymentFormSchema = z.object({
    cardNumber: z
        .string()
        .min(1, { message: "Card number is required." })
        .refine((value) => testCardNumberLength(value), { message: "Invalid card number length." })
        .refine((value) => testCardNumber(value), { message: "Invalid card number." }),

    cardHolderName: z
        .string()
        .min(1, { message: "Cardholder name is required." })
        .regex(cardNameRegex, { message: "Invalid cardholder name. Only letters, spaces, apostrophes, and hyphens are allowed." }),

    expiryDate: z
        .string()
        .min(1, { message: "Expiration date is required." })
        .regex(cardExpirationRegex, { message: "Invalid expiration date format. Use MM/YY." }),

    securityCode: z
        .string()
        .min(1, { message: "CVV is required." })
        .regex(cvvRegex, { message: "Invalid CVV format. Enter 3 or 4 digits." })
        .length(3, { message: "CVV must be 3 digits." })
        .refine((value) => value.length === 3 || value.length === 4, {
            message: "CVV must be 3 or 4 digits.",
        }),
});
