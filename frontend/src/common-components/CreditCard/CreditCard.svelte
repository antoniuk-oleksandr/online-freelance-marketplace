<script lang="ts">
    import {getFileServerHost} from "@/utils/utils.ts";
    import {isDarkMode} from "@/common-stores/theme-storage.ts";
    import CreditCardLayout from "@/common-components/CreditCard/CreditCardLayout.svelte";
    import CreditCardFrontSide
        from "@/common-components/CreditCard/components/CreditCardFrontSide/CreditCardFrontSide.svelte";
    import CreditCardBackSide
        from "@/common-components/CreditCard/components/CreditCardBackSide/CreditCardBackSide.svelte";

    type CreditCardProps = {
        flip: boolean,
        securityCode: string,
        cardNumber: string,
        cardHolderName: string,
        expiryDate: string,
    }

    const props: CreditCardProps = $props();

    const host = getFileServerHost();
    const mastercardLogo = `${host}/files/credit-card-logo.png`;
    const chip = `${host}/files/credit-card-chip.png`;

    let cardBg = $state("");

    isDarkMode.subscribe((value) => {
        cardBg = value
            ? `${host}/files/credit-card-bg.png`
            : `${host}/files/credit-card-light-bg.png`;
    });
</script>

<CreditCardLayout flip={props.flip}>
    <CreditCardFrontSide
            {...props}
            cardBg={cardBg}
            chip={chip}
            mastercardLogo={mastercardLogo}
    />
    <CreditCardBackSide
            {...props}
            cardBg={cardBg}
    />
</CreditCardLayout>