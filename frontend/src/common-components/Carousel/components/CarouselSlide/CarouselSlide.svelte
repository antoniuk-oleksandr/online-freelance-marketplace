<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";
    import CarouselSlideLayout from "@/common-components/Carousel/components/CarouselSlide/CarouselSlideLayout.svelte";


    const {children}: LayoutProps = $props();
</script>

<CarouselSlideLayout>
    {@render children()}
</CarouselSlideLayout>