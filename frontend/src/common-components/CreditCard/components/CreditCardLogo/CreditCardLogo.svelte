<script lang="ts">
    import DiscoverLogo from "@/common-components/CreditCard/components/DiscoverLogo/DiscoverLogo.svelte";
    import {CreditCardType} from "@/types/CreditCardType.ts";
    import Icon from "@iconify/svelte";

    type CreditCardLogoProps = {
        mastercardLogo: string,
        cardType: CreditCardType,
    }

    const {mastercardLogo, cardType}: CreditCardLogoProps = $props();
</script>

<div class="flex items-center h-8 ">
    {#if cardType === CreditCardType.MasterCard}
        <img src={mastercardLogo} alt="Card Logo" class="w-12 mr-2"/>
    {:else if cardType === CreditCardType.Visa}
        <Icon icon="ri:visa-line" width="56" height="56"/>
    {:else if cardType === CreditCardType.AmericanExpress}
        <Icon icon="fontisto:american-express" width="72" height="48"/>
    {:else if cardType === CreditCardType.Discover}
        <DiscoverLogo/>
    {:else}
        <Icon icon="famicons:card-outline" width="56" height="56"/>
    {/if}
</div>