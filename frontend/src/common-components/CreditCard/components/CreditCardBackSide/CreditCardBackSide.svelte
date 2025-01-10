<script lang="ts">
    import CreditCardBackSideLayout
        from "@/common-components/CreditCard/components/CreditCardBackSide/CreditCardBackSideLayout.svelte";
    import {formatCreditCardCVV} from "@/common-components/CreditCard/helpers.ts";

    type CreditCardBackSideProps = {
        cardBg: string,
        securityCode: string,
    }

    const {cardBg, securityCode}: CreditCardBackSideProps = $props();
</script>

<CreditCardBackSideLayout cardBg={cardBg}>
    <div class="bg-black w-full h-10 mb-4"></div>
    <div class="px-6">
        <h6 class="text-sm font-semibold mb-2">Authorized Signature</h6>
        <div class="flex items-center">
            <div class="flex-grow h-10 bg-light-palette-background-block dark:bg-dark-palette-background-block ring-1 ring-light-palette-divider dark:ring-dark-palette-divider flex items-center px-3">
                <span class="text-gray-500">CVV</span>
            </div>
            <div class="ml-4 h-10 w-16 bg-gray-300 text-black text-center flex items-center justify-center font-semibold">
                {formatCreditCardCVV(securityCode)}
            </div>
        </div>
    </div>
</CreditCardBackSideLayout>