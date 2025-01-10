<script lang="ts">
    import CreditCardFrontSideLayout
        from "@/common-components/CreditCard/components/CreditCardFrontSide/CreditCardFrontSideLayout.svelte";
    import CreditCardFrontSideHeader
        from "@/common-components/CreditCard/components/CreditCardFrontSideHeader/CreditCardFrontSideHeader.svelte";
    import CreditCardFrontSideBody
        from "@/common-components/CreditCard/components/CreditCardFrontSideBody/CreditCardFrontSideBody.svelte";

    type CreditCardFrontSideProps = {
        cardNumber: string,
        cardHolderName: string,
        expiryDate: string,
        cardBg: string,
        mastercardLogo: string,
        chip: string,
    }

    const props: CreditCardFrontSideProps = $props();
</script>

<CreditCardFrontSideLayout cardBg={props.cardBg}>
    <CreditCardFrontSideHeader {...props}/>
    <CreditCardFrontSideBody {...props}/>
</CreditCardFrontSideLayout>
