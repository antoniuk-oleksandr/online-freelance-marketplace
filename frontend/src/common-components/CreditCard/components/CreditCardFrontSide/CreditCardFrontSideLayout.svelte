<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    type CreditCardFrontSideLayoutProps = LayoutProps & {
        cardBg: string,
    }

    const {children, cardBg} : CreditCardFrontSideLayoutProps = $props();
</script>

<style>

</style>

<div
        class="card-side backface-hidden absolute rounded-2xl size-full text-neutral-800 dark:text-white bg-cover bg-no-repeat p-6"
        style="background-image: url({cardBg}); backface-visibility: hidden;"
>
    {@render  children()}
</div>