<script lang="ts">
    import {
        formatCreditCardExpiryDate,
        formatCreditCardHolderName,
        formatCreditCardNumber
    } from "@/common-components/CreditCard/helpers.ts";
    import CreditCardFrontSideBodyLayout
        from "@/common-components/CreditCard/components/CreditCardFrontSideBody/CreditCardFrontSideBodyLayout.svelte";

    type CreditCardFrontSideBodyProps = {
        cardNumber: string,
        cardHolderName: string,
        expiryDate: string,
    }

    const {cardNumber, cardHolderName, expiryDate}: CreditCardFrontSideBodyProps = $props();
</script>

<CreditCardFrontSideBodyLayout>
    <div>
        <h6 class="text-xs font-normal">Card Number</h6>
        <h5 class="text-lg font-normal mt-1 tracking-wide">{formatCreditCardNumber(cardNumber)}</h5>
        <h5 class="mt-5">{formatCreditCardHolderName(cardHolderName)}</h5>
    </div>
    <div>
        <h6 class="text-xs font-normal">Valid Thru</h6>
        <h5>{formatCreditCardExpiryDate(expiryDate)}</h5>
    </div>
</CreditCardFrontSideBodyLayout>