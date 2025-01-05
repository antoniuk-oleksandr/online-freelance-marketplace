<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    type CarouselSlideListProps = LayoutProps & {
        currentIndex: number,
    };

    const {children, currentIndex}: CarouselSlideListProps = $props();
</script>

<div
        class="flex transition-transform duration-500 ease-in-out"
        style="transform: translateX(-{currentIndex * 100}%)"
>
    {@render children()}
</div>