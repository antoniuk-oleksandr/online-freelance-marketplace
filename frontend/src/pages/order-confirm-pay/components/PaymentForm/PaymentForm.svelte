<script lang="ts">
    import CreditCard from "@/common-components/CreditCard/CreditCard.svelte";
    import CardInputElement
        from "@/pages/order-confirm-pay/components/PaymentForm/components/CardInput/CardInput.svelte";
    import {getContext, onDestroy} from "svelte";
    import {
        handleCardCvvInput,
        handleCardExpirationInput, handleCardHolderNameInput,
        handleCardNumberInput
    } from "@/pages/order-confirm-pay/handlers.ts";
    import type {PaymentData} from "@/types/PaymentData.ts";
    import PaymentFormLayout from "@/pages/order-confirm-pay/components/PaymentForm/PaymentFormLayout.svelte";

    let flip = false;
    let errorData: any = undefined;
    const handleCVV = () => flip = !flip;

    let inputs = [
        {id: "cardNumber", styles: "col-span-2", label: "Card number", placeholder: "1234 5678 9123 4567"},
        {id: "cardHolderName", styles: "col-span-2", label: "Card holder name", placeholder: "John Doe"},
        {id: "expiryDate", label: "Expiry date", placeholder: "01/25"},
        {id: "securityCode", label: "CVV", placeholder: "123", blurAction: handleCVV, focusAction: handleCVV},
    ];

    const setFields: (key: string, value: string) => void = getContext("setFields");
    const errorsStore: any = getContext("errors");
    const unsubscribeErrors = errorsStore.subscribe((value: any) => {
        errorData = value;
    });

    let formData: PaymentData = {
        cardNumber: "",
        cardHolderName: "",
        expiryDate: "",
        securityCode: "",
    };

    const feltDataStore: any = getContext("feltData");
    const unsubscribeData = feltDataStore.subscribe((value: PaymentData) => {
        handleCardNumberInput(value.cardNumber, setFields);
        handleCardExpirationInput(value.expiryDate, setFields);
        handleCardCvvInput(value.securityCode, value.cardNumber, setFields);
        handleCardHolderNameInput(value.cardHolderName, setFields);
        formData = value;
    })

    onDestroy(() => {
        unsubscribeData()
        unsubscribeErrors()
    });
</script>

<PaymentFormLayout>
    <CreditCard
            cardNumber={formData.cardNumber}
            cardHolderName={formData.cardHolderName}
            expiryDate={formData.expiryDate}
            securityCode={formData.securityCode}
            flip={flip}
    />
    <div class="grid grid-cols-2 gap-6 w-full">
        {#each inputs as input}
            <CardInputElement
                    error={errorData[input.id] ? errorData[input.id][0] : ""}
                    {...input}/>
        {/each}
    </div>
</PaymentFormLayout>
