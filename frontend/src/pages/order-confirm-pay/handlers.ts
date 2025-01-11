import {getCreditCardMaxLength, getCreditCardType} from "@/pages/order-confirm-pay/helpers.ts";
import {CreditCardType} from "@/types/CreditCardType.ts";
import type {PaymentData} from "@/types/PaymentData.ts";
import {JSEncrypt} from "jsencrypt";

export const handleCardNumberInput = (
    newValue: string,
    setFields: (key: string, value: string) => void,
) => {
    if (!newValue) return;

    const maxLength = getCreditCardMaxLength(newValue) + Math.floor(getCreditCardMaxLength(newValue) / 4);

    const formattedValue = newValue
        .replace(/\D/g, "")
        .replace(/(.{4})/g, "$1 ")
        .trim();

    const slicedValue = formattedValue.slice(0, maxLength);

    if (newValue !== slicedValue) {
        setFields("cardNumber", slicedValue);
    }
};


export const handleCardExpirationInput = (
    newValue: string,
    setFields: (key: string, value: string) => void,
) => {
    if (!newValue) return;

    const cleanedValue = newValue.replace(/[^0-9/]/g, "");

    let formattedValue = cleanedValue;

    if (cleanedValue.length === 1) {
        if (parseInt(cleanedValue, 10) > 1) formattedValue = `0${cleanedValue}`;
    } else if (cleanedValue.length === 2) {
        const month = parseInt(cleanedValue, 10);
        if (month > 12) formattedValue = "12";
    } else if (cleanedValue.length > 2) {
        const [month, year] = cleanedValue.split("/");
        const validMonth = month.slice(0, 2);
        const validYear = year?.slice(0, 2);
        const yearPart = month.slice(2, 3);

        formattedValue = `${validMonth}/${yearPart}${validYear || ""}`;
    }

    const [inputMonth, inputYear] = formattedValue.split("/").map((part) => parseInt(part, 10));
    const inputDecade = inputYear < 10 ? inputYear : Math.floor(inputYear / 10);
    const currentDate = new Date();
    const currentDecade = Math.floor(currentDate.getFullYear() / 10) % 10;

    if (inputDecade < currentDecade) {
        setFields("expiryDate", formattedValue.slice(0, 3));
        return;
    }

    if (inputMonth && inputYear) {
        const currentMonth = currentDate.getMonth() + 1;
        const currentYear = currentDate.getFullYear() % 100;

        if ((inputYear < currentYear || (inputYear === currentYear && inputMonth < currentMonth)) && inputYear >= 10) {
            setFields("expiryDate", formattedValue.slice(0, 4));
            return;
        }
    }

    if (newValue !== formattedValue) {
        setFields("expiryDate", formattedValue);
    }
};

export const handleCardCvvInput = (
    newCvvValue: string,
    cardNumber: string,
    setFields: (key: string, value: string) => void,
) => {
    if (!newCvvValue) return;

    const maxLength = getCreditCardType(cardNumber) === CreditCardType.AmericanExpress ? 4 : 3;

    const formattedValue = newCvvValue.replace(/\D/g, "").slice(0, maxLength);

    if (newCvvValue.trim() !== formattedValue) {
        setFields("securityCode", formattedValue);
    }
};

export const handleCardHolderNameInput = (
    newValue: string,
    setFields: (key: string, value: string) => void,
) => {
    if (!newValue) return;

    const maxLength = 26;

    const formattedValue = newValue
        .replace(/^[\s]+/, "")
        .replace(/[^a-zA-Z\s']/g, "")
        .replace(/\s+/g, " ")
        .slice(0, maxLength);

    if (newValue !== formattedValue) {
        setFields("cardHolderName", formattedValue);
    }
};

export const handlePaymentFormSubmit = (data: PaymentData) => {
    const publicKey = "-----BEGIN PUBLIC KEY-----\n" +
        "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA7C+L6+LMftNkTzIzg0iJ\n" +
        "w8Fvhlp0bjb29wLVu1Hyz4J4sjNsPqIMh6uw9hxHtpLwF9hPhxXHHLJQ/s1BpxQ3\n" +
        "oz0OU7nSgAOhmC3e23QbC5xhHnn9VTh/q9mpuLszXHl4XGEJhS3OnR5z7m6cQdzV\n" +
        "Bd+VoZSkX4PBLtJzYOlfZZuM7p8OkjwTnV/RVVGwleplKw+wgiPItpqu/KvHk6df\n" +
        "vUoA9HEz6mhGoJivEZrI3NRb3zqsUwmDZvH9fL7dxkRL+MOvGBY2yRAwxiFc07fF\n" +
        "6M3LsFTk9APanKhpMlZHo4hANct6lgT9FhZEF5KHlPvS3jrV6vjsftbEXgXKHZ2B\n" +
        "GQIDAQAB\n" +
        "-----END PUBLIC KEY-----\n"

    const encryptor = new JSEncrypt();
    encryptor.setPublicKey(publicKey);

    // Encrypt the card details
    const encrypted = encryptor.encrypt(JSON.stringify(data));
    if (!encrypted) {
        alert("Encryption failed!");
        return;
    }

    console.log(encrypted);
}



