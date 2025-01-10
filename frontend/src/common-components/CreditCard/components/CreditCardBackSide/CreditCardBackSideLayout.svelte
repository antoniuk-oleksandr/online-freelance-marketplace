<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    type CreditCardBackSideProps = LayoutProps & {
        cardBg: string,
    }

    const {cardBg, children}: CreditCardBackSideProps = $props();
</script>

<div
        class="card-side backface-hidden absolute rounded-2xl size-full card-back text-neutral-800 dark:text-white bg-cover bg-no-repeat py-6"
        style="background-image: url({cardBg}); backface-visibility: hidden; transform: rotateY(180deg);"
>
    {@render children()}
</div>