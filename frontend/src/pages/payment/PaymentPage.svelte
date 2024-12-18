<script lang="ts">
    import {getPaymentPageParams} from "@/pages/payment/helpers.ts";
    import type {OrderParams} from "@/types/OrderParams.ts";
    import NotFound from "@/common-components/NotFound/NotFound.svelte";

    let orderParams = $state<undefined | null | OrderParams>(undefined);

    $effect(() => {
        if (orderParams !== undefined) return;

        orderParams = getPaymentPageParams();
        console.log(orderParams);
    })


</script>

{#if orderParams === null}
    <NotFound/>
{:else if orderParams !== undefined}
    <div>
        <button>Pay</button>
    </div>
{/if}
