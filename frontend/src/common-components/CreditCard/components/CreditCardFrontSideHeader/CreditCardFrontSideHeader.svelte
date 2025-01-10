<script lang="ts">
    import CreditCardFrontSideHeaderLayout
        from "@/common-components/CreditCard/components/CreditCardFrontSideHeader/CreditCardFrontSideHeaderLayout.svelte";
    import CreditCardLogo from "@/common-components/CreditCard/components/CreditCardLogo/CreditCardLogo.svelte";
    import {getCreditCardType} from "@/pages/order-confirm-pay/helpers.ts";

    type CreditCardFrontSideHeaderProps = {
        mastercardLogo: string,
        chip: string,
        cardNumber: string,
    }

    const {mastercardLogo, chip, cardNumber}: CreditCardFrontSideHeaderProps = $props();
</script>

<CreditCardFrontSideHeaderLayout>
    <CreditCardLogo
            cardType={getCreditCardType(cardNumber)}
            mastercardLogo={mastercardLogo}/>
    <img src={chip} alt="Chip" class="w-16"/>
</CreditCardFrontSideHeaderLayout>