<script lang="ts">
    import CardInputElementLayout
        from "@/pages/order-confirm-pay/components/PaymentForm/components/CardInput/CardInputLayout.svelte";
    import CardInputElementLabel
        from "@/pages/order-confirm-pay/components/PaymentForm/components/CardInputElementLabel/CardInputElementLabel.svelte";
    import CardInputElement
        from "@/pages/order-confirm-pay/components/PaymentForm/components/CardInputElement/CardInputElement.svelte";
    import CardInputError
        from "@/pages/order-confirm-pay/components/PaymentForm/components/CardInputError/CardInputError.svelte";

    type CardInputElementProps = {
        label: string,
        id: string,
        placeholder: string,
        error?: string,
        styles?: string,
        focusAction?: () => void,
        blurAction?: () => void,
    }

    let props: CardInputElementProps = $props();
</script>

<CardInputElementLayout styles={props.styles}>
    <CardInputElementLabel {...props}/>
    <CardInputElement {...props}/>
    <CardInputError {...props}/>
</CardInputElementLayout>
