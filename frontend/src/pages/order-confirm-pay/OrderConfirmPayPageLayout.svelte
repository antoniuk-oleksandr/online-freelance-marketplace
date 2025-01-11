<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";
    import {createForm} from "felte";
    import {setContext} from "svelte";
    import {paymentFormSchema} from "@/pages/order-confirm-pay/helpers.ts";
    import {validator} from "@felte/validator-zod";
    import {handlePaymentFormSubmit} from "@/pages/order-confirm-pay/handlers.ts";

    const {children}: LayoutProps = $props();

    const {form, setFields, data, errors} = createForm({
        onSubmit: handlePaymentFormSubmit,
        extend: validator({schema: paymentFormSchema}),
    });

    setContext("feltData", data);
    setContext("errors", errors);
    setContext('setFields', setFields);
</script>

<form
        use:form
        class="flex flex-col gap-6 text-base h-full">
    {@render children()}
</form>