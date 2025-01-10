<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    type CreditCardLayoutProps = LayoutProps & {
        flip: boolean,
    }

    const {children, flip}: CreditCardLayoutProps = $props();
</script>

<style>
    .card {
        width: 100%;
        height: 100%;
        transform-style: preserve-3d;
        transition: transform 0.6s;
    }

    .card.flipped {
        transform: rotateY(180deg);
    }
</style>

<div style="perspective: 1000px" class="card-container h-[13.5rem] max-w-sm w-full">
    <div
            class="card relative shadow-lg rounded-2xl"
            class:flipped={flip}
    >
        {@render children()}
    </div>
</div>